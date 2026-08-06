// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package entity

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	xpv2 "github.com/crossplane/crossplane/apis/v2/core/v2"
	"github.com/crossplane/crossplane-runtime/v2/pkg/event"
	"github.com/crossplane/crossplane-runtime/v2/pkg/meta"
	"github.com/crossplane/crossplane-runtime/v2/pkg/ratelimiter"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	tjcontroller "github.com/crossplane/upjet/v2/pkg/controller"
	"github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/vikreinok/provider-dynatrace-all/apis/cluster/host/v1alpha1"
	clusterv1beta1 "github.com/vikreinok/provider-dynatrace-all/apis/cluster/v1beta1"
)

const (
	errNotHostEntity = "managed resource is not a HostEntity custom resource"
	errNewClient     = "cannot create Dynatrace API client"
)

// SetupGated adds a controller that reconciles HostEntity managed resources.
func SetupGated(mgr ctrl.Manager, o tjcontroller.Options) error {
	o.Gate.Register(func() {
		if err := Setup(mgr, o); err != nil {
			mgr.GetLogger().Error(err, "unable to setup reconciler", "gvk", v1alpha1.HostEntity_GroupVersionKind.String())
		}
	}, v1alpha1.HostEntity_GroupVersionKind)
	return nil
}

// Setup adds a controller that reconciles HostEntity managed resources.
func Setup(mgr ctrl.Manager, o tjcontroller.Options) error {
	name := managed.ControllerName(v1alpha1.HostEntity_GroupKind)

	r := managed.NewReconciler(mgr,
		resource.ManagedKind(v1alpha1.HostEntity_GroupVersionKind),
		managed.WithExternalConnecter(&connector{
			kube:         mgr.GetClient(),
			newServiceFn: newService,
		}),
		managed.WithLogger(o.Logger.WithValues("controller", name)),
		managed.WithRecorder(event.NewAPIRecorder(mgr.GetEventRecorderFor(name))),
		managed.WithPollInterval(o.PollInterval),
	)

	return ctrl.NewControllerManagedBy(mgr).
		Named(name).
		WithOptions(o.ForControllerRuntime()).
		For(&v1alpha1.HostEntity{}).
		Complete(ratelimiter.NewReconciler(name, r, o.GlobalRateLimiter))
}

type connector struct {
	kube         client.Client
	newServiceFn func(ctx context.Context, envURL, apiToken string) (service, error)
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.HostEntity)
	if !ok {
		return nil, errors.New(errNotHostEntity)
	}

	pc := &clusterv1beta1.ProviderConfig{}
	if err := c.kube.Get(ctx, client.ObjectKey{Name: cr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, "cannot get provider config")
	}

	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c.kube, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, "cannot extract credentials")
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal credentials")
	}

	svc, err := c.newServiceFn(ctx, creds["dt_env_url"], creds["dt_api_token"])
	if err != nil {
		return nil, errors.Wrap(err, errNewClient)
	}

	return &external{service: svc}, nil
}

type external struct {
	service service
}

func (e *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.HostEntity)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotHostEntity)
	}

	if meta.WasDeleted(cr) {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}

	entityType := "HOST"
	if cr.Spec.ForProvider.Type != nil {
		entityType = *cr.Spec.ForProvider.Type
	}
	entityName := ""
	if cr.Spec.ForProvider.Name != nil {
		entityName = *cr.Spec.ForProvider.Name
	}

	if entityName == "" {
		cr.SetConditions(xpv2.ReconcileError(errors.New("spec.forProvider.name is required")))
		return managed.ExternalObservation{
			ResourceExists:   true,
			ResourceUpToDate: true,
		}, nil
	}

	id, tags, count, err := e.service.Lookup(ctx, entityType, entityName)
	if err != nil {
		cr.SetConditions(xpv2.ReconcileError(err))
		return managed.ExternalObservation{}, err
	}

	if id != "" {
		cr.Status.AtProvider.EntityID = &id
		cr.Status.AtProvider.Tags = tags
		meta.SetExternalName(cr, id)
		if count > 1 {
			cr.SetConditions(xpv2.Available().WithMessage(fmt.Sprintf("Warning: Multiple entities (%d) matched query. Using first match: %s", count, id)))
		} else {
			cr.SetConditions(xpv2.Available())
		}
	} else {
		cr.Status.AtProvider.EntityID = nil
		cr.Status.AtProvider.Tags = nil
		cr.SetConditions(xpv2.Unavailable().WithMessage(fmt.Sprintf("Entity %s not found in Dynatrace yet", entityName)))
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil
}

func (e *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	return managed.ExternalCreation{}, nil
}

func (e *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	return managed.ExternalUpdate{}, nil
}

func (e *external) Delete(ctx context.Context, mg resource.Managed) (managed.ExternalDelete, error) {
	return managed.ExternalDelete{}, nil
}

func (e *external) Disconnect(ctx context.Context) error {
	return nil
}

type service interface {
	Lookup(ctx context.Context, entityType, entityName string) (string, []v1alpha1.HostEntityTag, int, error)
}

type apiService struct {
	client   *http.Client
	envURL   string
	apiToken string
}

func newService(ctx context.Context, envURL, apiToken string) (service, error) {
	return &apiService{
		client:   http.DefaultClient,
		envURL:   envURL,
		apiToken: apiToken,
	}, nil
}

type entityItem struct {
	EntityID    string                   `json:"entityId"`
	Type        string                   `json:"type"`
	DisplayName string                   `json:"displayName"`
	Tags        []v1alpha1.HostEntityTag `json:"tags"`
}

func (s *apiService) Lookup(ctx context.Context, entityType, entityName string) (string, []v1alpha1.HostEntityTag, int, error) {
	var selector string
	if entityName == "*" {
		selector = fmt.Sprintf(`type("%s")`, entityType)
	} else if len(entityName) > 1 && entityName[len(entityName)-1] == '*' {
		prefix := entityName[:len(entityName)-1]
		selector = fmt.Sprintf(`type("%s"),entityName.startsWith("%s")`, entityType, prefix)
	} else {
		selector = fmt.Sprintf(`type("%s"),entityName.equals("%s")`, entityType, entityName)
	}
	queryURL := fmt.Sprintf("%s/api/v2/entities?entitySelector=%s&fields=tags", s.envURL, url.QueryEscape(selector))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, queryURL, nil)
	if err != nil {
		return "", nil, 0, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Api-Token %s", s.apiToken))

	resp, err := s.client.Do(req)
	if err != nil {
		return "", nil, 0, err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return "", nil, 0, fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, string(b))
	}

	var response struct {
		Entities []entityItem `json:"entities"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", nil, 0, err
	}

	if len(response.Entities) > 0 {
		return response.Entities[0].EntityID, response.Entities[0].Tags, len(response.Entities), nil
	}

	return "", nil, 0, nil
}

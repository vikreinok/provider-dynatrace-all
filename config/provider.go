package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

const (
	resourcePrefix = "dynatrace"
	modulePath     = "github.com/vikreinok/provider-dynatrace-all"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

var resourceGroups = map[string]string{}

func init() {
	// Dynamically scan documentation directories to map resource names to subcategories
	docDirs := []string{
		".work/dynatrace-oss/dynatrace/docs/resources",
		".work/dynatrace-oss/dynatrace/website/docs/r",
		".work/dynatrace-oss/dynatrace/docs/r",
		".work/dynatrace-oss/dynatrace/website/docs/resources",
	}

	var foundDir string
	for _, dir := range docDirs {
		if _, err := os.Stat(dir); err == nil {
			foundDir = dir
			break
		}
	}

	if foundDir != "" {
		files, err := ioutil.ReadDir(foundDir)
		if err == nil {
			subcatRegex := regexp.MustCompile(`(?i)subcategory:\s*["']?([^"'\n\r]+)["']?`)
			for _, file := range files {
				if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
					continue
				}
				content, err := ioutil.ReadFile(filepath.Join(foundDir, file.Name()))
				if err != nil {
					continue
				}
				match := subcatRegex.FindStringSubmatch(string(content))
				if len(match) > 1 {
					subcat := strings.TrimSpace(match[1])
					// Normalize subcat to lowercase alphanumeric
					normalized := strings.ToLower(subcat)
					normalized = regexp.MustCompile(`[^a-z0-9]`).ReplaceAllString(normalized, "")

					// Resource name in TF has "dynatrace_" prefix
					resourceName := "dynatrace_" + strings.TrimSuffix(file.Name(), ".md")
					resourceGroups[resourceName] = normalized
				}
			}
		}
	}
}

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("dynatrace.crossplane.io"),
		ujconfig.WithIncludeList([]string{"dynatrace_.*"}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			DefaultResourceConfigurations(),
		))

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("dynatrace.m.crossplane.io"),
		ujconfig.WithIncludeList([]string{"dynatrace_.*"}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			DefaultResourceConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	pc.ConfigureResources()
	return pc
}

// DefaultResourceConfigurations sets the dynamic ShortGroup and version settings
func DefaultResourceConfigurations() ujconfig.ResourceOption {
	return func(r *ujconfig.Resource) {
		// Set dynamic ShortGroup based on extracted subcategories
		if group, ok := resourceGroups[r.Name]; ok {
			r.ShortGroup = group
		} else {
			// Fallback: use the second token of the resource name (e.g. dynatrace_iam_group -> iam)
			parts := strings.Split(r.Name, "_")
			if len(parts) > 1 {
				r.ShortGroup = parts[1]
			} else {
				r.ShortGroup = "dynatrace"
			}
		}

		// Apply custom kind/reference overrides to ensure compatibility with old manifests and references
		switch r.Name {
		case "dynatrace_openpipeline_v2_logs_pipelinegroups":
			r.References["member_pipelines"] = ujconfig.Reference{
				Type: "V2LogsPipelines",
			}
		case "dynatrace_openpipeline_v2_spans_pipelinegroups":
			r.References["member_pipelines"] = ujconfig.Reference{
				Type: "V2SpansPipelines",
			}
		case "dynatrace_openpipeline_v2_logs_routing":
			r.References["routing_entries.routing_entry.pipeline_id"] = ujconfig.Reference{
				Type: "V2LogsPipelines",
			}
		case "dynatrace_openpipeline_v2_spans_routing":
			r.References["routing_entries.routing_entry.pipeline_id"] = ujconfig.Reference{
				Type: "V2SpansPipelines",
			}
		case "dynatrace_iam_policy_bindings_v2":
			r.References["group"] = ujconfig.Reference{
				Type: "Group",
			}
			r.References["policy.id"] = ujconfig.Reference{
				Type: "Policy",
			}
			r.References["policy.boundaries"] = ujconfig.Reference{
				Type: "PolicyBoundary",
			}
		case "dynatrace_alerting":
			r.References["management_zone"] = ujconfig.Reference{
				Type: "github.com/vikreinok/provider-dynatrace-all/apis/cluster/management/v1alpha1.ZoneV2",
			}
		}
	}
}

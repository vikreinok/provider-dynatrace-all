// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	token "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/activegate/token"
	updates "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/activegate/updates"
	tokenag "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ag/token"
	extension "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aix/extension"
	alerting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/alerting/alerting"
	profile "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/alerting/profile"
	towernotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ansible/towernotification"
	detection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/api/detection"
	tokenapi "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/api/token"
	monitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/app/monitoring"
	anomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/application/anomalies"
	dataprivacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/application/dataprivacy"
	detectionrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/application/detectionrule"
	detectionrulev2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/application/detectionrulev2"
	errorrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/application/errorrules"
	notification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/appsec/notification"
	alertingattack "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attack/alerting"
	allowlist "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attack/allowlist"
	rules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attack/rules"
	settings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attack/settings"
	allowlistattribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attribute/allowlist"
	blocklist "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attribute/blocklist"
	masking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attribute/masking"
	preferences "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/attributes/preferences"
	log "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/audit/log"
	approval "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/approval"
	businesscalendar "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/businesscalendar"
	controllerconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/controllerconnections"
	schedulingrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/schedulingrule"
	workflow "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/workflow"
	workflowawsconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/workflowawsconnections"
	workflowjira "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/workflowjira"
	workflowk8sconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/workflowk8sconnections"
	workflowslack "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/automation/workflowslack"
	autotag "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/autotag/autotag"
	rulesautotag "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/autotag/rules"
	v2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/autotag/v2"
	anomaliesaws "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aws/anomalies"
	connection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aws/connection"
	connectionrolearn "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aws/connectionrolearn"
	credentials "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aws/credentials"
	service "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/aws/service"
	connectionazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/azure/connection"
	connectionauthentication "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/azure/connectionauthentication"
	credentialsazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/azure/credentials"
	serviceazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/azure/service"
	monitor "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/browser/monitor"
	monitoroutage "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/browser/monitoroutage"
	monitorperformance "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/browser/monitorperformance"
	processmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/builtin/processmonitoring"
	eventsbuckets "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventsbuckets"
	eventscapturingvariants "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventscapturingvariants"
	eventsmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventsmetrics"
	eventsoneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventsoneagent"
	eventsoneagentoutgoing "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventsoneagentoutgoing"
	eventsprocessing "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventsprocessing"
	eventssecuritycontext "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/business/eventssecuritycontext"
	mobilemetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/calculated/mobilemetric"
	servicemetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/calculated/servicemetric"
	syntheticmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/calculated/syntheticmetric"
	webmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/calculated/webmetric"
	developmentenvironments "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/cloud/developmentenvironments"
	foundry "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/cloud/foundry"
	workloaddetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/cloudapp/workloaddetection"
	credentialscloudfoundry "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/cloudfoundry/credentials"
	alerts "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/connectivity/alerts"
	builtinrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/container/builtinrule"
	registry "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/container/registry"
	rule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/container/rule"
	technology "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/container/technology"
	analytics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/crashdump/analytics"
	credentialscredentials "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/credentials/credentials"
	anomaliescustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/anomalies"
	appanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/appanomalies"
	appcrashrate "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/appcrashrate"
	appenablement "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/appenablement"
	device "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/device"
	servicecustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/service"
	serviceorder "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/serviceorder"
	tags "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/tags"
	units "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/custom/units"
	dashboard "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/dashboard/dashboard"
	sharing "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/dashboard/sharing"
	allowlistdashboards "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/dashboards/allowlist"
	general "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/dashboards/general"
	presets "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/dashboards/presets"
	privacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/data/privacy"
	anomaliesdatabase "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/database/anomalies"
	anomaliesv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/database/anomaliesv2"
	anomalydetectors "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/davis/anomalydetectors"
	copilot "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/davis/copilot"
	appfeatureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/db/appfeatureflags"
	pool "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ddu/pool"
	grouping "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/declarative/grouping"
	launchpad "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/default/launchpad"
	agentoptin "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/devobs/agentoptin"
	datamasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/devobs/datamasking"
	gitonprem "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/devobs/gitonprem"
	shares "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/direct/shares"
	defaultrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/discovery/defaultrules"
	featureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/discovery/featureflags"
	analyticsdisk "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/analytics"
	anomaliesdisk "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/anomalies"
	anomaliesv2disk "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/anomaliesv2"
	anomalyrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/anomalyrules"
	edgeanomalydetectors "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/edgeanomalydetectors"
	options "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/options"
	specificanomaliesv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/disk/specificanomaliesv2"
	document "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/document/document"
	servicediscovery "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ebpf/servicediscovery"
	notificationemail "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/email/notification"
	detectionrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/endpoint/detectionrules"
	detectionrulesoptin "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/endpoint/detectionrulesoptin"
	environment "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/environment/environment"
	settingseula "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/eula/settings"
	drivenansibleconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/event/drivenansibleconnections"
	executioncontroller "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/extension/executioncontroller"
	executionremote "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/extension/executionremote"
	detectionparameters "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/failure/detectionparameters"
	detectionrulesfailure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/failure/detectionrules"
	detectionrulesets "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/failure/detectionrulesets"
	issues "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/frequent/issues"
	relationships "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/generic/relationships"
	setting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/generic/setting"
	types "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/generic/types"
	geolocation "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/geolocation/geolocation"
	connectiongithub "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/github/connection"
	connectiongitlab "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/gitlab/connection"
	state "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/golden/state"
	metricsallowall "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/grail/metricsallowall"
	metricsallowlist "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/grail/metricsallowlist"
	securitycontext "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/grail/securitycontext"
	metrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/histogram/metrics"
	anomalieshost "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/anomalies"
	anomaliesv2host "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/anomaliesv2"
	monitoringhost "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/monitoring"
	monitoringadvanced "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/monitoringadvanced"
	monitoringmode "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/monitoringmode"
	naming "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/naming"
	namingorder "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/namingorder"
	processgroupmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/host/processgroupmonitoring"
	monitorhttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/http/monitor"
	monitorcookies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/http/monitorcookies"
	monitoroutagehttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/http/monitoroutage"
	monitorperformancehttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/http/monitorperformance"
	monitorscript "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/http/monitorscript"
	extensionactiveversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/hub/extensionactiveversion"
	extensionconfig "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/hub/extensionconfig"
	extensionv2config "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/hub/extensionv2config"
	permissions "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/hub/permissions"
	subscriptions "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/hub/subscriptions"
	group "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/group"
	permission "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/permission"
	policy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/policy"
	policybindings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/policybindings"
	policybindingsv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/policybindingsv2"
	policyboundary "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/policyboundary"
	serviceuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/serviceuser"
	user "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/iam/user"
	mqfilters "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ibm/mqfilters"
	bridges "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ims/bridges"
	appfeatureflagsinfraops "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/infraops/appfeatureflags"
	appsettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/infraops/appsettings"
	addressmasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ip/addressmasking"
	tracking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/issue/tracking"
	connectionjenkins "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/jenkins/connection"
	notificationjira "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/jira/notification"
	dashboardjson "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/json/dashboard"
	dashboardbase "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/json/dashboardbase"
	clusteranomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/clusteranomalies"
	credentialsk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/credentials"
	monitoringk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/monitoring"
	namespaceanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/namespaceanomalies"
	nodeanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/nodeanomalies"
	pvcanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/pvcanomalies"
	workloadanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/k8s/workloadanomalies"
	requests "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/key/requests"
	useraction "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/key/useraction"
	app "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/kubernetes/app"
	enrichment "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/kubernetes/enrichment"
	kubernetes "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/kubernetes/kubernetes"
	spm "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/kubernetes/spm"
	outboundconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/limit/outboundconnections"
	agentfeatureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/agentfeatureflags"
	buckets "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/buckets"
	customattribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/customattribute"
	customsource "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/customsource"
	debugsettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/debugsettings"
	events "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/events"
	grail "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/grail"
	metricslog "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/metrics"
	oneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/oneagent"
	processing "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/processing"
	securitycontextlog "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/securitycontext"
	sensitivedatamasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/sensitivedatamasking"
	storage "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/storage"
	timestamp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/log/timestamp"
	transactionmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mainframe/transactionmonitoring"
	maintenance "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/maintenance/maintenance"
	window "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/maintenance/window"
	backup "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/backup"
	internetproxy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/internetproxy"
	networkzones "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/networkzones"
	preferencesmanaged "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/preferences"
	publicendpoints "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/publicendpoints"
	remoteaccess "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/remoteaccess"
	smtp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/managed/smtp"
	zone "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/management/zone"
	zonev2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/management/zonev2"
	eventsmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/metric/events"
	metadata "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/metric/metadata"
	query "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/metric/query"
	permissionmgmz "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mgmz/permission"
	appanomaliesmobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/appanomalies"
	appcrashratemobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/appcrashrate"
	appenablementmobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/appenablement"
	appkeyperformance "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/appkeyperformance"
	application "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/application"
	apprequesterrors "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/apprequesterrors"
	notifications "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/mobile/notifications"
	technologiesapache "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesapache"
	technologiesdotnet "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesdotnet"
	technologiesgo "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesgo"
	technologiesiis "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesiis"
	technologiesjava "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesjava"
	technologiesnginx "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesnginx"
	technologiesnodejs "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesnodejs"
	technologiesopentracing "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesopentracing"
	technologiesphp "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesphp"
	technologiespython "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiespython"
	technologiesvarnish "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologiesvarnish"
	technologieswsmb "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/monitored/technologieswsmb"
	emailconnection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ms365/emailconnection"
	connectionmsentraid "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/msentraid/connection"
	connectionmsteams "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/msteams/connection"
	requestsmuted "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/muted/requests"
	nettracer "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/nettracer/nettracer"
	monitornetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/network/monitor"
	monitoroutagenetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/network/monitoroutage"
	traffic "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/network/traffic"
	zonenetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/network/zone"
	zones "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/network/zones"
	notificationnotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/notification/notification"
	defaultmode "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/oneagent/defaultmode"
	defaultversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/oneagent/defaultversion"
	features "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/oneagent/features"
	sidemasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/oneagent/sidemasking"
	updatesoneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/oneagent/updates"
	businessevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/businessevents"
	davisevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/davisevents"
	davisproblems "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/davisproblems"
	eventsopenpipeline "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/events"
	logs "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/logs"
	metricsopenpipeline "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/metrics"
	sdlcevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/sdlcevents"
	securityevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/securityevents"
	spans "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/spans"
	systemevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/systemevents"
	userevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/userevents"
	usersessions "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/usersessions"
	v2bizeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2bizeventsingestsources"
	v2bizeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2bizeventspipelinegroups"
	v2bizeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2bizeventspipelines"
	v2bizeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2bizeventsrouting"
	v2daviseventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2daviseventsingestsources"
	v2daviseventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2daviseventspipelinegroups"
	v2daviseventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2daviseventspipelines"
	v2daviseventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2daviseventsrouting"
	v2davisproblemsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2davisproblemsingestsources"
	v2davisproblemspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2davisproblemspipelinegroups"
	v2davisproblemspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2davisproblemspipelines"
	v2davisproblemsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2davisproblemsrouting"
	v2eventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventsingestsources"
	v2eventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventspipelinegroups"
	v2eventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventspipelines"
	v2eventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventsrouting"
	v2eventssdlcingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssdlcingestsources"
	v2eventssdlcpipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssdlcpipelinegroups"
	v2eventssdlcpipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssdlcpipelines"
	v2eventssdlcrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssdlcrouting"
	v2eventssecurityingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssecurityingestsources"
	v2eventssecuritypipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssecuritypipelinegroups"
	v2eventssecuritypipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssecuritypipelines"
	v2eventssecurityrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2eventssecurityrouting"
	v2logsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2logsingestsources"
	v2logspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2logspipelinegroups"
	v2logspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2logspipelines"
	v2logsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2logsrouting"
	v2metricsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2metricsingestsources"
	v2metricspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2metricspipelinegroups"
	v2metricspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2metricspipelines"
	v2metricsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2metricsrouting"
	v2securityeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2securityeventsingestsources"
	v2securityeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2securityeventspipelinegroups"
	v2securityeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2securityeventspipelines"
	v2securityeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2securityeventsrouting"
	v2spansingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2spansingestsources"
	v2spanspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2spanspipelinegroups"
	v2spanspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2spanspipelines"
	v2spansrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2spansrouting"
	v2systemeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2systemeventsingestsources"
	v2systemeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2systemeventspipelinegroups"
	v2systemeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2systemeventspipelines"
	v2systemeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2systemeventsrouting"
	v2usereventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usereventsingestsources"
	v2usereventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usereventspipelinegroups"
	v2usereventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usereventspipelines"
	v2usereventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usereventsrouting"
	v2usersessionsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usersessionsingestsources"
	v2usersessionspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usersessionspipelinegroups"
	v2usersessionspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usersessionspipelines"
	v2usersessionsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/openpipeline/v2usersessionsrouting"
	metricsopentelemetry "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/opentelemetry/metrics"
	genienotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ops/genienotification"
	services "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/os/services"
	config "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ownership/config"
	teams "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/ownership/teams"
	dutynotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/pager/dutynotification"
	connectionpagerduty "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/pagerduty/connection"
	alertingpg "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/pg/alerting"
	anomaliespg "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/pg/anomalies"
	bucket "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/platform/bucket"
	slo "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/platform/slo"
	bindings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/policy/bindings"
	policypolicy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/policy/policy"
	fields "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/problem/fields"
	recordpropagationrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/problem/recordpropagationrules"
	availability "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/availability"
	groupdetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/groupdetection"
	groupdetectionflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/groupdetectionflags"
	groupmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/groupmonitoring"
	grouprum "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/grouprum"
	groupsimpledetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/groupsimpledetection"
	monitoringprocess "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/monitoring"
	monitoringrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/monitoringrule"
	visibility "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/process/visibility"
	namingprocessgroup "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/processgroup/naming"
	namingorderprocessgroup "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/processgroup/namingorder"
	providerconfig "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/providerconfig"
	manager "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/queue/manager"
	sharinggroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/queue/sharinggroups"
	environments "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/remote/environments"
	report "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/report/report"
	attribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/request/attribute"
	namingrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/request/naming"
	namings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/request/namings"
	attributes "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/resource/attributes"
	basedsampling "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rpc/basedsampling"
	advancedcorrelation "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/advancedcorrelation"
	hostheaders "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/hostheaders"
	ipdetermination "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/ipdetermination"
	iplocations "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/iplocations"
	overloadprevention "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/overloadprevention"
	providerbreakdown "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/rum/providerbreakdown"
	context "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/security/context"
	segment "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/segment/segment"
	anomaliesservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/anomalies"
	anomaliesv2service "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/anomaliesv2"
	detectionrulesservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/detectionrules"
	externalwebrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/externalwebrequest"
	externalwebservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/externalwebservice"
	failure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/failure"
	fullwebrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/fullwebrequest"
	fullwebservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/fullwebservice"
	httpfailure "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/httpfailure"
	namingservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/naming"
	namingorderservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/namingorder"
	nownotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/nownotification"
	splitting "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/service/splitting"
	connectionservicenow "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/servicenow/connection"
	replayresourcecapture "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/session/replayresourcecapture"
	replaywebprivacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/session/replaywebprivacy"
	permissionssettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/settings/permissions"
	reliabilityguardian "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/site/reliabilityguardian"
	notificationslack "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/slack/notification"
	normalization "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/slo/normalization"
	sloslo "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/slo/slo"
	v2slo "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/slo/v2"
	attributespan "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/span/attribute"
	capturerule "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/span/capturerule"
	contextpropagation "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/span/contextpropagation"
	entrypoint "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/span/entrypoint"
	eventsspan "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/span/events"
	availabilitysynthetic "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/synthetic/availability"
	location "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/synthetic/location"
	settingstoken "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/token/settings"
	startfilters "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/transaction/startfilters"
	notificationtrello "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/trello/notification"
	servicesmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/unified/servicesmetrics"
	servicesopentel "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/unified/servicesopentel"
	windows "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/update/windows"
	basedsamplingurl "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/url/basedsampling"
	analyticsusability "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/usability/analytics"
	actionmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/actionmetrics"
	experiencescore "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/experiencescore"
	groupuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/group"
	sessionmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/sessionmetrics"
	settingsuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/settings"
	useruser "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/user/user"
	opsnotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/victor/opsnotification"
	anomaliesvmware "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vmware/anomalies"
	vmware "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vmware/vmware"
	alertingvulnerability "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/alerting"
	code "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/code"
	settingsvulnerability "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/settings"
	thirdparty "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/thirdparty"
	thirdpartyattr "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/thirdpartyattr"
	thirdpartyk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/vulnerability/thirdpartyk8s"
	appanomaliesweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appanomalies"
	appautoinjection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appautoinjection"
	appbeaconendpoint "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appbeaconendpoint"
	appbeaconorigins "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appbeaconorigins"
	appcustomconfigproperties "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appcustomconfigproperties"
	appcustomerrors "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appcustomerrors"
	appcustominjection "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appcustominjection"
	appcustomproprestrictions "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appcustomproprestrictions"
	appenablementweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appenablement"
	appinjectioncookie "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appinjectioncookie"
	appipaddressexclusion "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appipaddressexclusion"
	appjavascriptfilename "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appjavascriptfilename"
	appjavascriptupdates "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appjavascriptupdates"
	appjavascriptversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appjavascriptversion"
	appkeyperformancecustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appkeyperformancecustom"
	appkeyperformanceload "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appkeyperformanceload"
	appkeyperformancexhr "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appkeyperformancexhr"
	applicationweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/application"
	appmanualinsertion "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appmanualinsertion"
	apprequesterrorsweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/apprequesterrors"
	appresourcecleanup "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appresourcecleanup"
	appresourcetypes "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/web/appresourcetypes"
	notificationwebhook "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/webhook/notification"
	notificationxmatters "github.com/vikreinok/provider-dynatrace-all/internal/controller/namespaced/xmatters/notification"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		token.Setup,
		updates.Setup,
		tokenag.Setup,
		extension.Setup,
		alerting.Setup,
		profile.Setup,
		towernotification.Setup,
		detection.Setup,
		tokenapi.Setup,
		monitoring.Setup,
		anomalies.Setup,
		dataprivacy.Setup,
		detectionrule.Setup,
		detectionrulev2.Setup,
		errorrules.Setup,
		notification.Setup,
		alertingattack.Setup,
		allowlist.Setup,
		rules.Setup,
		settings.Setup,
		allowlistattribute.Setup,
		blocklist.Setup,
		masking.Setup,
		preferences.Setup,
		log.Setup,
		approval.Setup,
		businesscalendar.Setup,
		controllerconnections.Setup,
		schedulingrule.Setup,
		workflow.Setup,
		workflowawsconnections.Setup,
		workflowjira.Setup,
		workflowk8sconnections.Setup,
		workflowslack.Setup,
		autotag.Setup,
		rulesautotag.Setup,
		v2.Setup,
		anomaliesaws.Setup,
		connection.Setup,
		connectionrolearn.Setup,
		credentials.Setup,
		service.Setup,
		connectionazure.Setup,
		connectionauthentication.Setup,
		credentialsazure.Setup,
		serviceazure.Setup,
		monitor.Setup,
		monitoroutage.Setup,
		monitorperformance.Setup,
		processmonitoring.Setup,
		eventsbuckets.Setup,
		eventscapturingvariants.Setup,
		eventsmetrics.Setup,
		eventsoneagent.Setup,
		eventsoneagentoutgoing.Setup,
		eventsprocessing.Setup,
		eventssecuritycontext.Setup,
		mobilemetric.Setup,
		servicemetric.Setup,
		syntheticmetric.Setup,
		webmetric.Setup,
		developmentenvironments.Setup,
		foundry.Setup,
		workloaddetection.Setup,
		credentialscloudfoundry.Setup,
		alerts.Setup,
		builtinrule.Setup,
		registry.Setup,
		rule.Setup,
		technology.Setup,
		analytics.Setup,
		credentialscredentials.Setup,
		anomaliescustom.Setup,
		appanomalies.Setup,
		appcrashrate.Setup,
		appenablement.Setup,
		device.Setup,
		servicecustom.Setup,
		serviceorder.Setup,
		tags.Setup,
		units.Setup,
		dashboard.Setup,
		sharing.Setup,
		allowlistdashboards.Setup,
		general.Setup,
		presets.Setup,
		privacy.Setup,
		anomaliesdatabase.Setup,
		anomaliesv2.Setup,
		anomalydetectors.Setup,
		copilot.Setup,
		appfeatureflags.Setup,
		pool.Setup,
		grouping.Setup,
		launchpad.Setup,
		agentoptin.Setup,
		datamasking.Setup,
		gitonprem.Setup,
		shares.Setup,
		defaultrules.Setup,
		featureflags.Setup,
		analyticsdisk.Setup,
		anomaliesdisk.Setup,
		anomaliesv2disk.Setup,
		anomalyrules.Setup,
		edgeanomalydetectors.Setup,
		options.Setup,
		specificanomaliesv2.Setup,
		document.Setup,
		servicediscovery.Setup,
		notificationemail.Setup,
		detectionrules.Setup,
		detectionrulesoptin.Setup,
		environment.Setup,
		settingseula.Setup,
		drivenansibleconnections.Setup,
		executioncontroller.Setup,
		executionremote.Setup,
		detectionparameters.Setup,
		detectionrulesfailure.Setup,
		detectionrulesets.Setup,
		issues.Setup,
		relationships.Setup,
		setting.Setup,
		types.Setup,
		geolocation.Setup,
		connectiongithub.Setup,
		connectiongitlab.Setup,
		state.Setup,
		metricsallowall.Setup,
		metricsallowlist.Setup,
		securitycontext.Setup,
		metrics.Setup,
		anomalieshost.Setup,
		anomaliesv2host.Setup,
		monitoringhost.Setup,
		monitoringadvanced.Setup,
		monitoringmode.Setup,
		naming.Setup,
		namingorder.Setup,
		processgroupmonitoring.Setup,
		monitorhttp.Setup,
		monitorcookies.Setup,
		monitoroutagehttp.Setup,
		monitorperformancehttp.Setup,
		monitorscript.Setup,
		extensionactiveversion.Setup,
		extensionconfig.Setup,
		extensionv2config.Setup,
		permissions.Setup,
		subscriptions.Setup,
		group.Setup,
		permission.Setup,
		policy.Setup,
		policybindings.Setup,
		policybindingsv2.Setup,
		policyboundary.Setup,
		serviceuser.Setup,
		user.Setup,
		mqfilters.Setup,
		bridges.Setup,
		appfeatureflagsinfraops.Setup,
		appsettings.Setup,
		addressmasking.Setup,
		tracking.Setup,
		connectionjenkins.Setup,
		notificationjira.Setup,
		dashboardjson.Setup,
		dashboardbase.Setup,
		clusteranomalies.Setup,
		credentialsk8s.Setup,
		monitoringk8s.Setup,
		namespaceanomalies.Setup,
		nodeanomalies.Setup,
		pvcanomalies.Setup,
		workloadanomalies.Setup,
		requests.Setup,
		useraction.Setup,
		app.Setup,
		enrichment.Setup,
		kubernetes.Setup,
		spm.Setup,
		outboundconnections.Setup,
		agentfeatureflags.Setup,
		buckets.Setup,
		customattribute.Setup,
		customsource.Setup,
		debugsettings.Setup,
		events.Setup,
		grail.Setup,
		metricslog.Setup,
		oneagent.Setup,
		processing.Setup,
		securitycontextlog.Setup,
		sensitivedatamasking.Setup,
		storage.Setup,
		timestamp.Setup,
		transactionmonitoring.Setup,
		maintenance.Setup,
		window.Setup,
		backup.Setup,
		internetproxy.Setup,
		networkzones.Setup,
		preferencesmanaged.Setup,
		publicendpoints.Setup,
		remoteaccess.Setup,
		smtp.Setup,
		zone.Setup,
		zonev2.Setup,
		eventsmetric.Setup,
		metadata.Setup,
		query.Setup,
		permissionmgmz.Setup,
		appanomaliesmobile.Setup,
		appcrashratemobile.Setup,
		appenablementmobile.Setup,
		appkeyperformance.Setup,
		application.Setup,
		apprequesterrors.Setup,
		notifications.Setup,
		technologiesapache.Setup,
		technologiesdotnet.Setup,
		technologiesgo.Setup,
		technologiesiis.Setup,
		technologiesjava.Setup,
		technologiesnginx.Setup,
		technologiesnodejs.Setup,
		technologiesopentracing.Setup,
		technologiesphp.Setup,
		technologiespython.Setup,
		technologiesvarnish.Setup,
		technologieswsmb.Setup,
		emailconnection.Setup,
		connectionmsentraid.Setup,
		connectionmsteams.Setup,
		requestsmuted.Setup,
		nettracer.Setup,
		monitornetwork.Setup,
		monitoroutagenetwork.Setup,
		traffic.Setup,
		zonenetwork.Setup,
		zones.Setup,
		notificationnotification.Setup,
		defaultmode.Setup,
		defaultversion.Setup,
		features.Setup,
		sidemasking.Setup,
		updatesoneagent.Setup,
		businessevents.Setup,
		davisevents.Setup,
		davisproblems.Setup,
		eventsopenpipeline.Setup,
		logs.Setup,
		metricsopenpipeline.Setup,
		sdlcevents.Setup,
		securityevents.Setup,
		spans.Setup,
		systemevents.Setup,
		userevents.Setup,
		usersessions.Setup,
		v2bizeventsingestsources.Setup,
		v2bizeventspipelinegroups.Setup,
		v2bizeventspipelines.Setup,
		v2bizeventsrouting.Setup,
		v2daviseventsingestsources.Setup,
		v2daviseventspipelinegroups.Setup,
		v2daviseventspipelines.Setup,
		v2daviseventsrouting.Setup,
		v2davisproblemsingestsources.Setup,
		v2davisproblemspipelinegroups.Setup,
		v2davisproblemspipelines.Setup,
		v2davisproblemsrouting.Setup,
		v2eventsingestsources.Setup,
		v2eventspipelinegroups.Setup,
		v2eventspipelines.Setup,
		v2eventsrouting.Setup,
		v2eventssdlcingestsources.Setup,
		v2eventssdlcpipelinegroups.Setup,
		v2eventssdlcpipelines.Setup,
		v2eventssdlcrouting.Setup,
		v2eventssecurityingestsources.Setup,
		v2eventssecuritypipelinegroups.Setup,
		v2eventssecuritypipelines.Setup,
		v2eventssecurityrouting.Setup,
		v2logsingestsources.Setup,
		v2logspipelinegroups.Setup,
		v2logspipelines.Setup,
		v2logsrouting.Setup,
		v2metricsingestsources.Setup,
		v2metricspipelinegroups.Setup,
		v2metricspipelines.Setup,
		v2metricsrouting.Setup,
		v2securityeventsingestsources.Setup,
		v2securityeventspipelinegroups.Setup,
		v2securityeventspipelines.Setup,
		v2securityeventsrouting.Setup,
		v2spansingestsources.Setup,
		v2spanspipelinegroups.Setup,
		v2spanspipelines.Setup,
		v2spansrouting.Setup,
		v2systemeventsingestsources.Setup,
		v2systemeventspipelinegroups.Setup,
		v2systemeventspipelines.Setup,
		v2systemeventsrouting.Setup,
		v2usereventsingestsources.Setup,
		v2usereventspipelinegroups.Setup,
		v2usereventspipelines.Setup,
		v2usereventsrouting.Setup,
		v2usersessionsingestsources.Setup,
		v2usersessionspipelinegroups.Setup,
		v2usersessionspipelines.Setup,
		v2usersessionsrouting.Setup,
		metricsopentelemetry.Setup,
		genienotification.Setup,
		services.Setup,
		config.Setup,
		teams.Setup,
		dutynotification.Setup,
		connectionpagerduty.Setup,
		alertingpg.Setup,
		anomaliespg.Setup,
		bucket.Setup,
		slo.Setup,
		bindings.Setup,
		policypolicy.Setup,
		fields.Setup,
		recordpropagationrules.Setup,
		availability.Setup,
		groupdetection.Setup,
		groupdetectionflags.Setup,
		groupmonitoring.Setup,
		grouprum.Setup,
		groupsimpledetection.Setup,
		monitoringprocess.Setup,
		monitoringrule.Setup,
		visibility.Setup,
		namingprocessgroup.Setup,
		namingorderprocessgroup.Setup,
		providerconfig.Setup,
		manager.Setup,
		sharinggroups.Setup,
		environments.Setup,
		report.Setup,
		attribute.Setup,
		namingrequest.Setup,
		namings.Setup,
		attributes.Setup,
		basedsampling.Setup,
		advancedcorrelation.Setup,
		hostheaders.Setup,
		ipdetermination.Setup,
		iplocations.Setup,
		overloadprevention.Setup,
		providerbreakdown.Setup,
		context.Setup,
		segment.Setup,
		anomaliesservice.Setup,
		anomaliesv2service.Setup,
		detectionrulesservice.Setup,
		externalwebrequest.Setup,
		externalwebservice.Setup,
		failure.Setup,
		fullwebrequest.Setup,
		fullwebservice.Setup,
		httpfailure.Setup,
		namingservice.Setup,
		namingorderservice.Setup,
		nownotification.Setup,
		splitting.Setup,
		connectionservicenow.Setup,
		replayresourcecapture.Setup,
		replaywebprivacy.Setup,
		permissionssettings.Setup,
		reliabilityguardian.Setup,
		notificationslack.Setup,
		normalization.Setup,
		sloslo.Setup,
		v2slo.Setup,
		attributespan.Setup,
		capturerule.Setup,
		contextpropagation.Setup,
		entrypoint.Setup,
		eventsspan.Setup,
		availabilitysynthetic.Setup,
		location.Setup,
		settingstoken.Setup,
		startfilters.Setup,
		notificationtrello.Setup,
		servicesmetrics.Setup,
		servicesopentel.Setup,
		windows.Setup,
		basedsamplingurl.Setup,
		analyticsusability.Setup,
		actionmetrics.Setup,
		experiencescore.Setup,
		groupuser.Setup,
		sessionmetrics.Setup,
		settingsuser.Setup,
		useruser.Setup,
		opsnotification.Setup,
		anomaliesvmware.Setup,
		vmware.Setup,
		alertingvulnerability.Setup,
		code.Setup,
		settingsvulnerability.Setup,
		thirdparty.Setup,
		thirdpartyattr.Setup,
		thirdpartyk8s.Setup,
		appanomaliesweb.Setup,
		appautoinjection.Setup,
		appbeaconendpoint.Setup,
		appbeaconorigins.Setup,
		appcustomconfigproperties.Setup,
		appcustomerrors.Setup,
		appcustominjection.Setup,
		appcustomproprestrictions.Setup,
		appenablementweb.Setup,
		appinjectioncookie.Setup,
		appipaddressexclusion.Setup,
		appjavascriptfilename.Setup,
		appjavascriptupdates.Setup,
		appjavascriptversion.Setup,
		appkeyperformancecustom.Setup,
		appkeyperformanceload.Setup,
		appkeyperformancexhr.Setup,
		applicationweb.Setup,
		appmanualinsertion.Setup,
		apprequesterrorsweb.Setup,
		appresourcecleanup.Setup,
		appresourcetypes.Setup,
		notificationwebhook.Setup,
		notificationxmatters.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		token.SetupGated,
		updates.SetupGated,
		tokenag.SetupGated,
		extension.SetupGated,
		alerting.SetupGated,
		profile.SetupGated,
		towernotification.SetupGated,
		detection.SetupGated,
		tokenapi.SetupGated,
		monitoring.SetupGated,
		anomalies.SetupGated,
		dataprivacy.SetupGated,
		detectionrule.SetupGated,
		detectionrulev2.SetupGated,
		errorrules.SetupGated,
		notification.SetupGated,
		alertingattack.SetupGated,
		allowlist.SetupGated,
		rules.SetupGated,
		settings.SetupGated,
		allowlistattribute.SetupGated,
		blocklist.SetupGated,
		masking.SetupGated,
		preferences.SetupGated,
		log.SetupGated,
		approval.SetupGated,
		businesscalendar.SetupGated,
		controllerconnections.SetupGated,
		schedulingrule.SetupGated,
		workflow.SetupGated,
		workflowawsconnections.SetupGated,
		workflowjira.SetupGated,
		workflowk8sconnections.SetupGated,
		workflowslack.SetupGated,
		autotag.SetupGated,
		rulesautotag.SetupGated,
		v2.SetupGated,
		anomaliesaws.SetupGated,
		connection.SetupGated,
		connectionrolearn.SetupGated,
		credentials.SetupGated,
		service.SetupGated,
		connectionazure.SetupGated,
		connectionauthentication.SetupGated,
		credentialsazure.SetupGated,
		serviceazure.SetupGated,
		monitor.SetupGated,
		monitoroutage.SetupGated,
		monitorperformance.SetupGated,
		processmonitoring.SetupGated,
		eventsbuckets.SetupGated,
		eventscapturingvariants.SetupGated,
		eventsmetrics.SetupGated,
		eventsoneagent.SetupGated,
		eventsoneagentoutgoing.SetupGated,
		eventsprocessing.SetupGated,
		eventssecuritycontext.SetupGated,
		mobilemetric.SetupGated,
		servicemetric.SetupGated,
		syntheticmetric.SetupGated,
		webmetric.SetupGated,
		developmentenvironments.SetupGated,
		foundry.SetupGated,
		workloaddetection.SetupGated,
		credentialscloudfoundry.SetupGated,
		alerts.SetupGated,
		builtinrule.SetupGated,
		registry.SetupGated,
		rule.SetupGated,
		technology.SetupGated,
		analytics.SetupGated,
		credentialscredentials.SetupGated,
		anomaliescustom.SetupGated,
		appanomalies.SetupGated,
		appcrashrate.SetupGated,
		appenablement.SetupGated,
		device.SetupGated,
		servicecustom.SetupGated,
		serviceorder.SetupGated,
		tags.SetupGated,
		units.SetupGated,
		dashboard.SetupGated,
		sharing.SetupGated,
		allowlistdashboards.SetupGated,
		general.SetupGated,
		presets.SetupGated,
		privacy.SetupGated,
		anomaliesdatabase.SetupGated,
		anomaliesv2.SetupGated,
		anomalydetectors.SetupGated,
		copilot.SetupGated,
		appfeatureflags.SetupGated,
		pool.SetupGated,
		grouping.SetupGated,
		launchpad.SetupGated,
		agentoptin.SetupGated,
		datamasking.SetupGated,
		gitonprem.SetupGated,
		shares.SetupGated,
		defaultrules.SetupGated,
		featureflags.SetupGated,
		analyticsdisk.SetupGated,
		anomaliesdisk.SetupGated,
		anomaliesv2disk.SetupGated,
		anomalyrules.SetupGated,
		edgeanomalydetectors.SetupGated,
		options.SetupGated,
		specificanomaliesv2.SetupGated,
		document.SetupGated,
		servicediscovery.SetupGated,
		notificationemail.SetupGated,
		detectionrules.SetupGated,
		detectionrulesoptin.SetupGated,
		environment.SetupGated,
		settingseula.SetupGated,
		drivenansibleconnections.SetupGated,
		executioncontroller.SetupGated,
		executionremote.SetupGated,
		detectionparameters.SetupGated,
		detectionrulesfailure.SetupGated,
		detectionrulesets.SetupGated,
		issues.SetupGated,
		relationships.SetupGated,
		setting.SetupGated,
		types.SetupGated,
		geolocation.SetupGated,
		connectiongithub.SetupGated,
		connectiongitlab.SetupGated,
		state.SetupGated,
		metricsallowall.SetupGated,
		metricsallowlist.SetupGated,
		securitycontext.SetupGated,
		metrics.SetupGated,
		anomalieshost.SetupGated,
		anomaliesv2host.SetupGated,
		monitoringhost.SetupGated,
		monitoringadvanced.SetupGated,
		monitoringmode.SetupGated,
		naming.SetupGated,
		namingorder.SetupGated,
		processgroupmonitoring.SetupGated,
		monitorhttp.SetupGated,
		monitorcookies.SetupGated,
		monitoroutagehttp.SetupGated,
		monitorperformancehttp.SetupGated,
		monitorscript.SetupGated,
		extensionactiveversion.SetupGated,
		extensionconfig.SetupGated,
		extensionv2config.SetupGated,
		permissions.SetupGated,
		subscriptions.SetupGated,
		group.SetupGated,
		permission.SetupGated,
		policy.SetupGated,
		policybindings.SetupGated,
		policybindingsv2.SetupGated,
		policyboundary.SetupGated,
		serviceuser.SetupGated,
		user.SetupGated,
		mqfilters.SetupGated,
		bridges.SetupGated,
		appfeatureflagsinfraops.SetupGated,
		appsettings.SetupGated,
		addressmasking.SetupGated,
		tracking.SetupGated,
		connectionjenkins.SetupGated,
		notificationjira.SetupGated,
		dashboardjson.SetupGated,
		dashboardbase.SetupGated,
		clusteranomalies.SetupGated,
		credentialsk8s.SetupGated,
		monitoringk8s.SetupGated,
		namespaceanomalies.SetupGated,
		nodeanomalies.SetupGated,
		pvcanomalies.SetupGated,
		workloadanomalies.SetupGated,
		requests.SetupGated,
		useraction.SetupGated,
		app.SetupGated,
		enrichment.SetupGated,
		kubernetes.SetupGated,
		spm.SetupGated,
		outboundconnections.SetupGated,
		agentfeatureflags.SetupGated,
		buckets.SetupGated,
		customattribute.SetupGated,
		customsource.SetupGated,
		debugsettings.SetupGated,
		events.SetupGated,
		grail.SetupGated,
		metricslog.SetupGated,
		oneagent.SetupGated,
		processing.SetupGated,
		securitycontextlog.SetupGated,
		sensitivedatamasking.SetupGated,
		storage.SetupGated,
		timestamp.SetupGated,
		transactionmonitoring.SetupGated,
		maintenance.SetupGated,
		window.SetupGated,
		backup.SetupGated,
		internetproxy.SetupGated,
		networkzones.SetupGated,
		preferencesmanaged.SetupGated,
		publicendpoints.SetupGated,
		remoteaccess.SetupGated,
		smtp.SetupGated,
		zone.SetupGated,
		zonev2.SetupGated,
		eventsmetric.SetupGated,
		metadata.SetupGated,
		query.SetupGated,
		permissionmgmz.SetupGated,
		appanomaliesmobile.SetupGated,
		appcrashratemobile.SetupGated,
		appenablementmobile.SetupGated,
		appkeyperformance.SetupGated,
		application.SetupGated,
		apprequesterrors.SetupGated,
		notifications.SetupGated,
		technologiesapache.SetupGated,
		technologiesdotnet.SetupGated,
		technologiesgo.SetupGated,
		technologiesiis.SetupGated,
		technologiesjava.SetupGated,
		technologiesnginx.SetupGated,
		technologiesnodejs.SetupGated,
		technologiesopentracing.SetupGated,
		technologiesphp.SetupGated,
		technologiespython.SetupGated,
		technologiesvarnish.SetupGated,
		technologieswsmb.SetupGated,
		emailconnection.SetupGated,
		connectionmsentraid.SetupGated,
		connectionmsteams.SetupGated,
		requestsmuted.SetupGated,
		nettracer.SetupGated,
		monitornetwork.SetupGated,
		monitoroutagenetwork.SetupGated,
		traffic.SetupGated,
		zonenetwork.SetupGated,
		zones.SetupGated,
		notificationnotification.SetupGated,
		defaultmode.SetupGated,
		defaultversion.SetupGated,
		features.SetupGated,
		sidemasking.SetupGated,
		updatesoneagent.SetupGated,
		businessevents.SetupGated,
		davisevents.SetupGated,
		davisproblems.SetupGated,
		eventsopenpipeline.SetupGated,
		logs.SetupGated,
		metricsopenpipeline.SetupGated,
		sdlcevents.SetupGated,
		securityevents.SetupGated,
		spans.SetupGated,
		systemevents.SetupGated,
		userevents.SetupGated,
		usersessions.SetupGated,
		v2bizeventsingestsources.SetupGated,
		v2bizeventspipelinegroups.SetupGated,
		v2bizeventspipelines.SetupGated,
		v2bizeventsrouting.SetupGated,
		v2daviseventsingestsources.SetupGated,
		v2daviseventspipelinegroups.SetupGated,
		v2daviseventspipelines.SetupGated,
		v2daviseventsrouting.SetupGated,
		v2davisproblemsingestsources.SetupGated,
		v2davisproblemspipelinegroups.SetupGated,
		v2davisproblemspipelines.SetupGated,
		v2davisproblemsrouting.SetupGated,
		v2eventsingestsources.SetupGated,
		v2eventspipelinegroups.SetupGated,
		v2eventspipelines.SetupGated,
		v2eventsrouting.SetupGated,
		v2eventssdlcingestsources.SetupGated,
		v2eventssdlcpipelinegroups.SetupGated,
		v2eventssdlcpipelines.SetupGated,
		v2eventssdlcrouting.SetupGated,
		v2eventssecurityingestsources.SetupGated,
		v2eventssecuritypipelinegroups.SetupGated,
		v2eventssecuritypipelines.SetupGated,
		v2eventssecurityrouting.SetupGated,
		v2logsingestsources.SetupGated,
		v2logspipelinegroups.SetupGated,
		v2logspipelines.SetupGated,
		v2logsrouting.SetupGated,
		v2metricsingestsources.SetupGated,
		v2metricspipelinegroups.SetupGated,
		v2metricspipelines.SetupGated,
		v2metricsrouting.SetupGated,
		v2securityeventsingestsources.SetupGated,
		v2securityeventspipelinegroups.SetupGated,
		v2securityeventspipelines.SetupGated,
		v2securityeventsrouting.SetupGated,
		v2spansingestsources.SetupGated,
		v2spanspipelinegroups.SetupGated,
		v2spanspipelines.SetupGated,
		v2spansrouting.SetupGated,
		v2systemeventsingestsources.SetupGated,
		v2systemeventspipelinegroups.SetupGated,
		v2systemeventspipelines.SetupGated,
		v2systemeventsrouting.SetupGated,
		v2usereventsingestsources.SetupGated,
		v2usereventspipelinegroups.SetupGated,
		v2usereventspipelines.SetupGated,
		v2usereventsrouting.SetupGated,
		v2usersessionsingestsources.SetupGated,
		v2usersessionspipelinegroups.SetupGated,
		v2usersessionspipelines.SetupGated,
		v2usersessionsrouting.SetupGated,
		metricsopentelemetry.SetupGated,
		genienotification.SetupGated,
		services.SetupGated,
		config.SetupGated,
		teams.SetupGated,
		dutynotification.SetupGated,
		connectionpagerduty.SetupGated,
		alertingpg.SetupGated,
		anomaliespg.SetupGated,
		bucket.SetupGated,
		slo.SetupGated,
		bindings.SetupGated,
		policypolicy.SetupGated,
		fields.SetupGated,
		recordpropagationrules.SetupGated,
		availability.SetupGated,
		groupdetection.SetupGated,
		groupdetectionflags.SetupGated,
		groupmonitoring.SetupGated,
		grouprum.SetupGated,
		groupsimpledetection.SetupGated,
		monitoringprocess.SetupGated,
		monitoringrule.SetupGated,
		visibility.SetupGated,
		namingprocessgroup.SetupGated,
		namingorderprocessgroup.SetupGated,
		providerconfig.SetupGated,
		manager.SetupGated,
		sharinggroups.SetupGated,
		environments.SetupGated,
		report.SetupGated,
		attribute.SetupGated,
		namingrequest.SetupGated,
		namings.SetupGated,
		attributes.SetupGated,
		basedsampling.SetupGated,
		advancedcorrelation.SetupGated,
		hostheaders.SetupGated,
		ipdetermination.SetupGated,
		iplocations.SetupGated,
		overloadprevention.SetupGated,
		providerbreakdown.SetupGated,
		context.SetupGated,
		segment.SetupGated,
		anomaliesservice.SetupGated,
		anomaliesv2service.SetupGated,
		detectionrulesservice.SetupGated,
		externalwebrequest.SetupGated,
		externalwebservice.SetupGated,
		failure.SetupGated,
		fullwebrequest.SetupGated,
		fullwebservice.SetupGated,
		httpfailure.SetupGated,
		namingservice.SetupGated,
		namingorderservice.SetupGated,
		nownotification.SetupGated,
		splitting.SetupGated,
		connectionservicenow.SetupGated,
		replayresourcecapture.SetupGated,
		replaywebprivacy.SetupGated,
		permissionssettings.SetupGated,
		reliabilityguardian.SetupGated,
		notificationslack.SetupGated,
		normalization.SetupGated,
		sloslo.SetupGated,
		v2slo.SetupGated,
		attributespan.SetupGated,
		capturerule.SetupGated,
		contextpropagation.SetupGated,
		entrypoint.SetupGated,
		eventsspan.SetupGated,
		availabilitysynthetic.SetupGated,
		location.SetupGated,
		settingstoken.SetupGated,
		startfilters.SetupGated,
		notificationtrello.SetupGated,
		servicesmetrics.SetupGated,
		servicesopentel.SetupGated,
		windows.SetupGated,
		basedsamplingurl.SetupGated,
		analyticsusability.SetupGated,
		actionmetrics.SetupGated,
		experiencescore.SetupGated,
		groupuser.SetupGated,
		sessionmetrics.SetupGated,
		settingsuser.SetupGated,
		useruser.SetupGated,
		opsnotification.SetupGated,
		anomaliesvmware.SetupGated,
		vmware.SetupGated,
		alertingvulnerability.SetupGated,
		code.SetupGated,
		settingsvulnerability.SetupGated,
		thirdparty.SetupGated,
		thirdpartyattr.SetupGated,
		thirdpartyk8s.SetupGated,
		appanomaliesweb.SetupGated,
		appautoinjection.SetupGated,
		appbeaconendpoint.SetupGated,
		appbeaconorigins.SetupGated,
		appcustomconfigproperties.SetupGated,
		appcustomerrors.SetupGated,
		appcustominjection.SetupGated,
		appcustomproprestrictions.SetupGated,
		appenablementweb.SetupGated,
		appinjectioncookie.SetupGated,
		appipaddressexclusion.SetupGated,
		appjavascriptfilename.SetupGated,
		appjavascriptupdates.SetupGated,
		appjavascriptversion.SetupGated,
		appkeyperformancecustom.SetupGated,
		appkeyperformanceload.SetupGated,
		appkeyperformancexhr.SetupGated,
		applicationweb.SetupGated,
		appmanualinsertion.SetupGated,
		apprequesterrorsweb.SetupGated,
		appresourcecleanup.SetupGated,
		appresourcetypes.SetupGated,
		notificationwebhook.SetupGated,
		notificationxmatters.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

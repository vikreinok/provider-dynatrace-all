// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	token "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/activegate/token"
	updates "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/activegate/updates"
	tokenag "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ag/token"
	extension "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aix/extension"
	alerting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/alerting/alerting"
	profile "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/alerting/profile"
	towernotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ansible/towernotification"
	detection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/api/detection"
	tokenapi "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/api/token"
	monitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/app/monitoring"
	anomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/application/anomalies"
	dataprivacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/application/dataprivacy"
	detectionrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/application/detectionrule"
	detectionrulev2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/application/detectionrulev2"
	errorrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/application/errorrules"
	notification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/appsec/notification"
	alertingattack "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attack/alerting"
	allowlist "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attack/allowlist"
	rules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attack/rules"
	settings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attack/settings"
	allowlistattribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attribute/allowlist"
	blocklist "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attribute/blocklist"
	masking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attribute/masking"
	preferences "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/attributes/preferences"
	log "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/audit/log"
	approval "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/approval"
	businesscalendar "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/businesscalendar"
	controllerconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/controllerconnections"
	schedulingrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/schedulingrule"
	workflow "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/workflow"
	workflowawsconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/workflowawsconnections"
	workflowjira "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/workflowjira"
	workflowk8sconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/workflowk8sconnections"
	workflowslack "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/automation/workflowslack"
	autotag "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/autotag/autotag"
	rulesautotag "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/autotag/rules"
	v2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/autotag/v2"
	anomaliesaws "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aws/anomalies"
	connection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aws/connection"
	connectionrolearn "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aws/connectionrolearn"
	credentials "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aws/credentials"
	service "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/aws/service"
	connectionazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/azure/connection"
	connectionauthentication "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/azure/connectionauthentication"
	credentialsazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/azure/credentials"
	serviceazure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/azure/service"
	monitor "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/browser/monitor"
	monitoroutage "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/browser/monitoroutage"
	monitorperformance "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/browser/monitorperformance"
	processmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/builtin/processmonitoring"
	eventsbuckets "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventsbuckets"
	eventscapturingvariants "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventscapturingvariants"
	eventsmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventsmetrics"
	eventsoneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventsoneagent"
	eventsoneagentoutgoing "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventsoneagentoutgoing"
	eventsprocessing "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventsprocessing"
	eventssecuritycontext "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/business/eventssecuritycontext"
	mobilemetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/calculated/mobilemetric"
	servicemetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/calculated/servicemetric"
	syntheticmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/calculated/syntheticmetric"
	webmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/calculated/webmetric"
	developmentenvironments "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/cloud/developmentenvironments"
	foundry "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/cloud/foundry"
	workloaddetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/cloudapp/workloaddetection"
	credentialscloudfoundry "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/cloudfoundry/credentials"
	alerts "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/connectivity/alerts"
	builtinrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/container/builtinrule"
	registry "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/container/registry"
	rule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/container/rule"
	technology "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/container/technology"
	analytics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/crashdump/analytics"
	credentialscredentials "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/credentials/credentials"
	anomaliescustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/anomalies"
	appanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/appanomalies"
	appcrashrate "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/appcrashrate"
	appenablement "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/appenablement"
	device "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/device"
	servicecustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/service"
	serviceorder "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/serviceorder"
	tags "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/tags"
	units "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/custom/units"
	dashboard "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/dashboard/dashboard"
	sharing "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/dashboard/sharing"
	allowlistdashboards "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/dashboards/allowlist"
	general "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/dashboards/general"
	presets "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/dashboards/presets"
	privacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/data/privacy"
	anomaliesdatabase "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/database/anomalies"
	anomaliesv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/database/anomaliesv2"
	anomalydetectors "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/davis/anomalydetectors"
	copilot "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/davis/copilot"
	appfeatureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/db/appfeatureflags"
	pool "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ddu/pool"
	grouping "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/declarative/grouping"
	launchpad "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/default/launchpad"
	agentoptin "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/devobs/agentoptin"
	datamasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/devobs/datamasking"
	gitonprem "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/devobs/gitonprem"
	shares "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/direct/shares"
	defaultrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/discovery/defaultrules"
	featureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/discovery/featureflags"
	analyticsdisk "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/analytics"
	anomaliesdisk "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/anomalies"
	anomaliesv2disk "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/anomaliesv2"
	anomalyrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/anomalyrules"
	edgeanomalydetectors "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/edgeanomalydetectors"
	options "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/options"
	specificanomaliesv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/disk/specificanomaliesv2"
	document "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/document/document"
	servicediscovery "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ebpf/servicediscovery"
	notificationemail "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/email/notification"
	detectionrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/endpoint/detectionrules"
	detectionrulesoptin "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/endpoint/detectionrulesoptin"
	environment "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/environment/environment"
	settingseula "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/eula/settings"
	drivenansibleconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/event/drivenansibleconnections"
	executioncontroller "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/extension/executioncontroller"
	executionremote "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/extension/executionremote"
	detectionparameters "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/failure/detectionparameters"
	detectionrulesfailure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/failure/detectionrules"
	detectionrulesets "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/failure/detectionrulesets"
	issues "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/frequent/issues"
	relationships "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/generic/relationships"
	setting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/generic/setting"
	types "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/generic/types"
	geolocation "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/geolocation/geolocation"
	connectiongithub "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/github/connection"
	connectiongitlab "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/gitlab/connection"
	state "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/golden/state"
	metricsallowall "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/grail/metricsallowall"
	metricsallowlist "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/grail/metricsallowlist"
	securitycontext "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/grail/securitycontext"
	metrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/histogram/metrics"
	anomalieshost "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/anomalies"
	anomaliesv2host "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/anomaliesv2"
	monitoringhost "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/monitoring"
	monitoringadvanced "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/monitoringadvanced"
	monitoringmode "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/monitoringmode"
	naming "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/naming"
	namingorder "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/namingorder"
	processgroupmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/host/processgroupmonitoring"
	monitorhttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/http/monitor"
	monitorcookies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/http/monitorcookies"
	monitoroutagehttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/http/monitoroutage"
	monitorperformancehttp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/http/monitorperformance"
	monitorscript "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/http/monitorscript"
	extensionactiveversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/hub/extensionactiveversion"
	extensionconfig "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/hub/extensionconfig"
	extensionv2config "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/hub/extensionv2config"
	permissions "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/hub/permissions"
	subscriptions "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/hub/subscriptions"
	group "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/group"
	permission "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/permission"
	policy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/policy"
	policybindings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/policybindings"
	policybindingsv2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/policybindingsv2"
	policyboundary "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/policyboundary"
	serviceuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/serviceuser"
	user "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/iam/user"
	mqfilters "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ibm/mqfilters"
	bridges "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ims/bridges"
	appfeatureflagsinfraops "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/infraops/appfeatureflags"
	appsettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/infraops/appsettings"
	addressmasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ip/addressmasking"
	tracking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/issue/tracking"
	connectionjenkins "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/jenkins/connection"
	notificationjira "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/jira/notification"
	dashboardjson "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/json/dashboard"
	dashboardbase "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/json/dashboardbase"
	clusteranomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/clusteranomalies"
	credentialsk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/credentials"
	monitoringk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/monitoring"
	namespaceanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/namespaceanomalies"
	nodeanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/nodeanomalies"
	pvcanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/pvcanomalies"
	workloadanomalies "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/k8s/workloadanomalies"
	requests "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/key/requests"
	useraction "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/key/useraction"
	app "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/kubernetes/app"
	enrichment "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/kubernetes/enrichment"
	kubernetes "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/kubernetes/kubernetes"
	spm "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/kubernetes/spm"
	outboundconnections "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/limit/outboundconnections"
	agentfeatureflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/agentfeatureflags"
	buckets "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/buckets"
	customattribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/customattribute"
	customsource "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/customsource"
	debugsettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/debugsettings"
	events "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/events"
	grail "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/grail"
	metricslog "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/metrics"
	oneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/oneagent"
	processing "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/processing"
	securitycontextlog "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/securitycontext"
	sensitivedatamasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/sensitivedatamasking"
	storage "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/storage"
	timestamp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/log/timestamp"
	transactionmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mainframe/transactionmonitoring"
	maintenance "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/maintenance/maintenance"
	window "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/maintenance/window"
	backup "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/backup"
	internetproxy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/internetproxy"
	networkzones "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/networkzones"
	preferencesmanaged "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/preferences"
	publicendpoints "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/publicendpoints"
	remoteaccess "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/remoteaccess"
	smtp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/managed/smtp"
	zone "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/management/zone"
	zonev2 "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/management/zonev2"
	eventsmetric "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/metric/events"
	metadata "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/metric/metadata"
	query "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/metric/query"
	permissionmgmz "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mgmz/permission"
	appanomaliesmobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/appanomalies"
	appcrashratemobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/appcrashrate"
	appenablementmobile "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/appenablement"
	appkeyperformance "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/appkeyperformance"
	application "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/application"
	apprequesterrors "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/apprequesterrors"
	notifications "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/mobile/notifications"
	technologiesapache "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesapache"
	technologiesdotnet "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesdotnet"
	technologiesgo "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesgo"
	technologiesiis "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesiis"
	technologiesjava "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesjava"
	technologiesnginx "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesnginx"
	technologiesnodejs "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesnodejs"
	technologiesopentracing "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesopentracing"
	technologiesphp "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesphp"
	technologiespython "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiespython"
	technologiesvarnish "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologiesvarnish"
	technologieswsmb "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/monitored/technologieswsmb"
	emailconnection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ms365/emailconnection"
	connectionmsentraid "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/msentraid/connection"
	connectionmsteams "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/msteams/connection"
	requestsmuted "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/muted/requests"
	nettracer "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/nettracer/nettracer"
	monitornetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/network/monitor"
	monitoroutagenetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/network/monitoroutage"
	traffic "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/network/traffic"
	zonenetwork "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/network/zone"
	zones "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/network/zones"
	notificationnotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/notification/notification"
	defaultmode "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/oneagent/defaultmode"
	defaultversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/oneagent/defaultversion"
	features "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/oneagent/features"
	sidemasking "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/oneagent/sidemasking"
	updatesoneagent "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/oneagent/updates"
	businessevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/businessevents"
	davisevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/davisevents"
	davisproblems "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/davisproblems"
	eventsopenpipeline "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/events"
	logs "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/logs"
	metricsopenpipeline "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/metrics"
	sdlcevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/sdlcevents"
	securityevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/securityevents"
	spans "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/spans"
	systemevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/systemevents"
	userevents "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/userevents"
	usersessions "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/usersessions"
	v2bizeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2bizeventsingestsources"
	v2bizeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2bizeventspipelinegroups"
	v2bizeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2bizeventspipelines"
	v2bizeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2bizeventsrouting"
	v2daviseventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2daviseventsingestsources"
	v2daviseventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2daviseventspipelinegroups"
	v2daviseventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2daviseventspipelines"
	v2daviseventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2daviseventsrouting"
	v2davisproblemsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2davisproblemsingestsources"
	v2davisproblemspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2davisproblemspipelinegroups"
	v2davisproblemspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2davisproblemspipelines"
	v2davisproblemsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2davisproblemsrouting"
	v2eventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventsingestsources"
	v2eventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventspipelinegroups"
	v2eventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventspipelines"
	v2eventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventsrouting"
	v2eventssdlcingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssdlcingestsources"
	v2eventssdlcpipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssdlcpipelinegroups"
	v2eventssdlcpipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssdlcpipelines"
	v2eventssdlcrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssdlcrouting"
	v2eventssecurityingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssecurityingestsources"
	v2eventssecuritypipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssecuritypipelinegroups"
	v2eventssecuritypipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssecuritypipelines"
	v2eventssecurityrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2eventssecurityrouting"
	v2logsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2logsingestsources"
	v2logspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2logspipelinegroups"
	v2logspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2logspipelines"
	v2logsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2logsrouting"
	v2metricsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2metricsingestsources"
	v2metricspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2metricspipelinegroups"
	v2metricspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2metricspipelines"
	v2metricsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2metricsrouting"
	v2securityeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2securityeventsingestsources"
	v2securityeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2securityeventspipelinegroups"
	v2securityeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2securityeventspipelines"
	v2securityeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2securityeventsrouting"
	v2spansingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2spansingestsources"
	v2spanspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2spanspipelinegroups"
	v2spanspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2spanspipelines"
	v2spansrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2spansrouting"
	v2systemeventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2systemeventsingestsources"
	v2systemeventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2systemeventspipelinegroups"
	v2systemeventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2systemeventspipelines"
	v2systemeventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2systemeventsrouting"
	v2usereventsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usereventsingestsources"
	v2usereventspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usereventspipelinegroups"
	v2usereventspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usereventspipelines"
	v2usereventsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usereventsrouting"
	v2usersessionsingestsources "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usersessionsingestsources"
	v2usersessionspipelinegroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usersessionspipelinegroups"
	v2usersessionspipelines "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usersessionspipelines"
	v2usersessionsrouting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/openpipeline/v2usersessionsrouting"
	metricsopentelemetry "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/opentelemetry/metrics"
	genienotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ops/genienotification"
	services "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/os/services"
	config "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ownership/config"
	teams "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/ownership/teams"
	dutynotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/pager/dutynotification"
	connectionpagerduty "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/pagerduty/connection"
	alertingpg "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/pg/alerting"
	anomaliespg "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/pg/anomalies"
	bucket "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/platform/bucket"
	slo "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/platform/slo"
	bindings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/policy/bindings"
	policypolicy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/policy/policy"
	fields "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/problem/fields"
	recordpropagationrules "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/problem/recordpropagationrules"
	availability "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/availability"
	groupdetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/groupdetection"
	groupdetectionflags "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/groupdetectionflags"
	groupmonitoring "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/groupmonitoring"
	grouprum "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/grouprum"
	groupsimpledetection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/groupsimpledetection"
	monitoringprocess "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/monitoring"
	monitoringrule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/monitoringrule"
	visibility "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/process/visibility"
	namingprocessgroup "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/processgroup/naming"
	namingorderprocessgroup "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/processgroup/namingorder"
	providerconfig "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/providerconfig"
	manager "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/queue/manager"
	sharinggroups "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/queue/sharinggroups"
	environments "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/remote/environments"
	report "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/report/report"
	attribute "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/request/attribute"
	namingrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/request/naming"
	namings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/request/namings"
	attributes "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/resource/attributes"
	basedsampling "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rpc/basedsampling"
	advancedcorrelation "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/advancedcorrelation"
	hostheaders "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/hostheaders"
	ipdetermination "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/ipdetermination"
	iplocations "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/iplocations"
	overloadprevention "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/overloadprevention"
	providerbreakdown "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/rum/providerbreakdown"
	context "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/security/context"
	segment "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/segment/segment"
	anomaliesservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/anomalies"
	anomaliesv2service "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/anomaliesv2"
	detectionrulesservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/detectionrules"
	externalwebrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/externalwebrequest"
	externalwebservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/externalwebservice"
	failure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/failure"
	fullwebrequest "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/fullwebrequest"
	fullwebservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/fullwebservice"
	httpfailure "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/httpfailure"
	namingservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/naming"
	namingorderservice "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/namingorder"
	nownotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/nownotification"
	splitting "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/service/splitting"
	connectionservicenow "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/servicenow/connection"
	replayresourcecapture "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/session/replayresourcecapture"
	replaywebprivacy "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/session/replaywebprivacy"
	permissionssettings "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/settings/permissions"
	reliabilityguardian "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/site/reliabilityguardian"
	notificationslack "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/slack/notification"
	normalization "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/slo/normalization"
	sloslo "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/slo/slo"
	v2slo "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/slo/v2"
	attributespan "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/span/attribute"
	capturerule "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/span/capturerule"
	contextpropagation "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/span/contextpropagation"
	entrypoint "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/span/entrypoint"
	eventsspan "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/span/events"
	availabilitysynthetic "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/synthetic/availability"
	location "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/synthetic/location"
	settingstoken "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/token/settings"
	startfilters "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/transaction/startfilters"
	notificationtrello "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/trello/notification"
	servicesmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/unified/servicesmetrics"
	servicesopentel "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/unified/servicesopentel"
	windows "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/update/windows"
	basedsamplingurl "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/url/basedsampling"
	analyticsusability "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/usability/analytics"
	actionmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/actionmetrics"
	experiencescore "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/experiencescore"
	groupuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/group"
	sessionmetrics "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/sessionmetrics"
	settingsuser "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/settings"
	useruser "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/user/user"
	opsnotification "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/victor/opsnotification"
	anomaliesvmware "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vmware/anomalies"
	vmware "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vmware/vmware"
	alertingvulnerability "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/alerting"
	code "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/code"
	settingsvulnerability "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/settings"
	thirdparty "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/thirdparty"
	thirdpartyattr "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/thirdpartyattr"
	thirdpartyk8s "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/vulnerability/thirdpartyk8s"
	appanomaliesweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appanomalies"
	appautoinjection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appautoinjection"
	appbeaconendpoint "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appbeaconendpoint"
	appbeaconorigins "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appbeaconorigins"
	appcustomconfigproperties "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appcustomconfigproperties"
	appcustomerrors "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appcustomerrors"
	appcustominjection "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appcustominjection"
	appcustomproprestrictions "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appcustomproprestrictions"
	appenablementweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appenablement"
	appinjectioncookie "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appinjectioncookie"
	appipaddressexclusion "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appipaddressexclusion"
	appjavascriptfilename "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appjavascriptfilename"
	appjavascriptupdates "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appjavascriptupdates"
	appjavascriptversion "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appjavascriptversion"
	appkeyperformancecustom "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appkeyperformancecustom"
	appkeyperformanceload "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appkeyperformanceload"
	appkeyperformancexhr "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appkeyperformancexhr"
	applicationweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/application"
	appmanualinsertion "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appmanualinsertion"
	apprequesterrorsweb "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/apprequesterrors"
	appresourcecleanup "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appresourcecleanup"
	appresourcetypes "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/web/appresourcetypes"
	notificationwebhook "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/webhook/notification"
	notificationxmatters "github.com/vikreinok/provider-dynatrace-all/internal/controller/cluster/xmatters/notification"
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

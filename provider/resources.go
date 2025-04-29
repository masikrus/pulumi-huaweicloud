// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package huaweicloud

import (
	"fmt"
	"path/filepath"

	"github.com/huaweicloud/pulumi-huaweicloud/provider/pkg/version"
	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "huaweicloud"
	// modules:
	huaweicloudMod         = "index"
	modelArtsMod           = "ModelArts"
	dedicatedApigMod       = "DedicatedApig"    // ?
	sharedApigMod          = "SharedApig"       // ?
	advancedAntiDDosMod    = "AdvancedAntiDDos" // ?
	antiDDosMod            = "AntiDDos"         // ?
	aomMod                 = "Aom"
	asMod                  = "As"
	bmsMod                 = "Bms"
	bcsMod                 = "Bcs"
	cbrMod                 = "Cbr"
	ccMod                  = "Cc"
	cceMod                 = "Cce"
	cciMod                 = "Cci"
	cdmMod                 = "Cdm"
	cesMod                 = "Cse"
	cfwMod                 = "Cfw"
	cptsMod                = "Cpts"
	cssMod                 = "Css"
	cseMod                 = "Cse"
	ctsMod                 = "Cts"
	cloudTableMod          = "CloudTable"
	cdnMod                 = "Cdn"
	dewMod                 = "Dew" // ?
	disMod                 = "Dis"
	dliMod                 = "Dli"
	drsMod                 = "Drs"
	dwsMod                 = "Dws"
	dataArtsStudioMod      = "DataArtsStudio"
	dedicatedElbMod        = "DedicatedElb" // ?
	dcsMod                 = "Dcs"
	dmsMod                 = "Dms"
	ddsMod                 = "Dds"
	dnsMod                 = "Dns"
	ecsMod                 = "Ecs"
	eipMod                 = "Eip"
	elbMod                 = "Elb" // ?
	evsMod                 = "Evs"
	epsMod                 = "Eps"
	erMod                  = "Er"
	functionGraphMod       = "FunctionGraph"
	gaussDBMod             = "GaussDB"
	gaussDBforNoSQLMod     = "GaussDBforNoSQL"
	gaussDBforOpenGaussMod = "GaussDBforOpenGauss"
	iamMod                 = "Iam"
	imsMod                 = "Ims"
	iecMod                 = "Iec"
	ioTDAMod               = "IoTDA"
	liveMod                = "Live"
	ltsMod                 = "Lts"
	mrsMod                 = "Mrs"
	mpcMod                 = "Mpc"
	meetingMod             = "Meeting"
	natMod                 = "Nat"
	networkACLMod          = "NetworkAcl" // ?
	omsMod                 = "Oms"
	obsMod                 = "Obs"
	projectManMod          = "ProjectMan"
	rdsMod                 = "Rds"
	scmMod                 = "Scm"
	sfsMod                 = "Sfs" // ?
	smsMod                 = "Sms"
	serviceStageMod        = "ServiceStage"
	smnMod                 = "Smn"
	swrMod                 = "Swr"
	tmsMod                 = "Tms"
	vpcepMod               = "Vpcep"
	vodMod                 = "Vod"
	vpcMod                 = "Vpc"
	wafMod                 = "Waf"
	workspaceMod           = "Workspace"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(huaweicloud.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "huaweicloud",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Huawei",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/huaweicloud/pulumi-huaweicloud/releases/download/${VERSION}",
		Description:       "A Pulumi package for creating and managing Huaweicloud cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "huaweicloud", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://huaweicloud-pulumi-provider.readthedocs.io",
		Repository: "https://github.com/huaweicloud/pulumi-huaweicloud",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "huaweicloud",
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },

			"huaweicloud_aad_forward_rule": {Tok: tfbridge.MakeResource(mainPkg, advancedAntiDDosMod, "ForwardRule")},

			"huaweicloud_antiddos_basic": {Tok: tfbridge.MakeResource(mainPkg, antiDDosMod, "Basic")},

			"huaweicloud_aom_alarm_rule":             {Tok: tfbridge.MakeResource(mainPkg, aomMod, "AlarmRule")},
			"huaweicloud_aom_service_discovery_rule": {Tok: tfbridge.MakeResource(mainPkg, aomMod, "ServiceDiscoveryRule")},

			"huaweicloud_api_gateway_api":   {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Api")},
			"huaweicloud_api_gateway_group": {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Group")},

			"huaweicloud_apig_api":                         {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Api")},
			"huaweicloud_apig_api_publishment":             {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ApiPublishment")},
			"huaweicloud_apig_instance":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Instance")},
			"huaweicloud_apig_application":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Application")},
			"huaweicloud_apig_custom_authorizer":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "CustomAuthorizer")},
			"huaweicloud_apig_environment":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Environment")},
			"huaweicloud_apig_group":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Group")},
			"huaweicloud_apig_response":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Response")},
			"huaweicloud_apig_throttling_policy_associate": {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicyAssociate")},
			"huaweicloud_apig_throttling_policy":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicy")},
			"huaweicloud_apig_vpc_channel":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "VpcChannel")},

			"huaweicloud_as_configuration":    {Tok: tfbridge.MakeResource(mainPkg, asMod, "Configuration")},
			"huaweicloud_as_group":            {Tok: tfbridge.MakeResource(mainPkg, asMod, "Group")},
			"huaweicloud_as_lifecycle_hook":   {Tok: tfbridge.MakeResource(mainPkg, asMod, "LifecycleHook")},
			"huaweicloud_as_policy":           {Tok: tfbridge.MakeResource(mainPkg, asMod, "Policy")},
			"huaweicloud_as_bandwidth_policy": {Tok: tfbridge.MakeResource(mainPkg, asMod, "BandwidthPolicy")},

			"huaweicloud_bms_instance": {Tok: tfbridge.MakeResource(mainPkg, bmsMod, "Instance")},
			"huaweicloud_bcs_instance": {Tok: tfbridge.MakeResource(mainPkg, bcsMod, "Instance")},

			"huaweicloud_cbr_policy": {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Policy")},
			"huaweicloud_cbr_vault":  {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Vault")},

			"huaweicloud_cc_connection":       {Tok: tfbridge.MakeResource(mainPkg, ccMod, "Connection")},
			"huaweicloud_cc_network_instance": {Tok: tfbridge.MakeResource(mainPkg, ccMod, "NetworkInstance")},

			"huaweicloud_cce_cluster":                    {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Cluster")},
			"huaweicloud_cce_node":                       {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Node")},
			"huaweicloud_cce_node_attach":                {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodeAttach")},
			"huaweicloud_cce_addon":                      {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Addon")},
			"huaweicloud_cce_node_pool":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePool")},
			"huaweicloud_cce_namespace":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Namespace")},
			"huaweicloud_cce_pvc":                        {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Pvc")},
			"huaweicloud_cce_chart":                      {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Chart")},
			"huaweicloud_cce_cluster_certificate_revoke": {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterCertificateRevoke")},
			"huaweicloud_cce_cluster_log_config":         {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterLogConfig")},
			"huaweicloud_cce_cluster_upgrade":            {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterUpgrade")},
			"huaweicloud_cce_node_pool_nodes_add":        {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePoolNodesAdd")},
			"huaweicloud_cce_node_pool_scale":            {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePoolScale")},
			"huaweicloud_cce_node_sync":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodeSync")},

			"huaweicloud_cts_tracker":      {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Tracker")},
			"huaweicloud_cts_data_tracker": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "DataTracker")},
			"huaweicloud_cts_notification": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Notification")},
			"huaweicloud_cci_namespace":    {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Namespace")},
			"huaweicloud_cci_network":      {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Network")},
			"huaweicloud_cci_pvc":          {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Pvc")},

			"huaweicloud_cdm_cluster": {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Cluster")},
			"huaweicloud_cdm_job":     {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Job")},
			"huaweicloud_cdm_link":    {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Link")},

			"huaweicloud_cdn_domain":         {Tok: tfbridge.MakeResource(mainPkg, cdnMod, "Domain")},
			"huaweicloud_ces_alarmrule":      {Tok: tfbridge.MakeResource(mainPkg, cesMod, "Alarmrule")},
			"huaweicloud_cloudtable_cluster": {Tok: tfbridge.MakeResource(mainPkg, cloudTableMod, "Cluster")},

			"huaweicloud_compute_instance":         {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Instance")},
			"huaweicloud_compute_interface_attach": {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "InterfaceAttach")},
			"huaweicloud_compute_keypair":          {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Keypair")},
			"huaweicloud_compute_servergroup":      {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Servergroup")},
			"huaweicloud_compute_eip_associate":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "EipAssociate")},
			"huaweicloud_compute_volume_attach":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "VolumeAttach")},

			"huaweicloud_cfw_acl_rule":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AclRule")},
			"huaweicloud_cfw_address_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroup")},
			"huaweicloud_cfw_address_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroupMember")},
			"huaweicloud_cfw_alarm_config":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AlarmConfig")},
			"huaweicloud_cfw_anti_virus":           {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AntiVirus")},
			"huaweicloud_cfw_black_white_list":     {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "BlackWhiteList")},
			"huaweicloud_cfw_capture_task":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "CaptureTask")},
			"huaweicloud_cfw_dns_resolution":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DnsResolution")},
			"huaweicloud_cfw_domain_name_group":    {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DomainNameGroup")},
			"huaweicloud_cfw_eip_protection":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "EipProtection")},
			"huaweicloud_cfw_firewall":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "Firewall")},
			"huaweicloud_cfw_ips_rule_mode_change": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "IpsRuleModeChange")},
			"huaweicloud_cfw_lts_log":              {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "LtsLog")},
			"huaweicloud_cfw_service_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroup")},
			"huaweicloud_cfw_service_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroupMember")},

			"huaweicloud_cse_microservice":          {Tok: tfbridge.MakeResource(mainPkg, cseMod, "Microservice")},
			"huaweicloud_cse_microservice_engine":   {Tok: tfbridge.MakeResource(mainPkg, cseMod, "MicroserviceEngine")},
			"huaweicloud_cse_microservice_instance": {Tok: tfbridge.MakeResource(mainPkg, cseMod, "MicroserviceInstance")},

			"huaweicloud_csms_secret": {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Secret")},

			"huaweicloud_css_cluster":   {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Cluster")},
			"huaweicloud_css_snapshot":  {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Snapshot")},
			"huaweicloud_css_thesaurus": {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Thesaurus")},

			"huaweicloud_dcs_instance":      {Tok: tfbridge.MakeResource(mainPkg, dcsMod, "Instance")},
			"huaweicloud_dds_database_role": {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "DatabaseRole")},
			"huaweicloud_dds_database_user": {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "DatabaseUser")},
			"huaweicloud_dds_instance":      {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "Instance")},

			"huaweicloud_dis_stream": {Tok: tfbridge.MakeResource(mainPkg, disMod, "Stream")},

			"huaweicloud_dli_database":     {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Database")},
			"huaweicloud_dli_package":      {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Package")},
			"huaweicloud_dli_queue":        {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Queue")},
			"huaweicloud_dli_spark_job":    {Tok: tfbridge.MakeResource(mainPkg, dliMod, "SparkJob")},
			"huaweicloud_dli_sql_job":      {Tok: tfbridge.MakeResource(mainPkg, dliMod, "SqlJob")},
			"huaweicloud_dli_table":        {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Table")},
			"huaweicloud_dli_flinksql_job": {Tok: tfbridge.MakeResource(mainPkg, dliMod, "FlinksqlJob")},
			"huaweicloud_dli_flinkjar_job": {Tok: tfbridge.MakeResource(mainPkg, dliMod, "FlinkjarJob")},
			"huaweicloud_dli_permission":   {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Permission")},

			"huaweicloud_dms_kafka_user":        {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaUser")},
			"huaweicloud_dms_kafka_permissions": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaPermissions")},
			"huaweicloud_dms_kafka_instance":    {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaInstance")},
			"huaweicloud_dms_kafka_topic":       {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaTopic")},
			"huaweicloud_dms_rabbitmq_instance": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "RabbitmqInstance")},

			"huaweicloud_dns_ptrrecord": {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Ptrrecord")},
			"huaweicloud_dns_recordset": {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Recordset")},
			"huaweicloud_dns_zone":      {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Zone")},

			"huaweicloud_drs_job":     {Tok: tfbridge.MakeResource(mainPkg, drsMod, "Job")},
			"huaweicloud_dws_cluster": {Tok: tfbridge.MakeResource(mainPkg, dwsMod, "Cluster")},

			"huaweicloud_elb_certificate":                  {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Certificate")},
			"huaweicloud_elb_l7policy":                     {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "L7policy")},
			"huaweicloud_elb_l7rule":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "L7rule")},
			"huaweicloud_elb_listener":                     {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Listener")},
			"huaweicloud_elb_loadbalancer":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Loadbalancer")},
			"huaweicloud_elb_monitor":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Monitor")},
			"huaweicloud_elb_ipgroup":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Ipgroup")},
			"huaweicloud_elb_pool":                         {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Pool")},
			"huaweicloud_elb_member":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Member")},
			"huaweicloud_elb_active_standby_pool":          {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "ActiveStandbyPool")},
			"huaweicloud_elb_certificate_private_key_echo": {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "CertificatePrivateKeyEcho")},
			"huaweicloud_elb_loadbalancer_copy":            {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "LoadbalancerCopy")},
			"huaweicloud_elb_logtank":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Logtank")},
			"huaweicloud_elb_security_policy":              {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "SecurityPolicy")},

			"huaweicloud_enterprise_project": {Tok: tfbridge.MakeResource(mainPkg, epsMod, "Project")},

			"huaweicloud_er_association":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Association")},
			"huaweicloud_er_attachment_accepter": {Tok: tfbridge.MakeResource(mainPkg, erMod, "AttachmentAccepter")},
			"huaweicloud_er_flow_log":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "FlowLog")},
			"huaweicloud_er_instance":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "Instance")},
			"huaweicloud_er_propagation":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Propagation")},
			"huaweicloud_er_route_table":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "RouteTable")},
			"huaweicloud_er_static_route":        {Tok: tfbridge.MakeResource(mainPkg, erMod, "StaticRoute")},
			"huaweicloud_er_vpc_attachment":      {Tok: tfbridge.MakeResource(mainPkg, erMod, "VpcAttachment")},

			"huaweicloud_evs_snapshot": {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Snapshot")},
			"huaweicloud_evs_volume":   {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Volume")},

			"huaweicloud_fgs_dependency": {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Dependency")},
			"huaweicloud_fgs_function":   {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Function")},
			"huaweicloud_fgs_trigger":    {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Trigger")},

			"huaweicloud_gaussdb_mysql_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBMod, "MysqlInstance")},
			"huaweicloud_gaussdb_mysql_proxy":        {Tok: tfbridge.MakeResource(mainPkg, gaussDBMod, "MysqlProxy")},
			"huaweicloud_gaussdb_opengauss_instance": {Tok: tfbridge.MakeResource(mainPkg, gaussDBforOpenGaussMod, "OpengaussInstance")},
			"huaweicloud_gaussdb_cassandra_instance": {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "CassandraInstance")},
			"huaweicloud_gaussdb_redis_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "RedisInstance")},
			"huaweicloud_gaussdb_influx_instance":    {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "InfluxInstance")},
			"huaweicloud_gaussdb_mongo_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "MongoInstance")},

			"huaweicloud_identity_access_key":            {Tok: tfbridge.MakeResource(mainPkg, iamMod, "AccessKey")},
			"huaweicloud_identity_acl":                   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Acl")},
			"huaweicloud_identity_agency":                {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Agency")},
			"huaweicloud_identity_group":                 {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Group")},
			"huaweicloud_identity_group_membership":      {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupMembership")},
			"huaweicloud_identity_project":               {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Project")},
			"huaweicloud_identity_role":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Role")},
			"huaweicloud_identity_role_assignment":       {Tok: tfbridge.MakeResource(mainPkg, iamMod, "RoleAssignment")},
			"huaweicloud_identity_user":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "User")},
			"huaweicloud_identity_provider":              {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Provider")},
			"huaweicloud_identity_group_role_assignment": {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupRoleAssignment")},
			"huaweicloud_identity_login_policy":          {Tok: tfbridge.MakeResource(mainPkg, iamMod, "LoginPolicy")},
			"huaweicloud_identity_password_policy":       {Tok: tfbridge.MakeResource(mainPkg, iamMod, "PasswordPolicy")},
			"huaweicloud_identity_protection_policy":     {Tok: tfbridge.MakeResource(mainPkg, iamMod, "ProtectionPolicy")},
			"huaweicloud_identity_provider_conversion":   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "ProviderConversion")},
			"huaweicloud_identity_user_role_assignment":  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "UserRoleAssignment")},
			"huaweicloud_identity_user_token":            {Tok: tfbridge.MakeResource(mainPkg, iamMod, "UserToken")},
			"huaweicloud_identity_virtual_mfa_device":    {Tok: tfbridge.MakeResource(mainPkg, iamMod, "VirtualMfaDevice")},

			"huaweicloud_iec_eip":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Eip")},
			"huaweicloud_iec_keypair":             {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Keypair")},
			"huaweicloud_iec_network_acl":         {Tok: tfbridge.MakeResource(mainPkg, iecMod, "NetworkAcl")},
			"huaweicloud_iec_network_acl_rule":    {Tok: tfbridge.MakeResource(mainPkg, iecMod, "NetworkAclRule")},
			"huaweicloud_iec_security_group":      {Tok: tfbridge.MakeResource(mainPkg, iecMod, "SecurityGroup")},
			"huaweicloud_iec_security_group_rule": {Tok: tfbridge.MakeResource(mainPkg, iecMod, "SecurityGroupRule")},
			"huaweicloud_iec_server":              {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Server")},
			"huaweicloud_iec_vip":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Vip")},
			"huaweicloud_iec_vpc":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Vpc")},
			"huaweicloud_iec_vpc_subnet":          {Tok: tfbridge.MakeResource(mainPkg, iecMod, "VpcSubnet")},

			"huaweicloud_images_image":                {Tok: tfbridge.MakeResource(mainPkg, imsMod, "Image")},
			"huaweicloud_images_image_copy":           {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageCopy")},
			"huaweicloud_images_image_share":          {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageShare")},
			"huaweicloud_images_image_share_accepter": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageShareAccepter")},
			"huaweicloud_ims_cbr_whole_image":         {Tok: tfbridge.MakeResource(mainPkg, imsMod, "CbrWholeImage")},
			// "huaweicloud_ims_ecs_system_image":        {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EcsSystemImage")},
			"huaweicloud_ims_ecs_whole_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EcsWholeImage")},
			"huaweicloud_ims_evs_data_image":  {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EvsDataImage")},
			// "huaweicloud_ims_evs_system_image":         {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EvsSystemImage")},
			"huaweicloud_ims_image_export": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageExport")},
			// "huaweicloud_ims_image_metadata":           {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageMetadata")},
			// "huaweicloud_ims_image_registration":       {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageRegistration")},
			"huaweicloud_ims_obs_data_image":   {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsDataImage")},
			"huaweicloud_ims_obs_iso_image":    {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsIsoImage")},
			"huaweicloud_ims_obs_system_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsSystemImage")},
			// "huaweicloud_ims_quickimport_data_image":   {Tok: tfbridge.MakeResource(mainPkg, imsMod, "QuickimportDataImage")},
			// "huaweicloud_ims_quickimport_system_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "QuickimportSystemImage")},

			"huaweicloud_iotda_space":               {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Space")},
			"huaweicloud_iotda_product":             {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Product")},
			"huaweicloud_iotda_device":              {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Device")},
			"huaweicloud_iotda_device_group":        {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceGroup")},
			"huaweicloud_iotda_dataforwarding_rule": {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DataforwardingRule")},
			"huaweicloud_iotda_amqp":                {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Amqp")},
			"huaweicloud_iotda_device_certificate":  {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceCertificate")},
			"huaweicloud_iotda_device_linkage_rule": {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceLinkageRule")},

			"huaweicloud_kms_key":     {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Key")},
			"huaweicloud_kps_keypair": {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Keypair")},

			"huaweicloud_lb_certificate":  {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Certificate")},
			"huaweicloud_lb_l7policy":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7policy")},
			"huaweicloud_lb_l7rule":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7rule")},
			"huaweicloud_lb_listener":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Listener")},
			"huaweicloud_lb_loadbalancer": {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Loadbalancer")},
			"huaweicloud_lb_member":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Member")},
			"huaweicloud_lb_monitor":      {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Monitor")},
			"huaweicloud_lb_pool":         {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Pool")},
			"huaweicloud_lb_whitelist":    {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Whitelist")},

			"huaweicloud_live_domain":          {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Domain")},
			"huaweicloud_live_recording":       {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Recording")},
			"huaweicloud_live_record_callback": {Tok: tfbridge.MakeResource(mainPkg, liveMod, "RecordCallback")},
			"huaweicloud_live_transcoding":     {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Transcoding")},

			"huaweicloud_lts_group":  {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Group")},
			"huaweicloud_lts_stream": {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Stream")},

			"huaweicloud_mapreduce_cluster": {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "cluster")},
			"huaweicloud_mapreduce_job":     {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "Job")},

			"huaweicloud_meeting_admin_assignment": {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "AdminAssignment")},
			"huaweicloud_meeting_conference":       {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "Conference")},
			"huaweicloud_meeting_user":             {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "User")},

			"huaweicloud_modelarts_dataset":                {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "Dataset")},
			"huaweicloud_modelarts_dataset_version":        {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "DatasetVersion")},
			"huaweicloud_modelarts_notebook":               {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "Notebook")},
			"huaweicloud_modelarts_notebook_mount_storage": {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "NotebookMountStorage")},

			"huaweicloud_dataarts_studio_instance": {Tok: tfbridge.MakeResource(mainPkg, dataArtsStudioMod, "StudioInstance")},

			"huaweicloud_mpc_transcoding_template":       {Tok: tfbridge.MakeResource(mainPkg, mpcMod, "TranscodingTemplate")},
			"huaweicloud_mpc_transcoding_template_group": {Tok: tfbridge.MakeResource(mainPkg, mpcMod, "TranscodingTemplateGroup")},

			"huaweicloud_nat_dnat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "DnatRule")},
			"huaweicloud_nat_gateway":   {Tok: tfbridge.MakeResource(mainPkg, natMod, "Gateway")},
			"huaweicloud_nat_snat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "SnatRule")},

			"huaweicloud_network_acl":      {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "Acl")},
			"huaweicloud_network_acl_rule": {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "AclRule")},

			"huaweicloud_obs_bucket":        {Tok: tfbridge.MakeResource(mainPkg, obsMod, "Bucket")},
			"huaweicloud_obs_bucket_object": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketObject")},
			"huaweicloud_obs_bucket_policy": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketPolicy")},

			"huaweicloud_oms_migration_task": {Tok: tfbridge.MakeResource(mainPkg, omsMod, "MigrationTask")},

			"huaweicloud_rds_account":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Account")},
			"huaweicloud_rds_database":                     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Database")},
			"huaweicloud_rds_database_privilege":           {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Database_privilege")},
			"huaweicloud_rds_instance":                     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Instance")},
			"huaweicloud_rds_parametergroup":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Parametergroup")},
			"huaweicloud_rds_read_replica_instance":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "ReadReplicaInstance")},
			"huaweicloud_rds_backup":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Backup")},
			"huaweicloud_rds_cross_region_backup_strategy": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "CrossRegionBackupStrategy")},
			"huaweicloud_rds_database_logs_shrinking":      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "DatabaseLogsShrinking")},
			"huaweicloud_rds_extend_log_link":              {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "ExtendLogLink")},
			"huaweicloud_rds_instance_eip_associate":       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "InstanceEipAssociate")},
			"huaweicloud_rds_instant_task_delete":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "InstantTaskDelete")},
			"huaweicloud_rds_lts_log":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "LtsLog")},
			"huaweicloud_rds_mysql_account":                {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlAccount")},
			"huaweicloud_rds_mysql_binlog":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlBinlog")},
			"huaweicloud_rds_mysql_database":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabase")},
			"huaweicloud_rds_mysql_database_privilege":     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabasePrivilege")},
			"huaweicloud_rds_mysql_database_table_restore": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabaseTableRestore")},
			"huaweicloud_rds_mysql_proxy":                  {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlProxy")},
			"huaweicloud_rds_mysql_proxy_restart":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlProxyRestart")},
			"huaweicloud_rds_pg_account":                   {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccount")},
			"huaweicloud_rds_pg_account_privileges":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccountPrivileges")},
			"huaweicloud_rds_pg_account_roles":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccountRoles")},
			"huaweicloud_rds_pg_database":                  {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgDatabase")},
			"huaweicloud_rds_pg_database_privilege":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgDatabasePrivilege")},
			"huaweicloud_rds_pg_hba":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgHba")},
			"huaweicloud_rds_pg_plugin":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPlugin")},
			"huaweicloud_rds_pg_plugin_parameter":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginParameter")},
			"huaweicloud_rds_pg_plugin_update":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginUpdate")},
			"huaweicloud_rds_pg_sql_limit":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgSqlLimit")},
			"huaweicloud_rds_primary_standby_switch":       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PrimaryStandbySwitch")},
			"huaweicloud_rds_restore":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Restore")},
			"huaweicloud_rds_sql_audit":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlAudit")},
			"huaweicloud_rds_sqlserver_account":            {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverAccount")},
			"huaweicloud_rds_sqlserver_database":           {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabase")},
			"huaweicloud_rds_sqlserver_database_copy":      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabaseCopy")},
			"huaweicloud_rds_sqlserver_database_privilege": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabasePrivilege")},

			"huaweicloud_servicestage_application":                 {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Application")},
			"huaweicloud_servicestage_component_instance":          {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "ComponentInstance")},
			"huaweicloud_servicestage_component":                   {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Component")},
			"huaweicloud_servicestage_environment":                 {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Environment")},
			"huaweicloud_servicestage_repo_token_authorization":    {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "RepoTokenAuthorization")},
			"huaweicloud_servicestage_repo_password_authorization": {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "RepoPasswordAuthorization")},

			"huaweicloud_sfs_access_rule": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "AccessRule")},
			"huaweicloud_sfs_file_system": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "FileSystem")},
			"huaweicloud_sfs_turbo":       {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "Turbo")},

			"huaweicloud_smn_topic":        {Tok: tfbridge.MakeResource(mainPkg, smnMod, "Topic")},
			"huaweicloud_smn_subscription": {Tok: tfbridge.MakeResource(mainPkg, smnMod, "Subscription")},

			"huaweicloud_sms_server_template": {Tok: tfbridge.MakeResource(mainPkg, smsMod, "ServerTemplate")},
			"huaweicloud_sms_task":            {Tok: tfbridge.MakeResource(mainPkg, smsMod, "Task")},

			"huaweicloud_swr_organization":             {Tok: tfbridge.MakeResource(mainPkg, swrMod, "Organization")},
			"huaweicloud_swr_organization_permissions": {Tok: tfbridge.MakeResource(mainPkg, swrMod, "OrganizationPermissions")},
			"huaweicloud_swr_repository":               {Tok: tfbridge.MakeResource(mainPkg, swrMod, "Repository")},
			"huaweicloud_swr_repository_sharing":       {Tok: tfbridge.MakeResource(mainPkg, swrMod, "RepositorySharing")},

			"huaweicloud_tms_tags": {Tok: tfbridge.MakeResource(mainPkg, tmsMod, "Tags")},

			"huaweicloud_vod_media_asset":                {Tok: tfbridge.MakeResource(mainPkg, vodMod, "MediaAsset")},
			"huaweicloud_vod_media_category":             {Tok: tfbridge.MakeResource(mainPkg, vodMod, "MediaCategory")},
			"huaweicloud_vod_transcoding_template_group": {Tok: tfbridge.MakeResource(mainPkg, vodMod, "TranscodingTemplateGroup")},
			"huaweicloud_vod_watermark_template":         {Tok: tfbridge.MakeResource(mainPkg, vodMod, "WatermarkTemplate")},

			"huaweicloud_vpc_bandwidth":     {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Bandwidth")},
			"huaweicloud_vpc_eip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Eip")},
			"huaweicloud_vpc_eip_associate": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "EipAssociate")},

			"huaweicloud_networking_port":          {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Port")},
			"huaweicloud_networking_secgroup":      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Secgroup")},
			"huaweicloud_networking_secgroup_rule": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SecgroupRule")},
			"huaweicloud_networking_vip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vip")},
			"huaweicloud_networking_vip_associate": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "VipAssociate")},

			"huaweicloud_vpc_peering_connection":          {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnection")},
			"huaweicloud_vpc_peering_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnectionAccepter")},
			"huaweicloud_vpc_route_table":                 {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "RouteTable")},
			"huaweicloud_vpc":                             {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vpc")},
			"huaweicloud_vpc_route":                       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Route")},
			"huaweicloud_vpc_subnet":                      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Subnet")},
			"huaweicloud_vpc_address_group":               {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "AddressGroup")},
			"huaweicloud_vpc_flow_log":                    {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "FlowLog")},
			"huaweicloud_vpc_network_acl":                 {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "NetworkAcl")},
			"huaweicloud_vpc_network_interface":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "NetworkInterface")},
			"huaweicloud_vpc_sub_network_interface":       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SubNetworkInterface")},
			"huaweicloud_vpc_subnet_private_ip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SubnetPrivateIp")},
			"huaweicloud_vpc_traffic_mirror_filter":       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorFilter")},
			"huaweicloud_vpc_traffic_mirror_filter_rule":  {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorFilterRule")},
			"huaweicloud_vpc_traffic_mirror_session":      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorSession")},

			"huaweicloud_vpcep_approval": {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Approval")},
			"huaweicloud_vpcep_endpoint": {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Endpoint")},
			"huaweicloud_vpcep_service":  {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Service")},

			"huaweicloud_scm_certificate": {Tok: tfbridge.MakeResource(mainPkg, scmMod, "Certificate")},

			"huaweicloud_waf_certificate":                {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Certificate")},
			"huaweicloud_waf_domain":                     {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Domain")},
			"huaweicloud_waf_policy":                     {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Policy")},
			"huaweicloud_waf_rule_blacklist":             {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleBlacklist")},
			"huaweicloud_waf_rule_data_masking":          {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleDataMasking")},
			"huaweicloud_waf_rule_web_tamper_protection": {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleWebTamperProtection")},
			"huaweicloud_waf_dedicated_instance":         {Tok: tfbridge.MakeResource(mainPkg, wafMod, "DedicatedInstance")},
			"huaweicloud_waf_dedicated_domain":           {Tok: tfbridge.MakeResource(mainPkg, wafMod, "DedicatedDomain")},
			"huaweicloud_waf_instance_group":             {Tok: tfbridge.MakeResource(mainPkg, wafMod, "InstanceGroup")},
			"huaweicloud_waf_instance_group_associate":   {Tok: tfbridge.MakeResource(mainPkg, wafMod, "InstanceGroupAssociate")},
			"huaweicloud_waf_reference_table":            {Tok: tfbridge.MakeResource(mainPkg, wafMod, "ReferenceTable")},

			"huaweicloud_workspace_desktop": {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "Desktop")},
			"huaweicloud_workspace_service": {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "Service")},
			"huaweicloud_workspace_user":    {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "User")},

			"huaweicloud_cpts_project": {Tok: tfbridge.MakeResource(mainPkg, cptsMod, "Project")},
			"huaweicloud_cpts_task":    {Tok: tfbridge.MakeResource(mainPkg, cptsMod, "Task")},

			"huaweicloud_projectman_project": {Tok: tfbridge.MakeResource(mainPkg, projectManMod, "Project")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"huaweicloud_antiddos": {Tok: tfbridge.MakeDataSource(mainPkg, antiDDosMod, "getAntiddos")},

			"huaweicloud_apig_environments": {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedApigMod, "getEnvironments")},

			"huaweicloud_availability_zones": {Tok: tfbridge.MakeDataSource(mainPkg, huaweicloudMod, "getAvailabilityZones")},

			"huaweicloud_bms_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, bmsMod, "getFlavors")},

			"huaweicloud_cbr_vaults": {Tok: tfbridge.MakeDataSource(mainPkg, cbrMod, "getVaults")},

			"huaweicloud_cce_addon_template":         {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getAddonTemplate")},
			"huaweicloud_cce_cluster":                {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getCluster")},
			"huaweicloud_cce_clusters":               {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusters")},
			"huaweicloud_cce_node":                   {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNode")},
			"huaweicloud_cce_nodes":                  {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodes")},
			"huaweicloud_cce_node_pool":              {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodePool")},
			"huaweicloud_cce_addons":                 {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getAddons")},
			"huaweicloud_cce_charts":                 {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getCharts")},
			"huaweicloud_cce_cluster_certificate":    {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusterCertificates")},
			"huaweicloud_cce_cluster_configurations": {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusterConfigurations")},

			"huaweicloud_cci_namespaces": {Tok: tfbridge.MakeDataSource(mainPkg, cciMod, "getNamespaces")},

			"huaweicloud_cdm_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, cdmMod, "getFlavors")},

			"huaweicloud_cdn_domain_statistics": {Tok: tfbridge.MakeDataSource(mainPkg, cdnMod, "getDomainStatistics")},

			"huaweicloud_cfw_access_control_logs":       {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAccessControlLogs")},
			"huaweicloud_cfw_address_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroupMembers")},
			"huaweicloud_cfw_address_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroups")},
			"huaweicloud_cfw_attack_logs":               {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAttackLogs")},
			"huaweicloud_cfw_black_white_lists":         {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getBlackWhiteLists")},
			"huaweicloud_cfw_capture_task_results":      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTaskResults")},
			"huaweicloud_cfw_capture_tasks":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTasks")},
			"huaweicloud_cfw_domain_name_groups":        {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameGroups")},
			"huaweicloud_cfw_domain_name_parse_ip_list": {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameParseIpList")},
			"huaweicloud_cfw_firewalls":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFirewalls")},
			"huaweicloud_cfw_flow_logs":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFlowLogs")},
			"huaweicloud_cfw_ips_custom_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsCustomRules")},
			"huaweicloud_cfw_ips_rule_details":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRuleDetails")},
			"huaweicloud_cfw_ips_rules":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRules")},
			"huaweicloud_cfw_protection_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getProtectionRules")},
			"huaweicloud_cfw_regions":                   {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getRegions")},
			"huaweicloud_cfw_resource_tags":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getResourceTags")},
			"huaweicloud_cfw_service_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroupMembers")},
			"huaweicloud_cfw_service_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroups")},
			"huaweicloud_cfw_tags":                      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getTags")},

			"huaweicloud_compute_flavors":                 {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getFlavors")},
			"huaweicloud_compute_instance":                {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstance")},
			"huaweicloud_compute_instances":               {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstances")},
			"huaweicloud_compute_instance_remote_console": {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstanceRemoteConsole")},
			"huaweicloud_compute_servergroups":            {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getServergroups")},

			"huaweicloud_csms_secret_version": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getCsmsSecretVersion")},
			"huaweicloud_css_flavors":         {Tok: tfbridge.MakeDataSource(mainPkg, cssMod, "getFlavors")},

			"huaweicloud_dcs_flavors":        {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getFlavors")},
			"huaweicloud_dcs_maintainwindow": {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getMaintainwindow")},

			"huaweicloud_dds_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, ddsMod, "getFlavors")},

			"huaweicloud_dms_kafka_flavors":   {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getFlavors")},
			"huaweicloud_dms_kafka_instances": {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getInstances")},
			"huaweicloud_dms_product":         {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getProduct")},
			"huaweicloud_dms_maintainwindow":  {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getMaintainwindow")},

			"huaweicloud_enterprise_project": {Tok: tfbridge.MakeDataSource(mainPkg, epsMod, "getProject")},

			"huaweicloud_er_associations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAssociations")},
			"huaweicloud_er_attachments":        {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAttachments")},
			"huaweicloud_er_availability_zones": {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailabilityZones")},
			"huaweicloud_er_available_routes":   {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailableRoutes")},
			"huaweicloud_er_flow_logs":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getFlowLogs")},
			"huaweicloud_er_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getInstances")},
			"huaweicloud_er_propagations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getPropagations")},
			"huaweicloud_er_quotas":             {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getQuotas")},
			"huaweicloud_er_resource_tags":      {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getResourceTags")},
			"huaweicloud_er_route_tables":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getRouteTables")},
			"huaweicloud_er_tags":               {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getTags")},

			"huaweicloud_evs_volumes": {Tok: tfbridge.MakeDataSource(mainPkg, evsMod, "getVolumes")},

			"huaweicloud_fgs_dependencies": {Tok: tfbridge.MakeDataSource(mainPkg, functionGraphMod, "getDependencies")},

			"huaweicloud_gaussdb_cassandra_dedicated_resource": {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraDedicatedResource")},
			"huaweicloud_gaussdb_cassandra_flavors":            {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraFlavors")},
			"huaweicloud_gaussdb_nosql_flavors":                {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getNosqlFlavors")},
			"huaweicloud_gaussdb_cassandra_instance":           {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraInstance")},
			"huaweicloud_gaussdb_cassandra_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraInstances")},
			"huaweicloud_gaussdb_redis_instance":               {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getRedisInstance")},
			"huaweicloud_gaussdb_opengauss_instance":           {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforOpenGaussMod, "getOpengaussInstance")},
			"huaweicloud_gaussdb_opengauss_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforOpenGaussMod, "getOpengaussInstances")},
			"huaweicloud_gaussdb_mysql_configuration":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlConfiguration")},
			"huaweicloud_gaussdb_mysql_dedicated_resource":     {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlDedicatedResource")},
			"huaweicloud_gaussdb_mysql_flavors":                {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlFlavors")},
			"huaweicloud_gaussdb_mysql_instance":               {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlInstance")},
			"huaweicloud_gaussdb_mysql_instances":              {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlInstances")},

			"huaweicloud_identity_role":                {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getRole")},
			"huaweicloud_identity_custom_role":         {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getCustomRole")},
			"huaweicloud_identity_group":               {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getGroup")},
			"huaweicloud_identity_projects":            {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getProjects")},
			"huaweicloud_identity_users":               {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getUsers")},
			"huaweicloud_identity_agencies":            {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getAgencies")},
			"huaweicloud_identity_permissions":         {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getPermissions")},
			"huaweicloud_identity_providers":           {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getProviders")},
			"huaweicloud_identity_virtual_mfa_devices": {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getVirtualMfaDevices")},

			"huaweicloud_iec_bandwidths":     {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getBandwidths")},
			"huaweicloud_iec_eips":           {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getEips")},
			"huaweicloud_iec_flavors":        {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getFlavors")},
			"huaweicloud_iec_images":         {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getImages")},
			"huaweicloud_iec_keypair":        {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getKeypair")},
			"huaweicloud_iec_network_acl":    {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getNetwork_acl")},
			"huaweicloud_iec_port":           {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getPort")},
			"huaweicloud_iec_security_group": {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getSecurityGroup")},
			"huaweicloud_iec_server":         {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getServer")},
			"huaweicloud_iec_sites":          {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getSites")},
			"huaweicloud_iec_vpc":            {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getVpc")},
			"huaweicloud_iec_vpc_subnets":    {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getVpcSubnets")},

			"huaweicloud_images_image":    {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImage")},
			"huaweicloud_images_images":   {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImages")},
			"huaweicloud_ims_os_versions": {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getOsVersions")},
			"huaweicloud_ims_quotas":      {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getQuotas")},
			"huaweicloud_ims_tags":        {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getTags")},

			"huaweicloud_kms_key":      {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKey")},
			"huaweicloud_kms_data_key": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getDataKey")},
			"huaweicloud_kps_keypairs": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKeypairs")},

			"huaweicloud_lb_listeners":    {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getListeners")},
			"huaweicloud_lb_loadbalancer": {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getLoadbalancer")},
			"huaweicloud_lb_certificate":  {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getCertificate")},
			"huaweicloud_lb_pools":        {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getPools")},

			"huaweicloud_elb_certificate":                         {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getCertificate")},
			"huaweicloud_elb_flavors":                             {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getFlavors")},
			"huaweicloud_elb_pools":                               {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getPools")},
			"huaweicloud_elb_active_standby_pools":                {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getActiveStandbyPools")},
			"huaweicloud_elb_all_members":                         {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getAllMembers")},
			"huaweicloud_elb_availability_zones":                  {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getAvailabilityZones")},
			"huaweicloud_elb_feature_configurations":              {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getFeatureConfigurations")},
			"huaweicloud_elb_ipgroups":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getIpgroups")},
			"huaweicloud_elb_l7policies":                          {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getL7policies")},
			"huaweicloud_elb_l7rules":                             {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getL7rules")},
			"huaweicloud_elb_listeners":                           {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getListeners")},
			"huaweicloud_elb_loadbalancer_feature_configurations": {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLoadbalancerFeatureConfigurations")},
			"huaweicloud_elb_loadbalancers":                       {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLoadbalancers")},
			"huaweicloud_elb_logtanks":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLogtanks")},
			"huaweicloud_elb_monitors":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getMonitors")},
			"huaweicloud_elb_quotas":                              {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getQuotas")},
			"huaweicloud_elb_security_policies":                   {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getSecurityPolicies")},

			"huaweicloud_nat_gateway": {Tok: tfbridge.MakeDataSource(mainPkg, natMod, "getGateway")},

			"huaweicloud_modelarts_datasets":         {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getDatasets")},
			"huaweicloud_modelarts_dataset_versions": {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getDataset_versions")},
			"huaweicloud_modelarts_notebook_images":  {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getNotebookImages")},

			"huaweicloud_obs_buckets":       {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBuckets")},
			"huaweicloud_obs_bucket_object": {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBucketObject")},

			"huaweicloud_rds_flavors":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getFlavors")},
			"huaweicloud_rds_engine_versions":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getEngineVersions")},
			"huaweicloud_rds_instances":                       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getInstances")},
			"huaweicloud_rds_available_flavors":               {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getAvailableFlavors")},
			"huaweicloud_rds_backup_files":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getBackupFiles")},
			"huaweicloud_rds_backups":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getBackups")},
			"huaweicloud_rds_cross_region_backup_instances":   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getCrossRegionBackupInstances")},
			"huaweicloud_rds_cross_region_backups":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getCrossRegionBackups")},
			"huaweicloud_rds_error_log_link":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getErrorLogLink")},
			"huaweicloud_rds_error_logs":                      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getErrorLogs")},
			"huaweicloud_rds_extend_log_files":                {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getExtendLogFiles")},
			"huaweicloud_rds_extend_log_links":                {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getExtendLogLinks")},
			"huaweicloud_rds_mysql_accounts":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlAccounts")},
			"huaweicloud_rds_mysql_binlog":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlBinlog")},
			"huaweicloud_rds_mysql_database_privileges":       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlDatabasePrivileges")},
			"huaweicloud_rds_mysql_databases":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlDatabases")},
			"huaweicloud_rds_mysql_proxies":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlProxies")},
			"huaweicloud_rds_mysql_proxy_flavors":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlProxyFlavors")},
			"huaweicloud_rds_parametergroups":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getParametergroups")},
			"huaweicloud_rds_pg_accounts":                     {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgAccounts")},
			"huaweicloud_rds_pg_databases":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgDatabases")},
			"huaweicloud_rds_pg_plugin_parameter_value_range": {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValueRange")},
			"huaweicloud_rds_pg_plugin_parameter_values":      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValues")},
			"huaweicloud_rds_pg_plugins":                      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPlugins")},
			"huaweicloud_rds_pg_roles":                        {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgRoles")},
			"huaweicloud_rds_pg_sql_limits":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgSqlLimits")},
			"huaweicloud_rds_predefined_tags":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPredefinedTags")},
			"huaweicloud_rds_quotas":                          {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getQuotas")},
			"huaweicloud_rds_recycling_instances":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRecyclingInstances")},
			"huaweicloud_rds_restore_time_ranges":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoreTimeRanges")},
			"huaweicloud_rds_restored_databases":              {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoredDatabases")},
			"huaweicloud_rds_restored_tables":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoredTables")},
			"huaweicloud_rds_slow_log_files":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogFiles")},
			"huaweicloud_rds_slow_log_link":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogLink")},
			"huaweicloud_rds_slow_logs":                       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogs")},
			"huaweicloud_rds_sql_audit_log_links":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditLogLinks")},
			"huaweicloud_rds_sql_audit_logs":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditLogs")},
			"huaweicloud_rds_sql_audit_operations":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditOperations")},
			"huaweicloud_rds_sqlserver_accounts":              {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverAccounts")},
			"huaweicloud_rds_sqlserver_collations":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverCollations")},
			"huaweicloud_rds_sqlserver_database_privileges":   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverDatabasePrivileges")},
			"huaweicloud_rds_sqlserver_databases":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverDatabases")},
			"huaweicloud_rds_storage_types":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getStorageTypes")},
			"huaweicloud_rds_tags":                            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getTags")},

			"huaweicloud_servicestage_component_runtimes": {Tok: tfbridge.MakeDataSource(mainPkg, serviceStageMod, "getComponentRuntimes")},

			"huaweicloud_smn_topics": {Tok: tfbridge.MakeDataSource(mainPkg, smnMod, "getTopics")},

			"huaweicloud_sms_source_servers": {Tok: tfbridge.MakeDataSource(mainPkg, smsMod, "getSourceServers")},

			"huaweicloud_scm_certificates": {Tok: tfbridge.MakeDataSource(mainPkg, scmMod, "getCertificates")},

			"huaweicloud_sfs_file_system": {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getFileSystem")},
			"huaweicloud_sfs_turbos":      {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getTurbos")},

			"huaweicloud_vpc_bandwidth": {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getBandwidth")},
			"huaweicloud_vpc_eip":       {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEip")},
			"huaweicloud_vpc_eips":      {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEips")},

			"huaweicloud_networking_port":           {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPort")},
			"huaweicloud_networking_secgroup":       {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroup")},
			"huaweicloud_networking_secgroups":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroups")},
			"huaweicloud_networking_secgroup_rules": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroupRules")},

			"huaweicloud_vpc":                             {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpc")},
			"huaweicloud_vpcs":                            {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpcs")},
			"huaweicloud_vpc_ids":                         {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getIds")},
			"huaweicloud_vpc_peering_connection":          {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPeeringConnection")},
			"huaweicloud_vpc_route_table":                 {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRouteTable")},
			"huaweicloud_vpc_subnet":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnet")},
			"huaweicloud_vpc_subnets":                     {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnets")},
			"huaweicloud_vpc_subnet_ids":                  {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetIds")},
			"huaweicloud_vpc_address_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getAddressGroups")},
			"huaweicloud_vpc_flow_logs":                   {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getFlowLogs")},
			"huaweicloud_vpc_network_acls":                {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getNetworkAcls")},
			"huaweicloud_vpc_quotas":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getQuotas")},
			"huaweicloud_vpc_routes":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRoutes")},
			"huaweicloud_vpc_sub_network_interfaces":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubNetworkInterfaces")},
			"huaweicloud_vpc_subnet_ip_availabilities":    {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetIpAvailabilities")},
			"huaweicloud_vpc_subnet_private_ips":          {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetPrivateIps")},
			"huaweicloud_vpc_traffic_mirror_filter_rules": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorFilterRules")},
			"huaweicloud_vpc_traffic_mirror_filters":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorFilters")},
			"huaweicloud_vpc_traffic_mirror_sessions":     {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorSessions")},

			"huaweicloud_vpcep_public_services": {Tok: tfbridge.MakeDataSource(mainPkg, vpcepMod, "getPublicServices")},

			"huaweicloud_waf_certificate":         {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getCertificate")},
			"huaweicloud_waf_policies":            {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getPolicies")},
			"huaweicloud_waf_dedicated_instances": {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getDedicatedInstances")},
			"huaweicloud_waf_reference_tables":    {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getReferenceTables")},
			"huaweicloud_waf_instance_groups":     {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getInstanceGroups")},
			"huaweicloud_dws_flavors":             {Tok: tfbridge.MakeDataSource(mainPkg, dwsMod, "getFlaovrs")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@huaweicloudos/pulumi",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/huaweicloud/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

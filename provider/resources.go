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

	"github.com/huaweicloud/terraform-provider-huaweicloud/huaweicloud"
	"github.com/masikmos/pulumi-huaweicloud/provider/pkg/version"
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
		Publisher: "masikmos",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://console-static.huaweicloud.com/static/authui/20210202115135/public/custom/images/logo-en.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "https://github.com/masikmos/pulumi-huaweicloud/releases/download/${VERSION}",
		Description:       "A Pulumi package for creating and managing Huaweicloud cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "huaweicloud", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://huaweicloud-pulumi-provider.readthedocs.io",
		Repository: "https://github.com/masikmos/pulumi-huaweicloud",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "masikmos",
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
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// 	},
			// },

			"huaweicloud_aad_forward_rule": {Tok: tfbridge.MakeResource(mainPkg, advancedAntiDDosMod, "ForwardRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_antiddos_basic": {Tok: tfbridge.MakeResource(mainPkg, antiDDosMod, "Basic"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_aom_alarm_rule":             {Tok: tfbridge.MakeResource(mainPkg, aomMod, "AlarmRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_aom_service_discovery_rule": {Tok: tfbridge.MakeResource(mainPkg, aomMod, "ServiceDiscoveryRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_api_gateway_api":   {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Api"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_api_gateway_group": {Tok: tfbridge.MakeResource(mainPkg, sharedApigMod, "Group"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_apig_api":                         {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Api"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_api_publishment":             {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ApiPublishment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_instance":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_application":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Application"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_custom_authorizer":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "CustomAuthorizer"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_environment":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Environment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_group":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Group"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_response":                    {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "Response"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_throttling_policy_associate": {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicyAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_throttling_policy":           {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "ThrottlingPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_apig_vpc_channel":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedApigMod, "VpcChannel"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_as_configuration":    {Tok: tfbridge.MakeResource(mainPkg, asMod, "Configuration"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_as_group":            {Tok: tfbridge.MakeResource(mainPkg, asMod, "Group"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_as_lifecycle_hook":   {Tok: tfbridge.MakeResource(mainPkg, asMod, "LifecycleHook"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_as_policy":           {Tok: tfbridge.MakeResource(mainPkg, asMod, "Policy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_as_bandwidth_policy": {Tok: tfbridge.MakeResource(mainPkg, asMod, "BandwidthPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_bms_instance": {Tok: tfbridge.MakeResource(mainPkg, bmsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_bcs_instance": {Tok: tfbridge.MakeResource(mainPkg, bcsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cbr_policy": {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Policy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cbr_vault":  {Tok: tfbridge.MakeResource(mainPkg, cbrMod, "Vault"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cc_connection":       {Tok: tfbridge.MakeResource(mainPkg, ccMod, "Connection"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cc_network_instance": {Tok: tfbridge.MakeResource(mainPkg, ccMod, "NetworkInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cce_cluster":                    {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node":                       {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Node"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_attach":                {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodeAttach"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_addon":                      {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Addon"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_pool":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePool"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_namespace":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Namespace"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_pvc":                        {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Pvc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_chart":                      {Tok: tfbridge.MakeResource(mainPkg, cceMod, "Chart"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster_certificate_revoke": {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterCertificateRevoke"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster_log_config":         {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterLogConfig"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster_upgrade":            {Tok: tfbridge.MakeResource(mainPkg, cceMod, "ClusterUpgrade"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_pool_nodes_add":        {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePoolNodesAdd"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_pool_scale":            {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodePoolScale"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_sync":                  {Tok: tfbridge.MakeResource(mainPkg, cceMod, "NodeSync"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cts_tracker":      {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Tracker"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cts_data_tracker": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "DataTracker"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cts_notification": {Tok: tfbridge.MakeResource(mainPkg, ctsMod, "Notification"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cci_namespace":    {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Namespace"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cci_network":      {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Network"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cci_pvc":          {Tok: tfbridge.MakeResource(mainPkg, cciMod, "Pvc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cdm_cluster": {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cdm_job":     {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Job"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cdm_link":    {Tok: tfbridge.MakeResource(mainPkg, cdmMod, "Link"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cdn_domain":         {Tok: tfbridge.MakeResource(mainPkg, cdnMod, "Domain"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ces_alarmrule":      {Tok: tfbridge.MakeResource(mainPkg, cesMod, "Alarmrule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cloudtable_cluster": {Tok: tfbridge.MakeResource(mainPkg, cloudTableMod, "Cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_compute_instance":         {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_interface_attach": {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "InterfaceAttach"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_keypair":          {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Keypair"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_servergroup":      {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "Servergroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_eip_associate":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "EipAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_volume_attach":    {Tok: tfbridge.MakeResource(mainPkg, ecsMod, "VolumeAttach"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cfw_acl_rule":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AclRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_address_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_address_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AddressGroupMember"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_alarm_config":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AlarmConfig"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_anti_virus":           {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "AntiVirus"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_black_white_list":     {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "BlackWhiteList"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_capture_task":         {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "CaptureTask"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_dns_resolution":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DnsResolution"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_domain_name_group":    {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "DomainNameGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_eip_protection":       {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "EipProtection"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_firewall":             {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "Firewall"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_ips_rule_mode_change": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "IpsRuleModeChange"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_lts_log":              {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "LtsLog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_service_group":        {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_service_group_member": {Tok: tfbridge.MakeResource(mainPkg, cfwMod, "ServiceGroupMember"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cse_microservice":          {Tok: tfbridge.MakeResource(mainPkg, cseMod, "Microservice"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cse_microservice_engine":   {Tok: tfbridge.MakeResource(mainPkg, cseMod, "MicroserviceEngine"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cse_microservice_instance": {Tok: tfbridge.MakeResource(mainPkg, cseMod, "MicroserviceInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_csms_secret": {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Secret"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_css_cluster":   {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_css_snapshot":  {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Snapshot"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_css_thesaurus": {Tok: tfbridge.MakeResource(mainPkg, cssMod, "Thesaurus"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dcs_instance":      {Tok: tfbridge.MakeResource(mainPkg, dcsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dds_database_role": {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "DatabaseRole"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dds_database_user": {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "DatabaseUser"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dds_instance":      {Tok: tfbridge.MakeResource(mainPkg, ddsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dis_stream": {Tok: tfbridge.MakeResource(mainPkg, disMod, "Stream"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dli_database":     {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Database"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_package":      {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Package"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_queue":        {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Queue"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_spark_job":    {Tok: tfbridge.MakeResource(mainPkg, dliMod, "SparkJob"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_sql_job":      {Tok: tfbridge.MakeResource(mainPkg, dliMod, "SqlJob"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_table":        {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Table"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_flinksql_job": {Tok: tfbridge.MakeResource(mainPkg, dliMod, "FlinksqlJob"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_flinkjar_job": {Tok: tfbridge.MakeResource(mainPkg, dliMod, "FlinkjarJob"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dli_permission":   {Tok: tfbridge.MakeResource(mainPkg, dliMod, "Permission"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dms_kafka_user":        {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaUser"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_kafka_permissions": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaPermissions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_kafka_instance":    {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_kafka_topic":       {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "KafkaTopic"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_rabbitmq_instance": {Tok: tfbridge.MakeResource(mainPkg, dmsMod, "RabbitmqInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dns_ptrrecord": {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Ptrrecord"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dns_recordset": {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Recordset"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dns_zone":      {Tok: tfbridge.MakeResource(mainPkg, dnsMod, "Zone"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_drs_job":     {Tok: tfbridge.MakeResource(mainPkg, drsMod, "Job"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dws_cluster": {Tok: tfbridge.MakeResource(mainPkg, dwsMod, "Cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_elb_certificate":                  {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Certificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_l7policy":                     {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "L7policy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_l7rule":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "L7rule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_listener":                     {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Listener"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_loadbalancer":                 {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Loadbalancer"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_monitor":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Monitor"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_ipgroup":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Ipgroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_pool":                         {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Pool"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_member":                       {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Member"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_active_standby_pool":          {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "ActiveStandbyPool"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_certificate_private_key_echo": {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "CertificatePrivateKeyEcho"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_loadbalancer_copy":            {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "LoadbalancerCopy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_logtank":                      {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "Logtank"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_security_policy":              {Tok: tfbridge.MakeResource(mainPkg, dedicatedElbMod, "SecurityPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_enterprise_project": {Tok: tfbridge.MakeResource(mainPkg, epsMod, "Project"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_er_association":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Association"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_attachment_accepter": {Tok: tfbridge.MakeResource(mainPkg, erMod, "AttachmentAccepter"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_flow_log":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "FlowLog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_instance":            {Tok: tfbridge.MakeResource(mainPkg, erMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_propagation":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "Propagation"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_route_table":         {Tok: tfbridge.MakeResource(mainPkg, erMod, "RouteTable"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_static_route":        {Tok: tfbridge.MakeResource(mainPkg, erMod, "StaticRoute"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_vpc_attachment":      {Tok: tfbridge.MakeResource(mainPkg, erMod, "VpcAttachment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_evs_snapshot": {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Snapshot"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_evs_volume":   {Tok: tfbridge.MakeResource(mainPkg, evsMod, "Volume"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_fgs_dependency": {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Dependency"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_fgs_function":   {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Function"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_fgs_trigger":    {Tok: tfbridge.MakeResource(mainPkg, functionGraphMod, "Trigger"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_gaussdb_mysql_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBMod, "MysqlInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_proxy":        {Tok: tfbridge.MakeResource(mainPkg, gaussDBMod, "MysqlProxy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_opengauss_instance": {Tok: tfbridge.MakeResource(mainPkg, gaussDBforOpenGaussMod, "OpengaussInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_cassandra_instance": {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "CassandraInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_redis_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "RedisInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_influx_instance":    {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "InfluxInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mongo_instance":     {Tok: tfbridge.MakeResource(mainPkg, gaussDBforNoSQLMod, "MongoInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_identity_access_key":            {Tok: tfbridge.MakeResource(mainPkg, iamMod, "AccessKey"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_acl":                   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Acl"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_agency":                {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Agency"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_group":                 {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Group"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_group_membership":      {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupMembership"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_project":               {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Project"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_role":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Role"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_role_assignment":       {Tok: tfbridge.MakeResource(mainPkg, iamMod, "RoleAssignment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_user":                  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "User"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_provider":              {Tok: tfbridge.MakeResource(mainPkg, iamMod, "Provider"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_group_role_assignment": {Tok: tfbridge.MakeResource(mainPkg, iamMod, "GroupRoleAssignment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_login_policy":          {Tok: tfbridge.MakeResource(mainPkg, iamMod, "LoginPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_password_policy":       {Tok: tfbridge.MakeResource(mainPkg, iamMod, "PasswordPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_protection_policy":     {Tok: tfbridge.MakeResource(mainPkg, iamMod, "ProtectionPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_provider_conversion":   {Tok: tfbridge.MakeResource(mainPkg, iamMod, "ProviderConversion"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_user_role_assignment":  {Tok: tfbridge.MakeResource(mainPkg, iamMod, "UserRoleAssignment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_user_token":            {Tok: tfbridge.MakeResource(mainPkg, iamMod, "UserToken"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_virtual_mfa_device":    {Tok: tfbridge.MakeResource(mainPkg, iamMod, "VirtualMfaDevice"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_iec_eip":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Eip"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_keypair":             {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Keypair"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_network_acl":         {Tok: tfbridge.MakeResource(mainPkg, iecMod, "NetworkAcl"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_network_acl_rule":    {Tok: tfbridge.MakeResource(mainPkg, iecMod, "NetworkAclRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_security_group":      {Tok: tfbridge.MakeResource(mainPkg, iecMod, "SecurityGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_security_group_rule": {Tok: tfbridge.MakeResource(mainPkg, iecMod, "SecurityGroupRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_server":              {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Server"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_vip":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Vip"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_vpc":                 {Tok: tfbridge.MakeResource(mainPkg, iecMod, "Vpc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_vpc_subnet":          {Tok: tfbridge.MakeResource(mainPkg, iecMod, "VpcSubnet"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_images_image":                {Tok: tfbridge.MakeResource(mainPkg, imsMod, "Image"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_images_image_copy":           {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageCopy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_images_image_share":          {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageShare"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_images_image_share_accepter": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageShareAccepter"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_cbr_whole_image":         {Tok: tfbridge.MakeResource(mainPkg, imsMod, "CbrWholeImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_ecs_system_image":        {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EcsSystemImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_ecs_whole_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EcsWholeImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_evs_data_image":  {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EvsDataImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_evs_system_image":         {Tok: tfbridge.MakeResource(mainPkg, imsMod, "EvsSystemImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_image_export": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageExport"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_image_metadata":           {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageMetadata"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_image_registration":       {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ImageRegistration"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_obs_data_image":   {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsDataImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_obs_iso_image":    {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsIsoImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_obs_system_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "ObsSystemImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_quickimport_data_image":   {Tok: tfbridge.MakeResource(mainPkg, imsMod, "QuickimportDataImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			// "huaweicloud_ims_quickimport_system_image": {Tok: tfbridge.MakeResource(mainPkg, imsMod, "QuickimportSystemImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_iotda_space":               {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Space"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_product":             {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Product"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_device":              {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Device"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_device_group":        {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_dataforwarding_rule": {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DataforwardingRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_amqp":                {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "Amqp"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_device_certificate":  {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceCertificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iotda_device_linkage_rule": {Tok: tfbridge.MakeResource(mainPkg, ioTDAMod, "DeviceLinkageRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_kms_key":     {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Key"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_kps_keypair": {Tok: tfbridge.MakeResource(mainPkg, dewMod, "Keypair"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_lb_certificate":  {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Certificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_l7policy":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7policy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_l7rule":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "L7rule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_listener":     {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Listener"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_loadbalancer": {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Loadbalancer"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_member":       {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Member"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_monitor":      {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Monitor"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_pool":         {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Pool"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_whitelist":    {Tok: tfbridge.MakeResource(mainPkg, elbMod, "Whitelist"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_live_domain":          {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Domain"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_live_recording":       {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Recording"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_live_record_callback": {Tok: tfbridge.MakeResource(mainPkg, liveMod, "RecordCallback"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_live_transcoding":     {Tok: tfbridge.MakeResource(mainPkg, liveMod, "Transcoding"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_lts_group":  {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Group"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lts_stream": {Tok: tfbridge.MakeResource(mainPkg, ltsMod, "Stream"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_mapreduce_cluster": {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "cluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_mapreduce_job":     {Tok: tfbridge.MakeResource(mainPkg, mrsMod, "Job"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_meeting_admin_assignment": {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "AdminAssignment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_meeting_conference":       {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "Conference"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_meeting_user":             {Tok: tfbridge.MakeResource(mainPkg, meetingMod, "User"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_modelarts_dataset":                {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "Dataset"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_modelarts_dataset_version":        {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "DatasetVersion"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_modelarts_notebook":               {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "Notebook"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_modelarts_notebook_mount_storage": {Tok: tfbridge.MakeResource(mainPkg, modelArtsMod, "NotebookMountStorage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dataarts_studio_instance": {Tok: tfbridge.MakeResource(mainPkg, dataArtsStudioMod, "StudioInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_mpc_transcoding_template":       {Tok: tfbridge.MakeResource(mainPkg, mpcMod, "TranscodingTemplate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_mpc_transcoding_template_group": {Tok: tfbridge.MakeResource(mainPkg, mpcMod, "TranscodingTemplateGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_nat_dnat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "DnatRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_nat_gateway":   {Tok: tfbridge.MakeResource(mainPkg, natMod, "Gateway"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_nat_snat_rule": {Tok: tfbridge.MakeResource(mainPkg, natMod, "SnatRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_network_acl":      {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "Acl"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_network_acl_rule": {Tok: tfbridge.MakeResource(mainPkg, networkACLMod, "AclRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_obs_bucket":        {Tok: tfbridge.MakeResource(mainPkg, obsMod, "Bucket"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_obs_bucket_object": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketObject"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_obs_bucket_policy": {Tok: tfbridge.MakeResource(mainPkg, obsMod, "BucketPolicy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_oms_migration_task": {Tok: tfbridge.MakeResource(mainPkg, omsMod, "MigrationTask"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_rds_account":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Account"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_database":                     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Database"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_database_privilege":           {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Database_privilege"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_instance":                     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Instance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_parametergroup":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Parametergroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_read_replica_instance":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "ReadReplicaInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_backup":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Backup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_cross_region_backup_strategy": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "CrossRegionBackupStrategy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_database_logs_shrinking":      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "DatabaseLogsShrinking"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_extend_log_link":              {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "ExtendLogLink"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_instance_eip_associate":       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "InstanceEipAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_instant_task_delete":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "InstantTaskDelete"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_lts_log":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "LtsLog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_account":                {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlAccount"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_binlog":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlBinlog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_database":               {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabase"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_database_privilege":     {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabasePrivilege"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_database_table_restore": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlDatabaseTableRestore"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_proxy":                  {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlProxy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_proxy_restart":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "MysqlProxyRestart"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_account":                   {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccount"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_account_privileges":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccountPrivileges"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_account_roles":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgAccountRoles"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_database":                  {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgDatabase"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_database_privilege":        {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgDatabasePrivilege"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_hba":                       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgHba"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugin":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPlugin"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugin_parameter":          {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginParameter"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugin_update":             {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgPluginUpdate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_sql_limit":                 {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PgSqlLimit"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_primary_standby_switch":       {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "PrimaryStandbySwitch"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_restore":                      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "Restore"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sql_audit":                    {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlAudit"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_account":            {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverAccount"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_database":           {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabase"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_database_copy":      {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabaseCopy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_database_privilege": {Tok: tfbridge.MakeResource(mainPkg, rdsMod, "SqlserverDatabasePrivilege"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_servicestage_application":                 {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Application"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_servicestage_component_instance":          {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "ComponentInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_servicestage_component":                   {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Component"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_servicestage_environment":                 {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "Environment"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_servicestage_repo_token_authorization":    {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "RepoTokenAuthorization"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_servicestage_repo_password_authorization": {Tok: tfbridge.MakeResource(mainPkg, serviceStageMod, "RepoPasswordAuthorization"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_sfs_access_rule": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "AccessRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_sfs_file_system": {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "FileSystem"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_sfs_turbo":       {Tok: tfbridge.MakeResource(mainPkg, sfsMod, "Turbo"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_smn_topic":        {Tok: tfbridge.MakeResource(mainPkg, smnMod, "Topic"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_smn_subscription": {Tok: tfbridge.MakeResource(mainPkg, smnMod, "Subscription"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_sms_server_template": {Tok: tfbridge.MakeResource(mainPkg, smsMod, "ServerTemplate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_sms_task":            {Tok: tfbridge.MakeResource(mainPkg, smsMod, "Task"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_swr_organization":             {Tok: tfbridge.MakeResource(mainPkg, swrMod, "Organization"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_swr_organization_permissions": {Tok: tfbridge.MakeResource(mainPkg, swrMod, "OrganizationPermissions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_swr_repository":               {Tok: tfbridge.MakeResource(mainPkg, swrMod, "Repository"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_swr_repository_sharing":       {Tok: tfbridge.MakeResource(mainPkg, swrMod, "RepositorySharing"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_tms_tags": {Tok: tfbridge.MakeResource(mainPkg, tmsMod, "Tags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vod_media_asset":                {Tok: tfbridge.MakeResource(mainPkg, vodMod, "MediaAsset"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vod_media_category":             {Tok: tfbridge.MakeResource(mainPkg, vodMod, "MediaCategory"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vod_transcoding_template_group": {Tok: tfbridge.MakeResource(mainPkg, vodMod, "TranscodingTemplateGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vod_watermark_template":         {Tok: tfbridge.MakeResource(mainPkg, vodMod, "WatermarkTemplate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpc_bandwidth":     {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Bandwidth"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_eip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Eip"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_eip_associate": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "EipAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_networking_port":          {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Port"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_secgroup":      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Secgroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_secgroup_rule": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SecgroupRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_vip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vip"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_vip_associate": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "VipAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpc_peering_connection":          {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnection"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_peering_connection_accepter": {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "PeeringConnectionAccepter"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_route_table":                 {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "RouteTable"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc":                             {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Vpc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_route":                       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Route"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet":                      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "Subnet"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_address_group":               {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "AddressGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_flow_log":                    {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "FlowLog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_network_acl":                 {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "NetworkAcl"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_network_interface":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "NetworkInterface"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_sub_network_interface":       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SubNetworkInterface"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet_private_ip":           {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "SubnetPrivateIp"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_filter":       {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorFilter"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_filter_rule":  {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorFilterRule"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_session":      {Tok: tfbridge.MakeResource(mainPkg, vpcMod, "TrafficMirrorSession"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpcep_approval": {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Approval"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpcep_endpoint": {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Endpoint"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpcep_service":  {Tok: tfbridge.MakeResource(mainPkg, vpcepMod, "Service"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_scm_certificate": {Tok: tfbridge.MakeResource(mainPkg, scmMod, "Certificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_waf_certificate":                {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Certificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_domain":                     {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Domain"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_policy":                     {Tok: tfbridge.MakeResource(mainPkg, wafMod, "Policy"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_rule_blacklist":             {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleBlacklist"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_rule_data_masking":          {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleDataMasking"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_rule_web_tamper_protection": {Tok: tfbridge.MakeResource(mainPkg, wafMod, "RuleWebTamperProtection"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_dedicated_instance":         {Tok: tfbridge.MakeResource(mainPkg, wafMod, "DedicatedInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_dedicated_domain":           {Tok: tfbridge.MakeResource(mainPkg, wafMod, "DedicatedDomain"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_instance_group":             {Tok: tfbridge.MakeResource(mainPkg, wafMod, "InstanceGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_instance_group_associate":   {Tok: tfbridge.MakeResource(mainPkg, wafMod, "InstanceGroupAssociate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_reference_table":            {Tok: tfbridge.MakeResource(mainPkg, wafMod, "ReferenceTable"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_workspace_desktop": {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "Desktop"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_workspace_service": {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "Service"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_workspace_user":    {Tok: tfbridge.MakeResource(mainPkg, workspaceMod, "User"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cpts_project": {Tok: tfbridge.MakeResource(mainPkg, cptsMod, "Project"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cpts_task":    {Tok: tfbridge.MakeResource(mainPkg, cptsMod, "Task"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_projectman_project": {Tok: tfbridge.MakeResource(mainPkg, projectManMod, "Project"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_antiddos": {Tok: tfbridge.MakeDataSource(mainPkg, antiDDosMod, "getAntiddos"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_apig_environments": {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedApigMod, "getEnvironments"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_availability_zones": {Tok: tfbridge.MakeDataSource(mainPkg, huaweicloudMod, "getAvailabilityZones"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_bms_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, bmsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cbr_vaults": {Tok: tfbridge.MakeDataSource(mainPkg, cbrMod, "getVaults"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cce_addon_template":         {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getAddonTemplate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster":                {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getCluster"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_clusters":               {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusters"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node":                   {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNode"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_nodes":                  {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_node_pool":              {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getNodePool"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_addons":                 {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getAddons"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_charts":                 {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getCharts"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster_certificate":    {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusterCertificates"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cce_cluster_configurations": {Tok: tfbridge.MakeDataSource(mainPkg, cceMod, "getClusterConfigurations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cci_namespaces": {Tok: tfbridge.MakeDataSource(mainPkg, cciMod, "getNamespaces"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cdm_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, cdmMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cdn_domain_statistics": {Tok: tfbridge.MakeDataSource(mainPkg, cdnMod, "getDomainStatistics"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_cfw_access_control_logs":       {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAccessControlLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_address_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroupMembers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_address_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAddressGroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_attack_logs":               {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getAttackLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_black_white_lists":         {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getBlackWhiteLists"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_capture_task_results":      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTaskResults"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_capture_tasks":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getCaptureTasks"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_domain_name_groups":        {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameGroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_domain_name_parse_ip_list": {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getDomainNameParseIpList"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_firewalls":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFirewalls"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_flow_logs":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getFlowLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_ips_custom_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsCustomRules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_ips_rule_details":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRuleDetails"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_ips_rules":                 {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getIpsRules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_protection_rules":          {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getProtectionRules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_regions":                   {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getRegions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_resource_tags":             {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getResourceTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_service_group_members":     {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroupMembers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_service_groups":            {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getServiceGroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_cfw_tags":                      {Tok: tfbridge.MakeDataSource(mainPkg, cfwMod, "getTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_compute_flavors":                 {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_instance":                {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_instances":               {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_instance_remote_console": {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getInstanceRemoteConsole"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_compute_servergroups":            {Tok: tfbridge.MakeDataSource(mainPkg, ecsMod, "getServergroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_csms_secret_version": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getCsmsSecretVersion"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_css_flavors":         {Tok: tfbridge.MakeDataSource(mainPkg, cssMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dcs_flavors":        {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dcs_maintainwindow": {Tok: tfbridge.MakeDataSource(mainPkg, dcsMod, "getMaintainwindow"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dds_flavors": {Tok: tfbridge.MakeDataSource(mainPkg, ddsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_dms_kafka_flavors":   {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_kafka_instances": {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_product":         {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getProduct"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dms_maintainwindow":  {Tok: tfbridge.MakeDataSource(mainPkg, dmsMod, "getMaintainwindow"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_enterprise_project": {Tok: tfbridge.MakeDataSource(mainPkg, epsMod, "getProject"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_er_associations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAssociations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_attachments":        {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAttachments"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_availability_zones": {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailabilityZones"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_available_routes":   {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getAvailableRoutes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_flow_logs":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getFlowLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_propagations":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getPropagations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_quotas":             {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getQuotas"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_resource_tags":      {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getResourceTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_route_tables":       {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getRouteTables"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_er_tags":               {Tok: tfbridge.MakeDataSource(mainPkg, erMod, "getTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_evs_volumes": {Tok: tfbridge.MakeDataSource(mainPkg, evsMod, "getVolumes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_fgs_dependencies": {Tok: tfbridge.MakeDataSource(mainPkg, functionGraphMod, "getDependencies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_gaussdb_cassandra_dedicated_resource": {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraDedicatedResource"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_cassandra_flavors":            {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_nosql_flavors":                {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getNosqlFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_cassandra_instance":           {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_cassandra_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getCassandraInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_redis_instance":               {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforNoSQLMod, "getRedisInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_opengauss_instance":           {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforOpenGaussMod, "getOpengaussInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_opengauss_instances":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBforOpenGaussMod, "getOpengaussInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_configuration":          {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlConfiguration"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_dedicated_resource":     {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlDedicatedResource"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_flavors":                {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_instance":               {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlInstance"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_gaussdb_mysql_instances":              {Tok: tfbridge.MakeDataSource(mainPkg, gaussDBMod, "getMysqlInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_identity_role":                {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getRole"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_custom_role":         {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getCustomRole"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_group":               {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_projects":            {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getProjects"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_users":               {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getUsers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_agencies":            {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getAgencies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_permissions":         {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getPermissions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_providers":           {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getProviders"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_identity_virtual_mfa_devices": {Tok: tfbridge.MakeDataSource(mainPkg, iamMod, "getVirtualMfaDevices"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_iec_bandwidths":     {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getBandwidths"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_eips":           {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getEips"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_flavors":        {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_images":         {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getImages"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_keypair":        {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getKeypair"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_network_acl":    {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getNetwork_acl"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_port":           {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getPort"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_security_group": {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getSecurityGroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_server":         {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getServer"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_sites":          {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getSites"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_vpc":            {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getVpc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_iec_vpc_subnets":    {Tok: tfbridge.MakeDataSource(mainPkg, iecMod, "getVpcSubnets"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_images_image":    {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImage"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_images_images":   {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getImages"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_os_versions": {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getOsVersions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_quotas":      {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getQuotas"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_ims_tags":        {Tok: tfbridge.MakeDataSource(mainPkg, imsMod, "getTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_kms_key":      {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKey"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_kms_data_key": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getDataKey"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_kps_keypairs": {Tok: tfbridge.MakeDataSource(mainPkg, dewMod, "getKeypairs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_lb_listeners":    {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getListeners"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_loadbalancer": {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getLoadbalancer"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_certificate":  {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getCertificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_lb_pools":        {Tok: tfbridge.MakeDataSource(mainPkg, elbMod, "getPools"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_elb_certificate":                         {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getCertificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_flavors":                             {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_pools":                               {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getPools"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_active_standby_pools":                {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getActiveStandbyPools"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_all_members":                         {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getAllMembers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_availability_zones":                  {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getAvailabilityZones"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_feature_configurations":              {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getFeatureConfigurations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_ipgroups":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getIpgroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_l7policies":                          {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getL7policies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_l7rules":                             {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getL7rules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_listeners":                           {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getListeners"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_loadbalancer_feature_configurations": {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLoadbalancerFeatureConfigurations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_loadbalancers":                       {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLoadbalancers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_logtanks":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getLogtanks"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_monitors":                            {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getMonitors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_quotas":                              {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getQuotas"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_elb_security_policies":                   {Tok: tfbridge.MakeDataSource(mainPkg, dedicatedElbMod, "getSecurityPolicies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_nat_gateway": {Tok: tfbridge.MakeDataSource(mainPkg, natMod, "getGateway"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_modelarts_datasets":         {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getDatasets"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_modelarts_dataset_versions": {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getDataset_versions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_modelarts_notebook_images":  {Tok: tfbridge.MakeDataSource(mainPkg, modelArtsMod, "getNotebookImages"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_obs_buckets":       {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBuckets"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_obs_bucket_object": {Tok: tfbridge.MakeDataSource(mainPkg, obsMod, "getBucketObject"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_rds_flavors":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_engine_versions":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getEngineVersions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_instances":                       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_available_flavors":               {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getAvailableFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_backup_files":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getBackupFiles"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_backups":                         {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getBackups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_cross_region_backup_instances":   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getCrossRegionBackupInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_cross_region_backups":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getCrossRegionBackups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_error_log_link":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getErrorLogLink"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_error_logs":                      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getErrorLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_extend_log_files":                {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getExtendLogFiles"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_extend_log_links":                {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getExtendLogLinks"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_accounts":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlAccounts"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_binlog":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlBinlog"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_database_privileges":       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlDatabasePrivileges"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_databases":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlDatabases"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_proxies":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlProxies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_mysql_proxy_flavors":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getMysqlProxyFlavors"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_parametergroups":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getParametergroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_accounts":                     {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgAccounts"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_databases":                    {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgDatabases"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugin_parameter_value_range": {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValueRange"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugin_parameter_values":      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPluginParameterValues"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_plugins":                      {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgPlugins"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_roles":                        {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgRoles"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_pg_sql_limits":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPgSqlLimits"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_predefined_tags":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getPredefinedTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_quotas":                          {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getQuotas"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_recycling_instances":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRecyclingInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_restore_time_ranges":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoreTimeRanges"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_restored_databases":              {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoredDatabases"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_restored_tables":                 {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getRestoredTables"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_slow_log_files":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogFiles"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_slow_log_link":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogLink"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_slow_logs":                       {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSlowLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sql_audit_log_links":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditLogLinks"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sql_audit_logs":                  {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sql_audit_operations":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlAuditOperations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_accounts":              {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverAccounts"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_collations":            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverCollations"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_database_privileges":   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverDatabasePrivileges"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_sqlserver_databases":             {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getSqlserverDatabases"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_storage_types":                   {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getStorageTypes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_rds_tags":                            {Tok: tfbridge.MakeDataSource(mainPkg, rdsMod, "getTags"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_servicestage_component_runtimes": {Tok: tfbridge.MakeDataSource(mainPkg, serviceStageMod, "getComponentRuntimes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_smn_topics": {Tok: tfbridge.MakeDataSource(mainPkg, smnMod, "getTopics"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_sms_source_servers": {Tok: tfbridge.MakeDataSource(mainPkg, smsMod, "getSourceServers"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_scm_certificates": {Tok: tfbridge.MakeDataSource(mainPkg, scmMod, "getCertificates"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_sfs_file_system": {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getFileSystem"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_sfs_turbos":      {Tok: tfbridge.MakeDataSource(mainPkg, sfsMod, "getTurbos"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpc_bandwidth": {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getBandwidth"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_eip":       {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEip"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_eips":      {Tok: tfbridge.MakeDataSource(mainPkg, eipMod, "getEips"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_networking_port":           {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPort"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_secgroup":       {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroup"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_secgroups":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_networking_secgroup_rules": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSecgroupRules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpc":                             {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpc"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpcs":                            {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getVpcs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_ids":                         {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getIds"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_peering_connection":          {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getPeeringConnection"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_route_table":                 {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRouteTable"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnet"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnets":                     {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnets"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet_ids":                  {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetIds"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_address_groups":              {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getAddressGroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_flow_logs":                   {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getFlowLogs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_network_acls":                {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getNetworkAcls"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_quotas":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getQuotas"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_routes":                      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getRoutes"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_sub_network_interfaces":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubNetworkInterfaces"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet_ip_availabilities":    {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetIpAvailabilities"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_subnet_private_ips":          {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getSubnetPrivateIps"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_filter_rules": {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorFilterRules"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_filters":      {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorFilters"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_vpc_traffic_mirror_sessions":     {Tok: tfbridge.MakeDataSource(mainPkg, vpcMod, "getTrafficMirrorSessions"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_vpcep_public_services": {Tok: tfbridge.MakeDataSource(mainPkg, vpcepMod, "getPublicServices"), Docs: &tfbridge.DocInfo{AllowMissing: true}},

			"huaweicloud_waf_certificate":         {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getCertificate"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_policies":            {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getPolicies"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_dedicated_instances": {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getDedicatedInstances"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_reference_tables":    {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getReferenceTables"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_waf_instance_groups":     {Tok: tfbridge.MakeDataSource(mainPkg, wafMod, "getInstanceGroups"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
			"huaweicloud_dws_flavors":             {Tok: tfbridge.MakeDataSource(mainPkg, dwsMod, "getFlaovrs"), Docs: &tfbridge.DocInfo{AllowMissing: true}},
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
				fmt.Sprintf("github.com/masikmos/pulumi-%[1]s/sdk/", mainPkg),
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

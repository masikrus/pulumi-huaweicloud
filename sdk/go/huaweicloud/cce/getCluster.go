// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cce

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about the cluster and obtains certificate for accessing cluster information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Cce"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			clusterName := cfg.RequireObject("clusterName")
//			_, err := Cce.GetCluster(ctx, &cce.GetClusterArgs{
//				Name:   pulumi.StringRef(clusterName),
//				Status: pulumi.StringRef("Available"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupClusterResult
	err := ctx.Invoke("huaweicloud:Cce/getCluster:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCluster.
type LookupClusterArgs struct {
	// Specifies the type of the cluster. Possible values: **VirtualMachine**, **ARM64**.
	ClusterType *string `pulumi:"clusterType"`
	// Specifies the ID of the cluster.
	Id *string `pulumi:"id"`
	// Specifies the name of the cluster.
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain the CCE cluster. If omitted, the provider-level
	// region will be used.
	Region *string `pulumi:"region"`
	// Specifies the status of the cluster.
	Status *string `pulumi:"status"`
	// Specifies the VPC ID to which the cluster belongs.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getCluster.
type LookupClusterResult struct {
	// Authentication mode of the cluster, possible values are x509 and rbac. Defaults to **rbac**.
	AuthenticationMode string `pulumi:"authenticationMode"`
	// Charging mode of the cluster.
	BillingMode int `pulumi:"billingMode"`
	// The certificate clusters. Structure is documented below.
	CertificateClusters []GetClusterCertificateCluster `pulumi:"certificateClusters"`
	// The certificate users. Structure is documented below.
	CertificateUsers []GetClusterCertificateUser `pulumi:"certificateUsers"`
	ClusterType      string                      `pulumi:"clusterType"`
	// The version of cluster in string format.
	ClusterVersion string `pulumi:"clusterVersion"`
	// The container network segment.
	ContainerNetworkCidr string `pulumi:"containerNetworkCidr"`
	// The container network type: **overlay_l2** , **underlay_ipvlan**, **vpc-router** or **eni**.
	ContainerNetworkType string `pulumi:"containerNetworkType"`
	// Cluster description.
	Description string `pulumi:"description"`
	// The access addresses of kube-apiserver in the cluster. Structure is documented below.
	Endpoints []GetClusterEndpoint `pulumi:"endpoints"`
	// ENI network segment. Specified when creating a CCE Turbo cluster.
	EniSubnetCidr string `pulumi:"eniSubnetCidr"`
	// The **IPv4 subnet ID** of the subnet where the ENI resides.
	// Specified when creating a CCE Turbo cluster.
	EniSubnetId string `pulumi:"eniSubnetId"`
	// The enterprise project ID of the CCE cluster.
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	// The cluster specification in string format.
	FlavorId string `pulumi:"flavorId"`
	// The ID of the high speed network used to create bare metal nodes.
	HighwaySubnetId string `pulumi:"highwaySubnetId"`
	Id              string `pulumi:"id"`
	// Raw Kubernetes config to be used by kubectl and other compatible tools.
	KubeConfigRaw string `pulumi:"kubeConfigRaw"`
	// Advanced configuration of master nodes. Structure is documented below.
	Masters []GetClusterMaster `pulumi:"masters"`
	// The user name.
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
	// Security group ID of the cluster.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// The service network segment.
	ServiceNetworkCidr string `pulumi:"serviceNetworkCidr"`
	Status             string `pulumi:"status"`
	// The ID of the subnet used to create the node.
	SubnetId string `pulumi:"subnetId"`
	VpcId    string `pulumi:"vpcId"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

// A collection of arguments for invoking getCluster.
type LookupClusterOutputArgs struct {
	// Specifies the type of the cluster. Possible values: **VirtualMachine**, **ARM64**.
	ClusterType pulumi.StringPtrInput `pulumi:"clusterType"`
	// Specifies the ID of the cluster.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies the name of the cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain the CCE cluster. If omitted, the provider-level
	// region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the status of the cluster.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Specifies the VPC ID to which the cluster belongs.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCluster.
type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// Authentication mode of the cluster, possible values are x509 and rbac. Defaults to **rbac**.
func (o LookupClusterResultOutput) AuthenticationMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.AuthenticationMode }).(pulumi.StringOutput)
}

// Charging mode of the cluster.
func (o LookupClusterResultOutput) BillingMode() pulumi.IntOutput {
	return o.ApplyT(func(v LookupClusterResult) int { return v.BillingMode }).(pulumi.IntOutput)
}

// The certificate clusters. Structure is documented below.
func (o LookupClusterResultOutput) CertificateClusters() GetClusterCertificateClusterArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterCertificateCluster { return v.CertificateClusters }).(GetClusterCertificateClusterArrayOutput)
}

// The certificate users. Structure is documented below.
func (o LookupClusterResultOutput) CertificateUsers() GetClusterCertificateUserArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterCertificateUser { return v.CertificateUsers }).(GetClusterCertificateUserArrayOutput)
}

func (o LookupClusterResultOutput) ClusterType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterType }).(pulumi.StringOutput)
}

// The version of cluster in string format.
func (o LookupClusterResultOutput) ClusterVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterVersion }).(pulumi.StringOutput)
}

// The container network segment.
func (o LookupClusterResultOutput) ContainerNetworkCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ContainerNetworkCidr }).(pulumi.StringOutput)
}

// The container network type: **overlay_l2** , **underlay_ipvlan**, **vpc-router** or **eni**.
func (o LookupClusterResultOutput) ContainerNetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ContainerNetworkType }).(pulumi.StringOutput)
}

// Cluster description.
func (o LookupClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// The access addresses of kube-apiserver in the cluster. Structure is documented below.
func (o LookupClusterResultOutput) Endpoints() GetClusterEndpointArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterEndpoint { return v.Endpoints }).(GetClusterEndpointArrayOutput)
}

// ENI network segment. Specified when creating a CCE Turbo cluster.
func (o LookupClusterResultOutput) EniSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EniSubnetCidr }).(pulumi.StringOutput)
}

// The **IPv4 subnet ID** of the subnet where the ENI resides.
// Specified when creating a CCE Turbo cluster.
func (o LookupClusterResultOutput) EniSubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EniSubnetId }).(pulumi.StringOutput)
}

// The enterprise project ID of the CCE cluster.
func (o LookupClusterResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The cluster specification in string format.
func (o LookupClusterResultOutput) FlavorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.FlavorId }).(pulumi.StringOutput)
}

// The ID of the high speed network used to create bare metal nodes.
func (o LookupClusterResultOutput) HighwaySubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.HighwaySubnetId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Raw Kubernetes config to be used by kubectl and other compatible tools.
func (o LookupClusterResultOutput) KubeConfigRaw() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.KubeConfigRaw }).(pulumi.StringOutput)
}

// Advanced configuration of master nodes. Structure is documented below.
func (o LookupClusterResultOutput) Masters() GetClusterMasterArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []GetClusterMaster { return v.Masters }).(GetClusterMasterArrayOutput)
}

// The user name.
func (o LookupClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Region }).(pulumi.StringOutput)
}

// Security group ID of the cluster.
func (o LookupClusterResultOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The service network segment.
func (o LookupClusterResultOutput) ServiceNetworkCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ServiceNetworkCidr }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Status }).(pulumi.StringOutput)
}

// The ID of the subnet used to create the node.
func (o LookupClusterResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupClusterResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}

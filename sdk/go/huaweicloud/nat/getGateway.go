// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get an available public NAT gateway within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Nat"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Nat"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			gatewayName := cfg.RequireObject("gatewayName")
//			_, err := Nat.GetGateway(ctx, &nat.GetGatewayArgs{
//				Name: pulumi.StringRef(gatewayName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupGateway(ctx *pulumi.Context, args *LookupGatewayArgs, opts ...pulumi.InvokeOption) (*LookupGatewayResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupGatewayResult
	err := ctx.Invoke("huaweicloud:Nat/getGateway:getGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGateway.
type LookupGatewayArgs struct {
	// Specifies the description of the NAT gateway. The value contains 0 to 255
	// characters, and angle brackets (<)
	// and (>) are not allowed.
	Description *string `pulumi:"description"`
	// Specifies the enterprise project ID of the NAT gateway.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the ID of the NAT gateway.
	Id *string `pulumi:"id"`
	// Deprecated: use subnet_id instead
	InternalNetworkId *string `pulumi:"internalNetworkId"`
	// Specifies the public NAT gateway name.\
	// The valid length is limited from `1` to `64`, only letters, digits, hyphens (-) and underscores (_) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the NAT gateway resource. If omitted, the
	// provider-level region will be used.
	Region *string `pulumi:"region"`
	// Deprecated: use vpc_id instead
	RouterId *string `pulumi:"routerId"`
	// The public NAT gateway type. The valid values are as follows:
	// + **1**: Small type, which supports up to `10,000` SNAT connections.
	// + **2**: Medium type, which supports up to `50,000` SNAT connections.
	// + **3**: Large type, which supports up to `200,000` SNAT connections.
	// + **4**: Extra-large type, which supports up to `1,000,000` SNAT connections.
	Spec *string `pulumi:"spec"`
	// Specifies the status of the NAT gateway.
	Status *string `pulumi:"status"`
	// Specifies the subnet ID of the downstream interface (the next hop of the DVR) of the
	// public NAT gateway.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the ID of the VPC this public NAT gateway belongs to.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getGateway.
type LookupGatewayResult struct {
	Description         string `pulumi:"description"`
	EnterpriseProjectId string `pulumi:"enterpriseProjectId"`
	Id                  string `pulumi:"id"`
	// Deprecated: use subnet_id instead
	InternalNetworkId *string `pulumi:"internalNetworkId"`
	Name              string  `pulumi:"name"`
	Region            string  `pulumi:"region"`
	// Deprecated: use vpc_id instead
	RouterId *string `pulumi:"routerId"`
	Spec     string  `pulumi:"spec"`
	Status   string  `pulumi:"status"`
	SubnetId string  `pulumi:"subnetId"`
	VpcId    string  `pulumi:"vpcId"`
}

func LookupGatewayOutput(ctx *pulumi.Context, args LookupGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGatewayResult, error) {
			args := v.(LookupGatewayArgs)
			r, err := LookupGateway(ctx, &args, opts...)
			var s LookupGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGatewayResultOutput)
}

// A collection of arguments for invoking getGateway.
type LookupGatewayOutputArgs struct {
	// Specifies the description of the NAT gateway. The value contains 0 to 255
	// characters, and angle brackets (<)
	// and (>) are not allowed.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Specifies the enterprise project ID of the NAT gateway.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the ID of the NAT gateway.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Deprecated: use subnet_id instead
	InternalNetworkId pulumi.StringPtrInput `pulumi:"internalNetworkId"`
	// Specifies the public NAT gateway name.\
	// The valid length is limited from `1` to `64`, only letters, digits, hyphens (-) and underscores (_) are allowed.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to create the NAT gateway resource. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Deprecated: use vpc_id instead
	RouterId pulumi.StringPtrInput `pulumi:"routerId"`
	// The public NAT gateway type. The valid values are as follows:
	// + **1**: Small type, which supports up to `10,000` SNAT connections.
	// + **2**: Medium type, which supports up to `50,000` SNAT connections.
	// + **3**: Large type, which supports up to `200,000` SNAT connections.
	// + **4**: Extra-large type, which supports up to `1,000,000` SNAT connections.
	Spec pulumi.StringPtrInput `pulumi:"spec"`
	// Specifies the status of the NAT gateway.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// Specifies the subnet ID of the downstream interface (the next hop of the DVR) of the
	// public NAT gateway.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Specifies the ID of the VPC this public NAT gateway belongs to.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (LookupGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getGateway.
type LookupGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGatewayResult)(nil)).Elem()
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutput() LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) ToLookupGatewayResultOutputWithContext(ctx context.Context) LookupGatewayResultOutput {
	return o
}

func (o LookupGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

// Deprecated: use subnet_id instead
func (o LookupGatewayResultOutput) InternalNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.InternalNetworkId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Region }).(pulumi.StringOutput)
}

// Deprecated: use vpc_id instead
func (o LookupGatewayResultOutput) RouterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGatewayResult) *string { return v.RouterId }).(pulumi.StringPtrOutput)
}

func (o LookupGatewayResultOutput) Spec() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Spec }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupGatewayResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGatewayResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGatewayResultOutput{})
}

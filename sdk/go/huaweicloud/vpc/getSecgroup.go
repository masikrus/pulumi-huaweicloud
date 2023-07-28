// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available HuaweiCloud security group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Vpc.GetSecgroup(ctx, &vpc.GetSecgroupArgs{
//				Name: pulumi.StringRef("tf_test_secgroup"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupSecgroup(ctx *pulumi.Context, args *LookupSecgroupArgs, opts ...pulumi.InvokeOption) (*LookupSecgroupResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSecgroupResult
	err := ctx.Invoke("huaweicloud:Vpc/getSecgroup:getSecgroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecgroup.
type LookupSecgroupArgs struct {
	// Specifies the enterprise project ID of the security group.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the name of the security group.
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain the security group. If omitted, the
	// provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifiest he ID of the security group.
	SecgroupId *string `pulumi:"secgroupId"`
}

// A collection of values returned by getSecgroup.
type LookupSecgroupResult struct {
	// The creation time, in UTC format.
	CreatedAt string `pulumi:"createdAt"`
	// The supplementary information about the security group rule.
	Description         string  `pulumi:"description"`
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
	// The array of security group rules associating with the security group.
	// The rule object is documented below.
	Rules      []GetSecgroupRuleType `pulumi:"rules"`
	SecgroupId *string               `pulumi:"secgroupId"`
	// The last update time, in UTC format.
	UpdatedAt string `pulumi:"updatedAt"`
}

func LookupSecgroupOutput(ctx *pulumi.Context, args LookupSecgroupOutputArgs, opts ...pulumi.InvokeOption) LookupSecgroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecgroupResult, error) {
			args := v.(LookupSecgroupArgs)
			r, err := LookupSecgroup(ctx, &args, opts...)
			var s LookupSecgroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecgroupResultOutput)
}

// A collection of arguments for invoking getSecgroup.
type LookupSecgroupOutputArgs struct {
	// Specifies the enterprise project ID of the security group.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the name of the security group.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain the security group. If omitted, the
	// provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifiest he ID of the security group.
	SecgroupId pulumi.StringPtrInput `pulumi:"secgroupId"`
}

func (LookupSecgroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecgroupArgs)(nil)).Elem()
}

// A collection of values returned by getSecgroup.
type LookupSecgroupResultOutput struct{ *pulumi.OutputState }

func (LookupSecgroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecgroupResult)(nil)).Elem()
}

func (o LookupSecgroupResultOutput) ToLookupSecgroupResultOutput() LookupSecgroupResultOutput {
	return o
}

func (o LookupSecgroupResultOutput) ToLookupSecgroupResultOutputWithContext(ctx context.Context) LookupSecgroupResultOutput {
	return o
}

// The creation time, in UTC format.
func (o LookupSecgroupResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The supplementary information about the security group rule.
func (o LookupSecgroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupSecgroupResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecgroupResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupSecgroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecgroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecgroupResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.Region }).(pulumi.StringOutput)
}

// The array of security group rules associating with the security group.
// The rule object is documented below.
func (o LookupSecgroupResultOutput) Rules() GetSecgroupRuleTypeArrayOutput {
	return o.ApplyT(func(v LookupSecgroupResult) []GetSecgroupRuleType { return v.Rules }).(GetSecgroupRuleTypeArrayOutput)
}

func (o LookupSecgroupResultOutput) SecgroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecgroupResult) *string { return v.SecgroupId }).(pulumi.StringPtrOutput)
}

// The last update time, in UTC format.
func (o LookupSecgroupResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecgroupResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecgroupResultOutput{})
}

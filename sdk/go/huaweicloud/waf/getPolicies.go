// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of WAF policies.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			policyName := cfg.RequireObject("policyName")
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			_, err := Waf.GetPolicies(ctx, &waf.GetPoliciesArgs{
//				Name:                pulumi.StringRef(policyName),
//				EnterpriseProjectId: pulumi.StringRef(enterpriseProjectId),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPolicies(ctx *pulumi.Context, args *GetPoliciesArgs, opts ...pulumi.InvokeOption) (*GetPoliciesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPoliciesResult
	err := ctx.Invoke("huaweicloud:Waf/getPolicies:getPolicies", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesArgs struct {
	// Specifies the enterprise project ID of WAF policies.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Policy name used for matching. The value is case sensitive and supports fuzzy matching.
	Name *string `pulumi:"name"`
	// The region in which to obtain the WAF policies. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getPolicies.
type GetPoliciesResult struct {
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The WAF policy name.
	Name *string `pulumi:"name"`
	// A list of WAF policies.
	Policies []GetPoliciesPolicy `pulumi:"policies"`
	Region   string              `pulumi:"region"`
}

func GetPoliciesOutput(ctx *pulumi.Context, args GetPoliciesOutputArgs, opts ...pulumi.InvokeOption) GetPoliciesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPoliciesResult, error) {
			args := v.(GetPoliciesArgs)
			r, err := GetPolicies(ctx, &args, opts...)
			var s GetPoliciesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPoliciesResultOutput)
}

// A collection of arguments for invoking getPolicies.
type GetPoliciesOutputArgs struct {
	// Specifies the enterprise project ID of WAF policies.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Policy name used for matching. The value is case sensitive and supports fuzzy matching.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the WAF policies. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetPoliciesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesArgs)(nil)).Elem()
}

// A collection of values returned by getPolicies.
type GetPoliciesResultOutput struct{ *pulumi.OutputState }

func (GetPoliciesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoliciesResult)(nil)).Elem()
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutput() GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) ToGetPoliciesResultOutputWithContext(ctx context.Context) GetPoliciesResultOutput {
	return o
}

func (o GetPoliciesResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoliciesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoliciesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The WAF policy name.
func (o GetPoliciesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoliciesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of WAF policies.
func (o GetPoliciesResultOutput) Policies() GetPoliciesPolicyArrayOutput {
	return o.ApplyT(func(v GetPoliciesResult) []GetPoliciesPolicy { return v.Policies }).(GetPoliciesPolicyArrayOutput)
}

func (o GetPoliciesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoliciesResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoliciesResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the certificate in the WAF, including the one pushed from SCM.
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
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			_, err := Waf.GetCertificate(ctx, &waf.GetCertificateArgs{
//				Name:                "certificate name",
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
func LookupCertificate(ctx *pulumi.Context, args *LookupCertificateArgs, opts ...pulumi.InvokeOption) (*LookupCertificateResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCertificateResult
	err := ctx.Invoke("huaweicloud:Waf/getCertificate:getCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateArgs struct {
	// The enterprise project ID of WAF certificate.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The expire status of certificate. Defaults is `0`. The value can be:
	// + `0`: not expire
	// + `1`: has expired
	// + `2`: wil expired soon
	ExpireStatus *int `pulumi:"expireStatus"`
	// The name of certificate. The value is case sensitive and supports fuzzy matching.
	Name string `pulumi:"name"`
	// The region in which to obtain the WAF. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getCertificate.
type LookupCertificateResult struct {
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Indicates the time when the certificate expires.
	Expiration   string `pulumi:"expiration"`
	ExpireStatus *int   `pulumi:"expireStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Name   string `pulumi:"name"`
	Region string `pulumi:"region"`
}

func LookupCertificateOutput(ctx *pulumi.Context, args LookupCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCertificateResult, error) {
			args := v.(LookupCertificateArgs)
			r, err := LookupCertificate(ctx, &args, opts...)
			var s LookupCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCertificateResultOutput)
}

// A collection of arguments for invoking getCertificate.
type LookupCertificateOutputArgs struct {
	// The enterprise project ID of WAF certificate.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// The expire status of certificate. Defaults is `0`. The value can be:
	// + `0`: not expire
	// + `1`: has expired
	// + `2`: wil expired soon
	ExpireStatus pulumi.IntPtrInput `pulumi:"expireStatus"`
	// The name of certificate. The value is case sensitive and supports fuzzy matching.
	Name pulumi.StringInput `pulumi:"name"`
	// The region in which to obtain the WAF. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getCertificate.
type LookupCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateResult)(nil)).Elem()
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutput() LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) ToLookupCertificateResultOutputWithContext(ctx context.Context) LookupCertificateResultOutput {
	return o
}

func (o LookupCertificateResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// Indicates the time when the certificate expires.
func (o LookupCertificateResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Expiration }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) ExpireStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCertificateResult) *int { return v.ExpireStatus }).(pulumi.IntPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCertificateResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateResultOutput{})
}

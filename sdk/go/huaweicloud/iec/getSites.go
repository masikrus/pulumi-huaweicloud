// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the available of HuaweiCloud IEC sites.
//
// ## Example Usage
// ### Basic IEC Sites
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iec.GetSites(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSites(ctx *pulumi.Context, args *GetSitesArgs, opts ...pulumi.InvokeOption) (*GetSitesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSitesResult
	err := ctx.Invoke("huaweicloud:Iec/getSites:getSites", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSites.
type GetSitesArgs struct {
	// Specifies the area of the IEC sites located.
	Area *string `pulumi:"area"`
	// Specifies the city of the IEC sites located.
	City *string `pulumi:"city"`
	// Specifies the province of the IEC sites located.
	Province *string `pulumi:"province"`
}

// A collection of values returned by getSites.
type GetSitesResult struct {
	// The area of the IEC service site located.
	Area *string `pulumi:"area"`
	// The city of the IEC service site located.
	City *string `pulumi:"city"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The province of the IEC service site located.
	Province *string `pulumi:"province"`
	// An array of one or more IEC service sites. The sites object structure is documented below.
	Sites []GetSitesSite `pulumi:"sites"`
}

func GetSitesOutput(ctx *pulumi.Context, args GetSitesOutputArgs, opts ...pulumi.InvokeOption) GetSitesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSitesResult, error) {
			args := v.(GetSitesArgs)
			r, err := GetSites(ctx, &args, opts...)
			var s GetSitesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSitesResultOutput)
}

// A collection of arguments for invoking getSites.
type GetSitesOutputArgs struct {
	// Specifies the area of the IEC sites located.
	Area pulumi.StringPtrInput `pulumi:"area"`
	// Specifies the city of the IEC sites located.
	City pulumi.StringPtrInput `pulumi:"city"`
	// Specifies the province of the IEC sites located.
	Province pulumi.StringPtrInput `pulumi:"province"`
}

func (GetSitesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSitesArgs)(nil)).Elem()
}

// A collection of values returned by getSites.
type GetSitesResultOutput struct{ *pulumi.OutputState }

func (GetSitesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSitesResult)(nil)).Elem()
}

func (o GetSitesResultOutput) ToGetSitesResultOutput() GetSitesResultOutput {
	return o
}

func (o GetSitesResultOutput) ToGetSitesResultOutputWithContext(ctx context.Context) GetSitesResultOutput {
	return o
}

// The area of the IEC service site located.
func (o GetSitesResultOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSitesResult) *string { return v.Area }).(pulumi.StringPtrOutput)
}

// The city of the IEC service site located.
func (o GetSitesResultOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSitesResult) *string { return v.City }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetSitesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSitesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The province of the IEC service site located.
func (o GetSitesResultOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSitesResult) *string { return v.Province }).(pulumi.StringPtrOutput)
}

// An array of one or more IEC service sites. The sites object structure is documented below.
func (o GetSitesResultOutput) Sites() GetSitesSiteArrayOutput {
	return o.ApplyT(func(v GetSitesResult) []GetSitesSite { return v.Sites }).(GetSitesSiteArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSitesResultOutput{})
}

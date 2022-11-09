// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the available of HuaweiCloud IEC flavors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Iec"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			flavorName := "c6.large.2"
//			if param := cfg.Get("flavorName"); param != "" {
//				flavorName = param
//			}
//			_, err := Iec.GetFlavors(ctx, &iec.GetFlavorsArgs{
//				Name: pulumi.StringRef(flavorName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetFlavors(ctx *pulumi.Context, args *GetFlavorsArgs, opts ...pulumi.InvokeOption) (*GetFlavorsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFlavorsResult
	err := ctx.Invoke("huaweicloud:Iec/getFlavors:getFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlavors.
type GetFlavorsArgs struct {
	// Specifies the province of the iec instance located.
	Area *string `pulumi:"area"`
	// Specifies the province of the iec instance located.
	City *string `pulumi:"city"`
	// Specifies the flavor name, which can be queried with a regular expression.
	Name *string `pulumi:"name"`
	// Specifies the operator supported of the iec instance.
	Operator *string `pulumi:"operator"`
	// Specifies the province of the iec instance located.
	Province *string `pulumi:"province"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
	// Specifies the list of edge service site.
	SiteIds *string `pulumi:"siteIds"`
}

// A collection of values returned by getFlavors.
type GetFlavorsResult struct {
	Area *string `pulumi:"area"`
	City *string `pulumi:"city"`
	// An array of one or more flavors. The flavors object structure is documented below.
	Flavors []GetFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the iec flavor.
	Name     *string `pulumi:"name"`
	Operator *string `pulumi:"operator"`
	Province *string `pulumi:"province"`
	Region   string  `pulumi:"region"`
	SiteIds  *string `pulumi:"siteIds"`
}

func GetFlavorsOutput(ctx *pulumi.Context, args GetFlavorsOutputArgs, opts ...pulumi.InvokeOption) GetFlavorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFlavorsResult, error) {
			args := v.(GetFlavorsArgs)
			r, err := GetFlavors(ctx, &args, opts...)
			var s GetFlavorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFlavorsResultOutput)
}

// A collection of arguments for invoking getFlavors.
type GetFlavorsOutputArgs struct {
	// Specifies the province of the iec instance located.
	Area pulumi.StringPtrInput `pulumi:"area"`
	// Specifies the province of the iec instance located.
	City pulumi.StringPtrInput `pulumi:"city"`
	// Specifies the flavor name, which can be queried with a regular expression.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the operator supported of the iec instance.
	Operator pulumi.StringPtrInput `pulumi:"operator"`
	// Specifies the province of the iec instance located.
	Province pulumi.StringPtrInput `pulumi:"province"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the list of edge service site.
	SiteIds pulumi.StringPtrInput `pulumi:"siteIds"`
}

func (GetFlavorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlavorsArgs)(nil)).Elem()
}

// A collection of values returned by getFlavors.
type GetFlavorsResultOutput struct{ *pulumi.OutputState }

func (GetFlavorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFlavorsResult)(nil)).Elem()
}

func (o GetFlavorsResultOutput) ToGetFlavorsResultOutput() GetFlavorsResultOutput {
	return o
}

func (o GetFlavorsResultOutput) ToGetFlavorsResultOutputWithContext(ctx context.Context) GetFlavorsResultOutput {
	return o
}

func (o GetFlavorsResultOutput) Area() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Area }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.City }).(pulumi.StringPtrOutput)
}

// An array of one or more flavors. The flavors object structure is documented below.
func (o GetFlavorsResultOutput) Flavors() GetFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetFlavorsResult) []GetFlavorsFlavor { return v.Flavors }).(GetFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the iec flavor.
func (o GetFlavorsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) Province() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Province }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetFlavorsResultOutput) SiteIds() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.SiteIds }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlavorsResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gaussdb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get available HuaweiCloud gaussdb mysql flavors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/GaussDB"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := GaussDB.GetMysqlFlavors(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMysqlFlavors(ctx *pulumi.Context, args *GetMysqlFlavorsArgs, opts ...pulumi.InvokeOption) (*GetMysqlFlavorsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMysqlFlavorsResult
	err := ctx.Invoke("huaweicloud:GaussDB/getMysqlFlavors:getMysqlFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMysqlFlavors.
type GetMysqlFlavorsArgs struct {
	// Specifies the availability zone mode. Currently support `single` and '
	// multi'. Defaults to `single`.
	AvailabilityZoneMode *string `pulumi:"availabilityZoneMode"`
	// Specifies the database engine. Only "gaussdb-mysql" is supported now.
	Engine *string `pulumi:"engine"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
	// Specifies the database version. Only "8.0" is supported now.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getMysqlFlavors.
type GetMysqlFlavorsResult struct {
	AvailabilityZoneMode *string `pulumi:"availabilityZoneMode"`
	Engine               *string `pulumi:"engine"`
	// Indicates the flavors information. Structure is documented below.
	Flavors []GetMysqlFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Region string `pulumi:"region"`
	// Indicates the database version.
	Version *string `pulumi:"version"`
}

func GetMysqlFlavorsOutput(ctx *pulumi.Context, args GetMysqlFlavorsOutputArgs, opts ...pulumi.InvokeOption) GetMysqlFlavorsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMysqlFlavorsResult, error) {
			args := v.(GetMysqlFlavorsArgs)
			r, err := GetMysqlFlavors(ctx, &args, opts...)
			var s GetMysqlFlavorsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMysqlFlavorsResultOutput)
}

// A collection of arguments for invoking getMysqlFlavors.
type GetMysqlFlavorsOutputArgs struct {
	// Specifies the availability zone mode. Currently support `single` and '
	// multi'. Defaults to `single`.
	AvailabilityZoneMode pulumi.StringPtrInput `pulumi:"availabilityZoneMode"`
	// Specifies the database engine. Only "gaussdb-mysql" is supported now.
	Engine pulumi.StringPtrInput `pulumi:"engine"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the database version. Only "8.0" is supported now.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (GetMysqlFlavorsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlFlavorsArgs)(nil)).Elem()
}

// A collection of values returned by getMysqlFlavors.
type GetMysqlFlavorsResultOutput struct{ *pulumi.OutputState }

func (GetMysqlFlavorsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMysqlFlavorsResult)(nil)).Elem()
}

func (o GetMysqlFlavorsResultOutput) ToGetMysqlFlavorsResultOutput() GetMysqlFlavorsResultOutput {
	return o
}

func (o GetMysqlFlavorsResultOutput) ToGetMysqlFlavorsResultOutputWithContext(ctx context.Context) GetMysqlFlavorsResultOutput {
	return o
}

func (o GetMysqlFlavorsResultOutput) AvailabilityZoneMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) *string { return v.AvailabilityZoneMode }).(pulumi.StringPtrOutput)
}

func (o GetMysqlFlavorsResultOutput) Engine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) *string { return v.Engine }).(pulumi.StringPtrOutput)
}

// Indicates the flavors information. Structure is documented below.
func (o GetMysqlFlavorsResultOutput) Flavors() GetMysqlFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) []GetMysqlFlavorsFlavor { return v.Flavors }).(GetMysqlFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMysqlFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMysqlFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the database version.
func (o GetMysqlFlavorsResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMysqlFlavorsResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMysqlFlavorsResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicatedelb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the available ELB Flavors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/DedicatedElb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			flavors, err := DedicatedElb.GetFlavors(ctx, &dedicatedelb.GetFlavorsArgs{
//				Type:           pulumi.StringRef("L7"),
//				MaxConnections: pulumi.IntRef(200000),
//				Cps:            pulumi.IntRef(2000),
//				Bandwidth:      pulumi.IntRef(50),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = DedicatedElb.NewLoadbalancer(ctx, "lb", &DedicatedElb.LoadbalancerArgs{
//				L7FlavorId: pulumi.String(flavors.Ids[0]),
//			})
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
	err := ctx.Invoke("huaweicloud:DedicatedElb/getFlavors:getFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlavors.
type GetFlavorsArgs struct {
	// Specifies the bandwidth size(Mbit/s) in the flavor.
	Bandwidth *int `pulumi:"bandwidth"`
	// Specifies the cps in the flavor.
	Cps *int `pulumi:"cps"`
	// Specifies the maximum connections in the flavor.
	MaxConnections *int `pulumi:"maxConnections"`
	// Specifies the qps in the L7 flavor.
	Qps *int `pulumi:"qps"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
	// Specifies the flavor type. Valid values are L4 and L7.
	Type *string `pulumi:"type"`
}

// A collection of values returned by getFlavors.
type GetFlavorsResult struct {
	// Bandwidth size(Mbit/s) of the flavor.
	Bandwidth *int `pulumi:"bandwidth"`
	// Cps of the flavor.
	Cps *int `pulumi:"cps"`
	// A list of flavors. Each element contains the following attributes:
	Flavors []GetFlavorsFlavor `pulumi:"flavors"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of flavor IDs.
	Ids []string `pulumi:"ids"`
	// Maximum connections of the flavor.
	MaxConnections *int `pulumi:"maxConnections"`
	// Qps of the L7 flavor.
	Qps    *int   `pulumi:"qps"`
	Region string `pulumi:"region"`
	// Type of the flavor.
	Type *string `pulumi:"type"`
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
	// Specifies the bandwidth size(Mbit/s) in the flavor.
	Bandwidth pulumi.IntPtrInput `pulumi:"bandwidth"`
	// Specifies the cps in the flavor.
	Cps pulumi.IntPtrInput `pulumi:"cps"`
	// Specifies the maximum connections in the flavor.
	MaxConnections pulumi.IntPtrInput `pulumi:"maxConnections"`
	// Specifies the qps in the L7 flavor.
	Qps pulumi.IntPtrInput `pulumi:"qps"`
	// The region in which to obtain the flavors. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the flavor type. Valid values are L4 and L7.
	Type pulumi.StringPtrInput `pulumi:"type"`
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

// Bandwidth size(Mbit/s) of the flavor.
func (o GetFlavorsResultOutput) Bandwidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.Bandwidth }).(pulumi.IntPtrOutput)
}

// Cps of the flavor.
func (o GetFlavorsResultOutput) Cps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.Cps }).(pulumi.IntPtrOutput)
}

// A list of flavors. Each element contains the following attributes:
func (o GetFlavorsResultOutput) Flavors() GetFlavorsFlavorArrayOutput {
	return o.ApplyT(func(v GetFlavorsResult) []GetFlavorsFlavor { return v.Flavors }).(GetFlavorsFlavorArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of flavor IDs.
func (o GetFlavorsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlavorsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Maximum connections of the flavor.
func (o GetFlavorsResultOutput) MaxConnections() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.MaxConnections }).(pulumi.IntPtrOutput)
}

// Qps of the L7 flavor.
func (o GetFlavorsResultOutput) Qps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.Qps }).(pulumi.IntPtrOutput)
}

func (o GetFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

// Type of the flavor.
func (o GetFlavorsResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlavorsResultOutput{})
}

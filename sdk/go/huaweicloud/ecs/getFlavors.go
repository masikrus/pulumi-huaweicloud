// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the available Compute Flavors.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Ecs"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			flavors, err := Ecs.GetFlavors(ctx, &ecs.GetFlavorsArgs{
//				AvailabilityZone: pulumi.StringRef("cn-north-1a"),
//				PerformanceType:  pulumi.StringRef("normal"),
//				CpuCoreCount:     pulumi.IntRef(2),
//				MemorySize:       pulumi.IntRef(4),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Ecs.NewInstance(ctx, "instance", &Ecs.InstanceArgs{
//				FlavorId: pulumi.String(flavors.Ids[0]),
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
	err := ctx.Invoke("huaweicloud:Ecs/getFlavors:getFlavors", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFlavors.
type GetFlavorsArgs struct {
	// Specifies the AZ name.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the number of vCPUs in the ECS flavor.
	CpuCoreCount *int `pulumi:"cpuCoreCount"`
	// Specifies the generation of an ECS type. For example, **s3** indicates
	// the general-purpose third-generation ECSs. For details, see
	// [ECS Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-ecs/ecs_01_0014.html).
	Generation *string `pulumi:"generation"`
	// Specifies the memory size(GB) in the ECS flavor.
	MemorySize *int `pulumi:"memorySize"`
	// Specifies the ECS flavor type. Possible values are as follows:
	// + **normal**: General computing
	// + **computingv3**: General computing-plus
	// + **highmem**: Memory-optimized
	// + **saphana**: Large-memory HANA ECS
	// + **diskintensive**: Disk-intensive
	PerformanceType *string `pulumi:"performanceType"`
	// The region in which to obtain the flavors.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getFlavors.
type GetFlavorsResult struct {
	AvailabilityZone *string `pulumi:"availabilityZone"`
	CpuCoreCount     *int    `pulumi:"cpuCoreCount"`
	Generation       *string `pulumi:"generation"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of flavor IDs.
	Ids             []string `pulumi:"ids"`
	MemorySize      *int     `pulumi:"memorySize"`
	PerformanceType *string  `pulumi:"performanceType"`
	Region          string   `pulumi:"region"`
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
	// Specifies the AZ name.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Specifies the number of vCPUs in the ECS flavor.
	CpuCoreCount pulumi.IntPtrInput `pulumi:"cpuCoreCount"`
	// Specifies the generation of an ECS type. For example, **s3** indicates
	// the general-purpose third-generation ECSs. For details, see
	// [ECS Specifications](https://support.huaweicloud.com/intl/en-us/productdesc-ecs/ecs_01_0014.html).
	Generation pulumi.StringPtrInput `pulumi:"generation"`
	// Specifies the memory size(GB) in the ECS flavor.
	MemorySize pulumi.IntPtrInput `pulumi:"memorySize"`
	// Specifies the ECS flavor type. Possible values are as follows:
	// + **normal**: General computing
	// + **computingv3**: General computing-plus
	// + **highmem**: Memory-optimized
	// + **saphana**: Large-memory HANA ECS
	// + **diskintensive**: Disk-intensive
	PerformanceType pulumi.StringPtrInput `pulumi:"performanceType"`
	// The region in which to obtain the flavors.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
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

func (o GetFlavorsResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) CpuCoreCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.CpuCoreCount }).(pulumi.IntPtrOutput)
}

func (o GetFlavorsResultOutput) Generation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.Generation }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFlavorsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Id }).(pulumi.StringOutput)
}

// A list of flavor IDs.
func (o GetFlavorsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFlavorsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o GetFlavorsResultOutput) MemorySize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *int { return v.MemorySize }).(pulumi.IntPtrOutput)
}

func (o GetFlavorsResultOutput) PerformanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFlavorsResult) *string { return v.PerformanceType }).(pulumi.StringPtrOutput)
}

func (o GetFlavorsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetFlavorsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFlavorsResultOutput{})
}

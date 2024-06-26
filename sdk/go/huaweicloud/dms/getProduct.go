// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the ID of an available HuaweiCloud DMS product.
//
// ## Example Usage
// ### Filter DMS kafka product list by I/O specification
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dms"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dms.GetProduct(ctx, &dms.GetProductArgs{
//				Engine:          "kafka",
//				InstanceType:    "cluster",
//				PartitionNum:    pulumi.StringRef("300"),
//				Storage:         pulumi.StringRef("600"),
//				StorageSpecCode: pulumi.StringRef("dms.physical.storage.high"),
//				Version:         pulumi.StringRef("1.1.0"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Filter DMS kafka product list by underlying VM specification
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dms"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dms.GetProduct(ctx, &dms.GetProductArgs{
//				Engine:          "kafka",
//				InstanceType:    "cluster",
//				Version:         pulumi.StringRef("2.3.0"),
//				VmSpecification: pulumi.StringRef("c6.large.2"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProduct(ctx *pulumi.Context, args *GetProductArgs, opts ...pulumi.InvokeOption) (*GetProductResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetProductResult
	err := ctx.Invoke("huaweicloud:Dms/getProduct:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProduct.
type GetProductArgs struct {
	// Indicates the list of availability zones with available resources.
	AvailabilityZones []string `pulumi:"availabilityZones"`
	// Indicates the baseline bandwidth of a DMS instance.
	// The valid values are **100MB**, **300MB**, **600MB** and **1200MB**.
	Bandwidth *string `pulumi:"bandwidth"`
	// Indicates the name of a message engine. The valid values are **kafka**, **rabbitmq**.
	Engine string `pulumi:"engine"`
	// Indicates an instance type. The valid values are **single** and **cluster**.
	InstanceType string `pulumi:"instanceType"`
	// Deprecated: io_type has deprecated, please use storage_spec_code
	IoType *string `pulumi:"ioType"`
	// Indicates the number of nodes in a cluster.
	NodeNum *string `pulumi:"nodeNum"`
	// Indicates the maximum number of topics that can be created for a Kafka instance.
	// The valid values are **300**, **900** and **1800**.
	PartitionNum *string `pulumi:"partitionNum"`
	// The region in which to obtain the dms products. If omitted, the provider-level region
	// will be used.
	Region *string `pulumi:"region"`
	// Indicates the storage capacity of the resource.
	// The default value is the storage capacity of the product.
	Storage *string `pulumi:"storage"`
	// Indicates an I/O specification.
	// The valid values are **dms.physical.storage.high** and **dms.physical.storage.ultra**.
	StorageSpecCode *string `pulumi:"storageSpecCode"`
	// Indicates the version of a message engine.
	Version *string `pulumi:"version"`
	// Indicates underlying VM specification, such as **c6.large.2**.
	VmSpecification *string `pulumi:"vmSpecification"`
}

// A collection of values returned by getProduct.
type GetProductResult struct {
	AvailabilityZones []string `pulumi:"availabilityZones"`
	Bandwidth         string   `pulumi:"bandwidth"`
	Engine            string   `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	InstanceType string `pulumi:"instanceType"`
	// Deprecated: io_type has deprecated, please use storage_spec_code
	IoType          string `pulumi:"ioType"`
	NodeNum         string `pulumi:"nodeNum"`
	PartitionNum    string `pulumi:"partitionNum"`
	Region          string `pulumi:"region"`
	Storage         string `pulumi:"storage"`
	StorageSpecCode string `pulumi:"storageSpecCode"`
	// The available I/O specifications.
	StorageSpecCodes []string `pulumi:"storageSpecCodes"`
	Version          string   `pulumi:"version"`
	VmSpecification  string   `pulumi:"vmSpecification"`
}

func GetProductOutput(ctx *pulumi.Context, args GetProductOutputArgs, opts ...pulumi.InvokeOption) GetProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetProductResult, error) {
			args := v.(GetProductArgs)
			r, err := GetProduct(ctx, &args, opts...)
			var s GetProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetProductResultOutput)
}

// A collection of arguments for invoking getProduct.
type GetProductOutputArgs struct {
	// Indicates the list of availability zones with available resources.
	AvailabilityZones pulumi.StringArrayInput `pulumi:"availabilityZones"`
	// Indicates the baseline bandwidth of a DMS instance.
	// The valid values are **100MB**, **300MB**, **600MB** and **1200MB**.
	Bandwidth pulumi.StringPtrInput `pulumi:"bandwidth"`
	// Indicates the name of a message engine. The valid values are **kafka**, **rabbitmq**.
	Engine pulumi.StringInput `pulumi:"engine"`
	// Indicates an instance type. The valid values are **single** and **cluster**.
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
	// Deprecated: io_type has deprecated, please use storage_spec_code
	IoType pulumi.StringPtrInput `pulumi:"ioType"`
	// Indicates the number of nodes in a cluster.
	NodeNum pulumi.StringPtrInput `pulumi:"nodeNum"`
	// Indicates the maximum number of topics that can be created for a Kafka instance.
	// The valid values are **300**, **900** and **1800**.
	PartitionNum pulumi.StringPtrInput `pulumi:"partitionNum"`
	// The region in which to obtain the dms products. If omitted, the provider-level region
	// will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Indicates the storage capacity of the resource.
	// The default value is the storage capacity of the product.
	Storage pulumi.StringPtrInput `pulumi:"storage"`
	// Indicates an I/O specification.
	// The valid values are **dms.physical.storage.high** and **dms.physical.storage.ultra**.
	StorageSpecCode pulumi.StringPtrInput `pulumi:"storageSpecCode"`
	// Indicates the version of a message engine.
	Version pulumi.StringPtrInput `pulumi:"version"`
	// Indicates underlying VM specification, such as **c6.large.2**.
	VmSpecification pulumi.StringPtrInput `pulumi:"vmSpecification"`
}

func (GetProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductArgs)(nil)).Elem()
}

// A collection of values returned by getProduct.
type GetProductResultOutput struct{ *pulumi.OutputState }

func (GetProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProductResult)(nil)).Elem()
}

func (o GetProductResultOutput) ToGetProductResultOutput() GetProductResultOutput {
	return o
}

func (o GetProductResultOutput) ToGetProductResultOutputWithContext(ctx context.Context) GetProductResultOutput {
	return o
}

func (o GetProductResultOutput) AvailabilityZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductResult) []string { return v.AvailabilityZones }).(pulumi.StringArrayOutput)
}

func (o GetProductResultOutput) Bandwidth() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Bandwidth }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.InstanceType }).(pulumi.StringOutput)
}

// Deprecated: io_type has deprecated, please use storage_spec_code
func (o GetProductResultOutput) IoType() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.IoType }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) NodeNum() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.NodeNum }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) PartitionNum() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.PartitionNum }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) Storage() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Storage }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) StorageSpecCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.StorageSpecCode }).(pulumi.StringOutput)
}

// The available I/O specifications.
func (o GetProductResultOutput) StorageSpecCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetProductResult) []string { return v.StorageSpecCodes }).(pulumi.StringArrayOutput)
}

func (o GetProductResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o GetProductResultOutput) VmSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v GetProductResult) string { return v.VmSpecification }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProductResultOutput{})
}

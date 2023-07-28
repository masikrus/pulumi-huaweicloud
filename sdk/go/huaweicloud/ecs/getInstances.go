// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of the compute instances.
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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			nameRegex := cfg.RequireObject("nameRegex")
//			_, err := Ecs.GetInstances(ctx, &ecs.GetInstancesArgs{
//				Name: pulumi.StringRef(nameRegex),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetInstances(ctx *pulumi.Context, args *GetInstancesArgs, opts ...pulumi.InvokeOption) (*GetInstancesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetInstancesResult
	err := ctx.Invoke("huaweicloud:Ecs/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Specifies the availability zone where the instance is located.
	// Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?ECS) for this argument.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the flavor ID.
	FlavorId *string `pulumi:"flavorId"`
	// Specifies the flavor name of the instance.
	FlavorName *string `pulumi:"flavorName"`
	// Specifies the image ID of the instance.
	ImageId *string `pulumi:"imageId"`
	// Specifies the ECS ID.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the key pair that is used to authenticate the instance.
	KeyPair *string `pulumi:"keyPair"`
	// Specifies the instance name, which can be queried with a regular expression.
	// The instance name supports fuzzy matching query too.
	Name *string `pulumi:"name"`
	// Specifies the region in which to obtain the instances.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// Specifies the status of the instance. The valid values are as follows:
	// + **ACTIVE**: The instance is running properly.
	// + **SHUTOFF**: The instance has been properly stopped.
	// + **ERROR**: An error has occurred on the instance.
	Status *string `pulumi:"status"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	// The availability zone where the instance is located.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The enterprise project ID.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The flavor ID.
	FlavorId *string `pulumi:"flavorId"`
	// The flavor name of the instance.
	FlavorName *string `pulumi:"flavorName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The image ID of the instance.
	ImageId    *string `pulumi:"imageId"`
	InstanceId *string `pulumi:"instanceId"`
	// List of ECS instance details. The object structure of each ECS instance is documented below.
	Instances []GetInstancesInstance `pulumi:"instances"`
	// The key pair that is used to authenticate the instance.
	KeyPair *string `pulumi:"keyPair"`
	// The instance name.
	Name   *string `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The instance status.
	Status *string `pulumi:"status"`
}

func GetInstancesOutput(ctx *pulumi.Context, args GetInstancesOutputArgs, opts ...pulumi.InvokeOption) GetInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInstancesResult, error) {
			args := v.(GetInstancesArgs)
			r, err := GetInstances(ctx, &args, opts...)
			var s GetInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInstancesResultOutput)
}

// A collection of arguments for invoking getInstances.
type GetInstancesOutputArgs struct {
	// Specifies the availability zone where the instance is located.
	// Please following [reference](https://developer.huaweicloud.com/intl/en-us/endpoint?ECS) for this argument.
	AvailabilityZone pulumi.StringPtrInput `pulumi:"availabilityZone"`
	// Specifies the enterprise project ID.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the flavor ID.
	FlavorId pulumi.StringPtrInput `pulumi:"flavorId"`
	// Specifies the flavor name of the instance.
	FlavorName pulumi.StringPtrInput `pulumi:"flavorName"`
	// Specifies the image ID of the instance.
	ImageId pulumi.StringPtrInput `pulumi:"imageId"`
	// Specifies the ECS ID.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// Specifies the key pair that is used to authenticate the instance.
	KeyPair pulumi.StringPtrInput `pulumi:"keyPair"`
	// Specifies the instance name, which can be queried with a regular expression.
	// The instance name supports fuzzy matching query too.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the region in which to obtain the instances.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the status of the instance. The valid values are as follows:
	// + **ACTIVE**: The instance is running properly.
	// + **SHUTOFF**: The instance has been properly stopped.
	// + **ERROR**: An error has occurred on the instance.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesArgs)(nil)).Elem()
}

// A collection of values returned by getInstances.
type GetInstancesResultOutput struct{ *pulumi.OutputState }

func (GetInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstancesResult)(nil)).Elem()
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutput() GetInstancesResultOutput {
	return o
}

func (o GetInstancesResultOutput) ToGetInstancesResultOutputWithContext(ctx context.Context) GetInstancesResultOutput {
	return o
}

// The availability zone where the instance is located.
func (o GetInstancesResultOutput) AvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.AvailabilityZone }).(pulumi.StringPtrOutput)
}

// The enterprise project ID.
func (o GetInstancesResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The flavor ID.
func (o GetInstancesResultOutput) FlavorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.FlavorId }).(pulumi.StringPtrOutput)
}

// The flavor name of the instance.
func (o GetInstancesResultOutput) FlavorName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.FlavorName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The image ID of the instance.
func (o GetInstancesResultOutput) ImageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.ImageId }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

// List of ECS instance details. The object structure of each ECS instance is documented below.
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

// The key pair that is used to authenticate the instance.
func (o GetInstancesResultOutput) KeyPair() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.KeyPair }).(pulumi.StringPtrOutput)
}

// The instance name.
func (o GetInstancesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetInstancesResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// The instance status.
func (o GetInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}

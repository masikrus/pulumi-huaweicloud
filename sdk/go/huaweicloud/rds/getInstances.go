// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to list all available RDS instances.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Rds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Rds.GetInstances(ctx, &rds.GetInstancesArgs{
//				Name: pulumi.StringRef("rds-instance"),
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
	err := ctx.Invoke("huaweicloud:Rds/getInstances:getInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstances.
type GetInstancesArgs struct {
	// Specifies the type of the database. Valid values are: MySQL, PostgreSQL, and SQLServer.
	DatastoreType *string `pulumi:"datastoreType"`
	// Specifies the enterprise project id.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the name of the instance.
	Name *string `pulumi:"name"`
	// The region in which to obtain the instances. If omitted, the provider-level region will
	// be used.
	Region *string `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Specifies the type of the instance. Valid values are: Single, Ha, Replica, and Enterprise.
	Type *string `pulumi:"type"`
	// Specifies the VPC ID.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by getInstances.
type GetInstancesResult struct {
	DatastoreType *string `pulumi:"datastoreType"`
	// Indicates the enterprise project id.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An array of available instances.
	Instances []GetInstancesInstance `pulumi:"instances"`
	// Indicates the node name.
	Name *string `pulumi:"name"`
	// The region of the instance.
	Region string `pulumi:"region"`
	// Indicates the network ID of a subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Indicates the volume type.
	Type *string `pulumi:"type"`
	// Indicates the VPC ID.
	VpcId *string `pulumi:"vpcId"`
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
	// Specifies the type of the database. Valid values are: MySQL, PostgreSQL, and SQLServer.
	DatastoreType pulumi.StringPtrInput `pulumi:"datastoreType"`
	// Specifies the enterprise project id.
	EnterpriseProjectId pulumi.StringPtrInput `pulumi:"enterpriseProjectId"`
	// Specifies the name of the instance.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to obtain the instances. If omitted, the provider-level region will
	// be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// Specifies the network ID of a subnet.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Specifies the type of the instance. Valid values are: Single, Ha, Replica, and Enterprise.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Specifies the VPC ID.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
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

func (o GetInstancesResultOutput) DatastoreType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.DatastoreType }).(pulumi.StringPtrOutput)
}

// Indicates the enterprise project id.
func (o GetInstancesResultOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

// An array of available instances.
func (o GetInstancesResultOutput) Instances() GetInstancesInstanceArrayOutput {
	return o.ApplyT(func(v GetInstancesResult) []GetInstancesInstance { return v.Instances }).(GetInstancesInstanceArrayOutput)
}

// Indicates the node name.
func (o GetInstancesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The region of the instance.
func (o GetInstancesResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstancesResult) string { return v.Region }).(pulumi.StringOutput)
}

// Indicates the network ID of a subnet.
func (o GetInstancesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Indicates the volume type.
func (o GetInstancesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Indicates the VPC ID.
func (o GetInstancesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInstancesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstancesResultOutput{})
}

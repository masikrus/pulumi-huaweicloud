// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides details about a specific VPC route table.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			_, err := Vpc.GetRouteTable(ctx, &vpc.GetRouteTableArgs{
//				VpcId: vpcId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Vpc.GetRouteTable(ctx, &vpc.GetRouteTableArgs{
//				VpcId: vpcId,
//				Name:  pulumi.StringRef("demo"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRouteTable(ctx *pulumi.Context, args *LookupRouteTableArgs, opts ...pulumi.InvokeOption) (*LookupRouteTableResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupRouteTableResult
	err := ctx.Invoke("huaweicloud:Vpc/getRouteTable:getRouteTable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableArgs struct {
	// - Specifies the ID of the route table.
	Id *string `pulumi:"id"`
	// - Specifies the name of the route table.
	Name *string `pulumi:"name"`
	// The region in which to query the vpc route table.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
	// - Specifies the VPC ID where the route table resides.
	VpcId string `pulumi:"vpcId"`
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResult struct {
	Default     bool                 `pulumi:"default"`
	Description string               `pulumi:"description"`
	Id          string               `pulumi:"id"`
	Name        string               `pulumi:"name"`
	Region      string               `pulumi:"region"`
	Routes      []GetRouteTableRoute `pulumi:"routes"`
	Subnets     []string             `pulumi:"subnets"`
	VpcId       string               `pulumi:"vpcId"`
}

func LookupRouteTableOutput(ctx *pulumi.Context, args LookupRouteTableOutputArgs, opts ...pulumi.InvokeOption) LookupRouteTableResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteTableResult, error) {
			args := v.(LookupRouteTableArgs)
			r, err := LookupRouteTable(ctx, &args, opts...)
			var s LookupRouteTableResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteTableResultOutput)
}

// A collection of arguments for invoking getRouteTable.
type LookupRouteTableOutputArgs struct {
	// - Specifies the ID of the route table.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// - Specifies the name of the route table.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The region in which to query the vpc route table.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// - Specifies the VPC ID where the route table resides.
	VpcId pulumi.StringInput `pulumi:"vpcId"`
}

func (LookupRouteTableOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableArgs)(nil)).Elem()
}

// A collection of values returned by getRouteTable.
type LookupRouteTableResultOutput struct{ *pulumi.OutputState }

func (LookupRouteTableResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteTableResult)(nil)).Elem()
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutput() LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) ToLookupRouteTableResultOutputWithContext(ctx context.Context) LookupRouteTableResultOutput {
	return o
}

func (o LookupRouteTableResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRouteTableResult) bool { return v.Default }).(pulumi.BoolOutput)
}

func (o LookupRouteTableResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupRouteTableResultOutput) Routes() GetRouteTableRouteArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []GetRouteTableRoute { return v.Routes }).(GetRouteTableRouteArrayOutput)
}

func (o LookupRouteTableResultOutput) Subnets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRouteTableResult) []string { return v.Subnets }).(pulumi.StringArrayOutput)
}

func (o LookupRouteTableResultOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteTableResult) string { return v.VpcId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteTableResultOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of ELB pools.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			poolName := cfg.RequireObject("poolName")
//			_, err := Elb.GetPools(ctx, &elb.GetPoolsArgs{
//				Name: pulumi.StringRef(poolName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetPools(ctx *pulumi.Context, args *GetPoolsArgs, opts ...pulumi.InvokeOption) (*GetPoolsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetPoolsResult
	err := ctx.Invoke("huaweicloud:Elb/getPools:getPools", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPools.
type GetPoolsArgs struct {
	// Specifies the description of the ELB pool.
	Description *string `pulumi:"description"`
	// Specifies the health monitor ID of the ELB pool.
	HealthmonitorId *string `pulumi:"healthmonitorId"`
	// Specifies the method of the ELB pool. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
	// or SOURCE_IP.
	LbMethod *string `pulumi:"lbMethod"`
	// Specifies the loadbalancer ID of the ELB pool.
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// Specifies the name of the ELB pool.
	Name *string `pulumi:"name"`
	// Specifies the ID of the ELB pool.
	PoolId *string `pulumi:"poolId"`
	// Specifies the protocol of the ELB pool. This can either be TCP, UDP or HTTP.
	Protocol *string `pulumi:"protocol"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getPools.
type GetPoolsResult struct {
	// The description of pool.
	Description *string `pulumi:"description"`
	// The health monitor ID of the LB pool.
	HealthmonitorId *string `pulumi:"healthmonitorId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The load balancing algorithm to distribute traffic to the pool's members.
	LbMethod       *string `pulumi:"lbMethod"`
	LoadbalancerId *string `pulumi:"loadbalancerId"`
	// The pool name.
	Name   *string `pulumi:"name"`
	PoolId *string `pulumi:"poolId"`
	// The pool list. For details, see data structure of the pool field.
	// The object structure is documented below.
	Pools []GetPoolsPool `pulumi:"pools"`
	// The protocol of pool.
	Protocol *string `pulumi:"protocol"`
	Region   string  `pulumi:"region"`
}

func GetPoolsOutput(ctx *pulumi.Context, args GetPoolsOutputArgs, opts ...pulumi.InvokeOption) GetPoolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPoolsResult, error) {
			args := v.(GetPoolsArgs)
			r, err := GetPools(ctx, &args, opts...)
			var s GetPoolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPoolsResultOutput)
}

// A collection of arguments for invoking getPools.
type GetPoolsOutputArgs struct {
	// Specifies the description of the ELB pool.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Specifies the health monitor ID of the ELB pool.
	HealthmonitorId pulumi.StringPtrInput `pulumi:"healthmonitorId"`
	// Specifies the method of the ELB pool. Must be one of ROUND_ROBIN, LEAST_CONNECTIONS,
	// or SOURCE_IP.
	LbMethod pulumi.StringPtrInput `pulumi:"lbMethod"`
	// Specifies the loadbalancer ID of the ELB pool.
	LoadbalancerId pulumi.StringPtrInput `pulumi:"loadbalancerId"`
	// Specifies the name of the ELB pool.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the ID of the ELB pool.
	PoolId pulumi.StringPtrInput `pulumi:"poolId"`
	// Specifies the protocol of the ELB pool. This can either be TCP, UDP or HTTP.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
	// Specifies the region in which to query the data source.
	// If omitted, the provider-level region will be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetPoolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoolsArgs)(nil)).Elem()
}

// A collection of values returned by getPools.
type GetPoolsResultOutput struct{ *pulumi.OutputState }

func (GetPoolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPoolsResult)(nil)).Elem()
}

func (o GetPoolsResultOutput) ToGetPoolsResultOutput() GetPoolsResultOutput {
	return o
}

func (o GetPoolsResultOutput) ToGetPoolsResultOutputWithContext(ctx context.Context) GetPoolsResultOutput {
	return o
}

// The description of pool.
func (o GetPoolsResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The health monitor ID of the LB pool.
func (o GetPoolsResultOutput) HealthmonitorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.HealthmonitorId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetPoolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoolsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The load balancing algorithm to distribute traffic to the pool's members.
func (o GetPoolsResultOutput) LbMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.LbMethod }).(pulumi.StringPtrOutput)
}

func (o GetPoolsResultOutput) LoadbalancerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.LoadbalancerId }).(pulumi.StringPtrOutput)
}

// The pool name.
func (o GetPoolsResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetPoolsResultOutput) PoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.PoolId }).(pulumi.StringPtrOutput)
}

// The pool list. For details, see data structure of the pool field.
// The object structure is documented below.
func (o GetPoolsResultOutput) Pools() GetPoolsPoolArrayOutput {
	return o.ApplyT(func(v GetPoolsResult) []GetPoolsPool { return v.Pools }).(GetPoolsPoolArrayOutput)
}

// The protocol of pool.
func (o GetPoolsResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetPoolsResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o GetPoolsResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetPoolsResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPoolsResultOutput{})
}

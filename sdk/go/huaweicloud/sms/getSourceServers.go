// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sms

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of SMS source servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Sms"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Sms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			serverName := cfg.RequireObject("serverName")
//			_, err := Sms.GetSourceServers(ctx, &sms.GetSourceServersArgs{
//				Name: pulumi.StringRef(serverName),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSourceServers(ctx *pulumi.Context, args *GetSourceServersArgs, opts ...pulumi.InvokeOption) (*GetSourceServersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetSourceServersResult
	err := ctx.Invoke("huaweicloud:Sms/getSourceServers:getSourceServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSourceServers.
type GetSourceServersArgs struct {
	// Specifies the ID of the source server.
	Id *string `pulumi:"id"`
	// Specifies the IP address of the source server.
	Ip *string `pulumi:"ip"`
	// Specifies the name of the source server.
	Name *string `pulumi:"name"`
	// Specifies the status of the source server.
	State *string `pulumi:"state"`
}

// A collection of values returned by getSourceServers.
type GetSourceServersResult struct {
	// The ID of the source server.
	Id string `pulumi:"id"`
	// The IP address of the source server.
	Ip *string `pulumi:"ip"`
	// The disk name, for example, /dev/vda.
	Name *string `pulumi:"name"`
	// An array of SMS source servers found. Structure is documented below.
	Servers []GetSourceServersServer `pulumi:"servers"`
	// The status of the source server.
	State *string `pulumi:"state"`
}

func GetSourceServersOutput(ctx *pulumi.Context, args GetSourceServersOutputArgs, opts ...pulumi.InvokeOption) GetSourceServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSourceServersResult, error) {
			args := v.(GetSourceServersArgs)
			r, err := GetSourceServers(ctx, &args, opts...)
			var s GetSourceServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSourceServersResultOutput)
}

// A collection of arguments for invoking getSourceServers.
type GetSourceServersOutputArgs struct {
	// Specifies the ID of the source server.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Specifies the IP address of the source server.
	Ip pulumi.StringPtrInput `pulumi:"ip"`
	// Specifies the name of the source server.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Specifies the status of the source server.
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (GetSourceServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourceServersArgs)(nil)).Elem()
}

// A collection of values returned by getSourceServers.
type GetSourceServersResultOutput struct{ *pulumi.OutputState }

func (GetSourceServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSourceServersResult)(nil)).Elem()
}

func (o GetSourceServersResultOutput) ToGetSourceServersResultOutput() GetSourceServersResultOutput {
	return o
}

func (o GetSourceServersResultOutput) ToGetSourceServersResultOutputWithContext(ctx context.Context) GetSourceServersResultOutput {
	return o
}

// The ID of the source server.
func (o GetSourceServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSourceServersResult) string { return v.Id }).(pulumi.StringOutput)
}

// The IP address of the source server.
func (o GetSourceServersResultOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSourceServersResult) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

// The disk name, for example, /dev/vda.
func (o GetSourceServersResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSourceServersResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// An array of SMS source servers found. Structure is documented below.
func (o GetSourceServersResultOutput) Servers() GetSourceServersServerArrayOutput {
	return o.ApplyT(func(v GetSourceServersResult) []GetSourceServersServer { return v.Servers }).(GetSourceServersServerArrayOutput)
}

// The status of the source server.
func (o GetSourceServersResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSourceServersResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSourceServersResultOutput{})
}

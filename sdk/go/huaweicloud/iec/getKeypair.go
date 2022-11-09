// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a specific IEC keypair.
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
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Iec.GetKeypair(ctx, &iec.GetKeypairArgs{
//				Name: "iec-keypair-demo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupKeypair(ctx *pulumi.Context, args *LookupKeypairArgs, opts ...pulumi.InvokeOption) (*LookupKeypairResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupKeypairResult
	err := ctx.Invoke("huaweicloud:Iec/getKeypair:getKeypair", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKeypair.
type LookupKeypairArgs struct {
	// Specifies a unique name for the keypair. This parameter can contain a maximum of 64
	// characters, which may consist of letters, digits, underscores (_), and hyphens (-).
	Name string `pulumi:"name"`
}

// A collection of values returned by getKeypair.
type LookupKeypairResult struct {
	// The finger of iec keypair. The value contains a encoding type(SHA256) and a string of 43 characters.
	Fingerprint string `pulumi:"fingerprint"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// The pregenerated OpenSSH-formatted public key.
	PublicKey string `pulumi:"publicKey"`
}

func LookupKeypairOutput(ctx *pulumi.Context, args LookupKeypairOutputArgs, opts ...pulumi.InvokeOption) LookupKeypairResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKeypairResult, error) {
			args := v.(LookupKeypairArgs)
			r, err := LookupKeypair(ctx, &args, opts...)
			var s LookupKeypairResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKeypairResultOutput)
}

// A collection of arguments for invoking getKeypair.
type LookupKeypairOutputArgs struct {
	// Specifies a unique name for the keypair. This parameter can contain a maximum of 64
	// characters, which may consist of letters, digits, underscores (_), and hyphens (-).
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupKeypairOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeypairArgs)(nil)).Elem()
}

// A collection of values returned by getKeypair.
type LookupKeypairResultOutput struct{ *pulumi.OutputState }

func (LookupKeypairResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKeypairResult)(nil)).Elem()
}

func (o LookupKeypairResultOutput) ToLookupKeypairResultOutput() LookupKeypairResultOutput {
	return o
}

func (o LookupKeypairResultOutput) ToLookupKeypairResultOutputWithContext(ctx context.Context) LookupKeypairResultOutput {
	return o
}

// The finger of iec keypair. The value contains a encoding type(SHA256) and a string of 43 characters.
func (o LookupKeypairResultOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Fingerprint }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKeypairResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKeypairResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.Name }).(pulumi.StringOutput)
}

// The pregenerated OpenSSH-formatted public key.
func (o LookupKeypairResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKeypairResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKeypairResultOutput{})
}

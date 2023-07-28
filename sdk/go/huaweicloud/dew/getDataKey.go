// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dew

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the plaintext and the ciphertext of an available HuaweiCloud KMS DEK (data encryption key).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dew"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dew"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			key1, err := Dew.NewKey(ctx, "key1", &Dew.KeyArgs{
//				KeyAlias:       pulumi.String("key_1"),
//				PendingDays:    pulumi.String("7"),
//				KeyDescription: pulumi.String("first test key"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = Dew.GetDataKeyOutput(ctx, dew.GetDataKeyOutputArgs{
//				KeyId:         key1.ID(),
//				DatakeyLength: pulumi.String("512"),
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func GetDataKey(ctx *pulumi.Context, args *GetDataKeyArgs, opts ...pulumi.InvokeOption) (*GetDataKeyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDataKeyResult
	err := ctx.Invoke("huaweicloud:Dew/getDataKey:getDataKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataKey.
type GetDataKeyArgs struct {
	// Number of bits in the length of a DEK (data encryption keys). The maximum number
	// is 512. Changing this gets the new data encryption key.
	DatakeyLength string `pulumi:"datakeyLength"`
	// The value of this parameter must be a series of
	// "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive
	// information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext *string `pulumi:"encryptionContext"`
	// The globally unique identifier for the key. Changing this gets the new data encryption
	// key.
	KeyId string `pulumi:"keyId"`
	// The region in which to obtain the keys. If omitted, the provider-level region will be
	// used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getDataKey.
type GetDataKeyResult struct {
	// The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	CipherText        string  `pulumi:"cipherText"`
	DatakeyLength     string  `pulumi:"datakeyLength"`
	EncryptionContext *string `pulumi:"encryptionContext"`
	// The provider-assigned unique ID for this managed resource.
	Id    string `pulumi:"id"`
	KeyId string `pulumi:"keyId"`
	// The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
	PlainText string `pulumi:"plainText"`
	Region    string `pulumi:"region"`
}

func GetDataKeyOutput(ctx *pulumi.Context, args GetDataKeyOutputArgs, opts ...pulumi.InvokeOption) GetDataKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDataKeyResult, error) {
			args := v.(GetDataKeyArgs)
			r, err := GetDataKey(ctx, &args, opts...)
			var s GetDataKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDataKeyResultOutput)
}

// A collection of arguments for invoking getDataKey.
type GetDataKeyOutputArgs struct {
	// Number of bits in the length of a DEK (data encryption keys). The maximum number
	// is 512. Changing this gets the new data encryption key.
	DatakeyLength pulumi.StringInput `pulumi:"datakeyLength"`
	// The value of this parameter must be a series of
	// "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive
	// information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}
	EncryptionContext pulumi.StringPtrInput `pulumi:"encryptionContext"`
	// The globally unique identifier for the key. Changing this gets the new data encryption
	// key.
	KeyId pulumi.StringInput `pulumi:"keyId"`
	// The region in which to obtain the keys. If omitted, the provider-level region will be
	// used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (GetDataKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataKeyArgs)(nil)).Elem()
}

// A collection of values returned by getDataKey.
type GetDataKeyResultOutput struct{ *pulumi.OutputState }

func (GetDataKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDataKeyResult)(nil)).Elem()
}

func (o GetDataKeyResultOutput) ToGetDataKeyResultOutput() GetDataKeyResultOutput {
	return o
}

func (o GetDataKeyResultOutput) ToGetDataKeyResultOutputWithContext(ctx context.Context) GetDataKeyResultOutput {
	return o
}

// The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
func (o GetDataKeyResultOutput) CipherText() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.CipherText }).(pulumi.StringOutput)
}

func (o GetDataKeyResultOutput) DatakeyLength() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.DatakeyLength }).(pulumi.StringOutput)
}

func (o GetDataKeyResultOutput) EncryptionContext() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDataKeyResult) *string { return v.EncryptionContext }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDataKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDataKeyResultOutput) KeyId() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.KeyId }).(pulumi.StringOutput)
}

// The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.
func (o GetDataKeyResultOutput) PlainText() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.PlainText }).(pulumi.StringOutput)
}

func (o GetDataKeyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDataKeyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDataKeyResultOutput{})
}

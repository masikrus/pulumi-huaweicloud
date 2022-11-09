// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package obs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get info of special HuaweiCloud obs object.
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Obs"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Obs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Obs.GetBucketObject(ctx, &obs.GetBucketObjectArgs{
//				Bucket: "my-test-bucket",
//				Key:    "new-key",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBucketObject(ctx *pulumi.Context, args *LookupBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupBucketObjectResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupBucketObjectResult
	err := ctx.Invoke("huaweicloud:Obs/getBucketObject:getBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectArgs struct {
	// The name of the bucket to put the file in.
	Bucket string `pulumi:"bucket"`
	// The name of the object once it is in the bucket.
	Key string `pulumi:"key"`
	// The region in which to obtain the OBS object. If omitted, the provider-level region will
	// be used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getBucketObject.
type LookupBucketObjectResult struct {
	// The content of an object which is available only for objects which have a human-readable Content-Type
	// (text/* and application/json) and smaller than **64KB**. This is to prevent printing unsafe characters and
	// potentially downloading large amount of data.
	Body   string `pulumi:"body"`
	Bucket string `pulumi:"bucket"`
	// a standard MIME type describing the format of the object data, e.g. application/octet-stream. All
	// Valid MIME Types are valid for this input.
	ContentType string `pulumi:"contentType"`
	// the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the
	// server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the
	// server-side encryption.
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id     string `pulumi:"id"`
	Key    string `pulumi:"key"`
	Region string `pulumi:"region"`
	// the size of the object in bytes.
	Size int `pulumi:"size"`
	// specifies the storage class of the object.
	StorageClass string `pulumi:"storageClass"`
	// a unique version ID value for the object, if bucket versioning is enabled.
	VersionId string `pulumi:"versionId"`
}

func LookupBucketObjectOutput(ctx *pulumi.Context, args LookupBucketObjectOutputArgs, opts ...pulumi.InvokeOption) LookupBucketObjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketObjectResult, error) {
			args := v.(LookupBucketObjectArgs)
			r, err := LookupBucketObject(ctx, &args, opts...)
			var s LookupBucketObjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketObjectResultOutput)
}

// A collection of arguments for invoking getBucketObject.
type LookupBucketObjectOutputArgs struct {
	// The name of the bucket to put the file in.
	Bucket pulumi.StringInput `pulumi:"bucket"`
	// The name of the object once it is in the bucket.
	Key pulumi.StringInput `pulumi:"key"`
	// The region in which to obtain the OBS object. If omitted, the provider-level region will
	// be used.
	Region pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupBucketObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectArgs)(nil)).Elem()
}

// A collection of values returned by getBucketObject.
type LookupBucketObjectResultOutput struct{ *pulumi.OutputState }

func (LookupBucketObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectResult)(nil)).Elem()
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutput() LookupBucketObjectResultOutput {
	return o
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutputWithContext(ctx context.Context) LookupBucketObjectResultOutput {
	return o
}

// The content of an object which is available only for objects which have a human-readable Content-Type
// (text/* and application/json) and smaller than **64KB**. This is to prevent printing unsafe characters and
// potentially downloading large amount of data.
func (o LookupBucketObjectResultOutput) Body() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Body }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// a standard MIME type describing the format of the object data, e.g. application/octet-stream. All
// Valid MIME Types are valid for this input.
func (o LookupBucketObjectResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// the ETag generated for the object (an MD5 sum of the object content). When the object is encrypted on the
// server side, the ETag value is not the MD5 value of the object, but the unique identifier calculated through the
// server-side encryption.
func (o LookupBucketObjectResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupBucketObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupBucketObjectResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Region }).(pulumi.StringOutput)
}

// the size of the object in bytes.
func (o LookupBucketObjectResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) int { return v.Size }).(pulumi.IntOutput)
}

// specifies the storage class of the object.
func (o LookupBucketObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

// a unique version ID value for the object, if bucket versioning is enabled.
func (o LookupBucketObjectResultOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.VersionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketObjectResultOutput{})
}

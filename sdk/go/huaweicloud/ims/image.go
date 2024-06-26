// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ims

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image resource within HuaweiCloud IMS.
//
// ## Example Usage
// ### Creating an image from OBS bucket
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Ims"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Ims.NewImage(ctx, "imsTestFile", &Ims.ImageArgs{
//				Description: pulumi.String("Create an image from the OBS bucket."),
//				ImageUrl:    pulumi.String("ims-image:centos70.qcow2"),
//				MinDisk:     pulumi.Int(40),
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar1"),
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating a whole image from an existing ECS
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Ims"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vaultId := cfg.RequireObject("vaultId")
//			instanceId := cfg.RequireObject("instanceId")
//			_, err := Ims.NewImage(ctx, "test", &Ims.ImageArgs{
//				InstanceId: pulumi.Any(instanceId),
//				VaultId:    pulumi.Any(vaultId),
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar2"),
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating a whole image from CBR backup
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Ims"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			backupId := cfg.RequireObject("backupId")
//			_, err := Ims.NewImage(ctx, "test", &Ims.ImageArgs{
//				BackupId: pulumi.Any(backupId),
//				Tags: pulumi.StringMap{
//					"foo": pulumi.String("bar1"),
//					"key": pulumi.String("value"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Images can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Ims/image:Image my_image <id>
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response. The missing attributes include`vault_id`. It is generally recommended running `terraform plan` after importing the image. You can then decide if changes should be applied to the image, or the resource definition should be updated to align with the image. Also you can ignore changes as below. resource "huaweicloud_images_image" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	vault_id,
//
//	]
//
//	} }
type Image struct {
	pulumi.CustomResourceState

	// The ID of the CBR backup that needs to be converted into an image. This
	// parameter is mandatory when you create a private whole image from a CBR backup.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// The checksum of the data associated with the image.
	Checksum pulumi.StringOutput `pulumi:"checksum"`
	// The master key used for encrypting an image.
	CmkId pulumi.StringPtrOutput `pulumi:"cmkId"`
	// The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
	// or 'server_backup,*backup_id*'.
	DataOrigin pulumi.StringOutput `pulumi:"dataOrigin"`
	// A description of the image.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
	DiskFormat pulumi.StringOutput `pulumi:"diskFormat"`
	// The enterprise project id of the image. Changing this creates a
	// new image.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// The size(bytes) of the image file format.
	ImageSize pulumi.StringOutput `pulumi:"imageSize"`
	// The URL of the external image file in the OBS bucket. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
	// name:Image file name*.
	ImageUrl pulumi.StringPtrOutput `pulumi:"imageUrl"`
	// The ID of the ECS that needs to be converted into an image. This
	// parameter is mandatory when you create a private image or a private whole image from an ECS.
	// If the value of `vaultId` is not empty, then a whole image will be created.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// If automatic configuration is required, set the value to true. Otherwise, set
	// the value to false.
	IsConfig pulumi.BoolPtrOutput `pulumi:"isConfig"`
	// The maximum memory of the image in the unit of MB.
	MaxRam pulumi.IntOutput `pulumi:"maxRam"`
	// The minimum size of the system disk in the unit of GB. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
	// to 1024 GB.
	MinDisk pulumi.IntPtrOutput `pulumi:"minDisk"`
	// The minimum memory of the image in the unit of MB. The default value is 0,
	// indicating that the memory is not restricted.
	MinRam pulumi.IntOutput `pulumi:"minRam"`
	// The name of the image.
	Name pulumi.StringOutput `pulumi:"name"`
	// The OS version. This parameter is valid when you create a private image
	// from an external file uploaded to an OBS bucket.
	OsVersion pulumi.StringOutput `pulumi:"osVersion"`
	// The status of the image.
	Status pulumi.StringOutput `pulumi:"status"`
	// The tags of the image.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// The ID of the vault to which an ECS is to be added or has been added.
	// This parameter is mandatory when you create a private whole image from an ECS.
	VaultId pulumi.StringPtrOutput `pulumi:"vaultId"`
	// Whether the image is visible to other tenants.
	Visibility pulumi.StringOutput `pulumi:"visibility"`
}

// NewImage registers a new resource with the given unique name, arguments, and options.
func NewImage(ctx *pulumi.Context,
	name string, args *ImageArgs, opts ...pulumi.ResourceOption) (*Image, error) {
	if args == nil {
		args = &ImageArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Image
	err := ctx.RegisterResource("huaweicloud:Ims/image:Image", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImage gets an existing Image resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageState, opts ...pulumi.ResourceOption) (*Image, error) {
	var resource Image
	err := ctx.ReadResource("huaweicloud:Ims/image:Image", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Image resources.
type imageState struct {
	// The ID of the CBR backup that needs to be converted into an image. This
	// parameter is mandatory when you create a private whole image from a CBR backup.
	BackupId *string `pulumi:"backupId"`
	// The checksum of the data associated with the image.
	Checksum *string `pulumi:"checksum"`
	// The master key used for encrypting an image.
	CmkId *string `pulumi:"cmkId"`
	// The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
	// or 'server_backup,*backup_id*'.
	DataOrigin *string `pulumi:"dataOrigin"`
	// A description of the image.
	Description *string `pulumi:"description"`
	// The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
	DiskFormat *string `pulumi:"diskFormat"`
	// The enterprise project id of the image. Changing this creates a
	// new image.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The size(bytes) of the image file format.
	ImageSize *string `pulumi:"imageSize"`
	// The URL of the external image file in the OBS bucket. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
	// name:Image file name*.
	ImageUrl *string `pulumi:"imageUrl"`
	// The ID of the ECS that needs to be converted into an image. This
	// parameter is mandatory when you create a private image or a private whole image from an ECS.
	// If the value of `vaultId` is not empty, then a whole image will be created.
	InstanceId *string `pulumi:"instanceId"`
	// If automatic configuration is required, set the value to true. Otherwise, set
	// the value to false.
	IsConfig *bool `pulumi:"isConfig"`
	// The maximum memory of the image in the unit of MB.
	MaxRam *int `pulumi:"maxRam"`
	// The minimum size of the system disk in the unit of GB. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
	// to 1024 GB.
	MinDisk *int `pulumi:"minDisk"`
	// The minimum memory of the image in the unit of MB. The default value is 0,
	// indicating that the memory is not restricted.
	MinRam *int `pulumi:"minRam"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// The OS version. This parameter is valid when you create a private image
	// from an external file uploaded to an OBS bucket.
	OsVersion *string `pulumi:"osVersion"`
	// The status of the image.
	Status *string `pulumi:"status"`
	// The tags of the image.
	Tags map[string]string `pulumi:"tags"`
	// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
	Type *string `pulumi:"type"`
	// The ID of the vault to which an ECS is to be added or has been added.
	// This parameter is mandatory when you create a private whole image from an ECS.
	VaultId *string `pulumi:"vaultId"`
	// Whether the image is visible to other tenants.
	Visibility *string `pulumi:"visibility"`
}

type ImageState struct {
	// The ID of the CBR backup that needs to be converted into an image. This
	// parameter is mandatory when you create a private whole image from a CBR backup.
	BackupId pulumi.StringPtrInput
	// The checksum of the data associated with the image.
	Checksum pulumi.StringPtrInput
	// The master key used for encrypting an image.
	CmkId pulumi.StringPtrInput
	// The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
	// or 'server_backup,*backup_id*'.
	DataOrigin pulumi.StringPtrInput
	// A description of the image.
	Description pulumi.StringPtrInput
	// The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
	DiskFormat pulumi.StringPtrInput
	// The enterprise project id of the image. Changing this creates a
	// new image.
	EnterpriseProjectId pulumi.StringPtrInput
	// The size(bytes) of the image file format.
	ImageSize pulumi.StringPtrInput
	// The URL of the external image file in the OBS bucket. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
	// name:Image file name*.
	ImageUrl pulumi.StringPtrInput
	// The ID of the ECS that needs to be converted into an image. This
	// parameter is mandatory when you create a private image or a private whole image from an ECS.
	// If the value of `vaultId` is not empty, then a whole image will be created.
	InstanceId pulumi.StringPtrInput
	// If automatic configuration is required, set the value to true. Otherwise, set
	// the value to false.
	IsConfig pulumi.BoolPtrInput
	// The maximum memory of the image in the unit of MB.
	MaxRam pulumi.IntPtrInput
	// The minimum size of the system disk in the unit of GB. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
	// to 1024 GB.
	MinDisk pulumi.IntPtrInput
	// The minimum memory of the image in the unit of MB. The default value is 0,
	// indicating that the memory is not restricted.
	MinRam pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// The OS version. This parameter is valid when you create a private image
	// from an external file uploaded to an OBS bucket.
	OsVersion pulumi.StringPtrInput
	// The status of the image.
	Status pulumi.StringPtrInput
	// The tags of the image.
	Tags pulumi.StringMapInput
	// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
	Type pulumi.StringPtrInput
	// The ID of the vault to which an ECS is to be added or has been added.
	// This parameter is mandatory when you create a private whole image from an ECS.
	VaultId pulumi.StringPtrInput
	// Whether the image is visible to other tenants.
	Visibility pulumi.StringPtrInput
}

func (ImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageState)(nil)).Elem()
}

type imageArgs struct {
	// The ID of the CBR backup that needs to be converted into an image. This
	// parameter is mandatory when you create a private whole image from a CBR backup.
	BackupId *string `pulumi:"backupId"`
	// The master key used for encrypting an image.
	CmkId *string `pulumi:"cmkId"`
	// A description of the image.
	Description *string `pulumi:"description"`
	// The enterprise project id of the image. Changing this creates a
	// new image.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The URL of the external image file in the OBS bucket. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
	// name:Image file name*.
	ImageUrl *string `pulumi:"imageUrl"`
	// The ID of the ECS that needs to be converted into an image. This
	// parameter is mandatory when you create a private image or a private whole image from an ECS.
	// If the value of `vaultId` is not empty, then a whole image will be created.
	InstanceId *string `pulumi:"instanceId"`
	// If automatic configuration is required, set the value to true. Otherwise, set
	// the value to false.
	IsConfig *bool `pulumi:"isConfig"`
	// The maximum memory of the image in the unit of MB.
	MaxRam *int `pulumi:"maxRam"`
	// The minimum size of the system disk in the unit of GB. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
	// to 1024 GB.
	MinDisk *int `pulumi:"minDisk"`
	// The minimum memory of the image in the unit of MB. The default value is 0,
	// indicating that the memory is not restricted.
	MinRam *int `pulumi:"minRam"`
	// The name of the image.
	Name *string `pulumi:"name"`
	// The OS version. This parameter is valid when you create a private image
	// from an external file uploaded to an OBS bucket.
	OsVersion *string `pulumi:"osVersion"`
	// The tags of the image.
	Tags map[string]string `pulumi:"tags"`
	// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
	Type *string `pulumi:"type"`
	// The ID of the vault to which an ECS is to be added or has been added.
	// This parameter is mandatory when you create a private whole image from an ECS.
	VaultId *string `pulumi:"vaultId"`
}

// The set of arguments for constructing a Image resource.
type ImageArgs struct {
	// The ID of the CBR backup that needs to be converted into an image. This
	// parameter is mandatory when you create a private whole image from a CBR backup.
	BackupId pulumi.StringPtrInput
	// The master key used for encrypting an image.
	CmkId pulumi.StringPtrInput
	// A description of the image.
	Description pulumi.StringPtrInput
	// The enterprise project id of the image. Changing this creates a
	// new image.
	EnterpriseProjectId pulumi.StringPtrInput
	// The URL of the external image file in the OBS bucket. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
	// name:Image file name*.
	ImageUrl pulumi.StringPtrInput
	// The ID of the ECS that needs to be converted into an image. This
	// parameter is mandatory when you create a private image or a private whole image from an ECS.
	// If the value of `vaultId` is not empty, then a whole image will be created.
	InstanceId pulumi.StringPtrInput
	// If automatic configuration is required, set the value to true. Otherwise, set
	// the value to false.
	IsConfig pulumi.BoolPtrInput
	// The maximum memory of the image in the unit of MB.
	MaxRam pulumi.IntPtrInput
	// The minimum size of the system disk in the unit of GB. This parameter is
	// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
	// to 1024 GB.
	MinDisk pulumi.IntPtrInput
	// The minimum memory of the image in the unit of MB. The default value is 0,
	// indicating that the memory is not restricted.
	MinRam pulumi.IntPtrInput
	// The name of the image.
	Name pulumi.StringPtrInput
	// The OS version. This parameter is valid when you create a private image
	// from an external file uploaded to an OBS bucket.
	OsVersion pulumi.StringPtrInput
	// The tags of the image.
	Tags pulumi.StringMapInput
	// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
	Type pulumi.StringPtrInput
	// The ID of the vault to which an ECS is to be added or has been added.
	// This parameter is mandatory when you create a private whole image from an ECS.
	VaultId pulumi.StringPtrInput
}

func (ImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageArgs)(nil)).Elem()
}

type ImageInput interface {
	pulumi.Input

	ToImageOutput() ImageOutput
	ToImageOutputWithContext(ctx context.Context) ImageOutput
}

func (*Image) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (i *Image) ToImageOutput() ImageOutput {
	return i.ToImageOutputWithContext(context.Background())
}

func (i *Image) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageOutput)
}

// ImageArrayInput is an input type that accepts ImageArray and ImageArrayOutput values.
// You can construct a concrete instance of `ImageArrayInput` via:
//
//	ImageArray{ ImageArgs{...} }
type ImageArrayInput interface {
	pulumi.Input

	ToImageArrayOutput() ImageArrayOutput
	ToImageArrayOutputWithContext(context.Context) ImageArrayOutput
}

type ImageArray []ImageInput

func (ImageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (i ImageArray) ToImageArrayOutput() ImageArrayOutput {
	return i.ToImageArrayOutputWithContext(context.Background())
}

func (i ImageArray) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageArrayOutput)
}

// ImageMapInput is an input type that accepts ImageMap and ImageMapOutput values.
// You can construct a concrete instance of `ImageMapInput` via:
//
//	ImageMap{ "key": ImageArgs{...} }
type ImageMapInput interface {
	pulumi.Input

	ToImageMapOutput() ImageMapOutput
	ToImageMapOutputWithContext(context.Context) ImageMapOutput
}

type ImageMap map[string]ImageInput

func (ImageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (i ImageMap) ToImageMapOutput() ImageMapOutput {
	return i.ToImageMapOutputWithContext(context.Background())
}

func (i ImageMap) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageMapOutput)
}

type ImageOutput struct{ *pulumi.OutputState }

func (ImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Image)(nil)).Elem()
}

func (o ImageOutput) ToImageOutput() ImageOutput {
	return o
}

func (o ImageOutput) ToImageOutputWithContext(ctx context.Context) ImageOutput {
	return o
}

// The ID of the CBR backup that needs to be converted into an image. This
// parameter is mandatory when you create a private whole image from a CBR backup.
func (o ImageOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// The checksum of the data associated with the image.
func (o ImageOutput) Checksum() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Checksum }).(pulumi.StringOutput)
}

// The master key used for encrypting an image.
func (o ImageOutput) CmkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.CmkId }).(pulumi.StringPtrOutput)
}

// The image resource. The pattern can be 'instance,*instance_id*', 'file,*image_url*'
// or 'server_backup,*backup_id*'.
func (o ImageOutput) DataOrigin() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.DataOrigin }).(pulumi.StringOutput)
}

// A description of the image.
func (o ImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The image file format. The value can be `vhd`, `zvhd`, `raw`, `zvhd2`, or `qcow2`.
func (o ImageOutput) DiskFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.DiskFormat }).(pulumi.StringOutput)
}

// The enterprise project id of the image. Changing this creates a
// new image.
func (o ImageOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The size(bytes) of the image file format.
func (o ImageOutput) ImageSize() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.ImageSize }).(pulumi.StringOutput)
}

// The URL of the external image file in the OBS bucket. This parameter is
// mandatory when you create a private image from an external file uploaded to an OBS bucket. The format is *OBS bucket
// name:Image file name*.
func (o ImageOutput) ImageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.ImageUrl }).(pulumi.StringPtrOutput)
}

// The ID of the ECS that needs to be converted into an image. This
// parameter is mandatory when you create a private image or a private whole image from an ECS.
// If the value of `vaultId` is not empty, then a whole image will be created.
func (o ImageOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// If automatic configuration is required, set the value to true. Otherwise, set
// the value to false.
func (o ImageOutput) IsConfig() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.BoolPtrOutput { return v.IsConfig }).(pulumi.BoolPtrOutput)
}

// The maximum memory of the image in the unit of MB.
func (o ImageOutput) MaxRam() pulumi.IntOutput {
	return o.ApplyT(func(v *Image) pulumi.IntOutput { return v.MaxRam }).(pulumi.IntOutput)
}

// The minimum size of the system disk in the unit of GB. This parameter is
// mandatory when you create a private image from an external file uploaded to an OBS bucket. The value ranges from 1 GB
// to 1024 GB.
func (o ImageOutput) MinDisk() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.IntPtrOutput { return v.MinDisk }).(pulumi.IntPtrOutput)
}

// The minimum memory of the image in the unit of MB. The default value is 0,
// indicating that the memory is not restricted.
func (o ImageOutput) MinRam() pulumi.IntOutput {
	return o.ApplyT(func(v *Image) pulumi.IntOutput { return v.MinRam }).(pulumi.IntOutput)
}

// The name of the image.
func (o ImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The OS version. This parameter is valid when you create a private image
// from an external file uploaded to an OBS bucket.
func (o ImageOutput) OsVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.OsVersion }).(pulumi.StringOutput)
}

// The status of the image.
func (o ImageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The tags of the image.
func (o ImageOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Image) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The image type. Must be one of `ECS`, `FusionCompute`, `BMS`, or `Ironic`.
func (o ImageOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// The ID of the vault to which an ECS is to be added or has been added.
// This parameter is mandatory when you create a private whole image from an ECS.
func (o ImageOutput) VaultId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Image) pulumi.StringPtrOutput { return v.VaultId }).(pulumi.StringPtrOutput)
}

// Whether the image is visible to other tenants.
func (o ImageOutput) Visibility() pulumi.StringOutput {
	return o.ApplyT(func(v *Image) pulumi.StringOutput { return v.Visibility }).(pulumi.StringOutput)
}

type ImageArrayOutput struct{ *pulumi.OutputState }

func (ImageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Image)(nil)).Elem()
}

func (o ImageArrayOutput) ToImageArrayOutput() ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) ToImageArrayOutputWithContext(ctx context.Context) ImageArrayOutput {
	return o
}

func (o ImageArrayOutput) Index(i pulumi.IntInput) ImageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Image {
		return vs[0].([]*Image)[vs[1].(int)]
	}).(ImageOutput)
}

type ImageMapOutput struct{ *pulumi.OutputState }

func (ImageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Image)(nil)).Elem()
}

func (o ImageMapOutput) ToImageMapOutput() ImageMapOutput {
	return o
}

func (o ImageMapOutput) ToImageMapOutputWithContext(ctx context.Context) ImageMapOutput {
	return o
}

func (o ImageMapOutput) MapIndex(k pulumi.StringInput) ImageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Image {
		return vs[0].(map[string]*Image)[vs[1].(string)]
	}).(ImageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageInput)(nil)).Elem(), &Image{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageArrayInput)(nil)).Elem(), ImageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageMapInput)(nil)).Elem(), ImageMap{})
	pulumi.RegisterOutputType(ImageOutput{})
	pulumi.RegisterOutputType(ImageArrayOutput{})
	pulumi.RegisterOutputType(ImageMapOutput{})
}

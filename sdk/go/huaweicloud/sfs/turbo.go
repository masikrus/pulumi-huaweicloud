// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an Shared File System (SFS) Turbo resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Sfs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			subnetId := cfg.RequireObject("subnetId")
//			secgroupId := cfg.RequireObject("secgroupId")
//			testAz := cfg.RequireObject("testAz")
//			_, err := Sfs.NewTurbo(ctx, "sfs-turbo-1", &Sfs.TurboArgs{
//				Size:             pulumi.Int(500),
//				ShareProto:       pulumi.String("NFS"),
//				VpcId:            pulumi.Any(vpcId),
//				SubnetId:         pulumi.Any(subnetId),
//				SecurityGroupId:  pulumi.Any(secgroupId),
//				AvailabilityZone: pulumi.Any(testAz),
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
// SFS Turbo can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Sfs/turbo:Turbo huaweicloud_sfs_turbo 1e3d5306-24c9-4316-9185-70e9787d71ab
//
// ```
type Turbo struct {
	pulumi.CustomResourceState

	// Specifies the availability zone where the file system is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The available capacity of the SFS Turbo file system in the unit of GB.
	AvailableCapacity pulumi.StringOutput `pulumi:"availableCapacity"`
	// Specifies the ID of a KMS key to encrypt the file system. Changing this
	// will create a new resource.
	CryptKeyId pulumi.StringPtrOutput `pulumi:"cryptKeyId"`
	// Specifies whether the file system is enhanced or not. Changing this will
	// create a new resource.
	Enhanced pulumi.BoolOutput `pulumi:"enhanced"`
	// The enterprise project id of the file system. Changing this
	// will create a new resource.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Tthe mount point of the SFS Turbo file system.
	ExportLocation pulumi.StringOutput `pulumi:"exportLocation"`
	// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
	// characters and must start with a letter. Changing this will create a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The region in which to create the SFS Turbo resource. If omitted, the
	// provider-level region will be used. Changing this creates a new SFS Turbo resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the security group ID. Changing this will create a new
	// resource.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Specifies the protocol for sharing file systems. The valid value is NFS.
	// Changing this will create a new resource.
	ShareProto pulumi.StringPtrOutput `pulumi:"shareProto"`
	// Specifies the file system type. The valid values are STANDARD and
	// PERFORMANCE Changing this will create a new resource.
	ShareType pulumi.StringPtrOutput `pulumi:"shareType"`
	// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
	// and must be large than 10240 for an enhanced file system.
	Size pulumi.IntOutput `pulumi:"size"`
	// The status of the SFS Turbo file system.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the network ID of the subnet. Changing this will create a new
	// resource.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The version ID of the SFS Turbo file system.
	Version pulumi.StringOutput `pulumi:"version"`
	// Specifies the VPC ID. Changing this will create a new resource.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewTurbo registers a new resource with the given unique name, arguments, and options.
func NewTurbo(ctx *pulumi.Context,
	name string, args *TurboArgs, opts ...pulumi.ResourceOption) (*Turbo, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Turbo
	err := ctx.RegisterResource("huaweicloud:Sfs/turbo:Turbo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTurbo gets an existing Turbo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTurbo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TurboState, opts ...pulumi.ResourceOption) (*Turbo, error) {
	var resource Turbo
	err := ctx.ReadResource("huaweicloud:Sfs/turbo:Turbo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Turbo resources.
type turboState struct {
	// Specifies the availability zone where the file system is located.
	// Changing this will create a new resource.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The available capacity of the SFS Turbo file system in the unit of GB.
	AvailableCapacity *string `pulumi:"availableCapacity"`
	// Specifies the ID of a KMS key to encrypt the file system. Changing this
	// will create a new resource.
	CryptKeyId *string `pulumi:"cryptKeyId"`
	// Specifies whether the file system is enhanced or not. Changing this will
	// create a new resource.
	Enhanced *bool `pulumi:"enhanced"`
	// The enterprise project id of the file system. Changing this
	// will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Tthe mount point of the SFS Turbo file system.
	ExportLocation *string `pulumi:"exportLocation"`
	// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
	// characters and must start with a letter. Changing this will create a new resource.
	Name *string `pulumi:"name"`
	// The region in which to create the SFS Turbo resource. If omitted, the
	// provider-level region will be used. Changing this creates a new SFS Turbo resource.
	Region *string `pulumi:"region"`
	// Specifies the security group ID. Changing this will create a new
	// resource.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Specifies the protocol for sharing file systems. The valid value is NFS.
	// Changing this will create a new resource.
	ShareProto *string `pulumi:"shareProto"`
	// Specifies the file system type. The valid values are STANDARD and
	// PERFORMANCE Changing this will create a new resource.
	ShareType *string `pulumi:"shareType"`
	// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
	// and must be large than 10240 for an enhanced file system.
	Size *int `pulumi:"size"`
	// The status of the SFS Turbo file system.
	Status *string `pulumi:"status"`
	// Specifies the network ID of the subnet. Changing this will create a new
	// resource.
	SubnetId *string `pulumi:"subnetId"`
	// The version ID of the SFS Turbo file system.
	Version *string `pulumi:"version"`
	// Specifies the VPC ID. Changing this will create a new resource.
	VpcId *string `pulumi:"vpcId"`
}

type TurboState struct {
	// Specifies the availability zone where the file system is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringPtrInput
	// The available capacity of the SFS Turbo file system in the unit of GB.
	AvailableCapacity pulumi.StringPtrInput
	// Specifies the ID of a KMS key to encrypt the file system. Changing this
	// will create a new resource.
	CryptKeyId pulumi.StringPtrInput
	// Specifies whether the file system is enhanced or not. Changing this will
	// create a new resource.
	Enhanced pulumi.BoolPtrInput
	// The enterprise project id of the file system. Changing this
	// will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Tthe mount point of the SFS Turbo file system.
	ExportLocation pulumi.StringPtrInput
	// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
	// characters and must start with a letter. Changing this will create a new resource.
	Name pulumi.StringPtrInput
	// The region in which to create the SFS Turbo resource. If omitted, the
	// provider-level region will be used. Changing this creates a new SFS Turbo resource.
	Region pulumi.StringPtrInput
	// Specifies the security group ID. Changing this will create a new
	// resource.
	SecurityGroupId pulumi.StringPtrInput
	// Specifies the protocol for sharing file systems. The valid value is NFS.
	// Changing this will create a new resource.
	ShareProto pulumi.StringPtrInput
	// Specifies the file system type. The valid values are STANDARD and
	// PERFORMANCE Changing this will create a new resource.
	ShareType pulumi.StringPtrInput
	// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
	// and must be large than 10240 for an enhanced file system.
	Size pulumi.IntPtrInput
	// The status of the SFS Turbo file system.
	Status pulumi.StringPtrInput
	// Specifies the network ID of the subnet. Changing this will create a new
	// resource.
	SubnetId pulumi.StringPtrInput
	// The version ID of the SFS Turbo file system.
	Version pulumi.StringPtrInput
	// Specifies the VPC ID. Changing this will create a new resource.
	VpcId pulumi.StringPtrInput
}

func (TurboState) ElementType() reflect.Type {
	return reflect.TypeOf((*turboState)(nil)).Elem()
}

type turboArgs struct {
	// Specifies the availability zone where the file system is located.
	// Changing this will create a new resource.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies the ID of a KMS key to encrypt the file system. Changing this
	// will create a new resource.
	CryptKeyId *string `pulumi:"cryptKeyId"`
	// Specifies whether the file system is enhanced or not. Changing this will
	// create a new resource.
	Enhanced *bool `pulumi:"enhanced"`
	// The enterprise project id of the file system. Changing this
	// will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
	// characters and must start with a letter. Changing this will create a new resource.
	Name *string `pulumi:"name"`
	// The region in which to create the SFS Turbo resource. If omitted, the
	// provider-level region will be used. Changing this creates a new SFS Turbo resource.
	Region *string `pulumi:"region"`
	// Specifies the security group ID. Changing this will create a new
	// resource.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Specifies the protocol for sharing file systems. The valid value is NFS.
	// Changing this will create a new resource.
	ShareProto *string `pulumi:"shareProto"`
	// Specifies the file system type. The valid values are STANDARD and
	// PERFORMANCE Changing this will create a new resource.
	ShareType *string `pulumi:"shareType"`
	// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
	// and must be large than 10240 for an enhanced file system.
	Size int `pulumi:"size"`
	// Specifies the network ID of the subnet. Changing this will create a new
	// resource.
	SubnetId string `pulumi:"subnetId"`
	// Specifies the VPC ID. Changing this will create a new resource.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Turbo resource.
type TurboArgs struct {
	// Specifies the availability zone where the file system is located.
	// Changing this will create a new resource.
	AvailabilityZone pulumi.StringInput
	// Specifies the ID of a KMS key to encrypt the file system. Changing this
	// will create a new resource.
	CryptKeyId pulumi.StringPtrInput
	// Specifies whether the file system is enhanced or not. Changing this will
	// create a new resource.
	Enhanced pulumi.BoolPtrInput
	// The enterprise project id of the file system. Changing this
	// will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
	// characters and must start with a letter. Changing this will create a new resource.
	Name pulumi.StringPtrInput
	// The region in which to create the SFS Turbo resource. If omitted, the
	// provider-level region will be used. Changing this creates a new SFS Turbo resource.
	Region pulumi.StringPtrInput
	// Specifies the security group ID. Changing this will create a new
	// resource.
	SecurityGroupId pulumi.StringInput
	// Specifies the protocol for sharing file systems. The valid value is NFS.
	// Changing this will create a new resource.
	ShareProto pulumi.StringPtrInput
	// Specifies the file system type. The valid values are STANDARD and
	// PERFORMANCE Changing this will create a new resource.
	ShareType pulumi.StringPtrInput
	// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
	// and must be large than 10240 for an enhanced file system.
	Size pulumi.IntInput
	// Specifies the network ID of the subnet. Changing this will create a new
	// resource.
	SubnetId pulumi.StringInput
	// Specifies the VPC ID. Changing this will create a new resource.
	VpcId pulumi.StringInput
}

func (TurboArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*turboArgs)(nil)).Elem()
}

type TurboInput interface {
	pulumi.Input

	ToTurboOutput() TurboOutput
	ToTurboOutputWithContext(ctx context.Context) TurboOutput
}

func (*Turbo) ElementType() reflect.Type {
	return reflect.TypeOf((**Turbo)(nil)).Elem()
}

func (i *Turbo) ToTurboOutput() TurboOutput {
	return i.ToTurboOutputWithContext(context.Background())
}

func (i *Turbo) ToTurboOutputWithContext(ctx context.Context) TurboOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TurboOutput)
}

// TurboArrayInput is an input type that accepts TurboArray and TurboArrayOutput values.
// You can construct a concrete instance of `TurboArrayInput` via:
//
//	TurboArray{ TurboArgs{...} }
type TurboArrayInput interface {
	pulumi.Input

	ToTurboArrayOutput() TurboArrayOutput
	ToTurboArrayOutputWithContext(context.Context) TurboArrayOutput
}

type TurboArray []TurboInput

func (TurboArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Turbo)(nil)).Elem()
}

func (i TurboArray) ToTurboArrayOutput() TurboArrayOutput {
	return i.ToTurboArrayOutputWithContext(context.Background())
}

func (i TurboArray) ToTurboArrayOutputWithContext(ctx context.Context) TurboArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TurboArrayOutput)
}

// TurboMapInput is an input type that accepts TurboMap and TurboMapOutput values.
// You can construct a concrete instance of `TurboMapInput` via:
//
//	TurboMap{ "key": TurboArgs{...} }
type TurboMapInput interface {
	pulumi.Input

	ToTurboMapOutput() TurboMapOutput
	ToTurboMapOutputWithContext(context.Context) TurboMapOutput
}

type TurboMap map[string]TurboInput

func (TurboMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Turbo)(nil)).Elem()
}

func (i TurboMap) ToTurboMapOutput() TurboMapOutput {
	return i.ToTurboMapOutputWithContext(context.Background())
}

func (i TurboMap) ToTurboMapOutputWithContext(ctx context.Context) TurboMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TurboMapOutput)
}

type TurboOutput struct{ *pulumi.OutputState }

func (TurboOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Turbo)(nil)).Elem()
}

func (o TurboOutput) ToTurboOutput() TurboOutput {
	return o
}

func (o TurboOutput) ToTurboOutputWithContext(ctx context.Context) TurboOutput {
	return o
}

// Specifies the availability zone where the file system is located.
// Changing this will create a new resource.
func (o TurboOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The available capacity of the SFS Turbo file system in the unit of GB.
func (o TurboOutput) AvailableCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.AvailableCapacity }).(pulumi.StringOutput)
}

// Specifies the ID of a KMS key to encrypt the file system. Changing this
// will create a new resource.
func (o TurboOutput) CryptKeyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringPtrOutput { return v.CryptKeyId }).(pulumi.StringPtrOutput)
}

// Specifies whether the file system is enhanced or not. Changing this will
// create a new resource.
func (o TurboOutput) Enhanced() pulumi.BoolOutput {
	return o.ApplyT(func(v *Turbo) pulumi.BoolOutput { return v.Enhanced }).(pulumi.BoolOutput)
}

// The enterprise project id of the file system. Changing this
// will create a new resource.
func (o TurboOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Tthe mount point of the SFS Turbo file system.
func (o TurboOutput) ExportLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.ExportLocation }).(pulumi.StringOutput)
}

// Specifies the name of an SFS Turbo file system. The value contains 4 to 64
// characters and must start with a letter. Changing this will create a new resource.
func (o TurboOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The region in which to create the SFS Turbo resource. If omitted, the
// provider-level region will be used. Changing this creates a new SFS Turbo resource.
func (o TurboOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the security group ID. Changing this will create a new
// resource.
func (o TurboOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Specifies the protocol for sharing file systems. The valid value is NFS.
// Changing this will create a new resource.
func (o TurboOutput) ShareProto() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringPtrOutput { return v.ShareProto }).(pulumi.StringPtrOutput)
}

// Specifies the file system type. The valid values are STANDARD and
// PERFORMANCE Changing this will create a new resource.
func (o TurboOutput) ShareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringPtrOutput { return v.ShareType }).(pulumi.StringPtrOutput)
}

// Specifies the capacity of a common file system, in GB. The value ranges from 500 to 32768,
// and must be large than 10240 for an enhanced file system.
func (o TurboOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Turbo) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The status of the SFS Turbo file system.
func (o TurboOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the network ID of the subnet. Changing this will create a new
// resource.
func (o TurboOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The version ID of the SFS Turbo file system.
func (o TurboOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Specifies the VPC ID. Changing this will create a new resource.
func (o TurboOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Turbo) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type TurboArrayOutput struct{ *pulumi.OutputState }

func (TurboArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Turbo)(nil)).Elem()
}

func (o TurboArrayOutput) ToTurboArrayOutput() TurboArrayOutput {
	return o
}

func (o TurboArrayOutput) ToTurboArrayOutputWithContext(ctx context.Context) TurboArrayOutput {
	return o
}

func (o TurboArrayOutput) Index(i pulumi.IntInput) TurboOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Turbo {
		return vs[0].([]*Turbo)[vs[1].(int)]
	}).(TurboOutput)
}

type TurboMapOutput struct{ *pulumi.OutputState }

func (TurboMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Turbo)(nil)).Elem()
}

func (o TurboMapOutput) ToTurboMapOutput() TurboMapOutput {
	return o
}

func (o TurboMapOutput) ToTurboMapOutputWithContext(ctx context.Context) TurboMapOutput {
	return o
}

func (o TurboMapOutput) MapIndex(k pulumi.StringInput) TurboOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Turbo {
		return vs[0].(map[string]*Turbo)[vs[1].(string)]
	}).(TurboOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TurboInput)(nil)).Elem(), &Turbo{})
	pulumi.RegisterInputType(reflect.TypeOf((*TurboArrayInput)(nil)).Elem(), TurboArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TurboMapInput)(nil)).Elem(), TurboMap{})
	pulumi.RegisterOutputType(TurboOutput{})
	pulumi.RegisterOutputType(TurboArrayOutput{})
	pulumi.RegisterOutputType(TurboMapOutput{})
}

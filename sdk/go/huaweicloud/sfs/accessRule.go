// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package sfs

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an access rule resource of Scalable File Resource (SFS).
//
// ## Example Usage
// ### Usage in VPC authorization scenarios
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
//			shareName := cfg.RequireObject("shareName")
//			vpcId := cfg.RequireObject("vpcId")
//			_, err := Sfs.NewFileSystem(ctx, "share-file", &Sfs.FileSystemArgs{
//				Size:       pulumi.Int(100),
//				ShareProto: pulumi.String("NFS"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Sfs.NewAccessRule(ctx, "rule1", &Sfs.AccessRuleArgs{
//				SfsId:    share_file.ID(),
//				AccessTo: pulumi.Any(vpcId),
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
// SFS access rule can be imported by specifying the SFS ID and access rule ID separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Sfs/accessRule:AccessRule huaweicloud_sfs_access_rule <sfs_id>/<rule_id>
//
// ```
type AccessRule struct {
	pulumi.CustomResourceState

	// Specifies the access level of the shared file system. Possible values
	// are *ro* (read-only)
	// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
	AccessLevel pulumi.StringPtrOutput `pulumi:"accessLevel"`
	// Specifies the value that defines the access rule. The value contains 1 to
	// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
	// + Set the VPC ID in VPC authorization scenarios.
	// + Set this parameter in IP address authorization scenario:
	// - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
	//   For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
	AccessTo pulumi.StringOutput `pulumi:"accessTo"`
	// Specifies the type of the share access rule. The default value is *cert*.
	// Changing this will create a new access rule.
	AccessType pulumi.StringPtrOutput `pulumi:"accessType"`
	// The region in which to create the sfs access rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new access rule resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the UUID of the shared file system. Changing this will create a new
	// access rule.
	SfsId pulumi.StringOutput `pulumi:"sfsId"`
	// The status of the share access rule.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewAccessRule registers a new resource with the given unique name, arguments, and options.
func NewAccessRule(ctx *pulumi.Context,
	name string, args *AccessRuleArgs, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessTo == nil {
		return nil, errors.New("invalid value for required argument 'AccessTo'")
	}
	if args.SfsId == nil {
		return nil, errors.New("invalid value for required argument 'SfsId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource AccessRule
	err := ctx.RegisterResource("huaweicloud:Sfs/accessRule:AccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessRule gets an existing AccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessRuleState, opts ...pulumi.ResourceOption) (*AccessRule, error) {
	var resource AccessRule
	err := ctx.ReadResource("huaweicloud:Sfs/accessRule:AccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessRule resources.
type accessRuleState struct {
	// Specifies the access level of the shared file system. Possible values
	// are *ro* (read-only)
	// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
	AccessLevel *string `pulumi:"accessLevel"`
	// Specifies the value that defines the access rule. The value contains 1 to
	// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
	// + Set the VPC ID in VPC authorization scenarios.
	// + Set this parameter in IP address authorization scenario:
	// - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
	//   For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
	AccessTo *string `pulumi:"accessTo"`
	// Specifies the type of the share access rule. The default value is *cert*.
	// Changing this will create a new access rule.
	AccessType *string `pulumi:"accessType"`
	// The region in which to create the sfs access rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new access rule resource.
	Region *string `pulumi:"region"`
	// Specifies the UUID of the shared file system. Changing this will create a new
	// access rule.
	SfsId *string `pulumi:"sfsId"`
	// The status of the share access rule.
	Status *string `pulumi:"status"`
}

type AccessRuleState struct {
	// Specifies the access level of the shared file system. Possible values
	// are *ro* (read-only)
	// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
	AccessLevel pulumi.StringPtrInput
	// Specifies the value that defines the access rule. The value contains 1 to
	// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
	// + Set the VPC ID in VPC authorization scenarios.
	// + Set this parameter in IP address authorization scenario:
	// - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
	//   For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
	AccessTo pulumi.StringPtrInput
	// Specifies the type of the share access rule. The default value is *cert*.
	// Changing this will create a new access rule.
	AccessType pulumi.StringPtrInput
	// The region in which to create the sfs access rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new access rule resource.
	Region pulumi.StringPtrInput
	// Specifies the UUID of the shared file system. Changing this will create a new
	// access rule.
	SfsId pulumi.StringPtrInput
	// The status of the share access rule.
	Status pulumi.StringPtrInput
}

func (AccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleState)(nil)).Elem()
}

type accessRuleArgs struct {
	// Specifies the access level of the shared file system. Possible values
	// are *ro* (read-only)
	// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
	AccessLevel *string `pulumi:"accessLevel"`
	// Specifies the value that defines the access rule. The value contains 1 to
	// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
	// + Set the VPC ID in VPC authorization scenarios.
	// + Set this parameter in IP address authorization scenario:
	// - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
	//   For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
	AccessTo string `pulumi:"accessTo"`
	// Specifies the type of the share access rule. The default value is *cert*.
	// Changing this will create a new access rule.
	AccessType *string `pulumi:"accessType"`
	// The region in which to create the sfs access rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new access rule resource.
	Region *string `pulumi:"region"`
	// Specifies the UUID of the shared file system. Changing this will create a new
	// access rule.
	SfsId string `pulumi:"sfsId"`
}

// The set of arguments for constructing a AccessRule resource.
type AccessRuleArgs struct {
	// Specifies the access level of the shared file system. Possible values
	// are *ro* (read-only)
	// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
	AccessLevel pulumi.StringPtrInput
	// Specifies the value that defines the access rule. The value contains 1 to
	// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
	// + Set the VPC ID in VPC authorization scenarios.
	// + Set this parameter in IP address authorization scenario:
	// - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
	//   For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
	AccessTo pulumi.StringInput
	// Specifies the type of the share access rule. The default value is *cert*.
	// Changing this will create a new access rule.
	AccessType pulumi.StringPtrInput
	// The region in which to create the sfs access rule resource. If omitted, the
	// provider-level region will be used. Changing this creates a new access rule resource.
	Region pulumi.StringPtrInput
	// Specifies the UUID of the shared file system. Changing this will create a new
	// access rule.
	SfsId pulumi.StringInput
}

func (AccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessRuleArgs)(nil)).Elem()
}

type AccessRuleInput interface {
	pulumi.Input

	ToAccessRuleOutput() AccessRuleOutput
	ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput
}

func (*AccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil)).Elem()
}

func (i *AccessRule) ToAccessRuleOutput() AccessRuleOutput {
	return i.ToAccessRuleOutputWithContext(context.Background())
}

func (i *AccessRule) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleOutput)
}

// AccessRuleArrayInput is an input type that accepts AccessRuleArray and AccessRuleArrayOutput values.
// You can construct a concrete instance of `AccessRuleArrayInput` via:
//
//	AccessRuleArray{ AccessRuleArgs{...} }
type AccessRuleArrayInput interface {
	pulumi.Input

	ToAccessRuleArrayOutput() AccessRuleArrayOutput
	ToAccessRuleArrayOutputWithContext(context.Context) AccessRuleArrayOutput
}

type AccessRuleArray []AccessRuleInput

func (AccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessRule)(nil)).Elem()
}

func (i AccessRuleArray) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return i.ToAccessRuleArrayOutputWithContext(context.Background())
}

func (i AccessRuleArray) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleArrayOutput)
}

// AccessRuleMapInput is an input type that accepts AccessRuleMap and AccessRuleMapOutput values.
// You can construct a concrete instance of `AccessRuleMapInput` via:
//
//	AccessRuleMap{ "key": AccessRuleArgs{...} }
type AccessRuleMapInput interface {
	pulumi.Input

	ToAccessRuleMapOutput() AccessRuleMapOutput
	ToAccessRuleMapOutputWithContext(context.Context) AccessRuleMapOutput
}

type AccessRuleMap map[string]AccessRuleInput

func (AccessRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessRule)(nil)).Elem()
}

func (i AccessRuleMap) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return i.ToAccessRuleMapOutputWithContext(context.Background())
}

func (i AccessRuleMap) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessRuleMapOutput)
}

type AccessRuleOutput struct{ *pulumi.OutputState }

func (AccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRule)(nil)).Elem()
}

func (o AccessRuleOutput) ToAccessRuleOutput() AccessRuleOutput {
	return o
}

func (o AccessRuleOutput) ToAccessRuleOutputWithContext(ctx context.Context) AccessRuleOutput {
	return o
}

// Specifies the access level of the shared file system. Possible values
// are *ro* (read-only)
// and *rw* (read-write). The default value is *rw* (read/write). Changing this will create a new access rule.
func (o AccessRuleOutput) AccessLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringPtrOutput { return v.AccessLevel }).(pulumi.StringPtrOutput)
}

// Specifies the value that defines the access rule. The value contains 1 to
// 255 characters. Changing this will create a new access rule. The value varies according to the scenario:
//   - Set the VPC ID in VPC authorization scenarios.
//   - Set this parameter in IP address authorization scenario:
//   - For an NFS shared file system, the value in the format of  *VPC_ID#IP_address#priority#user_permission*.
//     For example, 0157b53f-4974-4e80-91c9-098532bcaf00#2.2.2.2/16#100#all_squash,root_squash.
func (o AccessRuleOutput) AccessTo() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.AccessTo }).(pulumi.StringOutput)
}

// Specifies the type of the share access rule. The default value is *cert*.
// Changing this will create a new access rule.
func (o AccessRuleOutput) AccessType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringPtrOutput { return v.AccessType }).(pulumi.StringPtrOutput)
}

// The region in which to create the sfs access rule resource. If omitted, the
// provider-level region will be used. Changing this creates a new access rule resource.
func (o AccessRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the UUID of the shared file system. Changing this will create a new
// access rule.
func (o AccessRuleOutput) SfsId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.SfsId }).(pulumi.StringOutput)
}

// The status of the share access rule.
func (o AccessRuleOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessRule) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type AccessRuleArrayOutput struct{ *pulumi.OutputState }

func (AccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AccessRule)(nil)).Elem()
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutput() AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) ToAccessRuleArrayOutputWithContext(ctx context.Context) AccessRuleArrayOutput {
	return o
}

func (o AccessRuleArrayOutput) Index(i pulumi.IntInput) AccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AccessRule {
		return vs[0].([]*AccessRule)[vs[1].(int)]
	}).(AccessRuleOutput)
}

type AccessRuleMapOutput struct{ *pulumi.OutputState }

func (AccessRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AccessRule)(nil)).Elem()
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutput() AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) ToAccessRuleMapOutputWithContext(ctx context.Context) AccessRuleMapOutput {
	return o
}

func (o AccessRuleMapOutput) MapIndex(k pulumi.StringInput) AccessRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AccessRule {
		return vs[0].(map[string]*AccessRule)[vs[1].(string)]
	}).(AccessRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleInput)(nil)).Elem(), &AccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleArrayInput)(nil)).Elem(), AccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRuleMapInput)(nil)).Elem(), AccessRuleMap{})
	pulumi.RegisterOutputType(AccessRuleOutput{})
	pulumi.RegisterOutputType(AccessRuleArrayOutput{})
	pulumi.RegisterOutputType(AccessRuleMapOutput{})
}

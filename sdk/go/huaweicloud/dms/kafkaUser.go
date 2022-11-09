// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dms

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a DMS kafka user resource within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			kafkaInstanceId := cfg.RequireObject("kafkaInstanceId")
//			_, err := Dms.NewKafkaUser(ctx, "user", &Dms.KafkaUserArgs{
//				InstanceId: pulumi.Any(kafkaInstanceId),
//				Password:   pulumi.String("Test@123"),
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
// DMS kafka users can be imported using the kafka instance ID and user name separated by a slash, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Dms/kafkaUser:KafkaUser user c8057fe5-23a8-46ef-ad83-c0055b4e0c5c/user_1
//
// ```
type KafkaUser struct {
	pulumi.CustomResourceState

	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Specifies the name of the user. Changing this creates a new resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	Password pulumi.StringOutput `pulumi:"password"`
	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewKafkaUser registers a new resource with the given unique name, arguments, and options.
func NewKafkaUser(ctx *pulumi.Context,
	name string, args *KafkaUserArgs, opts ...pulumi.ResourceOption) (*KafkaUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KafkaUser
	err := ctx.RegisterResource("huaweicloud:Dms/kafkaUser:KafkaUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaUser gets an existing KafkaUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaUserState, opts ...pulumi.ResourceOption) (*KafkaUser, error) {
	var resource KafkaUser
	err := ctx.ReadResource("huaweicloud:Dms/kafkaUser:KafkaUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaUser resources.
type kafkaUserState struct {
	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	InstanceId *string `pulumi:"instanceId"`
	// Specifies the name of the user. Changing this creates a new resource.
	Name *string `pulumi:"name"`
	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	Password *string `pulumi:"password"`
	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
}

type KafkaUserState struct {
	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	InstanceId pulumi.StringPtrInput
	// Specifies the name of the user. Changing this creates a new resource.
	Name pulumi.StringPtrInput
	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	Password pulumi.StringPtrInput
	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
}

func (KafkaUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaUserState)(nil)).Elem()
}

type kafkaUserArgs struct {
	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	InstanceId string `pulumi:"instanceId"`
	// Specifies the name of the user. Changing this creates a new resource.
	Name *string `pulumi:"name"`
	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	Password string `pulumi:"password"`
	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a KafkaUser resource.
type KafkaUserArgs struct {
	// Specifies the ID of the DMS kafka instance to which the user belongs.
	// Changing this creates a new resource.
	InstanceId pulumi.StringInput
	// Specifies the name of the user. Changing this creates a new resource.
	Name pulumi.StringPtrInput
	// Specifies the password of the user. The parameter must be 8 to 32 characters
	// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
	// The value must be different from name.
	Password pulumi.StringInput
	// The region in which to create the DMS kafka user resource. If omitted, the
	// provider-level region will be used. Changing this creates a new resource.
	Region pulumi.StringPtrInput
}

func (KafkaUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaUserArgs)(nil)).Elem()
}

type KafkaUserInput interface {
	pulumi.Input

	ToKafkaUserOutput() KafkaUserOutput
	ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput
}

func (*KafkaUser) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaUser)(nil)).Elem()
}

func (i *KafkaUser) ToKafkaUserOutput() KafkaUserOutput {
	return i.ToKafkaUserOutputWithContext(context.Background())
}

func (i *KafkaUser) ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaUserOutput)
}

// KafkaUserArrayInput is an input type that accepts KafkaUserArray and KafkaUserArrayOutput values.
// You can construct a concrete instance of `KafkaUserArrayInput` via:
//
//	KafkaUserArray{ KafkaUserArgs{...} }
type KafkaUserArrayInput interface {
	pulumi.Input

	ToKafkaUserArrayOutput() KafkaUserArrayOutput
	ToKafkaUserArrayOutputWithContext(context.Context) KafkaUserArrayOutput
}

type KafkaUserArray []KafkaUserInput

func (KafkaUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaUser)(nil)).Elem()
}

func (i KafkaUserArray) ToKafkaUserArrayOutput() KafkaUserArrayOutput {
	return i.ToKafkaUserArrayOutputWithContext(context.Background())
}

func (i KafkaUserArray) ToKafkaUserArrayOutputWithContext(ctx context.Context) KafkaUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaUserArrayOutput)
}

// KafkaUserMapInput is an input type that accepts KafkaUserMap and KafkaUserMapOutput values.
// You can construct a concrete instance of `KafkaUserMapInput` via:
//
//	KafkaUserMap{ "key": KafkaUserArgs{...} }
type KafkaUserMapInput interface {
	pulumi.Input

	ToKafkaUserMapOutput() KafkaUserMapOutput
	ToKafkaUserMapOutputWithContext(context.Context) KafkaUserMapOutput
}

type KafkaUserMap map[string]KafkaUserInput

func (KafkaUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaUser)(nil)).Elem()
}

func (i KafkaUserMap) ToKafkaUserMapOutput() KafkaUserMapOutput {
	return i.ToKafkaUserMapOutputWithContext(context.Background())
}

func (i KafkaUserMap) ToKafkaUserMapOutputWithContext(ctx context.Context) KafkaUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaUserMapOutput)
}

type KafkaUserOutput struct{ *pulumi.OutputState }

func (KafkaUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaUser)(nil)).Elem()
}

func (o KafkaUserOutput) ToKafkaUserOutput() KafkaUserOutput {
	return o
}

func (o KafkaUserOutput) ToKafkaUserOutputWithContext(ctx context.Context) KafkaUserOutput {
	return o
}

// Specifies the ID of the DMS kafka instance to which the user belongs.
// Changing this creates a new resource.
func (o KafkaUserOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Specifies the name of the user. Changing this creates a new resource.
func (o KafkaUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the password of the user. The parameter must be 8 to 32 characters
// long and contain only letters(case-sensitive), digits, and special characters(`~!@#$%^&*()-_=+|[{}]:'",<.>/?).
// The value must be different from name.
func (o KafkaUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The region in which to create the DMS kafka user resource. If omitted, the
// provider-level region will be used. Changing this creates a new resource.
func (o KafkaUserOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaUser) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type KafkaUserArrayOutput struct{ *pulumi.OutputState }

func (KafkaUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaUser)(nil)).Elem()
}

func (o KafkaUserArrayOutput) ToKafkaUserArrayOutput() KafkaUserArrayOutput {
	return o
}

func (o KafkaUserArrayOutput) ToKafkaUserArrayOutputWithContext(ctx context.Context) KafkaUserArrayOutput {
	return o
}

func (o KafkaUserArrayOutput) Index(i pulumi.IntInput) KafkaUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaUser {
		return vs[0].([]*KafkaUser)[vs[1].(int)]
	}).(KafkaUserOutput)
}

type KafkaUserMapOutput struct{ *pulumi.OutputState }

func (KafkaUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaUser)(nil)).Elem()
}

func (o KafkaUserMapOutput) ToKafkaUserMapOutput() KafkaUserMapOutput {
	return o
}

func (o KafkaUserMapOutput) ToKafkaUserMapOutputWithContext(ctx context.Context) KafkaUserMapOutput {
	return o
}

func (o KafkaUserMapOutput) MapIndex(k pulumi.StringInput) KafkaUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaUser {
		return vs[0].(map[string]*KafkaUser)[vs[1].(string)]
	}).(KafkaUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaUserInput)(nil)).Elem(), &KafkaUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaUserArrayInput)(nil)).Elem(), KafkaUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaUserMapInput)(nil)).Elem(), KafkaUserMap{})
	pulumi.RegisterOutputType(KafkaUserOutput{})
	pulumi.RegisterOutputType(KafkaUserArrayOutput{})
	pulumi.RegisterOutputType(KafkaUserMapOutput{})
}

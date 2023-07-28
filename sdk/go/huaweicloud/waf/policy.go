// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package waf

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a WAF policy resource within HuaweiCloud.
//
// > **NOTE:** All WAF resources depend on WAF instances, and the WAF instances need to be purchased before they can be
// used. The policy resource can be used in Cloud Mode, Dedicated Mode and ELB Mode.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Waf"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			enterpriseProjectId := cfg.RequireObject("enterpriseProjectId")
//			_, err := Waf.NewPolicy(ctx, "policy1", &Waf.PolicyArgs{
//				ProtectionMode:      pulumi.String("log"),
//				Level:               pulumi.Int(2),
//				EnterpriseProjectId: pulumi.Any(enterpriseProjectId),
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
// There are two ways to import WAF policy state. * Using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/policy:Policy test <id>
//
// ```
//
//   - Using `id` and `enterprise_project_id`, separated by a slash, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Waf/policy:Policy test <id>/<enterprise_project_id>
//
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Specifies the enterprise project ID of WAF policy.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrOutput `pulumi:"enterpriseProjectId"`
	// The detection mode in Precise Protection.
	// + `true`: full detection, Full detection finishes all threat detections before blocking requests that meet Precise
	//   Protection specified conditions.
	// + `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
	//   meets Precise Protection specified conditions.
	FullDetection pulumi.BoolOutput `pulumi:"fullDetection"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: low
	// + `2`: medium
	// + `3`: high
	Level pulumi.IntOutput `pulumi:"level"`
	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protection switches. The options object structure is documented below.
	Options PolicyOptionArrayOutput `pulumi:"options"`
	// Specifies the protective action after a rule is matched. Defaults to `log`.
	// Valid values are:
	// + `block`: WAF blocks and logs detected attacks.
	// + `log`: WAF logs detected attacks only.
	ProtectionMode pulumi.StringOutput `pulumi:"protectionMode"`
	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Policy
	err := ctx.RegisterResource("huaweicloud:Waf/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("huaweicloud:Waf/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Specifies the enterprise project ID of WAF policy.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The detection mode in Precise Protection.
	// + `true`: full detection, Full detection finishes all threat detections before blocking requests that meet Precise
	//   Protection specified conditions.
	// + `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
	//   meets Precise Protection specified conditions.
	FullDetection *bool `pulumi:"fullDetection"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: low
	// + `2`: medium
	// + `3`: high
	Level *int `pulumi:"level"`
	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	Name *string `pulumi:"name"`
	// The protection switches. The options object structure is documented below.
	Options []PolicyOption `pulumi:"options"`
	// Specifies the protective action after a rule is matched. Defaults to `log`.
	// Valid values are:
	// + `block`: WAF blocks and logs detected attacks.
	// + `log`: WAF logs detected attacks only.
	ProtectionMode *string `pulumi:"protectionMode"`
	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region *string `pulumi:"region"`
}

type PolicyState struct {
	// Specifies the enterprise project ID of WAF policy.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// The detection mode in Precise Protection.
	// + `true`: full detection, Full detection finishes all threat detections before blocking requests that meet Precise
	//   Protection specified conditions.
	// + `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
	//   meets Precise Protection specified conditions.
	FullDetection pulumi.BoolPtrInput
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: low
	// + `2`: medium
	// + `3`: high
	Level pulumi.IntPtrInput
	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringPtrInput
	// The protection switches. The options object structure is documented below.
	Options PolicyOptionArrayInput
	// Specifies the protective action after a rule is matched. Defaults to `log`.
	// Valid values are:
	// + `block`: WAF blocks and logs detected attacks.
	// + `log`: WAF logs detected attacks only.
	ProtectionMode pulumi.StringPtrInput
	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Specifies the enterprise project ID of WAF policy.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: low
	// + `2`: medium
	// + `3`: high
	Level *int `pulumi:"level"`
	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the protective action after a rule is matched. Defaults to `log`.
	// Valid values are:
	// + `block`: WAF blocks and logs detected attacks.
	// + `log`: WAF logs detected attacks only.
	ProtectionMode *string `pulumi:"protectionMode"`
	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Specifies the enterprise project ID of WAF policy.
	// Changing this parameter will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the protection level. Defaults to `2`. Valid values are:
	// + `1`: low
	// + `2`: medium
	// + `3`: high
	Level pulumi.IntPtrInput
	// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
	// underscores(_), and hyphens(-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the protective action after a rule is matched. Defaults to `log`.
	// Valid values are:
	// + `block`: WAF blocks and logs detected attacks.
	// + `log`: WAF logs detected attacks only.
	ProtectionMode pulumi.StringPtrInput
	// The region in which to create the WAF policy resource. If omitted, the
	// provider-level region will be used. Changing this setting will push a new certificate.
	Region pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//	PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//	PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct{ *pulumi.OutputState }

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil)).Elem()
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

// Specifies the enterprise project ID of WAF policy.
// Changing this parameter will create a new resource.
func (o PolicyOutput) EnterpriseProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringPtrOutput { return v.EnterpriseProjectId }).(pulumi.StringPtrOutput)
}

// The detection mode in Precise Protection.
//   - `true`: full detection, Full detection finishes all threat detections before blocking requests that meet Precise
//     Protection specified conditions.
//   - `false`: instant detection. Instant detection immediately ends threat detection after blocking a request that
//     meets Precise Protection specified conditions.
func (o PolicyOutput) FullDetection() pulumi.BoolOutput {
	return o.ApplyT(func(v *Policy) pulumi.BoolOutput { return v.FullDetection }).(pulumi.BoolOutput)
}

// Specifies the protection level. Defaults to `2`. Valid values are:
// + `1`: low
// + `2`: medium
// + `3`: high
func (o PolicyOutput) Level() pulumi.IntOutput {
	return o.ApplyT(func(v *Policy) pulumi.IntOutput { return v.Level }).(pulumi.IntOutput)
}

// Specifies the policy name. The maximum length is 256 characters. Only digits, letters,
// underscores(_), and hyphens(-) are allowed.
func (o PolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protection switches. The options object structure is documented below.
func (o PolicyOutput) Options() PolicyOptionArrayOutput {
	return o.ApplyT(func(v *Policy) PolicyOptionArrayOutput { return v.Options }).(PolicyOptionArrayOutput)
}

// Specifies the protective action after a rule is matched. Defaults to `log`.
// Valid values are:
// + `block`: WAF blocks and logs detected attacks.
// + `log`: WAF logs detected attacks only.
func (o PolicyOutput) ProtectionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.ProtectionMode }).(pulumi.StringOutput)
}

// The region in which to create the WAF policy resource. If omitted, the
// provider-level region will be used. Changing this setting will push a new certificate.
func (o PolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Policy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Policy)(nil)).Elem()
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].([]*Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Policy)(nil)).Elem()
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Policy {
		return vs[0].(map[string]*Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyInput)(nil)).Elem(), &Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyArrayInput)(nil)).Elem(), PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyMapInput)(nil)).Elem(), PolicyMap{})
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}

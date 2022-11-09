// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iotda

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an IoTDA data forwarding rule within HuaweiCloud.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/IoTDA"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/IoTDA"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			disRegion := cfg.RequireObject("disRegion")
//			disStreamId := cfg.RequireObject("disStreamId")
//			obsRegion := cfg.RequireObject("obsRegion")
//			obsbucket := cfg.RequireObject("obsbucket")
//			_, err := IoTDA.NewDataforwardingRule(ctx, "test", &IoTDA.DataforwardingRuleArgs{
//				Trigger: pulumi.String("product:create"),
//				Enabled: pulumi.Bool(true),
//				Targets: iotda.DataforwardingRuleTargetArray{
//					&iotda.DataforwardingRuleTargetArgs{
//						Type: pulumi.String("DIS_FORWARDING"),
//						DisForwarding: &iotda.DataforwardingRuleTargetDisForwardingArgs{
//							Region:   pulumi.Any(disRegion),
//							StreamId: pulumi.Any(disStreamId),
//						},
//					},
//					&iotda.DataforwardingRuleTargetArgs{
//						Type: pulumi.String("HTTP_FORWARDING"),
//						HttpForwarding: &iotda.DataforwardingRuleTargetHttpForwardingArgs{
//							Url: pulumi.String("http://www.yourDomain.com"),
//						},
//					},
//					&iotda.DataforwardingRuleTargetArgs{
//						Type: pulumi.String("OBS_FORWARDING"),
//						ObsForwarding: &iotda.DataforwardingRuleTargetObsForwardingArgs{
//							Region: pulumi.Any(obsRegion),
//							Bucket: pulumi.Any(obsbucket),
//						},
//					},
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
// Data forwarding rules can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:IoTDA/dataforwardingRule:DataforwardingRule test 10022532f4f94f26b01daa1e424853e1
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attrubutes missing from the API response, security or some other reason. The missing attributes include`password` of `kafka_forwarding`. It is generally recommended running `terraform plan` after importing the resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the group. Also you can ignore changes as below. resource "huaweicloud_iotda_device_group" "test" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	targets,
//
//	]
//
//	} }
type DataforwardingRule struct {
	pulumi.CustomResourceState

	// Specifies the description of data forwarding rule. The description contains
	// a maximum of 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies whether to enable the data forwarding rule. Defaults to `false`.
	// Can not enable without `targets`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
	// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
	// allowed: `?'#().,&%@!`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region to which the KAFKA belongs.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
	Select pulumi.StringPtrOutput `pulumi:"select"`
	// Specifies the resource space ID which uses the data forwarding rule.
	// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
	SpaceId pulumi.StringOutput `pulumi:"spaceId"`
	// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
	// want to forward the data. The targets structure is documented below.
	Targets DataforwardingRuleTargetArrayOutput `pulumi:"targets"`
	// Specifies the trigger event. The options are as follows:
	// + **device:create**: Device added.
	// + **device:delete**: Device deleted.
	// + **device:update**: Device updated.
	// + **device.status:update**: Device status changed.
	// + **device.property:report**: Device property reported.
	// + **device.message:report**: Device message reported.
	// + **device.message.status:update**: Device message status changed.
	// + **batchtask:update**: Batch task status changed.
	// + **product:create**: Product added.
	// + **product:delete**: Product deleted.
	// + **product:update**: Product updated.
	// + **device.command.status:update**: Update of the device asynchronous command status.
	Trigger pulumi.StringOutput `pulumi:"trigger"`
	// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
	Where pulumi.StringPtrOutput `pulumi:"where"`
}

// NewDataforwardingRule registers a new resource with the given unique name, arguments, and options.
func NewDataforwardingRule(ctx *pulumi.Context,
	name string, args *DataforwardingRuleArgs, opts ...pulumi.ResourceOption) (*DataforwardingRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Trigger == nil {
		return nil, errors.New("invalid value for required argument 'Trigger'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DataforwardingRule
	err := ctx.RegisterResource("huaweicloud:IoTDA/dataforwardingRule:DataforwardingRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataforwardingRule gets an existing DataforwardingRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataforwardingRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataforwardingRuleState, opts ...pulumi.ResourceOption) (*DataforwardingRule, error) {
	var resource DataforwardingRule
	err := ctx.ReadResource("huaweicloud:IoTDA/dataforwardingRule:DataforwardingRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataforwardingRule resources.
type dataforwardingRuleState struct {
	// Specifies the description of data forwarding rule. The description contains
	// a maximum of 256 characters.
	Description *string `pulumi:"description"`
	// Specifies whether to enable the data forwarding rule. Defaults to `false`.
	// Can not enable without `targets`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
	// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
	// allowed: `?'#().,&%@!`.
	Name *string `pulumi:"name"`
	// Specifies the region to which the KAFKA belongs.
	Region *string `pulumi:"region"`
	// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
	Select *string `pulumi:"select"`
	// Specifies the resource space ID which uses the data forwarding rule.
	// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
	SpaceId *string `pulumi:"spaceId"`
	// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
	// want to forward the data. The targets structure is documented below.
	Targets []DataforwardingRuleTarget `pulumi:"targets"`
	// Specifies the trigger event. The options are as follows:
	// + **device:create**: Device added.
	// + **device:delete**: Device deleted.
	// + **device:update**: Device updated.
	// + **device.status:update**: Device status changed.
	// + **device.property:report**: Device property reported.
	// + **device.message:report**: Device message reported.
	// + **device.message.status:update**: Device message status changed.
	// + **batchtask:update**: Batch task status changed.
	// + **product:create**: Product added.
	// + **product:delete**: Product deleted.
	// + **product:update**: Product updated.
	// + **device.command.status:update**: Update of the device asynchronous command status.
	Trigger *string `pulumi:"trigger"`
	// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
	Where *string `pulumi:"where"`
}

type DataforwardingRuleState struct {
	// Specifies the description of data forwarding rule. The description contains
	// a maximum of 256 characters.
	Description pulumi.StringPtrInput
	// Specifies whether to enable the data forwarding rule. Defaults to `false`.
	// Can not enable without `targets`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
	// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
	// allowed: `?'#().,&%@!`.
	Name pulumi.StringPtrInput
	// Specifies the region to which the KAFKA belongs.
	Region pulumi.StringPtrInput
	// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
	Select pulumi.StringPtrInput
	// Specifies the resource space ID which uses the data forwarding rule.
	// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
	SpaceId pulumi.StringPtrInput
	// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
	// want to forward the data. The targets structure is documented below.
	Targets DataforwardingRuleTargetArrayInput
	// Specifies the trigger event. The options are as follows:
	// + **device:create**: Device added.
	// + **device:delete**: Device deleted.
	// + **device:update**: Device updated.
	// + **device.status:update**: Device status changed.
	// + **device.property:report**: Device property reported.
	// + **device.message:report**: Device message reported.
	// + **device.message.status:update**: Device message status changed.
	// + **batchtask:update**: Batch task status changed.
	// + **product:create**: Product added.
	// + **product:delete**: Product deleted.
	// + **product:update**: Product updated.
	// + **device.command.status:update**: Update of the device asynchronous command status.
	Trigger pulumi.StringPtrInput
	// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
	Where pulumi.StringPtrInput
}

func (DataforwardingRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataforwardingRuleState)(nil)).Elem()
}

type dataforwardingRuleArgs struct {
	// Specifies the description of data forwarding rule. The description contains
	// a maximum of 256 characters.
	Description *string `pulumi:"description"`
	// Specifies whether to enable the data forwarding rule. Defaults to `false`.
	// Can not enable without `targets`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
	// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
	// allowed: `?'#().,&%@!`.
	Name *string `pulumi:"name"`
	// Specifies the region to which the KAFKA belongs.
	Region *string `pulumi:"region"`
	// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
	Select *string `pulumi:"select"`
	// Specifies the resource space ID which uses the data forwarding rule.
	// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
	SpaceId *string `pulumi:"spaceId"`
	// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
	// want to forward the data. The targets structure is documented below.
	Targets []DataforwardingRuleTarget `pulumi:"targets"`
	// Specifies the trigger event. The options are as follows:
	// + **device:create**: Device added.
	// + **device:delete**: Device deleted.
	// + **device:update**: Device updated.
	// + **device.status:update**: Device status changed.
	// + **device.property:report**: Device property reported.
	// + **device.message:report**: Device message reported.
	// + **device.message.status:update**: Device message status changed.
	// + **batchtask:update**: Batch task status changed.
	// + **product:create**: Product added.
	// + **product:delete**: Product deleted.
	// + **product:update**: Product updated.
	// + **device.command.status:update**: Update of the device asynchronous command status.
	Trigger string `pulumi:"trigger"`
	// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
	Where *string `pulumi:"where"`
}

// The set of arguments for constructing a DataforwardingRule resource.
type DataforwardingRuleArgs struct {
	// Specifies the description of data forwarding rule. The description contains
	// a maximum of 256 characters.
	Description pulumi.StringPtrInput
	// Specifies whether to enable the data forwarding rule. Defaults to `false`.
	// Can not enable without `targets`.
	Enabled pulumi.BoolPtrInput
	// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
	// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
	// allowed: `?'#().,&%@!`.
	Name pulumi.StringPtrInput
	// Specifies the region to which the KAFKA belongs.
	Region pulumi.StringPtrInput
	// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
	Select pulumi.StringPtrInput
	// Specifies the resource space ID which uses the data forwarding rule.
	// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
	SpaceId pulumi.StringPtrInput
	// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
	// want to forward the data. The targets structure is documented below.
	Targets DataforwardingRuleTargetArrayInput
	// Specifies the trigger event. The options are as follows:
	// + **device:create**: Device added.
	// + **device:delete**: Device deleted.
	// + **device:update**: Device updated.
	// + **device.status:update**: Device status changed.
	// + **device.property:report**: Device property reported.
	// + **device.message:report**: Device message reported.
	// + **device.message.status:update**: Device message status changed.
	// + **batchtask:update**: Batch task status changed.
	// + **product:create**: Product added.
	// + **product:delete**: Product deleted.
	// + **product:update**: Product updated.
	// + **device.command.status:update**: Update of the device asynchronous command status.
	Trigger pulumi.StringInput
	// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
	Where pulumi.StringPtrInput
}

func (DataforwardingRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataforwardingRuleArgs)(nil)).Elem()
}

type DataforwardingRuleInput interface {
	pulumi.Input

	ToDataforwardingRuleOutput() DataforwardingRuleOutput
	ToDataforwardingRuleOutputWithContext(ctx context.Context) DataforwardingRuleOutput
}

func (*DataforwardingRule) ElementType() reflect.Type {
	return reflect.TypeOf((**DataforwardingRule)(nil)).Elem()
}

func (i *DataforwardingRule) ToDataforwardingRuleOutput() DataforwardingRuleOutput {
	return i.ToDataforwardingRuleOutputWithContext(context.Background())
}

func (i *DataforwardingRule) ToDataforwardingRuleOutputWithContext(ctx context.Context) DataforwardingRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataforwardingRuleOutput)
}

// DataforwardingRuleArrayInput is an input type that accepts DataforwardingRuleArray and DataforwardingRuleArrayOutput values.
// You can construct a concrete instance of `DataforwardingRuleArrayInput` via:
//
//	DataforwardingRuleArray{ DataforwardingRuleArgs{...} }
type DataforwardingRuleArrayInput interface {
	pulumi.Input

	ToDataforwardingRuleArrayOutput() DataforwardingRuleArrayOutput
	ToDataforwardingRuleArrayOutputWithContext(context.Context) DataforwardingRuleArrayOutput
}

type DataforwardingRuleArray []DataforwardingRuleInput

func (DataforwardingRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataforwardingRule)(nil)).Elem()
}

func (i DataforwardingRuleArray) ToDataforwardingRuleArrayOutput() DataforwardingRuleArrayOutput {
	return i.ToDataforwardingRuleArrayOutputWithContext(context.Background())
}

func (i DataforwardingRuleArray) ToDataforwardingRuleArrayOutputWithContext(ctx context.Context) DataforwardingRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataforwardingRuleArrayOutput)
}

// DataforwardingRuleMapInput is an input type that accepts DataforwardingRuleMap and DataforwardingRuleMapOutput values.
// You can construct a concrete instance of `DataforwardingRuleMapInput` via:
//
//	DataforwardingRuleMap{ "key": DataforwardingRuleArgs{...} }
type DataforwardingRuleMapInput interface {
	pulumi.Input

	ToDataforwardingRuleMapOutput() DataforwardingRuleMapOutput
	ToDataforwardingRuleMapOutputWithContext(context.Context) DataforwardingRuleMapOutput
}

type DataforwardingRuleMap map[string]DataforwardingRuleInput

func (DataforwardingRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataforwardingRule)(nil)).Elem()
}

func (i DataforwardingRuleMap) ToDataforwardingRuleMapOutput() DataforwardingRuleMapOutput {
	return i.ToDataforwardingRuleMapOutputWithContext(context.Background())
}

func (i DataforwardingRuleMap) ToDataforwardingRuleMapOutputWithContext(ctx context.Context) DataforwardingRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataforwardingRuleMapOutput)
}

type DataforwardingRuleOutput struct{ *pulumi.OutputState }

func (DataforwardingRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataforwardingRule)(nil)).Elem()
}

func (o DataforwardingRuleOutput) ToDataforwardingRuleOutput() DataforwardingRuleOutput {
	return o
}

func (o DataforwardingRuleOutput) ToDataforwardingRuleOutputWithContext(ctx context.Context) DataforwardingRuleOutput {
	return o
}

// Specifies the description of data forwarding rule. The description contains
// a maximum of 256 characters.
func (o DataforwardingRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies whether to enable the data forwarding rule. Defaults to `false`.
// Can not enable without `targets`.
func (o DataforwardingRuleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Specifies the name of data forwarding rule. The name contains a maximum of 256 characters.
// Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following specail characters are
// allowed: `?'#().,&%@!`.
func (o DataforwardingRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region to which the KAFKA belongs.
func (o DataforwardingRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the SQL SELECT statement which contains a maximum of 500 characters.
func (o DataforwardingRuleOutput) Select() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringPtrOutput { return v.Select }).(pulumi.StringPtrOutput)
}

// Specifies the resource space ID which uses the data forwarding rule.
// If omitted, all resource space will use the data forwarding rule. Changing this parameter will create a new resource.
func (o DataforwardingRuleOutput) SpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringOutput { return v.SpaceId }).(pulumi.StringOutput)
}

// Specifies the list of the targets (HUAWEI CLOUD services or private servers) to which you
// want to forward the data. The targets structure is documented below.
func (o DataforwardingRuleOutput) Targets() DataforwardingRuleTargetArrayOutput {
	return o.ApplyT(func(v *DataforwardingRule) DataforwardingRuleTargetArrayOutput { return v.Targets }).(DataforwardingRuleTargetArrayOutput)
}

// Specifies the trigger event. The options are as follows:
// + **device:create**: Device added.
// + **device:delete**: Device deleted.
// + **device:update**: Device updated.
// + **device.status:update**: Device status changed.
// + **device.property:report**: Device property reported.
// + **device.message:report**: Device message reported.
// + **device.message.status:update**: Device message status changed.
// + **batchtask:update**: Batch task status changed.
// + **product:create**: Product added.
// + **product:delete**: Product deleted.
// + **product:update**: Product updated.
// + **device.command.status:update**: Update of the device asynchronous command status.
func (o DataforwardingRuleOutput) Trigger() pulumi.StringOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringOutput { return v.Trigger }).(pulumi.StringOutput)
}

// Specifies the SQL WHERE statement which contains a maximum of 500 characters.
func (o DataforwardingRuleOutput) Where() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataforwardingRule) pulumi.StringPtrOutput { return v.Where }).(pulumi.StringPtrOutput)
}

type DataforwardingRuleArrayOutput struct{ *pulumi.OutputState }

func (DataforwardingRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataforwardingRule)(nil)).Elem()
}

func (o DataforwardingRuleArrayOutput) ToDataforwardingRuleArrayOutput() DataforwardingRuleArrayOutput {
	return o
}

func (o DataforwardingRuleArrayOutput) ToDataforwardingRuleArrayOutputWithContext(ctx context.Context) DataforwardingRuleArrayOutput {
	return o
}

func (o DataforwardingRuleArrayOutput) Index(i pulumi.IntInput) DataforwardingRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataforwardingRule {
		return vs[0].([]*DataforwardingRule)[vs[1].(int)]
	}).(DataforwardingRuleOutput)
}

type DataforwardingRuleMapOutput struct{ *pulumi.OutputState }

func (DataforwardingRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataforwardingRule)(nil)).Elem()
}

func (o DataforwardingRuleMapOutput) ToDataforwardingRuleMapOutput() DataforwardingRuleMapOutput {
	return o
}

func (o DataforwardingRuleMapOutput) ToDataforwardingRuleMapOutputWithContext(ctx context.Context) DataforwardingRuleMapOutput {
	return o
}

func (o DataforwardingRuleMapOutput) MapIndex(k pulumi.StringInput) DataforwardingRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataforwardingRule {
		return vs[0].(map[string]*DataforwardingRule)[vs[1].(string)]
	}).(DataforwardingRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataforwardingRuleInput)(nil)).Elem(), &DataforwardingRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataforwardingRuleArrayInput)(nil)).Elem(), DataforwardingRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataforwardingRuleMapInput)(nil)).Elem(), DataforwardingRuleMap{})
	pulumi.RegisterOutputType(DataforwardingRuleOutput{})
	pulumi.RegisterOutputType(DataforwardingRuleArrayOutput{})
	pulumi.RegisterOutputType(DataforwardingRuleMapOutput{})
}

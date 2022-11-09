// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataartsstudio

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages DataArts Studio instance resource within HuaweiCloud.
//
// > Only **prePaid** charging mode is supported.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/DataArtsStudio"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			availabilityZone := cfg.RequireObject("availabilityZone")
//			vpcId := cfg.RequireObject("vpcId")
//			subnetId := cfg.RequireObject("subnetId")
//			secgroupId := cfg.RequireObject("secgroupId")
//			_, err := DataArtsStudio.NewStudioInstance(ctx, "myDemo", &DataArtsStudio.StudioInstanceArgs{
//				Version:             pulumi.String("dayu.starter"),
//				VpcId:               pulumi.Any(vpcId),
//				SubnetId:            pulumi.Any(subnetId),
//				SecurityGroupId:     pulumi.Any(secgroupId),
//				AvailabilityZone:    pulumi.Any(availabilityZone),
//				PeriodUnit:          pulumi.String("month"),
//				Period:              pulumi.Int(1),
//				EnterpriseProjectId: pulumi.String("0"),
//				Tags: pulumi.StringMap{
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
// <!--markdownlint-disable MD033-->
//
// ## Import
//
// DataArts Studio instances can be imported using their `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:DataArtsStudio/studioInstance:StudioInstance instance e60361de2cfd42d7a6b673f0ae58db82
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`tags`, `period_unit`, `period`, `auto_renew`. It is generally recommended running `terraform plan` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. resource "huaweicloud_dataarts_studio_instance" "instance" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	tags, period_unit, period, auto_renew,
//
//	]
//
//	} }
type StudioInstance struct {
	pulumi.CustomResourceState

	// Specifies whether auto renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrOutput `pulumi:"autoRenew"`
	// Specifies the AZ name. Changing this creates a new instance.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// The charging mode. The value is `prePaid` indicates the yearly/monthly billing mode.
	ChargingMode pulumi.StringOutput `pulumi:"chargingMode"`
	// Specifies the enterprise project id of the instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// The expire days to renew.
	ExpireDays pulumi.StringOutput `pulumi:"expireDays"`
	// Specifies the DataArts Studio instance name. Changing this creates a new instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The order ID of this DataArts Studio instance.
	OrderId pulumi.StringOutput `pulumi:"orderId"`
	// Specifies the charging period of the DataArts Studio instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// Changing this creates a new instance.
	Period pulumi.IntOutput `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringOutput `pulumi:"periodUnit"`
	// Specifies the region in which to manage the DataArts Studio instance.
	// Changing this creates a new instance.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the security group ID. Changing this creates a new instance.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// The status of this DataArts Studio instance.
	Status pulumi.IntOutput `pulumi:"status"`
	// Specifies the VPC subnet ID. Changing this creates a new instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The key/value pairs to associate with the DataArts Studio instance.
	// Changing this creates a new instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the DataArts Studio version version.
	// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
	// Changing this creates a new instance.
	Version pulumi.StringOutput `pulumi:"version"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewStudioInstance registers a new resource with the given unique name, arguments, and options.
func NewStudioInstance(ctx *pulumi.Context,
	name string, args *StudioInstanceArgs, opts ...pulumi.ResourceOption) (*StudioInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.Period == nil {
		return nil, errors.New("invalid value for required argument 'Period'")
	}
	if args.PeriodUnit == nil {
		return nil, errors.New("invalid value for required argument 'PeriodUnit'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource StudioInstance
	err := ctx.RegisterResource("huaweicloud:DataArtsStudio/studioInstance:StudioInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStudioInstance gets an existing StudioInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStudioInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StudioInstanceState, opts ...pulumi.ResourceOption) (*StudioInstance, error) {
	var resource StudioInstance
	err := ctx.ReadResource("huaweicloud:DataArtsStudio/studioInstance:StudioInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StudioInstance resources.
type studioInstanceState struct {
	// Specifies whether auto renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the AZ name. Changing this creates a new instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// The charging mode. The value is `prePaid` indicates the yearly/monthly billing mode.
	ChargingMode *string `pulumi:"chargingMode"`
	// Specifies the enterprise project id of the instance.
	// Changing this creates a new instance.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// The expire days to renew.
	ExpireDays *string `pulumi:"expireDays"`
	// Specifies the DataArts Studio instance name. Changing this creates a new instance.
	Name *string `pulumi:"name"`
	// The order ID of this DataArts Studio instance.
	OrderId *string `pulumi:"orderId"`
	// Specifies the charging period of the DataArts Studio instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// Changing this creates a new instance.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*.
	// Changing this creates a new instance.
	PeriodUnit *string `pulumi:"periodUnit"`
	// Specifies the region in which to manage the DataArts Studio instance.
	// Changing this creates a new instance.
	Region *string `pulumi:"region"`
	// Specifies the security group ID. Changing this creates a new instance.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The status of this DataArts Studio instance.
	Status *int `pulumi:"status"`
	// Specifies the VPC subnet ID. Changing this creates a new instance.
	SubnetId *string `pulumi:"subnetId"`
	// The key/value pairs to associate with the DataArts Studio instance.
	// Changing this creates a new instance.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the DataArts Studio version version.
	// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
	// Changing this creates a new instance.
	Version *string `pulumi:"version"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId *string `pulumi:"vpcId"`
}

type StudioInstanceState struct {
	// Specifies whether auto renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrInput
	// Specifies the AZ name. Changing this creates a new instance.
	AvailabilityZone pulumi.StringPtrInput
	// The charging mode. The value is `prePaid` indicates the yearly/monthly billing mode.
	ChargingMode pulumi.StringPtrInput
	// Specifies the enterprise project id of the instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringPtrInput
	// The expire days to renew.
	ExpireDays pulumi.StringPtrInput
	// Specifies the DataArts Studio instance name. Changing this creates a new instance.
	Name pulumi.StringPtrInput
	// The order ID of this DataArts Studio instance.
	OrderId pulumi.StringPtrInput
	// Specifies the charging period of the DataArts Studio instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// Changing this creates a new instance.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringPtrInput
	// Specifies the region in which to manage the DataArts Studio instance.
	// Changing this creates a new instance.
	Region pulumi.StringPtrInput
	// Specifies the security group ID. Changing this creates a new instance.
	SecurityGroupId pulumi.StringPtrInput
	// The status of this DataArts Studio instance.
	Status pulumi.IntPtrInput
	// Specifies the VPC subnet ID. Changing this creates a new instance.
	SubnetId pulumi.StringPtrInput
	// The key/value pairs to associate with the DataArts Studio instance.
	// Changing this creates a new instance.
	Tags pulumi.StringMapInput
	// Specifies the DataArts Studio version version.
	// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
	// Changing this creates a new instance.
	Version pulumi.StringPtrInput
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringPtrInput
}

func (StudioInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*studioInstanceState)(nil)).Elem()
}

type studioInstanceArgs struct {
	// Specifies whether auto renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the AZ name. Changing this creates a new instance.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies the enterprise project id of the instance.
	// Changing this creates a new instance.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the DataArts Studio instance name. Changing this creates a new instance.
	Name *string `pulumi:"name"`
	// Specifies the charging period of the DataArts Studio instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// Changing this creates a new instance.
	Period int `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*.
	// Changing this creates a new instance.
	PeriodUnit string `pulumi:"periodUnit"`
	// Specifies the region in which to manage the DataArts Studio instance.
	// Changing this creates a new instance.
	Region *string `pulumi:"region"`
	// Specifies the security group ID. Changing this creates a new instance.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Specifies the VPC subnet ID. Changing this creates a new instance.
	SubnetId string `pulumi:"subnetId"`
	// The key/value pairs to associate with the DataArts Studio instance.
	// Changing this creates a new instance.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the DataArts Studio version version.
	// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
	// Changing this creates a new instance.
	Version string `pulumi:"version"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a StudioInstance resource.
type StudioInstanceArgs struct {
	// Specifies whether auto renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrInput
	// Specifies the AZ name. Changing this creates a new instance.
	AvailabilityZone pulumi.StringInput
	// Specifies the enterprise project id of the instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the DataArts Studio instance name. Changing this creates a new instance.
	Name pulumi.StringPtrInput
	// Specifies the charging period of the DataArts Studio instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// Changing this creates a new instance.
	Period pulumi.IntInput
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringInput
	// Specifies the region in which to manage the DataArts Studio instance.
	// Changing this creates a new instance.
	Region pulumi.StringPtrInput
	// Specifies the security group ID. Changing this creates a new instance.
	SecurityGroupId pulumi.StringInput
	// Specifies the VPC subnet ID. Changing this creates a new instance.
	SubnetId pulumi.StringInput
	// The key/value pairs to associate with the DataArts Studio instance.
	// Changing this creates a new instance.
	Tags pulumi.StringMapInput
	// Specifies the DataArts Studio version version.
	// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
	// Changing this creates a new instance.
	Version pulumi.StringInput
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringInput
}

func (StudioInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*studioInstanceArgs)(nil)).Elem()
}

type StudioInstanceInput interface {
	pulumi.Input

	ToStudioInstanceOutput() StudioInstanceOutput
	ToStudioInstanceOutputWithContext(ctx context.Context) StudioInstanceOutput
}

func (*StudioInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioInstance)(nil)).Elem()
}

func (i *StudioInstance) ToStudioInstanceOutput() StudioInstanceOutput {
	return i.ToStudioInstanceOutputWithContext(context.Background())
}

func (i *StudioInstance) ToStudioInstanceOutputWithContext(ctx context.Context) StudioInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioInstanceOutput)
}

// StudioInstanceArrayInput is an input type that accepts StudioInstanceArray and StudioInstanceArrayOutput values.
// You can construct a concrete instance of `StudioInstanceArrayInput` via:
//
//	StudioInstanceArray{ StudioInstanceArgs{...} }
type StudioInstanceArrayInput interface {
	pulumi.Input

	ToStudioInstanceArrayOutput() StudioInstanceArrayOutput
	ToStudioInstanceArrayOutputWithContext(context.Context) StudioInstanceArrayOutput
}

type StudioInstanceArray []StudioInstanceInput

func (StudioInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StudioInstance)(nil)).Elem()
}

func (i StudioInstanceArray) ToStudioInstanceArrayOutput() StudioInstanceArrayOutput {
	return i.ToStudioInstanceArrayOutputWithContext(context.Background())
}

func (i StudioInstanceArray) ToStudioInstanceArrayOutputWithContext(ctx context.Context) StudioInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioInstanceArrayOutput)
}

// StudioInstanceMapInput is an input type that accepts StudioInstanceMap and StudioInstanceMapOutput values.
// You can construct a concrete instance of `StudioInstanceMapInput` via:
//
//	StudioInstanceMap{ "key": StudioInstanceArgs{...} }
type StudioInstanceMapInput interface {
	pulumi.Input

	ToStudioInstanceMapOutput() StudioInstanceMapOutput
	ToStudioInstanceMapOutputWithContext(context.Context) StudioInstanceMapOutput
}

type StudioInstanceMap map[string]StudioInstanceInput

func (StudioInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StudioInstance)(nil)).Elem()
}

func (i StudioInstanceMap) ToStudioInstanceMapOutput() StudioInstanceMapOutput {
	return i.ToStudioInstanceMapOutputWithContext(context.Background())
}

func (i StudioInstanceMap) ToStudioInstanceMapOutputWithContext(ctx context.Context) StudioInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StudioInstanceMapOutput)
}

type StudioInstanceOutput struct{ *pulumi.OutputState }

func (StudioInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StudioInstance)(nil)).Elem()
}

func (o StudioInstanceOutput) ToStudioInstanceOutput() StudioInstanceOutput {
	return o
}

func (o StudioInstanceOutput) ToStudioInstanceOutputWithContext(ctx context.Context) StudioInstanceOutput {
	return o
}

// Specifies whether auto renew is enabled.
// Valid values are `true` and `false`, defaults to `false`.
// Changing this creates a new instance.
func (o StudioInstanceOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

// Specifies the AZ name. Changing this creates a new instance.
func (o StudioInstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// The charging mode. The value is `prePaid` indicates the yearly/monthly billing mode.
func (o StudioInstanceOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

// Specifies the enterprise project id of the instance.
// Changing this creates a new instance.
func (o StudioInstanceOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// The expire days to renew.
func (o StudioInstanceOutput) ExpireDays() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.ExpireDays }).(pulumi.StringOutput)
}

// Specifies the DataArts Studio instance name. Changing this creates a new instance.
func (o StudioInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The order ID of this DataArts Studio instance.
func (o StudioInstanceOutput) OrderId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.OrderId }).(pulumi.StringOutput)
}

// Specifies the charging period of the DataArts Studio instance.
// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
// Changing this creates a new instance.
func (o StudioInstanceOutput) Period() pulumi.IntOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.IntOutput { return v.Period }).(pulumi.IntOutput)
}

// Specifies the charging period unit of the instance.
// Valid values are *month* and *year*.
// Changing this creates a new instance.
func (o StudioInstanceOutput) PeriodUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.PeriodUnit }).(pulumi.StringOutput)
}

// Specifies the region in which to manage the DataArts Studio instance.
// Changing this creates a new instance.
func (o StudioInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the security group ID. Changing this creates a new instance.
func (o StudioInstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// The status of this DataArts Studio instance.
func (o StudioInstanceOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Specifies the VPC subnet ID. Changing this creates a new instance.
func (o StudioInstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The key/value pairs to associate with the DataArts Studio instance.
// Changing this creates a new instance.
func (o StudioInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the DataArts Studio version version.
// The valid values are **dayu.starter**, **dayu.nb.professional** and **dayu.nb.enterprise**.
// Changing this creates a new instance.
func (o StudioInstanceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Specifies the VPC ID. Changing this creates a new instance.
func (o StudioInstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *StudioInstance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type StudioInstanceArrayOutput struct{ *pulumi.OutputState }

func (StudioInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StudioInstance)(nil)).Elem()
}

func (o StudioInstanceArrayOutput) ToStudioInstanceArrayOutput() StudioInstanceArrayOutput {
	return o
}

func (o StudioInstanceArrayOutput) ToStudioInstanceArrayOutputWithContext(ctx context.Context) StudioInstanceArrayOutput {
	return o
}

func (o StudioInstanceArrayOutput) Index(i pulumi.IntInput) StudioInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StudioInstance {
		return vs[0].([]*StudioInstance)[vs[1].(int)]
	}).(StudioInstanceOutput)
}

type StudioInstanceMapOutput struct{ *pulumi.OutputState }

func (StudioInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StudioInstance)(nil)).Elem()
}

func (o StudioInstanceMapOutput) ToStudioInstanceMapOutput() StudioInstanceMapOutput {
	return o
}

func (o StudioInstanceMapOutput) ToStudioInstanceMapOutputWithContext(ctx context.Context) StudioInstanceMapOutput {
	return o
}

func (o StudioInstanceMapOutput) MapIndex(k pulumi.StringInput) StudioInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StudioInstance {
		return vs[0].(map[string]*StudioInstance)[vs[1].(string)]
	}).(StudioInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StudioInstanceInput)(nil)).Elem(), &StudioInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*StudioInstanceArrayInput)(nil)).Elem(), StudioInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StudioInstanceMapInput)(nil)).Elem(), StudioInstanceMap{})
	pulumi.RegisterOutputType(StudioInstanceOutput{})
	pulumi.RegisterOutputType(StudioInstanceArrayOutput{})
	pulumi.RegisterOutputType(StudioInstanceMapOutput{})
}

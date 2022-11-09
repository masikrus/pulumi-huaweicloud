// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package servicestage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an application resource within HuaweiCloud ServiceStage.
//
// ## Example Usage
// ### Create an application and an environment variable
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/ServiceStage"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/ServiceStage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			envId := cfg.RequireObject("envId")
//			appName := cfg.RequireObject("appName")
//			vpcId := cfg.RequireObject("vpcId")
//			_, err := ServiceStage.NewApplication(ctx, "test", &ServiceStage.ApplicationArgs{
//				Environments: servicestage.ApplicationEnvironmentArray{
//					&servicestage.ApplicationEnvironmentArgs{
//						Id: pulumi.Any(envId),
//						Variables: servicestage.ApplicationEnvironmentVariableArray{
//							&servicestage.ApplicationEnvironmentVariableArgs{
//								Name:  pulumi.String("owner"),
//								Value: pulumi.String("terraform"),
//							},
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
// Applications can be imported using their `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:ServiceStage/application:Application test eeea08e7-c838-4794-926c-abc12f3e10e8
//
// ```
type Application struct {
	pulumi.CustomResourceState

	// The list of component IDs associated under the application.
	ComponentIds pulumi.StringArrayOutput `pulumi:"componentIds"`
	// Specifies the application description.
	// The description can contian a maximum of `128` characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies the enterprise projcet ID to which the application
	// belongs. Changing this will create a new resource.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Specifies the configurations of the environment variables.
	// The object structure is documented below.
	Environments ApplicationEnvironmentArrayOutput `pulumi:"environments"`
	// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
	// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region in which to create the ServiceStage application.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
}

// NewApplication registers a new resource with the given unique name, arguments, and options.
func NewApplication(ctx *pulumi.Context,
	name string, args *ApplicationArgs, opts ...pulumi.ResourceOption) (*Application, error) {
	if args == nil {
		args = &ApplicationArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Application
	err := ctx.RegisterResource("huaweicloud:ServiceStage/application:Application", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApplication gets an existing Application resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApplication(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationState, opts ...pulumi.ResourceOption) (*Application, error) {
	var resource Application
	err := ctx.ReadResource("huaweicloud:ServiceStage/application:Application", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Application resources.
type applicationState struct {
	// The list of component IDs associated under the application.
	ComponentIds []string `pulumi:"componentIds"`
	// Specifies the application description.
	// The description can contian a maximum of `128` characters.
	Description *string `pulumi:"description"`
	// Specifies the enterprise projcet ID to which the application
	// belongs. Changing this will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the configurations of the environment variables.
	// The object structure is documented below.
	Environments []ApplicationEnvironment `pulumi:"environments"`
	// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
	// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the ServiceStage application.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
}

type ApplicationState struct {
	// The list of component IDs associated under the application.
	ComponentIds pulumi.StringArrayInput
	// Specifies the application description.
	// The description can contian a maximum of `128` characters.
	Description pulumi.StringPtrInput
	// Specifies the enterprise projcet ID to which the application
	// belongs. Changing this will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the configurations of the environment variables.
	// The object structure is documented below.
	Environments ApplicationEnvironmentArrayInput
	// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
	// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the ServiceStage application.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
}

func (ApplicationState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationState)(nil)).Elem()
}

type applicationArgs struct {
	// Specifies the application description.
	// The description can contian a maximum of `128` characters.
	Description *string `pulumi:"description"`
	// Specifies the enterprise projcet ID to which the application
	// belongs. Changing this will create a new resource.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the configurations of the environment variables.
	// The object structure is documented below.
	Environments []ApplicationEnvironment `pulumi:"environments"`
	// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
	// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the ServiceStage application.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Application resource.
type ApplicationArgs struct {
	// Specifies the application description.
	// The description can contian a maximum of `128` characters.
	Description pulumi.StringPtrInput
	// Specifies the enterprise projcet ID to which the application
	// belongs. Changing this will create a new resource.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the configurations of the environment variables.
	// The object structure is documented below.
	Environments ApplicationEnvironmentArrayInput
	// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
	// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the ServiceStage application.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
}

func (ApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationArgs)(nil)).Elem()
}

type ApplicationInput interface {
	pulumi.Input

	ToApplicationOutput() ApplicationOutput
	ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput
}

func (*Application) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (i *Application) ToApplicationOutput() ApplicationOutput {
	return i.ToApplicationOutputWithContext(context.Background())
}

func (i *Application) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationOutput)
}

// ApplicationArrayInput is an input type that accepts ApplicationArray and ApplicationArrayOutput values.
// You can construct a concrete instance of `ApplicationArrayInput` via:
//
//	ApplicationArray{ ApplicationArgs{...} }
type ApplicationArrayInput interface {
	pulumi.Input

	ToApplicationArrayOutput() ApplicationArrayOutput
	ToApplicationArrayOutputWithContext(context.Context) ApplicationArrayOutput
}

type ApplicationArray []ApplicationInput

func (ApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (i ApplicationArray) ToApplicationArrayOutput() ApplicationArrayOutput {
	return i.ToApplicationArrayOutputWithContext(context.Background())
}

func (i ApplicationArray) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationArrayOutput)
}

// ApplicationMapInput is an input type that accepts ApplicationMap and ApplicationMapOutput values.
// You can construct a concrete instance of `ApplicationMapInput` via:
//
//	ApplicationMap{ "key": ApplicationArgs{...} }
type ApplicationMapInput interface {
	pulumi.Input

	ToApplicationMapOutput() ApplicationMapOutput
	ToApplicationMapOutputWithContext(context.Context) ApplicationMapOutput
}

type ApplicationMap map[string]ApplicationInput

func (ApplicationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (i ApplicationMap) ToApplicationMapOutput() ApplicationMapOutput {
	return i.ToApplicationMapOutputWithContext(context.Background())
}

func (i ApplicationMap) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationMapOutput)
}

type ApplicationOutput struct{ *pulumi.OutputState }

func (ApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Application)(nil)).Elem()
}

func (o ApplicationOutput) ToApplicationOutput() ApplicationOutput {
	return o
}

func (o ApplicationOutput) ToApplicationOutputWithContext(ctx context.Context) ApplicationOutput {
	return o
}

// The list of component IDs associated under the application.
func (o ApplicationOutput) ComponentIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Application) pulumi.StringArrayOutput { return v.ComponentIds }).(pulumi.StringArrayOutput)
}

// Specifies the application description.
// The description can contian a maximum of `128` characters.
func (o ApplicationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Application) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the enterprise projcet ID to which the application
// belongs. Changing this will create a new resource.
func (o ApplicationOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Specifies the configurations of the environment variables.
// The object structure is documented below.
func (o ApplicationOutput) Environments() ApplicationEnvironmentArrayOutput {
	return o.ApplyT(func(v *Application) ApplicationEnvironmentArrayOutput { return v.Environments }).(ApplicationEnvironmentArrayOutput)
}

// Specifies the variable name. The name can contain `1` to `64` characters, only letters,
// digits, underscores (_), hyphens (-) and dots (.) are allowed. The name cannot start with a digit or dot.
func (o ApplicationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region in which to create the ServiceStage application.
// If omitted, the provider-level region will be used. Changing this will create a new resource.
func (o ApplicationOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Application) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

type ApplicationArrayOutput struct{ *pulumi.OutputState }

func (ApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Application)(nil)).Elem()
}

func (o ApplicationArrayOutput) ToApplicationArrayOutput() ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) ToApplicationArrayOutputWithContext(ctx context.Context) ApplicationArrayOutput {
	return o
}

func (o ApplicationArrayOutput) Index(i pulumi.IntInput) ApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Application {
		return vs[0].([]*Application)[vs[1].(int)]
	}).(ApplicationOutput)
}

type ApplicationMapOutput struct{ *pulumi.OutputState }

func (ApplicationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Application)(nil)).Elem()
}

func (o ApplicationMapOutput) ToApplicationMapOutput() ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) ToApplicationMapOutputWithContext(ctx context.Context) ApplicationMapOutput {
	return o
}

func (o ApplicationMapOutput) MapIndex(k pulumi.StringInput) ApplicationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Application {
		return vs[0].(map[string]*Application)[vs[1].(string)]
	}).(ApplicationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationInput)(nil)).Elem(), &Application{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationArrayInput)(nil)).Elem(), ApplicationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApplicationMapInput)(nil)).Elem(), ApplicationMap{})
	pulumi.RegisterOutputType(ApplicationOutput{})
	pulumi.RegisterOutputType(ApplicationArrayOutput{})
	pulumi.RegisterOutputType(ApplicationMapOutput{})
}

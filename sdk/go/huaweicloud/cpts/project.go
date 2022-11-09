// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cpts

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a project resource within HuaweiCloud CPTS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Cpts"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Cpts.NewProject(ctx, "test", &Cpts.ProjectArgs{
//				Description: pulumi.String("cpts project description"),
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
// Projects can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Cpts/project:Project test 1090
//
// ```
type Project struct {
	pulumi.CustomResourceState

	// The creation time, in UTC format.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Specifies the description of the project, which can contain a maximum of
	// 50 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies a name for the project, which can contain a maximum of 42 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the region in which to create the project resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The last update time, in UTC format.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		args = &ProjectArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterResource("huaweicloud:Cpts/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("huaweicloud:Cpts/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// The creation time, in UTC format.
	CreatedAt *string `pulumi:"createdAt"`
	// Specifies the description of the project, which can contain a maximum of
	// 50 characters.
	Description *string `pulumi:"description"`
	// Specifies a name for the project, which can contain a maximum of 42 characters.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the project resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
	// The last update time, in UTC format.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type ProjectState struct {
	// The creation time, in UTC format.
	CreatedAt pulumi.StringPtrInput
	// Specifies the description of the project, which can contain a maximum of
	// 50 characters.
	Description pulumi.StringPtrInput
	// Specifies a name for the project, which can contain a maximum of 42 characters.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the project resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
	// The last update time, in UTC format.
	UpdatedAt pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Specifies the description of the project, which can contain a maximum of
	// 50 characters.
	Description *string `pulumi:"description"`
	// Specifies a name for the project, which can contain a maximum of 42 characters.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create the project resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Specifies the description of the project, which can contain a maximum of
	// 50 characters.
	Description pulumi.StringPtrInput
	// Specifies a name for the project, which can contain a maximum of 42 characters.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create the project resource. If omitted, the
	// provider-level region will be used. Changing this parameter will create a new resource.
	Region pulumi.StringPtrInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The creation time, in UTC format.
func (o ProjectOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Specifies the description of the project, which can contain a maximum of
// 50 characters.
func (o ProjectOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Project) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies a name for the project, which can contain a maximum of 42 characters.
func (o ProjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the region in which to create the project resource. If omitted, the
// provider-level region will be used. Changing this parameter will create a new resource.
func (o ProjectOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The last update time, in UTC format.
func (o ProjectOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Project) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}

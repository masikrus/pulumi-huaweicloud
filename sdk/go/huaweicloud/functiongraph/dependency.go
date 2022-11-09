// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package functiongraph

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a custom dependency package within HuaweiCloud FunctionGraph.
//
// ## Example Usage
//
// ## Import
//
// Dependencies can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:FunctionGraph/dependency:Dependency test 795e722f-0c23-41b6-a189-dcd56f889cf6
//
// ```
type Dependency struct {
	pulumi.CustomResourceState

	// Specifies the dependency description.
	// The description can contain a maximum of 512 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// The unique ID of the dependency package.
	Etag pulumi.StringOutput `pulumi:"etag"`
	Link pulumi.StringOutput `pulumi:"link"`
	// Specifies the dependeny name.
	// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
	// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// The base64 encoded digest of the dependency after encryption by MD5.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Specifies the region in which to create a custom dependency package.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the dependency package runtime.
	// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
	// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
	// **PHP7.3**.
	Runtime pulumi.StringOutput `pulumi:"runtime"`
	// The dependency package size in bytes.
	Size pulumi.IntOutput `pulumi:"size"`
}

// NewDependency registers a new resource with the given unique name, arguments, and options.
func NewDependency(ctx *pulumi.Context,
	name string, args *DependencyArgs, opts ...pulumi.ResourceOption) (*Dependency, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Link == nil {
		return nil, errors.New("invalid value for required argument 'Link'")
	}
	if args.Runtime == nil {
		return nil, errors.New("invalid value for required argument 'Runtime'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Dependency
	err := ctx.RegisterResource("huaweicloud:FunctionGraph/dependency:Dependency", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDependency gets an existing Dependency resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDependency(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DependencyState, opts ...pulumi.ResourceOption) (*Dependency, error) {
	var resource Dependency
	err := ctx.ReadResource("huaweicloud:FunctionGraph/dependency:Dependency", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dependency resources.
type dependencyState struct {
	// Specifies the dependency description.
	// The description can contain a maximum of 512 characters.
	Description *string `pulumi:"description"`
	// The unique ID of the dependency package.
	Etag *string `pulumi:"etag"`
	Link *string `pulumi:"link"`
	// Specifies the dependeny name.
	// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
	// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
	Name *string `pulumi:"name"`
	// The base64 encoded digest of the dependency after encryption by MD5.
	Owner *string `pulumi:"owner"`
	// Specifies the region in which to create a custom dependency package.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the dependency package runtime.
	// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
	// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
	// **PHP7.3**.
	Runtime *string `pulumi:"runtime"`
	// The dependency package size in bytes.
	Size *int `pulumi:"size"`
}

type DependencyState struct {
	// Specifies the dependency description.
	// The description can contain a maximum of 512 characters.
	Description pulumi.StringPtrInput
	// The unique ID of the dependency package.
	Etag pulumi.StringPtrInput
	Link pulumi.StringPtrInput
	// Specifies the dependeny name.
	// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
	// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
	Name pulumi.StringPtrInput
	// The base64 encoded digest of the dependency after encryption by MD5.
	Owner pulumi.StringPtrInput
	// Specifies the region in which to create a custom dependency package.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the dependency package runtime.
	// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
	// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
	// **PHP7.3**.
	Runtime pulumi.StringPtrInput
	// The dependency package size in bytes.
	Size pulumi.IntPtrInput
}

func (DependencyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dependencyState)(nil)).Elem()
}

type dependencyArgs struct {
	// Specifies the dependency description.
	// The description can contain a maximum of 512 characters.
	Description *string `pulumi:"description"`
	Link        string  `pulumi:"link"`
	// Specifies the dependeny name.
	// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
	// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
	Name *string `pulumi:"name"`
	// Specifies the region in which to create a custom dependency package.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region *string `pulumi:"region"`
	// Specifies the dependency package runtime.
	// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
	// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
	// **PHP7.3**.
	Runtime string `pulumi:"runtime"`
}

// The set of arguments for constructing a Dependency resource.
type DependencyArgs struct {
	// Specifies the dependency description.
	// The description can contain a maximum of 512 characters.
	Description pulumi.StringPtrInput
	Link        pulumi.StringInput
	// Specifies the dependeny name.
	// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
	// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
	Name pulumi.StringPtrInput
	// Specifies the region in which to create a custom dependency package.
	// If omitted, the provider-level region will be used. Changing this will create a new resource.
	Region pulumi.StringPtrInput
	// Specifies the dependency package runtime.
	// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
	// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
	// **PHP7.3**.
	Runtime pulumi.StringInput
}

func (DependencyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dependencyArgs)(nil)).Elem()
}

type DependencyInput interface {
	pulumi.Input

	ToDependencyOutput() DependencyOutput
	ToDependencyOutputWithContext(ctx context.Context) DependencyOutput
}

func (*Dependency) ElementType() reflect.Type {
	return reflect.TypeOf((**Dependency)(nil)).Elem()
}

func (i *Dependency) ToDependencyOutput() DependencyOutput {
	return i.ToDependencyOutputWithContext(context.Background())
}

func (i *Dependency) ToDependencyOutputWithContext(ctx context.Context) DependencyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependencyOutput)
}

// DependencyArrayInput is an input type that accepts DependencyArray and DependencyArrayOutput values.
// You can construct a concrete instance of `DependencyArrayInput` via:
//
//	DependencyArray{ DependencyArgs{...} }
type DependencyArrayInput interface {
	pulumi.Input

	ToDependencyArrayOutput() DependencyArrayOutput
	ToDependencyArrayOutputWithContext(context.Context) DependencyArrayOutput
}

type DependencyArray []DependencyInput

func (DependencyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dependency)(nil)).Elem()
}

func (i DependencyArray) ToDependencyArrayOutput() DependencyArrayOutput {
	return i.ToDependencyArrayOutputWithContext(context.Background())
}

func (i DependencyArray) ToDependencyArrayOutputWithContext(ctx context.Context) DependencyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependencyArrayOutput)
}

// DependencyMapInput is an input type that accepts DependencyMap and DependencyMapOutput values.
// You can construct a concrete instance of `DependencyMapInput` via:
//
//	DependencyMap{ "key": DependencyArgs{...} }
type DependencyMapInput interface {
	pulumi.Input

	ToDependencyMapOutput() DependencyMapOutput
	ToDependencyMapOutputWithContext(context.Context) DependencyMapOutput
}

type DependencyMap map[string]DependencyInput

func (DependencyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dependency)(nil)).Elem()
}

func (i DependencyMap) ToDependencyMapOutput() DependencyMapOutput {
	return i.ToDependencyMapOutputWithContext(context.Background())
}

func (i DependencyMap) ToDependencyMapOutputWithContext(ctx context.Context) DependencyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DependencyMapOutput)
}

type DependencyOutput struct{ *pulumi.OutputState }

func (DependencyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dependency)(nil)).Elem()
}

func (o DependencyOutput) ToDependencyOutput() DependencyOutput {
	return o
}

func (o DependencyOutput) ToDependencyOutputWithContext(ctx context.Context) DependencyOutput {
	return o
}

// Specifies the dependency description.
// The description can contain a maximum of 512 characters.
func (o DependencyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The unique ID of the dependency package.
func (o DependencyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DependencyOutput) Link() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Link }).(pulumi.StringOutput)
}

// Specifies the dependeny name.
// The name can contain a maximum of 96 characters and must start with a letter and end with a letter or digit.
// Only letters, digits, underscores (_), periods (.), and hyphens (-) are allowed.
func (o DependencyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The base64 encoded digest of the dependency after encryption by MD5.
func (o DependencyOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

// Specifies the region in which to create a custom dependency package.
// If omitted, the provider-level region will be used. Changing this will create a new resource.
func (o DependencyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the dependency package runtime.
// The valid values are **Java8**, **Node.js6.10**, **Node.js8.10**, **Node.js10.16**, **Node.js12.13**, **Python2.7**,
// **Python3.6**, **Go1.8**, **Go1.x**, **C#(.NET Core 2.0)**, **C#(.NET Core 2.1)**, **C#(.NET Core 3.1)** and
// **PHP7.3**.
func (o DependencyOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v *Dependency) pulumi.StringOutput { return v.Runtime }).(pulumi.StringOutput)
}

// The dependency package size in bytes.
func (o DependencyOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *Dependency) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

type DependencyArrayOutput struct{ *pulumi.OutputState }

func (DependencyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dependency)(nil)).Elem()
}

func (o DependencyArrayOutput) ToDependencyArrayOutput() DependencyArrayOutput {
	return o
}

func (o DependencyArrayOutput) ToDependencyArrayOutputWithContext(ctx context.Context) DependencyArrayOutput {
	return o
}

func (o DependencyArrayOutput) Index(i pulumi.IntInput) DependencyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dependency {
		return vs[0].([]*Dependency)[vs[1].(int)]
	}).(DependencyOutput)
}

type DependencyMapOutput struct{ *pulumi.OutputState }

func (DependencyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dependency)(nil)).Elem()
}

func (o DependencyMapOutput) ToDependencyMapOutput() DependencyMapOutput {
	return o
}

func (o DependencyMapOutput) ToDependencyMapOutputWithContext(ctx context.Context) DependencyMapOutput {
	return o
}

func (o DependencyMapOutput) MapIndex(k pulumi.StringInput) DependencyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dependency {
		return vs[0].(map[string]*Dependency)[vs[1].(string)]
	}).(DependencyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyInput)(nil)).Elem(), &Dependency{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyArrayInput)(nil)).Elem(), DependencyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DependencyMapInput)(nil)).Elem(), DependencyMap{})
	pulumi.RegisterOutputType(DependencyOutput{})
	pulumi.RegisterOutputType(DependencyArrayOutput{})
	pulumi.RegisterOutputType(DependencyMapOutput{})
}

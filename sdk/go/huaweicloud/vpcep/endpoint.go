// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpcep

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage a VPC endpoint resource.
//
// ## Example Usage
// ### Access to the public service
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpcep"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpcep"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			vpcId := cfg.RequireObject("vpcId")
//			networkId := cfg.RequireObject("networkId")
//			cloudService, err := Vpcep.GetPublicServices(ctx, &vpcep.GetPublicServicesArgs{
//				ServiceName: pulumi.StringRef("dis"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Vpcep.NewEndpoint(ctx, "myendpoint", &Vpcep.EndpointArgs{
//				ServiceId:       pulumi.String(cloudService.Services[0].Id),
//				VpcId:           pulumi.Any(vpcId),
//				NetworkId:       pulumi.Any(networkId),
//				EnableDns:       pulumi.Bool(true),
//				EnableWhitelist: pulumi.Bool(true),
//				Whitelists: pulumi.StringArray{
//					pulumi.String("192.168.0.0/24"),
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
// ### Access to the private service
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Vpcep"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Vpcep"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			serviceVpcId := cfg.RequireObject("serviceVpcId")
//			vmPort := cfg.RequireObject("vmPort")
//			vpcId := cfg.RequireObject("vpcId")
//			networkId := cfg.RequireObject("networkId")
//			demoService, err := Vpcep.NewService(ctx, "demoService", &Vpcep.ServiceArgs{
//				ServerType: pulumi.String("VM"),
//				VpcId:      pulumi.Any(serviceVpcId),
//				PortId:     pulumi.Any(vmPort),
//				PortMappings: vpcep.ServicePortMappingArray{
//					&vpcep.ServicePortMappingArgs{
//						ServicePort:  pulumi.Int(8080),
//						TerminalPort: pulumi.Int(80),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = Vpcep.NewEndpoint(ctx, "demoEndpoint", &Vpcep.EndpointArgs{
//				ServiceId:   demoService.ID(),
//				VpcId:       pulumi.Any(vpcId),
//				NetworkId:   pulumi.Any(networkId),
//				EnableDns:   pulumi.Bool(true),
//				Description: pulumi.String("test description"),
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
// VPC endpoint can be imported using the `id`, e.g. bash
//
// ```sh
//
//	$ pulumi import huaweicloud:Vpcep/endpoint:Endpoint test <id>
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// Specifies the description of the VPC endpoint.
	Description pulumi.StringOutput `pulumi:"description"`
	// Specifies whether to create a private domain name. The default value is
	// true. Changing this creates a new VPC endpoint.
	EnableDns pulumi.BoolPtrOutput `pulumi:"enableDns"`
	// Specifies whether to enable access control. The default value is
	// false.
	EnableWhitelist pulumi.BoolPtrOutput `pulumi:"enableWhitelist"`
	// Specifies the IP address for accessing the associated VPC endpoint
	// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
	// Changing this creates a new VPC endpoint.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The packet ID of the VPC endpoint.
	PacketId pulumi.IntOutput `pulumi:"packetId"`
	// The domain name for accessing the associated VPC endpoint service. This parameter is only
	// available when enableDns is set to true.
	PrivateDomainName pulumi.StringOutput `pulumi:"privateDomainName"`
	// The region in which to create the VPC endpoint. If omitted, the provider-level
	// region will be used. Changing this creates a new VPC endpoint.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the ID of the VPC endpoint service.
	// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The name of the VPC endpoint service.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The type of the VPC endpoint service.
	ServiceType pulumi.StringOutput `pulumi:"serviceType"`
	// The status of the VPC endpoint. The value can be **accepted**, **pendingAcceptance** or **rejected**.
	Status pulumi.StringOutput `pulumi:"status"`
	// The key/value pairs to associate with the VPC endpoint.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
	// this creates a new VPC endpoint.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// Specifies the list of IP address or CIDR block which can be accessed to the
	// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
	Whitelists pulumi.StringArrayOutput `pulumi:"whitelists"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("huaweicloud:Vpcep/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("huaweicloud:Vpcep/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// Specifies the description of the VPC endpoint.
	Description *string `pulumi:"description"`
	// Specifies whether to create a private domain name. The default value is
	// true. Changing this creates a new VPC endpoint.
	EnableDns *bool `pulumi:"enableDns"`
	// Specifies whether to enable access control. The default value is
	// false.
	EnableWhitelist *bool `pulumi:"enableWhitelist"`
	// Specifies the IP address for accessing the associated VPC endpoint
	// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	IpAddress *string `pulumi:"ipAddress"`
	// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
	// Changing this creates a new VPC endpoint.
	NetworkId *string `pulumi:"networkId"`
	// The packet ID of the VPC endpoint.
	PacketId *int `pulumi:"packetId"`
	// The domain name for accessing the associated VPC endpoint service. This parameter is only
	// available when enableDns is set to true.
	PrivateDomainName *string `pulumi:"privateDomainName"`
	// The region in which to create the VPC endpoint. If omitted, the provider-level
	// region will be used. Changing this creates a new VPC endpoint.
	Region *string `pulumi:"region"`
	// Specifies the ID of the VPC endpoint service.
	// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
	ServiceId *string `pulumi:"serviceId"`
	// The name of the VPC endpoint service.
	ServiceName *string `pulumi:"serviceName"`
	// The type of the VPC endpoint service.
	ServiceType *string `pulumi:"serviceType"`
	// The status of the VPC endpoint. The value can be **accepted**, **pendingAcceptance** or **rejected**.
	Status *string `pulumi:"status"`
	// The key/value pairs to associate with the VPC endpoint.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
	// this creates a new VPC endpoint.
	VpcId *string `pulumi:"vpcId"`
	// Specifies the list of IP address or CIDR block which can be accessed to the
	// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
	Whitelists []string `pulumi:"whitelists"`
}

type EndpointState struct {
	// Specifies the description of the VPC endpoint.
	Description pulumi.StringPtrInput
	// Specifies whether to create a private domain name. The default value is
	// true. Changing this creates a new VPC endpoint.
	EnableDns pulumi.BoolPtrInput
	// Specifies whether to enable access control. The default value is
	// false.
	EnableWhitelist pulumi.BoolPtrInput
	// Specifies the IP address for accessing the associated VPC endpoint
	// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	IpAddress pulumi.StringPtrInput
	// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
	// Changing this creates a new VPC endpoint.
	NetworkId pulumi.StringPtrInput
	// The packet ID of the VPC endpoint.
	PacketId pulumi.IntPtrInput
	// The domain name for accessing the associated VPC endpoint service. This parameter is only
	// available when enableDns is set to true.
	PrivateDomainName pulumi.StringPtrInput
	// The region in which to create the VPC endpoint. If omitted, the provider-level
	// region will be used. Changing this creates a new VPC endpoint.
	Region pulumi.StringPtrInput
	// Specifies the ID of the VPC endpoint service.
	// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
	ServiceId pulumi.StringPtrInput
	// The name of the VPC endpoint service.
	ServiceName pulumi.StringPtrInput
	// The type of the VPC endpoint service.
	ServiceType pulumi.StringPtrInput
	// The status of the VPC endpoint. The value can be **accepted**, **pendingAcceptance** or **rejected**.
	Status pulumi.StringPtrInput
	// The key/value pairs to associate with the VPC endpoint.
	Tags pulumi.StringMapInput
	// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
	// this creates a new VPC endpoint.
	VpcId pulumi.StringPtrInput
	// Specifies the list of IP address or CIDR block which can be accessed to the
	// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
	Whitelists pulumi.StringArrayInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// Specifies the description of the VPC endpoint.
	Description *string `pulumi:"description"`
	// Specifies whether to create a private domain name. The default value is
	// true. Changing this creates a new VPC endpoint.
	EnableDns *bool `pulumi:"enableDns"`
	// Specifies whether to enable access control. The default value is
	// false.
	EnableWhitelist *bool `pulumi:"enableWhitelist"`
	// Specifies the IP address for accessing the associated VPC endpoint
	// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	IpAddress *string `pulumi:"ipAddress"`
	// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
	// Changing this creates a new VPC endpoint.
	NetworkId string `pulumi:"networkId"`
	// The region in which to create the VPC endpoint. If omitted, the provider-level
	// region will be used. Changing this creates a new VPC endpoint.
	Region *string `pulumi:"region"`
	// Specifies the ID of the VPC endpoint service.
	// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
	ServiceId string `pulumi:"serviceId"`
	// The key/value pairs to associate with the VPC endpoint.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
	// this creates a new VPC endpoint.
	VpcId string `pulumi:"vpcId"`
	// Specifies the list of IP address or CIDR block which can be accessed to the
	// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
	Whitelists []string `pulumi:"whitelists"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// Specifies the description of the VPC endpoint.
	Description pulumi.StringPtrInput
	// Specifies whether to create a private domain name. The default value is
	// true. Changing this creates a new VPC endpoint.
	EnableDns pulumi.BoolPtrInput
	// Specifies whether to enable access control. The default value is
	// false.
	EnableWhitelist pulumi.BoolPtrInput
	// Specifies the IP address for accessing the associated VPC endpoint
	// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
	IpAddress pulumi.StringPtrInput
	// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
	// Changing this creates a new VPC endpoint.
	NetworkId pulumi.StringInput
	// The region in which to create the VPC endpoint. If omitted, the provider-level
	// region will be used. Changing this creates a new VPC endpoint.
	Region pulumi.StringPtrInput
	// Specifies the ID of the VPC endpoint service.
	// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
	ServiceId pulumi.StringInput
	// The key/value pairs to associate with the VPC endpoint.
	Tags pulumi.StringMapInput
	// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
	// this creates a new VPC endpoint.
	VpcId pulumi.StringInput
	// Specifies the list of IP address or CIDR block which can be accessed to the
	// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
	Whitelists pulumi.StringArrayInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// Specifies the description of the VPC endpoint.
func (o EndpointOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Specifies whether to create a private domain name. The default value is
// true. Changing this creates a new VPC endpoint.
func (o EndpointOutput) EnableDns() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.EnableDns }).(pulumi.BoolPtrOutput)
}

// Specifies whether to enable access control. The default value is
// false.
func (o EndpointOutput) EnableWhitelist() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.BoolPtrOutput { return v.EnableWhitelist }).(pulumi.BoolPtrOutput)
}

// Specifies the IP address for accessing the associated VPC endpoint
// service. Only IPv4 addresses are supported. Changing this creates a new VPC endpoint.
func (o EndpointOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// Specifies the network ID of the subnet in the VPC specified by `vpcId`.
// Changing this creates a new VPC endpoint.
func (o EndpointOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The packet ID of the VPC endpoint.
func (o EndpointOutput) PacketId() pulumi.IntOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.IntOutput { return v.PacketId }).(pulumi.IntOutput)
}

// The domain name for accessing the associated VPC endpoint service. This parameter is only
// available when enableDns is set to true.
func (o EndpointOutput) PrivateDomainName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.PrivateDomainName }).(pulumi.StringOutput)
}

// The region in which to create the VPC endpoint. If omitted, the provider-level
// region will be used. Changing this creates a new VPC endpoint.
func (o EndpointOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the ID of the VPC endpoint service.
// The VPC endpoint service could be private or public. Changing this creates a new VPC endpoint.
func (o EndpointOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The name of the VPC endpoint service.
func (o EndpointOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The type of the VPC endpoint service.
func (o EndpointOutput) ServiceType() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ServiceType }).(pulumi.StringOutput)
}

// The status of the VPC endpoint. The value can be **accepted**, **pendingAcceptance** or **rejected**.
func (o EndpointOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The key/value pairs to associate with the VPC endpoint.
func (o EndpointOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the ID of the VPC where the VPC endpoint is to be created. Changing
// this creates a new VPC endpoint.
func (o EndpointOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// Specifies the list of IP address or CIDR block which can be accessed to the
// VPC endpoint. This field is valid when `enableWhitelist` is set to **true**. The max length of whitelist is 20.
func (o EndpointOutput) Whitelists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.Whitelists }).(pulumi.StringArrayOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}

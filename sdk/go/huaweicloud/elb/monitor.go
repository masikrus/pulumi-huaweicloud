// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an ELB monitor resource within HuaweiCloud.
//
// ## Example Usage
// ### TCP Health Check
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Elb.NewMonitor(ctx, "monitorTcp", &Elb.MonitorArgs{
//				PoolId:     pulumi.Any(_var.Pool_id),
//				Type:       pulumi.String("TCP"),
//				Delay:      pulumi.Int(5),
//				Timeout:    pulumi.Int(3),
//				MaxRetries: pulumi.Int(3),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### UDP Health Check
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Elb.NewMonitor(ctx, "monitorUdp", &Elb.MonitorArgs{
//				PoolId:     pulumi.Any(_var.Pool_id),
//				Type:       pulumi.String("UDP_CONNECT"),
//				Delay:      pulumi.Int(5),
//				Timeout:    pulumi.Int(3),
//				MaxRetries: pulumi.Int(3),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### HTTP Health Check
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Elb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Elb.NewMonitor(ctx, "monitorHttp", &Elb.MonitorArgs{
//				PoolId:        pulumi.Any(_var.Pool_id),
//				Type:          pulumi.String("HTTP"),
//				Delay:         pulumi.Int(5),
//				Timeout:       pulumi.Int(3),
//				MaxRetries:    pulumi.Int(3),
//				UrlPath:       pulumi.String("/test"),
//				ExpectedCodes: pulumi.String("200-202"),
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
// ELB monitor can be imported using the monitor ID, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Elb/monitor:Monitor monitor_1 5c20fdad-7288-11eb-b817-0255ac10158b
//
// ```
type Monitor struct {
	pulumi.CustomResourceState

	AdminStateUp pulumi.BoolPtrOutput `pulumi:"adminStateUp"`
	// Specifies the maximum time between health checks in the unit of second. The value ranges
	// from 1 to 50.
	Delay pulumi.IntOutput `pulumi:"delay"`
	// Specifies the expected HTTP status code. Required for HTTP type.
	// You can either specify a single status like "200", or a range like "200-202".
	ExpectedCodes pulumi.StringOutput `pulumi:"expectedCodes"`
	// Specifies the HTTP request method. Required for HTTP type.
	// The default value is *GET*.
	HttpMethod pulumi.StringOutput `pulumi:"httpMethod"`
	// Specifies the maximum number of consecutive health checks after which the backend
	// servers are declared *healthy*. The value ranges from 1 to 10.
	MaxRetries pulumi.IntOutput `pulumi:"maxRetries"`
	// Specifies the health check name. The value contains a maximum of 255 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the id of the pool that this monitor will be assigned to. Changing
	// this creates a new monitor.
	PoolId pulumi.StringOutput `pulumi:"poolId"`
	// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
	// the port of the backend server will be used as the health check port.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// The region in which to create the ELB monitor resource. If omitted, the
	// provider-level region will be used. Changing this creates a new monitor.
	Region pulumi.StringOutput `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// Specifies the health check timeout duration in the unit of second.
	// The value ranges from 1 to 50 and must be less than the `delay` value.
	Timeout pulumi.IntOutput `pulumi:"timeout"`
	// Specifies the monitor protocol.
	// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
	// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
	Type pulumi.StringOutput `pulumi:"type"`
	// Specifies the HTTP request path for the health check. Required for HTTP type.
	// The value starts with a slash (/) and contains a maximum of 255 characters.
	UrlPath pulumi.StringOutput `pulumi:"urlPath"`
}

// NewMonitor registers a new resource with the given unique name, arguments, and options.
func NewMonitor(ctx *pulumi.Context,
	name string, args *MonitorArgs, opts ...pulumi.ResourceOption) (*Monitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Delay == nil {
		return nil, errors.New("invalid value for required argument 'Delay'")
	}
	if args.MaxRetries == nil {
		return nil, errors.New("invalid value for required argument 'MaxRetries'")
	}
	if args.PoolId == nil {
		return nil, errors.New("invalid value for required argument 'PoolId'")
	}
	if args.Timeout == nil {
		return nil, errors.New("invalid value for required argument 'Timeout'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Monitor
	err := ctx.RegisterResource("huaweicloud:Elb/monitor:Monitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitor gets an existing Monitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitorState, opts ...pulumi.ResourceOption) (*Monitor, error) {
	var resource Monitor
	err := ctx.ReadResource("huaweicloud:Elb/monitor:Monitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Monitor resources.
type monitorState struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Specifies the maximum time between health checks in the unit of second. The value ranges
	// from 1 to 50.
	Delay *int `pulumi:"delay"`
	// Specifies the expected HTTP status code. Required for HTTP type.
	// You can either specify a single status like "200", or a range like "200-202".
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Specifies the HTTP request method. Required for HTTP type.
	// The default value is *GET*.
	HttpMethod *string `pulumi:"httpMethod"`
	// Specifies the maximum number of consecutive health checks after which the backend
	// servers are declared *healthy*. The value ranges from 1 to 10.
	MaxRetries *int `pulumi:"maxRetries"`
	// Specifies the health check name. The value contains a maximum of 255 characters.
	Name *string `pulumi:"name"`
	// Specifies the id of the pool that this monitor will be assigned to. Changing
	// this creates a new monitor.
	PoolId *string `pulumi:"poolId"`
	// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
	// the port of the backend server will be used as the health check port.
	Port *int `pulumi:"port"`
	// The region in which to create the ELB monitor resource. If omitted, the
	// provider-level region will be used. Changing this creates a new monitor.
	Region *string `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
	// Specifies the health check timeout duration in the unit of second.
	// The value ranges from 1 to 50 and must be less than the `delay` value.
	Timeout *int `pulumi:"timeout"`
	// Specifies the monitor protocol.
	// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
	// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
	Type *string `pulumi:"type"`
	// Specifies the HTTP request path for the health check. Required for HTTP type.
	// The value starts with a slash (/) and contains a maximum of 255 characters.
	UrlPath *string `pulumi:"urlPath"`
}

type MonitorState struct {
	AdminStateUp pulumi.BoolPtrInput
	// Specifies the maximum time between health checks in the unit of second. The value ranges
	// from 1 to 50.
	Delay pulumi.IntPtrInput
	// Specifies the expected HTTP status code. Required for HTTP type.
	// You can either specify a single status like "200", or a range like "200-202".
	ExpectedCodes pulumi.StringPtrInput
	// Specifies the HTTP request method. Required for HTTP type.
	// The default value is *GET*.
	HttpMethod pulumi.StringPtrInput
	// Specifies the maximum number of consecutive health checks after which the backend
	// servers are declared *healthy*. The value ranges from 1 to 10.
	MaxRetries pulumi.IntPtrInput
	// Specifies the health check name. The value contains a maximum of 255 characters.
	Name pulumi.StringPtrInput
	// Specifies the id of the pool that this monitor will be assigned to. Changing
	// this creates a new monitor.
	PoolId pulumi.StringPtrInput
	// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
	// the port of the backend server will be used as the health check port.
	Port pulumi.IntPtrInput
	// The region in which to create the ELB monitor resource. If omitted, the
	// provider-level region will be used. Changing this creates a new monitor.
	Region pulumi.StringPtrInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
	// Specifies the health check timeout duration in the unit of second.
	// The value ranges from 1 to 50 and must be less than the `delay` value.
	Timeout pulumi.IntPtrInput
	// Specifies the monitor protocol.
	// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
	// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
	Type pulumi.StringPtrInput
	// Specifies the HTTP request path for the health check. Required for HTTP type.
	// The value starts with a slash (/) and contains a maximum of 255 characters.
	UrlPath pulumi.StringPtrInput
}

func (MonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorState)(nil)).Elem()
}

type monitorArgs struct {
	AdminStateUp *bool `pulumi:"adminStateUp"`
	// Specifies the maximum time between health checks in the unit of second. The value ranges
	// from 1 to 50.
	Delay int `pulumi:"delay"`
	// Specifies the expected HTTP status code. Required for HTTP type.
	// You can either specify a single status like "200", or a range like "200-202".
	ExpectedCodes *string `pulumi:"expectedCodes"`
	// Specifies the HTTP request method. Required for HTTP type.
	// The default value is *GET*.
	HttpMethod *string `pulumi:"httpMethod"`
	// Specifies the maximum number of consecutive health checks after which the backend
	// servers are declared *healthy*. The value ranges from 1 to 10.
	MaxRetries int `pulumi:"maxRetries"`
	// Specifies the health check name. The value contains a maximum of 255 characters.
	Name *string `pulumi:"name"`
	// Specifies the id of the pool that this monitor will be assigned to. Changing
	// this creates a new monitor.
	PoolId string `pulumi:"poolId"`
	// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
	// the port of the backend server will be used as the health check port.
	Port *int `pulumi:"port"`
	// The region in which to create the ELB monitor resource. If omitted, the
	// provider-level region will be used. Changing this creates a new monitor.
	Region *string `pulumi:"region"`
	// Deprecated: tenant_id is deprecated
	TenantId *string `pulumi:"tenantId"`
	// Specifies the health check timeout duration in the unit of second.
	// The value ranges from 1 to 50 and must be less than the `delay` value.
	Timeout int `pulumi:"timeout"`
	// Specifies the monitor protocol.
	// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
	// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
	Type string `pulumi:"type"`
	// Specifies the HTTP request path for the health check. Required for HTTP type.
	// The value starts with a slash (/) and contains a maximum of 255 characters.
	UrlPath *string `pulumi:"urlPath"`
}

// The set of arguments for constructing a Monitor resource.
type MonitorArgs struct {
	AdminStateUp pulumi.BoolPtrInput
	// Specifies the maximum time between health checks in the unit of second. The value ranges
	// from 1 to 50.
	Delay pulumi.IntInput
	// Specifies the expected HTTP status code. Required for HTTP type.
	// You can either specify a single status like "200", or a range like "200-202".
	ExpectedCodes pulumi.StringPtrInput
	// Specifies the HTTP request method. Required for HTTP type.
	// The default value is *GET*.
	HttpMethod pulumi.StringPtrInput
	// Specifies the maximum number of consecutive health checks after which the backend
	// servers are declared *healthy*. The value ranges from 1 to 10.
	MaxRetries pulumi.IntInput
	// Specifies the health check name. The value contains a maximum of 255 characters.
	Name pulumi.StringPtrInput
	// Specifies the id of the pool that this monitor will be assigned to. Changing
	// this creates a new monitor.
	PoolId pulumi.StringInput
	// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
	// the port of the backend server will be used as the health check port.
	Port pulumi.IntPtrInput
	// The region in which to create the ELB monitor resource. If omitted, the
	// provider-level region will be used. Changing this creates a new monitor.
	Region pulumi.StringPtrInput
	// Deprecated: tenant_id is deprecated
	TenantId pulumi.StringPtrInput
	// Specifies the health check timeout duration in the unit of second.
	// The value ranges from 1 to 50 and must be less than the `delay` value.
	Timeout pulumi.IntInput
	// Specifies the monitor protocol.
	// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
	// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
	Type pulumi.StringInput
	// Specifies the HTTP request path for the health check. Required for HTTP type.
	// The value starts with a slash (/) and contains a maximum of 255 characters.
	UrlPath pulumi.StringPtrInput
}

func (MonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitorArgs)(nil)).Elem()
}

type MonitorInput interface {
	pulumi.Input

	ToMonitorOutput() MonitorOutput
	ToMonitorOutputWithContext(ctx context.Context) MonitorOutput
}

func (*Monitor) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (i *Monitor) ToMonitorOutput() MonitorOutput {
	return i.ToMonitorOutputWithContext(context.Background())
}

func (i *Monitor) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorOutput)
}

// MonitorArrayInput is an input type that accepts MonitorArray and MonitorArrayOutput values.
// You can construct a concrete instance of `MonitorArrayInput` via:
//
//	MonitorArray{ MonitorArgs{...} }
type MonitorArrayInput interface {
	pulumi.Input

	ToMonitorArrayOutput() MonitorArrayOutput
	ToMonitorArrayOutputWithContext(context.Context) MonitorArrayOutput
}

type MonitorArray []MonitorInput

func (MonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (i MonitorArray) ToMonitorArrayOutput() MonitorArrayOutput {
	return i.ToMonitorArrayOutputWithContext(context.Background())
}

func (i MonitorArray) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorArrayOutput)
}

// MonitorMapInput is an input type that accepts MonitorMap and MonitorMapOutput values.
// You can construct a concrete instance of `MonitorMapInput` via:
//
//	MonitorMap{ "key": MonitorArgs{...} }
type MonitorMapInput interface {
	pulumi.Input

	ToMonitorMapOutput() MonitorMapOutput
	ToMonitorMapOutputWithContext(context.Context) MonitorMapOutput
}

type MonitorMap map[string]MonitorInput

func (MonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (i MonitorMap) ToMonitorMapOutput() MonitorMapOutput {
	return i.ToMonitorMapOutputWithContext(context.Background())
}

func (i MonitorMap) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitorMapOutput)
}

type MonitorOutput struct{ *pulumi.OutputState }

func (MonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Monitor)(nil)).Elem()
}

func (o MonitorOutput) ToMonitorOutput() MonitorOutput {
	return o
}

func (o MonitorOutput) ToMonitorOutputWithContext(ctx context.Context) MonitorOutput {
	return o
}

func (o MonitorOutput) AdminStateUp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.BoolPtrOutput { return v.AdminStateUp }).(pulumi.BoolPtrOutput)
}

// Specifies the maximum time between health checks in the unit of second. The value ranges
// from 1 to 50.
func (o MonitorOutput) Delay() pulumi.IntOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntOutput { return v.Delay }).(pulumi.IntOutput)
}

// Specifies the expected HTTP status code. Required for HTTP type.
// You can either specify a single status like "200", or a range like "200-202".
func (o MonitorOutput) ExpectedCodes() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.ExpectedCodes }).(pulumi.StringOutput)
}

// Specifies the HTTP request method. Required for HTTP type.
// The default value is *GET*.
func (o MonitorOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.HttpMethod }).(pulumi.StringOutput)
}

// Specifies the maximum number of consecutive health checks after which the backend
// servers are declared *healthy*. The value ranges from 1 to 10.
func (o MonitorOutput) MaxRetries() pulumi.IntOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntOutput { return v.MaxRetries }).(pulumi.IntOutput)
}

// Specifies the health check name. The value contains a maximum of 255 characters.
func (o MonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies the id of the pool that this monitor will be assigned to. Changing
// this creates a new monitor.
func (o MonitorOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.PoolId }).(pulumi.StringOutput)
}

// Specifies the health check port. The port number ranges from 1 to 65535. If not specified,
// the port of the backend server will be used as the health check port.
func (o MonitorOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// The region in which to create the ELB monitor resource. If omitted, the
// provider-level region will be used. Changing this creates a new monitor.
func (o MonitorOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Deprecated: tenant_id is deprecated
func (o MonitorOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// Specifies the health check timeout duration in the unit of second.
// The value ranges from 1 to 50 and must be less than the `delay` value.
func (o MonitorOutput) Timeout() pulumi.IntOutput {
	return o.ApplyT(func(v *Monitor) pulumi.IntOutput { return v.Timeout }).(pulumi.IntOutput)
}

// Specifies the monitor protocol.
// The value can be *TCP*, *UDP_CONNECT*, or *HTTP*.
// If the listener protocol is UDP, the monitor protocol must be *UDP_CONNECT*. Changing this creates a new monitor.
func (o MonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Specifies the HTTP request path for the health check. Required for HTTP type.
// The value starts with a slash (/) and contains a maximum of 255 characters.
func (o MonitorOutput) UrlPath() pulumi.StringOutput {
	return o.ApplyT(func(v *Monitor) pulumi.StringOutput { return v.UrlPath }).(pulumi.StringOutput)
}

type MonitorArrayOutput struct{ *pulumi.OutputState }

func (MonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Monitor)(nil)).Elem()
}

func (o MonitorArrayOutput) ToMonitorArrayOutput() MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) ToMonitorArrayOutputWithContext(ctx context.Context) MonitorArrayOutput {
	return o
}

func (o MonitorArrayOutput) Index(i pulumi.IntInput) MonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].([]*Monitor)[vs[1].(int)]
	}).(MonitorOutput)
}

type MonitorMapOutput struct{ *pulumi.OutputState }

func (MonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Monitor)(nil)).Elem()
}

func (o MonitorMapOutput) ToMonitorMapOutput() MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) ToMonitorMapOutputWithContext(ctx context.Context) MonitorMapOutput {
	return o
}

func (o MonitorMapOutput) MapIndex(k pulumi.StringInput) MonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Monitor {
		return vs[0].(map[string]*Monitor)[vs[1].(string)]
	}).(MonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorInput)(nil)).Elem(), &Monitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorArrayInput)(nil)).Elem(), MonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonitorMapInput)(nil)).Elem(), MonitorMap{})
	pulumi.RegisterOutputType(MonitorOutput{})
	pulumi.RegisterOutputType(MonitorArrayOutput{})
	pulumi.RegisterOutputType(MonitorMapOutput{})
}

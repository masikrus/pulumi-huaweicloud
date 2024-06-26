// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages dds instance resource within HuaweiCloud.
//
// ## Example Usage
// ### Creating A Cluster Community Edition
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dds.NewInstance(ctx, "instance", &Dds.InstanceArgs{
//				AvailabilityZone: pulumi.String("{{ availability_zone }}"),
//				BackupStrategy: &dds.InstanceBackupStrategyArgs{
//					KeepDays:  pulumi.Int(8),
//					StartTime: pulumi.String("08:00-09:00"),
//				},
//				Datastore: &dds.InstanceDatastoreArgs{
//					StorageEngine: pulumi.String("wiredTiger"),
//					Type:          pulumi.String("DDS-Community"),
//					Version:       pulumi.String("3.4"),
//				},
//				Flavors: dds.InstanceFlavorArray{
//					&dds.InstanceFlavorArgs{
//						Num:      pulumi.Int(2),
//						SpecCode: pulumi.String("dds.mongodb.c3.medium.4.mongos"),
//						Type:     pulumi.String("mongos"),
//					},
//					&dds.InstanceFlavorArgs{
//						Num:      pulumi.Int(2),
//						Size:     pulumi.Int(20),
//						SpecCode: pulumi.String("dds.mongodb.c3.medium.4.shard"),
//						Storage:  pulumi.String("ULTRAHIGH"),
//						Type:     pulumi.String("shard"),
//					},
//					&dds.InstanceFlavorArgs{
//						Num:      pulumi.Int(1),
//						Size:     pulumi.Int(20),
//						SpecCode: pulumi.String("dds.mongodb.c3.large.2.config"),
//						Storage:  pulumi.String("ULTRAHIGH"),
//						Type:     pulumi.String("config"),
//					},
//				},
//				Mode:            pulumi.String("Sharding"),
//				Password:        pulumi.String("Test@123"),
//				SecurityGroupId: pulumi.String("{{ security_group_id }}"),
//				SubnetId:        pulumi.String("{{ subnet_network_id }}}"),
//				VpcId:           pulumi.String("{{ vpc_id }}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating A Replica Set Community Edition
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dds.NewInstance(ctx, "instance", &Dds.InstanceArgs{
//				AvailabilityZone: pulumi.String("{{ availability_zone }}"),
//				Datastore: &dds.InstanceDatastoreArgs{
//					StorageEngine: pulumi.String("wiredTiger"),
//					Type:          pulumi.String("DDS-Community"),
//					Version:       pulumi.String("3.4"),
//				},
//				Flavors: dds.InstanceFlavorArray{
//					&dds.InstanceFlavorArgs{
//						Num:      pulumi.Int(1),
//						Size:     pulumi.Int(30),
//						SpecCode: pulumi.String("dds.mongodb.c3.medium.4.repset"),
//						Storage:  pulumi.String("ULTRAHIGH"),
//						Type:     pulumi.String("replica"),
//					},
//				},
//				Mode:            pulumi.String("ReplicaSet"),
//				Password:        pulumi.String("Test@123"),
//				SecurityGroupId: pulumi.String("{{ security_group_id }}"),
//				SubnetId:        pulumi.String("{{ subnet_network_id }}}"),
//				VpcId:           pulumi.String("{{ vpc_id }}"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating A Single Community Edition
//
// ```go
// package main
//
// import (
//
//	"github.com/huaweicloud/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi-huaweicloud/sdk/go/huaweicloud/Dds"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dds.NewInstance(ctx, "instance", &Dds.InstanceArgs{
//				AvailabilityZone: pulumi.String("{{ availability_zone }}"),
//				Datastore: &dds.InstanceDatastoreArgs{
//					StorageEngine: pulumi.String("wiredTiger"),
//					Type:          pulumi.String("DDS-Community"),
//					Version:       pulumi.String("3.4"),
//				},
//				Flavors: dds.InstanceFlavorArray{
//					&dds.InstanceFlavorArgs{
//						Num:      pulumi.Int(1),
//						Size:     pulumi.Int(30),
//						SpecCode: pulumi.String("dds.mongodb.s6.large.2.single"),
//						Storage:  pulumi.String("ULTRAHIGH"),
//						Type:     pulumi.String("single"),
//					},
//				},
//				Mode:            pulumi.String("Single"),
//				Password:        pulumi.String("Test@123"),
//				SecurityGroupId: pulumi.String("{{ security_group_id }}"),
//				SubnetId:        pulumi.String("{{ subnet_network_id }}}"),
//				VpcId:           pulumi.String("{{ vpc_id }}"),
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
// DDS instance can be imported using the `id`, e.g.
//
// ```sh
//
//	$ pulumi import huaweicloud:Dds/instance:Instance instance 9c6d6ff2cba3434293fd479571517e16in02
//
// ```
//
//	Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`password`, `availability_zone`, `flavor`, configuration. It is generally recommended running `terraform plan` after importing an instance. You can then decide if changes should be applied to the instance, or the resource definition should be updated to align with the instance. Also you can ignore changes as below. resource "huaweicloud_dds_instance" "instance" {
//
//	...
//
//	lifecycle {
//
//	ignore_changes = [
//
//	password, availability_zone, flavor, configuration,
//
//	]
//
//	} }
type Instance struct {
	pulumi.CustomResourceState

	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrOutput `pulumi:"autoPay"`
	// Specifies whether auto-renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrOutput `pulumi:"autoRenew"`
	// Specifies the ID of the availability zone. Changing this creates a
	// new instance.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// Specifies the advanced backup policy. The structure is described below.
	BackupStrategy InstanceBackupStrategyOutput `pulumi:"backupStrategy"`
	// Specifies the charging mode of the instance.
	// The valid values are as follows:
	// + `prePaid`: indicates the yearly/monthly billing mode.
	// + `postPaid`: indicates the pay-per-use billing mode.
	ChargingMode pulumi.StringOutput `pulumi:"chargingMode"`
	// Specifies the configuration information.
	// The structure is described below. Changing this creates a new instance.
	Configurations InstanceConfigurationArrayOutput `pulumi:"configurations"`
	// Specifies database information. The structure is described below. Changing
	// this creates a new instance.
	Datastore InstanceDatastoreOutput `pulumi:"datastore"`
	// Indicates the DB Administator name.
	DbUsername pulumi.StringOutput `pulumi:"dbUsername"`
	// Specifies the disk encryption ID of the instance. Changing this
	// creates a new instance.
	DiskEncryptionId pulumi.StringPtrOutput `pulumi:"diskEncryptionId"`
	// Specifies the enterprise project id of the dds instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringOutput `pulumi:"enterpriseProjectId"`
	// Specifies the flavors information. The structure is described below. Changing
	// this creates a new instance.
	Flavors InstanceFlavorArrayOutput `pulumi:"flavors"`
	// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
	// **Single** are supported. Changing this creates a new instance.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Specifies the DB instance name. The DB instance name of the same type is unique in the
	// same tenant.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates the instance nodes information. Structure is documented below.
	Nodes InstanceNodeArrayOutput `pulumi:"nodes"`
	// Specifies the Administrator password of the database instance.
	Password pulumi.StringOutput `pulumi:"password"`
	// Specifies the charging period of the instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// Specifies the database access port. The valid values are range from `2100` to `9500` and
	// `27017`, `27018`, `27019`. Defaults to `8635`.
	Port pulumi.IntOutput `pulumi:"port"`
	// Specifies the region of the DDS instance. Changing this creates a new
	// instance.
	Region pulumi.StringOutput `pulumi:"region"`
	// Specifies the security group ID of the DDS instance.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
	// Specifies whether to enable or disable SSL. Defaults to true.
	Ssl pulumi.BoolPtrOutput `pulumi:"ssl"`
	// Indicates the node status.
	Status pulumi.StringOutput `pulumi:"status"`
	// Specifies the subnet Network ID. Changing this creates a new instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// The key/value pairs to associate with the DDS instance.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AvailabilityZone == nil {
		return nil, errors.New("invalid value for required argument 'AvailabilityZone'")
	}
	if args.Datastore == nil {
		return nil, errors.New("invalid value for required argument 'Datastore'")
	}
	if args.Flavors == nil {
		return nil, errors.New("invalid value for required argument 'Flavors'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.SecurityGroupId == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("huaweicloud:Dds/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("huaweicloud:Dds/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Deprecated: Deprecated
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto-renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the ID of the availability zone. Changing this creates a
	// new instance.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// Specifies the advanced backup policy. The structure is described below.
	BackupStrategy *InstanceBackupStrategy `pulumi:"backupStrategy"`
	// Specifies the charging mode of the instance.
	// The valid values are as follows:
	// + `prePaid`: indicates the yearly/monthly billing mode.
	// + `postPaid`: indicates the pay-per-use billing mode.
	ChargingMode *string `pulumi:"chargingMode"`
	// Specifies the configuration information.
	// The structure is described below. Changing this creates a new instance.
	Configurations []InstanceConfiguration `pulumi:"configurations"`
	// Specifies database information. The structure is described below. Changing
	// this creates a new instance.
	Datastore *InstanceDatastore `pulumi:"datastore"`
	// Indicates the DB Administator name.
	DbUsername *string `pulumi:"dbUsername"`
	// Specifies the disk encryption ID of the instance. Changing this
	// creates a new instance.
	DiskEncryptionId *string `pulumi:"diskEncryptionId"`
	// Specifies the enterprise project id of the dds instance.
	// Changing this creates a new instance.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the flavors information. The structure is described below. Changing
	// this creates a new instance.
	Flavors []InstanceFlavor `pulumi:"flavors"`
	// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
	// **Single** are supported. Changing this creates a new instance.
	Mode *string `pulumi:"mode"`
	// Specifies the DB instance name. The DB instance name of the same type is unique in the
	// same tenant.
	Name *string `pulumi:"name"`
	// Indicates the instance nodes information. Structure is documented below.
	Nodes []InstanceNode `pulumi:"nodes"`
	// Specifies the Administrator password of the database instance.
	Password *string `pulumi:"password"`
	// Specifies the charging period of the instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	PeriodUnit *string `pulumi:"periodUnit"`
	// Specifies the database access port. The valid values are range from `2100` to `9500` and
	// `27017`, `27018`, `27019`. Defaults to `8635`.
	Port *int `pulumi:"port"`
	// Specifies the region of the DDS instance. Changing this creates a new
	// instance.
	Region *string `pulumi:"region"`
	// Specifies the security group ID of the DDS instance.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// Specifies whether to enable or disable SSL. Defaults to true.
	Ssl *bool `pulumi:"ssl"`
	// Indicates the node status.
	Status *string `pulumi:"status"`
	// Specifies the subnet Network ID. Changing this creates a new instance.
	SubnetId *string `pulumi:"subnetId"`
	// The key/value pairs to associate with the DDS instance.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId *string `pulumi:"vpcId"`
}

type InstanceState struct {
	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto-renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrInput
	// Specifies the ID of the availability zone. Changing this creates a
	// new instance.
	AvailabilityZone pulumi.StringPtrInput
	// Specifies the advanced backup policy. The structure is described below.
	BackupStrategy InstanceBackupStrategyPtrInput
	// Specifies the charging mode of the instance.
	// The valid values are as follows:
	// + `prePaid`: indicates the yearly/monthly billing mode.
	// + `postPaid`: indicates the pay-per-use billing mode.
	ChargingMode pulumi.StringPtrInput
	// Specifies the configuration information.
	// The structure is described below. Changing this creates a new instance.
	Configurations InstanceConfigurationArrayInput
	// Specifies database information. The structure is described below. Changing
	// this creates a new instance.
	Datastore InstanceDatastorePtrInput
	// Indicates the DB Administator name.
	DbUsername pulumi.StringPtrInput
	// Specifies the disk encryption ID of the instance. Changing this
	// creates a new instance.
	DiskEncryptionId pulumi.StringPtrInput
	// Specifies the enterprise project id of the dds instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the flavors information. The structure is described below. Changing
	// this creates a new instance.
	Flavors InstanceFlavorArrayInput
	// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
	// **Single** are supported. Changing this creates a new instance.
	Mode pulumi.StringPtrInput
	// Specifies the DB instance name. The DB instance name of the same type is unique in the
	// same tenant.
	Name pulumi.StringPtrInput
	// Indicates the instance nodes information. Structure is documented below.
	Nodes InstanceNodeArrayInput
	// Specifies the Administrator password of the database instance.
	Password pulumi.StringPtrInput
	// Specifies the charging period of the instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringPtrInput
	// Specifies the database access port. The valid values are range from `2100` to `9500` and
	// `27017`, `27018`, `27019`. Defaults to `8635`.
	Port pulumi.IntPtrInput
	// Specifies the region of the DDS instance. Changing this creates a new
	// instance.
	Region pulumi.StringPtrInput
	// Specifies the security group ID of the DDS instance.
	SecurityGroupId pulumi.StringPtrInput
	// Specifies whether to enable or disable SSL. Defaults to true.
	Ssl pulumi.BoolPtrInput
	// Indicates the node status.
	Status pulumi.StringPtrInput
	// Specifies the subnet Network ID. Changing this creates a new instance.
	SubnetId pulumi.StringPtrInput
	// The key/value pairs to associate with the DDS instance.
	Tags pulumi.StringMapInput
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Deprecated: Deprecated
	AutoPay *string `pulumi:"autoPay"`
	// Specifies whether auto-renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew *string `pulumi:"autoRenew"`
	// Specifies the ID of the availability zone. Changing this creates a
	// new instance.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// Specifies the advanced backup policy. The structure is described below.
	BackupStrategy *InstanceBackupStrategy `pulumi:"backupStrategy"`
	// Specifies the charging mode of the instance.
	// The valid values are as follows:
	// + `prePaid`: indicates the yearly/monthly billing mode.
	// + `postPaid`: indicates the pay-per-use billing mode.
	ChargingMode *string `pulumi:"chargingMode"`
	// Specifies the configuration information.
	// The structure is described below. Changing this creates a new instance.
	Configurations []InstanceConfiguration `pulumi:"configurations"`
	// Specifies database information. The structure is described below. Changing
	// this creates a new instance.
	Datastore InstanceDatastore `pulumi:"datastore"`
	// Specifies the disk encryption ID of the instance. Changing this
	// creates a new instance.
	DiskEncryptionId *string `pulumi:"diskEncryptionId"`
	// Specifies the enterprise project id of the dds instance.
	// Changing this creates a new instance.
	EnterpriseProjectId *string `pulumi:"enterpriseProjectId"`
	// Specifies the flavors information. The structure is described below. Changing
	// this creates a new instance.
	Flavors []InstanceFlavor `pulumi:"flavors"`
	// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
	// **Single** are supported. Changing this creates a new instance.
	Mode string `pulumi:"mode"`
	// Specifies the DB instance name. The DB instance name of the same type is unique in the
	// same tenant.
	Name *string `pulumi:"name"`
	// Specifies the Administrator password of the database instance.
	Password string `pulumi:"password"`
	// Specifies the charging period of the instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	Period *int `pulumi:"period"`
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	PeriodUnit *string `pulumi:"periodUnit"`
	// Specifies the database access port. The valid values are range from `2100` to `9500` and
	// `27017`, `27018`, `27019`. Defaults to `8635`.
	Port *int `pulumi:"port"`
	// Specifies the region of the DDS instance. Changing this creates a new
	// instance.
	Region *string `pulumi:"region"`
	// Specifies the security group ID of the DDS instance.
	SecurityGroupId string `pulumi:"securityGroupId"`
	// Specifies whether to enable or disable SSL. Defaults to true.
	Ssl *bool `pulumi:"ssl"`
	// Specifies the subnet Network ID. Changing this creates a new instance.
	SubnetId string `pulumi:"subnetId"`
	// The key/value pairs to associate with the DDS instance.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Deprecated: Deprecated
	AutoPay pulumi.StringPtrInput
	// Specifies whether auto-renew is enabled.
	// Valid values are `true` and `false`, defaults to `false`.
	// Changing this creates a new instance.
	AutoRenew pulumi.StringPtrInput
	// Specifies the ID of the availability zone. Changing this creates a
	// new instance.
	AvailabilityZone pulumi.StringInput
	// Specifies the advanced backup policy. The structure is described below.
	BackupStrategy InstanceBackupStrategyPtrInput
	// Specifies the charging mode of the instance.
	// The valid values are as follows:
	// + `prePaid`: indicates the yearly/monthly billing mode.
	// + `postPaid`: indicates the pay-per-use billing mode.
	ChargingMode pulumi.StringPtrInput
	// Specifies the configuration information.
	// The structure is described below. Changing this creates a new instance.
	Configurations InstanceConfigurationArrayInput
	// Specifies database information. The structure is described below. Changing
	// this creates a new instance.
	Datastore InstanceDatastoreInput
	// Specifies the disk encryption ID of the instance. Changing this
	// creates a new instance.
	DiskEncryptionId pulumi.StringPtrInput
	// Specifies the enterprise project id of the dds instance.
	// Changing this creates a new instance.
	EnterpriseProjectId pulumi.StringPtrInput
	// Specifies the flavors information. The structure is described below. Changing
	// this creates a new instance.
	Flavors InstanceFlavorArrayInput
	// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
	// **Single** are supported. Changing this creates a new instance.
	Mode pulumi.StringInput
	// Specifies the DB instance name. The DB instance name of the same type is unique in the
	// same tenant.
	Name pulumi.StringPtrInput
	// Specifies the Administrator password of the database instance.
	Password pulumi.StringInput
	// Specifies the charging period of the instance.
	// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
	// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
	// This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	Period pulumi.IntPtrInput
	// Specifies the charging period unit of the instance.
	// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
	// Changing this creates a new instance.
	PeriodUnit pulumi.StringPtrInput
	// Specifies the database access port. The valid values are range from `2100` to `9500` and
	// `27017`, `27018`, `27019`. Defaults to `8635`.
	Port pulumi.IntPtrInput
	// Specifies the region of the DDS instance. Changing this creates a new
	// instance.
	Region pulumi.StringPtrInput
	// Specifies the security group ID of the DDS instance.
	SecurityGroupId pulumi.StringInput
	// Specifies whether to enable or disable SSL. Defaults to true.
	Ssl pulumi.BoolPtrInput
	// Specifies the subnet Network ID. Changing this creates a new instance.
	SubnetId pulumi.StringInput
	// The key/value pairs to associate with the DDS instance.
	Tags pulumi.StringMapInput
	// Specifies the VPC ID. Changing this creates a new instance.
	VpcId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Deprecated: Deprecated
func (o InstanceOutput) AutoPay() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.AutoPay }).(pulumi.StringPtrOutput)
}

// Specifies whether auto-renew is enabled.
// Valid values are `true` and `false`, defaults to `false`.
// Changing this creates a new instance.
func (o InstanceOutput) AutoRenew() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.AutoRenew }).(pulumi.StringPtrOutput)
}

// Specifies the ID of the availability zone. Changing this creates a
// new instance.
func (o InstanceOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// Specifies the advanced backup policy. The structure is described below.
func (o InstanceOutput) BackupStrategy() InstanceBackupStrategyOutput {
	return o.ApplyT(func(v *Instance) InstanceBackupStrategyOutput { return v.BackupStrategy }).(InstanceBackupStrategyOutput)
}

// Specifies the charging mode of the instance.
// The valid values are as follows:
// + `prePaid`: indicates the yearly/monthly billing mode.
// + `postPaid`: indicates the pay-per-use billing mode.
func (o InstanceOutput) ChargingMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ChargingMode }).(pulumi.StringOutput)
}

// Specifies the configuration information.
// The structure is described below. Changing this creates a new instance.
func (o InstanceOutput) Configurations() InstanceConfigurationArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceConfigurationArrayOutput { return v.Configurations }).(InstanceConfigurationArrayOutput)
}

// Specifies database information. The structure is described below. Changing
// this creates a new instance.
func (o InstanceOutput) Datastore() InstanceDatastoreOutput {
	return o.ApplyT(func(v *Instance) InstanceDatastoreOutput { return v.Datastore }).(InstanceDatastoreOutput)
}

// Indicates the DB Administator name.
func (o InstanceOutput) DbUsername() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DbUsername }).(pulumi.StringOutput)
}

// Specifies the disk encryption ID of the instance. Changing this
// creates a new instance.
func (o InstanceOutput) DiskEncryptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.DiskEncryptionId }).(pulumi.StringPtrOutput)
}

// Specifies the enterprise project id of the dds instance.
// Changing this creates a new instance.
func (o InstanceOutput) EnterpriseProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.EnterpriseProjectId }).(pulumi.StringOutput)
}

// Specifies the flavors information. The structure is described below. Changing
// this creates a new instance.
func (o InstanceOutput) Flavors() InstanceFlavorArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceFlavorArrayOutput { return v.Flavors }).(InstanceFlavorArrayOutput)
}

// Specifies the mode of the database instance. **Sharding**, **ReplicaSet**,
// **Single** are supported. Changing this creates a new instance.
func (o InstanceOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Specifies the DB instance name. The DB instance name of the same type is unique in the
// same tenant.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Indicates the instance nodes information. Structure is documented below.
func (o InstanceOutput) Nodes() InstanceNodeArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceNodeArrayOutput { return v.Nodes }).(InstanceNodeArrayOutput)
}

// Specifies the Administrator password of the database instance.
func (o InstanceOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Specifies the charging period of the instance.
// If `periodUnit` is set to *month*, the value ranges from 1 to 9.
// If `periodUnit` is set to *year*, the value ranges from 1 to 3.
// This parameter is mandatory if `chargingMode` is set to *prePaid*.
// Changing this creates a new instance.
func (o InstanceOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// Specifies the charging period unit of the instance.
// Valid values are *month* and *year*. This parameter is mandatory if `chargingMode` is set to *prePaid*.
// Changing this creates a new instance.
func (o InstanceOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// Specifies the database access port. The valid values are range from `2100` to `9500` and
// `27017`, `27018`, `27019`. Defaults to `8635`.
func (o InstanceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Specifies the region of the DDS instance. Changing this creates a new
// instance.
func (o InstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Specifies the security group ID of the DDS instance.
func (o InstanceOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

// Specifies whether to enable or disable SSL. Defaults to true.
func (o InstanceOutput) Ssl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolPtrOutput { return v.Ssl }).(pulumi.BoolPtrOutput)
}

// Indicates the node status.
func (o InstanceOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Specifies the subnet Network ID. Changing this creates a new instance.
func (o InstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// The key/value pairs to associate with the DDS instance.
func (o InstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the VPC ID. Changing this creates a new instance.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}

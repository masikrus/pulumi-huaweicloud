# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServerTemplateArgs', 'ServerTemplate']

@pulumi.input_type
class ServerTemplateArgs:
    def __init__(__self__, *,
                 availability_zone: pulumi.Input[str],
                 bandwidth_size: Optional[pulumi.Input[int]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_server_name: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServerTemplate resource.
        :param pulumi.Input[str] availability_zone: Specifies the availability zone where the target server is located.
        :param pulumi.Input[int] bandwidth_size: Specifies the bandwidth size in Mbit/s about the public IP address
               that will be used for migration.
        :param pulumi.Input[str] flavor: Specifies the flavor ID for the target server.
        :param pulumi.Input[str] name: Specifies the server template name.
        :param pulumi.Input[str] project_id: Specifies the project ID where the target server is located.
               If omitted, the default project in the region will be used.
        :param pulumi.Input[str] region: Specifies the region where the target server is located.
               If omitted, the provider-level region will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies an array of one or more security group IDs to associate with
               the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: Specifies an array of one or more subnet IDs to attach to the target server.
               If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        :param pulumi.Input[str] target_server_name: Specifies the name of the target server. Defaults to the template name.
        :param pulumi.Input[str] volume_type: Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
               defaults to **SAS**.
        :param pulumi.Input[str] vpc_id: Specifies the ID of the VPC which the target server belongs to.
               If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        """
        pulumi.set(__self__, "availability_zone", availability_zone)
        if bandwidth_size is not None:
            pulumi.set(__self__, "bandwidth_size", bandwidth_size)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if target_server_name is not None:
            pulumi.set(__self__, "target_server_name", target_server_name)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Input[str]:
        """
        Specifies the availability zone where the target server is located.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: pulumi.Input[str]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="bandwidthSize")
    def bandwidth_size(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the bandwidth size in Mbit/s about the public IP address
        that will be used for migration.
        """
        return pulumi.get(self, "bandwidth_size")

    @bandwidth_size.setter
    def bandwidth_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth_size", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the flavor ID for the target server.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the server template name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the project ID where the target server is located.
        If omitted, the default project in the region will be used.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the target server is located.
        If omitted, the provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more security group IDs to associate with
        the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more subnet IDs to attach to the target server.
        If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="targetServerName")
    def target_server_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the target server. Defaults to the template name.
        """
        return pulumi.get(self, "target_server_name")

    @target_server_name.setter
    def target_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_server_name", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
        defaults to **SAS**.
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the VPC which the target server belongs to.
        If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class _ServerTemplateState:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 bandwidth_size: Optional[pulumi.Input[int]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_server_name: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 vpc_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServerTemplate resources.
        :param pulumi.Input[str] availability_zone: Specifies the availability zone where the target server is located.
        :param pulumi.Input[int] bandwidth_size: Specifies the bandwidth size in Mbit/s about the public IP address
               that will be used for migration.
        :param pulumi.Input[str] flavor: Specifies the flavor ID for the target server.
        :param pulumi.Input[str] name: Specifies the server template name.
        :param pulumi.Input[str] project_id: Specifies the project ID where the target server is located.
               If omitted, the default project in the region will be used.
        :param pulumi.Input[str] region: Specifies the region where the target server is located.
               If omitted, the provider-level region will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies an array of one or more security group IDs to associate with
               the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: Specifies an array of one or more subnet IDs to attach to the target server.
               If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        :param pulumi.Input[str] target_server_name: Specifies the name of the target server. Defaults to the template name.
        :param pulumi.Input[str] volume_type: Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
               defaults to **SAS**.
        :param pulumi.Input[str] vpc_id: Specifies the ID of the VPC which the target server belongs to.
               If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        :param pulumi.Input[str] vpc_name: The name of the VPC which the target server belongs to.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if bandwidth_size is not None:
            pulumi.set(__self__, "bandwidth_size", bandwidth_size)
        if flavor is not None:
            pulumi.set(__self__, "flavor", flavor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if subnet_ids is not None:
            pulumi.set(__self__, "subnet_ids", subnet_ids)
        if target_server_name is not None:
            pulumi.set(__self__, "target_server_name", target_server_name)
        if volume_type is not None:
            pulumi.set(__self__, "volume_type", volume_type)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)
        if vpc_name is not None:
            pulumi.set(__self__, "vpc_name", vpc_name)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the availability zone where the target server is located.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter(name="bandwidthSize")
    def bandwidth_size(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the bandwidth size in Mbit/s about the public IP address
        that will be used for migration.
        """
        return pulumi.get(self, "bandwidth_size")

    @bandwidth_size.setter
    def bandwidth_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth_size", value)

    @property
    @pulumi.getter
    def flavor(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the flavor ID for the target server.
        """
        return pulumi.get(self, "flavor")

    @flavor.setter
    def flavor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "flavor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the server template name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the project ID where the target server is located.
        If omitted, the default project in the region will be used.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region where the target server is located.
        If omitted, the provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more security group IDs to associate with
        the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of one or more subnet IDs to attach to the target server.
        If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        """
        return pulumi.get(self, "subnet_ids")

    @subnet_ids.setter
    def subnet_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "subnet_ids", value)

    @property
    @pulumi.getter(name="targetServerName")
    def target_server_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the target server. Defaults to the template name.
        """
        return pulumi.get(self, "target_server_name")

    @target_server_name.setter
    def target_server_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_server_name", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
        defaults to **SAS**.
        """
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "volume_type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the VPC which the target server belongs to.
        If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the VPC which the target server belongs to.
        """
        return pulumi.get(self, "vpc_name")

    @vpc_name.setter
    def vpc_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_name", value)


class ServerTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 bandwidth_size: Optional[pulumi.Input[int]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_server_name: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an SMS server template resource within HuaweiCloud.

        ## Example Usage
        ### A template will create networks during migration

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        demo_availability_zones = huaweicloud.get_availability_zones()
        demo_server_template = huaweicloud.sms.ServerTemplate("demoServerTemplate", availability_zone=demo_availability_zones.names[0])
        ```
        ### A template will use the existing networks during migration

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        vpc_id = config.require_object("vpcId")
        subent_id = config.require_object("subentId")
        secgroup_id = config.require_object("secgroupId")
        demo_availability_zones = huaweicloud.get_availability_zones()
        demo_server_template = huaweicloud.sms.ServerTemplate("demoServerTemplate",
            availability_zone=demo_availability_zones.names[0],
            vpc_id=vpc_id,
            subnet_ids=[subent_id],
            security_group_ids=[secgroup_id])
        ```

        ## Import

        SMS server templates can be imported by `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Sms/serverTemplate:ServerTemplate demo 4618ccaf-b4d7-43b9-b958-3df3b885126d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: Specifies the availability zone where the target server is located.
        :param pulumi.Input[int] bandwidth_size: Specifies the bandwidth size in Mbit/s about the public IP address
               that will be used for migration.
        :param pulumi.Input[str] flavor: Specifies the flavor ID for the target server.
        :param pulumi.Input[str] name: Specifies the server template name.
        :param pulumi.Input[str] project_id: Specifies the project ID where the target server is located.
               If omitted, the default project in the region will be used.
        :param pulumi.Input[str] region: Specifies the region where the target server is located.
               If omitted, the provider-level region will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies an array of one or more security group IDs to associate with
               the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: Specifies an array of one or more subnet IDs to attach to the target server.
               If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        :param pulumi.Input[str] target_server_name: Specifies the name of the target server. Defaults to the template name.
        :param pulumi.Input[str] volume_type: Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
               defaults to **SAS**.
        :param pulumi.Input[str] vpc_id: Specifies the ID of the VPC which the target server belongs to.
               If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an SMS server template resource within HuaweiCloud.

        ## Example Usage
        ### A template will create networks during migration

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        demo_availability_zones = huaweicloud.get_availability_zones()
        demo_server_template = huaweicloud.sms.ServerTemplate("demoServerTemplate", availability_zone=demo_availability_zones.names[0])
        ```
        ### A template will use the existing networks during migration

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        vpc_id = config.require_object("vpcId")
        subent_id = config.require_object("subentId")
        secgroup_id = config.require_object("secgroupId")
        demo_availability_zones = huaweicloud.get_availability_zones()
        demo_server_template = huaweicloud.sms.ServerTemplate("demoServerTemplate",
            availability_zone=demo_availability_zones.names[0],
            vpc_id=vpc_id,
            subnet_ids=[subent_id],
            security_group_ids=[secgroup_id])
        ```

        ## Import

        SMS server templates can be imported by `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Sms/serverTemplate:ServerTemplate demo 4618ccaf-b4d7-43b9-b958-3df3b885126d
        ```

        :param str resource_name: The name of the resource.
        :param ServerTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 bandwidth_size: Optional[pulumi.Input[int]] = None,
                 flavor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 target_server_name: Optional[pulumi.Input[str]] = None,
                 volume_type: Optional[pulumi.Input[str]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServerTemplateArgs.__new__(ServerTemplateArgs)

            if availability_zone is None and not opts.urn:
                raise TypeError("Missing required property 'availability_zone'")
            __props__.__dict__["availability_zone"] = availability_zone
            __props__.__dict__["bandwidth_size"] = bandwidth_size
            __props__.__dict__["flavor"] = flavor
            __props__.__dict__["name"] = name
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["region"] = region
            __props__.__dict__["security_group_ids"] = security_group_ids
            __props__.__dict__["subnet_ids"] = subnet_ids
            __props__.__dict__["target_server_name"] = target_server_name
            __props__.__dict__["volume_type"] = volume_type
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["vpc_name"] = None
        super(ServerTemplate, __self__).__init__(
            'huaweicloud:Sms/serverTemplate:ServerTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            availability_zone: Optional[pulumi.Input[str]] = None,
            bandwidth_size: Optional[pulumi.Input[int]] = None,
            flavor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            subnet_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            target_server_name: Optional[pulumi.Input[str]] = None,
            volume_type: Optional[pulumi.Input[str]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None,
            vpc_name: Optional[pulumi.Input[str]] = None) -> 'ServerTemplate':
        """
        Get an existing ServerTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] availability_zone: Specifies the availability zone where the target server is located.
        :param pulumi.Input[int] bandwidth_size: Specifies the bandwidth size in Mbit/s about the public IP address
               that will be used for migration.
        :param pulumi.Input[str] flavor: Specifies the flavor ID for the target server.
        :param pulumi.Input[str] name: Specifies the server template name.
        :param pulumi.Input[str] project_id: Specifies the project ID where the target server is located.
               If omitted, the default project in the region will be used.
        :param pulumi.Input[str] region: Specifies the region where the target server is located.
               If omitted, the provider-level region will be used.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: Specifies an array of one or more security group IDs to associate with
               the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] subnet_ids: Specifies an array of one or more subnet IDs to attach to the target server.
               If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        :param pulumi.Input[str] target_server_name: Specifies the name of the target server. Defaults to the template name.
        :param pulumi.Input[str] volume_type: Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
               defaults to **SAS**.
        :param pulumi.Input[str] vpc_id: Specifies the ID of the VPC which the target server belongs to.
               If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        :param pulumi.Input[str] vpc_name: The name of the VPC which the target server belongs to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerTemplateState.__new__(_ServerTemplateState)

        __props__.__dict__["availability_zone"] = availability_zone
        __props__.__dict__["bandwidth_size"] = bandwidth_size
        __props__.__dict__["flavor"] = flavor
        __props__.__dict__["name"] = name
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["region"] = region
        __props__.__dict__["security_group_ids"] = security_group_ids
        __props__.__dict__["subnet_ids"] = subnet_ids
        __props__.__dict__["target_server_name"] = target_server_name
        __props__.__dict__["volume_type"] = volume_type
        __props__.__dict__["vpc_id"] = vpc_id
        __props__.__dict__["vpc_name"] = vpc_name
        return ServerTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> pulumi.Output[str]:
        """
        Specifies the availability zone where the target server is located.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="bandwidthSize")
    def bandwidth_size(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the bandwidth size in Mbit/s about the public IP address
        that will be used for migration.
        """
        return pulumi.get(self, "bandwidth_size")

    @property
    @pulumi.getter
    def flavor(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the flavor ID for the target server.
        """
        return pulumi.get(self, "flavor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the server template name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        Specifies the project ID where the target server is located.
        If omitted, the default project in the region will be used.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region where the target server is located.
        If omitted, the provider-level region will be used.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies an array of one or more security group IDs to associate with
        the target server. If omitted or set to ["autoCreate"], a new security group will be created automatically during migration.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter(name="subnetIds")
    def subnet_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies an array of one or more subnet IDs to attach to the target server.
        If omitted or set to ["autoCreate"], a new subnet will be created automatically during migration.
        """
        return pulumi.get(self, "subnet_ids")

    @property
    @pulumi.getter(name="targetServerName")
    def target_server_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the target server. Defaults to the template name.
        """
        return pulumi.get(self, "target_server_name")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the disk type of the target server. Available values are: **SAS**, **SSD**,
        defaults to **SAS**.
        """
        return pulumi.get(self, "volume_type")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the VPC which the target server belongs to.
        If omitted or set to "autoCreate", a new VPC will be created automatically during migration.
        """
        return pulumi.get(self, "vpc_id")

    @property
    @pulumi.getter(name="vpcName")
    def vpc_name(self) -> pulumi.Output[str]:
        """
        The name of the VPC which the target server belongs to.
        """
        return pulumi.get(self, "vpc_name")


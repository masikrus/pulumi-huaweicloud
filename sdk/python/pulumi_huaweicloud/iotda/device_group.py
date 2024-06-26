# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DeviceGroupArgs', 'DeviceGroup']

@pulumi.input_type
class DeviceGroupArgs:
    def __init__(__self__, *,
                 space_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 device_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DeviceGroup resource.
        :param pulumi.Input[str] space_id: Specifies the resource space ID to which the device group belongs.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of device group. The description contains a maximum of 64
               characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
               are allowed: `?'#().,&%@!`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] device_ids: Specifies the list of device IDs bound to the group.
        :param pulumi.Input[str] name: Specifies the name of device group. The name contains a maximum of 64 characters.
               Only letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] parent_group_id: Specifies the parent group id.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA device group resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "space_id", space_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_ids is not None:
            pulumi.set(__self__, "device_ids", device_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_group_id is not None:
            pulumi.set(__self__, "parent_group_id", parent_group_id)
        if region is not None:
            pulumi.set(__self__, "region", region)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Input[str]:
        """
        Specifies the resource space ID to which the device group belongs.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "space_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of device group. The description contains a maximum of 64
        characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
        are allowed: `?'#().,&%@!`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceIds")
    def device_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of device IDs bound to the group.
        """
        return pulumi.get(self, "device_ids")

    @device_ids.setter
    def device_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "device_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of device group. The name contains a maximum of 64 characters.
        Only letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the parent group id.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "parent_group_id")

    @parent_group_id.setter
    def parent_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_group_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the IoTDA device group resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)


@pulumi.input_type
class _DeviceGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 device_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DeviceGroup resources.
        :param pulumi.Input[str] description: Specifies the description of device group. The description contains a maximum of 64
               characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
               are allowed: `?'#().,&%@!`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] device_ids: Specifies the list of device IDs bound to the group.
        :param pulumi.Input[str] name: Specifies the name of device group. The name contains a maximum of 64 characters.
               Only letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] parent_group_id: Specifies the parent group id.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA device group resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] space_id: Specifies the resource space ID to which the device group belongs.
               Changing this parameter will create a new resource.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device_ids is not None:
            pulumi.set(__self__, "device_ids", device_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parent_group_id is not None:
            pulumi.set(__self__, "parent_group_id", parent_group_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if space_id is not None:
            pulumi.set(__self__, "space_id", space_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of device group. The description contains a maximum of 64
        characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
        are allowed: `?'#().,&%@!`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="deviceIds")
    def device_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of device IDs bound to the group.
        """
        return pulumi.get(self, "device_ids")

    @device_ids.setter
    def device_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "device_ids", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of device group. The name contains a maximum of 64 characters.
        Only letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the parent group id.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "parent_group_id")

    @parent_group_id.setter
    def parent_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_group_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the IoTDA device group resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resource space ID to which the device group belongs.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "space_id")

    @space_id.setter
    def space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "space_id", value)


class DeviceGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an IoTDA device group within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        space_id = config.require_object("spaceId")
        device_id = config.require_object("deviceId")
        group = huaweicloud.io_tda.DeviceGroup("group",
            space_id=space_id,
            device_ids=[device_id])
        ```

        ## Import

        Groups can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:IoTDA/deviceGroup:DeviceGroup test 10022532f4f94f26b01daa1e424853e1
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`space_id`. It is generally recommended running `terraform plan` after importing the resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the group. Also you can ignore changes as below. resource "huaweicloud_iotda_device_group" "test" {

         ...

         lifecycle {

         ignore_changes = [

         space_id,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of device group. The description contains a maximum of 64
               characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
               are allowed: `?'#().,&%@!`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] device_ids: Specifies the list of device IDs bound to the group.
        :param pulumi.Input[str] name: Specifies the name of device group. The name contains a maximum of 64 characters.
               Only letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] parent_group_id: Specifies the parent group id.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA device group resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] space_id: Specifies the resource space ID to which the device group belongs.
               Changing this parameter will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an IoTDA device group within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        space_id = config.require_object("spaceId")
        device_id = config.require_object("deviceId")
        group = huaweicloud.io_tda.DeviceGroup("group",
            space_id=space_id,
            device_ids=[device_id])
        ```

        ## Import

        Groups can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:IoTDA/deviceGroup:DeviceGroup test 10022532f4f94f26b01daa1e424853e1
        ```

         Note that the imported state may not be identical to your resource definition, due to some attributes missing from the API response, security or some other reason. The missing attributes include`space_id`. It is generally recommended running `terraform plan` after importing the resource. You can then decide if changes should be applied to the resource, or the resource definition should be updated to align with the group. Also you can ignore changes as below. resource "huaweicloud_iotda_device_group" "test" {

         ...

         lifecycle {

         ignore_changes = [

         space_id,

         ]

         } }

        :param str resource_name: The name of the resource.
        :param DeviceGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_group_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 space_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceGroupArgs.__new__(DeviceGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["device_ids"] = device_ids
            __props__.__dict__["name"] = name
            __props__.__dict__["parent_group_id"] = parent_group_id
            __props__.__dict__["region"] = region
            if space_id is None and not opts.urn:
                raise TypeError("Missing required property 'space_id'")
            __props__.__dict__["space_id"] = space_id
        super(DeviceGroup, __self__).__init__(
            'huaweicloud:IoTDA/deviceGroup:DeviceGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            device_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_group_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            space_id: Optional[pulumi.Input[str]] = None) -> 'DeviceGroup':
        """
        Get an existing DeviceGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of device group. The description contains a maximum of 64
               characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
               are allowed: `?'#().,&%@!`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] device_ids: Specifies the list of device IDs bound to the group.
        :param pulumi.Input[str] name: Specifies the name of device group. The name contains a maximum of 64 characters.
               Only letters, digits, hyphens (-) and underscores (_) are allowed.
        :param pulumi.Input[str] parent_group_id: Specifies the parent group id.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] region: Specifies the region in which to create the IoTDA device group resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] space_id: Specifies the resource space ID to which the device group belongs.
               Changing this parameter will create a new resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceGroupState.__new__(_DeviceGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["device_ids"] = device_ids
        __props__.__dict__["name"] = name
        __props__.__dict__["parent_group_id"] = parent_group_id
        __props__.__dict__["region"] = region
        __props__.__dict__["space_id"] = space_id
        return DeviceGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description of device group. The description contains a maximum of 64
        characters. Only letters, Chinese characters, digits, hyphens (-), underscores (_) and the following special characters
        are allowed: `?'#().,&%@!`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceIds")
    def device_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the list of device IDs bound to the group.
        """
        return pulumi.get(self, "device_ids")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of device group. The name contains a maximum of 64 characters.
        Only letters, digits, hyphens (-) and underscores (_) are allowed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentGroupId")
    def parent_group_id(self) -> pulumi.Output[str]:
        """
        Specifies the parent group id.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "parent_group_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the IoTDA device group resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="spaceId")
    def space_id(self) -> pulumi.Output[str]:
        """
        Specifies the resource space ID to which the device group belongs.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "space_id")


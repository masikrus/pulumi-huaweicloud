# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RecordsetArgs', 'Recordset']

@pulumi.input_type
class RecordsetArgs:
    def __init__(__self__, *,
                 records: pulumi.Input[Sequence[pulumi.Input[str]]],
                 type: pulumi.Input[str],
                 zone_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 line_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Recordset resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: Specifies an array of DNS records. The value rules vary depending on the record set type.
        :param pulumi.Input[str] type: Specifies the type of the record set.
               Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        :param pulumi.Input[str] zone_id: Specifies the zone ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the record set.
        :param pulumi.Input[str] line_id: Specifies the resolution line ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the record set.
               The name suffixed with a zone name, which is a complete host name ended with a dot.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] status: Specifies the status of the record set.
               Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the DNS recordset.
        :param pulumi.Input[int] ttl: Specifies the time to live (TTL) of the record set (in seconds).
               The value range is 1–2147483647. The default value is 300.
        :param pulumi.Input[int] weight: Specifies the weight of the record set.
               Only public zone support. The value range is 0–1000.
        """
        pulumi.set(__self__, "records", records)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "zone_id", zone_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if line_id is not None:
            pulumi.set(__self__, "line_id", line_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def records(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies an array of DNS records. The value rules vary depending on the record set type.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of the record set.
        Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Input[str]:
        """
        Specifies the zone ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the record set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lineId")
    def line_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resolution line ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "line_id")

    @line_id.setter
    def line_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "line_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the record set.
        The name suffixed with a zone name, which is a complete host name ended with a dot.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the status of the record set.
        Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs to associate with the DNS recordset.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the time to live (TTL) of the record set (in seconds).
        The value range is 1–2147483647. The default value is 300.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the weight of the record set.
        Only public zone support. The value range is 0–1000.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _RecordsetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 line_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 zone_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Recordset resources.
        :param pulumi.Input[str] description: Specifies the description of the record set.
        :param pulumi.Input[str] line_id: Specifies the resolution line ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the record set.
               The name suffixed with a zone name, which is a complete host name ended with a dot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: Specifies an array of DNS records. The value rules vary depending on the record set type.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] status: Specifies the status of the record set.
               Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the DNS recordset.
        :param pulumi.Input[int] ttl: Specifies the time to live (TTL) of the record set (in seconds).
               The value range is 1–2147483647. The default value is 300.
        :param pulumi.Input[str] type: Specifies the type of the record set.
               Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        :param pulumi.Input[int] weight: Specifies the weight of the record set.
               Only public zone support. The value range is 0–1000.
        :param pulumi.Input[str] zone_id: Specifies the zone ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] zone_name: The zone name of the record set.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if line_id is not None:
            pulumi.set(__self__, "line_id", line_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if records is not None:
            pulumi.set(__self__, "records", records)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)
        if zone_name is not None:
            pulumi.set(__self__, "zone_name", zone_name)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the record set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="lineId")
    def line_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the resolution line ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "line_id")

    @line_id.setter
    def line_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "line_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the record set.
        The name suffixed with a zone name, which is a complete host name ended with a dot.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def records(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies an array of DNS records. The value rules vary depending on the record set type.
        """
        return pulumi.get(self, "records")

    @records.setter
    def records(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "records", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the status of the record set.
        Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Specifies the key/value pairs to associate with the DNS recordset.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the time to live (TTL) of the record set (in seconds).
        The value range is 1–2147483647. The default value is 300.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of the record set.
        Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the weight of the record set.
        Only public zone support. The value range is 0–1000.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the zone ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> Optional[pulumi.Input[str]]:
        """
        The zone name of the record set.
        """
        return pulumi.get(self, "zone_name")

    @zone_name.setter
    def zone_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_name", value)


class Recordset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 line_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages a DNS record set resource within HuaweiCloud.

        ## Example Usage
        ### Record Set with Multi-line

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a zone",
            ttl=6000,
            zone_type="public")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            type="A",
            description="a recordset description",
            status="ENABLE",
            ttl=300,
            records=["10.1.0.0"],
            line_id="Dianxin_Shanxi",
            weight=3,
            tags={
                "key1": "value1",
                "key2": "value2",
            })
        ```
        ### Record Set with Public Zone

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a public zone",
            ttl=6000,
            zone_type="public")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```
        ### Record Set with Private Zone

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a private zone",
            ttl=6000,
            zone_type="private")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```

        ## Import

        The DNS recordset can be imported using `zone_id`, `recordset_id`, separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Dns/recordset:Recordset test <zone_id>/<recordset_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of the record set.
        :param pulumi.Input[str] line_id: Specifies the resolution line ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the record set.
               The name suffixed with a zone name, which is a complete host name ended with a dot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: Specifies an array of DNS records. The value rules vary depending on the record set type.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] status: Specifies the status of the record set.
               Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the DNS recordset.
        :param pulumi.Input[int] ttl: Specifies the time to live (TTL) of the record set (in seconds).
               The value range is 1–2147483647. The default value is 300.
        :param pulumi.Input[str] type: Specifies the type of the record set.
               Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        :param pulumi.Input[int] weight: Specifies the weight of the record set.
               Only public zone support. The value range is 0–1000.
        :param pulumi.Input[str] zone_id: Specifies the zone ID.
               Changing this parameter will create a new resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RecordsetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a DNS record set resource within HuaweiCloud.

        ## Example Usage
        ### Record Set with Multi-line

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a zone",
            ttl=6000,
            zone_type="public")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            type="A",
            description="a recordset description",
            status="ENABLE",
            ttl=300,
            records=["10.1.0.0"],
            line_id="Dianxin_Shanxi",
            weight=3,
            tags={
                "key1": "value1",
                "key2": "value2",
            })
        ```
        ### Record Set with Public Zone

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a public zone",
            ttl=6000,
            zone_type="public")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```
        ### Record Set with Private Zone

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        example_zone = huaweicloud.dns.Zone("exampleZone",
            email="email2@example.com",
            description="a private zone",
            ttl=6000,
            zone_type="private")
        test = huaweicloud.dns.Recordset("test",
            zone_id=example_zone.id,
            description="An example record set",
            ttl=3000,
            type="A",
            records=["10.0.0.1"])
        ```

        ## Import

        The DNS recordset can be imported using `zone_id`, `recordset_id`, separated by slashes, e.g. bash

        ```sh
         $ pulumi import huaweicloud:Dns/recordset:Recordset test <zone_id>/<recordset_id>
        ```

        :param str resource_name: The name of the resource.
        :param RecordsetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RecordsetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 line_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RecordsetArgs.__new__(RecordsetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["line_id"] = line_id
            __props__.__dict__["name"] = name
            if records is None and not opts.urn:
                raise TypeError("Missing required property 'records'")
            __props__.__dict__["records"] = records
            __props__.__dict__["region"] = region
            __props__.__dict__["status"] = status
            __props__.__dict__["tags"] = tags
            __props__.__dict__["ttl"] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["weight"] = weight
            if zone_id is None and not opts.urn:
                raise TypeError("Missing required property 'zone_id'")
            __props__.__dict__["zone_id"] = zone_id
            __props__.__dict__["zone_name"] = None
        super(Recordset, __self__).__init__(
            'huaweicloud:Dns/recordset:Recordset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            line_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            records: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            region: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None,
            zone_id: Optional[pulumi.Input[str]] = None,
            zone_name: Optional[pulumi.Input[str]] = None) -> 'Recordset':
        """
        Get an existing Recordset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Specifies the description of the record set.
        :param pulumi.Input[str] line_id: Specifies the resolution line ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] name: Specifies the name of the record set.
               The name suffixed with a zone name, which is a complete host name ended with a dot.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] records: Specifies an array of DNS records. The value rules vary depending on the record set type.
        :param pulumi.Input[str] region: Specifies the region in which to create the resource.
               If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        :param pulumi.Input[str] status: Specifies the status of the record set.
               Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: Specifies the key/value pairs to associate with the DNS recordset.
        :param pulumi.Input[int] ttl: Specifies the time to live (TTL) of the record set (in seconds).
               The value range is 1–2147483647. The default value is 300.
        :param pulumi.Input[str] type: Specifies the type of the record set.
               Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        :param pulumi.Input[int] weight: Specifies the weight of the record set.
               Only public zone support. The value range is 0–1000.
        :param pulumi.Input[str] zone_id: Specifies the zone ID.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] zone_name: The zone name of the record set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RecordsetState.__new__(_RecordsetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["line_id"] = line_id
        __props__.__dict__["name"] = name
        __props__.__dict__["records"] = records
        __props__.__dict__["region"] = region
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["type"] = type
        __props__.__dict__["weight"] = weight
        __props__.__dict__["zone_id"] = zone_id
        __props__.__dict__["zone_name"] = zone_name
        return Recordset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Specifies the description of the record set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="lineId")
    def line_id(self) -> pulumi.Output[str]:
        """
        Specifies the resolution line ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "line_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the record set.
        The name suffixed with a zone name, which is a complete host name ended with a dot.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def records(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies an array of DNS records. The value rules vary depending on the record set type.
        """
        return pulumi.get(self, "records")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to create the resource.
        If omitted, the provider-level region will be used. Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the status of the record set.
        Value options: **ENABLE**, **DISABLE**. The default value is **ENABLE**.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Specifies the key/value pairs to associate with the DNS recordset.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the time to live (TTL) of the record set (in seconds).
        The value range is 1–2147483647. The default value is 300.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the type of the record set.
        Value options: **A**, **AAAA**, **MX**, **CNAME**, **TXT**, **NS**, **SRV**, **CAA**.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[int]:
        """
        Specifies the weight of the record set.
        Only public zone support. The value range is 0–1000.
        """
        return pulumi.get(self, "weight")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        Specifies the zone ID.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneName")
    def zone_name(self) -> pulumi.Output[str]:
        """
        The zone name of the record set.
        """
        return pulumi.get(self, "zone_name")


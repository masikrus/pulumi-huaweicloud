# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetVpcsResult',
    'AwaitableGetVpcsResult',
    'get_vpcs',
    'get_vpcs_output',
]

@pulumi.output_type
class GetVpcsResult:
    """
    A collection of values returned by getVpcs.
    """
    def __init__(__self__, cidr=None, enterprise_project_id=None, id=None, name=None, region=None, status=None, tags=None, vpcs=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if enterprise_project_id and not isinstance(enterprise_project_id, str):
            raise TypeError("Expected argument 'enterprise_project_id' to be a str")
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if tags and not isinstance(tags, dict):
            raise TypeError("Expected argument 'tags' to be a dict")
        pulumi.set(__self__, "tags", tags)
        if vpcs and not isinstance(vpcs, list):
            raise TypeError("Expected argument 'vpcs' to be a list")
        pulumi.set(__self__, "vpcs", vpcs)

    @property
    @pulumi.getter
    def cidr(self) -> Optional[str]:
        """
        Indicates the cidr block of the VPC.
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> Optional[str]:
        """
        Indicates the the enterprise project ID of the VPC.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Indicates the ID of the VPC.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        Indicates the name of the VPC.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        Indicates the current status of the VPC.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> Optional[Mapping[str, str]]:
        """
        Indicates the key/value pairs which associated with the VPC.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def vpcs(self) -> Sequence['outputs.GetVpcsVpcResult']:
        """
        Indicates a list of all VPCs found. Structure is documented below.
        """
        return pulumi.get(self, "vpcs")


class AwaitableGetVpcsResult(GetVpcsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcsResult(
            cidr=self.cidr,
            enterprise_project_id=self.enterprise_project_id,
            id=self.id,
            name=self.name,
            region=self.region,
            status=self.status,
            tags=self.tags,
            vpcs=self.vpcs)


def get_vpcs(cidr: Optional[str] = None,
             enterprise_project_id: Optional[str] = None,
             id: Optional[str] = None,
             name: Optional[str] = None,
             region: Optional[str] = None,
             status: Optional[str] = None,
             tags: Optional[Mapping[str, str]] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcsResult:
    """
    Use this data source to get a list of VPC.

    ## Example Usage

    An example filter by name and tag

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    vpc_name = config.require_object("vpcName")
    vpc = huaweicloud.Vpc.get_vpcs(name=vpc_name,
        tags={
            "foo": "bar",
        })
    pulumi.export("vpcIds", [__item.id for __item in [vpc.vpcs]])
    ```


    :param str cidr: Specifies the cidr block of the desired VPC.
    :param str enterprise_project_id: Specifies the enterprise project ID which the desired VPC belongs to.
    :param str id: Specifies the id of the desired VPC.
    :param str name: Specifies the name of the desired VPC. The value is a string of no more than 64 characters
           and can contain digits, letters, underscores (_) and hyphens (-).
    :param str region: Specifies the region in which to obtain the VPC. If omitted, the provider-level region
           will be used.
    :param str status: Specifies the current status of the desired VPC. The value can be CREATING, OK or ERROR.
    :param Mapping[str, str] tags: Specifies the included key/value pairs which associated with the desired VPC.
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['enterpriseProjectId'] = enterprise_project_id
    __args__['id'] = id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    __args__['tags'] = tags
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Vpc/getVpcs:getVpcs', __args__, opts=opts, typ=GetVpcsResult).value

    return AwaitableGetVpcsResult(
        cidr=__ret__.cidr,
        enterprise_project_id=__ret__.enterprise_project_id,
        id=__ret__.id,
        name=__ret__.name,
        region=__ret__.region,
        status=__ret__.status,
        tags=__ret__.tags,
        vpcs=__ret__.vpcs)


@_utilities.lift_output_func(get_vpcs)
def get_vpcs_output(cidr: Optional[pulumi.Input[Optional[str]]] = None,
                    enterprise_project_id: Optional[pulumi.Input[Optional[str]]] = None,
                    id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[Optional[str]]] = None,
                    region: Optional[pulumi.Input[Optional[str]]] = None,
                    status: Optional[pulumi.Input[Optional[str]]] = None,
                    tags: Optional[pulumi.Input[Optional[Mapping[str, str]]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcsResult]:
    """
    Use this data source to get a list of VPC.

    ## Example Usage

    An example filter by name and tag

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    vpc_name = config.require_object("vpcName")
    vpc = huaweicloud.Vpc.get_vpcs(name=vpc_name,
        tags={
            "foo": "bar",
        })
    pulumi.export("vpcIds", [__item.id for __item in [vpc.vpcs]])
    ```


    :param str cidr: Specifies the cidr block of the desired VPC.
    :param str enterprise_project_id: Specifies the enterprise project ID which the desired VPC belongs to.
    :param str id: Specifies the id of the desired VPC.
    :param str name: Specifies the name of the desired VPC. The value is a string of no more than 64 characters
           and can contain digits, letters, underscores (_) and hyphens (-).
    :param str region: Specifies the region in which to obtain the VPC. If omitted, the provider-level region
           will be used.
    :param str status: Specifies the current status of the desired VPC. The value can be CREATING, OK or ERROR.
    :param Mapping[str, str] tags: Specifies the included key/value pairs which associated with the desired VPC.
    """
    ...

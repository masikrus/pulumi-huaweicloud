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
    'GetFlaovrsResult',
    'AwaitableGetFlaovrsResult',
    'get_flaovrs',
    'get_flaovrs_output',
]

@pulumi.output_type
class GetFlaovrsResult:
    """
    A collection of values returned by getFlaovrs.
    """
    def __init__(__self__, availability_zone=None, datastore_type=None, flavors=None, id=None, memory=None, region=None, vcpus=None):
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if datastore_type and not isinstance(datastore_type, str):
            raise TypeError("Expected argument 'datastore_type' to be a str")
        pulumi.set(__self__, "datastore_type", datastore_type)
        if flavors and not isinstance(flavors, list):
            raise TypeError("Expected argument 'flavors' to be a list")
        pulumi.set(__self__, "flavors", flavors)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if memory and not isinstance(memory, int):
            raise TypeError("Expected argument 'memory' to be a int")
        pulumi.set(__self__, "memory", memory)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if vcpus and not isinstance(vcpus, int):
            raise TypeError("Expected argument 'vcpus' to be a int")
        pulumi.set(__self__, "vcpus", vcpus)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[str]:
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter(name="datastoreType")
    def datastore_type(self) -> Optional[str]:
        """
        The type of datastore.  
        The options are as follows:
        - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
        - **hybrid**: a single data warehouse used for transaction and analytics workloads,
        in single-node or cluster mode.
        - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
        """
        return pulumi.get(self, "datastore_type")

    @property
    @pulumi.getter
    def flavors(self) -> Sequence['outputs.GetFlaovrsFlavorResult']:
        """
        The list of flavor detail.
        The Flavors structure is documented below.
        """
        return pulumi.get(self, "flavors")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def memory(self) -> Optional[int]:
        """
        The ram of the dws node flavor in GB.
        """
        return pulumi.get(self, "memory")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def vcpus(self) -> Optional[int]:
        """
        The vcpus of the dws node flavor.
        """
        return pulumi.get(self, "vcpus")


class AwaitableGetFlaovrsResult(GetFlaovrsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFlaovrsResult(
            availability_zone=self.availability_zone,
            datastore_type=self.datastore_type,
            flavors=self.flavors,
            id=self.id,
            memory=self.memory,
            region=self.region,
            vcpus=self.vcpus)


def get_flaovrs(availability_zone: Optional[str] = None,
                datastore_type: Optional[str] = None,
                memory: Optional[int] = None,
                region: Optional[str] = None,
                vcpus: Optional[int] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFlaovrsResult:
    """
    Use this data source to get available flavors of DWS cluster node.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavor = huaweicloud.Dws.get_flaovrs(vcpus=8)
    ```


    :param str availability_zone: The availability zone name.
    :param str datastore_type: The type of datastore.  
           The options are as follows:
           - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
           - **hybrid**: a single data warehouse used for transaction and analytics workloads,
           in single-node or cluster mode.
           - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
    :param int memory: The ram of the dws node flavor in GB.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param int vcpus: The vcpus of the dws node flavor.
    """
    __args__ = dict()
    __args__['availabilityZone'] = availability_zone
    __args__['datastoreType'] = datastore_type
    __args__['memory'] = memory
    __args__['region'] = region
    __args__['vcpus'] = vcpus
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Dws/getFlaovrs:getFlaovrs', __args__, opts=opts, typ=GetFlaovrsResult).value

    return AwaitableGetFlaovrsResult(
        availability_zone=__ret__.availability_zone,
        datastore_type=__ret__.datastore_type,
        flavors=__ret__.flavors,
        id=__ret__.id,
        memory=__ret__.memory,
        region=__ret__.region,
        vcpus=__ret__.vcpus)


@_utilities.lift_output_func(get_flaovrs)
def get_flaovrs_output(availability_zone: Optional[pulumi.Input[Optional[str]]] = None,
                       datastore_type: Optional[pulumi.Input[Optional[str]]] = None,
                       memory: Optional[pulumi.Input[Optional[int]]] = None,
                       region: Optional[pulumi.Input[Optional[str]]] = None,
                       vcpus: Optional[pulumi.Input[Optional[int]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFlaovrsResult]:
    """
    Use this data source to get available flavors of DWS cluster node.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    flavor = huaweicloud.Dws.get_flaovrs(vcpus=8)
    ```


    :param str availability_zone: The availability zone name.
    :param str datastore_type: The type of datastore.  
           The options are as follows:
           - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
           - **hybrid**: a single data warehouse used for transaction and analytics workloads,
           in single-node or cluster mode.
           - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
    :param int memory: The ram of the dws node flavor in GB.
    :param str region: Specifies the region in which to query the data source.
           If omitted, the provider-level region will be used.
    :param int vcpus: The vcpus of the dws node flavor.
    """
    ...

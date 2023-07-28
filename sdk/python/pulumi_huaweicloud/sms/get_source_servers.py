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
    'GetSourceServersResult',
    'AwaitableGetSourceServersResult',
    'get_source_servers',
    'get_source_servers_output',
]

@pulumi.output_type
class GetSourceServersResult:
    """
    A collection of values returned by getSourceServers.
    """
    def __init__(__self__, id=None, ip=None, name=None, servers=None, state=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip and not isinstance(ip, str):
            raise TypeError("Expected argument 'ip' to be a str")
        pulumi.set(__self__, "ip", ip)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if servers and not isinstance(servers, list):
            raise TypeError("Expected argument 'servers' to be a list")
        pulumi.set(__self__, "servers", servers)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The ID of the source server.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ip(self) -> Optional[str]:
        """
        The IP address of the source server.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        """
        The disk name, for example, /dev/vda.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def servers(self) -> Sequence['outputs.GetSourceServersServerResult']:
        """
        An array of SMS source servers found. Structure is documented below.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter
    def state(self) -> Optional[str]:
        """
        The status of the source server.
        """
        return pulumi.get(self, "state")


class AwaitableGetSourceServersResult(GetSourceServersResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSourceServersResult(
            id=self.id,
            ip=self.ip,
            name=self.name,
            servers=self.servers,
            state=self.state)


def get_source_servers(id: Optional[str] = None,
                       ip: Optional[str] = None,
                       name: Optional[str] = None,
                       state: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSourceServersResult:
    """
    Use this data source to get a list of SMS source servers.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    server_name = config.require_object("serverName")
    demo = huaweicloud.Sms.get_source_servers(name=server_name)
    ```


    :param str id: Specifies the ID of the source server.
    :param str ip: Specifies the IP address of the source server.
    :param str name: Specifies the name of the source server.
    :param str state: Specifies the status of the source server.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['ip'] = ip
    __args__['name'] = name
    __args__['state'] = state
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Sms/getSourceServers:getSourceServers', __args__, opts=opts, typ=GetSourceServersResult).value

    return AwaitableGetSourceServersResult(
        id=__ret__.id,
        ip=__ret__.ip,
        name=__ret__.name,
        servers=__ret__.servers,
        state=__ret__.state)


@_utilities.lift_output_func(get_source_servers)
def get_source_servers_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                              ip: Optional[pulumi.Input[Optional[str]]] = None,
                              name: Optional[pulumi.Input[Optional[str]]] = None,
                              state: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSourceServersResult]:
    """
    Use this data source to get a list of SMS source servers.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    server_name = config.require_object("serverName")
    demo = huaweicloud.Sms.get_source_servers(name=server_name)
    ```


    :param str id: Specifies the ID of the source server.
    :param str ip: Specifies the IP address of the source server.
    :param str name: Specifies the name of the source server.
    :param str state: Specifies the status of the source server.
    """
    ...

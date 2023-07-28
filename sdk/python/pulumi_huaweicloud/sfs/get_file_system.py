# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFileSystemResult',
    'AwaitableGetFileSystemResult',
    'get_file_system',
    'get_file_system_output',
]

@pulumi.output_type
class GetFileSystemResult:
    """
    A collection of values returned by getFileSystem.
    """
    def __init__(__self__, access_level=None, access_to=None, access_type=None, availability_zone=None, description=None, export_location=None, id=None, is_public=None, metadata=None, mount_id=None, name=None, preferred=None, region=None, share_access_id=None, share_instance_id=None, share_proto=None, size=None, state=None, status=None):
        if access_level and not isinstance(access_level, str):
            raise TypeError("Expected argument 'access_level' to be a str")
        pulumi.set(__self__, "access_level", access_level)
        if access_to and not isinstance(access_to, str):
            raise TypeError("Expected argument 'access_to' to be a str")
        pulumi.set(__self__, "access_to", access_to)
        if access_type and not isinstance(access_type, str):
            raise TypeError("Expected argument 'access_type' to be a str")
        pulumi.set(__self__, "access_type", access_type)
        if availability_zone and not isinstance(availability_zone, str):
            raise TypeError("Expected argument 'availability_zone' to be a str")
        pulumi.set(__self__, "availability_zone", availability_zone)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if export_location and not isinstance(export_location, str):
            raise TypeError("Expected argument 'export_location' to be a str")
        pulumi.set(__self__, "export_location", export_location)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_public and not isinstance(is_public, bool):
            raise TypeError("Expected argument 'is_public' to be a bool")
        pulumi.set(__self__, "is_public", is_public)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if mount_id and not isinstance(mount_id, str):
            raise TypeError("Expected argument 'mount_id' to be a str")
        pulumi.set(__self__, "mount_id", mount_id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if preferred and not isinstance(preferred, bool):
            raise TypeError("Expected argument 'preferred' to be a bool")
        pulumi.set(__self__, "preferred", preferred)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if share_access_id and not isinstance(share_access_id, str):
            raise TypeError("Expected argument 'share_access_id' to be a str")
        pulumi.set(__self__, "share_access_id", share_access_id)
        if share_instance_id and not isinstance(share_instance_id, str):
            raise TypeError("Expected argument 'share_instance_id' to be a str")
        pulumi.set(__self__, "share_instance_id", share_instance_id)
        if share_proto and not isinstance(share_proto, str):
            raise TypeError("Expected argument 'share_proto' to be a str")
        pulumi.set(__self__, "share_proto", share_proto)
        if size and not isinstance(size, int):
            raise TypeError("Expected argument 'size' to be a int")
        pulumi.set(__self__, "size", size)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> str:
        """
        The level of the access rule.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="accessTo")
    def access_to(self) -> str:
        """
        The access that the back end grants or denies.
        """
        return pulumi.get(self, "access_to")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> str:
        """
        The type of the share access rule.
        """
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> str:
        """
        The availability zone name.
        """
        return pulumi.get(self, "availability_zone")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the shared file system.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="exportLocation")
    def export_location(self) -> str:
        """
        The path for accessing the shared file system.
        """
        return pulumi.get(self, "export_location")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isPublic")
    def is_public(self) -> bool:
        """
        Whether a file system can be publicly seen.
        """
        return pulumi.get(self, "is_public")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        The key and value pairs information of the shared file system.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mountId")
    def mount_id(self) -> str:
        """
        The UUID of the mount location of the shared file system.
        """
        return pulumi.get(self, "mount_id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def preferred(self) -> bool:
        """
        Identifies which mount locations are most efficient and are used preferentially when multiple mount
        locations exist.
        """
        return pulumi.get(self, "preferred")

    @property
    @pulumi.getter
    def region(self) -> str:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="shareAccessId")
    def share_access_id(self) -> str:
        """
        The UUID of the share access rule.
        """
        return pulumi.get(self, "share_access_id")

    @property
    @pulumi.getter(name="shareInstanceId")
    def share_instance_id(self) -> str:
        """
        The access that the back end grants or denies.
        """
        return pulumi.get(self, "share_instance_id")

    @property
    @pulumi.getter(name="shareProto")
    def share_proto(self) -> str:
        """
        The protocol for sharing file systems.
        """
        return pulumi.get(self, "share_proto")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The size (GB) of the shared file system.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state of the shared file system.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")


class AwaitableGetFileSystemResult(GetFileSystemResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFileSystemResult(
            access_level=self.access_level,
            access_to=self.access_to,
            access_type=self.access_type,
            availability_zone=self.availability_zone,
            description=self.description,
            export_location=self.export_location,
            id=self.id,
            is_public=self.is_public,
            metadata=self.metadata,
            mount_id=self.mount_id,
            name=self.name,
            preferred=self.preferred,
            region=self.region,
            share_access_id=self.share_access_id,
            share_instance_id=self.share_instance_id,
            share_proto=self.share_proto,
            size=self.size,
            state=self.state,
            status=self.status)


def get_file_system(id: Optional[str] = None,
                    name: Optional[str] = None,
                    region: Optional[str] = None,
                    status: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFileSystemResult:
    """
    Provides information about an Shared File System (SFS) within HuaweiCloud.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    share_name = config.require_object("shareName")
    shared_file = huaweicloud.Sfs.get_file_system(name=share_name)
    ```


    :param str id: The UUID of the shared file system.
    :param str name: The name of the shared file system.
    :param str region: Specifies the region in which to obtain the shared file system.
           If omitted, the provider-level region will be used.
    :param str status: The status of the shared file system.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['region'] = region
    __args__['status'] = status
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('huaweicloud:Sfs/getFileSystem:getFileSystem', __args__, opts=opts, typ=GetFileSystemResult).value

    return AwaitableGetFileSystemResult(
        access_level=__ret__.access_level,
        access_to=__ret__.access_to,
        access_type=__ret__.access_type,
        availability_zone=__ret__.availability_zone,
        description=__ret__.description,
        export_location=__ret__.export_location,
        id=__ret__.id,
        is_public=__ret__.is_public,
        metadata=__ret__.metadata,
        mount_id=__ret__.mount_id,
        name=__ret__.name,
        preferred=__ret__.preferred,
        region=__ret__.region,
        share_access_id=__ret__.share_access_id,
        share_instance_id=__ret__.share_instance_id,
        share_proto=__ret__.share_proto,
        size=__ret__.size,
        state=__ret__.state,
        status=__ret__.status)


@_utilities.lift_output_func(get_file_system)
def get_file_system_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                           name: Optional[pulumi.Input[Optional[str]]] = None,
                           region: Optional[pulumi.Input[Optional[str]]] = None,
                           status: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFileSystemResult]:
    """
    Provides information about an Shared File System (SFS) within HuaweiCloud.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_huaweicloud as huaweicloud

    config = pulumi.Config()
    share_name = config.require_object("shareName")
    shared_file = huaweicloud.Sfs.get_file_system(name=share_name)
    ```


    :param str id: The UUID of the shared file system.
    :param str name: The name of the shared file system.
    :param str region: Specifies the region in which to obtain the shared file system.
           If omitted, the provider-level region will be used.
    :param str status: The status of the shared file system.
    """
    ...

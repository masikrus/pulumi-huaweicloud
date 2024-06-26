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
    'MysqlInstanceBackupStrategyArgs',
    'MysqlInstanceDatastoreArgs',
    'MysqlInstanceNodeArgs',
]

@pulumi.input_type
class MysqlInstanceBackupStrategyArgs:
    def __init__(__self__, *,
                 start_time: pulumi.Input[str],
                 keep_days: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input[str] start_time: Specifies the backup time window. Automated backups will be triggered during the
               backup time window. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format. The
               HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to 00. Example
               value: 08:00-09:00, 03:00-04:00.
        :param pulumi.Input[int] keep_days: Specifies the number of days to retain the generated backup files. The value ranges from
               0 to 35. If this parameter is set to 0, the automated backup policy is not set. If this parameter is not transferred,
               the automated backup policy is enabled by default. Backup files are stored for seven days by default.
        """
        pulumi.set(__self__, "start_time", start_time)
        if keep_days is not None:
            pulumi.set(__self__, "keep_days", keep_days)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input[str]:
        """
        Specifies the backup time window. Automated backups will be triggered during the
        backup time window. It must be a valid value in the "hh:mm-HH:MM" format. The current time is in the UTC format. The
        HH value must be 1 greater than the hh value. The values of mm and MM must be the same and must be set to 00. Example
        value: 08:00-09:00, 03:00-04:00.
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter(name="keepDays")
    def keep_days(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of days to retain the generated backup files. The value ranges from
        0 to 35. If this parameter is set to 0, the automated backup policy is not set. If this parameter is not transferred,
        the automated backup policy is enabled by default. Backup files are stored for seven days by default.
        """
        return pulumi.get(self, "keep_days")

    @keep_days.setter
    def keep_days(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keep_days", value)


@pulumi.input_type
class MysqlInstanceDatastoreArgs:
    def __init__(__self__, *,
                 engine: pulumi.Input[str],
                 version: pulumi.Input[str]):
        """
        :param pulumi.Input[str] engine: Specifies the database engine. Only "gaussdb-mysql" is supported now.
               Changing this parameter will create a new resource.
        :param pulumi.Input[str] version: Specifies the database version. Only "8.0" is supported now.
               Changing this parameter will create a new resource.
        """
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        Specifies the database engine. Only "gaussdb-mysql" is supported now.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Specifies the database version. Only "8.0" is supported now.
        Changing this parameter will create a new resource.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class MysqlInstanceNodeArgs:
    def __init__(__self__, *,
                 availability_zone: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_read_ip: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] availability_zone: Indicates the availability zone where the node resides.
        :param pulumi.Input[str] id: Indicates the node ID.
        :param pulumi.Input[str] name: Specifies the instance name, which can be the same as an existing instance name. The value
               must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters,
               digits, hyphens (-), and underscores (_).
        :param pulumi.Input[str] private_read_ip: Indicates the private IP address of a node.
        :param pulumi.Input[str] status: Indicates the node status.
        :param pulumi.Input[str] type: Indicates the node type: master or slave.
        """
        if availability_zone is not None:
            pulumi.set(__self__, "availability_zone", availability_zone)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_read_ip is not None:
            pulumi.set(__self__, "private_read_ip", private_read_ip)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="availabilityZone")
    def availability_zone(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the availability zone where the node resides.
        """
        return pulumi.get(self, "availability_zone")

    @availability_zone.setter
    def availability_zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "availability_zone", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the node ID.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the instance name, which can be the same as an existing instance name. The value
        must be 4 to 64 characters in length and start with a letter. It is case-sensitive and can contain only letters,
        digits, hyphens (-), and underscores (_).
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateReadIp")
    def private_read_ip(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the private IP address of a node.
        """
        return pulumi.get(self, "private_read_ip")

    @private_read_ip.setter
    def private_read_ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_read_ip", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the node status.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates the node type: master or slave.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)



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
    'PolicyBackupCycle',
    'PolicyLongTermRetention',
    'VaultPolicy',
    'VaultResource',
    'GetVaultsVaultResult',
    'GetVaultsVaultResourceResult',
]

@pulumi.output_type
class PolicyBackupCycle(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "executionTimes":
            suggest = "execution_times"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyBackupCycle. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyBackupCycle.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyBackupCycle.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 execution_times: Sequence[str],
                 days: Optional[str] = None,
                 interval: Optional[int] = None):
        """
        :param Sequence[str] execution_times: Specifies the backup time. Automated backups will be triggered at the backup
               time. The current time is in the UTC format (HH:MM). The minutes in the list must be set to **00** and the hours
               cannot be repeated. In the replication policy, you are advised to set one time point for one day.
        :param str days: Specifies the weekly backup day of backup schedule. It supports seven days a week (MO, TU,
               WE, TH, FR, SA, SU) and this parameter is separated by a comma (,) without spaces, between date and date during the
               configuration.
        :param int interval: Specifies the interval (in days) of backup schedule. The value range is `1` to `30`. This
               parameter and `days` are alternative.
        """
        pulumi.set(__self__, "execution_times", execution_times)
        if days is not None:
            pulumi.set(__self__, "days", days)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)

    @property
    @pulumi.getter(name="executionTimes")
    def execution_times(self) -> Sequence[str]:
        """
        Specifies the backup time. Automated backups will be triggered at the backup
        time. The current time is in the UTC format (HH:MM). The minutes in the list must be set to **00** and the hours
        cannot be repeated. In the replication policy, you are advised to set one time point for one day.
        """
        return pulumi.get(self, "execution_times")

    @property
    @pulumi.getter
    def days(self) -> Optional[str]:
        """
        Specifies the weekly backup day of backup schedule. It supports seven days a week (MO, TU,
        WE, TH, FR, SA, SU) and this parameter is separated by a comma (,) without spaces, between date and date during the
        configuration.
        """
        return pulumi.get(self, "days")

    @property
    @pulumi.getter
    def interval(self) -> Optional[int]:
        """
        Specifies the interval (in days) of backup schedule. The value range is `1` to `30`. This
        parameter and `days` are alternative.
        """
        return pulumi.get(self, "interval")


@pulumi.output_type
class PolicyLongTermRetention(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "fullBackupInterval":
            suggest = "full_backup_interval"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PolicyLongTermRetention. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PolicyLongTermRetention.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PolicyLongTermRetention.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 daily: Optional[int] = None,
                 full_backup_interval: Optional[int] = None,
                 monthly: Optional[int] = None,
                 weekly: Optional[int] = None,
                 yearly: Optional[int] = None):
        """
        :param int daily: - Specifies the latest backup of each day is saved in the long term.
        :param int full_backup_interval: Specifies how often (after how many incremental backups) a full backup is
               performed. The valid value ranges from `-1` to `100`.
               If `-1` is specified, full backup will not be performed.
        :param int monthly: - Specifies the latest backup of each month is saved in the long term.
        :param int weekly: - Specifies the latest backup of each week is saved in the long term.
        :param int yearly: - Specifies the latest backup of each year is saved in the long term.
        """
        if daily is not None:
            pulumi.set(__self__, "daily", daily)
        if full_backup_interval is not None:
            pulumi.set(__self__, "full_backup_interval", full_backup_interval)
        if monthly is not None:
            pulumi.set(__self__, "monthly", monthly)
        if weekly is not None:
            pulumi.set(__self__, "weekly", weekly)
        if yearly is not None:
            pulumi.set(__self__, "yearly", yearly)

    @property
    @pulumi.getter
    def daily(self) -> Optional[int]:
        """
        - Specifies the latest backup of each day is saved in the long term.
        """
        return pulumi.get(self, "daily")

    @property
    @pulumi.getter(name="fullBackupInterval")
    def full_backup_interval(self) -> Optional[int]:
        """
        Specifies how often (after how many incremental backups) a full backup is
        performed. The valid value ranges from `-1` to `100`.
        If `-1` is specified, full backup will not be performed.
        """
        return pulumi.get(self, "full_backup_interval")

    @property
    @pulumi.getter
    def monthly(self) -> Optional[int]:
        """
        - Specifies the latest backup of each month is saved in the long term.
        """
        return pulumi.get(self, "monthly")

    @property
    @pulumi.getter
    def weekly(self) -> Optional[int]:
        """
        - Specifies the latest backup of each week is saved in the long term.
        """
        return pulumi.get(self, "weekly")

    @property
    @pulumi.getter
    def yearly(self) -> Optional[int]:
        """
        - Specifies the latest backup of each year is saved in the long term.
        """
        return pulumi.get(self, "yearly")


@pulumi.output_type
class VaultPolicy(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "destinationVaultId":
            suggest = "destination_vault_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VaultPolicy. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VaultPolicy.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VaultPolicy.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 id: str,
                 destination_vault_id: Optional[str] = None):
        """
        :param str id: Specifies the policy ID.
        :param str destination_vault_id: Specifies the ID of destination vault to which the replication policy
               will associated.
        """
        pulumi.set(__self__, "id", id)
        if destination_vault_id is not None:
            pulumi.set(__self__, "destination_vault_id", destination_vault_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Specifies the policy ID.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="destinationVaultId")
    def destination_vault_id(self) -> Optional[str]:
        """
        Specifies the ID of destination vault to which the replication policy
        will associated.
        """
        return pulumi.get(self, "destination_vault_id")


@pulumi.output_type
class VaultResource(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "serverId":
            suggest = "server_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in VaultResource. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        VaultResource.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        VaultResource.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 excludes: Optional[Sequence[str]] = None,
                 includes: Optional[Sequence[str]] = None,
                 server_id: Optional[str] = None):
        """
        :param Sequence[str] excludes: Specifies the array of disk IDs which will be excluded in the backup.
               Only **server** vault support this parameter.
        :param Sequence[str] includes: Specifies the array of disk or SFS file system IDs which will be included in the backup.
               Only **disk** and **turbo** vault support this parameter.
        :param str server_id: Specifies the ID of the ECS instance to be backed up.
        """
        if excludes is not None:
            pulumi.set(__self__, "excludes", excludes)
        if includes is not None:
            pulumi.set(__self__, "includes", includes)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter
    def excludes(self) -> Optional[Sequence[str]]:
        """
        Specifies the array of disk IDs which will be excluded in the backup.
        Only **server** vault support this parameter.
        """
        return pulumi.get(self, "excludes")

    @property
    @pulumi.getter
    def includes(self) -> Optional[Sequence[str]]:
        """
        Specifies the array of disk or SFS file system IDs which will be included in the backup.
        Only **disk** and **turbo** vault support this parameter.
        """
        return pulumi.get(self, "includes")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[str]:
        """
        Specifies the ID of the ECS instance to be backed up.
        """
        return pulumi.get(self, "server_id")


@pulumi.output_type
class GetVaultsVaultResult(dict):
    def __init__(__self__, *,
                 allocated: float,
                 auto_expand_enabled: bool,
                 consistent_level: str,
                 enterprise_project_id: str,
                 id: str,
                 name: str,
                 policy_id: str,
                 protection_type: str,
                 resources: Sequence['outputs.GetVaultsVaultResourceResult'],
                 size: int,
                 spec_code: str,
                 status: str,
                 storage: str,
                 tags: Mapping[str, str],
                 type: str,
                 used: float):
        """
        :param float allocated: The allocated capacity of the vault, in GB.
        :param bool auto_expand_enabled: Specifies whether to enable automatic expansion of the backup protection
               type vault. Defaults to **false**.
        :param str consistent_level: Specifies the consistent level (specification) of the vault.
               The valid values are as follows:
               + **[crash_consistent](https://support.huaweicloud.com/intl/en-us/usermanual-cbr/cbr_03_0109.html)**
               + **[app_consistent](https://support.huaweicloud.com/intl/en-us/usermanual-cbr/cbr_03_0109.html)**
        :param str enterprise_project_id: Specifies the ID of the enterprise project to which the vault belongs.
        :param str id: The vault ID in UUID format.
        :param str name: Specifies the vault name. This parameter can contain a maximum of 64
               characters, which may consist of letters, digits, underscores(_) and hyphens (-).
        :param str policy_id: Specifies the ID of the policy associated with the vault.
               The `policy_id` cannot be used with the vault of replicate protection type.
        :param str protection_type: Specifies the protection type of the vault.
               The valid values are **backup** and **replication**. Vaults of type **disk** don't support **replication**.
        :param Sequence['GetVaultsVaultResourceArgs'] resources: The array of one or more resources to attach to the vault.
               The object structure is documented below.
        :param int size: Specifies the vault sapacity, in GB. The valid value range is `1` to `10,485,760`.
        :param str spec_code: The specification code.
        :param str status: Specifies the vault status, including **available**, **lock**, **frozen** and **error**.
        :param str storage: The name of the bucket for the vault.
        :param Mapping[str, str] tags: The key/value pairs to associate with the vault.
        :param str type: Specifies the object type of the vault. The vaild values are as follows:
               + **server** (Cloud Servers)
               + **disk** (EVS Disks)
               + **turbo** (SFS Turbo file systems)
        :param float used: The used capacity, in GB.
        """
        pulumi.set(__self__, "allocated", allocated)
        pulumi.set(__self__, "auto_expand_enabled", auto_expand_enabled)
        pulumi.set(__self__, "consistent_level", consistent_level)
        pulumi.set(__self__, "enterprise_project_id", enterprise_project_id)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "policy_id", policy_id)
        pulumi.set(__self__, "protection_type", protection_type)
        pulumi.set(__self__, "resources", resources)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "spec_code", spec_code)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "storage", storage)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "used", used)

    @property
    @pulumi.getter
    def allocated(self) -> float:
        """
        The allocated capacity of the vault, in GB.
        """
        return pulumi.get(self, "allocated")

    @property
    @pulumi.getter(name="autoExpandEnabled")
    def auto_expand_enabled(self) -> bool:
        """
        Specifies whether to enable automatic expansion of the backup protection
        type vault. Defaults to **false**.
        """
        return pulumi.get(self, "auto_expand_enabled")

    @property
    @pulumi.getter(name="consistentLevel")
    def consistent_level(self) -> str:
        """
        Specifies the consistent level (specification) of the vault.
        The valid values are as follows:
        + **[crash_consistent](https://support.huaweicloud.com/intl/en-us/usermanual-cbr/cbr_03_0109.html)**
        + **[app_consistent](https://support.huaweicloud.com/intl/en-us/usermanual-cbr/cbr_03_0109.html)**
        """
        return pulumi.get(self, "consistent_level")

    @property
    @pulumi.getter(name="enterpriseProjectId")
    def enterprise_project_id(self) -> str:
        """
        Specifies the ID of the enterprise project to which the vault belongs.
        """
        return pulumi.get(self, "enterprise_project_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The vault ID in UUID format.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Specifies the vault name. This parameter can contain a maximum of 64
        characters, which may consist of letters, digits, underscores(_) and hyphens (-).
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> str:
        """
        Specifies the ID of the policy associated with the vault.
        The `policy_id` cannot be used with the vault of replicate protection type.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="protectionType")
    def protection_type(self) -> str:
        """
        Specifies the protection type of the vault.
        The valid values are **backup** and **replication**. Vaults of type **disk** don't support **replication**.
        """
        return pulumi.get(self, "protection_type")

    @property
    @pulumi.getter
    def resources(self) -> Sequence['outputs.GetVaultsVaultResourceResult']:
        """
        The array of one or more resources to attach to the vault.
        The object structure is documented below.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        Specifies the vault sapacity, in GB. The valid value range is `1` to `10,485,760`.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="specCode")
    def spec_code(self) -> str:
        """
        The specification code.
        """
        return pulumi.get(self, "spec_code")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Specifies the vault status, including **available**, **lock**, **frozen** and **error**.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storage(self) -> str:
        """
        The name of the bucket for the vault.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def tags(self) -> Mapping[str, str]:
        """
        The key/value pairs to associate with the vault.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Specifies the object type of the vault. The vaild values are as follows:
        + **server** (Cloud Servers)
        + **disk** (EVS Disks)
        + **turbo** (SFS Turbo file systems)
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def used(self) -> float:
        """
        The used capacity, in GB.
        """
        return pulumi.get(self, "used")


@pulumi.output_type
class GetVaultsVaultResourceResult(dict):
    def __init__(__self__, *,
                 excludes: Sequence[str],
                 includes: Sequence[str],
                 server_id: str):
        """
        :param Sequence[str] excludes: The array of disk IDs which will be excluded in the backup.
        :param Sequence[str] includes: The array of disk or SFS file system IDs which will be included in the backup.
        :param str server_id: The ID of the ECS instance to be backed up.
        """
        pulumi.set(__self__, "excludes", excludes)
        pulumi.set(__self__, "includes", includes)
        pulumi.set(__self__, "server_id", server_id)

    @property
    @pulumi.getter
    def excludes(self) -> Sequence[str]:
        """
        The array of disk IDs which will be excluded in the backup.
        """
        return pulumi.get(self, "excludes")

    @property
    @pulumi.getter
    def includes(self) -> Sequence[str]:
        """
        The array of disk or SFS file system IDs which will be included in the backup.
        """
        return pulumi.get(self, "includes")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> str:
        """
        The ID of the ECS instance to be backed up.
        """
        return pulumi.get(self, "server_id")



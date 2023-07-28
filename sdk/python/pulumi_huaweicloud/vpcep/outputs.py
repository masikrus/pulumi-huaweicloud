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
    'ApprovalConnection',
    'ServiceConnection',
    'ServicePortMapping',
    'GetPublicServicesServiceResult',
]

@pulumi.output_type
class ApprovalConnection(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "domainId":
            suggest = "domain_id"
        elif key == "endpointId":
            suggest = "endpoint_id"
        elif key == "packetId":
            suggest = "packet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ApprovalConnection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ApprovalConnection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ApprovalConnection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 domain_id: Optional[str] = None,
                 endpoint_id: Optional[str] = None,
                 packet_id: Optional[int] = None,
                 status: Optional[str] = None):
        """
        :param str description: The description of the VPC endpoint service connection.
        :param str domain_id: The user's domain ID.
        :param str endpoint_id: The unique ID of the VPC endpoint.
        :param int packet_id: The packet ID of the VPC endpoint.
        :param str status: The connection status of the VPC endpoint.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if packet_id is not None:
            pulumi.set(__self__, "packet_id", packet_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        The description of the VPC endpoint service connection.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        """
        The user's domain ID.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[str]:
        """
        The unique ID of the VPC endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="packetId")
    def packet_id(self) -> Optional[int]:
        """
        The packet ID of the VPC endpoint.
        """
        return pulumi.get(self, "packet_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The connection status of the VPC endpoint.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ServiceConnection(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "domainId":
            suggest = "domain_id"
        elif key == "endpointId":
            suggest = "endpoint_id"
        elif key == "packetId":
            suggest = "packet_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServiceConnection. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServiceConnection.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServiceConnection.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 description: Optional[str] = None,
                 domain_id: Optional[str] = None,
                 endpoint_id: Optional[str] = None,
                 packet_id: Optional[int] = None,
                 status: Optional[str] = None):
        """
        :param str description: Specifies the description of the VPC endpoint service.
        :param str domain_id: The user's domain ID.
        :param str endpoint_id: The unique ID of the VPC endpoint.
        :param int packet_id: The packet ID of the VPC endpoint.
        :param str status: The connection status of the VPC endpoint.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain_id is not None:
            pulumi.set(__self__, "domain_id", domain_id)
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if packet_id is not None:
            pulumi.set(__self__, "packet_id", packet_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Specifies the description of the VPC endpoint service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="domainId")
    def domain_id(self) -> Optional[str]:
        """
        The user's domain ID.
        """
        return pulumi.get(self, "domain_id")

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[str]:
        """
        The unique ID of the VPC endpoint.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="packetId")
    def packet_id(self) -> Optional[int]:
        """
        The packet ID of the VPC endpoint.
        """
        return pulumi.get(self, "packet_id")

    @property
    @pulumi.getter
    def status(self) -> Optional[str]:
        """
        The connection status of the VPC endpoint.
        """
        return pulumi.get(self, "status")


@pulumi.output_type
class ServicePortMapping(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "servicePort":
            suggest = "service_port"
        elif key == "terminalPort":
            suggest = "terminal_port"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ServicePortMapping. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ServicePortMapping.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ServicePortMapping.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 protocol: Optional[str] = None,
                 service_port: Optional[int] = None,
                 terminal_port: Optional[int] = None):
        """
        :param str protocol: Specifies the protocol used in port mappings. Only **TCP** is supported.
        :param int service_port: Specifies the port for accessing the VPC endpoint service. This port is provided by
               the backend service to provide services. The value ranges from 1 to 65535.
        :param int terminal_port: Specifies the port for accessing the VPC endpoint. This port is provided by the VPC
               endpoint, allowing you to access the VPC endpoint service. The value ranges from 1 to 65535.
        """
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if service_port is not None:
            pulumi.set(__self__, "service_port", service_port)
        if terminal_port is not None:
            pulumi.set(__self__, "terminal_port", terminal_port)

    @property
    @pulumi.getter
    def protocol(self) -> Optional[str]:
        """
        Specifies the protocol used in port mappings. Only **TCP** is supported.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="servicePort")
    def service_port(self) -> Optional[int]:
        """
        Specifies the port for accessing the VPC endpoint service. This port is provided by
        the backend service to provide services. The value ranges from 1 to 65535.
        """
        return pulumi.get(self, "service_port")

    @property
    @pulumi.getter(name="terminalPort")
    def terminal_port(self) -> Optional[int]:
        """
        Specifies the port for accessing the VPC endpoint. This port is provided by the VPC
        endpoint, allowing you to access the VPC endpoint service. The value ranges from 1 to 65535.
        """
        return pulumi.get(self, "terminal_port")


@pulumi.output_type
class GetPublicServicesServiceResult(dict):
    def __init__(__self__, *,
                 id: str,
                 is_charge: bool,
                 owner: str,
                 service_name: str,
                 service_type: str):
        """
        :param str id: The unique ID of the public VPC endpoint service.
        :param bool is_charge: Indicates whether the associated VPC endpoint carries a charge.
        :param str owner: The owner of the VPC endpoint service.
        :param str service_name: Specifies the name of the public VPC endpoint service. The value is not
               case-sensitive and supports fuzzy match.
        :param str service_type: The type of the VPC endpoint service.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "is_charge", is_charge)
        pulumi.set(__self__, "owner", owner)
        pulumi.set(__self__, "service_name", service_name)
        pulumi.set(__self__, "service_type", service_type)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The unique ID of the public VPC endpoint service.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isCharge")
    def is_charge(self) -> bool:
        """
        Indicates whether the associated VPC endpoint carries a charge.
        """
        return pulumi.get(self, "is_charge")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        The owner of the VPC endpoint service.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Specifies the name of the public VPC endpoint service. The value is not
        case-sensitive and supports fuzzy match.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="serviceType")
    def service_type(self) -> str:
        """
        The type of the VPC endpoint service.
        """
        return pulumi.get(self, "service_type")



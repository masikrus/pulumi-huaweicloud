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
    'DedicatedDomainServerArgs',
    'DomainServerArgs',
    'PolicyOptionArgs',
]

@pulumi.input_type
class DedicatedDomainServerArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 client_protocol: pulumi.Input[str],
                 port: pulumi.Input[int],
                 server_protocol: pulumi.Input[str],
                 type: pulumi.Input[str],
                 vpc_id: pulumi.Input[str]):
        """
        :param pulumi.Input[str] address: IP address or domain name of the web server that the client accesses. For
               example, `192.168.1.1` or `www.example.com`. Changing this creates a new service.
        :param pulumi.Input[str] client_protocol: Protocol type of the client. The options include `HTTP` and `HTTPS`.
               Changing this creates a new service.
        :param pulumi.Input[int] port: Port number used by the web server. The value ranges from 0 to 65535. Changing this
               creates a new service.
        :param pulumi.Input[str] server_protocol: Protocol used by WAF to forward client requests to the server. The
               options include `HTTP` and `HTTPS`. Changing this creates a new service.
        :param pulumi.Input[str] type: Server network type, IPv4 or IPv6. Valid values are: `ipv4` and `ipv6`. Changing
               this creates a new service.
        :param pulumi.Input[str] vpc_id: The id of the vpc used by the server. Changing this creates a service.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "client_protocol", client_protocol)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "server_protocol", server_protocol)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        IP address or domain name of the web server that the client accesses. For
        example, `192.168.1.1` or `www.example.com`. Changing this creates a new service.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> pulumi.Input[str]:
        """
        Protocol type of the client. The options include `HTTP` and `HTTPS`.
        Changing this creates a new service.
        """
        return pulumi.get(self, "client_protocol")

    @client_protocol.setter
    def client_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_protocol", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Port number used by the web server. The value ranges from 0 to 65535. Changing this
        creates a new service.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverProtocol")
    def server_protocol(self) -> pulumi.Input[str]:
        """
        Protocol used by WAF to forward client requests to the server. The
        options include `HTTP` and `HTTPS`. Changing this creates a new service.
        """
        return pulumi.get(self, "server_protocol")

    @server_protocol.setter
    def server_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_protocol", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Server network type, IPv4 or IPv6. Valid values are: `ipv4` and `ipv6`. Changing
        this creates a new service.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The id of the vpc used by the server. Changing this creates a service.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)


@pulumi.input_type
class DomainServerArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 client_protocol: pulumi.Input[str],
                 port: pulumi.Input[int],
                 server_protocol: pulumi.Input[str]):
        """
        :param pulumi.Input[str] address: IP address or domain name of the web server that the client accesses. For example,
               `192.168.1.1` or `www.a.com`.
        :param pulumi.Input[str] client_protocol: Protocol type of the client. The options include `HTTP` and `HTTPS`.
        :param pulumi.Input[int] port: Port number used by the web server. The value ranges from 0 to 65535, for example, 8080.
        :param pulumi.Input[str] server_protocol: Protocol used by WAF to forward client requests to the server. The options
               include `HTTP` and `HTTPS`.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "client_protocol", client_protocol)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "server_protocol", server_protocol)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        IP address or domain name of the web server that the client accesses. For example,
        `192.168.1.1` or `www.a.com`.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="clientProtocol")
    def client_protocol(self) -> pulumi.Input[str]:
        """
        Protocol type of the client. The options include `HTTP` and `HTTPS`.
        """
        return pulumi.get(self, "client_protocol")

    @client_protocol.setter
    def client_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_protocol", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        Port number used by the web server. The value ranges from 0 to 65535, for example, 8080.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverProtocol")
    def server_protocol(self) -> pulumi.Input[str]:
        """
        Protocol used by WAF to forward client requests to the server. The options
        include `HTTP` and `HTTPS`.
        """
        return pulumi.get(self, "server_protocol")

    @server_protocol.setter
    def server_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_protocol", value)


@pulumi.input_type
class PolicyOptionArgs:
    def __init__(__self__, *,
                 basic_web_protection: Optional[pulumi.Input[bool]] = None,
                 blacklist: Optional[pulumi.Input[bool]] = None,
                 cc_attack_protection: Optional[pulumi.Input[bool]] = None,
                 crawler: Optional[pulumi.Input[bool]] = None,
                 crawler_engine: Optional[pulumi.Input[bool]] = None,
                 crawler_other: Optional[pulumi.Input[bool]] = None,
                 crawler_scanner: Optional[pulumi.Input[bool]] = None,
                 crawler_script: Optional[pulumi.Input[bool]] = None,
                 data_masking: Optional[pulumi.Input[bool]] = None,
                 false_alarm_masking: Optional[pulumi.Input[bool]] = None,
                 general_check: Optional[pulumi.Input[bool]] = None,
                 precise_protection: Optional[pulumi.Input[bool]] = None,
                 web_tamper_protection: Optional[pulumi.Input[bool]] = None,
                 webshell: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[bool] basic_web_protection: Indicates whether Basic Web Protection is enabled.
        :param pulumi.Input[bool] blacklist: Indicates whether Blacklist and Whitelist is enabled.
        :param pulumi.Input[bool] cc_attack_protection: Indicates whether CC Attack Protection is enabled.
        :param pulumi.Input[bool] crawler: Indicates whether the master crawler detection switch in Basic Web Protection is enabled.
        :param pulumi.Input[bool] crawler_engine: Indicates whether the Search Engine switch in Basic Web Protection is enabled.
        :param pulumi.Input[bool] crawler_other: Indicates whether detection of other crawlers in Basic Web Protection is enabled.
        :param pulumi.Input[bool] crawler_scanner: Indicates whether the Scanner switch in Basic Web Protection is enabled.
        :param pulumi.Input[bool] crawler_script: Indicates whether the Script Tool switch in Basic Web Protection is enabled.
        :param pulumi.Input[bool] data_masking: Indicates whether Data Masking is enabled.
        :param pulumi.Input[bool] false_alarm_masking: Indicates whether False Alarm Masking is enabled.
        :param pulumi.Input[bool] general_check: Indicates whether General Check in Basic Web Protection is enabled.
        :param pulumi.Input[bool] precise_protection: Indicates whether Precise Protection is enabled.
        :param pulumi.Input[bool] web_tamper_protection: Indicates whether Web Tamper Protection is enabled.
        :param pulumi.Input[bool] webshell: Indicates whether webshell detection in Basic Web Protection is enabled.
        """
        if basic_web_protection is not None:
            pulumi.set(__self__, "basic_web_protection", basic_web_protection)
        if blacklist is not None:
            pulumi.set(__self__, "blacklist", blacklist)
        if cc_attack_protection is not None:
            pulumi.set(__self__, "cc_attack_protection", cc_attack_protection)
        if crawler is not None:
            pulumi.set(__self__, "crawler", crawler)
        if crawler_engine is not None:
            pulumi.set(__self__, "crawler_engine", crawler_engine)
        if crawler_other is not None:
            pulumi.set(__self__, "crawler_other", crawler_other)
        if crawler_scanner is not None:
            pulumi.set(__self__, "crawler_scanner", crawler_scanner)
        if crawler_script is not None:
            pulumi.set(__self__, "crawler_script", crawler_script)
        if data_masking is not None:
            pulumi.set(__self__, "data_masking", data_masking)
        if false_alarm_masking is not None:
            pulumi.set(__self__, "false_alarm_masking", false_alarm_masking)
        if general_check is not None:
            pulumi.set(__self__, "general_check", general_check)
        if precise_protection is not None:
            pulumi.set(__self__, "precise_protection", precise_protection)
        if web_tamper_protection is not None:
            pulumi.set(__self__, "web_tamper_protection", web_tamper_protection)
        if webshell is not None:
            pulumi.set(__self__, "webshell", webshell)

    @property
    @pulumi.getter(name="basicWebProtection")
    def basic_web_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Basic Web Protection is enabled.
        """
        return pulumi.get(self, "basic_web_protection")

    @basic_web_protection.setter
    def basic_web_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "basic_web_protection", value)

    @property
    @pulumi.getter
    def blacklist(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Blacklist and Whitelist is enabled.
        """
        return pulumi.get(self, "blacklist")

    @blacklist.setter
    def blacklist(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "blacklist", value)

    @property
    @pulumi.getter(name="ccAttackProtection")
    def cc_attack_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether CC Attack Protection is enabled.
        """
        return pulumi.get(self, "cc_attack_protection")

    @cc_attack_protection.setter
    def cc_attack_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cc_attack_protection", value)

    @property
    @pulumi.getter
    def crawler(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the master crawler detection switch in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "crawler")

    @crawler.setter
    def crawler(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "crawler", value)

    @property
    @pulumi.getter(name="crawlerEngine")
    def crawler_engine(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the Search Engine switch in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "crawler_engine")

    @crawler_engine.setter
    def crawler_engine(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "crawler_engine", value)

    @property
    @pulumi.getter(name="crawlerOther")
    def crawler_other(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether detection of other crawlers in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "crawler_other")

    @crawler_other.setter
    def crawler_other(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "crawler_other", value)

    @property
    @pulumi.getter(name="crawlerScanner")
    def crawler_scanner(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the Scanner switch in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "crawler_scanner")

    @crawler_scanner.setter
    def crawler_scanner(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "crawler_scanner", value)

    @property
    @pulumi.getter(name="crawlerScript")
    def crawler_script(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the Script Tool switch in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "crawler_script")

    @crawler_script.setter
    def crawler_script(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "crawler_script", value)

    @property
    @pulumi.getter(name="dataMasking")
    def data_masking(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Data Masking is enabled.
        """
        return pulumi.get(self, "data_masking")

    @data_masking.setter
    def data_masking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "data_masking", value)

    @property
    @pulumi.getter(name="falseAlarmMasking")
    def false_alarm_masking(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether False Alarm Masking is enabled.
        """
        return pulumi.get(self, "false_alarm_masking")

    @false_alarm_masking.setter
    def false_alarm_masking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "false_alarm_masking", value)

    @property
    @pulumi.getter(name="generalCheck")
    def general_check(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether General Check in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "general_check")

    @general_check.setter
    def general_check(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "general_check", value)

    @property
    @pulumi.getter(name="preciseProtection")
    def precise_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Precise Protection is enabled.
        """
        return pulumi.get(self, "precise_protection")

    @precise_protection.setter
    def precise_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "precise_protection", value)

    @property
    @pulumi.getter(name="webTamperProtection")
    def web_tamper_protection(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether Web Tamper Protection is enabled.
        """
        return pulumi.get(self, "web_tamper_protection")

    @web_tamper_protection.setter
    def web_tamper_protection(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "web_tamper_protection", value)

    @property
    @pulumi.getter
    def webshell(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether webshell detection in Basic Web Protection is enabled.
        """
        return pulumi.get(self, "webshell")

    @webshell.setter
    def webshell(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "webshell", value)



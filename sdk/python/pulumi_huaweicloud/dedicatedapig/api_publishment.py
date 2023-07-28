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
from ._inputs import *

__all__ = ['ApiPublishmentArgs', 'ApiPublishment']

@pulumi.input_type
class ApiPublishmentArgs:
    def __init__(__self__, *,
                 api_id: pulumi.Input[str],
                 env_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiPublishment resource.
        :param pulumi.Input[str] api_id: Specifies the ID of the API to be published or already published.  
               Changing this will create a new resource.
        :param pulumi.Input[str] env_id: Specifies the ID of the environmentto which the current version of the API
               will be published or has been published.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API belongs
               to. Changing this will create a new publishment resource.
        :param pulumi.Input[str] description: Specifies the description of the current publishment.
        :param pulumi.Input[str] region: Specifies the region in which to publish APIs.  
               If omitted, the provider-level region will be used.
               Changing this will create a new resource.
        :param pulumi.Input[str] version_id: Specifies the version ID of the current publishment.
        """
        pulumi.set(__self__, "api_id", api_id)
        pulumi.set(__self__, "env_id", env_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if version_id is not None:
            pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the API to be published or already published.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Input[str]:
        """
        Specifies the ID of the environmentto which the current version of the API
        will be published or has been published.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the API belongs
        to. Changing this will create a new publishment resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the current publishment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to publish APIs.  
        If omitted, the provider-level region will be used.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version ID of the current publishment.
        """
        return pulumi.get(self, "version_id")

    @version_id.setter
    def version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_id", value)


@pulumi.input_type
class _ApiPublishmentState:
    def __init__(__self__, *,
                 api_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 env_name: Optional[pulumi.Input[str]] = None,
                 histories: Optional[pulumi.Input[Sequence[pulumi.Input['ApiPublishmentHistoryArgs']]]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 publish_id: Optional[pulumi.Input[str]] = None,
                 published_at: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiPublishment resources.
        :param pulumi.Input[str] api_id: Specifies the ID of the API to be published or already published.  
               Changing this will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the current publishment.
        :param pulumi.Input[str] env_id: Specifies the ID of the environmentto which the current version of the API
               will be published or has been published.
               Changing this will create a new resource.
        :param pulumi.Input[str] env_name: The name of the environment to which the current version of the API is published.
        :param pulumi.Input[Sequence[pulumi.Input['ApiPublishmentHistoryArgs']]] histories: All publish informations of the API.  
               The object structure is documented below.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API belongs
               to. Changing this will create a new publishment resource.
        :param pulumi.Input[str] publish_id: The publish ID of the API in current environment.
        :param pulumi.Input[str] published_at: Time when the current version was published.
        :param pulumi.Input[str] region: Specifies the region in which to publish APIs.  
               If omitted, the provider-level region will be used.
               Changing this will create a new resource.
        :param pulumi.Input[str] version_id: Specifies the version ID of the current publishment.
        """
        if api_id is not None:
            pulumi.set(__self__, "api_id", api_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if env_id is not None:
            pulumi.set(__self__, "env_id", env_id)
        if env_name is not None:
            pulumi.set(__self__, "env_name", env_name)
        if histories is not None:
            pulumi.set(__self__, "histories", histories)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if publish_id is not None:
            pulumi.set(__self__, "publish_id", publish_id)
        if published_at is not None:
            pulumi.set(__self__, "published_at", published_at)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if version_id is not None:
            pulumi.set(__self__, "version_id", version_id)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the API to be published or already published.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "api_id")

    @api_id.setter
    def api_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the description of the current publishment.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the ID of the environmentto which the current version of the API
        will be published or has been published.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "env_id")

    @env_id.setter
    def env_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_id", value)

    @property
    @pulumi.getter(name="envName")
    def env_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the environment to which the current version of the API is published.
        """
        return pulumi.get(self, "env_name")

    @env_name.setter
    def env_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "env_name", value)

    @property
    @pulumi.getter
    def histories(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ApiPublishmentHistoryArgs']]]]:
        """
        All publish informations of the API.  
        The object structure is documented below.
        """
        return pulumi.get(self, "histories")

    @histories.setter
    def histories(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ApiPublishmentHistoryArgs']]]]):
        pulumi.set(self, "histories", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies an ID of the APIG dedicated instance to which the API belongs
        to. Changing this will create a new publishment resource.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="publishId")
    def publish_id(self) -> Optional[pulumi.Input[str]]:
        """
        The publish ID of the API in current environment.
        """
        return pulumi.get(self, "publish_id")

    @publish_id.setter
    def publish_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publish_id", value)

    @property
    @pulumi.getter(name="publishedAt")
    def published_at(self) -> Optional[pulumi.Input[str]]:
        """
        Time when the current version was published.
        """
        return pulumi.get(self, "published_at")

    @published_at.setter
    def published_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "published_at", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the region in which to publish APIs.  
        If omitted, the provider-level region will be used.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the version ID of the current publishment.
        """
        return pulumi.get(self, "version_id")

    @version_id.setter
    def version_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version_id", value)


class ApiPublishment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage
        ### Publish a new version of the API

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        env_id = config.require_object("envId")
        api_id = config.require_object("apiId")
        default = huaweicloud.dedicated_apig.ApiPublishment("default",
            instance_id=instance_id,
            env_id=env_id,
            api_id=api_id)
        ```
        ### Switch to a specified version of the API which is published

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        env_id = config.require_object("envId")
        api_id = config.require_object("apiId")
        version_id = config.require_object("versionId")
        default = huaweicloud.dedicated_apig.ApiPublishment("default",
            instance_id=instance_id,
            env_id=env_id,
            api_id=api_id,
            version_id=version_id)
        ```

        ## Import

        The publishments can be imported using their related `instance_id`, `env_id` and `api_id`, separated by slashes, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/apiPublishment:ApiPublishment test <instance_id>/<env_id>/<api_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Specifies the ID of the API to be published or already published.  
               Changing this will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the current publishment.
        :param pulumi.Input[str] env_id: Specifies the ID of the environmentto which the current version of the API
               will be published or has been published.
               Changing this will create a new resource.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API belongs
               to. Changing this will create a new publishment resource.
        :param pulumi.Input[str] region: Specifies the region in which to publish APIs.  
               If omitted, the provider-level region will be used.
               Changing this will create a new resource.
        :param pulumi.Input[str] version_id: Specifies the version ID of the current publishment.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiPublishmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage
        ### Publish a new version of the API

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        env_id = config.require_object("envId")
        api_id = config.require_object("apiId")
        default = huaweicloud.dedicated_apig.ApiPublishment("default",
            instance_id=instance_id,
            env_id=env_id,
            api_id=api_id)
        ```
        ### Switch to a specified version of the API which is published

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        config = pulumi.Config()
        instance_id = config.require_object("instanceId")
        env_id = config.require_object("envId")
        api_id = config.require_object("apiId")
        version_id = config.require_object("versionId")
        default = huaweicloud.dedicated_apig.ApiPublishment("default",
            instance_id=instance_id,
            env_id=env_id,
            api_id=api_id,
            version_id=version_id)
        ```

        ## Import

        The publishments can be imported using their related `instance_id`, `env_id` and `api_id`, separated by slashes, e.g.

        ```sh
         $ pulumi import huaweicloud:DedicatedApig/apiPublishment:ApiPublishment test <instance_id>/<env_id>/<api_id>
        ```

        :param str resource_name: The name of the resource.
        :param ApiPublishmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiPublishmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 env_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 version_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiPublishmentArgs.__new__(ApiPublishmentArgs)

            if api_id is None and not opts.urn:
                raise TypeError("Missing required property 'api_id'")
            __props__.__dict__["api_id"] = api_id
            __props__.__dict__["description"] = description
            if env_id is None and not opts.urn:
                raise TypeError("Missing required property 'env_id'")
            __props__.__dict__["env_id"] = env_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["region"] = region
            __props__.__dict__["version_id"] = version_id
            __props__.__dict__["env_name"] = None
            __props__.__dict__["histories"] = None
            __props__.__dict__["publish_id"] = None
            __props__.__dict__["published_at"] = None
        super(ApiPublishment, __self__).__init__(
            'huaweicloud:DedicatedApig/apiPublishment:ApiPublishment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            env_id: Optional[pulumi.Input[str]] = None,
            env_name: Optional[pulumi.Input[str]] = None,
            histories: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiPublishmentHistoryArgs']]]]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            publish_id: Optional[pulumi.Input[str]] = None,
            published_at: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            version_id: Optional[pulumi.Input[str]] = None) -> 'ApiPublishment':
        """
        Get an existing ApiPublishment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_id: Specifies the ID of the API to be published or already published.  
               Changing this will create a new resource.
        :param pulumi.Input[str] description: Specifies the description of the current publishment.
        :param pulumi.Input[str] env_id: Specifies the ID of the environmentto which the current version of the API
               will be published or has been published.
               Changing this will create a new resource.
        :param pulumi.Input[str] env_name: The name of the environment to which the current version of the API is published.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ApiPublishmentHistoryArgs']]]] histories: All publish informations of the API.  
               The object structure is documented below.
        :param pulumi.Input[str] instance_id: Specifies an ID of the APIG dedicated instance to which the API belongs
               to. Changing this will create a new publishment resource.
        :param pulumi.Input[str] publish_id: The publish ID of the API in current environment.
        :param pulumi.Input[str] published_at: Time when the current version was published.
        :param pulumi.Input[str] region: Specifies the region in which to publish APIs.  
               If omitted, the provider-level region will be used.
               Changing this will create a new resource.
        :param pulumi.Input[str] version_id: Specifies the version ID of the current publishment.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiPublishmentState.__new__(_ApiPublishmentState)

        __props__.__dict__["api_id"] = api_id
        __props__.__dict__["description"] = description
        __props__.__dict__["env_id"] = env_id
        __props__.__dict__["env_name"] = env_name
        __props__.__dict__["histories"] = histories
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["publish_id"] = publish_id
        __props__.__dict__["published_at"] = published_at
        __props__.__dict__["region"] = region
        __props__.__dict__["version_id"] = version_id
        return ApiPublishment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiId")
    def api_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the API to be published or already published.  
        Changing this will create a new resource.
        """
        return pulumi.get(self, "api_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the description of the current publishment.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="envId")
    def env_id(self) -> pulumi.Output[str]:
        """
        Specifies the ID of the environmentto which the current version of the API
        will be published or has been published.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "env_id")

    @property
    @pulumi.getter(name="envName")
    def env_name(self) -> pulumi.Output[str]:
        """
        The name of the environment to which the current version of the API is published.
        """
        return pulumi.get(self, "env_name")

    @property
    @pulumi.getter
    def histories(self) -> pulumi.Output[Sequence['outputs.ApiPublishmentHistory']]:
        """
        All publish informations of the API.  
        The object structure is documented below.
        """
        return pulumi.get(self, "histories")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Specifies an ID of the APIG dedicated instance to which the API belongs
        to. Changing this will create a new publishment resource.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="publishId")
    def publish_id(self) -> pulumi.Output[str]:
        """
        The publish ID of the API in current environment.
        """
        return pulumi.get(self, "publish_id")

    @property
    @pulumi.getter(name="publishedAt")
    def published_at(self) -> pulumi.Output[str]:
        """
        Time when the current version was published.
        """
        return pulumi.get(self, "published_at")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        Specifies the region in which to publish APIs.  
        If omitted, the provider-level region will be used.
        Changing this will create a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="versionId")
    def version_id(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the version ID of the current publishment.
        """
        return pulumi.get(self, "version_id")


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

__all__ = ['TranscodingTemplateGroupArgs', 'TranscodingTemplateGroup']

@pulumi.input_type
class TranscodingTemplateGroupArgs:
    def __init__(__self__, *,
                 output_format: pulumi.Input[int],
                 audio: Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']] = None,
                 dash_segment_duration: Optional[pulumi.Input[int]] = None,
                 hls_segment_duration: Optional[pulumi.Input[int]] = None,
                 low_bitrate_hd: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 video_common: Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']] = None,
                 videos: Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]] = None):
        """
        The set of arguments for constructing a TranscodingTemplateGroup resource.
        :param pulumi.Input[int] output_format: Specifies the packaging type. Possible values are:
               + **1**: HLS
               + **2**: DASH
               + **3**: HLS+DASH
               + **4**: MP4
               + **5**: MP3
               + **6**: ADTS
        :param pulumi.Input['TranscodingTemplateGroupAudioArgs'] audio: Specifies the audio parameters. The object structure is documented below.
        :param pulumi.Input[int] dash_segment_duration: Specifies the dash segment duration. This parameter is used only when `output_format`
               is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[int] hls_segment_duration: Specifies the HLS segment duration. This parameter is used only
               when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[bool] low_bitrate_hd: Specifies Whether to enable low bitrate HD. The default value is false.
        :param pulumi.Input[str] name: Specifies the name of a transcoding template group.
        :param pulumi.Input[str] region: The region in which to create the transcoding template group resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input['TranscodingTemplateGroupVideoCommonArgs'] video_common: Specifies the video parameters.
               The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]] videos: Specifies the list of output video configurations.
               The object structure is documented below.
        """
        pulumi.set(__self__, "output_format", output_format)
        if audio is not None:
            pulumi.set(__self__, "audio", audio)
        if dash_segment_duration is not None:
            pulumi.set(__self__, "dash_segment_duration", dash_segment_duration)
        if hls_segment_duration is not None:
            pulumi.set(__self__, "hls_segment_duration", hls_segment_duration)
        if low_bitrate_hd is not None:
            pulumi.set(__self__, "low_bitrate_hd", low_bitrate_hd)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if video_common is not None:
            pulumi.set(__self__, "video_common", video_common)
        if videos is not None:
            pulumi.set(__self__, "videos", videos)

    @property
    @pulumi.getter(name="outputFormat")
    def output_format(self) -> pulumi.Input[int]:
        """
        Specifies the packaging type. Possible values are:
        + **1**: HLS
        + **2**: DASH
        + **3**: HLS+DASH
        + **4**: MP4
        + **5**: MP3
        + **6**: ADTS
        """
        return pulumi.get(self, "output_format")

    @output_format.setter
    def output_format(self, value: pulumi.Input[int]):
        pulumi.set(self, "output_format", value)

    @property
    @pulumi.getter
    def audio(self) -> Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']]:
        """
        Specifies the audio parameters. The object structure is documented below.
        """
        return pulumi.get(self, "audio")

    @audio.setter
    def audio(self, value: Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']]):
        pulumi.set(self, "audio", value)

    @property
    @pulumi.getter(name="dashSegmentDuration")
    def dash_segment_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the dash segment duration. This parameter is used only when `output_format`
        is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "dash_segment_duration")

    @dash_segment_duration.setter
    def dash_segment_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dash_segment_duration", value)

    @property
    @pulumi.getter(name="hlsSegmentDuration")
    def hls_segment_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the HLS segment duration. This parameter is used only
        when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "hls_segment_duration")

    @hls_segment_duration.setter
    def hls_segment_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hls_segment_duration", value)

    @property
    @pulumi.getter(name="lowBitrateHd")
    def low_bitrate_hd(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies Whether to enable low bitrate HD. The default value is false.
        """
        return pulumi.get(self, "low_bitrate_hd")

    @low_bitrate_hd.setter
    def low_bitrate_hd(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "low_bitrate_hd", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of a transcoding template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the transcoding template group resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="videoCommon")
    def video_common(self) -> Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']]:
        """
        Specifies the video parameters.
        The object structure is documented below.
        """
        return pulumi.get(self, "video_common")

    @video_common.setter
    def video_common(self, value: Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']]):
        pulumi.set(self, "video_common", value)

    @property
    @pulumi.getter
    def videos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]]:
        """
        Specifies the list of output video configurations.
        The object structure is documented below.
        """
        return pulumi.get(self, "videos")

    @videos.setter
    def videos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]]):
        pulumi.set(self, "videos", value)


@pulumi.input_type
class _TranscodingTemplateGroupState:
    def __init__(__self__, *,
                 audio: Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']] = None,
                 dash_segment_duration: Optional[pulumi.Input[int]] = None,
                 hls_segment_duration: Optional[pulumi.Input[int]] = None,
                 low_bitrate_hd: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 video_common: Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']] = None,
                 videos: Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]] = None):
        """
        Input properties used for looking up and filtering TranscodingTemplateGroup resources.
        :param pulumi.Input['TranscodingTemplateGroupAudioArgs'] audio: Specifies the audio parameters. The object structure is documented below.
        :param pulumi.Input[int] dash_segment_duration: Specifies the dash segment duration. This parameter is used only when `output_format`
               is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[int] hls_segment_duration: Specifies the HLS segment duration. This parameter is used only
               when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[bool] low_bitrate_hd: Specifies Whether to enable low bitrate HD. The default value is false.
        :param pulumi.Input[str] name: Specifies the name of a transcoding template group.
        :param pulumi.Input[int] output_format: Specifies the packaging type. Possible values are:
               + **1**: HLS
               + **2**: DASH
               + **3**: HLS+DASH
               + **4**: MP4
               + **5**: MP3
               + **6**: ADTS
        :param pulumi.Input[str] region: The region in which to create the transcoding template group resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Indicates the IDs of templates in the template group
        :param pulumi.Input['TranscodingTemplateGroupVideoCommonArgs'] video_common: Specifies the video parameters.
               The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]] videos: Specifies the list of output video configurations.
               The object structure is documented below.
        """
        if audio is not None:
            pulumi.set(__self__, "audio", audio)
        if dash_segment_duration is not None:
            pulumi.set(__self__, "dash_segment_duration", dash_segment_duration)
        if hls_segment_duration is not None:
            pulumi.set(__self__, "hls_segment_duration", hls_segment_duration)
        if low_bitrate_hd is not None:
            pulumi.set(__self__, "low_bitrate_hd", low_bitrate_hd)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if output_format is not None:
            pulumi.set(__self__, "output_format", output_format)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if template_ids is not None:
            pulumi.set(__self__, "template_ids", template_ids)
        if video_common is not None:
            pulumi.set(__self__, "video_common", video_common)
        if videos is not None:
            pulumi.set(__self__, "videos", videos)

    @property
    @pulumi.getter
    def audio(self) -> Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']]:
        """
        Specifies the audio parameters. The object structure is documented below.
        """
        return pulumi.get(self, "audio")

    @audio.setter
    def audio(self, value: Optional[pulumi.Input['TranscodingTemplateGroupAudioArgs']]):
        pulumi.set(self, "audio", value)

    @property
    @pulumi.getter(name="dashSegmentDuration")
    def dash_segment_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the dash segment duration. This parameter is used only when `output_format`
        is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "dash_segment_duration")

    @dash_segment_duration.setter
    def dash_segment_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dash_segment_duration", value)

    @property
    @pulumi.getter(name="hlsSegmentDuration")
    def hls_segment_duration(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the HLS segment duration. This parameter is used only
        when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "hls_segment_duration")

    @hls_segment_duration.setter
    def hls_segment_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hls_segment_duration", value)

    @property
    @pulumi.getter(name="lowBitrateHd")
    def low_bitrate_hd(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies Whether to enable low bitrate HD. The default value is false.
        """
        return pulumi.get(self, "low_bitrate_hd")

    @low_bitrate_hd.setter
    def low_bitrate_hd(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "low_bitrate_hd", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of a transcoding template group.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="outputFormat")
    def output_format(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the packaging type. Possible values are:
        + **1**: HLS
        + **2**: DASH
        + **3**: HLS+DASH
        + **4**: MP4
        + **5**: MP3
        + **6**: ADTS
        """
        return pulumi.get(self, "output_format")

    @output_format.setter
    def output_format(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "output_format", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which to create the transcoding template group resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Indicates the IDs of templates in the template group
        """
        return pulumi.get(self, "template_ids")

    @template_ids.setter
    def template_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "template_ids", value)

    @property
    @pulumi.getter(name="videoCommon")
    def video_common(self) -> Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']]:
        """
        Specifies the video parameters.
        The object structure is documented below.
        """
        return pulumi.get(self, "video_common")

    @video_common.setter
    def video_common(self, value: Optional[pulumi.Input['TranscodingTemplateGroupVideoCommonArgs']]):
        pulumi.set(self, "video_common", value)

    @property
    @pulumi.getter
    def videos(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]]:
        """
        Specifies the list of output video configurations.
        The object structure is documented below.
        """
        return pulumi.get(self, "videos")

    @videos.setter
    def videos(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TranscodingTemplateGroupVideoArgs']]]]):
        pulumi.set(self, "videos", value)


class TranscodingTemplateGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audio: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupAudioArgs']]] = None,
                 dash_segment_duration: Optional[pulumi.Input[int]] = None,
                 hls_segment_duration: Optional[pulumi.Input[int]] = None,
                 low_bitrate_hd: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 video_common: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoCommonArgs']]] = None,
                 videos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoArgs']]]]] = None,
                 __props__=None):
        """
        Manages an MPC transcoding template group resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test = huaweicloud.mpc.TranscodingTemplateGroup("test",
            audio=huaweicloud.mpc.TranscodingTemplateGroupAudioArgs(
                bitrate=0,
                channels=2,
                codec=2,
                output_policy="transcode",
                sample_rate=1,
            ),
            dash_segment_duration=5,
            hls_segment_duration=5,
            low_bitrate_hd=True,
            output_format=1,
            video_common=huaweicloud.mpc.TranscodingTemplateGroupVideoCommonArgs(
                black_bar_removal=0,
                codec=2,
                fps=0,
                level=15,
                max_consecutive_bframes=7,
                max_iframes_interval=5,
                output_policy="transcode",
                profile=4,
                quality=1,
            ),
            videos=[huaweicloud.mpc.TranscodingTemplateGroupVideoArgs(
                bitrate=0,
                height=2160,
                width=3840,
            )])
        ```

        ## Import

        MPC transcoding template groups can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup test 589e49809bb84447a759f6fa9aa19949
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TranscodingTemplateGroupAudioArgs']] audio: Specifies the audio parameters. The object structure is documented below.
        :param pulumi.Input[int] dash_segment_duration: Specifies the dash segment duration. This parameter is used only when `output_format`
               is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[int] hls_segment_duration: Specifies the HLS segment duration. This parameter is used only
               when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[bool] low_bitrate_hd: Specifies Whether to enable low bitrate HD. The default value is false.
        :param pulumi.Input[str] name: Specifies the name of a transcoding template group.
        :param pulumi.Input[int] output_format: Specifies the packaging type. Possible values are:
               + **1**: HLS
               + **2**: DASH
               + **3**: HLS+DASH
               + **4**: MP4
               + **5**: MP3
               + **6**: ADTS
        :param pulumi.Input[str] region: The region in which to create the transcoding template group resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoCommonArgs']] video_common: Specifies the video parameters.
               The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoArgs']]]] videos: Specifies the list of output video configurations.
               The object structure is documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TranscodingTemplateGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an MPC transcoding template group resource within HuaweiCloud.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_huaweicloud as huaweicloud

        test = huaweicloud.mpc.TranscodingTemplateGroup("test",
            audio=huaweicloud.mpc.TranscodingTemplateGroupAudioArgs(
                bitrate=0,
                channels=2,
                codec=2,
                output_policy="transcode",
                sample_rate=1,
            ),
            dash_segment_duration=5,
            hls_segment_duration=5,
            low_bitrate_hd=True,
            output_format=1,
            video_common=huaweicloud.mpc.TranscodingTemplateGroupVideoCommonArgs(
                black_bar_removal=0,
                codec=2,
                fps=0,
                level=15,
                max_consecutive_bframes=7,
                max_iframes_interval=5,
                output_policy="transcode",
                profile=4,
                quality=1,
            ),
            videos=[huaweicloud.mpc.TranscodingTemplateGroupVideoArgs(
                bitrate=0,
                height=2160,
                width=3840,
            )])
        ```

        ## Import

        MPC transcoding template groups can be imported using the `id`, e.g.

        ```sh
         $ pulumi import huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup test 589e49809bb84447a759f6fa9aa19949
        ```

        :param str resource_name: The name of the resource.
        :param TranscodingTemplateGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TranscodingTemplateGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audio: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupAudioArgs']]] = None,
                 dash_segment_duration: Optional[pulumi.Input[int]] = None,
                 hls_segment_duration: Optional[pulumi.Input[int]] = None,
                 low_bitrate_hd: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 output_format: Optional[pulumi.Input[int]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 video_common: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoCommonArgs']]] = None,
                 videos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TranscodingTemplateGroupArgs.__new__(TranscodingTemplateGroupArgs)

            __props__.__dict__["audio"] = audio
            __props__.__dict__["dash_segment_duration"] = dash_segment_duration
            __props__.__dict__["hls_segment_duration"] = hls_segment_duration
            __props__.__dict__["low_bitrate_hd"] = low_bitrate_hd
            __props__.__dict__["name"] = name
            if output_format is None and not opts.urn:
                raise TypeError("Missing required property 'output_format'")
            __props__.__dict__["output_format"] = output_format
            __props__.__dict__["region"] = region
            __props__.__dict__["video_common"] = video_common
            __props__.__dict__["videos"] = videos
            __props__.__dict__["template_ids"] = None
        super(TranscodingTemplateGroup, __self__).__init__(
            'huaweicloud:Mpc/transcodingTemplateGroup:TranscodingTemplateGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audio: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupAudioArgs']]] = None,
            dash_segment_duration: Optional[pulumi.Input[int]] = None,
            hls_segment_duration: Optional[pulumi.Input[int]] = None,
            low_bitrate_hd: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            output_format: Optional[pulumi.Input[int]] = None,
            region: Optional[pulumi.Input[str]] = None,
            template_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            video_common: Optional[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoCommonArgs']]] = None,
            videos: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoArgs']]]]] = None) -> 'TranscodingTemplateGroup':
        """
        Get an existing TranscodingTemplateGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['TranscodingTemplateGroupAudioArgs']] audio: Specifies the audio parameters. The object structure is documented below.
        :param pulumi.Input[int] dash_segment_duration: Specifies the dash segment duration. This parameter is used only when `output_format`
               is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[int] hls_segment_duration: Specifies the HLS segment duration. This parameter is used only
               when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        :param pulumi.Input[bool] low_bitrate_hd: Specifies Whether to enable low bitrate HD. The default value is false.
        :param pulumi.Input[str] name: Specifies the name of a transcoding template group.
        :param pulumi.Input[int] output_format: Specifies the packaging type. Possible values are:
               + **1**: HLS
               + **2**: DASH
               + **3**: HLS+DASH
               + **4**: MP4
               + **5**: MP3
               + **6**: ADTS
        :param pulumi.Input[str] region: The region in which to create the transcoding template group resource. If omitted,
               the provider-level region will be used. Changing this creates a new resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] template_ids: Indicates the IDs of templates in the template group
        :param pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoCommonArgs']] video_common: Specifies the video parameters.
               The object structure is documented below.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TranscodingTemplateGroupVideoArgs']]]] videos: Specifies the list of output video configurations.
               The object structure is documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TranscodingTemplateGroupState.__new__(_TranscodingTemplateGroupState)

        __props__.__dict__["audio"] = audio
        __props__.__dict__["dash_segment_duration"] = dash_segment_duration
        __props__.__dict__["hls_segment_duration"] = hls_segment_duration
        __props__.__dict__["low_bitrate_hd"] = low_bitrate_hd
        __props__.__dict__["name"] = name
        __props__.__dict__["output_format"] = output_format
        __props__.__dict__["region"] = region
        __props__.__dict__["template_ids"] = template_ids
        __props__.__dict__["video_common"] = video_common
        __props__.__dict__["videos"] = videos
        return TranscodingTemplateGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def audio(self) -> pulumi.Output[Optional['outputs.TranscodingTemplateGroupAudio']]:
        """
        Specifies the audio parameters. The object structure is documented below.
        """
        return pulumi.get(self, "audio")

    @property
    @pulumi.getter(name="dashSegmentDuration")
    def dash_segment_duration(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the dash segment duration. This parameter is used only when `output_format`
        is set to 2 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "dash_segment_duration")

    @property
    @pulumi.getter(name="hlsSegmentDuration")
    def hls_segment_duration(self) -> pulumi.Output[Optional[int]]:
        """
        Specifies the HLS segment duration. This parameter is used only
        when `output_format` is set to 1 or 3. The value ranges from 2 to 10. The default value is 5. The unit is second.
        """
        return pulumi.get(self, "hls_segment_duration")

    @property
    @pulumi.getter(name="lowBitrateHd")
    def low_bitrate_hd(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies Whether to enable low bitrate HD. The default value is false.
        """
        return pulumi.get(self, "low_bitrate_hd")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of a transcoding template group.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="outputFormat")
    def output_format(self) -> pulumi.Output[int]:
        """
        Specifies the packaging type. Possible values are:
        + **1**: HLS
        + **2**: DASH
        + **3**: HLS+DASH
        + **4**: MP4
        + **5**: MP3
        + **6**: ADTS
        """
        return pulumi.get(self, "output_format")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The region in which to create the transcoding template group resource. If omitted,
        the provider-level region will be used. Changing this creates a new resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="templateIds")
    def template_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        Indicates the IDs of templates in the template group
        """
        return pulumi.get(self, "template_ids")

    @property
    @pulumi.getter(name="videoCommon")
    def video_common(self) -> pulumi.Output[Optional['outputs.TranscodingTemplateGroupVideoCommon']]:
        """
        Specifies the video parameters.
        The object structure is documented below.
        """
        return pulumi.get(self, "video_common")

    @property
    @pulumi.getter
    def videos(self) -> pulumi.Output[Optional[Sequence['outputs.TranscodingTemplateGroupVideo']]]:
        """
        Specifies the list of output video configurations.
        The object structure is documented below.
        """
        return pulumi.get(self, "videos")


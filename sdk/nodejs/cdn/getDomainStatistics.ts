// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get the statistics of CDN domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const config = new pulumi.Config();
 * const domainName = config.require("domainName");
 *
 * const test = pulumi.output(huaweicloud.Cdn.getDomainStatistics({
 *     action: "location_detail",
 *     domainName: "terraform.test.huaweicloud.com",
 *     endTime: 1662021000000,
 *     startTime: 1662019200000,
 *     statType: "req_num",
 * }));
 * ```
 */
export function getDomainStatistics(args: GetDomainStatisticsArgs, opts?: pulumi.InvokeOptions): Promise<GetDomainStatisticsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Cdn/getDomainStatistics:getDomainStatistics", {
        "action": args.action,
        "country": args.country,
        "domainName": args.domainName,
        "endTime": args.endTime,
        "enterpriseProjectId": args.enterpriseProjectId,
        "groupBy": args.groupBy,
        "interval": args.interval,
        "isp": args.isp,
        "province": args.province,
        "startTime": args.startTime,
        "statType": args.statType,
    }, opts);
}

/**
 * A collection of arguments for invoking getDomainStatistics.
 */
export interface GetDomainStatisticsArgs {
    /**
     * Specifies the action name. Possible values are: **location_summary** and **location_detail**.
     */
    action: string;
    /**
     * Specifies the country or region code. Use commas (,) to separate multiple codes.
     * The value all indicates all country/region codes.
     * See the [country and region](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0089.html) for values.
     */
    country?: string;
    /**
     * Specifies the domain name list.
     * Domain names are separated by commas (,), for example, `www.test1.com,www.test2.com`.
     * The value all indicates that all domain names under your account are queried.
     */
    domainName: string;
    /**
     * Specifies the end timestamp of the query.
     * The timestamp must be set to a multiple of 5 minutes.
     * + If the value of interval is 300, set this parameter to a multiple of 5 minutes,
     * for example, 1631243700000, which means 2021-09-10 11:15:00.
     * + If the value of interval is 3600, set this parameter to a multiple of 1 hour,
     * for example, 1631325600000, which means 2021-09-11 10:00:00.
     * + If the value of interval is 86400, set this parameter to 00:00:00 (GMT+08:00),
     * for example, 1631376000000, which means 2021-09-12 00:00:00.
     */
    endTime: number;
    /**
     * Specifies the enterprise project that the resource belongs to.
     * This parameter is valid only when the enterprise project function is enabled.
     * The value all indicates all projects. This parameter is mandatory when you use an IAM user.
     */
    enterpriseProjectId?: string;
    /**
     * Specifies the data grouping mode. Use commas (,) to separate multiple groups.
     * Available data groups are **domain**, **country**, **province**, and **isp**. By default, data is not grouped.
     */
    groupBy?: string;
    /**
     * Specifies the query time interval, in seconds, the value can be,
     * + **300**(5 minutes): Maximum query span 2 days
     * + **3600**(1 hour): Maximum query span 7 days
     * + **86400**(1 day): Maximum query span 31 days
     */
    interval?: number;
    /**
     * Specifies the carrier code. Use commas (,) to separate multiple codes.
     * The value all indicates all carrier codes.
     * See the [carriers](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0075.html) for values.
     */
    isp?: string;
    /**
     * Specifies the province code. This parameter is valid when country is set to **cn**.
     * Use commas (,) to separate multiple codes. The value all indicates all provinces.
     * See the [areas](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0074.html) for values.
     */
    province?: string;
    /**
     * Specifies the start timestamp of the query.
     * The timestamp must be set to a multiple of 5 minutes.
     * + If the value of interval is **300**, set this parameter to a multiple of 5 minutes,
     * for example, 1631240100000, which means 2021-09-10 10:15:00.
     * + If the value of interval is **3600**, set this parameter to a multiple of 1 hour,
     * for example, 1631239200000, which means 2021-09-10 10:00:00.
     * + If the value of interval is **86400**, set this parameter to 00:00:00 (GMT+08:00),
     * for example, 1631203200000, which means 2021-09-10 00:00:00.
     */
    startTime: number;
    /**
     * The statistic type.
     */
    statType: string;
}

/**
 * A collection of values returned by getDomainStatistics.
 */
export interface GetDomainStatisticsResult {
    readonly action: string;
    readonly country?: string;
    readonly domainName: string;
    readonly endTime: number;
    readonly enterpriseProjectId?: string;
    readonly groupBy?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interval?: number;
    readonly isp?: string;
    readonly province?: string;
    /**
     * The data organized according to the specified grouping mode.
     */
    readonly result: string;
    readonly startTime: number;
    readonly statType: string;
}

export function getDomainStatisticsOutput(args: GetDomainStatisticsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDomainStatisticsResult> {
    return pulumi.output(args).apply(a => getDomainStatistics(a, opts))
}

/**
 * A collection of arguments for invoking getDomainStatistics.
 */
export interface GetDomainStatisticsOutputArgs {
    /**
     * Specifies the action name. Possible values are: **location_summary** and **location_detail**.
     */
    action: pulumi.Input<string>;
    /**
     * Specifies the country or region code. Use commas (,) to separate multiple codes.
     * The value all indicates all country/region codes.
     * See the [country and region](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0089.html) for values.
     */
    country?: pulumi.Input<string>;
    /**
     * Specifies the domain name list.
     * Domain names are separated by commas (,), for example, `www.test1.com,www.test2.com`.
     * The value all indicates that all domain names under your account are queried.
     */
    domainName: pulumi.Input<string>;
    /**
     * Specifies the end timestamp of the query.
     * The timestamp must be set to a multiple of 5 minutes.
     * + If the value of interval is 300, set this parameter to a multiple of 5 minutes,
     * for example, 1631243700000, which means 2021-09-10 11:15:00.
     * + If the value of interval is 3600, set this parameter to a multiple of 1 hour,
     * for example, 1631325600000, which means 2021-09-11 10:00:00.
     * + If the value of interval is 86400, set this parameter to 00:00:00 (GMT+08:00),
     * for example, 1631376000000, which means 2021-09-12 00:00:00.
     */
    endTime: pulumi.Input<number>;
    /**
     * Specifies the enterprise project that the resource belongs to.
     * This parameter is valid only when the enterprise project function is enabled.
     * The value all indicates all projects. This parameter is mandatory when you use an IAM user.
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Specifies the data grouping mode. Use commas (,) to separate multiple groups.
     * Available data groups are **domain**, **country**, **province**, and **isp**. By default, data is not grouped.
     */
    groupBy?: pulumi.Input<string>;
    /**
     * Specifies the query time interval, in seconds, the value can be,
     * + **300**(5 minutes): Maximum query span 2 days
     * + **3600**(1 hour): Maximum query span 7 days
     * + **86400**(1 day): Maximum query span 31 days
     */
    interval?: pulumi.Input<number>;
    /**
     * Specifies the carrier code. Use commas (,) to separate multiple codes.
     * The value all indicates all carrier codes.
     * See the [carriers](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0075.html) for values.
     */
    isp?: pulumi.Input<string>;
    /**
     * Specifies the province code. This parameter is valid when country is set to **cn**.
     * Use commas (,) to separate multiple codes. The value all indicates all provinces.
     * See the [areas](https://support.huaweicloud.com/intl/en-us/api-cdn/cdn_02_0074.html) for values.
     */
    province?: pulumi.Input<string>;
    /**
     * Specifies the start timestamp of the query.
     * The timestamp must be set to a multiple of 5 minutes.
     * + If the value of interval is **300**, set this parameter to a multiple of 5 minutes,
     * for example, 1631240100000, which means 2021-09-10 10:15:00.
     * + If the value of interval is **3600**, set this parameter to a multiple of 1 hour,
     * for example, 1631239200000, which means 2021-09-10 10:00:00.
     * + If the value of interval is **86400**, set this parameter to 00:00:00 (GMT+08:00),
     * for example, 1631203200000, which means 2021-09-10 00:00:00.
     */
    startTime: pulumi.Input<number>;
    /**
     * The statistic type.
     */
    statType: pulumi.Input<string>;
}

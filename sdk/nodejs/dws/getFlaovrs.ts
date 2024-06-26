// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get available flavors of DWS cluster node.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const flavor = pulumi.output(huaweicloud.Dws.getFlaovrs({
 *     vcpus: 8,
 * }));
 * ```
 */
export function getFlaovrs(args?: GetFlaovrsArgs, opts?: pulumi.InvokeOptions): Promise<GetFlaovrsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Dws/getFlaovrs:getFlaovrs", {
        "availabilityZone": args.availabilityZone,
        "datastoreType": args.datastoreType,
        "memory": args.memory,
        "region": args.region,
        "vcpus": args.vcpus,
    }, opts);
}

/**
 * A collection of arguments for invoking getFlaovrs.
 */
export interface GetFlaovrsArgs {
    /**
     * The availability zone name.
     */
    availabilityZone?: string;
    /**
     * The type of datastore.  
     * The options are as follows:
     * - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
     * - **hybrid**: a single data warehouse used for transaction and analytics workloads,
     * in single-node or cluster mode.
     * - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
     */
    datastoreType?: string;
    /**
     * The ram of the dws node flavor in GB.
     */
    memory?: number;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: string;
    /**
     * The vcpus of the dws node flavor.
     */
    vcpus?: number;
}

/**
 * A collection of values returned by getFlaovrs.
 */
export interface GetFlaovrsResult {
    readonly availabilityZone?: string;
    /**
     * The type of datastore.  
     * The options are as follows:
     * - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
     * - **hybrid**: a single data warehouse used for transaction and analytics workloads,
     * in single-node or cluster mode.
     * - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
     */
    readonly datastoreType?: string;
    /**
     * The list of flavor detail.
     * The Flavors structure is documented below.
     */
    readonly flavors: outputs.Dws.GetFlaovrsFlavor[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The ram of the dws node flavor in GB.
     */
    readonly memory?: number;
    readonly region: string;
    /**
     * The vcpus of the dws node flavor.
     */
    readonly vcpus?: number;
}

export function getFlaovrsOutput(args?: GetFlaovrsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlaovrsResult> {
    return pulumi.output(args).apply(a => getFlaovrs(a, opts))
}

/**
 * A collection of arguments for invoking getFlaovrs.
 */
export interface GetFlaovrsOutputArgs {
    /**
     * The availability zone name.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * The type of datastore.  
     * The options are as follows:
     * - **dws**: OLAP, elastic scaling, unlimited scaling of compute and storage capacity.
     * - **hybrid**: a single data warehouse used for transaction and analytics workloads,
     * in single-node or cluster mode.
     * - **stream**: built-in time series operators; up to 40:1 compression ratio; applicable to IoT services.
     */
    datastoreType?: pulumi.Input<string>;
    /**
     * The ram of the dws node flavor in GB.
     */
    memory?: pulumi.Input<number>;
    /**
     * Specifies the region in which to query the data source.
     * If omitted, the provider-level region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * The vcpus of the dws node flavor.
     */
    vcpus?: pulumi.Input<number>;
}

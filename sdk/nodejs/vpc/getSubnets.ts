// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get a list of VPC subnet.
 *
 * ## Example Usage
 *
 * An example filter by name and tag
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as huaweicloud from "@pulumi/huaweicloud";
 *
 * const subnet = huaweicloud.Vpc.getSubnets({
 *     name: _var.subnet_name,
 *     tags: {
 *         foo: "bar",
 *     },
 * });
 * export const subnetVpcIds = [subnet.then(subnet => subnet.subnets)].map(__item => __item?.vpcId);
 * ```
 * ## **Attributes Reference**
 *
 * The following attributes are exported:
 *
 * * `id` - Indicates a data source ID.
 * * `subnets` - Indicates a list of all subnets found. Structure is documented below.
 *
 * The `subnets` block supports:
 *
 * * `id` - Indicates the ID of the subnet.
 * * `name` - Indicates the name of the subnet.
 * * `description` - Indicates the description of the subnet.
 * * `cidr` - Indicates the cidr block of the subnet.
 * * `status` - Indicates the current status of the subnet.
 * * `vpcId` - Indicates the Id of the VPC that the subnet belongs to.
 * * `gatewayIp` - Indicates the subnet gateway address of the subnet.
 * * `primaryDns` - Indicates the IP address of DNS server 1 on the subnet.
 * * `secondaryDns` - Indicates the IP address of DNS server 2 on the subnet.
 * * `availabilityZone` - Indicates the availability zone (AZ) to which the subnet belongs to.
 * * `dhcpEnable` - Indicates whether the DHCP is enabled.
 * * `dnsList` - Indicates The IP address list of DNS servers on the subnet.
 * * `ipv4SubnetId` - Indicates the ID of the IPv4 subnet (Native OpenStack API).
 * * `ipv6Enable` - Indicates whether the IPv6 is enabled.
 * * `ipv6SubnetId` - Indicates the ID of the IPv6 subnet (Native OpenStack API).
 * * `ipv6Cidr` - Indicates the IPv6 subnet CIDR block.
 * * `ipv6Gateway` - Indicates the IPv6 subnet gateway.
 * * `tags` - Indicates the key/value pairs which associated with the subnet.
 */
export function getSubnets(args?: GetSubnetsArgs, opts?: pulumi.InvokeOptions): Promise<GetSubnetsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("huaweicloud:Vpc/getSubnets:getSubnets", {
        "availabilityZone": args.availabilityZone,
        "cidr": args.cidr,
        "gatewayIp": args.gatewayIp,
        "id": args.id,
        "name": args.name,
        "primaryDns": args.primaryDns,
        "region": args.region,
        "secondaryDns": args.secondaryDns,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubnets.
 */
export interface GetSubnetsArgs {
    /**
     * Specifies the availability zone (AZ) to which the desired subnet belongs to.
     */
    availabilityZone?: string;
    /**
     * Specifies the network segment of desired subnet. The value must be in CIDR format.
     */
    cidr?: string;
    /**
     * Specifies the subnet gateway address of desired subnet.
     */
    gatewayIp?: string;
    /**
     * - Specifies the id of the desired subnet.
     */
    id?: string;
    /**
     * Specifies the name of the desired subnet.
     */
    name?: string;
    /**
     * Specifies the IP address of DNS server 1 on the desired subnet.
     */
    primaryDns?: string;
    /**
     * Specifies the region in which to obtain the subnet. If omitted, the provider-level
     * region will be used.
     */
    region?: string;
    /**
     * Specifies the IP address of DNS server 2 on the desired subnet.
     */
    secondaryDns?: string;
    /**
     * Specifies the current status of the desired subnet.
     * the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
     */
    status?: string;
    /**
     * Specifies the included key/value pairs which associated with the desired subnet.
     */
    tags?: {[key: string]: string};
    /**
     * Specifies the id of the VPC that the desired subnet belongs to.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by getSubnets.
 */
export interface GetSubnetsResult {
    readonly availabilityZone?: string;
    readonly cidr?: string;
    readonly gatewayIp?: string;
    readonly id: string;
    readonly name?: string;
    readonly primaryDns?: string;
    readonly region: string;
    readonly secondaryDns?: string;
    readonly status?: string;
    readonly subnets: outputs.Vpc.GetSubnetsSubnet[];
    readonly tags?: {[key: string]: string};
    readonly vpcId?: string;
}

export function getSubnetsOutput(args?: GetSubnetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubnetsResult> {
    return pulumi.output(args).apply(a => getSubnets(a, opts))
}

/**
 * A collection of arguments for invoking getSubnets.
 */
export interface GetSubnetsOutputArgs {
    /**
     * Specifies the availability zone (AZ) to which the desired subnet belongs to.
     */
    availabilityZone?: pulumi.Input<string>;
    /**
     * Specifies the network segment of desired subnet. The value must be in CIDR format.
     */
    cidr?: pulumi.Input<string>;
    /**
     * Specifies the subnet gateway address of desired subnet.
     */
    gatewayIp?: pulumi.Input<string>;
    /**
     * - Specifies the id of the desired subnet.
     */
    id?: pulumi.Input<string>;
    /**
     * Specifies the name of the desired subnet.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the IP address of DNS server 1 on the desired subnet.
     */
    primaryDns?: pulumi.Input<string>;
    /**
     * Specifies the region in which to obtain the subnet. If omitted, the provider-level
     * region will be used.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the IP address of DNS server 2 on the desired subnet.
     */
    secondaryDns?: pulumi.Input<string>;
    /**
     * Specifies the current status of the desired subnet.
     * the value can be ACTIVE, DOWN, UNKNOWN, or ERROR.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the included key/value pairs which associated with the desired subnet.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the id of the VPC that the desired subnet belongs to.
     */
    vpcId?: pulumi.Input<string>;
}

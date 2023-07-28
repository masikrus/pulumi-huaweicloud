// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an ELB L7 Policy resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const listenerId = config.requireObject("listenerId");
 * const poolId = config.requireObject("poolId");
 * const policy1 = new huaweicloud.dedicatedelb.L7policy("policy1", {
 *     action: "REDIRECT_TO_POOL",
 *     description: "test description",
 *     listenerId: listenerId,
 *     redirectPoolId: poolId,
 * });
 * ```
 *
 * ## Import
 *
 * ELB policy can be imported using the `id`, e.g. bash
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedElb/l7policy:L7policy policy_1 <id>
 * ```
 */
export class L7policy extends pulumi.CustomResource {
    /**
     * Get an existing L7policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: L7policyState, opts?: pulumi.CustomResourceOptions): L7policy {
        return new L7policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedElb/l7policy:L7policy';

    /**
     * Returns true if the given object is an instance of L7policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is L7policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === L7policy.__pulumiType;
    }

    /**
     * Specifies whether requests are forwarded to another backend server group
     * or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
     * + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirectPoolId`.
     * + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listenerId` to the
     * HTTPS listener specified by `redirectListenerId`.
     * Defaults to **REDIRECT_TO_POOL**.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * Human-readable description for the L7 Policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Listener on which the L7 Policy will be associated with. Changing
     * this creates a new L7 Policy.
     */
    public readonly listenerId!: pulumi.Output<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have to be unique.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the ID of the listener to which the traffic is redirected.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
     * following requirements:
     * + Can only be an HTTPS listener.
     * + Can only be a listener of the same load balancer.
     */
    public readonly redirectListenerId!: pulumi.Output<string>;
    /**
     * Specifies the ID of the backend server group to which traffic is forwarded.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
     * following requirements:
     * + Cannot be the default backend server group of the listener.
     * + Cannot be the backend server group used by forwarding policies of other listeners.
     */
    public readonly redirectPoolId!: pulumi.Output<string>;
    /**
     * The region in which to create the L7 Policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Policy.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a L7policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: L7policyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: L7policyArgs | L7policyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as L7policyState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["listenerId"] = state ? state.listenerId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["redirectListenerId"] = state ? state.redirectListenerId : undefined;
            resourceInputs["redirectPoolId"] = state ? state.redirectPoolId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as L7policyArgs | undefined;
            if ((!args || args.listenerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'listenerId'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["listenerId"] = args ? args.listenerId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["redirectListenerId"] = args ? args.redirectListenerId : undefined;
            resourceInputs["redirectPoolId"] = args ? args.redirectPoolId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(L7policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering L7policy resources.
 */
export interface L7policyState {
    /**
     * Specifies whether requests are forwarded to another backend server group
     * or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
     * + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirectPoolId`.
     * + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listenerId` to the
     * HTTPS listener specified by `redirectListenerId`.
     * Defaults to **REDIRECT_TO_POOL**.
     */
    action?: pulumi.Input<string>;
    /**
     * Human-readable description for the L7 Policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The Listener on which the L7 Policy will be associated with. Changing
     * this creates a new L7 Policy.
     */
    listenerId?: pulumi.Input<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the ID of the listener to which the traffic is redirected.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
     * following requirements:
     * + Can only be an HTTPS listener.
     * + Can only be a listener of the same load balancer.
     */
    redirectListenerId?: pulumi.Input<string>;
    /**
     * Specifies the ID of the backend server group to which traffic is forwarded.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
     * following requirements:
     * + Cannot be the default backend server group of the listener.
     * + Cannot be the backend server group used by forwarding policies of other listeners.
     */
    redirectPoolId?: pulumi.Input<string>;
    /**
     * The region in which to create the L7 Policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Policy.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a L7policy resource.
 */
export interface L7policyArgs {
    /**
     * Specifies whether requests are forwarded to another backend server group
     * or redirected to an HTTPS listener. Changing this creates a new L7 Policy. The value ranges:
     * + **REDIRECT_TO_POOL**: Requests are forwarded to the backend server group specified by `redirectPoolId`.
     * + **REDIRECT_TO_LISTENER**: Requests are redirected from the HTTP listener specified by `listenerId` to the
     * HTTPS listener specified by `redirectListenerId`.
     * Defaults to **REDIRECT_TO_POOL**.
     */
    action?: pulumi.Input<string>;
    /**
     * Human-readable description for the L7 Policy.
     */
    description?: pulumi.Input<string>;
    /**
     * The Listener on which the L7 Policy will be associated with. Changing
     * this creates a new L7 Policy.
     */
    listenerId: pulumi.Input<string>;
    /**
     * Human-readable name for the L7 Policy. Does not have to be unique.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the ID of the listener to which the traffic is redirected.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_LISTENER**. The listener must meet the
     * following requirements:
     * + Can only be an HTTPS listener.
     * + Can only be a listener of the same load balancer.
     */
    redirectListenerId?: pulumi.Input<string>;
    /**
     * Specifies the ID of the backend server group to which traffic is forwarded.
     * This parameter is mandatory when `action` is set to **REDIRECT_TO_POOL**. The backend server group must meet the
     * following requirements:
     * + Cannot be the default backend server group of the listener.
     * + Cannot be the backend server group used by forwarding policies of other listeners.
     */
    redirectPoolId?: pulumi.Input<string>;
    /**
     * The region in which to create the L7 Policy resource. If omitted, the
     * provider-level region will be used. Changing this creates a new L7 Policy.
     */
    region?: pulumi.Input<string>;
}

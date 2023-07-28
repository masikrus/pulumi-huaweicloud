// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an APIG (API) custom response resource within HuaweiCloud.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const instanceId = config.requireObject("instanceId");
 * const groupId = config.requireObject("groupId");
 * const responseName = config.requireObject("responseName");
 * const test = new huaweicloud.dedicatedapig.Response("test", {
 *     instanceId: instanceId,
 *     groupId: groupId,
 *     rules: [{
 *         errorType: "AUTHORIZER_FAILURE",
 *         body: `{"code":"$context.authorizer.frontend.code","message":"$context.authorizer.frontend.message"}`,
 *         statusCode: 401,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * API Responses can be imported using their `name` and IDs of the APIG dedicated instances and API groups to which the API response belongs, separated by slashes, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/response:Response test <instance_id>/<group_id>/<name>
 * ```
 */
export class Response extends pulumi.CustomResource {
    /**
     * Get an existing Response resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResponseState, opts?: pulumi.CustomResourceOptions): Response {
        return new Response(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/response:Response';

    /**
     * Returns true if the given object is an instance of Response.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Response {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Response.__pulumiType;
    }

    /**
     * The creation time of the API custom response.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Specifies the ID of the API group to which the API custom response
     * belongs.
     * Changing this will create a new resource.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * Specifies the ID of the dedicated instance to which the API group and the
     * API custom response belong.
     * Changing this will create a new resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the name of the API custom response.  
     * The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region where the API custom response is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Specifies the API custom response rules definition.  
     * The object structure is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.DedicatedApig.ResponseRule[]>;
    /**
     * The latest update time of the API custom response.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Response resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResponseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResponseArgs | ResponseState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ResponseState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
        } else {
            const args = argsOrState as ResponseArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Response.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Response resources.
 */
export interface ResponseState {
    /**
     * The creation time of the API custom response.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Specifies the ID of the API group to which the API custom response
     * belongs.
     * Changing this will create a new resource.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated instance to which the API group and the
     * API custom response belong.
     * Changing this will create a new resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the name of the API custom response.  
     * The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the API custom response is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the API custom response rules definition.  
     * The object structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ResponseRule>[]>;
    /**
     * The latest update time of the API custom response.
     */
    updatedAt?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Response resource.
 */
export interface ResponseArgs {
    /**
     * Specifies the ID of the API group to which the API custom response
     * belongs.
     * Changing this will create a new resource.
     */
    groupId: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated instance to which the API group and the
     * API custom response belong.
     * Changing this will create a new resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the name of the API custom response.  
     * The valid length is limited from `1` to `64`, letters, digits, hyphens (-) and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the API custom response is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the API custom response rules definition.  
     * The object structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ResponseRule>[]>;
}

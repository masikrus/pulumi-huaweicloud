// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Using this resource to assign an administrator role to a user within HuaweiCloud.
 *
 * ## Example Usage
 * ### Assign an administrator role to a user
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const appId = config.requireObject("appId");
 * const appKey = config.requireObject("appKey");
 * const userAccount = config.requireObject("userAccount");
 * const test = new huaweicloud.meeting.AdminAssignment("test", {
 *     appId: appId,
 *     appKey: appKey,
 *     account: userAccount,
 * });
 * ```
 *
 * ## Import
 *
 * The assignment relationships can be imported using their `id` and authorization parameters, separated by slashes, e.g. Import an administrator assignment and authenticated by account.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test &ltid&gt/&ltaccount_name&gt/&ltaccount_password&gt
 * ```
 *
 *  Import an administrator assignment and authenticated by `APP ID`/`APP Key`.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/adminAssignment:AdminAssignment test &ltid&gt/&ltapp_id&gt/&ltapp_key&gt/&ltcorp_id&gt/&ltuser_id&gt
 * ```
 *
 *  For this resource, the `corp_id` and `user_id` are never used, you can omit them but the slashes cannot be missing.
 */
export class AdminAssignment extends pulumi.CustomResource {
    /**
     * Get an existing AdminAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AdminAssignmentState, opts?: pulumi.CustomResourceOptions): AdminAssignment {
        return new AdminAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Meeting/adminAssignment:AdminAssignment';

    /**
     * Returns true if the given object is an instance of AdminAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AdminAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AdminAssignment.__pulumiType;
    }

    /**
     * Specifies the user account to be assigned the administrator role.
     * The value can contain **1** to **64** characters.
     * Changing this parameter will create a new resource.
     */
    public readonly account!: pulumi.Output<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * default administrator belongs. Changing this parameter will create a new resource.
     */
    public readonly accountName!: pulumi.Output<string | undefined>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    public readonly accountPassword!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    public readonly appId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    public readonly appKey!: pulumi.Output<string | undefined>;

    /**
     * Create a AdminAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AdminAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AdminAssignmentArgs | AdminAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AdminAssignmentState | undefined;
            resourceInputs["account"] = state ? state.account : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["appKey"] = state ? state.appKey : undefined;
        } else {
            const args = argsOrState as AdminAssignmentArgs | undefined;
            if ((!args || args.account === undefined) && !opts.urn) {
                throw new Error("Missing required property 'account'");
            }
            resourceInputs["account"] = args ? args.account : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args ? args.accountPassword : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["appKey"] = args ? args.appKey : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AdminAssignment.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AdminAssignment resources.
 */
export interface AdminAssignmentState {
    /**
     * Specifies the user account to be assigned the administrator role.
     * The value can contain **1** to **64** characters.
     * Changing this parameter will create a new resource.
     */
    account?: pulumi.Input<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * default administrator belongs. Changing this parameter will create a new resource.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    appId?: pulumi.Input<string>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    appKey?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AdminAssignment resource.
 */
export interface AdminAssignmentArgs {
    /**
     * Specifies the user account to be assigned the administrator role.
     * The value can contain **1** to **64** characters.
     * Changing this parameter will create a new resource.
     */
    account: pulumi.Input<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * default administrator belongs. Changing this parameter will create a new resource.
     */
    accountName?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * Required if `accountName` is set. Changing this parameter will create a new resource.
     */
    accountPassword?: pulumi.Input<string>;
    /**
     * Specifies the ID of the Third-party application.
     * Changing this parameter will create a new resource.
     */
    appId?: pulumi.Input<string>;
    /**
     * Specifies the Key information of the Third-party APP.
     * Required if `appId` is set. Changing this parameter will create a new resource.
     */
    appKey?: pulumi.Input<string>;
}

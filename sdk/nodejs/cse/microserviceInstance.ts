// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a dedicated microservice instance resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a microservice instance under a microservice with RBAC authentication of engine disabled
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const engineConnAddr = config.requireObject("engineConnAddr");
 * const microserviceId = config.requireObject("microserviceId");
 * const regionName = config.requireObject("regionName");
 * const azName = config.requireObject("azName");
 * const test = new huaweicloud.cse.MicroserviceInstance("test", {
 *     connectAddress: engineConnAddr,
 *     microserviceId: microserviceId,
 *     hostName: "localhost",
 *     endpoints: [
 *         "grpc://127.0.1.132:9980",
 *         "rest://127.0.0.111:8081",
 *     ],
 *     version: "1.0.0",
 *     properties: {
 *         _TAGS: "A, B",
 *         attr1: "a",
 *         nodeIP: "127.0.0.1",
 *     },
 *     healthCheck: {
 *         mode: "push",
 *         interval: 30,
 *         maxRetries: 3,
 *     },
 *     dataCenter: {
 *         name: "dc",
 *         region: regionName,
 *         availabilityZone: azName,
 *     },
 * });
 * ```
 * ### Create a microservice instance under a microservice with RBAC authentication of engine enabled
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const engineConnAddr = config.requireObject("engineConnAddr");
 * const microserviceId = config.requireObject("microserviceId");
 * const regionName = config.requireObject("regionName");
 * const azName = config.requireObject("azName");
 * const test = new huaweicloud.cse.MicroserviceInstance("test", {
 *     connectAddress: engineConnAddr,
 *     microserviceId: microserviceId,
 *     hostName: "localhost",
 *     endpoints: [
 *         "grpc://127.0.1.132:9980",
 *         "rest://127.0.0.111:8081",
 *     ],
 *     version: "1.0.0",
 *     properties: {
 *         _TAGS: "A, B",
 *         attr1: "a",
 *         nodeIP: "127.0.0.1",
 *     },
 *     healthCheck: {
 *         mode: "push",
 *         interval: 30,
 *         maxRetries: 3,
 *     },
 *     dataCenter: {
 *         name: "dc",
 *         region: regionName,
 *         availabilityZone: azName,
 *     },
 *     adminUser: "root",
 *     adminPass: "Huawei!123",
 * });
 * ```
 *
 * ## Import
 *
 * Microservices can be imported using related `connect_address`, `microservice_id` and their `id`, separated by a slash (/), e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cse/microserviceInstance:MicroserviceInstance test https://124.70.26.32:30100/f14960ba495e03f59f85aacaaafbdef3fbff3f0d/336e7428dd9411eca913fa163e7364b7
 * ```
 *
 *  If you enabled the **RBAC** authorization, you also need to provide the account name and password, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Cse/microserviceInstance:MicroserviceInstance test 'https://124.70.26.32:30100/f14960ba495e03f59f85aacaaafbdef3fbff3f0d/336e7428dd9411eca913fa163e7364b7/root/Test!123'
 * ```
 *
 *  The single quotes can help you solve the problem of special characters reporting errors on bash.
 */
export class MicroserviceInstance extends pulumi.CustomResource {
    /**
     * Get an existing MicroserviceInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MicroserviceInstanceState, opts?: pulumi.CustomResourceOptions): MicroserviceInstance {
        return new MicroserviceInstance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Cse/microserviceInstance:MicroserviceInstance';

    /**
     * Returns true if the given object is an instance of MicroserviceInstance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MicroserviceInstance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MicroserviceInstance.__pulumiType;
    }

    /**
     * Specifies the account password.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     * The password format must meet the following conditions:
     * + Must be `8` to `32` characters long.
     * + A password must contain at least one digit, one uppercase letter, one lowercase letter, and one special character
     * (-~!@#%^*_=+?$&()|<>{}[]).
     * + Cannot be the account name or account name spelled backwards.
     * + The password can only start with a letter.
     */
    public readonly adminPass!: pulumi.Output<string | undefined>;
    /**
     * Specifies the account name. The initial account name is **root**.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     */
    public readonly adminUser!: pulumi.Output<string | undefined>;
    /**
     * Specifies the connection address of service registry center for the
     * specified dedicated CSE engine. Changing this will create a new microservice instance.
     */
    public readonly connectAddress!: pulumi.Output<string>;
    /**
     * Specifies the data center configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    public readonly dataCenter!: pulumi.Output<outputs.Cse.MicroserviceInstanceDataCenter>;
    /**
     * Specifies the access addresses information.
     * Changing this will create a new microservice instance.
     */
    public readonly endpoints!: pulumi.Output<string[]>;
    /**
     * Specifies the health check configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    public readonly healthCheck!: pulumi.Output<outputs.Cse.MicroserviceInstanceHealthCheck>;
    /**
     * Specifies the host name, such as `localhost`.
     * Changing this will create a new microservice instance.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * Specifies the ID of the dedicated microservice to which the instance
     * belongs. Changing this will create a new microservice instance.
     */
    public readonly microserviceId!: pulumi.Output<string>;
    /**
     * Specifies the extended attributes.
     * Changing this will create a new microservice instance.
     */
    public readonly properties!: pulumi.Output<{[key: string]: string}>;
    /**
     * The microservice instance status. The values supports **UP**, **DOWN**, **STARTING** and **OUTOFSERVICE**.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Specifies the version of the dedicated microservice instance.
     * Changing this will create a new microservice instance.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a MicroserviceInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MicroserviceInstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MicroserviceInstanceArgs | MicroserviceInstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MicroserviceInstanceState | undefined;
            resourceInputs["adminPass"] = state ? state.adminPass : undefined;
            resourceInputs["adminUser"] = state ? state.adminUser : undefined;
            resourceInputs["connectAddress"] = state ? state.connectAddress : undefined;
            resourceInputs["dataCenter"] = state ? state.dataCenter : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["healthCheck"] = state ? state.healthCheck : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["microserviceId"] = state ? state.microserviceId : undefined;
            resourceInputs["properties"] = state ? state.properties : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as MicroserviceInstanceArgs | undefined;
            if ((!args || args.connectAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectAddress'");
            }
            if ((!args || args.endpoints === undefined) && !opts.urn) {
                throw new Error("Missing required property 'endpoints'");
            }
            if ((!args || args.hostName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'hostName'");
            }
            if ((!args || args.microserviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'microserviceId'");
            }
            resourceInputs["adminPass"] = args ? args.adminPass : undefined;
            resourceInputs["adminUser"] = args ? args.adminUser : undefined;
            resourceInputs["connectAddress"] = args ? args.connectAddress : undefined;
            resourceInputs["dataCenter"] = args ? args.dataCenter : undefined;
            resourceInputs["endpoints"] = args ? args.endpoints : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["microserviceId"] = args ? args.microserviceId : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MicroserviceInstance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MicroserviceInstance resources.
 */
export interface MicroserviceInstanceState {
    /**
     * Specifies the account password.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     * The password format must meet the following conditions:
     * + Must be `8` to `32` characters long.
     * + A password must contain at least one digit, one uppercase letter, one lowercase letter, and one special character
     * (-~!@#%^*_=+?$&()|<>{}[]).
     * + Cannot be the account name or account name spelled backwards.
     * + The password can only start with a letter.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies the account name. The initial account name is **root**.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     */
    adminUser?: pulumi.Input<string>;
    /**
     * Specifies the connection address of service registry center for the
     * specified dedicated CSE engine. Changing this will create a new microservice instance.
     */
    connectAddress?: pulumi.Input<string>;
    /**
     * Specifies the data center configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    dataCenter?: pulumi.Input<inputs.Cse.MicroserviceInstanceDataCenter>;
    /**
     * Specifies the access addresses information.
     * Changing this will create a new microservice instance.
     */
    endpoints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the health check configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    healthCheck?: pulumi.Input<inputs.Cse.MicroserviceInstanceHealthCheck>;
    /**
     * Specifies the host name, such as `localhost`.
     * Changing this will create a new microservice instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated microservice to which the instance
     * belongs. Changing this will create a new microservice instance.
     */
    microserviceId?: pulumi.Input<string>;
    /**
     * Specifies the extended attributes.
     * Changing this will create a new microservice instance.
     */
    properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The microservice instance status. The values supports **UP**, **DOWN**, **STARTING** and **OUTOFSERVICE**.
     */
    status?: pulumi.Input<string>;
    /**
     * Specifies the version of the dedicated microservice instance.
     * Changing this will create a new microservice instance.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MicroserviceInstance resource.
 */
export interface MicroserviceInstanceArgs {
    /**
     * Specifies the account password.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     * The password format must meet the following conditions:
     * + Must be `8` to `32` characters long.
     * + A password must contain at least one digit, one uppercase letter, one lowercase letter, and one special character
     * (-~!@#%^*_=+?$&()|<>{}[]).
     * + Cannot be the account name or account name spelled backwards.
     * + The password can only start with a letter.
     */
    adminPass?: pulumi.Input<string>;
    /**
     * Specifies the account name. The initial account name is **root**.
     * Required if the `authType` of engine is **RBAC**. Changing this will create a new microservice instance.
     */
    adminUser?: pulumi.Input<string>;
    /**
     * Specifies the connection address of service registry center for the
     * specified dedicated CSE engine. Changing this will create a new microservice instance.
     */
    connectAddress: pulumi.Input<string>;
    /**
     * Specifies the data center configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    dataCenter?: pulumi.Input<inputs.Cse.MicroserviceInstanceDataCenter>;
    /**
     * Specifies the access addresses information.
     * Changing this will create a new microservice instance.
     */
    endpoints: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the health check configuration.
     * The object structure is documented below.
     * Changing this will create a new microservice instance.
     */
    healthCheck?: pulumi.Input<inputs.Cse.MicroserviceInstanceHealthCheck>;
    /**
     * Specifies the host name, such as `localhost`.
     * Changing this will create a new microservice instance.
     */
    hostName: pulumi.Input<string>;
    /**
     * Specifies the ID of the dedicated microservice to which the instance
     * belongs. Changing this will create a new microservice instance.
     */
    microserviceId: pulumi.Input<string>;
    /**
     * Specifies the extended attributes.
     * Changing this will create a new microservice instance.
     */
    properties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the version of the dedicated microservice instance.
     * Changing this will create a new microservice instance.
     */
    version?: pulumi.Input<string>;
}

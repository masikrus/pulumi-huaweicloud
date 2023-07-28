// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an APIG API resource within HuaweiCloud.
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
 * const apiName = config.requireObject("apiName");
 * const customResponseId = config.requireObject("customResponseId");
 * const customAuthId = config.requireObject("customAuthId");
 * const vpcChannelId = config.requireObject("vpcChannelId");
 * const test = new huaweicloud.dedicatedapig.Api("test", {
 *     instanceId: instanceId,
 *     groupId: groupId,
 *     type: "Public",
 *     requestProtocol: "HTTP",
 *     requestMethod: "POST",
 *     requestPath: "/terraform/users",
 *     securityAuthentication: "AUTHORIZER",
 *     matching: "Exact",
 *     successResponse: "Successful",
 *     responseId: customResponseId,
 *     authorizerId: customAuthId,
 *     backendParams: [{
 *         type: "SYSTEM",
 *         name: "X-User-Auth",
 *         location: "HEADER",
 *         value: "user_name",
 *     }],
 *     web: {
 *         path: "/backend/users",
 *         vpcChannelId: vpcChannelId,
 *         requestMethod: "POST",
 *         requestProtocol: "HTTP",
 *         timeout: 5000,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * APIs can be imported using their `name` and the related dedicated instance IDs, separated by a slash, e.g.
 *
 * ```sh
 *  $ pulumi import huaweicloud:DedicatedApig/api:Api test <instance_id>/<name>
 * ```
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:DedicatedApig/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * Specifies the ID of the backend custom authorization.
     */
    public readonly authorizerId!: pulumi.Output<string | undefined>;
    /**
     * Specifies an array of one or more backend parameters. The maximum of request
     * parameters is 50. The object structure is documented above.
     */
    public readonly backendParams!: pulumi.Output<outputs.DedicatedApig.ApiBackendParam[] | undefined>;
    /**
     * Specifies the description of the API request body, which can be an example
     * request body, media type or parameters.
     * The request body does not exceed `20,480` characters.
     */
    public readonly bodyDescription!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether CORS is supported, defaults to **false**.
     */
    public readonly cors!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the description of the constant or system parameter.  
     * The description contains a maximum of `255` characters and the angle brackets (< and >) are not allowed.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the example response for a failure request.  
     * The response contains a maximum of `20,480` characters.
     */
    public readonly failureResponse!: pulumi.Output<string | undefined>;
    /**
     * Specifies the function graph backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    public readonly funcGraph!: pulumi.Output<outputs.DedicatedApig.ApiFuncGraph>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    public readonly funcGraphPolicies!: pulumi.Output<outputs.DedicatedApig.ApiFuncGraphPolicy[] | undefined>;
    /**
     * Specifies an ID of the APIG group to which the API belongs to.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the API belongs
     * to. Changing this will create a new API resource.
     */
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * Specifies the route matching mode.  
     * The valid values are **Exact** and **Prefix**, defaults to **Exact**.
     */
    public readonly matching!: pulumi.Output<string | undefined>;
    /**
     * Specifies the mock backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    public readonly mock!: pulumi.Output<outputs.DedicatedApig.ApiMock>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    public readonly mockPolicies!: pulumi.Output<outputs.DedicatedApig.ApiMockPolicy[] | undefined>;
    /**
     * Specifies the backend policy name.  
     * The valid length is limited from can contain `3` to `64`, only letters, digits and underscores (_) are allowed.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the region where the API is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new API resource.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The registered time of the API.
     */
    public /*out*/ readonly registeredAt!: pulumi.Output<string>;
    /**
     * Specifies the backend request method of the API.  
     * The valid types are **GET**, **POST**, **PUT**, **DELETE**, **HEAD**, **PATCH**, **OPTIONS** and **ANY**.
     */
    public readonly requestMethod!: pulumi.Output<string>;
    /**
     * Specifies the configurations of the front-end parameters.  
     * The object structure is documented below.
     */
    public readonly requestParams!: pulumi.Output<outputs.DedicatedApig.ApiRequestParam[] | undefined>;
    /**
     * Specifies the request address, which can contain a maximum of `512` characters,
     * the request parameters enclosed with brackets ({}).
     * + The address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-), and
     * underscores (_) and must comply with URI specifications.
     * + The address can contain environment variables, each starting with a letter and consisting of `3` to `32` characters.
     */
    public readonly requestPath!: pulumi.Output<string>;
    /**
     * Specifies the backend request protocol. The valid values are **HTTP** and
     * **HTTPS**, defaults to **HTTPS**.
     */
    public readonly requestProtocol!: pulumi.Output<string>;
    /**
     * Specifies the APIG group response ID.
     */
    public readonly responseId!: pulumi.Output<string | undefined>;
    /**
     * Specifies the security authentication mode of the API request.  
     * The valid values are **NONE**, **APP** and **IAM**, defaults to **NONE**.
     */
    public readonly securityAuthentication!: pulumi.Output<string | undefined>;
    /**
     * Specifies whether the authentication of the application code is enabled.  
     * The application code must located in the header when `simpleAuthentication` is true.
     */
    public readonly simpleAuthentication!: pulumi.Output<boolean>;
    /**
     * Specifies the example response for a successful request.  
     * The response contains a maximum of `20,480` characters.
     */
    public readonly successResponse!: pulumi.Output<string | undefined>;
    /**
     * Specifies the condition type of the backend policy.  
     * The valid values are **Equal**, **Enumerated** and **Matching**, defaults to **Equal**.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The latest update time of the API.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Specifies the web backend details.  
     * The object structure is documented below. Changing this will create a new API resource.
     */
    public readonly web!: pulumi.Output<outputs.DedicatedApig.ApiWeb>;
    /**
     * Specifies the example response for a failed request.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    public readonly webPolicies!: pulumi.Output<outputs.DedicatedApig.ApiWebPolicy[] | undefined>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiState | undefined;
            resourceInputs["authorizerId"] = state ? state.authorizerId : undefined;
            resourceInputs["backendParams"] = state ? state.backendParams : undefined;
            resourceInputs["bodyDescription"] = state ? state.bodyDescription : undefined;
            resourceInputs["cors"] = state ? state.cors : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["failureResponse"] = state ? state.failureResponse : undefined;
            resourceInputs["funcGraph"] = state ? state.funcGraph : undefined;
            resourceInputs["funcGraphPolicies"] = state ? state.funcGraphPolicies : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["matching"] = state ? state.matching : undefined;
            resourceInputs["mock"] = state ? state.mock : undefined;
            resourceInputs["mockPolicies"] = state ? state.mockPolicies : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["registeredAt"] = state ? state.registeredAt : undefined;
            resourceInputs["requestMethod"] = state ? state.requestMethod : undefined;
            resourceInputs["requestParams"] = state ? state.requestParams : undefined;
            resourceInputs["requestPath"] = state ? state.requestPath : undefined;
            resourceInputs["requestProtocol"] = state ? state.requestProtocol : undefined;
            resourceInputs["responseId"] = state ? state.responseId : undefined;
            resourceInputs["securityAuthentication"] = state ? state.securityAuthentication : undefined;
            resourceInputs["simpleAuthentication"] = state ? state.simpleAuthentication : undefined;
            resourceInputs["successResponse"] = state ? state.successResponse : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["web"] = state ? state.web : undefined;
            resourceInputs["webPolicies"] = state ? state.webPolicies : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.requestMethod === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestMethod'");
            }
            if ((!args || args.requestPath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestPath'");
            }
            if ((!args || args.requestProtocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requestProtocol'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["authorizerId"] = args ? args.authorizerId : undefined;
            resourceInputs["backendParams"] = args ? args.backendParams : undefined;
            resourceInputs["bodyDescription"] = args ? args.bodyDescription : undefined;
            resourceInputs["cors"] = args ? args.cors : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["failureResponse"] = args ? args.failureResponse : undefined;
            resourceInputs["funcGraph"] = args ? args.funcGraph : undefined;
            resourceInputs["funcGraphPolicies"] = args ? args.funcGraphPolicies : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["matching"] = args ? args.matching : undefined;
            resourceInputs["mock"] = args ? args.mock : undefined;
            resourceInputs["mockPolicies"] = args ? args.mockPolicies : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestMethod"] = args ? args.requestMethod : undefined;
            resourceInputs["requestParams"] = args ? args.requestParams : undefined;
            resourceInputs["requestPath"] = args ? args.requestPath : undefined;
            resourceInputs["requestProtocol"] = args ? args.requestProtocol : undefined;
            resourceInputs["responseId"] = args ? args.responseId : undefined;
            resourceInputs["securityAuthentication"] = args ? args.securityAuthentication : undefined;
            resourceInputs["simpleAuthentication"] = args ? args.simpleAuthentication : undefined;
            resourceInputs["successResponse"] = args ? args.successResponse : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["web"] = args ? args.web : undefined;
            resourceInputs["webPolicies"] = args ? args.webPolicies : undefined;
            resourceInputs["registeredAt"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Api.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * Specifies the ID of the backend custom authorization.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more backend parameters. The maximum of request
     * parameters is 50. The object structure is documented above.
     */
    backendParams?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiBackendParam>[]>;
    /**
     * Specifies the description of the API request body, which can be an example
     * request body, media type or parameters.
     * The request body does not exceed `20,480` characters.
     */
    bodyDescription?: pulumi.Input<string>;
    /**
     * Specifies whether CORS is supported, defaults to **false**.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the constant or system parameter.  
     * The description contains a maximum of `255` characters and the angle brackets (< and >) are not allowed.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the example response for a failure request.  
     * The response contains a maximum of `20,480` characters.
     */
    failureResponse?: pulumi.Input<string>;
    /**
     * Specifies the function graph backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    funcGraph?: pulumi.Input<inputs.DedicatedApig.ApiFuncGraph>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    funcGraphPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiFuncGraphPolicy>[]>;
    /**
     * Specifies an ID of the APIG group to which the API belongs to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the API belongs
     * to. Changing this will create a new API resource.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * Specifies the route matching mode.  
     * The valid values are **Exact** and **Prefix**, defaults to **Exact**.
     */
    matching?: pulumi.Input<string>;
    /**
     * Specifies the mock backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    mock?: pulumi.Input<inputs.DedicatedApig.ApiMock>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    mockPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiMockPolicy>[]>;
    /**
     * Specifies the backend policy name.  
     * The valid length is limited from can contain `3` to `64`, only letters, digits and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the API is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new API resource.
     */
    region?: pulumi.Input<string>;
    /**
     * The registered time of the API.
     */
    registeredAt?: pulumi.Input<string>;
    /**
     * Specifies the backend request method of the API.  
     * The valid types are **GET**, **POST**, **PUT**, **DELETE**, **HEAD**, **PATCH**, **OPTIONS** and **ANY**.
     */
    requestMethod?: pulumi.Input<string>;
    /**
     * Specifies the configurations of the front-end parameters.  
     * The object structure is documented below.
     */
    requestParams?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiRequestParam>[]>;
    /**
     * Specifies the request address, which can contain a maximum of `512` characters,
     * the request parameters enclosed with brackets ({}).
     * + The address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-), and
     * underscores (_) and must comply with URI specifications.
     * + The address can contain environment variables, each starting with a letter and consisting of `3` to `32` characters.
     */
    requestPath?: pulumi.Input<string>;
    /**
     * Specifies the backend request protocol. The valid values are **HTTP** and
     * **HTTPS**, defaults to **HTTPS**.
     */
    requestProtocol?: pulumi.Input<string>;
    /**
     * Specifies the APIG group response ID.
     */
    responseId?: pulumi.Input<string>;
    /**
     * Specifies the security authentication mode of the API request.  
     * The valid values are **NONE**, **APP** and **IAM**, defaults to **NONE**.
     */
    securityAuthentication?: pulumi.Input<string>;
    /**
     * Specifies whether the authentication of the application code is enabled.  
     * The application code must located in the header when `simpleAuthentication` is true.
     */
    simpleAuthentication?: pulumi.Input<boolean>;
    /**
     * Specifies the example response for a successful request.  
     * The response contains a maximum of `20,480` characters.
     */
    successResponse?: pulumi.Input<string>;
    /**
     * Specifies the condition type of the backend policy.  
     * The valid values are **Equal**, **Enumerated** and **Matching**, defaults to **Equal**.
     */
    type?: pulumi.Input<string>;
    /**
     * The latest update time of the API.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Specifies the web backend details.  
     * The object structure is documented below. Changing this will create a new API resource.
     */
    web?: pulumi.Input<inputs.DedicatedApig.ApiWeb>;
    /**
     * Specifies the example response for a failed request.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    webPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiWebPolicy>[]>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * Specifies the ID of the backend custom authorization.
     */
    authorizerId?: pulumi.Input<string>;
    /**
     * Specifies an array of one or more backend parameters. The maximum of request
     * parameters is 50. The object structure is documented above.
     */
    backendParams?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiBackendParam>[]>;
    /**
     * Specifies the description of the API request body, which can be an example
     * request body, media type or parameters.
     * The request body does not exceed `20,480` characters.
     */
    bodyDescription?: pulumi.Input<string>;
    /**
     * Specifies whether CORS is supported, defaults to **false**.
     */
    cors?: pulumi.Input<boolean>;
    /**
     * Specifies the description of the constant or system parameter.  
     * The description contains a maximum of `255` characters and the angle brackets (< and >) are not allowed.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the example response for a failure request.  
     * The response contains a maximum of `20,480` characters.
     */
    failureResponse?: pulumi.Input<string>;
    /**
     * Specifies the function graph backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    funcGraph?: pulumi.Input<inputs.DedicatedApig.ApiFuncGraph>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    funcGraphPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiFuncGraphPolicy>[]>;
    /**
     * Specifies an ID of the APIG group to which the API belongs to.
     */
    groupId: pulumi.Input<string>;
    /**
     * Specifies an ID of the APIG dedicated instance to which the API belongs
     * to. Changing this will create a new API resource.
     */
    instanceId: pulumi.Input<string>;
    /**
     * Specifies the route matching mode.  
     * The valid values are **Exact** and **Prefix**, defaults to **Exact**.
     */
    matching?: pulumi.Input<string>;
    /**
     * Specifies the mock backend details.  
     * The object structure is documented below.
     * Changing this will create a new API resource.
     */
    mock?: pulumi.Input<inputs.DedicatedApig.ApiMock>;
    /**
     * Specifies the Mock policy backends.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    mockPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiMockPolicy>[]>;
    /**
     * Specifies the backend policy name.  
     * The valid length is limited from can contain `3` to `64`, only letters, digits and underscores (_) are allowed.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the region where the API is located.  
     * If omitted, the provider-level region will be used. Changing this will create a new API resource.
     */
    region?: pulumi.Input<string>;
    /**
     * Specifies the backend request method of the API.  
     * The valid types are **GET**, **POST**, **PUT**, **DELETE**, **HEAD**, **PATCH**, **OPTIONS** and **ANY**.
     */
    requestMethod: pulumi.Input<string>;
    /**
     * Specifies the configurations of the front-end parameters.  
     * The object structure is documented below.
     */
    requestParams?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiRequestParam>[]>;
    /**
     * Specifies the request address, which can contain a maximum of `512` characters,
     * the request parameters enclosed with brackets ({}).
     * + The address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-), and
     * underscores (_) and must comply with URI specifications.
     * + The address can contain environment variables, each starting with a letter and consisting of `3` to `32` characters.
     */
    requestPath: pulumi.Input<string>;
    /**
     * Specifies the backend request protocol. The valid values are **HTTP** and
     * **HTTPS**, defaults to **HTTPS**.
     */
    requestProtocol: pulumi.Input<string>;
    /**
     * Specifies the APIG group response ID.
     */
    responseId?: pulumi.Input<string>;
    /**
     * Specifies the security authentication mode of the API request.  
     * The valid values are **NONE**, **APP** and **IAM**, defaults to **NONE**.
     */
    securityAuthentication?: pulumi.Input<string>;
    /**
     * Specifies whether the authentication of the application code is enabled.  
     * The application code must located in the header when `simpleAuthentication` is true.
     */
    simpleAuthentication?: pulumi.Input<boolean>;
    /**
     * Specifies the example response for a successful request.  
     * The response contains a maximum of `20,480` characters.
     */
    successResponse?: pulumi.Input<string>;
    /**
     * Specifies the condition type of the backend policy.  
     * The valid values are **Equal**, **Enumerated** and **Matching**, defaults to **Equal**.
     */
    type: pulumi.Input<string>;
    /**
     * Specifies the web backend details.  
     * The object structure is documented below. Changing this will create a new API resource.
     */
    web?: pulumi.Input<inputs.DedicatedApig.ApiWeb>;
    /**
     * Specifies the example response for a failed request.  
     * The maximum blocks of the policy is 5.
     * The object structure is documented below.
     */
    webPolicies?: pulumi.Input<pulumi.Input<inputs.DedicatedApig.ApiWebPolicy>[]>;
}

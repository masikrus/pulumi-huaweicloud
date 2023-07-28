// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The provider type for the huaweicloud package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'huaweicloud';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    /**
     * The access key of the HuaweiCloud to use.
     */
    public readonly accessKey!: pulumi.Output<string | undefined>;
    /**
     * The name of domain who created the agency (Identity v3).
     */
    public readonly agencyDomainName!: pulumi.Output<string | undefined>;
    /**
     * The name of agency
     */
    public readonly agencyName!: pulumi.Output<string | undefined>;
    /**
     * The Identity authentication URL.
     */
    public readonly authUrl!: pulumi.Output<string | undefined>;
    /**
     * A Custom CA certificate.
     */
    public readonly cacertFile!: pulumi.Output<string | undefined>;
    /**
     * A client certificate to authenticate with.
     */
    public readonly cert!: pulumi.Output<string | undefined>;
    /**
     * The endpoint of cloud provider, defaults to myhuaweicloud.com
     */
    public readonly cloud!: pulumi.Output<string | undefined>;
    /**
     * The name of delegated project (Identity v3).
     */
    public readonly delegatedProject!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Domain to scope to.
     */
    public readonly domainId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Domain to scope to.
     */
    public readonly domainName!: pulumi.Output<string | undefined>;
    /**
     * enterprise project id
     */
    public readonly enterpriseProjectId!: pulumi.Output<string | undefined>;
    /**
     * A client private key to authenticate with.
     */
    public readonly key!: pulumi.Output<string | undefined>;
    /**
     * Password to login with.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The profile name as set in the shared config file.
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project to login with.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * The name of the project to login with.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The HuaweiCloud region to connect to.
     */
    public readonly region!: pulumi.Output<string | undefined>;
    /**
     * The secret key of the HuaweiCloud to use.
     */
    public readonly secretKey!: pulumi.Output<string | undefined>;
    /**
     * The security token to authenticate with a temporary security credential.
     */
    public readonly securityToken!: pulumi.Output<string | undefined>;
    /**
     * The path to the shared config file. If not set, the default is ~/.hcloud/config.json.
     */
    public readonly sharedConfigFile!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Tenant (Identity v2) to login with.
     */
    public readonly tenantId!: pulumi.Output<string | undefined>;
    /**
     * The name of the Tenant (Identity v2) to login with.
     */
    public readonly tenantName!: pulumi.Output<string | undefined>;
    /**
     * Authentication token to use as an alternative to username/password.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * User ID to login with.
     */
    public readonly userId!: pulumi.Output<string | undefined>;
    /**
     * Username to login with.
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            resourceInputs["accessKey"] = args ? args.accessKey : undefined;
            resourceInputs["agencyDomainName"] = args ? args.agencyDomainName : undefined;
            resourceInputs["agencyName"] = args ? args.agencyName : undefined;
            resourceInputs["assumeRole"] = pulumi.output(args ? args.assumeRole : undefined).apply(JSON.stringify);
            resourceInputs["authUrl"] = args ? args.authUrl : undefined;
            resourceInputs["cacertFile"] = args ? args.cacertFile : undefined;
            resourceInputs["cert"] = args ? args.cert : undefined;
            resourceInputs["cloud"] = args ? args.cloud : undefined;
            resourceInputs["delegatedProject"] = args ? args.delegatedProject : undefined;
            resourceInputs["domainId"] = args ? args.domainId : undefined;
            resourceInputs["domainName"] = args ? args.domainName : undefined;
            resourceInputs["endpoints"] = pulumi.output(args ? args.endpoints : undefined).apply(JSON.stringify);
            resourceInputs["enterpriseProjectId"] = args ? args.enterpriseProjectId : undefined;
            resourceInputs["insecure"] = pulumi.output(args ? args.insecure : undefined).apply(JSON.stringify);
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["maxRetries"] = pulumi.output(args ? args.maxRetries : undefined).apply(JSON.stringify);
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["regional"] = pulumi.output(args ? args.regional : undefined).apply(JSON.stringify);
            resourceInputs["secretKey"] = args ? args.secretKey : undefined;
            resourceInputs["securityToken"] = args ? args.securityToken : undefined;
            resourceInputs["sharedConfigFile"] = args ? args.sharedConfigFile : undefined;
            resourceInputs["tenantId"] = args ? args.tenantId : undefined;
            resourceInputs["tenantName"] = args ? args.tenantName : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * The access key of the HuaweiCloud to use.
     */
    accessKey?: pulumi.Input<string>;
    /**
     * The name of domain who created the agency (Identity v3).
     */
    agencyDomainName?: pulumi.Input<string>;
    /**
     * The name of agency
     */
    agencyName?: pulumi.Input<string>;
    assumeRole?: pulumi.Input<inputs.ProviderAssumeRole>;
    /**
     * The Identity authentication URL.
     */
    authUrl?: pulumi.Input<string>;
    /**
     * A Custom CA certificate.
     */
    cacertFile?: pulumi.Input<string>;
    /**
     * A client certificate to authenticate with.
     */
    cert?: pulumi.Input<string>;
    /**
     * The endpoint of cloud provider, defaults to myhuaweicloud.com
     */
    cloud?: pulumi.Input<string>;
    /**
     * The name of delegated project (Identity v3).
     */
    delegatedProject?: pulumi.Input<string>;
    /**
     * The ID of the Domain to scope to.
     */
    domainId?: pulumi.Input<string>;
    /**
     * The name of the Domain to scope to.
     */
    domainName?: pulumi.Input<string>;
    /**
     * The custom endpoints used to override the default endpoint URL.
     */
    endpoints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * enterprise project id
     */
    enterpriseProjectId?: pulumi.Input<string>;
    /**
     * Trust self-signed certificates.
     */
    insecure?: pulumi.Input<boolean>;
    /**
     * A client private key to authenticate with.
     */
    key?: pulumi.Input<string>;
    /**
     * How many times HTTP connection should be retried until giving up.
     */
    maxRetries?: pulumi.Input<number>;
    /**
     * Password to login with.
     */
    password?: pulumi.Input<string>;
    /**
     * The profile name as set in the shared config file.
     */
    profile?: pulumi.Input<string>;
    /**
     * The ID of the project to login with.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of the project to login with.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The HuaweiCloud region to connect to.
     */
    region?: pulumi.Input<string>;
    /**
     * Whether the service endpoints are regional
     */
    regional?: pulumi.Input<boolean>;
    /**
     * The secret key of the HuaweiCloud to use.
     */
    secretKey?: pulumi.Input<string>;
    /**
     * The security token to authenticate with a temporary security credential.
     */
    securityToken?: pulumi.Input<string>;
    /**
     * The path to the shared config file. If not set, the default is ~/.hcloud/config.json.
     */
    sharedConfigFile?: pulumi.Input<string>;
    /**
     * The ID of the Tenant (Identity v2) to login with.
     */
    tenantId?: pulumi.Input<string>;
    /**
     * The name of the Tenant (Identity v2) to login with.
     */
    tenantName?: pulumi.Input<string>;
    /**
     * Authentication token to use as an alternative to username/password.
     */
    token?: pulumi.Input<string>;
    /**
     * User ID to login with.
     */
    userId?: pulumi.Input<string>;
    /**
     * Username to login with.
     */
    userName?: pulumi.Input<string>;
}

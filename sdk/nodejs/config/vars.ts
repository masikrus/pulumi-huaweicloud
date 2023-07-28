// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("huaweicloud");

/**
 * The access key of the HuaweiCloud to use.
 */
export declare const accessKey: string | undefined;
Object.defineProperty(exports, "accessKey", {
    get() {
        return __config.get("accessKey");
    },
    enumerable: true,
});

/**
 * The name of domain who created the agency (Identity v3).
 */
export declare const agencyDomainName: string | undefined;
Object.defineProperty(exports, "agencyDomainName", {
    get() {
        return __config.get("agencyDomainName");
    },
    enumerable: true,
});

/**
 * The name of agency
 */
export declare const agencyName: string | undefined;
Object.defineProperty(exports, "agencyName", {
    get() {
        return __config.get("agencyName");
    },
    enumerable: true,
});

export declare const assumeRole: outputs.config.AssumeRole | undefined;
Object.defineProperty(exports, "assumeRole", {
    get() {
        return __config.getObject<outputs.config.AssumeRole>("assumeRole");
    },
    enumerable: true,
});

/**
 * The Identity authentication URL.
 */
export declare const authUrl: string | undefined;
Object.defineProperty(exports, "authUrl", {
    get() {
        return __config.get("authUrl");
    },
    enumerable: true,
});

/**
 * A Custom CA certificate.
 */
export declare const cacertFile: string | undefined;
Object.defineProperty(exports, "cacertFile", {
    get() {
        return __config.get("cacertFile");
    },
    enumerable: true,
});

/**
 * A client certificate to authenticate with.
 */
export declare const cert: string | undefined;
Object.defineProperty(exports, "cert", {
    get() {
        return __config.get("cert");
    },
    enumerable: true,
});

/**
 * The endpoint of cloud provider, defaults to myhuaweicloud.com
 */
export declare const cloud: string | undefined;
Object.defineProperty(exports, "cloud", {
    get() {
        return __config.get("cloud");
    },
    enumerable: true,
});

/**
 * The name of delegated project (Identity v3).
 */
export declare const delegatedProject: string | undefined;
Object.defineProperty(exports, "delegatedProject", {
    get() {
        return __config.get("delegatedProject");
    },
    enumerable: true,
});

/**
 * The ID of the Domain to scope to.
 */
export declare const domainId: string | undefined;
Object.defineProperty(exports, "domainId", {
    get() {
        return __config.get("domainId");
    },
    enumerable: true,
});

/**
 * The name of the Domain to scope to.
 */
export declare const domainName: string | undefined;
Object.defineProperty(exports, "domainName", {
    get() {
        return __config.get("domainName");
    },
    enumerable: true,
});

/**
 * The custom endpoints used to override the default endpoint URL.
 */
export declare const endpoints: {[key: string]: string} | undefined;
Object.defineProperty(exports, "endpoints", {
    get() {
        return __config.getObject<{[key: string]: string}>("endpoints");
    },
    enumerable: true,
});

/**
 * enterprise project id
 */
export declare const enterpriseProjectId: string | undefined;
Object.defineProperty(exports, "enterpriseProjectId", {
    get() {
        return __config.get("enterpriseProjectId");
    },
    enumerable: true,
});

/**
 * Trust self-signed certificates.
 */
export declare const insecure: boolean | undefined;
Object.defineProperty(exports, "insecure", {
    get() {
        return __config.getObject<boolean>("insecure");
    },
    enumerable: true,
});

/**
 * A client private key to authenticate with.
 */
export declare const key: string | undefined;
Object.defineProperty(exports, "key", {
    get() {
        return __config.get("key");
    },
    enumerable: true,
});

/**
 * How many times HTTP connection should be retried until giving up.
 */
export declare const maxRetries: number | undefined;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries");
    },
    enumerable: true,
});

/**
 * Password to login with.
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

/**
 * The profile name as set in the shared config file.
 */
export declare const profile: string | undefined;
Object.defineProperty(exports, "profile", {
    get() {
        return __config.get("profile");
    },
    enumerable: true,
});

/**
 * The ID of the project to login with.
 */
export declare const projectId: string | undefined;
Object.defineProperty(exports, "projectId", {
    get() {
        return __config.get("projectId");
    },
    enumerable: true,
});

/**
 * The name of the project to login with.
 */
export declare const projectName: string | undefined;
Object.defineProperty(exports, "projectName", {
    get() {
        return __config.get("projectName");
    },
    enumerable: true,
});

/**
 * The HuaweiCloud region to connect to.
 */
export declare const region: string | undefined;
Object.defineProperty(exports, "region", {
    get() {
        return __config.get("region");
    },
    enumerable: true,
});

/**
 * Whether the service endpoints are regional
 */
export declare const regional: boolean | undefined;
Object.defineProperty(exports, "regional", {
    get() {
        return __config.getObject<boolean>("regional");
    },
    enumerable: true,
});

/**
 * The secret key of the HuaweiCloud to use.
 */
export declare const secretKey: string | undefined;
Object.defineProperty(exports, "secretKey", {
    get() {
        return __config.get("secretKey");
    },
    enumerable: true,
});

/**
 * The security token to authenticate with a temporary security credential.
 */
export declare const securityToken: string | undefined;
Object.defineProperty(exports, "securityToken", {
    get() {
        return __config.get("securityToken");
    },
    enumerable: true,
});

/**
 * The path to the shared config file. If not set, the default is ~/.hcloud/config.json.
 */
export declare const sharedConfigFile: string | undefined;
Object.defineProperty(exports, "sharedConfigFile", {
    get() {
        return __config.get("sharedConfigFile");
    },
    enumerable: true,
});

/**
 * The ID of the Tenant (Identity v2) to login with.
 */
export declare const tenantId: string | undefined;
Object.defineProperty(exports, "tenantId", {
    get() {
        return __config.get("tenantId");
    },
    enumerable: true,
});

/**
 * The name of the Tenant (Identity v2) to login with.
 */
export declare const tenantName: string | undefined;
Object.defineProperty(exports, "tenantName", {
    get() {
        return __config.get("tenantName");
    },
    enumerable: true,
});

/**
 * Authentication token to use as an alternative to username/password.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

/**
 * User ID to login with.
 */
export declare const userId: string | undefined;
Object.defineProperty(exports, "userId", {
    get() {
        return __config.get("userId");
    },
    enumerable: true,
});

/**
 * Username to login with.
 */
export declare const userName: string | undefined;
Object.defineProperty(exports, "userName", {
    get() {
        return __config.get("userName");
    },
    enumerable: true,
});


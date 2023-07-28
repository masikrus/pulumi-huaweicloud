// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a meeting user resource within HuaweiCloud.
 *
 * ## Example Usage
 * ### Create a user using third-party application authorization and specifies the account parameters
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const appId = config.requireObject("appId");
 * const appKey = config.requireObject("appKey");
 * const userAccount = config.requireObject("userAccount");
 * const thirdAccount = config.requireObject("thirdAccount");
 * const userName = config.requireObject("userName");
 * const userPassword = config.requireObject("userPassword");
 * const englishName = config.requireObject("englishName");
 * const signature = config.requireObject("signature");
 * const title = config.requireObject("title");
 * const test = new huaweicloud.meeting.User("test", {
 *     appId: appId,
 *     appKey: appKey,
 *     account: userAccount,
 *     thirdAccount: thirdAccount,
 *     password: userPassword,
 *     country: "chinaPR",
 *     description: "Created by script",
 *     email: "123456789@example.com",
 *     englishName: englishName,
 *     phone: "+8612345678987",
 *     signature: signature,
 *     title: title,
 *     sortLevel: 5,
 * });
 * ```
 * ### Create a user using account authorization
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@huaweicloudos/pulumi";
 *
 * const config = new pulumi.Config();
 * const accountName = config.requireObject("accountName");
 * const accountPassword = config.requireObject("accountPassword");
 * const userName = config.requireObject("userName");
 * const userPassword = config.requireObject("userPassword");
 * const englishName = config.requireObject("englishName");
 * const signature = config.requireObject("signature");
 * const title = config.requireObject("title");
 * const test = new huaweicloud.meeting.User("test", {
 *     accountName: accountName,
 *     accountPassword: accountPassword,
 *     password: userPassword,
 *     country: "chinaPR",
 *     email: "123456789@example.com",
 *     englishName: englishName,
 *     phone: "+8612345678987",
 *     signature: signature,
 *     title: title,
 *     sortLevel: 5,
 * });
 * ```
 * ## Appendix
 *
 * <a name="phoneNumberMapping"></a>
 * The countries (or regions) and phone numbers mapping relationship supports:
 *
 * | Country Or Region | Country Code |
 * | ---- | ---- |
 * | chinaPR | +86 |
 * | chinaHKG | +852|
 * | chinaOMA | +853 |
 * | chinaTPE | +886 |
 * | BVl | +1284  |
 * | Bolivia | +591 |
 * | CZ | +420 |
 * | GB | +245 |
 * | SVGrenadines | +1784 |
 * | TAT | +1868 |
 * | UK | +44 |
 * | afghanistan | +93 |
 * | albania | +355 |
 * | algeria | +213 |
 * | andorra | +376 |
 * | angola | +244 |
 * | argentina | +54 |
 * | armenia | +374 |
 * | australia | +61 |
 * | austria | +43 |
 * | azerbaijan | +994 |
 * | bahamas | +1242 |
 * | bahrain | +973 |
 * | bangladesh | +880 |
 * | belarus | +375 |
 * | belgium | +32 |
 * | belize | +501 |
 * | benin | +229 |
 * | bosniaAndHerzegovina | +387 |
 * | botswana | +267 |
 * | brazil | +55 |
 * | brunei | +673 |
 * | bulgaria | +359 |
 * | burkinaFaso | +226 |
 * | burundi | +257 |
 * | cambodia | +855 |
 * | cameroon | +237 |
 * | canada | +1 |
 * | capeVerde | +238 |
 * | caymanIslands | +1345 |
 * | centralAfrican | +236 |
 * | chad | +235 |
 * | chile | +56 |
 * | columbia | +57 |
 * | comoros | +269 |
 * | congoB | +242 |
 * | congoJ | +243 |
 * | costarica | +506 |
 * | croatia | +385 |
 * | curacao | +599 |
 * | cyprus | +357 |
 * | denmark | +45 |
 * | djibouti | +253 |
 * | dominica | +1809 |
 * | ecuador | +593 |
 * | egypt | +20 |
 * | equatorialGuinea | +240 |
 * | estonia | +372 |
 * | finland | +358 |
 * | france | +33 |
 * | gabon | +241 |
 * | gambia | +220 |
 * | georgia | +995 |
 * | germany | +49 |
 * | ghana | +233 |
 * | greece | +30 |
 * | grenada | +1473 |
 * | guatemala | +502 |
 * | guinea | +224 |
 * | guyana | +592 |
 * | honduras | +504 |
 * | hungary | +36 |
 * | india | +91 |
 * | indonesia | +62 |
 * | iraq | +964 |
 * | ireland | +353 |
 * | israel | +972 |
 * | italy | +39 |
 * | ivoryCoast | +225 |
 * | jamaica | +1876 |
 * | japan | +81 |
 * | jordan | +962 |
 * | kazakhstan | +7 |
 * | kenya | +254 |
 * | kuwait | +965 |
 * | kyrgyzstan | +996 |
 * | laos | +856 |
 * | latvia | +371 |
 * | lebanon | +961 |
 * | lesotho | +266 |
 * | liberia | +231 |
 * | libya | +218 |
 * | lithuania | +370 |
 * | luxembourg | +352 |
 * | macedonia | +389 |
 * | madagascar | +261 |
 * | malawi | +265 |
 * | malaysia | +60 |
 * | maldives | +960 |
 * | mali | +223 |
 * | malta | +356 |
 * | mauritania | +222 |
 * | mauritius | +230 |
 * | mexico | +52 |
 * | moldova | +373 |
 * | mongolia | +976 |
 * | montenegro | +382  |
 * | morocco | +212 |
 * | mozambique | +258 |
 * | myanmar | +95 |
 * | namibia | +264 |
 * | nepal | +977 |
 * | netherlands | +31 |
 * | newZealand | +64 |
 * | nicaragua | +505 |
 * | niger | +227 |
 * | nigeria | +234 |
 * | norway | +47 |
 * | oman | +968 |
 * | pakistan | +92 |
 * | palestine | +970 |
 * | panama | +507 |
 * | papuaNewGuinea | +675 |
 * | peru | +51 |
 * | philippines | +63 |
 * | poland | +48 |
 * | portugal | +351 |
 * | puertoRico | +1787 |
 * | qatar | +974 |
 * | romania | +40 |
 * | russia | +7 |
 * | rwanda | +250 |
 * | saintMartin | +590 |
 * | salvatore | +503 |
 * | saudiArabia | +966 |
 * | senegal | +221 |
 * | serbia | +381 |
 * | seychelles | +248 |
 * | sierraLeone | +232 |
 * | singapore | +65 |
 * | slovakia | +421 |
 * | slovenia | +386 |
 * | somalia | +252 |
 * | southAfrica | +27 |
 * | southKorea | +82 |
 * | spain | +34 |
 * | sriLanka | +94 |
 * | suriname | +597 |
 * | swaziland | +268 |
 * | sweden | +46 |
 * | switzerland | +41 |
 * | tajikistan | +992 |
 * | tanzania | +255 |
 * | thailand | +66 |
 * | togo | +228 |
 * | tunisia | +216 |
 * | turkey | +90 |
 * | turkmenistan | +993 |
 * | uae | +971 |
 * | uganda | +256 |
 * | ukraine | +380 |
 * | uruguay | +598 |
 * | usa | +1 |
 * | uzbekistan | +998 |
 * | venezuela | +58 |
 * | vietNam | +84 |
 * | yemen | +967 |
 * | zambia | +260 |
 * | zimbabwe | +263 |
 *
 * ## Import
 *
 * Users can be imported using their `id` and authorization parameters, separated by slashes, e.g. Import a user and authenticated by account.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/user:User test &ltid&gt/&ltaccount_name&gt/&ltaccount_password&gt
 * ```
 *
 *  Import a user and authenticated by `APP ID`/`APP Key`.
 *
 * ```sh
 *  $ pulumi import huaweicloud:Meeting/user:User test &ltid&gt/&ltapp_id&gt/&ltapp_key&gt/&ltcorp_id&gt/&ltuser_id&gt
 * ```
 *
 *  The slashes cannot be missing even corporation ID and user ID are empty. Note that some parameters do not support import due to missing API responses or privacy, such as `password`, `is_send_notify` and `is_admin`. You can ignore this change as below. resource "huaweicloud_meeting_user" "test" {
 *
 *  ...
 *
 *  lifecycle {
 *
 *  ignore_changes = [
 *
 *  password, is_send_notify,
 *
 *  ]
 *
 *  } }
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'huaweicloud:Meeting/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Specifies the user account. The value can contain **1** to **64** characters.
     * If omitted, the service will automatically generate a value.
     * Changing this parameter will create a new resource.
     */
    public readonly account!: pulumi.Output<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * administrator belongs. Changing this parameter will create a new resource.
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
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    public readonly corpId!: pulumi.Output<string>;
    /**
     * Specifies the country to which the phone number belongs to.
     */
    public readonly country!: pulumi.Output<string>;
    /**
     * Specifies the department code. Defaults to **1** (Root department).
     */
    public readonly departmentCode!: pulumi.Output<string>;
    /**
     * The department name.
     */
    public /*out*/ readonly departmentName!: pulumi.Output<string>;
    /**
     * The department full name.
     */
    public /*out*/ readonly departmentNamePath!: pulumi.Output<string>;
    /**
     * Specifies the description. The value can contain **0** to **128** characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the email address.
     */
    public readonly email!: pulumi.Output<string>;
    /**
     * Specifies the english name. The value can contain **0** to **64** characters.
     */
    public readonly englishName!: pulumi.Output<string>;
    /**
     * Specifies whether to hide the phone number.
     */
    public readonly hidePhone!: pulumi.Output<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    public readonly isAdmin!: pulumi.Output<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    public readonly isSendNotify!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the user name. The value can contain **1** to **64** characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Specifies the user password.
     * The following conditions must be met:
     * + **8** to **32** characters
     * + It cannot be consistent with the positive and reverse order of the `account` parameter.
     * + Contains at least two character types: lowercase letters, uppercase letters, numbers, special characters
     * (**`~!@#$%^&*()-_=+|[{}];:",'<.>/?**).
     */
    public readonly password!: pulumi.Output<string>;
    /**
     * Specifies the phone number.
     * The phone number must start with a country (region) code.
     */
    public readonly phone!: pulumi.Output<string>;
    /**
     * Specifies the signature. The value can contain **0** to **512** characters.
     */
    public readonly signature!: pulumi.Output<string>;
    /**
     * The SIP number.
     */
    public /*out*/ readonly sipNumber!: pulumi.Output<string>;
    /**
     * Specifies the address book sorting level.
     * The lower the serial number, the higher the priority.
     * The valid value is range from **1** to **10000**. Defaults to **10000**.
     */
    public readonly sortLevel!: pulumi.Output<number>;
    /**
     * Specifies the status. The valid values are as follows:
     * + **0**: Normal.
     * + **1**: Disable.
     */
    public readonly status!: pulumi.Output<number>;
    /**
     * Specifies the third-party account name.
     */
    public readonly thirdAccount!: pulumi.Output<string>;
    /**
     * Specifies the title name. The value can contain **0** to **32** characters.
     */
    public readonly title!: pulumi.Output<string>;
    /**
     * The user type.
     * + **2**: Enterprise member account.
     */
    public /*out*/ readonly type!: pulumi.Output<number>;
    /**
     * Specifies the user ID of the administrator.
     * Only available if `appId` is set. If omitted, the user ID of default administrator will be used.
     * Changing this parameter will create a new resource.
     */
    public readonly userId!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            resourceInputs["account"] = state ? state.account : undefined;
            resourceInputs["accountName"] = state ? state.accountName : undefined;
            resourceInputs["accountPassword"] = state ? state.accountPassword : undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["appKey"] = state ? state.appKey : undefined;
            resourceInputs["corpId"] = state ? state.corpId : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["departmentCode"] = state ? state.departmentCode : undefined;
            resourceInputs["departmentName"] = state ? state.departmentName : undefined;
            resourceInputs["departmentNamePath"] = state ? state.departmentNamePath : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["englishName"] = state ? state.englishName : undefined;
            resourceInputs["hidePhone"] = state ? state.hidePhone : undefined;
            resourceInputs["isAdmin"] = state ? state.isAdmin : undefined;
            resourceInputs["isSendNotify"] = state ? state.isSendNotify : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["phone"] = state ? state.phone : undefined;
            resourceInputs["signature"] = state ? state.signature : undefined;
            resourceInputs["sipNumber"] = state ? state.sipNumber : undefined;
            resourceInputs["sortLevel"] = state ? state.sortLevel : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["thirdAccount"] = state ? state.thirdAccount : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["userId"] = state ? state.userId : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.password === undefined) && !opts.urn) {
                throw new Error("Missing required property 'password'");
            }
            resourceInputs["account"] = args ? args.account : undefined;
            resourceInputs["accountName"] = args ? args.accountName : undefined;
            resourceInputs["accountPassword"] = args ? args.accountPassword : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["appKey"] = args ? args.appKey : undefined;
            resourceInputs["corpId"] = args ? args.corpId : undefined;
            resourceInputs["country"] = args ? args.country : undefined;
            resourceInputs["departmentCode"] = args ? args.departmentCode : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["englishName"] = args ? args.englishName : undefined;
            resourceInputs["hidePhone"] = args ? args.hidePhone : undefined;
            resourceInputs["isAdmin"] = args ? args.isAdmin : undefined;
            resourceInputs["isSendNotify"] = args ? args.isSendNotify : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["phone"] = args ? args.phone : undefined;
            resourceInputs["signature"] = args ? args.signature : undefined;
            resourceInputs["sortLevel"] = args ? args.sortLevel : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["thirdAccount"] = args ? args.thirdAccount : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["userId"] = args ? args.userId : undefined;
            resourceInputs["departmentName"] = undefined /*out*/;
            resourceInputs["departmentNamePath"] = undefined /*out*/;
            resourceInputs["sipNumber"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(User.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Specifies the user account. The value can contain **1** to **64** characters.
     * If omitted, the service will automatically generate a value.
     * Changing this parameter will create a new resource.
     */
    account?: pulumi.Input<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * administrator belongs. Changing this parameter will create a new resource.
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
    /**
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    corpId?: pulumi.Input<string>;
    /**
     * Specifies the country to which the phone number belongs to.
     */
    country?: pulumi.Input<string>;
    /**
     * Specifies the department code. Defaults to **1** (Root department).
     */
    departmentCode?: pulumi.Input<string>;
    /**
     * The department name.
     */
    departmentName?: pulumi.Input<string>;
    /**
     * The department full name.
     */
    departmentNamePath?: pulumi.Input<string>;
    /**
     * Specifies the description. The value can contain **0** to **128** characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Specifies the english name. The value can contain **0** to **64** characters.
     */
    englishName?: pulumi.Input<string>;
    /**
     * Specifies whether to hide the phone number.
     */
    hidePhone?: pulumi.Input<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    isSendNotify?: pulumi.Input<boolean>;
    /**
     * Specifies the user name. The value can contain **1** to **64** characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * The following conditions must be met:
     * + **8** to **32** characters
     * + It cannot be consistent with the positive and reverse order of the `account` parameter.
     * + Contains at least two character types: lowercase letters, uppercase letters, numbers, special characters
     * (**`~!@#$%^&*()-_=+|[{}];:",'<.>/?**).
     */
    password?: pulumi.Input<string>;
    /**
     * Specifies the phone number.
     * The phone number must start with a country (region) code.
     */
    phone?: pulumi.Input<string>;
    /**
     * Specifies the signature. The value can contain **0** to **512** characters.
     */
    signature?: pulumi.Input<string>;
    /**
     * The SIP number.
     */
    sipNumber?: pulumi.Input<string>;
    /**
     * Specifies the address book sorting level.
     * The lower the serial number, the higher the priority.
     * The valid value is range from **1** to **10000**. Defaults to **10000**.
     */
    sortLevel?: pulumi.Input<number>;
    /**
     * Specifies the status. The valid values are as follows:
     * + **0**: Normal.
     * + **1**: Disable.
     */
    status?: pulumi.Input<number>;
    /**
     * Specifies the third-party account name.
     */
    thirdAccount?: pulumi.Input<string>;
    /**
     * Specifies the title name. The value can contain **0** to **32** characters.
     */
    title?: pulumi.Input<string>;
    /**
     * The user type.
     * + **2**: Enterprise member account.
     */
    type?: pulumi.Input<number>;
    /**
     * Specifies the user ID of the administrator.
     * Only available if `appId` is set. If omitted, the user ID of default administrator will be used.
     * Changing this parameter will create a new resource.
     */
    userId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Specifies the user account. The value can contain **1** to **64** characters.
     * If omitted, the service will automatically generate a value.
     * Changing this parameter will create a new resource.
     */
    account?: pulumi.Input<string>;
    /**
     * Specifies the (HUAWEI Cloud meeting) user account name to which the
     * administrator belongs. Changing this parameter will create a new resource.
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
    /**
     * Specifies the corporation ID.
     * Required if the application is used in multiple enterprises. Only available if `appId` is set.
     * Changing this parameter will create a new resource.
     */
    corpId?: pulumi.Input<string>;
    /**
     * Specifies the country to which the phone number belongs to.
     */
    country?: pulumi.Input<string>;
    /**
     * Specifies the department code. Defaults to **1** (Root department).
     */
    departmentCode?: pulumi.Input<string>;
    /**
     * Specifies the description. The value can contain **0** to **128** characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies the email address.
     */
    email?: pulumi.Input<string>;
    /**
     * Specifies the english name. The value can contain **0** to **64** characters.
     */
    englishName?: pulumi.Input<string>;
    /**
     * Specifies whether to hide the phone number.
     */
    hidePhone?: pulumi.Input<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    isAdmin?: pulumi.Input<boolean>;
    /**
     * Specifies whether to send email and SMS notifications for account opening.
     * Defaults to **true**.
     */
    isSendNotify?: pulumi.Input<boolean>;
    /**
     * Specifies the user name. The value can contain **1** to **64** characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the user password.
     * The following conditions must be met:
     * + **8** to **32** characters
     * + It cannot be consistent with the positive and reverse order of the `account` parameter.
     * + Contains at least two character types: lowercase letters, uppercase letters, numbers, special characters
     * (**`~!@#$%^&*()-_=+|[{}];:",'<.>/?**).
     */
    password: pulumi.Input<string>;
    /**
     * Specifies the phone number.
     * The phone number must start with a country (region) code.
     */
    phone?: pulumi.Input<string>;
    /**
     * Specifies the signature. The value can contain **0** to **512** characters.
     */
    signature?: pulumi.Input<string>;
    /**
     * Specifies the address book sorting level.
     * The lower the serial number, the higher the priority.
     * The valid value is range from **1** to **10000**. Defaults to **10000**.
     */
    sortLevel?: pulumi.Input<number>;
    /**
     * Specifies the status. The valid values are as follows:
     * + **0**: Normal.
     * + **1**: Disable.
     */
    status?: pulumi.Input<number>;
    /**
     * Specifies the third-party account name.
     */
    thirdAccount?: pulumi.Input<string>;
    /**
     * Specifies the title name. The value can contain **0** to **32** characters.
     */
    title?: pulumi.Input<string>;
    /**
     * Specifies the user ID of the administrator.
     * Only available if `appId` is set. If omitted, the user ID of default administrator will be used.
     * Changing this parameter will create a new resource.
     */
    userId?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing a Google IDP on the instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getIdpGoogle({
 *     id: "123456789012345678",
 * });
 * ```
 */
export function getIdpGoogle(args: GetIdpGoogleArgs, opts?: pulumi.InvokeOptions): Promise<GetIdpGoogleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getIdpGoogle:getIdpGoogle", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdpGoogle.
 */
export interface GetIdpGoogleArgs {
    /**
     * The ID of this resource.
     */
    id: string;
}

/**
 * A collection of values returned by getIdpGoogle.
 */
export interface GetIdpGoogleResult {
    /**
     * client id generated by the identity provider
     */
    readonly clientId: string;
    /**
     * client secret generated by the identity provider
     */
    readonly clientSecret: string;
    /**
     * The ID of this resource.
     */
    readonly id: string;
    /**
     * enabled if a new account in ZITADEL are created automatically on login with an external account
     */
    readonly isAutoCreation: boolean;
    /**
     * enabled if a the ZITADEL account fields are updated automatically on each login
     */
    readonly isAutoUpdate: boolean;
    /**
     * enabled if users are able to create a new account in ZITADEL when using an external account
     */
    readonly isCreationAllowed: boolean;
    /**
     * enabled if users are able to link an existing ZITADEL user with an external account
     */
    readonly isLinkingAllowed: boolean;
    /**
     * Name of the IDP
     */
    readonly name: string;
    /**
     * the scopes requested by ZITADEL during the request on the identity provider
     */
    readonly scopes: string[];
}
/**
 * Datasource representing a Google IDP on the instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getIdpGoogle({
 *     id: "123456789012345678",
 * });
 * ```
 */
export function getIdpGoogleOutput(args: GetIdpGoogleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIdpGoogleResult> {
    return pulumi.output(args).apply((a: any) => getIdpGoogle(a, opts))
}

/**
 * A collection of arguments for invoking getIdpGoogle.
 */
export interface GetIdpGoogleOutputArgs {
    /**
     * The ID of this resource.
     */
    id: pulumi.Input<string>;
}

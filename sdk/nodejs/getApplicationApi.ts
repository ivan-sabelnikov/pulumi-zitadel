// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Datasource representing an API application belonging to a project, with all configuration possibilities.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getApplicationApi({
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: data.zitadel_project["default"].id,
 *     appId: "123456789012345678",
 * });
 * export const applicationApi = _default;
 * ```
 */
export function getApplicationApi(args: GetApplicationApiArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationApiResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zitadel:index/getApplicationApi:getApplicationApi", {
        "appId": args.appId,
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApplicationApi.
 */
export interface GetApplicationApiArgs {
    /**
     * The ID of this resource.
     */
    appId: string;
    /**
     * ID of the organization
     */
    orgId?: string;
    /**
     * ID of the project
     */
    projectId: string;
}

/**
 * A collection of values returned by getApplicationApi.
 */
export interface GetApplicationApiResult {
    /**
     * The ID of this resource.
     */
    readonly appId: string;
    /**
     * Auth method type
     */
    readonly authMethodType: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Name of the application
     */
    readonly name: string;
    /**
     * ID of the organization
     */
    readonly orgId?: string;
    /**
     * ID of the project
     */
    readonly projectId: string;
}
/**
 * Datasource representing an API application belonging to a project, with all configuration possibilities.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const default = zitadel.getApplicationApi({
 *     orgId: data.zitadel_org["default"].id,
 *     projectId: data.zitadel_project["default"].id,
 *     appId: "123456789012345678",
 * });
 * export const applicationApi = _default;
 * ```
 */
export function getApplicationApiOutput(args: GetApplicationApiOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetApplicationApiResult> {
    return pulumi.output(args).apply((a: any) => getApplicationApi(a, opts))
}

/**
 * A collection of arguments for invoking getApplicationApi.
 */
export interface GetApplicationApiOutputArgs {
    /**
     * The ID of this resource.
     */
    appId: pulumi.Input<string>;
    /**
     * ID of the organization
     */
    orgId?: pulumi.Input<string>;
    /**
     * ID of the project
     */
    projectId: pulumi.Input<string>;
}

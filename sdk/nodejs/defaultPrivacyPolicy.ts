// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource representing the default privacy policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zitadel from "@pulumi/zitadel";
 *
 * const defaultDefaultPrivacyPolicy = new zitadel.DefaultPrivacyPolicy("default", {
 *     helpLink: "https://example.com/help",
 *     privacyLink: "https://example.com/privacy",
 *     supportEmail: "support@example.com",
 *     tosLink: "https://example.com/tos",
 * });
 * ```
 *
 * ## Import
 *
 * terraform # The resource can be imported using the ID format `<>`, e.g.
 *
 * ```sh
 *  $ pulumi import zitadel:index/defaultPrivacyPolicy:DefaultPrivacyPolicy imported ''
 * ```
 */
export class DefaultPrivacyPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DefaultPrivacyPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultPrivacyPolicyState, opts?: pulumi.CustomResourceOptions): DefaultPrivacyPolicy {
        return new DefaultPrivacyPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zitadel:index/defaultPrivacyPolicy:DefaultPrivacyPolicy';

    /**
     * Returns true if the given object is an instance of DefaultPrivacyPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultPrivacyPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultPrivacyPolicy.__pulumiType;
    }

    public readonly helpLink!: pulumi.Output<string | undefined>;
    public readonly privacyLink!: pulumi.Output<string | undefined>;
    public readonly supportEmail!: pulumi.Output<string | undefined>;
    public readonly tosLink!: pulumi.Output<string | undefined>;

    /**
     * Create a DefaultPrivacyPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DefaultPrivacyPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultPrivacyPolicyArgs | DefaultPrivacyPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultPrivacyPolicyState | undefined;
            resourceInputs["helpLink"] = state ? state.helpLink : undefined;
            resourceInputs["privacyLink"] = state ? state.privacyLink : undefined;
            resourceInputs["supportEmail"] = state ? state.supportEmail : undefined;
            resourceInputs["tosLink"] = state ? state.tosLink : undefined;
        } else {
            const args = argsOrState as DefaultPrivacyPolicyArgs | undefined;
            resourceInputs["helpLink"] = args ? args.helpLink : undefined;
            resourceInputs["privacyLink"] = args ? args.privacyLink : undefined;
            resourceInputs["supportEmail"] = args ? args.supportEmail : undefined;
            resourceInputs["tosLink"] = args ? args.tosLink : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultPrivacyPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultPrivacyPolicy resources.
 */
export interface DefaultPrivacyPolicyState {
    helpLink?: pulumi.Input<string>;
    privacyLink?: pulumi.Input<string>;
    supportEmail?: pulumi.Input<string>;
    tosLink?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DefaultPrivacyPolicy resource.
 */
export interface DefaultPrivacyPolicyArgs {
    helpLink?: pulumi.Input<string>;
    privacyLink?: pulumi.Input<string>;
    supportEmail?: pulumi.Input<string>;
    tosLink?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the Certificate Authentication Profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getCertificateAuthenticationProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getCertificateAuthenticationProfile(args?: GetCertificateAuthenticationProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetCertificateAuthenticationProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificateAuthenticationProfile.
 */
export interface GetCertificateAuthenticationProfileArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the certificate profile
     */
    name?: string;
}

/**
 * A collection of values returned by getCertificateAuthenticationProfile.
 */
export interface GetCertificateAuthenticationProfileResult {
    /**
     * Allow as username
     */
    readonly allowedAsUserName: boolean;
    /**
     * Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in `usernameFrom`.
     */
    readonly certificateAttributeName: string;
    /**
     * Description
     */
    readonly description: string;
    /**
     * Referred IDStore name for the Certificate Profile or `[not applicable]` in case no identity store is chosen
     */
    readonly externalIdentityStoreName: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Match mode of the Certificate Profile. Allowed values: NEVER, RESOLVE*IDENTITY*AMBIGUITY, BINARY_COMPARISON
     */
    readonly matchMode: string;
    /**
     * The name of the certificate profile
     */
    readonly name: string;
    /**
     * The attribute in the certificate where the user name should be taken from. Allowed values: `CERTIFICATE` (for a specific attribute as defined in certificateAttributeName), `UPN` (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
     */
    readonly usernameFrom: string;
}
/**
 * This data source can read the Certificate Authentication Profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getCertificateAuthenticationProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getCertificateAuthenticationProfileOutput(args?: GetCertificateAuthenticationProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCertificateAuthenticationProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCertificateAuthenticationProfile.
 */
export interface GetCertificateAuthenticationProfileOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the certificate profile
     */
    name?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the TACACS Profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.Tacacs.getProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getProfile(args?: GetProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetProfileResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:Tacacs/getProfile:getProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProfile.
 */
export interface GetProfileArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the TACACS profile
     */
    name?: string;
}

/**
 * A collection of values returned by getProfile.
 */
export interface GetProfileResult {
    /**
     * Description
     */
    readonly description: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * The name of the TACACS profile
     */
    readonly name: string;
    readonly sessionAttributes: outputs.Tacacs.GetProfileSessionAttribute[];
}
/**
 * This data source can read the TACACS Profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.Tacacs.getProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getProfileOutput(args?: GetProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProfileResult> {
    return pulumi.output(args).apply((a: any) => getProfile(a, opts))
}

/**
 * A collection of arguments for invoking getProfile.
 */
export interface GetProfileOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the TACACS profile
     */
    name?: pulumi.Input<string>;
}

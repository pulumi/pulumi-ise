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
 * const example = ise.deviceadmin.getTacacsProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTacacsProfile(args?: GetTacacsProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetTacacsProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:deviceadmin/getTacacsProfile:getTacacsProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTacacsProfile.
 */
export interface GetTacacsProfileArgs {
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
 * A collection of values returned by getTacacsProfile.
 */
export interface GetTacacsProfileResult {
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
    readonly sessionAttributes: outputs.deviceadmin.GetTacacsProfileSessionAttribute[];
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
 * const example = ise.deviceadmin.getTacacsProfile({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTacacsProfileOutput(args?: GetTacacsProfileOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTacacsProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:deviceadmin/getTacacsProfile:getTacacsProfile", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTacacsProfile.
 */
export interface GetTacacsProfileOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the TACACS profile
     */
    name?: pulumi.Input<string>;
}

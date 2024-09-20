// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the TACACS Command Set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.deviceadmin.getTacacsCommandSet({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTacacsCommandSet(args?: GetTacacsCommandSetArgs, opts?: pulumi.InvokeOptions): Promise<GetTacacsCommandSetResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:deviceadmin/getTacacsCommandSet:getTacacsCommandSet", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTacacsCommandSet.
 */
export interface GetTacacsCommandSetArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the TACACS command set
     */
    name?: string;
}

/**
 * A collection of values returned by getTacacsCommandSet.
 */
export interface GetTacacsCommandSetResult {
    readonly commands: outputs.deviceadmin.GetTacacsCommandSetCommand[];
    /**
     * Description
     */
    readonly description: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * The name of the TACACS command set
     */
    readonly name: string;
    /**
     * Permit unmatched commands
     */
    readonly permitUnmatched: boolean;
}
/**
 * This data source can read the TACACS Command Set.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.deviceadmin.getTacacsCommandSet({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTacacsCommandSetOutput(args?: GetTacacsCommandSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTacacsCommandSetResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:deviceadmin/getTacacsCommandSet:getTacacsCommandSet", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTacacsCommandSet.
 */
export interface GetTacacsCommandSetOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the TACACS command set
     */
    name?: pulumi.Input<string>;
}

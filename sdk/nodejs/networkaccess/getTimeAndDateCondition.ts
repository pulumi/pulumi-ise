// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the Network Access Time And Date Condition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getTimeAndDateCondition({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTimeAndDateCondition(args?: GetTimeAndDateConditionArgs, opts?: pulumi.InvokeOptions): Promise<GetTimeAndDateConditionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTimeAndDateCondition.
 */
export interface GetTimeAndDateConditionArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * Condition name
     */
    name?: string;
}

/**
 * A collection of values returned by getTimeAndDateCondition.
 */
export interface GetTimeAndDateConditionResult {
    /**
     * Condition description
     */
    readonly description: string;
    /**
     * End date
     */
    readonly endDate: string;
    /**
     * End time
     */
    readonly endTime: string;
    /**
     * Exception end date
     */
    readonly exceptionEndDate: string;
    /**
     * Exception end time
     */
    readonly exceptionEndTime: string;
    /**
     * Exception start date
     */
    readonly exceptionStartDate: string;
    /**
     * Exception start time
     */
    readonly exceptionStartTime: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Indicates whereas this condition is in negate mode
     */
    readonly isNegate: boolean;
    /**
     * Condition name
     */
    readonly name: string;
    /**
     * Start date
     */
    readonly startDate: string;
    /**
     * Start time
     */
    readonly startTime: string;
    /**
     * Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
     */
    readonly weekDays: string[];
    /**
     * Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
     */
    readonly weekDaysExceptions: string[];
}
/**
 * This data source can read the Network Access Time And Date Condition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getTimeAndDateCondition({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getTimeAndDateConditionOutput(args?: GetTimeAndDateConditionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTimeAndDateConditionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getTimeAndDateCondition.
 */
export interface GetTimeAndDateConditionOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * Condition name
     */
    name?: pulumi.Input<string>;
}

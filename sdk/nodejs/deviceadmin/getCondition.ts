// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Device Admin Condition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.deviceadmin.getCondition({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getCondition(args?: GetConditionArgs, opts?: pulumi.InvokeOptions): Promise<GetConditionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:deviceadmin/getCondition:getCondition", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCondition.
 */
export interface GetConditionArgs {
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
 * A collection of values returned by getCondition.
 */
export interface GetConditionResult {
    /**
     * Dictionary attribute name
     */
    readonly attributeName: string;
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     */
    readonly attributeValue: string;
    /**
     * List of child conditions. `conditionType` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
     */
    readonly childrens: outputs.deviceadmin.GetConditionChildren[];
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
     */
    readonly conditionType: string;
    /**
     * Condition description
     */
    readonly description: string;
    /**
     * Dictionary name
     */
    readonly dictionaryName: string;
    /**
     * Dictionary value
     */
    readonly dictionaryValue: string;
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
     * Equality operator
     */
    readonly operator: string;
}
/**
 * This data source can read the Device Admin Condition.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.deviceadmin.getCondition({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getConditionOutput(args?: GetConditionOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetConditionResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:deviceadmin/getCondition:getCondition", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getCondition.
 */
export interface GetConditionOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * Condition name
     */
    name?: pulumi.Input<string>;
}

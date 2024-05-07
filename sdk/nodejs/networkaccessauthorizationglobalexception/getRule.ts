// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Network Access Authorization Global Exception Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.NetworkAccessAuthorizationGlobalException.getRule({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getRule(args?: GetRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetRuleResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:NetworkAccessAuthorizationGlobalException/getRule:getRule", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRule.
 */
export interface GetRuleArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: string;
}

/**
 * A collection of values returned by getRule.
 */
export interface GetRuleResult {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    readonly childrens: outputs.NetworkAccessAuthorizationGlobalException.GetRuleChildren[];
    /**
     * Dictionary attribute name
     */
    readonly conditionAttributeName: string;
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     */
    readonly conditionAttributeValue: string;
    /**
     * Dictionary name
     */
    readonly conditionDictionaryName: string;
    /**
     * Dictionary value
     */
    readonly conditionDictionaryValue: string;
    /**
     * UUID for condition
     */
    readonly conditionId: string;
    /**
     * Indicates whereas this condition is in negate mode
     */
    readonly conditionIsNegate: boolean;
    /**
     * Equality operator
     */
    readonly conditionOperator: string;
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
     */
    readonly conditionType: string;
    /**
     * Indicates if this rule is the default one
     */
    readonly default: boolean;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    readonly name: string;
    /**
     * The authorization profile(s)
     */
    readonly profiles: string[];
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    readonly rank: number;
    /**
     * Security group used in authorization policies
     */
    readonly securityGroup: string;
    /**
     * The state that the rule is in. A disabled rule cannot be matched.
     */
    readonly state: string;
}
/**
 * This data source can read the Network Access Authorization Global Exception Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.NetworkAccessAuthorizationGlobalException.getRule({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getRuleOutput(args?: GetRuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuleResult> {
    return pulumi.output(args).apply((a: any) => getRule(a, opts))
}

/**
 * A collection of arguments for invoking getRule.
 */
export interface GetRuleOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
}

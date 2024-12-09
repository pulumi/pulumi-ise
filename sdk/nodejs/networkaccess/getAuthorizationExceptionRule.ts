// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Network Access Authorization Exception Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getAuthorizationExceptionRule({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 *     policySetId: "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
 * });
 * ```
 */
export function getAuthorizationExceptionRule(args: GetAuthorizationExceptionRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthorizationExceptionRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:networkaccess/getAuthorizationExceptionRule:getAuthorizationExceptionRule", {
        "id": args.id,
        "name": args.name,
        "policySetId": args.policySetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthorizationExceptionRule.
 */
export interface GetAuthorizationExceptionRuleArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: string;
    /**
     * Policy set ID
     */
    policySetId: string;
}

/**
 * A collection of values returned by getAuthorizationExceptionRule.
 */
export interface GetAuthorizationExceptionRuleResult {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    readonly childrens: outputs.networkaccess.GetAuthorizationExceptionRuleChildren[];
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
     * Policy set ID
     */
    readonly policySetId: string;
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
 * This data source can read the Network Access Authorization Exception Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getAuthorizationExceptionRule({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 *     policySetId: "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
 * });
 * ```
 */
export function getAuthorizationExceptionRuleOutput(args: GetAuthorizationExceptionRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetAuthorizationExceptionRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:networkaccess/getAuthorizationExceptionRule:getAuthorizationExceptionRule", {
        "id": args.id,
        "name": args.name,
        "policySetId": args.policySetId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthorizationExceptionRule.
 */
export interface GetAuthorizationExceptionRuleOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
    /**
     * Policy set ID
     */
    policySetId: pulumi.Input<string>;
}

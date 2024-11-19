// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource can manage a Network Access Authorization Global Exception Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.networkaccess.AuthorizationGlobalExceptionRule("example", {
 *     name: "Rule1",
 *     rank: 0,
 *     state: "enabled",
 *     conditionType: "ConditionAttributes",
 *     conditionIsNegate: false,
 *     conditionAttributeName: "Location",
 *     conditionAttributeValue: "All Locations",
 *     conditionDictionaryName: "DEVICE",
 *     conditionOperator: "equals",
 *     profiles: ["PermitAccess"],
 *     securityGroup: "BYOD",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:networkaccess/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class AuthorizationGlobalExceptionRule extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationGlobalExceptionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationGlobalExceptionRuleState, opts?: pulumi.CustomResourceOptions): AuthorizationGlobalExceptionRule {
        return new AuthorizationGlobalExceptionRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:networkaccess/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule';

    /**
     * Returns true if the given object is an instance of AuthorizationGlobalExceptionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationGlobalExceptionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationGlobalExceptionRule.__pulumiType;
    }

    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    public readonly childrens!: pulumi.Output<outputs.networkaccess.AuthorizationGlobalExceptionRuleChildren[] | undefined>;
    /**
     * Dictionary attribute name
     */
    public readonly conditionAttributeName!: pulumi.Output<string | undefined>;
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     */
    public readonly conditionAttributeValue!: pulumi.Output<string | undefined>;
    /**
     * Dictionary name
     */
    public readonly conditionDictionaryName!: pulumi.Output<string | undefined>;
    /**
     * Dictionary value
     */
    public readonly conditionDictionaryValue!: pulumi.Output<string | undefined>;
    /**
     * UUID for condition
     */
    public readonly conditionId!: pulumi.Output<string | undefined>;
    /**
     * Indicates whereas this condition is in negate mode
     */
    public readonly conditionIsNegate!: pulumi.Output<boolean | undefined>;
    /**
     * Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     */
    public readonly conditionOperator!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     */
    public readonly conditionType!: pulumi.Output<string | undefined>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The authorization profile(s)
     */
    public readonly profiles!: pulumi.Output<string[] | undefined>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    public readonly rank!: pulumi.Output<number | undefined>;
    /**
     * Security group used in authorization policies
     */
    public readonly securityGroup!: pulumi.Output<string | undefined>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    public readonly state!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthorizationGlobalExceptionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: AuthorizationGlobalExceptionRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationGlobalExceptionRuleArgs | AuthorizationGlobalExceptionRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorizationGlobalExceptionRuleState | undefined;
            resourceInputs["childrens"] = state ? state.childrens : undefined;
            resourceInputs["conditionAttributeName"] = state ? state.conditionAttributeName : undefined;
            resourceInputs["conditionAttributeValue"] = state ? state.conditionAttributeValue : undefined;
            resourceInputs["conditionDictionaryName"] = state ? state.conditionDictionaryName : undefined;
            resourceInputs["conditionDictionaryValue"] = state ? state.conditionDictionaryValue : undefined;
            resourceInputs["conditionId"] = state ? state.conditionId : undefined;
            resourceInputs["conditionIsNegate"] = state ? state.conditionIsNegate : undefined;
            resourceInputs["conditionOperator"] = state ? state.conditionOperator : undefined;
            resourceInputs["conditionType"] = state ? state.conditionType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["profiles"] = state ? state.profiles : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["securityGroup"] = state ? state.securityGroup : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as AuthorizationGlobalExceptionRuleArgs | undefined;
            resourceInputs["childrens"] = args ? args.childrens : undefined;
            resourceInputs["conditionAttributeName"] = args ? args.conditionAttributeName : undefined;
            resourceInputs["conditionAttributeValue"] = args ? args.conditionAttributeValue : undefined;
            resourceInputs["conditionDictionaryName"] = args ? args.conditionDictionaryName : undefined;
            resourceInputs["conditionDictionaryValue"] = args ? args.conditionDictionaryValue : undefined;
            resourceInputs["conditionId"] = args ? args.conditionId : undefined;
            resourceInputs["conditionIsNegate"] = args ? args.conditionIsNegate : undefined;
            resourceInputs["conditionOperator"] = args ? args.conditionOperator : undefined;
            resourceInputs["conditionType"] = args ? args.conditionType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["profiles"] = args ? args.profiles : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["securityGroup"] = args ? args.securityGroup : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthorizationGlobalExceptionRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorizationGlobalExceptionRule resources.
 */
export interface AuthorizationGlobalExceptionRuleState {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.networkaccess.AuthorizationGlobalExceptionRuleChildren>[]>;
    /**
     * Dictionary attribute name
     */
    conditionAttributeName?: pulumi.Input<string>;
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     */
    conditionAttributeValue?: pulumi.Input<string>;
    /**
     * Dictionary name
     */
    conditionDictionaryName?: pulumi.Input<string>;
    /**
     * Dictionary value
     */
    conditionDictionaryValue?: pulumi.Input<string>;
    /**
     * UUID for condition
     */
    conditionId?: pulumi.Input<string>;
    /**
     * Indicates whereas this condition is in negate mode
     */
    conditionIsNegate?: pulumi.Input<boolean>;
    /**
     * Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     */
    conditionOperator?: pulumi.Input<string>;
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     */
    conditionType?: pulumi.Input<string>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
    /**
     * The authorization profile(s)
     */
    profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank?: pulumi.Input<number>;
    /**
     * Security group used in authorization policies
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthorizationGlobalExceptionRule resource.
 */
export interface AuthorizationGlobalExceptionRuleArgs {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.networkaccess.AuthorizationGlobalExceptionRuleChildren>[]>;
    /**
     * Dictionary attribute name
     */
    conditionAttributeName?: pulumi.Input<string>;
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     */
    conditionAttributeValue?: pulumi.Input<string>;
    /**
     * Dictionary name
     */
    conditionDictionaryName?: pulumi.Input<string>;
    /**
     * Dictionary value
     */
    conditionDictionaryValue?: pulumi.Input<string>;
    /**
     * UUID for condition
     */
    conditionId?: pulumi.Input<string>;
    /**
     * Indicates whereas this condition is in negate mode
     */
    conditionIsNegate?: pulumi.Input<boolean>;
    /**
     * Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     */
    conditionOperator?: pulumi.Input<string>;
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     */
    conditionType?: pulumi.Input<string>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
    /**
     * The authorization profile(s)
     */
    profiles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank?: pulumi.Input<number>;
    /**
     * Security group used in authorization policies
     */
    securityGroup?: pulumi.Input<string>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    state?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource can manage a Device Admin Authorization Rule.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.deviceadmin.AuthorizationRule("example", {
 *     policySetId: "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
 *     name: "Rule1",
 *     "default": false,
 *     rank: 0,
 *     state: "enabled",
 *     conditionType: "ConditionAttributes",
 *     conditionIsNegate: false,
 *     conditionAttributeName: "Location",
 *     conditionAttributeValue: "All Locations",
 *     conditionDictionaryName: "DEVICE",
 *     conditionOperator: "equals",
 *     commandSets: ["DenyAllCommands"],
 *     profile: "Default Shell Profile",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:deviceadmin/authorizationRule:AuthorizationRule example "76d24097-41c4-4558-a4d0-a8c07ac08470,76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class AuthorizationRule extends pulumi.CustomResource {
    /**
     * Get an existing AuthorizationRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthorizationRuleState, opts?: pulumi.CustomResourceOptions): AuthorizationRule {
        return new AuthorizationRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:deviceadmin/authorizationRule:AuthorizationRule';

    /**
     * Returns true if the given object is an instance of AuthorizationRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthorizationRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthorizationRule.__pulumiType;
    }

    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    public readonly childrens!: pulumi.Output<outputs.deviceadmin.AuthorizationRuleChildren[] | undefined>;
    /**
     * Command sets enforce the specified list of commands that can be executed by a device administrator
     */
    public readonly commandSets!: pulumi.Output<string[] | undefined>;
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
     * Indicates if this rule is the default one
     */
    public readonly default!: pulumi.Output<boolean | undefined>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Policy set ID
     */
    public readonly policySetId!: pulumi.Output<string>;
    /**
     * Device admin profiles control the initial login session of the device administrator
     */
    public readonly profile!: pulumi.Output<string | undefined>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    public readonly rank!: pulumi.Output<number | undefined>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    public readonly state!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthorizationRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthorizationRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthorizationRuleArgs | AuthorizationRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthorizationRuleState | undefined;
            resourceInputs["childrens"] = state ? state.childrens : undefined;
            resourceInputs["commandSets"] = state ? state.commandSets : undefined;
            resourceInputs["conditionAttributeName"] = state ? state.conditionAttributeName : undefined;
            resourceInputs["conditionAttributeValue"] = state ? state.conditionAttributeValue : undefined;
            resourceInputs["conditionDictionaryName"] = state ? state.conditionDictionaryName : undefined;
            resourceInputs["conditionDictionaryValue"] = state ? state.conditionDictionaryValue : undefined;
            resourceInputs["conditionId"] = state ? state.conditionId : undefined;
            resourceInputs["conditionIsNegate"] = state ? state.conditionIsNegate : undefined;
            resourceInputs["conditionOperator"] = state ? state.conditionOperator : undefined;
            resourceInputs["conditionType"] = state ? state.conditionType : undefined;
            resourceInputs["default"] = state ? state.default : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["profile"] = state ? state.profile : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
        } else {
            const args = argsOrState as AuthorizationRuleArgs | undefined;
            if ((!args || args.policySetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policySetId'");
            }
            resourceInputs["childrens"] = args ? args.childrens : undefined;
            resourceInputs["commandSets"] = args ? args.commandSets : undefined;
            resourceInputs["conditionAttributeName"] = args ? args.conditionAttributeName : undefined;
            resourceInputs["conditionAttributeValue"] = args ? args.conditionAttributeValue : undefined;
            resourceInputs["conditionDictionaryName"] = args ? args.conditionDictionaryName : undefined;
            resourceInputs["conditionDictionaryValue"] = args ? args.conditionDictionaryValue : undefined;
            resourceInputs["conditionId"] = args ? args.conditionId : undefined;
            resourceInputs["conditionIsNegate"] = args ? args.conditionIsNegate : undefined;
            resourceInputs["conditionOperator"] = args ? args.conditionOperator : undefined;
            resourceInputs["conditionType"] = args ? args.conditionType : undefined;
            resourceInputs["default"] = args ? args.default : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["policySetId"] = args ? args.policySetId : undefined;
            resourceInputs["profile"] = args ? args.profile : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthorizationRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthorizationRule resources.
 */
export interface AuthorizationRuleState {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.deviceadmin.AuthorizationRuleChildren>[]>;
    /**
     * Command sets enforce the specified list of commands that can be executed by a device administrator
     */
    commandSets?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Indicates if this rule is the default one
     */
    default?: pulumi.Input<boolean>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
    /**
     * Policy set ID
     */
    policySetId?: pulumi.Input<string>;
    /**
     * Device admin profiles control the initial login session of the device administrator
     */
    profile?: pulumi.Input<string>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank?: pulumi.Input<number>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    state?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthorizationRule resource.
 */
export interface AuthorizationRuleArgs {
    /**
     * List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.deviceadmin.AuthorizationRuleChildren>[]>;
    /**
     * Command sets enforce the specified list of commands that can be executed by a device administrator
     */
    commandSets?: pulumi.Input<pulumi.Input<string>[]>;
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
     * Indicates if this rule is the default one
     */
    default?: pulumi.Input<boolean>;
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     */
    name?: pulumi.Input<string>;
    /**
     * Policy set ID
     */
    policySetId: pulumi.Input<string>;
    /**
     * Device admin profiles control the initial login session of the device administrator
     */
    profile?: pulumi.Input<string>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank?: pulumi.Input<number>;
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     */
    state?: pulumi.Input<string>;
}

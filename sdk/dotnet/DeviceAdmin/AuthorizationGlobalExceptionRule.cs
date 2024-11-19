// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.DeviceAdmin
{
    /// <summary>
    /// This resource can manage a Device Admin Authorization Global Exception Rule.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ise = Pulumi.Ise;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Ise.DeviceAdmin.AuthorizationGlobalExceptionRule("example", new()
    ///     {
    ///         Name = "Rule1",
    ///         Rank = 0,
    ///         State = "enabled",
    ///         ConditionType = "ConditionAttributes",
    ///         ConditionIsNegate = false,
    ///         ConditionAttributeName = "Location",
    ///         ConditionAttributeValue = "All Locations",
    ///         ConditionDictionaryName = "DEVICE",
    ///         ConditionOperator = "equals",
    ///         CommandSets = new[]
    ///         {
    ///             "DenyAllCommands",
    ///         },
    ///         Profile = "Default Shell Profile",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:deviceadmin/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:deviceadmin/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule")]
    public partial class AuthorizationGlobalExceptionRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        [Output("childrens")]
        public Output<ImmutableArray<Outputs.AuthorizationGlobalExceptionRuleChildren>> Childrens { get; private set; } = null!;

        /// <summary>
        /// Command sets enforce the specified list of commands that can be executed by a device administrator
        /// </summary>
        [Output("commandSets")]
        public Output<ImmutableArray<string>> CommandSets { get; private set; } = null!;

        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        [Output("conditionAttributeName")]
        public Output<string?> ConditionAttributeName { get; private set; } = null!;

        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        [Output("conditionAttributeValue")]
        public Output<string?> ConditionAttributeValue { get; private set; } = null!;

        /// <summary>
        /// Dictionary name
        /// </summary>
        [Output("conditionDictionaryName")]
        public Output<string?> ConditionDictionaryName { get; private set; } = null!;

        /// <summary>
        /// Dictionary value
        /// </summary>
        [Output("conditionDictionaryValue")]
        public Output<string?> ConditionDictionaryValue { get; private set; } = null!;

        /// <summary>
        /// UUID for condition
        /// </summary>
        [Output("conditionId")]
        public Output<string?> ConditionId { get; private set; } = null!;

        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Output("conditionIsNegate")]
        public Output<bool?> ConditionIsNegate { get; private set; } = null!;

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Output("conditionOperator")]
        public Output<string?> ConditionOperator { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
        /// `ConditionOrBlock`, `ConditionReference`
        /// </summary>
        [Output("conditionType")]
        public Output<string?> ConditionType { get; private set; } = null!;

        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Device admin profiles control the initial login session of the device administrator
        /// </summary>
        [Output("profile")]
        public Output<string?> Profile { get; private set; } = null!;

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Output("rank")]
        public Output<int?> Rank { get; private set; } = null!;

        /// <summary>
        /// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;


        /// <summary>
        /// Create a AuthorizationGlobalExceptionRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthorizationGlobalExceptionRule(string name, AuthorizationGlobalExceptionRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("ise:deviceadmin/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule", name, args ?? new AuthorizationGlobalExceptionRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthorizationGlobalExceptionRule(string name, Input<string> id, AuthorizationGlobalExceptionRuleState? state = null, CustomResourceOptions? options = null)
            : base("ise:deviceadmin/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AuthorizationGlobalExceptionRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthorizationGlobalExceptionRule Get(string name, Input<string> id, AuthorizationGlobalExceptionRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthorizationGlobalExceptionRule(name, id, state, options);
        }
    }

    public sealed class AuthorizationGlobalExceptionRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("childrens")]
        private InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenArgs>? _childrens;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenArgs> Childrens
        {
            get => _childrens ?? (_childrens = new InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenArgs>());
            set => _childrens = value;
        }

        [Input("commandSets")]
        private InputList<string>? _commandSets;

        /// <summary>
        /// Command sets enforce the specified list of commands that can be executed by a device administrator
        /// </summary>
        public InputList<string> CommandSets
        {
            get => _commandSets ?? (_commandSets = new InputList<string>());
            set => _commandSets = value;
        }

        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        [Input("conditionAttributeName")]
        public Input<string>? ConditionAttributeName { get; set; }

        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        [Input("conditionAttributeValue")]
        public Input<string>? ConditionAttributeValue { get; set; }

        /// <summary>
        /// Dictionary name
        /// </summary>
        [Input("conditionDictionaryName")]
        public Input<string>? ConditionDictionaryName { get; set; }

        /// <summary>
        /// Dictionary value
        /// </summary>
        [Input("conditionDictionaryValue")]
        public Input<string>? ConditionDictionaryValue { get; set; }

        /// <summary>
        /// UUID for condition
        /// </summary>
        [Input("conditionId")]
        public Input<string>? ConditionId { get; set; }

        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Input("conditionIsNegate")]
        public Input<bool>? ConditionIsNegate { get; set; }

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Input("conditionOperator")]
        public Input<string>? ConditionOperator { get; set; }

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
        /// `ConditionOrBlock`, `ConditionReference`
        /// </summary>
        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Device admin profiles control the initial login session of the device administrator
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        /// <summary>
        /// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public AuthorizationGlobalExceptionRuleArgs()
        {
        }
        public static new AuthorizationGlobalExceptionRuleArgs Empty => new AuthorizationGlobalExceptionRuleArgs();
    }

    public sealed class AuthorizationGlobalExceptionRuleState : global::Pulumi.ResourceArgs
    {
        [Input("childrens")]
        private InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenGetArgs>? _childrens;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenGetArgs> Childrens
        {
            get => _childrens ?? (_childrens = new InputList<Inputs.AuthorizationGlobalExceptionRuleChildrenGetArgs>());
            set => _childrens = value;
        }

        [Input("commandSets")]
        private InputList<string>? _commandSets;

        /// <summary>
        /// Command sets enforce the specified list of commands that can be executed by a device administrator
        /// </summary>
        public InputList<string> CommandSets
        {
            get => _commandSets ?? (_commandSets = new InputList<string>());
            set => _commandSets = value;
        }

        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        [Input("conditionAttributeName")]
        public Input<string>? ConditionAttributeName { get; set; }

        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        [Input("conditionAttributeValue")]
        public Input<string>? ConditionAttributeValue { get; set; }

        /// <summary>
        /// Dictionary name
        /// </summary>
        [Input("conditionDictionaryName")]
        public Input<string>? ConditionDictionaryName { get; set; }

        /// <summary>
        /// Dictionary value
        /// </summary>
        [Input("conditionDictionaryValue")]
        public Input<string>? ConditionDictionaryValue { get; set; }

        /// <summary>
        /// UUID for condition
        /// </summary>
        [Input("conditionId")]
        public Input<string>? ConditionId { get; set; }

        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Input("conditionIsNegate")]
        public Input<bool>? ConditionIsNegate { get; set; }

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Input("conditionOperator")]
        public Input<string>? ConditionOperator { get; set; }

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
        /// `ConditionOrBlock`, `ConditionReference`
        /// </summary>
        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Device admin profiles control the initial login session of the device administrator
        /// </summary>
        [Input("profile")]
        public Input<string>? Profile { get; set; }

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        /// <summary>
        /// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public AuthorizationGlobalExceptionRuleState()
        {
        }
        public static new AuthorizationGlobalExceptionRuleState Empty => new AuthorizationGlobalExceptionRuleState();
    }
}

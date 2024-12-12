// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccess
{
    public static class GetAuthorizationRule
    {
        /// <summary>
        /// This data source can read the Network Access Authorization Rule.
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
        ///     var example = Ise.NetworkAccess.GetAuthorizationRule.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAuthorizationRuleResult> InvokeAsync(GetAuthorizationRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationRuleResult>("ise:networkaccess/getAuthorizationRule:getAuthorizationRule", args ?? new GetAuthorizationRuleArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Network Access Authorization Rule.
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
        ///     var example = Ise.NetworkAccess.GetAuthorizationRule.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthorizationRuleResult> Invoke(GetAuthorizationRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationRuleResult>("ise:networkaccess/getAuthorizationRule:getAuthorizationRule", args ?? new GetAuthorizationRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationRuleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Policy set ID
        /// </summary>
        [Input("policySetId", required: true)]
        public string PolicySetId { get; set; } = null!;

        public GetAuthorizationRuleArgs()
        {
        }
        public static new GetAuthorizationRuleArgs Empty => new GetAuthorizationRuleArgs();
    }

    public sealed class GetAuthorizationRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy set ID
        /// </summary>
        [Input("policySetId", required: true)]
        public Input<string> PolicySetId { get; set; } = null!;

        public GetAuthorizationRuleInvokeArgs()
        {
        }
        public static new GetAuthorizationRuleInvokeArgs Empty => new GetAuthorizationRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizationRuleResult
    {
        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAuthorizationRuleChildrenResult> Childrens;
        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        public readonly string ConditionAttributeName;
        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        public readonly string ConditionAttributeValue;
        /// <summary>
        /// Dictionary name
        /// </summary>
        public readonly string ConditionDictionaryName;
        /// <summary>
        /// Dictionary value
        /// </summary>
        public readonly string ConditionDictionaryValue;
        /// <summary>
        /// UUID for condition
        /// </summary>
        public readonly string ConditionId;
        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        public readonly bool ConditionIsNegate;
        /// <summary>
        /// Equality operator
        /// </summary>
        public readonly string ConditionOperator;
        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        /// </summary>
        public readonly string ConditionType;
        /// <summary>
        /// Indicates if this rule is the default one
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Policy set ID
        /// </summary>
        public readonly string PolicySetId;
        /// <summary>
        /// The authorization profile(s)
        /// </summary>
        public readonly ImmutableArray<string> Profiles;
        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        public readonly int Rank;
        /// <summary>
        /// Security group used in authorization policies
        /// </summary>
        public readonly string SecurityGroup;
        /// <summary>
        /// The state that the rule is in. A disabled rule cannot be matched.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetAuthorizationRuleResult(
            ImmutableArray<Outputs.GetAuthorizationRuleChildrenResult> childrens,

            string conditionAttributeName,

            string conditionAttributeValue,

            string conditionDictionaryName,

            string conditionDictionaryValue,

            string conditionId,

            bool conditionIsNegate,

            string conditionOperator,

            string conditionType,

            bool @default,

            string id,

            string name,

            string policySetId,

            ImmutableArray<string> profiles,

            int rank,

            string securityGroup,

            string state)
        {
            Childrens = childrens;
            ConditionAttributeName = conditionAttributeName;
            ConditionAttributeValue = conditionAttributeValue;
            ConditionDictionaryName = conditionDictionaryName;
            ConditionDictionaryValue = conditionDictionaryValue;
            ConditionId = conditionId;
            ConditionIsNegate = conditionIsNegate;
            ConditionOperator = conditionOperator;
            ConditionType = conditionType;
            Default = @default;
            Id = id;
            Name = name;
            PolicySetId = policySetId;
            Profiles = profiles;
            Rank = rank;
            SecurityGroup = securityGroup;
            State = state;
        }
    }
}
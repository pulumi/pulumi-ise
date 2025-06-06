// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccess
{
    public static class GetAuthenticationRule
    {
        /// <summary>
        /// This data source can read the Network Access Authentication Rule.
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
        ///     var example = Ise.NetworkAccess.GetAuthenticationRule.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAuthenticationRuleResult> InvokeAsync(GetAuthenticationRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthenticationRuleResult>("ise:networkaccess/getAuthenticationRule:getAuthenticationRule", args ?? new GetAuthenticationRuleArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Network Access Authentication Rule.
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
        ///     var example = Ise.NetworkAccess.GetAuthenticationRule.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthenticationRuleResult> Invoke(GetAuthenticationRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthenticationRuleResult>("ise:networkaccess/getAuthenticationRule:getAuthenticationRule", args ?? new GetAuthenticationRuleInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Network Access Authentication Rule.
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
        ///     var example = Ise.NetworkAccess.GetAuthenticationRule.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthenticationRuleResult> Invoke(GetAuthenticationRuleInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthenticationRuleResult>("ise:networkaccess/getAuthenticationRule:getAuthenticationRule", args ?? new GetAuthenticationRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthenticationRuleArgs : global::Pulumi.InvokeArgs
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

        public GetAuthenticationRuleArgs()
        {
        }
        public static new GetAuthenticationRuleArgs Empty => new GetAuthenticationRuleArgs();
    }

    public sealed class GetAuthenticationRuleInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetAuthenticationRuleInvokeArgs()
        {
        }
        public static new GetAuthenticationRuleInvokeArgs Empty => new GetAuthenticationRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthenticationRuleResult
    {
        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAuthenticationRuleChildrenResult> Childrens;
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
        /// Identity source name from the identity stores
        /// </summary>
        public readonly string IdentitySourceName;
        /// <summary>
        /// Action to perform when authentication fails such as Bad credentials, disabled user and so on
        /// </summary>
        public readonly string IfAuthFail;
        /// <summary>
        /// Action to perform when ISE is uanble to access the identity database
        /// </summary>
        public readonly string IfProcessFail;
        /// <summary>
        /// Action to perform when user is not found in any of identity stores
        /// </summary>
        public readonly string IfUserNotFound;
        /// <summary>
        /// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Policy set ID
        /// </summary>
        public readonly string PolicySetId;
        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        public readonly int Rank;
        /// <summary>
        /// The state that the rule is in. A disabled rule cannot be matched.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetAuthenticationRuleResult(
            ImmutableArray<Outputs.GetAuthenticationRuleChildrenResult> childrens,

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

            string identitySourceName,

            string ifAuthFail,

            string ifProcessFail,

            string ifUserNotFound,

            string name,

            string policySetId,

            int rank,

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
            IdentitySourceName = identitySourceName;
            IfAuthFail = ifAuthFail;
            IfProcessFail = ifProcessFail;
            IfUserNotFound = ifUserNotFound;
            Name = name;
            PolicySetId = policySetId;
            Rank = rank;
            State = state;
        }
    }
}

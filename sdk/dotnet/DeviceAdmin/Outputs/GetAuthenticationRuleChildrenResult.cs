// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.DeviceAdmin.Outputs
{

    [OutputType]
    public sealed class GetAuthenticationRuleChildrenResult
    {
        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        public readonly string AttributeName;
        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        public readonly string AttributeValue;
        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAuthenticationRuleChildrenChildrenResult> Childrens;
        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        /// </summary>
        public readonly string ConditionType;
        /// <summary>
        /// Dictionary name
        /// </summary>
        public readonly string DictionaryName;
        /// <summary>
        /// Dictionary value
        /// </summary>
        public readonly string DictionaryValue;
        /// <summary>
        /// UUID for condition
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        public readonly bool IsNegate;
        /// <summary>
        /// Equality operator
        /// </summary>
        public readonly string Operator;

        [OutputConstructor]
        private GetAuthenticationRuleChildrenResult(
            string attributeName,

            string attributeValue,

            ImmutableArray<Outputs.GetAuthenticationRuleChildrenChildrenResult> childrens,

            string conditionType,

            string dictionaryName,

            string dictionaryValue,

            string id,

            bool isNegate,

            string @operator)
        {
            AttributeName = attributeName;
            AttributeValue = attributeValue;
            Childrens = childrens;
            ConditionType = conditionType;
            DictionaryName = dictionaryName;
            DictionaryValue = dictionaryValue;
            Id = id;
            IsNegate = isNegate;
            Operator = @operator;
        }
    }
}
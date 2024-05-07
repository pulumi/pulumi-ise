// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccessPolicy.Outputs
{

    [OutputType]
    public sealed class SetChildrenChildren
    {
        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        public readonly string? AttributeName;
        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        public readonly string? AttributeValue;
        /// <summary>
        /// Condition type.
        ///   - Choices: `ConditionAttributes`, `ConditionReference`
        /// </summary>
        public readonly string ConditionType;
        /// <summary>
        /// Dictionary name
        /// </summary>
        public readonly string? DictionaryName;
        /// <summary>
        /// Dictionary value
        /// </summary>
        public readonly string? DictionaryValue;
        /// <summary>
        /// UUID for condition
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        public readonly bool? IsNegate;
        /// <summary>
        /// Equality operator
        ///   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        public readonly string? Operator;

        [OutputConstructor]
        private SetChildrenChildren(
            string? attributeName,

            string? attributeValue,

            string conditionType,

            string? dictionaryName,

            string? dictionaryValue,

            string? id,

            bool? isNegate,

            string? @operator)
        {
            AttributeName = attributeName;
            AttributeValue = attributeValue;
            ConditionType = conditionType;
            DictionaryName = dictionaryName;
            DictionaryValue = dictionaryValue;
            Id = id;
            IsNegate = isNegate;
            Operator = @operator;
        }
    }
}

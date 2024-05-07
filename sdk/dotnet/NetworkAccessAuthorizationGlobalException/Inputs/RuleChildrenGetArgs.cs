// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccessAuthorizationGlobalException.Inputs
{

    public sealed class RuleChildrenGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        [Input("attributeValue")]
        public Input<string>? AttributeValue { get; set; }

        [Input("childrens")]
        private InputList<Inputs.RuleChildrenChildrenGetArgs>? _childrens;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        /// </summary>
        public InputList<Inputs.RuleChildrenChildrenGetArgs> Childrens
        {
            get => _childrens ?? (_childrens = new InputList<Inputs.RuleChildrenChildrenGetArgs>());
            set => _childrens = value;
        }

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        ///   - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
        /// </summary>
        [Input("conditionType", required: true)]
        public Input<string> ConditionType { get; set; } = null!;

        /// <summary>
        /// Dictionary name
        /// </summary>
        [Input("dictionaryName")]
        public Input<string>? DictionaryName { get; set; }

        /// <summary>
        /// Dictionary value
        /// </summary>
        [Input("dictionaryValue")]
        public Input<string>? DictionaryValue { get; set; }

        /// <summary>
        /// UUID for condition
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Input("isNegate")]
        public Input<bool>? IsNegate { get; set; }

        /// <summary>
        /// Equality operator
        ///   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public RuleChildrenGetArgs()
        {
        }
        public static new RuleChildrenGetArgs Empty => new RuleChildrenGetArgs();
    }
}

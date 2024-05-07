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
    /// This resource can manage a Device Admin Condition.
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
    ///     var example = new Ise.DeviceAdmin.Condition("example", new()
    ///     {
    ///         Name = "Cond1",
    ///         Description = "My description",
    ///         ConditionType = "LibraryConditionAttributes",
    ///         IsNegate = false,
    ///         AttributeName = "User",
    ///         AttributeValue = "User1",
    ///         DictionaryName = "TACACS",
    ///         Operator = "equals",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:DeviceAdmin/condition:Condition example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:DeviceAdmin/condition:Condition")]
    public partial class Condition : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dictionary attribute name
        /// </summary>
        [Output("attributeName")]
        public Output<string?> AttributeName { get; private set; } = null!;

        /// <summary>
        /// Attribute value for condition. Value type is specified in dictionary object.
        /// </summary>
        [Output("attributeValue")]
        public Output<string?> AttributeValue { get; private set; } = null!;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        /// </summary>
        [Output("childrens")]
        public Output<ImmutableArray<Outputs.ConditionChildren>> Childrens { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        /// `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        /// </summary>
        [Output("conditionType")]
        public Output<string> ConditionType { get; private set; } = null!;

        /// <summary>
        /// Condition description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Dictionary name
        /// </summary>
        [Output("dictionaryName")]
        public Output<string?> DictionaryName { get; private set; } = null!;

        /// <summary>
        /// Dictionary value
        /// </summary>
        [Output("dictionaryValue")]
        public Output<string?> DictionaryValue { get; private set; } = null!;

        /// <summary>
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Output("isNegate")]
        public Output<bool?> IsNegate { get; private set; } = null!;

        /// <summary>
        /// Condition name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Output("operator")]
        public Output<string?> Operator { get; private set; } = null!;


        /// <summary>
        /// Create a Condition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Condition(string name, ConditionArgs args, CustomResourceOptions? options = null)
            : base("ise:DeviceAdmin/condition:Condition", name, args ?? new ConditionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Condition(string name, Input<string> id, ConditionState? state = null, CustomResourceOptions? options = null)
            : base("ise:DeviceAdmin/condition:Condition", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Condition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Condition Get(string name, Input<string> id, ConditionState? state = null, CustomResourceOptions? options = null)
        {
            return new Condition(name, id, state, options);
        }
    }

    public sealed class ConditionArgs : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.ConditionChildrenArgs>? _childrens;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        /// </summary>
        public InputList<Inputs.ConditionChildrenArgs> Childrens
        {
            get => _childrens ?? (_childrens = new InputList<Inputs.ConditionChildrenArgs>());
            set => _childrens = value;
        }

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        /// `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        /// </summary>
        [Input("conditionType", required: true)]
        public Input<string> ConditionType { get; set; } = null!;

        /// <summary>
        /// Condition description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Input("isNegate")]
        public Input<bool>? IsNegate { get; set; }

        /// <summary>
        /// Condition name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public ConditionArgs()
        {
        }
        public static new ConditionArgs Empty => new ConditionArgs();
    }

    public sealed class ConditionState : global::Pulumi.ResourceArgs
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
        private InputList<Inputs.ConditionChildrenGetArgs>? _childrens;

        /// <summary>
        /// List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        /// </summary>
        public InputList<Inputs.ConditionChildrenGetArgs> Childrens
        {
            get => _childrens ?? (_childrens = new InputList<Inputs.ConditionChildrenGetArgs>());
            set => _childrens = value;
        }

        /// <summary>
        /// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        /// additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        /// `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        /// </summary>
        [Input("conditionType")]
        public Input<string>? ConditionType { get; set; }

        /// <summary>
        /// Condition description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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
        /// Indicates whereas this condition is in negate mode
        /// </summary>
        [Input("isNegate")]
        public Input<bool>? IsNegate { get; set; }

        /// <summary>
        /// Condition name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        /// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        /// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        public ConditionState()
        {
        }
        public static new ConditionState Empty => new ConditionState();
    }
}

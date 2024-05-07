// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.Authorization.Inputs
{

    public sealed class ProfileAdvancedAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dictionary name
        /// </summary>
        [Input("attributeLeftDictionaryName")]
        public Input<string>? AttributeLeftDictionaryName { get; set; }

        /// <summary>
        /// Attribute name
        /// </summary>
        [Input("attributeLeftName")]
        public Input<string>? AttributeLeftName { get; set; }

        /// <summary>
        /// Dictionary name, only required when `attribute_right_value_type` is `AdvancedDictionaryAttribute`
        /// </summary>
        [Input("attributeRightDictionaryName")]
        public Input<string>? AttributeRightDictionaryName { get; set; }

        /// <summary>
        /// Attribute name, only required when `attribute_right_value_type` is `AdvancedDictionaryAttribute`
        /// </summary>
        [Input("attributeRightName")]
        public Input<string>? AttributeRightName { get; set; }

        /// <summary>
        /// Attribute value, only required when `attribute_right_value_type` is `AttributeValue`
        /// </summary>
        [Input("attributeRightValue")]
        public Input<string>? AttributeRightValue { get; set; }

        /// <summary>
        /// Advanced attribute value type
        ///   - Choices: `AdvancedDictionaryAttribute`, `AttributeValue`
        /// </summary>
        [Input("attributeRightValueType")]
        public Input<string>? AttributeRightValueType { get; set; }

        public ProfileAdvancedAttributeArgs()
        {
        }
        public static new ProfileAdvancedAttributeArgs Empty => new ProfileAdvancedAttributeArgs();
    }
}

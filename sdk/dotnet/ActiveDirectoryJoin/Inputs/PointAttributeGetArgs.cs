// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.ActiveDirectoryJoin.Inputs
{

    public sealed class PointAttributeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required for each attribute in the attribute list. Can contain an empty string.
        /// </summary>
        [Input("defaultValue", required: true)]
        public Input<string> DefaultValue { get; set; } = null!;

        /// <summary>
        /// Required for each attribute in the attribute list
        /// </summary>
        [Input("internalName", required: true)]
        public Input<string> InternalName { get; set; } = null!;

        /// <summary>
        /// Required for each attribute in the attribute list with no duplication between attributes
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Required for each group in the group list
        ///   - Choices: `STRING`, `IP`, `BOOLEAN`, `INT`, `OCTET_STRING`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public PointAttributeGetArgs()
        {
        }
        public static new PointAttributeGetArgs Empty => new PointAttributeGetArgs();
    }
}

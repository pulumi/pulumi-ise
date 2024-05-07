// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.ActiveDirectoryJoin.Outputs
{

    [OutputType]
    public sealed class PointAttribute
    {
        /// <summary>
        /// Required for each attribute in the attribute list. Can contain an empty string.
        /// </summary>
        public readonly string DefaultValue;
        /// <summary>
        /// Required for each attribute in the attribute list
        /// </summary>
        public readonly string InternalName;
        /// <summary>
        /// Required for each attribute in the attribute list with no duplication between attributes
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Required for each group in the group list
        ///   - Choices: `STRING`, `IP`, `BOOLEAN`, `INT`, `OCTET_STRING`
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private PointAttribute(
            string defaultValue,

            string internalName,

            string name,

            string type)
        {
            DefaultValue = defaultValue;
            InternalName = internalName;
            Name = name;
            Type = type;
        }
    }
}

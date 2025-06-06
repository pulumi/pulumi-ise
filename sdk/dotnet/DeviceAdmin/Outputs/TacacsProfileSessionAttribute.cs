// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.DeviceAdmin.Outputs
{

    [OutputType]
    public sealed class TacacsProfileSessionAttribute
    {
        /// <summary>
        /// Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type
        ///   - Choices: `MANDATORY`, `OPTIONAL`
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private TacacsProfileSessionAttribute(
            string name,

            string type,

            string value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }
}

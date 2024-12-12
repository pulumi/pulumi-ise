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
    public sealed class GetTacacsProfileSessionAttributeResult
    {
        /// <summary>
        /// Name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetTacacsProfileSessionAttributeResult(
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
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement.Inputs
{

    public sealed class IdentitySourceSequenceIdentitySourceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the identity source
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Order of the identity source in the sequence
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        public IdentitySourceSequenceIdentitySourceArgs()
        {
        }
        public static new IdentitySourceSequenceIdentitySourceArgs Empty => new IdentitySourceSequenceIdentitySourceArgs();
    }
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TacacsCommand.Inputs
{

    public sealed class SetCommandArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command arguments
        /// </summary>
        [Input("arguments", required: true)]
        public Input<string> Arguments { get; set; } = null!;

        /// <summary>
        /// Command
        /// </summary>
        [Input("command", required: true)]
        public Input<string> Command { get; set; } = null!;

        /// <summary>
        /// Grant
        ///   - Choices: `PERMIT`, `DENY`, `DENY_ALWAYS`
        /// </summary>
        [Input("grant", required: true)]
        public Input<string> Grant { get; set; } = null!;

        public SetCommandArgs()
        {
        }
        public static new SetCommandArgs Empty => new SetCommandArgs();
    }
}

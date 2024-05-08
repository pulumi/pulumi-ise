// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.System.Inputs
{

    public sealed class LicenseTierStateLicenseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// License name
        ///   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// License status
        ///   - Choices: `ENABLED`, `DISABLED`
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public LicenseTierStateLicenseArgs()
        {
        }
        public static new LicenseTierStateLicenseArgs Empty => new LicenseTierStateLicenseArgs();
    }
}
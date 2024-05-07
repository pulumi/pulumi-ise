// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.ActiveDirectoryAdd.Inputs
{

    public sealed class GroupsGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required for each group in the group list with no duplication between groups
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Required for each group in the group list with no duplication between groups
        /// </summary>
        [Input("sid", required: true)]
        public Input<string> Sid { get; set; } = null!;

        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupsGroupArgs()
        {
        }
        public static new GroupsGroupArgs Empty => new GroupsGroupArgs();
    }
}

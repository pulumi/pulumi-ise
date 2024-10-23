// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccess
{
    /// <summary>
    /// This resource is used to update rank field in network access policy set. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and network_access_policy_set resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.
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
    ///     var example = new Ise.NetworkAccess.PolicySetUpdateRank("example", new()
    ///     {
    ///         PolicySetId = "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
    ///         Rank = 0,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [IseResourceType("ise:networkaccess/policySetUpdateRank:PolicySetUpdateRank")]
    public partial class PolicySetUpdateRank : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Policy set ID
        /// </summary>
        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Output("rank")]
        public Output<int> Rank { get; private set; } = null!;


        /// <summary>
        /// Create a PolicySetUpdateRank resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicySetUpdateRank(string name, PolicySetUpdateRankArgs args, CustomResourceOptions? options = null)
            : base("ise:networkaccess/policySetUpdateRank:PolicySetUpdateRank", name, args ?? new PolicySetUpdateRankArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicySetUpdateRank(string name, Input<string> id, PolicySetUpdateRankState? state = null, CustomResourceOptions? options = null)
            : base("ise:networkaccess/policySetUpdateRank:PolicySetUpdateRank", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicySetUpdateRank resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicySetUpdateRank Get(string name, Input<string> id, PolicySetUpdateRankState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicySetUpdateRank(name, id, state, options);
        }
    }

    public sealed class PolicySetUpdateRankArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy set ID
        /// </summary>
        [Input("policySetId", required: true)]
        public Input<string> PolicySetId { get; set; } = null!;

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Input("rank", required: true)]
        public Input<int> Rank { get; set; } = null!;

        public PolicySetUpdateRankArgs()
        {
        }
        public static new PolicySetUpdateRankArgs Empty => new PolicySetUpdateRankArgs();
    }

    public sealed class PolicySetUpdateRankState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Policy set ID
        /// </summary>
        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        /// <summary>
        /// The rank (priority) in relation to other rules. Lower rank is higher priority.
        /// </summary>
        [Input("rank")]
        public Input<int>? Rank { get; set; }

        public PolicySetUpdateRankState()
        {
        }
        public static new PolicySetUpdateRankState Empty => new PolicySetUpdateRankState();
    }
}

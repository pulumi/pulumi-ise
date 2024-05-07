// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustsecIpToSgtMapping
{
    /// <summary>
    /// This resource can manage a TrustSec IP to SGT Mapping Group.
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
    ///     var example = new Ise.TrustsecIpToSgtMapping.Group("example", new()
    ///     {
    ///         Name = "groupA",
    ///         DeployType = "ALL",
    ///         Sgt = "93e1bf00-8c01-11e6-996c-525400b48521",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:TrustsecIpToSgtMapping/group:Group example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:TrustsecIpToSgtMapping/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Mandatory unless `deploy_type` is `ALL`
        /// </summary>
        [Output("deployTo")]
        public Output<string?> DeployTo { get; private set; } = null!;

        /// <summary>
        /// Deploy Type - Choices: `ALL`, `ND`, `NDG`
        /// </summary>
        [Output("deployType")]
        public Output<string> DeployType { get; private set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the IP to SGT mapping Group
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Trustsec Security Group ID
        /// </summary>
        [Output("sgt")]
        public Output<string> Sgt { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("ise:TrustsecIpToSgtMapping/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("ise:TrustsecIpToSgtMapping/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mandatory unless `deploy_type` is `ALL`
        /// </summary>
        [Input("deployTo")]
        public Input<string>? DeployTo { get; set; }

        /// <summary>
        /// Deploy Type - Choices: `ALL`, `ND`, `NDG`
        /// </summary>
        [Input("deployType", required: true)]
        public Input<string> DeployType { get; set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the IP to SGT mapping Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Trustsec Security Group ID
        /// </summary>
        [Input("sgt", required: true)]
        public Input<string> Sgt { get; set; } = null!;

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mandatory unless `deploy_type` is `ALL`
        /// </summary>
        [Input("deployTo")]
        public Input<string>? DeployTo { get; set; }

        /// <summary>
        /// Deploy Type - Choices: `ALL`, `ND`, `NDG`
        /// </summary>
        [Input("deployType")]
        public Input<string>? DeployType { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the IP to SGT mapping Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Trustsec Security Group ID
        /// </summary>
        [Input("sgt")]
        public Input<string>? Sgt { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}

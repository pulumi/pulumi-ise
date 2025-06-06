// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement
{
    /// <summary>
    /// This resource can manage an Endpoint Identity Group.
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
    ///     var example = new Ise.IdentityManagement.EndpointIdentityGroup("example", new()
    ///     {
    ///         Name = "Group1",
    ///         Description = "My endpoint identity group",
    ///         SystemDefined = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup")]
    public partial class EndpointIdentityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the endpoint identity group
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Parent endpoint identity group ID
        /// </summary>
        [Output("parentEndpointIdentityGroupId")]
        public Output<string?> ParentEndpointIdentityGroupId { get; private set; } = null!;

        /// <summary>
        /// System defined endpoint identity group
        /// </summary>
        [Output("systemDefined")]
        public Output<bool?> SystemDefined { get; private set; } = null!;


        /// <summary>
        /// Create a EndpointIdentityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EndpointIdentityGroup(string name, EndpointIdentityGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, args ?? new EndpointIdentityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EndpointIdentityGroup(string name, Input<string> id, EndpointIdentityGroupState? state = null, CustomResourceOptions? options = null)
            : base("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EndpointIdentityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EndpointIdentityGroup Get(string name, Input<string> id, EndpointIdentityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new EndpointIdentityGroup(name, id, state, options);
        }
    }

    public sealed class EndpointIdentityGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the endpoint identity group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parent endpoint identity group ID
        /// </summary>
        [Input("parentEndpointIdentityGroupId")]
        public Input<string>? ParentEndpointIdentityGroupId { get; set; }

        /// <summary>
        /// System defined endpoint identity group
        /// </summary>
        [Input("systemDefined")]
        public Input<bool>? SystemDefined { get; set; }

        public EndpointIdentityGroupArgs()
        {
        }
        public static new EndpointIdentityGroupArgs Empty => new EndpointIdentityGroupArgs();
    }

    public sealed class EndpointIdentityGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the endpoint identity group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Parent endpoint identity group ID
        /// </summary>
        [Input("parentEndpointIdentityGroupId")]
        public Input<string>? ParentEndpointIdentityGroupId { get; set; }

        /// <summary>
        /// System defined endpoint identity group
        /// </summary>
        [Input("systemDefined")]
        public Input<bool>? SystemDefined { get; set; }

        public EndpointIdentityGroupState()
        {
        }
        public static new EndpointIdentityGroupState Empty => new EndpointIdentityGroupState();
    }
}

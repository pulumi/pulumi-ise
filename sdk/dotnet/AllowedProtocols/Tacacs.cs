// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.AllowedProtocols
{
    /// <summary>
    /// This resource can manage a TACACS allowed protocols policy element.
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
    ///     var example = new Ise.AllowedProtocols.Tacacs("example", new()
    ///     {
    ///         Name = "Protocols1",
    ///         Description = "My allowed TACACS protocols",
    ///         AllowPapAscii = true,
    ///         AllowChap = true,
    ///         AllowMsChapV1 = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:AllowedProtocols/tacacs:Tacacs example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:AllowedProtocols/tacacs:Tacacs")]
    public partial class Tacacs : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Allow CHAP
        /// </summary>
        [Output("allowChap")]
        public Output<bool> AllowChap { get; private set; } = null!;

        /// <summary>
        /// Allow MS CHAP v1
        /// </summary>
        [Output("allowMsChapV1")]
        public Output<bool> AllowMsChapV1 { get; private set; } = null!;

        /// <summary>
        /// Allow PAP ASCII
        /// </summary>
        [Output("allowPapAscii")]
        public Output<bool> AllowPapAscii { get; private set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The name of the allowed protocols
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Tacacs resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Tacacs(string name, TacacsArgs args, CustomResourceOptions? options = null)
            : base("ise:AllowedProtocols/tacacs:Tacacs", name, args ?? new TacacsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Tacacs(string name, Input<string> id, TacacsState? state = null, CustomResourceOptions? options = null)
            : base("ise:AllowedProtocols/tacacs:Tacacs", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Tacacs resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Tacacs Get(string name, Input<string> id, TacacsState? state = null, CustomResourceOptions? options = null)
        {
            return new Tacacs(name, id, state, options);
        }
    }

    public sealed class TacacsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow CHAP
        /// </summary>
        [Input("allowChap", required: true)]
        public Input<bool> AllowChap { get; set; } = null!;

        /// <summary>
        /// Allow MS CHAP v1
        /// </summary>
        [Input("allowMsChapV1", required: true)]
        public Input<bool> AllowMsChapV1 { get; set; } = null!;

        /// <summary>
        /// Allow PAP ASCII
        /// </summary>
        [Input("allowPapAscii", required: true)]
        public Input<bool> AllowPapAscii { get; set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the allowed protocols
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TacacsArgs()
        {
        }
        public static new TacacsArgs Empty => new TacacsArgs();
    }

    public sealed class TacacsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow CHAP
        /// </summary>
        [Input("allowChap")]
        public Input<bool>? AllowChap { get; set; }

        /// <summary>
        /// Allow MS CHAP v1
        /// </summary>
        [Input("allowMsChapV1")]
        public Input<bool>? AllowMsChapV1 { get; set; }

        /// <summary>
        /// Allow PAP ASCII
        /// </summary>
        [Input("allowPapAscii")]
        public Input<bool>? AllowPapAscii { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the allowed protocols
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public TacacsState()
        {
        }
        public static new TacacsState Empty => new TacacsState();
    }
}

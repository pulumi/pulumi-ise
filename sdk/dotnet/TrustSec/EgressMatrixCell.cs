// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustSec
{
    /// <summary>
    /// This resource can manage a TrustSec Egress Matrix Cell.
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
    ///     var example = new Ise.TrustSec.EgressMatrixCell("example", new()
    ///     {
    ///         Description = "EgressMatrixCell Description",
    ///         DefaultRule = "NONE",
    ///         MatrixCellStatus = "ENABLED",
    ///         Sgacls = new[]
    ///         {
    ///             "26b76b10-66e6-11ee-9cc1-9eb2a3ecc82a, 9d64dcd0-6384-11ee-9cc1-9eb2a3ecc82a",
    ///         },
    ///         SourceSgtId = "93c66ed0-8c01-11e6-996c-525400b48521",
    ///         DestinationSgtId = "93e1bf00-8c01-11e6-996c-525400b48521",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:trustsec/egressMatrixCell:EgressMatrixCell example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:trustsec/egressMatrixCell:EgressMatrixCell")]
    public partial class EgressMatrixCell : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
        /// </summary>
        [Output("defaultRule")]
        public Output<string> DefaultRule { get; private set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Destination Trustsec Security Group ID
        /// </summary>
        [Output("destinationSgtId")]
        public Output<string> DestinationSgtId { get; private set; } = null!;

        /// <summary>
        /// Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
        /// </summary>
        [Output("matrixCellStatus")]
        public Output<string> MatrixCellStatus { get; private set; } = null!;

        /// <summary>
        /// List of TrustSec Security Groups ACLs
        /// </summary>
        [Output("sgacls")]
        public Output<ImmutableArray<string>> Sgacls { get; private set; } = null!;

        /// <summary>
        /// Source Trustsec Security Group ID
        /// </summary>
        [Output("sourceSgtId")]
        public Output<string> SourceSgtId { get; private set; } = null!;


        /// <summary>
        /// Create a EgressMatrixCell resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EgressMatrixCell(string name, EgressMatrixCellArgs args, CustomResourceOptions? options = null)
            : base("ise:trustsec/egressMatrixCell:EgressMatrixCell", name, args ?? new EgressMatrixCellArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EgressMatrixCell(string name, Input<string> id, EgressMatrixCellState? state = null, CustomResourceOptions? options = null)
            : base("ise:trustsec/egressMatrixCell:EgressMatrixCell", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EgressMatrixCell resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EgressMatrixCell Get(string name, Input<string> id, EgressMatrixCellState? state = null, CustomResourceOptions? options = null)
        {
            return new EgressMatrixCell(name, id, state, options);
        }
    }

    public sealed class EgressMatrixCellArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
        /// </summary>
        [Input("defaultRule")]
        public Input<string>? DefaultRule { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination Trustsec Security Group ID
        /// </summary>
        [Input("destinationSgtId", required: true)]
        public Input<string> DestinationSgtId { get; set; } = null!;

        /// <summary>
        /// Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
        /// </summary>
        [Input("matrixCellStatus")]
        public Input<string>? MatrixCellStatus { get; set; }

        [Input("sgacls")]
        private InputList<string>? _sgacls;

        /// <summary>
        /// List of TrustSec Security Groups ACLs
        /// </summary>
        public InputList<string> Sgacls
        {
            get => _sgacls ?? (_sgacls = new InputList<string>());
            set => _sgacls = value;
        }

        /// <summary>
        /// Source Trustsec Security Group ID
        /// </summary>
        [Input("sourceSgtId", required: true)]
        public Input<string> SourceSgtId { get; set; } = null!;

        public EgressMatrixCellArgs()
        {
        }
        public static new EgressMatrixCellArgs Empty => new EgressMatrixCellArgs();
    }

    public sealed class EgressMatrixCellState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
        /// </summary>
        [Input("defaultRule")]
        public Input<string>? DefaultRule { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Destination Trustsec Security Group ID
        /// </summary>
        [Input("destinationSgtId")]
        public Input<string>? DestinationSgtId { get; set; }

        /// <summary>
        /// Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
        /// </summary>
        [Input("matrixCellStatus")]
        public Input<string>? MatrixCellStatus { get; set; }

        [Input("sgacls")]
        private InputList<string>? _sgacls;

        /// <summary>
        /// List of TrustSec Security Groups ACLs
        /// </summary>
        public InputList<string> Sgacls
        {
            get => _sgacls ?? (_sgacls = new InputList<string>());
            set => _sgacls = value;
        }

        /// <summary>
        /// Source Trustsec Security Group ID
        /// </summary>
        [Input("sourceSgtId")]
        public Input<string>? SourceSgtId { get; set; }

        public EgressMatrixCellState()
        {
        }
        public static new EgressMatrixCellState Empty => new EgressMatrixCellState();
    }
}

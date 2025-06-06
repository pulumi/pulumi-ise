// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustSec
{
    public static class GetEgressMatrixCell
    {
        /// <summary>
        /// This data source can read the TrustSec Egress Matrix Cell.
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
        ///     var example = Ise.TrustSec.GetEgressMatrixCell.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetEgressMatrixCellResult> InvokeAsync(GetEgressMatrixCellArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEgressMatrixCellResult>("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", args ?? new GetEgressMatrixCellArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TrustSec Egress Matrix Cell.
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
        ///     var example = Ise.TrustSec.GetEgressMatrixCell.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEgressMatrixCellResult> Invoke(GetEgressMatrixCellInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEgressMatrixCellResult>("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", args ?? new GetEgressMatrixCellInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TrustSec Egress Matrix Cell.
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
        ///     var example = Ise.TrustSec.GetEgressMatrixCell.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetEgressMatrixCellResult> Invoke(GetEgressMatrixCellInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetEgressMatrixCellResult>("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", args ?? new GetEgressMatrixCellInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEgressMatrixCellArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetEgressMatrixCellArgs()
        {
        }
        public static new GetEgressMatrixCellArgs Empty => new GetEgressMatrixCellArgs();
    }

    public sealed class GetEgressMatrixCellInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetEgressMatrixCellInvokeArgs()
        {
        }
        public static new GetEgressMatrixCellInvokeArgs Empty => new GetEgressMatrixCellInvokeArgs();
    }


    [OutputType]
    public sealed class GetEgressMatrixCellResult
    {
        /// <summary>
        /// Can be used only if sgacls not specified.
        /// </summary>
        public readonly string DefaultRule;
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Destination Trustsec Security Group ID
        /// </summary>
        public readonly string DestinationSgtId;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Matrix Cell Status
        /// </summary>
        public readonly string MatrixCellStatus;
        /// <summary>
        /// List of TrustSec Security Groups ACLs
        /// </summary>
        public readonly ImmutableArray<string> Sgacls;
        /// <summary>
        /// Source Trustsec Security Group ID
        /// </summary>
        public readonly string SourceSgtId;

        [OutputConstructor]
        private GetEgressMatrixCellResult(
            string defaultRule,

            string description,

            string destinationSgtId,

            string id,

            string matrixCellStatus,

            ImmutableArray<string> sgacls,

            string sourceSgtId)
        {
            DefaultRule = defaultRule;
            Description = description;
            DestinationSgtId = destinationSgtId;
            Id = id;
            MatrixCellStatus = matrixCellStatus;
            Sgacls = sgacls;
            SourceSgtId = sourceSgtId;
        }
    }
}

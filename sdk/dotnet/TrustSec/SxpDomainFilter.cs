// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustSec
{
    /// <summary>
    /// This resource can manage a SXP Domain Filter.
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
    ///     var example = new Ise.TrustSec.SxpDomainFilter("example", new()
    ///     {
    ///         Subnet = "1.0.0.0/24",
    ///         Vn = "VN1",
    ///         Domains = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:trustsec/sxpDomainFilter:SxpDomainFilter example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:trustsec/sxpDomainFilter:SxpDomainFilter")]
    public partial class SxpDomainFilter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of SXP Domains, separated with comma
        /// </summary>
        [Output("domains")]
        public Output<string> Domains { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// SGT name or ID. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Output("sgt")]
        public Output<string?> Sgt { get; private set; } = null!;

        /// <summary>
        /// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Output("subnet")]
        public Output<string?> Subnet { get; private set; } = null!;

        /// <summary>
        /// Virtual Network. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Output("vn")]
        public Output<string?> Vn { get; private set; } = null!;


        /// <summary>
        /// Create a SxpDomainFilter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SxpDomainFilter(string name, SxpDomainFilterArgs args, CustomResourceOptions? options = null)
            : base("ise:trustsec/sxpDomainFilter:SxpDomainFilter", name, args ?? new SxpDomainFilterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SxpDomainFilter(string name, Input<string> id, SxpDomainFilterState? state = null, CustomResourceOptions? options = null)
            : base("ise:trustsec/sxpDomainFilter:SxpDomainFilter", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SxpDomainFilter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SxpDomainFilter Get(string name, Input<string> id, SxpDomainFilterState? state = null, CustomResourceOptions? options = null)
        {
            return new SxpDomainFilter(name, id, state, options);
        }
    }

    public sealed class SxpDomainFilterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of SXP Domains, separated with comma
        /// </summary>
        [Input("domains", required: true)]
        public Input<string> Domains { get; set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SGT name or ID. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("sgt")]
        public Input<string>? Sgt { get; set; }

        /// <summary>
        /// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Virtual Network. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("vn")]
        public Input<string>? Vn { get; set; }

        public SxpDomainFilterArgs()
        {
        }
        public static new SxpDomainFilterArgs Empty => new SxpDomainFilterArgs();
    }

    public sealed class SxpDomainFilterState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// List of SXP Domains, separated with comma
        /// </summary>
        [Input("domains")]
        public Input<string>? Domains { get; set; }

        /// <summary>
        /// Resource name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// SGT name or ID. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("sgt")]
        public Input<string>? Sgt { get; set; }

        /// <summary>
        /// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("subnet")]
        public Input<string>? Subnet { get; set; }

        /// <summary>
        /// Virtual Network. At least one of subnet or sgt or vn should be defined
        /// </summary>
        [Input("vn")]
        public Input<string>? Vn { get; set; }

        public SxpDomainFilterState()
        {
        }
        public static new SxpDomainFilterState Empty => new SxpDomainFilterState();
    }
}
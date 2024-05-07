// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustsecIpToSgt
{
    public static class GetMapping
    {
        /// <summary>
        /// This data source can read the TrustSec IP to SGT Mapping.
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
        ///     var example = Ise.TrustsecIpToSgt.GetMapping.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetMappingResult> InvokeAsync(GetMappingArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetMappingResult>("ise:TrustsecIpToSgt/getMapping:getMapping", args ?? new GetMappingArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TrustSec IP to SGT Mapping.
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
        ///     var example = Ise.TrustsecIpToSgt.GetMapping.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetMappingResult> Invoke(GetMappingInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetMappingResult>("ise:TrustsecIpToSgt/getMapping:getMapping", args ?? new GetMappingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetMappingArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the IP to SGT mapping
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetMappingArgs()
        {
        }
        public static new GetMappingArgs Empty => new GetMappingArgs();
    }

    public sealed class GetMappingInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the IP to SGT mapping
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetMappingInvokeArgs()
        {
        }
        public static new GetMappingInvokeArgs Empty => new GetMappingInvokeArgs();
    }


    [OutputType]
    public sealed class GetMappingResult
    {
        /// <summary>
        /// Mandatory unless `mapping_group` is set or unless `deploy_type` is `ALL`
        /// </summary>
        public readonly string DeployTo;
        /// <summary>
        /// Deploy Type
        /// </summary>
        public readonly string DeployType;
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Mandatory if `host_name` is empty
        /// </summary>
        public readonly string HostIp;
        /// <summary>
        /// Mandatory if `host_ip` is empty
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deploy_to` and `deploy_type` are set
        /// </summary>
        public readonly string MappingGroup;
        /// <summary>
        /// The name of the IP to SGT mapping
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Trustsec Security Group ID. Mandatory unless `mapping_group` is set
        /// </summary>
        public readonly string Sgt;

        [OutputConstructor]
        private GetMappingResult(
            string deployTo,

            string deployType,

            string description,

            string hostIp,

            string hostName,

            string id,

            string mappingGroup,

            string name,

            string sgt)
        {
            DeployTo = deployTo;
            DeployType = deployType;
            Description = description;
            HostIp = hostIp;
            HostName = hostName;
            Id = id;
            MappingGroup = mappingGroup;
            Name = name;
            Sgt = sgt;
        }
    }
}

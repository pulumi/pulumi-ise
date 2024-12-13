// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.TrustSec
{
    public static class GetSecurityGroupAcl
    {
        /// <summary>
        /// This data source can read the TrustSec Security Group ACL.
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
        ///     var example = Ise.TrustSec.GetSecurityGroupAcl.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetSecurityGroupAclResult> InvokeAsync(GetSecurityGroupAclArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecurityGroupAclResult>("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", args ?? new GetSecurityGroupAclArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TrustSec Security Group ACL.
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
        ///     var example = Ise.TrustSec.GetSecurityGroupAcl.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSecurityGroupAclResult> Invoke(GetSecurityGroupAclInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupAclResult>("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", args ?? new GetSecurityGroupAclInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TrustSec Security Group ACL.
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
        ///     var example = Ise.TrustSec.GetSecurityGroupAcl.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetSecurityGroupAclResult> Invoke(GetSecurityGroupAclInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecurityGroupAclResult>("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", args ?? new GetSecurityGroupAclInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecurityGroupAclArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the security group ACL
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetSecurityGroupAclArgs()
        {
        }
        public static new GetSecurityGroupAclArgs Empty => new GetSecurityGroupAclArgs();
    }

    public sealed class GetSecurityGroupAclInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the security group ACL
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetSecurityGroupAclInvokeArgs()
        {
        }
        public static new GetSecurityGroupAclInvokeArgs Empty => new GetSecurityGroupAclInvokeArgs();
    }


    [OutputType]
    public sealed class GetSecurityGroupAclResult
    {
        /// <summary>
        /// Content of ACL
        /// </summary>
        public readonly string AclContent;
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP Version
        /// </summary>
        public readonly string IpVersion;
        /// <summary>
        /// The name of the security group ACL
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Read-only
        /// </summary>
        public readonly bool ReadOnly;

        [OutputConstructor]
        private GetSecurityGroupAclResult(
            string aclContent,

            string description,

            string id,

            string ipVersion,

            string name,

            bool readOnly)
        {
            AclContent = aclContent;
            Description = description;
            Id = id;
            IpVersion = ipVersion;
            Name = name;
            ReadOnly = readOnly;
        }
    }
}

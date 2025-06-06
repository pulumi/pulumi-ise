// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement
{
    public static class GetCertificateAuthenticationProfile
    {
        /// <summary>
        /// This data source can read the Certificate Authentication Profile.
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
        ///     var example = Ise.IdentityManagement.GetCertificateAuthenticationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCertificateAuthenticationProfileResult> InvokeAsync(GetCertificateAuthenticationProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateAuthenticationProfileResult>("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", args ?? new GetCertificateAuthenticationProfileArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Certificate Authentication Profile.
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
        ///     var example = Ise.IdentityManagement.GetCertificateAuthenticationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCertificateAuthenticationProfileResult> Invoke(GetCertificateAuthenticationProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateAuthenticationProfileResult>("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", args ?? new GetCertificateAuthenticationProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Certificate Authentication Profile.
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
        ///     var example = Ise.IdentityManagement.GetCertificateAuthenticationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCertificateAuthenticationProfileResult> Invoke(GetCertificateAuthenticationProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateAuthenticationProfileResult>("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", args ?? new GetCertificateAuthenticationProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateAuthenticationProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the certificate profile
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetCertificateAuthenticationProfileArgs()
        {
        }
        public static new GetCertificateAuthenticationProfileArgs Empty => new GetCertificateAuthenticationProfileArgs();
    }

    public sealed class GetCertificateAuthenticationProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the certificate profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetCertificateAuthenticationProfileInvokeArgs()
        {
        }
        public static new GetCertificateAuthenticationProfileInvokeArgs Empty => new GetCertificateAuthenticationProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateAuthenticationProfileResult
    {
        /// <summary>
        /// Allow as username
        /// </summary>
        public readonly bool AllowedAsUserName;
        /// <summary>
        /// Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in `username_from`.
        /// </summary>
        public readonly string CertificateAttributeName;
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Referred IDStore name for the Certificate Profile or `[not applicable]` in case no identity store is chosen
        /// </summary>
        public readonly string ExternalIdentityStoreName;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Match mode of the Certificate Profile. Allowed values: NEVER, RESOLVE*IDENTITY*AMBIGUITY, BINARY_COMPARISON
        /// </summary>
        public readonly string MatchMode;
        /// <summary>
        /// The name of the certificate profile
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The attribute in the certificate where the user name should be taken from. Allowed values: `CERTIFICATE` (for a specific attribute as defined in certificateAttributeName), `UPN` (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
        /// </summary>
        public readonly string UsernameFrom;

        [OutputConstructor]
        private GetCertificateAuthenticationProfileResult(
            bool allowedAsUserName,

            string certificateAttributeName,

            string description,

            string externalIdentityStoreName,

            string id,

            string matchMode,

            string name,

            string usernameFrom)
        {
            AllowedAsUserName = allowedAsUserName;
            CertificateAttributeName = certificateAttributeName;
            Description = description;
            ExternalIdentityStoreName = externalIdentityStoreName;
            Id = id;
            MatchMode = matchMode;
            Name = name;
            UsernameFrom = usernameFrom;
        }
    }
}

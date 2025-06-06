// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.DeviceAdmin
{
    public static class GetTacacsProfile
    {
        /// <summary>
        /// This data source can read the TACACS Profile.
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
        ///     var example = Ise.DeviceAdmin.GetTacacsProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTacacsProfileResult> InvokeAsync(GetTacacsProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTacacsProfileResult>("ise:deviceadmin/getTacacsProfile:getTacacsProfile", args ?? new GetTacacsProfileArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TACACS Profile.
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
        ///     var example = Ise.DeviceAdmin.GetTacacsProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTacacsProfileResult> Invoke(GetTacacsProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTacacsProfileResult>("ise:deviceadmin/getTacacsProfile:getTacacsProfile", args ?? new GetTacacsProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the TACACS Profile.
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
        ///     var example = Ise.DeviceAdmin.GetTacacsProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTacacsProfileResult> Invoke(GetTacacsProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetTacacsProfileResult>("ise:deviceadmin/getTacacsProfile:getTacacsProfile", args ?? new GetTacacsProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetTacacsProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the TACACS profile
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetTacacsProfileArgs()
        {
        }
        public static new GetTacacsProfileArgs Empty => new GetTacacsProfileArgs();
    }

    public sealed class GetTacacsProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the TACACS profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetTacacsProfileInvokeArgs()
        {
        }
        public static new GetTacacsProfileInvokeArgs Empty => new GetTacacsProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetTacacsProfileResult
    {
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the TACACS profile
        /// </summary>
        public readonly string Name;
        public readonly ImmutableArray<Outputs.GetTacacsProfileSessionAttributeResult> SessionAttributes;

        [OutputConstructor]
        private GetTacacsProfileResult(
            string description,

            string id,

            string name,

            ImmutableArray<Outputs.GetTacacsProfileSessionAttributeResult> sessionAttributes)
        {
            Description = description;
            Id = id;
            Name = name;
            SessionAttributes = sessionAttributes;
        }
    }
}

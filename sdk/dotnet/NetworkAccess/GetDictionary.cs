// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccess
{
    public static class GetDictionary
    {
        /// <summary>
        /// This data source can read the Network Access Dictionary.
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
        ///     var example = Ise.NetworkAccess.GetDictionary.Invoke(new()
        ///     {
        ///         Id = "Dict1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDictionaryResult> InvokeAsync(GetDictionaryArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDictionaryResult>("ise:networkaccess/getDictionary:getDictionary", args ?? new GetDictionaryArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Network Access Dictionary.
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
        ///     var example = Ise.NetworkAccess.GetDictionary.Invoke(new()
        ///     {
        ///         Id = "Dict1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDictionaryResult> Invoke(GetDictionaryInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDictionaryResult>("ise:networkaccess/getDictionary:getDictionary", args ?? new GetDictionaryInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Network Access Dictionary.
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
        ///     var example = Ise.NetworkAccess.GetDictionary.Invoke(new()
        ///     {
        ///         Id = "Dict1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDictionaryResult> Invoke(GetDictionaryInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDictionaryResult>("ise:networkaccess/getDictionary:getDictionary", args ?? new GetDictionaryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDictionaryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The dictionary name
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetDictionaryArgs()
        {
        }
        public static new GetDictionaryArgs Empty => new GetDictionaryArgs();
    }

    public sealed class GetDictionaryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The dictionary name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetDictionaryInvokeArgs()
        {
        }
        public static new GetDictionaryInvokeArgs Empty => new GetDictionaryInvokeArgs();
    }


    [OutputType]
    public sealed class GetDictionaryResult
    {
        /// <summary>
        /// The description of the dictionary
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The dictionary attribute type
        /// </summary>
        public readonly string DictionaryAttrType;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The dictionary name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The version of the dictionary
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetDictionaryResult(
            string description,

            string dictionaryAttrType,

            string id,

            string name,

            string version)
        {
            Description = description;
            DictionaryAttrType = dictionaryAttrType;
            Id = id;
            Name = name;
            Version = version;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.System
{
    public static class GetRepository
    {
        /// <summary>
        /// This data source can read the Repository.
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
        ///     var example = Ise.System.GetRepository.Invoke(new()
        ///     {
        ///         Id = "repo1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetRepositoryResult> InvokeAsync(GetRepositoryArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRepositoryResult>("ise:system/getRepository:getRepository", args ?? new GetRepositoryArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Repository.
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
        ///     var example = Ise.System.GetRepository.Invoke(new()
        ///     {
        ///         Id = "repo1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetRepositoryResult> Invoke(GetRepositoryInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRepositoryResult>("ise:system/getRepository:getRepository", args ?? new GetRepositoryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRepositoryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetRepositoryArgs()
        {
        }
        public static new GetRepositoryArgs Empty => new GetRepositoryArgs();
    }

    public sealed class GetRepositoryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetRepositoryInvokeArgs()
        {
        }
        public static new GetRepositoryInvokeArgs Empty => new GetRepositoryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRepositoryResult
    {
        /// <summary>
        /// Enable PKI
        /// </summary>
        public readonly bool EnablePki;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Password can contain alphanumeric and/or special characters.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Protocol
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Name of the server
        /// </summary>
        public readonly string ServerName;
        /// <summary>
        /// User name
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private GetRepositoryResult(
            bool enablePki,

            string id,

            string name,

            string password,

            string path,

            string protocol,

            string serverName,

            string userName)
        {
            EnablePki = enablePki;
            Id = id;
            Name = name;
            Password = password;
            Path = path;
            Protocol = protocol;
            ServerName = serverName;
            UserName = userName;
        }
    }
}
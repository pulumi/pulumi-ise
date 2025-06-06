// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.System
{
    /// <summary>
    /// This resource can manage a Repository.
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
    ///     var example = new Ise.System.Repository("example", new()
    ///     {
    ///         Name = "repo1",
    ///         Protocol = "SFTP",
    ///         Path = "/dir",
    ///         ServerName = "server1",
    ///         UserName = "user9",
    ///         Password = "cisco123",
    ///         EnablePki = false,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:system/repository:Repository example "repo1"
    /// ```
    /// </summary>
    [IseResourceType("ise:system/repository:Repository")]
    public partial class Repository : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Enable PKI
        /// </summary>
        [Output("enablePki")]
        public Output<bool?> EnablePki { get; private set; } = null!;

        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Password can contain alphanumeric and/or special characters.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// Name of the server
        /// </summary>
        [Output("serverName")]
        public Output<string?> ServerName { get; private set; } = null!;

        /// <summary>
        /// User name
        /// </summary>
        [Output("userName")]
        public Output<string?> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a Repository resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Repository(string name, RepositoryArgs args, CustomResourceOptions? options = null)
            : base("ise:system/repository:Repository", name, args ?? new RepositoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Repository(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
            : base("ise:system/repository:Repository", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Repository resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Repository Get(string name, Input<string> id, RepositoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Repository(name, id, state, options);
        }
    }

    public sealed class RepositoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable PKI
        /// </summary>
        [Input("enablePki")]
        public Input<bool>? EnablePki { get; set; }

        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password can contain alphanumeric and/or special characters.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// Name of the server
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// User name
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public RepositoryArgs()
        {
        }
        public static new RepositoryArgs Empty => new RepositoryArgs();
    }

    public sealed class RepositoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable PKI
        /// </summary>
        [Input("enablePki")]
        public Input<bool>? EnablePki { get; set; }

        /// <summary>
        /// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Password can contain alphanumeric and/or special characters.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Name of the server
        /// </summary>
        [Input("serverName")]
        public Input<string>? ServerName { get; set; }

        /// <summary>
        /// User name
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public RepositoryState()
        {
        }
        public static new RepositoryState Empty => new RepositoryState();
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.Internal
{
    /// <summary>
    /// This resource can manage an Internal User.
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
    ///     var example = new Ise.Internal.User("example", new()
    ///     {
    ///         Name = "UserTF",
    ///         Password = "Cisco123",
    ///         ChangePassword = true,
    ///         Email = "aaa@cisco.com",
    ///         AccountNameAlias = "User 1",
    ///         EnablePassword = "Cisco123",
    ///         Enabled = true,
    ///         PasswordNeverExpires = false,
    ///         FirstName = "John",
    ///         LastName = "Doe",
    ///         PasswordIdStore = "Internal Users",
    ///         Description = "My first Terraform user",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:Internal/user:User example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:Internal/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
        /// from ISE 3.2.
        /// </summary>
        [Output("accountNameAlias")]
        public Output<string?> AccountNameAlias { get; private set; } = null!;

        /// <summary>
        /// Requires the user to change the password - Default value: `true`
        /// </summary>
        [Output("changePassword")]
        public Output<bool> ChangePassword { get; private set; } = null!;

        /// <summary>
        /// Key value map
        /// </summary>
        [Output("customAttributes")]
        public Output<string?> CustomAttributes { get; private set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Email address
        /// </summary>
        [Output("email")]
        public Output<string?> Email { get; private set; } = null!;

        /// <summary>
        /// This field is added in ISE 2.0 to support TACACS+
        /// </summary>
        [Output("enablePassword")]
        public Output<string?> EnablePassword { get; private set; } = null!;

        /// <summary>
        /// Whether the user is enabled/disabled
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// First name of the internal user
        /// </summary>
        [Output("firstName")]
        public Output<string?> FirstName { get; private set; } = null!;

        /// <summary>
        /// Comma separated list of identity group IDs.
        /// </summary>
        [Output("identityGroups")]
        public Output<string?> IdentityGroups { get; private set; } = null!;

        /// <summary>
        /// Last name of the internal user
        /// </summary>
        [Output("lastName")]
        public Output<string?> LastName { get; private set; } = null!;

        /// <summary>
        /// The name of the internal user
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password of the internal user
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The ID store where the internal user's password is kept - Default value: `Internal Users`
        /// </summary>
        [Output("passwordIdStore")]
        public Output<string> PasswordIdStore { get; private set; } = null!;

        /// <summary>
        /// Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
        /// field is only supported from ISE 3.2. - Default value: `false`
        /// </summary>
        [Output("passwordNeverExpires")]
        public Output<bool> PasswordNeverExpires { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("ise:Internal/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("ise:Internal/user:User", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
        /// from ISE 3.2.
        /// </summary>
        [Input("accountNameAlias")]
        public Input<string>? AccountNameAlias { get; set; }

        /// <summary>
        /// Requires the user to change the password - Default value: `true`
        /// </summary>
        [Input("changePassword")]
        public Input<bool>? ChangePassword { get; set; }

        /// <summary>
        /// Key value map
        /// </summary>
        [Input("customAttributes")]
        public Input<string>? CustomAttributes { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Email address
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// This field is added in ISE 2.0 to support TACACS+
        /// </summary>
        [Input("enablePassword")]
        public Input<string>? EnablePassword { get; set; }

        /// <summary>
        /// Whether the user is enabled/disabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// First name of the internal user
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Comma separated list of identity group IDs.
        /// </summary>
        [Input("identityGroups")]
        public Input<string>? IdentityGroups { get; set; }

        /// <summary>
        /// Last name of the internal user
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// The name of the internal user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password of the internal user
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// The ID store where the internal user's password is kept - Default value: `Internal Users`
        /// </summary>
        [Input("passwordIdStore")]
        public Input<string>? PasswordIdStore { get; set; }

        /// <summary>
        /// Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
        /// field is only supported from ISE 3.2. - Default value: `false`
        /// </summary>
        [Input("passwordNeverExpires")]
        public Input<bool>? PasswordNeverExpires { get; set; }

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
        /// from ISE 3.2.
        /// </summary>
        [Input("accountNameAlias")]
        public Input<string>? AccountNameAlias { get; set; }

        /// <summary>
        /// Requires the user to change the password - Default value: `true`
        /// </summary>
        [Input("changePassword")]
        public Input<bool>? ChangePassword { get; set; }

        /// <summary>
        /// Key value map
        /// </summary>
        [Input("customAttributes")]
        public Input<string>? CustomAttributes { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Email address
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// This field is added in ISE 2.0 to support TACACS+
        /// </summary>
        [Input("enablePassword")]
        public Input<string>? EnablePassword { get; set; }

        /// <summary>
        /// Whether the user is enabled/disabled
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// First name of the internal user
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Comma separated list of identity group IDs.
        /// </summary>
        [Input("identityGroups")]
        public Input<string>? IdentityGroups { get; set; }

        /// <summary>
        /// Last name of the internal user
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// The name of the internal user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password of the internal user
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The ID store where the internal user's password is kept - Default value: `Internal Users`
        /// </summary>
        [Input("passwordIdStore")]
        public Input<string>? PasswordIdStore { get; set; }

        /// <summary>
        /// Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
        /// field is only supported from ISE 3.2. - Default value: `false`
        /// </summary>
        [Input("passwordNeverExpires")]
        public Input<bool>? PasswordNeverExpires { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}

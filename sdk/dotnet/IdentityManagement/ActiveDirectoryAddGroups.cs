// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement
{
    /// <summary>
    /// This resource can manage an Active Directory Add Groups.
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
    ///     var example = new Ise.IdentityManagement.ActiveDirectoryAddGroups("example", new()
    ///     {
    ///         JoinPointId = "73808580-b6e6-11ee-8960-de6d7692bc40",
    ///         Name = "cisco.local",
    ///         Description = "My AD join point",
    ///         Domain = "cisco.local",
    ///         AdScopesNames = "Default_Scope",
    ///         EnableDomainAllowedList = true,
    ///         GroupsValue = new[]
    ///         {
    ///             new Ise.IdentityManagement.Inputs.ActiveDirectoryAddGroupsGroupArgs
    ///             {
    ///                 Name = "cisco.local/operators",
    ///                 Sid = "S-1-5-32-548",
    ///                 Type = "GLOBAL",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [IseResourceType("ise:identitymanagement/activeDirectoryAddGroups:ActiveDirectoryAddGroups")]
    public partial class ActiveDirectoryAddGroups : global::Pulumi.CustomResource
    {
        /// <summary>
        /// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        /// value: `Default_Scope`
        /// </summary>
        [Output("adScopesNames")]
        public Output<string> AdScopesNames { get; private set; } = null!;

        /// <summary>
        /// Join point Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// AD domain associated with the join point
        /// </summary>
        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        /// <summary>
        /// - Default value: `true`
        /// </summary>
        [Output("enableDomainAllowedList")]
        public Output<bool> EnableDomainAllowedList { get; private set; } = null!;

        /// <summary>
        /// List of AD Groups
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<Outputs.ActiveDirectoryAddGroupsGroup>> GroupsValue { get; private set; } = null!;

        /// <summary>
        /// Active Directory Join Point ID
        /// </summary>
        [Output("joinPointId")]
        public Output<string> JoinPointId { get; private set; } = null!;

        /// <summary>
        /// The name of the active directory join point
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a ActiveDirectoryAddGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ActiveDirectoryAddGroups(string name, ActiveDirectoryAddGroupsArgs args, CustomResourceOptions? options = null)
            : base("ise:identitymanagement/activeDirectoryAddGroups:ActiveDirectoryAddGroups", name, args ?? new ActiveDirectoryAddGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ActiveDirectoryAddGroups(string name, Input<string> id, ActiveDirectoryAddGroupsState? state = null, CustomResourceOptions? options = null)
            : base("ise:identitymanagement/activeDirectoryAddGroups:ActiveDirectoryAddGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ActiveDirectoryAddGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ActiveDirectoryAddGroups Get(string name, Input<string> id, ActiveDirectoryAddGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new ActiveDirectoryAddGroups(name, id, state, options);
        }
    }

    public sealed class ActiveDirectoryAddGroupsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        /// value: `Default_Scope`
        /// </summary>
        [Input("adScopesNames")]
        public Input<string>? AdScopesNames { get; set; }

        /// <summary>
        /// Join point Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// AD domain associated with the join point
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// - Default value: `true`
        /// </summary>
        [Input("enableDomainAllowedList")]
        public Input<bool>? EnableDomainAllowedList { get; set; }

        [Input("groups")]
        private InputList<Inputs.ActiveDirectoryAddGroupsGroupArgs>? _groups;

        /// <summary>
        /// List of AD Groups
        /// </summary>
        public InputList<Inputs.ActiveDirectoryAddGroupsGroupArgs> GroupsValue
        {
            get => _groups ?? (_groups = new InputList<Inputs.ActiveDirectoryAddGroupsGroupArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Active Directory Join Point ID
        /// </summary>
        [Input("joinPointId", required: true)]
        public Input<string> JoinPointId { get; set; } = null!;

        /// <summary>
        /// The name of the active directory join point
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ActiveDirectoryAddGroupsArgs()
        {
        }
        public static new ActiveDirectoryAddGroupsArgs Empty => new ActiveDirectoryAddGroupsArgs();
    }

    public sealed class ActiveDirectoryAddGroupsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        /// value: `Default_Scope`
        /// </summary>
        [Input("adScopesNames")]
        public Input<string>? AdScopesNames { get; set; }

        /// <summary>
        /// Join point Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// AD domain associated with the join point
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// - Default value: `true`
        /// </summary>
        [Input("enableDomainAllowedList")]
        public Input<bool>? EnableDomainAllowedList { get; set; }

        [Input("groups")]
        private InputList<Inputs.ActiveDirectoryAddGroupsGroupGetArgs>? _groups;

        /// <summary>
        /// List of AD Groups
        /// </summary>
        public InputList<Inputs.ActiveDirectoryAddGroupsGroupGetArgs> GroupsValue
        {
            get => _groups ?? (_groups = new InputList<Inputs.ActiveDirectoryAddGroupsGroupGetArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// Active Directory Join Point ID
        /// </summary>
        [Input("joinPointId")]
        public Input<string>? JoinPointId { get; set; }

        /// <summary>
        /// The name of the active directory join point
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ActiveDirectoryAddGroupsState()
        {
        }
        public static new ActiveDirectoryAddGroupsState Empty => new ActiveDirectoryAddGroupsState();
    }
}
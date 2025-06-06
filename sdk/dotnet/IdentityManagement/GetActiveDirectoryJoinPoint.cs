// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.IdentityManagement
{
    public static class GetActiveDirectoryJoinPoint
    {
        /// <summary>
        /// This data source can read the Active Directory Join Point.
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
        ///     var example = Ise.IdentityManagement.GetActiveDirectoryJoinPoint.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetActiveDirectoryJoinPointResult> InvokeAsync(GetActiveDirectoryJoinPointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetActiveDirectoryJoinPointResult>("ise:identitymanagement/getActiveDirectoryJoinPoint:getActiveDirectoryJoinPoint", args ?? new GetActiveDirectoryJoinPointArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Active Directory Join Point.
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
        ///     var example = Ise.IdentityManagement.GetActiveDirectoryJoinPoint.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActiveDirectoryJoinPointResult> Invoke(GetActiveDirectoryJoinPointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetActiveDirectoryJoinPointResult>("ise:identitymanagement/getActiveDirectoryJoinPoint:getActiveDirectoryJoinPoint", args ?? new GetActiveDirectoryJoinPointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the Active Directory Join Point.
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
        ///     var example = Ise.IdentityManagement.GetActiveDirectoryJoinPoint.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetActiveDirectoryJoinPointResult> Invoke(GetActiveDirectoryJoinPointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetActiveDirectoryJoinPointResult>("ise:identitymanagement/getActiveDirectoryJoinPoint:getActiveDirectoryJoinPoint", args ?? new GetActiveDirectoryJoinPointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetActiveDirectoryJoinPointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetActiveDirectoryJoinPointArgs()
        {
        }
        public static new GetActiveDirectoryJoinPointArgs Empty => new GetActiveDirectoryJoinPointArgs();
    }

    public sealed class GetActiveDirectoryJoinPointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetActiveDirectoryJoinPointInvokeArgs()
        {
        }
        public static new GetActiveDirectoryJoinPointInvokeArgs Empty => new GetActiveDirectoryJoinPointInvokeArgs();
    }


    [OutputType]
    public sealed class GetActiveDirectoryJoinPointResult
    {
        /// <summary>
        /// String that contains the names of the scopes that the active directory belongs to. Names are separated by comma.
        /// </summary>
        public readonly string AdScopesNames;
        /// <summary>
        /// Aging Time
        /// </summary>
        public readonly int AgingTime;
        /// <summary>
        /// List of AD attributes
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActiveDirectoryJoinPointAttributeResult> Attributes;
        /// <summary>
        /// Enable prevent AD account lockout for WIRELESS/WIRED/BOTH
        /// </summary>
        public readonly string AuthProtectionType;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string Country;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string Department;
        /// <summary>
        /// Join point description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// AD domain associated with the join point
        /// </summary>
        public readonly string Domain;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Enable Callback For Dial In Client
        /// </summary>
        public readonly bool EnableCallbackForDialinClient;
        /// <summary>
        /// Enable Dial In Permission Check
        /// </summary>
        public readonly bool EnableDialinPermissionCheck;
        public readonly bool EnableDomainAllowedList;
        /// <summary>
        /// Enable prevent AD account lockout due to too many bad password attempts
        /// </summary>
        public readonly bool EnableFailedAuthProtection;
        /// <summary>
        /// Enable Machine Access
        /// </summary>
        public readonly bool EnableMachineAccess;
        /// <summary>
        /// Enable Machine Authentication
        /// </summary>
        public readonly bool EnableMachineAuth;
        /// <summary>
        /// Enable Password Change
        /// </summary>
        public readonly bool EnablePassChange;
        /// <summary>
        /// Enable Rewrites
        /// </summary>
        public readonly bool EnableRewrites;
        /// <summary>
        /// Number of bad password attempts
        /// </summary>
        public readonly int FailedAuthThreshold;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// List of AD Groups
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActiveDirectoryJoinPointGroupResult> Groups;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Identity Not In AD Behaviour
        /// </summary>
        public readonly string IdentityNotInAdBehaviour;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string JobTitle;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string Locality;
        /// <summary>
        /// The name of the active directory join point
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string OrganizationalUnit;
        /// <summary>
        /// Plain Text Authentication
        /// </summary>
        public readonly bool PlaintextAuth;
        /// <summary>
        /// List of Rewrite rules
        /// </summary>
        public readonly ImmutableArray<Outputs.GetActiveDirectoryJoinPointRewriteRuleResult> RewriteRules;
        /// <summary>
        /// Schema
        /// </summary>
        public readonly string Schema;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string StateOrProvince;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string StreetAddress;
        /// <summary>
        /// User info attribute
        /// </summary>
        public readonly string Telephone;
        /// <summary>
        /// Unreachable Domains Behaviour
        /// </summary>
        public readonly string UnreachableDomainsBehaviour;

        [OutputConstructor]
        private GetActiveDirectoryJoinPointResult(
            string adScopesNames,

            int agingTime,

            ImmutableArray<Outputs.GetActiveDirectoryJoinPointAttributeResult> attributes,

            string authProtectionType,

            string country,

            string department,

            string description,

            string domain,

            string email,

            bool enableCallbackForDialinClient,

            bool enableDialinPermissionCheck,

            bool enableDomainAllowedList,

            bool enableFailedAuthProtection,

            bool enableMachineAccess,

            bool enableMachineAuth,

            bool enablePassChange,

            bool enableRewrites,

            int failedAuthThreshold,

            string firstName,

            ImmutableArray<Outputs.GetActiveDirectoryJoinPointGroupResult> groups,

            string id,

            string identityNotInAdBehaviour,

            string jobTitle,

            string lastName,

            string locality,

            string name,

            string organizationalUnit,

            bool plaintextAuth,

            ImmutableArray<Outputs.GetActiveDirectoryJoinPointRewriteRuleResult> rewriteRules,

            string schema,

            string stateOrProvince,

            string streetAddress,

            string telephone,

            string unreachableDomainsBehaviour)
        {
            AdScopesNames = adScopesNames;
            AgingTime = agingTime;
            Attributes = attributes;
            AuthProtectionType = authProtectionType;
            Country = country;
            Department = department;
            Description = description;
            Domain = domain;
            Email = email;
            EnableCallbackForDialinClient = enableCallbackForDialinClient;
            EnableDialinPermissionCheck = enableDialinPermissionCheck;
            EnableDomainAllowedList = enableDomainAllowedList;
            EnableFailedAuthProtection = enableFailedAuthProtection;
            EnableMachineAccess = enableMachineAccess;
            EnableMachineAuth = enableMachineAuth;
            EnablePassChange = enablePassChange;
            EnableRewrites = enableRewrites;
            FailedAuthThreshold = failedAuthThreshold;
            FirstName = firstName;
            Groups = groups;
            Id = id;
            IdentityNotInAdBehaviour = identityNotInAdBehaviour;
            JobTitle = jobTitle;
            LastName = lastName;
            Locality = locality;
            Name = name;
            OrganizationalUnit = organizationalUnit;
            PlaintextAuth = plaintextAuth;
            RewriteRules = rewriteRules;
            Schema = schema;
            StateOrProvince = stateOrProvince;
            StreetAddress = streetAddress;
            Telephone = telephone;
            UnreachableDomainsBehaviour = unreachableDomainsBehaviour;
        }
    }
}

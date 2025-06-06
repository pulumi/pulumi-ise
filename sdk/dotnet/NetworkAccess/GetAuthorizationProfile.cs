// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.NetworkAccess
{
    public static class GetAuthorizationProfile
    {
        /// <summary>
        /// This data source can read an authorization profiles policy element.
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
        ///     var example = Ise.NetworkAccess.GetAuthorizationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAuthorizationProfileResult> InvokeAsync(GetAuthorizationProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthorizationProfileResult>("ise:networkaccess/getAuthorizationProfile:getAuthorizationProfile", args ?? new GetAuthorizationProfileArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an authorization profiles policy element.
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
        ///     var example = Ise.NetworkAccess.GetAuthorizationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthorizationProfileResult> Invoke(GetAuthorizationProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationProfileResult>("ise:networkaccess/getAuthorizationProfile:getAuthorizationProfile", args ?? new GetAuthorizationProfileInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an authorization profiles policy element.
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
        ///     var example = Ise.NetworkAccess.GetAuthorizationProfile.Invoke(new()
        ///     {
        ///         Id = "76d24097-41c4-4558-a4d0-a8c07ac08470",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAuthorizationProfileResult> Invoke(GetAuthorizationProfileInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthorizationProfileResult>("ise:networkaccess/getAuthorizationProfile:getAuthorizationProfile", args ?? new GetAuthorizationProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorizationProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetAuthorizationProfileArgs()
        {
        }
        public static new GetAuthorizationProfileArgs Empty => new GetAuthorizationProfileArgs();
    }

    public sealed class GetAuthorizationProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the object
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetAuthorizationProfileInvokeArgs()
        {
        }
        public static new GetAuthorizationProfileInvokeArgs Empty => new GetAuthorizationProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthorizationProfileResult
    {
        /// <summary>
        /// Access type
        /// </summary>
        public readonly string AccessType;
        /// <summary>
        /// ACL
        /// </summary>
        public readonly string Acl;
        /// <summary>
        /// List of advanced attributes
        /// </summary>
        public readonly ImmutableArray<Outputs.GetAuthorizationProfileAdvancedAttributeResult> AdvancedAttributes;
        /// <summary>
        /// Agentless Posture.
        /// </summary>
        public readonly bool AgentlessPosture;
        /// <summary>
        /// Airespace ACL
        /// </summary>
        public readonly string AirespaceAcl;
        /// <summary>
        /// Airespace IPv6 ACL
        /// </summary>
        public readonly string AirespaceIpv6Acl;
        /// <summary>
        /// ASA VPN
        /// </summary>
        public readonly string AsaVpn;
        /// <summary>
        /// Auto smart port
        /// </summary>
        public readonly string AutoSmartPort;
        /// <summary>
        /// AVC profile
        /// </summary>
        public readonly string AvcProfile;
        /// <summary>
        /// DACL name
        /// </summary>
        public readonly string DaclName;
        /// <summary>
        /// Description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Easy wired session candidate
        /// </summary>
        public readonly bool EasywiredSessionCandidate;
        /// <summary>
        /// The id of the object
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Interface template
        /// </summary>
        public readonly string InterfaceTemplate;
        /// <summary>
        /// IPv6 ACL
        /// </summary>
        public readonly string Ipv6AclFilter;
        /// <summary>
        /// IPv6 DACL name
        /// </summary>
        public readonly string Ipv6DaclName;
        /// <summary>
        /// MacSec policy
        /// </summary>
        public readonly string MacSecPolicy;
        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// NEAT
        /// </summary>
        public readonly bool Neat;
        /// <summary>
        /// Value needs to be an existing Network Device Profile
        /// </summary>
        public readonly string ProfileName;
        /// <summary>
        /// Maintain Connectivity During Reauthentication
        /// </summary>
        public readonly string ReauthenticationConnectivity;
        /// <summary>
        /// Reauthentication timer
        /// </summary>
        public readonly int ReauthenticationTimer;
        /// <summary>
        /// Service template
        /// </summary>
        public readonly bool ServiceTemplate;
        /// <summary>
        /// Track movement
        /// </summary>
        public readonly bool TrackMovement;
        /// <summary>
        /// Unique identifier
        /// </summary>
        public readonly string UniqueIdentifier;
        /// <summary>
        /// Vlan name or ID
        /// </summary>
        public readonly string VlanNameId;
        /// <summary>
        /// Vlan tag ID
        /// </summary>
        public readonly int VlanTagId;
        /// <summary>
        /// Voice domain permission
        /// </summary>
        public readonly bool VoiceDomainPermission;
        /// <summary>
        /// Web authentication (local)
        /// </summary>
        public readonly bool WebAuth;
        /// <summary>
        /// Web redirection ACL
        /// </summary>
        public readonly string WebRedirectionAcl;
        /// <summary>
        /// This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other `web_redirection_type` values the field must be ignored.
        /// </summary>
        public readonly bool WebRedirectionDisplayCertificatesRenewalMessages;
        /// <summary>
        /// A portal that exist in the DB and fits the `web_redirection_type`
        /// </summary>
        public readonly string WebRedirectionPortalName;
        /// <summary>
        /// IP, hostname or FQDN
        /// </summary>
        public readonly string WebRedirectionStaticIpHostNameFqdn;
        /// <summary>
        /// This type must fit the `web_redirection_portal_name`
        /// </summary>
        public readonly string WebRedirectionType;

        [OutputConstructor]
        private GetAuthorizationProfileResult(
            string accessType,

            string acl,

            ImmutableArray<Outputs.GetAuthorizationProfileAdvancedAttributeResult> advancedAttributes,

            bool agentlessPosture,

            string airespaceAcl,

            string airespaceIpv6Acl,

            string asaVpn,

            string autoSmartPort,

            string avcProfile,

            string daclName,

            string description,

            bool easywiredSessionCandidate,

            string id,

            string interfaceTemplate,

            string ipv6AclFilter,

            string ipv6DaclName,

            string macSecPolicy,

            string name,

            bool neat,

            string profileName,

            string reauthenticationConnectivity,

            int reauthenticationTimer,

            bool serviceTemplate,

            bool trackMovement,

            string uniqueIdentifier,

            string vlanNameId,

            int vlanTagId,

            bool voiceDomainPermission,

            bool webAuth,

            string webRedirectionAcl,

            bool webRedirectionDisplayCertificatesRenewalMessages,

            string webRedirectionPortalName,

            string webRedirectionStaticIpHostNameFqdn,

            string webRedirectionType)
        {
            AccessType = accessType;
            Acl = acl;
            AdvancedAttributes = advancedAttributes;
            AgentlessPosture = agentlessPosture;
            AirespaceAcl = airespaceAcl;
            AirespaceIpv6Acl = airespaceIpv6Acl;
            AsaVpn = asaVpn;
            AutoSmartPort = autoSmartPort;
            AvcProfile = avcProfile;
            DaclName = daclName;
            Description = description;
            EasywiredSessionCandidate = easywiredSessionCandidate;
            Id = id;
            InterfaceTemplate = interfaceTemplate;
            Ipv6AclFilter = ipv6AclFilter;
            Ipv6DaclName = ipv6DaclName;
            MacSecPolicy = macSecPolicy;
            Name = name;
            Neat = neat;
            ProfileName = profileName;
            ReauthenticationConnectivity = reauthenticationConnectivity;
            ReauthenticationTimer = reauthenticationTimer;
            ServiceTemplate = serviceTemplate;
            TrackMovement = trackMovement;
            UniqueIdentifier = uniqueIdentifier;
            VlanNameId = vlanNameId;
            VlanTagId = vlanTagId;
            VoiceDomainPermission = voiceDomainPermission;
            WebAuth = webAuth;
            WebRedirectionAcl = webRedirectionAcl;
            WebRedirectionDisplayCertificatesRenewalMessages = webRedirectionDisplayCertificatesRenewalMessages;
            WebRedirectionPortalName = webRedirectionPortalName;
            WebRedirectionStaticIpHostNameFqdn = webRedirectionStaticIpHostNameFqdn;
            WebRedirectionType = webRedirectionType;
        }
    }
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ise.Authorization
{
    /// <summary>
    /// This resource can manage an authorization profiles policy element.
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import ise:Authorization/profile:Profile example "76d24097-41c4-4558-a4d0-a8c07ac08470"
    /// ```
    /// </summary>
    [IseResourceType("ise:Authorization/profile:Profile")]
    public partial class Profile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access type - Choices: `ACCESS_ACCEPT`, `ACCESS_REJECT` - Default value: `ACCESS_ACCEPT`
        /// </summary>
        [Output("accessType")]
        public Output<string> AccessType { get; private set; } = null!;

        /// <summary>
        /// ACL
        /// </summary>
        [Output("acl")]
        public Output<string?> Acl { get; private set; } = null!;

        /// <summary>
        /// List of advanced attributes
        /// </summary>
        [Output("advancedAttributes")]
        public Output<ImmutableArray<Outputs.ProfileAdvancedAttribute>> AdvancedAttributes { get; private set; } = null!;

        /// <summary>
        /// Agentless Posture.
        /// </summary>
        [Output("agentlessPosture")]
        public Output<bool?> AgentlessPosture { get; private set; } = null!;

        /// <summary>
        /// Airespace ACL
        /// </summary>
        [Output("airespaceAcl")]
        public Output<string?> AirespaceAcl { get; private set; } = null!;

        /// <summary>
        /// Airespace IPv6 ACL
        /// </summary>
        [Output("airespaceIpv6Acl")]
        public Output<string?> AirespaceIpv6Acl { get; private set; } = null!;

        /// <summary>
        /// ASA VPN
        /// </summary>
        [Output("asaVpn")]
        public Output<string?> AsaVpn { get; private set; } = null!;

        /// <summary>
        /// Auto smart port
        /// </summary>
        [Output("autoSmartPort")]
        public Output<string?> AutoSmartPort { get; private set; } = null!;

        /// <summary>
        /// AVC profile
        /// </summary>
        [Output("avcProfile")]
        public Output<string?> AvcProfile { get; private set; } = null!;

        /// <summary>
        /// DACL name
        /// </summary>
        [Output("daclName")]
        public Output<string?> DaclName { get; private set; } = null!;

        /// <summary>
        /// Description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Easy wired session candidate - Default value: `false`
        /// </summary>
        [Output("easywiredSessionCandidate")]
        public Output<bool> EasywiredSessionCandidate { get; private set; } = null!;

        /// <summary>
        /// Interface template
        /// </summary>
        [Output("interfaceTemplate")]
        public Output<string?> InterfaceTemplate { get; private set; } = null!;

        /// <summary>
        /// IPv6 ACL
        /// </summary>
        [Output("ipv6AclFilter")]
        public Output<string?> Ipv6AclFilter { get; private set; } = null!;

        /// <summary>
        /// IPv6 DACL name
        /// </summary>
        [Output("ipv6DaclName")]
        public Output<string?> Ipv6DaclName { get; private set; } = null!;

        /// <summary>
        /// MacSec policy - Choices: `MUST_SECURE`, `MUST_NOT_SECURE`, `SHOULD_SECURE`
        /// </summary>
        [Output("macSecPolicy")]
        public Output<string?> MacSecPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// NEAT - Default value: `false`
        /// </summary>
        [Output("neat")]
        public Output<bool> Neat { get; private set; } = null!;

        /// <summary>
        /// Value needs to be an existing Network Device Profile - Default value: `Cisco`
        /// </summary>
        [Output("profileName")]
        public Output<string> ProfileName { get; private set; } = null!;

        /// <summary>
        /// Maintain Connectivity During Reauthentication - Choices: `DEFAULT`, `RADIUS_REQUEST`
        /// </summary>
        [Output("reauthenticationConnectivity")]
        public Output<string?> ReauthenticationConnectivity { get; private set; } = null!;

        /// <summary>
        /// Reauthentication timer - Range: `1`-`65535`
        /// </summary>
        [Output("reauthenticationTimer")]
        public Output<int?> ReauthenticationTimer { get; private set; } = null!;

        /// <summary>
        /// Service template - Default value: `false`
        /// </summary>
        [Output("serviceTemplate")]
        public Output<bool> ServiceTemplate { get; private set; } = null!;

        /// <summary>
        /// Track movement - Default value: `false`
        /// </summary>
        [Output("trackMovement")]
        public Output<bool> TrackMovement { get; private set; } = null!;

        /// <summary>
        /// Unique identifier
        /// </summary>
        [Output("uniqueIdentifier")]
        public Output<string?> UniqueIdentifier { get; private set; } = null!;

        /// <summary>
        /// Vlan name or ID
        /// </summary>
        [Output("vlanNameId")]
        public Output<string?> VlanNameId { get; private set; } = null!;

        /// <summary>
        /// Vlan tag ID - Range: `0`-`31`
        /// </summary>
        [Output("vlanTagId")]
        public Output<int?> VlanTagId { get; private set; } = null!;

        /// <summary>
        /// Voice domain permission - Default value: `false`
        /// </summary>
        [Output("voiceDomainPermission")]
        public Output<bool> VoiceDomainPermission { get; private set; } = null!;

        /// <summary>
        /// Web authentication (local) - Default value: `false`
        /// </summary>
        [Output("webAuth")]
        public Output<bool> WebAuth { get; private set; } = null!;

        /// <summary>
        /// Web redirection ACL
        /// </summary>
        [Output("webRedirectionAcl")]
        public Output<string?> WebRedirectionAcl { get; private set; } = null!;

        /// <summary>
        /// This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other
        /// `web_redirection_type` values the field must be ignored.
        /// </summary>
        [Output("webRedirectionDisplayCertificatesRenewalMessages")]
        public Output<bool?> WebRedirectionDisplayCertificatesRenewalMessages { get; private set; } = null!;

        /// <summary>
        /// A portal that exist in the DB and fits the `web_redirection_type`
        /// </summary>
        [Output("webRedirectionPortalName")]
        public Output<string?> WebRedirectionPortalName { get; private set; } = null!;

        /// <summary>
        /// IP, hostname or FQDN
        /// </summary>
        [Output("webRedirectionStaticIpHostNameFqdn")]
        public Output<string?> WebRedirectionStaticIpHostNameFqdn { get; private set; } = null!;

        /// <summary>
        /// This type must fit the `web_redirection_portal_name` - Choices: `CentralizedWebAuth`, `HotSpot`,
        /// `NativeSupplicanProvisioning`, `ClientProvisioning`
        /// </summary>
        [Output("webRedirectionType")]
        public Output<string?> WebRedirectionType { get; private set; } = null!;


        /// <summary>
        /// Create a Profile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Profile(string name, ProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("ise:Authorization/profile:Profile", name, args ?? new ProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Profile(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
            : base("ise:Authorization/profile:Profile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Profile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Profile Get(string name, Input<string> id, ProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new Profile(name, id, state, options);
        }
    }

    public sealed class ProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access type - Choices: `ACCESS_ACCEPT`, `ACCESS_REJECT` - Default value: `ACCESS_ACCEPT`
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// ACL
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("advancedAttributes")]
        private InputList<Inputs.ProfileAdvancedAttributeArgs>? _advancedAttributes;

        /// <summary>
        /// List of advanced attributes
        /// </summary>
        public InputList<Inputs.ProfileAdvancedAttributeArgs> AdvancedAttributes
        {
            get => _advancedAttributes ?? (_advancedAttributes = new InputList<Inputs.ProfileAdvancedAttributeArgs>());
            set => _advancedAttributes = value;
        }

        /// <summary>
        /// Agentless Posture.
        /// </summary>
        [Input("agentlessPosture")]
        public Input<bool>? AgentlessPosture { get; set; }

        /// <summary>
        /// Airespace ACL
        /// </summary>
        [Input("airespaceAcl")]
        public Input<string>? AirespaceAcl { get; set; }

        /// <summary>
        /// Airespace IPv6 ACL
        /// </summary>
        [Input("airespaceIpv6Acl")]
        public Input<string>? AirespaceIpv6Acl { get; set; }

        /// <summary>
        /// ASA VPN
        /// </summary>
        [Input("asaVpn")]
        public Input<string>? AsaVpn { get; set; }

        /// <summary>
        /// Auto smart port
        /// </summary>
        [Input("autoSmartPort")]
        public Input<string>? AutoSmartPort { get; set; }

        /// <summary>
        /// AVC profile
        /// </summary>
        [Input("avcProfile")]
        public Input<string>? AvcProfile { get; set; }

        /// <summary>
        /// DACL name
        /// </summary>
        [Input("daclName")]
        public Input<string>? DaclName { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Easy wired session candidate - Default value: `false`
        /// </summary>
        [Input("easywiredSessionCandidate")]
        public Input<bool>? EasywiredSessionCandidate { get; set; }

        /// <summary>
        /// Interface template
        /// </summary>
        [Input("interfaceTemplate")]
        public Input<string>? InterfaceTemplate { get; set; }

        /// <summary>
        /// IPv6 ACL
        /// </summary>
        [Input("ipv6AclFilter")]
        public Input<string>? Ipv6AclFilter { get; set; }

        /// <summary>
        /// IPv6 DACL name
        /// </summary>
        [Input("ipv6DaclName")]
        public Input<string>? Ipv6DaclName { get; set; }

        /// <summary>
        /// MacSec policy - Choices: `MUST_SECURE`, `MUST_NOT_SECURE`, `SHOULD_SECURE`
        /// </summary>
        [Input("macSecPolicy")]
        public Input<string>? MacSecPolicy { get; set; }

        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NEAT - Default value: `false`
        /// </summary>
        [Input("neat")]
        public Input<bool>? Neat { get; set; }

        /// <summary>
        /// Value needs to be an existing Network Device Profile - Default value: `Cisco`
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        /// <summary>
        /// Maintain Connectivity During Reauthentication - Choices: `DEFAULT`, `RADIUS_REQUEST`
        /// </summary>
        [Input("reauthenticationConnectivity")]
        public Input<string>? ReauthenticationConnectivity { get; set; }

        /// <summary>
        /// Reauthentication timer - Range: `1`-`65535`
        /// </summary>
        [Input("reauthenticationTimer")]
        public Input<int>? ReauthenticationTimer { get; set; }

        /// <summary>
        /// Service template - Default value: `false`
        /// </summary>
        [Input("serviceTemplate")]
        public Input<bool>? ServiceTemplate { get; set; }

        /// <summary>
        /// Track movement - Default value: `false`
        /// </summary>
        [Input("trackMovement")]
        public Input<bool>? TrackMovement { get; set; }

        /// <summary>
        /// Unique identifier
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        /// <summary>
        /// Vlan name or ID
        /// </summary>
        [Input("vlanNameId")]
        public Input<string>? VlanNameId { get; set; }

        /// <summary>
        /// Vlan tag ID - Range: `0`-`31`
        /// </summary>
        [Input("vlanTagId")]
        public Input<int>? VlanTagId { get; set; }

        /// <summary>
        /// Voice domain permission - Default value: `false`
        /// </summary>
        [Input("voiceDomainPermission")]
        public Input<bool>? VoiceDomainPermission { get; set; }

        /// <summary>
        /// Web authentication (local) - Default value: `false`
        /// </summary>
        [Input("webAuth")]
        public Input<bool>? WebAuth { get; set; }

        /// <summary>
        /// Web redirection ACL
        /// </summary>
        [Input("webRedirectionAcl")]
        public Input<string>? WebRedirectionAcl { get; set; }

        /// <summary>
        /// This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other
        /// `web_redirection_type` values the field must be ignored.
        /// </summary>
        [Input("webRedirectionDisplayCertificatesRenewalMessages")]
        public Input<bool>? WebRedirectionDisplayCertificatesRenewalMessages { get; set; }

        /// <summary>
        /// A portal that exist in the DB and fits the `web_redirection_type`
        /// </summary>
        [Input("webRedirectionPortalName")]
        public Input<string>? WebRedirectionPortalName { get; set; }

        /// <summary>
        /// IP, hostname or FQDN
        /// </summary>
        [Input("webRedirectionStaticIpHostNameFqdn")]
        public Input<string>? WebRedirectionStaticIpHostNameFqdn { get; set; }

        /// <summary>
        /// This type must fit the `web_redirection_portal_name` - Choices: `CentralizedWebAuth`, `HotSpot`,
        /// `NativeSupplicanProvisioning`, `ClientProvisioning`
        /// </summary>
        [Input("webRedirectionType")]
        public Input<string>? WebRedirectionType { get; set; }

        public ProfileArgs()
        {
        }
        public static new ProfileArgs Empty => new ProfileArgs();
    }

    public sealed class ProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access type - Choices: `ACCESS_ACCEPT`, `ACCESS_REJECT` - Default value: `ACCESS_ACCEPT`
        /// </summary>
        [Input("accessType")]
        public Input<string>? AccessType { get; set; }

        /// <summary>
        /// ACL
        /// </summary>
        [Input("acl")]
        public Input<string>? Acl { get; set; }

        [Input("advancedAttributes")]
        private InputList<Inputs.ProfileAdvancedAttributeGetArgs>? _advancedAttributes;

        /// <summary>
        /// List of advanced attributes
        /// </summary>
        public InputList<Inputs.ProfileAdvancedAttributeGetArgs> AdvancedAttributes
        {
            get => _advancedAttributes ?? (_advancedAttributes = new InputList<Inputs.ProfileAdvancedAttributeGetArgs>());
            set => _advancedAttributes = value;
        }

        /// <summary>
        /// Agentless Posture.
        /// </summary>
        [Input("agentlessPosture")]
        public Input<bool>? AgentlessPosture { get; set; }

        /// <summary>
        /// Airespace ACL
        /// </summary>
        [Input("airespaceAcl")]
        public Input<string>? AirespaceAcl { get; set; }

        /// <summary>
        /// Airespace IPv6 ACL
        /// </summary>
        [Input("airespaceIpv6Acl")]
        public Input<string>? AirespaceIpv6Acl { get; set; }

        /// <summary>
        /// ASA VPN
        /// </summary>
        [Input("asaVpn")]
        public Input<string>? AsaVpn { get; set; }

        /// <summary>
        /// Auto smart port
        /// </summary>
        [Input("autoSmartPort")]
        public Input<string>? AutoSmartPort { get; set; }

        /// <summary>
        /// AVC profile
        /// </summary>
        [Input("avcProfile")]
        public Input<string>? AvcProfile { get; set; }

        /// <summary>
        /// DACL name
        /// </summary>
        [Input("daclName")]
        public Input<string>? DaclName { get; set; }

        /// <summary>
        /// Description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Easy wired session candidate - Default value: `false`
        /// </summary>
        [Input("easywiredSessionCandidate")]
        public Input<bool>? EasywiredSessionCandidate { get; set; }

        /// <summary>
        /// Interface template
        /// </summary>
        [Input("interfaceTemplate")]
        public Input<string>? InterfaceTemplate { get; set; }

        /// <summary>
        /// IPv6 ACL
        /// </summary>
        [Input("ipv6AclFilter")]
        public Input<string>? Ipv6AclFilter { get; set; }

        /// <summary>
        /// IPv6 DACL name
        /// </summary>
        [Input("ipv6DaclName")]
        public Input<string>? Ipv6DaclName { get; set; }

        /// <summary>
        /// MacSec policy - Choices: `MUST_SECURE`, `MUST_NOT_SECURE`, `SHOULD_SECURE`
        /// </summary>
        [Input("macSecPolicy")]
        public Input<string>? MacSecPolicy { get; set; }

        /// <summary>
        /// The name of the authorization profile
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// NEAT - Default value: `false`
        /// </summary>
        [Input("neat")]
        public Input<bool>? Neat { get; set; }

        /// <summary>
        /// Value needs to be an existing Network Device Profile - Default value: `Cisco`
        /// </summary>
        [Input("profileName")]
        public Input<string>? ProfileName { get; set; }

        /// <summary>
        /// Maintain Connectivity During Reauthentication - Choices: `DEFAULT`, `RADIUS_REQUEST`
        /// </summary>
        [Input("reauthenticationConnectivity")]
        public Input<string>? ReauthenticationConnectivity { get; set; }

        /// <summary>
        /// Reauthentication timer - Range: `1`-`65535`
        /// </summary>
        [Input("reauthenticationTimer")]
        public Input<int>? ReauthenticationTimer { get; set; }

        /// <summary>
        /// Service template - Default value: `false`
        /// </summary>
        [Input("serviceTemplate")]
        public Input<bool>? ServiceTemplate { get; set; }

        /// <summary>
        /// Track movement - Default value: `false`
        /// </summary>
        [Input("trackMovement")]
        public Input<bool>? TrackMovement { get; set; }

        /// <summary>
        /// Unique identifier
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        /// <summary>
        /// Vlan name or ID
        /// </summary>
        [Input("vlanNameId")]
        public Input<string>? VlanNameId { get; set; }

        /// <summary>
        /// Vlan tag ID - Range: `0`-`31`
        /// </summary>
        [Input("vlanTagId")]
        public Input<int>? VlanTagId { get; set; }

        /// <summary>
        /// Voice domain permission - Default value: `false`
        /// </summary>
        [Input("voiceDomainPermission")]
        public Input<bool>? VoiceDomainPermission { get; set; }

        /// <summary>
        /// Web authentication (local) - Default value: `false`
        /// </summary>
        [Input("webAuth")]
        public Input<bool>? WebAuth { get; set; }

        /// <summary>
        /// Web redirection ACL
        /// </summary>
        [Input("webRedirectionAcl")]
        public Input<string>? WebRedirectionAcl { get; set; }

        /// <summary>
        /// This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other
        /// `web_redirection_type` values the field must be ignored.
        /// </summary>
        [Input("webRedirectionDisplayCertificatesRenewalMessages")]
        public Input<bool>? WebRedirectionDisplayCertificatesRenewalMessages { get; set; }

        /// <summary>
        /// A portal that exist in the DB and fits the `web_redirection_type`
        /// </summary>
        [Input("webRedirectionPortalName")]
        public Input<string>? WebRedirectionPortalName { get; set; }

        /// <summary>
        /// IP, hostname or FQDN
        /// </summary>
        [Input("webRedirectionStaticIpHostNameFqdn")]
        public Input<string>? WebRedirectionStaticIpHostNameFqdn { get; set; }

        /// <summary>
        /// This type must fit the `web_redirection_portal_name` - Choices: `CentralizedWebAuth`, `HotSpot`,
        /// `NativeSupplicanProvisioning`, `ClientProvisioning`
        /// </summary>
        [Input("webRedirectionType")]
        public Input<string>? WebRedirectionType { get; set; }

        public ProfileState()
        {
        }
        public static new ProfileState Empty => new ProfileState();
    }
}

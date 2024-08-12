// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.networkaccess.AuthorizationProfileArgs;
import com.pulumi.ise.networkaccess.inputs.AuthorizationProfileState;
import com.pulumi.ise.networkaccess.outputs.AuthorizationProfileAdvancedAttribute;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage an authorization profiles policy element.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ise.networkaccess.AuthorizationProfile;
 * import com.pulumi.ise.networkaccess.AuthorizationProfileArgs;
 * import com.pulumi.ise.networkaccess.inputs.AuthorizationProfileAdvancedAttributeArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new AuthorizationProfile("example", AuthorizationProfileArgs.builder()
 *             .name("AuthzProfile1")
 *             .description("My Authorization Profile")
 *             .vlanNameId("VLAN10")
 *             .vlanTagId(0)
 *             .webRedirectionType("CentralizedWebAuth")
 *             .webRedirectionAcl("TEST_ACL")
 *             .webRedirectionPortalName("Sponsored Guest Portal (default)")
 *             .webRedirectionStaticIpHostNameFqdn("1.2.3.4")
 *             .webRedirectionDisplayCertificatesRenewalMessages(true)
 *             .agentlessPosture(false)
 *             .accessType("ACCESS_ACCEPT")
 *             .profileName("Cisco")
 *             .airespaceAcl("ACL1")
 *             .acl("ACL1")
 *             .autoSmartPort("PROFILE1")
 *             .interfaceTemplate("TEMP1")
 *             .ipv6AclFilter("ACL1")
 *             .avcProfile("PROF1")
 *             .asaVpn("1")
 *             .uniqueIdentifier("ID1234")
 *             .trackMovement(false)
 *             .serviceTemplate(false)
 *             .easywiredSessionCandidate(false)
 *             .voiceDomainPermission(false)
 *             .neat(false)
 *             .webAuth(false)
 *             .macSecPolicy("MUST_SECURE")
 *             .reauthenticationConnectivity("DEFAULT")
 *             .reauthenticationTimer(1)
 *             .advancedAttributes(AuthorizationProfileAdvancedAttributeArgs.builder()
 *                 .attribute_left_dictionary_name("Cisco")
 *                 .attribute_left_name("cisco-av-pair")
 *                 .attribute_right_value_type("AttributeValue")
 *                 .attribute_right_value("set_nadprofile_vlan=true,vlan=TEST,tag=1")
 *                 .build())
 *             .airespaceIpv6Acl("ACL1")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ise:networkaccess/authorizationProfile:AuthorizationProfile example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:networkaccess/authorizationProfile:AuthorizationProfile")
public class AuthorizationProfile extends com.pulumi.resources.CustomResource {
    /**
     * Access type - Choices: `ACCESS_ACCEPT`, `ACCESS_REJECT` - Default value: `ACCESS_ACCEPT`
     * 
     */
    @Export(name="accessType", refs={String.class}, tree="[0]")
    private Output<String> accessType;

    /**
     * @return Access type - Choices: `ACCESS_ACCEPT`, `ACCESS_REJECT` - Default value: `ACCESS_ACCEPT`
     * 
     */
    public Output<String> accessType() {
        return this.accessType;
    }
    /**
     * ACL
     * 
     */
    @Export(name="acl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> acl;

    /**
     * @return ACL
     * 
     */
    public Output<Optional<String>> acl() {
        return Codegen.optional(this.acl);
    }
    /**
     * List of advanced attributes
     * 
     */
    @Export(name="advancedAttributes", refs={List.class,AuthorizationProfileAdvancedAttribute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AuthorizationProfileAdvancedAttribute>> advancedAttributes;

    /**
     * @return List of advanced attributes
     * 
     */
    public Output<Optional<List<AuthorizationProfileAdvancedAttribute>>> advancedAttributes() {
        return Codegen.optional(this.advancedAttributes);
    }
    /**
     * Agentless Posture.
     * 
     */
    @Export(name="agentlessPosture", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> agentlessPosture;

    /**
     * @return Agentless Posture.
     * 
     */
    public Output<Optional<Boolean>> agentlessPosture() {
        return Codegen.optional(this.agentlessPosture);
    }
    /**
     * Airespace ACL
     * 
     */
    @Export(name="airespaceAcl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> airespaceAcl;

    /**
     * @return Airespace ACL
     * 
     */
    public Output<Optional<String>> airespaceAcl() {
        return Codegen.optional(this.airespaceAcl);
    }
    /**
     * Airespace IPv6 ACL
     * 
     */
    @Export(name="airespaceIpv6Acl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> airespaceIpv6Acl;

    /**
     * @return Airespace IPv6 ACL
     * 
     */
    public Output<Optional<String>> airespaceIpv6Acl() {
        return Codegen.optional(this.airespaceIpv6Acl);
    }
    /**
     * ASA VPN
     * 
     */
    @Export(name="asaVpn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> asaVpn;

    /**
     * @return ASA VPN
     * 
     */
    public Output<Optional<String>> asaVpn() {
        return Codegen.optional(this.asaVpn);
    }
    /**
     * Auto smart port
     * 
     */
    @Export(name="autoSmartPort", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoSmartPort;

    /**
     * @return Auto smart port
     * 
     */
    public Output<Optional<String>> autoSmartPort() {
        return Codegen.optional(this.autoSmartPort);
    }
    /**
     * AVC profile
     * 
     */
    @Export(name="avcProfile", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> avcProfile;

    /**
     * @return AVC profile
     * 
     */
    public Output<Optional<String>> avcProfile() {
        return Codegen.optional(this.avcProfile);
    }
    /**
     * DACL name
     * 
     */
    @Export(name="daclName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> daclName;

    /**
     * @return DACL name
     * 
     */
    public Output<Optional<String>> daclName() {
        return Codegen.optional(this.daclName);
    }
    /**
     * Description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Easy wired session candidate - Default value: `false`
     * 
     */
    @Export(name="easywiredSessionCandidate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> easywiredSessionCandidate;

    /**
     * @return Easy wired session candidate - Default value: `false`
     * 
     */
    public Output<Boolean> easywiredSessionCandidate() {
        return this.easywiredSessionCandidate;
    }
    /**
     * Interface template
     * 
     */
    @Export(name="interfaceTemplate", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> interfaceTemplate;

    /**
     * @return Interface template
     * 
     */
    public Output<Optional<String>> interfaceTemplate() {
        return Codegen.optional(this.interfaceTemplate);
    }
    /**
     * IPv6 ACL
     * 
     */
    @Export(name="ipv6AclFilter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipv6AclFilter;

    /**
     * @return IPv6 ACL
     * 
     */
    public Output<Optional<String>> ipv6AclFilter() {
        return Codegen.optional(this.ipv6AclFilter);
    }
    /**
     * IPv6 DACL name
     * 
     */
    @Export(name="ipv6DaclName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipv6DaclName;

    /**
     * @return IPv6 DACL name
     * 
     */
    public Output<Optional<String>> ipv6DaclName() {
        return Codegen.optional(this.ipv6DaclName);
    }
    /**
     * MacSec policy - Choices: `MUST_SECURE`, `MUST_NOT_SECURE`, `SHOULD_SECURE`
     * 
     */
    @Export(name="macSecPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> macSecPolicy;

    /**
     * @return MacSec policy - Choices: `MUST_SECURE`, `MUST_NOT_SECURE`, `SHOULD_SECURE`
     * 
     */
    public Output<Optional<String>> macSecPolicy() {
        return Codegen.optional(this.macSecPolicy);
    }
    /**
     * The name of the authorization profile
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the authorization profile
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * NEAT - Default value: `false`
     * 
     */
    @Export(name="neat", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> neat;

    /**
     * @return NEAT - Default value: `false`
     * 
     */
    public Output<Boolean> neat() {
        return this.neat;
    }
    /**
     * Value needs to be an existing Network Device Profile - Default value: `Cisco`
     * 
     */
    @Export(name="profileName", refs={String.class}, tree="[0]")
    private Output<String> profileName;

    /**
     * @return Value needs to be an existing Network Device Profile - Default value: `Cisco`
     * 
     */
    public Output<String> profileName() {
        return this.profileName;
    }
    /**
     * Maintain Connectivity During Reauthentication - Choices: `DEFAULT`, `RADIUS_REQUEST`
     * 
     */
    @Export(name="reauthenticationConnectivity", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> reauthenticationConnectivity;

    /**
     * @return Maintain Connectivity During Reauthentication - Choices: `DEFAULT`, `RADIUS_REQUEST`
     * 
     */
    public Output<Optional<String>> reauthenticationConnectivity() {
        return Codegen.optional(this.reauthenticationConnectivity);
    }
    /**
     * Reauthentication timer - Range: `1`-`65535`
     * 
     */
    @Export(name="reauthenticationTimer", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> reauthenticationTimer;

    /**
     * @return Reauthentication timer - Range: `1`-`65535`
     * 
     */
    public Output<Optional<Integer>> reauthenticationTimer() {
        return Codegen.optional(this.reauthenticationTimer);
    }
    /**
     * Service template - Default value: `false`
     * 
     */
    @Export(name="serviceTemplate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> serviceTemplate;

    /**
     * @return Service template - Default value: `false`
     * 
     */
    public Output<Boolean> serviceTemplate() {
        return this.serviceTemplate;
    }
    /**
     * Track movement - Default value: `false`
     * 
     */
    @Export(name="trackMovement", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> trackMovement;

    /**
     * @return Track movement - Default value: `false`
     * 
     */
    public Output<Boolean> trackMovement() {
        return this.trackMovement;
    }
    /**
     * Unique identifier
     * 
     */
    @Export(name="uniqueIdentifier", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> uniqueIdentifier;

    /**
     * @return Unique identifier
     * 
     */
    public Output<Optional<String>> uniqueIdentifier() {
        return Codegen.optional(this.uniqueIdentifier);
    }
    /**
     * Vlan name or ID
     * 
     */
    @Export(name="vlanNameId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vlanNameId;

    /**
     * @return Vlan name or ID
     * 
     */
    public Output<Optional<String>> vlanNameId() {
        return Codegen.optional(this.vlanNameId);
    }
    /**
     * Vlan tag ID - Range: `0`-`31`
     * 
     */
    @Export(name="vlanTagId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> vlanTagId;

    /**
     * @return Vlan tag ID - Range: `0`-`31`
     * 
     */
    public Output<Optional<Integer>> vlanTagId() {
        return Codegen.optional(this.vlanTagId);
    }
    /**
     * Voice domain permission - Default value: `false`
     * 
     */
    @Export(name="voiceDomainPermission", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> voiceDomainPermission;

    /**
     * @return Voice domain permission - Default value: `false`
     * 
     */
    public Output<Boolean> voiceDomainPermission() {
        return this.voiceDomainPermission;
    }
    /**
     * Web authentication (local) - Default value: `false`
     * 
     */
    @Export(name="webAuth", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> webAuth;

    /**
     * @return Web authentication (local) - Default value: `false`
     * 
     */
    public Output<Boolean> webAuth() {
        return this.webAuth;
    }
    /**
     * Web redirection ACL
     * 
     */
    @Export(name="webRedirectionAcl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> webRedirectionAcl;

    /**
     * @return Web redirection ACL
     * 
     */
    public Output<Optional<String>> webRedirectionAcl() {
        return Codegen.optional(this.webRedirectionAcl);
    }
    /**
     * This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other
     * `web_redirection_type` values the field must be ignored.
     * 
     */
    @Export(name="webRedirectionDisplayCertificatesRenewalMessages", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> webRedirectionDisplayCertificatesRenewalMessages;

    /**
     * @return This attribute is mandatory when `web_redirection_type` value is `CentralizedWebAuth`. For all other
     * `web_redirection_type` values the field must be ignored.
     * 
     */
    public Output<Optional<Boolean>> webRedirectionDisplayCertificatesRenewalMessages() {
        return Codegen.optional(this.webRedirectionDisplayCertificatesRenewalMessages);
    }
    /**
     * A portal that exist in the DB and fits the `web_redirection_type`
     * 
     */
    @Export(name="webRedirectionPortalName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> webRedirectionPortalName;

    /**
     * @return A portal that exist in the DB and fits the `web_redirection_type`
     * 
     */
    public Output<Optional<String>> webRedirectionPortalName() {
        return Codegen.optional(this.webRedirectionPortalName);
    }
    /**
     * IP, hostname or FQDN
     * 
     */
    @Export(name="webRedirectionStaticIpHostNameFqdn", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> webRedirectionStaticIpHostNameFqdn;

    /**
     * @return IP, hostname or FQDN
     * 
     */
    public Output<Optional<String>> webRedirectionStaticIpHostNameFqdn() {
        return Codegen.optional(this.webRedirectionStaticIpHostNameFqdn);
    }
    /**
     * This type must fit the `web_redirection_portal_name` - Choices: `CentralizedWebAuth`, `HotSpot`,
     * `NativeSupplicanProvisioning`, `ClientProvisioning`
     * 
     */
    @Export(name="webRedirectionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> webRedirectionType;

    /**
     * @return This type must fit the `web_redirection_portal_name` - Choices: `CentralizedWebAuth`, `HotSpot`,
     * `NativeSupplicanProvisioning`, `ClientProvisioning`
     * 
     */
    public Output<Optional<String>> webRedirectionType() {
        return Codegen.optional(this.webRedirectionType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthorizationProfile(java.lang.String name) {
        this(name, AuthorizationProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthorizationProfile(java.lang.String name, @Nullable AuthorizationProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthorizationProfile(java.lang.String name, @Nullable AuthorizationProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:networkaccess/authorizationProfile:AuthorizationProfile", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private AuthorizationProfile(java.lang.String name, Output<java.lang.String> id, @Nullable AuthorizationProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:networkaccess/authorizationProfile:AuthorizationProfile", name, state, makeResourceOptions(options, id), false);
    }

    private static AuthorizationProfileArgs makeArgs(@Nullable AuthorizationProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? AuthorizationProfileArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AuthorizationProfile get(java.lang.String name, Output<java.lang.String> id, @Nullable AuthorizationProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthorizationProfile(name, id, state, options);
    }
}

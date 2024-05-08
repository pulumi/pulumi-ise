// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.network;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.network.DeviceArgs;
import com.pulumi.ise.network.inputs.DeviceState;
import com.pulumi.ise.network.outputs.DeviceIp;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a Network Device.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ise.network.Device;
 * import com.pulumi.ise.network.DeviceArgs;
 * import com.pulumi.ise.network.inputs.DeviceIpArgs;
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
 *         var example = new Device(&#34;example&#34;, DeviceArgs.builder()        
 *             .name(&#34;Device1&#34;)
 *             .description(&#34;My device&#34;)
 *             .authenticationEnableKeyWrap(true)
 *             .authenticationEncryptionKey(&#34;cisco123cisco123&#34;)
 *             .authenticationEncryptionKeyFormat(&#34;ASCII&#34;)
 *             .authenticationMessageAuthenticatorCodeKey(&#34;cisco123cisco1235678&#34;)
 *             .authenticationNetworkProtocol(&#34;RADIUS&#34;)
 *             .authenticationRadiusSharedSecret(&#34;cisco123&#34;)
 *             .authenticationEnableMultiSecret(true)
 *             .authenticationSecondRadiusSharedSecret(&#34;cisco12345&#34;)
 *             .authenticationDtlsRequired(true)
 *             .coaPort(12345)
 *             .dtlsDnsName(&#34;cisco.com&#34;)
 *             .ips(DeviceIpArgs.builder()
 *                 .ipaddress(&#34;2.3.4.5&#34;)
 *                 .mask(&#34;32&#34;)
 *                 .build())
 *             .modelName(&#34;Unknown&#34;)
 *             .softwareVersion(&#34;Unknown&#34;)
 *             .profileName(&#34;Cisco&#34;)
 *             .snmpLinkTrapQuery(true)
 *             .snmpMacTrapQuery(true)
 *             .snmpPollingInterval(1200)
 *             .snmpRoCommunity(&#34;rocom&#34;)
 *             .snmpVersion(&#34;TWO_C&#34;)
 *             .tacacsConnectModeOptions(&#34;OFF&#34;)
 *             .tacacsSharedSecret(&#34;cisco123&#34;)
 *             .trustsecDeviceId(&#34;device123&#34;)
 *             .trustsecDevicePassword(&#34;cisco123&#34;)
 *             .trustsecRestApiUsername(&#34;user123&#34;)
 *             .trustsecRestApiPassword(&#34;Cisco123&#34;)
 *             .trustsecEnableModePassword(&#34;cisco123&#34;)
 *             .trustsecExecModePassword(&#34;cisco123&#34;)
 *             .trustsecExecModeUsername(&#34;user456&#34;)
 *             .trustsecIncludeWhenDeployingSgtUpdates(true)
 *             .trustsecDownloadEnviromentDataEveryXSeconds(1000)
 *             .trustsecDownloadPeerAuthorizationPolicyEveryXSeconds(1000)
 *             .trustsecDownloadSgaclListsEveryXSeconds(1000)
 *             .trustsecOtherSgaDevicesToTrustThisDevice(true)
 *             .trustsecReAuthenticationEveryXSeconds(1000)
 *             .trustsecSendConfigurationToDevice(true)
 *             .trustsecSendConfigurationToDeviceUsing(&#34;ENABLE_USING_COA&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ise:network/device:Device example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:network/device:Device")
public class Device extends com.pulumi.resources.CustomResource {
    /**
     * Enforce use of DTLS
     * 
     */
    @Export(name="authenticationDtlsRequired", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authenticationDtlsRequired;

    /**
     * @return Enforce use of DTLS
     * 
     */
    public Output<Optional<Boolean>> authenticationDtlsRequired() {
        return Codegen.optional(this.authenticationDtlsRequired);
    }
    /**
     * Enable key wrap
     * 
     */
    @Export(name="authenticationEnableKeyWrap", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authenticationEnableKeyWrap;

    /**
     * @return Enable key wrap
     * 
     */
    public Output<Optional<Boolean>> authenticationEnableKeyWrap() {
        return Codegen.optional(this.authenticationEnableKeyWrap);
    }
    /**
     * Enable multiple RADIUS shared secrets
     * 
     */
    @Export(name="authenticationEnableMultiSecret", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> authenticationEnableMultiSecret;

    /**
     * @return Enable multiple RADIUS shared secrets
     * 
     */
    public Output<Optional<Boolean>> authenticationEnableMultiSecret() {
        return Codegen.optional(this.authenticationEnableMultiSecret);
    }
    /**
     * Encryption key
     * 
     */
    @Export(name="authenticationEncryptionKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationEncryptionKey;

    /**
     * @return Encryption key
     * 
     */
    public Output<Optional<String>> authenticationEncryptionKey() {
        return Codegen.optional(this.authenticationEncryptionKey);
    }
    /**
     * Key input format - Choices: `ASCII`, `HEXADECIMAL`
     * 
     */
    @Export(name="authenticationEncryptionKeyFormat", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationEncryptionKeyFormat;

    /**
     * @return Key input format - Choices: `ASCII`, `HEXADECIMAL`
     * 
     */
    public Output<Optional<String>> authenticationEncryptionKeyFormat() {
        return Codegen.optional(this.authenticationEncryptionKeyFormat);
    }
    /**
     * Message authenticator code key
     * 
     */
    @Export(name="authenticationMessageAuthenticatorCodeKey", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationMessageAuthenticatorCodeKey;

    /**
     * @return Message authenticator code key
     * 
     */
    public Output<Optional<String>> authenticationMessageAuthenticatorCodeKey() {
        return Codegen.optional(this.authenticationMessageAuthenticatorCodeKey);
    }
    /**
     * Network protocol - Choices: `RADIUS`, `TACACS_PLUS`
     * 
     */
    @Export(name="authenticationNetworkProtocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationNetworkProtocol;

    /**
     * @return Network protocol - Choices: `RADIUS`, `TACACS_PLUS`
     * 
     */
    public Output<Optional<String>> authenticationNetworkProtocol() {
        return Codegen.optional(this.authenticationNetworkProtocol);
    }
    /**
     * RADIUS shared secret
     * 
     */
    @Export(name="authenticationRadiusSharedSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationRadiusSharedSecret;

    /**
     * @return RADIUS shared secret
     * 
     */
    public Output<Optional<String>> authenticationRadiusSharedSecret() {
        return Codegen.optional(this.authenticationRadiusSharedSecret);
    }
    /**
     * Second RADIUS shared secret
     * 
     */
    @Export(name="authenticationSecondRadiusSharedSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationSecondRadiusSharedSecret;

    /**
     * @return Second RADIUS shared secret
     * 
     */
    public Output<Optional<String>> authenticationSecondRadiusSharedSecret() {
        return Codegen.optional(this.authenticationSecondRadiusSharedSecret);
    }
    /**
     * CoA port - Default value: `1700`
     * 
     */
    @Export(name="coaPort", refs={Integer.class}, tree="[0]")
    private Output<Integer> coaPort;

    /**
     * @return CoA port - Default value: `1700`
     * 
     */
    public Output<Integer> coaPort() {
        return this.coaPort;
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
     * This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
     * 
     */
    @Export(name="dtlsDnsName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> dtlsDnsName;

    /**
     * @return This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
     * 
     */
    public Output<Optional<String>> dtlsDnsName() {
        return Codegen.optional(this.dtlsDnsName);
    }
    /**
     * List of IP subnets
     * 
     */
    @Export(name="ips", refs={List.class,DeviceIp.class}, tree="[0,1]")
    private Output<List<DeviceIp>> ips;

    /**
     * @return List of IP subnets
     * 
     */
    public Output<List<DeviceIp>> ips() {
        return this.ips;
    }
    /**
     * Model name
     * 
     */
    @Export(name="modelName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> modelName;

    /**
     * @return Model name
     * 
     */
    public Output<Optional<String>> modelName() {
        return Codegen.optional(this.modelName);
    }
    /**
     * The name of the network device
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the network device
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * List of network device groups, e.g. `Device Type#All Device Types#ACCESS`
     * 
     */
    @Export(name="networkDeviceGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> networkDeviceGroups;

    /**
     * @return List of network device groups, e.g. `Device Type#All Device Types#ACCESS`
     * 
     */
    public Output<Optional<List<String>>> networkDeviceGroups() {
        return Codegen.optional(this.networkDeviceGroups);
    }
    /**
     * Profile name - Default value: `Cisco`
     * 
     */
    @Export(name="profileName", refs={String.class}, tree="[0]")
    private Output<String> profileName;

    /**
     * @return Profile name - Default value: `Cisco`
     * 
     */
    public Output<String> profileName() {
        return this.profileName;
    }
    /**
     * SNMP link Trap Query
     * 
     */
    @Export(name="snmpLinkTrapQuery", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> snmpLinkTrapQuery;

    /**
     * @return SNMP link Trap Query
     * 
     */
    public Output<Optional<Boolean>> snmpLinkTrapQuery() {
        return Codegen.optional(this.snmpLinkTrapQuery);
    }
    /**
     * SNMP MAC Trap Query
     * 
     */
    @Export(name="snmpMacTrapQuery", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> snmpMacTrapQuery;

    /**
     * @return SNMP MAC Trap Query
     * 
     */
    public Output<Optional<Boolean>> snmpMacTrapQuery() {
        return Codegen.optional(this.snmpMacTrapQuery);
    }
    /**
     * Originating Policy Services Node
     * 
     */
    @Export(name="snmpOriginatingPolicyServiceNode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snmpOriginatingPolicyServiceNode;

    /**
     * @return Originating Policy Services Node
     * 
     */
    public Output<Optional<String>> snmpOriginatingPolicyServiceNode() {
        return Codegen.optional(this.snmpOriginatingPolicyServiceNode);
    }
    /**
     * SNMP Polling Interval in seconds - Range: `600`-`86400`
     * 
     */
    @Export(name="snmpPollingInterval", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> snmpPollingInterval;

    /**
     * @return SNMP Polling Interval in seconds - Range: `600`-`86400`
     * 
     */
    public Output<Optional<Integer>> snmpPollingInterval() {
        return Codegen.optional(this.snmpPollingInterval);
    }
    /**
     * SNMP RO Community
     * 
     */
    @Export(name="snmpRoCommunity", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snmpRoCommunity;

    /**
     * @return SNMP RO Community
     * 
     */
    public Output<Optional<String>> snmpRoCommunity() {
        return Codegen.optional(this.snmpRoCommunity);
    }
    /**
     * SNMP version - Choices: `ONE`, `TWO_C`, `THREE`
     * 
     */
    @Export(name="snmpVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> snmpVersion;

    /**
     * @return SNMP version - Choices: `ONE`, `TWO_C`, `THREE`
     * 
     */
    public Output<Optional<String>> snmpVersion() {
        return Codegen.optional(this.snmpVersion);
    }
    /**
     * Software version
     * 
     */
    @Export(name="softwareVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> softwareVersion;

    /**
     * @return Software version
     * 
     */
    public Output<Optional<String>> softwareVersion() {
        return Codegen.optional(this.softwareVersion);
    }
    /**
     * Connect mode options - Choices: `OFF`, `ON_LEGACY`, `ON_DRAFT_COMPLIANT`
     * 
     */
    @Export(name="tacacsConnectModeOptions", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tacacsConnectModeOptions;

    /**
     * @return Connect mode options - Choices: `OFF`, `ON_LEGACY`, `ON_DRAFT_COMPLIANT`
     * 
     */
    public Output<Optional<String>> tacacsConnectModeOptions() {
        return Codegen.optional(this.tacacsConnectModeOptions);
    }
    /**
     * Shared secret
     * 
     */
    @Export(name="tacacsSharedSecret", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tacacsSharedSecret;

    /**
     * @return Shared secret
     * 
     */
    public Output<Optional<String>> tacacsSharedSecret() {
        return Codegen.optional(this.tacacsSharedSecret);
    }
    /**
     * CoA source host
     * 
     */
    @Export(name="trustsecCoaSourceHost", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecCoaSourceHost;

    /**
     * @return CoA source host
     * 
     */
    public Output<Optional<String>> trustsecCoaSourceHost() {
        return Codegen.optional(this.trustsecCoaSourceHost);
    }
    /**
     * TrustSec device ID
     * 
     */
    @Export(name="trustsecDeviceId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecDeviceId;

    /**
     * @return TrustSec device ID
     * 
     */
    public Output<Optional<String>> trustsecDeviceId() {
        return Codegen.optional(this.trustsecDeviceId);
    }
    /**
     * TrustSec device password
     * 
     */
    @Export(name="trustsecDevicePassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecDevicePassword;

    /**
     * @return TrustSec device password
     * 
     */
    public Output<Optional<String>> trustsecDevicePassword() {
        return Codegen.optional(this.trustsecDevicePassword);
    }
    /**
     * Download environment data every X seconds
     * 
     */
    @Export(name="trustsecDownloadEnviromentDataEveryXSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> trustsecDownloadEnviromentDataEveryXSeconds;

    /**
     * @return Download environment data every X seconds
     * 
     */
    public Output<Optional<Integer>> trustsecDownloadEnviromentDataEveryXSeconds() {
        return Codegen.optional(this.trustsecDownloadEnviromentDataEveryXSeconds);
    }
    /**
     * Download peer authorization policy every X seconds
     * 
     */
    @Export(name="trustsecDownloadPeerAuthorizationPolicyEveryXSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> trustsecDownloadPeerAuthorizationPolicyEveryXSeconds;

    /**
     * @return Download peer authorization policy every X seconds
     * 
     */
    public Output<Optional<Integer>> trustsecDownloadPeerAuthorizationPolicyEveryXSeconds() {
        return Codegen.optional(this.trustsecDownloadPeerAuthorizationPolicyEveryXSeconds);
    }
    /**
     * Download SGACL lists every X seconds
     * 
     */
    @Export(name="trustsecDownloadSgaclListsEveryXSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> trustsecDownloadSgaclListsEveryXSeconds;

    /**
     * @return Download SGACL lists every X seconds
     * 
     */
    public Output<Optional<Integer>> trustsecDownloadSgaclListsEveryXSeconds() {
        return Codegen.optional(this.trustsecDownloadSgaclListsEveryXSeconds);
    }
    /**
     * Enable mode password
     * 
     */
    @Export(name="trustsecEnableModePassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecEnableModePassword;

    /**
     * @return Enable mode password
     * 
     */
    public Output<Optional<String>> trustsecEnableModePassword() {
        return Codegen.optional(this.trustsecEnableModePassword);
    }
    /**
     * EXEC mode password
     * 
     */
    @Export(name="trustsecExecModePassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecExecModePassword;

    /**
     * @return EXEC mode password
     * 
     */
    public Output<Optional<String>> trustsecExecModePassword() {
        return Codegen.optional(this.trustsecExecModePassword);
    }
    /**
     * EXEC mode username
     * 
     */
    @Export(name="trustsecExecModeUsername", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecExecModeUsername;

    /**
     * @return EXEC mode username
     * 
     */
    public Output<Optional<String>> trustsecExecModeUsername() {
        return Codegen.optional(this.trustsecExecModeUsername);
    }
    /**
     * Include this device when deploying Security Group Tag Mapping Updates
     * 
     */
    @Export(name="trustsecIncludeWhenDeployingSgtUpdates", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> trustsecIncludeWhenDeployingSgtUpdates;

    /**
     * @return Include this device when deploying Security Group Tag Mapping Updates
     * 
     */
    public Output<Optional<Boolean>> trustsecIncludeWhenDeployingSgtUpdates() {
        return Codegen.optional(this.trustsecIncludeWhenDeployingSgtUpdates);
    }
    /**
     * Other TrustSec devices to trust this device
     * 
     */
    @Export(name="trustsecOtherSgaDevicesToTrustThisDevice", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> trustsecOtherSgaDevicesToTrustThisDevice;

    /**
     * @return Other TrustSec devices to trust this device
     * 
     */
    public Output<Optional<Boolean>> trustsecOtherSgaDevicesToTrustThisDevice() {
        return Codegen.optional(this.trustsecOtherSgaDevicesToTrustThisDevice);
    }
    /**
     * Re-authenticate every X seconds
     * 
     */
    @Export(name="trustsecReAuthenticationEveryXSeconds", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> trustsecReAuthenticationEveryXSeconds;

    /**
     * @return Re-authenticate every X seconds
     * 
     */
    public Output<Optional<Integer>> trustsecReAuthenticationEveryXSeconds() {
        return Codegen.optional(this.trustsecReAuthenticationEveryXSeconds);
    }
    /**
     * REST API password
     * 
     */
    @Export(name="trustsecRestApiPassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecRestApiPassword;

    /**
     * @return REST API password
     * 
     */
    public Output<Optional<String>> trustsecRestApiPassword() {
        return Codegen.optional(this.trustsecRestApiPassword);
    }
    /**
     * REST API username
     * 
     */
    @Export(name="trustsecRestApiUsername", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecRestApiUsername;

    /**
     * @return REST API username
     * 
     */
    public Output<Optional<String>> trustsecRestApiUsername() {
        return Codegen.optional(this.trustsecRestApiUsername);
    }
    /**
     * Send configuration to device
     * 
     */
    @Export(name="trustsecSendConfigurationToDevice", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> trustsecSendConfigurationToDevice;

    /**
     * @return Send configuration to device
     * 
     */
    public Output<Optional<Boolean>> trustsecSendConfigurationToDevice() {
        return Codegen.optional(this.trustsecSendConfigurationToDevice);
    }
    /**
     * Send configuration to device using - Choices: `DISABLE_ALL`, `ENABLE_USING_CLI`, `ENABLE_USING_COA`
     * 
     */
    @Export(name="trustsecSendConfigurationToDeviceUsing", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> trustsecSendConfigurationToDeviceUsing;

    /**
     * @return Send configuration to device using - Choices: `DISABLE_ALL`, `ENABLE_USING_CLI`, `ENABLE_USING_COA`
     * 
     */
    public Output<Optional<String>> trustsecSendConfigurationToDeviceUsing() {
        return Codegen.optional(this.trustsecSendConfigurationToDeviceUsing);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Device(String name) {
        this(name, DeviceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Device(String name, DeviceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Device(String name, DeviceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:network/device:Device", name, args == null ? DeviceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Device(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:network/device:Device", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Device get(String name, Output<String> id, @Nullable DeviceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Device(name, id, state, options);
    }
}

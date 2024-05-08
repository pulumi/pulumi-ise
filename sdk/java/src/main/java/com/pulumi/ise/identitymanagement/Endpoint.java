// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.identitymanagement.EndpointArgs;
import com.pulumi.ise.identitymanagement.inputs.EndpointState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage an Endpoint.
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
 * import com.pulumi.ise.identitymanagement.Endpoint;
 * import com.pulumi.ise.identitymanagement.EndpointArgs;
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
 *         var example = new Endpoint(&#34;example&#34;, EndpointArgs.builder()        
 *             .name(&#34;00:11:22:33:44:55&#34;)
 *             .description(&#34;My endpoint&#34;)
 *             .mac(&#34;00:11:22:33:44:55&#34;)
 *             .groupId(&#34;3a88eec0-8c00-11e6-996c-525400b48521&#34;)
 *             .profileId(&#34;3a91a150-8c00-11e6-996c-525400b48521&#34;)
 *             .staticProfileAssignment(true)
 *             .staticProfileAssignmentDefined(true)
 *             .staticGroupAssignment(true)
 *             .staticGroupAssignmentDefined(true)
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
 * $ pulumi import ise:identitymanagement/endpoint:Endpoint example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:identitymanagement/endpoint:Endpoint")
public class Endpoint extends com.pulumi.resources.CustomResource {
    /**
     * Custom Attributes
     * 
     */
    @Export(name="customAttributes", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> customAttributes;

    /**
     * @return Custom Attributes
     * 
     */
    public Output<Optional<Map<String,String>>> customAttributes() {
        return Codegen.optional(this.customAttributes);
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
     * Identity Group ID
     * 
     */
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> groupId;

    /**
     * @return Identity Group ID
     * 
     */
    public Output<Optional<String>> groupId() {
        return Codegen.optional(this.groupId);
    }
    /**
     * Identity Store
     * 
     */
    @Export(name="identityStore", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityStore;

    /**
     * @return Identity Store
     * 
     */
    public Output<Optional<String>> identityStore() {
        return Codegen.optional(this.identityStore);
    }
    /**
     * Identity Store Id
     * 
     */
    @Export(name="identityStoreId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityStoreId;

    /**
     * @return Identity Store Id
     * 
     */
    public Output<Optional<String>> identityStoreId() {
        return Codegen.optional(this.identityStoreId);
    }
    /**
     * MAC address of the endpoint
     * 
     */
    @Export(name="mac", refs={String.class}, tree="[0]")
    private Output<String> mac;

    /**
     * @return MAC address of the endpoint
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }
    /**
     * Mdm Compliance Status
     * 
     */
    @Export(name="mdmComplianceStatus", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmComplianceStatus;

    /**
     * @return Mdm Compliance Status
     * 
     */
    public Output<Optional<Boolean>> mdmComplianceStatus() {
        return Codegen.optional(this.mdmComplianceStatus);
    }
    /**
     * Mdm Encrypted
     * 
     */
    @Export(name="mdmEncrypted", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmEncrypted;

    /**
     * @return Mdm Encrypted
     * 
     */
    public Output<Optional<Boolean>> mdmEncrypted() {
        return Codegen.optional(this.mdmEncrypted);
    }
    /**
     * Mdm Enrolled
     * 
     */
    @Export(name="mdmEnrolled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmEnrolled;

    /**
     * @return Mdm Enrolled
     * 
     */
    public Output<Optional<Boolean>> mdmEnrolled() {
        return Codegen.optional(this.mdmEnrolled);
    }
    /**
     * Mdm IMEI
     * 
     */
    @Export(name="mdmImei", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmImei;

    /**
     * @return Mdm IMEI
     * 
     */
    public Output<Optional<String>> mdmImei() {
        return Codegen.optional(this.mdmImei);
    }
    /**
     * Mdm JailBroken
     * 
     */
    @Export(name="mdmJailBroken", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmJailBroken;

    /**
     * @return Mdm JailBroken
     * 
     */
    public Output<Optional<Boolean>> mdmJailBroken() {
        return Codegen.optional(this.mdmJailBroken);
    }
    /**
     * Mdm Manufacturer
     * 
     */
    @Export(name="mdmManufacturer", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmManufacturer;

    /**
     * @return Mdm Manufacturer
     * 
     */
    public Output<Optional<String>> mdmManufacturer() {
        return Codegen.optional(this.mdmManufacturer);
    }
    /**
     * Mdm Model
     * 
     */
    @Export(name="mdmModel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmModel;

    /**
     * @return Mdm Model
     * 
     */
    public Output<Optional<String>> mdmModel() {
        return Codegen.optional(this.mdmModel);
    }
    /**
     * Mdm OS
     * 
     */
    @Export(name="mdmOs", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmOs;

    /**
     * @return Mdm OS
     * 
     */
    public Output<Optional<String>> mdmOs() {
        return Codegen.optional(this.mdmOs);
    }
    /**
     * Mdm PhoneNumber
     * 
     */
    @Export(name="mdmPhoneNumber", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmPhoneNumber;

    /**
     * @return Mdm PhoneNumber
     * 
     */
    public Output<Optional<String>> mdmPhoneNumber() {
        return Codegen.optional(this.mdmPhoneNumber);
    }
    /**
     * Mdm Pinlock
     * 
     */
    @Export(name="mdmPinlock", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmPinlock;

    /**
     * @return Mdm Pinlock
     * 
     */
    public Output<Optional<Boolean>> mdmPinlock() {
        return Codegen.optional(this.mdmPinlock);
    }
    /**
     * Mdm Reachable
     * 
     */
    @Export(name="mdmReachable", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> mdmReachable;

    /**
     * @return Mdm Reachable
     * 
     */
    public Output<Optional<Boolean>> mdmReachable() {
        return Codegen.optional(this.mdmReachable);
    }
    /**
     * Mdm Serial
     * 
     */
    @Export(name="mdmSerial", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmSerial;

    /**
     * @return Mdm Serial
     * 
     */
    public Output<Optional<String>> mdmSerial() {
        return Codegen.optional(this.mdmSerial);
    }
    /**
     * Mdm Server Name
     * 
     */
    @Export(name="mdmServerName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mdmServerName;

    /**
     * @return Mdm Server Name
     * 
     */
    public Output<Optional<String>> mdmServerName() {
        return Codegen.optional(this.mdmServerName);
    }
    /**
     * The name of the endpoint
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the endpoint
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Portal User
     * 
     */
    @Export(name="portalUser", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> portalUser;

    /**
     * @return Portal User
     * 
     */
    public Output<Optional<String>> portalUser() {
        return Codegen.optional(this.portalUser);
    }
    /**
     * Profile ID
     * 
     */
    @Export(name="profileId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> profileId;

    /**
     * @return Profile ID
     * 
     */
    public Output<Optional<String>> profileId() {
        return Codegen.optional(this.profileId);
    }
    /**
     * Static Group Assignment
     * 
     */
    @Export(name="staticGroupAssignment", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> staticGroupAssignment;

    /**
     * @return Static Group Assignment
     * 
     */
    public Output<Boolean> staticGroupAssignment() {
        return this.staticGroupAssignment;
    }
    /**
     * staticGroupAssignmentDefined - Default value: `true`
     * 
     */
    @Export(name="staticGroupAssignmentDefined", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> staticGroupAssignmentDefined;

    /**
     * @return staticGroupAssignmentDefined - Default value: `true`
     * 
     */
    public Output<Boolean> staticGroupAssignmentDefined() {
        return this.staticGroupAssignmentDefined;
    }
    /**
     * Static Profile Assignment
     * 
     */
    @Export(name="staticProfileAssignment", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> staticProfileAssignment;

    /**
     * @return Static Profile Assignment
     * 
     */
    public Output<Boolean> staticProfileAssignment() {
        return this.staticProfileAssignment;
    }
    /**
     * Static Profile Assignment Defined - Default value: `true`
     * 
     */
    @Export(name="staticProfileAssignmentDefined", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> staticProfileAssignmentDefined;

    /**
     * @return Static Profile Assignment Defined - Default value: `true`
     * 
     */
    public Output<Boolean> staticProfileAssignmentDefined() {
        return this.staticProfileAssignmentDefined;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Endpoint(String name) {
        this(name, EndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Endpoint(String name, EndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Endpoint(String name, EndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/endpoint:Endpoint", name, args == null ? EndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Endpoint(String name, Output<String> id, @Nullable EndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/endpoint:Endpoint", name, state, makeResourceOptions(options, id));
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
    public static Endpoint get(String name, Output<String> id, @Nullable EndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Endpoint(name, id, state, options);
    }
}

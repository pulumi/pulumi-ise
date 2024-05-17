// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.identitymanagement.InternalUserArgs;
import com.pulumi.ise.identitymanagement.inputs.InternalUserState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage an Internal User.
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
 * import com.pulumi.ise.identitymanagement.InternalUser;
 * import com.pulumi.ise.identitymanagement.InternalUserArgs;
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
 *         var example = new InternalUser("example", InternalUserArgs.builder()
 *             .name("UserTF")
 *             .password("Cisco123")
 *             .changePassword(true)
 *             .email("aaa{@literal @}cisco.com")
 *             .accountNameAlias("User 1")
 *             .enablePassword("Cisco123")
 *             .enabled(true)
 *             .passwordNeverExpires(false)
 *             .firstName("John")
 *             .lastName("Doe")
 *             .passwordIdStore("Internal Users")
 *             .description("My first Terraform user")
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
 * $ pulumi import ise:identitymanagement/internalUser:InternalUser example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:identitymanagement/internalUser:InternalUser")
public class InternalUser extends com.pulumi.resources.CustomResource {
    /**
     * The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
     * from ISE 3.2.
     * 
     */
    @Export(name="accountNameAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accountNameAlias;

    /**
     * @return The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
     * from ISE 3.2.
     * 
     */
    public Output<Optional<String>> accountNameAlias() {
        return Codegen.optional(this.accountNameAlias);
    }
    /**
     * Requires the user to change the password - Default value: `true`
     * 
     */
    @Export(name="changePassword", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> changePassword;

    /**
     * @return Requires the user to change the password - Default value: `true`
     * 
     */
    public Output<Boolean> changePassword() {
        return this.changePassword;
    }
    /**
     * Key value map
     * 
     */
    @Export(name="customAttributes", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> customAttributes;

    /**
     * @return Key value map
     * 
     */
    public Output<Optional<String>> customAttributes() {
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
     * Email address
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> email;

    /**
     * @return Email address
     * 
     */
    public Output<Optional<String>> email() {
        return Codegen.optional(this.email);
    }
    /**
     * This field is added in ISE 2.0 to support TACACS+
     * 
     */
    @Export(name="enablePassword", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enablePassword;

    /**
     * @return This field is added in ISE 2.0 to support TACACS+
     * 
     */
    public Output<Optional<String>> enablePassword() {
        return Codegen.optional(this.enablePassword);
    }
    /**
     * Whether the user is enabled/disabled
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether the user is enabled/disabled
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * First name of the internal user
     * 
     */
    @Export(name="firstName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> firstName;

    /**
     * @return First name of the internal user
     * 
     */
    public Output<Optional<String>> firstName() {
        return Codegen.optional(this.firstName);
    }
    /**
     * Comma separated list of identity group IDs.
     * 
     */
    @Export(name="identityGroups", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identityGroups;

    /**
     * @return Comma separated list of identity group IDs.
     * 
     */
    public Output<Optional<String>> identityGroups() {
        return Codegen.optional(this.identityGroups);
    }
    /**
     * Last name of the internal user
     * 
     */
    @Export(name="lastName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lastName;

    /**
     * @return Last name of the internal user
     * 
     */
    public Output<Optional<String>> lastName() {
        return Codegen.optional(this.lastName);
    }
    /**
     * The name of the internal user
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the internal user
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password of the internal user
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return The password of the internal user
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
     * 
     */
    @Export(name="passwordIdStore", refs={String.class}, tree="[0]")
    private Output<String> passwordIdStore;

    /**
     * @return The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
     * 
     */
    public Output<String> passwordIdStore() {
        return this.passwordIdStore;
    }
    /**
     * Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
     * field is only supported from ISE 3.2. - Default value: `false`
     * 
     */
    @Export(name="passwordNeverExpires", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> passwordNeverExpires;

    /**
     * @return Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
     * field is only supported from ISE 3.2. - Default value: `false`
     * 
     */
    public Output<Boolean> passwordNeverExpires() {
        return this.passwordNeverExpires;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InternalUser(String name) {
        this(name, InternalUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InternalUser(String name, InternalUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InternalUser(String name, InternalUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/internalUser:InternalUser", name, args == null ? InternalUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InternalUser(String name, Output<String> id, @Nullable InternalUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/internalUser:InternalUser", name, state, makeResourceOptions(options, id));
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
    public static InternalUser get(String name, Output<String> id, @Nullable InternalUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InternalUser(name, id, state, options);
    }
}

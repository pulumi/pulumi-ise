// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.trustsec.SecurityGroupArgs;
import com.pulumi.ise.trustsec.inputs.SecurityGroupState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a TrustSec Security Group.
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
 * import com.pulumi.ise.trustsec.SecurityGroup;
 * import com.pulumi.ise.trustsec.SecurityGroupArgs;
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
 *         var example = new SecurityGroup("example", SecurityGroupArgs.builder()
 *             .name("SGT1234")
 *             .description("My SGT")
 *             .value(1234)
 *             .propogateToApic(true)
 *             .isReadOnly(false)
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
 * $ pulumi import ise:trustsec/securityGroup:SecurityGroup example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:trustsec/securityGroup:SecurityGroup")
public class SecurityGroup extends com.pulumi.resources.CustomResource {
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
     * Read-only - Default value: `false`
     * 
     */
    @Export(name="isReadOnly", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isReadOnly;

    /**
     * @return Read-only - Default value: `false`
     * 
     */
    public Output<Boolean> isReadOnly() {
        return this.isReadOnly;
    }
    /**
     * The name of the security group
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the security group
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Propagate to APIC (ACI)
     * 
     */
    @Export(name="propogateToApic", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> propogateToApic;

    /**
     * @return Propagate to APIC (ACI)
     * 
     */
    public Output<Optional<Boolean>> propogateToApic() {
        return Codegen.optional(this.propogateToApic);
    }
    /**
     * `-1` to auto-generate - Range: `-1`-`65519`
     * 
     */
    @Export(name="value", refs={Integer.class}, tree="[0]")
    private Output<Integer> value;

    /**
     * @return `-1` to auto-generate - Range: `-1`-`65519`
     * 
     */
    public Output<Integer> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecurityGroup(java.lang.String name) {
        this(name, SecurityGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecurityGroup(java.lang.String name, SecurityGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecurityGroup(java.lang.String name, SecurityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:trustsec/securityGroup:SecurityGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private SecurityGroup(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:trustsec/securityGroup:SecurityGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static SecurityGroupArgs makeArgs(SecurityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SecurityGroupArgs.Empty : args;
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
    public static SecurityGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable SecurityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecurityGroup(name, id, state, options);
    }
}

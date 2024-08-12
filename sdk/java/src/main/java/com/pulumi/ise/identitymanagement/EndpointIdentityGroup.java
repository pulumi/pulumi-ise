// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.identitymanagement.EndpointIdentityGroupArgs;
import com.pulumi.ise.identitymanagement.inputs.EndpointIdentityGroupState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage an Endpoint Identity Group.
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
 * import com.pulumi.ise.identitymanagement.EndpointIdentityGroup;
 * import com.pulumi.ise.identitymanagement.EndpointIdentityGroupArgs;
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
 *         var example = new EndpointIdentityGroup("example", EndpointIdentityGroupArgs.builder()
 *             .name("Group1")
 *             .description("My endpoint identity group")
 *             .systemDefined(false)
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
 * $ pulumi import ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup")
public class EndpointIdentityGroup extends com.pulumi.resources.CustomResource {
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
     * The name of the endpoint identity group
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the endpoint identity group
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Parent endpoint identity group ID
     * 
     */
    @Export(name="parentEndpointIdentityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> parentEndpointIdentityGroupId;

    /**
     * @return Parent endpoint identity group ID
     * 
     */
    public Output<Optional<String>> parentEndpointIdentityGroupId() {
        return Codegen.optional(this.parentEndpointIdentityGroupId);
    }
    /**
     * System defined endpoint identity group - Default value: `false`
     * 
     */
    @Export(name="systemDefined", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> systemDefined;

    /**
     * @return System defined endpoint identity group - Default value: `false`
     * 
     */
    public Output<Boolean> systemDefined() {
        return this.systemDefined;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EndpointIdentityGroup(java.lang.String name) {
        this(name, EndpointIdentityGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EndpointIdentityGroup(java.lang.String name, @Nullable EndpointIdentityGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EndpointIdentityGroup(java.lang.String name, @Nullable EndpointIdentityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private EndpointIdentityGroup(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointIdentityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, state, makeResourceOptions(options, id), false);
    }

    private static EndpointIdentityGroupArgs makeArgs(@Nullable EndpointIdentityGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EndpointIdentityGroupArgs.Empty : args;
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
    public static EndpointIdentityGroup get(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointIdentityGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EndpointIdentityGroup(name, id, state, options);
    }
}

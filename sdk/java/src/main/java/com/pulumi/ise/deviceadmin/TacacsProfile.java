// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.deviceadmin.TacacsProfileArgs;
import com.pulumi.ise.deviceadmin.inputs.TacacsProfileState;
import com.pulumi.ise.deviceadmin.outputs.TacacsProfileSessionAttribute;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a TACACS Profile.
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
 * import com.pulumi.ise.deviceadmin.TacacsProfile;
 * import com.pulumi.ise.deviceadmin.TacacsProfileArgs;
 * import com.pulumi.ise.deviceadmin.inputs.TacacsProfileSessionAttributeArgs;
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
 *         var example = new TacacsProfile("example", TacacsProfileArgs.builder()
 *             .name("Profile1")
 *             .description("My TACACS profile")
 *             .sessionAttributes(TacacsProfileSessionAttributeArgs.builder()
 *                 .type("MANDATORY")
 *                 .name("attr1")
 *                 .value("value")
 *                 .build())
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
 * $ pulumi import ise:deviceadmin/tacacsProfile:TacacsProfile example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:deviceadmin/tacacsProfile:TacacsProfile")
public class TacacsProfile extends com.pulumi.resources.CustomResource {
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
     * The name of the TACACS profile
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the TACACS profile
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="sessionAttributes", refs={List.class,TacacsProfileSessionAttribute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<TacacsProfileSessionAttribute>> sessionAttributes;

    public Output<Optional<List<TacacsProfileSessionAttribute>>> sessionAttributes() {
        return Codegen.optional(this.sessionAttributes);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TacacsProfile(String name) {
        this(name, TacacsProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TacacsProfile(String name, @Nullable TacacsProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TacacsProfile(String name, @Nullable TacacsProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/tacacsProfile:TacacsProfile", name, args == null ? TacacsProfileArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TacacsProfile(String name, Output<String> id, @Nullable TacacsProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/tacacsProfile:TacacsProfile", name, state, makeResourceOptions(options, id));
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
    public static TacacsProfile get(String name, Output<String> id, @Nullable TacacsProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TacacsProfile(name, id, state, options);
    }
}

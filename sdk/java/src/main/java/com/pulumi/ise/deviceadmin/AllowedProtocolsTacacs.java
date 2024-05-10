// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.deviceadmin.AllowedProtocolsTacacsArgs;
import com.pulumi.ise.deviceadmin.inputs.AllowedProtocolsTacacsState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a TACACS allowed protocols policy element.
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
 * import com.pulumi.ise.deviceadmin.AllowedProtocolsTacacs;
 * import com.pulumi.ise.deviceadmin.AllowedProtocolsTacacsArgs;
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
 *         var example = new AllowedProtocolsTacacs("example", AllowedProtocolsTacacsArgs.builder()        
 *             .name("Protocols1")
 *             .description("My allowed TACACS protocols")
 *             .allowPapAscii(true)
 *             .allowChap(true)
 *             .allowMsChapV1(true)
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
 * $ pulumi import ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs")
public class AllowedProtocolsTacacs extends com.pulumi.resources.CustomResource {
    /**
     * Allow CHAP
     * 
     */
    @Export(name="allowChap", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowChap;

    /**
     * @return Allow CHAP
     * 
     */
    public Output<Boolean> allowChap() {
        return this.allowChap;
    }
    /**
     * Allow MS CHAP v1
     * 
     */
    @Export(name="allowMsChapV1", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowMsChapV1;

    /**
     * @return Allow MS CHAP v1
     * 
     */
    public Output<Boolean> allowMsChapV1() {
        return this.allowMsChapV1;
    }
    /**
     * Allow PAP ASCII
     * 
     */
    @Export(name="allowPapAscii", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowPapAscii;

    /**
     * @return Allow PAP ASCII
     * 
     */
    public Output<Boolean> allowPapAscii() {
        return this.allowPapAscii;
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
     * The name of the allowed protocols
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the allowed protocols
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AllowedProtocolsTacacs(String name) {
        this(name, AllowedProtocolsTacacsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AllowedProtocolsTacacs(String name, AllowedProtocolsTacacsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AllowedProtocolsTacacs(String name, AllowedProtocolsTacacsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs", name, args == null ? AllowedProtocolsTacacsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AllowedProtocolsTacacs(String name, Output<String> id, @Nullable AllowedProtocolsTacacsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs", name, state, makeResourceOptions(options, id));
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
    public static AllowedProtocolsTacacs get(String name, Output<String> id, @Nullable AllowedProtocolsTacacsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AllowedProtocolsTacacs(name, id, state, options);
    }
}

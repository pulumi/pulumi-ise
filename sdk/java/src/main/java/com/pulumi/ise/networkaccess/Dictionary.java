// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.networkaccess.DictionaryArgs;
import com.pulumi.ise.networkaccess.inputs.DictionaryState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a Network Access Dictionary.
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
 * import com.pulumi.ise.networkaccess.Dictionary;
 * import com.pulumi.ise.networkaccess.DictionaryArgs;
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
 *         var example = new Dictionary("example", DictionaryArgs.builder()
 *             .name("Dict1")
 *             .description("My description")
 *             .version("1.1")
 *             .dictionaryAttrType("ENTITY_ATTR")
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
 * $ pulumi import ise:networkaccess/dictionary:Dictionary example &#34;Dict1&#34;
 * ```
 * 
 */
@ResourceType(type="ise:networkaccess/dictionary:Dictionary")
public class Dictionary extends com.pulumi.resources.CustomResource {
    /**
     * The description of the dictionary
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the dictionary
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
     * 
     */
    @Export(name="dictionaryAttrType", refs={String.class}, tree="[0]")
    private Output<String> dictionaryAttrType;

    /**
     * @return The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
     * 
     */
    public Output<String> dictionaryAttrType() {
        return this.dictionaryAttrType;
    }
    /**
     * The dictionary name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The dictionary name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The version of the dictionary
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The version of the dictionary
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Dictionary(String name) {
        this(name, DictionaryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Dictionary(String name, DictionaryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Dictionary(String name, DictionaryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:networkaccess/dictionary:Dictionary", name, args == null ? DictionaryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Dictionary(String name, Output<String> id, @Nullable DictionaryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:networkaccess/dictionary:Dictionary", name, state, makeResourceOptions(options, id));
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
    public static Dictionary get(String name, Output<String> id, @Nullable DictionaryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Dictionary(name, id, state, options);
    }
}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.trustsec.IpToSgtMappingGroupArgs;
import com.pulumi.ise.trustsec.inputs.IpToSgtMappingGroupState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a TrustSec IP to SGT Mapping Group.
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
 * import com.pulumi.ise.trustsec.IpToSgtMappingGroup;
 * import com.pulumi.ise.trustsec.IpToSgtMappingGroupArgs;
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
 *         var example = new IpToSgtMappingGroup("example", IpToSgtMappingGroupArgs.builder()
 *             .name("groupA")
 *             .deployType("ALL")
 *             .sgt("93e1bf00-8c01-11e6-996c-525400b48521")
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
 * $ pulumi import ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup")
public class IpToSgtMappingGroup extends com.pulumi.resources.CustomResource {
    /**
     * Mandatory unless `deploy_type` is `ALL`
     * 
     */
    @Export(name="deployTo", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deployTo;

    /**
     * @return Mandatory unless `deploy_type` is `ALL`
     * 
     */
    public Output<Optional<String>> deployTo() {
        return Codegen.optional(this.deployTo);
    }
    /**
     * Deploy Type - Choices: `ALL`, `ND`, `NDG`
     * 
     */
    @Export(name="deployType", refs={String.class}, tree="[0]")
    private Output<String> deployType;

    /**
     * @return Deploy Type - Choices: `ALL`, `ND`, `NDG`
     * 
     */
    public Output<String> deployType() {
        return this.deployType;
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
     * The name of the IP to SGT mapping Group
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the IP to SGT mapping Group
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Trustsec Security Group ID
     * 
     */
    @Export(name="sgt", refs={String.class}, tree="[0]")
    private Output<String> sgt;

    /**
     * @return Trustsec Security Group ID
     * 
     */
    public Output<String> sgt() {
        return this.sgt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpToSgtMappingGroup(String name) {
        this(name, IpToSgtMappingGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpToSgtMappingGroup(String name, IpToSgtMappingGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpToSgtMappingGroup(String name, IpToSgtMappingGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup", name, args == null ? IpToSgtMappingGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpToSgtMappingGroup(String name, Output<String> id, @Nullable IpToSgtMappingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup", name, state, makeResourceOptions(options, id));
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
    public static IpToSgtMappingGroup get(String name, Output<String> id, @Nullable IpToSgtMappingGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpToSgtMappingGroup(name, id, state, options);
    }
}

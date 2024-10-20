// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.deviceadmin.PolicySetUpdateRankArgs;
import com.pulumi.ise.deviceadmin.inputs.PolicySetUpdateRankState;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * This resource is used to update rank field in device admin policy set. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and device_admin_policy_set resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.
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
 * import com.pulumi.ise.deviceadmin.PolicySetUpdateRank;
 * import com.pulumi.ise.deviceadmin.PolicySetUpdateRankArgs;
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
 *         var example = new PolicySetUpdateRank("example", PolicySetUpdateRankArgs.builder()
 *             .policySetId("d82952cb-b901-4b09-b363-5ebf39bdbaf9")
 *             .rank(0)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ise:deviceadmin/policySetUpdateRank:PolicySetUpdateRank")
public class PolicySetUpdateRank extends com.pulumi.resources.CustomResource {
    /**
     * Policy set ID
     * 
     */
    @Export(name="policySetId", refs={String.class}, tree="[0]")
    private Output<String> policySetId;

    /**
     * @return Policy set ID
     * 
     */
    public Output<String> policySetId() {
        return this.policySetId;
    }
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    @Export(name="rank", refs={Integer.class}, tree="[0]")
    private Output<Integer> rank;

    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Output<Integer> rank() {
        return this.rank;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PolicySetUpdateRank(java.lang.String name) {
        this(name, PolicySetUpdateRankArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PolicySetUpdateRank(java.lang.String name, PolicySetUpdateRankArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PolicySetUpdateRank(java.lang.String name, PolicySetUpdateRankArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/policySetUpdateRank:PolicySetUpdateRank", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PolicySetUpdateRank(java.lang.String name, Output<java.lang.String> id, @Nullable PolicySetUpdateRankState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/policySetUpdateRank:PolicySetUpdateRank", name, state, makeResourceOptions(options, id), false);
    }

    private static PolicySetUpdateRankArgs makeArgs(PolicySetUpdateRankArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PolicySetUpdateRankArgs.Empty : args;
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
    public static PolicySetUpdateRank get(java.lang.String name, Output<java.lang.String> id, @Nullable PolicySetUpdateRankState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PolicySetUpdateRank(name, id, state, options);
    }
}

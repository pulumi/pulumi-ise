// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.deviceadmin.PolicySetArgs;
import com.pulumi.ise.deviceadmin.inputs.PolicySetState;
import com.pulumi.ise.deviceadmin.outputs.PolicySetChildren;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a Device Admin Policy Set.
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
 * import com.pulumi.ise.deviceadmin.PolicySet;
 * import com.pulumi.ise.deviceadmin.PolicySetArgs;
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
 *         var example = new PolicySet("example", PolicySetArgs.builder()
 *             .name("PolicySet1")
 *             .description("My description")
 *             .isProxy(false)
 *             .rank(0)
 *             .serviceName("Default Device Admin")
 *             .state("enabled")
 *             .conditionType("ConditionAttributes")
 *             .conditionIsNegate(false)
 *             .conditionAttributeName("Location")
 *             .conditionAttributeValue("All Locations")
 *             .conditionDictionaryName("DEVICE")
 *             .conditionOperator("equals")
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
 * $ pulumi import ise:deviceadmin/policySet:PolicySet example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:deviceadmin/policySet:PolicySet")
public class PolicySet extends com.pulumi.resources.CustomResource {
    /**
     * List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    @Export(name="childrens", refs={List.class,PolicySetChildren.class}, tree="[0,1]")
    private Output</* @Nullable */ List<PolicySetChildren>> childrens;

    /**
     * @return List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    public Output<Optional<List<PolicySetChildren>>> childrens() {
        return Codegen.optional(this.childrens);
    }
    /**
     * Dictionary attribute name
     * 
     */
    @Export(name="conditionAttributeName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionAttributeName;

    /**
     * @return Dictionary attribute name
     * 
     */
    public Output<Optional<String>> conditionAttributeName() {
        return Codegen.optional(this.conditionAttributeName);
    }
    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    @Export(name="conditionAttributeValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionAttributeValue;

    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public Output<Optional<String>> conditionAttributeValue() {
        return Codegen.optional(this.conditionAttributeValue);
    }
    /**
     * Dictionary name
     * 
     */
    @Export(name="conditionDictionaryName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionDictionaryName;

    /**
     * @return Dictionary name
     * 
     */
    public Output<Optional<String>> conditionDictionaryName() {
        return Codegen.optional(this.conditionDictionaryName);
    }
    /**
     * Dictionary value
     * 
     */
    @Export(name="conditionDictionaryValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionDictionaryValue;

    /**
     * @return Dictionary value
     * 
     */
    public Output<Optional<String>> conditionDictionaryValue() {
        return Codegen.optional(this.conditionDictionaryValue);
    }
    /**
     * UUID for condition
     * 
     */
    @Export(name="conditionId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionId;

    /**
     * @return UUID for condition
     * 
     */
    public Output<Optional<String>> conditionId() {
        return Codegen.optional(this.conditionId);
    }
    /**
     * Indicates whereas this condition is in negate mode
     * 
     */
    @Export(name="conditionIsNegate", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> conditionIsNegate;

    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Output<Optional<Boolean>> conditionIsNegate() {
        return Codegen.optional(this.conditionIsNegate);
    }
    /**
     * Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    @Export(name="conditionOperator", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionOperator;

    /**
     * @return Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    public Output<Optional<String>> conditionOperator() {
        return Codegen.optional(this.conditionOperator);
    }
    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     * 
     */
    @Export(name="conditionType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> conditionType;

    /**
     * @return Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     * 
     */
    public Output<Optional<String>> conditionType() {
        return Codegen.optional(this.conditionType);
    }
    /**
     * Indicates if this policy set is the default one
     * 
     */
    @Export(name="default", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> default_;

    /**
     * @return Indicates if this policy set is the default one
     * 
     */
    public Output<Optional<Boolean>> default_() {
        return Codegen.optional(this.default_);
    }
    /**
     * The description of the policy set
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The description of the policy set
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * Flag which indicates if the policy set service is of type &#39;Proxy Sequence&#39; or &#39;Allowed Protocols&#39;
     * 
     */
    @Export(name="isProxy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> isProxy;

    /**
     * @return Flag which indicates if the policy set service is of type &#39;Proxy Sequence&#39; or &#39;Allowed Protocols&#39;
     * 
     */
    public Output<Optional<Boolean>> isProxy() {
        return Codegen.optional(this.isProxy);
    }
    /**
     * Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The rank (priority) in relation to other policy sets. Lower rank is higher priority.
     * 
     */
    @Export(name="rank", refs={Integer.class}, tree="[0]")
    private Output<Integer> rank;

    /**
     * @return The rank (priority) in relation to other policy sets. Lower rank is higher priority.
     * 
     */
    public Output<Integer> rank() {
        return this.rank;
    }
    /**
     * Policy set service identifier. &#39;Allowed Protocols&#39; or &#39;Server Sequence&#39;.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return Policy set service identifier. &#39;Allowed Protocols&#39; or &#39;Server Sequence&#39;.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
     * `monitor`
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
     * `monitor`
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PolicySet(java.lang.String name) {
        this(name, PolicySetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PolicySet(java.lang.String name, PolicySetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PolicySet(java.lang.String name, PolicySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/policySet:PolicySet", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PolicySet(java.lang.String name, Output<java.lang.String> id, @Nullable PolicySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/policySet:PolicySet", name, state, makeResourceOptions(options, id), false);
    }

    private static PolicySetArgs makeArgs(PolicySetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PolicySetArgs.Empty : args;
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
    public static PolicySet get(java.lang.String name, Output<java.lang.String> id, @Nullable PolicySetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PolicySet(name, id, state, options);
    }
}

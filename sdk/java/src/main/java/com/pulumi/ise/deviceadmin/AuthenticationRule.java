// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.deviceadmin.AuthenticationRuleArgs;
import com.pulumi.ise.deviceadmin.inputs.AuthenticationRuleState;
import com.pulumi.ise.deviceadmin.outputs.AuthenticationRuleChildren;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a Device Admin Authentication Rule.
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
 * import com.pulumi.ise.deviceadmin.AuthenticationRule;
 * import com.pulumi.ise.deviceadmin.AuthenticationRuleArgs;
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
 *         var example = new AuthenticationRule("example", AuthenticationRuleArgs.builder()        
 *             .policySetId("d82952cb-b901-4b09-b363-5ebf39bdbaf9")
 *             .name("Rule1")
 *             .default_(false)
 *             .rank(0)
 *             .state("enabled")
 *             .conditionType("ConditionAttributes")
 *             .conditionIsNegate(false)
 *             .conditionAttributeName("Location")
 *             .conditionAttributeValue("All Locations")
 *             .conditionDictionaryName("DEVICE")
 *             .conditionOperator("equals")
 *             .identitySourceName("Internal Endpoints")
 *             .ifAuthFail("REJECT")
 *             .ifProcessFail("DROP")
 *             .ifUserNotFound("REJECT")
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
 * $ pulumi import ise:deviceadmin/authenticationRule:AuthenticationRule example &#34;76d24097-41c4-4558-a4d0-a8c07ac08470,76d24097-41c4-4558-a4d0-a8c07ac08470&#34;
 * ```
 * 
 */
@ResourceType(type="ise:deviceadmin/authenticationRule:AuthenticationRule")
public class AuthenticationRule extends com.pulumi.resources.CustomResource {
    /**
     * List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    @Export(name="childrens", refs={List.class,AuthenticationRuleChildren.class}, tree="[0,1]")
    private Output</* @Nullable */ List<AuthenticationRuleChildren>> childrens;

    /**
     * @return List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    public Output<Optional<List<AuthenticationRuleChildren>>> childrens() {
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
     * Indicates if this rule is the default one
     * 
     */
    @Export(name="default", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> default_;

    /**
     * @return Indicates if this rule is the default one
     * 
     */
    public Output<Optional<Boolean>> default_() {
        return Codegen.optional(this.default_);
    }
    /**
     * Identity source name from the identity stores
     * 
     */
    @Export(name="identitySourceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> identitySourceName;

    /**
     * @return Identity source name from the identity stores
     * 
     */
    public Output<Optional<String>> identitySourceName() {
        return Codegen.optional(this.identitySourceName);
    }
    /**
     * Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
     * `DROP`, `CONTINUE`
     * 
     */
    @Export(name="ifAuthFail", refs={String.class}, tree="[0]")
    private Output<String> ifAuthFail;

    /**
     * @return Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
     * `DROP`, `CONTINUE`
     * 
     */
    public Output<String> ifAuthFail() {
        return this.ifAuthFail;
    }
    /**
     * Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
     * 
     */
    @Export(name="ifProcessFail", refs={String.class}, tree="[0]")
    private Output<String> ifProcessFail;

    /**
     * @return Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
     * 
     */
    public Output<String> ifProcessFail() {
        return this.ifProcessFail;
    }
    /**
     * Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
     * 
     */
    @Export(name="ifUserNotFound", refs={String.class}, tree="[0]")
    private Output<String> ifUserNotFound;

    /**
     * @return Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
     * 
     */
    public Output<String> ifUserNotFound() {
        return this.ifUserNotFound;
    }
    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Output<String> name() {
        return this.name;
    }
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
    private Output</* @Nullable */ Integer> rank;

    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Output<Optional<Integer>> rank() {
        return Codegen.optional(this.rank);
    }
    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> state;

    /**
     * @return The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     * 
     */
    public Output<Optional<String>> state() {
        return Codegen.optional(this.state);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthenticationRule(String name) {
        this(name, AuthenticationRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthenticationRule(String name, AuthenticationRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthenticationRule(String name, AuthenticationRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/authenticationRule:AuthenticationRule", name, args == null ? AuthenticationRuleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthenticationRule(String name, Output<String> id, @Nullable AuthenticationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:deviceadmin/authenticationRule:AuthenticationRule", name, state, makeResourceOptions(options, id));
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
    public static AuthenticationRule get(String name, Output<String> id, @Nullable AuthenticationRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthenticationRule(name, id, state, options);
    }
}

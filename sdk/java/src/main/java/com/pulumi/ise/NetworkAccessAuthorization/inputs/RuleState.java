// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.NetworkAccessAuthorization.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ise.NetworkAccessAuthorization.inputs.RuleChildrenArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuleState extends com.pulumi.resources.ResourceArgs {

    public static final RuleState Empty = new RuleState();

    /**
     * List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    @Import(name="childrens")
    private @Nullable Output<List<RuleChildrenArgs>> childrens;

    /**
     * @return List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    public Optional<Output<List<RuleChildrenArgs>>> childrens() {
        return Optional.ofNullable(this.childrens);
    }

    /**
     * Dictionary attribute name
     * 
     */
    @Import(name="conditionAttributeName")
    private @Nullable Output<String> conditionAttributeName;

    /**
     * @return Dictionary attribute name
     * 
     */
    public Optional<Output<String>> conditionAttributeName() {
        return Optional.ofNullable(this.conditionAttributeName);
    }

    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    @Import(name="conditionAttributeValue")
    private @Nullable Output<String> conditionAttributeValue;

    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public Optional<Output<String>> conditionAttributeValue() {
        return Optional.ofNullable(this.conditionAttributeValue);
    }

    /**
     * Dictionary name
     * 
     */
    @Import(name="conditionDictionaryName")
    private @Nullable Output<String> conditionDictionaryName;

    /**
     * @return Dictionary name
     * 
     */
    public Optional<Output<String>> conditionDictionaryName() {
        return Optional.ofNullable(this.conditionDictionaryName);
    }

    /**
     * Dictionary value
     * 
     */
    @Import(name="conditionDictionaryValue")
    private @Nullable Output<String> conditionDictionaryValue;

    /**
     * @return Dictionary value
     * 
     */
    public Optional<Output<String>> conditionDictionaryValue() {
        return Optional.ofNullable(this.conditionDictionaryValue);
    }

    /**
     * UUID for condition
     * 
     */
    @Import(name="conditionId")
    private @Nullable Output<String> conditionId;

    /**
     * @return UUID for condition
     * 
     */
    public Optional<Output<String>> conditionId() {
        return Optional.ofNullable(this.conditionId);
    }

    /**
     * Indicates whereas this condition is in negate mode
     * 
     */
    @Import(name="conditionIsNegate")
    private @Nullable Output<Boolean> conditionIsNegate;

    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Optional<Output<Boolean>> conditionIsNegate() {
        return Optional.ofNullable(this.conditionIsNegate);
    }

    /**
     * Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    @Import(name="conditionOperator")
    private @Nullable Output<String> conditionOperator;

    /**
     * @return Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
     * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
     * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    public Optional<Output<String>> conditionOperator() {
        return Optional.ofNullable(this.conditionOperator);
    }

    /**
     * Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     * 
     */
    @Import(name="conditionType")
    private @Nullable Output<String> conditionType;

    /**
     * @return Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
     * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
     * `ConditionOrBlock`, `ConditionReference`
     * 
     */
    public Optional<Output<String>> conditionType() {
        return Optional.ofNullable(this.conditionType);
    }

    /**
     * Indicates if this rule is the default one
     * 
     */
    @Import(name="default")
    private @Nullable Output<Boolean> default_;

    /**
     * @return Indicates if this rule is the default one
     * 
     */
    public Optional<Output<Boolean>> default_() {
        return Optional.ofNullable(this.default_);
    }

    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Policy set ID
     * 
     */
    @Import(name="policySetId")
    private @Nullable Output<String> policySetId;

    /**
     * @return Policy set ID
     * 
     */
    public Optional<Output<String>> policySetId() {
        return Optional.ofNullable(this.policySetId);
    }

    /**
     * The authorization profile(s)
     * 
     */
    @Import(name="profiles")
    private @Nullable Output<List<String>> profiles;

    /**
     * @return The authorization profile(s)
     * 
     */
    public Optional<Output<List<String>>> profiles() {
        return Optional.ofNullable(this.profiles);
    }

    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    @Import(name="rank")
    private @Nullable Output<Integer> rank;

    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Optional<Output<Integer>> rank() {
        return Optional.ofNullable(this.rank);
    }

    /**
     * Security group used in authorization policies
     * 
     */
    @Import(name="securityGroup")
    private @Nullable Output<String> securityGroup;

    /**
     * @return Security group used in authorization policies
     * 
     */
    public Optional<Output<String>> securityGroup() {
        return Optional.ofNullable(this.securityGroup);
    }

    /**
     * The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     * 
     */
    @Import(name="state")
    private @Nullable Output<String> state;

    /**
     * @return The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
     * 
     */
    public Optional<Output<String>> state() {
        return Optional.ofNullable(this.state);
    }

    private RuleState() {}

    private RuleState(RuleState $) {
        this.childrens = $.childrens;
        this.conditionAttributeName = $.conditionAttributeName;
        this.conditionAttributeValue = $.conditionAttributeValue;
        this.conditionDictionaryName = $.conditionDictionaryName;
        this.conditionDictionaryValue = $.conditionDictionaryValue;
        this.conditionId = $.conditionId;
        this.conditionIsNegate = $.conditionIsNegate;
        this.conditionOperator = $.conditionOperator;
        this.conditionType = $.conditionType;
        this.default_ = $.default_;
        this.name = $.name;
        this.policySetId = $.policySetId;
        this.profiles = $.profiles;
        this.rank = $.rank;
        this.securityGroup = $.securityGroup;
        this.state = $.state;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuleState $;

        public Builder() {
            $ = new RuleState();
        }

        public Builder(RuleState defaults) {
            $ = new RuleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param childrens List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
         * 
         * @return builder
         * 
         */
        public Builder childrens(@Nullable Output<List<RuleChildrenArgs>> childrens) {
            $.childrens = childrens;
            return this;
        }

        /**
         * @param childrens List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
         * 
         * @return builder
         * 
         */
        public Builder childrens(List<RuleChildrenArgs> childrens) {
            return childrens(Output.of(childrens));
        }

        /**
         * @param childrens List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
         * 
         * @return builder
         * 
         */
        public Builder childrens(RuleChildrenArgs... childrens) {
            return childrens(List.of(childrens));
        }

        /**
         * @param conditionAttributeName Dictionary attribute name
         * 
         * @return builder
         * 
         */
        public Builder conditionAttributeName(@Nullable Output<String> conditionAttributeName) {
            $.conditionAttributeName = conditionAttributeName;
            return this;
        }

        /**
         * @param conditionAttributeName Dictionary attribute name
         * 
         * @return builder
         * 
         */
        public Builder conditionAttributeName(String conditionAttributeName) {
            return conditionAttributeName(Output.of(conditionAttributeName));
        }

        /**
         * @param conditionAttributeValue Attribute value for condition. Value type is specified in dictionary object.
         * 
         * @return builder
         * 
         */
        public Builder conditionAttributeValue(@Nullable Output<String> conditionAttributeValue) {
            $.conditionAttributeValue = conditionAttributeValue;
            return this;
        }

        /**
         * @param conditionAttributeValue Attribute value for condition. Value type is specified in dictionary object.
         * 
         * @return builder
         * 
         */
        public Builder conditionAttributeValue(String conditionAttributeValue) {
            return conditionAttributeValue(Output.of(conditionAttributeValue));
        }

        /**
         * @param conditionDictionaryName Dictionary name
         * 
         * @return builder
         * 
         */
        public Builder conditionDictionaryName(@Nullable Output<String> conditionDictionaryName) {
            $.conditionDictionaryName = conditionDictionaryName;
            return this;
        }

        /**
         * @param conditionDictionaryName Dictionary name
         * 
         * @return builder
         * 
         */
        public Builder conditionDictionaryName(String conditionDictionaryName) {
            return conditionDictionaryName(Output.of(conditionDictionaryName));
        }

        /**
         * @param conditionDictionaryValue Dictionary value
         * 
         * @return builder
         * 
         */
        public Builder conditionDictionaryValue(@Nullable Output<String> conditionDictionaryValue) {
            $.conditionDictionaryValue = conditionDictionaryValue;
            return this;
        }

        /**
         * @param conditionDictionaryValue Dictionary value
         * 
         * @return builder
         * 
         */
        public Builder conditionDictionaryValue(String conditionDictionaryValue) {
            return conditionDictionaryValue(Output.of(conditionDictionaryValue));
        }

        /**
         * @param conditionId UUID for condition
         * 
         * @return builder
         * 
         */
        public Builder conditionId(@Nullable Output<String> conditionId) {
            $.conditionId = conditionId;
            return this;
        }

        /**
         * @param conditionId UUID for condition
         * 
         * @return builder
         * 
         */
        public Builder conditionId(String conditionId) {
            return conditionId(Output.of(conditionId));
        }

        /**
         * @param conditionIsNegate Indicates whereas this condition is in negate mode
         * 
         * @return builder
         * 
         */
        public Builder conditionIsNegate(@Nullable Output<Boolean> conditionIsNegate) {
            $.conditionIsNegate = conditionIsNegate;
            return this;
        }

        /**
         * @param conditionIsNegate Indicates whereas this condition is in negate mode
         * 
         * @return builder
         * 
         */
        public Builder conditionIsNegate(Boolean conditionIsNegate) {
            return conditionIsNegate(Output.of(conditionIsNegate));
        }

        /**
         * @param conditionOperator Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
         * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
         * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
         * 
         * @return builder
         * 
         */
        public Builder conditionOperator(@Nullable Output<String> conditionOperator) {
            $.conditionOperator = conditionOperator;
            return this;
        }

        /**
         * @param conditionOperator Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
         * `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
         * `notEquals`, `notIn`, `notStartsWith`, `startsWith`
         * 
         * @return builder
         * 
         */
        public Builder conditionOperator(String conditionOperator) {
            return conditionOperator(Output.of(conditionOperator));
        }

        /**
         * @param conditionType Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
         * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
         * `ConditionOrBlock`, `ConditionReference`
         * 
         * @return builder
         * 
         */
        public Builder conditionType(@Nullable Output<String> conditionType) {
            $.conditionType = conditionType;
            return this;
        }

        /**
         * @param conditionType Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
         * additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
         * `ConditionOrBlock`, `ConditionReference`
         * 
         * @return builder
         * 
         */
        public Builder conditionType(String conditionType) {
            return conditionType(Output.of(conditionType));
        }

        /**
         * @param default_ Indicates if this rule is the default one
         * 
         * @return builder
         * 
         */
        public Builder default_(@Nullable Output<Boolean> default_) {
            $.default_ = default_;
            return this;
        }

        /**
         * @param default_ Indicates if this rule is the default one
         * 
         * @return builder
         * 
         */
        public Builder default_(Boolean default_) {
            return default_(Output.of(default_));
        }

        /**
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(@Nullable Output<String> policySetId) {
            $.policySetId = policySetId;
            return this;
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(String policySetId) {
            return policySetId(Output.of(policySetId));
        }

        /**
         * @param profiles The authorization profile(s)
         * 
         * @return builder
         * 
         */
        public Builder profiles(@Nullable Output<List<String>> profiles) {
            $.profiles = profiles;
            return this;
        }

        /**
         * @param profiles The authorization profile(s)
         * 
         * @return builder
         * 
         */
        public Builder profiles(List<String> profiles) {
            return profiles(Output.of(profiles));
        }

        /**
         * @param profiles The authorization profile(s)
         * 
         * @return builder
         * 
         */
        public Builder profiles(String... profiles) {
            return profiles(List.of(profiles));
        }

        /**
         * @param rank The rank (priority) in relation to other rules. Lower rank is higher priority.
         * 
         * @return builder
         * 
         */
        public Builder rank(@Nullable Output<Integer> rank) {
            $.rank = rank;
            return this;
        }

        /**
         * @param rank The rank (priority) in relation to other rules. Lower rank is higher priority.
         * 
         * @return builder
         * 
         */
        public Builder rank(Integer rank) {
            return rank(Output.of(rank));
        }

        /**
         * @param securityGroup Security group used in authorization policies
         * 
         * @return builder
         * 
         */
        public Builder securityGroup(@Nullable Output<String> securityGroup) {
            $.securityGroup = securityGroup;
            return this;
        }

        /**
         * @param securityGroup Security group used in authorization policies
         * 
         * @return builder
         * 
         */
        public Builder securityGroup(String securityGroup) {
            return securityGroup(Output.of(securityGroup));
        }

        /**
         * @param state The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
         * 
         * @return builder
         * 
         */
        public Builder state(@Nullable Output<String> state) {
            $.state = state;
            return this;
        }

        /**
         * @param state The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
         * 
         * @return builder
         * 
         */
        public Builder state(String state) {
            return state(Output.of(state));
        }

        public RuleState build() {
            return $;
        }
    }

}

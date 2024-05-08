// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthorizationExceptionRuleChildrenChildrenArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationExceptionRuleChildrenChildrenArgs Empty = new AuthorizationExceptionRuleChildrenChildrenArgs();

    /**
     * Dictionary attribute name
     * 
     */
    @Import(name="attributeName")
    private @Nullable Output<String> attributeName;

    /**
     * @return Dictionary attribute name
     * 
     */
    public Optional<Output<String>> attributeName() {
        return Optional.ofNullable(this.attributeName);
    }

    /**
     * Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    @Import(name="attributeValue")
    private @Nullable Output<String> attributeValue;

    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public Optional<Output<String>> attributeValue() {
        return Optional.ofNullable(this.attributeValue);
    }

    /**
     * Condition type.
     *   - Choices: `ConditionAttributes`, `ConditionReference`
     * 
     */
    @Import(name="conditionType", required=true)
    private Output<String> conditionType;

    /**
     * @return Condition type.
     *   - Choices: `ConditionAttributes`, `ConditionReference`
     * 
     */
    public Output<String> conditionType() {
        return this.conditionType;
    }

    /**
     * Dictionary name
     * 
     */
    @Import(name="dictionaryName")
    private @Nullable Output<String> dictionaryName;

    /**
     * @return Dictionary name
     * 
     */
    public Optional<Output<String>> dictionaryName() {
        return Optional.ofNullable(this.dictionaryName);
    }

    /**
     * Dictionary value
     * 
     */
    @Import(name="dictionaryValue")
    private @Nullable Output<String> dictionaryValue;

    /**
     * @return Dictionary value
     * 
     */
    public Optional<Output<String>> dictionaryValue() {
        return Optional.ofNullable(this.dictionaryValue);
    }

    /**
     * UUID for condition
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return UUID for condition
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Indicates whereas this condition is in negate mode
     * 
     */
    @Import(name="isNegate")
    private @Nullable Output<Boolean> isNegate;

    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Optional<Output<Boolean>> isNegate() {
        return Optional.ofNullable(this.isNegate);
    }

    /**
     * Equality operator
     *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    @Import(name="operator")
    private @Nullable Output<String> operator;

    /**
     * @return Equality operator
     *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    public Optional<Output<String>> operator() {
        return Optional.ofNullable(this.operator);
    }

    private AuthorizationExceptionRuleChildrenChildrenArgs() {}

    private AuthorizationExceptionRuleChildrenChildrenArgs(AuthorizationExceptionRuleChildrenChildrenArgs $) {
        this.attributeName = $.attributeName;
        this.attributeValue = $.attributeValue;
        this.conditionType = $.conditionType;
        this.dictionaryName = $.dictionaryName;
        this.dictionaryValue = $.dictionaryValue;
        this.id = $.id;
        this.isNegate = $.isNegate;
        this.operator = $.operator;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationExceptionRuleChildrenChildrenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationExceptionRuleChildrenChildrenArgs $;

        public Builder() {
            $ = new AuthorizationExceptionRuleChildrenChildrenArgs();
        }

        public Builder(AuthorizationExceptionRuleChildrenChildrenArgs defaults) {
            $ = new AuthorizationExceptionRuleChildrenChildrenArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributeName Dictionary attribute name
         * 
         * @return builder
         * 
         */
        public Builder attributeName(@Nullable Output<String> attributeName) {
            $.attributeName = attributeName;
            return this;
        }

        /**
         * @param attributeName Dictionary attribute name
         * 
         * @return builder
         * 
         */
        public Builder attributeName(String attributeName) {
            return attributeName(Output.of(attributeName));
        }

        /**
         * @param attributeValue Attribute value for condition. Value type is specified in dictionary object.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(@Nullable Output<String> attributeValue) {
            $.attributeValue = attributeValue;
            return this;
        }

        /**
         * @param attributeValue Attribute value for condition. Value type is specified in dictionary object.
         * 
         * @return builder
         * 
         */
        public Builder attributeValue(String attributeValue) {
            return attributeValue(Output.of(attributeValue));
        }

        /**
         * @param conditionType Condition type.
         *   - Choices: `ConditionAttributes`, `ConditionReference`
         * 
         * @return builder
         * 
         */
        public Builder conditionType(Output<String> conditionType) {
            $.conditionType = conditionType;
            return this;
        }

        /**
         * @param conditionType Condition type.
         *   - Choices: `ConditionAttributes`, `ConditionReference`
         * 
         * @return builder
         * 
         */
        public Builder conditionType(String conditionType) {
            return conditionType(Output.of(conditionType));
        }

        /**
         * @param dictionaryName Dictionary name
         * 
         * @return builder
         * 
         */
        public Builder dictionaryName(@Nullable Output<String> dictionaryName) {
            $.dictionaryName = dictionaryName;
            return this;
        }

        /**
         * @param dictionaryName Dictionary name
         * 
         * @return builder
         * 
         */
        public Builder dictionaryName(String dictionaryName) {
            return dictionaryName(Output.of(dictionaryName));
        }

        /**
         * @param dictionaryValue Dictionary value
         * 
         * @return builder
         * 
         */
        public Builder dictionaryValue(@Nullable Output<String> dictionaryValue) {
            $.dictionaryValue = dictionaryValue;
            return this;
        }

        /**
         * @param dictionaryValue Dictionary value
         * 
         * @return builder
         * 
         */
        public Builder dictionaryValue(String dictionaryValue) {
            return dictionaryValue(Output.of(dictionaryValue));
        }

        /**
         * @param id UUID for condition
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id UUID for condition
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param isNegate Indicates whereas this condition is in negate mode
         * 
         * @return builder
         * 
         */
        public Builder isNegate(@Nullable Output<Boolean> isNegate) {
            $.isNegate = isNegate;
            return this;
        }

        /**
         * @param isNegate Indicates whereas this condition is in negate mode
         * 
         * @return builder
         * 
         */
        public Builder isNegate(Boolean isNegate) {
            return isNegate(Output.of(isNegate));
        }

        /**
         * @param operator Equality operator
         *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
         * 
         * @return builder
         * 
         */
        public Builder operator(@Nullable Output<String> operator) {
            $.operator = operator;
            return this;
        }

        /**
         * @param operator Equality operator
         *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
         * 
         * @return builder
         * 
         */
        public Builder operator(String operator) {
            return operator(Output.of(operator));
        }

        public AuthorizationExceptionRuleChildrenChildrenArgs build() {
            if ($.conditionType == null) {
                throw new MissingRequiredPropertyException("AuthorizationExceptionRuleChildrenChildrenArgs", "conditionType");
            }
            return $;
        }
    }

}

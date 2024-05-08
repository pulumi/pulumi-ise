// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class PolicySetChildrenChildren {
    /**
     * @return Dictionary attribute name
     * 
     */
    private @Nullable String attributeName;
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    private @Nullable String attributeValue;
    /**
     * @return Condition type.
     *   - Choices: `ConditionAttributes`, `ConditionReference`
     * 
     */
    private String conditionType;
    /**
     * @return Dictionary name
     * 
     */
    private @Nullable String dictionaryName;
    /**
     * @return Dictionary value
     * 
     */
    private @Nullable String dictionaryValue;
    /**
     * @return UUID for condition
     * 
     */
    private @Nullable String id;
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    private @Nullable Boolean isNegate;
    /**
     * @return Equality operator
     *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    private @Nullable String operator;

    private PolicySetChildrenChildren() {}
    /**
     * @return Dictionary attribute name
     * 
     */
    public Optional<String> attributeName() {
        return Optional.ofNullable(this.attributeName);
    }
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public Optional<String> attributeValue() {
        return Optional.ofNullable(this.attributeValue);
    }
    /**
     * @return Condition type.
     *   - Choices: `ConditionAttributes`, `ConditionReference`
     * 
     */
    public String conditionType() {
        return this.conditionType;
    }
    /**
     * @return Dictionary name
     * 
     */
    public Optional<String> dictionaryName() {
        return Optional.ofNullable(this.dictionaryName);
    }
    /**
     * @return Dictionary value
     * 
     */
    public Optional<String> dictionaryValue() {
        return Optional.ofNullable(this.dictionaryValue);
    }
    /**
     * @return UUID for condition
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Optional<Boolean> isNegate() {
        return Optional.ofNullable(this.isNegate);
    }
    /**
     * @return Equality operator
     *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    public Optional<String> operator() {
        return Optional.ofNullable(this.operator);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(PolicySetChildrenChildren defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String attributeName;
        private @Nullable String attributeValue;
        private String conditionType;
        private @Nullable String dictionaryName;
        private @Nullable String dictionaryValue;
        private @Nullable String id;
        private @Nullable Boolean isNegate;
        private @Nullable String operator;
        public Builder() {}
        public Builder(PolicySetChildrenChildren defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributeName = defaults.attributeName;
    	      this.attributeValue = defaults.attributeValue;
    	      this.conditionType = defaults.conditionType;
    	      this.dictionaryName = defaults.dictionaryName;
    	      this.dictionaryValue = defaults.dictionaryValue;
    	      this.id = defaults.id;
    	      this.isNegate = defaults.isNegate;
    	      this.operator = defaults.operator;
        }

        @CustomType.Setter
        public Builder attributeName(@Nullable String attributeName) {

            this.attributeName = attributeName;
            return this;
        }
        @CustomType.Setter
        public Builder attributeValue(@Nullable String attributeValue) {

            this.attributeValue = attributeValue;
            return this;
        }
        @CustomType.Setter
        public Builder conditionType(String conditionType) {
            if (conditionType == null) {
              throw new MissingRequiredPropertyException("PolicySetChildrenChildren", "conditionType");
            }
            this.conditionType = conditionType;
            return this;
        }
        @CustomType.Setter
        public Builder dictionaryName(@Nullable String dictionaryName) {

            this.dictionaryName = dictionaryName;
            return this;
        }
        @CustomType.Setter
        public Builder dictionaryValue(@Nullable String dictionaryValue) {

            this.dictionaryValue = dictionaryValue;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isNegate(@Nullable Boolean isNegate) {

            this.isNegate = isNegate;
            return this;
        }
        @CustomType.Setter
        public Builder operator(@Nullable String operator) {

            this.operator = operator;
            return this;
        }
        public PolicySetChildrenChildren build() {
            final var _resultValue = new PolicySetChildrenChildren();
            _resultValue.attributeName = attributeName;
            _resultValue.attributeValue = attributeValue;
            _resultValue.conditionType = conditionType;
            _resultValue.dictionaryName = dictionaryName;
            _resultValue.dictionaryValue = dictionaryValue;
            _resultValue.id = id;
            _resultValue.isNegate = isNegate;
            _resultValue.operator = operator;
            return _resultValue;
        }
    }
}

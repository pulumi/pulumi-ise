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
public final class ConditionChildrenChildren {
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
     * @return Condition description
     * 
     */
    private @Nullable String description;
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
     * @return Condition name
     * 
     */
    private @Nullable String name;
    /**
     * @return Equality operator
     *   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
     * 
     */
    private @Nullable String operator;

    private ConditionChildrenChildren() {}
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
     * @return Condition description
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
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
     * @return Condition name
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
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

    public static Builder builder(ConditionChildrenChildren defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String attributeName;
        private @Nullable String attributeValue;
        private String conditionType;
        private @Nullable String description;
        private @Nullable String dictionaryName;
        private @Nullable String dictionaryValue;
        private @Nullable String id;
        private @Nullable Boolean isNegate;
        private @Nullable String name;
        private @Nullable String operator;
        public Builder() {}
        public Builder(ConditionChildrenChildren defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributeName = defaults.attributeName;
    	      this.attributeValue = defaults.attributeValue;
    	      this.conditionType = defaults.conditionType;
    	      this.description = defaults.description;
    	      this.dictionaryName = defaults.dictionaryName;
    	      this.dictionaryValue = defaults.dictionaryValue;
    	      this.id = defaults.id;
    	      this.isNegate = defaults.isNegate;
    	      this.name = defaults.name;
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
              throw new MissingRequiredPropertyException("ConditionChildrenChildren", "conditionType");
            }
            this.conditionType = conditionType;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
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
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder operator(@Nullable String operator) {

            this.operator = operator;
            return this;
        }
        public ConditionChildrenChildren build() {
            final var _resultValue = new ConditionChildrenChildren();
            _resultValue.attributeName = attributeName;
            _resultValue.attributeValue = attributeValue;
            _resultValue.conditionType = conditionType;
            _resultValue.description = description;
            _resultValue.dictionaryName = dictionaryName;
            _resultValue.dictionaryValue = dictionaryValue;
            _resultValue.id = id;
            _resultValue.isNegate = isNegate;
            _resultValue.name = name;
            _resultValue.operator = operator;
            return _resultValue;
        }
    }
}

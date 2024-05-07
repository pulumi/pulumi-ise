// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.DeviceAdminAuthorization.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRuleChildrenChildren {
    /**
     * @return Dictionary attribute name
     * 
     */
    private String attributeName;
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    private String attributeValue;
    /**
     * @return Condition type.
     * 
     */
    private String conditionType;
    /**
     * @return Dictionary name
     * 
     */
    private String dictionaryName;
    /**
     * @return Dictionary value
     * 
     */
    private String dictionaryValue;
    /**
     * @return UUID for condition
     * 
     */
    private String id;
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    private Boolean isNegate;
    /**
     * @return Equality operator
     * 
     */
    private String operator;

    private GetRuleChildrenChildren() {}
    /**
     * @return Dictionary attribute name
     * 
     */
    public String attributeName() {
        return this.attributeName;
    }
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public String attributeValue() {
        return this.attributeValue;
    }
    /**
     * @return Condition type.
     * 
     */
    public String conditionType() {
        return this.conditionType;
    }
    /**
     * @return Dictionary name
     * 
     */
    public String dictionaryName() {
        return this.dictionaryName;
    }
    /**
     * @return Dictionary value
     * 
     */
    public String dictionaryValue() {
        return this.dictionaryValue;
    }
    /**
     * @return UUID for condition
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Boolean isNegate() {
        return this.isNegate;
    }
    /**
     * @return Equality operator
     * 
     */
    public String operator() {
        return this.operator;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRuleChildrenChildren defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String attributeName;
        private String attributeValue;
        private String conditionType;
        private String dictionaryName;
        private String dictionaryValue;
        private String id;
        private Boolean isNegate;
        private String operator;
        public Builder() {}
        public Builder(GetRuleChildrenChildren defaults) {
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
        public Builder attributeName(String attributeName) {
            if (attributeName == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "attributeName");
            }
            this.attributeName = attributeName;
            return this;
        }
        @CustomType.Setter
        public Builder attributeValue(String attributeValue) {
            if (attributeValue == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "attributeValue");
            }
            this.attributeValue = attributeValue;
            return this;
        }
        @CustomType.Setter
        public Builder conditionType(String conditionType) {
            if (conditionType == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "conditionType");
            }
            this.conditionType = conditionType;
            return this;
        }
        @CustomType.Setter
        public Builder dictionaryName(String dictionaryName) {
            if (dictionaryName == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "dictionaryName");
            }
            this.dictionaryName = dictionaryName;
            return this;
        }
        @CustomType.Setter
        public Builder dictionaryValue(String dictionaryValue) {
            if (dictionaryValue == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "dictionaryValue");
            }
            this.dictionaryValue = dictionaryValue;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isNegate(Boolean isNegate) {
            if (isNegate == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "isNegate");
            }
            this.isNegate = isNegate;
            return this;
        }
        @CustomType.Setter
        public Builder operator(String operator) {
            if (operator == null) {
              throw new MissingRequiredPropertyException("GetRuleChildrenChildren", "operator");
            }
            this.operator = operator;
            return this;
        }
        public GetRuleChildrenChildren build() {
            final var _resultValue = new GetRuleChildrenChildren();
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

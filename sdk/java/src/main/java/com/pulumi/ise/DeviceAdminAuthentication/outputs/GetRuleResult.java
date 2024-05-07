// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.DeviceAdminAuthentication.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.DeviceAdminAuthentication.outputs.GetRuleChildren;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRuleResult {
    /**
     * @return List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    private List<GetRuleChildren> childrens;
    /**
     * @return Dictionary attribute name
     * 
     */
    private String conditionAttributeName;
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    private String conditionAttributeValue;
    /**
     * @return Dictionary name
     * 
     */
    private String conditionDictionaryName;
    /**
     * @return Dictionary value
     * 
     */
    private String conditionDictionaryValue;
    /**
     * @return UUID for condition
     * 
     */
    private String conditionId;
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    private Boolean conditionIsNegate;
    /**
     * @return Equality operator
     * 
     */
    private String conditionOperator;
    /**
     * @return Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
     * 
     */
    private String conditionType;
    /**
     * @return Indicates if this rule is the default one
     * 
     */
    private Boolean default_;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return Identity source name from the identity stores
     * 
     */
    private String identitySourceName;
    /**
     * @return Action to perform when authentication fails such as Bad credentials, disabled user and so on
     * 
     */
    private String ifAuthFail;
    /**
     * @return Action to perform when ISE is unable to access the identity database
     * 
     */
    private String ifProcessFail;
    /**
     * @return Action to perform when user is not found in any of identity stores
     * 
     */
    private String ifUserNotFound;
    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    private String name;
    /**
     * @return Policy set ID
     * 
     */
    private String policySetId;
    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    private Integer rank;
    /**
     * @return The state that the rule is in. A disabled rule cannot be matched.
     * 
     */
    private String state;

    private GetRuleResult() {}
    /**
     * @return List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
     * 
     */
    public List<GetRuleChildren> childrens() {
        return this.childrens;
    }
    /**
     * @return Dictionary attribute name
     * 
     */
    public String conditionAttributeName() {
        return this.conditionAttributeName;
    }
    /**
     * @return Attribute value for condition. Value type is specified in dictionary object.
     * 
     */
    public String conditionAttributeValue() {
        return this.conditionAttributeValue;
    }
    /**
     * @return Dictionary name
     * 
     */
    public String conditionDictionaryName() {
        return this.conditionDictionaryName;
    }
    /**
     * @return Dictionary value
     * 
     */
    public String conditionDictionaryValue() {
        return this.conditionDictionaryValue;
    }
    /**
     * @return UUID for condition
     * 
     */
    public String conditionId() {
        return this.conditionId;
    }
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    public Boolean conditionIsNegate() {
        return this.conditionIsNegate;
    }
    /**
     * @return Equality operator
     * 
     */
    public String conditionOperator() {
        return this.conditionOperator;
    }
    /**
     * @return Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
     * 
     */
    public String conditionType() {
        return this.conditionType;
    }
    /**
     * @return Indicates if this rule is the default one
     * 
     */
    public Boolean default_() {
        return this.default_;
    }
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Identity source name from the identity stores
     * 
     */
    public String identitySourceName() {
        return this.identitySourceName;
    }
    /**
     * @return Action to perform when authentication fails such as Bad credentials, disabled user and so on
     * 
     */
    public String ifAuthFail() {
        return this.ifAuthFail;
    }
    /**
     * @return Action to perform when ISE is unable to access the identity database
     * 
     */
    public String ifProcessFail() {
        return this.ifProcessFail;
    }
    /**
     * @return Action to perform when user is not found in any of identity stores
     * 
     */
    public String ifUserNotFound() {
        return this.ifUserNotFound;
    }
    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Policy set ID
     * 
     */
    public String policySetId() {
        return this.policySetId;
    }
    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Integer rank() {
        return this.rank;
    }
    /**
     * @return The state that the rule is in. A disabled rule cannot be matched.
     * 
     */
    public String state() {
        return this.state;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRuleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRuleChildren> childrens;
        private String conditionAttributeName;
        private String conditionAttributeValue;
        private String conditionDictionaryName;
        private String conditionDictionaryValue;
        private String conditionId;
        private Boolean conditionIsNegate;
        private String conditionOperator;
        private String conditionType;
        private Boolean default_;
        private String id;
        private String identitySourceName;
        private String ifAuthFail;
        private String ifProcessFail;
        private String ifUserNotFound;
        private String name;
        private String policySetId;
        private Integer rank;
        private String state;
        public Builder() {}
        public Builder(GetRuleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.childrens = defaults.childrens;
    	      this.conditionAttributeName = defaults.conditionAttributeName;
    	      this.conditionAttributeValue = defaults.conditionAttributeValue;
    	      this.conditionDictionaryName = defaults.conditionDictionaryName;
    	      this.conditionDictionaryValue = defaults.conditionDictionaryValue;
    	      this.conditionId = defaults.conditionId;
    	      this.conditionIsNegate = defaults.conditionIsNegate;
    	      this.conditionOperator = defaults.conditionOperator;
    	      this.conditionType = defaults.conditionType;
    	      this.default_ = defaults.default_;
    	      this.id = defaults.id;
    	      this.identitySourceName = defaults.identitySourceName;
    	      this.ifAuthFail = defaults.ifAuthFail;
    	      this.ifProcessFail = defaults.ifProcessFail;
    	      this.ifUserNotFound = defaults.ifUserNotFound;
    	      this.name = defaults.name;
    	      this.policySetId = defaults.policySetId;
    	      this.rank = defaults.rank;
    	      this.state = defaults.state;
        }

        @CustomType.Setter
        public Builder childrens(List<GetRuleChildren> childrens) {
            if (childrens == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "childrens");
            }
            this.childrens = childrens;
            return this;
        }
        public Builder childrens(GetRuleChildren... childrens) {
            return childrens(List.of(childrens));
        }
        @CustomType.Setter
        public Builder conditionAttributeName(String conditionAttributeName) {
            if (conditionAttributeName == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionAttributeName");
            }
            this.conditionAttributeName = conditionAttributeName;
            return this;
        }
        @CustomType.Setter
        public Builder conditionAttributeValue(String conditionAttributeValue) {
            if (conditionAttributeValue == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionAttributeValue");
            }
            this.conditionAttributeValue = conditionAttributeValue;
            return this;
        }
        @CustomType.Setter
        public Builder conditionDictionaryName(String conditionDictionaryName) {
            if (conditionDictionaryName == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionDictionaryName");
            }
            this.conditionDictionaryName = conditionDictionaryName;
            return this;
        }
        @CustomType.Setter
        public Builder conditionDictionaryValue(String conditionDictionaryValue) {
            if (conditionDictionaryValue == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionDictionaryValue");
            }
            this.conditionDictionaryValue = conditionDictionaryValue;
            return this;
        }
        @CustomType.Setter
        public Builder conditionId(String conditionId) {
            if (conditionId == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionId");
            }
            this.conditionId = conditionId;
            return this;
        }
        @CustomType.Setter
        public Builder conditionIsNegate(Boolean conditionIsNegate) {
            if (conditionIsNegate == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionIsNegate");
            }
            this.conditionIsNegate = conditionIsNegate;
            return this;
        }
        @CustomType.Setter
        public Builder conditionOperator(String conditionOperator) {
            if (conditionOperator == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionOperator");
            }
            this.conditionOperator = conditionOperator;
            return this;
        }
        @CustomType.Setter
        public Builder conditionType(String conditionType) {
            if (conditionType == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "conditionType");
            }
            this.conditionType = conditionType;
            return this;
        }
        @CustomType.Setter("default")
        public Builder default_(Boolean default_) {
            if (default_ == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "default_");
            }
            this.default_ = default_;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder identitySourceName(String identitySourceName) {
            if (identitySourceName == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "identitySourceName");
            }
            this.identitySourceName = identitySourceName;
            return this;
        }
        @CustomType.Setter
        public Builder ifAuthFail(String ifAuthFail) {
            if (ifAuthFail == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "ifAuthFail");
            }
            this.ifAuthFail = ifAuthFail;
            return this;
        }
        @CustomType.Setter
        public Builder ifProcessFail(String ifProcessFail) {
            if (ifProcessFail == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "ifProcessFail");
            }
            this.ifProcessFail = ifProcessFail;
            return this;
        }
        @CustomType.Setter
        public Builder ifUserNotFound(String ifUserNotFound) {
            if (ifUserNotFound == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "ifUserNotFound");
            }
            this.ifUserNotFound = ifUserNotFound;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder policySetId(String policySetId) {
            if (policySetId == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "policySetId");
            }
            this.policySetId = policySetId;
            return this;
        }
        @CustomType.Setter
        public Builder rank(Integer rank) {
            if (rank == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "rank");
            }
            this.rank = rank;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetRuleResult", "state");
            }
            this.state = state;
            return this;
        }
        public GetRuleResult build() {
            final var _resultValue = new GetRuleResult();
            _resultValue.childrens = childrens;
            _resultValue.conditionAttributeName = conditionAttributeName;
            _resultValue.conditionAttributeValue = conditionAttributeValue;
            _resultValue.conditionDictionaryName = conditionDictionaryName;
            _resultValue.conditionDictionaryValue = conditionDictionaryValue;
            _resultValue.conditionId = conditionId;
            _resultValue.conditionIsNegate = conditionIsNegate;
            _resultValue.conditionOperator = conditionOperator;
            _resultValue.conditionType = conditionType;
            _resultValue.default_ = default_;
            _resultValue.id = id;
            _resultValue.identitySourceName = identitySourceName;
            _resultValue.ifAuthFail = ifAuthFail;
            _resultValue.ifProcessFail = ifProcessFail;
            _resultValue.ifUserNotFound = ifUserNotFound;
            _resultValue.name = name;
            _resultValue.policySetId = policySetId;
            _resultValue.rank = rank;
            _resultValue.state = state;
            return _resultValue;
        }
    }
}

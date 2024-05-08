// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTimeAndDateConditionResult {
    /**
     * @return Condition description
     * 
     */
    private String description;
    /**
     * @return End date
     * 
     */
    private String endDate;
    /**
     * @return End time
     * 
     */
    private String endTime;
    /**
     * @return Exception end date
     * 
     */
    private String exceptionEndDate;
    /**
     * @return Exception end time
     * 
     */
    private String exceptionEndTime;
    /**
     * @return Exception start date
     * 
     */
    private String exceptionStartDate;
    /**
     * @return Exception start time
     * 
     */
    private String exceptionStartTime;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return Indicates whereas this condition is in negate mode
     * 
     */
    private Boolean isNegate;
    /**
     * @return Condition name
     * 
     */
    private String name;
    /**
     * @return Start date
     * 
     */
    private String startDate;
    /**
     * @return Start time
     * 
     */
    private String startTime;
    /**
     * @return Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
     * 
     */
    private List<String> weekDays;
    /**
     * @return Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
     * 
     */
    private List<String> weekDaysExceptions;

    private GetTimeAndDateConditionResult() {}
    /**
     * @return Condition description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return End date
     * 
     */
    public String endDate() {
        return this.endDate;
    }
    /**
     * @return End time
     * 
     */
    public String endTime() {
        return this.endTime;
    }
    /**
     * @return Exception end date
     * 
     */
    public String exceptionEndDate() {
        return this.exceptionEndDate;
    }
    /**
     * @return Exception end time
     * 
     */
    public String exceptionEndTime() {
        return this.exceptionEndTime;
    }
    /**
     * @return Exception start date
     * 
     */
    public String exceptionStartDate() {
        return this.exceptionStartDate;
    }
    /**
     * @return Exception start time
     * 
     */
    public String exceptionStartTime() {
        return this.exceptionStartTime;
    }
    /**
     * @return The id of the object
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
     * @return Condition name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Start date
     * 
     */
    public String startDate() {
        return this.startDate;
    }
    /**
     * @return Start time
     * 
     */
    public String startTime() {
        return this.startTime;
    }
    /**
     * @return Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
     * 
     */
    public List<String> weekDays() {
        return this.weekDays;
    }
    /**
     * @return Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
     * 
     */
    public List<String> weekDaysExceptions() {
        return this.weekDaysExceptions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTimeAndDateConditionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String endDate;
        private String endTime;
        private String exceptionEndDate;
        private String exceptionEndTime;
        private String exceptionStartDate;
        private String exceptionStartTime;
        private String id;
        private Boolean isNegate;
        private String name;
        private String startDate;
        private String startTime;
        private List<String> weekDays;
        private List<String> weekDaysExceptions;
        public Builder() {}
        public Builder(GetTimeAndDateConditionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.endDate = defaults.endDate;
    	      this.endTime = defaults.endTime;
    	      this.exceptionEndDate = defaults.exceptionEndDate;
    	      this.exceptionEndTime = defaults.exceptionEndTime;
    	      this.exceptionStartDate = defaults.exceptionStartDate;
    	      this.exceptionStartTime = defaults.exceptionStartTime;
    	      this.id = defaults.id;
    	      this.isNegate = defaults.isNegate;
    	      this.name = defaults.name;
    	      this.startDate = defaults.startDate;
    	      this.startTime = defaults.startTime;
    	      this.weekDays = defaults.weekDays;
    	      this.weekDaysExceptions = defaults.weekDaysExceptions;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder endDate(String endDate) {
            if (endDate == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "endDate");
            }
            this.endDate = endDate;
            return this;
        }
        @CustomType.Setter
        public Builder endTime(String endTime) {
            if (endTime == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "endTime");
            }
            this.endTime = endTime;
            return this;
        }
        @CustomType.Setter
        public Builder exceptionEndDate(String exceptionEndDate) {
            if (exceptionEndDate == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "exceptionEndDate");
            }
            this.exceptionEndDate = exceptionEndDate;
            return this;
        }
        @CustomType.Setter
        public Builder exceptionEndTime(String exceptionEndTime) {
            if (exceptionEndTime == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "exceptionEndTime");
            }
            this.exceptionEndTime = exceptionEndTime;
            return this;
        }
        @CustomType.Setter
        public Builder exceptionStartDate(String exceptionStartDate) {
            if (exceptionStartDate == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "exceptionStartDate");
            }
            this.exceptionStartDate = exceptionStartDate;
            return this;
        }
        @CustomType.Setter
        public Builder exceptionStartTime(String exceptionStartTime) {
            if (exceptionStartTime == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "exceptionStartTime");
            }
            this.exceptionStartTime = exceptionStartTime;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isNegate(Boolean isNegate) {
            if (isNegate == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "isNegate");
            }
            this.isNegate = isNegate;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder startDate(String startDate) {
            if (startDate == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "startDate");
            }
            this.startDate = startDate;
            return this;
        }
        @CustomType.Setter
        public Builder startTime(String startTime) {
            if (startTime == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "startTime");
            }
            this.startTime = startTime;
            return this;
        }
        @CustomType.Setter
        public Builder weekDays(List<String> weekDays) {
            if (weekDays == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "weekDays");
            }
            this.weekDays = weekDays;
            return this;
        }
        public Builder weekDays(String... weekDays) {
            return weekDays(List.of(weekDays));
        }
        @CustomType.Setter
        public Builder weekDaysExceptions(List<String> weekDaysExceptions) {
            if (weekDaysExceptions == null) {
              throw new MissingRequiredPropertyException("GetTimeAndDateConditionResult", "weekDaysExceptions");
            }
            this.weekDaysExceptions = weekDaysExceptions;
            return this;
        }
        public Builder weekDaysExceptions(String... weekDaysExceptions) {
            return weekDaysExceptions(List.of(weekDaysExceptions));
        }
        public GetTimeAndDateConditionResult build() {
            final var _resultValue = new GetTimeAndDateConditionResult();
            _resultValue.description = description;
            _resultValue.endDate = endDate;
            _resultValue.endTime = endTime;
            _resultValue.exceptionEndDate = exceptionEndDate;
            _resultValue.exceptionEndTime = exceptionEndTime;
            _resultValue.exceptionStartDate = exceptionStartDate;
            _resultValue.exceptionStartTime = exceptionStartTime;
            _resultValue.id = id;
            _resultValue.isNegate = isNegate;
            _resultValue.name = name;
            _resultValue.startDate = startDate;
            _resultValue.startTime = startTime;
            _resultValue.weekDays = weekDays;
            _resultValue.weekDaysExceptions = weekDaysExceptions;
            return _resultValue;
        }
    }
}

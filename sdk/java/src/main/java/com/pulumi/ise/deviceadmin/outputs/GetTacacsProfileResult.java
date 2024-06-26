// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.deviceadmin.outputs.GetTacacsProfileSessionAttribute;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetTacacsProfileResult {
    /**
     * @return Description
     * 
     */
    private String description;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return The name of the TACACS profile
     * 
     */
    private String name;
    private List<GetTacacsProfileSessionAttribute> sessionAttributes;

    private GetTacacsProfileResult() {}
    /**
     * @return Description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the TACACS profile
     * 
     */
    public String name() {
        return this.name;
    }
    public List<GetTacacsProfileSessionAttribute> sessionAttributes() {
        return this.sessionAttributes;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTacacsProfileResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String name;
        private List<GetTacacsProfileSessionAttribute> sessionAttributes;
        public Builder() {}
        public Builder(GetTacacsProfileResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.sessionAttributes = defaults.sessionAttributes;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetTacacsProfileResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetTacacsProfileResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetTacacsProfileResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sessionAttributes(List<GetTacacsProfileSessionAttribute> sessionAttributes) {
            if (sessionAttributes == null) {
              throw new MissingRequiredPropertyException("GetTacacsProfileResult", "sessionAttributes");
            }
            this.sessionAttributes = sessionAttributes;
            return this;
        }
        public Builder sessionAttributes(GetTacacsProfileSessionAttribute... sessionAttributes) {
            return sessionAttributes(List.of(sessionAttributes));
        }
        public GetTacacsProfileResult build() {
            final var _resultValue = new GetTacacsProfileResult();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.sessionAttributes = sessionAttributes;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.NetworkDevice.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupResult {
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
     * @return The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     * 
     */
    private String name;
    /**
     * @return The name of the root device group.
     * 
     */
    private String rootGroup;

    private GetGroupResult() {}
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
     * @return The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The name of the root device group.
     * 
     */
    public String rootGroup() {
        return this.rootGroup;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private String name;
        private String rootGroup;
        public Builder() {}
        public Builder(GetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.rootGroup = defaults.rootGroup;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder rootGroup(String rootGroup) {
            if (rootGroup == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "rootGroup");
            }
            this.rootGroup = rootGroup;
            return this;
        }
        public GetGroupResult build() {
            final var _resultValue = new GetGroupResult();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.rootGroup = rootGroup;
            return _resultValue;
        }
    }
}

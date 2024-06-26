// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpToSgtMappingGroupResult {
    /**
     * @return Mandatory unless `deploy_type` is `ALL`
     * 
     */
    private String deployTo;
    /**
     * @return Deploy Type
     * 
     */
    private String deployType;
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
     * @return The name of the IP to SGT mapping Group
     * 
     */
    private String name;
    /**
     * @return Trustsec Security Group ID
     * 
     */
    private String sgt;

    private GetIpToSgtMappingGroupResult() {}
    /**
     * @return Mandatory unless `deploy_type` is `ALL`
     * 
     */
    public String deployTo() {
        return this.deployTo;
    }
    /**
     * @return Deploy Type
     * 
     */
    public String deployType() {
        return this.deployType;
    }
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
     * @return The name of the IP to SGT mapping Group
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Trustsec Security Group ID
     * 
     */
    public String sgt() {
        return this.sgt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpToSgtMappingGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String deployTo;
        private String deployType;
        private String description;
        private String id;
        private String name;
        private String sgt;
        public Builder() {}
        public Builder(GetIpToSgtMappingGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deployTo = defaults.deployTo;
    	      this.deployType = defaults.deployType;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.sgt = defaults.sgt;
        }

        @CustomType.Setter
        public Builder deployTo(String deployTo) {
            if (deployTo == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "deployTo");
            }
            this.deployTo = deployTo;
            return this;
        }
        @CustomType.Setter
        public Builder deployType(String deployType) {
            if (deployType == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "deployType");
            }
            this.deployType = deployType;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sgt(String sgt) {
            if (sgt == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingGroupResult", "sgt");
            }
            this.sgt = sgt;
            return this;
        }
        public GetIpToSgtMappingGroupResult build() {
            final var _resultValue = new GetIpToSgtMappingGroupResult();
            _resultValue.deployTo = deployTo;
            _resultValue.deployType = deployType;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.sgt = sgt;
            return _resultValue;
        }
    }
}

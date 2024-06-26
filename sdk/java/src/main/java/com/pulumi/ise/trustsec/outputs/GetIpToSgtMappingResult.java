// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpToSgtMappingResult {
    /**
     * @return Mandatory unless `mapping_group` is set or unless `deploy_type` is `ALL`
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
     * @return Mandatory if `host_name` is empty
     * 
     */
    private String hostIp;
    /**
     * @return Mandatory if `host_ip` is empty
     * 
     */
    private String hostName;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deploy_to` and `deploy_type` are set
     * 
     */
    private String mappingGroup;
    /**
     * @return The name of the IP to SGT mapping
     * 
     */
    private String name;
    /**
     * @return Trustsec Security Group ID. Mandatory unless `mapping_group` is set
     * 
     */
    private String sgt;

    private GetIpToSgtMappingResult() {}
    /**
     * @return Mandatory unless `mapping_group` is set or unless `deploy_type` is `ALL`
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
     * @return Mandatory if `host_name` is empty
     * 
     */
    public String hostIp() {
        return this.hostIp;
    }
    /**
     * @return Mandatory if `host_ip` is empty
     * 
     */
    public String hostName() {
        return this.hostName;
    }
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deploy_to` and `deploy_type` are set
     * 
     */
    public String mappingGroup() {
        return this.mappingGroup;
    }
    /**
     * @return The name of the IP to SGT mapping
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Trustsec Security Group ID. Mandatory unless `mapping_group` is set
     * 
     */
    public String sgt() {
        return this.sgt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpToSgtMappingResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String deployTo;
        private String deployType;
        private String description;
        private String hostIp;
        private String hostName;
        private String id;
        private String mappingGroup;
        private String name;
        private String sgt;
        public Builder() {}
        public Builder(GetIpToSgtMappingResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deployTo = defaults.deployTo;
    	      this.deployType = defaults.deployType;
    	      this.description = defaults.description;
    	      this.hostIp = defaults.hostIp;
    	      this.hostName = defaults.hostName;
    	      this.id = defaults.id;
    	      this.mappingGroup = defaults.mappingGroup;
    	      this.name = defaults.name;
    	      this.sgt = defaults.sgt;
        }

        @CustomType.Setter
        public Builder deployTo(String deployTo) {
            if (deployTo == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "deployTo");
            }
            this.deployTo = deployTo;
            return this;
        }
        @CustomType.Setter
        public Builder deployType(String deployType) {
            if (deployType == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "deployType");
            }
            this.deployType = deployType;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder hostIp(String hostIp) {
            if (hostIp == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "hostIp");
            }
            this.hostIp = hostIp;
            return this;
        }
        @CustomType.Setter
        public Builder hostName(String hostName) {
            if (hostName == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "hostName");
            }
            this.hostName = hostName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder mappingGroup(String mappingGroup) {
            if (mappingGroup == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "mappingGroup");
            }
            this.mappingGroup = mappingGroup;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sgt(String sgt) {
            if (sgt == null) {
              throw new MissingRequiredPropertyException("GetIpToSgtMappingResult", "sgt");
            }
            this.sgt = sgt;
            return this;
        }
        public GetIpToSgtMappingResult build() {
            final var _resultValue = new GetIpToSgtMappingResult();
            _resultValue.deployTo = deployTo;
            _resultValue.deployType = deployType;
            _resultValue.description = description;
            _resultValue.hostIp = hostIp;
            _resultValue.hostName = hostName;
            _resultValue.id = id;
            _resultValue.mappingGroup = mappingGroup;
            _resultValue.name = name;
            _resultValue.sgt = sgt;
            return _resultValue;
        }
    }
}

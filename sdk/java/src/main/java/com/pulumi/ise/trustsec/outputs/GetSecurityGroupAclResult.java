// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSecurityGroupAclResult {
    /**
     * @return Content of ACL
     * 
     */
    private String aclContent;
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
     * @return IP Version
     * 
     */
    private String ipVersion;
    /**
     * @return The name of the security group ACL
     * 
     */
    private String name;
    /**
     * @return Read-only
     * 
     */
    private Boolean readOnly;

    private GetSecurityGroupAclResult() {}
    /**
     * @return Content of ACL
     * 
     */
    public String aclContent() {
        return this.aclContent;
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
     * @return IP Version
     * 
     */
    public String ipVersion() {
        return this.ipVersion;
    }
    /**
     * @return The name of the security group ACL
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Read-only
     * 
     */
    public Boolean readOnly() {
        return this.readOnly;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecurityGroupAclResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String aclContent;
        private String description;
        private String id;
        private String ipVersion;
        private String name;
        private Boolean readOnly;
        public Builder() {}
        public Builder(GetSecurityGroupAclResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aclContent = defaults.aclContent;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.ipVersion = defaults.ipVersion;
    	      this.name = defaults.name;
    	      this.readOnly = defaults.readOnly;
        }

        @CustomType.Setter
        public Builder aclContent(String aclContent) {
            if (aclContent == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "aclContent");
            }
            this.aclContent = aclContent;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipVersion(String ipVersion) {
            if (ipVersion == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "ipVersion");
            }
            this.ipVersion = ipVersion;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder readOnly(Boolean readOnly) {
            if (readOnly == null) {
              throw new MissingRequiredPropertyException("GetSecurityGroupAclResult", "readOnly");
            }
            this.readOnly = readOnly;
            return this;
        }
        public GetSecurityGroupAclResult build() {
            final var _resultValue = new GetSecurityGroupAclResult();
            _resultValue.aclContent = aclContent;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.ipVersion = ipVersion;
            _resultValue.name = name;
            _resultValue.readOnly = readOnly;
            return _resultValue;
        }
    }
}

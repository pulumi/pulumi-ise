// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSxpDomainFilterResult {
    /**
     * @return Description
     * 
     */
    private String description;
    /**
     * @return List of SXP Domains, separated with comma
     * 
     */
    private String domains;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return Resource name
     * 
     */
    private String name;
    /**
     * @return SGT name or ID. At least one of subnet or sgt or vn should be defined
     * 
     */
    private String sgt;
    /**
     * @return Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
     * 
     */
    private String subnet;
    /**
     * @return Virtual Network. At least one of subnet or sgt or vn should be defined
     * 
     */
    private String vn;

    private GetSxpDomainFilterResult() {}
    /**
     * @return Description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return List of SXP Domains, separated with comma
     * 
     */
    public String domains() {
        return this.domains;
    }
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Resource name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return SGT name or ID. At least one of subnet or sgt or vn should be defined
     * 
     */
    public String sgt() {
        return this.sgt;
    }
    /**
     * @return Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
     * 
     */
    public String subnet() {
        return this.subnet;
    }
    /**
     * @return Virtual Network. At least one of subnet or sgt or vn should be defined
     * 
     */
    public String vn() {
        return this.vn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSxpDomainFilterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String domains;
        private String id;
        private String name;
        private String sgt;
        private String subnet;
        private String vn;
        public Builder() {}
        public Builder(GetSxpDomainFilterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.domains = defaults.domains;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.sgt = defaults.sgt;
    	      this.subnet = defaults.subnet;
    	      this.vn = defaults.vn;
        }

        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder domains(String domains) {
            if (domains == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "domains");
            }
            this.domains = domains;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sgt(String sgt) {
            if (sgt == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "sgt");
            }
            this.sgt = sgt;
            return this;
        }
        @CustomType.Setter
        public Builder subnet(String subnet) {
            if (subnet == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "subnet");
            }
            this.subnet = subnet;
            return this;
        }
        @CustomType.Setter
        public Builder vn(String vn) {
            if (vn == null) {
              throw new MissingRequiredPropertyException("GetSxpDomainFilterResult", "vn");
            }
            this.vn = vn;
            return this;
        }
        public GetSxpDomainFilterResult build() {
            final var _resultValue = new GetSxpDomainFilterResult();
            _resultValue.description = description;
            _resultValue.domains = domains;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.sgt = sgt;
            _resultValue.subnet = subnet;
            _resultValue.vn = vn;
            return _resultValue;
        }
    }
}

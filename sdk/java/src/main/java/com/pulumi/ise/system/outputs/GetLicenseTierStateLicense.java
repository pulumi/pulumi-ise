// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.system.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLicenseTierStateLicense {
    /**
     * @return License name
     * 
     */
    private String name;
    /**
     * @return License status
     * 
     */
    private String status;

    private GetLicenseTierStateLicense() {}
    /**
     * @return License name
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return License status
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLicenseTierStateLicense defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String status;
        public Builder() {}
        public Builder(GetLicenseTierStateLicense defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetLicenseTierStateLicense", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetLicenseTierStateLicense", "status");
            }
            this.status = status;
            return this;
        }
        public GetLicenseTierStateLicense build() {
            final var _resultValue = new GetLicenseTierStateLicense();
            _resultValue.name = name;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.system.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.system.outputs.GetLicenseTierStateLicense;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetLicenseTierStateResult {
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return List of licenses
     * 
     */
    private List<GetLicenseTierStateLicense> licenses;

    private GetLicenseTierStateResult() {}
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return List of licenses
     * 
     */
    public List<GetLicenseTierStateLicense> licenses() {
        return this.licenses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLicenseTierStateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<GetLicenseTierStateLicense> licenses;
        public Builder() {}
        public Builder(GetLicenseTierStateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.licenses = defaults.licenses;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLicenseTierStateResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder licenses(List<GetLicenseTierStateLicense> licenses) {
            if (licenses == null) {
              throw new MissingRequiredPropertyException("GetLicenseTierStateResult", "licenses");
            }
            this.licenses = licenses;
            return this;
        }
        public Builder licenses(GetLicenseTierStateLicense... licenses) {
            return licenses(List.of(licenses));
        }
        public GetLicenseTierStateResult build() {
            final var _resultValue = new GetLicenseTierStateResult();
            _resultValue.id = id;
            _resultValue.licenses = licenses;
            return _resultValue;
        }
    }
}
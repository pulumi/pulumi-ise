// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.identitymanagement.outputs.GetIdentitySourceSequenceIdentitySource;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIdentitySourceSequenceResult {
    /**
     * @return Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     * 
     */
    private Boolean breakOnStoreFail;
    /**
     * @return Certificate Authentication Profile, empty if doesn&#39;t exist
     * 
     */
    private String certificateAuthenticationProfile;
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
    private List<GetIdentitySourceSequenceIdentitySource> identitySources;
    /**
     * @return The name of the identity source sequence
     * 
     */
    private String name;

    private GetIdentitySourceSequenceResult() {}
    /**
     * @return Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     * 
     */
    public Boolean breakOnStoreFail() {
        return this.breakOnStoreFail;
    }
    /**
     * @return Certificate Authentication Profile, empty if doesn&#39;t exist
     * 
     */
    public String certificateAuthenticationProfile() {
        return this.certificateAuthenticationProfile;
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
    public List<GetIdentitySourceSequenceIdentitySource> identitySources() {
        return this.identitySources;
    }
    /**
     * @return The name of the identity source sequence
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentitySourceSequenceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean breakOnStoreFail;
        private String certificateAuthenticationProfile;
        private String description;
        private String id;
        private List<GetIdentitySourceSequenceIdentitySource> identitySources;
        private String name;
        public Builder() {}
        public Builder(GetIdentitySourceSequenceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.breakOnStoreFail = defaults.breakOnStoreFail;
    	      this.certificateAuthenticationProfile = defaults.certificateAuthenticationProfile;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.identitySources = defaults.identitySources;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder breakOnStoreFail(Boolean breakOnStoreFail) {
            if (breakOnStoreFail == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "breakOnStoreFail");
            }
            this.breakOnStoreFail = breakOnStoreFail;
            return this;
        }
        @CustomType.Setter
        public Builder certificateAuthenticationProfile(String certificateAuthenticationProfile) {
            if (certificateAuthenticationProfile == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "certificateAuthenticationProfile");
            }
            this.certificateAuthenticationProfile = certificateAuthenticationProfile;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder identitySources(List<GetIdentitySourceSequenceIdentitySource> identitySources) {
            if (identitySources == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "identitySources");
            }
            this.identitySources = identitySources;
            return this;
        }
        public Builder identitySources(GetIdentitySourceSequenceIdentitySource... identitySources) {
            return identitySources(List.of(identitySources));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIdentitySourceSequenceResult", "name");
            }
            this.name = name;
            return this;
        }
        public GetIdentitySourceSequenceResult build() {
            final var _resultValue = new GetIdentitySourceSequenceResult();
            _resultValue.breakOnStoreFail = breakOnStoreFail;
            _resultValue.certificateAuthenticationProfile = certificateAuthenticationProfile;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.identitySources = identitySources;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}

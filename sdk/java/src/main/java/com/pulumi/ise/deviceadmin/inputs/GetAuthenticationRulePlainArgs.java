// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthenticationRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthenticationRulePlainArgs Empty = new GetAuthenticationRulePlainArgs();

    /**
     * The id of the object
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return The id of the object
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Policy set ID
     * 
     */
    @Import(name="policySetId", required=true)
    private String policySetId;

    /**
     * @return Policy set ID
     * 
     */
    public String policySetId() {
        return this.policySetId;
    }

    private GetAuthenticationRulePlainArgs() {}

    private GetAuthenticationRulePlainArgs(GetAuthenticationRulePlainArgs $) {
        this.id = $.id;
        this.name = $.name;
        this.policySetId = $.policySetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthenticationRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthenticationRulePlainArgs $;

        public Builder() {
            $ = new GetAuthenticationRulePlainArgs();
        }

        public Builder(GetAuthenticationRulePlainArgs defaults) {
            $ = new GetAuthenticationRulePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The id of the object
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        /**
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(String policySetId) {
            $.policySetId = policySetId;
            return this;
        }

        public GetAuthenticationRulePlainArgs build() {
            if ($.policySetId == null) {
                throw new MissingRequiredPropertyException("GetAuthenticationRulePlainArgs", "policySetId");
            }
            return $;
        }
    }

}

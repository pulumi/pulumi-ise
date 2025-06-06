// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.deviceadmin.inputs.AuthorizationRuleUpdateRanksRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthorizationRuleUpdateRanksArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationRuleUpdateRanksArgs Empty = new AuthorizationRuleUpdateRanksArgs();

    /**
     * Policy set ID
     * 
     */
    @Import(name="policySetId", required=true)
    private Output<String> policySetId;

    /**
     * @return Policy set ID
     * 
     */
    public Output<String> policySetId() {
        return this.policySetId;
    }

    @Import(name="rules")
    private @Nullable Output<List<AuthorizationRuleUpdateRanksRuleArgs>> rules;

    public Optional<Output<List<AuthorizationRuleUpdateRanksRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    private AuthorizationRuleUpdateRanksArgs() {}

    private AuthorizationRuleUpdateRanksArgs(AuthorizationRuleUpdateRanksArgs $) {
        this.policySetId = $.policySetId;
        this.rules = $.rules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationRuleUpdateRanksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationRuleUpdateRanksArgs $;

        public Builder() {
            $ = new AuthorizationRuleUpdateRanksArgs();
        }

        public Builder(AuthorizationRuleUpdateRanksArgs defaults) {
            $ = new AuthorizationRuleUpdateRanksArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(Output<String> policySetId) {
            $.policySetId = policySetId;
            return this;
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(String policySetId) {
            return policySetId(Output.of(policySetId));
        }

        public Builder rules(@Nullable Output<List<AuthorizationRuleUpdateRanksRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        public Builder rules(List<AuthorizationRuleUpdateRanksRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        public Builder rules(AuthorizationRuleUpdateRanksRuleArgs... rules) {
            return rules(List.of(rules));
        }

        public AuthorizationRuleUpdateRanksArgs build() {
            if ($.policySetId == null) {
                throw new MissingRequiredPropertyException("AuthorizationRuleUpdateRanksArgs", "policySetId");
            }
            return $;
        }
    }

}

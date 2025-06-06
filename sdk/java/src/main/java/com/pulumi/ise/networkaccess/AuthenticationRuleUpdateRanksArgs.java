// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.networkaccess.inputs.AuthenticationRuleUpdateRanksRuleArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthenticationRuleUpdateRanksArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthenticationRuleUpdateRanksArgs Empty = new AuthenticationRuleUpdateRanksArgs();

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
    private @Nullable Output<List<AuthenticationRuleUpdateRanksRuleArgs>> rules;

    public Optional<Output<List<AuthenticationRuleUpdateRanksRuleArgs>>> rules() {
        return Optional.ofNullable(this.rules);
    }

    private AuthenticationRuleUpdateRanksArgs() {}

    private AuthenticationRuleUpdateRanksArgs(AuthenticationRuleUpdateRanksArgs $) {
        this.policySetId = $.policySetId;
        this.rules = $.rules;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthenticationRuleUpdateRanksArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthenticationRuleUpdateRanksArgs $;

        public Builder() {
            $ = new AuthenticationRuleUpdateRanksArgs();
        }

        public Builder(AuthenticationRuleUpdateRanksArgs defaults) {
            $ = new AuthenticationRuleUpdateRanksArgs(Objects.requireNonNull(defaults));
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

        public Builder rules(@Nullable Output<List<AuthenticationRuleUpdateRanksRuleArgs>> rules) {
            $.rules = rules;
            return this;
        }

        public Builder rules(List<AuthenticationRuleUpdateRanksRuleArgs> rules) {
            return rules(Output.of(rules));
        }

        public Builder rules(AuthenticationRuleUpdateRanksRuleArgs... rules) {
            return rules(List.of(rules));
        }

        public AuthenticationRuleUpdateRanksArgs build() {
            if ($.policySetId == null) {
                throw new MissingRequiredPropertyException("AuthenticationRuleUpdateRanksArgs", "policySetId");
            }
            return $;
        }
    }

}

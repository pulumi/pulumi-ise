// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthorizationRuleUpdateRankState extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationRuleUpdateRankState Empty = new AuthorizationRuleUpdateRankState();

    /**
     * Policy set ID
     * 
     */
    @Import(name="policySetId")
    private @Nullable Output<String> policySetId;

    /**
     * @return Policy set ID
     * 
     */
    public Optional<Output<String>> policySetId() {
        return Optional.ofNullable(this.policySetId);
    }

    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    @Import(name="rank")
    private @Nullable Output<Integer> rank;

    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Optional<Output<Integer>> rank() {
        return Optional.ofNullable(this.rank);
    }

    /**
     * Authorization rule ID
     * 
     */
    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    /**
     * @return Authorization rule ID
     * 
     */
    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    private AuthorizationRuleUpdateRankState() {}

    private AuthorizationRuleUpdateRankState(AuthorizationRuleUpdateRankState $) {
        this.policySetId = $.policySetId;
        this.rank = $.rank;
        this.ruleId = $.ruleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationRuleUpdateRankState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationRuleUpdateRankState $;

        public Builder() {
            $ = new AuthorizationRuleUpdateRankState();
        }

        public Builder(AuthorizationRuleUpdateRankState defaults) {
            $ = new AuthorizationRuleUpdateRankState(Objects.requireNonNull(defaults));
        }

        /**
         * @param policySetId Policy set ID
         * 
         * @return builder
         * 
         */
        public Builder policySetId(@Nullable Output<String> policySetId) {
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

        /**
         * @param rank The rank (priority) in relation to other rules. Lower rank is higher priority.
         * 
         * @return builder
         * 
         */
        public Builder rank(@Nullable Output<Integer> rank) {
            $.rank = rank;
            return this;
        }

        /**
         * @param rank The rank (priority) in relation to other rules. Lower rank is higher priority.
         * 
         * @return builder
         * 
         */
        public Builder rank(Integer rank) {
            return rank(Output.of(rank));
        }

        /**
         * @param ruleId Authorization rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId Authorization rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        public AuthorizationRuleUpdateRankState build() {
            return $;
        }
    }

}

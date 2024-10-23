// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AuthorizationExceptionRuleUpdateRankArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationExceptionRuleUpdateRankArgs Empty = new AuthorizationExceptionRuleUpdateRankArgs();

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

    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    @Import(name="rank", required=true)
    private Output<Integer> rank;

    /**
     * @return The rank (priority) in relation to other rules. Lower rank is higher priority.
     * 
     */
    public Output<Integer> rank() {
        return this.rank;
    }

    /**
     * Authorization exception rule ID
     * 
     */
    @Import(name="ruleId", required=true)
    private Output<String> ruleId;

    /**
     * @return Authorization exception rule ID
     * 
     */
    public Output<String> ruleId() {
        return this.ruleId;
    }

    private AuthorizationExceptionRuleUpdateRankArgs() {}

    private AuthorizationExceptionRuleUpdateRankArgs(AuthorizationExceptionRuleUpdateRankArgs $) {
        this.policySetId = $.policySetId;
        this.rank = $.rank;
        this.ruleId = $.ruleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationExceptionRuleUpdateRankArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationExceptionRuleUpdateRankArgs $;

        public Builder() {
            $ = new AuthorizationExceptionRuleUpdateRankArgs();
        }

        public Builder(AuthorizationExceptionRuleUpdateRankArgs defaults) {
            $ = new AuthorizationExceptionRuleUpdateRankArgs(Objects.requireNonNull(defaults));
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

        /**
         * @param rank The rank (priority) in relation to other rules. Lower rank is higher priority.
         * 
         * @return builder
         * 
         */
        public Builder rank(Output<Integer> rank) {
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
         * @param ruleId Authorization exception rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId Authorization exception rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        public AuthorizationExceptionRuleUpdateRankArgs build() {
            if ($.policySetId == null) {
                throw new MissingRequiredPropertyException("AuthorizationExceptionRuleUpdateRankArgs", "policySetId");
            }
            if ($.rank == null) {
                throw new MissingRequiredPropertyException("AuthorizationExceptionRuleUpdateRankArgs", "rank");
            }
            if ($.ruleId == null) {
                throw new MissingRequiredPropertyException("AuthorizationExceptionRuleUpdateRankArgs", "ruleId");
            }
            return $;
        }
    }

}

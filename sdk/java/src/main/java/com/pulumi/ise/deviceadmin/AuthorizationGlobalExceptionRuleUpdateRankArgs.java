// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class AuthorizationGlobalExceptionRuleUpdateRankArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationGlobalExceptionRuleUpdateRankArgs Empty = new AuthorizationGlobalExceptionRuleUpdateRankArgs();

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
     * Authorization global exception rule ID
     * 
     */
    @Import(name="ruleId", required=true)
    private Output<String> ruleId;

    /**
     * @return Authorization global exception rule ID
     * 
     */
    public Output<String> ruleId() {
        return this.ruleId;
    }

    private AuthorizationGlobalExceptionRuleUpdateRankArgs() {}

    private AuthorizationGlobalExceptionRuleUpdateRankArgs(AuthorizationGlobalExceptionRuleUpdateRankArgs $) {
        this.rank = $.rank;
        this.ruleId = $.ruleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationGlobalExceptionRuleUpdateRankArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationGlobalExceptionRuleUpdateRankArgs $;

        public Builder() {
            $ = new AuthorizationGlobalExceptionRuleUpdateRankArgs();
        }

        public Builder(AuthorizationGlobalExceptionRuleUpdateRankArgs defaults) {
            $ = new AuthorizationGlobalExceptionRuleUpdateRankArgs(Objects.requireNonNull(defaults));
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
         * @param ruleId Authorization global exception rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(Output<String> ruleId) {
            $.ruleId = ruleId;
            return this;
        }

        /**
         * @param ruleId Authorization global exception rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(String ruleId) {
            return ruleId(Output.of(ruleId));
        }

        public AuthorizationGlobalExceptionRuleUpdateRankArgs build() {
            if ($.rank == null) {
                throw new MissingRequiredPropertyException("AuthorizationGlobalExceptionRuleUpdateRankArgs", "rank");
            }
            if ($.ruleId == null) {
                throw new MissingRequiredPropertyException("AuthorizationGlobalExceptionRuleUpdateRankArgs", "ruleId");
            }
            return $;
        }
    }

}
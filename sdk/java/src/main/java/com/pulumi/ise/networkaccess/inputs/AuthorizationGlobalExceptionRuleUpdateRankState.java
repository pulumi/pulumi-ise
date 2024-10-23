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


public final class AuthorizationGlobalExceptionRuleUpdateRankState extends com.pulumi.resources.ResourceArgs {

    public static final AuthorizationGlobalExceptionRuleUpdateRankState Empty = new AuthorizationGlobalExceptionRuleUpdateRankState();

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
     * Authorization global exception rule ID
     * 
     */
    @Import(name="ruleId")
    private @Nullable Output<String> ruleId;

    /**
     * @return Authorization global exception rule ID
     * 
     */
    public Optional<Output<String>> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }

    private AuthorizationGlobalExceptionRuleUpdateRankState() {}

    private AuthorizationGlobalExceptionRuleUpdateRankState(AuthorizationGlobalExceptionRuleUpdateRankState $) {
        this.rank = $.rank;
        this.ruleId = $.ruleId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthorizationGlobalExceptionRuleUpdateRankState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthorizationGlobalExceptionRuleUpdateRankState $;

        public Builder() {
            $ = new AuthorizationGlobalExceptionRuleUpdateRankState();
        }

        public Builder(AuthorizationGlobalExceptionRuleUpdateRankState defaults) {
            $ = new AuthorizationGlobalExceptionRuleUpdateRankState(Objects.requireNonNull(defaults));
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
         * @param ruleId Authorization global exception rule ID
         * 
         * @return builder
         * 
         */
        public Builder ruleId(@Nullable Output<String> ruleId) {
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

        public AuthorizationGlobalExceptionRuleUpdateRankState build() {
            return $;
        }
    }

}

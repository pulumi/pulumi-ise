// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ActiveDirectoryJoinPointRewriteRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActiveDirectoryJoinPointRewriteRuleArgs Empty = new ActiveDirectoryJoinPointRewriteRuleArgs();

    /**
     * Required for each rule in the list with no duplication between rules
     * 
     */
    @Import(name="rewriteMatch", required=true)
    private Output<String> rewriteMatch;

    /**
     * @return Required for each rule in the list with no duplication between rules
     * 
     */
    public Output<String> rewriteMatch() {
        return this.rewriteMatch;
    }

    /**
     * Required for each rule in the list
     * 
     */
    @Import(name="rewriteResult", required=true)
    private Output<String> rewriteResult;

    /**
     * @return Required for each rule in the list
     * 
     */
    public Output<String> rewriteResult() {
        return this.rewriteResult;
    }

    /**
     * Required for each rule in the list in serial order
     * 
     */
    @Import(name="rowId", required=true)
    private Output<String> rowId;

    /**
     * @return Required for each rule in the list in serial order
     * 
     */
    public Output<String> rowId() {
        return this.rowId;
    }

    private ActiveDirectoryJoinPointRewriteRuleArgs() {}

    private ActiveDirectoryJoinPointRewriteRuleArgs(ActiveDirectoryJoinPointRewriteRuleArgs $) {
        this.rewriteMatch = $.rewriteMatch;
        this.rewriteResult = $.rewriteResult;
        this.rowId = $.rowId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActiveDirectoryJoinPointRewriteRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActiveDirectoryJoinPointRewriteRuleArgs $;

        public Builder() {
            $ = new ActiveDirectoryJoinPointRewriteRuleArgs();
        }

        public Builder(ActiveDirectoryJoinPointRewriteRuleArgs defaults) {
            $ = new ActiveDirectoryJoinPointRewriteRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param rewriteMatch Required for each rule in the list with no duplication between rules
         * 
         * @return builder
         * 
         */
        public Builder rewriteMatch(Output<String> rewriteMatch) {
            $.rewriteMatch = rewriteMatch;
            return this;
        }

        /**
         * @param rewriteMatch Required for each rule in the list with no duplication between rules
         * 
         * @return builder
         * 
         */
        public Builder rewriteMatch(String rewriteMatch) {
            return rewriteMatch(Output.of(rewriteMatch));
        }

        /**
         * @param rewriteResult Required for each rule in the list
         * 
         * @return builder
         * 
         */
        public Builder rewriteResult(Output<String> rewriteResult) {
            $.rewriteResult = rewriteResult;
            return this;
        }

        /**
         * @param rewriteResult Required for each rule in the list
         * 
         * @return builder
         * 
         */
        public Builder rewriteResult(String rewriteResult) {
            return rewriteResult(Output.of(rewriteResult));
        }

        /**
         * @param rowId Required for each rule in the list in serial order
         * 
         * @return builder
         * 
         */
        public Builder rowId(Output<String> rowId) {
            $.rowId = rowId;
            return this;
        }

        /**
         * @param rowId Required for each rule in the list in serial order
         * 
         * @return builder
         * 
         */
        public Builder rowId(String rowId) {
            return rowId(Output.of(rowId));
        }

        public ActiveDirectoryJoinPointRewriteRuleArgs build() {
            if ($.rewriteMatch == null) {
                throw new MissingRequiredPropertyException("ActiveDirectoryJoinPointRewriteRuleArgs", "rewriteMatch");
            }
            if ($.rewriteResult == null) {
                throw new MissingRequiredPropertyException("ActiveDirectoryJoinPointRewriteRuleArgs", "rewriteResult");
            }
            if ($.rowId == null) {
                throw new MissingRequiredPropertyException("ActiveDirectoryJoinPointRewriteRuleArgs", "rowId");
            }
            return $;
        }
    }

}
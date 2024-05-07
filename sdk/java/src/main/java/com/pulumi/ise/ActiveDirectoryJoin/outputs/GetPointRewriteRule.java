// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.ActiveDirectoryJoin.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPointRewriteRule {
    /**
     * @return Required for each rule in the list with no duplication between rules
     * 
     */
    private String rewriteMatch;
    /**
     * @return Required for each rule in the list
     * 
     */
    private String rewriteResult;
    /**
     * @return Required for each rule in the list in serial order
     * 
     */
    private String rowId;

    private GetPointRewriteRule() {}
    /**
     * @return Required for each rule in the list with no duplication between rules
     * 
     */
    public String rewriteMatch() {
        return this.rewriteMatch;
    }
    /**
     * @return Required for each rule in the list
     * 
     */
    public String rewriteResult() {
        return this.rewriteResult;
    }
    /**
     * @return Required for each rule in the list in serial order
     * 
     */
    public String rowId() {
        return this.rowId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPointRewriteRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String rewriteMatch;
        private String rewriteResult;
        private String rowId;
        public Builder() {}
        public Builder(GetPointRewriteRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.rewriteMatch = defaults.rewriteMatch;
    	      this.rewriteResult = defaults.rewriteResult;
    	      this.rowId = defaults.rowId;
        }

        @CustomType.Setter
        public Builder rewriteMatch(String rewriteMatch) {
            if (rewriteMatch == null) {
              throw new MissingRequiredPropertyException("GetPointRewriteRule", "rewriteMatch");
            }
            this.rewriteMatch = rewriteMatch;
            return this;
        }
        @CustomType.Setter
        public Builder rewriteResult(String rewriteResult) {
            if (rewriteResult == null) {
              throw new MissingRequiredPropertyException("GetPointRewriteRule", "rewriteResult");
            }
            this.rewriteResult = rewriteResult;
            return this;
        }
        @CustomType.Setter
        public Builder rowId(String rowId) {
            if (rowId == null) {
              throw new MissingRequiredPropertyException("GetPointRewriteRule", "rowId");
            }
            this.rowId = rowId;
            return this;
        }
        public GetPointRewriteRule build() {
            final var _resultValue = new GetPointRewriteRule();
            _resultValue.rewriteMatch = rewriteMatch;
            _resultValue.rewriteResult = rewriteResult;
            _resultValue.rowId = rowId;
            return _resultValue;
        }
    }
}

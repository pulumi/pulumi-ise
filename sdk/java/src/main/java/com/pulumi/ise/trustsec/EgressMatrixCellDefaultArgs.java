// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EgressMatrixCellDefaultArgs extends com.pulumi.resources.ResourceArgs {

    public static final EgressMatrixCellDefaultArgs Empty = new EgressMatrixCellDefaultArgs();

    /**
     * Can be used only if sgacls not specified. Final Catch All Rule - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
     * 
     */
    @Import(name="defaultRule")
    private @Nullable Output<String> defaultRule;

    /**
     * @return Can be used only if sgacls not specified. Final Catch All Rule - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
     * 
     */
    public Optional<Output<String>> defaultRule() {
        return Optional.ofNullable(this.defaultRule);
    }

    /**
     * Description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
     * 
     */
    @Import(name="matrixCellStatus")
    private @Nullable Output<String> matrixCellStatus;

    /**
     * @return Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
     * 
     */
    public Optional<Output<String>> matrixCellStatus() {
        return Optional.ofNullable(this.matrixCellStatus);
    }

    /**
     * List of TrustSec Security Groups ACLs
     * 
     */
    @Import(name="sgacls")
    private @Nullable Output<List<String>> sgacls;

    /**
     * @return List of TrustSec Security Groups ACLs
     * 
     */
    public Optional<Output<List<String>>> sgacls() {
        return Optional.ofNullable(this.sgacls);
    }

    private EgressMatrixCellDefaultArgs() {}

    private EgressMatrixCellDefaultArgs(EgressMatrixCellDefaultArgs $) {
        this.defaultRule = $.defaultRule;
        this.description = $.description;
        this.matrixCellStatus = $.matrixCellStatus;
        this.sgacls = $.sgacls;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EgressMatrixCellDefaultArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EgressMatrixCellDefaultArgs $;

        public Builder() {
            $ = new EgressMatrixCellDefaultArgs();
        }

        public Builder(EgressMatrixCellDefaultArgs defaults) {
            $ = new EgressMatrixCellDefaultArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultRule Can be used only if sgacls not specified. Final Catch All Rule - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
         * 
         * @return builder
         * 
         */
        public Builder defaultRule(@Nullable Output<String> defaultRule) {
            $.defaultRule = defaultRule;
            return this;
        }

        /**
         * @param defaultRule Can be used only if sgacls not specified. Final Catch All Rule - Choices: `NONE`, `DENY_IP`, `PERMIT_IP`
         * 
         * @return builder
         * 
         */
        public Builder defaultRule(String defaultRule) {
            return defaultRule(Output.of(defaultRule));
        }

        /**
         * @param description Description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param matrixCellStatus Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder matrixCellStatus(@Nullable Output<String> matrixCellStatus) {
            $.matrixCellStatus = matrixCellStatus;
            return this;
        }

        /**
         * @param matrixCellStatus Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder matrixCellStatus(String matrixCellStatus) {
            return matrixCellStatus(Output.of(matrixCellStatus));
        }

        /**
         * @param sgacls List of TrustSec Security Groups ACLs
         * 
         * @return builder
         * 
         */
        public Builder sgacls(@Nullable Output<List<String>> sgacls) {
            $.sgacls = sgacls;
            return this;
        }

        /**
         * @param sgacls List of TrustSec Security Groups ACLs
         * 
         * @return builder
         * 
         */
        public Builder sgacls(List<String> sgacls) {
            return sgacls(Output.of(sgacls));
        }

        /**
         * @param sgacls List of TrustSec Security Groups ACLs
         * 
         * @return builder
         * 
         */
        public Builder sgacls(String... sgacls) {
            return sgacls(List.of(sgacls));
        }

        public EgressMatrixCellDefaultArgs build() {
            return $;
        }
    }

}

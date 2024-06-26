// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EgressMatrixCellArgs extends com.pulumi.resources.ResourceArgs {

    public static final EgressMatrixCellArgs Empty = new EgressMatrixCellArgs();

    /**
     * Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
     * 
     */
    @Import(name="defaultRule")
    private @Nullable Output<String> defaultRule;

    /**
     * @return Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
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
     * Destination Trustsec Security Group ID
     * 
     */
    @Import(name="destinationSgtId", required=true)
    private Output<String> destinationSgtId;

    /**
     * @return Destination Trustsec Security Group ID
     * 
     */
    public Output<String> destinationSgtId() {
        return this.destinationSgtId;
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

    /**
     * Source Trustsec Security Group ID
     * 
     */
    @Import(name="sourceSgtId", required=true)
    private Output<String> sourceSgtId;

    /**
     * @return Source Trustsec Security Group ID
     * 
     */
    public Output<String> sourceSgtId() {
        return this.sourceSgtId;
    }

    private EgressMatrixCellArgs() {}

    private EgressMatrixCellArgs(EgressMatrixCellArgs $) {
        this.defaultRule = $.defaultRule;
        this.description = $.description;
        this.destinationSgtId = $.destinationSgtId;
        this.matrixCellStatus = $.matrixCellStatus;
        this.sgacls = $.sgacls;
        this.sourceSgtId = $.sourceSgtId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EgressMatrixCellArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EgressMatrixCellArgs $;

        public Builder() {
            $ = new EgressMatrixCellArgs();
        }

        public Builder(EgressMatrixCellArgs defaults) {
            $ = new EgressMatrixCellArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param defaultRule Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
         * 
         * @return builder
         * 
         */
        public Builder defaultRule(@Nullable Output<String> defaultRule) {
            $.defaultRule = defaultRule;
            return this;
        }

        /**
         * @param defaultRule Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
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
         * @param destinationSgtId Destination Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder destinationSgtId(Output<String> destinationSgtId) {
            $.destinationSgtId = destinationSgtId;
            return this;
        }

        /**
         * @param destinationSgtId Destination Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder destinationSgtId(String destinationSgtId) {
            return destinationSgtId(Output.of(destinationSgtId));
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

        /**
         * @param sourceSgtId Source Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder sourceSgtId(Output<String> sourceSgtId) {
            $.sourceSgtId = sourceSgtId;
            return this;
        }

        /**
         * @param sourceSgtId Source Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder sourceSgtId(String sourceSgtId) {
            return sourceSgtId(Output.of(sourceSgtId));
        }

        public EgressMatrixCellArgs build() {
            if ($.destinationSgtId == null) {
                throw new MissingRequiredPropertyException("EgressMatrixCellArgs", "destinationSgtId");
            }
            if ($.sourceSgtId == null) {
                throw new MissingRequiredPropertyException("EgressMatrixCellArgs", "sourceSgtId");
            }
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpToSgtMappingGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpToSgtMappingGroupArgs Empty = new IpToSgtMappingGroupArgs();

    /**
     * Mandatory unless `deploy_type` is `ALL`
     * 
     */
    @Import(name="deployTo")
    private @Nullable Output<String> deployTo;

    /**
     * @return Mandatory unless `deploy_type` is `ALL`
     * 
     */
    public Optional<Output<String>> deployTo() {
        return Optional.ofNullable(this.deployTo);
    }

    /**
     * Deploy Type - Choices: `ALL`, `ND`, `NDG`
     * 
     */
    @Import(name="deployType", required=true)
    private Output<String> deployType;

    /**
     * @return Deploy Type - Choices: `ALL`, `ND`, `NDG`
     * 
     */
    public Output<String> deployType() {
        return this.deployType;
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
     * The name of the IP to SGT mapping Group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the IP to SGT mapping Group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Trustsec Security Group ID
     * 
     */
    @Import(name="sgt", required=true)
    private Output<String> sgt;

    /**
     * @return Trustsec Security Group ID
     * 
     */
    public Output<String> sgt() {
        return this.sgt;
    }

    private IpToSgtMappingGroupArgs() {}

    private IpToSgtMappingGroupArgs(IpToSgtMappingGroupArgs $) {
        this.deployTo = $.deployTo;
        this.deployType = $.deployType;
        this.description = $.description;
        this.name = $.name;
        this.sgt = $.sgt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpToSgtMappingGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpToSgtMappingGroupArgs $;

        public Builder() {
            $ = new IpToSgtMappingGroupArgs();
        }

        public Builder(IpToSgtMappingGroupArgs defaults) {
            $ = new IpToSgtMappingGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deployTo Mandatory unless `deploy_type` is `ALL`
         * 
         * @return builder
         * 
         */
        public Builder deployTo(@Nullable Output<String> deployTo) {
            $.deployTo = deployTo;
            return this;
        }

        /**
         * @param deployTo Mandatory unless `deploy_type` is `ALL`
         * 
         * @return builder
         * 
         */
        public Builder deployTo(String deployTo) {
            return deployTo(Output.of(deployTo));
        }

        /**
         * @param deployType Deploy Type - Choices: `ALL`, `ND`, `NDG`
         * 
         * @return builder
         * 
         */
        public Builder deployType(Output<String> deployType) {
            $.deployType = deployType;
            return this;
        }

        /**
         * @param deployType Deploy Type - Choices: `ALL`, `ND`, `NDG`
         * 
         * @return builder
         * 
         */
        public Builder deployType(String deployType) {
            return deployType(Output.of(deployType));
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
         * @param name The name of the IP to SGT mapping Group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the IP to SGT mapping Group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sgt Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder sgt(Output<String> sgt) {
            $.sgt = sgt;
            return this;
        }

        /**
         * @param sgt Trustsec Security Group ID
         * 
         * @return builder
         * 
         */
        public Builder sgt(String sgt) {
            return sgt(Output.of(sgt));
        }

        public IpToSgtMappingGroupArgs build() {
            if ($.deployType == null) {
                throw new MissingRequiredPropertyException("IpToSgtMappingGroupArgs", "deployType");
            }
            if ($.sgt == null) {
                throw new MissingRequiredPropertyException("IpToSgtMappingGroupArgs", "sgt");
            }
            return $;
        }
    }

}

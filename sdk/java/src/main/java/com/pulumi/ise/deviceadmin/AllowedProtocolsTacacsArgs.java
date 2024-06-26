// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AllowedProtocolsTacacsArgs extends com.pulumi.resources.ResourceArgs {

    public static final AllowedProtocolsTacacsArgs Empty = new AllowedProtocolsTacacsArgs();

    /**
     * Allow CHAP
     * 
     */
    @Import(name="allowChap", required=true)
    private Output<Boolean> allowChap;

    /**
     * @return Allow CHAP
     * 
     */
    public Output<Boolean> allowChap() {
        return this.allowChap;
    }

    /**
     * Allow MS CHAP v1
     * 
     */
    @Import(name="allowMsChapV1", required=true)
    private Output<Boolean> allowMsChapV1;

    /**
     * @return Allow MS CHAP v1
     * 
     */
    public Output<Boolean> allowMsChapV1() {
        return this.allowMsChapV1;
    }

    /**
     * Allow PAP ASCII
     * 
     */
    @Import(name="allowPapAscii", required=true)
    private Output<Boolean> allowPapAscii;

    /**
     * @return Allow PAP ASCII
     * 
     */
    public Output<Boolean> allowPapAscii() {
        return this.allowPapAscii;
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
     * The name of the allowed protocols
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the allowed protocols
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private AllowedProtocolsTacacsArgs() {}

    private AllowedProtocolsTacacsArgs(AllowedProtocolsTacacsArgs $) {
        this.allowChap = $.allowChap;
        this.allowMsChapV1 = $.allowMsChapV1;
        this.allowPapAscii = $.allowPapAscii;
        this.description = $.description;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AllowedProtocolsTacacsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AllowedProtocolsTacacsArgs $;

        public Builder() {
            $ = new AllowedProtocolsTacacsArgs();
        }

        public Builder(AllowedProtocolsTacacsArgs defaults) {
            $ = new AllowedProtocolsTacacsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowChap Allow CHAP
         * 
         * @return builder
         * 
         */
        public Builder allowChap(Output<Boolean> allowChap) {
            $.allowChap = allowChap;
            return this;
        }

        /**
         * @param allowChap Allow CHAP
         * 
         * @return builder
         * 
         */
        public Builder allowChap(Boolean allowChap) {
            return allowChap(Output.of(allowChap));
        }

        /**
         * @param allowMsChapV1 Allow MS CHAP v1
         * 
         * @return builder
         * 
         */
        public Builder allowMsChapV1(Output<Boolean> allowMsChapV1) {
            $.allowMsChapV1 = allowMsChapV1;
            return this;
        }

        /**
         * @param allowMsChapV1 Allow MS CHAP v1
         * 
         * @return builder
         * 
         */
        public Builder allowMsChapV1(Boolean allowMsChapV1) {
            return allowMsChapV1(Output.of(allowMsChapV1));
        }

        /**
         * @param allowPapAscii Allow PAP ASCII
         * 
         * @return builder
         * 
         */
        public Builder allowPapAscii(Output<Boolean> allowPapAscii) {
            $.allowPapAscii = allowPapAscii;
            return this;
        }

        /**
         * @param allowPapAscii Allow PAP ASCII
         * 
         * @return builder
         * 
         */
        public Builder allowPapAscii(Boolean allowPapAscii) {
            return allowPapAscii(Output.of(allowPapAscii));
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
         * @param name The name of the allowed protocols
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the allowed protocols
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public AllowedProtocolsTacacsArgs build() {
            if ($.allowChap == null) {
                throw new MissingRequiredPropertyException("AllowedProtocolsTacacsArgs", "allowChap");
            }
            if ($.allowMsChapV1 == null) {
                throw new MissingRequiredPropertyException("AllowedProtocolsTacacsArgs", "allowMsChapV1");
            }
            if ($.allowPapAscii == null) {
                throw new MissingRequiredPropertyException("AllowedProtocolsTacacsArgs", "allowPapAscii");
            }
            return $;
        }
    }

}

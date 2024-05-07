// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.ActiveDirectoryJoin.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PointGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final PointGroupArgs Empty = new PointGroupArgs();

    /**
     * Required for each group in the group list with no duplication between groups
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Required for each group in the group list with no duplication between groups
     * 
     */
    @Import(name="sid", required=true)
    private Output<String> sid;

    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    public Output<String> sid() {
        return this.sid;
    }

    @Import(name="type")
    private @Nullable Output<String> type;

    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private PointGroupArgs() {}

    private PointGroupArgs(PointGroupArgs $) {
        this.name = $.name;
        this.sid = $.sid;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PointGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PointGroupArgs $;

        public Builder() {
            $ = new PointGroupArgs();
        }

        public Builder(PointGroupArgs defaults) {
            $ = new PointGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Required for each group in the group list with no duplication between groups
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Required for each group in the group list with no duplication between groups
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sid Required for each group in the group list with no duplication between groups
         * 
         * @return builder
         * 
         */
        public Builder sid(Output<String> sid) {
            $.sid = sid;
            return this;
        }

        /**
         * @param sid Required for each group in the group list with no duplication between groups
         * 
         * @return builder
         * 
         */
        public Builder sid(String sid) {
            return sid(Output.of(sid));
        }

        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public PointGroupArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("PointGroupArgs", "name");
            }
            if ($.sid == null) {
                throw new MissingRequiredPropertyException("PointGroupArgs", "sid");
            }
            return $;
        }
    }

}

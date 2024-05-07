// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.DeviceAdmin.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetConditionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetConditionPlainArgs Empty = new GetConditionPlainArgs();

    /**
     * The id of the object
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return The id of the object
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Condition name
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Condition name
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetConditionPlainArgs() {}

    private GetConditionPlainArgs(GetConditionPlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetConditionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetConditionPlainArgs $;

        public Builder() {
            $ = new GetConditionPlainArgs();
        }

        public Builder(GetConditionPlainArgs defaults) {
            $ = new GetConditionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The id of the object
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        /**
         * @param name Condition name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetConditionPlainArgs build() {
            return $;
        }
    }

}

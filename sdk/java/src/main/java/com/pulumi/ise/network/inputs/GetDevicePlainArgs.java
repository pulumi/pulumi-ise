// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.network.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDevicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDevicePlainArgs Empty = new GetDevicePlainArgs();

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
     * The name of the network device
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the network device
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetDevicePlainArgs() {}

    private GetDevicePlainArgs(GetDevicePlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDevicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDevicePlainArgs $;

        public Builder() {
            $ = new GetDevicePlainArgs();
        }

        public Builder(GetDevicePlainArgs defaults) {
            $ = new GetDevicePlainArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the network device
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetDevicePlainArgs build() {
            return $;
        }
    }

}

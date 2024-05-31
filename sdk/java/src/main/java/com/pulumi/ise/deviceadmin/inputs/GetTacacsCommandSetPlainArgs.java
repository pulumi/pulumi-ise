// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetTacacsCommandSetPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetTacacsCommandSetPlainArgs Empty = new GetTacacsCommandSetPlainArgs();

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
     * The name of the TACACS command set
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the TACACS command set
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetTacacsCommandSetPlainArgs() {}

    private GetTacacsCommandSetPlainArgs(GetTacacsCommandSetPlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetTacacsCommandSetPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetTacacsCommandSetPlainArgs $;

        public Builder() {
            $ = new GetTacacsCommandSetPlainArgs();
        }

        public Builder(GetTacacsCommandSetPlainArgs defaults) {
            $ = new GetTacacsCommandSetPlainArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the TACACS command set
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetTacacsCommandSetPlainArgs build() {
            return $;
        }
    }

}
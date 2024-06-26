// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthorizationProfilePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthorizationProfilePlainArgs Empty = new GetAuthorizationProfilePlainArgs();

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
     * The name of the authorization profile
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the authorization profile
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetAuthorizationProfilePlainArgs() {}

    private GetAuthorizationProfilePlainArgs(GetAuthorizationProfilePlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthorizationProfilePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthorizationProfilePlainArgs $;

        public Builder() {
            $ = new GetAuthorizationProfilePlainArgs();
        }

        public Builder(GetAuthorizationProfilePlainArgs defaults) {
            $ = new GetAuthorizationProfilePlainArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the authorization profile
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetAuthorizationProfilePlainArgs build() {
            return $;
        }
    }

}

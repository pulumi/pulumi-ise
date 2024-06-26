// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEndpointIdentityGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEndpointIdentityGroupPlainArgs Empty = new GetEndpointIdentityGroupPlainArgs();

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
     * The name of the endpoint identity group
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the endpoint identity group
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetEndpointIdentityGroupPlainArgs() {}

    private GetEndpointIdentityGroupPlainArgs(GetEndpointIdentityGroupPlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEndpointIdentityGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEndpointIdentityGroupPlainArgs $;

        public Builder() {
            $ = new GetEndpointIdentityGroupPlainArgs();
        }

        public Builder(GetEndpointIdentityGroupPlainArgs defaults) {
            $ = new GetEndpointIdentityGroupPlainArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the endpoint identity group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetEndpointIdentityGroupPlainArgs build() {
            return $;
        }
    }

}

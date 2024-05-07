// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRepositoryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRepositoryArgs Empty = new GetRepositoryArgs();

    /**
     * The id of the object
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return The id of the object
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetRepositoryArgs() {}

    private GetRepositoryArgs(GetRepositoryArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRepositoryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRepositoryArgs $;

        public Builder() {
            $ = new GetRepositoryArgs();
        }

        public Builder(GetRepositoryArgs defaults) {
            $ = new GetRepositoryArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The id of the object
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id The id of the object
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param name Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetRepositoryArgs build() {
            return $;
        }
    }

}

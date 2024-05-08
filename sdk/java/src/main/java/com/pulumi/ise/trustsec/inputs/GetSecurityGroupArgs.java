// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecurityGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecurityGroupArgs Empty = new GetSecurityGroupArgs();

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
     * The name of the security group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the security group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetSecurityGroupArgs() {}

    private GetSecurityGroupArgs(GetSecurityGroupArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecurityGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecurityGroupArgs $;

        public Builder() {
            $ = new GetSecurityGroupArgs();
        }

        public Builder(GetSecurityGroupArgs defaults) {
            $ = new GetSecurityGroupArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the security group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the security group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetSecurityGroupArgs build() {
            return $;
        }
    }

}

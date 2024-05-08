// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.networkaccess.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthorizationGlobalExceptionRuleArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthorizationGlobalExceptionRuleArgs Empty = new GetAuthorizationGlobalExceptionRuleArgs();

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
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private GetAuthorizationGlobalExceptionRuleArgs() {}

    private GetAuthorizationGlobalExceptionRuleArgs(GetAuthorizationGlobalExceptionRuleArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthorizationGlobalExceptionRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthorizationGlobalExceptionRuleArgs $;

        public Builder() {
            $ = new GetAuthorizationGlobalExceptionRuleArgs();
        }

        public Builder(GetAuthorizationGlobalExceptionRuleArgs defaults) {
            $ = new GetAuthorizationGlobalExceptionRuleArgs(Objects.requireNonNull(defaults));
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
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetAuthorizationGlobalExceptionRuleArgs build() {
            return $;
        }
    }

}

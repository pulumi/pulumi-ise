// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.deviceadmin.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthorizationGlobalExceptionRulePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthorizationGlobalExceptionRulePlainArgs Empty = new GetAuthorizationGlobalExceptionRulePlainArgs();

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
     * Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    private GetAuthorizationGlobalExceptionRulePlainArgs() {}

    private GetAuthorizationGlobalExceptionRulePlainArgs(GetAuthorizationGlobalExceptionRulePlainArgs $) {
        this.id = $.id;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthorizationGlobalExceptionRulePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthorizationGlobalExceptionRulePlainArgs $;

        public Builder() {
            $ = new GetAuthorizationGlobalExceptionRulePlainArgs();
        }

        public Builder(GetAuthorizationGlobalExceptionRulePlainArgs defaults) {
            $ = new GetAuthorizationGlobalExceptionRulePlainArgs(Objects.requireNonNull(defaults));
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
         * @param name Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        public GetAuthorizationGlobalExceptionRulePlainArgs build() {
            return $;
        }
    }

}

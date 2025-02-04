// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointIdentityGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointIdentityGroupArgs Empty = new EndpointIdentityGroupArgs();

    /**
     * Description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The name of the endpoint identity group
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the endpoint identity group
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Parent endpoint identity group ID
     * 
     */
    @Import(name="parentEndpointIdentityGroupId")
    private @Nullable Output<String> parentEndpointIdentityGroupId;

    /**
     * @return Parent endpoint identity group ID
     * 
     */
    public Optional<Output<String>> parentEndpointIdentityGroupId() {
        return Optional.ofNullable(this.parentEndpointIdentityGroupId);
    }

    /**
     * System defined endpoint identity group
     * 
     */
    @Import(name="systemDefined")
    private @Nullable Output<Boolean> systemDefined;

    /**
     * @return System defined endpoint identity group
     * 
     */
    public Optional<Output<Boolean>> systemDefined() {
        return Optional.ofNullable(this.systemDefined);
    }

    private EndpointIdentityGroupArgs() {}

    private EndpointIdentityGroupArgs(EndpointIdentityGroupArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.parentEndpointIdentityGroupId = $.parentEndpointIdentityGroupId;
        this.systemDefined = $.systemDefined;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointIdentityGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointIdentityGroupArgs $;

        public Builder() {
            $ = new EndpointIdentityGroupArgs();
        }

        public Builder(EndpointIdentityGroupArgs defaults) {
            $ = new EndpointIdentityGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name The name of the endpoint identity group
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the endpoint identity group
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parentEndpointIdentityGroupId Parent endpoint identity group ID
         * 
         * @return builder
         * 
         */
        public Builder parentEndpointIdentityGroupId(@Nullable Output<String> parentEndpointIdentityGroupId) {
            $.parentEndpointIdentityGroupId = parentEndpointIdentityGroupId;
            return this;
        }

        /**
         * @param parentEndpointIdentityGroupId Parent endpoint identity group ID
         * 
         * @return builder
         * 
         */
        public Builder parentEndpointIdentityGroupId(String parentEndpointIdentityGroupId) {
            return parentEndpointIdentityGroupId(Output.of(parentEndpointIdentityGroupId));
        }

        /**
         * @param systemDefined System defined endpoint identity group
         * 
         * @return builder
         * 
         */
        public Builder systemDefined(@Nullable Output<Boolean> systemDefined) {
            $.systemDefined = systemDefined;
            return this;
        }

        /**
         * @param systemDefined System defined endpoint identity group
         * 
         * @return builder
         * 
         */
        public Builder systemDefined(Boolean systemDefined) {
            return systemDefined(Output.of(systemDefined));
        }

        public EndpointIdentityGroupArgs build() {
            return $;
        }
    }

}

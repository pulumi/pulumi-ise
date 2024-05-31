// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.network;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DeviceGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final DeviceGroupArgs Empty = new DeviceGroupArgs();

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
     * The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the root device group.
     * 
     */
    @Import(name="rootGroup", required=true)
    private Output<String> rootGroup;

    /**
     * @return The name of the root device group.
     * 
     */
    public Output<String> rootGroup() {
        return this.rootGroup;
    }

    private DeviceGroupArgs() {}

    private DeviceGroupArgs(DeviceGroupArgs $) {
        this.description = $.description;
        this.name = $.name;
        this.rootGroup = $.rootGroup;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DeviceGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DeviceGroupArgs $;

        public Builder() {
            $ = new DeviceGroupArgs();
        }

        public Builder(DeviceGroupArgs defaults) {
            $ = new DeviceGroupArgs(Objects.requireNonNull(defaults));
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
         * @param name The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param rootGroup The name of the root device group.
         * 
         * @return builder
         * 
         */
        public Builder rootGroup(Output<String> rootGroup) {
            $.rootGroup = rootGroup;
            return this;
        }

        /**
         * @param rootGroup The name of the root device group.
         * 
         * @return builder
         * 
         */
        public Builder rootGroup(String rootGroup) {
            return rootGroup(Output.of(rootGroup));
        }

        public DeviceGroupArgs build() {
            if ($.rootGroup == null) {
                throw new MissingRequiredPropertyException("DeviceGroupArgs", "rootGroup");
            }
            return $;
        }
    }

}
// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityGroupArgs Empty = new SecurityGroupArgs();

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
     * Read-only - Default value: `false`
     * 
     */
    @Import(name="isReadOnly")
    private @Nullable Output<Boolean> isReadOnly;

    /**
     * @return Read-only - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> isReadOnly() {
        return Optional.ofNullable(this.isReadOnly);
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

    /**
     * Propagate to APIC (ACI)
     * 
     */
    @Import(name="propogateToApic")
    private @Nullable Output<Boolean> propogateToApic;

    /**
     * @return Propagate to APIC (ACI)
     * 
     */
    public Optional<Output<Boolean>> propogateToApic() {
        return Optional.ofNullable(this.propogateToApic);
    }

    /**
     * `-1` to auto-generate - Range: `-1`-`65519`
     * 
     */
    @Import(name="value", required=true)
    private Output<Integer> value;

    /**
     * @return `-1` to auto-generate - Range: `-1`-`65519`
     * 
     */
    public Output<Integer> value() {
        return this.value;
    }

    private SecurityGroupArgs() {}

    private SecurityGroupArgs(SecurityGroupArgs $) {
        this.description = $.description;
        this.isReadOnly = $.isReadOnly;
        this.name = $.name;
        this.propogateToApic = $.propogateToApic;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityGroupArgs $;

        public Builder() {
            $ = new SecurityGroupArgs();
        }

        public Builder(SecurityGroupArgs defaults) {
            $ = new SecurityGroupArgs(Objects.requireNonNull(defaults));
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
         * @param isReadOnly Read-only - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder isReadOnly(@Nullable Output<Boolean> isReadOnly) {
            $.isReadOnly = isReadOnly;
            return this;
        }

        /**
         * @param isReadOnly Read-only - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder isReadOnly(Boolean isReadOnly) {
            return isReadOnly(Output.of(isReadOnly));
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

        /**
         * @param propogateToApic Propagate to APIC (ACI)
         * 
         * @return builder
         * 
         */
        public Builder propogateToApic(@Nullable Output<Boolean> propogateToApic) {
            $.propogateToApic = propogateToApic;
            return this;
        }

        /**
         * @param propogateToApic Propagate to APIC (ACI)
         * 
         * @return builder
         * 
         */
        public Builder propogateToApic(Boolean propogateToApic) {
            return propogateToApic(Output.of(propogateToApic));
        }

        /**
         * @param value `-1` to auto-generate - Range: `-1`-`65519`
         * 
         * @return builder
         * 
         */
        public Builder value(Output<Integer> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value `-1` to auto-generate - Range: `-1`-`65519`
         * 
         * @return builder
         * 
         */
        public Builder value(Integer value) {
            return value(Output.of(value));
        }

        public SecurityGroupArgs build() {
            if ($.value == null) {
                throw new MissingRequiredPropertyException("SecurityGroupArgs", "value");
            }
            return $;
        }
    }

}
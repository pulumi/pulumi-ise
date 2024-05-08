// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.system.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class LicenseTierStateLicenseArgs extends com.pulumi.resources.ResourceArgs {

    public static final LicenseTierStateLicenseArgs Empty = new LicenseTierStateLicenseArgs();

    /**
     * License name
     *   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return License name
     *   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * License status
     *   - Choices: `ENABLED`, `DISABLED`
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return License status
     *   - Choices: `ENABLED`, `DISABLED`
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    private LicenseTierStateLicenseArgs() {}

    private LicenseTierStateLicenseArgs(LicenseTierStateLicenseArgs $) {
        this.name = $.name;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LicenseTierStateLicenseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LicenseTierStateLicenseArgs $;

        public Builder() {
            $ = new LicenseTierStateLicenseArgs();
        }

        public Builder(LicenseTierStateLicenseArgs defaults) {
            $ = new LicenseTierStateLicenseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name License name
         *   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name License name
         *   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param status License status
         *   - Choices: `ENABLED`, `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status License status
         *   - Choices: `ENABLED`, `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public LicenseTierStateLicenseArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("LicenseTierStateLicenseArgs", "name");
            }
            if ($.status == null) {
                throw new MissingRequiredPropertyException("LicenseTierStateLicenseArgs", "status");
            }
            return $;
        }
    }

}

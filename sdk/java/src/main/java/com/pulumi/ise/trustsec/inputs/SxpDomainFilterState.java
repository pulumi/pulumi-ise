// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SxpDomainFilterState extends com.pulumi.resources.ResourceArgs {

    public static final SxpDomainFilterState Empty = new SxpDomainFilterState();

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
     * List of SXP Domains, separated with comma
     * 
     */
    @Import(name="domains")
    private @Nullable Output<String> domains;

    /**
     * @return List of SXP Domains, separated with comma
     * 
     */
    public Optional<Output<String>> domains() {
        return Optional.ofNullable(this.domains);
    }

    /**
     * Resource name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Resource name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * SGT name or ID. At least one of subnet or sgt or vn should be defined
     * 
     */
    @Import(name="sgt")
    private @Nullable Output<String> sgt;

    /**
     * @return SGT name or ID. At least one of subnet or sgt or vn should be defined
     * 
     */
    public Optional<Output<String>> sgt() {
        return Optional.ofNullable(this.sgt);
    }

    /**
     * Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
     * 
     */
    @Import(name="subnet")
    private @Nullable Output<String> subnet;

    /**
     * @return Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
     * 
     */
    public Optional<Output<String>> subnet() {
        return Optional.ofNullable(this.subnet);
    }

    /**
     * Virtual Network. At least one of subnet or sgt or vn should be defined
     * 
     */
    @Import(name="vn")
    private @Nullable Output<String> vn;

    /**
     * @return Virtual Network. At least one of subnet or sgt or vn should be defined
     * 
     */
    public Optional<Output<String>> vn() {
        return Optional.ofNullable(this.vn);
    }

    private SxpDomainFilterState() {}

    private SxpDomainFilterState(SxpDomainFilterState $) {
        this.description = $.description;
        this.domains = $.domains;
        this.name = $.name;
        this.sgt = $.sgt;
        this.subnet = $.subnet;
        this.vn = $.vn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SxpDomainFilterState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SxpDomainFilterState $;

        public Builder() {
            $ = new SxpDomainFilterState();
        }

        public Builder(SxpDomainFilterState defaults) {
            $ = new SxpDomainFilterState(Objects.requireNonNull(defaults));
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
         * @param domains List of SXP Domains, separated with comma
         * 
         * @return builder
         * 
         */
        public Builder domains(@Nullable Output<String> domains) {
            $.domains = domains;
            return this;
        }

        /**
         * @param domains List of SXP Domains, separated with comma
         * 
         * @return builder
         * 
         */
        public Builder domains(String domains) {
            return domains(Output.of(domains));
        }

        /**
         * @param name Resource name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Resource name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param sgt SGT name or ID. At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder sgt(@Nullable Output<String> sgt) {
            $.sgt = sgt;
            return this;
        }

        /**
         * @param sgt SGT name or ID. At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder sgt(String sgt) {
            return sgt(Output.of(sgt));
        }

        /**
         * @param subnet Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder subnet(@Nullable Output<String> subnet) {
            $.subnet = subnet;
            return this;
        }

        /**
         * @param subnet Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder subnet(String subnet) {
            return subnet(Output.of(subnet));
        }

        /**
         * @param vn Virtual Network. At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder vn(@Nullable Output<String> vn) {
            $.vn = vn;
            return this;
        }

        /**
         * @param vn Virtual Network. At least one of subnet or sgt or vn should be defined
         * 
         * @return builder
         * 
         */
        public Builder vn(String vn) {
            return vn(Output.of(vn));
        }

        public SxpDomainFilterState build() {
            return $;
        }
    }

}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.network.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DeviceIp {
    /**
     * @return It can be either single ip address or ip range address
     * 
     */
    private String ipaddress;
    /**
     * @return It can be either single ip address or ip range address
     * 
     */
    private @Nullable String ipaddressExclude;
    /**
     * @return Subnet mask length
     * 
     */
    private @Nullable String mask;

    private DeviceIp() {}
    /**
     * @return It can be either single ip address or ip range address
     * 
     */
    public String ipaddress() {
        return this.ipaddress;
    }
    /**
     * @return It can be either single ip address or ip range address
     * 
     */
    public Optional<String> ipaddressExclude() {
        return Optional.ofNullable(this.ipaddressExclude);
    }
    /**
     * @return Subnet mask length
     * 
     */
    public Optional<String> mask() {
        return Optional.ofNullable(this.mask);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DeviceIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ipaddress;
        private @Nullable String ipaddressExclude;
        private @Nullable String mask;
        public Builder() {}
        public Builder(DeviceIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipaddress = defaults.ipaddress;
    	      this.ipaddressExclude = defaults.ipaddressExclude;
    	      this.mask = defaults.mask;
        }

        @CustomType.Setter
        public Builder ipaddress(String ipaddress) {
            if (ipaddress == null) {
              throw new MissingRequiredPropertyException("DeviceIp", "ipaddress");
            }
            this.ipaddress = ipaddress;
            return this;
        }
        @CustomType.Setter
        public Builder ipaddressExclude(@Nullable String ipaddressExclude) {

            this.ipaddressExclude = ipaddressExclude;
            return this;
        }
        @CustomType.Setter
        public Builder mask(@Nullable String mask) {

            this.mask = mask;
            return this;
        }
        public DeviceIp build() {
            final var _resultValue = new DeviceIp();
            _resultValue.ipaddress = ipaddress;
            _resultValue.ipaddressExclude = ipaddressExclude;
            _resultValue.mask = mask;
            return _resultValue;
        }
    }
}

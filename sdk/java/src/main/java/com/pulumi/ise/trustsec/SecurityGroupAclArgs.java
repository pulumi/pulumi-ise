// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.trustsec;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecurityGroupAclArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecurityGroupAclArgs Empty = new SecurityGroupAclArgs();

    /**
     * Content of ACL
     * 
     */
    @Import(name="aclContent", required=true)
    private Output<String> aclContent;

    /**
     * @return Content of ACL
     * 
     */
    public Output<String> aclContent() {
        return this.aclContent;
    }

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
     * IP Version - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IP_AGNOSTIC`
     * 
     */
    @Import(name="ipVersion")
    private @Nullable Output<String> ipVersion;

    /**
     * @return IP Version - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IP_AGNOSTIC`
     * 
     */
    public Optional<Output<String>> ipVersion() {
        return Optional.ofNullable(this.ipVersion);
    }

    /**
     * The name of the security group ACL
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the security group ACL
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Read-only
     * 
     */
    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    /**
     * @return Read-only
     * 
     */
    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    private SecurityGroupAclArgs() {}

    private SecurityGroupAclArgs(SecurityGroupAclArgs $) {
        this.aclContent = $.aclContent;
        this.description = $.description;
        this.ipVersion = $.ipVersion;
        this.name = $.name;
        this.readOnly = $.readOnly;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecurityGroupAclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecurityGroupAclArgs $;

        public Builder() {
            $ = new SecurityGroupAclArgs();
        }

        public Builder(SecurityGroupAclArgs defaults) {
            $ = new SecurityGroupAclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param aclContent Content of ACL
         * 
         * @return builder
         * 
         */
        public Builder aclContent(Output<String> aclContent) {
            $.aclContent = aclContent;
            return this;
        }

        /**
         * @param aclContent Content of ACL
         * 
         * @return builder
         * 
         */
        public Builder aclContent(String aclContent) {
            return aclContent(Output.of(aclContent));
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
         * @param ipVersion IP Version - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IP_AGNOSTIC`
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(@Nullable Output<String> ipVersion) {
            $.ipVersion = ipVersion;
            return this;
        }

        /**
         * @param ipVersion IP Version - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IP_AGNOSTIC`
         * 
         * @return builder
         * 
         */
        public Builder ipVersion(String ipVersion) {
            return ipVersion(Output.of(ipVersion));
        }

        /**
         * @param name The name of the security group ACL
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the security group ACL
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param readOnly Read-only
         * 
         * @return builder
         * 
         */
        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        /**
         * @param readOnly Read-only
         * 
         * @return builder
         * 
         */
        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        public SecurityGroupAclArgs build() {
            if ($.aclContent == null) {
                throw new MissingRequiredPropertyException("SecurityGroupAclArgs", "aclContent");
            }
            return $;
        }
    }

}

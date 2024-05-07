// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.Internal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserState extends com.pulumi.resources.ResourceArgs {

    public static final UserState Empty = new UserState();

    /**
     * The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
     * from ISE 3.2.
     * 
     */
    @Import(name="accountNameAlias")
    private @Nullable Output<String> accountNameAlias;

    /**
     * @return The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
     * from ISE 3.2.
     * 
     */
    public Optional<Output<String>> accountNameAlias() {
        return Optional.ofNullable(this.accountNameAlias);
    }

    /**
     * Requires the user to change the password - Default value: `true`
     * 
     */
    @Import(name="changePassword")
    private @Nullable Output<Boolean> changePassword;

    /**
     * @return Requires the user to change the password - Default value: `true`
     * 
     */
    public Optional<Output<Boolean>> changePassword() {
        return Optional.ofNullable(this.changePassword);
    }

    /**
     * Key value map
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<String> customAttributes;

    /**
     * @return Key value map
     * 
     */
    public Optional<Output<String>> customAttributes() {
        return Optional.ofNullable(this.customAttributes);
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
     * Email address
     * 
     */
    @Import(name="email")
    private @Nullable Output<String> email;

    /**
     * @return Email address
     * 
     */
    public Optional<Output<String>> email() {
        return Optional.ofNullable(this.email);
    }

    /**
     * This field is added in ISE 2.0 to support TACACS+
     * 
     */
    @Import(name="enablePassword")
    private @Nullable Output<String> enablePassword;

    /**
     * @return This field is added in ISE 2.0 to support TACACS+
     * 
     */
    public Optional<Output<String>> enablePassword() {
        return Optional.ofNullable(this.enablePassword);
    }

    /**
     * Whether the user is enabled/disabled
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether the user is enabled/disabled
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * First name of the internal user
     * 
     */
    @Import(name="firstName")
    private @Nullable Output<String> firstName;

    /**
     * @return First name of the internal user
     * 
     */
    public Optional<Output<String>> firstName() {
        return Optional.ofNullable(this.firstName);
    }

    /**
     * Comma separated list of identity group IDs.
     * 
     */
    @Import(name="identityGroups")
    private @Nullable Output<String> identityGroups;

    /**
     * @return Comma separated list of identity group IDs.
     * 
     */
    public Optional<Output<String>> identityGroups() {
        return Optional.ofNullable(this.identityGroups);
    }

    /**
     * Last name of the internal user
     * 
     */
    @Import(name="lastName")
    private @Nullable Output<String> lastName;

    /**
     * @return Last name of the internal user
     * 
     */
    public Optional<Output<String>> lastName() {
        return Optional.ofNullable(this.lastName);
    }

    /**
     * The name of the internal user
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the internal user
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The password of the internal user
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password of the internal user
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
     * 
     */
    @Import(name="passwordIdStore")
    private @Nullable Output<String> passwordIdStore;

    /**
     * @return The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
     * 
     */
    public Optional<Output<String>> passwordIdStore() {
        return Optional.ofNullable(this.passwordIdStore);
    }

    /**
     * Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
     * field is only supported from ISE 3.2. - Default value: `false`
     * 
     */
    @Import(name="passwordNeverExpires")
    private @Nullable Output<Boolean> passwordNeverExpires;

    /**
     * @return Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
     * field is only supported from ISE 3.2. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> passwordNeverExpires() {
        return Optional.ofNullable(this.passwordNeverExpires);
    }

    private UserState() {}

    private UserState(UserState $) {
        this.accountNameAlias = $.accountNameAlias;
        this.changePassword = $.changePassword;
        this.customAttributes = $.customAttributes;
        this.description = $.description;
        this.email = $.email;
        this.enablePassword = $.enablePassword;
        this.enabled = $.enabled;
        this.firstName = $.firstName;
        this.identityGroups = $.identityGroups;
        this.lastName = $.lastName;
        this.name = $.name;
        this.password = $.password;
        this.passwordIdStore = $.passwordIdStore;
        this.passwordNeverExpires = $.passwordNeverExpires;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserState $;

        public Builder() {
            $ = new UserState();
        }

        public Builder(UserState defaults) {
            $ = new UserState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountNameAlias The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
         * from ISE 3.2.
         * 
         * @return builder
         * 
         */
        public Builder accountNameAlias(@Nullable Output<String> accountNameAlias) {
            $.accountNameAlias = accountNameAlias;
            return this;
        }

        /**
         * @param accountNameAlias The Account Name Alias will be used to send email notifications about password expiration. This field is only supported
         * from ISE 3.2.
         * 
         * @return builder
         * 
         */
        public Builder accountNameAlias(String accountNameAlias) {
            return accountNameAlias(Output.of(accountNameAlias));
        }

        /**
         * @param changePassword Requires the user to change the password - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder changePassword(@Nullable Output<Boolean> changePassword) {
            $.changePassword = changePassword;
            return this;
        }

        /**
         * @param changePassword Requires the user to change the password - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder changePassword(Boolean changePassword) {
            return changePassword(Output.of(changePassword));
        }

        /**
         * @param customAttributes Key value map
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<String> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes Key value map
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(String customAttributes) {
            return customAttributes(Output.of(customAttributes));
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
         * @param email Email address
         * 
         * @return builder
         * 
         */
        public Builder email(@Nullable Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email Email address
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param enablePassword This field is added in ISE 2.0 to support TACACS+
         * 
         * @return builder
         * 
         */
        public Builder enablePassword(@Nullable Output<String> enablePassword) {
            $.enablePassword = enablePassword;
            return this;
        }

        /**
         * @param enablePassword This field is added in ISE 2.0 to support TACACS+
         * 
         * @return builder
         * 
         */
        public Builder enablePassword(String enablePassword) {
            return enablePassword(Output.of(enablePassword));
        }

        /**
         * @param enabled Whether the user is enabled/disabled
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether the user is enabled/disabled
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param firstName First name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder firstName(@Nullable Output<String> firstName) {
            $.firstName = firstName;
            return this;
        }

        /**
         * @param firstName First name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder firstName(String firstName) {
            return firstName(Output.of(firstName));
        }

        /**
         * @param identityGroups Comma separated list of identity group IDs.
         * 
         * @return builder
         * 
         */
        public Builder identityGroups(@Nullable Output<String> identityGroups) {
            $.identityGroups = identityGroups;
            return this;
        }

        /**
         * @param identityGroups Comma separated list of identity group IDs.
         * 
         * @return builder
         * 
         */
        public Builder identityGroups(String identityGroups) {
            return identityGroups(Output.of(identityGroups));
        }

        /**
         * @param lastName Last name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder lastName(@Nullable Output<String> lastName) {
            $.lastName = lastName;
            return this;
        }

        /**
         * @param lastName Last name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder lastName(String lastName) {
            return lastName(Output.of(lastName));
        }

        /**
         * @param name The name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the internal user
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The password of the internal user
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password of the internal user
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param passwordIdStore The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
         * 
         * @return builder
         * 
         */
        public Builder passwordIdStore(@Nullable Output<String> passwordIdStore) {
            $.passwordIdStore = passwordIdStore;
            return this;
        }

        /**
         * @param passwordIdStore The ID store where the internal user&#39;s password is kept - Default value: `Internal Users`
         * 
         * @return builder
         * 
         */
        public Builder passwordIdStore(String passwordIdStore) {
            return passwordIdStore(Output.of(passwordIdStore));
        }

        /**
         * @param passwordNeverExpires Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
         * field is only supported from ISE 3.2. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder passwordNeverExpires(@Nullable Output<Boolean> passwordNeverExpires) {
            $.passwordNeverExpires = passwordNeverExpires;
            return this;
        }

        /**
         * @param passwordNeverExpires Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This
         * field is only supported from ISE 3.2. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder passwordNeverExpires(Boolean passwordNeverExpires) {
            return passwordNeverExpires(Output.of(passwordNeverExpires));
        }

        public UserState build() {
            return $;
        }
    }

}

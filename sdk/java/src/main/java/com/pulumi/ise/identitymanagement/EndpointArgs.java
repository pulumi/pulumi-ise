// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final EndpointArgs Empty = new EndpointArgs();

    /**
     * Custom Attributes
     * 
     */
    @Import(name="customAttributes")
    private @Nullable Output<Map<String,String>> customAttributes;

    /**
     * @return Custom Attributes
     * 
     */
    public Optional<Output<Map<String,String>>> customAttributes() {
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
     * Identity Group ID
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return Identity Group ID
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * Identity Store
     * 
     */
    @Import(name="identityStore")
    private @Nullable Output<String> identityStore;

    /**
     * @return Identity Store
     * 
     */
    public Optional<Output<String>> identityStore() {
        return Optional.ofNullable(this.identityStore);
    }

    /**
     * Identity Store Id
     * 
     */
    @Import(name="identityStoreId")
    private @Nullable Output<String> identityStoreId;

    /**
     * @return Identity Store Id
     * 
     */
    public Optional<Output<String>> identityStoreId() {
        return Optional.ofNullable(this.identityStoreId);
    }

    /**
     * MAC address of the endpoint
     * 
     */
    @Import(name="mac", required=true)
    private Output<String> mac;

    /**
     * @return MAC address of the endpoint
     * 
     */
    public Output<String> mac() {
        return this.mac;
    }

    /**
     * Mdm Compliance Status
     * 
     */
    @Import(name="mdmComplianceStatus")
    private @Nullable Output<Boolean> mdmComplianceStatus;

    /**
     * @return Mdm Compliance Status
     * 
     */
    public Optional<Output<Boolean>> mdmComplianceStatus() {
        return Optional.ofNullable(this.mdmComplianceStatus);
    }

    /**
     * Mdm Encrypted
     * 
     */
    @Import(name="mdmEncrypted")
    private @Nullable Output<Boolean> mdmEncrypted;

    /**
     * @return Mdm Encrypted
     * 
     */
    public Optional<Output<Boolean>> mdmEncrypted() {
        return Optional.ofNullable(this.mdmEncrypted);
    }

    /**
     * Mdm Enrolled
     * 
     */
    @Import(name="mdmEnrolled")
    private @Nullable Output<Boolean> mdmEnrolled;

    /**
     * @return Mdm Enrolled
     * 
     */
    public Optional<Output<Boolean>> mdmEnrolled() {
        return Optional.ofNullable(this.mdmEnrolled);
    }

    /**
     * Mdm IMEI
     * 
     */
    @Import(name="mdmImei")
    private @Nullable Output<String> mdmImei;

    /**
     * @return Mdm IMEI
     * 
     */
    public Optional<Output<String>> mdmImei() {
        return Optional.ofNullable(this.mdmImei);
    }

    /**
     * Mdm JailBroken
     * 
     */
    @Import(name="mdmJailBroken")
    private @Nullable Output<Boolean> mdmJailBroken;

    /**
     * @return Mdm JailBroken
     * 
     */
    public Optional<Output<Boolean>> mdmJailBroken() {
        return Optional.ofNullable(this.mdmJailBroken);
    }

    /**
     * Mdm Manufacturer
     * 
     */
    @Import(name="mdmManufacturer")
    private @Nullable Output<String> mdmManufacturer;

    /**
     * @return Mdm Manufacturer
     * 
     */
    public Optional<Output<String>> mdmManufacturer() {
        return Optional.ofNullable(this.mdmManufacturer);
    }

    /**
     * Mdm Model
     * 
     */
    @Import(name="mdmModel")
    private @Nullable Output<String> mdmModel;

    /**
     * @return Mdm Model
     * 
     */
    public Optional<Output<String>> mdmModel() {
        return Optional.ofNullable(this.mdmModel);
    }

    /**
     * Mdm OS
     * 
     */
    @Import(name="mdmOs")
    private @Nullable Output<String> mdmOs;

    /**
     * @return Mdm OS
     * 
     */
    public Optional<Output<String>> mdmOs() {
        return Optional.ofNullable(this.mdmOs);
    }

    /**
     * Mdm PhoneNumber
     * 
     */
    @Import(name="mdmPhoneNumber")
    private @Nullable Output<String> mdmPhoneNumber;

    /**
     * @return Mdm PhoneNumber
     * 
     */
    public Optional<Output<String>> mdmPhoneNumber() {
        return Optional.ofNullable(this.mdmPhoneNumber);
    }

    /**
     * Mdm Pinlock
     * 
     */
    @Import(name="mdmPinlock")
    private @Nullable Output<Boolean> mdmPinlock;

    /**
     * @return Mdm Pinlock
     * 
     */
    public Optional<Output<Boolean>> mdmPinlock() {
        return Optional.ofNullable(this.mdmPinlock);
    }

    /**
     * Mdm Reachable
     * 
     */
    @Import(name="mdmReachable")
    private @Nullable Output<Boolean> mdmReachable;

    /**
     * @return Mdm Reachable
     * 
     */
    public Optional<Output<Boolean>> mdmReachable() {
        return Optional.ofNullable(this.mdmReachable);
    }

    /**
     * Mdm Serial
     * 
     */
    @Import(name="mdmSerial")
    private @Nullable Output<String> mdmSerial;

    /**
     * @return Mdm Serial
     * 
     */
    public Optional<Output<String>> mdmSerial() {
        return Optional.ofNullable(this.mdmSerial);
    }

    /**
     * Mdm Server Name
     * 
     */
    @Import(name="mdmServerName")
    private @Nullable Output<String> mdmServerName;

    /**
     * @return Mdm Server Name
     * 
     */
    public Optional<Output<String>> mdmServerName() {
        return Optional.ofNullable(this.mdmServerName);
    }

    /**
     * The name of the endpoint
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the endpoint
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Portal User
     * 
     */
    @Import(name="portalUser")
    private @Nullable Output<String> portalUser;

    /**
     * @return Portal User
     * 
     */
    public Optional<Output<String>> portalUser() {
        return Optional.ofNullable(this.portalUser);
    }

    /**
     * Profile ID
     * 
     */
    @Import(name="profileId")
    private @Nullable Output<String> profileId;

    /**
     * @return Profile ID
     * 
     */
    public Optional<Output<String>> profileId() {
        return Optional.ofNullable(this.profileId);
    }

    /**
     * Static Group Assignment
     * 
     */
    @Import(name="staticGroupAssignment", required=true)
    private Output<Boolean> staticGroupAssignment;

    /**
     * @return Static Group Assignment
     * 
     */
    public Output<Boolean> staticGroupAssignment() {
        return this.staticGroupAssignment;
    }

    /**
     * staticGroupAssignmentDefined - Default value: `true`
     * 
     */
    @Import(name="staticGroupAssignmentDefined")
    private @Nullable Output<Boolean> staticGroupAssignmentDefined;

    /**
     * @return staticGroupAssignmentDefined - Default value: `true`
     * 
     */
    public Optional<Output<Boolean>> staticGroupAssignmentDefined() {
        return Optional.ofNullable(this.staticGroupAssignmentDefined);
    }

    /**
     * Static Profile Assignment
     * 
     */
    @Import(name="staticProfileAssignment", required=true)
    private Output<Boolean> staticProfileAssignment;

    /**
     * @return Static Profile Assignment
     * 
     */
    public Output<Boolean> staticProfileAssignment() {
        return this.staticProfileAssignment;
    }

    /**
     * Static Profile Assignment Defined - Default value: `true`
     * 
     */
    @Import(name="staticProfileAssignmentDefined")
    private @Nullable Output<Boolean> staticProfileAssignmentDefined;

    /**
     * @return Static Profile Assignment Defined - Default value: `true`
     * 
     */
    public Optional<Output<Boolean>> staticProfileAssignmentDefined() {
        return Optional.ofNullable(this.staticProfileAssignmentDefined);
    }

    private EndpointArgs() {}

    private EndpointArgs(EndpointArgs $) {
        this.customAttributes = $.customAttributes;
        this.description = $.description;
        this.groupId = $.groupId;
        this.identityStore = $.identityStore;
        this.identityStoreId = $.identityStoreId;
        this.mac = $.mac;
        this.mdmComplianceStatus = $.mdmComplianceStatus;
        this.mdmEncrypted = $.mdmEncrypted;
        this.mdmEnrolled = $.mdmEnrolled;
        this.mdmImei = $.mdmImei;
        this.mdmJailBroken = $.mdmJailBroken;
        this.mdmManufacturer = $.mdmManufacturer;
        this.mdmModel = $.mdmModel;
        this.mdmOs = $.mdmOs;
        this.mdmPhoneNumber = $.mdmPhoneNumber;
        this.mdmPinlock = $.mdmPinlock;
        this.mdmReachable = $.mdmReachable;
        this.mdmSerial = $.mdmSerial;
        this.mdmServerName = $.mdmServerName;
        this.name = $.name;
        this.portalUser = $.portalUser;
        this.profileId = $.profileId;
        this.staticGroupAssignment = $.staticGroupAssignment;
        this.staticGroupAssignmentDefined = $.staticGroupAssignmentDefined;
        this.staticProfileAssignment = $.staticProfileAssignment;
        this.staticProfileAssignmentDefined = $.staticProfileAssignmentDefined;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EndpointArgs $;

        public Builder() {
            $ = new EndpointArgs();
        }

        public Builder(EndpointArgs defaults) {
            $ = new EndpointArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customAttributes Custom Attributes
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(@Nullable Output<Map<String,String>> customAttributes) {
            $.customAttributes = customAttributes;
            return this;
        }

        /**
         * @param customAttributes Custom Attributes
         * 
         * @return builder
         * 
         */
        public Builder customAttributes(Map<String,String> customAttributes) {
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
         * @param groupId Identity Group ID
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId Identity Group ID
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param identityStore Identity Store
         * 
         * @return builder
         * 
         */
        public Builder identityStore(@Nullable Output<String> identityStore) {
            $.identityStore = identityStore;
            return this;
        }

        /**
         * @param identityStore Identity Store
         * 
         * @return builder
         * 
         */
        public Builder identityStore(String identityStore) {
            return identityStore(Output.of(identityStore));
        }

        /**
         * @param identityStoreId Identity Store Id
         * 
         * @return builder
         * 
         */
        public Builder identityStoreId(@Nullable Output<String> identityStoreId) {
            $.identityStoreId = identityStoreId;
            return this;
        }

        /**
         * @param identityStoreId Identity Store Id
         * 
         * @return builder
         * 
         */
        public Builder identityStoreId(String identityStoreId) {
            return identityStoreId(Output.of(identityStoreId));
        }

        /**
         * @param mac MAC address of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder mac(Output<String> mac) {
            $.mac = mac;
            return this;
        }

        /**
         * @param mac MAC address of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder mac(String mac) {
            return mac(Output.of(mac));
        }

        /**
         * @param mdmComplianceStatus Mdm Compliance Status
         * 
         * @return builder
         * 
         */
        public Builder mdmComplianceStatus(@Nullable Output<Boolean> mdmComplianceStatus) {
            $.mdmComplianceStatus = mdmComplianceStatus;
            return this;
        }

        /**
         * @param mdmComplianceStatus Mdm Compliance Status
         * 
         * @return builder
         * 
         */
        public Builder mdmComplianceStatus(Boolean mdmComplianceStatus) {
            return mdmComplianceStatus(Output.of(mdmComplianceStatus));
        }

        /**
         * @param mdmEncrypted Mdm Encrypted
         * 
         * @return builder
         * 
         */
        public Builder mdmEncrypted(@Nullable Output<Boolean> mdmEncrypted) {
            $.mdmEncrypted = mdmEncrypted;
            return this;
        }

        /**
         * @param mdmEncrypted Mdm Encrypted
         * 
         * @return builder
         * 
         */
        public Builder mdmEncrypted(Boolean mdmEncrypted) {
            return mdmEncrypted(Output.of(mdmEncrypted));
        }

        /**
         * @param mdmEnrolled Mdm Enrolled
         * 
         * @return builder
         * 
         */
        public Builder mdmEnrolled(@Nullable Output<Boolean> mdmEnrolled) {
            $.mdmEnrolled = mdmEnrolled;
            return this;
        }

        /**
         * @param mdmEnrolled Mdm Enrolled
         * 
         * @return builder
         * 
         */
        public Builder mdmEnrolled(Boolean mdmEnrolled) {
            return mdmEnrolled(Output.of(mdmEnrolled));
        }

        /**
         * @param mdmImei Mdm IMEI
         * 
         * @return builder
         * 
         */
        public Builder mdmImei(@Nullable Output<String> mdmImei) {
            $.mdmImei = mdmImei;
            return this;
        }

        /**
         * @param mdmImei Mdm IMEI
         * 
         * @return builder
         * 
         */
        public Builder mdmImei(String mdmImei) {
            return mdmImei(Output.of(mdmImei));
        }

        /**
         * @param mdmJailBroken Mdm JailBroken
         * 
         * @return builder
         * 
         */
        public Builder mdmJailBroken(@Nullable Output<Boolean> mdmJailBroken) {
            $.mdmJailBroken = mdmJailBroken;
            return this;
        }

        /**
         * @param mdmJailBroken Mdm JailBroken
         * 
         * @return builder
         * 
         */
        public Builder mdmJailBroken(Boolean mdmJailBroken) {
            return mdmJailBroken(Output.of(mdmJailBroken));
        }

        /**
         * @param mdmManufacturer Mdm Manufacturer
         * 
         * @return builder
         * 
         */
        public Builder mdmManufacturer(@Nullable Output<String> mdmManufacturer) {
            $.mdmManufacturer = mdmManufacturer;
            return this;
        }

        /**
         * @param mdmManufacturer Mdm Manufacturer
         * 
         * @return builder
         * 
         */
        public Builder mdmManufacturer(String mdmManufacturer) {
            return mdmManufacturer(Output.of(mdmManufacturer));
        }

        /**
         * @param mdmModel Mdm Model
         * 
         * @return builder
         * 
         */
        public Builder mdmModel(@Nullable Output<String> mdmModel) {
            $.mdmModel = mdmModel;
            return this;
        }

        /**
         * @param mdmModel Mdm Model
         * 
         * @return builder
         * 
         */
        public Builder mdmModel(String mdmModel) {
            return mdmModel(Output.of(mdmModel));
        }

        /**
         * @param mdmOs Mdm OS
         * 
         * @return builder
         * 
         */
        public Builder mdmOs(@Nullable Output<String> mdmOs) {
            $.mdmOs = mdmOs;
            return this;
        }

        /**
         * @param mdmOs Mdm OS
         * 
         * @return builder
         * 
         */
        public Builder mdmOs(String mdmOs) {
            return mdmOs(Output.of(mdmOs));
        }

        /**
         * @param mdmPhoneNumber Mdm PhoneNumber
         * 
         * @return builder
         * 
         */
        public Builder mdmPhoneNumber(@Nullable Output<String> mdmPhoneNumber) {
            $.mdmPhoneNumber = mdmPhoneNumber;
            return this;
        }

        /**
         * @param mdmPhoneNumber Mdm PhoneNumber
         * 
         * @return builder
         * 
         */
        public Builder mdmPhoneNumber(String mdmPhoneNumber) {
            return mdmPhoneNumber(Output.of(mdmPhoneNumber));
        }

        /**
         * @param mdmPinlock Mdm Pinlock
         * 
         * @return builder
         * 
         */
        public Builder mdmPinlock(@Nullable Output<Boolean> mdmPinlock) {
            $.mdmPinlock = mdmPinlock;
            return this;
        }

        /**
         * @param mdmPinlock Mdm Pinlock
         * 
         * @return builder
         * 
         */
        public Builder mdmPinlock(Boolean mdmPinlock) {
            return mdmPinlock(Output.of(mdmPinlock));
        }

        /**
         * @param mdmReachable Mdm Reachable
         * 
         * @return builder
         * 
         */
        public Builder mdmReachable(@Nullable Output<Boolean> mdmReachable) {
            $.mdmReachable = mdmReachable;
            return this;
        }

        /**
         * @param mdmReachable Mdm Reachable
         * 
         * @return builder
         * 
         */
        public Builder mdmReachable(Boolean mdmReachable) {
            return mdmReachable(Output.of(mdmReachable));
        }

        /**
         * @param mdmSerial Mdm Serial
         * 
         * @return builder
         * 
         */
        public Builder mdmSerial(@Nullable Output<String> mdmSerial) {
            $.mdmSerial = mdmSerial;
            return this;
        }

        /**
         * @param mdmSerial Mdm Serial
         * 
         * @return builder
         * 
         */
        public Builder mdmSerial(String mdmSerial) {
            return mdmSerial(Output.of(mdmSerial));
        }

        /**
         * @param mdmServerName Mdm Server Name
         * 
         * @return builder
         * 
         */
        public Builder mdmServerName(@Nullable Output<String> mdmServerName) {
            $.mdmServerName = mdmServerName;
            return this;
        }

        /**
         * @param mdmServerName Mdm Server Name
         * 
         * @return builder
         * 
         */
        public Builder mdmServerName(String mdmServerName) {
            return mdmServerName(Output.of(mdmServerName));
        }

        /**
         * @param name The name of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the endpoint
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param portalUser Portal User
         * 
         * @return builder
         * 
         */
        public Builder portalUser(@Nullable Output<String> portalUser) {
            $.portalUser = portalUser;
            return this;
        }

        /**
         * @param portalUser Portal User
         * 
         * @return builder
         * 
         */
        public Builder portalUser(String portalUser) {
            return portalUser(Output.of(portalUser));
        }

        /**
         * @param profileId Profile ID
         * 
         * @return builder
         * 
         */
        public Builder profileId(@Nullable Output<String> profileId) {
            $.profileId = profileId;
            return this;
        }

        /**
         * @param profileId Profile ID
         * 
         * @return builder
         * 
         */
        public Builder profileId(String profileId) {
            return profileId(Output.of(profileId));
        }

        /**
         * @param staticGroupAssignment Static Group Assignment
         * 
         * @return builder
         * 
         */
        public Builder staticGroupAssignment(Output<Boolean> staticGroupAssignment) {
            $.staticGroupAssignment = staticGroupAssignment;
            return this;
        }

        /**
         * @param staticGroupAssignment Static Group Assignment
         * 
         * @return builder
         * 
         */
        public Builder staticGroupAssignment(Boolean staticGroupAssignment) {
            return staticGroupAssignment(Output.of(staticGroupAssignment));
        }

        /**
         * @param staticGroupAssignmentDefined staticGroupAssignmentDefined - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder staticGroupAssignmentDefined(@Nullable Output<Boolean> staticGroupAssignmentDefined) {
            $.staticGroupAssignmentDefined = staticGroupAssignmentDefined;
            return this;
        }

        /**
         * @param staticGroupAssignmentDefined staticGroupAssignmentDefined - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder staticGroupAssignmentDefined(Boolean staticGroupAssignmentDefined) {
            return staticGroupAssignmentDefined(Output.of(staticGroupAssignmentDefined));
        }

        /**
         * @param staticProfileAssignment Static Profile Assignment
         * 
         * @return builder
         * 
         */
        public Builder staticProfileAssignment(Output<Boolean> staticProfileAssignment) {
            $.staticProfileAssignment = staticProfileAssignment;
            return this;
        }

        /**
         * @param staticProfileAssignment Static Profile Assignment
         * 
         * @return builder
         * 
         */
        public Builder staticProfileAssignment(Boolean staticProfileAssignment) {
            return staticProfileAssignment(Output.of(staticProfileAssignment));
        }

        /**
         * @param staticProfileAssignmentDefined Static Profile Assignment Defined - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder staticProfileAssignmentDefined(@Nullable Output<Boolean> staticProfileAssignmentDefined) {
            $.staticProfileAssignmentDefined = staticProfileAssignmentDefined;
            return this;
        }

        /**
         * @param staticProfileAssignmentDefined Static Profile Assignment Defined - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder staticProfileAssignmentDefined(Boolean staticProfileAssignmentDefined) {
            return staticProfileAssignmentDefined(Output.of(staticProfileAssignmentDefined));
        }

        public EndpointArgs build() {
            if ($.mac == null) {
                throw new MissingRequiredPropertyException("EndpointArgs", "mac");
            }
            if ($.staticGroupAssignment == null) {
                throw new MissingRequiredPropertyException("EndpointArgs", "staticGroupAssignment");
            }
            if ($.staticProfileAssignment == null) {
                throw new MissingRequiredPropertyException("EndpointArgs", "staticProfileAssignment");
            }
            return $;
        }
    }

}

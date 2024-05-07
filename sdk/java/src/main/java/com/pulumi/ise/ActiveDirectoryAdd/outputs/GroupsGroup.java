// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.ActiveDirectoryAdd.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GroupsGroup {
    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    private String name;
    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    private String sid;
    private @Nullable String type;

    private GroupsGroup() {}
    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Required for each group in the group list with no duplication between groups
     * 
     */
    public String sid() {
        return this.sid;
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String sid;
        private @Nullable String type;
        public Builder() {}
        public Builder(GroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.sid = defaults.sid;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GroupsGroup", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder sid(String sid) {
            if (sid == null) {
              throw new MissingRequiredPropertyException("GroupsGroup", "sid");
            }
            this.sid = sid;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        public GroupsGroup build() {
            final var _resultValue = new GroupsGroup();
            _resultValue.name = name;
            _resultValue.sid = sid;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}

// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.TacacsCommand.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.TacacsCommand.outputs.GetSetCommand;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSetResult {
    private List<GetSetCommand> commands;
    /**
     * @return Description
     * 
     */
    private String description;
    /**
     * @return The id of the object
     * 
     */
    private String id;
    /**
     * @return The name of the TACACS command set
     * 
     */
    private String name;
    /**
     * @return Permit unmatched commands
     * 
     */
    private Boolean permitUnmatched;

    private GetSetResult() {}
    public List<GetSetCommand> commands() {
        return this.commands;
    }
    /**
     * @return Description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The name of the TACACS command set
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Permit unmatched commands
     * 
     */
    public Boolean permitUnmatched() {
        return this.permitUnmatched;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSetResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetSetCommand> commands;
        private String description;
        private String id;
        private String name;
        private Boolean permitUnmatched;
        public Builder() {}
        public Builder(GetSetResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.commands = defaults.commands;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.permitUnmatched = defaults.permitUnmatched;
        }

        @CustomType.Setter
        public Builder commands(List<GetSetCommand> commands) {
            if (commands == null) {
              throw new MissingRequiredPropertyException("GetSetResult", "commands");
            }
            this.commands = commands;
            return this;
        }
        public Builder commands(GetSetCommand... commands) {
            return commands(List.of(commands));
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetSetResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetSetResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetSetResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder permitUnmatched(Boolean permitUnmatched) {
            if (permitUnmatched == null) {
              throw new MissingRequiredPropertyException("GetSetResult", "permitUnmatched");
            }
            this.permitUnmatched = permitUnmatched;
            return this;
        }
        public GetSetResult build() {
            final var _resultValue = new GetSetResult();
            _resultValue.commands = commands;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.permitUnmatched = permitUnmatched;
            return _resultValue;
        }
    }
}

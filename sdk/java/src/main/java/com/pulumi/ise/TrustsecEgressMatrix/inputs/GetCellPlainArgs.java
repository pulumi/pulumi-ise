// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.TrustsecEgressMatrix.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetCellPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCellPlainArgs Empty = new GetCellPlainArgs();

    /**
     * The id of the object
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return The id of the object
     * 
     */
    public String id() {
        return this.id;
    }

    private GetCellPlainArgs() {}

    private GetCellPlainArgs(GetCellPlainArgs $) {
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCellPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCellPlainArgs $;

        public Builder() {
            $ = new GetCellPlainArgs();
        }

        public Builder(GetCellPlainArgs defaults) {
            $ = new GetCellPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id The id of the object
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public GetCellPlainArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetCellPlainArgs", "id");
            }
            return $;
        }
    }

}

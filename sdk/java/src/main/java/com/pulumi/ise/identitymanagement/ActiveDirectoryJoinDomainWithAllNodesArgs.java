// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.identitymanagement;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.ise.identitymanagement.inputs.ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class ActiveDirectoryJoinDomainWithAllNodesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ActiveDirectoryJoinDomainWithAllNodesArgs Empty = new ActiveDirectoryJoinDomainWithAllNodesArgs();

    @Import(name="additionalDatas", required=true)
    private Output<List<ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs>> additionalDatas;

    public Output<List<ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs>> additionalDatas() {
        return this.additionalDatas;
    }

    /**
     * Active Directory Join Point ID
     * 
     */
    @Import(name="joinPointId", required=true)
    private Output<String> joinPointId;

    /**
     * @return Active Directory Join Point ID
     * 
     */
    public Output<String> joinPointId() {
        return this.joinPointId;
    }

    private ActiveDirectoryJoinDomainWithAllNodesArgs() {}

    private ActiveDirectoryJoinDomainWithAllNodesArgs(ActiveDirectoryJoinDomainWithAllNodesArgs $) {
        this.additionalDatas = $.additionalDatas;
        this.joinPointId = $.joinPointId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ActiveDirectoryJoinDomainWithAllNodesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ActiveDirectoryJoinDomainWithAllNodesArgs $;

        public Builder() {
            $ = new ActiveDirectoryJoinDomainWithAllNodesArgs();
        }

        public Builder(ActiveDirectoryJoinDomainWithAllNodesArgs defaults) {
            $ = new ActiveDirectoryJoinDomainWithAllNodesArgs(Objects.requireNonNull(defaults));
        }

        public Builder additionalDatas(Output<List<ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs>> additionalDatas) {
            $.additionalDatas = additionalDatas;
            return this;
        }

        public Builder additionalDatas(List<ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs> additionalDatas) {
            return additionalDatas(Output.of(additionalDatas));
        }

        public Builder additionalDatas(ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs... additionalDatas) {
            return additionalDatas(List.of(additionalDatas));
        }

        /**
         * @param joinPointId Active Directory Join Point ID
         * 
         * @return builder
         * 
         */
        public Builder joinPointId(Output<String> joinPointId) {
            $.joinPointId = joinPointId;
            return this;
        }

        /**
         * @param joinPointId Active Directory Join Point ID
         * 
         * @return builder
         * 
         */
        public Builder joinPointId(String joinPointId) {
            return joinPointId(Output.of(joinPointId));
        }

        public ActiveDirectoryJoinDomainWithAllNodesArgs build() {
            if ($.additionalDatas == null) {
                throw new MissingRequiredPropertyException("ActiveDirectoryJoinDomainWithAllNodesArgs", "additionalDatas");
            }
            if ($.joinPointId == null) {
                throw new MissingRequiredPropertyException("ActiveDirectoryJoinDomainWithAllNodesArgs", "joinPointId");
            }
            return $;
        }
    }

}

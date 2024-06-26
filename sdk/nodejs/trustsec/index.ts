// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { EgressMatrixCellArgs, EgressMatrixCellState } from "./egressMatrixCell";
export type EgressMatrixCell = import("./egressMatrixCell").EgressMatrixCell;
export const EgressMatrixCell: typeof import("./egressMatrixCell").EgressMatrixCell = null as any;
utilities.lazyLoad(exports, ["EgressMatrixCell"], () => require("./egressMatrixCell"));

export { GetEgressMatrixCellArgs, GetEgressMatrixCellResult, GetEgressMatrixCellOutputArgs } from "./getEgressMatrixCell";
export const getEgressMatrixCell: typeof import("./getEgressMatrixCell").getEgressMatrixCell = null as any;
export const getEgressMatrixCellOutput: typeof import("./getEgressMatrixCell").getEgressMatrixCellOutput = null as any;
utilities.lazyLoad(exports, ["getEgressMatrixCell","getEgressMatrixCellOutput"], () => require("./getEgressMatrixCell"));

export { GetIpToSgtMappingArgs, GetIpToSgtMappingResult, GetIpToSgtMappingOutputArgs } from "./getIpToSgtMapping";
export const getIpToSgtMapping: typeof import("./getIpToSgtMapping").getIpToSgtMapping = null as any;
export const getIpToSgtMappingOutput: typeof import("./getIpToSgtMapping").getIpToSgtMappingOutput = null as any;
utilities.lazyLoad(exports, ["getIpToSgtMapping","getIpToSgtMappingOutput"], () => require("./getIpToSgtMapping"));

export { GetIpToSgtMappingGroupArgs, GetIpToSgtMappingGroupResult, GetIpToSgtMappingGroupOutputArgs } from "./getIpToSgtMappingGroup";
export const getIpToSgtMappingGroup: typeof import("./getIpToSgtMappingGroup").getIpToSgtMappingGroup = null as any;
export const getIpToSgtMappingGroupOutput: typeof import("./getIpToSgtMappingGroup").getIpToSgtMappingGroupOutput = null as any;
utilities.lazyLoad(exports, ["getIpToSgtMappingGroup","getIpToSgtMappingGroupOutput"], () => require("./getIpToSgtMappingGroup"));

export { GetSecurityGroupArgs, GetSecurityGroupResult, GetSecurityGroupOutputArgs } from "./getSecurityGroup";
export const getSecurityGroup: typeof import("./getSecurityGroup").getSecurityGroup = null as any;
export const getSecurityGroupOutput: typeof import("./getSecurityGroup").getSecurityGroupOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityGroup","getSecurityGroupOutput"], () => require("./getSecurityGroup"));

export { GetSecurityGroupAclArgs, GetSecurityGroupAclResult, GetSecurityGroupAclOutputArgs } from "./getSecurityGroupAcl";
export const getSecurityGroupAcl: typeof import("./getSecurityGroupAcl").getSecurityGroupAcl = null as any;
export const getSecurityGroupAclOutput: typeof import("./getSecurityGroupAcl").getSecurityGroupAclOutput = null as any;
utilities.lazyLoad(exports, ["getSecurityGroupAcl","getSecurityGroupAclOutput"], () => require("./getSecurityGroupAcl"));

export { IpToSgtMappingArgs, IpToSgtMappingState } from "./ipToSgtMapping";
export type IpToSgtMapping = import("./ipToSgtMapping").IpToSgtMapping;
export const IpToSgtMapping: typeof import("./ipToSgtMapping").IpToSgtMapping = null as any;
utilities.lazyLoad(exports, ["IpToSgtMapping"], () => require("./ipToSgtMapping"));

export { IpToSgtMappingGroupArgs, IpToSgtMappingGroupState } from "./ipToSgtMappingGroup";
export type IpToSgtMappingGroup = import("./ipToSgtMappingGroup").IpToSgtMappingGroup;
export const IpToSgtMappingGroup: typeof import("./ipToSgtMappingGroup").IpToSgtMappingGroup = null as any;
utilities.lazyLoad(exports, ["IpToSgtMappingGroup"], () => require("./ipToSgtMappingGroup"));

export { SecurityGroupArgs, SecurityGroupState } from "./securityGroup";
export type SecurityGroup = import("./securityGroup").SecurityGroup;
export const SecurityGroup: typeof import("./securityGroup").SecurityGroup = null as any;
utilities.lazyLoad(exports, ["SecurityGroup"], () => require("./securityGroup"));

export { SecurityGroupAclArgs, SecurityGroupAclState } from "./securityGroupAcl";
export type SecurityGroupAcl = import("./securityGroupAcl").SecurityGroupAcl;
export const SecurityGroupAcl: typeof import("./securityGroupAcl").SecurityGroupAcl = null as any;
utilities.lazyLoad(exports, ["SecurityGroupAcl"], () => require("./securityGroupAcl"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:trustsec/egressMatrixCell:EgressMatrixCell":
                return new EgressMatrixCell(name, <any>undefined, { urn })
            case "ise:trustsec/ipToSgtMapping:IpToSgtMapping":
                return new IpToSgtMapping(name, <any>undefined, { urn })
            case "ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup":
                return new IpToSgtMappingGroup(name, <any>undefined, { urn })
            case "ise:trustsec/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "ise:trustsec/securityGroupAcl:SecurityGroupAcl":
                return new SecurityGroupAcl(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "trustsec/egressMatrixCell", _module)
pulumi.runtime.registerResourceModule("ise", "trustsec/ipToSgtMapping", _module)
pulumi.runtime.registerResourceModule("ise", "trustsec/ipToSgtMappingGroup", _module)
pulumi.runtime.registerResourceModule("ise", "trustsec/securityGroup", _module)
pulumi.runtime.registerResourceModule("ise", "trustsec/securityGroupAcl", _module)

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetLicenseTierStateArgs, GetLicenseTierStateResult, GetLicenseTierStateOutputArgs } from "./getLicenseTierState";
export const getLicenseTierState: typeof import("./getLicenseTierState").getLicenseTierState = null as any;
export const getLicenseTierStateOutput: typeof import("./getLicenseTierState").getLicenseTierStateOutput = null as any;
utilities.lazyLoad(exports, ["getLicenseTierState","getLicenseTierStateOutput"], () => require("./getLicenseTierState"));

export { GetRepositoryArgs, GetRepositoryResult, GetRepositoryOutputArgs } from "./getRepository";
export const getRepository: typeof import("./getRepository").getRepository = null as any;
export const getRepositoryOutput: typeof import("./getRepository").getRepositoryOutput = null as any;
utilities.lazyLoad(exports, ["getRepository","getRepositoryOutput"], () => require("./getRepository"));

export { LicenseTierStateArgs, LicenseTierStateState } from "./licenseTierState";
export type LicenseTierState = import("./licenseTierState").LicenseTierState;
export const LicenseTierState: typeof import("./licenseTierState").LicenseTierState = null as any;
utilities.lazyLoad(exports, ["LicenseTierState"], () => require("./licenseTierState"));

export { RepositoryArgs, RepositoryState } from "./repository";
export type Repository = import("./repository").Repository;
export const Repository: typeof import("./repository").Repository = null as any;
utilities.lazyLoad(exports, ["Repository"], () => require("./repository"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:system/licenseTierState:LicenseTierState":
                return new LicenseTierState(name, <any>undefined, { urn })
            case "ise:system/repository:Repository":
                return new Repository(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "system/licenseTierState", _module)
pulumi.runtime.registerResourceModule("ise", "system/repository", _module)
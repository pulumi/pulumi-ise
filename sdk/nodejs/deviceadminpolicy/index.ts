// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetSetArgs, GetSetResult, GetSetOutputArgs } from "./getSet";
export const getSet: typeof import("./getSet").getSet = null as any;
export const getSetOutput: typeof import("./getSet").getSetOutput = null as any;
utilities.lazyLoad(exports, ["getSet","getSetOutput"], () => require("./getSet"));

export { SetArgs, SetState } from "./set";
export type Set = import("./set").Set;
export const Set: typeof import("./set").Set = null as any;
utilities.lazyLoad(exports, ["Set"], () => require("./set"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:DeviceAdminPolicy/set:Set":
                return new Set(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "DeviceAdminPolicy/set", _module)

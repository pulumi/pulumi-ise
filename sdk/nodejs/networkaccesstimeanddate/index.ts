// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ConditionArgs, ConditionState } from "./condition";
export type Condition = import("./condition").Condition;
export const Condition: typeof import("./condition").Condition = null as any;
utilities.lazyLoad(exports, ["Condition"], () => require("./condition"));

export { GetConditionArgs, GetConditionResult, GetConditionOutputArgs } from "./getCondition";
export const getCondition: typeof import("./getCondition").getCondition = null as any;
export const getConditionOutput: typeof import("./getCondition").getConditionOutput = null as any;
utilities.lazyLoad(exports, ["getCondition","getConditionOutput"], () => require("./getCondition"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:NetworkAccessTimeAndDate/condition:Condition":
                return new Condition(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "NetworkAccessTimeAndDate/condition", _module)

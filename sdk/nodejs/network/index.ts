// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DeviceArgs, DeviceState } from "./device";
export type Device = import("./device").Device;
export const Device: typeof import("./device").Device = null as any;
utilities.lazyLoad(exports, ["Device"], () => require("./device"));

export { DeviceGroupArgs, DeviceGroupState } from "./deviceGroup";
export type DeviceGroup = import("./deviceGroup").DeviceGroup;
export const DeviceGroup: typeof import("./deviceGroup").DeviceGroup = null as any;
utilities.lazyLoad(exports, ["DeviceGroup"], () => require("./deviceGroup"));

export { GetDeviceArgs, GetDeviceResult, GetDeviceOutputArgs } from "./getDevice";
export const getDevice: typeof import("./getDevice").getDevice = null as any;
export const getDeviceOutput: typeof import("./getDevice").getDeviceOutput = null as any;
utilities.lazyLoad(exports, ["getDevice","getDeviceOutput"], () => require("./getDevice"));

export { GetDeviceGroupArgs, GetDeviceGroupResult, GetDeviceGroupOutputArgs } from "./getDeviceGroup";
export const getDeviceGroup: typeof import("./getDeviceGroup").getDeviceGroup = null as any;
export const getDeviceGroupOutput: typeof import("./getDeviceGroup").getDeviceGroupOutput = null as any;
utilities.lazyLoad(exports, ["getDeviceGroup","getDeviceGroupOutput"], () => require("./getDeviceGroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:network/device:Device":
                return new Device(name, <any>undefined, { urn })
            case "ise:network/deviceGroup:DeviceGroup":
                return new DeviceGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "network/device", _module)
pulumi.runtime.registerResourceModule("ise", "network/deviceGroup", _module)

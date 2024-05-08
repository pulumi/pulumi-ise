// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AllowedProtocolsArgs, AllowedProtocolsState } from "./allowedProtocols";
export type AllowedProtocols = import("./allowedProtocols").AllowedProtocols;
export const AllowedProtocols: typeof import("./allowedProtocols").AllowedProtocols = null as any;
utilities.lazyLoad(exports, ["AllowedProtocols"], () => require("./allowedProtocols"));

export { CommandSetArgs, CommandSetState } from "./commandSet";
export type CommandSet = import("./commandSet").CommandSet;
export const CommandSet: typeof import("./commandSet").CommandSet = null as any;
utilities.lazyLoad(exports, ["CommandSet"], () => require("./commandSet"));

export { GetAllowedProtocolsArgs, GetAllowedProtocolsResult, GetAllowedProtocolsOutputArgs } from "./getAllowedProtocols";
export const getAllowedProtocols: typeof import("./getAllowedProtocols").getAllowedProtocols = null as any;
export const getAllowedProtocolsOutput: typeof import("./getAllowedProtocols").getAllowedProtocolsOutput = null as any;
utilities.lazyLoad(exports, ["getAllowedProtocols","getAllowedProtocolsOutput"], () => require("./getAllowedProtocols"));

export { GetCommandSetArgs, GetCommandSetResult, GetCommandSetOutputArgs } from "./getCommandSet";
export const getCommandSet: typeof import("./getCommandSet").getCommandSet = null as any;
export const getCommandSetOutput: typeof import("./getCommandSet").getCommandSetOutput = null as any;
utilities.lazyLoad(exports, ["getCommandSet","getCommandSetOutput"], () => require("./getCommandSet"));

export { GetProfileArgs, GetProfileResult, GetProfileOutputArgs } from "./getProfile";
export const getProfile: typeof import("./getProfile").getProfile = null as any;
export const getProfileOutput: typeof import("./getProfile").getProfileOutput = null as any;
utilities.lazyLoad(exports, ["getProfile","getProfileOutput"], () => require("./getProfile"));

export { ProfileArgs, ProfileState } from "./profile";
export type Profile = import("./profile").Profile;
export const Profile: typeof import("./profile").Profile = null as any;
utilities.lazyLoad(exports, ["Profile"], () => require("./profile"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "ise:tacacs/allowedProtocols:AllowedProtocols":
                return new AllowedProtocols(name, <any>undefined, { urn })
            case "ise:tacacs/commandSet:CommandSet":
                return new CommandSet(name, <any>undefined, { urn })
            case "ise:tacacs/profile:Profile":
                return new Profile(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("ise", "tacacs/allowedProtocols", _module)
pulumi.runtime.registerResourceModule("ise", "tacacs/commandSet", _module)
pulumi.runtime.registerResourceModule("ise", "tacacs/profile", _module)

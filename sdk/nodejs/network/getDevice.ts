// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Network Device.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.Network.getDevice({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getDevice(args?: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:Network/getDevice:getDevice", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDevice.
 */
export interface GetDeviceArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the network device
     */
    name?: string;
}

/**
 * A collection of values returned by getDevice.
 */
export interface GetDeviceResult {
    /**
     * Enforce use of DTLS
     */
    readonly authenticationDtlsRequired: boolean;
    /**
     * Enable key wrap
     */
    readonly authenticationEnableKeyWrap: boolean;
    /**
     * Enable multiple RADIUS shared secrets
     */
    readonly authenticationEnableMultiSecret: boolean;
    /**
     * Encryption key
     */
    readonly authenticationEncryptionKey: string;
    /**
     * Key input format
     */
    readonly authenticationEncryptionKeyFormat: string;
    /**
     * Message authenticator code key
     */
    readonly authenticationMessageAuthenticatorCodeKey: string;
    /**
     * Network protocol
     */
    readonly authenticationNetworkProtocol: string;
    /**
     * RADIUS shared secret
     */
    readonly authenticationRadiusSharedSecret: string;
    /**
     * Second RADIUS shared secret
     */
    readonly authenticationSecondRadiusSharedSecret: string;
    /**
     * CoA port
     */
    readonly coaPort: number;
    /**
     * Description
     */
    readonly description: string;
    /**
     * This value is used to verify the client identity contained in the X.509 RADIUS/DTLS client certificate
     */
    readonly dtlsDnsName: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * List of IP subnets
     */
    readonly ips: outputs.Network.GetDeviceIp[];
    /**
     * Model name
     */
    readonly modelName: string;
    /**
     * The name of the network device
     */
    readonly name: string;
    /**
     * List of network device groups, e.g. `Device Type#All Device Types#ACCESS`
     */
    readonly networkDeviceGroups: string[];
    /**
     * Profile name
     */
    readonly profileName: string;
    /**
     * SNMP link Trap Query
     */
    readonly snmpLinkTrapQuery: boolean;
    /**
     * SNMP MAC Trap Query
     */
    readonly snmpMacTrapQuery: boolean;
    /**
     * Originating Policy Services Node
     */
    readonly snmpOriginatingPolicyServiceNode: string;
    /**
     * SNMP Polling Interval in seconds
     */
    readonly snmpPollingInterval: number;
    /**
     * SNMP RO Community
     */
    readonly snmpRoCommunity: string;
    /**
     * SNMP version
     */
    readonly snmpVersion: string;
    /**
     * Software version
     */
    readonly softwareVersion: string;
    /**
     * Connect mode options
     */
    readonly tacacsConnectModeOptions: string;
    /**
     * Shared secret
     */
    readonly tacacsSharedSecret: string;
    /**
     * CoA source host
     */
    readonly trustsecCoaSourceHost: string;
    /**
     * TrustSec device ID
     */
    readonly trustsecDeviceId: string;
    /**
     * TrustSec device password
     */
    readonly trustsecDevicePassword: string;
    /**
     * Download environment data every X seconds
     */
    readonly trustsecDownloadEnviromentDataEveryXSeconds: number;
    /**
     * Download peer authorization policy every X seconds
     */
    readonly trustsecDownloadPeerAuthorizationPolicyEveryXSeconds: number;
    /**
     * Download SGACL lists every X seconds
     */
    readonly trustsecDownloadSgaclListsEveryXSeconds: number;
    /**
     * Enable mode password
     */
    readonly trustsecEnableModePassword: string;
    /**
     * EXEC mode password
     */
    readonly trustsecExecModePassword: string;
    /**
     * EXEC mode username
     */
    readonly trustsecExecModeUsername: string;
    /**
     * Include this device when deploying Security Group Tag Mapping Updates
     */
    readonly trustsecIncludeWhenDeployingSgtUpdates: boolean;
    /**
     * Other TrustSec devices to trust this device
     */
    readonly trustsecOtherSgaDevicesToTrustThisDevice: boolean;
    /**
     * Re-authenticate every X seconds
     */
    readonly trustsecReAuthenticationEveryXSeconds: number;
    /**
     * REST API password
     */
    readonly trustsecRestApiPassword: string;
    /**
     * REST API username
     */
    readonly trustsecRestApiUsername: string;
    /**
     * Send configuration to device
     */
    readonly trustsecSendConfigurationToDevice: boolean;
    /**
     * Send configuration to device using
     */
    readonly trustsecSendConfigurationToDeviceUsing: string;
}
/**
 * This data source can read the Network Device.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.Network.getDevice({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getDeviceOutput(args?: GetDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceResult> {
    return pulumi.output(args).apply((a: any) => getDevice(a, opts))
}

/**
 * A collection of arguments for invoking getDevice.
 */
export interface GetDeviceOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the network device
     */
    name?: pulumi.Input<string>;
}

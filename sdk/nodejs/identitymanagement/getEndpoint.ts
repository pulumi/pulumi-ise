// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getEndpoint({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getEndpoint(args?: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:identitymanagement/getEndpoint:getEndpoint", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoint.
 */
export interface GetEndpointArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the endpoint
     */
    name?: string;
}

/**
 * A collection of values returned by getEndpoint.
 */
export interface GetEndpointResult {
    /**
     * Custom Attributes
     */
    readonly customAttributes: {[key: string]: string};
    /**
     * Description
     */
    readonly description: string;
    /**
     * Identity Group ID
     */
    readonly groupId: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Identity Store
     */
    readonly identityStore: string;
    /**
     * Identity Store Id
     */
    readonly identityStoreId: string;
    /**
     * MAC address of the endpoint
     */
    readonly mac: string;
    /**
     * Mdm Compliance Status
     */
    readonly mdmComplianceStatus: boolean;
    /**
     * Mdm Encrypted
     */
    readonly mdmEncrypted: boolean;
    /**
     * Mdm Enrolled
     */
    readonly mdmEnrolled: boolean;
    /**
     * Mdm IMEI
     */
    readonly mdmImei: string;
    /**
     * Mdm JailBroken
     */
    readonly mdmJailBroken: boolean;
    /**
     * Mdm Manufacturer
     */
    readonly mdmManufacturer: string;
    /**
     * Mdm Model
     */
    readonly mdmModel: string;
    /**
     * Mdm OS
     */
    readonly mdmOs: string;
    /**
     * Mdm PhoneNumber
     */
    readonly mdmPhoneNumber: string;
    /**
     * Mdm Pinlock
     */
    readonly mdmPinlock: boolean;
    /**
     * Mdm Reachable
     */
    readonly mdmReachable: boolean;
    /**
     * Mdm Serial
     */
    readonly mdmSerial: string;
    /**
     * Mdm Server Name
     */
    readonly mdmServerName: string;
    /**
     * The name of the endpoint
     */
    readonly name: string;
    /**
     * Portal User
     */
    readonly portalUser: string;
    /**
     * Profile ID
     */
    readonly profileId: string;
    /**
     * Static Group Assignment
     */
    readonly staticGroupAssignment: boolean;
    /**
     * staticGroupAssignmentDefined
     */
    readonly staticGroupAssignmentDefined: boolean;
    /**
     * Static Profile Assignment
     */
    readonly staticProfileAssignment: boolean;
    /**
     * Static Profile Assignment Defined
     */
    readonly staticProfileAssignmentDefined: boolean;
}
/**
 * This data source can read the Endpoint.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getEndpoint({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getEndpointOutput(args?: GetEndpointOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEndpointResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:identitymanagement/getEndpoint:getEndpoint", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getEndpoint.
 */
export interface GetEndpointOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the endpoint
     */
    name?: pulumi.Input<string>;
}

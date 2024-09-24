// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the TrustSec Egress Matrix Cell.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getEgressMatrixCell({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getEgressMatrixCell(args: GetEgressMatrixCellArgs, opts?: pulumi.InvokeOptions): Promise<GetEgressMatrixCellResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getEgressMatrixCell.
 */
export interface GetEgressMatrixCellArgs {
    /**
     * The id of the object
     */
    id: string;
}

/**
 * A collection of values returned by getEgressMatrixCell.
 */
export interface GetEgressMatrixCellResult {
    /**
     * Can be used only if sgacls not specified.
     */
    readonly defaultRule: string;
    /**
     * Description
     */
    readonly description: string;
    /**
     * Destination Trustsec Security Group ID
     */
    readonly destinationSgtId: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Matrix Cell Status
     */
    readonly matrixCellStatus: string;
    /**
     * List of TrustSec Security Groups ACLs
     */
    readonly sgacls: string[];
    /**
     * Source Trustsec Security Group ID
     */
    readonly sourceSgtId: string;
}
/**
 * This data source can read the TrustSec Egress Matrix Cell.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getEgressMatrixCell({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getEgressMatrixCellOutput(args: GetEgressMatrixCellOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEgressMatrixCellResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getEgressMatrixCell.
 */
export interface GetEgressMatrixCellOutputArgs {
    /**
     * The id of the object
     */
    id: pulumi.Input<string>;
}
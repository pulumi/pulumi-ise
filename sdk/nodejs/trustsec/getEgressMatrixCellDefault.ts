// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the TrustSec Egress Matrix Cell Default.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getEgressMatrixCellDefault({
 *     id: "92c1a900-8c01-11e6-996c-525400b48521",
 * });
 * ```
 */
export function getEgressMatrixCellDefault(args: GetEgressMatrixCellDefaultArgs, opts?: pulumi.InvokeOptions): Promise<GetEgressMatrixCellDefaultResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:trustsec/getEgressMatrixCellDefault:getEgressMatrixCellDefault", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getEgressMatrixCellDefault.
 */
export interface GetEgressMatrixCellDefaultArgs {
    /**
     * The id of the object
     */
    id: string;
}

/**
 * A collection of values returned by getEgressMatrixCellDefault.
 */
export interface GetEgressMatrixCellDefaultResult {
    /**
     * Can be used only if sgacls not specified. Final Catch All Rule
     */
    readonly defaultRule: string;
    /**
     * Description
     */
    readonly description: string;
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
}
/**
 * This data source can read the TrustSec Egress Matrix Cell Default.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getEgressMatrixCellDefault({
 *     id: "92c1a900-8c01-11e6-996c-525400b48521",
 * });
 * ```
 */
export function getEgressMatrixCellDefaultOutput(args: GetEgressMatrixCellDefaultOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetEgressMatrixCellDefaultResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:trustsec/getEgressMatrixCellDefault:getEgressMatrixCellDefault", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getEgressMatrixCellDefault.
 */
export interface GetEgressMatrixCellDefaultOutputArgs {
    /**
     * The id of the object
     */
    id: pulumi.Input<string>;
}

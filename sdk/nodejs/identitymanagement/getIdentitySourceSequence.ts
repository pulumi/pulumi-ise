// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Identity Source Sequence.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getIdentitySourceSequence({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getIdentitySourceSequence(args?: GetIdentitySourceSequenceArgs, opts?: pulumi.InvokeOptions): Promise<GetIdentitySourceSequenceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:identitymanagement/getIdentitySourceSequence:getIdentitySourceSequence", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentitySourceSequence.
 */
export interface GetIdentitySourceSequenceArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the identity source sequence
     */
    name?: string;
}

/**
 * A collection of values returned by getIdentitySourceSequence.
 */
export interface GetIdentitySourceSequenceResult {
    /**
     * Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     */
    readonly breakOnStoreFail: boolean;
    /**
     * Certificate Authentication Profile, empty if doesn't exist
     */
    readonly certificateAuthenticationProfile: string;
    /**
     * Description
     */
    readonly description: string;
    /**
     * The id of the object
     */
    readonly id: string;
    readonly identitySources: outputs.identitymanagement.GetIdentitySourceSequenceIdentitySource[];
    /**
     * The name of the identity source sequence
     */
    readonly name: string;
}
/**
 * This data source can read the Identity Source Sequence.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getIdentitySourceSequence({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getIdentitySourceSequenceOutput(args?: GetIdentitySourceSequenceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIdentitySourceSequenceResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:identitymanagement/getIdentitySourceSequence:getIdentitySourceSequence", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIdentitySourceSequence.
 */
export interface GetIdentitySourceSequenceOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the identity source sequence
     */
    name?: pulumi.Input<string>;
}

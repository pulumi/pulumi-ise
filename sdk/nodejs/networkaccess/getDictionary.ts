// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the Network Access Dictionary.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getDictionary({
 *     id: "Dict1",
 * });
 * ```
 */
export function getDictionary(args?: GetDictionaryArgs, opts?: pulumi.InvokeOptions): Promise<GetDictionaryResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:networkaccess/getDictionary:getDictionary", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDictionary.
 */
export interface GetDictionaryArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The dictionary name
     */
    name?: string;
}

/**
 * A collection of values returned by getDictionary.
 */
export interface GetDictionaryResult {
    /**
     * The description of the dictionary
     */
    readonly description: string;
    /**
     * The dictionary attribute type
     */
    readonly dictionaryAttrType: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * The dictionary name
     */
    readonly name: string;
    /**
     * The version of the dictionary
     */
    readonly version: string;
}
/**
 * This data source can read the Network Access Dictionary.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getDictionary({
 *     id: "Dict1",
 * });
 * ```
 */
export function getDictionaryOutput(args?: GetDictionaryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDictionaryResult> {
    return pulumi.output(args).apply((a: any) => getDictionary(a, opts))
}

/**
 * A collection of arguments for invoking getDictionary.
 */
export interface GetDictionaryOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The dictionary name
     */
    name?: pulumi.Input<string>;
}

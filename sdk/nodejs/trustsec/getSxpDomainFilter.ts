// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the SXP Domain Filter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getSxpDomainFilter({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getSxpDomainFilter(args?: GetSxpDomainFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetSxpDomainFilterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:trustsec/getSxpDomainFilter:getSxpDomainFilter", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSxpDomainFilter.
 */
export interface GetSxpDomainFilterArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * Resource name
     */
    name?: string;
}

/**
 * A collection of values returned by getSxpDomainFilter.
 */
export interface GetSxpDomainFilterResult {
    /**
     * Description
     */
    readonly description: string;
    /**
     * List of SXP Domains, separated with comma
     */
    readonly domains: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * SGT name or ID. At least one of subnet or sgt or vn should be defined
     */
    readonly sgt: string;
    /**
     * Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
     */
    readonly subnet: string;
    /**
     * Virtual Network. At least one of subnet or sgt or vn should be defined
     */
    readonly vn: string;
}
/**
 * This data source can read the SXP Domain Filter.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getSxpDomainFilter({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getSxpDomainFilterOutput(args?: GetSxpDomainFilterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSxpDomainFilterResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:trustsec/getSxpDomainFilter:getSxpDomainFilter", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSxpDomainFilter.
 */
export interface GetSxpDomainFilterOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * Resource name
     */
    name?: pulumi.Input<string>;
}

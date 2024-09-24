// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This data source can read the Active Directory Groups By Domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getActiveDirectoryGroupsByDomain({
 *     joinPointId: "73808580-b6e6-11ee-8960-de6d7692bc40",
 *     domain: "cisco.com",
 *     filter: "CN=ISE Admins",
 *     sidFilter: "cisco.com/S-1-5-33-544",
 *     typeFilter: "UNIVERSAL",
 * });
 * ```
 */
export function getActiveDirectoryGroupsByDomain(args: GetActiveDirectoryGroupsByDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetActiveDirectoryGroupsByDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:identitymanagement/getActiveDirectoryGroupsByDomain:getActiveDirectoryGroupsByDomain", {
        "domain": args.domain,
        "filter": args.filter,
        "joinPointId": args.joinPointId,
        "sidFilter": args.sidFilter,
        "typeFilter": args.typeFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getActiveDirectoryGroupsByDomain.
 */
export interface GetActiveDirectoryGroupsByDomainArgs {
    /**
     * The domain whose groups we want to fetch
     */
    domain: string;
    /**
     * Exact match filter on group's CN
     */
    filter?: string;
    /**
     * Active Directory Join Point ID
     */
    joinPointId: string;
    /**
     * Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
     */
    sidFilter?: string;
    /**
     * Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
     */
    typeFilter?: string;
}

/**
 * A collection of values returned by getActiveDirectoryGroupsByDomain.
 */
export interface GetActiveDirectoryGroupsByDomainResult {
    /**
     * The domain whose groups we want to fetch
     */
    readonly domain: string;
    /**
     * Exact match filter on group's CN
     */
    readonly filter?: string;
    /**
     * List of groups
     */
    readonly groups: outputs.identitymanagement.GetActiveDirectoryGroupsByDomainGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Active Directory Join Point ID
     */
    readonly joinPointId: string;
    /**
     * Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
     */
    readonly sidFilter?: string;
    /**
     * Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
     */
    readonly typeFilter?: string;
}
/**
 * This data source can read the Active Directory Groups By Domain.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.identitymanagement.getActiveDirectoryGroupsByDomain({
 *     joinPointId: "73808580-b6e6-11ee-8960-de6d7692bc40",
 *     domain: "cisco.com",
 *     filter: "CN=ISE Admins",
 *     sidFilter: "cisco.com/S-1-5-33-544",
 *     typeFilter: "UNIVERSAL",
 * });
 * ```
 */
export function getActiveDirectoryGroupsByDomainOutput(args: GetActiveDirectoryGroupsByDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetActiveDirectoryGroupsByDomainResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:identitymanagement/getActiveDirectoryGroupsByDomain:getActiveDirectoryGroupsByDomain", {
        "domain": args.domain,
        "filter": args.filter,
        "joinPointId": args.joinPointId,
        "sidFilter": args.sidFilter,
        "typeFilter": args.typeFilter,
    }, opts);
}

/**
 * A collection of arguments for invoking getActiveDirectoryGroupsByDomain.
 */
export interface GetActiveDirectoryGroupsByDomainOutputArgs {
    /**
     * The domain whose groups we want to fetch
     */
    domain: pulumi.Input<string>;
    /**
     * Exact match filter on group's CN
     */
    filter?: pulumi.Input<string>;
    /**
     * Active Directory Join Point ID
     */
    joinPointId: pulumi.Input<string>;
    /**
     * Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
     */
    sidFilter?: pulumi.Input<string>;
    /**
     * Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
     */
    typeFilter?: pulumi.Input<string>;
}
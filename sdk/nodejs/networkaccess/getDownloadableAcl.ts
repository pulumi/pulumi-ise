// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the Downloadable ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getDownloadableAcl({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getDownloadableAcl(args?: GetDownloadableAclArgs, opts?: pulumi.InvokeOptions): Promise<GetDownloadableAclResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:networkaccess/getDownloadableAcl:getDownloadableAcl", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDownloadableAcl.
 */
export interface GetDownloadableAclArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the downloadable ACL
     */
    name?: string;
}

/**
 * A collection of values returned by getDownloadableAcl.
 */
export interface GetDownloadableAclResult {
    /**
     * The DACL content
     */
    readonly dacl: string;
    /**
     * The type of ACL
     */
    readonly daclType: string;
    /**
     * Description
     */
    readonly description: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * The name of the downloadable ACL
     */
    readonly name: string;
}
/**
 * This data source can read the Downloadable ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.networkaccess.getDownloadableAcl({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getDownloadableAclOutput(args?: GetDownloadableAclOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDownloadableAclResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:networkaccess/getDownloadableAcl:getDownloadableAcl", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getDownloadableAcl.
 */
export interface GetDownloadableAclOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the downloadable ACL
     */
    name?: pulumi.Input<string>;
}

// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This data source can read the TrustSec Security Group ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getSecurityGroupAcl({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getSecurityGroupAcl(args?: GetSecurityGroupAclArgs, opts?: pulumi.InvokeOptions): Promise<GetSecurityGroupAclResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroupAcl.
 */
export interface GetSecurityGroupAclArgs {
    /**
     * The id of the object
     */
    id?: string;
    /**
     * The name of the security group ACL
     */
    name?: string;
}

/**
 * A collection of values returned by getSecurityGroupAcl.
 */
export interface GetSecurityGroupAclResult {
    /**
     * Content of ACL
     */
    readonly aclContent: string;
    /**
     * Description
     */
    readonly description: string;
    /**
     * The id of the object
     */
    readonly id: string;
    /**
     * IP Version
     */
    readonly ipVersion: string;
    /**
     * The name of the security group ACL
     */
    readonly name: string;
    /**
     * Read-only
     */
    readonly readOnly: boolean;
}
/**
 * This data source can read the TrustSec Security Group ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = ise.trustsec.getSecurityGroupAcl({
 *     id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
 * });
 * ```
 */
export function getSecurityGroupAclOutput(args?: GetSecurityGroupAclOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetSecurityGroupAclResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecurityGroupAcl.
 */
export interface GetSecurityGroupAclOutputArgs {
    /**
     * The id of the object
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the security group ACL
     */
    name?: pulumi.Input<string>;
}

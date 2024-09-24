// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can manage a Downloadable ACL.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.networkaccess.DownloadableAcl("example", {
 *     name: "MyACL",
 *     description: "My first downloadable ACL",
 *     dacl: "permit ip any any",
 *     daclType: "IPV4",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:networkaccess/downloadableAcl:DownloadableAcl example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class DownloadableAcl extends pulumi.CustomResource {
    /**
     * Get an existing DownloadableAcl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DownloadableAclState, opts?: pulumi.CustomResourceOptions): DownloadableAcl {
        return new DownloadableAcl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:networkaccess/downloadableAcl:DownloadableAcl';

    /**
     * Returns true if the given object is an instance of DownloadableAcl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DownloadableAcl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DownloadableAcl.__pulumiType;
    }

    /**
     * The DACL content
     */
    public readonly dacl!: pulumi.Output<string>;
    /**
     * The type of ACL - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IPV4`
     */
    public readonly daclType!: pulumi.Output<string>;
    /**
     * Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the downloadable ACL
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a DownloadableAcl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DownloadableAclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DownloadableAclArgs | DownloadableAclState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DownloadableAclState | undefined;
            resourceInputs["dacl"] = state ? state.dacl : undefined;
            resourceInputs["daclType"] = state ? state.daclType : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as DownloadableAclArgs | undefined;
            if ((!args || args.dacl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dacl'");
            }
            resourceInputs["dacl"] = args ? args.dacl : undefined;
            resourceInputs["daclType"] = args ? args.daclType : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DownloadableAcl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DownloadableAcl resources.
 */
export interface DownloadableAclState {
    /**
     * The DACL content
     */
    dacl?: pulumi.Input<string>;
    /**
     * The type of ACL - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IPV4`
     */
    daclType?: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the downloadable ACL
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DownloadableAcl resource.
 */
export interface DownloadableAclArgs {
    /**
     * The DACL content
     */
    dacl: pulumi.Input<string>;
    /**
     * The type of ACL - Choices: `IPV4`, `IPV6`, `IP_AGNOSTIC` - Default value: `IPV4`
     */
    daclType?: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the downloadable ACL
     */
    name?: pulumi.Input<string>;
}
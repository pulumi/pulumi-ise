// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can manage a Network Device Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.networkdevice.Group("example", {
 *     name: "Device Type#All Device Types#Group1",
 *     description: "My network device group",
 *     rootGroup: "Device Type",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:NetworkDevice/group:Group example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:NetworkDevice/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the root device group.
     */
    public readonly rootGroup!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rootGroup"] = state ? state.rootGroup : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.rootGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rootGroup'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rootGroup"] = args ? args.rootGroup : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the root device group.
     */
    rootGroup?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the root device group.
     */
    rootGroup: pulumi.Input<string>;
}

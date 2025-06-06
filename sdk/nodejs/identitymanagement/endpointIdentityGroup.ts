// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can manage an Endpoint Identity Group.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.identitymanagement.EndpointIdentityGroup("example", {
 *     name: "Group1",
 *     description: "My endpoint identity group",
 *     systemDefined: false,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class EndpointIdentityGroup extends pulumi.CustomResource {
    /**
     * Get an existing EndpointIdentityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EndpointIdentityGroupState, opts?: pulumi.CustomResourceOptions): EndpointIdentityGroup {
        return new EndpointIdentityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup';

    /**
     * Returns true if the given object is an instance of EndpointIdentityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EndpointIdentityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EndpointIdentityGroup.__pulumiType;
    }

    /**
     * Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the endpoint identity group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Parent endpoint identity group ID
     */
    public readonly parentEndpointIdentityGroupId!: pulumi.Output<string | undefined>;
    /**
     * System defined endpoint identity group
     */
    public readonly systemDefined!: pulumi.Output<boolean | undefined>;

    /**
     * Create a EndpointIdentityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: EndpointIdentityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EndpointIdentityGroupArgs | EndpointIdentityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EndpointIdentityGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentEndpointIdentityGroupId"] = state ? state.parentEndpointIdentityGroupId : undefined;
            resourceInputs["systemDefined"] = state ? state.systemDefined : undefined;
        } else {
            const args = argsOrState as EndpointIdentityGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parentEndpointIdentityGroupId"] = args ? args.parentEndpointIdentityGroupId : undefined;
            resourceInputs["systemDefined"] = args ? args.systemDefined : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EndpointIdentityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EndpointIdentityGroup resources.
 */
export interface EndpointIdentityGroupState {
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the endpoint identity group
     */
    name?: pulumi.Input<string>;
    /**
     * Parent endpoint identity group ID
     */
    parentEndpointIdentityGroupId?: pulumi.Input<string>;
    /**
     * System defined endpoint identity group
     */
    systemDefined?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a EndpointIdentityGroup resource.
 */
export interface EndpointIdentityGroupArgs {
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the endpoint identity group
     */
    name?: pulumi.Input<string>;
    /**
     * Parent endpoint identity group ID
     */
    parentEndpointIdentityGroupId?: pulumi.Input<string>;
    /**
     * System defined endpoint identity group
     */
    systemDefined?: pulumi.Input<boolean>;
}

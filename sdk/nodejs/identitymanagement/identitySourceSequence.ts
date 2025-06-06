// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * This resource can manage an Identity Source Sequence.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.identitymanagement.IdentitySourceSequence("example", {
 *     name: "Sequence1",
 *     description: "My identity source sequence",
 *     breakOnStoreFail: true,
 *     certificateAuthenticationProfile: "Preloaded_Certificate_Profile",
 *     identitySources: [{
 *         name: "Internal Users",
 *         order: 1,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:identitymanagement/identitySourceSequence:IdentitySourceSequence example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class IdentitySourceSequence extends pulumi.CustomResource {
    /**
     * Get an existing IdentitySourceSequence resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentitySourceSequenceState, opts?: pulumi.CustomResourceOptions): IdentitySourceSequence {
        return new IdentitySourceSequence(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:identitymanagement/identitySourceSequence:IdentitySourceSequence';

    /**
     * Returns true if the given object is an instance of IdentitySourceSequence.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentitySourceSequence {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentitySourceSequence.__pulumiType;
    }

    /**
     * Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     */
    public readonly breakOnStoreFail!: pulumi.Output<boolean>;
    /**
     * Certificate Authentication Profile, empty if doesn't exist
     */
    public readonly certificateAuthenticationProfile!: pulumi.Output<string>;
    /**
     * Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly identitySources!: pulumi.Output<outputs.identitymanagement.IdentitySourceSequenceIdentitySource[]>;
    /**
     * The name of the identity source sequence
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a IdentitySourceSequence resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentitySourceSequenceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentitySourceSequenceArgs | IdentitySourceSequenceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IdentitySourceSequenceState | undefined;
            resourceInputs["breakOnStoreFail"] = state ? state.breakOnStoreFail : undefined;
            resourceInputs["certificateAuthenticationProfile"] = state ? state.certificateAuthenticationProfile : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["identitySources"] = state ? state.identitySources : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as IdentitySourceSequenceArgs | undefined;
            if ((!args || args.breakOnStoreFail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'breakOnStoreFail'");
            }
            if ((!args || args.certificateAuthenticationProfile === undefined) && !opts.urn) {
                throw new Error("Missing required property 'certificateAuthenticationProfile'");
            }
            if ((!args || args.identitySources === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identitySources'");
            }
            resourceInputs["breakOnStoreFail"] = args ? args.breakOnStoreFail : undefined;
            resourceInputs["certificateAuthenticationProfile"] = args ? args.certificateAuthenticationProfile : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identitySources"] = args ? args.identitySources : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IdentitySourceSequence.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentitySourceSequence resources.
 */
export interface IdentitySourceSequenceState {
    /**
     * Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     */
    breakOnStoreFail?: pulumi.Input<boolean>;
    /**
     * Certificate Authentication Profile, empty if doesn't exist
     */
    certificateAuthenticationProfile?: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    identitySources?: pulumi.Input<pulumi.Input<inputs.identitymanagement.IdentitySourceSequenceIdentitySource>[]>;
    /**
     * The name of the identity source sequence
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentitySourceSequence resource.
 */
export interface IdentitySourceSequenceArgs {
    /**
     * Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
     */
    breakOnStoreFail: pulumi.Input<boolean>;
    /**
     * Certificate Authentication Profile, empty if doesn't exist
     */
    certificateAuthenticationProfile: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    identitySources: pulumi.Input<pulumi.Input<inputs.identitymanagement.IdentitySourceSequenceIdentitySource>[]>;
    /**
     * The name of the identity source sequence
     */
    name?: pulumi.Input<string>;
}

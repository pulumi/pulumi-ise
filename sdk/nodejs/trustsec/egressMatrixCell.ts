// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can manage a TrustSec Egress Matrix Cell.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.trustsec.EgressMatrixCell("example", {
 *     description: "EgressMatrixCell Description",
 *     defaultRule: "NONE",
 *     matrixCellStatus: "ENABLED",
 *     sgacls: ["26b76b10-66e6-11ee-9cc1-9eb2a3ecc82a, 9d64dcd0-6384-11ee-9cc1-9eb2a3ecc82a"],
 *     sourceSgtId: "93c66ed0-8c01-11e6-996c-525400b48521",
 *     destinationSgtId: "93e1bf00-8c01-11e6-996c-525400b48521",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:trustsec/egressMatrixCell:EgressMatrixCell example "76d24097-41c4-4558-a4d0-a8c07ac08470"
 * ```
 */
export class EgressMatrixCell extends pulumi.CustomResource {
    /**
     * Get an existing EgressMatrixCell resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EgressMatrixCellState, opts?: pulumi.CustomResourceOptions): EgressMatrixCell {
        return new EgressMatrixCell(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:trustsec/egressMatrixCell:EgressMatrixCell';

    /**
     * Returns true if the given object is an instance of EgressMatrixCell.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EgressMatrixCell {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EgressMatrixCell.__pulumiType;
    }

    /**
     * Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
     */
    public readonly defaultRule!: pulumi.Output<string>;
    /**
     * Description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Destination Trustsec Security Group ID
     */
    public readonly destinationSgtId!: pulumi.Output<string>;
    /**
     * Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
     */
    public readonly matrixCellStatus!: pulumi.Output<string>;
    /**
     * List of TrustSec Security Groups ACLs
     */
    public readonly sgacls!: pulumi.Output<string[] | undefined>;
    /**
     * Source Trustsec Security Group ID
     */
    public readonly sourceSgtId!: pulumi.Output<string>;

    /**
     * Create a EgressMatrixCell resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EgressMatrixCellArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EgressMatrixCellArgs | EgressMatrixCellState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EgressMatrixCellState | undefined;
            resourceInputs["defaultRule"] = state ? state.defaultRule : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["destinationSgtId"] = state ? state.destinationSgtId : undefined;
            resourceInputs["matrixCellStatus"] = state ? state.matrixCellStatus : undefined;
            resourceInputs["sgacls"] = state ? state.sgacls : undefined;
            resourceInputs["sourceSgtId"] = state ? state.sourceSgtId : undefined;
        } else {
            const args = argsOrState as EgressMatrixCellArgs | undefined;
            if ((!args || args.destinationSgtId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationSgtId'");
            }
            if ((!args || args.sourceSgtId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceSgtId'");
            }
            resourceInputs["defaultRule"] = args ? args.defaultRule : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["destinationSgtId"] = args ? args.destinationSgtId : undefined;
            resourceInputs["matrixCellStatus"] = args ? args.matrixCellStatus : undefined;
            resourceInputs["sgacls"] = args ? args.sgacls : undefined;
            resourceInputs["sourceSgtId"] = args ? args.sourceSgtId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EgressMatrixCell.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EgressMatrixCell resources.
 */
export interface EgressMatrixCellState {
    /**
     * Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
     */
    defaultRule?: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * Destination Trustsec Security Group ID
     */
    destinationSgtId?: pulumi.Input<string>;
    /**
     * Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
     */
    matrixCellStatus?: pulumi.Input<string>;
    /**
     * List of TrustSec Security Groups ACLs
     */
    sgacls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Source Trustsec Security Group ID
     */
    sourceSgtId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EgressMatrixCell resource.
 */
export interface EgressMatrixCellArgs {
    /**
     * Can be used only if sgacls not specified. - Choices: `NONE`, `DENY_IP`, `PERMIT_IP` - Default value: `NONE`
     */
    defaultRule?: pulumi.Input<string>;
    /**
     * Description
     */
    description?: pulumi.Input<string>;
    /**
     * Destination Trustsec Security Group ID
     */
    destinationSgtId: pulumi.Input<string>;
    /**
     * Matrix Cell Status - Choices: `DISABLED`, `ENABLED`, `MONITOR` - Default value: `DISABLED`
     */
    matrixCellStatus?: pulumi.Input<string>;
    /**
     * List of TrustSec Security Groups ACLs
     */
    sgacls?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Source Trustsec Security Group ID
     */
    sourceSgtId: pulumi.Input<string>;
}
// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource is used to update rank field in network access policy set. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and networkAccessPolicySet resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.networkaccess.PolicySetUpdateRank("example", {
 *     policySetId: "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
 *     rank: 0,
 * });
 * ```
 */
export class PolicySetUpdateRank extends pulumi.CustomResource {
    /**
     * Get an existing PolicySetUpdateRank resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicySetUpdateRankState, opts?: pulumi.CustomResourceOptions): PolicySetUpdateRank {
        return new PolicySetUpdateRank(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:networkaccess/policySetUpdateRank:PolicySetUpdateRank';

    /**
     * Returns true if the given object is an instance of PolicySetUpdateRank.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PolicySetUpdateRank {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PolicySetUpdateRank.__pulumiType;
    }

    /**
     * Policy set ID
     */
    public readonly policySetId!: pulumi.Output<string>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    public readonly rank!: pulumi.Output<number>;

    /**
     * Create a PolicySetUpdateRank resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicySetUpdateRankArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicySetUpdateRankArgs | PolicySetUpdateRankState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicySetUpdateRankState | undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["rank"] = state ? state.rank : undefined;
        } else {
            const args = argsOrState as PolicySetUpdateRankArgs | undefined;
            if ((!args || args.policySetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policySetId'");
            }
            if ((!args || args.rank === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rank'");
            }
            resourceInputs["policySetId"] = args ? args.policySetId : undefined;
            resourceInputs["rank"] = args ? args.rank : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PolicySetUpdateRank.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PolicySetUpdateRank resources.
 */
export interface PolicySetUpdateRankState {
    /**
     * Policy set ID
     */
    policySetId?: pulumi.Input<string>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a PolicySetUpdateRank resource.
 */
export interface PolicySetUpdateRankArgs {
    /**
     * Policy set ID
     */
    policySetId: pulumi.Input<string>;
    /**
     * The rank (priority) in relation to other rules. Lower rank is higher priority.
     */
    rank: pulumi.Input<number>;
}

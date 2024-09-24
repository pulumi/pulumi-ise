// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * This resource can manage a Repository.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ise from "@pulumi/ise";
 *
 * const example = new ise.system.Repository("example", {
 *     name: "repo1",
 *     protocol: "SFTP",
 *     path: "/dir",
 *     serverName: "server1",
 *     userName: "user9",
 *     password: "cisco123",
 *     enablePki: false,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import ise:system/repository:Repository example "repo1"
 * ```
 */
export class Repository extends pulumi.CustomResource {
    /**
     * Get an existing Repository resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RepositoryState, opts?: pulumi.CustomResourceOptions): Repository {
        return new Repository(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ise:system/repository:Repository';

    /**
     * Returns true if the given object is an instance of Repository.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Repository {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Repository.__pulumiType;
    }

    /**
     * Enable PKI
     */
    public readonly enablePki!: pulumi.Output<boolean | undefined>;
    /**
     * Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Password can contain alphanumeric and/or special characters.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    public readonly path!: pulumi.Output<string>;
    /**
     * Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * Name of the server
     */
    public readonly serverName!: pulumi.Output<string | undefined>;
    /**
     * User name
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a Repository resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RepositoryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RepositoryArgs | RepositoryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RepositoryState | undefined;
            resourceInputs["enablePki"] = state ? state.enablePki : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["serverName"] = state ? state.serverName : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as RepositoryArgs | undefined;
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["enablePki"] = args ? args.enablePki : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["serverName"] = args ? args.serverName : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Repository.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Repository resources.
 */
export interface RepositoryState {
    /**
     * Enable PKI
     */
    enablePki?: pulumi.Input<boolean>;
    /**
     * Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Password can contain alphanumeric and/or special characters.
     */
    password?: pulumi.Input<string>;
    /**
     * Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    path?: pulumi.Input<string>;
    /**
     * Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
     */
    protocol?: pulumi.Input<string>;
    /**
     * Name of the server
     */
    serverName?: pulumi.Input<string>;
    /**
     * User name
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Repository resource.
 */
export interface RepositoryArgs {
    /**
     * Enable PKI
     */
    enablePki?: pulumi.Input<boolean>;
    /**
     * Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    name?: pulumi.Input<string>;
    /**
     * Password can contain alphanumeric and/or special characters.
     */
    password?: pulumi.Input<string>;
    /**
     * Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
     */
    path: pulumi.Input<string>;
    /**
     * Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
     */
    protocol: pulumi.Input<string>;
    /**
     * Name of the server
     */
    serverName?: pulumi.Input<string>;
    /**
     * User name
     */
    userName?: pulumi.Input<string>;
}
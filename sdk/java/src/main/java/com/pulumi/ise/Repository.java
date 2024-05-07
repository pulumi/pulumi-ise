// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ise.RepositoryArgs;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.inputs.RepositoryState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a Repository.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ise.Repository;
 * import com.pulumi.ise.RepositoryArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new Repository(&#34;example&#34;, RepositoryArgs.builder()        
 *             .name(&#34;repo1&#34;)
 *             .protocol(&#34;SFTP&#34;)
 *             .path(&#34;/dir&#34;)
 *             .serverName(&#34;server1&#34;)
 *             .userName(&#34;user9&#34;)
 *             .password(&#34;cisco123&#34;)
 *             .enablePki(false)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ```sh
 * $ pulumi import ise:index/repository:Repository example &#34;repo1&#34;
 * ```
 * 
 */
@ResourceType(type="ise:index/repository:Repository")
public class Repository extends com.pulumi.resources.CustomResource {
    /**
     * Enable PKI
     * 
     */
    @Export(name="enablePki", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enablePki;

    /**
     * @return Enable PKI
     * 
     */
    public Output<Optional<Boolean>> enablePki() {
        return Codegen.optional(this.enablePki);
    }
    /**
     * Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Password can contain alphanumeric and/or special characters.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return Password can contain alphanumeric and/or special characters.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Path should always start with &#34;/&#34; and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output<String> path;

    /**
     * @return Path should always start with &#34;/&#34; and can contain alphanumeric, underscore, hyphen and dot characters.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * Name of the server
     * 
     */
    @Export(name="serverName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverName;

    /**
     * @return Name of the server
     * 
     */
    public Output<Optional<String>> serverName() {
        return Codegen.optional(this.serverName);
    }
    /**
     * User name
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userName;

    /**
     * @return User name
     * 
     */
    public Output<Optional<String>> userName() {
        return Codegen.optional(this.userName);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Repository(String name) {
        this(name, RepositoryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Repository(String name, RepositoryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Repository(String name, RepositoryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:index/repository:Repository", name, args == null ? RepositoryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Repository(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ise:index/repository:Repository", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Repository get(String name, Output<String> id, @Nullable RepositoryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Repository(name, id, state, options);
    }
}

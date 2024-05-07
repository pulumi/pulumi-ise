// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.TrustsecSecurityGroup;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclPlainArgs;
import com.pulumi.ise.TrustsecSecurityGroup.outputs.GetAclResult;
import com.pulumi.ise.Utilities;
import java.util.concurrent.CompletableFuture;

public final class TrustsecSecurityGroupFunctions {
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAclResult> getAcl() {
        return getAcl(GetAclArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAclResult> getAclPlain() {
        return getAclPlain(GetAclPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAclResult> getAcl(GetAclArgs args) {
        return getAcl(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAclResult> getAclPlain(GetAclPlainArgs args) {
        return getAclPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAclResult> getAcl(GetAclArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:TrustsecSecurityGroup/getAcl:getAcl", TypeShape.of(GetAclResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the TrustSec Security Group ACL.
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
     * import com.pulumi.ise.TrustsecSecurityGroup.TrustsecSecurityGroupFunctions;
     * import com.pulumi.ise.TrustsecSecurityGroup.inputs.GetAclArgs;
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
     *         final var example = TrustsecSecurityGroupFunctions.getAcl(GetAclArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAclResult> getAclPlain(GetAclPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:TrustsecSecurityGroup/getAcl:getAcl", TypeShape.of(GetAclResult.class), args, Utilities.withVersion(options));
    }
}

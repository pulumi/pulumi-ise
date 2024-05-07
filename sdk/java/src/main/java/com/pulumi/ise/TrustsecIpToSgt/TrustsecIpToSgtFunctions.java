// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.TrustsecIpToSgt;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingPlainArgs;
import com.pulumi.ise.TrustsecIpToSgt.outputs.GetMappingResult;
import com.pulumi.ise.Utilities;
import java.util.concurrent.CompletableFuture;

public final class TrustsecIpToSgtFunctions {
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetMappingResult> getMapping() {
        return getMapping(GetMappingArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetMappingResult> getMappingPlain() {
        return getMappingPlain(GetMappingPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetMappingResult> getMapping(GetMappingArgs args) {
        return getMapping(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetMappingResult> getMappingPlain(GetMappingPlainArgs args) {
        return getMappingPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetMappingResult> getMapping(GetMappingArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:TrustsecIpToSgt/getMapping:getMapping", TypeShape.of(GetMappingResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the TrustSec IP to SGT Mapping.
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
     * import com.pulumi.ise.TrustsecIpToSgt.TrustsecIpToSgtFunctions;
     * import com.pulumi.ise.TrustsecIpToSgt.inputs.GetMappingArgs;
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
     *         final var example = TrustsecIpToSgtFunctions.getMapping(GetMappingArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetMappingResult> getMappingPlain(GetMappingPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:TrustsecIpToSgt/getMapping:getMapping", TypeShape.of(GetMappingResult.class), args, Utilities.withVersion(options));
    }
}

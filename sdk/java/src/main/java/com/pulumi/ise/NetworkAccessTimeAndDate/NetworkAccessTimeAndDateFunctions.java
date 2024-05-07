// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.NetworkAccessTimeAndDate;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionPlainArgs;
import com.pulumi.ise.NetworkAccessTimeAndDate.outputs.GetConditionResult;
import com.pulumi.ise.Utilities;
import java.util.concurrent.CompletableFuture;

public final class NetworkAccessTimeAndDateFunctions {
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConditionResult> getCondition() {
        return getCondition(GetConditionArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetConditionResult> getConditionPlain() {
        return getConditionPlain(GetConditionPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConditionResult> getCondition(GetConditionArgs args) {
        return getCondition(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetConditionResult> getConditionPlain(GetConditionPlainArgs args) {
        return getConditionPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetConditionResult> getCondition(GetConditionArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:NetworkAccessTimeAndDate/getCondition:getCondition", TypeShape.of(GetConditionResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Network Access Time And Date Condition.
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
     * import com.pulumi.ise.NetworkAccessTimeAndDate.NetworkAccessTimeAndDateFunctions;
     * import com.pulumi.ise.NetworkAccessTimeAndDate.inputs.GetConditionArgs;
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
     *         final var example = NetworkAccessTimeAndDateFunctions.getCondition(GetConditionArgs.builder()
     *             .id(&#34;76d24097-41c4-4558-a4d0-a8c07ac08470&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetConditionResult> getConditionPlain(GetConditionPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:NetworkAccessTimeAndDate/getCondition:getCondition", TypeShape.of(GetConditionResult.class), args, Utilities.withVersion(options));
    }
}

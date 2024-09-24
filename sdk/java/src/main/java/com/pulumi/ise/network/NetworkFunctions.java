// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.network;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.network.inputs.GetDeviceArgs;
import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
import com.pulumi.ise.network.inputs.GetDeviceGroupPlainArgs;
import com.pulumi.ise.network.inputs.GetDevicePlainArgs;
import com.pulumi.ise.network.outputs.GetDeviceGroupResult;
import com.pulumi.ise.network.outputs.GetDeviceResult;
import java.util.concurrent.CompletableFuture;

public final class NetworkFunctions {
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceResult> getDevice() {
        return getDevice(GetDeviceArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceResult> getDevicePlain() {
        return getDevicePlain(GetDevicePlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceResult> getDevice(GetDeviceArgs args) {
        return getDevice(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceResult> getDevicePlain(GetDevicePlainArgs args) {
        return getDevicePlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceResult> getDevice(GetDeviceArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:network/getDevice:getDevice", TypeShape.of(GetDeviceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Network Device.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceArgs;
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
     *         final var example = NetworkFunctions.getDevice(GetDeviceArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceResult> getDevicePlain(GetDevicePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:network/getDevice:getDevice", TypeShape.of(GetDeviceResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceGroupResult> getDeviceGroup() {
        return getDeviceGroup(GetDeviceGroupArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceGroupResult> getDeviceGroupPlain() {
        return getDeviceGroupPlain(GetDeviceGroupPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceGroupResult> getDeviceGroup(GetDeviceGroupArgs args) {
        return getDeviceGroup(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceGroupResult> getDeviceGroupPlain(GetDeviceGroupPlainArgs args) {
        return getDeviceGroupPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetDeviceGroupResult> getDeviceGroup(GetDeviceGroupArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:network/getDeviceGroup:getDeviceGroup", TypeShape.of(GetDeviceGroupResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Network Device Group.
     * 
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.ise.network.NetworkFunctions;
     * import com.pulumi.ise.network.inputs.GetDeviceGroupArgs;
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
     *         final var example = NetworkFunctions.getDeviceGroup(GetDeviceGroupArgs.builder()
     *             .id("76d24097-41c4-4558-a4d0-a8c07ac08470")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetDeviceGroupResult> getDeviceGroupPlain(GetDeviceGroupPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:network/getDeviceGroup:getDeviceGroup", TypeShape.of(GetDeviceGroupResult.class), args, Utilities.withVersion(options));
    }
}
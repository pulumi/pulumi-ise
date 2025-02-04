// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise.system;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.ise.Utilities;
import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
import com.pulumi.ise.system.inputs.GetLicenseTierStatePlainArgs;
import com.pulumi.ise.system.inputs.GetRepositoryArgs;
import com.pulumi.ise.system.inputs.GetRepositoryPlainArgs;
import com.pulumi.ise.system.outputs.GetLicenseTierStateResult;
import com.pulumi.ise.system.outputs.GetRepositoryResult;
import java.util.concurrent.CompletableFuture;

public final class SystemFunctions {
    /**
     * This data source can read the License Tier State.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
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
     *         final var example = SystemFunctions.getLicenseTierState(GetLicenseTierStateArgs.builder()
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
    public static Output<GetLicenseTierStateResult> getLicenseTierState(GetLicenseTierStateArgs args) {
        return getLicenseTierState(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the License Tier State.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
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
     *         final var example = SystemFunctions.getLicenseTierState(GetLicenseTierStateArgs.builder()
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
    public static CompletableFuture<GetLicenseTierStateResult> getLicenseTierStatePlain(GetLicenseTierStatePlainArgs args) {
        return getLicenseTierStatePlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the License Tier State.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
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
     *         final var example = SystemFunctions.getLicenseTierState(GetLicenseTierStateArgs.builder()
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
    public static Output<GetLicenseTierStateResult> getLicenseTierState(GetLicenseTierStateArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:system/getLicenseTierState:getLicenseTierState", TypeShape.of(GetLicenseTierStateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the License Tier State.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
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
     *         final var example = SystemFunctions.getLicenseTierState(GetLicenseTierStateArgs.builder()
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
    public static Output<GetLicenseTierStateResult> getLicenseTierState(GetLicenseTierStateArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("ise:system/getLicenseTierState:getLicenseTierState", TypeShape.of(GetLicenseTierStateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the License Tier State.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetLicenseTierStateArgs;
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
     *         final var example = SystemFunctions.getLicenseTierState(GetLicenseTierStateArgs.builder()
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
    public static CompletableFuture<GetLicenseTierStateResult> getLicenseTierStatePlain(GetLicenseTierStatePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:system/getLicenseTierState:getLicenseTierState", TypeShape.of(GetLicenseTierStateResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRepositoryResult> getRepository() {
        return getRepository(GetRepositoryArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRepositoryResult> getRepositoryPlain() {
        return getRepositoryPlain(GetRepositoryPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRepositoryResult> getRepository(GetRepositoryArgs args) {
        return getRepository(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRepositoryResult> getRepositoryPlain(GetRepositoryPlainArgs args) {
        return getRepositoryPlain(args, InvokeOptions.Empty);
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRepositoryResult> getRepository(GetRepositoryArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("ise:system/getRepository:getRepository", TypeShape.of(GetRepositoryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetRepositoryResult> getRepository(GetRepositoryArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("ise:system/getRepository:getRepository", TypeShape.of(GetRepositoryResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This data source can read the Repository.
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
     * import com.pulumi.ise.system.SystemFunctions;
     * import com.pulumi.ise.system.inputs.GetRepositoryArgs;
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
     *         final var example = SystemFunctions.getRepository(GetRepositoryArgs.builder()
     *             .id("repo1")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetRepositoryResult> getRepositoryPlain(GetRepositoryPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("ise:system/getRepository:getRepository", TypeShape.of(GetRepositoryResult.class), args, Utilities.withVersion(options));
    }
}

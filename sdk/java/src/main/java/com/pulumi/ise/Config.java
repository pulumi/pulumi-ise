// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ise;

import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("ise");
/**
 * Allow insecure HTTPS client. This can also be set as the ISE_INSECURE environment variable. Defaults to `true`.
 * 
 */
    public Optional<Boolean> insecure() {
        return Codegen.booleanProp("insecure").config(config).get();
    }
/**
 * Password for the ISE instance. This can also be set as the ISE_PASSWORD environment variable.
 * 
 */
    public Optional<String> password() {
        return Codegen.stringProp("password").config(config).get();
    }
/**
 * Number of retries for REST API calls. This can also be set as the ISE_RETRIES environment variable. Defaults to `3`.
 * 
 */
    public Optional<Integer> retries() {
        return Codegen.integerProp("retries").config(config).get();
    }
/**
 * URL of the Cisco ISE instance. This can also be set as the ISE_URL environment variable.
 * 
 */
    public Optional<String> url() {
        return Codegen.stringProp("url").config(config).get();
    }
/**
 * Username for the ISE instance. This can also be set as the ISE_USERNAME environment variable.
 * 
 */
    public Optional<String> username() {
        return Codegen.stringProp("username").config(config).get();
    }
}

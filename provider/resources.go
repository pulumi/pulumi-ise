// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"fmt"
	"path/filepath"
	"strings"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/CiscoDevNet/terraform-provider-ise/ise"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-ise/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-ise/bridge-metadata.json
var bridgeMetadata []byte

const (
	mainPkg = "ise"

	mainMod = "index"
)

func makeDataSource(mod string, res string) tokens.ModuleMember {
	mod = strings.ToLower(mod)
	return tfbridge.MakeDataSource(mainPkg, mod, res)
}

func makeResource(mod string, res string) tokens.Type {
	mod = strings.ToLower(mod)
	return tfbridge.MakeResource(mainPkg, mod, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pf.ShimProvider(ise.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "ise",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Ise",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "pulumi",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL:     "https://raw.githubusercontent.com/pulumi/pulumi-ise/main/docs/ise.png",
		Description: "A Pulumi package for managing resources on a Cisco ISE (Identity Service Engine) instance.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords: []string{
			"pulumi",
			"ise",
			"category/network",
		},
		License:    "Apache-2.0",
		Homepage:   "https://github.com/pulumi/pulumi-ise",
		Repository: "https://github.com/pulumi/pulumi-ise",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:         "CiscoDevNet",
		UpstreamRepoPath:  "./upstream",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		TFProviderVersion: "0.2.1",
		Version:           version.Version,
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"ise_allowed_protocols_tacacs":           {Tok: makeResource("tacacs", "AllowedProtocols")},
			"ise_license_tier_state":                 {Tok: makeResource("system", "LicenseTierState")},
			"ise_repository":                         {Tok: makeResource("system", "Repository")},
			"ise_certificate_authentication_profile": {Tok: makeResource("identity", "CertificateAuthenticationProfile")},
			"ise_endpoint":                           {Tok: makeResource("identity", "Endpoint")},
			"ise_endpoint_identity_group":            {Tok: makeResource("identity", "EndpointIdentityGroup")},
			"ise_identity_source_sequence":           {Tok: makeResource("identity", "IdentitySourceSequence")},
			"ise_internal_user":                      {Tok: makeResource("identity", "InternalUser")},
			"ise_user_identity_group":                {Tok: makeResource("identity", "UserIdentityGroup")},
			"ise_allowed_protocols":                  {Tok: makeResource("network", "AllowedProtocols")},
			"ise_authorization_profile":              {Tok: makeResource("network", "AuthorizationProfile")},
			"ise_downloadable_acl":                   {Tok: makeResource("network", "DownloadableAcl")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"ise_allowed_protocols_tacacs":           {Tok: makeDataSource("tacacs", "getAllowedProtocols")},
			"ise_license_tier_state":                 {Tok: makeDataSource("system", "getLicenseTierState")},
			"ise_repository":                         {Tok: makeDataSource("system", "getRepository")},
			"ise_certificate_authentication_profile": {Tok: makeDataSource("identity", "getCertificateAuthenticationProfile")},
			"ise_endpoint":                           {Tok: makeDataSource("identity", "getEndpoint")},
			"ise_endpoint_identity_group":            {Tok: makeDataSource("identity", "getEndpointIdentityGroup")},
			"ise_identity_source_sequence":           {Tok: makeDataSource("identity", "getIdentitySourceSequence")},
			"ise_internal_user":                      {Tok: makeDataSource("identity", "getInternalUser")},
			"ise_user_identity_group":                {Tok: makeDataSource("identity", "getUserIdentityGroup")},
			"ise_allowed_protocols":                  {Tok: makeDataSource("network", "getAllowedProtocols")},
			"ise_authorization_profile":              {Tok: makeDataSource("network", "getAuthorizationProfile")},
			"ise_downloadable_acl":                   {Tok: makeDataSource("network", "getDownloadableAcl")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumi/ise",

			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{
				Requires: map[string]string{
					"pulumi": ">=3.0.0,<4.0.0",
				}}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", "ise"),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				"ise",
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "Pulumi",

			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumi",
		},
	}

	prov.MustComputeTokens(tks.KnownModules("ise_", mainMod, []string{
		"active_directory_",
		"network",
		"system_",
		"trustsec_",
		"device_admin_",
		"tacacs_",
	}, tks.MakeStandard(mainPkg)))
	prov.SetAutonaming(255, "-")

	return prov
}

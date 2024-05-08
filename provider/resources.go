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
	"unicode"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/CiscoDevNet/terraform-provider-ise/ise"
	"github.com/ettle/strcase"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"

	"github.com/pulumi/pulumi-ise/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-ise/bridge-metadata.json
var bridgeMetadata []byte

const (
	modDeviceAdmin        = "DeviceAdmin"
	modIdentityManagement = "IdentityManagement"
	modNetwork            = "Network"
	modNetworkAccess      = "NetworkAccess"
	modSystem             = "System"
	modTrustSec           = "TrustSec"
)

var moduleNames = map[string]string{
	"active_directory": modIdentityManagement,
	"certificate":      modIdentityManagement,
	"device_admin":     modDeviceAdmin,
	"endpoint":         modIdentityManagement,
	"identity":         modIdentityManagement,
	"network_access":   modNetworkAccess,
	"tacacs":           modDeviceAdmin,
	"trustsec":         modTrustSec,
}

var resourceModules = map[string]string{
	"allowed_protocols":        modNetworkAccess,
	"allowed_protocols_tacacs": modDeviceAdmin,
	"authorization_profile":    modNetworkAccess,
	"downloadable_acl":         modNetworkAccess,
	"internal_user":            modIdentityManagement,
	"license_tier_state":       modSystem,
	"network_device":           modNetwork,
	"network_device_group":     modNetwork,
	"repository":               modSystem,
	"user_identity_group":      modIdentityManagement,
}

var namespaceMap = map[string]string{}

func convertName(tfname string) (module string, name string) {
	tfNameItems := strings.Split(tfname, "_")
	contract.Assertf(len(tfNameItems) >= 2, "Invalid snake case name %s", tfname)
	contract.Assertf(tfNameItems[0] == "ise", "Invalid snake case name %s. Does not start with ise", tfname)
	fullName := strings.Join(tfNameItems[1:], "_")
	moduleTitle := ""
	if mod, ok := resourceModules[fullName]; ok {
		moduleTitle = mod
	} else {
		for mod, val := range moduleNames {
			if strings.HasPrefix(fullName, mod) {
				moduleTitle = val
				break
			}
		}
		if moduleTitle == "" {
			contract.Failf("could not determine module for %q", tfname)
		}
	}
	name = strcase.ToPascal(strings.Join(tfNameItems[1:], "_"))
	name = trimPrefixCaseInsensitive(name, moduleTitle)
	module = strings.ToLower(moduleTitle)
	namespaceMap[module] = moduleTitle

	contract.Assertf(!unicode.IsDigit(rune(module[0])), "Pulumi namespace must not start with a digit: %s", name)
	contract.Assertf(!unicode.IsDigit(rune(name[0])), "Pulumi name must not start with a digit: %s", name)
	return
}

func trimPrefixCaseInsensitive(s, prefix string) string {
	if strings.HasPrefix(strings.ToLower(s), strings.ToLower(prefix)) {
		return s[len(prefix):]
	}
	return s
}

func makeDataSource(ds string) tokens.ModuleMember {
	mod, name := convertName(ds)
	return tfbridge.MakeDataSource("ise", mod, "get"+name)
}

func makeResource(res string) tokens.Type {
	mod, name := convertName(res)
	return tfbridge.MakeResource("ise", mod, name)
}

func moduleComputeStrategy() tfbridge.Strategy {
	return tfbridge.Strategy{
		Resource: func(tfToken string, elem *tfbridge.ResourceInfo) error {
			elem.Tok = makeResource(tfToken)
			return nil
		},
		DataSource: func(tfToken string, elem *tfbridge.DataSourceInfo) error {
			elem.Tok = makeDataSource(tfToken)
			return nil
		},
	}
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
		DisplayName: "Cisco ISE",
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
			"ise_active_directory_add_groups": {
				Tok: makeResource("ise_active_directory_add_groups"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"groups": {
						CSharpName: "GroupsValue",
					},
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
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
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: namespaceMap,
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumi",
		},
	}

	prov.MustComputeTokens(moduleComputeStrategy())
	prov.SetAutonaming(255, "-")

	return prov
}

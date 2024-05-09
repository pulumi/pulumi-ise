// Copyright 2016-2024, Pulumi Corporation.
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
	"path"
	"strings"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/CiscoDevNet/terraform-provider-ise/ise"

	pf "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"

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
	"network":          modNetwork,
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
	"repository":               modSystem,
	"user_identity_group":      modIdentityManagement,
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:           pf.ShimProvider(ise.Provider()),
		Name:        "ise",
		DisplayName: "Cisco ISE",
		Publisher:   "pulumi",
		LogoURL:     "https://raw.githubusercontent.com/pulumi/pulumi-ise/main/docs/ise.png",
		Description: "A Pulumi package for managing resources on a Cisco ISE (Identity Service Engine) instance.",
		Keywords: []string{
			"pulumi",
			"ise",
			"category/network",
		},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/pulumi/pulumi-ise",
		Repository:        "https://github.com/pulumi/pulumi-ise",
		GitHubOrg:         "CiscoDevNet",
		UpstreamRepoPath:  "./upstream",
		MetadataInfo:      tfbridge.NewProviderMetadata(bridgeMetadata),
		TFProviderVersion: "0.2.1",
		Version:           version.Version,
		Config:            map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*tfbridge.ResourceInfo{
			"ise_active_directory_add_groups": {
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
			ImportBasePath: path.Join(
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
		},
		Java: &tfbridge.JavaInfo{
			BasePackage: "com.pulumi",
		},
	}

	cSharpNamespaces := map[string]string{}
	prov.MustComputeTokens(tokens.MappedModules("ise_",
		"index", moduleNames, noGet(func(module, name string) (string, error) {
			switch name {
			case "":
				name = module
			case "AllowedProtocols":
				module = modNetworkAccess
			case "AllowedProtocolsTacacs":
				module = modDeviceAdmin
			case "AuthorizationProfile":
				module = modNetworkAccess
			case "DownloadableAcl":
				module = modNetworkAccess
			case "InternalUser":
				module = modIdentityManagement
			case "LicenseTierState":
				module = modSystem
			case "NetworkDevice":
				module = modNetwork
			case "NetworkDeviceGroup":
				module = modNetwork
			case "Repository":
				module = modSystem
			case "UserIdentityGroup":
				module = modIdentityManagement
			}
			if module == "index" {
				return "", fmt.Errorf("Missing module mapping for %s", name)
			}

			cSharpNamespaces[strings.ToLower(module)] = module

			return string(tfbridge.MakeResource(prov.Name, strings.ToLower(module), name)), nil
		})))
	prov.CSharp.Namespaces = cSharpNamespaces
	prov.SetAutonaming(255, "-")

	return prov
}

func noGet(f tokens.Make) tokens.Make {
	return func(module, name string) (string, error) {
		if rest, ok := strings.CutPrefix(name, "get"); ok {
			tk, err := f(module, rest)
			if err != nil {
				return "", err
			}
			parts := strings.Split(tk, ":")
			parts[2] = "get" + parts[2]

			return strings.Join(parts, ":"), nil
		}
		return f(module, name)
	}
}

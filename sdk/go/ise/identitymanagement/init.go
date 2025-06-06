// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ise:identitymanagement/activeDirectoryAddGroups:ActiveDirectoryAddGroups":
		r = &ActiveDirectoryAddGroups{}
	case "ise:identitymanagement/activeDirectoryJoinDomainWithAllNodes:ActiveDirectoryJoinDomainWithAllNodes":
		r = &ActiveDirectoryJoinDomainWithAllNodes{}
	case "ise:identitymanagement/activeDirectoryJoinPoint:ActiveDirectoryJoinPoint":
		r = &ActiveDirectoryJoinPoint{}
	case "ise:identitymanagement/certificateAuthenticationProfile:CertificateAuthenticationProfile":
		r = &CertificateAuthenticationProfile{}
	case "ise:identitymanagement/endpoint:Endpoint":
		r = &Endpoint{}
	case "ise:identitymanagement/endpointCustomAttribute:EndpointCustomAttribute":
		r = &EndpointCustomAttribute{}
	case "ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup":
		r = &EndpointIdentityGroup{}
	case "ise:identitymanagement/identitySourceSequence:IdentitySourceSequence":
		r = &IdentitySourceSequence{}
	case "ise:identitymanagement/internalUser:InternalUser":
		r = &InternalUser{}
	case "ise:identitymanagement/userIdentityGroup:UserIdentityGroup":
		r = &UserIdentityGroup{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/activeDirectoryAddGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/activeDirectoryJoinDomainWithAllNodes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/activeDirectoryJoinPoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/certificateAuthenticationProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/endpoint",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/endpointCustomAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/endpointIdentityGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/identitySourceSequence",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/internalUser",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"identitymanagement/userIdentityGroup",
		&module{version},
	)
}

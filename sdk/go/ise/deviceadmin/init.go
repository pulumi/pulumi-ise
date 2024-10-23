// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmin

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
	case "ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs":
		r = &AllowedProtocolsTacacs{}
	case "ise:deviceadmin/authenticationRule:AuthenticationRule":
		r = &AuthenticationRule{}
	case "ise:deviceadmin/authenticationRuleUpdateRank:AuthenticationRuleUpdateRank":
		r = &AuthenticationRuleUpdateRank{}
	case "ise:deviceadmin/authorizationExceptionRule:AuthorizationExceptionRule":
		r = &AuthorizationExceptionRule{}
	case "ise:deviceadmin/authorizationExceptionRuleUpdateRank:AuthorizationExceptionRuleUpdateRank":
		r = &AuthorizationExceptionRuleUpdateRank{}
	case "ise:deviceadmin/authorizationGlobalExceptionRule:AuthorizationGlobalExceptionRule":
		r = &AuthorizationGlobalExceptionRule{}
	case "ise:deviceadmin/authorizationGlobalExceptionRuleUpdateRank:AuthorizationGlobalExceptionRuleUpdateRank":
		r = &AuthorizationGlobalExceptionRuleUpdateRank{}
	case "ise:deviceadmin/authorizationRule:AuthorizationRule":
		r = &AuthorizationRule{}
	case "ise:deviceadmin/authorizationRuleUpdateRank:AuthorizationRuleUpdateRank":
		r = &AuthorizationRuleUpdateRank{}
	case "ise:deviceadmin/condition:Condition":
		r = &Condition{}
	case "ise:deviceadmin/policySet:PolicySet":
		r = &PolicySet{}
	case "ise:deviceadmin/policySetUpdateRank:PolicySetUpdateRank":
		r = &PolicySetUpdateRank{}
	case "ise:deviceadmin/tacacsCommandSet:TacacsCommandSet":
		r = &TacacsCommandSet{}
	case "ise:deviceadmin/tacacsProfile:TacacsProfile":
		r = &TacacsProfile{}
	case "ise:deviceadmin/timeAndDateCondition:TimeAndDateCondition":
		r = &TimeAndDateCondition{}
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
		"deviceadmin/allowedProtocolsTacacs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authenticationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authenticationRuleUpdateRank",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationExceptionRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationExceptionRuleUpdateRank",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationGlobalExceptionRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationGlobalExceptionRuleUpdateRank",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/authorizationRuleUpdateRank",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/condition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/policySet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/policySetUpdateRank",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/tacacsCommandSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/tacacsProfile",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ise",
		"deviceadmin/timeAndDateCondition",
		&module{version},
	)
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read an authorization profiles policy element.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/networkaccess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkaccess.LookupAuthorizationProfile(ctx, &networkaccess.LookupAuthorizationProfileArgs{
//				Id: pulumi.StringRef("76d24097-41c4-4558-a4d0-a8c07ac08470"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAuthorizationProfile(ctx *pulumi.Context, args *LookupAuthorizationProfileArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationProfileResult
	err := ctx.Invoke("ise:networkaccess/getAuthorizationProfile:getAuthorizationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthorizationProfile.
type LookupAuthorizationProfileArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the authorization profile
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAuthorizationProfile.
type LookupAuthorizationProfileResult struct {
	// Access type
	AccessType string `pulumi:"accessType"`
	// ACL
	Acl string `pulumi:"acl"`
	// List of advanced attributes
	AdvancedAttributes []GetAuthorizationProfileAdvancedAttribute `pulumi:"advancedAttributes"`
	// Agentless Posture.
	AgentlessPosture bool `pulumi:"agentlessPosture"`
	// Airespace ACL
	AirespaceAcl string `pulumi:"airespaceAcl"`
	// Airespace IPv6 ACL
	AirespaceIpv6Acl string `pulumi:"airespaceIpv6Acl"`
	// ASA VPN
	AsaVpn string `pulumi:"asaVpn"`
	// Auto smart port
	AutoSmartPort string `pulumi:"autoSmartPort"`
	// AVC profile
	AvcProfile string `pulumi:"avcProfile"`
	// DACL name
	DaclName string `pulumi:"daclName"`
	// Description
	Description string `pulumi:"description"`
	// Easy wired session candidate
	EasywiredSessionCandidate bool `pulumi:"easywiredSessionCandidate"`
	// The id of the object
	Id string `pulumi:"id"`
	// Interface template
	InterfaceTemplate string `pulumi:"interfaceTemplate"`
	// IPv6 ACL
	Ipv6AclFilter string `pulumi:"ipv6AclFilter"`
	// IPv6 DACL name
	Ipv6DaclName string `pulumi:"ipv6DaclName"`
	// MacSec policy
	MacSecPolicy string `pulumi:"macSecPolicy"`
	// The name of the authorization profile
	Name string `pulumi:"name"`
	// NEAT
	Neat bool `pulumi:"neat"`
	// Value needs to be an existing Network Device Profile
	ProfileName string `pulumi:"profileName"`
	// Maintain Connectivity During Reauthentication
	ReauthenticationConnectivity string `pulumi:"reauthenticationConnectivity"`
	// Reauthentication timer
	ReauthenticationTimer int `pulumi:"reauthenticationTimer"`
	// Service template
	ServiceTemplate bool `pulumi:"serviceTemplate"`
	// Track movement
	TrackMovement bool `pulumi:"trackMovement"`
	// Unique identifier
	UniqueIdentifier string `pulumi:"uniqueIdentifier"`
	// Vlan name or ID
	VlanNameId string `pulumi:"vlanNameId"`
	// Vlan tag ID
	VlanTagId int `pulumi:"vlanTagId"`
	// Voice domain permission
	VoiceDomainPermission bool `pulumi:"voiceDomainPermission"`
	// Web authentication (local)
	WebAuth bool `pulumi:"webAuth"`
	// Web redirection ACL
	WebRedirectionAcl string `pulumi:"webRedirectionAcl"`
	// This attribute is mandatory when `webRedirectionType` value is `CentralizedWebAuth`. For all other `webRedirectionType` values the field must be ignored.
	WebRedirectionDisplayCertificatesRenewalMessages bool `pulumi:"webRedirectionDisplayCertificatesRenewalMessages"`
	// A portal that exist in the DB and fits the `webRedirectionType`
	WebRedirectionPortalName string `pulumi:"webRedirectionPortalName"`
	// IP, hostname or FQDN
	WebRedirectionStaticIpHostNameFqdn string `pulumi:"webRedirectionStaticIpHostNameFqdn"`
	// This type must fit the `webRedirectionPortalName`
	WebRedirectionType string `pulumi:"webRedirectionType"`
}

func LookupAuthorizationProfileOutput(ctx *pulumi.Context, args LookupAuthorizationProfileOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationProfileResult, error) {
			args := v.(LookupAuthorizationProfileArgs)
			r, err := LookupAuthorizationProfile(ctx, &args, opts...)
			var s LookupAuthorizationProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthorizationProfileResultOutput)
}

// A collection of arguments for invoking getAuthorizationProfile.
type LookupAuthorizationProfileOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the authorization profile
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAuthorizationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProfileArgs)(nil)).Elem()
}

// A collection of values returned by getAuthorizationProfile.
type LookupAuthorizationProfileResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationProfileResult)(nil)).Elem()
}

func (o LookupAuthorizationProfileResultOutput) ToLookupAuthorizationProfileResultOutput() LookupAuthorizationProfileResultOutput {
	return o
}

func (o LookupAuthorizationProfileResultOutput) ToLookupAuthorizationProfileResultOutputWithContext(ctx context.Context) LookupAuthorizationProfileResultOutput {
	return o
}

// Access type
func (o LookupAuthorizationProfileResultOutput) AccessType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AccessType }).(pulumi.StringOutput)
}

// ACL
func (o LookupAuthorizationProfileResultOutput) Acl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Acl }).(pulumi.StringOutput)
}

// List of advanced attributes
func (o LookupAuthorizationProfileResultOutput) AdvancedAttributes() GetAuthorizationProfileAdvancedAttributeArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) []GetAuthorizationProfileAdvancedAttribute {
		return v.AdvancedAttributes
	}).(GetAuthorizationProfileAdvancedAttributeArrayOutput)
}

// Agentless Posture.
func (o LookupAuthorizationProfileResultOutput) AgentlessPosture() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.AgentlessPosture }).(pulumi.BoolOutput)
}

// Airespace ACL
func (o LookupAuthorizationProfileResultOutput) AirespaceAcl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AirespaceAcl }).(pulumi.StringOutput)
}

// Airespace IPv6 ACL
func (o LookupAuthorizationProfileResultOutput) AirespaceIpv6Acl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AirespaceIpv6Acl }).(pulumi.StringOutput)
}

// ASA VPN
func (o LookupAuthorizationProfileResultOutput) AsaVpn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AsaVpn }).(pulumi.StringOutput)
}

// Auto smart port
func (o LookupAuthorizationProfileResultOutput) AutoSmartPort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AutoSmartPort }).(pulumi.StringOutput)
}

// AVC profile
func (o LookupAuthorizationProfileResultOutput) AvcProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.AvcProfile }).(pulumi.StringOutput)
}

// DACL name
func (o LookupAuthorizationProfileResultOutput) DaclName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.DaclName }).(pulumi.StringOutput)
}

// Description
func (o LookupAuthorizationProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// Easy wired session candidate
func (o LookupAuthorizationProfileResultOutput) EasywiredSessionCandidate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.EasywiredSessionCandidate }).(pulumi.BoolOutput)
}

// The id of the object
func (o LookupAuthorizationProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Interface template
func (o LookupAuthorizationProfileResultOutput) InterfaceTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.InterfaceTemplate }).(pulumi.StringOutput)
}

// IPv6 ACL
func (o LookupAuthorizationProfileResultOutput) Ipv6AclFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Ipv6AclFilter }).(pulumi.StringOutput)
}

// IPv6 DACL name
func (o LookupAuthorizationProfileResultOutput) Ipv6DaclName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Ipv6DaclName }).(pulumi.StringOutput)
}

// MacSec policy
func (o LookupAuthorizationProfileResultOutput) MacSecPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.MacSecPolicy }).(pulumi.StringOutput)
}

// The name of the authorization profile
func (o LookupAuthorizationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// NEAT
func (o LookupAuthorizationProfileResultOutput) Neat() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.Neat }).(pulumi.BoolOutput)
}

// Value needs to be an existing Network Device Profile
func (o LookupAuthorizationProfileResultOutput) ProfileName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.ProfileName }).(pulumi.StringOutput)
}

// Maintain Connectivity During Reauthentication
func (o LookupAuthorizationProfileResultOutput) ReauthenticationConnectivity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.ReauthenticationConnectivity }).(pulumi.StringOutput)
}

// Reauthentication timer
func (o LookupAuthorizationProfileResultOutput) ReauthenticationTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) int { return v.ReauthenticationTimer }).(pulumi.IntOutput)
}

// Service template
func (o LookupAuthorizationProfileResultOutput) ServiceTemplate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.ServiceTemplate }).(pulumi.BoolOutput)
}

// Track movement
func (o LookupAuthorizationProfileResultOutput) TrackMovement() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.TrackMovement }).(pulumi.BoolOutput)
}

// Unique identifier
func (o LookupAuthorizationProfileResultOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

// Vlan name or ID
func (o LookupAuthorizationProfileResultOutput) VlanNameId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.VlanNameId }).(pulumi.StringOutput)
}

// Vlan tag ID
func (o LookupAuthorizationProfileResultOutput) VlanTagId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) int { return v.VlanTagId }).(pulumi.IntOutput)
}

// Voice domain permission
func (o LookupAuthorizationProfileResultOutput) VoiceDomainPermission() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.VoiceDomainPermission }).(pulumi.BoolOutput)
}

// Web authentication (local)
func (o LookupAuthorizationProfileResultOutput) WebAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool { return v.WebAuth }).(pulumi.BoolOutput)
}

// Web redirection ACL
func (o LookupAuthorizationProfileResultOutput) WebRedirectionAcl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.WebRedirectionAcl }).(pulumi.StringOutput)
}

// This attribute is mandatory when `webRedirectionType` value is `CentralizedWebAuth`. For all other `webRedirectionType` values the field must be ignored.
func (o LookupAuthorizationProfileResultOutput) WebRedirectionDisplayCertificatesRenewalMessages() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) bool {
		return v.WebRedirectionDisplayCertificatesRenewalMessages
	}).(pulumi.BoolOutput)
}

// A portal that exist in the DB and fits the `webRedirectionType`
func (o LookupAuthorizationProfileResultOutput) WebRedirectionPortalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.WebRedirectionPortalName }).(pulumi.StringOutput)
}

// IP, hostname or FQDN
func (o LookupAuthorizationProfileResultOutput) WebRedirectionStaticIpHostNameFqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.WebRedirectionStaticIpHostNameFqdn }).(pulumi.StringOutput)
}

// This type must fit the `webRedirectionPortalName`
func (o LookupAuthorizationProfileResultOutput) WebRedirectionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationProfileResult) string { return v.WebRedirectionType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationProfileResultOutput{})
}
// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read an allowed protocols policy element.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/network"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := network.LookupAllowedProtocols(ctx, &network.LookupAllowedProtocolsArgs{
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
func LookupAllowedProtocols(ctx *pulumi.Context, args *LookupAllowedProtocolsArgs, opts ...pulumi.InvokeOption) (*LookupAllowedProtocolsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAllowedProtocolsResult
	err := ctx.Invoke("ise:network/getAllowedProtocols:getAllowedProtocols", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAllowedProtocols.
type LookupAllowedProtocolsArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the allowed protocols
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAllowedProtocols.
type LookupAllowedProtocolsResult struct {
	// Allow 5G. This field is only supported from ISE 3.2.
	Allow5g bool `pulumi:"allow5g"`
	// Allow CHAP
	AllowChap bool `pulumi:"allowChap"`
	// Allow EAP Fast
	AllowEapFast bool `pulumi:"allowEapFast"`
	// Allow EAP MD5
	AllowEapMd5 bool `pulumi:"allowEapMd5"`
	// Allow EAP TLS
	AllowEapTls bool `pulumi:"allowEapTls"`
	// Allow EAP TTLS
	AllowEapTtls bool `pulumi:"allowEapTtls"`
	// Allow LEAP
	AllowLeap bool `pulumi:"allowLeap"`
	// Allow MS CHAP v1
	AllowMsChapV1 bool `pulumi:"allowMsChapV1"`
	// Allow MS CHAP v2
	AllowMsChapV2 bool `pulumi:"allowMsChapV2"`
	// Allow PAP ASCII
	AllowPapAscii bool `pulumi:"allowPapAscii"`
	// Allow PEAP
	AllowPeap bool `pulumi:"allowPeap"`
	// Allow preferred EAP protocol
	AllowPreferredEapProtocol bool `pulumi:"allowPreferredEapProtocol"`
	// Allow TEAP
	AllowTeap bool `pulumi:"allowTeap"`
	// Allow weak ciphers for EAP
	AllowWeakCiphersForEap bool `pulumi:"allowWeakCiphersForEap"`
	// Description
	Description string `pulumi:"description"`
	// Accept client certificates. Is required only if `eapFastUsePacs` is `false`.
	EapFastAcceptClientCert bool `pulumi:"eapFastAcceptClientCert"`
	// Allow machine authentication. Is required only if `eapFastUsePacs` is `false`.
	EapFastAllowMachineAuthentication bool `pulumi:"eapFastAllowMachineAuthentication"`
	// Allow EAP GTC
	EapFastEapGtc bool `pulumi:"eapFastEapGtc"`
	// Allow EAP GTC password change. Is required only if `eapFastEapGtc` is `true`.
	EapFastEapGtcPwdChange bool `pulumi:"eapFastEapGtcPwdChange"`
	// EAP GTC password change retries. Is required only if `eapFastEapGtc` is `true`.
	EapFastEapGtcPwdChangeRetries int `pulumi:"eapFastEapGtcPwdChangeRetries"`
	// Allow EAP MS CHAP v2
	EapFastEapMsChapV2 bool `pulumi:"eapFastEapMsChapV2"`
	// Allow EAP MS CHAP v2 password change. Is required only if `eapFastEapMsChapV2` is `true`.
	EapFastEapMsChapV2PwdChange bool `pulumi:"eapFastEapMsChapV2PwdChange"`
	// EAP MS CHAP v2 password change retries. Is required only if `eapFastEapMsChapV2` is `true`.
	EapFastEapMsChapV2PwdChangeRetries int `pulumi:"eapFastEapMsChapV2PwdChangeRetries"`
	// Allow EAP TLS
	EapFastEapTls bool `pulumi:"eapFastEapTls"`
	// Allow EAP TLS authentication of expired certificates. Is required only if `eapFastEapTls` is `true`.
	EapFastEapTlsAuthOfExpiredCerts bool `pulumi:"eapFastEapTlsAuthOfExpiredCerts"`
	// Enable EAP chaining
	EapFastEnableEapChaining bool `pulumi:"eapFastEnableEapChaining"`
	// Allow anonymous provisioning. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsAllowAnonymousProvisioning bool `pulumi:"eapFastPacsAllowAnonymousProvisioning"`
	// Allow authenticated provisioning. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsAllowAuthenticatedProvisioning bool `pulumi:"eapFastPacsAllowAuthenticatedProvisioning"`
	// Accept client certification for provisioning. Is required only if `eapFastPacsAllowAuthenticatedProvisioning` is `true`.
	EapFastPacsAllowClientCert bool `pulumi:"eapFastPacsAllowClientCert"`
	// Allow machine authentication. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsAllowMachineAuthentication bool `pulumi:"eapFastPacsAllowMachineAuthentication"`
	// Authorization PAC TTL. Is required only if `eapFastPacsStatelessSessionResume` is `true`.
	EapFastPacsAuthorizationPacTtl int `pulumi:"eapFastPacsAuthorizationPacTtl"`
	// Authorization PAC TTL unit. Is required only if `eapFastPacsStatelessSessionResume` is `true`.
	EapFastPacsAuthorizationPacTtlUnit string `pulumi:"eapFastPacsAuthorizationPacTtlUnit"`
	// Machine PAC TTL. Is required only if `eapFastPacsAllowMachineAuthentication` is `true`.
	EapFastPacsMachinePacTtl int `pulumi:"eapFastPacsMachinePacTtl"`
	// Machine PAC TTL unit. Is required only if `eapFastPacsAllowMachineAuthentication` is `true`.
	EapFastPacsMachinePacTtlUnit string `pulumi:"eapFastPacsMachinePacTtlUnit"`
	// Server returns access accept after authenticated provisioning. Is required only if `eapFastPacsAllowAuthenticatedProvisioning` is `true`.
	EapFastPacsServerReturns bool `pulumi:"eapFastPacsServerReturns"`
	// Stateless session resume. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsStatelessSessionResume bool `pulumi:"eapFastPacsStatelessSessionResume"`
	// PACs tunnel PAC time to live. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsTunnelPacTtl int `pulumi:"eapFastPacsTunnelPacTtl"`
	// PACs tunnel PAC time to live unit. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsTunnelPacTtlUnit string `pulumi:"eapFastPacsTunnelPacTtlUnit"`
	// Use proactive pac update percentage. Is required only if `eapFastUsePacs` is `true`.
	EapFastPacsUseProactivePacUpdatePercentage int `pulumi:"eapFastPacsUseProactivePacUpdatePercentage"`
	// Use PACs
	EapFastUsePacs bool `pulumi:"eapFastUsePacs"`
	// Allow authentication of expired certificates
	EapTlsAllowAuthOfExpiredCerts bool `pulumi:"eapTlsAllowAuthOfExpiredCerts"`
	// Enable stateless session resume
	EapTlsEnableStatelessSessionResume bool `pulumi:"eapTlsEnableStatelessSessionResume"`
	// EAP TLS L-Bit
	EapTlsLBit bool `pulumi:"eapTlsLBit"`
	// Session ticket percentage. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
	EapTlsSessionTicketPercentage int `pulumi:"eapTlsSessionTicketPercentage"`
	// Session ticket TTL. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
	EapTlsSessionTicketTtl int `pulumi:"eapTlsSessionTicketTtl"`
	// Session ticket TTL unit. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
	EapTlsSessionTicketTtlUnit string `pulumi:"eapTlsSessionTicketTtlUnit"`
	// Allow CHAP
	EapTtlsChap bool `pulumi:"eapTtlsChap"`
	// Allow EAP MD5
	EapTtlsEapMd5 bool `pulumi:"eapTtlsEapMd5"`
	// Allow EAP MS CHAP v2
	EapTtlsEapMsChapV2 bool `pulumi:"eapTtlsEapMsChapV2"`
	// Allow EAP MS CHAP v2 password change. Is required only if `eapTtlsEapMsChapV2` is `true`.
	EapTtlsEapMsChapV2PwdChange bool `pulumi:"eapTtlsEapMsChapV2PwdChange"`
	// EAP MS CHAP v2 password change retries. Is required only if `eapTtlsEapMsChapV2` is `true`.
	EapTtlsEapMsChapV2PwdChangeRetries int `pulumi:"eapTtlsEapMsChapV2PwdChangeRetries"`
	// Allow MS CHAP v1
	EapTtlsMsChapV1 bool `pulumi:"eapTtlsMsChapV1"`
	// Allow MS CHAP v2
	EapTtlsMsChapV2 bool `pulumi:"eapTtlsMsChapV2"`
	// Allow PAP ASCII
	EapTtlsPapAscii bool `pulumi:"eapTtlsPapAscii"`
	// The id of the object
	Id string `pulumi:"id"`
	// The name of the allowed protocols
	Name string `pulumi:"name"`
	// Allow PEAP EAP GTC
	PeapAllowPeapEapGtc bool `pulumi:"peapAllowPeapEapGtc"`
	// Allow PEAP EAP GTC password change. Is required only if `allowPeapEapGtc` is `true`.
	PeapAllowPeapEapGtcPwdChange bool `pulumi:"peapAllowPeapEapGtcPwdChange"`
	// PEAP EAP GTC password change retries. Is required only if `allowPeapEapGtc` is `true`.
	PeapAllowPeapEapGtcPwdChangeRetries int `pulumi:"peapAllowPeapEapGtcPwdChangeRetries"`
	// Allow PEAP EAP MS CHAP v2
	PeapAllowPeapEapMsChapV2 bool `pulumi:"peapAllowPeapEapMsChapV2"`
	// Allow PEAP EAP MS CHAP v2 password change. Is required only if `allowPeapEapMsChapV2` is `true`.
	PeapAllowPeapEapMsChapV2PwdChange bool `pulumi:"peapAllowPeapEapMsChapV2PwdChange"`
	// Allow PEAP EAP MS CHAP v2 password change retries. Is required only if `allowPeapEapMsChapV2` is `true`.
	PeapAllowPeapEapMsChapV2PwdChangeRetries int `pulumi:"peapAllowPeapEapMsChapV2PwdChangeRetries"`
	// Allow PEAP EAP TLS
	PeapAllowPeapEapTls bool `pulumi:"peapAllowPeapEapTls"`
	// Allow PEAP EAP TLS authentication of expired certificates. Is required only if `peapAllowPeapEapTls` is `true`.
	PeapAllowPeapEapTlsAuthOfExpiredCerts bool `pulumi:"peapAllowPeapEapTlsAuthOfExpiredCerts"`
	// Allow PEAP v0
	PeapPeapV0 bool `pulumi:"peapPeapV0"`
	// Preferred EAP protocol
	PreferredEapProtocol string `pulumi:"preferredEapProtocol"`
	// Process host lookup
	ProcessHostLookup bool `pulumi:"processHostLookup"`
	// Require cryptobinding
	RequireCryptobinding bool `pulumi:"requireCryptobinding"`
	// Require message authentication
	RequireMessageAuth bool `pulumi:"requireMessageAuth"`
	// Allow downgrade to MSK
	TeapDowngradeMsk bool `pulumi:"teapDowngradeMsk"`
	// Accept client certificate during tunnel establishment
	TeapEapAcceptClientCertDuringTunnelEst bool `pulumi:"teapEapAcceptClientCertDuringTunnelEst"`
	// Allow EAP chaining
	TeapEapChaining bool `pulumi:"teapEapChaining"`
	// Allow EAP MS CHAP v2
	TeapEapMsChapV2 bool `pulumi:"teapEapMsChapV2"`
	// Allow EAP MS CHAP v2 password change. Is required only if `teapEapMsChapV2` is `true`.
	TeapEapMsChapV2PwdChange bool `pulumi:"teapEapMsChapV2PwdChange"`
	// EAP MS CHAP v2 password change retries. Is required only if `teapEapMsChapV2` is `true`.
	TeapEapMsChapV2PwdChangeRetries int `pulumi:"teapEapMsChapV2PwdChangeRetries"`
	// Allow EAP TLS
	TeapEapTls bool `pulumi:"teapEapTls"`
	// Allow EAP TLS authentication of expired certs. Is required only if `teapEapTls` is `true`.
	TeapEapTlsAuthOfExpiredCerts bool `pulumi:"teapEapTlsAuthOfExpiredCerts"`
	// Request basic password authentication
	TeapRequestBasicPwdAuth bool `pulumi:"teapRequestBasicPwdAuth"`
}

func LookupAllowedProtocolsOutput(ctx *pulumi.Context, args LookupAllowedProtocolsOutputArgs, opts ...pulumi.InvokeOption) LookupAllowedProtocolsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAllowedProtocolsResult, error) {
			args := v.(LookupAllowedProtocolsArgs)
			r, err := LookupAllowedProtocols(ctx, &args, opts...)
			var s LookupAllowedProtocolsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAllowedProtocolsResultOutput)
}

// A collection of arguments for invoking getAllowedProtocols.
type LookupAllowedProtocolsOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the allowed protocols
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAllowedProtocolsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAllowedProtocolsArgs)(nil)).Elem()
}

// A collection of values returned by getAllowedProtocols.
type LookupAllowedProtocolsResultOutput struct{ *pulumi.OutputState }

func (LookupAllowedProtocolsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAllowedProtocolsResult)(nil)).Elem()
}

func (o LookupAllowedProtocolsResultOutput) ToLookupAllowedProtocolsResultOutput() LookupAllowedProtocolsResultOutput {
	return o
}

func (o LookupAllowedProtocolsResultOutput) ToLookupAllowedProtocolsResultOutputWithContext(ctx context.Context) LookupAllowedProtocolsResultOutput {
	return o
}

// Allow 5G. This field is only supported from ISE 3.2.
func (o LookupAllowedProtocolsResultOutput) Allow5g() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.Allow5g }).(pulumi.BoolOutput)
}

// Allow CHAP
func (o LookupAllowedProtocolsResultOutput) AllowChap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowChap }).(pulumi.BoolOutput)
}

// Allow EAP Fast
func (o LookupAllowedProtocolsResultOutput) AllowEapFast() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowEapFast }).(pulumi.BoolOutput)
}

// Allow EAP MD5
func (o LookupAllowedProtocolsResultOutput) AllowEapMd5() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowEapMd5 }).(pulumi.BoolOutput)
}

// Allow EAP TLS
func (o LookupAllowedProtocolsResultOutput) AllowEapTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowEapTls }).(pulumi.BoolOutput)
}

// Allow EAP TTLS
func (o LookupAllowedProtocolsResultOutput) AllowEapTtls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowEapTtls }).(pulumi.BoolOutput)
}

// Allow LEAP
func (o LookupAllowedProtocolsResultOutput) AllowLeap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowLeap }).(pulumi.BoolOutput)
}

// Allow MS CHAP v1
func (o LookupAllowedProtocolsResultOutput) AllowMsChapV1() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowMsChapV1 }).(pulumi.BoolOutput)
}

// Allow MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) AllowMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowMsChapV2 }).(pulumi.BoolOutput)
}

// Allow PAP ASCII
func (o LookupAllowedProtocolsResultOutput) AllowPapAscii() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowPapAscii }).(pulumi.BoolOutput)
}

// Allow PEAP
func (o LookupAllowedProtocolsResultOutput) AllowPeap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowPeap }).(pulumi.BoolOutput)
}

// Allow preferred EAP protocol
func (o LookupAllowedProtocolsResultOutput) AllowPreferredEapProtocol() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowPreferredEapProtocol }).(pulumi.BoolOutput)
}

// Allow TEAP
func (o LookupAllowedProtocolsResultOutput) AllowTeap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowTeap }).(pulumi.BoolOutput)
}

// Allow weak ciphers for EAP
func (o LookupAllowedProtocolsResultOutput) AllowWeakCiphersForEap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.AllowWeakCiphersForEap }).(pulumi.BoolOutput)
}

// Description
func (o LookupAllowedProtocolsResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.Description }).(pulumi.StringOutput)
}

// Accept client certificates. Is required only if `eapFastUsePacs` is `false`.
func (o LookupAllowedProtocolsResultOutput) EapFastAcceptClientCert() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastAcceptClientCert }).(pulumi.BoolOutput)
}

// Allow machine authentication. Is required only if `eapFastUsePacs` is `false`.
func (o LookupAllowedProtocolsResultOutput) EapFastAllowMachineAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastAllowMachineAuthentication }).(pulumi.BoolOutput)
}

// Allow EAP GTC
func (o LookupAllowedProtocolsResultOutput) EapFastEapGtc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapGtc }).(pulumi.BoolOutput)
}

// Allow EAP GTC password change. Is required only if `eapFastEapGtc` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastEapGtcPwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapGtcPwdChange }).(pulumi.BoolOutput)
}

// EAP GTC password change retries. Is required only if `eapFastEapGtc` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastEapGtcPwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastEapGtcPwdChangeRetries }).(pulumi.IntOutput)
}

// Allow EAP MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) EapFastEapMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapMsChapV2 }).(pulumi.BoolOutput)
}

// Allow EAP MS CHAP v2 password change. Is required only if `eapFastEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastEapMsChapV2PwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapMsChapV2PwdChange }).(pulumi.BoolOutput)
}

// EAP MS CHAP v2 password change retries. Is required only if `eapFastEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastEapMsChapV2PwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastEapMsChapV2PwdChangeRetries }).(pulumi.IntOutput)
}

// Allow EAP TLS
func (o LookupAllowedProtocolsResultOutput) EapFastEapTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapTls }).(pulumi.BoolOutput)
}

// Allow EAP TLS authentication of expired certificates. Is required only if `eapFastEapTls` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastEapTlsAuthOfExpiredCerts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEapTlsAuthOfExpiredCerts }).(pulumi.BoolOutput)
}

// Enable EAP chaining
func (o LookupAllowedProtocolsResultOutput) EapFastEnableEapChaining() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastEnableEapChaining }).(pulumi.BoolOutput)
}

// Allow anonymous provisioning. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAllowAnonymousProvisioning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsAllowAnonymousProvisioning }).(pulumi.BoolOutput)
}

// Allow authenticated provisioning. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAllowAuthenticatedProvisioning() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsAllowAuthenticatedProvisioning }).(pulumi.BoolOutput)
}

// Accept client certification for provisioning. Is required only if `eapFastPacsAllowAuthenticatedProvisioning` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAllowClientCert() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsAllowClientCert }).(pulumi.BoolOutput)
}

// Allow machine authentication. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAllowMachineAuthentication() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsAllowMachineAuthentication }).(pulumi.BoolOutput)
}

// Authorization PAC TTL. Is required only if `eapFastPacsStatelessSessionResume` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAuthorizationPacTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastPacsAuthorizationPacTtl }).(pulumi.IntOutput)
}

// Authorization PAC TTL unit. Is required only if `eapFastPacsStatelessSessionResume` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsAuthorizationPacTtlUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.EapFastPacsAuthorizationPacTtlUnit }).(pulumi.StringOutput)
}

// Machine PAC TTL. Is required only if `eapFastPacsAllowMachineAuthentication` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsMachinePacTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastPacsMachinePacTtl }).(pulumi.IntOutput)
}

// Machine PAC TTL unit. Is required only if `eapFastPacsAllowMachineAuthentication` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsMachinePacTtlUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.EapFastPacsMachinePacTtlUnit }).(pulumi.StringOutput)
}

// Server returns access accept after authenticated provisioning. Is required only if `eapFastPacsAllowAuthenticatedProvisioning` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsServerReturns() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsServerReturns }).(pulumi.BoolOutput)
}

// Stateless session resume. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsStatelessSessionResume() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastPacsStatelessSessionResume }).(pulumi.BoolOutput)
}

// PACs tunnel PAC time to live. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsTunnelPacTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastPacsTunnelPacTtl }).(pulumi.IntOutput)
}

// PACs tunnel PAC time to live unit. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsTunnelPacTtlUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.EapFastPacsTunnelPacTtlUnit }).(pulumi.StringOutput)
}

// Use proactive pac update percentage. Is required only if `eapFastUsePacs` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapFastPacsUseProactivePacUpdatePercentage() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapFastPacsUseProactivePacUpdatePercentage }).(pulumi.IntOutput)
}

// Use PACs
func (o LookupAllowedProtocolsResultOutput) EapFastUsePacs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapFastUsePacs }).(pulumi.BoolOutput)
}

// Allow authentication of expired certificates
func (o LookupAllowedProtocolsResultOutput) EapTlsAllowAuthOfExpiredCerts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTlsAllowAuthOfExpiredCerts }).(pulumi.BoolOutput)
}

// Enable stateless session resume
func (o LookupAllowedProtocolsResultOutput) EapTlsEnableStatelessSessionResume() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTlsEnableStatelessSessionResume }).(pulumi.BoolOutput)
}

// EAP TLS L-Bit
func (o LookupAllowedProtocolsResultOutput) EapTlsLBit() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTlsLBit }).(pulumi.BoolOutput)
}

// Session ticket percentage. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapTlsSessionTicketPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapTlsSessionTicketPercentage }).(pulumi.IntOutput)
}

// Session ticket TTL. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapTlsSessionTicketTtl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapTlsSessionTicketTtl }).(pulumi.IntOutput)
}

// Session ticket TTL unit. Is required only if `eapTlsEnableStatelessSessionResume` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapTlsSessionTicketTtlUnit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.EapTlsSessionTicketTtlUnit }).(pulumi.StringOutput)
}

// Allow CHAP
func (o LookupAllowedProtocolsResultOutput) EapTtlsChap() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsChap }).(pulumi.BoolOutput)
}

// Allow EAP MD5
func (o LookupAllowedProtocolsResultOutput) EapTtlsEapMd5() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsEapMd5 }).(pulumi.BoolOutput)
}

// Allow EAP MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) EapTtlsEapMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsEapMsChapV2 }).(pulumi.BoolOutput)
}

// Allow EAP MS CHAP v2 password change. Is required only if `eapTtlsEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapTtlsEapMsChapV2PwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsEapMsChapV2PwdChange }).(pulumi.BoolOutput)
}

// EAP MS CHAP v2 password change retries. Is required only if `eapTtlsEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) EapTtlsEapMsChapV2PwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.EapTtlsEapMsChapV2PwdChangeRetries }).(pulumi.IntOutput)
}

// Allow MS CHAP v1
func (o LookupAllowedProtocolsResultOutput) EapTtlsMsChapV1() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsMsChapV1 }).(pulumi.BoolOutput)
}

// Allow MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) EapTtlsMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsMsChapV2 }).(pulumi.BoolOutput)
}

// Allow PAP ASCII
func (o LookupAllowedProtocolsResultOutput) EapTtlsPapAscii() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.EapTtlsPapAscii }).(pulumi.BoolOutput)
}

// The id of the object
func (o LookupAllowedProtocolsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the allowed protocols
func (o LookupAllowedProtocolsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.Name }).(pulumi.StringOutput)
}

// Allow PEAP EAP GTC
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapGtc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapGtc }).(pulumi.BoolOutput)
}

// Allow PEAP EAP GTC password change. Is required only if `allowPeapEapGtc` is `true`.
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapGtcPwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapGtcPwdChange }).(pulumi.BoolOutput)
}

// PEAP EAP GTC password change retries. Is required only if `allowPeapEapGtc` is `true`.
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapGtcPwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.PeapAllowPeapEapGtcPwdChangeRetries }).(pulumi.IntOutput)
}

// Allow PEAP EAP MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapMsChapV2 }).(pulumi.BoolOutput)
}

// Allow PEAP EAP MS CHAP v2 password change. Is required only if `allowPeapEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapMsChapV2PwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapMsChapV2PwdChange }).(pulumi.BoolOutput)
}

// Allow PEAP EAP MS CHAP v2 password change retries. Is required only if `allowPeapEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapMsChapV2PwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.PeapAllowPeapEapMsChapV2PwdChangeRetries }).(pulumi.IntOutput)
}

// Allow PEAP EAP TLS
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapTls }).(pulumi.BoolOutput)
}

// Allow PEAP EAP TLS authentication of expired certificates. Is required only if `peapAllowPeapEapTls` is `true`.
func (o LookupAllowedProtocolsResultOutput) PeapAllowPeapEapTlsAuthOfExpiredCerts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapAllowPeapEapTlsAuthOfExpiredCerts }).(pulumi.BoolOutput)
}

// Allow PEAP v0
func (o LookupAllowedProtocolsResultOutput) PeapPeapV0() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.PeapPeapV0 }).(pulumi.BoolOutput)
}

// Preferred EAP protocol
func (o LookupAllowedProtocolsResultOutput) PreferredEapProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) string { return v.PreferredEapProtocol }).(pulumi.StringOutput)
}

// Process host lookup
func (o LookupAllowedProtocolsResultOutput) ProcessHostLookup() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.ProcessHostLookup }).(pulumi.BoolOutput)
}

// Require cryptobinding
func (o LookupAllowedProtocolsResultOutput) RequireCryptobinding() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.RequireCryptobinding }).(pulumi.BoolOutput)
}

// Require message authentication
func (o LookupAllowedProtocolsResultOutput) RequireMessageAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.RequireMessageAuth }).(pulumi.BoolOutput)
}

// Allow downgrade to MSK
func (o LookupAllowedProtocolsResultOutput) TeapDowngradeMsk() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapDowngradeMsk }).(pulumi.BoolOutput)
}

// Accept client certificate during tunnel establishment
func (o LookupAllowedProtocolsResultOutput) TeapEapAcceptClientCertDuringTunnelEst() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapAcceptClientCertDuringTunnelEst }).(pulumi.BoolOutput)
}

// Allow EAP chaining
func (o LookupAllowedProtocolsResultOutput) TeapEapChaining() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapChaining }).(pulumi.BoolOutput)
}

// Allow EAP MS CHAP v2
func (o LookupAllowedProtocolsResultOutput) TeapEapMsChapV2() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapMsChapV2 }).(pulumi.BoolOutput)
}

// Allow EAP MS CHAP v2 password change. Is required only if `teapEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) TeapEapMsChapV2PwdChange() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapMsChapV2PwdChange }).(pulumi.BoolOutput)
}

// EAP MS CHAP v2 password change retries. Is required only if `teapEapMsChapV2` is `true`.
func (o LookupAllowedProtocolsResultOutput) TeapEapMsChapV2PwdChangeRetries() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) int { return v.TeapEapMsChapV2PwdChangeRetries }).(pulumi.IntOutput)
}

// Allow EAP TLS
func (o LookupAllowedProtocolsResultOutput) TeapEapTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapTls }).(pulumi.BoolOutput)
}

// Allow EAP TLS authentication of expired certs. Is required only if `teapEapTls` is `true`.
func (o LookupAllowedProtocolsResultOutput) TeapEapTlsAuthOfExpiredCerts() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapEapTlsAuthOfExpiredCerts }).(pulumi.BoolOutput)
}

// Request basic password authentication
func (o LookupAllowedProtocolsResultOutput) TeapRequestBasicPwdAuth() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAllowedProtocolsResult) bool { return v.TeapRequestBasicPwdAuth }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAllowedProtocolsResultOutput{})
}

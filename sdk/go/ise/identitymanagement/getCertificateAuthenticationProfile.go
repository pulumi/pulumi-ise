// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Certificate Authentication Profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/identitymanagement"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identitymanagement.LookupCertificateAuthenticationProfile(ctx, &identitymanagement.LookupCertificateAuthenticationProfileArgs{
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
func LookupCertificateAuthenticationProfile(ctx *pulumi.Context, args *LookupCertificateAuthenticationProfileArgs, opts ...pulumi.InvokeOption) (*LookupCertificateAuthenticationProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCertificateAuthenticationProfileResult
	err := ctx.Invoke("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificateAuthenticationProfile.
type LookupCertificateAuthenticationProfileArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the certificate profile
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCertificateAuthenticationProfile.
type LookupCertificateAuthenticationProfileResult struct {
	// Allow as username
	AllowedAsUserName bool `pulumi:"allowedAsUserName"`
	// Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in `usernameFrom`.
	CertificateAttributeName string `pulumi:"certificateAttributeName"`
	// Description
	Description string `pulumi:"description"`
	// Referred IDStore name for the Certificate Profile or `[not applicable]` in case no identity store is chosen
	ExternalIdentityStoreName string `pulumi:"externalIdentityStoreName"`
	// The id of the object
	Id string `pulumi:"id"`
	// Match mode of the Certificate Profile. Allowed values: NEVER, RESOLVE*IDENTITY*AMBIGUITY, BINARY_COMPARISON
	MatchMode string `pulumi:"matchMode"`
	// The name of the certificate profile
	Name string `pulumi:"name"`
	// The attribute in the certificate where the user name should be taken from. Allowed values: `CERTIFICATE` (for a specific attribute as defined in certificateAttributeName), `UPN` (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
	UsernameFrom string `pulumi:"usernameFrom"`
}

func LookupCertificateAuthenticationProfileOutput(ctx *pulumi.Context, args LookupCertificateAuthenticationProfileOutputArgs, opts ...pulumi.InvokeOption) LookupCertificateAuthenticationProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCertificateAuthenticationProfileResultOutput, error) {
			args := v.(LookupCertificateAuthenticationProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:identitymanagement/getCertificateAuthenticationProfile:getCertificateAuthenticationProfile", args, LookupCertificateAuthenticationProfileResultOutput{}, options).(LookupCertificateAuthenticationProfileResultOutput), nil
		}).(LookupCertificateAuthenticationProfileResultOutput)
}

// A collection of arguments for invoking getCertificateAuthenticationProfile.
type LookupCertificateAuthenticationProfileOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the certificate profile
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupCertificateAuthenticationProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthenticationProfileArgs)(nil)).Elem()
}

// A collection of values returned by getCertificateAuthenticationProfile.
type LookupCertificateAuthenticationProfileResultOutput struct{ *pulumi.OutputState }

func (LookupCertificateAuthenticationProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCertificateAuthenticationProfileResult)(nil)).Elem()
}

func (o LookupCertificateAuthenticationProfileResultOutput) ToLookupCertificateAuthenticationProfileResultOutput() LookupCertificateAuthenticationProfileResultOutput {
	return o
}

func (o LookupCertificateAuthenticationProfileResultOutput) ToLookupCertificateAuthenticationProfileResultOutputWithContext(ctx context.Context) LookupCertificateAuthenticationProfileResultOutput {
	return o
}

// Allow as username
func (o LookupCertificateAuthenticationProfileResultOutput) AllowedAsUserName() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) bool { return v.AllowedAsUserName }).(pulumi.BoolOutput)
}

// Attribute name of the Certificate Profile - used only when CERTIFICATE is chosen in `usernameFrom`.
func (o LookupCertificateAuthenticationProfileResultOutput) CertificateAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.CertificateAttributeName }).(pulumi.StringOutput)
}

// Description
func (o LookupCertificateAuthenticationProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// Referred IDStore name for the Certificate Profile or `[not applicable]` in case no identity store is chosen
func (o LookupCertificateAuthenticationProfileResultOutput) ExternalIdentityStoreName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.ExternalIdentityStoreName }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupCertificateAuthenticationProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Match mode of the Certificate Profile. Allowed values: NEVER, RESOLVE*IDENTITY*AMBIGUITY, BINARY_COMPARISON
func (o LookupCertificateAuthenticationProfileResultOutput) MatchMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.MatchMode }).(pulumi.StringOutput)
}

// The name of the certificate profile
func (o LookupCertificateAuthenticationProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// The attribute in the certificate where the user name should be taken from. Allowed values: `CERTIFICATE` (for a specific attribute as defined in certificateAttributeName), `UPN` (for using any Subject or Alternative Name Attributes in the Certificate - an option only in AD)
func (o LookupCertificateAuthenticationProfileResultOutput) UsernameFrom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCertificateAuthenticationProfileResult) string { return v.UsernameFrom }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCertificateAuthenticationProfileResultOutput{})
}

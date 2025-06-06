// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Internal User.
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
//			_, err := identitymanagement.LookupInternalUser(ctx, &identitymanagement.LookupInternalUserArgs{
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
func LookupInternalUser(ctx *pulumi.Context, args *LookupInternalUserArgs, opts ...pulumi.InvokeOption) (*LookupInternalUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupInternalUserResult
	err := ctx.Invoke("ise:identitymanagement/getInternalUser:getInternalUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInternalUser.
type LookupInternalUserArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the internal user
	Name *string `pulumi:"name"`
}

// A collection of values returned by getInternalUser.
type LookupInternalUserResult struct {
	// The Account Name Alias will be used to send email notifications about password expiration. This field is only supported from ISE 3.2.
	AccountNameAlias string `pulumi:"accountNameAlias"`
	// Requires the user to change the password
	ChangePassword bool `pulumi:"changePassword"`
	// Key value map
	CustomAttributes string `pulumi:"customAttributes"`
	// Description
	Description string `pulumi:"description"`
	// Email address
	Email string `pulumi:"email"`
	// This field is added in ISE 2.0 to support TACACS+
	EnablePassword string `pulumi:"enablePassword"`
	// Whether the user is enabled/disabled
	Enabled bool `pulumi:"enabled"`
	// First name of the internal user
	FirstName string `pulumi:"firstName"`
	// The id of the object
	Id string `pulumi:"id"`
	// Comma separated list of identity group IDs.
	IdentityGroups string `pulumi:"identityGroups"`
	// Last name of the internal user
	LastName string `pulumi:"lastName"`
	// The name of the internal user
	Name string `pulumi:"name"`
	// The password of the internal user
	Password string `pulumi:"password"`
	// The ID store where the internal user's password is kept
	PasswordIdStore string `pulumi:"passwordIdStore"`
	// Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This field is only supported from ISE 3.2.
	PasswordNeverExpires bool `pulumi:"passwordNeverExpires"`
}

func LookupInternalUserOutput(ctx *pulumi.Context, args LookupInternalUserOutputArgs, opts ...pulumi.InvokeOption) LookupInternalUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupInternalUserResultOutput, error) {
			args := v.(LookupInternalUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:identitymanagement/getInternalUser:getInternalUser", args, LookupInternalUserResultOutput{}, options).(LookupInternalUserResultOutput), nil
		}).(LookupInternalUserResultOutput)
}

// A collection of arguments for invoking getInternalUser.
type LookupInternalUserOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the internal user
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupInternalUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalUserArgs)(nil)).Elem()
}

// A collection of values returned by getInternalUser.
type LookupInternalUserResultOutput struct{ *pulumi.OutputState }

func (LookupInternalUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupInternalUserResult)(nil)).Elem()
}

func (o LookupInternalUserResultOutput) ToLookupInternalUserResultOutput() LookupInternalUserResultOutput {
	return o
}

func (o LookupInternalUserResultOutput) ToLookupInternalUserResultOutputWithContext(ctx context.Context) LookupInternalUserResultOutput {
	return o
}

// The Account Name Alias will be used to send email notifications about password expiration. This field is only supported from ISE 3.2.
func (o LookupInternalUserResultOutput) AccountNameAlias() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.AccountNameAlias }).(pulumi.StringOutput)
}

// Requires the user to change the password
func (o LookupInternalUserResultOutput) ChangePassword() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInternalUserResult) bool { return v.ChangePassword }).(pulumi.BoolOutput)
}

// Key value map
func (o LookupInternalUserResultOutput) CustomAttributes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.CustomAttributes }).(pulumi.StringOutput)
}

// Description
func (o LookupInternalUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// Email address
func (o LookupInternalUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// This field is added in ISE 2.0 to support TACACS+
func (o LookupInternalUserResultOutput) EnablePassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.EnablePassword }).(pulumi.StringOutput)
}

// Whether the user is enabled/disabled
func (o LookupInternalUserResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInternalUserResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// First name of the internal user
func (o LookupInternalUserResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupInternalUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Comma separated list of identity group IDs.
func (o LookupInternalUserResultOutput) IdentityGroups() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.IdentityGroups }).(pulumi.StringOutput)
}

// Last name of the internal user
func (o LookupInternalUserResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.LastName }).(pulumi.StringOutput)
}

// The name of the internal user
func (o LookupInternalUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// The password of the internal user
func (o LookupInternalUserResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.Password }).(pulumi.StringOutput)
}

// The ID store where the internal user's password is kept
func (o LookupInternalUserResultOutput) PasswordIdStore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupInternalUserResult) string { return v.PasswordIdStore }).(pulumi.StringOutput)
}

// Set to `true` to indicate the user password never expires. This will not apply to Users who are also ISE Admins. This field is only supported from ISE 3.2.
func (o LookupInternalUserResultOutput) PasswordNeverExpires() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupInternalUserResult) bool { return v.PasswordNeverExpires }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupInternalUserResultOutput{})
}

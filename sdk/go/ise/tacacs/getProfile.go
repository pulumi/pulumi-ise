// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tacacs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TACACS Profile.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/tacacs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tacacs.LookupProfile(ctx, &tacacs.LookupProfileArgs{
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
func LookupProfile(ctx *pulumi.Context, args *LookupProfileArgs, opts ...pulumi.InvokeOption) (*LookupProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupProfileResult
	err := ctx.Invoke("ise:tacacs/getProfile:getProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProfile.
type LookupProfileArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the TACACS profile
	Name *string `pulumi:"name"`
}

// A collection of values returned by getProfile.
type LookupProfileResult struct {
	// Description
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// The name of the TACACS profile
	Name              string                       `pulumi:"name"`
	SessionAttributes []GetProfileSessionAttribute `pulumi:"sessionAttributes"`
}

func LookupProfileOutput(ctx *pulumi.Context, args LookupProfileOutputArgs, opts ...pulumi.InvokeOption) LookupProfileResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupProfileResult, error) {
			args := v.(LookupProfileArgs)
			r, err := LookupProfile(ctx, &args, opts...)
			var s LookupProfileResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupProfileResultOutput)
}

// A collection of arguments for invoking getProfile.
type LookupProfileOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the TACACS profile
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileArgs)(nil)).Elem()
}

// A collection of values returned by getProfile.
type LookupProfileResultOutput struct{ *pulumi.OutputState }

func (LookupProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProfileResult)(nil)).Elem()
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutput() LookupProfileResultOutput {
	return o
}

func (o LookupProfileResultOutput) ToLookupProfileResultOutputWithContext(ctx context.Context) LookupProfileResultOutput {
	return o
}

// Description
func (o LookupProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the TACACS profile
func (o LookupProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupProfileResultOutput) SessionAttributes() GetProfileSessionAttributeArrayOutput {
	return o.ApplyT(func(v LookupProfileResult) []GetProfileSessionAttribute { return v.SessionAttributes }).(GetProfileSessionAttributeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupProfileResultOutput{})
}

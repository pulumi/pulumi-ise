// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the License Tier State.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/system"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.LookupLicenseTierState(ctx, &system.LookupLicenseTierStateArgs{
//				Id: "76d24097-41c4-4558-a4d0-a8c07ac08470",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLicenseTierState(ctx *pulumi.Context, args *LookupLicenseTierStateArgs, opts ...pulumi.InvokeOption) (*LookupLicenseTierStateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLicenseTierStateResult
	err := ctx.Invoke("ise:system/getLicenseTierState:getLicenseTierState", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLicenseTierState.
type LookupLicenseTierStateArgs struct {
	// The id of the object
	Id string `pulumi:"id"`
}

// A collection of values returned by getLicenseTierState.
type LookupLicenseTierStateResult struct {
	// The id of the object
	Id string `pulumi:"id"`
	// List of licenses
	Licenses []GetLicenseTierStateLicense `pulumi:"licenses"`
}

func LookupLicenseTierStateOutput(ctx *pulumi.Context, args LookupLicenseTierStateOutputArgs, opts ...pulumi.InvokeOption) LookupLicenseTierStateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLicenseTierStateResultOutput, error) {
			args := v.(LookupLicenseTierStateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:system/getLicenseTierState:getLicenseTierState", args, LookupLicenseTierStateResultOutput{}, options).(LookupLicenseTierStateResultOutput), nil
		}).(LookupLicenseTierStateResultOutput)
}

// A collection of arguments for invoking getLicenseTierState.
type LookupLicenseTierStateOutputArgs struct {
	// The id of the object
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupLicenseTierStateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseTierStateArgs)(nil)).Elem()
}

// A collection of values returned by getLicenseTierState.
type LookupLicenseTierStateResultOutput struct{ *pulumi.OutputState }

func (LookupLicenseTierStateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLicenseTierStateResult)(nil)).Elem()
}

func (o LookupLicenseTierStateResultOutput) ToLookupLicenseTierStateResultOutput() LookupLicenseTierStateResultOutput {
	return o
}

func (o LookupLicenseTierStateResultOutput) ToLookupLicenseTierStateResultOutputWithContext(ctx context.Context) LookupLicenseTierStateResultOutput {
	return o
}

// The id of the object
func (o LookupLicenseTierStateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLicenseTierStateResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of licenses
func (o LookupLicenseTierStateResultOutput) Licenses() GetLicenseTierStateLicenseArrayOutput {
	return o.ApplyT(func(v LookupLicenseTierStateResult) []GetLicenseTierStateLicense { return v.Licenses }).(GetLicenseTierStateLicenseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLicenseTierStateResultOutput{})
}

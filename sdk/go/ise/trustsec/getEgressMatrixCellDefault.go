// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TrustSec Egress Matrix Cell Default.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/trustsec"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := trustsec.LookupEgressMatrixCellDefault(ctx, &trustsec.LookupEgressMatrixCellDefaultArgs{
//				Id: "92c1a900-8c01-11e6-996c-525400b48521",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupEgressMatrixCellDefault(ctx *pulumi.Context, args *LookupEgressMatrixCellDefaultArgs, opts ...pulumi.InvokeOption) (*LookupEgressMatrixCellDefaultResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEgressMatrixCellDefaultResult
	err := ctx.Invoke("ise:trustsec/getEgressMatrixCellDefault:getEgressMatrixCellDefault", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEgressMatrixCellDefault.
type LookupEgressMatrixCellDefaultArgs struct {
	// The id of the object
	Id string `pulumi:"id"`
}

// A collection of values returned by getEgressMatrixCellDefault.
type LookupEgressMatrixCellDefaultResult struct {
	// Can be used only if sgacls not specified. Final Catch All Rule
	DefaultRule string `pulumi:"defaultRule"`
	// Description
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// Matrix Cell Status
	MatrixCellStatus string `pulumi:"matrixCellStatus"`
	// List of TrustSec Security Groups ACLs
	Sgacls []string `pulumi:"sgacls"`
}

func LookupEgressMatrixCellDefaultOutput(ctx *pulumi.Context, args LookupEgressMatrixCellDefaultOutputArgs, opts ...pulumi.InvokeOption) LookupEgressMatrixCellDefaultResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEgressMatrixCellDefaultResultOutput, error) {
			args := v.(LookupEgressMatrixCellDefaultArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:trustsec/getEgressMatrixCellDefault:getEgressMatrixCellDefault", args, LookupEgressMatrixCellDefaultResultOutput{}, options).(LookupEgressMatrixCellDefaultResultOutput), nil
		}).(LookupEgressMatrixCellDefaultResultOutput)
}

// A collection of arguments for invoking getEgressMatrixCellDefault.
type LookupEgressMatrixCellDefaultOutputArgs struct {
	// The id of the object
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEgressMatrixCellDefaultOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressMatrixCellDefaultArgs)(nil)).Elem()
}

// A collection of values returned by getEgressMatrixCellDefault.
type LookupEgressMatrixCellDefaultResultOutput struct{ *pulumi.OutputState }

func (LookupEgressMatrixCellDefaultResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressMatrixCellDefaultResult)(nil)).Elem()
}

func (o LookupEgressMatrixCellDefaultResultOutput) ToLookupEgressMatrixCellDefaultResultOutput() LookupEgressMatrixCellDefaultResultOutput {
	return o
}

func (o LookupEgressMatrixCellDefaultResultOutput) ToLookupEgressMatrixCellDefaultResultOutputWithContext(ctx context.Context) LookupEgressMatrixCellDefaultResultOutput {
	return o
}

// Can be used only if sgacls not specified. Final Catch All Rule
func (o LookupEgressMatrixCellDefaultResultOutput) DefaultRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellDefaultResult) string { return v.DefaultRule }).(pulumi.StringOutput)
}

// Description
func (o LookupEgressMatrixCellDefaultResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellDefaultResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupEgressMatrixCellDefaultResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellDefaultResult) string { return v.Id }).(pulumi.StringOutput)
}

// Matrix Cell Status
func (o LookupEgressMatrixCellDefaultResultOutput) MatrixCellStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellDefaultResult) string { return v.MatrixCellStatus }).(pulumi.StringOutput)
}

// List of TrustSec Security Groups ACLs
func (o LookupEgressMatrixCellDefaultResultOutput) Sgacls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellDefaultResult) []string { return v.Sgacls }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEgressMatrixCellDefaultResultOutput{})
}

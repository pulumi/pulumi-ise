// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TrustSec Egress Matrix Cell.
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
//			_, err := trustsec.LookupEgressMatrixCell(ctx, &trustsec.LookupEgressMatrixCellArgs{
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
func LookupEgressMatrixCell(ctx *pulumi.Context, args *LookupEgressMatrixCellArgs, opts ...pulumi.InvokeOption) (*LookupEgressMatrixCellResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEgressMatrixCellResult
	err := ctx.Invoke("ise:trustsec/getEgressMatrixCell:getEgressMatrixCell", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEgressMatrixCell.
type LookupEgressMatrixCellArgs struct {
	// The id of the object
	Id string `pulumi:"id"`
}

// A collection of values returned by getEgressMatrixCell.
type LookupEgressMatrixCellResult struct {
	// Can be used only if sgacls not specified.
	DefaultRule string `pulumi:"defaultRule"`
	// Description
	Description string `pulumi:"description"`
	// Destination Trustsec Security Group ID
	DestinationSgtId string `pulumi:"destinationSgtId"`
	// The id of the object
	Id string `pulumi:"id"`
	// Matrix Cell Status
	MatrixCellStatus string `pulumi:"matrixCellStatus"`
	// List of TrustSec Security Groups ACLs
	Sgacls []string `pulumi:"sgacls"`
	// Source Trustsec Security Group ID
	SourceSgtId string `pulumi:"sourceSgtId"`
}

func LookupEgressMatrixCellOutput(ctx *pulumi.Context, args LookupEgressMatrixCellOutputArgs, opts ...pulumi.InvokeOption) LookupEgressMatrixCellResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEgressMatrixCellResult, error) {
			args := v.(LookupEgressMatrixCellArgs)
			r, err := LookupEgressMatrixCell(ctx, &args, opts...)
			var s LookupEgressMatrixCellResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEgressMatrixCellResultOutput)
}

// A collection of arguments for invoking getEgressMatrixCell.
type LookupEgressMatrixCellOutputArgs struct {
	// The id of the object
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupEgressMatrixCellOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressMatrixCellArgs)(nil)).Elem()
}

// A collection of values returned by getEgressMatrixCell.
type LookupEgressMatrixCellResultOutput struct{ *pulumi.OutputState }

func (LookupEgressMatrixCellResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEgressMatrixCellResult)(nil)).Elem()
}

func (o LookupEgressMatrixCellResultOutput) ToLookupEgressMatrixCellResultOutput() LookupEgressMatrixCellResultOutput {
	return o
}

func (o LookupEgressMatrixCellResultOutput) ToLookupEgressMatrixCellResultOutputWithContext(ctx context.Context) LookupEgressMatrixCellResultOutput {
	return o
}

// Can be used only if sgacls not specified.
func (o LookupEgressMatrixCellResultOutput) DefaultRule() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.DefaultRule }).(pulumi.StringOutput)
}

// Description
func (o LookupEgressMatrixCellResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.Description }).(pulumi.StringOutput)
}

// Destination Trustsec Security Group ID
func (o LookupEgressMatrixCellResultOutput) DestinationSgtId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.DestinationSgtId }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupEgressMatrixCellResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.Id }).(pulumi.StringOutput)
}

// Matrix Cell Status
func (o LookupEgressMatrixCellResultOutput) MatrixCellStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.MatrixCellStatus }).(pulumi.StringOutput)
}

// List of TrustSec Security Groups ACLs
func (o LookupEgressMatrixCellResultOutput) Sgacls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) []string { return v.Sgacls }).(pulumi.StringArrayOutput)
}

// Source Trustsec Security Group ID
func (o LookupEgressMatrixCellResultOutput) SourceSgtId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEgressMatrixCellResult) string { return v.SourceSgtId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEgressMatrixCellResultOutput{})
}
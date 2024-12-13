// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the SXP Domain Filter.
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
//			_, err := trustsec.LookupSxpDomainFilter(ctx, &trustsec.LookupSxpDomainFilterArgs{
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
func LookupSxpDomainFilter(ctx *pulumi.Context, args *LookupSxpDomainFilterArgs, opts ...pulumi.InvokeOption) (*LookupSxpDomainFilterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSxpDomainFilterResult
	err := ctx.Invoke("ise:trustsec/getSxpDomainFilter:getSxpDomainFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSxpDomainFilter.
type LookupSxpDomainFilterArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Resource name
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSxpDomainFilter.
type LookupSxpDomainFilterResult struct {
	// Description
	Description string `pulumi:"description"`
	// List of SXP Domains, separated with comma
	Domains string `pulumi:"domains"`
	// The id of the object
	Id string `pulumi:"id"`
	// Resource name
	Name string `pulumi:"name"`
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt string `pulumi:"sgt"`
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet string `pulumi:"subnet"`
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn string `pulumi:"vn"`
}

func LookupSxpDomainFilterOutput(ctx *pulumi.Context, args LookupSxpDomainFilterOutputArgs, opts ...pulumi.InvokeOption) LookupSxpDomainFilterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSxpDomainFilterResultOutput, error) {
			args := v.(LookupSxpDomainFilterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:trustsec/getSxpDomainFilter:getSxpDomainFilter", args, LookupSxpDomainFilterResultOutput{}, options).(LookupSxpDomainFilterResultOutput), nil
		}).(LookupSxpDomainFilterResultOutput)
}

// A collection of arguments for invoking getSxpDomainFilter.
type LookupSxpDomainFilterOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Resource name
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSxpDomainFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSxpDomainFilterArgs)(nil)).Elem()
}

// A collection of values returned by getSxpDomainFilter.
type LookupSxpDomainFilterResultOutput struct{ *pulumi.OutputState }

func (LookupSxpDomainFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSxpDomainFilterResult)(nil)).Elem()
}

func (o LookupSxpDomainFilterResultOutput) ToLookupSxpDomainFilterResultOutput() LookupSxpDomainFilterResultOutput {
	return o
}

func (o LookupSxpDomainFilterResultOutput) ToLookupSxpDomainFilterResultOutputWithContext(ctx context.Context) LookupSxpDomainFilterResultOutput {
	return o
}

// Description
func (o LookupSxpDomainFilterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Description }).(pulumi.StringOutput)
}

// List of SXP Domains, separated with comma
func (o LookupSxpDomainFilterResultOutput) Domains() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Domains }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupSxpDomainFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource name
func (o LookupSxpDomainFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

// SGT name or ID. At least one of subnet or sgt or vn should be defined
func (o LookupSxpDomainFilterResultOutput) Sgt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Sgt }).(pulumi.StringOutput)
}

// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
func (o LookupSxpDomainFilterResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// Virtual Network. At least one of subnet or sgt or vn should be defined
func (o LookupSxpDomainFilterResultOutput) Vn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSxpDomainFilterResult) string { return v.Vn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSxpDomainFilterResultOutput{})
}

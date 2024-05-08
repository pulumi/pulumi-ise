// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TrustSec Security Group.
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
//			_, err := trustsec.LookupSecurityGroup(ctx, &trustsec.LookupSecurityGroupArgs{
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
func LookupSecurityGroup(ctx *pulumi.Context, args *LookupSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupResult
	err := ctx.Invoke("ise:trustsec/getSecurityGroup:getSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the security group
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResult struct {
	// Description
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// Read-only
	IsReadOnly bool `pulumi:"isReadOnly"`
	// The name of the security group
	Name string `pulumi:"name"`
	// Propagate to APIC (ACI)
	PropogateToApic bool `pulumi:"propogateToApic"`
	// `-1` to auto-generate
	Value int `pulumi:"value"`
}

func LookupSecurityGroupOutput(ctx *pulumi.Context, args LookupSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupResult, error) {
			args := v.(LookupSecurityGroupArgs)
			r, err := LookupSecurityGroup(ctx, &args, opts...)
			var s LookupSecurityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityGroupResultOutput)
}

// A collection of arguments for invoking getSecurityGroup.
type LookupSecurityGroupOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the security group
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityGroup.
type LookupSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupResult)(nil)).Elem()
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutput() LookupSecurityGroupResultOutput {
	return o
}

func (o LookupSecurityGroupResultOutput) ToLookupSecurityGroupResultOutputWithContext(ctx context.Context) LookupSecurityGroupResultOutput {
	return o
}

// Description
func (o LookupSecurityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// Read-only
func (o LookupSecurityGroupResultOutput) IsReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) bool { return v.IsReadOnly }).(pulumi.BoolOutput)
}

// The name of the security group
func (o LookupSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Propagate to APIC (ACI)
func (o LookupSecurityGroupResultOutput) PropogateToApic() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) bool { return v.PropogateToApic }).(pulumi.BoolOutput)
}

// `-1` to auto-generate
func (o LookupSecurityGroupResultOutput) Value() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecurityGroupResult) int { return v.Value }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupResultOutput{})
}
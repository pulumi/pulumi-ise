// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Endpoint Identity Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.LookupEndpointIdentityGroup(ctx, &identity.LookupEndpointIdentityGroupArgs{
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
func LookupEndpointIdentityGroup(ctx *pulumi.Context, args *LookupEndpointIdentityGroupArgs, opts ...pulumi.InvokeOption) (*LookupEndpointIdentityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointIdentityGroupResult
	err := ctx.Invoke("ise:identity/getEndpointIdentityGroup:getEndpointIdentityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpointIdentityGroup.
type LookupEndpointIdentityGroupArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the endpoint identity group
	Name *string `pulumi:"name"`
}

// A collection of values returned by getEndpointIdentityGroup.
type LookupEndpointIdentityGroupResult struct {
	// Description
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// The name of the endpoint identity group
	Name string `pulumi:"name"`
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId string `pulumi:"parentEndpointIdentityGroupId"`
	// System defined endpoint identity group
	SystemDefined bool `pulumi:"systemDefined"`
}

func LookupEndpointIdentityGroupOutput(ctx *pulumi.Context, args LookupEndpointIdentityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointIdentityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointIdentityGroupResult, error) {
			args := v.(LookupEndpointIdentityGroupArgs)
			r, err := LookupEndpointIdentityGroup(ctx, &args, opts...)
			var s LookupEndpointIdentityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointIdentityGroupResultOutput)
}

// A collection of arguments for invoking getEndpointIdentityGroup.
type LookupEndpointIdentityGroupOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the endpoint identity group
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupEndpointIdentityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointIdentityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getEndpointIdentityGroup.
type LookupEndpointIdentityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointIdentityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointIdentityGroupResult)(nil)).Elem()
}

func (o LookupEndpointIdentityGroupResultOutput) ToLookupEndpointIdentityGroupResultOutput() LookupEndpointIdentityGroupResultOutput {
	return o
}

func (o LookupEndpointIdentityGroupResultOutput) ToLookupEndpointIdentityGroupResultOutputWithContext(ctx context.Context) LookupEndpointIdentityGroupResultOutput {
	return o
}

// Description
func (o LookupEndpointIdentityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointIdentityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupEndpointIdentityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointIdentityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the endpoint identity group
func (o LookupEndpointIdentityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointIdentityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Parent endpoint identity group ID
func (o LookupEndpointIdentityGroupResultOutput) ParentEndpointIdentityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointIdentityGroupResult) string { return v.ParentEndpointIdentityGroupId }).(pulumi.StringOutput)
}

// System defined endpoint identity group
func (o LookupEndpointIdentityGroupResultOutput) SystemDefined() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointIdentityGroupResult) bool { return v.SystemDefined }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointIdentityGroupResultOutput{})
}

// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TrustSec Security Group ACL.
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
//			_, err := trustsec.LookupSecurityGroupAcl(ctx, &trustsec.LookupSecurityGroupAclArgs{
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
func LookupSecurityGroupAcl(ctx *pulumi.Context, args *LookupSecurityGroupAclArgs, opts ...pulumi.InvokeOption) (*LookupSecurityGroupAclResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSecurityGroupAclResult
	err := ctx.Invoke("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSecurityGroupAcl.
type LookupSecurityGroupAclArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the security group ACL
	Name *string `pulumi:"name"`
}

// A collection of values returned by getSecurityGroupAcl.
type LookupSecurityGroupAclResult struct {
	// Content of ACL
	AclContent string `pulumi:"aclContent"`
	// Description
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// IP Version
	IpVersion string `pulumi:"ipVersion"`
	// The name of the security group ACL
	Name string `pulumi:"name"`
	// Read-only
	ReadOnly bool `pulumi:"readOnly"`
}

func LookupSecurityGroupAclOutput(ctx *pulumi.Context, args LookupSecurityGroupAclOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityGroupAclResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSecurityGroupAclResultOutput, error) {
			args := v.(LookupSecurityGroupAclArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl", args, LookupSecurityGroupAclResultOutput{}, options).(LookupSecurityGroupAclResultOutput), nil
		}).(LookupSecurityGroupAclResultOutput)
}

// A collection of arguments for invoking getSecurityGroupAcl.
type LookupSecurityGroupAclOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the security group ACL
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupSecurityGroupAclOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupAclArgs)(nil)).Elem()
}

// A collection of values returned by getSecurityGroupAcl.
type LookupSecurityGroupAclResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityGroupAclResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityGroupAclResult)(nil)).Elem()
}

func (o LookupSecurityGroupAclResultOutput) ToLookupSecurityGroupAclResultOutput() LookupSecurityGroupAclResultOutput {
	return o
}

func (o LookupSecurityGroupAclResultOutput) ToLookupSecurityGroupAclResultOutputWithContext(ctx context.Context) LookupSecurityGroupAclResultOutput {
	return o
}

// Content of ACL
func (o LookupSecurityGroupAclResultOutput) AclContent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) string { return v.AclContent }).(pulumi.StringOutput)
}

// Description
func (o LookupSecurityGroupAclResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupSecurityGroupAclResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP Version
func (o LookupSecurityGroupAclResultOutput) IpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) string { return v.IpVersion }).(pulumi.StringOutput)
}

// The name of the security group ACL
func (o LookupSecurityGroupAclResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) string { return v.Name }).(pulumi.StringOutput)
}

// Read-only
func (o LookupSecurityGroupAclResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSecurityGroupAclResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityGroupAclResultOutput{})
}

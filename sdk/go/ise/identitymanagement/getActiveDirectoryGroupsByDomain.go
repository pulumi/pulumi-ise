// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Active Directory Groups By Domain.
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
//			_, err := identitymanagement.GetActiveDirectoryGroupsByDomain(ctx, &identitymanagement.GetActiveDirectoryGroupsByDomainArgs{
//				JoinPointId: "73808580-b6e6-11ee-8960-de6d7692bc40",
//				Domain:      "cisco.com",
//				Filter:      pulumi.StringRef("CN=ISE Admins"),
//				SidFilter:   pulumi.StringRef("cisco.com/S-1-5-33-544"),
//				TypeFilter:  pulumi.StringRef("UNIVERSAL"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetActiveDirectoryGroupsByDomain(ctx *pulumi.Context, args *GetActiveDirectoryGroupsByDomainArgs, opts ...pulumi.InvokeOption) (*GetActiveDirectoryGroupsByDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetActiveDirectoryGroupsByDomainResult
	err := ctx.Invoke("ise:identitymanagement/getActiveDirectoryGroupsByDomain:getActiveDirectoryGroupsByDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActiveDirectoryGroupsByDomain.
type GetActiveDirectoryGroupsByDomainArgs struct {
	// The domain whose groups we want to fetch
	Domain string `pulumi:"domain"`
	// Exact match filter on group's CN
	Filter *string `pulumi:"filter"`
	// Active Directory Join Point ID
	JoinPointId string `pulumi:"joinPointId"`
	// Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
	SidFilter *string `pulumi:"sidFilter"`
	// Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
	TypeFilter *string `pulumi:"typeFilter"`
}

// A collection of values returned by getActiveDirectoryGroupsByDomain.
type GetActiveDirectoryGroupsByDomainResult struct {
	// The domain whose groups we want to fetch
	Domain string `pulumi:"domain"`
	// Exact match filter on group's CN
	Filter *string `pulumi:"filter"`
	// List of groups
	Groups []GetActiveDirectoryGroupsByDomainGroup `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Active Directory Join Point ID
	JoinPointId string `pulumi:"joinPointId"`
	// Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
	SidFilter *string `pulumi:"sidFilter"`
	// Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
	TypeFilter *string `pulumi:"typeFilter"`
}

func GetActiveDirectoryGroupsByDomainOutput(ctx *pulumi.Context, args GetActiveDirectoryGroupsByDomainOutputArgs, opts ...pulumi.InvokeOption) GetActiveDirectoryGroupsByDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetActiveDirectoryGroupsByDomainResult, error) {
			args := v.(GetActiveDirectoryGroupsByDomainArgs)
			r, err := GetActiveDirectoryGroupsByDomain(ctx, &args, opts...)
			var s GetActiveDirectoryGroupsByDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetActiveDirectoryGroupsByDomainResultOutput)
}

// A collection of arguments for invoking getActiveDirectoryGroupsByDomain.
type GetActiveDirectoryGroupsByDomainOutputArgs struct {
	// The domain whose groups we want to fetch
	Domain pulumi.StringInput `pulumi:"domain"`
	// Exact match filter on group's CN
	Filter pulumi.StringPtrInput `pulumi:"filter"`
	// Active Directory Join Point ID
	JoinPointId pulumi.StringInput `pulumi:"joinPointId"`
	// Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
	SidFilter pulumi.StringPtrInput `pulumi:"sidFilter"`
	// Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
	TypeFilter pulumi.StringPtrInput `pulumi:"typeFilter"`
}

func (GetActiveDirectoryGroupsByDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveDirectoryGroupsByDomainArgs)(nil)).Elem()
}

// A collection of values returned by getActiveDirectoryGroupsByDomain.
type GetActiveDirectoryGroupsByDomainResultOutput struct{ *pulumi.OutputState }

func (GetActiveDirectoryGroupsByDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetActiveDirectoryGroupsByDomainResult)(nil)).Elem()
}

func (o GetActiveDirectoryGroupsByDomainResultOutput) ToGetActiveDirectoryGroupsByDomainResultOutput() GetActiveDirectoryGroupsByDomainResultOutput {
	return o
}

func (o GetActiveDirectoryGroupsByDomainResultOutput) ToGetActiveDirectoryGroupsByDomainResultOutputWithContext(ctx context.Context) GetActiveDirectoryGroupsByDomainResultOutput {
	return o
}

// The domain whose groups we want to fetch
func (o GetActiveDirectoryGroupsByDomainResultOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) string { return v.Domain }).(pulumi.StringOutput)
}

// Exact match filter on group's CN
func (o GetActiveDirectoryGroupsByDomainResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

// List of groups
func (o GetActiveDirectoryGroupsByDomainResultOutput) Groups() GetActiveDirectoryGroupsByDomainGroupArrayOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) []GetActiveDirectoryGroupsByDomainGroup {
		return v.Groups
	}).(GetActiveDirectoryGroupsByDomainGroupArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetActiveDirectoryGroupsByDomainResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) string { return v.Id }).(pulumi.StringOutput)
}

// Active Directory Join Point ID
func (o GetActiveDirectoryGroupsByDomainResultOutput) JoinPointId() pulumi.StringOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) string { return v.JoinPointId }).(pulumi.StringOutput)
}

// Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
func (o GetActiveDirectoryGroupsByDomainResultOutput) SidFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) *string { return v.SidFilter }).(pulumi.StringPtrOutput)
}

// Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
func (o GetActiveDirectoryGroupsByDomainResultOutput) TypeFilter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetActiveDirectoryGroupsByDomainResult) *string { return v.TypeFilter }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetActiveDirectoryGroupsByDomainResultOutput{})
}
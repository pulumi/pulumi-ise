// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Endpoint.
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
//			_, err := identitymanagement.LookupEndpoint(ctx, &identitymanagement.LookupEndpointArgs{
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
func LookupEndpoint(ctx *pulumi.Context, args *LookupEndpointArgs, opts ...pulumi.InvokeOption) (*LookupEndpointResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointResult
	err := ctx.Invoke("ise:identitymanagement/getEndpoint:getEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEndpoint.
type LookupEndpointArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the endpoint
	Name *string `pulumi:"name"`
}

// A collection of values returned by getEndpoint.
type LookupEndpointResult struct {
	// Custom Attributes
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// Description
	Description string `pulumi:"description"`
	// Identity Group ID
	GroupId string `pulumi:"groupId"`
	// The id of the object
	Id string `pulumi:"id"`
	// Identity Store
	IdentityStore string `pulumi:"identityStore"`
	// Identity Store Id
	IdentityStoreId string `pulumi:"identityStoreId"`
	// MAC address of the endpoint
	Mac string `pulumi:"mac"`
	// Mdm Compliance Status
	MdmComplianceStatus bool `pulumi:"mdmComplianceStatus"`
	// Mdm Encrypted
	MdmEncrypted bool `pulumi:"mdmEncrypted"`
	// Mdm Enrolled
	MdmEnrolled bool `pulumi:"mdmEnrolled"`
	// Mdm IMEI
	MdmImei string `pulumi:"mdmImei"`
	// Mdm JailBroken
	MdmJailBroken bool `pulumi:"mdmJailBroken"`
	// Mdm Manufacturer
	MdmManufacturer string `pulumi:"mdmManufacturer"`
	// Mdm Model
	MdmModel string `pulumi:"mdmModel"`
	// Mdm OS
	MdmOs string `pulumi:"mdmOs"`
	// Mdm PhoneNumber
	MdmPhoneNumber string `pulumi:"mdmPhoneNumber"`
	// Mdm Pinlock
	MdmPinlock bool `pulumi:"mdmPinlock"`
	// Mdm Reachable
	MdmReachable bool `pulumi:"mdmReachable"`
	// Mdm Serial
	MdmSerial string `pulumi:"mdmSerial"`
	// Mdm Server Name
	MdmServerName string `pulumi:"mdmServerName"`
	// The name of the endpoint
	Name string `pulumi:"name"`
	// Portal User
	PortalUser string `pulumi:"portalUser"`
	// Profile ID
	ProfileId string `pulumi:"profileId"`
	// Static Group Assignment
	StaticGroupAssignment bool `pulumi:"staticGroupAssignment"`
	// staticGroupAssignmentDefined
	StaticGroupAssignmentDefined bool `pulumi:"staticGroupAssignmentDefined"`
	// Static Profile Assignment
	StaticProfileAssignment bool `pulumi:"staticProfileAssignment"`
	// Static Profile Assignment Defined
	StaticProfileAssignmentDefined bool `pulumi:"staticProfileAssignmentDefined"`
}

func LookupEndpointOutput(ctx *pulumi.Context, args LookupEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointResultOutput, error) {
			args := v.(LookupEndpointArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv LookupEndpointResult
			secret, err := ctx.InvokePackageRaw("ise:identitymanagement/getEndpoint:getEndpoint", args, &rv, "", opts...)
			if err != nil {
				return LookupEndpointResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(LookupEndpointResultOutput)
			if secret {
				return pulumi.ToSecret(output).(LookupEndpointResultOutput), nil
			}
			return output, nil
		}).(LookupEndpointResultOutput)
}

// A collection of arguments for invoking getEndpoint.
type LookupEndpointOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the endpoint
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointArgs)(nil)).Elem()
}

// A collection of values returned by getEndpoint.
type LookupEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointResult)(nil)).Elem()
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutput() LookupEndpointResultOutput {
	return o
}

func (o LookupEndpointResultOutput) ToLookupEndpointResultOutputWithContext(ctx context.Context) LookupEndpointResultOutput {
	return o
}

// Custom Attributes
func (o LookupEndpointResultOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointResult) map[string]string { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// Description
func (o LookupEndpointResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Description }).(pulumi.StringOutput)
}

// Identity Group ID
func (o LookupEndpointResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity Store
func (o LookupEndpointResultOutput) IdentityStore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.IdentityStore }).(pulumi.StringOutput)
}

// Identity Store Id
func (o LookupEndpointResultOutput) IdentityStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.IdentityStoreId }).(pulumi.StringOutput)
}

// MAC address of the endpoint
func (o LookupEndpointResultOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Mac }).(pulumi.StringOutput)
}

// Mdm Compliance Status
func (o LookupEndpointResultOutput) MdmComplianceStatus() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmComplianceStatus }).(pulumi.BoolOutput)
}

// Mdm Encrypted
func (o LookupEndpointResultOutput) MdmEncrypted() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmEncrypted }).(pulumi.BoolOutput)
}

// Mdm Enrolled
func (o LookupEndpointResultOutput) MdmEnrolled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmEnrolled }).(pulumi.BoolOutput)
}

// Mdm IMEI
func (o LookupEndpointResultOutput) MdmImei() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmImei }).(pulumi.StringOutput)
}

// Mdm JailBroken
func (o LookupEndpointResultOutput) MdmJailBroken() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmJailBroken }).(pulumi.BoolOutput)
}

// Mdm Manufacturer
func (o LookupEndpointResultOutput) MdmManufacturer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmManufacturer }).(pulumi.StringOutput)
}

// Mdm Model
func (o LookupEndpointResultOutput) MdmModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmModel }).(pulumi.StringOutput)
}

// Mdm OS
func (o LookupEndpointResultOutput) MdmOs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmOs }).(pulumi.StringOutput)
}

// Mdm PhoneNumber
func (o LookupEndpointResultOutput) MdmPhoneNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmPhoneNumber }).(pulumi.StringOutput)
}

// Mdm Pinlock
func (o LookupEndpointResultOutput) MdmPinlock() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmPinlock }).(pulumi.BoolOutput)
}

// Mdm Reachable
func (o LookupEndpointResultOutput) MdmReachable() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.MdmReachable }).(pulumi.BoolOutput)
}

// Mdm Serial
func (o LookupEndpointResultOutput) MdmSerial() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmSerial }).(pulumi.StringOutput)
}

// Mdm Server Name
func (o LookupEndpointResultOutput) MdmServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.MdmServerName }).(pulumi.StringOutput)
}

// The name of the endpoint
func (o LookupEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// Portal User
func (o LookupEndpointResultOutput) PortalUser() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.PortalUser }).(pulumi.StringOutput)
}

// Profile ID
func (o LookupEndpointResultOutput) ProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointResult) string { return v.ProfileId }).(pulumi.StringOutput)
}

// Static Group Assignment
func (o LookupEndpointResultOutput) StaticGroupAssignment() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.StaticGroupAssignment }).(pulumi.BoolOutput)
}

// staticGroupAssignmentDefined
func (o LookupEndpointResultOutput) StaticGroupAssignmentDefined() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.StaticGroupAssignmentDefined }).(pulumi.BoolOutput)
}

// Static Profile Assignment
func (o LookupEndpointResultOutput) StaticProfileAssignment() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.StaticProfileAssignment }).(pulumi.BoolOutput)
}

// Static Profile Assignment Defined
func (o LookupEndpointResultOutput) StaticProfileAssignmentDefined() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEndpointResult) bool { return v.StaticProfileAssignmentDefined }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointResultOutput{})
}

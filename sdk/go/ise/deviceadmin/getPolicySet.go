// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmin

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Device Admin Policy Set.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/deviceadmin"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := deviceadmin.LookupPolicySet(ctx, &deviceadmin.LookupPolicySetArgs{
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
func LookupPolicySet(ctx *pulumi.Context, args *LookupPolicySetArgs, opts ...pulumi.InvokeOption) (*LookupPolicySetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPolicySetResult
	err := ctx.Invoke("ise:deviceadmin/getPolicySet:getPolicySet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPolicySet.
type LookupPolicySetArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
}

// A collection of values returned by getPolicySet.
type LookupPolicySetResult struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []GetPolicySetChildren `pulumi:"childrens"`
	// Dictionary attribute name
	ConditionAttributeName string `pulumi:"conditionAttributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue string `pulumi:"conditionAttributeValue"`
	// Dictionary name
	ConditionDictionaryName string `pulumi:"conditionDictionaryName"`
	// Dictionary value
	ConditionDictionaryValue string `pulumi:"conditionDictionaryValue"`
	// UUID for condition
	ConditionId string `pulumi:"conditionId"`
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate bool `pulumi:"conditionIsNegate"`
	// Equality operator
	ConditionOperator string `pulumi:"conditionOperator"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
	ConditionType string `pulumi:"conditionType"`
	// Indicates if this policy set is the default one
	Default bool `pulumi:"default"`
	// The description of the policy set
	Description string `pulumi:"description"`
	// The id of the object
	Id string `pulumi:"id"`
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy bool `pulumi:"isProxy"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name string `pulumi:"name"`
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank int `pulumi:"rank"`
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName string `pulumi:"serviceName"`
	// The state that the policy set is in. A disabled policy set cannot be matched.
	State string `pulumi:"state"`
}

func LookupPolicySetOutput(ctx *pulumi.Context, args LookupPolicySetOutputArgs, opts ...pulumi.InvokeOption) LookupPolicySetResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupPolicySetResultOutput, error) {
			args := v.(LookupPolicySetArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:deviceadmin/getPolicySet:getPolicySet", args, LookupPolicySetResultOutput{}, options).(LookupPolicySetResultOutput), nil
		}).(LookupPolicySetResultOutput)
}

// A collection of arguments for invoking getPolicySet.
type LookupPolicySetOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupPolicySetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetArgs)(nil)).Elem()
}

// A collection of values returned by getPolicySet.
type LookupPolicySetResultOutput struct{ *pulumi.OutputState }

func (LookupPolicySetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPolicySetResult)(nil)).Elem()
}

func (o LookupPolicySetResultOutput) ToLookupPolicySetResultOutput() LookupPolicySetResultOutput {
	return o
}

func (o LookupPolicySetResultOutput) ToLookupPolicySetResultOutputWithContext(ctx context.Context) LookupPolicySetResultOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o LookupPolicySetResultOutput) Childrens() GetPolicySetChildrenArrayOutput {
	return o.ApplyT(func(v LookupPolicySetResult) []GetPolicySetChildren { return v.Childrens }).(GetPolicySetChildrenArrayOutput)
}

// Dictionary attribute name
func (o LookupPolicySetResultOutput) ConditionAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionAttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o LookupPolicySetResultOutput) ConditionAttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionAttributeValue }).(pulumi.StringOutput)
}

// Dictionary name
func (o LookupPolicySetResultOutput) ConditionDictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionDictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o LookupPolicySetResultOutput) ConditionDictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionDictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o LookupPolicySetResultOutput) ConditionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionId }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupPolicySetResultOutput) ConditionIsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicySetResult) bool { return v.ConditionIsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o LookupPolicySetResultOutput) ConditionOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionOperator }).(pulumi.StringOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
func (o LookupPolicySetResultOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Indicates if this policy set is the default one
func (o LookupPolicySetResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicySetResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// The description of the policy set
func (o LookupPolicySetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupPolicySetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
func (o LookupPolicySetResultOutput) IsProxy() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPolicySetResult) bool { return v.IsProxy }).(pulumi.BoolOutput)
}

// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o LookupPolicySetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.Name }).(pulumi.StringOutput)
}

// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
func (o LookupPolicySetResultOutput) Rank() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPolicySetResult) int { return v.Rank }).(pulumi.IntOutput)
}

// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
func (o LookupPolicySetResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The state that the policy set is in. A disabled policy set cannot be matched.
func (o LookupPolicySetResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPolicySetResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPolicySetResultOutput{})
}

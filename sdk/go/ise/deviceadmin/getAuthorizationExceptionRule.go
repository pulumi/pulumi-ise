// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmin

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Device Admin Authorization Exception Rule.
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
//			_, err := deviceadmin.LookupAuthorizationExceptionRule(ctx, &deviceadmin.LookupAuthorizationExceptionRuleArgs{
//				Id:          pulumi.StringRef("76d24097-41c4-4558-a4d0-a8c07ac08470"),
//				PolicySetId: "d82952cb-b901-4b09-b363-5ebf39bdbaf9",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupAuthorizationExceptionRule(ctx *pulumi.Context, args *LookupAuthorizationExceptionRuleArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationExceptionRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationExceptionRuleResult
	err := ctx.Invoke("ise:deviceadmin/getAuthorizationExceptionRule:getAuthorizationExceptionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthorizationExceptionRule.
type LookupAuthorizationExceptionRuleArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// Policy set ID
	PolicySetId string `pulumi:"policySetId"`
}

// A collection of values returned by getAuthorizationExceptionRule.
type LookupAuthorizationExceptionRuleResult struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []GetAuthorizationExceptionRuleChildren `pulumi:"childrens"`
	// Command sets enforce the specified list of commands that can be executed by a device administrator
	CommandSets []string `pulumi:"commandSets"`
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
	// Indicates if this rule is the default one
	Default bool `pulumi:"default"`
	// The id of the object
	Id string `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name string `pulumi:"name"`
	// Policy set ID
	PolicySetId string `pulumi:"policySetId"`
	// Device admin profiles control the initial login session of the device administrator
	Profile string `pulumi:"profile"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank int `pulumi:"rank"`
	// The state that the rule is in. A disabled rule cannot be matched.
	State string `pulumi:"state"`
}

func LookupAuthorizationExceptionRuleOutput(ctx *pulumi.Context, args LookupAuthorizationExceptionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationExceptionRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationExceptionRuleResultOutput, error) {
			args := v.(LookupAuthorizationExceptionRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:deviceadmin/getAuthorizationExceptionRule:getAuthorizationExceptionRule", args, LookupAuthorizationExceptionRuleResultOutput{}, options).(LookupAuthorizationExceptionRuleResultOutput), nil
		}).(LookupAuthorizationExceptionRuleResultOutput)
}

// A collection of arguments for invoking getAuthorizationExceptionRule.
type LookupAuthorizationExceptionRuleOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Policy set ID
	PolicySetId pulumi.StringInput `pulumi:"policySetId"`
}

func (LookupAuthorizationExceptionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationExceptionRuleArgs)(nil)).Elem()
}

// A collection of values returned by getAuthorizationExceptionRule.
type LookupAuthorizationExceptionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationExceptionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationExceptionRuleResult)(nil)).Elem()
}

func (o LookupAuthorizationExceptionRuleResultOutput) ToLookupAuthorizationExceptionRuleResultOutput() LookupAuthorizationExceptionRuleResultOutput {
	return o
}

func (o LookupAuthorizationExceptionRuleResultOutput) ToLookupAuthorizationExceptionRuleResultOutputWithContext(ctx context.Context) LookupAuthorizationExceptionRuleResultOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o LookupAuthorizationExceptionRuleResultOutput) Childrens() GetAuthorizationExceptionRuleChildrenArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) []GetAuthorizationExceptionRuleChildren {
		return v.Childrens
	}).(GetAuthorizationExceptionRuleChildrenArrayOutput)
}

// Command sets enforce the specified list of commands that can be executed by a device administrator
func (o LookupAuthorizationExceptionRuleResultOutput) CommandSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) []string { return v.CommandSets }).(pulumi.StringArrayOutput)
}

// Dictionary attribute name
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionAttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionAttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionAttributeValue }).(pulumi.StringOutput)
}

// Dictionary name
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionDictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionDictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionDictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionDictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionId }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionIsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) bool { return v.ConditionIsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionOperator }).(pulumi.StringOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
func (o LookupAuthorizationExceptionRuleResultOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Indicates if this rule is the default one
func (o LookupAuthorizationExceptionRuleResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// The id of the object
func (o LookupAuthorizationExceptionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o LookupAuthorizationExceptionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy set ID
func (o LookupAuthorizationExceptionRuleResultOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.PolicySetId }).(pulumi.StringOutput)
}

// Device admin profiles control the initial login session of the device administrator
func (o LookupAuthorizationExceptionRuleResultOutput) Profile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.Profile }).(pulumi.StringOutput)
}

// The rank (priority) in relation to other rules. Lower rank is higher priority.
func (o LookupAuthorizationExceptionRuleResultOutput) Rank() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) int { return v.Rank }).(pulumi.IntOutput)
}

// The state that the rule is in. A disabled rule cannot be matched.
func (o LookupAuthorizationExceptionRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationExceptionRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationExceptionRuleResultOutput{})
}

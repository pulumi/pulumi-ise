// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Network Access Authorization Global Exception Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/networkaccess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkaccess.LookupAuthorizationGlobalExceptionRule(ctx, &networkaccess.LookupAuthorizationGlobalExceptionRuleArgs{
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
func LookupAuthorizationGlobalExceptionRule(ctx *pulumi.Context, args *LookupAuthorizationGlobalExceptionRuleArgs, opts ...pulumi.InvokeOption) (*LookupAuthorizationGlobalExceptionRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthorizationGlobalExceptionRuleResult
	err := ctx.Invoke("ise:networkaccess/getAuthorizationGlobalExceptionRule:getAuthorizationGlobalExceptionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthorizationGlobalExceptionRule.
type LookupAuthorizationGlobalExceptionRuleArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAuthorizationGlobalExceptionRule.
type LookupAuthorizationGlobalExceptionRuleResult struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []GetAuthorizationGlobalExceptionRuleChildren `pulumi:"childrens"`
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
	// The id of the object
	Id string `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name string `pulumi:"name"`
	// The authorization profile(s)
	Profiles []string `pulumi:"profiles"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank int `pulumi:"rank"`
	// Security group used in authorization policies
	SecurityGroup string `pulumi:"securityGroup"`
	// The state that the rule is in. A disabled rule cannot be matched.
	State string `pulumi:"state"`
}

func LookupAuthorizationGlobalExceptionRuleOutput(ctx *pulumi.Context, args LookupAuthorizationGlobalExceptionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAuthorizationGlobalExceptionRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupAuthorizationGlobalExceptionRuleResultOutput, error) {
			args := v.(LookupAuthorizationGlobalExceptionRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:networkaccess/getAuthorizationGlobalExceptionRule:getAuthorizationGlobalExceptionRule", args, LookupAuthorizationGlobalExceptionRuleResultOutput{}, options).(LookupAuthorizationGlobalExceptionRuleResultOutput), nil
		}).(LookupAuthorizationGlobalExceptionRuleResultOutput)
}

// A collection of arguments for invoking getAuthorizationGlobalExceptionRule.
type LookupAuthorizationGlobalExceptionRuleOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAuthorizationGlobalExceptionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationGlobalExceptionRuleArgs)(nil)).Elem()
}

// A collection of values returned by getAuthorizationGlobalExceptionRule.
type LookupAuthorizationGlobalExceptionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAuthorizationGlobalExceptionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthorizationGlobalExceptionRuleResult)(nil)).Elem()
}

func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ToLookupAuthorizationGlobalExceptionRuleResultOutput() LookupAuthorizationGlobalExceptionRuleResultOutput {
	return o
}

func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ToLookupAuthorizationGlobalExceptionRuleResultOutputWithContext(ctx context.Context) LookupAuthorizationGlobalExceptionRuleResultOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) Childrens() GetAuthorizationGlobalExceptionRuleChildrenArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) []GetAuthorizationGlobalExceptionRuleChildren {
		return v.Childrens
	}).(GetAuthorizationGlobalExceptionRuleChildrenArrayOutput)
}

// Dictionary attribute name
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionAttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionAttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionAttributeValue }).(pulumi.StringOutput)
}

// Dictionary name
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionDictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionDictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionDictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionDictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionId }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionIsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) bool { return v.ConditionIsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionOperator }).(pulumi.StringOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.ConditionType }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// The authorization profile(s)
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) Profiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) []string { return v.Profiles }).(pulumi.StringArrayOutput)
}

// The rank (priority) in relation to other rules. Lower rank is higher priority.
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) Rank() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) int { return v.Rank }).(pulumi.IntOutput)
}

// Security group used in authorization policies
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) SecurityGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.SecurityGroup }).(pulumi.StringOutput)
}

// The state that the rule is in. A disabled rule cannot be matched.
func (o LookupAuthorizationGlobalExceptionRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthorizationGlobalExceptionRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthorizationGlobalExceptionRuleResultOutput{})
}

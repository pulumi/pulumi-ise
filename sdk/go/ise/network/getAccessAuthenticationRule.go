// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Network Access Authentication Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/network"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := network.LookupAccessAuthenticationRule(ctx, &network.LookupAccessAuthenticationRuleArgs{
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
func LookupAccessAuthenticationRule(ctx *pulumi.Context, args *LookupAccessAuthenticationRuleArgs, opts ...pulumi.InvokeOption) (*LookupAccessAuthenticationRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAccessAuthenticationRuleResult
	err := ctx.Invoke("ise:network/getAccessAuthenticationRule:getAccessAuthenticationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAccessAuthenticationRule.
type LookupAccessAuthenticationRuleArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// Policy set ID
	PolicySetId string `pulumi:"policySetId"`
}

// A collection of values returned by getAccessAuthenticationRule.
type LookupAccessAuthenticationRuleResult struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []GetAccessAuthenticationRuleChildren `pulumi:"childrens"`
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
	// Identity source name from the identity stores
	IdentitySourceName string `pulumi:"identitySourceName"`
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on
	IfAuthFail string `pulumi:"ifAuthFail"`
	// Action to perform when ISE is uanble to access the identity database
	IfProcessFail string `pulumi:"ifProcessFail"`
	// Action to perform when user is not found in any of identity stores
	IfUserNotFound string `pulumi:"ifUserNotFound"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name string `pulumi:"name"`
	// Policy set ID
	PolicySetId string `pulumi:"policySetId"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank int `pulumi:"rank"`
	// The state that the rule is in. A disabled rule cannot be matched.
	State string `pulumi:"state"`
}

func LookupAccessAuthenticationRuleOutput(ctx *pulumi.Context, args LookupAccessAuthenticationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAccessAuthenticationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAccessAuthenticationRuleResult, error) {
			args := v.(LookupAccessAuthenticationRuleArgs)
			r, err := LookupAccessAuthenticationRule(ctx, &args, opts...)
			var s LookupAccessAuthenticationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAccessAuthenticationRuleResultOutput)
}

// A collection of arguments for invoking getAccessAuthenticationRule.
type LookupAccessAuthenticationRuleOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Policy set ID
	PolicySetId pulumi.StringInput `pulumi:"policySetId"`
}

func (LookupAccessAuthenticationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessAuthenticationRuleArgs)(nil)).Elem()
}

// A collection of values returned by getAccessAuthenticationRule.
type LookupAccessAuthenticationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAccessAuthenticationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAccessAuthenticationRuleResult)(nil)).Elem()
}

func (o LookupAccessAuthenticationRuleResultOutput) ToLookupAccessAuthenticationRuleResultOutput() LookupAccessAuthenticationRuleResultOutput {
	return o
}

func (o LookupAccessAuthenticationRuleResultOutput) ToLookupAccessAuthenticationRuleResultOutputWithContext(ctx context.Context) LookupAccessAuthenticationRuleResultOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o LookupAccessAuthenticationRuleResultOutput) Childrens() GetAccessAuthenticationRuleChildrenArrayOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) []GetAccessAuthenticationRuleChildren { return v.Childrens }).(GetAccessAuthenticationRuleChildrenArrayOutput)
}

// Dictionary attribute name
func (o LookupAccessAuthenticationRuleResultOutput) ConditionAttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionAttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o LookupAccessAuthenticationRuleResultOutput) ConditionAttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionAttributeValue }).(pulumi.StringOutput)
}

// Dictionary name
func (o LookupAccessAuthenticationRuleResultOutput) ConditionDictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionDictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o LookupAccessAuthenticationRuleResultOutput) ConditionDictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionDictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o LookupAccessAuthenticationRuleResultOutput) ConditionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionId }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupAccessAuthenticationRuleResultOutput) ConditionIsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) bool { return v.ConditionIsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o LookupAccessAuthenticationRuleResultOutput) ConditionOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionOperator }).(pulumi.StringOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
func (o LookupAccessAuthenticationRuleResultOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Indicates if this rule is the default one
func (o LookupAccessAuthenticationRuleResultOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) bool { return v.Default }).(pulumi.BoolOutput)
}

// The id of the object
func (o LookupAccessAuthenticationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity source name from the identity stores
func (o LookupAccessAuthenticationRuleResultOutput) IdentitySourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.IdentitySourceName }).(pulumi.StringOutput)
}

// Action to perform when authentication fails such as Bad credentials, disabled user and so on
func (o LookupAccessAuthenticationRuleResultOutput) IfAuthFail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.IfAuthFail }).(pulumi.StringOutput)
}

// Action to perform when ISE is uanble to access the identity database
func (o LookupAccessAuthenticationRuleResultOutput) IfProcessFail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.IfProcessFail }).(pulumi.StringOutput)
}

// Action to perform when user is not found in any of identity stores
func (o LookupAccessAuthenticationRuleResultOutput) IfUserNotFound() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.IfUserNotFound }).(pulumi.StringOutput)
}

// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o LookupAccessAuthenticationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy set ID
func (o LookupAccessAuthenticationRuleResultOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.PolicySetId }).(pulumi.StringOutput)
}

// The rank (priority) in relation to other rules. Lower rank is higher priority.
func (o LookupAccessAuthenticationRuleResultOutput) Rank() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) int { return v.Rank }).(pulumi.IntOutput)
}

// The state that the rule is in. A disabled rule cannot be matched.
func (o LookupAccessAuthenticationRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAccessAuthenticationRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAccessAuthenticationRuleResultOutput{})
}

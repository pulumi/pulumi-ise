// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccessauthorizationglobalexception

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type RuleChildren struct {
	// Dictionary attribute name
	AttributeName *string `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue *string `pulumi:"attributeValue"`
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []RuleChildrenChildren `pulumi:"childrens"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
	//   - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
	ConditionType string `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName *string `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue *string `pulumi:"dictionaryValue"`
	// UUID for condition
	Id *string `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate *bool `pulumi:"isNegate"`
	// Equality operator
	//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	Operator *string `pulumi:"operator"`
}

// RuleChildrenInput is an input type that accepts RuleChildrenArgs and RuleChildrenOutput values.
// You can construct a concrete instance of `RuleChildrenInput` via:
//
//	RuleChildrenArgs{...}
type RuleChildrenInput interface {
	pulumi.Input

	ToRuleChildrenOutput() RuleChildrenOutput
	ToRuleChildrenOutputWithContext(context.Context) RuleChildrenOutput
}

type RuleChildrenArgs struct {
	// Dictionary attribute name
	AttributeName pulumi.StringPtrInput `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue pulumi.StringPtrInput `pulumi:"attributeValue"`
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens RuleChildrenChildrenArrayInput `pulumi:"childrens"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
	//   - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
	ConditionType pulumi.StringInput `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName pulumi.StringPtrInput `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue pulumi.StringPtrInput `pulumi:"dictionaryValue"`
	// UUID for condition
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate pulumi.BoolPtrInput `pulumi:"isNegate"`
	// Equality operator
	//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	Operator pulumi.StringPtrInput `pulumi:"operator"`
}

func (RuleChildrenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleChildren)(nil)).Elem()
}

func (i RuleChildrenArgs) ToRuleChildrenOutput() RuleChildrenOutput {
	return i.ToRuleChildrenOutputWithContext(context.Background())
}

func (i RuleChildrenArgs) ToRuleChildrenOutputWithContext(ctx context.Context) RuleChildrenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleChildrenOutput)
}

// RuleChildrenArrayInput is an input type that accepts RuleChildrenArray and RuleChildrenArrayOutput values.
// You can construct a concrete instance of `RuleChildrenArrayInput` via:
//
//	RuleChildrenArray{ RuleChildrenArgs{...} }
type RuleChildrenArrayInput interface {
	pulumi.Input

	ToRuleChildrenArrayOutput() RuleChildrenArrayOutput
	ToRuleChildrenArrayOutputWithContext(context.Context) RuleChildrenArrayOutput
}

type RuleChildrenArray []RuleChildrenInput

func (RuleChildrenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleChildren)(nil)).Elem()
}

func (i RuleChildrenArray) ToRuleChildrenArrayOutput() RuleChildrenArrayOutput {
	return i.ToRuleChildrenArrayOutputWithContext(context.Background())
}

func (i RuleChildrenArray) ToRuleChildrenArrayOutputWithContext(ctx context.Context) RuleChildrenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleChildrenArrayOutput)
}

type RuleChildrenOutput struct{ *pulumi.OutputState }

func (RuleChildrenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleChildren)(nil)).Elem()
}

func (o RuleChildrenOutput) ToRuleChildrenOutput() RuleChildrenOutput {
	return o
}

func (o RuleChildrenOutput) ToRuleChildrenOutputWithContext(ctx context.Context) RuleChildrenOutput {
	return o
}

// Dictionary attribute name
func (o RuleChildrenOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o RuleChildrenOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o RuleChildrenOutput) Childrens() RuleChildrenChildrenArrayOutput {
	return o.ApplyT(func(v RuleChildren) []RuleChildrenChildren { return v.Childrens }).(RuleChildrenChildrenArrayOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
//   - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
func (o RuleChildrenOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleChildren) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Dictionary name
func (o RuleChildrenOutput) DictionaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.DictionaryName }).(pulumi.StringPtrOutput)
}

// Dictionary value
func (o RuleChildrenOutput) DictionaryValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.DictionaryValue }).(pulumi.StringPtrOutput)
}

// UUID for condition
func (o RuleChildrenOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Indicates whereas this condition is in negate mode
func (o RuleChildrenOutput) IsNegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleChildren) *bool { return v.IsNegate }).(pulumi.BoolPtrOutput)
}

// Equality operator
//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
func (o RuleChildrenOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildren) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

type RuleChildrenArrayOutput struct{ *pulumi.OutputState }

func (RuleChildrenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleChildren)(nil)).Elem()
}

func (o RuleChildrenArrayOutput) ToRuleChildrenArrayOutput() RuleChildrenArrayOutput {
	return o
}

func (o RuleChildrenArrayOutput) ToRuleChildrenArrayOutputWithContext(ctx context.Context) RuleChildrenArrayOutput {
	return o
}

func (o RuleChildrenArrayOutput) Index(i pulumi.IntInput) RuleChildrenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleChildren {
		return vs[0].([]RuleChildren)[vs[1].(int)]
	}).(RuleChildrenOutput)
}

type RuleChildrenChildren struct {
	// Dictionary attribute name
	AttributeName *string `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue *string `pulumi:"attributeValue"`
	// Condition type.
	//   - Choices: `ConditionAttributes`, `ConditionReference`
	ConditionType string `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName *string `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue *string `pulumi:"dictionaryValue"`
	// UUID for condition
	Id *string `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate *bool `pulumi:"isNegate"`
	// Equality operator
	//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	Operator *string `pulumi:"operator"`
}

// RuleChildrenChildrenInput is an input type that accepts RuleChildrenChildrenArgs and RuleChildrenChildrenOutput values.
// You can construct a concrete instance of `RuleChildrenChildrenInput` via:
//
//	RuleChildrenChildrenArgs{...}
type RuleChildrenChildrenInput interface {
	pulumi.Input

	ToRuleChildrenChildrenOutput() RuleChildrenChildrenOutput
	ToRuleChildrenChildrenOutputWithContext(context.Context) RuleChildrenChildrenOutput
}

type RuleChildrenChildrenArgs struct {
	// Dictionary attribute name
	AttributeName pulumi.StringPtrInput `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue pulumi.StringPtrInput `pulumi:"attributeValue"`
	// Condition type.
	//   - Choices: `ConditionAttributes`, `ConditionReference`
	ConditionType pulumi.StringInput `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName pulumi.StringPtrInput `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue pulumi.StringPtrInput `pulumi:"dictionaryValue"`
	// UUID for condition
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate pulumi.BoolPtrInput `pulumi:"isNegate"`
	// Equality operator
	//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	Operator pulumi.StringPtrInput `pulumi:"operator"`
}

func (RuleChildrenChildrenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleChildrenChildren)(nil)).Elem()
}

func (i RuleChildrenChildrenArgs) ToRuleChildrenChildrenOutput() RuleChildrenChildrenOutput {
	return i.ToRuleChildrenChildrenOutputWithContext(context.Background())
}

func (i RuleChildrenChildrenArgs) ToRuleChildrenChildrenOutputWithContext(ctx context.Context) RuleChildrenChildrenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleChildrenChildrenOutput)
}

// RuleChildrenChildrenArrayInput is an input type that accepts RuleChildrenChildrenArray and RuleChildrenChildrenArrayOutput values.
// You can construct a concrete instance of `RuleChildrenChildrenArrayInput` via:
//
//	RuleChildrenChildrenArray{ RuleChildrenChildrenArgs{...} }
type RuleChildrenChildrenArrayInput interface {
	pulumi.Input

	ToRuleChildrenChildrenArrayOutput() RuleChildrenChildrenArrayOutput
	ToRuleChildrenChildrenArrayOutputWithContext(context.Context) RuleChildrenChildrenArrayOutput
}

type RuleChildrenChildrenArray []RuleChildrenChildrenInput

func (RuleChildrenChildrenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleChildrenChildren)(nil)).Elem()
}

func (i RuleChildrenChildrenArray) ToRuleChildrenChildrenArrayOutput() RuleChildrenChildrenArrayOutput {
	return i.ToRuleChildrenChildrenArrayOutputWithContext(context.Background())
}

func (i RuleChildrenChildrenArray) ToRuleChildrenChildrenArrayOutputWithContext(ctx context.Context) RuleChildrenChildrenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleChildrenChildrenArrayOutput)
}

type RuleChildrenChildrenOutput struct{ *pulumi.OutputState }

func (RuleChildrenChildrenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RuleChildrenChildren)(nil)).Elem()
}

func (o RuleChildrenChildrenOutput) ToRuleChildrenChildrenOutput() RuleChildrenChildrenOutput {
	return o
}

func (o RuleChildrenChildrenOutput) ToRuleChildrenChildrenOutputWithContext(ctx context.Context) RuleChildrenChildrenOutput {
	return o
}

// Dictionary attribute name
func (o RuleChildrenChildrenOutput) AttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.AttributeName }).(pulumi.StringPtrOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o RuleChildrenChildrenOutput) AttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.AttributeValue }).(pulumi.StringPtrOutput)
}

// Condition type.
//   - Choices: `ConditionAttributes`, `ConditionReference`
func (o RuleChildrenChildrenOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v RuleChildrenChildren) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Dictionary name
func (o RuleChildrenChildrenOutput) DictionaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.DictionaryName }).(pulumi.StringPtrOutput)
}

// Dictionary value
func (o RuleChildrenChildrenOutput) DictionaryValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.DictionaryValue }).(pulumi.StringPtrOutput)
}

// UUID for condition
func (o RuleChildrenChildrenOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Indicates whereas this condition is in negate mode
func (o RuleChildrenChildrenOutput) IsNegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *bool { return v.IsNegate }).(pulumi.BoolPtrOutput)
}

// Equality operator
//   - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
func (o RuleChildrenChildrenOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RuleChildrenChildren) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

type RuleChildrenChildrenArrayOutput struct{ *pulumi.OutputState }

func (RuleChildrenChildrenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RuleChildrenChildren)(nil)).Elem()
}

func (o RuleChildrenChildrenArrayOutput) ToRuleChildrenChildrenArrayOutput() RuleChildrenChildrenArrayOutput {
	return o
}

func (o RuleChildrenChildrenArrayOutput) ToRuleChildrenChildrenArrayOutputWithContext(ctx context.Context) RuleChildrenChildrenArrayOutput {
	return o
}

func (o RuleChildrenChildrenArrayOutput) Index(i pulumi.IntInput) RuleChildrenChildrenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RuleChildrenChildren {
		return vs[0].([]RuleChildrenChildren)[vs[1].(int)]
	}).(RuleChildrenChildrenOutput)
}

type GetRuleChildren struct {
	// Dictionary attribute name
	AttributeName string `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue string `pulumi:"attributeValue"`
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []GetRuleChildrenChildren `pulumi:"childrens"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
	ConditionType string `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName string `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue string `pulumi:"dictionaryValue"`
	// UUID for condition
	Id string `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate bool `pulumi:"isNegate"`
	// Equality operator
	Operator string `pulumi:"operator"`
}

// GetRuleChildrenInput is an input type that accepts GetRuleChildrenArgs and GetRuleChildrenOutput values.
// You can construct a concrete instance of `GetRuleChildrenInput` via:
//
//	GetRuleChildrenArgs{...}
type GetRuleChildrenInput interface {
	pulumi.Input

	ToGetRuleChildrenOutput() GetRuleChildrenOutput
	ToGetRuleChildrenOutputWithContext(context.Context) GetRuleChildrenOutput
}

type GetRuleChildrenArgs struct {
	// Dictionary attribute name
	AttributeName pulumi.StringInput `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue pulumi.StringInput `pulumi:"attributeValue"`
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens GetRuleChildrenChildrenArrayInput `pulumi:"childrens"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
	ConditionType pulumi.StringInput `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName pulumi.StringInput `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue pulumi.StringInput `pulumi:"dictionaryValue"`
	// UUID for condition
	Id pulumi.StringInput `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate pulumi.BoolInput `pulumi:"isNegate"`
	// Equality operator
	Operator pulumi.StringInput `pulumi:"operator"`
}

func (GetRuleChildrenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleChildren)(nil)).Elem()
}

func (i GetRuleChildrenArgs) ToGetRuleChildrenOutput() GetRuleChildrenOutput {
	return i.ToGetRuleChildrenOutputWithContext(context.Background())
}

func (i GetRuleChildrenArgs) ToGetRuleChildrenOutputWithContext(ctx context.Context) GetRuleChildrenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRuleChildrenOutput)
}

// GetRuleChildrenArrayInput is an input type that accepts GetRuleChildrenArray and GetRuleChildrenArrayOutput values.
// You can construct a concrete instance of `GetRuleChildrenArrayInput` via:
//
//	GetRuleChildrenArray{ GetRuleChildrenArgs{...} }
type GetRuleChildrenArrayInput interface {
	pulumi.Input

	ToGetRuleChildrenArrayOutput() GetRuleChildrenArrayOutput
	ToGetRuleChildrenArrayOutputWithContext(context.Context) GetRuleChildrenArrayOutput
}

type GetRuleChildrenArray []GetRuleChildrenInput

func (GetRuleChildrenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRuleChildren)(nil)).Elem()
}

func (i GetRuleChildrenArray) ToGetRuleChildrenArrayOutput() GetRuleChildrenArrayOutput {
	return i.ToGetRuleChildrenArrayOutputWithContext(context.Background())
}

func (i GetRuleChildrenArray) ToGetRuleChildrenArrayOutputWithContext(ctx context.Context) GetRuleChildrenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRuleChildrenArrayOutput)
}

type GetRuleChildrenOutput struct{ *pulumi.OutputState }

func (GetRuleChildrenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleChildren)(nil)).Elem()
}

func (o GetRuleChildrenOutput) ToGetRuleChildrenOutput() GetRuleChildrenOutput {
	return o
}

func (o GetRuleChildrenOutput) ToGetRuleChildrenOutputWithContext(ctx context.Context) GetRuleChildrenOutput {
	return o
}

// Dictionary attribute name
func (o GetRuleChildrenOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.AttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o GetRuleChildrenOutput) AttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.AttributeValue }).(pulumi.StringOutput)
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o GetRuleChildrenOutput) Childrens() GetRuleChildrenChildrenArrayOutput {
	return o.ApplyT(func(v GetRuleChildren) []GetRuleChildrenChildren { return v.Childrens }).(GetRuleChildrenChildrenArrayOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
func (o GetRuleChildrenOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Dictionary name
func (o GetRuleChildrenOutput) DictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.DictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o GetRuleChildrenOutput) DictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.DictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o GetRuleChildrenOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o GetRuleChildrenOutput) IsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v GetRuleChildren) bool { return v.IsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o GetRuleChildrenOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildren) string { return v.Operator }).(pulumi.StringOutput)
}

type GetRuleChildrenArrayOutput struct{ *pulumi.OutputState }

func (GetRuleChildrenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRuleChildren)(nil)).Elem()
}

func (o GetRuleChildrenArrayOutput) ToGetRuleChildrenArrayOutput() GetRuleChildrenArrayOutput {
	return o
}

func (o GetRuleChildrenArrayOutput) ToGetRuleChildrenArrayOutputWithContext(ctx context.Context) GetRuleChildrenArrayOutput {
	return o
}

func (o GetRuleChildrenArrayOutput) Index(i pulumi.IntInput) GetRuleChildrenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetRuleChildren {
		return vs[0].([]GetRuleChildren)[vs[1].(int)]
	}).(GetRuleChildrenOutput)
}

type GetRuleChildrenChildren struct {
	// Dictionary attribute name
	AttributeName string `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue string `pulumi:"attributeValue"`
	// Condition type.
	ConditionType string `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName string `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue string `pulumi:"dictionaryValue"`
	// UUID for condition
	Id string `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate bool `pulumi:"isNegate"`
	// Equality operator
	Operator string `pulumi:"operator"`
}

// GetRuleChildrenChildrenInput is an input type that accepts GetRuleChildrenChildrenArgs and GetRuleChildrenChildrenOutput values.
// You can construct a concrete instance of `GetRuleChildrenChildrenInput` via:
//
//	GetRuleChildrenChildrenArgs{...}
type GetRuleChildrenChildrenInput interface {
	pulumi.Input

	ToGetRuleChildrenChildrenOutput() GetRuleChildrenChildrenOutput
	ToGetRuleChildrenChildrenOutputWithContext(context.Context) GetRuleChildrenChildrenOutput
}

type GetRuleChildrenChildrenArgs struct {
	// Dictionary attribute name
	AttributeName pulumi.StringInput `pulumi:"attributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	AttributeValue pulumi.StringInput `pulumi:"attributeValue"`
	// Condition type.
	ConditionType pulumi.StringInput `pulumi:"conditionType"`
	// Dictionary name
	DictionaryName pulumi.StringInput `pulumi:"dictionaryName"`
	// Dictionary value
	DictionaryValue pulumi.StringInput `pulumi:"dictionaryValue"`
	// UUID for condition
	Id pulumi.StringInput `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate pulumi.BoolInput `pulumi:"isNegate"`
	// Equality operator
	Operator pulumi.StringInput `pulumi:"operator"`
}

func (GetRuleChildrenChildrenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleChildrenChildren)(nil)).Elem()
}

func (i GetRuleChildrenChildrenArgs) ToGetRuleChildrenChildrenOutput() GetRuleChildrenChildrenOutput {
	return i.ToGetRuleChildrenChildrenOutputWithContext(context.Background())
}

func (i GetRuleChildrenChildrenArgs) ToGetRuleChildrenChildrenOutputWithContext(ctx context.Context) GetRuleChildrenChildrenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRuleChildrenChildrenOutput)
}

// GetRuleChildrenChildrenArrayInput is an input type that accepts GetRuleChildrenChildrenArray and GetRuleChildrenChildrenArrayOutput values.
// You can construct a concrete instance of `GetRuleChildrenChildrenArrayInput` via:
//
//	GetRuleChildrenChildrenArray{ GetRuleChildrenChildrenArgs{...} }
type GetRuleChildrenChildrenArrayInput interface {
	pulumi.Input

	ToGetRuleChildrenChildrenArrayOutput() GetRuleChildrenChildrenArrayOutput
	ToGetRuleChildrenChildrenArrayOutputWithContext(context.Context) GetRuleChildrenChildrenArrayOutput
}

type GetRuleChildrenChildrenArray []GetRuleChildrenChildrenInput

func (GetRuleChildrenChildrenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRuleChildrenChildren)(nil)).Elem()
}

func (i GetRuleChildrenChildrenArray) ToGetRuleChildrenChildrenArrayOutput() GetRuleChildrenChildrenArrayOutput {
	return i.ToGetRuleChildrenChildrenArrayOutputWithContext(context.Background())
}

func (i GetRuleChildrenChildrenArray) ToGetRuleChildrenChildrenArrayOutputWithContext(ctx context.Context) GetRuleChildrenChildrenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetRuleChildrenChildrenArrayOutput)
}

type GetRuleChildrenChildrenOutput struct{ *pulumi.OutputState }

func (GetRuleChildrenChildrenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRuleChildrenChildren)(nil)).Elem()
}

func (o GetRuleChildrenChildrenOutput) ToGetRuleChildrenChildrenOutput() GetRuleChildrenChildrenOutput {
	return o
}

func (o GetRuleChildrenChildrenOutput) ToGetRuleChildrenChildrenOutputWithContext(ctx context.Context) GetRuleChildrenChildrenOutput {
	return o
}

// Dictionary attribute name
func (o GetRuleChildrenChildrenOutput) AttributeName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.AttributeName }).(pulumi.StringOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o GetRuleChildrenChildrenOutput) AttributeValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.AttributeValue }).(pulumi.StringOutput)
}

// Condition type.
func (o GetRuleChildrenChildrenOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.ConditionType }).(pulumi.StringOutput)
}

// Dictionary name
func (o GetRuleChildrenChildrenOutput) DictionaryName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.DictionaryName }).(pulumi.StringOutput)
}

// Dictionary value
func (o GetRuleChildrenChildrenOutput) DictionaryValue() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.DictionaryValue }).(pulumi.StringOutput)
}

// UUID for condition
func (o GetRuleChildrenChildrenOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o GetRuleChildrenChildrenOutput) IsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) bool { return v.IsNegate }).(pulumi.BoolOutput)
}

// Equality operator
func (o GetRuleChildrenChildrenOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v GetRuleChildrenChildren) string { return v.Operator }).(pulumi.StringOutput)
}

type GetRuleChildrenChildrenArrayOutput struct{ *pulumi.OutputState }

func (GetRuleChildrenChildrenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetRuleChildrenChildren)(nil)).Elem()
}

func (o GetRuleChildrenChildrenArrayOutput) ToGetRuleChildrenChildrenArrayOutput() GetRuleChildrenChildrenArrayOutput {
	return o
}

func (o GetRuleChildrenChildrenArrayOutput) ToGetRuleChildrenChildrenArrayOutputWithContext(ctx context.Context) GetRuleChildrenChildrenArrayOutput {
	return o
}

func (o GetRuleChildrenChildrenArrayOutput) Index(i pulumi.IntInput) GetRuleChildrenChildrenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetRuleChildrenChildren {
		return vs[0].([]GetRuleChildrenChildren)[vs[1].(int)]
	}).(GetRuleChildrenChildrenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleChildrenInput)(nil)).Elem(), RuleChildrenArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleChildrenArrayInput)(nil)).Elem(), RuleChildrenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleChildrenChildrenInput)(nil)).Elem(), RuleChildrenChildrenArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleChildrenChildrenArrayInput)(nil)).Elem(), RuleChildrenChildrenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRuleChildrenInput)(nil)).Elem(), GetRuleChildrenArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRuleChildrenArrayInput)(nil)).Elem(), GetRuleChildrenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRuleChildrenChildrenInput)(nil)).Elem(), GetRuleChildrenChildrenArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetRuleChildrenChildrenArrayInput)(nil)).Elem(), GetRuleChildrenChildrenArray{})
	pulumi.RegisterOutputType(RuleChildrenOutput{})
	pulumi.RegisterOutputType(RuleChildrenArrayOutput{})
	pulumi.RegisterOutputType(RuleChildrenChildrenOutput{})
	pulumi.RegisterOutputType(RuleChildrenChildrenArrayOutput{})
	pulumi.RegisterOutputType(GetRuleChildrenOutput{})
	pulumi.RegisterOutputType(GetRuleChildrenArrayOutput{})
	pulumi.RegisterOutputType(GetRuleChildrenChildrenOutput{})
	pulumi.RegisterOutputType(GetRuleChildrenChildrenArrayOutput{})
}

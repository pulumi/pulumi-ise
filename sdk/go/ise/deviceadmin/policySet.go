// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmin

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a Device Admin Policy Set.
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
//			_, err := deviceadmin.NewPolicySet(ctx, "example", &deviceadmin.PolicySetArgs{
//				Name:                    pulumi.String("PolicySet1"),
//				Description:             pulumi.String("My description"),
//				IsProxy:                 pulumi.Bool(false),
//				Rank:                    pulumi.Int(0),
//				ServiceName:             pulumi.String("Default Device Admin"),
//				State:                   pulumi.String("enabled"),
//				ConditionType:           pulumi.String("ConditionAttributes"),
//				ConditionIsNegate:       pulumi.Bool(false),
//				ConditionAttributeName:  pulumi.String("Location"),
//				ConditionAttributeValue: pulumi.String("All Locations"),
//				ConditionDictionaryName: pulumi.String("DEVICE"),
//				ConditionOperator:       pulumi.String("equals"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import ise:deviceadmin/policySet:PolicySet example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type PolicySet struct {
	pulumi.CustomResourceState

	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens PolicySetChildrenArrayOutput `pulumi:"childrens"`
	// Dictionary attribute name
	ConditionAttributeName pulumi.StringPtrOutput `pulumi:"conditionAttributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue pulumi.StringPtrOutput `pulumi:"conditionAttributeValue"`
	// Dictionary name
	ConditionDictionaryName pulumi.StringPtrOutput `pulumi:"conditionDictionaryName"`
	// Dictionary value
	ConditionDictionaryValue pulumi.StringPtrOutput `pulumi:"conditionDictionaryValue"`
	// UUID for condition
	ConditionId pulumi.StringPtrOutput `pulumi:"conditionId"`
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate pulumi.BoolPtrOutput `pulumi:"conditionIsNegate"`
	// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
	// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
	// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	ConditionOperator pulumi.StringPtrOutput `pulumi:"conditionOperator"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
	// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
	// `ConditionOrBlock`, `ConditionReference`
	ConditionType pulumi.StringPtrOutput `pulumi:"conditionType"`
	// Indicates if this policy set is the default one
	Default pulumi.BoolPtrOutput `pulumi:"default"`
	// The description of the policy set
	Description pulumi.StringOutput `pulumi:"description"`
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy pulumi.BoolPtrOutput `pulumi:"isProxy"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringOutput `pulumi:"name"`
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank pulumi.IntPtrOutput `pulumi:"rank"`
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
	// `monitor`
	State pulumi.StringPtrOutput `pulumi:"state"`
}

// NewPolicySet registers a new resource with the given unique name, arguments, and options.
func NewPolicySet(ctx *pulumi.Context,
	name string, args *PolicySetArgs, opts ...pulumi.ResourceOption) (*PolicySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PolicySet
	err := ctx.RegisterResource("ise:deviceadmin/policySet:PolicySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySet gets an existing PolicySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetState, opts ...pulumi.ResourceOption) (*PolicySet, error) {
	var resource PolicySet
	err := ctx.ReadResource("ise:deviceadmin/policySet:PolicySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySet resources.
type policySetState struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []PolicySetChildren `pulumi:"childrens"`
	// Dictionary attribute name
	ConditionAttributeName *string `pulumi:"conditionAttributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue *string `pulumi:"conditionAttributeValue"`
	// Dictionary name
	ConditionDictionaryName *string `pulumi:"conditionDictionaryName"`
	// Dictionary value
	ConditionDictionaryValue *string `pulumi:"conditionDictionaryValue"`
	// UUID for condition
	ConditionId *string `pulumi:"conditionId"`
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate *bool `pulumi:"conditionIsNegate"`
	// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
	// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
	// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	ConditionOperator *string `pulumi:"conditionOperator"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
	// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
	// `ConditionOrBlock`, `ConditionReference`
	ConditionType *string `pulumi:"conditionType"`
	// Indicates if this policy set is the default one
	Default *bool `pulumi:"default"`
	// The description of the policy set
	Description *string `pulumi:"description"`
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy *bool `pulumi:"isProxy"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank *int `pulumi:"rank"`
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName *string `pulumi:"serviceName"`
	// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
	// `monitor`
	State *string `pulumi:"state"`
}

type PolicySetState struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens PolicySetChildrenArrayInput
	// Dictionary attribute name
	ConditionAttributeName pulumi.StringPtrInput
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue pulumi.StringPtrInput
	// Dictionary name
	ConditionDictionaryName pulumi.StringPtrInput
	// Dictionary value
	ConditionDictionaryValue pulumi.StringPtrInput
	// UUID for condition
	ConditionId pulumi.StringPtrInput
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate pulumi.BoolPtrInput
	// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
	// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
	// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	ConditionOperator pulumi.StringPtrInput
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
	// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
	// `ConditionOrBlock`, `ConditionReference`
	ConditionType pulumi.StringPtrInput
	// Indicates if this policy set is the default one
	Default pulumi.BoolPtrInput
	// The description of the policy set
	Description pulumi.StringPtrInput
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy pulumi.BoolPtrInput
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank pulumi.IntPtrInput
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName pulumi.StringPtrInput
	// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
	// `monitor`
	State pulumi.StringPtrInput
}

func (PolicySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetState)(nil)).Elem()
}

type policySetArgs struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []PolicySetChildren `pulumi:"childrens"`
	// Dictionary attribute name
	ConditionAttributeName *string `pulumi:"conditionAttributeName"`
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue *string `pulumi:"conditionAttributeValue"`
	// Dictionary name
	ConditionDictionaryName *string `pulumi:"conditionDictionaryName"`
	// Dictionary value
	ConditionDictionaryValue *string `pulumi:"conditionDictionaryValue"`
	// UUID for condition
	ConditionId *string `pulumi:"conditionId"`
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate *bool `pulumi:"conditionIsNegate"`
	// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
	// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
	// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	ConditionOperator *string `pulumi:"conditionOperator"`
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
	// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
	// `ConditionOrBlock`, `ConditionReference`
	ConditionType *string `pulumi:"conditionType"`
	// Indicates if this policy set is the default one
	Default *bool `pulumi:"default"`
	// The description of the policy set
	Description *string `pulumi:"description"`
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy *bool `pulumi:"isProxy"`
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank *int `pulumi:"rank"`
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName string `pulumi:"serviceName"`
	// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
	// `monitor`
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a PolicySet resource.
type PolicySetArgs struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens PolicySetChildrenArrayInput
	// Dictionary attribute name
	ConditionAttributeName pulumi.StringPtrInput
	// Attribute value for condition. Value type is specified in dictionary object.
	ConditionAttributeValue pulumi.StringPtrInput
	// Dictionary name
	ConditionDictionaryName pulumi.StringPtrInput
	// Dictionary value
	ConditionDictionaryValue pulumi.StringPtrInput
	// UUID for condition
	ConditionId pulumi.StringPtrInput
	// Indicates whereas this condition is in negate mode
	ConditionIsNegate pulumi.BoolPtrInput
	// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
	// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
	// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
	ConditionOperator pulumi.StringPtrInput
	// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
	// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
	// `ConditionOrBlock`, `ConditionReference`
	ConditionType pulumi.StringPtrInput
	// Indicates if this policy set is the default one
	Default pulumi.BoolPtrInput
	// The description of the policy set
	Description pulumi.StringPtrInput
	// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
	IsProxy pulumi.BoolPtrInput
	// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput
	// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
	Rank pulumi.IntPtrInput
	// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
	ServiceName pulumi.StringInput
	// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
	// `monitor`
	State pulumi.StringPtrInput
}

func (PolicySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetArgs)(nil)).Elem()
}

type PolicySetInput interface {
	pulumi.Input

	ToPolicySetOutput() PolicySetOutput
	ToPolicySetOutputWithContext(ctx context.Context) PolicySetOutput
}

func (*PolicySet) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySet)(nil)).Elem()
}

func (i *PolicySet) ToPolicySetOutput() PolicySetOutput {
	return i.ToPolicySetOutputWithContext(context.Background())
}

func (i *PolicySet) ToPolicySetOutputWithContext(ctx context.Context) PolicySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetOutput)
}

// PolicySetArrayInput is an input type that accepts PolicySetArray and PolicySetArrayOutput values.
// You can construct a concrete instance of `PolicySetArrayInput` via:
//
//	PolicySetArray{ PolicySetArgs{...} }
type PolicySetArrayInput interface {
	pulumi.Input

	ToPolicySetArrayOutput() PolicySetArrayOutput
	ToPolicySetArrayOutputWithContext(context.Context) PolicySetArrayOutput
}

type PolicySetArray []PolicySetInput

func (PolicySetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicySet)(nil)).Elem()
}

func (i PolicySetArray) ToPolicySetArrayOutput() PolicySetArrayOutput {
	return i.ToPolicySetArrayOutputWithContext(context.Background())
}

func (i PolicySetArray) ToPolicySetArrayOutputWithContext(ctx context.Context) PolicySetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetArrayOutput)
}

// PolicySetMapInput is an input type that accepts PolicySetMap and PolicySetMapOutput values.
// You can construct a concrete instance of `PolicySetMapInput` via:
//
//	PolicySetMap{ "key": PolicySetArgs{...} }
type PolicySetMapInput interface {
	pulumi.Input

	ToPolicySetMapOutput() PolicySetMapOutput
	ToPolicySetMapOutputWithContext(context.Context) PolicySetMapOutput
}

type PolicySetMap map[string]PolicySetInput

func (PolicySetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicySet)(nil)).Elem()
}

func (i PolicySetMap) ToPolicySetMapOutput() PolicySetMapOutput {
	return i.ToPolicySetMapOutputWithContext(context.Background())
}

func (i PolicySetMap) ToPolicySetMapOutputWithContext(ctx context.Context) PolicySetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicySetMapOutput)
}

type PolicySetOutput struct{ *pulumi.OutputState }

func (PolicySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicySet)(nil)).Elem()
}

func (o PolicySetOutput) ToPolicySetOutput() PolicySetOutput {
	return o
}

func (o PolicySetOutput) ToPolicySetOutputWithContext(ctx context.Context) PolicySetOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o PolicySetOutput) Childrens() PolicySetChildrenArrayOutput {
	return o.ApplyT(func(v *PolicySet) PolicySetChildrenArrayOutput { return v.Childrens }).(PolicySetChildrenArrayOutput)
}

// Dictionary attribute name
func (o PolicySetOutput) ConditionAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionAttributeName }).(pulumi.StringPtrOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o PolicySetOutput) ConditionAttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionAttributeValue }).(pulumi.StringPtrOutput)
}

// Dictionary name
func (o PolicySetOutput) ConditionDictionaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionDictionaryName }).(pulumi.StringPtrOutput)
}

// Dictionary value
func (o PolicySetOutput) ConditionDictionaryValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionDictionaryValue }).(pulumi.StringPtrOutput)
}

// UUID for condition
func (o PolicySetOutput) ConditionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionId }).(pulumi.StringPtrOutput)
}

// Indicates whereas this condition is in negate mode
func (o PolicySetOutput) ConditionIsNegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.BoolPtrOutput { return v.ConditionIsNegate }).(pulumi.BoolPtrOutput)
}

// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
func (o PolicySetOutput) ConditionOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionOperator }).(pulumi.StringPtrOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
// `ConditionOrBlock`, `ConditionReference`
func (o PolicySetOutput) ConditionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.ConditionType }).(pulumi.StringPtrOutput)
}

// Indicates if this policy set is the default one
func (o PolicySetOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.BoolPtrOutput { return v.Default }).(pulumi.BoolPtrOutput)
}

// The description of the policy set
func (o PolicySetOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
func (o PolicySetOutput) IsProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.BoolPtrOutput { return v.IsProxy }).(pulumi.BoolPtrOutput)
}

// Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o PolicySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The rank (priority) in relation to other policy sets. Lower rank is higher priority.
func (o PolicySetOutput) Rank() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.IntPtrOutput { return v.Rank }).(pulumi.IntPtrOutput)
}

// Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
func (o PolicySetOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The state that the policy set is in. A disabled policy set cannot be matched. - Choices: `disabled`, `enabled`,
// `monitor`
func (o PolicySetOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicySet) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

type PolicySetArrayOutput struct{ *pulumi.OutputState }

func (PolicySetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PolicySet)(nil)).Elem()
}

func (o PolicySetArrayOutput) ToPolicySetArrayOutput() PolicySetArrayOutput {
	return o
}

func (o PolicySetArrayOutput) ToPolicySetArrayOutputWithContext(ctx context.Context) PolicySetArrayOutput {
	return o
}

func (o PolicySetArrayOutput) Index(i pulumi.IntInput) PolicySetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PolicySet {
		return vs[0].([]*PolicySet)[vs[1].(int)]
	}).(PolicySetOutput)
}

type PolicySetMapOutput struct{ *pulumi.OutputState }

func (PolicySetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PolicySet)(nil)).Elem()
}

func (o PolicySetMapOutput) ToPolicySetMapOutput() PolicySetMapOutput {
	return o
}

func (o PolicySetMapOutput) ToPolicySetMapOutputWithContext(ctx context.Context) PolicySetMapOutput {
	return o
}

func (o PolicySetMapOutput) MapIndex(k pulumi.StringInput) PolicySetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PolicySet {
		return vs[0].(map[string]*PolicySet)[vs[1].(string)]
	}).(PolicySetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySetInput)(nil)).Elem(), &PolicySet{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySetArrayInput)(nil)).Elem(), PolicySetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicySetMapInput)(nil)).Elem(), PolicySetMap{})
	pulumi.RegisterOutputType(PolicySetOutput{})
	pulumi.RegisterOutputType(PolicySetArrayOutput{})
	pulumi.RegisterOutputType(PolicySetMapOutput{})
}

// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmin

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a Device Admin Authentication Rule.
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
//			_, err := deviceadmin.NewAuthenticationRule(ctx, "example", &deviceadmin.AuthenticationRuleArgs{
//				PolicySetId:             pulumi.String("d82952cb-b901-4b09-b363-5ebf39bdbaf9"),
//				Name:                    pulumi.String("Rule1"),
//				Default:                 pulumi.Bool(false),
//				Rank:                    pulumi.Int(0),
//				State:                   pulumi.String("enabled"),
//				ConditionType:           pulumi.String("ConditionAttributes"),
//				ConditionIsNegate:       pulumi.Bool(false),
//				ConditionAttributeName:  pulumi.String("Location"),
//				ConditionAttributeValue: pulumi.String("All Locations"),
//				ConditionDictionaryName: pulumi.String("DEVICE"),
//				ConditionOperator:       pulumi.String("equals"),
//				IdentitySourceName:      pulumi.String("Internal Endpoints"),
//				IfAuthFail:              pulumi.String("REJECT"),
//				IfProcessFail:           pulumi.String("DROP"),
//				IfUserNotFound:          pulumi.String("REJECT"),
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
// $ pulumi import ise:deviceadmin/authenticationRule:AuthenticationRule example "76d24097-41c4-4558-a4d0-a8c07ac08470,76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type AuthenticationRule struct {
	pulumi.CustomResourceState

	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens AuthenticationRuleChildrenArrayOutput `pulumi:"childrens"`
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
	// Indicates if this rule is the default one
	Default pulumi.BoolPtrOutput `pulumi:"default"`
	// Identity source name from the identity stores
	IdentitySourceName pulumi.StringPtrOutput `pulumi:"identitySourceName"`
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
	// `DROP`, `CONTINUE`
	IfAuthFail pulumi.StringOutput `pulumi:"ifAuthFail"`
	// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfProcessFail pulumi.StringOutput `pulumi:"ifProcessFail"`
	// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfUserNotFound pulumi.StringOutput `pulumi:"ifUserNotFound"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringOutput `pulumi:"name"`
	// Policy set ID
	PolicySetId pulumi.StringOutput `pulumi:"policySetId"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntPtrOutput `pulumi:"rank"`
	// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
	State pulumi.StringPtrOutput `pulumi:"state"`
}

// NewAuthenticationRule registers a new resource with the given unique name, arguments, and options.
func NewAuthenticationRule(ctx *pulumi.Context,
	name string, args *AuthenticationRuleArgs, opts ...pulumi.ResourceOption) (*AuthenticationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IfAuthFail == nil {
		return nil, errors.New("invalid value for required argument 'IfAuthFail'")
	}
	if args.IfProcessFail == nil {
		return nil, errors.New("invalid value for required argument 'IfProcessFail'")
	}
	if args.IfUserNotFound == nil {
		return nil, errors.New("invalid value for required argument 'IfUserNotFound'")
	}
	if args.PolicySetId == nil {
		return nil, errors.New("invalid value for required argument 'PolicySetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthenticationRule
	err := ctx.RegisterResource("ise:deviceadmin/authenticationRule:AuthenticationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthenticationRule gets an existing AuthenticationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthenticationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthenticationRuleState, opts ...pulumi.ResourceOption) (*AuthenticationRule, error) {
	var resource AuthenticationRule
	err := ctx.ReadResource("ise:deviceadmin/authenticationRule:AuthenticationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthenticationRule resources.
type authenticationRuleState struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []AuthenticationRuleChildren `pulumi:"childrens"`
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
	// Indicates if this rule is the default one
	Default *bool `pulumi:"default"`
	// Identity source name from the identity stores
	IdentitySourceName *string `pulumi:"identitySourceName"`
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
	// `DROP`, `CONTINUE`
	IfAuthFail *string `pulumi:"ifAuthFail"`
	// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfProcessFail *string `pulumi:"ifProcessFail"`
	// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfUserNotFound *string `pulumi:"ifUserNotFound"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// Policy set ID
	PolicySetId *string `pulumi:"policySetId"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank *int `pulumi:"rank"`
	// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
	State *string `pulumi:"state"`
}

type AuthenticationRuleState struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens AuthenticationRuleChildrenArrayInput
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
	// Indicates if this rule is the default one
	Default pulumi.BoolPtrInput
	// Identity source name from the identity stores
	IdentitySourceName pulumi.StringPtrInput
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
	// `DROP`, `CONTINUE`
	IfAuthFail pulumi.StringPtrInput
	// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfProcessFail pulumi.StringPtrInput
	// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfUserNotFound pulumi.StringPtrInput
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput
	// Policy set ID
	PolicySetId pulumi.StringPtrInput
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntPtrInput
	// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
	State pulumi.StringPtrInput
}

func (AuthenticationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationRuleState)(nil)).Elem()
}

type authenticationRuleArgs struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens []AuthenticationRuleChildren `pulumi:"childrens"`
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
	// Indicates if this rule is the default one
	Default *bool `pulumi:"default"`
	// Identity source name from the identity stores
	IdentitySourceName *string `pulumi:"identitySourceName"`
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
	// `DROP`, `CONTINUE`
	IfAuthFail string `pulumi:"ifAuthFail"`
	// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfProcessFail string `pulumi:"ifProcessFail"`
	// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfUserNotFound string `pulumi:"ifUserNotFound"`
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name *string `pulumi:"name"`
	// Policy set ID
	PolicySetId string `pulumi:"policySetId"`
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank *int `pulumi:"rank"`
	// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
	State *string `pulumi:"state"`
}

// The set of arguments for constructing a AuthenticationRule resource.
type AuthenticationRuleArgs struct {
	// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
	Childrens AuthenticationRuleChildrenArrayInput
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
	// Indicates if this rule is the default one
	Default pulumi.BoolPtrInput
	// Identity source name from the identity stores
	IdentitySourceName pulumi.StringPtrInput
	// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
	// `DROP`, `CONTINUE`
	IfAuthFail pulumi.StringInput
	// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfProcessFail pulumi.StringInput
	// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
	IfUserNotFound pulumi.StringInput
	// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
	Name pulumi.StringPtrInput
	// Policy set ID
	PolicySetId pulumi.StringInput
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntPtrInput
	// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
	State pulumi.StringPtrInput
}

func (AuthenticationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authenticationRuleArgs)(nil)).Elem()
}

type AuthenticationRuleInput interface {
	pulumi.Input

	ToAuthenticationRuleOutput() AuthenticationRuleOutput
	ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput
}

func (*AuthenticationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationRule)(nil)).Elem()
}

func (i *AuthenticationRule) ToAuthenticationRuleOutput() AuthenticationRuleOutput {
	return i.ToAuthenticationRuleOutputWithContext(context.Background())
}

func (i *AuthenticationRule) ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleOutput)
}

// AuthenticationRuleArrayInput is an input type that accepts AuthenticationRuleArray and AuthenticationRuleArrayOutput values.
// You can construct a concrete instance of `AuthenticationRuleArrayInput` via:
//
//	AuthenticationRuleArray{ AuthenticationRuleArgs{...} }
type AuthenticationRuleArrayInput interface {
	pulumi.Input

	ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput
	ToAuthenticationRuleArrayOutputWithContext(context.Context) AuthenticationRuleArrayOutput
}

type AuthenticationRuleArray []AuthenticationRuleInput

func (AuthenticationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationRule)(nil)).Elem()
}

func (i AuthenticationRuleArray) ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput {
	return i.ToAuthenticationRuleArrayOutputWithContext(context.Background())
}

func (i AuthenticationRuleArray) ToAuthenticationRuleArrayOutputWithContext(ctx context.Context) AuthenticationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleArrayOutput)
}

// AuthenticationRuleMapInput is an input type that accepts AuthenticationRuleMap and AuthenticationRuleMapOutput values.
// You can construct a concrete instance of `AuthenticationRuleMapInput` via:
//
//	AuthenticationRuleMap{ "key": AuthenticationRuleArgs{...} }
type AuthenticationRuleMapInput interface {
	pulumi.Input

	ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput
	ToAuthenticationRuleMapOutputWithContext(context.Context) AuthenticationRuleMapOutput
}

type AuthenticationRuleMap map[string]AuthenticationRuleInput

func (AuthenticationRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationRule)(nil)).Elem()
}

func (i AuthenticationRuleMap) ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput {
	return i.ToAuthenticationRuleMapOutputWithContext(context.Background())
}

func (i AuthenticationRuleMap) ToAuthenticationRuleMapOutputWithContext(ctx context.Context) AuthenticationRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthenticationRuleMapOutput)
}

type AuthenticationRuleOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleOutput) ToAuthenticationRuleOutput() AuthenticationRuleOutput {
	return o
}

func (o AuthenticationRuleOutput) ToAuthenticationRuleOutputWithContext(ctx context.Context) AuthenticationRuleOutput {
	return o
}

// List of child conditions. `conditionType` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
func (o AuthenticationRuleOutput) Childrens() AuthenticationRuleChildrenArrayOutput {
	return o.ApplyT(func(v *AuthenticationRule) AuthenticationRuleChildrenArrayOutput { return v.Childrens }).(AuthenticationRuleChildrenArrayOutput)
}

// Dictionary attribute name
func (o AuthenticationRuleOutput) ConditionAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionAttributeName }).(pulumi.StringPtrOutput)
}

// Attribute value for condition. Value type is specified in dictionary object.
func (o AuthenticationRuleOutput) ConditionAttributeValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionAttributeValue }).(pulumi.StringPtrOutput)
}

// Dictionary name
func (o AuthenticationRuleOutput) ConditionDictionaryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionDictionaryName }).(pulumi.StringPtrOutput)
}

// Dictionary value
func (o AuthenticationRuleOutput) ConditionDictionaryValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionDictionaryValue }).(pulumi.StringPtrOutput)
}

// UUID for condition
func (o AuthenticationRuleOutput) ConditionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionId }).(pulumi.StringPtrOutput)
}

// Indicates whereas this condition is in negate mode
func (o AuthenticationRuleOutput) ConditionIsNegate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.BoolPtrOutput { return v.ConditionIsNegate }).(pulumi.BoolPtrOutput)
}

// Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
// `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
// `notEquals`, `notIn`, `notStartsWith`, `startsWith`
func (o AuthenticationRuleOutput) ConditionOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionOperator }).(pulumi.StringPtrOutput)
}

// Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
// additional conditions are present under the children attribute. - Choices: `ConditionAndBlock`, `ConditionAttributes`,
// `ConditionOrBlock`, `ConditionReference`
func (o AuthenticationRuleOutput) ConditionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.ConditionType }).(pulumi.StringPtrOutput)
}

// Indicates if this rule is the default one
func (o AuthenticationRuleOutput) Default() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.BoolPtrOutput { return v.Default }).(pulumi.BoolPtrOutput)
}

// Identity source name from the identity stores
func (o AuthenticationRuleOutput) IdentitySourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.IdentitySourceName }).(pulumi.StringPtrOutput)
}

// Action to perform when authentication fails such as Bad credentials, disabled user and so on - Choices: `REJECT`,
// `DROP`, `CONTINUE`
func (o AuthenticationRuleOutput) IfAuthFail() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.IfAuthFail }).(pulumi.StringOutput)
}

// Action to perform when ISE is unable to access the identity database - Choices: `REJECT`, `DROP`, `CONTINUE`
func (o AuthenticationRuleOutput) IfProcessFail() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.IfProcessFail }).(pulumi.StringOutput)
}

// Action to perform when user is not found in any of identity stores - Choices: `REJECT`, `DROP`, `CONTINUE`
func (o AuthenticationRuleOutput) IfUserNotFound() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.IfUserNotFound }).(pulumi.StringOutput)
}

// Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
func (o AuthenticationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Policy set ID
func (o AuthenticationRuleOutput) PolicySetId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringOutput { return v.PolicySetId }).(pulumi.StringOutput)
}

// The rank (priority) in relation to other rules. Lower rank is higher priority.
func (o AuthenticationRuleOutput) Rank() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.IntPtrOutput { return v.Rank }).(pulumi.IntPtrOutput)
}

// The state that the rule is in. A disabled rule cannot be matched. - Choices: `disabled`, `enabled`, `monitor`
func (o AuthenticationRuleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AuthenticationRule) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

type AuthenticationRuleArrayOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleArrayOutput) ToAuthenticationRuleArrayOutput() AuthenticationRuleArrayOutput {
	return o
}

func (o AuthenticationRuleArrayOutput) ToAuthenticationRuleArrayOutputWithContext(ctx context.Context) AuthenticationRuleArrayOutput {
	return o
}

func (o AuthenticationRuleArrayOutput) Index(i pulumi.IntInput) AuthenticationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthenticationRule {
		return vs[0].([]*AuthenticationRule)[vs[1].(int)]
	}).(AuthenticationRuleOutput)
}

type AuthenticationRuleMapOutput struct{ *pulumi.OutputState }

func (AuthenticationRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthenticationRule)(nil)).Elem()
}

func (o AuthenticationRuleMapOutput) ToAuthenticationRuleMapOutput() AuthenticationRuleMapOutput {
	return o
}

func (o AuthenticationRuleMapOutput) ToAuthenticationRuleMapOutputWithContext(ctx context.Context) AuthenticationRuleMapOutput {
	return o
}

func (o AuthenticationRuleMapOutput) MapIndex(k pulumi.StringInput) AuthenticationRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthenticationRule {
		return vs[0].(map[string]*AuthenticationRule)[vs[1].(string)]
	}).(AuthenticationRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleInput)(nil)).Elem(), &AuthenticationRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleArrayInput)(nil)).Elem(), AuthenticationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationRuleMapInput)(nil)).Elem(), AuthenticationRuleMap{})
	pulumi.RegisterOutputType(AuthenticationRuleOutput{})
	pulumi.RegisterOutputType(AuthenticationRuleArrayOutput{})
	pulumi.RegisterOutputType(AuthenticationRuleMapOutput{})
}

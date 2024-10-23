// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccess

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource is used to update rank field in network access authorization global exception rule. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and networkAccessAuthorizationGlobalExceptionRule resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.
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
//			_, err := networkaccess.NewAuthorizationGlobalExceptionRuleUpdateRank(ctx, "example", &networkaccess.AuthorizationGlobalExceptionRuleUpdateRankArgs{
//				RuleId: pulumi.String("d82952cb-b901-4b09-b363-5ebf39bdbaf9"),
//				Rank:   pulumi.Int(0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type AuthorizationGlobalExceptionRuleUpdateRank struct {
	pulumi.CustomResourceState

	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntOutput `pulumi:"rank"`
	// Authorization global exception rule ID
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
}

// NewAuthorizationGlobalExceptionRuleUpdateRank registers a new resource with the given unique name, arguments, and options.
func NewAuthorizationGlobalExceptionRuleUpdateRank(ctx *pulumi.Context,
	name string, args *AuthorizationGlobalExceptionRuleUpdateRankArgs, opts ...pulumi.ResourceOption) (*AuthorizationGlobalExceptionRuleUpdateRank, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rank == nil {
		return nil, errors.New("invalid value for required argument 'Rank'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AuthorizationGlobalExceptionRuleUpdateRank
	err := ctx.RegisterResource("ise:networkaccess/authorizationGlobalExceptionRuleUpdateRank:AuthorizationGlobalExceptionRuleUpdateRank", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthorizationGlobalExceptionRuleUpdateRank gets an existing AuthorizationGlobalExceptionRuleUpdateRank resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthorizationGlobalExceptionRuleUpdateRank(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthorizationGlobalExceptionRuleUpdateRankState, opts ...pulumi.ResourceOption) (*AuthorizationGlobalExceptionRuleUpdateRank, error) {
	var resource AuthorizationGlobalExceptionRuleUpdateRank
	err := ctx.ReadResource("ise:networkaccess/authorizationGlobalExceptionRuleUpdateRank:AuthorizationGlobalExceptionRuleUpdateRank", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthorizationGlobalExceptionRuleUpdateRank resources.
type authorizationGlobalExceptionRuleUpdateRankState struct {
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank *int `pulumi:"rank"`
	// Authorization global exception rule ID
	RuleId *string `pulumi:"ruleId"`
}

type AuthorizationGlobalExceptionRuleUpdateRankState struct {
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntPtrInput
	// Authorization global exception rule ID
	RuleId pulumi.StringPtrInput
}

func (AuthorizationGlobalExceptionRuleUpdateRankState) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationGlobalExceptionRuleUpdateRankState)(nil)).Elem()
}

type authorizationGlobalExceptionRuleUpdateRankArgs struct {
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank int `pulumi:"rank"`
	// Authorization global exception rule ID
	RuleId string `pulumi:"ruleId"`
}

// The set of arguments for constructing a AuthorizationGlobalExceptionRuleUpdateRank resource.
type AuthorizationGlobalExceptionRuleUpdateRankArgs struct {
	// The rank (priority) in relation to other rules. Lower rank is higher priority.
	Rank pulumi.IntInput
	// Authorization global exception rule ID
	RuleId pulumi.StringInput
}

func (AuthorizationGlobalExceptionRuleUpdateRankArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authorizationGlobalExceptionRuleUpdateRankArgs)(nil)).Elem()
}

type AuthorizationGlobalExceptionRuleUpdateRankInput interface {
	pulumi.Input

	ToAuthorizationGlobalExceptionRuleUpdateRankOutput() AuthorizationGlobalExceptionRuleUpdateRankOutput
	ToAuthorizationGlobalExceptionRuleUpdateRankOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankOutput
}

func (*AuthorizationGlobalExceptionRuleUpdateRank) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (i *AuthorizationGlobalExceptionRuleUpdateRank) ToAuthorizationGlobalExceptionRuleUpdateRankOutput() AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return i.ToAuthorizationGlobalExceptionRuleUpdateRankOutputWithContext(context.Background())
}

func (i *AuthorizationGlobalExceptionRuleUpdateRank) ToAuthorizationGlobalExceptionRuleUpdateRankOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationGlobalExceptionRuleUpdateRankOutput)
}

// AuthorizationGlobalExceptionRuleUpdateRankArrayInput is an input type that accepts AuthorizationGlobalExceptionRuleUpdateRankArray and AuthorizationGlobalExceptionRuleUpdateRankArrayOutput values.
// You can construct a concrete instance of `AuthorizationGlobalExceptionRuleUpdateRankArrayInput` via:
//
//	AuthorizationGlobalExceptionRuleUpdateRankArray{ AuthorizationGlobalExceptionRuleUpdateRankArgs{...} }
type AuthorizationGlobalExceptionRuleUpdateRankArrayInput interface {
	pulumi.Input

	ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutput() AuthorizationGlobalExceptionRuleUpdateRankArrayOutput
	ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutputWithContext(context.Context) AuthorizationGlobalExceptionRuleUpdateRankArrayOutput
}

type AuthorizationGlobalExceptionRuleUpdateRankArray []AuthorizationGlobalExceptionRuleUpdateRankInput

func (AuthorizationGlobalExceptionRuleUpdateRankArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (i AuthorizationGlobalExceptionRuleUpdateRankArray) ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutput() AuthorizationGlobalExceptionRuleUpdateRankArrayOutput {
	return i.ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutputWithContext(context.Background())
}

func (i AuthorizationGlobalExceptionRuleUpdateRankArray) ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationGlobalExceptionRuleUpdateRankArrayOutput)
}

// AuthorizationGlobalExceptionRuleUpdateRankMapInput is an input type that accepts AuthorizationGlobalExceptionRuleUpdateRankMap and AuthorizationGlobalExceptionRuleUpdateRankMapOutput values.
// You can construct a concrete instance of `AuthorizationGlobalExceptionRuleUpdateRankMapInput` via:
//
//	AuthorizationGlobalExceptionRuleUpdateRankMap{ "key": AuthorizationGlobalExceptionRuleUpdateRankArgs{...} }
type AuthorizationGlobalExceptionRuleUpdateRankMapInput interface {
	pulumi.Input

	ToAuthorizationGlobalExceptionRuleUpdateRankMapOutput() AuthorizationGlobalExceptionRuleUpdateRankMapOutput
	ToAuthorizationGlobalExceptionRuleUpdateRankMapOutputWithContext(context.Context) AuthorizationGlobalExceptionRuleUpdateRankMapOutput
}

type AuthorizationGlobalExceptionRuleUpdateRankMap map[string]AuthorizationGlobalExceptionRuleUpdateRankInput

func (AuthorizationGlobalExceptionRuleUpdateRankMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (i AuthorizationGlobalExceptionRuleUpdateRankMap) ToAuthorizationGlobalExceptionRuleUpdateRankMapOutput() AuthorizationGlobalExceptionRuleUpdateRankMapOutput {
	return i.ToAuthorizationGlobalExceptionRuleUpdateRankMapOutputWithContext(context.Background())
}

func (i AuthorizationGlobalExceptionRuleUpdateRankMap) ToAuthorizationGlobalExceptionRuleUpdateRankMapOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthorizationGlobalExceptionRuleUpdateRankMapOutput)
}

type AuthorizationGlobalExceptionRuleUpdateRankOutput struct{ *pulumi.OutputState }

func (AuthorizationGlobalExceptionRuleUpdateRankOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (o AuthorizationGlobalExceptionRuleUpdateRankOutput) ToAuthorizationGlobalExceptionRuleUpdateRankOutput() AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return o
}

func (o AuthorizationGlobalExceptionRuleUpdateRankOutput) ToAuthorizationGlobalExceptionRuleUpdateRankOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return o
}

// The rank (priority) in relation to other rules. Lower rank is higher priority.
func (o AuthorizationGlobalExceptionRuleUpdateRankOutput) Rank() pulumi.IntOutput {
	return o.ApplyT(func(v *AuthorizationGlobalExceptionRuleUpdateRank) pulumi.IntOutput { return v.Rank }).(pulumi.IntOutput)
}

// Authorization global exception rule ID
func (o AuthorizationGlobalExceptionRuleUpdateRankOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *AuthorizationGlobalExceptionRuleUpdateRank) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

type AuthorizationGlobalExceptionRuleUpdateRankArrayOutput struct{ *pulumi.OutputState }

func (AuthorizationGlobalExceptionRuleUpdateRankArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (o AuthorizationGlobalExceptionRuleUpdateRankArrayOutput) ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutput() AuthorizationGlobalExceptionRuleUpdateRankArrayOutput {
	return o
}

func (o AuthorizationGlobalExceptionRuleUpdateRankArrayOutput) ToAuthorizationGlobalExceptionRuleUpdateRankArrayOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankArrayOutput {
	return o
}

func (o AuthorizationGlobalExceptionRuleUpdateRankArrayOutput) Index(i pulumi.IntInput) AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthorizationGlobalExceptionRuleUpdateRank {
		return vs[0].([]*AuthorizationGlobalExceptionRuleUpdateRank)[vs[1].(int)]
	}).(AuthorizationGlobalExceptionRuleUpdateRankOutput)
}

type AuthorizationGlobalExceptionRuleUpdateRankMapOutput struct{ *pulumi.OutputState }

func (AuthorizationGlobalExceptionRuleUpdateRankMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthorizationGlobalExceptionRuleUpdateRank)(nil)).Elem()
}

func (o AuthorizationGlobalExceptionRuleUpdateRankMapOutput) ToAuthorizationGlobalExceptionRuleUpdateRankMapOutput() AuthorizationGlobalExceptionRuleUpdateRankMapOutput {
	return o
}

func (o AuthorizationGlobalExceptionRuleUpdateRankMapOutput) ToAuthorizationGlobalExceptionRuleUpdateRankMapOutputWithContext(ctx context.Context) AuthorizationGlobalExceptionRuleUpdateRankMapOutput {
	return o
}

func (o AuthorizationGlobalExceptionRuleUpdateRankMapOutput) MapIndex(k pulumi.StringInput) AuthorizationGlobalExceptionRuleUpdateRankOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthorizationGlobalExceptionRuleUpdateRank {
		return vs[0].(map[string]*AuthorizationGlobalExceptionRuleUpdateRank)[vs[1].(string)]
	}).(AuthorizationGlobalExceptionRuleUpdateRankOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationGlobalExceptionRuleUpdateRankInput)(nil)).Elem(), &AuthorizationGlobalExceptionRuleUpdateRank{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationGlobalExceptionRuleUpdateRankArrayInput)(nil)).Elem(), AuthorizationGlobalExceptionRuleUpdateRankArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthorizationGlobalExceptionRuleUpdateRankMapInput)(nil)).Elem(), AuthorizationGlobalExceptionRuleUpdateRankMap{})
	pulumi.RegisterOutputType(AuthorizationGlobalExceptionRuleUpdateRankOutput{})
	pulumi.RegisterOutputType(AuthorizationGlobalExceptionRuleUpdateRankArrayOutput{})
	pulumi.RegisterOutputType(AuthorizationGlobalExceptionRuleUpdateRankMapOutput{})
}

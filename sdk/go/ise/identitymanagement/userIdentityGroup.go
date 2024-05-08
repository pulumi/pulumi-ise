// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage an User Identity Group.
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
//			_, err := identitymanagement.NewUserIdentityGroup(ctx, "example", &identitymanagement.UserIdentityGroupArgs{
//				Name:        pulumi.String("Group1"),
//				Description: pulumi.String("My endpoint identity group"),
//				Parent:      pulumi.String("NAC Group:NAC:IdentityGroups:User Identity Groups"),
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
// $ pulumi import ise:identitymanagement/userIdentityGroup:UserIdentityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type UserIdentityGroup struct {
	pulumi.CustomResourceState

	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the user identity group
	Name pulumi.StringOutput `pulumi:"name"`
	// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
}

// NewUserIdentityGroup registers a new resource with the given unique name, arguments, and options.
func NewUserIdentityGroup(ctx *pulumi.Context,
	name string, args *UserIdentityGroupArgs, opts ...pulumi.ResourceOption) (*UserIdentityGroup, error) {
	if args == nil {
		args = &UserIdentityGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserIdentityGroup
	err := ctx.RegisterResource("ise:identitymanagement/userIdentityGroup:UserIdentityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserIdentityGroup gets an existing UserIdentityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserIdentityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserIdentityGroupState, opts ...pulumi.ResourceOption) (*UserIdentityGroup, error) {
	var resource UserIdentityGroup
	err := ctx.ReadResource("ise:identitymanagement/userIdentityGroup:UserIdentityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserIdentityGroup resources.
type userIdentityGroupState struct {
	// Description
	Description *string `pulumi:"description"`
	// The name of the user identity group
	Name *string `pulumi:"name"`
	// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
	Parent *string `pulumi:"parent"`
}

type UserIdentityGroupState struct {
	// Description
	Description pulumi.StringPtrInput
	// The name of the user identity group
	Name pulumi.StringPtrInput
	// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
	Parent pulumi.StringPtrInput
}

func (UserIdentityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*userIdentityGroupState)(nil)).Elem()
}

type userIdentityGroupArgs struct {
	// Description
	Description *string `pulumi:"description"`
	// The name of the user identity group
	Name *string `pulumi:"name"`
	// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
	Parent *string `pulumi:"parent"`
}

// The set of arguments for constructing a UserIdentityGroup resource.
type UserIdentityGroupArgs struct {
	// Description
	Description pulumi.StringPtrInput
	// The name of the user identity group
	Name pulumi.StringPtrInput
	// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
	Parent pulumi.StringPtrInput
}

func (UserIdentityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userIdentityGroupArgs)(nil)).Elem()
}

type UserIdentityGroupInput interface {
	pulumi.Input

	ToUserIdentityGroupOutput() UserIdentityGroupOutput
	ToUserIdentityGroupOutputWithContext(ctx context.Context) UserIdentityGroupOutput
}

func (*UserIdentityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentityGroup)(nil)).Elem()
}

func (i *UserIdentityGroup) ToUserIdentityGroupOutput() UserIdentityGroupOutput {
	return i.ToUserIdentityGroupOutputWithContext(context.Background())
}

func (i *UserIdentityGroup) ToUserIdentityGroupOutputWithContext(ctx context.Context) UserIdentityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityGroupOutput)
}

// UserIdentityGroupArrayInput is an input type that accepts UserIdentityGroupArray and UserIdentityGroupArrayOutput values.
// You can construct a concrete instance of `UserIdentityGroupArrayInput` via:
//
//	UserIdentityGroupArray{ UserIdentityGroupArgs{...} }
type UserIdentityGroupArrayInput interface {
	pulumi.Input

	ToUserIdentityGroupArrayOutput() UserIdentityGroupArrayOutput
	ToUserIdentityGroupArrayOutputWithContext(context.Context) UserIdentityGroupArrayOutput
}

type UserIdentityGroupArray []UserIdentityGroupInput

func (UserIdentityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserIdentityGroup)(nil)).Elem()
}

func (i UserIdentityGroupArray) ToUserIdentityGroupArrayOutput() UserIdentityGroupArrayOutput {
	return i.ToUserIdentityGroupArrayOutputWithContext(context.Background())
}

func (i UserIdentityGroupArray) ToUserIdentityGroupArrayOutputWithContext(ctx context.Context) UserIdentityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityGroupArrayOutput)
}

// UserIdentityGroupMapInput is an input type that accepts UserIdentityGroupMap and UserIdentityGroupMapOutput values.
// You can construct a concrete instance of `UserIdentityGroupMapInput` via:
//
//	UserIdentityGroupMap{ "key": UserIdentityGroupArgs{...} }
type UserIdentityGroupMapInput interface {
	pulumi.Input

	ToUserIdentityGroupMapOutput() UserIdentityGroupMapOutput
	ToUserIdentityGroupMapOutputWithContext(context.Context) UserIdentityGroupMapOutput
}

type UserIdentityGroupMap map[string]UserIdentityGroupInput

func (UserIdentityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserIdentityGroup)(nil)).Elem()
}

func (i UserIdentityGroupMap) ToUserIdentityGroupMapOutput() UserIdentityGroupMapOutput {
	return i.ToUserIdentityGroupMapOutputWithContext(context.Background())
}

func (i UserIdentityGroupMap) ToUserIdentityGroupMapOutputWithContext(ctx context.Context) UserIdentityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityGroupMapOutput)
}

type UserIdentityGroupOutput struct{ *pulumi.OutputState }

func (UserIdentityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserIdentityGroup)(nil)).Elem()
}

func (o UserIdentityGroupOutput) ToUserIdentityGroupOutput() UserIdentityGroupOutput {
	return o
}

func (o UserIdentityGroupOutput) ToUserIdentityGroupOutputWithContext(ctx context.Context) UserIdentityGroupOutput {
	return o
}

// Description
func (o UserIdentityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the user identity group
func (o UserIdentityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserIdentityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parent user identity group, e.g. `NAC Group:NAC:IdentityGroups:User Identity Groups`
func (o UserIdentityGroupOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserIdentityGroup) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

type UserIdentityGroupArrayOutput struct{ *pulumi.OutputState }

func (UserIdentityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserIdentityGroup)(nil)).Elem()
}

func (o UserIdentityGroupArrayOutput) ToUserIdentityGroupArrayOutput() UserIdentityGroupArrayOutput {
	return o
}

func (o UserIdentityGroupArrayOutput) ToUserIdentityGroupArrayOutputWithContext(ctx context.Context) UserIdentityGroupArrayOutput {
	return o
}

func (o UserIdentityGroupArrayOutput) Index(i pulumi.IntInput) UserIdentityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserIdentityGroup {
		return vs[0].([]*UserIdentityGroup)[vs[1].(int)]
	}).(UserIdentityGroupOutput)
}

type UserIdentityGroupMapOutput struct{ *pulumi.OutputState }

func (UserIdentityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserIdentityGroup)(nil)).Elem()
}

func (o UserIdentityGroupMapOutput) ToUserIdentityGroupMapOutput() UserIdentityGroupMapOutput {
	return o
}

func (o UserIdentityGroupMapOutput) ToUserIdentityGroupMapOutputWithContext(ctx context.Context) UserIdentityGroupMapOutput {
	return o
}

func (o UserIdentityGroupMapOutput) MapIndex(k pulumi.StringInput) UserIdentityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserIdentityGroup {
		return vs[0].(map[string]*UserIdentityGroup)[vs[1].(string)]
	}).(UserIdentityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityGroupInput)(nil)).Elem(), &UserIdentityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityGroupArrayInput)(nil)).Elem(), UserIdentityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserIdentityGroupMapInput)(nil)).Elem(), UserIdentityGroupMap{})
	pulumi.RegisterOutputType(UserIdentityGroupOutput{})
	pulumi.RegisterOutputType(UserIdentityGroupArrayOutput{})
	pulumi.RegisterOutputType(UserIdentityGroupMapOutput{})
}

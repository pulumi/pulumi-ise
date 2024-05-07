// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package activedirectoryadd

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage an Active Directory Add Groups.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/ActiveDirectoryAdd"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ActiveDirectoryAdd.NewGroups(ctx, "example", &ActiveDirectoryAdd.GroupsArgs{
//				JoinPointId:             pulumi.String("73808580-b6e6-11ee-8960-de6d7692bc40"),
//				Name:                    pulumi.String("cisco.local"),
//				Description:             pulumi.String("My AD join point"),
//				Domain:                  pulumi.String("cisco.local"),
//				AdScopesNames:           pulumi.String("Default_Scope"),
//				EnableDomainAllowedList: pulumi.Bool(true),
//				Groups: activedirectoryadd.GroupsGroupArray{
//					&activedirectoryadd.GroupsGroupArgs{
//						Name: pulumi.String("cisco.local/operators"),
//						Sid:  pulumi.String("S-1-5-32-548"),
//						Type: pulumi.String("GLOBAL"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Groups struct {
	pulumi.CustomResourceState

	// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
	// value: `Default_Scope`
	AdScopesNames pulumi.StringOutput `pulumi:"adScopesNames"`
	// Join point Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// AD domain associated with the join point
	Domain pulumi.StringOutput `pulumi:"domain"`
	// - Default value: `true`
	EnableDomainAllowedList pulumi.BoolOutput `pulumi:"enableDomainAllowedList"`
	// List of AD Groups
	Groups GroupsGroupArrayOutput `pulumi:"groups"`
	// Active Directory Join Point ID
	JoinPointId pulumi.StringOutput `pulumi:"joinPointId"`
	// The name of the active directory join point
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewGroups registers a new resource with the given unique name, arguments, and options.
func NewGroups(ctx *pulumi.Context,
	name string, args *GroupsArgs, opts ...pulumi.ResourceOption) (*Groups, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.JoinPointId == nil {
		return nil, errors.New("invalid value for required argument 'JoinPointId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Groups
	err := ctx.RegisterResource("ise:ActiveDirectoryAdd/groups:Groups", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroups gets an existing Groups resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroups(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupsState, opts ...pulumi.ResourceOption) (*Groups, error) {
	var resource Groups
	err := ctx.ReadResource("ise:ActiveDirectoryAdd/groups:Groups", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Groups resources.
type groupsState struct {
	// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
	// value: `Default_Scope`
	AdScopesNames *string `pulumi:"adScopesNames"`
	// Join point Description
	Description *string `pulumi:"description"`
	// AD domain associated with the join point
	Domain *string `pulumi:"domain"`
	// - Default value: `true`
	EnableDomainAllowedList *bool `pulumi:"enableDomainAllowedList"`
	// List of AD Groups
	Groups []GroupsGroup `pulumi:"groups"`
	// Active Directory Join Point ID
	JoinPointId *string `pulumi:"joinPointId"`
	// The name of the active directory join point
	Name *string `pulumi:"name"`
}

type GroupsState struct {
	// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
	// value: `Default_Scope`
	AdScopesNames pulumi.StringPtrInput
	// Join point Description
	Description pulumi.StringPtrInput
	// AD domain associated with the join point
	Domain pulumi.StringPtrInput
	// - Default value: `true`
	EnableDomainAllowedList pulumi.BoolPtrInput
	// List of AD Groups
	Groups GroupsGroupArrayInput
	// Active Directory Join Point ID
	JoinPointId pulumi.StringPtrInput
	// The name of the active directory join point
	Name pulumi.StringPtrInput
}

func (GroupsState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsState)(nil)).Elem()
}

type groupsArgs struct {
	// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
	// value: `Default_Scope`
	AdScopesNames *string `pulumi:"adScopesNames"`
	// Join point Description
	Description *string `pulumi:"description"`
	// AD domain associated with the join point
	Domain string `pulumi:"domain"`
	// - Default value: `true`
	EnableDomainAllowedList *bool `pulumi:"enableDomainAllowedList"`
	// List of AD Groups
	Groups []GroupsGroup `pulumi:"groups"`
	// Active Directory Join Point ID
	JoinPointId string `pulumi:"joinPointId"`
	// The name of the active directory join point
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Groups resource.
type GroupsArgs struct {
	// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
	// value: `Default_Scope`
	AdScopesNames pulumi.StringPtrInput
	// Join point Description
	Description pulumi.StringPtrInput
	// AD domain associated with the join point
	Domain pulumi.StringInput
	// - Default value: `true`
	EnableDomainAllowedList pulumi.BoolPtrInput
	// List of AD Groups
	Groups GroupsGroupArrayInput
	// Active Directory Join Point ID
	JoinPointId pulumi.StringInput
	// The name of the active directory join point
	Name pulumi.StringPtrInput
}

func (GroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupsArgs)(nil)).Elem()
}

type GroupsInput interface {
	pulumi.Input

	ToGroupsOutput() GroupsOutput
	ToGroupsOutputWithContext(ctx context.Context) GroupsOutput
}

func (*Groups) ElementType() reflect.Type {
	return reflect.TypeOf((**Groups)(nil)).Elem()
}

func (i *Groups) ToGroupsOutput() GroupsOutput {
	return i.ToGroupsOutputWithContext(context.Background())
}

func (i *Groups) ToGroupsOutputWithContext(ctx context.Context) GroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupsOutput)
}

// GroupsArrayInput is an input type that accepts GroupsArray and GroupsArrayOutput values.
// You can construct a concrete instance of `GroupsArrayInput` via:
//
//	GroupsArray{ GroupsArgs{...} }
type GroupsArrayInput interface {
	pulumi.Input

	ToGroupsArrayOutput() GroupsArrayOutput
	ToGroupsArrayOutputWithContext(context.Context) GroupsArrayOutput
}

type GroupsArray []GroupsInput

func (GroupsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Groups)(nil)).Elem()
}

func (i GroupsArray) ToGroupsArrayOutput() GroupsArrayOutput {
	return i.ToGroupsArrayOutputWithContext(context.Background())
}

func (i GroupsArray) ToGroupsArrayOutputWithContext(ctx context.Context) GroupsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupsArrayOutput)
}

// GroupsMapInput is an input type that accepts GroupsMap and GroupsMapOutput values.
// You can construct a concrete instance of `GroupsMapInput` via:
//
//	GroupsMap{ "key": GroupsArgs{...} }
type GroupsMapInput interface {
	pulumi.Input

	ToGroupsMapOutput() GroupsMapOutput
	ToGroupsMapOutputWithContext(context.Context) GroupsMapOutput
}

type GroupsMap map[string]GroupsInput

func (GroupsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Groups)(nil)).Elem()
}

func (i GroupsMap) ToGroupsMapOutput() GroupsMapOutput {
	return i.ToGroupsMapOutputWithContext(context.Background())
}

func (i GroupsMap) ToGroupsMapOutputWithContext(ctx context.Context) GroupsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupsMapOutput)
}

type GroupsOutput struct{ *pulumi.OutputState }

func (GroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Groups)(nil)).Elem()
}

func (o GroupsOutput) ToGroupsOutput() GroupsOutput {
	return o
}

func (o GroupsOutput) ToGroupsOutputWithContext(ctx context.Context) GroupsOutput {
	return o
}

// String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
// value: `Default_Scope`
func (o GroupsOutput) AdScopesNames() pulumi.StringOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringOutput { return v.AdScopesNames }).(pulumi.StringOutput)
}

// Join point Description
func (o GroupsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// AD domain associated with the join point
func (o GroupsOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// - Default value: `true`
func (o GroupsOutput) EnableDomainAllowedList() pulumi.BoolOutput {
	return o.ApplyT(func(v *Groups) pulumi.BoolOutput { return v.EnableDomainAllowedList }).(pulumi.BoolOutput)
}

// List of AD Groups
func (o GroupsOutput) Groups() GroupsGroupArrayOutput {
	return o.ApplyT(func(v *Groups) GroupsGroupArrayOutput { return v.Groups }).(GroupsGroupArrayOutput)
}

// Active Directory Join Point ID
func (o GroupsOutput) JoinPointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringOutput { return v.JoinPointId }).(pulumi.StringOutput)
}

// The name of the active directory join point
func (o GroupsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Groups) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type GroupsArrayOutput struct{ *pulumi.OutputState }

func (GroupsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Groups)(nil)).Elem()
}

func (o GroupsArrayOutput) ToGroupsArrayOutput() GroupsArrayOutput {
	return o
}

func (o GroupsArrayOutput) ToGroupsArrayOutputWithContext(ctx context.Context) GroupsArrayOutput {
	return o
}

func (o GroupsArrayOutput) Index(i pulumi.IntInput) GroupsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Groups {
		return vs[0].([]*Groups)[vs[1].(int)]
	}).(GroupsOutput)
}

type GroupsMapOutput struct{ *pulumi.OutputState }

func (GroupsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Groups)(nil)).Elem()
}

func (o GroupsMapOutput) ToGroupsMapOutput() GroupsMapOutput {
	return o
}

func (o GroupsMapOutput) ToGroupsMapOutputWithContext(ctx context.Context) GroupsMapOutput {
	return o
}

func (o GroupsMapOutput) MapIndex(k pulumi.StringInput) GroupsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Groups {
		return vs[0].(map[string]*Groups)[vs[1].(string)]
	}).(GroupsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GroupsInput)(nil)).Elem(), &Groups{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupsArrayInput)(nil)).Elem(), GroupsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupsMapInput)(nil)).Elem(), GroupsMap{})
	pulumi.RegisterOutputType(GroupsOutput{})
	pulumi.RegisterOutputType(GroupsArrayOutput{})
	pulumi.RegisterOutputType(GroupsMapOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identitymanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage an Endpoint Identity Group.
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
//			_, err := identitymanagement.NewEndpointIdentityGroup(ctx, "example", &identitymanagement.EndpointIdentityGroupArgs{
//				Name:          pulumi.String("Group1"),
//				Description:   pulumi.String("My endpoint identity group"),
//				SystemDefined: pulumi.Bool(false),
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
// $ pulumi import ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type EndpointIdentityGroup struct {
	pulumi.CustomResourceState

	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the endpoint identity group
	Name pulumi.StringOutput `pulumi:"name"`
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId pulumi.StringPtrOutput `pulumi:"parentEndpointIdentityGroupId"`
	// System defined endpoint identity group
	SystemDefined pulumi.BoolPtrOutput `pulumi:"systemDefined"`
}

// NewEndpointIdentityGroup registers a new resource with the given unique name, arguments, and options.
func NewEndpointIdentityGroup(ctx *pulumi.Context,
	name string, args *EndpointIdentityGroupArgs, opts ...pulumi.ResourceOption) (*EndpointIdentityGroup, error) {
	if args == nil {
		args = &EndpointIdentityGroupArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EndpointIdentityGroup
	err := ctx.RegisterResource("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpointIdentityGroup gets an existing EndpointIdentityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpointIdentityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointIdentityGroupState, opts ...pulumi.ResourceOption) (*EndpointIdentityGroup, error) {
	var resource EndpointIdentityGroup
	err := ctx.ReadResource("ise:identitymanagement/endpointIdentityGroup:EndpointIdentityGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EndpointIdentityGroup resources.
type endpointIdentityGroupState struct {
	// Description
	Description *string `pulumi:"description"`
	// The name of the endpoint identity group
	Name *string `pulumi:"name"`
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId *string `pulumi:"parentEndpointIdentityGroupId"`
	// System defined endpoint identity group
	SystemDefined *bool `pulumi:"systemDefined"`
}

type EndpointIdentityGroupState struct {
	// Description
	Description pulumi.StringPtrInput
	// The name of the endpoint identity group
	Name pulumi.StringPtrInput
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId pulumi.StringPtrInput
	// System defined endpoint identity group
	SystemDefined pulumi.BoolPtrInput
}

func (EndpointIdentityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointIdentityGroupState)(nil)).Elem()
}

type endpointIdentityGroupArgs struct {
	// Description
	Description *string `pulumi:"description"`
	// The name of the endpoint identity group
	Name *string `pulumi:"name"`
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId *string `pulumi:"parentEndpointIdentityGroupId"`
	// System defined endpoint identity group
	SystemDefined *bool `pulumi:"systemDefined"`
}

// The set of arguments for constructing a EndpointIdentityGroup resource.
type EndpointIdentityGroupArgs struct {
	// Description
	Description pulumi.StringPtrInput
	// The name of the endpoint identity group
	Name pulumi.StringPtrInput
	// Parent endpoint identity group ID
	ParentEndpointIdentityGroupId pulumi.StringPtrInput
	// System defined endpoint identity group
	SystemDefined pulumi.BoolPtrInput
}

func (EndpointIdentityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointIdentityGroupArgs)(nil)).Elem()
}

type EndpointIdentityGroupInput interface {
	pulumi.Input

	ToEndpointIdentityGroupOutput() EndpointIdentityGroupOutput
	ToEndpointIdentityGroupOutputWithContext(ctx context.Context) EndpointIdentityGroupOutput
}

func (*EndpointIdentityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointIdentityGroup)(nil)).Elem()
}

func (i *EndpointIdentityGroup) ToEndpointIdentityGroupOutput() EndpointIdentityGroupOutput {
	return i.ToEndpointIdentityGroupOutputWithContext(context.Background())
}

func (i *EndpointIdentityGroup) ToEndpointIdentityGroupOutputWithContext(ctx context.Context) EndpointIdentityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointIdentityGroupOutput)
}

// EndpointIdentityGroupArrayInput is an input type that accepts EndpointIdentityGroupArray and EndpointIdentityGroupArrayOutput values.
// You can construct a concrete instance of `EndpointIdentityGroupArrayInput` via:
//
//	EndpointIdentityGroupArray{ EndpointIdentityGroupArgs{...} }
type EndpointIdentityGroupArrayInput interface {
	pulumi.Input

	ToEndpointIdentityGroupArrayOutput() EndpointIdentityGroupArrayOutput
	ToEndpointIdentityGroupArrayOutputWithContext(context.Context) EndpointIdentityGroupArrayOutput
}

type EndpointIdentityGroupArray []EndpointIdentityGroupInput

func (EndpointIdentityGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointIdentityGroup)(nil)).Elem()
}

func (i EndpointIdentityGroupArray) ToEndpointIdentityGroupArrayOutput() EndpointIdentityGroupArrayOutput {
	return i.ToEndpointIdentityGroupArrayOutputWithContext(context.Background())
}

func (i EndpointIdentityGroupArray) ToEndpointIdentityGroupArrayOutputWithContext(ctx context.Context) EndpointIdentityGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointIdentityGroupArrayOutput)
}

// EndpointIdentityGroupMapInput is an input type that accepts EndpointIdentityGroupMap and EndpointIdentityGroupMapOutput values.
// You can construct a concrete instance of `EndpointIdentityGroupMapInput` via:
//
//	EndpointIdentityGroupMap{ "key": EndpointIdentityGroupArgs{...} }
type EndpointIdentityGroupMapInput interface {
	pulumi.Input

	ToEndpointIdentityGroupMapOutput() EndpointIdentityGroupMapOutput
	ToEndpointIdentityGroupMapOutputWithContext(context.Context) EndpointIdentityGroupMapOutput
}

type EndpointIdentityGroupMap map[string]EndpointIdentityGroupInput

func (EndpointIdentityGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointIdentityGroup)(nil)).Elem()
}

func (i EndpointIdentityGroupMap) ToEndpointIdentityGroupMapOutput() EndpointIdentityGroupMapOutput {
	return i.ToEndpointIdentityGroupMapOutputWithContext(context.Background())
}

func (i EndpointIdentityGroupMap) ToEndpointIdentityGroupMapOutputWithContext(ctx context.Context) EndpointIdentityGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointIdentityGroupMapOutput)
}

type EndpointIdentityGroupOutput struct{ *pulumi.OutputState }

func (EndpointIdentityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointIdentityGroup)(nil)).Elem()
}

func (o EndpointIdentityGroupOutput) ToEndpointIdentityGroupOutput() EndpointIdentityGroupOutput {
	return o
}

func (o EndpointIdentityGroupOutput) ToEndpointIdentityGroupOutputWithContext(ctx context.Context) EndpointIdentityGroupOutput {
	return o
}

// Description
func (o EndpointIdentityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointIdentityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the endpoint identity group
func (o EndpointIdentityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EndpointIdentityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parent endpoint identity group ID
func (o EndpointIdentityGroupOutput) ParentEndpointIdentityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointIdentityGroup) pulumi.StringPtrOutput { return v.ParentEndpointIdentityGroupId }).(pulumi.StringPtrOutput)
}

// System defined endpoint identity group
func (o EndpointIdentityGroupOutput) SystemDefined() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EndpointIdentityGroup) pulumi.BoolPtrOutput { return v.SystemDefined }).(pulumi.BoolPtrOutput)
}

type EndpointIdentityGroupArrayOutput struct{ *pulumi.OutputState }

func (EndpointIdentityGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EndpointIdentityGroup)(nil)).Elem()
}

func (o EndpointIdentityGroupArrayOutput) ToEndpointIdentityGroupArrayOutput() EndpointIdentityGroupArrayOutput {
	return o
}

func (o EndpointIdentityGroupArrayOutput) ToEndpointIdentityGroupArrayOutputWithContext(ctx context.Context) EndpointIdentityGroupArrayOutput {
	return o
}

func (o EndpointIdentityGroupArrayOutput) Index(i pulumi.IntInput) EndpointIdentityGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EndpointIdentityGroup {
		return vs[0].([]*EndpointIdentityGroup)[vs[1].(int)]
	}).(EndpointIdentityGroupOutput)
}

type EndpointIdentityGroupMapOutput struct{ *pulumi.OutputState }

func (EndpointIdentityGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EndpointIdentityGroup)(nil)).Elem()
}

func (o EndpointIdentityGroupMapOutput) ToEndpointIdentityGroupMapOutput() EndpointIdentityGroupMapOutput {
	return o
}

func (o EndpointIdentityGroupMapOutput) ToEndpointIdentityGroupMapOutputWithContext(ctx context.Context) EndpointIdentityGroupMapOutput {
	return o
}

func (o EndpointIdentityGroupMapOutput) MapIndex(k pulumi.StringInput) EndpointIdentityGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EndpointIdentityGroup {
		return vs[0].(map[string]*EndpointIdentityGroup)[vs[1].(string)]
	}).(EndpointIdentityGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointIdentityGroupInput)(nil)).Elem(), &EndpointIdentityGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointIdentityGroupArrayInput)(nil)).Elem(), EndpointIdentityGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointIdentityGroupMapInput)(nil)).Elem(), EndpointIdentityGroupMap{})
	pulumi.RegisterOutputType(EndpointIdentityGroupOutput{})
	pulumi.RegisterOutputType(EndpointIdentityGroupArrayOutput{})
	pulumi.RegisterOutputType(EndpointIdentityGroupMapOutput{})
}

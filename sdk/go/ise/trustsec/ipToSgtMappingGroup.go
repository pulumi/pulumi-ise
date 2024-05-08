// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a TrustSec IP to SGT Mapping Group.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/trustsec"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := trustsec.NewIpToSgtMappingGroup(ctx, "example", &trustsec.IpToSgtMappingGroupArgs{
//				Name:       pulumi.String("groupA"),
//				DeployType: pulumi.String("ALL"),
//				Sgt:        pulumi.String("93e1bf00-8c01-11e6-996c-525400b48521"),
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
// $ pulumi import ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type IpToSgtMappingGroup struct {
	pulumi.CustomResourceState

	// Mandatory unless `deploy_type` is `ALL`
	DeployTo pulumi.StringPtrOutput `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringOutput `pulumi:"deployType"`
	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the IP to SGT mapping Group
	Name pulumi.StringOutput `pulumi:"name"`
	// Trustsec Security Group ID
	Sgt pulumi.StringOutput `pulumi:"sgt"`
}

// NewIpToSgtMappingGroup registers a new resource with the given unique name, arguments, and options.
func NewIpToSgtMappingGroup(ctx *pulumi.Context,
	name string, args *IpToSgtMappingGroupArgs, opts ...pulumi.ResourceOption) (*IpToSgtMappingGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeployType == nil {
		return nil, errors.New("invalid value for required argument 'DeployType'")
	}
	if args.Sgt == nil {
		return nil, errors.New("invalid value for required argument 'Sgt'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpToSgtMappingGroup
	err := ctx.RegisterResource("ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpToSgtMappingGroup gets an existing IpToSgtMappingGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpToSgtMappingGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpToSgtMappingGroupState, opts ...pulumi.ResourceOption) (*IpToSgtMappingGroup, error) {
	var resource IpToSgtMappingGroup
	err := ctx.ReadResource("ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpToSgtMappingGroup resources.
type ipToSgtMappingGroupState struct {
	// Mandatory unless `deploy_type` is `ALL`
	DeployTo *string `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType *string `pulumi:"deployType"`
	// Description
	Description *string `pulumi:"description"`
	// The name of the IP to SGT mapping Group
	Name *string `pulumi:"name"`
	// Trustsec Security Group ID
	Sgt *string `pulumi:"sgt"`
}

type IpToSgtMappingGroupState struct {
	// Mandatory unless `deploy_type` is `ALL`
	DeployTo pulumi.StringPtrInput
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringPtrInput
	// Description
	Description pulumi.StringPtrInput
	// The name of the IP to SGT mapping Group
	Name pulumi.StringPtrInput
	// Trustsec Security Group ID
	Sgt pulumi.StringPtrInput
}

func (IpToSgtMappingGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipToSgtMappingGroupState)(nil)).Elem()
}

type ipToSgtMappingGroupArgs struct {
	// Mandatory unless `deploy_type` is `ALL`
	DeployTo *string `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType string `pulumi:"deployType"`
	// Description
	Description *string `pulumi:"description"`
	// The name of the IP to SGT mapping Group
	Name *string `pulumi:"name"`
	// Trustsec Security Group ID
	Sgt string `pulumi:"sgt"`
}

// The set of arguments for constructing a IpToSgtMappingGroup resource.
type IpToSgtMappingGroupArgs struct {
	// Mandatory unless `deploy_type` is `ALL`
	DeployTo pulumi.StringPtrInput
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringInput
	// Description
	Description pulumi.StringPtrInput
	// The name of the IP to SGT mapping Group
	Name pulumi.StringPtrInput
	// Trustsec Security Group ID
	Sgt pulumi.StringInput
}

func (IpToSgtMappingGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipToSgtMappingGroupArgs)(nil)).Elem()
}

type IpToSgtMappingGroupInput interface {
	pulumi.Input

	ToIpToSgtMappingGroupOutput() IpToSgtMappingGroupOutput
	ToIpToSgtMappingGroupOutputWithContext(ctx context.Context) IpToSgtMappingGroupOutput
}

func (*IpToSgtMappingGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**IpToSgtMappingGroup)(nil)).Elem()
}

func (i *IpToSgtMappingGroup) ToIpToSgtMappingGroupOutput() IpToSgtMappingGroupOutput {
	return i.ToIpToSgtMappingGroupOutputWithContext(context.Background())
}

func (i *IpToSgtMappingGroup) ToIpToSgtMappingGroupOutputWithContext(ctx context.Context) IpToSgtMappingGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingGroupOutput)
}

// IpToSgtMappingGroupArrayInput is an input type that accepts IpToSgtMappingGroupArray and IpToSgtMappingGroupArrayOutput values.
// You can construct a concrete instance of `IpToSgtMappingGroupArrayInput` via:
//
//	IpToSgtMappingGroupArray{ IpToSgtMappingGroupArgs{...} }
type IpToSgtMappingGroupArrayInput interface {
	pulumi.Input

	ToIpToSgtMappingGroupArrayOutput() IpToSgtMappingGroupArrayOutput
	ToIpToSgtMappingGroupArrayOutputWithContext(context.Context) IpToSgtMappingGroupArrayOutput
}

type IpToSgtMappingGroupArray []IpToSgtMappingGroupInput

func (IpToSgtMappingGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpToSgtMappingGroup)(nil)).Elem()
}

func (i IpToSgtMappingGroupArray) ToIpToSgtMappingGroupArrayOutput() IpToSgtMappingGroupArrayOutput {
	return i.ToIpToSgtMappingGroupArrayOutputWithContext(context.Background())
}

func (i IpToSgtMappingGroupArray) ToIpToSgtMappingGroupArrayOutputWithContext(ctx context.Context) IpToSgtMappingGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingGroupArrayOutput)
}

// IpToSgtMappingGroupMapInput is an input type that accepts IpToSgtMappingGroupMap and IpToSgtMappingGroupMapOutput values.
// You can construct a concrete instance of `IpToSgtMappingGroupMapInput` via:
//
//	IpToSgtMappingGroupMap{ "key": IpToSgtMappingGroupArgs{...} }
type IpToSgtMappingGroupMapInput interface {
	pulumi.Input

	ToIpToSgtMappingGroupMapOutput() IpToSgtMappingGroupMapOutput
	ToIpToSgtMappingGroupMapOutputWithContext(context.Context) IpToSgtMappingGroupMapOutput
}

type IpToSgtMappingGroupMap map[string]IpToSgtMappingGroupInput

func (IpToSgtMappingGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpToSgtMappingGroup)(nil)).Elem()
}

func (i IpToSgtMappingGroupMap) ToIpToSgtMappingGroupMapOutput() IpToSgtMappingGroupMapOutput {
	return i.ToIpToSgtMappingGroupMapOutputWithContext(context.Background())
}

func (i IpToSgtMappingGroupMap) ToIpToSgtMappingGroupMapOutputWithContext(ctx context.Context) IpToSgtMappingGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingGroupMapOutput)
}

type IpToSgtMappingGroupOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpToSgtMappingGroup)(nil)).Elem()
}

func (o IpToSgtMappingGroupOutput) ToIpToSgtMappingGroupOutput() IpToSgtMappingGroupOutput {
	return o
}

func (o IpToSgtMappingGroupOutput) ToIpToSgtMappingGroupOutputWithContext(ctx context.Context) IpToSgtMappingGroupOutput {
	return o
}

// Mandatory unless `deploy_type` is `ALL`
func (o IpToSgtMappingGroupOutput) DeployTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMappingGroup) pulumi.StringPtrOutput { return v.DeployTo }).(pulumi.StringPtrOutput)
}

// Deploy Type - Choices: `ALL`, `ND`, `NDG`
func (o IpToSgtMappingGroupOutput) DeployType() pulumi.StringOutput {
	return o.ApplyT(func(v *IpToSgtMappingGroup) pulumi.StringOutput { return v.DeployType }).(pulumi.StringOutput)
}

// Description
func (o IpToSgtMappingGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMappingGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the IP to SGT mapping Group
func (o IpToSgtMappingGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpToSgtMappingGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Trustsec Security Group ID
func (o IpToSgtMappingGroupOutput) Sgt() pulumi.StringOutput {
	return o.ApplyT(func(v *IpToSgtMappingGroup) pulumi.StringOutput { return v.Sgt }).(pulumi.StringOutput)
}

type IpToSgtMappingGroupArrayOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpToSgtMappingGroup)(nil)).Elem()
}

func (o IpToSgtMappingGroupArrayOutput) ToIpToSgtMappingGroupArrayOutput() IpToSgtMappingGroupArrayOutput {
	return o
}

func (o IpToSgtMappingGroupArrayOutput) ToIpToSgtMappingGroupArrayOutputWithContext(ctx context.Context) IpToSgtMappingGroupArrayOutput {
	return o
}

func (o IpToSgtMappingGroupArrayOutput) Index(i pulumi.IntInput) IpToSgtMappingGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpToSgtMappingGroup {
		return vs[0].([]*IpToSgtMappingGroup)[vs[1].(int)]
	}).(IpToSgtMappingGroupOutput)
}

type IpToSgtMappingGroupMapOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpToSgtMappingGroup)(nil)).Elem()
}

func (o IpToSgtMappingGroupMapOutput) ToIpToSgtMappingGroupMapOutput() IpToSgtMappingGroupMapOutput {
	return o
}

func (o IpToSgtMappingGroupMapOutput) ToIpToSgtMappingGroupMapOutputWithContext(ctx context.Context) IpToSgtMappingGroupMapOutput {
	return o
}

func (o IpToSgtMappingGroupMapOutput) MapIndex(k pulumi.StringInput) IpToSgtMappingGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpToSgtMappingGroup {
		return vs[0].(map[string]*IpToSgtMappingGroup)[vs[1].(string)]
	}).(IpToSgtMappingGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingGroupInput)(nil)).Elem(), &IpToSgtMappingGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingGroupArrayInput)(nil)).Elem(), IpToSgtMappingGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingGroupMapInput)(nil)).Elem(), IpToSgtMappingGroupMap{})
	pulumi.RegisterOutputType(IpToSgtMappingGroupOutput{})
	pulumi.RegisterOutputType(IpToSgtMappingGroupArrayOutput{})
	pulumi.RegisterOutputType(IpToSgtMappingGroupMapOutput{})
}
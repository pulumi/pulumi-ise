// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a TrustSec IP to SGT Mapping.
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
//			_, err := trustsec.NewIpToSgtMapping(ctx, "example", &trustsec.IpToSgtMappingArgs{
//				Name:       pulumi.String("10.0.0.1/32"),
//				DeployType: pulumi.String("ALL"),
//				HostIp:     pulumi.String("10.0.0.1/32"),
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
// $ pulumi import ise:trustsec/ipToSgtMapping:IpToSgtMapping example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type IpToSgtMapping struct {
	pulumi.CustomResourceState

	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo pulumi.StringPtrOutput `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringPtrOutput `pulumi:"deployType"`
	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Mandatory if `hostName` is empty
	HostIp pulumi.StringPtrOutput `pulumi:"hostIp"`
	// Mandatory if `hostIp` is empty
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup pulumi.StringPtrOutput `pulumi:"mappingGroup"`
	// The name of the IP to SGT mapping
	Name pulumi.StringOutput `pulumi:"name"`
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt pulumi.StringPtrOutput `pulumi:"sgt"`
}

// NewIpToSgtMapping registers a new resource with the given unique name, arguments, and options.
func NewIpToSgtMapping(ctx *pulumi.Context,
	name string, args *IpToSgtMappingArgs, opts ...pulumi.ResourceOption) (*IpToSgtMapping, error) {
	if args == nil {
		args = &IpToSgtMappingArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpToSgtMapping
	err := ctx.RegisterResource("ise:trustsec/ipToSgtMapping:IpToSgtMapping", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpToSgtMapping gets an existing IpToSgtMapping resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpToSgtMapping(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpToSgtMappingState, opts ...pulumi.ResourceOption) (*IpToSgtMapping, error) {
	var resource IpToSgtMapping
	err := ctx.ReadResource("ise:trustsec/ipToSgtMapping:IpToSgtMapping", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpToSgtMapping resources.
type ipToSgtMappingState struct {
	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo *string `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType *string `pulumi:"deployType"`
	// Description
	Description *string `pulumi:"description"`
	// Mandatory if `hostName` is empty
	HostIp *string `pulumi:"hostIp"`
	// Mandatory if `hostIp` is empty
	HostName *string `pulumi:"hostName"`
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup *string `pulumi:"mappingGroup"`
	// The name of the IP to SGT mapping
	Name *string `pulumi:"name"`
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt *string `pulumi:"sgt"`
}

type IpToSgtMappingState struct {
	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo pulumi.StringPtrInput
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringPtrInput
	// Description
	Description pulumi.StringPtrInput
	// Mandatory if `hostName` is empty
	HostIp pulumi.StringPtrInput
	// Mandatory if `hostIp` is empty
	HostName pulumi.StringPtrInput
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup pulumi.StringPtrInput
	// The name of the IP to SGT mapping
	Name pulumi.StringPtrInput
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt pulumi.StringPtrInput
}

func (IpToSgtMappingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipToSgtMappingState)(nil)).Elem()
}

type ipToSgtMappingArgs struct {
	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo *string `pulumi:"deployTo"`
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType *string `pulumi:"deployType"`
	// Description
	Description *string `pulumi:"description"`
	// Mandatory if `hostName` is empty
	HostIp *string `pulumi:"hostIp"`
	// Mandatory if `hostIp` is empty
	HostName *string `pulumi:"hostName"`
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup *string `pulumi:"mappingGroup"`
	// The name of the IP to SGT mapping
	Name *string `pulumi:"name"`
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt *string `pulumi:"sgt"`
}

// The set of arguments for constructing a IpToSgtMapping resource.
type IpToSgtMappingArgs struct {
	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo pulumi.StringPtrInput
	// Deploy Type - Choices: `ALL`, `ND`, `NDG`
	DeployType pulumi.StringPtrInput
	// Description
	Description pulumi.StringPtrInput
	// Mandatory if `hostName` is empty
	HostIp pulumi.StringPtrInput
	// Mandatory if `hostIp` is empty
	HostName pulumi.StringPtrInput
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup pulumi.StringPtrInput
	// The name of the IP to SGT mapping
	Name pulumi.StringPtrInput
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt pulumi.StringPtrInput
}

func (IpToSgtMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipToSgtMappingArgs)(nil)).Elem()
}

type IpToSgtMappingInput interface {
	pulumi.Input

	ToIpToSgtMappingOutput() IpToSgtMappingOutput
	ToIpToSgtMappingOutputWithContext(ctx context.Context) IpToSgtMappingOutput
}

func (*IpToSgtMapping) ElementType() reflect.Type {
	return reflect.TypeOf((**IpToSgtMapping)(nil)).Elem()
}

func (i *IpToSgtMapping) ToIpToSgtMappingOutput() IpToSgtMappingOutput {
	return i.ToIpToSgtMappingOutputWithContext(context.Background())
}

func (i *IpToSgtMapping) ToIpToSgtMappingOutputWithContext(ctx context.Context) IpToSgtMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingOutput)
}

// IpToSgtMappingArrayInput is an input type that accepts IpToSgtMappingArray and IpToSgtMappingArrayOutput values.
// You can construct a concrete instance of `IpToSgtMappingArrayInput` via:
//
//	IpToSgtMappingArray{ IpToSgtMappingArgs{...} }
type IpToSgtMappingArrayInput interface {
	pulumi.Input

	ToIpToSgtMappingArrayOutput() IpToSgtMappingArrayOutput
	ToIpToSgtMappingArrayOutputWithContext(context.Context) IpToSgtMappingArrayOutput
}

type IpToSgtMappingArray []IpToSgtMappingInput

func (IpToSgtMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpToSgtMapping)(nil)).Elem()
}

func (i IpToSgtMappingArray) ToIpToSgtMappingArrayOutput() IpToSgtMappingArrayOutput {
	return i.ToIpToSgtMappingArrayOutputWithContext(context.Background())
}

func (i IpToSgtMappingArray) ToIpToSgtMappingArrayOutputWithContext(ctx context.Context) IpToSgtMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingArrayOutput)
}

// IpToSgtMappingMapInput is an input type that accepts IpToSgtMappingMap and IpToSgtMappingMapOutput values.
// You can construct a concrete instance of `IpToSgtMappingMapInput` via:
//
//	IpToSgtMappingMap{ "key": IpToSgtMappingArgs{...} }
type IpToSgtMappingMapInput interface {
	pulumi.Input

	ToIpToSgtMappingMapOutput() IpToSgtMappingMapOutput
	ToIpToSgtMappingMapOutputWithContext(context.Context) IpToSgtMappingMapOutput
}

type IpToSgtMappingMap map[string]IpToSgtMappingInput

func (IpToSgtMappingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpToSgtMapping)(nil)).Elem()
}

func (i IpToSgtMappingMap) ToIpToSgtMappingMapOutput() IpToSgtMappingMapOutput {
	return i.ToIpToSgtMappingMapOutputWithContext(context.Background())
}

func (i IpToSgtMappingMap) ToIpToSgtMappingMapOutputWithContext(ctx context.Context) IpToSgtMappingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpToSgtMappingMapOutput)
}

type IpToSgtMappingOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpToSgtMapping)(nil)).Elem()
}

func (o IpToSgtMappingOutput) ToIpToSgtMappingOutput() IpToSgtMappingOutput {
	return o
}

func (o IpToSgtMappingOutput) ToIpToSgtMappingOutputWithContext(ctx context.Context) IpToSgtMappingOutput {
	return o
}

// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
func (o IpToSgtMappingOutput) DeployTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.DeployTo }).(pulumi.StringPtrOutput)
}

// Deploy Type - Choices: `ALL`, `ND`, `NDG`
func (o IpToSgtMappingOutput) DeployType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.DeployType }).(pulumi.StringPtrOutput)
}

// Description
func (o IpToSgtMappingOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Mandatory if `hostName` is empty
func (o IpToSgtMappingOutput) HostIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.HostIp }).(pulumi.StringPtrOutput)
}

// Mandatory if `hostIp` is empty
func (o IpToSgtMappingOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.HostName }).(pulumi.StringPtrOutput)
}

// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
func (o IpToSgtMappingOutput) MappingGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.MappingGroup }).(pulumi.StringPtrOutput)
}

// The name of the IP to SGT mapping
func (o IpToSgtMappingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
func (o IpToSgtMappingOutput) Sgt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpToSgtMapping) pulumi.StringPtrOutput { return v.Sgt }).(pulumi.StringPtrOutput)
}

type IpToSgtMappingArrayOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpToSgtMapping)(nil)).Elem()
}

func (o IpToSgtMappingArrayOutput) ToIpToSgtMappingArrayOutput() IpToSgtMappingArrayOutput {
	return o
}

func (o IpToSgtMappingArrayOutput) ToIpToSgtMappingArrayOutputWithContext(ctx context.Context) IpToSgtMappingArrayOutput {
	return o
}

func (o IpToSgtMappingArrayOutput) Index(i pulumi.IntInput) IpToSgtMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpToSgtMapping {
		return vs[0].([]*IpToSgtMapping)[vs[1].(int)]
	}).(IpToSgtMappingOutput)
}

type IpToSgtMappingMapOutput struct{ *pulumi.OutputState }

func (IpToSgtMappingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpToSgtMapping)(nil)).Elem()
}

func (o IpToSgtMappingMapOutput) ToIpToSgtMappingMapOutput() IpToSgtMappingMapOutput {
	return o
}

func (o IpToSgtMappingMapOutput) ToIpToSgtMappingMapOutputWithContext(ctx context.Context) IpToSgtMappingMapOutput {
	return o
}

func (o IpToSgtMappingMapOutput) MapIndex(k pulumi.StringInput) IpToSgtMappingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpToSgtMapping {
		return vs[0].(map[string]*IpToSgtMapping)[vs[1].(string)]
	}).(IpToSgtMappingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingInput)(nil)).Elem(), &IpToSgtMapping{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingArrayInput)(nil)).Elem(), IpToSgtMappingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpToSgtMappingMapInput)(nil)).Elem(), IpToSgtMappingMap{})
	pulumi.RegisterOutputType(IpToSgtMappingOutput{})
	pulumi.RegisterOutputType(IpToSgtMappingArrayOutput{})
	pulumi.RegisterOutputType(IpToSgtMappingMapOutput{})
}

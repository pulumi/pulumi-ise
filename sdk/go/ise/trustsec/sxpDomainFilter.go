// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a SXP Domain Filter.
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
//			_, err := trustsec.NewSxpDomainFilter(ctx, "example", &trustsec.SxpDomainFilterArgs{
//				Subnet:  pulumi.String("1.0.0.0/24"),
//				Vn:      pulumi.String("VN1"),
//				Domains: pulumi.String("default"),
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
// $ pulumi import ise:trustsec/sxpDomainFilter:SxpDomainFilter example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type SxpDomainFilter struct {
	pulumi.CustomResourceState

	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of SXP Domains, separated with comma
	Domains pulumi.StringOutput `pulumi:"domains"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt pulumi.StringPtrOutput `pulumi:"sgt"`
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet pulumi.StringPtrOutput `pulumi:"subnet"`
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn pulumi.StringPtrOutput `pulumi:"vn"`
}

// NewSxpDomainFilter registers a new resource with the given unique name, arguments, and options.
func NewSxpDomainFilter(ctx *pulumi.Context,
	name string, args *SxpDomainFilterArgs, opts ...pulumi.ResourceOption) (*SxpDomainFilter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domains == nil {
		return nil, errors.New("invalid value for required argument 'Domains'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SxpDomainFilter
	err := ctx.RegisterResource("ise:trustsec/sxpDomainFilter:SxpDomainFilter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSxpDomainFilter gets an existing SxpDomainFilter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSxpDomainFilter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SxpDomainFilterState, opts ...pulumi.ResourceOption) (*SxpDomainFilter, error) {
	var resource SxpDomainFilter
	err := ctx.ReadResource("ise:trustsec/sxpDomainFilter:SxpDomainFilter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SxpDomainFilter resources.
type sxpDomainFilterState struct {
	// Description
	Description *string `pulumi:"description"`
	// List of SXP Domains, separated with comma
	Domains *string `pulumi:"domains"`
	// Resource name
	Name *string `pulumi:"name"`
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt *string `pulumi:"sgt"`
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet *string `pulumi:"subnet"`
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn *string `pulumi:"vn"`
}

type SxpDomainFilterState struct {
	// Description
	Description pulumi.StringPtrInput
	// List of SXP Domains, separated with comma
	Domains pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt pulumi.StringPtrInput
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet pulumi.StringPtrInput
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn pulumi.StringPtrInput
}

func (SxpDomainFilterState) ElementType() reflect.Type {
	return reflect.TypeOf((*sxpDomainFilterState)(nil)).Elem()
}

type sxpDomainFilterArgs struct {
	// Description
	Description *string `pulumi:"description"`
	// List of SXP Domains, separated with comma
	Domains string `pulumi:"domains"`
	// Resource name
	Name *string `pulumi:"name"`
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt *string `pulumi:"sgt"`
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet *string `pulumi:"subnet"`
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn *string `pulumi:"vn"`
}

// The set of arguments for constructing a SxpDomainFilter resource.
type SxpDomainFilterArgs struct {
	// Description
	Description pulumi.StringPtrInput
	// List of SXP Domains, separated with comma
	Domains pulumi.StringInput
	// Resource name
	Name pulumi.StringPtrInput
	// SGT name or ID. At least one of subnet or sgt or vn should be defined
	Sgt pulumi.StringPtrInput
	// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
	Subnet pulumi.StringPtrInput
	// Virtual Network. At least one of subnet or sgt or vn should be defined
	Vn pulumi.StringPtrInput
}

func (SxpDomainFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sxpDomainFilterArgs)(nil)).Elem()
}

type SxpDomainFilterInput interface {
	pulumi.Input

	ToSxpDomainFilterOutput() SxpDomainFilterOutput
	ToSxpDomainFilterOutputWithContext(ctx context.Context) SxpDomainFilterOutput
}

func (*SxpDomainFilter) ElementType() reflect.Type {
	return reflect.TypeOf((**SxpDomainFilter)(nil)).Elem()
}

func (i *SxpDomainFilter) ToSxpDomainFilterOutput() SxpDomainFilterOutput {
	return i.ToSxpDomainFilterOutputWithContext(context.Background())
}

func (i *SxpDomainFilter) ToSxpDomainFilterOutputWithContext(ctx context.Context) SxpDomainFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SxpDomainFilterOutput)
}

// SxpDomainFilterArrayInput is an input type that accepts SxpDomainFilterArray and SxpDomainFilterArrayOutput values.
// You can construct a concrete instance of `SxpDomainFilterArrayInput` via:
//
//	SxpDomainFilterArray{ SxpDomainFilterArgs{...} }
type SxpDomainFilterArrayInput interface {
	pulumi.Input

	ToSxpDomainFilterArrayOutput() SxpDomainFilterArrayOutput
	ToSxpDomainFilterArrayOutputWithContext(context.Context) SxpDomainFilterArrayOutput
}

type SxpDomainFilterArray []SxpDomainFilterInput

func (SxpDomainFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SxpDomainFilter)(nil)).Elem()
}

func (i SxpDomainFilterArray) ToSxpDomainFilterArrayOutput() SxpDomainFilterArrayOutput {
	return i.ToSxpDomainFilterArrayOutputWithContext(context.Background())
}

func (i SxpDomainFilterArray) ToSxpDomainFilterArrayOutputWithContext(ctx context.Context) SxpDomainFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SxpDomainFilterArrayOutput)
}

// SxpDomainFilterMapInput is an input type that accepts SxpDomainFilterMap and SxpDomainFilterMapOutput values.
// You can construct a concrete instance of `SxpDomainFilterMapInput` via:
//
//	SxpDomainFilterMap{ "key": SxpDomainFilterArgs{...} }
type SxpDomainFilterMapInput interface {
	pulumi.Input

	ToSxpDomainFilterMapOutput() SxpDomainFilterMapOutput
	ToSxpDomainFilterMapOutputWithContext(context.Context) SxpDomainFilterMapOutput
}

type SxpDomainFilterMap map[string]SxpDomainFilterInput

func (SxpDomainFilterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SxpDomainFilter)(nil)).Elem()
}

func (i SxpDomainFilterMap) ToSxpDomainFilterMapOutput() SxpDomainFilterMapOutput {
	return i.ToSxpDomainFilterMapOutputWithContext(context.Background())
}

func (i SxpDomainFilterMap) ToSxpDomainFilterMapOutputWithContext(ctx context.Context) SxpDomainFilterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SxpDomainFilterMapOutput)
}

type SxpDomainFilterOutput struct{ *pulumi.OutputState }

func (SxpDomainFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SxpDomainFilter)(nil)).Elem()
}

func (o SxpDomainFilterOutput) ToSxpDomainFilterOutput() SxpDomainFilterOutput {
	return o
}

func (o SxpDomainFilterOutput) ToSxpDomainFilterOutputWithContext(ctx context.Context) SxpDomainFilterOutput {
	return o
}

// Description
func (o SxpDomainFilterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of SXP Domains, separated with comma
func (o SxpDomainFilterOutput) Domains() pulumi.StringOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringOutput { return v.Domains }).(pulumi.StringOutput)
}

// Resource name
func (o SxpDomainFilterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// SGT name or ID. At least one of subnet or sgt or vn should be defined
func (o SxpDomainFilterOutput) Sgt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringPtrOutput { return v.Sgt }).(pulumi.StringPtrOutput)
}

// Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
func (o SxpDomainFilterOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringPtrOutput { return v.Subnet }).(pulumi.StringPtrOutput)
}

// Virtual Network. At least one of subnet or sgt or vn should be defined
func (o SxpDomainFilterOutput) Vn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SxpDomainFilter) pulumi.StringPtrOutput { return v.Vn }).(pulumi.StringPtrOutput)
}

type SxpDomainFilterArrayOutput struct{ *pulumi.OutputState }

func (SxpDomainFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SxpDomainFilter)(nil)).Elem()
}

func (o SxpDomainFilterArrayOutput) ToSxpDomainFilterArrayOutput() SxpDomainFilterArrayOutput {
	return o
}

func (o SxpDomainFilterArrayOutput) ToSxpDomainFilterArrayOutputWithContext(ctx context.Context) SxpDomainFilterArrayOutput {
	return o
}

func (o SxpDomainFilterArrayOutput) Index(i pulumi.IntInput) SxpDomainFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SxpDomainFilter {
		return vs[0].([]*SxpDomainFilter)[vs[1].(int)]
	}).(SxpDomainFilterOutput)
}

type SxpDomainFilterMapOutput struct{ *pulumi.OutputState }

func (SxpDomainFilterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SxpDomainFilter)(nil)).Elem()
}

func (o SxpDomainFilterMapOutput) ToSxpDomainFilterMapOutput() SxpDomainFilterMapOutput {
	return o
}

func (o SxpDomainFilterMapOutput) ToSxpDomainFilterMapOutputWithContext(ctx context.Context) SxpDomainFilterMapOutput {
	return o
}

func (o SxpDomainFilterMapOutput) MapIndex(k pulumi.StringInput) SxpDomainFilterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SxpDomainFilter {
		return vs[0].(map[string]*SxpDomainFilter)[vs[1].(string)]
	}).(SxpDomainFilterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SxpDomainFilterInput)(nil)).Elem(), &SxpDomainFilter{})
	pulumi.RegisterInputType(reflect.TypeOf((*SxpDomainFilterArrayInput)(nil)).Elem(), SxpDomainFilterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SxpDomainFilterMapInput)(nil)).Elem(), SxpDomainFilterMap{})
	pulumi.RegisterOutputType(SxpDomainFilterOutput{})
	pulumi.RegisterOutputType(SxpDomainFilterArrayOutput{})
	pulumi.RegisterOutputType(SxpDomainFilterMapOutput{})
}

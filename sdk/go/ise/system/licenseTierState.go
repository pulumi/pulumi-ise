// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a License Tier State.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/system"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := system.NewLicenseTierState(ctx, "example", &system.LicenseTierStateArgs{
//				Licenses: system.LicenseTierStateLicenseArray{
//					&system.LicenseTierStateLicenseArgs{
//						Name:   pulumi.String("ESSENTIAL"),
//						Status: pulumi.String("ENABLED"),
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
//
// ## Import
//
// ```sh
// $ pulumi import ise:system/licenseTierState:LicenseTierState example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type LicenseTierState struct {
	pulumi.CustomResourceState

	// List of licenses
	Licenses LicenseTierStateLicenseArrayOutput `pulumi:"licenses"`
}

// NewLicenseTierState registers a new resource with the given unique name, arguments, and options.
func NewLicenseTierState(ctx *pulumi.Context,
	name string, args *LicenseTierStateArgs, opts ...pulumi.ResourceOption) (*LicenseTierState, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Licenses == nil {
		return nil, errors.New("invalid value for required argument 'Licenses'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LicenseTierState
	err := ctx.RegisterResource("ise:system/licenseTierState:LicenseTierState", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLicenseTierState gets an existing LicenseTierState resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLicenseTierState(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LicenseTierStateState, opts ...pulumi.ResourceOption) (*LicenseTierState, error) {
	var resource LicenseTierState
	err := ctx.ReadResource("ise:system/licenseTierState:LicenseTierState", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LicenseTierState resources.
type licenseTierStateState struct {
	// List of licenses
	Licenses []LicenseTierStateLicense `pulumi:"licenses"`
}

type LicenseTierStateState struct {
	// List of licenses
	Licenses LicenseTierStateLicenseArrayInput
}

func (LicenseTierStateState) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseTierStateState)(nil)).Elem()
}

type licenseTierStateArgs struct {
	// List of licenses
	Licenses []LicenseTierStateLicense `pulumi:"licenses"`
}

// The set of arguments for constructing a LicenseTierState resource.
type LicenseTierStateArgs struct {
	// List of licenses
	Licenses LicenseTierStateLicenseArrayInput
}

func (LicenseTierStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*licenseTierStateArgs)(nil)).Elem()
}

type LicenseTierStateInput interface {
	pulumi.Input

	ToLicenseTierStateOutput() LicenseTierStateOutput
	ToLicenseTierStateOutputWithContext(ctx context.Context) LicenseTierStateOutput
}

func (*LicenseTierState) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseTierState)(nil)).Elem()
}

func (i *LicenseTierState) ToLicenseTierStateOutput() LicenseTierStateOutput {
	return i.ToLicenseTierStateOutputWithContext(context.Background())
}

func (i *LicenseTierState) ToLicenseTierStateOutputWithContext(ctx context.Context) LicenseTierStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseTierStateOutput)
}

// LicenseTierStateArrayInput is an input type that accepts LicenseTierStateArray and LicenseTierStateArrayOutput values.
// You can construct a concrete instance of `LicenseTierStateArrayInput` via:
//
//	LicenseTierStateArray{ LicenseTierStateArgs{...} }
type LicenseTierStateArrayInput interface {
	pulumi.Input

	ToLicenseTierStateArrayOutput() LicenseTierStateArrayOutput
	ToLicenseTierStateArrayOutputWithContext(context.Context) LicenseTierStateArrayOutput
}

type LicenseTierStateArray []LicenseTierStateInput

func (LicenseTierStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LicenseTierState)(nil)).Elem()
}

func (i LicenseTierStateArray) ToLicenseTierStateArrayOutput() LicenseTierStateArrayOutput {
	return i.ToLicenseTierStateArrayOutputWithContext(context.Background())
}

func (i LicenseTierStateArray) ToLicenseTierStateArrayOutputWithContext(ctx context.Context) LicenseTierStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseTierStateArrayOutput)
}

// LicenseTierStateMapInput is an input type that accepts LicenseTierStateMap and LicenseTierStateMapOutput values.
// You can construct a concrete instance of `LicenseTierStateMapInput` via:
//
//	LicenseTierStateMap{ "key": LicenseTierStateArgs{...} }
type LicenseTierStateMapInput interface {
	pulumi.Input

	ToLicenseTierStateMapOutput() LicenseTierStateMapOutput
	ToLicenseTierStateMapOutputWithContext(context.Context) LicenseTierStateMapOutput
}

type LicenseTierStateMap map[string]LicenseTierStateInput

func (LicenseTierStateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LicenseTierState)(nil)).Elem()
}

func (i LicenseTierStateMap) ToLicenseTierStateMapOutput() LicenseTierStateMapOutput {
	return i.ToLicenseTierStateMapOutputWithContext(context.Background())
}

func (i LicenseTierStateMap) ToLicenseTierStateMapOutputWithContext(ctx context.Context) LicenseTierStateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseTierStateMapOutput)
}

type LicenseTierStateOutput struct{ *pulumi.OutputState }

func (LicenseTierStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LicenseTierState)(nil)).Elem()
}

func (o LicenseTierStateOutput) ToLicenseTierStateOutput() LicenseTierStateOutput {
	return o
}

func (o LicenseTierStateOutput) ToLicenseTierStateOutputWithContext(ctx context.Context) LicenseTierStateOutput {
	return o
}

// List of licenses
func (o LicenseTierStateOutput) Licenses() LicenseTierStateLicenseArrayOutput {
	return o.ApplyT(func(v *LicenseTierState) LicenseTierStateLicenseArrayOutput { return v.Licenses }).(LicenseTierStateLicenseArrayOutput)
}

type LicenseTierStateArrayOutput struct{ *pulumi.OutputState }

func (LicenseTierStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LicenseTierState)(nil)).Elem()
}

func (o LicenseTierStateArrayOutput) ToLicenseTierStateArrayOutput() LicenseTierStateArrayOutput {
	return o
}

func (o LicenseTierStateArrayOutput) ToLicenseTierStateArrayOutputWithContext(ctx context.Context) LicenseTierStateArrayOutput {
	return o
}

func (o LicenseTierStateArrayOutput) Index(i pulumi.IntInput) LicenseTierStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LicenseTierState {
		return vs[0].([]*LicenseTierState)[vs[1].(int)]
	}).(LicenseTierStateOutput)
}

type LicenseTierStateMapOutput struct{ *pulumi.OutputState }

func (LicenseTierStateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LicenseTierState)(nil)).Elem()
}

func (o LicenseTierStateMapOutput) ToLicenseTierStateMapOutput() LicenseTierStateMapOutput {
	return o
}

func (o LicenseTierStateMapOutput) ToLicenseTierStateMapOutputWithContext(ctx context.Context) LicenseTierStateMapOutput {
	return o
}

func (o LicenseTierStateMapOutput) MapIndex(k pulumi.StringInput) LicenseTierStateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LicenseTierState {
		return vs[0].(map[string]*LicenseTierState)[vs[1].(string)]
	}).(LicenseTierStateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTierStateInput)(nil)).Elem(), &LicenseTierState{})
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTierStateArrayInput)(nil)).Elem(), LicenseTierStateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTierStateMapInput)(nil)).Elem(), LicenseTierStateMap{})
	pulumi.RegisterOutputType(LicenseTierStateOutput{})
	pulumi.RegisterOutputType(LicenseTierStateArrayOutput{})
	pulumi.RegisterOutputType(LicenseTierStateMapOutput{})
}

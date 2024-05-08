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

// This resource can manage a TACACS allowed protocols policy element.
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
//			_, err := deviceadmin.NewAllowedProtocolsTacacs(ctx, "example", &deviceadmin.AllowedProtocolsTacacsArgs{
//				Name:          pulumi.String("Protocols1"),
//				Description:   pulumi.String("My allowed TACACS protocols"),
//				AllowPapAscii: pulumi.Bool(true),
//				AllowChap:     pulumi.Bool(true),
//				AllowMsChapV1: pulumi.Bool(true),
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
// $ pulumi import ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs example "76d24097-41c4-4558-a4d0-a8c07ac08470"
// ```
type AllowedProtocolsTacacs struct {
	pulumi.CustomResourceState

	// Allow CHAP
	AllowChap pulumi.BoolOutput `pulumi:"allowChap"`
	// Allow MS CHAP v1
	AllowMsChapV1 pulumi.BoolOutput `pulumi:"allowMsChapV1"`
	// Allow PAP ASCII
	AllowPapAscii pulumi.BoolOutput `pulumi:"allowPapAscii"`
	// Description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the allowed protocols
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAllowedProtocolsTacacs registers a new resource with the given unique name, arguments, and options.
func NewAllowedProtocolsTacacs(ctx *pulumi.Context,
	name string, args *AllowedProtocolsTacacsArgs, opts ...pulumi.ResourceOption) (*AllowedProtocolsTacacs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowChap == nil {
		return nil, errors.New("invalid value for required argument 'AllowChap'")
	}
	if args.AllowMsChapV1 == nil {
		return nil, errors.New("invalid value for required argument 'AllowMsChapV1'")
	}
	if args.AllowPapAscii == nil {
		return nil, errors.New("invalid value for required argument 'AllowPapAscii'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AllowedProtocolsTacacs
	err := ctx.RegisterResource("ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAllowedProtocolsTacacs gets an existing AllowedProtocolsTacacs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAllowedProtocolsTacacs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AllowedProtocolsTacacsState, opts ...pulumi.ResourceOption) (*AllowedProtocolsTacacs, error) {
	var resource AllowedProtocolsTacacs
	err := ctx.ReadResource("ise:deviceadmin/allowedProtocolsTacacs:AllowedProtocolsTacacs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AllowedProtocolsTacacs resources.
type allowedProtocolsTacacsState struct {
	// Allow CHAP
	AllowChap *bool `pulumi:"allowChap"`
	// Allow MS CHAP v1
	AllowMsChapV1 *bool `pulumi:"allowMsChapV1"`
	// Allow PAP ASCII
	AllowPapAscii *bool `pulumi:"allowPapAscii"`
	// Description
	Description *string `pulumi:"description"`
	// The name of the allowed protocols
	Name *string `pulumi:"name"`
}

type AllowedProtocolsTacacsState struct {
	// Allow CHAP
	AllowChap pulumi.BoolPtrInput
	// Allow MS CHAP v1
	AllowMsChapV1 pulumi.BoolPtrInput
	// Allow PAP ASCII
	AllowPapAscii pulumi.BoolPtrInput
	// Description
	Description pulumi.StringPtrInput
	// The name of the allowed protocols
	Name pulumi.StringPtrInput
}

func (AllowedProtocolsTacacsState) ElementType() reflect.Type {
	return reflect.TypeOf((*allowedProtocolsTacacsState)(nil)).Elem()
}

type allowedProtocolsTacacsArgs struct {
	// Allow CHAP
	AllowChap bool `pulumi:"allowChap"`
	// Allow MS CHAP v1
	AllowMsChapV1 bool `pulumi:"allowMsChapV1"`
	// Allow PAP ASCII
	AllowPapAscii bool `pulumi:"allowPapAscii"`
	// Description
	Description *string `pulumi:"description"`
	// The name of the allowed protocols
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a AllowedProtocolsTacacs resource.
type AllowedProtocolsTacacsArgs struct {
	// Allow CHAP
	AllowChap pulumi.BoolInput
	// Allow MS CHAP v1
	AllowMsChapV1 pulumi.BoolInput
	// Allow PAP ASCII
	AllowPapAscii pulumi.BoolInput
	// Description
	Description pulumi.StringPtrInput
	// The name of the allowed protocols
	Name pulumi.StringPtrInput
}

func (AllowedProtocolsTacacsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*allowedProtocolsTacacsArgs)(nil)).Elem()
}

type AllowedProtocolsTacacsInput interface {
	pulumi.Input

	ToAllowedProtocolsTacacsOutput() AllowedProtocolsTacacsOutput
	ToAllowedProtocolsTacacsOutputWithContext(ctx context.Context) AllowedProtocolsTacacsOutput
}

func (*AllowedProtocolsTacacs) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedProtocolsTacacs)(nil)).Elem()
}

func (i *AllowedProtocolsTacacs) ToAllowedProtocolsTacacsOutput() AllowedProtocolsTacacsOutput {
	return i.ToAllowedProtocolsTacacsOutputWithContext(context.Background())
}

func (i *AllowedProtocolsTacacs) ToAllowedProtocolsTacacsOutputWithContext(ctx context.Context) AllowedProtocolsTacacsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedProtocolsTacacsOutput)
}

// AllowedProtocolsTacacsArrayInput is an input type that accepts AllowedProtocolsTacacsArray and AllowedProtocolsTacacsArrayOutput values.
// You can construct a concrete instance of `AllowedProtocolsTacacsArrayInput` via:
//
//	AllowedProtocolsTacacsArray{ AllowedProtocolsTacacsArgs{...} }
type AllowedProtocolsTacacsArrayInput interface {
	pulumi.Input

	ToAllowedProtocolsTacacsArrayOutput() AllowedProtocolsTacacsArrayOutput
	ToAllowedProtocolsTacacsArrayOutputWithContext(context.Context) AllowedProtocolsTacacsArrayOutput
}

type AllowedProtocolsTacacsArray []AllowedProtocolsTacacsInput

func (AllowedProtocolsTacacsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowedProtocolsTacacs)(nil)).Elem()
}

func (i AllowedProtocolsTacacsArray) ToAllowedProtocolsTacacsArrayOutput() AllowedProtocolsTacacsArrayOutput {
	return i.ToAllowedProtocolsTacacsArrayOutputWithContext(context.Background())
}

func (i AllowedProtocolsTacacsArray) ToAllowedProtocolsTacacsArrayOutputWithContext(ctx context.Context) AllowedProtocolsTacacsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedProtocolsTacacsArrayOutput)
}

// AllowedProtocolsTacacsMapInput is an input type that accepts AllowedProtocolsTacacsMap and AllowedProtocolsTacacsMapOutput values.
// You can construct a concrete instance of `AllowedProtocolsTacacsMapInput` via:
//
//	AllowedProtocolsTacacsMap{ "key": AllowedProtocolsTacacsArgs{...} }
type AllowedProtocolsTacacsMapInput interface {
	pulumi.Input

	ToAllowedProtocolsTacacsMapOutput() AllowedProtocolsTacacsMapOutput
	ToAllowedProtocolsTacacsMapOutputWithContext(context.Context) AllowedProtocolsTacacsMapOutput
}

type AllowedProtocolsTacacsMap map[string]AllowedProtocolsTacacsInput

func (AllowedProtocolsTacacsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowedProtocolsTacacs)(nil)).Elem()
}

func (i AllowedProtocolsTacacsMap) ToAllowedProtocolsTacacsMapOutput() AllowedProtocolsTacacsMapOutput {
	return i.ToAllowedProtocolsTacacsMapOutputWithContext(context.Background())
}

func (i AllowedProtocolsTacacsMap) ToAllowedProtocolsTacacsMapOutputWithContext(ctx context.Context) AllowedProtocolsTacacsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AllowedProtocolsTacacsMapOutput)
}

type AllowedProtocolsTacacsOutput struct{ *pulumi.OutputState }

func (AllowedProtocolsTacacsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllowedProtocolsTacacs)(nil)).Elem()
}

func (o AllowedProtocolsTacacsOutput) ToAllowedProtocolsTacacsOutput() AllowedProtocolsTacacsOutput {
	return o
}

func (o AllowedProtocolsTacacsOutput) ToAllowedProtocolsTacacsOutputWithContext(ctx context.Context) AllowedProtocolsTacacsOutput {
	return o
}

// Allow CHAP
func (o AllowedProtocolsTacacsOutput) AllowChap() pulumi.BoolOutput {
	return o.ApplyT(func(v *AllowedProtocolsTacacs) pulumi.BoolOutput { return v.AllowChap }).(pulumi.BoolOutput)
}

// Allow MS CHAP v1
func (o AllowedProtocolsTacacsOutput) AllowMsChapV1() pulumi.BoolOutput {
	return o.ApplyT(func(v *AllowedProtocolsTacacs) pulumi.BoolOutput { return v.AllowMsChapV1 }).(pulumi.BoolOutput)
}

// Allow PAP ASCII
func (o AllowedProtocolsTacacsOutput) AllowPapAscii() pulumi.BoolOutput {
	return o.ApplyT(func(v *AllowedProtocolsTacacs) pulumi.BoolOutput { return v.AllowPapAscii }).(pulumi.BoolOutput)
}

// Description
func (o AllowedProtocolsTacacsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AllowedProtocolsTacacs) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The name of the allowed protocols
func (o AllowedProtocolsTacacsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AllowedProtocolsTacacs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type AllowedProtocolsTacacsArrayOutput struct{ *pulumi.OutputState }

func (AllowedProtocolsTacacsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AllowedProtocolsTacacs)(nil)).Elem()
}

func (o AllowedProtocolsTacacsArrayOutput) ToAllowedProtocolsTacacsArrayOutput() AllowedProtocolsTacacsArrayOutput {
	return o
}

func (o AllowedProtocolsTacacsArrayOutput) ToAllowedProtocolsTacacsArrayOutputWithContext(ctx context.Context) AllowedProtocolsTacacsArrayOutput {
	return o
}

func (o AllowedProtocolsTacacsArrayOutput) Index(i pulumi.IntInput) AllowedProtocolsTacacsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AllowedProtocolsTacacs {
		return vs[0].([]*AllowedProtocolsTacacs)[vs[1].(int)]
	}).(AllowedProtocolsTacacsOutput)
}

type AllowedProtocolsTacacsMapOutput struct{ *pulumi.OutputState }

func (AllowedProtocolsTacacsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AllowedProtocolsTacacs)(nil)).Elem()
}

func (o AllowedProtocolsTacacsMapOutput) ToAllowedProtocolsTacacsMapOutput() AllowedProtocolsTacacsMapOutput {
	return o
}

func (o AllowedProtocolsTacacsMapOutput) ToAllowedProtocolsTacacsMapOutputWithContext(ctx context.Context) AllowedProtocolsTacacsMapOutput {
	return o
}

func (o AllowedProtocolsTacacsMapOutput) MapIndex(k pulumi.StringInput) AllowedProtocolsTacacsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AllowedProtocolsTacacs {
		return vs[0].(map[string]*AllowedProtocolsTacacs)[vs[1].(string)]
	}).(AllowedProtocolsTacacsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedProtocolsTacacsInput)(nil)).Elem(), &AllowedProtocolsTacacs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedProtocolsTacacsArrayInput)(nil)).Elem(), AllowedProtocolsTacacsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AllowedProtocolsTacacsMapInput)(nil)).Elem(), AllowedProtocolsTacacsMap{})
	pulumi.RegisterOutputType(AllowedProtocolsTacacsOutput{})
	pulumi.RegisterOutputType(AllowedProtocolsTacacsArrayOutput{})
	pulumi.RegisterOutputType(AllowedProtocolsTacacsMapOutput{})
}

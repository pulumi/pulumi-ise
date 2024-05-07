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

// This resource can manage a Network Access Dictionary.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/NetworkAccess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := NetworkAccess.NewDictionary(ctx, "example", &NetworkAccess.DictionaryArgs{
//				Name:               pulumi.String("Dict1"),
//				Description:        pulumi.String("My description"),
//				Version:            pulumi.String("1.1"),
//				DictionaryAttrType: pulumi.String("ENTITY_ATTR"),
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
// $ pulumi import ise:NetworkAccess/dictionary:Dictionary example "Dict1"
// ```
type Dictionary struct {
	pulumi.CustomResourceState

	// The description of the dictionary
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
	DictionaryAttrType pulumi.StringOutput `pulumi:"dictionaryAttrType"`
	// The dictionary name
	Name pulumi.StringOutput `pulumi:"name"`
	// The version of the dictionary
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDictionary registers a new resource with the given unique name, arguments, and options.
func NewDictionary(ctx *pulumi.Context,
	name string, args *DictionaryArgs, opts ...pulumi.ResourceOption) (*Dictionary, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DictionaryAttrType == nil {
		return nil, errors.New("invalid value for required argument 'DictionaryAttrType'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Dictionary
	err := ctx.RegisterResource("ise:NetworkAccess/dictionary:Dictionary", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDictionary gets an existing Dictionary resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDictionary(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DictionaryState, opts ...pulumi.ResourceOption) (*Dictionary, error) {
	var resource Dictionary
	err := ctx.ReadResource("ise:NetworkAccess/dictionary:Dictionary", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dictionary resources.
type dictionaryState struct {
	// The description of the dictionary
	Description *string `pulumi:"description"`
	// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
	DictionaryAttrType *string `pulumi:"dictionaryAttrType"`
	// The dictionary name
	Name *string `pulumi:"name"`
	// The version of the dictionary
	Version *string `pulumi:"version"`
}

type DictionaryState struct {
	// The description of the dictionary
	Description pulumi.StringPtrInput
	// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
	DictionaryAttrType pulumi.StringPtrInput
	// The dictionary name
	Name pulumi.StringPtrInput
	// The version of the dictionary
	Version pulumi.StringPtrInput
}

func (DictionaryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dictionaryState)(nil)).Elem()
}

type dictionaryArgs struct {
	// The description of the dictionary
	Description *string `pulumi:"description"`
	// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
	DictionaryAttrType string `pulumi:"dictionaryAttrType"`
	// The dictionary name
	Name *string `pulumi:"name"`
	// The version of the dictionary
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Dictionary resource.
type DictionaryArgs struct {
	// The description of the dictionary
	Description pulumi.StringPtrInput
	// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
	DictionaryAttrType pulumi.StringInput
	// The dictionary name
	Name pulumi.StringPtrInput
	// The version of the dictionary
	Version pulumi.StringInput
}

func (DictionaryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dictionaryArgs)(nil)).Elem()
}

type DictionaryInput interface {
	pulumi.Input

	ToDictionaryOutput() DictionaryOutput
	ToDictionaryOutputWithContext(ctx context.Context) DictionaryOutput
}

func (*Dictionary) ElementType() reflect.Type {
	return reflect.TypeOf((**Dictionary)(nil)).Elem()
}

func (i *Dictionary) ToDictionaryOutput() DictionaryOutput {
	return i.ToDictionaryOutputWithContext(context.Background())
}

func (i *Dictionary) ToDictionaryOutputWithContext(ctx context.Context) DictionaryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DictionaryOutput)
}

// DictionaryArrayInput is an input type that accepts DictionaryArray and DictionaryArrayOutput values.
// You can construct a concrete instance of `DictionaryArrayInput` via:
//
//	DictionaryArray{ DictionaryArgs{...} }
type DictionaryArrayInput interface {
	pulumi.Input

	ToDictionaryArrayOutput() DictionaryArrayOutput
	ToDictionaryArrayOutputWithContext(context.Context) DictionaryArrayOutput
}

type DictionaryArray []DictionaryInput

func (DictionaryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dictionary)(nil)).Elem()
}

func (i DictionaryArray) ToDictionaryArrayOutput() DictionaryArrayOutput {
	return i.ToDictionaryArrayOutputWithContext(context.Background())
}

func (i DictionaryArray) ToDictionaryArrayOutputWithContext(ctx context.Context) DictionaryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DictionaryArrayOutput)
}

// DictionaryMapInput is an input type that accepts DictionaryMap and DictionaryMapOutput values.
// You can construct a concrete instance of `DictionaryMapInput` via:
//
//	DictionaryMap{ "key": DictionaryArgs{...} }
type DictionaryMapInput interface {
	pulumi.Input

	ToDictionaryMapOutput() DictionaryMapOutput
	ToDictionaryMapOutputWithContext(context.Context) DictionaryMapOutput
}

type DictionaryMap map[string]DictionaryInput

func (DictionaryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dictionary)(nil)).Elem()
}

func (i DictionaryMap) ToDictionaryMapOutput() DictionaryMapOutput {
	return i.ToDictionaryMapOutputWithContext(context.Background())
}

func (i DictionaryMap) ToDictionaryMapOutputWithContext(ctx context.Context) DictionaryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DictionaryMapOutput)
}

type DictionaryOutput struct{ *pulumi.OutputState }

func (DictionaryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dictionary)(nil)).Elem()
}

func (o DictionaryOutput) ToDictionaryOutput() DictionaryOutput {
	return o
}

func (o DictionaryOutput) ToDictionaryOutputWithContext(ctx context.Context) DictionaryOutput {
	return o
}

// The description of the dictionary
func (o DictionaryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dictionary) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The dictionary attribute type - Choices: `ENTITY_ATTR`, `MSG_ATTR`, `PIP_ATTR`
func (o DictionaryOutput) DictionaryAttrType() pulumi.StringOutput {
	return o.ApplyT(func(v *Dictionary) pulumi.StringOutput { return v.DictionaryAttrType }).(pulumi.StringOutput)
}

// The dictionary name
func (o DictionaryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dictionary) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The version of the dictionary
func (o DictionaryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Dictionary) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type DictionaryArrayOutput struct{ *pulumi.OutputState }

func (DictionaryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dictionary)(nil)).Elem()
}

func (o DictionaryArrayOutput) ToDictionaryArrayOutput() DictionaryArrayOutput {
	return o
}

func (o DictionaryArrayOutput) ToDictionaryArrayOutputWithContext(ctx context.Context) DictionaryArrayOutput {
	return o
}

func (o DictionaryArrayOutput) Index(i pulumi.IntInput) DictionaryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dictionary {
		return vs[0].([]*Dictionary)[vs[1].(int)]
	}).(DictionaryOutput)
}

type DictionaryMapOutput struct{ *pulumi.OutputState }

func (DictionaryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dictionary)(nil)).Elem()
}

func (o DictionaryMapOutput) ToDictionaryMapOutput() DictionaryMapOutput {
	return o
}

func (o DictionaryMapOutput) ToDictionaryMapOutputWithContext(ctx context.Context) DictionaryMapOutput {
	return o
}

func (o DictionaryMapOutput) MapIndex(k pulumi.StringInput) DictionaryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dictionary {
		return vs[0].(map[string]*Dictionary)[vs[1].(string)]
	}).(DictionaryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DictionaryInput)(nil)).Elem(), &Dictionary{})
	pulumi.RegisterInputType(reflect.TypeOf((*DictionaryArrayInput)(nil)).Elem(), DictionaryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DictionaryMapInput)(nil)).Elem(), DictionaryMap{})
	pulumi.RegisterOutputType(DictionaryOutput{})
	pulumi.RegisterOutputType(DictionaryArrayOutput{})
	pulumi.RegisterOutputType(DictionaryMapOutput{})
}

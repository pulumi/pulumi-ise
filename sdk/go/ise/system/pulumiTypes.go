// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type LicenseTierStateLicense struct {
	// License name
	//   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
	Name string `pulumi:"name"`
	// License status
	//   - Choices: `ENABLED`, `DISABLED`
	Status string `pulumi:"status"`
}

// LicenseTierStateLicenseInput is an input type that accepts LicenseTierStateLicenseArgs and LicenseTierStateLicenseOutput values.
// You can construct a concrete instance of `LicenseTierStateLicenseInput` via:
//
//	LicenseTierStateLicenseArgs{...}
type LicenseTierStateLicenseInput interface {
	pulumi.Input

	ToLicenseTierStateLicenseOutput() LicenseTierStateLicenseOutput
	ToLicenseTierStateLicenseOutputWithContext(context.Context) LicenseTierStateLicenseOutput
}

type LicenseTierStateLicenseArgs struct {
	// License name
	//   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
	Name pulumi.StringInput `pulumi:"name"`
	// License status
	//   - Choices: `ENABLED`, `DISABLED`
	Status pulumi.StringInput `pulumi:"status"`
}

func (LicenseTierStateLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseTierStateLicense)(nil)).Elem()
}

func (i LicenseTierStateLicenseArgs) ToLicenseTierStateLicenseOutput() LicenseTierStateLicenseOutput {
	return i.ToLicenseTierStateLicenseOutputWithContext(context.Background())
}

func (i LicenseTierStateLicenseArgs) ToLicenseTierStateLicenseOutputWithContext(ctx context.Context) LicenseTierStateLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseTierStateLicenseOutput)
}

// LicenseTierStateLicenseArrayInput is an input type that accepts LicenseTierStateLicenseArray and LicenseTierStateLicenseArrayOutput values.
// You can construct a concrete instance of `LicenseTierStateLicenseArrayInput` via:
//
//	LicenseTierStateLicenseArray{ LicenseTierStateLicenseArgs{...} }
type LicenseTierStateLicenseArrayInput interface {
	pulumi.Input

	ToLicenseTierStateLicenseArrayOutput() LicenseTierStateLicenseArrayOutput
	ToLicenseTierStateLicenseArrayOutputWithContext(context.Context) LicenseTierStateLicenseArrayOutput
}

type LicenseTierStateLicenseArray []LicenseTierStateLicenseInput

func (LicenseTierStateLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LicenseTierStateLicense)(nil)).Elem()
}

func (i LicenseTierStateLicenseArray) ToLicenseTierStateLicenseArrayOutput() LicenseTierStateLicenseArrayOutput {
	return i.ToLicenseTierStateLicenseArrayOutputWithContext(context.Background())
}

func (i LicenseTierStateLicenseArray) ToLicenseTierStateLicenseArrayOutputWithContext(ctx context.Context) LicenseTierStateLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LicenseTierStateLicenseArrayOutput)
}

type LicenseTierStateLicenseOutput struct{ *pulumi.OutputState }

func (LicenseTierStateLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LicenseTierStateLicense)(nil)).Elem()
}

func (o LicenseTierStateLicenseOutput) ToLicenseTierStateLicenseOutput() LicenseTierStateLicenseOutput {
	return o
}

func (o LicenseTierStateLicenseOutput) ToLicenseTierStateLicenseOutputWithContext(ctx context.Context) LicenseTierStateLicenseOutput {
	return o
}

// License name
//   - Choices: `ESSENTIAL`, `ADVANTAGE`, `PREMIER`, `DEVICEADMIN`, `VM`
func (o LicenseTierStateLicenseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LicenseTierStateLicense) string { return v.Name }).(pulumi.StringOutput)
}

// License status
//   - Choices: `ENABLED`, `DISABLED`
func (o LicenseTierStateLicenseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LicenseTierStateLicense) string { return v.Status }).(pulumi.StringOutput)
}

type LicenseTierStateLicenseArrayOutput struct{ *pulumi.OutputState }

func (LicenseTierStateLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LicenseTierStateLicense)(nil)).Elem()
}

func (o LicenseTierStateLicenseArrayOutput) ToLicenseTierStateLicenseArrayOutput() LicenseTierStateLicenseArrayOutput {
	return o
}

func (o LicenseTierStateLicenseArrayOutput) ToLicenseTierStateLicenseArrayOutputWithContext(ctx context.Context) LicenseTierStateLicenseArrayOutput {
	return o
}

func (o LicenseTierStateLicenseArrayOutput) Index(i pulumi.IntInput) LicenseTierStateLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LicenseTierStateLicense {
		return vs[0].([]LicenseTierStateLicense)[vs[1].(int)]
	}).(LicenseTierStateLicenseOutput)
}

type GetLicenseTierStateLicense struct {
	// License name
	Name string `pulumi:"name"`
	// License status
	Status string `pulumi:"status"`
}

// GetLicenseTierStateLicenseInput is an input type that accepts GetLicenseTierStateLicenseArgs and GetLicenseTierStateLicenseOutput values.
// You can construct a concrete instance of `GetLicenseTierStateLicenseInput` via:
//
//	GetLicenseTierStateLicenseArgs{...}
type GetLicenseTierStateLicenseInput interface {
	pulumi.Input

	ToGetLicenseTierStateLicenseOutput() GetLicenseTierStateLicenseOutput
	ToGetLicenseTierStateLicenseOutputWithContext(context.Context) GetLicenseTierStateLicenseOutput
}

type GetLicenseTierStateLicenseArgs struct {
	// License name
	Name pulumi.StringInput `pulumi:"name"`
	// License status
	Status pulumi.StringInput `pulumi:"status"`
}

func (GetLicenseTierStateLicenseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLicenseTierStateLicense)(nil)).Elem()
}

func (i GetLicenseTierStateLicenseArgs) ToGetLicenseTierStateLicenseOutput() GetLicenseTierStateLicenseOutput {
	return i.ToGetLicenseTierStateLicenseOutputWithContext(context.Background())
}

func (i GetLicenseTierStateLicenseArgs) ToGetLicenseTierStateLicenseOutputWithContext(ctx context.Context) GetLicenseTierStateLicenseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLicenseTierStateLicenseOutput)
}

// GetLicenseTierStateLicenseArrayInput is an input type that accepts GetLicenseTierStateLicenseArray and GetLicenseTierStateLicenseArrayOutput values.
// You can construct a concrete instance of `GetLicenseTierStateLicenseArrayInput` via:
//
//	GetLicenseTierStateLicenseArray{ GetLicenseTierStateLicenseArgs{...} }
type GetLicenseTierStateLicenseArrayInput interface {
	pulumi.Input

	ToGetLicenseTierStateLicenseArrayOutput() GetLicenseTierStateLicenseArrayOutput
	ToGetLicenseTierStateLicenseArrayOutputWithContext(context.Context) GetLicenseTierStateLicenseArrayOutput
}

type GetLicenseTierStateLicenseArray []GetLicenseTierStateLicenseInput

func (GetLicenseTierStateLicenseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLicenseTierStateLicense)(nil)).Elem()
}

func (i GetLicenseTierStateLicenseArray) ToGetLicenseTierStateLicenseArrayOutput() GetLicenseTierStateLicenseArrayOutput {
	return i.ToGetLicenseTierStateLicenseArrayOutputWithContext(context.Background())
}

func (i GetLicenseTierStateLicenseArray) ToGetLicenseTierStateLicenseArrayOutputWithContext(ctx context.Context) GetLicenseTierStateLicenseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetLicenseTierStateLicenseArrayOutput)
}

type GetLicenseTierStateLicenseOutput struct{ *pulumi.OutputState }

func (GetLicenseTierStateLicenseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLicenseTierStateLicense)(nil)).Elem()
}

func (o GetLicenseTierStateLicenseOutput) ToGetLicenseTierStateLicenseOutput() GetLicenseTierStateLicenseOutput {
	return o
}

func (o GetLicenseTierStateLicenseOutput) ToGetLicenseTierStateLicenseOutputWithContext(ctx context.Context) GetLicenseTierStateLicenseOutput {
	return o
}

// License name
func (o GetLicenseTierStateLicenseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetLicenseTierStateLicense) string { return v.Name }).(pulumi.StringOutput)
}

// License status
func (o GetLicenseTierStateLicenseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetLicenseTierStateLicense) string { return v.Status }).(pulumi.StringOutput)
}

type GetLicenseTierStateLicenseArrayOutput struct{ *pulumi.OutputState }

func (GetLicenseTierStateLicenseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetLicenseTierStateLicense)(nil)).Elem()
}

func (o GetLicenseTierStateLicenseArrayOutput) ToGetLicenseTierStateLicenseArrayOutput() GetLicenseTierStateLicenseArrayOutput {
	return o
}

func (o GetLicenseTierStateLicenseArrayOutput) ToGetLicenseTierStateLicenseArrayOutputWithContext(ctx context.Context) GetLicenseTierStateLicenseArrayOutput {
	return o
}

func (o GetLicenseTierStateLicenseArrayOutput) Index(i pulumi.IntInput) GetLicenseTierStateLicenseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetLicenseTierStateLicense {
		return vs[0].([]GetLicenseTierStateLicense)[vs[1].(int)]
	}).(GetLicenseTierStateLicenseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTierStateLicenseInput)(nil)).Elem(), LicenseTierStateLicenseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LicenseTierStateLicenseArrayInput)(nil)).Elem(), LicenseTierStateLicenseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetLicenseTierStateLicenseInput)(nil)).Elem(), GetLicenseTierStateLicenseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetLicenseTierStateLicenseArrayInput)(nil)).Elem(), GetLicenseTierStateLicenseArray{})
	pulumi.RegisterOutputType(LicenseTierStateLicenseOutput{})
	pulumi.RegisterOutputType(LicenseTierStateLicenseArrayOutput{})
	pulumi.RegisterOutputType(GetLicenseTierStateLicenseOutput{})
	pulumi.RegisterOutputType(GetLicenseTierStateLicenseArrayOutput{})
}

// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ise

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This resource can manage a Repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ise.NewRepository(ctx, "example", &ise.RepositoryArgs{
//				Name:       pulumi.String("repo1"),
//				Protocol:   pulumi.String("SFTP"),
//				Path:       pulumi.String("/dir"),
//				ServerName: pulumi.String("server1"),
//				UserName:   pulumi.String("user9"),
//				Password:   pulumi.String("cisco123"),
//				EnablePki:  pulumi.Bool(false),
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
// $ pulumi import ise:index/repository:Repository example "repo1"
// ```
type Repository struct {
	pulumi.CustomResourceState

	// Enable PKI
	EnablePki pulumi.BoolPtrOutput `pulumi:"enablePki"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password can contain alphanumeric and/or special characters.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path pulumi.StringOutput `pulumi:"path"`
	// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Name of the server
	ServerName pulumi.StringPtrOutput `pulumi:"serverName"`
	// User name
	UserName pulumi.StringPtrOutput `pulumi:"userName"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Repository
	err := ctx.RegisterResource("ise:index/repository:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("ise:index/repository:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
	// Enable PKI
	EnablePki *bool `pulumi:"enablePki"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name *string `pulumi:"name"`
	// Password can contain alphanumeric and/or special characters.
	Password *string `pulumi:"password"`
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path *string `pulumi:"path"`
	// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
	Protocol *string `pulumi:"protocol"`
	// Name of the server
	ServerName *string `pulumi:"serverName"`
	// User name
	UserName *string `pulumi:"userName"`
}

type RepositoryState struct {
	// Enable PKI
	EnablePki pulumi.BoolPtrInput
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name pulumi.StringPtrInput
	// Password can contain alphanumeric and/or special characters.
	Password pulumi.StringPtrInput
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path pulumi.StringPtrInput
	// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
	Protocol pulumi.StringPtrInput
	// Name of the server
	ServerName pulumi.StringPtrInput
	// User name
	UserName pulumi.StringPtrInput
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// Enable PKI
	EnablePki *bool `pulumi:"enablePki"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name *string `pulumi:"name"`
	// Password can contain alphanumeric and/or special characters.
	Password *string `pulumi:"password"`
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path string `pulumi:"path"`
	// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
	Protocol string `pulumi:"protocol"`
	// Name of the server
	ServerName *string `pulumi:"serverName"`
	// User name
	UserName *string `pulumi:"userName"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// Enable PKI
	EnablePki pulumi.BoolPtrInput
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name pulumi.StringPtrInput
	// Password can contain alphanumeric and/or special characters.
	Password pulumi.StringPtrInput
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path pulumi.StringInput
	// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
	Protocol pulumi.StringInput
	// Name of the server
	ServerName pulumi.StringPtrInput
	// User name
	UserName pulumi.StringPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

// RepositoryArrayInput is an input type that accepts RepositoryArray and RepositoryArrayOutput values.
// You can construct a concrete instance of `RepositoryArrayInput` via:
//
//	RepositoryArray{ RepositoryArgs{...} }
type RepositoryArrayInput interface {
	pulumi.Input

	ToRepositoryArrayOutput() RepositoryArrayOutput
	ToRepositoryArrayOutputWithContext(context.Context) RepositoryArrayOutput
}

type RepositoryArray []RepositoryInput

func (RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (i RepositoryArray) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return i.ToRepositoryArrayOutputWithContext(context.Background())
}

func (i RepositoryArray) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryArrayOutput)
}

// RepositoryMapInput is an input type that accepts RepositoryMap and RepositoryMapOutput values.
// You can construct a concrete instance of `RepositoryMapInput` via:
//
//	RepositoryMap{ "key": RepositoryArgs{...} }
type RepositoryMapInput interface {
	pulumi.Input

	ToRepositoryMapOutput() RepositoryMapOutput
	ToRepositoryMapOutputWithContext(context.Context) RepositoryMapOutput
}

type RepositoryMap map[string]RepositoryInput

func (RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (i RepositoryMap) ToRepositoryMapOutput() RepositoryMapOutput {
	return i.ToRepositoryMapOutputWithContext(context.Background())
}

func (i RepositoryMap) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryMapOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

// Enable PKI
func (o RepositoryOutput) EnablePki() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.BoolPtrOutput { return v.EnablePki }).(pulumi.BoolPtrOutput)
}

// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
func (o RepositoryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Password can contain alphanumeric and/or special characters.
func (o RepositoryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
func (o RepositoryOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Protocol - Choices: `CDROM`, `DISK`, `FTP`, `HTTP`, `HTTPS`, `NFS`, `SFTP`, `TFTP`
func (o RepositoryOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Name of the server
func (o RepositoryOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.ServerName }).(pulumi.StringPtrOutput)
}

// User name
func (o RepositoryOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) pulumi.StringPtrOutput { return v.UserName }).(pulumi.StringPtrOutput)
}

type RepositoryArrayOutput struct{ *pulumi.OutputState }

func (RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Repository)(nil)).Elem()
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutput() RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) ToRepositoryArrayOutputWithContext(ctx context.Context) RepositoryArrayOutput {
	return o
}

func (o RepositoryArrayOutput) Index(i pulumi.IntInput) RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].([]*Repository)[vs[1].(int)]
	}).(RepositoryOutput)
}

type RepositoryMapOutput struct{ *pulumi.OutputState }

func (RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Repository)(nil)).Elem()
}

func (o RepositoryMapOutput) ToRepositoryMapOutput() RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) ToRepositoryMapOutputWithContext(ctx context.Context) RepositoryMapOutput {
	return o
}

func (o RepositoryMapOutput) MapIndex(k pulumi.StringInput) RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Repository {
		return vs[0].(map[string]*Repository)[vs[1].(string)]
	}).(RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryArrayInput)(nil)).Elem(), RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryMapInput)(nil)).Elem(), RepositoryMap{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryArrayOutput{})
	pulumi.RegisterOutputType(RepositoryMapOutput{})
}

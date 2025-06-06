// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package system

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Repository.
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
//			_, err := system.LookupRepository(ctx, &system.LookupRepositoryArgs{
//				Id: pulumi.StringRef("repo1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRepositoryResult
	err := ctx.Invoke("ise:system/getRepository:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getRepository.
type LookupRepositoryResult struct {
	// Enable PKI
	EnablePki bool `pulumi:"enablePki"`
	// The id of the object
	Id string `pulumi:"id"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name string `pulumi:"name"`
	// Password can contain alphanumeric and/or special characters.
	Password string `pulumi:"password"`
	// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
	Path string `pulumi:"path"`
	// Protocol
	Protocol string `pulumi:"protocol"`
	// Name of the server
	ServerName string `pulumi:"serverName"`
	// User name
	UserName string `pulumi:"userName"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResultOutput, error) {
			args := v.(LookupRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:system/getRepository:getRepository", args, LookupRepositoryResultOutput{}, options).(LookupRepositoryResultOutput), nil
		}).(LookupRepositoryResultOutput)
}

// A collection of arguments for invoking getRepository.
type LookupRepositoryOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRepository.
type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// Enable PKI
func (o LookupRepositoryResultOutput) EnablePki() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRepositoryResult) bool { return v.EnablePki }).(pulumi.BoolOutput)
}

// The id of the object
func (o LookupRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Repository name should be less than 80 characters and can contain alphanumeric, underscore, hyphen and dot characters.
func (o LookupRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// Password can contain alphanumeric and/or special characters.
func (o LookupRepositoryResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Password }).(pulumi.StringOutput)
}

// Path should always start with "/" and can contain alphanumeric, underscore, hyphen and dot characters.
func (o LookupRepositoryResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Path }).(pulumi.StringOutput)
}

// Protocol
func (o LookupRepositoryResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Name of the server
func (o LookupRepositoryResultOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.ServerName }).(pulumi.StringOutput)
}

// User name
func (o LookupRepositoryResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}

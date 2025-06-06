// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Network Access Dictionary.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/networkaccess"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkaccess.LookupDictionary(ctx, &networkaccess.LookupDictionaryArgs{
//				Id: pulumi.StringRef("Dict1"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDictionary(ctx *pulumi.Context, args *LookupDictionaryArgs, opts ...pulumi.InvokeOption) (*LookupDictionaryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDictionaryResult
	err := ctx.Invoke("ise:networkaccess/getDictionary:getDictionary", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDictionary.
type LookupDictionaryArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The dictionary name
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDictionary.
type LookupDictionaryResult struct {
	// The description of the dictionary
	Description string `pulumi:"description"`
	// The dictionary attribute type
	DictionaryAttrType string `pulumi:"dictionaryAttrType"`
	// The id of the object
	Id string `pulumi:"id"`
	// The dictionary name
	Name string `pulumi:"name"`
	// The version of the dictionary
	Version string `pulumi:"version"`
}

func LookupDictionaryOutput(ctx *pulumi.Context, args LookupDictionaryOutputArgs, opts ...pulumi.InvokeOption) LookupDictionaryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDictionaryResultOutput, error) {
			args := v.(LookupDictionaryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:networkaccess/getDictionary:getDictionary", args, LookupDictionaryResultOutput{}, options).(LookupDictionaryResultOutput), nil
		}).(LookupDictionaryResultOutput)
}

// A collection of arguments for invoking getDictionary.
type LookupDictionaryOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The dictionary name
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupDictionaryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDictionaryArgs)(nil)).Elem()
}

// A collection of values returned by getDictionary.
type LookupDictionaryResultOutput struct{ *pulumi.OutputState }

func (LookupDictionaryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDictionaryResult)(nil)).Elem()
}

func (o LookupDictionaryResultOutput) ToLookupDictionaryResultOutput() LookupDictionaryResultOutput {
	return o
}

func (o LookupDictionaryResultOutput) ToLookupDictionaryResultOutputWithContext(ctx context.Context) LookupDictionaryResultOutput {
	return o
}

// The description of the dictionary
func (o LookupDictionaryResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDictionaryResult) string { return v.Description }).(pulumi.StringOutput)
}

// The dictionary attribute type
func (o LookupDictionaryResultOutput) DictionaryAttrType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDictionaryResult) string { return v.DictionaryAttrType }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupDictionaryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDictionaryResult) string { return v.Id }).(pulumi.StringOutput)
}

// The dictionary name
func (o LookupDictionaryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDictionaryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The version of the dictionary
func (o LookupDictionaryResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDictionaryResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDictionaryResultOutput{})
}

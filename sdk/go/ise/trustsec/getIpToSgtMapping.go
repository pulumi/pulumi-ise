// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package trustsec

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the TrustSec IP to SGT Mapping.
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
//			_, err := trustsec.LookupIpToSgtMapping(ctx, &trustsec.LookupIpToSgtMappingArgs{
//				Id: pulumi.StringRef("76d24097-41c4-4558-a4d0-a8c07ac08470"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpToSgtMapping(ctx *pulumi.Context, args *LookupIpToSgtMappingArgs, opts ...pulumi.InvokeOption) (*LookupIpToSgtMappingResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpToSgtMappingResult
	err := ctx.Invoke("ise:trustsec/getIpToSgtMapping:getIpToSgtMapping", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpToSgtMapping.
type LookupIpToSgtMappingArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// The name of the IP to SGT mapping
	Name *string `pulumi:"name"`
}

// A collection of values returned by getIpToSgtMapping.
type LookupIpToSgtMappingResult struct {
	// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
	DeployTo string `pulumi:"deployTo"`
	// Deploy Type
	DeployType string `pulumi:"deployType"`
	// Description
	Description string `pulumi:"description"`
	// Mandatory if `hostName` is empty
	HostIp string `pulumi:"hostIp"`
	// Mandatory if `hostIp` is empty
	HostName string `pulumi:"hostName"`
	// The id of the object
	Id string `pulumi:"id"`
	// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
	MappingGroup string `pulumi:"mappingGroup"`
	// The name of the IP to SGT mapping
	Name string `pulumi:"name"`
	// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
	Sgt string `pulumi:"sgt"`
}

func LookupIpToSgtMappingOutput(ctx *pulumi.Context, args LookupIpToSgtMappingOutputArgs, opts ...pulumi.InvokeOption) LookupIpToSgtMappingResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIpToSgtMappingResultOutput, error) {
			args := v.(LookupIpToSgtMappingArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:trustsec/getIpToSgtMapping:getIpToSgtMapping", args, LookupIpToSgtMappingResultOutput{}, options).(LookupIpToSgtMappingResultOutput), nil
		}).(LookupIpToSgtMappingResultOutput)
}

// A collection of arguments for invoking getIpToSgtMapping.
type LookupIpToSgtMappingOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the IP to SGT mapping
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupIpToSgtMappingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpToSgtMappingArgs)(nil)).Elem()
}

// A collection of values returned by getIpToSgtMapping.
type LookupIpToSgtMappingResultOutput struct{ *pulumi.OutputState }

func (LookupIpToSgtMappingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpToSgtMappingResult)(nil)).Elem()
}

func (o LookupIpToSgtMappingResultOutput) ToLookupIpToSgtMappingResultOutput() LookupIpToSgtMappingResultOutput {
	return o
}

func (o LookupIpToSgtMappingResultOutput) ToLookupIpToSgtMappingResultOutputWithContext(ctx context.Context) LookupIpToSgtMappingResultOutput {
	return o
}

// Mandatory unless `mappingGroup` is set or unless `deployType` is `ALL`
func (o LookupIpToSgtMappingResultOutput) DeployTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.DeployTo }).(pulumi.StringOutput)
}

// Deploy Type
func (o LookupIpToSgtMappingResultOutput) DeployType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.DeployType }).(pulumi.StringOutput)
}

// Description
func (o LookupIpToSgtMappingResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.Description }).(pulumi.StringOutput)
}

// Mandatory if `hostName` is empty
func (o LookupIpToSgtMappingResultOutput) HostIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.HostIp }).(pulumi.StringOutput)
}

// Mandatory if `hostIp` is empty
func (o LookupIpToSgtMappingResultOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.HostName }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupIpToSgtMappingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.Id }).(pulumi.StringOutput)
}

// IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deployTo` and `deployType` are set
func (o LookupIpToSgtMappingResultOutput) MappingGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.MappingGroup }).(pulumi.StringOutput)
}

// The name of the IP to SGT mapping
func (o LookupIpToSgtMappingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.Name }).(pulumi.StringOutput)
}

// Trustsec Security Group ID. Mandatory unless `mappingGroup` is set
func (o LookupIpToSgtMappingResultOutput) Sgt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpToSgtMappingResult) string { return v.Sgt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpToSgtMappingResultOutput{})
}

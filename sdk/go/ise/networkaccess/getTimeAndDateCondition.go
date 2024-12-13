// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkaccess

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Network Access Time And Date Condition.
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
//			_, err := networkaccess.LookupTimeAndDateCondition(ctx, &networkaccess.LookupTimeAndDateConditionArgs{
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
func LookupTimeAndDateCondition(ctx *pulumi.Context, args *LookupTimeAndDateConditionArgs, opts ...pulumi.InvokeOption) (*LookupTimeAndDateConditionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupTimeAndDateConditionResult
	err := ctx.Invoke("ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTimeAndDateCondition.
type LookupTimeAndDateConditionArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Condition name
	Name *string `pulumi:"name"`
}

// A collection of values returned by getTimeAndDateCondition.
type LookupTimeAndDateConditionResult struct {
	// Condition description
	Description string `pulumi:"description"`
	// End date
	EndDate string `pulumi:"endDate"`
	// End time
	EndTime string `pulumi:"endTime"`
	// Exception end date
	ExceptionEndDate string `pulumi:"exceptionEndDate"`
	// Exception end time
	ExceptionEndTime string `pulumi:"exceptionEndTime"`
	// Exception start date
	ExceptionStartDate string `pulumi:"exceptionStartDate"`
	// Exception start time
	ExceptionStartTime string `pulumi:"exceptionStartTime"`
	// The id of the object
	Id string `pulumi:"id"`
	// Indicates whereas this condition is in negate mode
	IsNegate bool `pulumi:"isNegate"`
	// Condition name
	Name string `pulumi:"name"`
	// Start date
	StartDate string `pulumi:"startDate"`
	// Start time
	StartTime string `pulumi:"startTime"`
	// Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
	WeekDays []string `pulumi:"weekDays"`
	// Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
	WeekDaysExceptions []string `pulumi:"weekDaysExceptions"`
}

func LookupTimeAndDateConditionOutput(ctx *pulumi.Context, args LookupTimeAndDateConditionOutputArgs, opts ...pulumi.InvokeOption) LookupTimeAndDateConditionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTimeAndDateConditionResultOutput, error) {
			args := v.(LookupTimeAndDateConditionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition", args, LookupTimeAndDateConditionResultOutput{}, options).(LookupTimeAndDateConditionResultOutput), nil
		}).(LookupTimeAndDateConditionResultOutput)
}

// A collection of arguments for invoking getTimeAndDateCondition.
type LookupTimeAndDateConditionOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Condition name
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupTimeAndDateConditionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeAndDateConditionArgs)(nil)).Elem()
}

// A collection of values returned by getTimeAndDateCondition.
type LookupTimeAndDateConditionResultOutput struct{ *pulumi.OutputState }

func (LookupTimeAndDateConditionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTimeAndDateConditionResult)(nil)).Elem()
}

func (o LookupTimeAndDateConditionResultOutput) ToLookupTimeAndDateConditionResultOutput() LookupTimeAndDateConditionResultOutput {
	return o
}

func (o LookupTimeAndDateConditionResultOutput) ToLookupTimeAndDateConditionResultOutputWithContext(ctx context.Context) LookupTimeAndDateConditionResultOutput {
	return o
}

// Condition description
func (o LookupTimeAndDateConditionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.Description }).(pulumi.StringOutput)
}

// End date
func (o LookupTimeAndDateConditionResultOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.EndDate }).(pulumi.StringOutput)
}

// End time
func (o LookupTimeAndDateConditionResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Exception end date
func (o LookupTimeAndDateConditionResultOutput) ExceptionEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.ExceptionEndDate }).(pulumi.StringOutput)
}

// Exception end time
func (o LookupTimeAndDateConditionResultOutput) ExceptionEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.ExceptionEndTime }).(pulumi.StringOutput)
}

// Exception start date
func (o LookupTimeAndDateConditionResultOutput) ExceptionStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.ExceptionStartDate }).(pulumi.StringOutput)
}

// Exception start time
func (o LookupTimeAndDateConditionResultOutput) ExceptionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.ExceptionStartTime }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupTimeAndDateConditionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupTimeAndDateConditionResultOutput) IsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) bool { return v.IsNegate }).(pulumi.BoolOutput)
}

// Condition name
func (o LookupTimeAndDateConditionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Start date
func (o LookupTimeAndDateConditionResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.StartDate }).(pulumi.StringOutput)
}

// Start time
func (o LookupTimeAndDateConditionResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
func (o LookupTimeAndDateConditionResultOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

// Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
func (o LookupTimeAndDateConditionResultOutput) WeekDaysExceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupTimeAndDateConditionResult) []string { return v.WeekDaysExceptions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTimeAndDateConditionResultOutput{})
}

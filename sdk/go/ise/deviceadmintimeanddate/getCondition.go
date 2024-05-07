// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package deviceadmintimeanddate

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-ise/sdk/go/ise/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// This data source can read the Device Admin Time And Date Condition.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-ise/sdk/go/ise/DeviceAdminTimeAndDate"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := DeviceAdminTimeAndDate.GetCondition(ctx, &deviceadmintimeanddate.GetConditionArgs{
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
func LookupCondition(ctx *pulumi.Context, args *LookupConditionArgs, opts ...pulumi.InvokeOption) (*LookupConditionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupConditionResult
	err := ctx.Invoke("ise:DeviceAdminTimeAndDate/getCondition:getCondition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCondition.
type LookupConditionArgs struct {
	// The id of the object
	Id *string `pulumi:"id"`
	// Condition name
	Name *string `pulumi:"name"`
}

// A collection of values returned by getCondition.
type LookupConditionResult struct {
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

func LookupConditionOutput(ctx *pulumi.Context, args LookupConditionOutputArgs, opts ...pulumi.InvokeOption) LookupConditionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupConditionResult, error) {
			args := v.(LookupConditionArgs)
			r, err := LookupCondition(ctx, &args, opts...)
			var s LookupConditionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupConditionResultOutput)
}

// A collection of arguments for invoking getCondition.
type LookupConditionOutputArgs struct {
	// The id of the object
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Condition name
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupConditionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConditionArgs)(nil)).Elem()
}

// A collection of values returned by getCondition.
type LookupConditionResultOutput struct{ *pulumi.OutputState }

func (LookupConditionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupConditionResult)(nil)).Elem()
}

func (o LookupConditionResultOutput) ToLookupConditionResultOutput() LookupConditionResultOutput {
	return o
}

func (o LookupConditionResultOutput) ToLookupConditionResultOutputWithContext(ctx context.Context) LookupConditionResultOutput {
	return o
}

// Condition description
func (o LookupConditionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.Description }).(pulumi.StringOutput)
}

// End date
func (o LookupConditionResultOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.EndDate }).(pulumi.StringOutput)
}

// End time
func (o LookupConditionResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// Exception end date
func (o LookupConditionResultOutput) ExceptionEndDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.ExceptionEndDate }).(pulumi.StringOutput)
}

// Exception end time
func (o LookupConditionResultOutput) ExceptionEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.ExceptionEndTime }).(pulumi.StringOutput)
}

// Exception start date
func (o LookupConditionResultOutput) ExceptionStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.ExceptionStartDate }).(pulumi.StringOutput)
}

// Exception start time
func (o LookupConditionResultOutput) ExceptionStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.ExceptionStartTime }).(pulumi.StringOutput)
}

// The id of the object
func (o LookupConditionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whereas this condition is in negate mode
func (o LookupConditionResultOutput) IsNegate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupConditionResult) bool { return v.IsNegate }).(pulumi.BoolOutput)
}

// Condition name
func (o LookupConditionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Start date
func (o LookupConditionResultOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.StartDate }).(pulumi.StringOutput)
}

// Start time
func (o LookupConditionResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupConditionResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
func (o LookupConditionResultOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConditionResult) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

// Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
func (o LookupConditionResultOutput) WeekDaysExceptions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupConditionResult) []string { return v.WeekDaysExceptions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupConditionResultOutput{})
}

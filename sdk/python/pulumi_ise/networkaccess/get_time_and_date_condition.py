# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetTimeAndDateConditionResult',
    'AwaitableGetTimeAndDateConditionResult',
    'get_time_and_date_condition',
    'get_time_and_date_condition_output',
]

@pulumi.output_type
class GetTimeAndDateConditionResult:
    """
    A collection of values returned by getTimeAndDateCondition.
    """
    def __init__(__self__, description=None, end_date=None, end_time=None, exception_end_date=None, exception_end_time=None, exception_start_date=None, exception_start_time=None, id=None, is_negate=None, name=None, start_date=None, start_time=None, week_days=None, week_days_exceptions=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if end_date and not isinstance(end_date, str):
            raise TypeError("Expected argument 'end_date' to be a str")
        pulumi.set(__self__, "end_date", end_date)
        if end_time and not isinstance(end_time, str):
            raise TypeError("Expected argument 'end_time' to be a str")
        pulumi.set(__self__, "end_time", end_time)
        if exception_end_date and not isinstance(exception_end_date, str):
            raise TypeError("Expected argument 'exception_end_date' to be a str")
        pulumi.set(__self__, "exception_end_date", exception_end_date)
        if exception_end_time and not isinstance(exception_end_time, str):
            raise TypeError("Expected argument 'exception_end_time' to be a str")
        pulumi.set(__self__, "exception_end_time", exception_end_time)
        if exception_start_date and not isinstance(exception_start_date, str):
            raise TypeError("Expected argument 'exception_start_date' to be a str")
        pulumi.set(__self__, "exception_start_date", exception_start_date)
        if exception_start_time and not isinstance(exception_start_time, str):
            raise TypeError("Expected argument 'exception_start_time' to be a str")
        pulumi.set(__self__, "exception_start_time", exception_start_time)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_negate and not isinstance(is_negate, bool):
            raise TypeError("Expected argument 'is_negate' to be a bool")
        pulumi.set(__self__, "is_negate", is_negate)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if start_date and not isinstance(start_date, str):
            raise TypeError("Expected argument 'start_date' to be a str")
        pulumi.set(__self__, "start_date", start_date)
        if start_time and not isinstance(start_time, str):
            raise TypeError("Expected argument 'start_time' to be a str")
        pulumi.set(__self__, "start_time", start_time)
        if week_days and not isinstance(week_days, list):
            raise TypeError("Expected argument 'week_days' to be a list")
        pulumi.set(__self__, "week_days", week_days)
        if week_days_exceptions and not isinstance(week_days_exceptions, list):
            raise TypeError("Expected argument 'week_days_exceptions' to be a list")
        pulumi.set(__self__, "week_days_exceptions", week_days_exceptions)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Condition description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="endDate")
    def end_date(self) -> str:
        """
        End date
        """
        return pulumi.get(self, "end_date")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> str:
        """
        End time
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter(name="exceptionEndDate")
    def exception_end_date(self) -> str:
        """
        Exception end date
        """
        return pulumi.get(self, "exception_end_date")

    @property
    @pulumi.getter(name="exceptionEndTime")
    def exception_end_time(self) -> str:
        """
        Exception end time
        """
        return pulumi.get(self, "exception_end_time")

    @property
    @pulumi.getter(name="exceptionStartDate")
    def exception_start_date(self) -> str:
        """
        Exception start date
        """
        return pulumi.get(self, "exception_start_date")

    @property
    @pulumi.getter(name="exceptionStartTime")
    def exception_start_time(self) -> str:
        """
        Exception start time
        """
        return pulumi.get(self, "exception_start_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isNegate")
    def is_negate(self) -> bool:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "is_negate")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Condition name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="startDate")
    def start_date(self) -> str:
        """
        Start date
        """
        return pulumi.get(self, "start_date")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> str:
        """
        Start time
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter(name="weekDays")
    def week_days(self) -> Sequence[str]:
        """
        Defines for which days this condition will be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`. Default - List of all week days.
        """
        return pulumi.get(self, "week_days")

    @property
    @pulumi.getter(name="weekDaysExceptions")
    def week_days_exceptions(self) -> Sequence[str]:
        """
        Defines for which days this condition will NOT be matched. List of weekdays - `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`, `Sunday`.
        """
        return pulumi.get(self, "week_days_exceptions")


class AwaitableGetTimeAndDateConditionResult(GetTimeAndDateConditionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTimeAndDateConditionResult(
            description=self.description,
            end_date=self.end_date,
            end_time=self.end_time,
            exception_end_date=self.exception_end_date,
            exception_end_time=self.exception_end_time,
            exception_start_date=self.exception_start_date,
            exception_start_time=self.exception_start_time,
            id=self.id,
            is_negate=self.is_negate,
            name=self.name,
            start_date=self.start_date,
            start_time=self.start_time,
            week_days=self.week_days,
            week_days_exceptions=self.week_days_exceptions)


def get_time_and_date_condition(id: Optional[str] = None,
                                name: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTimeAndDateConditionResult:
    """
    This data source can read the Network Access Time And Date Condition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_time_and_date_condition(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Condition name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition', __args__, opts=opts, typ=GetTimeAndDateConditionResult).value

    return AwaitableGetTimeAndDateConditionResult(
        description=pulumi.get(__ret__, 'description'),
        end_date=pulumi.get(__ret__, 'end_date'),
        end_time=pulumi.get(__ret__, 'end_time'),
        exception_end_date=pulumi.get(__ret__, 'exception_end_date'),
        exception_end_time=pulumi.get(__ret__, 'exception_end_time'),
        exception_start_date=pulumi.get(__ret__, 'exception_start_date'),
        exception_start_time=pulumi.get(__ret__, 'exception_start_time'),
        id=pulumi.get(__ret__, 'id'),
        is_negate=pulumi.get(__ret__, 'is_negate'),
        name=pulumi.get(__ret__, 'name'),
        start_date=pulumi.get(__ret__, 'start_date'),
        start_time=pulumi.get(__ret__, 'start_time'),
        week_days=pulumi.get(__ret__, 'week_days'),
        week_days_exceptions=pulumi.get(__ret__, 'week_days_exceptions'))
def get_time_and_date_condition_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                       name: Optional[pulumi.Input[Optional[str]]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetTimeAndDateConditionResult]:
    """
    This data source can read the Network Access Time And Date Condition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_time_and_date_condition(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Condition name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:networkaccess/getTimeAndDateCondition:getTimeAndDateCondition', __args__, opts=opts, typ=GetTimeAndDateConditionResult)
    return __ret__.apply(lambda __response__: GetTimeAndDateConditionResult(
        description=pulumi.get(__response__, 'description'),
        end_date=pulumi.get(__response__, 'end_date'),
        end_time=pulumi.get(__response__, 'end_time'),
        exception_end_date=pulumi.get(__response__, 'exception_end_date'),
        exception_end_time=pulumi.get(__response__, 'exception_end_time'),
        exception_start_date=pulumi.get(__response__, 'exception_start_date'),
        exception_start_time=pulumi.get(__response__, 'exception_start_time'),
        id=pulumi.get(__response__, 'id'),
        is_negate=pulumi.get(__response__, 'is_negate'),
        name=pulumi.get(__response__, 'name'),
        start_date=pulumi.get(__response__, 'start_date'),
        start_time=pulumi.get(__response__, 'start_time'),
        week_days=pulumi.get(__response__, 'week_days'),
        week_days_exceptions=pulumi.get(__response__, 'week_days_exceptions')))

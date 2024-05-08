# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetConditionResult',
    'AwaitableGetConditionResult',
    'get_condition',
    'get_condition_output',
]

@pulumi.output_type
class GetConditionResult:
    """
    A collection of values returned by getCondition.
    """
    def __init__(__self__, attribute_name=None, attribute_value=None, childrens=None, condition_type=None, description=None, dictionary_name=None, dictionary_value=None, id=None, is_negate=None, name=None, operator=None):
        if attribute_name and not isinstance(attribute_name, str):
            raise TypeError("Expected argument 'attribute_name' to be a str")
        pulumi.set(__self__, "attribute_name", attribute_name)
        if attribute_value and not isinstance(attribute_value, str):
            raise TypeError("Expected argument 'attribute_value' to be a str")
        pulumi.set(__self__, "attribute_value", attribute_value)
        if childrens and not isinstance(childrens, list):
            raise TypeError("Expected argument 'childrens' to be a list")
        pulumi.set(__self__, "childrens", childrens)
        if condition_type and not isinstance(condition_type, str):
            raise TypeError("Expected argument 'condition_type' to be a str")
        pulumi.set(__self__, "condition_type", condition_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dictionary_name and not isinstance(dictionary_name, str):
            raise TypeError("Expected argument 'dictionary_name' to be a str")
        pulumi.set(__self__, "dictionary_name", dictionary_name)
        if dictionary_value and not isinstance(dictionary_value, str):
            raise TypeError("Expected argument 'dictionary_value' to be a str")
        pulumi.set(__self__, "dictionary_value", dictionary_value)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_negate and not isinstance(is_negate, bool):
            raise TypeError("Expected argument 'is_negate' to be a bool")
        pulumi.set(__self__, "is_negate", is_negate)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operator and not isinstance(operator, str):
            raise TypeError("Expected argument 'operator' to be a str")
        pulumi.set(__self__, "operator", operator)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> str:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "attribute_name")

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> str:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "attribute_value")

    @property
    @pulumi.getter
    def childrens(self) -> Sequence['outputs.GetConditionChildrenResult']:
        """
        List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> str:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        """
        return pulumi.get(self, "condition_type")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Condition description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dictionaryName")
    def dictionary_name(self) -> str:
        """
        Dictionary name
        """
        return pulumi.get(self, "dictionary_name")

    @property
    @pulumi.getter(name="dictionaryValue")
    def dictionary_value(self) -> str:
        """
        Dictionary value
        """
        return pulumi.get(self, "dictionary_value")

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
    @pulumi.getter
    def operator(self) -> str:
        """
        Equality operator
        """
        return pulumi.get(self, "operator")


class AwaitableGetConditionResult(GetConditionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConditionResult(
            attribute_name=self.attribute_name,
            attribute_value=self.attribute_value,
            childrens=self.childrens,
            condition_type=self.condition_type,
            description=self.description,
            dictionary_name=self.dictionary_name,
            dictionary_value=self.dictionary_value,
            id=self.id,
            is_negate=self.is_negate,
            name=self.name,
            operator=self.operator)


def get_condition(id: Optional[str] = None,
                  name: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConditionResult:
    """
    This data source can read the Device Admin Condition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_condition(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Condition name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:deviceadmin/getCondition:getCondition', __args__, opts=opts, typ=GetConditionResult).value

    return AwaitableGetConditionResult(
        attribute_name=pulumi.get(__ret__, 'attribute_name'),
        attribute_value=pulumi.get(__ret__, 'attribute_value'),
        childrens=pulumi.get(__ret__, 'childrens'),
        condition_type=pulumi.get(__ret__, 'condition_type'),
        description=pulumi.get(__ret__, 'description'),
        dictionary_name=pulumi.get(__ret__, 'dictionary_name'),
        dictionary_value=pulumi.get(__ret__, 'dictionary_value'),
        id=pulumi.get(__ret__, 'id'),
        is_negate=pulumi.get(__ret__, 'is_negate'),
        name=pulumi.get(__ret__, 'name'),
        operator=pulumi.get(__ret__, 'operator'))


@_utilities.lift_output_func(get_condition)
def get_condition_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                         name: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConditionResult]:
    """
    This data source can read the Device Admin Condition.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_condition(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Condition name
    """
    ...

# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ConditionChildrenArgs',
    'ConditionChildrenChildrenArgs',
]

@pulumi.input_type
class ConditionChildrenArgs:
    def __init__(__self__, *,
                 condition_type: pulumi.Input[str],
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 childrens: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenChildrenArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] condition_type: Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
                 - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[Sequence[pulumi.Input['ConditionChildrenChildrenArgs']]] childrens: List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[str] id: UUID for condition
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator
                 - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        pulumi.set(__self__, "condition_type", condition_type)
        if attribute_name is not None:
            pulumi.set(__self__, "attribute_name", attribute_name)
        if attribute_value is not None:
            pulumi.set(__self__, "attribute_value", attribute_value)
        if childrens is not None:
            pulumi.set(__self__, "childrens", childrens)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dictionary_name is not None:
            pulumi.set(__self__, "dictionary_name", dictionary_name)
        if dictionary_value is not None:
            pulumi.set(__self__, "dictionary_value", dictionary_value)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_negate is not None:
            pulumi.set(__self__, "is_negate", is_negate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> pulumi.Input[str]:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
          - Choices: `ConditionAndBlock`, `ConditionAttributes`, `ConditionOrBlock`, `ConditionReference`
        """
        return pulumi.get(self, "condition_type")

    @condition_type.setter
    def condition_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "condition_type", value)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> Optional[pulumi.Input[str]]:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "attribute_value")

    @attribute_value.setter
    def attribute_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_value", value)

    @property
    @pulumi.getter
    def childrens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenChildrenArgs']]]]:
        """
        List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @childrens.setter
    def childrens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenChildrenArgs']]]]):
        pulumi.set(self, "childrens", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Condition description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dictionaryName")
    def dictionary_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary name
        """
        return pulumi.get(self, "dictionary_name")

    @dictionary_name.setter
    def dictionary_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dictionary_name", value)

    @property
    @pulumi.getter(name="dictionaryValue")
    def dictionary_value(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary value
        """
        return pulumi.get(self, "dictionary_value")

    @dictionary_value.setter
    def dictionary_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dictionary_value", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID for condition
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="isNegate")
    def is_negate(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "is_negate")

    @is_negate.setter
    def is_negate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_negate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Condition name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        Equality operator
          - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)


@pulumi.input_type
class ConditionChildrenChildrenArgs:
    def __init__(__self__, *,
                 condition_type: pulumi.Input[str],
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 id: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] condition_type: Condition type.
                 - Choices: `ConditionAttributes`, `ConditionReference`
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[str] id: UUID for condition
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator
                 - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        pulumi.set(__self__, "condition_type", condition_type)
        if attribute_name is not None:
            pulumi.set(__self__, "attribute_name", attribute_name)
        if attribute_value is not None:
            pulumi.set(__self__, "attribute_value", attribute_value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dictionary_name is not None:
            pulumi.set(__self__, "dictionary_name", dictionary_name)
        if dictionary_value is not None:
            pulumi.set(__self__, "dictionary_value", dictionary_value)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if is_negate is not None:
            pulumi.set(__self__, "is_negate", is_negate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> pulumi.Input[str]:
        """
        Condition type.
          - Choices: `ConditionAttributes`, `ConditionReference`
        """
        return pulumi.get(self, "condition_type")

    @condition_type.setter
    def condition_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "condition_type", value)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> Optional[pulumi.Input[str]]:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "attribute_value")

    @attribute_value.setter
    def attribute_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Condition description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dictionaryName")
    def dictionary_name(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary name
        """
        return pulumi.get(self, "dictionary_name")

    @dictionary_name.setter
    def dictionary_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dictionary_name", value)

    @property
    @pulumi.getter(name="dictionaryValue")
    def dictionary_value(self) -> Optional[pulumi.Input[str]]:
        """
        Dictionary value
        """
        return pulumi.get(self, "dictionary_value")

    @dictionary_value.setter
    def dictionary_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dictionary_value", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID for condition
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter(name="isNegate")
    def is_negate(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "is_negate")

    @is_negate.setter
    def is_negate(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_negate", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Condition name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def operator(self) -> Optional[pulumi.Input[str]]:
        """
        Equality operator
          - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`, `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`, `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)



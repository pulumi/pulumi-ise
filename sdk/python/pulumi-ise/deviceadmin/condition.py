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
from ._inputs import *

__all__ = ['ConditionArgs', 'Condition']

@pulumi.input_type
class ConditionArgs:
    def __init__(__self__, *,
                 condition_type: pulumi.Input[str],
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 childrens: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Condition resource.
        :param pulumi.Input[str] condition_type: Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
               additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
               `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]] childrens: List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
               `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
               `notEquals`, `notIn`, `notStartsWith`, `startsWith`
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
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        `LibraryConditionAttributes`, `LibraryConditionOrBlock`
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
    def childrens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]]:
        """
        List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @childrens.setter
    def childrens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]]):
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
        Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)


@pulumi.input_type
class _ConditionState:
    def __init__(__self__, *,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 childrens: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]] = None,
                 condition_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Condition resources.
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]] childrens: List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        :param pulumi.Input[str] condition_type: Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
               additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
               `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
               `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
               `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        if attribute_name is not None:
            pulumi.set(__self__, "attribute_name", attribute_name)
        if attribute_value is not None:
            pulumi.set(__self__, "attribute_value", attribute_value)
        if childrens is not None:
            pulumi.set(__self__, "childrens", childrens)
        if condition_type is not None:
            pulumi.set(__self__, "condition_type", condition_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dictionary_name is not None:
            pulumi.set(__self__, "dictionary_name", dictionary_name)
        if dictionary_value is not None:
            pulumi.set(__self__, "dictionary_value", dictionary_value)
        if is_negate is not None:
            pulumi.set(__self__, "is_negate", is_negate)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if operator is not None:
            pulumi.set(__self__, "operator", operator)

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
    def childrens(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]]:
        """
        List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @childrens.setter
    def childrens(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ConditionChildrenArgs']]]]):
        pulumi.set(self, "childrens", value)

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> Optional[pulumi.Input[str]]:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        """
        return pulumi.get(self, "condition_type")

    @condition_type.setter
    def condition_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "condition_type", value)

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
        Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        return pulumi.get(self, "operator")

    @operator.setter
    def operator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "operator", value)


class Condition(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 childrens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionChildrenArgs']]]]] = None,
                 condition_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage a Device Admin Condition.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-ise as ise

        example = ise.device_admin.Condition("example",
            name="Cond1",
            description="My description",
            condition_type="LibraryConditionAttributes",
            is_negate=False,
            attribute_name="User",
            attribute_value="User1",
            dictionary_name="TACACS",
            operator="equals")
        ```

        ## Import

        ```sh
        $ pulumi import ise:DeviceAdmin/condition:Condition example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionChildrenArgs']]]] childrens: List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        :param pulumi.Input[str] condition_type: Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
               additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
               `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
               `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
               `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConditionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage a Device Admin Condition.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-ise as ise

        example = ise.device_admin.Condition("example",
            name="Cond1",
            description="My description",
            condition_type="LibraryConditionAttributes",
            is_negate=False,
            attribute_name="User",
            attribute_value="User1",
            dictionary_name="TACACS",
            operator="equals")
        ```

        ## Import

        ```sh
        $ pulumi import ise:DeviceAdmin/condition:Condition example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param ConditionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConditionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 childrens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionChildrenArgs']]]]] = None,
                 condition_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dictionary_name: Optional[pulumi.Input[str]] = None,
                 dictionary_value: Optional[pulumi.Input[str]] = None,
                 is_negate: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 operator: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConditionArgs.__new__(ConditionArgs)

            __props__.__dict__["attribute_name"] = attribute_name
            __props__.__dict__["attribute_value"] = attribute_value
            __props__.__dict__["childrens"] = childrens
            if condition_type is None and not opts.urn:
                raise TypeError("Missing required property 'condition_type'")
            __props__.__dict__["condition_type"] = condition_type
            __props__.__dict__["description"] = description
            __props__.__dict__["dictionary_name"] = dictionary_name
            __props__.__dict__["dictionary_value"] = dictionary_value
            __props__.__dict__["is_negate"] = is_negate
            __props__.__dict__["name"] = name
            __props__.__dict__["operator"] = operator
        super(Condition, __self__).__init__(
            'ise:DeviceAdmin/condition:Condition',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attribute_name: Optional[pulumi.Input[str]] = None,
            attribute_value: Optional[pulumi.Input[str]] = None,
            childrens: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionChildrenArgs']]]]] = None,
            condition_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            dictionary_name: Optional[pulumi.Input[str]] = None,
            dictionary_value: Optional[pulumi.Input[str]] = None,
            is_negate: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            operator: Optional[pulumi.Input[str]] = None) -> 'Condition':
        """
        Get an existing Condition resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_name: Dictionary attribute name
        :param pulumi.Input[str] attribute_value: Attribute value for condition. Value type is specified in dictionary object.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ConditionChildrenArgs']]]] childrens: List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        :param pulumi.Input[str] condition_type: Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
               additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
               `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        :param pulumi.Input[str] description: Condition description
        :param pulumi.Input[str] dictionary_name: Dictionary name
        :param pulumi.Input[str] dictionary_value: Dictionary value
        :param pulumi.Input[bool] is_negate: Indicates whereas this condition is in negate mode
        :param pulumi.Input[str] name: Condition name
        :param pulumi.Input[str] operator: Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
               `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
               `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConditionState.__new__(_ConditionState)

        __props__.__dict__["attribute_name"] = attribute_name
        __props__.__dict__["attribute_value"] = attribute_value
        __props__.__dict__["childrens"] = childrens
        __props__.__dict__["condition_type"] = condition_type
        __props__.__dict__["description"] = description
        __props__.__dict__["dictionary_name"] = dictionary_name
        __props__.__dict__["dictionary_value"] = dictionary_value
        __props__.__dict__["is_negate"] = is_negate
        __props__.__dict__["name"] = name
        __props__.__dict__["operator"] = operator
        return Condition(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> pulumi.Output[Optional[str]]:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "attribute_name")

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> pulumi.Output[Optional[str]]:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "attribute_value")

    @property
    @pulumi.getter
    def childrens(self) -> pulumi.Output[Optional[Sequence['outputs.ConditionChildren']]]:
        """
        List of child conditions. `condition_type` must be one of `LibraryConditionAndBlock` or `LibraryConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> pulumi.Output[str]:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that
        additional conditions are present under the children attribute. - Choices: `LibraryConditionAndBlock`,
        `LibraryConditionAttributes`, `LibraryConditionOrBlock`
        """
        return pulumi.get(self, "condition_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Condition description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dictionaryName")
    def dictionary_name(self) -> pulumi.Output[Optional[str]]:
        """
        Dictionary name
        """
        return pulumi.get(self, "dictionary_name")

    @property
    @pulumi.getter(name="dictionaryValue")
    def dictionary_value(self) -> pulumi.Output[Optional[str]]:
        """
        Dictionary value
        """
        return pulumi.get(self, "dictionary_value")

    @property
    @pulumi.getter(name="isNegate")
    def is_negate(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "is_negate")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Condition name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def operator(self) -> pulumi.Output[Optional[str]]:
        """
        Equality operator - Choices: `contains`, `endsWith`, `equals`, `greaterOrEquals`, `greaterThan`, `in`, `ipEquals`,
        `ipGreaterThan`, `ipLessThan`, `ipNotEquals`, `lessOrEquals`, `lessThan`, `matches`, `notContains`, `notEndsWith`,
        `notEquals`, `notIn`, `notStartsWith`, `startsWith`
        """
        return pulumi.get(self, "operator")


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
from . import outputs

__all__ = [
    'GetPolicySetResult',
    'AwaitableGetPolicySetResult',
    'get_policy_set',
    'get_policy_set_output',
]

@pulumi.output_type
class GetPolicySetResult:
    """
    A collection of values returned by getPolicySet.
    """
    def __init__(__self__, childrens=None, condition_attribute_name=None, condition_attribute_value=None, condition_dictionary_name=None, condition_dictionary_value=None, condition_id=None, condition_is_negate=None, condition_operator=None, condition_type=None, default=None, description=None, id=None, is_proxy=None, name=None, rank=None, service_name=None, state=None):
        if childrens and not isinstance(childrens, list):
            raise TypeError("Expected argument 'childrens' to be a list")
        pulumi.set(__self__, "childrens", childrens)
        if condition_attribute_name and not isinstance(condition_attribute_name, str):
            raise TypeError("Expected argument 'condition_attribute_name' to be a str")
        pulumi.set(__self__, "condition_attribute_name", condition_attribute_name)
        if condition_attribute_value and not isinstance(condition_attribute_value, str):
            raise TypeError("Expected argument 'condition_attribute_value' to be a str")
        pulumi.set(__self__, "condition_attribute_value", condition_attribute_value)
        if condition_dictionary_name and not isinstance(condition_dictionary_name, str):
            raise TypeError("Expected argument 'condition_dictionary_name' to be a str")
        pulumi.set(__self__, "condition_dictionary_name", condition_dictionary_name)
        if condition_dictionary_value and not isinstance(condition_dictionary_value, str):
            raise TypeError("Expected argument 'condition_dictionary_value' to be a str")
        pulumi.set(__self__, "condition_dictionary_value", condition_dictionary_value)
        if condition_id and not isinstance(condition_id, str):
            raise TypeError("Expected argument 'condition_id' to be a str")
        pulumi.set(__self__, "condition_id", condition_id)
        if condition_is_negate and not isinstance(condition_is_negate, bool):
            raise TypeError("Expected argument 'condition_is_negate' to be a bool")
        pulumi.set(__self__, "condition_is_negate", condition_is_negate)
        if condition_operator and not isinstance(condition_operator, str):
            raise TypeError("Expected argument 'condition_operator' to be a str")
        pulumi.set(__self__, "condition_operator", condition_operator)
        if condition_type and not isinstance(condition_type, str):
            raise TypeError("Expected argument 'condition_type' to be a str")
        pulumi.set(__self__, "condition_type", condition_type)
        if default and not isinstance(default, bool):
            raise TypeError("Expected argument 'default' to be a bool")
        pulumi.set(__self__, "default", default)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_proxy and not isinstance(is_proxy, bool):
            raise TypeError("Expected argument 'is_proxy' to be a bool")
        pulumi.set(__self__, "is_proxy", is_proxy)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rank and not isinstance(rank, int):
            raise TypeError("Expected argument 'rank' to be a int")
        pulumi.set(__self__, "rank", rank)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def childrens(self) -> Sequence['outputs.GetPolicySetChildrenResult']:
        """
        List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @property
    @pulumi.getter(name="conditionAttributeName")
    def condition_attribute_name(self) -> str:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "condition_attribute_name")

    @property
    @pulumi.getter(name="conditionAttributeValue")
    def condition_attribute_value(self) -> str:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "condition_attribute_value")

    @property
    @pulumi.getter(name="conditionDictionaryName")
    def condition_dictionary_name(self) -> str:
        """
        Dictionary name
        """
        return pulumi.get(self, "condition_dictionary_name")

    @property
    @pulumi.getter(name="conditionDictionaryValue")
    def condition_dictionary_value(self) -> str:
        """
        Dictionary value
        """
        return pulumi.get(self, "condition_dictionary_value")

    @property
    @pulumi.getter(name="conditionId")
    def condition_id(self) -> str:
        """
        UUID for condition
        """
        return pulumi.get(self, "condition_id")

    @property
    @pulumi.getter(name="conditionIsNegate")
    def condition_is_negate(self) -> bool:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "condition_is_negate")

    @property
    @pulumi.getter(name="conditionOperator")
    def condition_operator(self) -> str:
        """
        Equality operator
        """
        return pulumi.get(self, "condition_operator")

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> str:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        """
        return pulumi.get(self, "condition_type")

    @property
    @pulumi.getter
    def default(self) -> bool:
        """
        Indicates if this policy set is the default one
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the policy set
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isProxy")
    def is_proxy(self) -> bool:
        """
        Flag which indicates if the policy set service is of type 'Proxy Sequence' or 'Allowed Protocols'
        """
        return pulumi.get(self, "is_proxy")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def rank(self) -> int:
        """
        The rank (priority) in relation to other policy sets. Lower rank is higher priority.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        Policy set service identifier. 'Allowed Protocols' or 'Server Sequence'.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The state that the policy set is in. A disabled policy set cannot be matched.
        """
        return pulumi.get(self, "state")


class AwaitableGetPolicySetResult(GetPolicySetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicySetResult(
            childrens=self.childrens,
            condition_attribute_name=self.condition_attribute_name,
            condition_attribute_value=self.condition_attribute_value,
            condition_dictionary_name=self.condition_dictionary_name,
            condition_dictionary_value=self.condition_dictionary_value,
            condition_id=self.condition_id,
            condition_is_negate=self.condition_is_negate,
            condition_operator=self.condition_operator,
            condition_type=self.condition_type,
            default=self.default,
            description=self.description,
            id=self.id,
            is_proxy=self.is_proxy,
            name=self.name,
            rank=self.rank,
            service_name=self.service_name,
            state=self.state)


def get_policy_set(id: Optional[str] = None,
                   name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPolicySetResult:
    """
    This data source can read the Device Admin Policy Set.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_policy_set(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:deviceadmin/getPolicySet:getPolicySet', __args__, opts=opts, typ=GetPolicySetResult).value

    return AwaitableGetPolicySetResult(
        childrens=pulumi.get(__ret__, 'childrens'),
        condition_attribute_name=pulumi.get(__ret__, 'condition_attribute_name'),
        condition_attribute_value=pulumi.get(__ret__, 'condition_attribute_value'),
        condition_dictionary_name=pulumi.get(__ret__, 'condition_dictionary_name'),
        condition_dictionary_value=pulumi.get(__ret__, 'condition_dictionary_value'),
        condition_id=pulumi.get(__ret__, 'condition_id'),
        condition_is_negate=pulumi.get(__ret__, 'condition_is_negate'),
        condition_operator=pulumi.get(__ret__, 'condition_operator'),
        condition_type=pulumi.get(__ret__, 'condition_type'),
        default=pulumi.get(__ret__, 'default'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        is_proxy=pulumi.get(__ret__, 'is_proxy'),
        name=pulumi.get(__ret__, 'name'),
        rank=pulumi.get(__ret__, 'rank'),
        service_name=pulumi.get(__ret__, 'service_name'),
        state=pulumi.get(__ret__, 'state'))
def get_policy_set_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetPolicySetResult]:
    """
    This data source can read the Device Admin Policy Set.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_policy_set(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: Given name for the policy set, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:deviceadmin/getPolicySet:getPolicySet', __args__, opts=opts, typ=GetPolicySetResult)
    return __ret__.apply(lambda __response__: GetPolicySetResult(
        childrens=pulumi.get(__response__, 'childrens'),
        condition_attribute_name=pulumi.get(__response__, 'condition_attribute_name'),
        condition_attribute_value=pulumi.get(__response__, 'condition_attribute_value'),
        condition_dictionary_name=pulumi.get(__response__, 'condition_dictionary_name'),
        condition_dictionary_value=pulumi.get(__response__, 'condition_dictionary_value'),
        condition_id=pulumi.get(__response__, 'condition_id'),
        condition_is_negate=pulumi.get(__response__, 'condition_is_negate'),
        condition_operator=pulumi.get(__response__, 'condition_operator'),
        condition_type=pulumi.get(__response__, 'condition_type'),
        default=pulumi.get(__response__, 'default'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        is_proxy=pulumi.get(__response__, 'is_proxy'),
        name=pulumi.get(__response__, 'name'),
        rank=pulumi.get(__response__, 'rank'),
        service_name=pulumi.get(__response__, 'service_name'),
        state=pulumi.get(__response__, 'state')))

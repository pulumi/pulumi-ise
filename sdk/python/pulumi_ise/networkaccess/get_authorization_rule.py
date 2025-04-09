# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
    'GetAuthorizationRuleResult',
    'AwaitableGetAuthorizationRuleResult',
    'get_authorization_rule',
    'get_authorization_rule_output',
]

@pulumi.output_type
class GetAuthorizationRuleResult:
    """
    A collection of values returned by getAuthorizationRule.
    """
    def __init__(__self__, childrens=None, condition_attribute_name=None, condition_attribute_value=None, condition_dictionary_name=None, condition_dictionary_value=None, condition_id=None, condition_is_negate=None, condition_operator=None, condition_type=None, default=None, id=None, name=None, policy_set_id=None, profiles=None, rank=None, security_group=None, state=None):
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
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if policy_set_id and not isinstance(policy_set_id, str):
            raise TypeError("Expected argument 'policy_set_id' to be a str")
        pulumi.set(__self__, "policy_set_id", policy_set_id)
        if profiles and not isinstance(profiles, list):
            raise TypeError("Expected argument 'profiles' to be a list")
        pulumi.set(__self__, "profiles", profiles)
        if rank and not isinstance(rank, int):
            raise TypeError("Expected argument 'rank' to be a int")
        pulumi.set(__self__, "rank", rank)
        if security_group and not isinstance(security_group, str):
            raise TypeError("Expected argument 'security_group' to be a str")
        pulumi.set(__self__, "security_group", security_group)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def childrens(self) -> Sequence['outputs.GetAuthorizationRuleChildrenResult']:
        """
        List of child conditions. `condition_type` must be one of `ConditionAndBlock` or `ConditionOrBlock`.
        """
        return pulumi.get(self, "childrens")

    @property
    @pulumi.getter(name="conditionAttributeName")
    def condition_attribute_name(self) -> builtins.str:
        """
        Dictionary attribute name
        """
        return pulumi.get(self, "condition_attribute_name")

    @property
    @pulumi.getter(name="conditionAttributeValue")
    def condition_attribute_value(self) -> builtins.str:
        """
        Attribute value for condition. Value type is specified in dictionary object.
        """
        return pulumi.get(self, "condition_attribute_value")

    @property
    @pulumi.getter(name="conditionDictionaryName")
    def condition_dictionary_name(self) -> builtins.str:
        """
        Dictionary name
        """
        return pulumi.get(self, "condition_dictionary_name")

    @property
    @pulumi.getter(name="conditionDictionaryValue")
    def condition_dictionary_value(self) -> builtins.str:
        """
        Dictionary value
        """
        return pulumi.get(self, "condition_dictionary_value")

    @property
    @pulumi.getter(name="conditionId")
    def condition_id(self) -> builtins.str:
        """
        UUID for condition
        """
        return pulumi.get(self, "condition_id")

    @property
    @pulumi.getter(name="conditionIsNegate")
    def condition_is_negate(self) -> builtins.bool:
        """
        Indicates whereas this condition is in negate mode
        """
        return pulumi.get(self, "condition_is_negate")

    @property
    @pulumi.getter(name="conditionOperator")
    def condition_operator(self) -> builtins.str:
        """
        Equality operator
        """
        return pulumi.get(self, "condition_operator")

    @property
    @pulumi.getter(name="conditionType")
    def condition_type(self) -> builtins.str:
        """
        Indicates whether the record is the condition itself or a logical aggregation. Logical aggreation indicates that additional conditions are present under the children attribute.
        """
        return pulumi.get(self, "condition_type")

    @property
    @pulumi.getter
    def default(self) -> builtins.bool:
        """
        Indicates if this rule is the default one
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> builtins.str:
        """
        Policy set ID
        """
        return pulumi.get(self, "policy_set_id")

    @property
    @pulumi.getter
    def profiles(self) -> Sequence[builtins.str]:
        """
        The authorization profile(s)
        """
        return pulumi.get(self, "profiles")

    @property
    @pulumi.getter
    def rank(self) -> builtins.int:
        """
        The rank (priority) in relation to other rules. Lower rank is higher priority.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="securityGroup")
    def security_group(self) -> builtins.str:
        """
        Security group used in authorization policies
        """
        return pulumi.get(self, "security_group")

    @property
    @pulumi.getter
    def state(self) -> builtins.str:
        """
        The state that the rule is in. A disabled rule cannot be matched.
        """
        return pulumi.get(self, "state")


class AwaitableGetAuthorizationRuleResult(GetAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthorizationRuleResult(
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
            id=self.id,
            name=self.name,
            policy_set_id=self.policy_set_id,
            profiles=self.profiles,
            rank=self.rank,
            security_group=self.security_group,
            state=self.state)


def get_authorization_rule(id: Optional[builtins.str] = None,
                           name: Optional[builtins.str] = None,
                           policy_set_id: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthorizationRuleResult:
    """
    This data source can read the Network Access Authorization Rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_authorization_rule(id="76d24097-41c4-4558-a4d0-a8c07ac08470",
        policy_set_id="d82952cb-b901-4b09-b363-5ebf39bdbaf9")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
    :param builtins.str policy_set_id: Policy set ID
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['policySetId'] = policy_set_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:networkaccess/getAuthorizationRule:getAuthorizationRule', __args__, opts=opts, typ=GetAuthorizationRuleResult).value

    return AwaitableGetAuthorizationRuleResult(
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
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        policy_set_id=pulumi.get(__ret__, 'policy_set_id'),
        profiles=pulumi.get(__ret__, 'profiles'),
        rank=pulumi.get(__ret__, 'rank'),
        security_group=pulumi.get(__ret__, 'security_group'),
        state=pulumi.get(__ret__, 'state'))
def get_authorization_rule_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  policy_set_id: Optional[pulumi.Input[builtins.str]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAuthorizationRuleResult]:
    """
    This data source can read the Network Access Authorization Rule.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_authorization_rule(id="76d24097-41c4-4558-a4d0-a8c07ac08470",
        policy_set_id="d82952cb-b901-4b09-b363-5ebf39bdbaf9")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: Rule name, [Valid characters are alphanumerics, underscore, hyphen, space, period, parentheses]
    :param builtins.str policy_set_id: Policy set ID
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['policySetId'] = policy_set_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:networkaccess/getAuthorizationRule:getAuthorizationRule', __args__, opts=opts, typ=GetAuthorizationRuleResult)
    return __ret__.apply(lambda __response__: GetAuthorizationRuleResult(
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
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        policy_set_id=pulumi.get(__response__, 'policy_set_id'),
        profiles=pulumi.get(__response__, 'profiles'),
        rank=pulumi.get(__response__, 'rank'),
        security_group=pulumi.get(__response__, 'security_group'),
        state=pulumi.get(__response__, 'state')))

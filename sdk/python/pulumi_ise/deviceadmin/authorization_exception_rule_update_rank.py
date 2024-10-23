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

__all__ = ['AuthorizationExceptionRuleUpdateRankArgs', 'AuthorizationExceptionRuleUpdateRank']

@pulumi.input_type
class AuthorizationExceptionRuleUpdateRankArgs:
    def __init__(__self__, *,
                 policy_set_id: pulumi.Input[str],
                 rank: pulumi.Input[int],
                 rule_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AuthorizationExceptionRuleUpdateRank resource.
        :param pulumi.Input[str] policy_set_id: Policy set ID
        :param pulumi.Input[int] rank: The rank (priority) in relation to other rules. Lower rank is higher priority.
        :param pulumi.Input[str] rule_id: Authorization exception rule ID
        """
        pulumi.set(__self__, "policy_set_id", policy_set_id)
        pulumi.set(__self__, "rank", rank)
        pulumi.set(__self__, "rule_id", rule_id)

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Input[str]:
        """
        Policy set ID
        """
        return pulumi.get(self, "policy_set_id")

    @policy_set_id.setter
    def policy_set_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_set_id", value)

    @property
    @pulumi.getter
    def rank(self) -> pulumi.Input[int]:
        """
        The rank (priority) in relation to other rules. Lower rank is higher priority.
        """
        return pulumi.get(self, "rank")

    @rank.setter
    def rank(self, value: pulumi.Input[int]):
        pulumi.set(self, "rank", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Input[str]:
        """
        Authorization exception rule ID
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_id", value)


@pulumi.input_type
class _AuthorizationExceptionRuleUpdateRankState:
    def __init__(__self__, *,
                 policy_set_id: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthorizationExceptionRuleUpdateRank resources.
        :param pulumi.Input[str] policy_set_id: Policy set ID
        :param pulumi.Input[int] rank: The rank (priority) in relation to other rules. Lower rank is higher priority.
        :param pulumi.Input[str] rule_id: Authorization exception rule ID
        """
        if policy_set_id is not None:
            pulumi.set(__self__, "policy_set_id", policy_set_id)
        if rank is not None:
            pulumi.set(__self__, "rank", rank)
        if rule_id is not None:
            pulumi.set(__self__, "rule_id", rule_id)

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        Policy set ID
        """
        return pulumi.get(self, "policy_set_id")

    @policy_set_id.setter
    def policy_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_set_id", value)

    @property
    @pulumi.getter
    def rank(self) -> Optional[pulumi.Input[int]]:
        """
        The rank (priority) in relation to other rules. Lower rank is higher priority.
        """
        return pulumi.get(self, "rank")

    @rank.setter
    def rank(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rank", value)

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> Optional[pulumi.Input[str]]:
        """
        Authorization exception rule ID
        """
        return pulumi.get(self, "rule_id")

    @rule_id.setter
    def rule_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_id", value)


class AuthorizationExceptionRuleUpdateRank(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_set_id: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource is used to update rank field in device admin Authorization exception rule. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and device_admin_authorization_exception_rule resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.deviceadmin.AuthorizationExceptionRuleUpdateRank("example",
            rule_id="9b3680da-0165-44f6-9cff-88e778d98020",
            policy_set_id="d82952cb-b901-4b09-b363-5ebf39bdbaf9",
            rank=0)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_set_id: Policy set ID
        :param pulumi.Input[int] rank: The rank (priority) in relation to other rules. Lower rank is higher priority.
        :param pulumi.Input[str] rule_id: Authorization exception rule ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthorizationExceptionRuleUpdateRankArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource is used to update rank field in device admin Authorization exception rule. It serves as a workaround for the ISE API/Backend limitation which restricts rank assignments to a strictly incremental sequence. By utilizing this resource and device_admin_authorization_exception_rule resource, you can bypass the APIs limitation. Creation of this resource is performing PUT operation (Update) and it only tracks rank field. When this resource is destroyed, no action is performed on ISE and resource is just removed from state.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.deviceadmin.AuthorizationExceptionRuleUpdateRank("example",
            rule_id="9b3680da-0165-44f6-9cff-88e778d98020",
            policy_set_id="d82952cb-b901-4b09-b363-5ebf39bdbaf9",
            rank=0)
        ```

        :param str resource_name: The name of the resource.
        :param AuthorizationExceptionRuleUpdateRankArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthorizationExceptionRuleUpdateRankArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_set_id: Optional[pulumi.Input[str]] = None,
                 rank: Optional[pulumi.Input[int]] = None,
                 rule_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthorizationExceptionRuleUpdateRankArgs.__new__(AuthorizationExceptionRuleUpdateRankArgs)

            if policy_set_id is None and not opts.urn:
                raise TypeError("Missing required property 'policy_set_id'")
            __props__.__dict__["policy_set_id"] = policy_set_id
            if rank is None and not opts.urn:
                raise TypeError("Missing required property 'rank'")
            __props__.__dict__["rank"] = rank
            if rule_id is None and not opts.urn:
                raise TypeError("Missing required property 'rule_id'")
            __props__.__dict__["rule_id"] = rule_id
        super(AuthorizationExceptionRuleUpdateRank, __self__).__init__(
            'ise:deviceadmin/authorizationExceptionRuleUpdateRank:AuthorizationExceptionRuleUpdateRank',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_set_id: Optional[pulumi.Input[str]] = None,
            rank: Optional[pulumi.Input[int]] = None,
            rule_id: Optional[pulumi.Input[str]] = None) -> 'AuthorizationExceptionRuleUpdateRank':
        """
        Get an existing AuthorizationExceptionRuleUpdateRank resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_set_id: Policy set ID
        :param pulumi.Input[int] rank: The rank (priority) in relation to other rules. Lower rank is higher priority.
        :param pulumi.Input[str] rule_id: Authorization exception rule ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthorizationExceptionRuleUpdateRankState.__new__(_AuthorizationExceptionRuleUpdateRankState)

        __props__.__dict__["policy_set_id"] = policy_set_id
        __props__.__dict__["rank"] = rank
        __props__.__dict__["rule_id"] = rule_id
        return AuthorizationExceptionRuleUpdateRank(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policySetId")
    def policy_set_id(self) -> pulumi.Output[str]:
        """
        Policy set ID
        """
        return pulumi.get(self, "policy_set_id")

    @property
    @pulumi.getter
    def rank(self) -> pulumi.Output[int]:
        """
        The rank (priority) in relation to other rules. Lower rank is higher priority.
        """
        return pulumi.get(self, "rank")

    @property
    @pulumi.getter(name="ruleId")
    def rule_id(self) -> pulumi.Output[str]:
        """
        Authorization exception rule ID
        """
        return pulumi.get(self, "rule_id")


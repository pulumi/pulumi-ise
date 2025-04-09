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

__all__ = [
    'ActiveDirectoryAddGroupsGroup',
    'ActiveDirectoryJoinDomainWithAllNodesAdditionalData',
    'ActiveDirectoryJoinPointAttribute',
    'ActiveDirectoryJoinPointGroup',
    'ActiveDirectoryJoinPointRewriteRule',
    'IdentitySourceSequenceIdentitySource',
    'GetActiveDirectoryGroupsByDomainGroupResult',
    'GetActiveDirectoryJoinPointAttributeResult',
    'GetActiveDirectoryJoinPointGroupResult',
    'GetActiveDirectoryJoinPointRewriteRuleResult',
    'GetIdentitySourceSequenceIdentitySourceResult',
]

@pulumi.output_type
class ActiveDirectoryAddGroupsGroup(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 sid: builtins.str,
                 type: Optional[builtins.str] = None):
        """
        :param builtins.str name: Required for each group in the group list with no duplication between groups
        :param builtins.str sid: Required for each group in the group list with no duplication between groups
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sid", sid)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sid(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class ActiveDirectoryJoinDomainWithAllNodesAdditionalData(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 value: builtins.str):
        """
        :param builtins.str name: Additional attribute name
        :param builtins.str value: Additional attribute value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Additional attribute name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> builtins.str:
        """
        Additional attribute value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class ActiveDirectoryJoinPointAttribute(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "defaultValue":
            suggest = "default_value"
        elif key == "internalName":
            suggest = "internal_name"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ActiveDirectoryJoinPointAttribute. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ActiveDirectoryJoinPointAttribute.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ActiveDirectoryJoinPointAttribute.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 default_value: builtins.str,
                 internal_name: builtins.str,
                 name: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str default_value: Required for each attribute in the attribute list. Can contain an empty string.
        :param builtins.str internal_name: Required for each attribute in the attribute list
        :param builtins.str name: Required for each attribute in the attribute list with no duplication between attributes
        :param builtins.str type: Required for each group in the group list
                 - Choices: `STRING`, `IP`, `BOOLEAN`, `INT`, `OCTET_STRING`
        """
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "internal_name", internal_name)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> builtins.str:
        """
        Required for each attribute in the attribute list. Can contain an empty string.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="internalName")
    def internal_name(self) -> builtins.str:
        """
        Required for each attribute in the attribute list
        """
        return pulumi.get(self, "internal_name")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Required for each attribute in the attribute list with no duplication between attributes
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Required for each group in the group list
          - Choices: `STRING`, `IP`, `BOOLEAN`, `INT`, `OCTET_STRING`
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class ActiveDirectoryJoinPointGroup(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 sid: builtins.str,
                 type: Optional[builtins.str] = None):
        """
        :param builtins.str name: Required for each group in the group list with no duplication between groups
        :param builtins.str sid: Required for each group in the group list with no duplication between groups
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sid", sid)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sid(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


@pulumi.output_type
class ActiveDirectoryJoinPointRewriteRule(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "rewriteMatch":
            suggest = "rewrite_match"
        elif key == "rewriteResult":
            suggest = "rewrite_result"
        elif key == "rowId":
            suggest = "row_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ActiveDirectoryJoinPointRewriteRule. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ActiveDirectoryJoinPointRewriteRule.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ActiveDirectoryJoinPointRewriteRule.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 rewrite_match: builtins.str,
                 rewrite_result: builtins.str,
                 row_id: builtins.str):
        """
        :param builtins.str rewrite_match: Required for each rule in the list with no duplication between rules
        :param builtins.str rewrite_result: Required for each rule in the list
        :param builtins.str row_id: Required for each rule in the list in serial order
        """
        pulumi.set(__self__, "rewrite_match", rewrite_match)
        pulumi.set(__self__, "rewrite_result", rewrite_result)
        pulumi.set(__self__, "row_id", row_id)

    @property
    @pulumi.getter(name="rewriteMatch")
    def rewrite_match(self) -> builtins.str:
        """
        Required for each rule in the list with no duplication between rules
        """
        return pulumi.get(self, "rewrite_match")

    @property
    @pulumi.getter(name="rewriteResult")
    def rewrite_result(self) -> builtins.str:
        """
        Required for each rule in the list
        """
        return pulumi.get(self, "rewrite_result")

    @property
    @pulumi.getter(name="rowId")
    def row_id(self) -> builtins.str:
        """
        Required for each rule in the list in serial order
        """
        return pulumi.get(self, "row_id")


@pulumi.output_type
class IdentitySourceSequenceIdentitySource(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 order: builtins.int):
        """
        :param builtins.str name: Name of the identity source
        :param builtins.int order: Order of the identity source in the sequence
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the identity source
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> builtins.int:
        """
        Order of the identity source in the sequence
        """
        return pulumi.get(self, "order")


@pulumi.output_type
class GetActiveDirectoryGroupsByDomainGroupResult(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 sid: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str name: Group name
        :param builtins.str sid: Group SID
        :param builtins.str type: Group type
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sid", sid)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Group name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sid(self) -> builtins.str:
        """
        Group SID
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Group type
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetActiveDirectoryJoinPointAttributeResult(dict):
    def __init__(__self__, *,
                 default_value: builtins.str,
                 internal_name: builtins.str,
                 name: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str default_value: Required for each attribute in the attribute list. Can contain an empty string.
        :param builtins.str internal_name: Required for each attribute in the attribute list
        :param builtins.str name: Required for each attribute in the attribute list with no duplication between attributes
        :param builtins.str type: Required for each group in the group list
        """
        pulumi.set(__self__, "default_value", default_value)
        pulumi.set(__self__, "internal_name", internal_name)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="defaultValue")
    def default_value(self) -> builtins.str:
        """
        Required for each attribute in the attribute list. Can contain an empty string.
        """
        return pulumi.get(self, "default_value")

    @property
    @pulumi.getter(name="internalName")
    def internal_name(self) -> builtins.str:
        """
        Required for each attribute in the attribute list
        """
        return pulumi.get(self, "internal_name")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Required for each attribute in the attribute list with no duplication between attributes
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        """
        Required for each group in the group list
        """
        return pulumi.get(self, "type")


@pulumi.output_type
class GetActiveDirectoryJoinPointGroupResult(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 sid: builtins.str,
                 type: builtins.str):
        """
        :param builtins.str name: Required for each group in the group list with no duplication between groups
        :param builtins.str sid: Required for each group in the group list with no duplication between groups
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "sid", sid)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sid(self) -> builtins.str:
        """
        Required for each group in the group list with no duplication between groups
        """
        return pulumi.get(self, "sid")

    @property
    @pulumi.getter
    def type(self) -> builtins.str:
        return pulumi.get(self, "type")


@pulumi.output_type
class GetActiveDirectoryJoinPointRewriteRuleResult(dict):
    def __init__(__self__, *,
                 rewrite_match: builtins.str,
                 rewrite_result: builtins.str,
                 row_id: builtins.str):
        """
        :param builtins.str rewrite_match: Required for each rule in the list with no duplication between rules
        :param builtins.str rewrite_result: Required for each rule in the list
        :param builtins.str row_id: Required for each rule in the list in serial order
        """
        pulumi.set(__self__, "rewrite_match", rewrite_match)
        pulumi.set(__self__, "rewrite_result", rewrite_result)
        pulumi.set(__self__, "row_id", row_id)

    @property
    @pulumi.getter(name="rewriteMatch")
    def rewrite_match(self) -> builtins.str:
        """
        Required for each rule in the list with no duplication between rules
        """
        return pulumi.get(self, "rewrite_match")

    @property
    @pulumi.getter(name="rewriteResult")
    def rewrite_result(self) -> builtins.str:
        """
        Required for each rule in the list
        """
        return pulumi.get(self, "rewrite_result")

    @property
    @pulumi.getter(name="rowId")
    def row_id(self) -> builtins.str:
        """
        Required for each rule in the list in serial order
        """
        return pulumi.get(self, "row_id")


@pulumi.output_type
class GetIdentitySourceSequenceIdentitySourceResult(dict):
    def __init__(__self__, *,
                 name: builtins.str,
                 order: builtins.int):
        """
        :param builtins.str name: Name of the identity source
        :param builtins.int order: Order of the identity source in the sequence
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        Name of the identity source
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> builtins.int:
        """
        Order of the identity source in the sequence
        """
        return pulumi.get(self, "order")



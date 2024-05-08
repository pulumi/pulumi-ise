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
    'CommandSetCommand',
    'ProfileSessionAttribute',
    'GetCommandSetCommandResult',
    'GetProfileSessionAttributeResult',
]

@pulumi.output_type
class CommandSetCommand(dict):
    def __init__(__self__, *,
                 arguments: str,
                 command: str,
                 grant: str):
        """
        :param str arguments: Command arguments
        :param str command: Command
        :param str grant: Grant
                 - Choices: `PERMIT`, `DENY`, `DENY_ALWAYS`
        """
        pulumi.set(__self__, "arguments", arguments)
        pulumi.set(__self__, "command", command)
        pulumi.set(__self__, "grant", grant)

    @property
    @pulumi.getter
    def arguments(self) -> str:
        """
        Command arguments
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter
    def command(self) -> str:
        """
        Command
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def grant(self) -> str:
        """
        Grant
          - Choices: `PERMIT`, `DENY`, `DENY_ALWAYS`
        """
        return pulumi.get(self, "grant")


@pulumi.output_type
class ProfileSessionAttribute(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str,
                 value: str):
        """
        :param str name: Name
        :param str type: Type
                 - Choices: `MANDATORY`, `OPTIONAL`
        :param str value: Value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type
          - Choices: `MANDATORY`, `OPTIONAL`
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value
        """
        return pulumi.get(self, "value")


@pulumi.output_type
class GetCommandSetCommandResult(dict):
    def __init__(__self__, *,
                 arguments: str,
                 command: str,
                 grant: str):
        """
        :param str arguments: Command arguments
        :param str command: Command
        :param str grant: Grant
        """
        pulumi.set(__self__, "arguments", arguments)
        pulumi.set(__self__, "command", command)
        pulumi.set(__self__, "grant", grant)

    @property
    @pulumi.getter
    def arguments(self) -> str:
        """
        Command arguments
        """
        return pulumi.get(self, "arguments")

    @property
    @pulumi.getter
    def command(self) -> str:
        """
        Command
        """
        return pulumi.get(self, "command")

    @property
    @pulumi.getter
    def grant(self) -> str:
        """
        Grant
        """
        return pulumi.get(self, "grant")


@pulumi.output_type
class GetProfileSessionAttributeResult(dict):
    def __init__(__self__, *,
                 name: str,
                 type: str,
                 value: str):
        """
        :param str name: Name
        :param str type: Type
        :param str value: Value
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def value(self) -> str:
        """
        Value
        """
        return pulumi.get(self, "value")



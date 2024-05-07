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
    'SetCommand',
    'GetSetCommandResult',
]

@pulumi.output_type
class SetCommand(dict):
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
class GetSetCommandResult(dict):
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



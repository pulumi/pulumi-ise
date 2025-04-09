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
    'DeviceIpArgs',
    'DeviceIpArgsDict',
]

MYPY = False

if not MYPY:
    class DeviceIpArgsDict(TypedDict):
        ipaddress: pulumi.Input[builtins.str]
        """
        It can be either single ip address or ip range address
        """
        ipaddress_exclude: NotRequired[pulumi.Input[builtins.str]]
        """
        It can be either single ip address or ip range address
        """
        mask: NotRequired[pulumi.Input[builtins.str]]
        """
        Subnet mask length
        """
elif False:
    DeviceIpArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class DeviceIpArgs:
    def __init__(__self__, *,
                 ipaddress: pulumi.Input[builtins.str],
                 ipaddress_exclude: Optional[pulumi.Input[builtins.str]] = None,
                 mask: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] ipaddress: It can be either single ip address or ip range address
        :param pulumi.Input[builtins.str] ipaddress_exclude: It can be either single ip address or ip range address
        :param pulumi.Input[builtins.str] mask: Subnet mask length
        """
        pulumi.set(__self__, "ipaddress", ipaddress)
        if ipaddress_exclude is not None:
            pulumi.set(__self__, "ipaddress_exclude", ipaddress_exclude)
        if mask is not None:
            pulumi.set(__self__, "mask", mask)

    @property
    @pulumi.getter
    def ipaddress(self) -> pulumi.Input[builtins.str]:
        """
        It can be either single ip address or ip range address
        """
        return pulumi.get(self, "ipaddress")

    @ipaddress.setter
    def ipaddress(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "ipaddress", value)

    @property
    @pulumi.getter(name="ipaddressExclude")
    def ipaddress_exclude(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        It can be either single ip address or ip range address
        """
        return pulumi.get(self, "ipaddress_exclude")

    @ipaddress_exclude.setter
    def ipaddress_exclude(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ipaddress_exclude", value)

    @property
    @pulumi.getter
    def mask(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Subnet mask length
        """
        return pulumi.get(self, "mask")

    @mask.setter
    def mask(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "mask", value)



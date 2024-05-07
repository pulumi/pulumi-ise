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
    'SequenceIdentitySource',
    'GetSequenceIdentitySourceResult',
]

@pulumi.output_type
class SequenceIdentitySource(dict):
    def __init__(__self__, *,
                 name: str,
                 order: int):
        """
        :param str name: Name of the identity source
        :param int order: Order of the identity source in the sequence
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the identity source
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        Order of the identity source in the sequence
        """
        return pulumi.get(self, "order")


@pulumi.output_type
class GetSequenceIdentitySourceResult(dict):
    def __init__(__self__, *,
                 name: str,
                 order: int):
        """
        :param str name: Name of the identity source
        :param int order: Order of the identity source in the sequence
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "order", order)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the identity source
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def order(self) -> int:
        """
        Order of the identity source in the sequence
        """
        return pulumi.get(self, "order")



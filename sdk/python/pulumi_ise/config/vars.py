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

import types

__config__ = pulumi.Config('ise')


class _ExportableConfig(types.ModuleType):
    @property
    def insecure(self) -> Optional[bool]:
        """
        Allow insecure HTTPS client. This can also be set as the ISE_INSECURE environment variable. Defaults to `true`.
        """
        return __config__.get_bool('insecure')

    @property
    def password(self) -> Optional[str]:
        """
        Password for the ISE instance. This can also be set as the ISE_PASSWORD environment variable.
        """
        return __config__.get('password')

    @property
    def retries(self) -> Optional[int]:
        """
        Number of retries for REST API calls. This can also be set as the ISE_RETRIES environment variable. Defaults to `3`.
        """
        return __config__.get_int('retries')

    @property
    def url(self) -> Optional[str]:
        """
        URL of the Cisco ISE instance. This can also be set as the ISE_URL environment variable.
        """
        return __config__.get('url')

    @property
    def username(self) -> Optional[str]:
        """
        Username for the ISE instance. This can also be set as the ISE_USERNAME environment variable.
        """
        return __config__.get('username')

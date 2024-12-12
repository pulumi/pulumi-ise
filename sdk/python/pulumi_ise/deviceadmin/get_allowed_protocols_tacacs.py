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

__all__ = [
    'GetAllowedProtocolsTacacsResult',
    'AwaitableGetAllowedProtocolsTacacsResult',
    'get_allowed_protocols_tacacs',
    'get_allowed_protocols_tacacs_output',
]

@pulumi.output_type
class GetAllowedProtocolsTacacsResult:
    """
    A collection of values returned by getAllowedProtocolsTacacs.
    """
    def __init__(__self__, allow_chap=None, allow_ms_chap_v1=None, allow_pap_ascii=None, description=None, id=None, name=None):
        if allow_chap and not isinstance(allow_chap, bool):
            raise TypeError("Expected argument 'allow_chap' to be a bool")
        pulumi.set(__self__, "allow_chap", allow_chap)
        if allow_ms_chap_v1 and not isinstance(allow_ms_chap_v1, bool):
            raise TypeError("Expected argument 'allow_ms_chap_v1' to be a bool")
        pulumi.set(__self__, "allow_ms_chap_v1", allow_ms_chap_v1)
        if allow_pap_ascii and not isinstance(allow_pap_ascii, bool):
            raise TypeError("Expected argument 'allow_pap_ascii' to be a bool")
        pulumi.set(__self__, "allow_pap_ascii", allow_pap_ascii)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="allowChap")
    def allow_chap(self) -> bool:
        """
        Allow CHAP
        """
        return pulumi.get(self, "allow_chap")

    @property
    @pulumi.getter(name="allowMsChapV1")
    def allow_ms_chap_v1(self) -> bool:
        """
        Allow MS CHAP v1
        """
        return pulumi.get(self, "allow_ms_chap_v1")

    @property
    @pulumi.getter(name="allowPapAscii")
    def allow_pap_ascii(self) -> bool:
        """
        Allow PAP ASCII
        """
        return pulumi.get(self, "allow_pap_ascii")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description
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
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the allowed protocols
        """
        return pulumi.get(self, "name")


class AwaitableGetAllowedProtocolsTacacsResult(GetAllowedProtocolsTacacsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAllowedProtocolsTacacsResult(
            allow_chap=self.allow_chap,
            allow_ms_chap_v1=self.allow_ms_chap_v1,
            allow_pap_ascii=self.allow_pap_ascii,
            description=self.description,
            id=self.id,
            name=self.name)


def get_allowed_protocols_tacacs(id: Optional[str] = None,
                                 name: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAllowedProtocolsTacacsResult:
    """
    This data source can read a TACACS allowed protocols policy element.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_allowed_protocols_tacacs(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the allowed protocols
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:deviceadmin/getAllowedProtocolsTacacs:getAllowedProtocolsTacacs', __args__, opts=opts, typ=GetAllowedProtocolsTacacsResult).value

    return AwaitableGetAllowedProtocolsTacacsResult(
        allow_chap=pulumi.get(__ret__, 'allow_chap'),
        allow_ms_chap_v1=pulumi.get(__ret__, 'allow_ms_chap_v1'),
        allow_pap_ascii=pulumi.get(__ret__, 'allow_pap_ascii'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_allowed_protocols_tacacs_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                        name: Optional[pulumi.Input[Optional[str]]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAllowedProtocolsTacacsResult]:
    """
    This data source can read a TACACS allowed protocols policy element.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.deviceadmin.get_allowed_protocols_tacacs(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the allowed protocols
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:deviceadmin/getAllowedProtocolsTacacs:getAllowedProtocolsTacacs', __args__, opts=opts, typ=GetAllowedProtocolsTacacsResult)
    return __ret__.apply(lambda __response__: GetAllowedProtocolsTacacsResult(
        allow_chap=pulumi.get(__response__, 'allow_chap'),
        allow_ms_chap_v1=pulumi.get(__response__, 'allow_ms_chap_v1'),
        allow_pap_ascii=pulumi.get(__response__, 'allow_pap_ascii'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
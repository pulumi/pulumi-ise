# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetProfileResult',
    'AwaitableGetProfileResult',
    'get_profile',
    'get_profile_output',
]

@pulumi.output_type
class GetProfileResult:
    """
    A collection of values returned by getProfile.
    """
    def __init__(__self__, description=None, id=None, name=None, session_attributes=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if session_attributes and not isinstance(session_attributes, list):
            raise TypeError("Expected argument 'session_attributes' to be a list")
        pulumi.set(__self__, "session_attributes", session_attributes)

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
        The name of the TACACS profile
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sessionAttributes")
    def session_attributes(self) -> Sequence['outputs.GetProfileSessionAttributeResult']:
        return pulumi.get(self, "session_attributes")


class AwaitableGetProfileResult(GetProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProfileResult(
            description=self.description,
            id=self.id,
            name=self.name,
            session_attributes=self.session_attributes)


def get_profile(id: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProfileResult:
    """
    This data source can read the TACACS Profile.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.tacacs.get_profile(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the TACACS profile
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:tacacs/getProfile:getProfile', __args__, opts=opts, typ=GetProfileResult).value

    return AwaitableGetProfileResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        session_attributes=pulumi.get(__ret__, 'session_attributes'))


@_utilities.lift_output_func(get_profile)
def get_profile_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                       name: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProfileResult]:
    """
    This data source can read the TACACS Profile.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.tacacs.get_profile(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the TACACS profile
    """
    ...

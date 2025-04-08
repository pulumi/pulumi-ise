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
    'GetIdentitySourceSequenceResult',
    'AwaitableGetIdentitySourceSequenceResult',
    'get_identity_source_sequence',
    'get_identity_source_sequence_output',
]

@pulumi.output_type
class GetIdentitySourceSequenceResult:
    """
    A collection of values returned by getIdentitySourceSequence.
    """
    def __init__(__self__, break_on_store_fail=None, certificate_authentication_profile=None, description=None, id=None, identity_sources=None, name=None):
        if break_on_store_fail and not isinstance(break_on_store_fail, bool):
            raise TypeError("Expected argument 'break_on_store_fail' to be a bool")
        pulumi.set(__self__, "break_on_store_fail", break_on_store_fail)
        if certificate_authentication_profile and not isinstance(certificate_authentication_profile, str):
            raise TypeError("Expected argument 'certificate_authentication_profile' to be a str")
        pulumi.set(__self__, "certificate_authentication_profile", certificate_authentication_profile)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_sources and not isinstance(identity_sources, list):
            raise TypeError("Expected argument 'identity_sources' to be a list")
        pulumi.set(__self__, "identity_sources", identity_sources)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="breakOnStoreFail")
    def break_on_store_fail(self) -> builtins.bool:
        """
        Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        """
        return pulumi.get(self, "break_on_store_fail")

    @property
    @pulumi.getter(name="certificateAuthenticationProfile")
    def certificate_authentication_profile(self) -> builtins.str:
        """
        Certificate Authentication Profile, empty if doesn't exist
        """
        return pulumi.get(self, "certificate_authentication_profile")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> Sequence['outputs.GetIdentitySourceSequenceIdentitySourceResult']:
        return pulumi.get(self, "identity_sources")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the identity source sequence
        """
        return pulumi.get(self, "name")


class AwaitableGetIdentitySourceSequenceResult(GetIdentitySourceSequenceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIdentitySourceSequenceResult(
            break_on_store_fail=self.break_on_store_fail,
            certificate_authentication_profile=self.certificate_authentication_profile,
            description=self.description,
            id=self.id,
            identity_sources=self.identity_sources,
            name=self.name)


def get_identity_source_sequence(id: Optional[builtins.str] = None,
                                 name: Optional[builtins.str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIdentitySourceSequenceResult:
    """
    This data source can read the Identity Source Sequence.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_identity_source_sequence(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the identity source sequence
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:identitymanagement/getIdentitySourceSequence:getIdentitySourceSequence', __args__, opts=opts, typ=GetIdentitySourceSequenceResult).value

    return AwaitableGetIdentitySourceSequenceResult(
        break_on_store_fail=pulumi.get(__ret__, 'break_on_store_fail'),
        certificate_authentication_profile=pulumi.get(__ret__, 'certificate_authentication_profile'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        identity_sources=pulumi.get(__ret__, 'identity_sources'),
        name=pulumi.get(__ret__, 'name'))
def get_identity_source_sequence_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                        name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIdentitySourceSequenceResult]:
    """
    This data source can read the Identity Source Sequence.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_identity_source_sequence(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the identity source sequence
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:identitymanagement/getIdentitySourceSequence:getIdentitySourceSequence', __args__, opts=opts, typ=GetIdentitySourceSequenceResult)
    return __ret__.apply(lambda __response__: GetIdentitySourceSequenceResult(
        break_on_store_fail=pulumi.get(__response__, 'break_on_store_fail'),
        certificate_authentication_profile=pulumi.get(__response__, 'certificate_authentication_profile'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        identity_sources=pulumi.get(__response__, 'identity_sources'),
        name=pulumi.get(__response__, 'name')))

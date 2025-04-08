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
    'GetDownloadableAclResult',
    'AwaitableGetDownloadableAclResult',
    'get_downloadable_acl',
    'get_downloadable_acl_output',
]

@pulumi.output_type
class GetDownloadableAclResult:
    """
    A collection of values returned by getDownloadableAcl.
    """
    def __init__(__self__, dacl=None, dacl_type=None, description=None, id=None, name=None):
        if dacl and not isinstance(dacl, str):
            raise TypeError("Expected argument 'dacl' to be a str")
        pulumi.set(__self__, "dacl", dacl)
        if dacl_type and not isinstance(dacl_type, str):
            raise TypeError("Expected argument 'dacl_type' to be a str")
        pulumi.set(__self__, "dacl_type", dacl_type)
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
    @pulumi.getter
    def dacl(self) -> builtins.str:
        """
        The DACL content
        """
        return pulumi.get(self, "dacl")

    @property
    @pulumi.getter(name="daclType")
    def dacl_type(self) -> builtins.str:
        """
        The type of ACL
        """
        return pulumi.get(self, "dacl_type")

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
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the downloadable ACL
        """
        return pulumi.get(self, "name")


class AwaitableGetDownloadableAclResult(GetDownloadableAclResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDownloadableAclResult(
            dacl=self.dacl,
            dacl_type=self.dacl_type,
            description=self.description,
            id=self.id,
            name=self.name)


def get_downloadable_acl(id: Optional[builtins.str] = None,
                         name: Optional[builtins.str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDownloadableAclResult:
    """
    This data source can read the Downloadable ACL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_downloadable_acl(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the downloadable ACL
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:networkaccess/getDownloadableAcl:getDownloadableAcl', __args__, opts=opts, typ=GetDownloadableAclResult).value

    return AwaitableGetDownloadableAclResult(
        dacl=pulumi.get(__ret__, 'dacl'),
        dacl_type=pulumi.get(__ret__, 'dacl_type'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_downloadable_acl_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDownloadableAclResult]:
    """
    This data source can read the Downloadable ACL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_downloadable_acl(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the downloadable ACL
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:networkaccess/getDownloadableAcl:getDownloadableAcl', __args__, opts=opts, typ=GetDownloadableAclResult)
    return __ret__.apply(lambda __response__: GetDownloadableAclResult(
        dacl=pulumi.get(__response__, 'dacl'),
        dacl_type=pulumi.get(__response__, 'dacl_type'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))

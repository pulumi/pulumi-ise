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
    'GetEndpointIdentityGroupResult',
    'AwaitableGetEndpointIdentityGroupResult',
    'get_endpoint_identity_group',
    'get_endpoint_identity_group_output',
]

@pulumi.output_type
class GetEndpointIdentityGroupResult:
    """
    A collection of values returned by getEndpointIdentityGroup.
    """
    def __init__(__self__, description=None, id=None, name=None, parent_endpoint_identity_group_id=None, system_defined=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_endpoint_identity_group_id and not isinstance(parent_endpoint_identity_group_id, str):
            raise TypeError("Expected argument 'parent_endpoint_identity_group_id' to be a str")
        pulumi.set(__self__, "parent_endpoint_identity_group_id", parent_endpoint_identity_group_id)
        if system_defined and not isinstance(system_defined, bool):
            raise TypeError("Expected argument 'system_defined' to be a bool")
        pulumi.set(__self__, "system_defined", system_defined)

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
        The name of the endpoint identity group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentEndpointIdentityGroupId")
    def parent_endpoint_identity_group_id(self) -> builtins.str:
        """
        Parent endpoint identity group ID
        """
        return pulumi.get(self, "parent_endpoint_identity_group_id")

    @property
    @pulumi.getter(name="systemDefined")
    def system_defined(self) -> builtins.bool:
        """
        System defined endpoint identity group
        """
        return pulumi.get(self, "system_defined")


class AwaitableGetEndpointIdentityGroupResult(GetEndpointIdentityGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointIdentityGroupResult(
            description=self.description,
            id=self.id,
            name=self.name,
            parent_endpoint_identity_group_id=self.parent_endpoint_identity_group_id,
            system_defined=self.system_defined)


def get_endpoint_identity_group(id: Optional[builtins.str] = None,
                                name: Optional[builtins.str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointIdentityGroupResult:
    """
    This data source can read the Endpoint Identity Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_endpoint_identity_group(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the endpoint identity group
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:identitymanagement/getEndpointIdentityGroup:getEndpointIdentityGroup', __args__, opts=opts, typ=GetEndpointIdentityGroupResult).value

    return AwaitableGetEndpointIdentityGroupResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        parent_endpoint_identity_group_id=pulumi.get(__ret__, 'parent_endpoint_identity_group_id'),
        system_defined=pulumi.get(__ret__, 'system_defined'))
def get_endpoint_identity_group_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                       name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                       opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEndpointIdentityGroupResult]:
    """
    This data source can read the Endpoint Identity Group.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_endpoint_identity_group(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the endpoint identity group
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:identitymanagement/getEndpointIdentityGroup:getEndpointIdentityGroup', __args__, opts=opts, typ=GetEndpointIdentityGroupResult)
    return __ret__.apply(lambda __response__: GetEndpointIdentityGroupResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        parent_endpoint_identity_group_id=pulumi.get(__response__, 'parent_endpoint_identity_group_id'),
        system_defined=pulumi.get(__response__, 'system_defined')))

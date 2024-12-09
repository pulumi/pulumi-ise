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
    'GetSecurityGroupAclResult',
    'AwaitableGetSecurityGroupAclResult',
    'get_security_group_acl',
    'get_security_group_acl_output',
]

@pulumi.output_type
class GetSecurityGroupAclResult:
    """
    A collection of values returned by getSecurityGroupAcl.
    """
    def __init__(__self__, acl_content=None, description=None, id=None, ip_version=None, name=None, read_only=None):
        if acl_content and not isinstance(acl_content, str):
            raise TypeError("Expected argument 'acl_content' to be a str")
        pulumi.set(__self__, "acl_content", acl_content)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_version and not isinstance(ip_version, str):
            raise TypeError("Expected argument 'ip_version' to be a str")
        pulumi.set(__self__, "ip_version", ip_version)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if read_only and not isinstance(read_only, bool):
            raise TypeError("Expected argument 'read_only' to be a bool")
        pulumi.set(__self__, "read_only", read_only)

    @property
    @pulumi.getter(name="aclContent")
    def acl_content(self) -> str:
        """
        Content of ACL
        """
        return pulumi.get(self, "acl_content")

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
    @pulumi.getter(name="ipVersion")
    def ip_version(self) -> str:
        """
        IP Version
        """
        return pulumi.get(self, "ip_version")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the security group ACL
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> bool:
        """
        Read-only
        """
        return pulumi.get(self, "read_only")


class AwaitableGetSecurityGroupAclResult(GetSecurityGroupAclResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSecurityGroupAclResult(
            acl_content=self.acl_content,
            description=self.description,
            id=self.id,
            ip_version=self.ip_version,
            name=self.name,
            read_only=self.read_only)


def get_security_group_acl(id: Optional[str] = None,
                           name: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSecurityGroupAclResult:
    """
    This data source can read the TrustSec Security Group ACL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_security_group_acl(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the security group ACL
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl', __args__, opts=opts, typ=GetSecurityGroupAclResult).value

    return AwaitableGetSecurityGroupAclResult(
        acl_content=pulumi.get(__ret__, 'acl_content'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_version=pulumi.get(__ret__, 'ip_version'),
        name=pulumi.get(__ret__, 'name'),
        read_only=pulumi.get(__ret__, 'read_only'))
def get_security_group_acl_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                  name: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSecurityGroupAclResult]:
    """
    This data source can read the TrustSec Security Group ACL.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_security_group_acl(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param str id: The id of the object
    :param str name: The name of the security group ACL
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:trustsec/getSecurityGroupAcl:getSecurityGroupAcl', __args__, opts=opts, typ=GetSecurityGroupAclResult)
    return __ret__.apply(lambda __response__: GetSecurityGroupAclResult(
        acl_content=pulumi.get(__response__, 'acl_content'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        ip_version=pulumi.get(__response__, 'ip_version'),
        name=pulumi.get(__response__, 'name'),
        read_only=pulumi.get(__response__, 'read_only')))

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
from . import outputs

__all__ = [
    'GetActiveDirectoryGroupsByDomainResult',
    'AwaitableGetActiveDirectoryGroupsByDomainResult',
    'get_active_directory_groups_by_domain',
    'get_active_directory_groups_by_domain_output',
]

@pulumi.output_type
class GetActiveDirectoryGroupsByDomainResult:
    """
    A collection of values returned by getActiveDirectoryGroupsByDomain.
    """
    def __init__(__self__, domain=None, filter=None, groups=None, id=None, join_point_id=None, sid_filter=None, type_filter=None):
        if domain and not isinstance(domain, str):
            raise TypeError("Expected argument 'domain' to be a str")
        pulumi.set(__self__, "domain", domain)
        if filter and not isinstance(filter, str):
            raise TypeError("Expected argument 'filter' to be a str")
        pulumi.set(__self__, "filter", filter)
        if groups and not isinstance(groups, list):
            raise TypeError("Expected argument 'groups' to be a list")
        pulumi.set(__self__, "groups", groups)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if join_point_id and not isinstance(join_point_id, str):
            raise TypeError("Expected argument 'join_point_id' to be a str")
        pulumi.set(__self__, "join_point_id", join_point_id)
        if sid_filter and not isinstance(sid_filter, str):
            raise TypeError("Expected argument 'sid_filter' to be a str")
        pulumi.set(__self__, "sid_filter", sid_filter)
        if type_filter and not isinstance(type_filter, str):
            raise TypeError("Expected argument 'type_filter' to be a str")
        pulumi.set(__self__, "type_filter", type_filter)

    @property
    @pulumi.getter
    def domain(self) -> str:
        """
        The domain whose groups we want to fetch
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter
    def filter(self) -> Optional[str]:
        """
        Exact match filter on group's CN
        """
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def groups(self) -> Sequence['outputs.GetActiveDirectoryGroupsByDomainGroupResult']:
        """
        List of groups
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> str:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @property
    @pulumi.getter(name="sidFilter")
    def sid_filter(self) -> Optional[str]:
        """
        Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
        """
        return pulumi.get(self, "sid_filter")

    @property
    @pulumi.getter(name="typeFilter")
    def type_filter(self) -> Optional[str]:
        """
        Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
        """
        return pulumi.get(self, "type_filter")


class AwaitableGetActiveDirectoryGroupsByDomainResult(GetActiveDirectoryGroupsByDomainResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetActiveDirectoryGroupsByDomainResult(
            domain=self.domain,
            filter=self.filter,
            groups=self.groups,
            id=self.id,
            join_point_id=self.join_point_id,
            sid_filter=self.sid_filter,
            type_filter=self.type_filter)


def get_active_directory_groups_by_domain(domain: Optional[str] = None,
                                          filter: Optional[str] = None,
                                          join_point_id: Optional[str] = None,
                                          sid_filter: Optional[str] = None,
                                          type_filter: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetActiveDirectoryGroupsByDomainResult:
    """
    This data source can read the Active Directory Groups By Domain.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_active_directory_groups_by_domain(join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
        domain="cisco.com",
        filter="CN=ISE Admins",
        sid_filter="cisco.com/S-1-5-33-544",
        type_filter="UNIVERSAL")
    ```


    :param str domain: The domain whose groups we want to fetch
    :param str filter: Exact match filter on group's CN
    :param str join_point_id: Active Directory Join Point ID
    :param str sid_filter: Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
    :param str type_filter: Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['filter'] = filter
    __args__['joinPointId'] = join_point_id
    __args__['sidFilter'] = sid_filter
    __args__['typeFilter'] = type_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:identitymanagement/getActiveDirectoryGroupsByDomain:getActiveDirectoryGroupsByDomain', __args__, opts=opts, typ=GetActiveDirectoryGroupsByDomainResult).value

    return AwaitableGetActiveDirectoryGroupsByDomainResult(
        domain=pulumi.get(__ret__, 'domain'),
        filter=pulumi.get(__ret__, 'filter'),
        groups=pulumi.get(__ret__, 'groups'),
        id=pulumi.get(__ret__, 'id'),
        join_point_id=pulumi.get(__ret__, 'join_point_id'),
        sid_filter=pulumi.get(__ret__, 'sid_filter'),
        type_filter=pulumi.get(__ret__, 'type_filter'))
def get_active_directory_groups_by_domain_output(domain: Optional[pulumi.Input[str]] = None,
                                                 filter: Optional[pulumi.Input[Optional[str]]] = None,
                                                 join_point_id: Optional[pulumi.Input[str]] = None,
                                                 sid_filter: Optional[pulumi.Input[Optional[str]]] = None,
                                                 type_filter: Optional[pulumi.Input[Optional[str]]] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetActiveDirectoryGroupsByDomainResult]:
    """
    This data source can read the Active Directory Groups By Domain.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_active_directory_groups_by_domain(join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
        domain="cisco.com",
        filter="CN=ISE Admins",
        sid_filter="cisco.com/S-1-5-33-544",
        type_filter="UNIVERSAL")
    ```


    :param str domain: The domain whose groups we want to fetch
    :param str filter: Exact match filter on group's CN
    :param str join_point_id: Active Directory Join Point ID
    :param str sid_filter: Exact match filter on group's SID, optionally specifying the domain as prefix. e.g. S-1-5-33-544 and R1.dom/S-1-5-33-544 are legal.
    :param str type_filter: Can be exactly one of: BUILTIN, DOMAIN LOCAL, GLOBAL, UNIVERSAL.
    """
    __args__ = dict()
    __args__['domain'] = domain
    __args__['filter'] = filter
    __args__['joinPointId'] = join_point_id
    __args__['sidFilter'] = sid_filter
    __args__['typeFilter'] = type_filter
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:identitymanagement/getActiveDirectoryGroupsByDomain:getActiveDirectoryGroupsByDomain', __args__, opts=opts, typ=GetActiveDirectoryGroupsByDomainResult)
    return __ret__.apply(lambda __response__: GetActiveDirectoryGroupsByDomainResult(
        domain=pulumi.get(__response__, 'domain'),
        filter=pulumi.get(__response__, 'filter'),
        groups=pulumi.get(__response__, 'groups'),
        id=pulumi.get(__response__, 'id'),
        join_point_id=pulumi.get(__response__, 'join_point_id'),
        sid_filter=pulumi.get(__response__, 'sid_filter'),
        type_filter=pulumi.get(__response__, 'type_filter')))

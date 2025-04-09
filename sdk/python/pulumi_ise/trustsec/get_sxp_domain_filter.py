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
    'GetSxpDomainFilterResult',
    'AwaitableGetSxpDomainFilterResult',
    'get_sxp_domain_filter',
    'get_sxp_domain_filter_output',
]

@pulumi.output_type
class GetSxpDomainFilterResult:
    """
    A collection of values returned by getSxpDomainFilter.
    """
    def __init__(__self__, description=None, domains=None, id=None, name=None, sgt=None, subnet=None, vn=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if domains and not isinstance(domains, str):
            raise TypeError("Expected argument 'domains' to be a str")
        pulumi.set(__self__, "domains", domains)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sgt and not isinstance(sgt, str):
            raise TypeError("Expected argument 'sgt' to be a str")
        pulumi.set(__self__, "sgt", sgt)
        if subnet and not isinstance(subnet, str):
            raise TypeError("Expected argument 'subnet' to be a str")
        pulumi.set(__self__, "subnet", subnet)
        if vn and not isinstance(vn, str):
            raise TypeError("Expected argument 'vn' to be a str")
        pulumi.set(__self__, "vn", vn)

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domains(self) -> builtins.str:
        """
        List of SXP Domains, separated with comma
        """
        return pulumi.get(self, "domains")

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
        Resource name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sgt(self) -> builtins.str:
        """
        SGT name or ID. At least one of subnet or sgt or vn should be defined
        """
        return pulumi.get(self, "sgt")

    @property
    @pulumi.getter
    def subnet(self) -> builtins.str:
        """
        Subnet for filter policy (hostname is not supported). At least one of subnet or sgt or vn should be defined
        """
        return pulumi.get(self, "subnet")

    @property
    @pulumi.getter
    def vn(self) -> builtins.str:
        """
        Virtual Network. At least one of subnet or sgt or vn should be defined
        """
        return pulumi.get(self, "vn")


class AwaitableGetSxpDomainFilterResult(GetSxpDomainFilterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSxpDomainFilterResult(
            description=self.description,
            domains=self.domains,
            id=self.id,
            name=self.name,
            sgt=self.sgt,
            subnet=self.subnet,
            vn=self.vn)


def get_sxp_domain_filter(id: Optional[builtins.str] = None,
                          name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSxpDomainFilterResult:
    """
    This data source can read the SXP Domain Filter.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_sxp_domain_filter(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: Resource name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:trustsec/getSxpDomainFilter:getSxpDomainFilter', __args__, opts=opts, typ=GetSxpDomainFilterResult).value

    return AwaitableGetSxpDomainFilterResult(
        description=pulumi.get(__ret__, 'description'),
        domains=pulumi.get(__ret__, 'domains'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        sgt=pulumi.get(__ret__, 'sgt'),
        subnet=pulumi.get(__ret__, 'subnet'),
        vn=pulumi.get(__ret__, 'vn'))
def get_sxp_domain_filter_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetSxpDomainFilterResult]:
    """
    This data source can read the SXP Domain Filter.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_sxp_domain_filter(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: Resource name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:trustsec/getSxpDomainFilter:getSxpDomainFilter', __args__, opts=opts, typ=GetSxpDomainFilterResult)
    return __ret__.apply(lambda __response__: GetSxpDomainFilterResult(
        description=pulumi.get(__response__, 'description'),
        domains=pulumi.get(__response__, 'domains'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        sgt=pulumi.get(__response__, 'sgt'),
        subnet=pulumi.get(__response__, 'subnet'),
        vn=pulumi.get(__response__, 'vn')))

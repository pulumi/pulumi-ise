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
    'GetIpToSgtMappingResult',
    'AwaitableGetIpToSgtMappingResult',
    'get_ip_to_sgt_mapping',
    'get_ip_to_sgt_mapping_output',
]

@pulumi.output_type
class GetIpToSgtMappingResult:
    """
    A collection of values returned by getIpToSgtMapping.
    """
    def __init__(__self__, deploy_to=None, deploy_type=None, description=None, host_ip=None, host_name=None, id=None, mapping_group=None, name=None, sgt=None):
        if deploy_to and not isinstance(deploy_to, str):
            raise TypeError("Expected argument 'deploy_to' to be a str")
        pulumi.set(__self__, "deploy_to", deploy_to)
        if deploy_type and not isinstance(deploy_type, str):
            raise TypeError("Expected argument 'deploy_type' to be a str")
        pulumi.set(__self__, "deploy_type", deploy_type)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if host_ip and not isinstance(host_ip, str):
            raise TypeError("Expected argument 'host_ip' to be a str")
        pulumi.set(__self__, "host_ip", host_ip)
        if host_name and not isinstance(host_name, str):
            raise TypeError("Expected argument 'host_name' to be a str")
        pulumi.set(__self__, "host_name", host_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mapping_group and not isinstance(mapping_group, str):
            raise TypeError("Expected argument 'mapping_group' to be a str")
        pulumi.set(__self__, "mapping_group", mapping_group)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if sgt and not isinstance(sgt, str):
            raise TypeError("Expected argument 'sgt' to be a str")
        pulumi.set(__self__, "sgt", sgt)

    @property
    @pulumi.getter(name="deployTo")
    def deploy_to(self) -> builtins.str:
        """
        Mandatory unless `mapping_group` is set or unless `deploy_type` is `ALL`
        """
        return pulumi.get(self, "deploy_to")

    @property
    @pulumi.getter(name="deployType")
    def deploy_type(self) -> builtins.str:
        """
        Deploy Type
        """
        return pulumi.get(self, "deploy_type")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="hostIp")
    def host_ip(self) -> builtins.str:
        """
        Mandatory if `host_name` is empty
        """
        return pulumi.get(self, "host_ip")

    @property
    @pulumi.getter(name="hostName")
    def host_name(self) -> builtins.str:
        """
        Mandatory if `host_ip` is empty
        """
        return pulumi.get(self, "host_name")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="mappingGroup")
    def mapping_group(self) -> builtins.str:
        """
        IP to SGT Mapping Group ID. Mandatory unless `sgt` and `deploy_to` and `deploy_type` are set
        """
        return pulumi.get(self, "mapping_group")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the IP to SGT mapping
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sgt(self) -> builtins.str:
        """
        Trustsec Security Group ID. Mandatory unless `mapping_group` is set
        """
        return pulumi.get(self, "sgt")


class AwaitableGetIpToSgtMappingResult(GetIpToSgtMappingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpToSgtMappingResult(
            deploy_to=self.deploy_to,
            deploy_type=self.deploy_type,
            description=self.description,
            host_ip=self.host_ip,
            host_name=self.host_name,
            id=self.id,
            mapping_group=self.mapping_group,
            name=self.name,
            sgt=self.sgt)


def get_ip_to_sgt_mapping(id: Optional[builtins.str] = None,
                          name: Optional[builtins.str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpToSgtMappingResult:
    """
    This data source can read the TrustSec IP to SGT Mapping.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_ip_to_sgt_mapping(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the IP to SGT mapping
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:trustsec/getIpToSgtMapping:getIpToSgtMapping', __args__, opts=opts, typ=GetIpToSgtMappingResult).value

    return AwaitableGetIpToSgtMappingResult(
        deploy_to=pulumi.get(__ret__, 'deploy_to'),
        deploy_type=pulumi.get(__ret__, 'deploy_type'),
        description=pulumi.get(__ret__, 'description'),
        host_ip=pulumi.get(__ret__, 'host_ip'),
        host_name=pulumi.get(__ret__, 'host_name'),
        id=pulumi.get(__ret__, 'id'),
        mapping_group=pulumi.get(__ret__, 'mapping_group'),
        name=pulumi.get(__ret__, 'name'),
        sgt=pulumi.get(__ret__, 'sgt'))
def get_ip_to_sgt_mapping_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetIpToSgtMappingResult]:
    """
    This data source can read the TrustSec IP to SGT Mapping.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.trustsec.get_ip_to_sgt_mapping(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the IP to SGT mapping
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:trustsec/getIpToSgtMapping:getIpToSgtMapping', __args__, opts=opts, typ=GetIpToSgtMappingResult)
    return __ret__.apply(lambda __response__: GetIpToSgtMappingResult(
        deploy_to=pulumi.get(__response__, 'deploy_to'),
        deploy_type=pulumi.get(__response__, 'deploy_type'),
        description=pulumi.get(__response__, 'description'),
        host_ip=pulumi.get(__response__, 'host_ip'),
        host_name=pulumi.get(__response__, 'host_name'),
        id=pulumi.get(__response__, 'id'),
        mapping_group=pulumi.get(__response__, 'mapping_group'),
        name=pulumi.get(__response__, 'name'),
        sgt=pulumi.get(__response__, 'sgt')))

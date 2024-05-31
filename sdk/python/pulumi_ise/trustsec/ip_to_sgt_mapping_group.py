# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IpToSgtMappingGroupArgs', 'IpToSgtMappingGroup']

@pulumi.input_type
class IpToSgtMappingGroupArgs:
    def __init__(__self__, *,
                 deploy_type: pulumi.Input[str],
                 sgt: pulumi.Input[str],
                 deploy_to: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpToSgtMappingGroup resource.
        :param pulumi.Input[str] deploy_type: Deploy Type - Choices: `ALL`, `ND`, `NDG`
        :param pulumi.Input[str] sgt: Trustsec Security Group ID
        :param pulumi.Input[str] deploy_to: Mandatory unless `deploy_type` is `ALL`
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the IP to SGT mapping Group
        """
        pulumi.set(__self__, "deploy_type", deploy_type)
        pulumi.set(__self__, "sgt", sgt)
        if deploy_to is not None:
            pulumi.set(__self__, "deploy_to", deploy_to)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="deployType")
    def deploy_type(self) -> pulumi.Input[str]:
        """
        Deploy Type - Choices: `ALL`, `ND`, `NDG`
        """
        return pulumi.get(self, "deploy_type")

    @deploy_type.setter
    def deploy_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "deploy_type", value)

    @property
    @pulumi.getter
    def sgt(self) -> pulumi.Input[str]:
        """
        Trustsec Security Group ID
        """
        return pulumi.get(self, "sgt")

    @sgt.setter
    def sgt(self, value: pulumi.Input[str]):
        pulumi.set(self, "sgt", value)

    @property
    @pulumi.getter(name="deployTo")
    def deploy_to(self) -> Optional[pulumi.Input[str]]:
        """
        Mandatory unless `deploy_type` is `ALL`
        """
        return pulumi.get(self, "deploy_to")

    @deploy_to.setter
    def deploy_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_to", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IP to SGT mapping Group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IpToSgtMappingGroupState:
    def __init__(__self__, *,
                 deploy_to: Optional[pulumi.Input[str]] = None,
                 deploy_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sgt: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpToSgtMappingGroup resources.
        :param pulumi.Input[str] deploy_to: Mandatory unless `deploy_type` is `ALL`
        :param pulumi.Input[str] deploy_type: Deploy Type - Choices: `ALL`, `ND`, `NDG`
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the IP to SGT mapping Group
        :param pulumi.Input[str] sgt: Trustsec Security Group ID
        """
        if deploy_to is not None:
            pulumi.set(__self__, "deploy_to", deploy_to)
        if deploy_type is not None:
            pulumi.set(__self__, "deploy_type", deploy_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if sgt is not None:
            pulumi.set(__self__, "sgt", sgt)

    @property
    @pulumi.getter(name="deployTo")
    def deploy_to(self) -> Optional[pulumi.Input[str]]:
        """
        Mandatory unless `deploy_type` is `ALL`
        """
        return pulumi.get(self, "deploy_to")

    @deploy_to.setter
    def deploy_to(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_to", value)

    @property
    @pulumi.getter(name="deployType")
    def deploy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Deploy Type - Choices: `ALL`, `ND`, `NDG`
        """
        return pulumi.get(self, "deploy_type")

    @deploy_type.setter
    def deploy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the IP to SGT mapping Group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def sgt(self) -> Optional[pulumi.Input[str]]:
        """
        Trustsec Security Group ID
        """
        return pulumi.get(self, "sgt")

    @sgt.setter
    def sgt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sgt", value)


class IpToSgtMappingGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_to: Optional[pulumi.Input[str]] = None,
                 deploy_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sgt: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage a TrustSec IP to SGT Mapping Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.trustsec.IpToSgtMappingGroup("example",
            name="groupA",
            deploy_type="ALL",
            sgt="93e1bf00-8c01-11e6-996c-525400b48521")
        ```

        ## Import

        ```sh
        $ pulumi import ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deploy_to: Mandatory unless `deploy_type` is `ALL`
        :param pulumi.Input[str] deploy_type: Deploy Type - Choices: `ALL`, `ND`, `NDG`
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the IP to SGT mapping Group
        :param pulumi.Input[str] sgt: Trustsec Security Group ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpToSgtMappingGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage a TrustSec IP to SGT Mapping Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.trustsec.IpToSgtMappingGroup("example",
            name="groupA",
            deploy_type="ALL",
            sgt="93e1bf00-8c01-11e6-996c-525400b48521")
        ```

        ## Import

        ```sh
        $ pulumi import ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param IpToSgtMappingGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpToSgtMappingGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 deploy_to: Optional[pulumi.Input[str]] = None,
                 deploy_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 sgt: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpToSgtMappingGroupArgs.__new__(IpToSgtMappingGroupArgs)

            __props__.__dict__["deploy_to"] = deploy_to
            if deploy_type is None and not opts.urn:
                raise TypeError("Missing required property 'deploy_type'")
            __props__.__dict__["deploy_type"] = deploy_type
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if sgt is None and not opts.urn:
                raise TypeError("Missing required property 'sgt'")
            __props__.__dict__["sgt"] = sgt
        super(IpToSgtMappingGroup, __self__).__init__(
            'ise:trustsec/ipToSgtMappingGroup:IpToSgtMappingGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            deploy_to: Optional[pulumi.Input[str]] = None,
            deploy_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            sgt: Optional[pulumi.Input[str]] = None) -> 'IpToSgtMappingGroup':
        """
        Get an existing IpToSgtMappingGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] deploy_to: Mandatory unless `deploy_type` is `ALL`
        :param pulumi.Input[str] deploy_type: Deploy Type - Choices: `ALL`, `ND`, `NDG`
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the IP to SGT mapping Group
        :param pulumi.Input[str] sgt: Trustsec Security Group ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpToSgtMappingGroupState.__new__(_IpToSgtMappingGroupState)

        __props__.__dict__["deploy_to"] = deploy_to
        __props__.__dict__["deploy_type"] = deploy_type
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["sgt"] = sgt
        return IpToSgtMappingGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deployTo")
    def deploy_to(self) -> pulumi.Output[Optional[str]]:
        """
        Mandatory unless `deploy_type` is `ALL`
        """
        return pulumi.get(self, "deploy_to")

    @property
    @pulumi.getter(name="deployType")
    def deploy_type(self) -> pulumi.Output[str]:
        """
        Deploy Type - Choices: `ALL`, `ND`, `NDG`
        """
        return pulumi.get(self, "deploy_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the IP to SGT mapping Group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def sgt(self) -> pulumi.Output[str]:
        """
        Trustsec Security Group ID
        """
        return pulumi.get(self, "sgt")

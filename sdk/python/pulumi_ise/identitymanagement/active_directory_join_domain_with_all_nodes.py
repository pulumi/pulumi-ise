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
from ._inputs import *

__all__ = ['ActiveDirectoryJoinDomainWithAllNodesArgs', 'ActiveDirectoryJoinDomainWithAllNodes']

@pulumi.input_type
class ActiveDirectoryJoinDomainWithAllNodesArgs:
    def __init__(__self__, *,
                 additional_datas: pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]],
                 join_point_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ActiveDirectoryJoinDomainWithAllNodes resource.
        :param pulumi.Input[str] join_point_id: Active Directory Join Point ID
        """
        pulumi.set(__self__, "additional_datas", additional_datas)
        pulumi.set(__self__, "join_point_id", join_point_id)

    @property
    @pulumi.getter(name="additionalDatas")
    def additional_datas(self) -> pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]]:
        return pulumi.get(self, "additional_datas")

    @additional_datas.setter
    def additional_datas(self, value: pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]]):
        pulumi.set(self, "additional_datas", value)

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> pulumi.Input[str]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @join_point_id.setter
    def join_point_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "join_point_id", value)


@pulumi.input_type
class _ActiveDirectoryJoinDomainWithAllNodesState:
    def __init__(__self__, *,
                 additional_datas: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]]] = None,
                 join_point_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ActiveDirectoryJoinDomainWithAllNodes resources.
        :param pulumi.Input[str] join_point_id: Active Directory Join Point ID
        """
        if additional_datas is not None:
            pulumi.set(__self__, "additional_datas", additional_datas)
        if join_point_id is not None:
            pulumi.set(__self__, "join_point_id", join_point_id)

    @property
    @pulumi.getter(name="additionalDatas")
    def additional_datas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]]]:
        return pulumi.get(self, "additional_datas")

    @additional_datas.setter
    def additional_datas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs']]]]):
        pulumi.set(self, "additional_datas", value)

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> Optional[pulumi.Input[str]]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @join_point_id.setter
    def join_point_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "join_point_id", value)


class ActiveDirectoryJoinDomainWithAllNodes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_datas: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs', 'ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgsDict']]]]] = None,
                 join_point_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage an Active Directory Join Domain with All Nodes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.ActiveDirectoryJoinDomainWithAllNodes("example",
            join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
            additional_datas=[{
                "name": "username",
                "value": "administrator",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] join_point_id: Active Directory Join Point ID
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActiveDirectoryJoinDomainWithAllNodesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage an Active Directory Join Domain with All Nodes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.ActiveDirectoryJoinDomainWithAllNodes("example",
            join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
            additional_datas=[{
                "name": "username",
                "value": "administrator",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param ActiveDirectoryJoinDomainWithAllNodesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActiveDirectoryJoinDomainWithAllNodesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_datas: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs', 'ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgsDict']]]]] = None,
                 join_point_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActiveDirectoryJoinDomainWithAllNodesArgs.__new__(ActiveDirectoryJoinDomainWithAllNodesArgs)

            if additional_datas is None and not opts.urn:
                raise TypeError("Missing required property 'additional_datas'")
            __props__.__dict__["additional_datas"] = additional_datas
            if join_point_id is None and not opts.urn:
                raise TypeError("Missing required property 'join_point_id'")
            __props__.__dict__["join_point_id"] = join_point_id
        super(ActiveDirectoryJoinDomainWithAllNodes, __self__).__init__(
            'ise:identitymanagement/activeDirectoryJoinDomainWithAllNodes:ActiveDirectoryJoinDomainWithAllNodes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_datas: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgs', 'ActiveDirectoryJoinDomainWithAllNodesAdditionalDataArgsDict']]]]] = None,
            join_point_id: Optional[pulumi.Input[str]] = None) -> 'ActiveDirectoryJoinDomainWithAllNodes':
        """
        Get an existing ActiveDirectoryJoinDomainWithAllNodes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] join_point_id: Active Directory Join Point ID
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActiveDirectoryJoinDomainWithAllNodesState.__new__(_ActiveDirectoryJoinDomainWithAllNodesState)

        __props__.__dict__["additional_datas"] = additional_datas
        __props__.__dict__["join_point_id"] = join_point_id
        return ActiveDirectoryJoinDomainWithAllNodes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalDatas")
    def additional_datas(self) -> pulumi.Output[Sequence['outputs.ActiveDirectoryJoinDomainWithAllNodesAdditionalData']]:
        return pulumi.get(self, "additional_datas")

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> pulumi.Output[str]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")


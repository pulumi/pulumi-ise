# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GroupArgs', 'Group']

@pulumi.input_type
class GroupArgs:
    def __init__(__self__, *,
                 root_group: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Group resource.
        :param pulumi.Input[str] root_group: The name of the root device group.
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        """
        pulumi.set(__self__, "root_group", root_group)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="rootGroup")
    def root_group(self) -> pulumi.Input[str]:
        """
        The name of the root device group.
        """
        return pulumi.get(self, "root_group")

    @root_group.setter
    def root_group(self, value: pulumi.Input[str]):
        pulumi.set(self, "root_group", value)

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
        The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _GroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_group: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Group resources.
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        :param pulumi.Input[str] root_group: The name of the root device group.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if root_group is not None:
            pulumi.set(__self__, "root_group", root_group)

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
        The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rootGroup")
    def root_group(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the root device group.
        """
        return pulumi.get(self, "root_group")

    @root_group.setter
    def root_group(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_group", value)


class Group(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_group: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage a Network Device Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-ise as ise

        example = ise.network_device.Group("example",
            name="Device Type#All Device Types#Group1",
            description="My network device group",
            root_group="Device Type")
        ```

        ## Import

        ```sh
        $ pulumi import ise:NetworkDevice/group:Group example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        :param pulumi.Input[str] root_group: The name of the root device group.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage a Network Device Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi-ise as ise

        example = ise.network_device.Group("example",
            name="Device Type#All Device Types#Group1",
            description="My network device group",
            root_group="Device Type")
        ```

        ## Import

        ```sh
        $ pulumi import ise:NetworkDevice/group:Group example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param GroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 root_group: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupArgs.__new__(GroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if root_group is None and not opts.urn:
                raise TypeError("Missing required property 'root_group'")
            __props__.__dict__["root_group"] = root_group
        super(Group, __self__).__init__(
            'ise:NetworkDevice/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            root_group: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        :param pulumi.Input[str] root_group: The name of the root device group.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupState.__new__(_GroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["root_group"] = root_group
        return Group(resource_name, opts=opts, __props__=__props__)

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
        The name of the network device group including its hierarchy, e.g. `Device Type#All Device Types#ACCESS`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rootGroup")
    def root_group(self) -> pulumi.Output[str]:
        """
        The name of the root device group.
        """
        return pulumi.get(self, "root_group")


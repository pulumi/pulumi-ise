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
from ._inputs import *

__all__ = ['ActiveDirectoryAddGroupsArgs', 'ActiveDirectoryAddGroups']

@pulumi.input_type
class ActiveDirectoryAddGroupsArgs:
    def __init__(__self__, *,
                 domain: pulumi.Input[builtins.str],
                 join_point_id: pulumi.Input[builtins.str],
                 ad_scopes_names: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 enable_domain_allowed_list: Optional[pulumi.Input[builtins.bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ActiveDirectoryAddGroups resource.
        :param pulumi.Input[builtins.str] domain: AD domain associated with the join point
        :param pulumi.Input[builtins.str] join_point_id: Active Directory Join Point ID
        :param pulumi.Input[builtins.str] ad_scopes_names: String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
               value: `Default_Scope`
        :param pulumi.Input[builtins.str] description: Join point Description
        :param pulumi.Input[builtins.bool] enable_domain_allowed_list: - Default value: `true`
        :param pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]] groups: List of AD Groups
        :param pulumi.Input[builtins.str] name: The name of the active directory join point
        """
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "join_point_id", join_point_id)
        if ad_scopes_names is not None:
            pulumi.set(__self__, "ad_scopes_names", ad_scopes_names)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_domain_allowed_list is not None:
            pulumi.set(__self__, "enable_domain_allowed_list", enable_domain_allowed_list)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[builtins.str]:
        """
        AD domain associated with the join point
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> pulumi.Input[builtins.str]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @join_point_id.setter
    def join_point_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "join_point_id", value)

    @property
    @pulumi.getter(name="adScopesNames")
    def ad_scopes_names(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        value: `Default_Scope`
        """
        return pulumi.get(self, "ad_scopes_names")

    @ad_scopes_names.setter
    def ad_scopes_names(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ad_scopes_names", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Join point Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableDomainAllowedList")
    def enable_domain_allowed_list(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        - Default value: `true`
        """
        return pulumi.get(self, "enable_domain_allowed_list")

    @enable_domain_allowed_list.setter
    def enable_domain_allowed_list(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_domain_allowed_list", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]]:
        """
        List of AD Groups
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the active directory join point
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ActiveDirectoryAddGroupsState:
    def __init__(__self__, *,
                 ad_scopes_names: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 enable_domain_allowed_list: Optional[pulumi.Input[builtins.bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]] = None,
                 join_point_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ActiveDirectoryAddGroups resources.
        :param pulumi.Input[builtins.str] ad_scopes_names: String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
               value: `Default_Scope`
        :param pulumi.Input[builtins.str] description: Join point Description
        :param pulumi.Input[builtins.str] domain: AD domain associated with the join point
        :param pulumi.Input[builtins.bool] enable_domain_allowed_list: - Default value: `true`
        :param pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]] groups: List of AD Groups
        :param pulumi.Input[builtins.str] join_point_id: Active Directory Join Point ID
        :param pulumi.Input[builtins.str] name: The name of the active directory join point
        """
        if ad_scopes_names is not None:
            pulumi.set(__self__, "ad_scopes_names", ad_scopes_names)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if domain is not None:
            pulumi.set(__self__, "domain", domain)
        if enable_domain_allowed_list is not None:
            pulumi.set(__self__, "enable_domain_allowed_list", enable_domain_allowed_list)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if join_point_id is not None:
            pulumi.set(__self__, "join_point_id", join_point_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="adScopesNames")
    def ad_scopes_names(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        value: `Default_Scope`
        """
        return pulumi.get(self, "ad_scopes_names")

    @ad_scopes_names.setter
    def ad_scopes_names(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ad_scopes_names", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Join point Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def domain(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        AD domain associated with the join point
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="enableDomainAllowedList")
    def enable_domain_allowed_list(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        - Default value: `true`
        """
        return pulumi.get(self, "enable_domain_allowed_list")

    @enable_domain_allowed_list.setter
    def enable_domain_allowed_list(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_domain_allowed_list", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]]:
        """
        List of AD Groups
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ActiveDirectoryAddGroupsGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @join_point_id.setter
    def join_point_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "join_point_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the active directory join point
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


class ActiveDirectoryAddGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_scopes_names: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 enable_domain_allowed_list: Optional[pulumi.Input[builtins.bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryAddGroupsGroupArgs', 'ActiveDirectoryAddGroupsGroupArgsDict']]]]] = None,
                 join_point_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        This resource can manage an Active Directory Add Groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.ActiveDirectoryAddGroups("example",
            join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
            name="cisco.local",
            description="My AD join point",
            domain="cisco.local",
            ad_scopes_names="Default_Scope",
            enable_domain_allowed_list=True,
            groups=[{
                "name": "cisco.local/operators",
                "sid": "S-1-5-32-548",
                "type": "GLOBAL",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] ad_scopes_names: String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
               value: `Default_Scope`
        :param pulumi.Input[builtins.str] description: Join point Description
        :param pulumi.Input[builtins.str] domain: AD domain associated with the join point
        :param pulumi.Input[builtins.bool] enable_domain_allowed_list: - Default value: `true`
        :param pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryAddGroupsGroupArgs', 'ActiveDirectoryAddGroupsGroupArgsDict']]]] groups: List of AD Groups
        :param pulumi.Input[builtins.str] join_point_id: Active Directory Join Point ID
        :param pulumi.Input[builtins.str] name: The name of the active directory join point
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ActiveDirectoryAddGroupsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage an Active Directory Add Groups.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.ActiveDirectoryAddGroups("example",
            join_point_id="73808580-b6e6-11ee-8960-de6d7692bc40",
            name="cisco.local",
            description="My AD join point",
            domain="cisco.local",
            ad_scopes_names="Default_Scope",
            enable_domain_allowed_list=True,
            groups=[{
                "name": "cisco.local/operators",
                "sid": "S-1-5-32-548",
                "type": "GLOBAL",
            }])
        ```

        :param str resource_name: The name of the resource.
        :param ActiveDirectoryAddGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ActiveDirectoryAddGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ad_scopes_names: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 domain: Optional[pulumi.Input[builtins.str]] = None,
                 enable_domain_allowed_list: Optional[pulumi.Input[builtins.bool]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryAddGroupsGroupArgs', 'ActiveDirectoryAddGroupsGroupArgsDict']]]]] = None,
                 join_point_id: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ActiveDirectoryAddGroupsArgs.__new__(ActiveDirectoryAddGroupsArgs)

            __props__.__dict__["ad_scopes_names"] = ad_scopes_names
            __props__.__dict__["description"] = description
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["enable_domain_allowed_list"] = enable_domain_allowed_list
            __props__.__dict__["groups"] = groups
            if join_point_id is None and not opts.urn:
                raise TypeError("Missing required property 'join_point_id'")
            __props__.__dict__["join_point_id"] = join_point_id
            __props__.__dict__["name"] = name
        super(ActiveDirectoryAddGroups, __self__).__init__(
            'ise:identitymanagement/activeDirectoryAddGroups:ActiveDirectoryAddGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ad_scopes_names: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            domain: Optional[pulumi.Input[builtins.str]] = None,
            enable_domain_allowed_list: Optional[pulumi.Input[builtins.bool]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryAddGroupsGroupArgs', 'ActiveDirectoryAddGroupsGroupArgsDict']]]]] = None,
            join_point_id: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None) -> 'ActiveDirectoryAddGroups':
        """
        Get an existing ActiveDirectoryAddGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] ad_scopes_names: String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
               value: `Default_Scope`
        :param pulumi.Input[builtins.str] description: Join point Description
        :param pulumi.Input[builtins.str] domain: AD domain associated with the join point
        :param pulumi.Input[builtins.bool] enable_domain_allowed_list: - Default value: `true`
        :param pulumi.Input[Sequence[pulumi.Input[Union['ActiveDirectoryAddGroupsGroupArgs', 'ActiveDirectoryAddGroupsGroupArgsDict']]]] groups: List of AD Groups
        :param pulumi.Input[builtins.str] join_point_id: Active Directory Join Point ID
        :param pulumi.Input[builtins.str] name: The name of the active directory join point
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ActiveDirectoryAddGroupsState.__new__(_ActiveDirectoryAddGroupsState)

        __props__.__dict__["ad_scopes_names"] = ad_scopes_names
        __props__.__dict__["description"] = description
        __props__.__dict__["domain"] = domain
        __props__.__dict__["enable_domain_allowed_list"] = enable_domain_allowed_list
        __props__.__dict__["groups"] = groups
        __props__.__dict__["join_point_id"] = join_point_id
        __props__.__dict__["name"] = name
        return ActiveDirectoryAddGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adScopesNames")
    def ad_scopes_names(self) -> pulumi.Output[builtins.str]:
        """
        String that contains the names of the scopes that the active directory belongs to. Names are separated by comm - Default
        value: `Default_Scope`
        """
        return pulumi.get(self, "ad_scopes_names")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Join point Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[builtins.str]:
        """
        AD domain associated with the join point
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="enableDomainAllowedList")
    def enable_domain_allowed_list(self) -> pulumi.Output[builtins.bool]:
        """
        - Default value: `true`
        """
        return pulumi.get(self, "enable_domain_allowed_list")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Optional[Sequence['outputs.ActiveDirectoryAddGroupsGroup']]]:
        """
        List of AD Groups
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="joinPointId")
    def join_point_id(self) -> pulumi.Output[builtins.str]:
        """
        Active Directory Join Point ID
        """
        return pulumi.get(self, "join_point_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the active directory join point
        """
        return pulumi.get(self, "name")


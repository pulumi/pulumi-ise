# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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

__all__ = ['SecurityGroupArgs', 'SecurityGroup']

@pulumi.input_type
class SecurityGroupArgs:
    def __init__(__self__, *,
                 value: pulumi.Input[builtins.int],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_read_only: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 propogate_to_apic: Optional[pulumi.Input[builtins.bool]] = None):
        """
        The set of arguments for constructing a SecurityGroup resource.
        :param pulumi.Input[builtins.int] value: `-1` to auto-generate - Range: `-1`-`65519`
        :param pulumi.Input[builtins.str] description: Description
        :param pulumi.Input[builtins.bool] is_read_only: Read-only
        :param pulumi.Input[builtins.str] name: The name of the security group
        :param pulumi.Input[builtins.bool] propogate_to_apic: Propagate to APIC (ACI)
        """
        pulumi.set(__self__, "value", value)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_read_only is not None:
            pulumi.set(__self__, "is_read_only", is_read_only)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if propogate_to_apic is not None:
            pulumi.set(__self__, "propogate_to_apic", propogate_to_apic)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[builtins.int]:
        """
        `-1` to auto-generate - Range: `-1`-`65519`
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[builtins.int]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isReadOnly")
    def is_read_only(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Read-only
        """
        return pulumi.get(self, "is_read_only")

    @is_read_only.setter
    def is_read_only(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_read_only", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the security group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="propogateToApic")
    def propogate_to_apic(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Propagate to APIC (ACI)
        """
        return pulumi.get(self, "propogate_to_apic")

    @propogate_to_apic.setter
    def propogate_to_apic(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "propogate_to_apic", value)


@pulumi.input_type
class _SecurityGroupState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_read_only: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 propogate_to_apic: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.int]] = None):
        """
        Input properties used for looking up and filtering SecurityGroup resources.
        :param pulumi.Input[builtins.str] description: Description
        :param pulumi.Input[builtins.bool] is_read_only: Read-only
        :param pulumi.Input[builtins.str] name: The name of the security group
        :param pulumi.Input[builtins.bool] propogate_to_apic: Propagate to APIC (ACI)
        :param pulumi.Input[builtins.int] value: `-1` to auto-generate - Range: `-1`-`65519`
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_read_only is not None:
            pulumi.set(__self__, "is_read_only", is_read_only)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if propogate_to_apic is not None:
            pulumi.set(__self__, "propogate_to_apic", propogate_to_apic)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isReadOnly")
    def is_read_only(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Read-only
        """
        return pulumi.get(self, "is_read_only")

    @is_read_only.setter
    def is_read_only(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "is_read_only", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the security group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="propogateToApic")
    def propogate_to_apic(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Propagate to APIC (ACI)
        """
        return pulumi.get(self, "propogate_to_apic")

    @propogate_to_apic.setter
    def propogate_to_apic(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "propogate_to_apic", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        `-1` to auto-generate - Range: `-1`-`65519`
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "value", value)


@pulumi.type_token("ise:trustsec/securityGroup:SecurityGroup")
class SecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_read_only: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 propogate_to_apic: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        """
        This resource can manage a TrustSec Security Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.trustsec.SecurityGroup("example",
            name="SGT1234",
            description="My SGT",
            value=1234,
            propogate_to_apic=True,
            is_read_only=False)
        ```

        ## Import

        ```sh
        $ pulumi import ise:trustsec/securityGroup:SecurityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description
        :param pulumi.Input[builtins.bool] is_read_only: Read-only
        :param pulumi.Input[builtins.str] name: The name of the security group
        :param pulumi.Input[builtins.bool] propogate_to_apic: Propagate to APIC (ACI)
        :param pulumi.Input[builtins.int] value: `-1` to auto-generate - Range: `-1`-`65519`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage a TrustSec Security Group.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.trustsec.SecurityGroup("example",
            name="SGT1234",
            description="My SGT",
            value=1234,
            propogate_to_apic=True,
            is_read_only=False)
        ```

        ## Import

        ```sh
        $ pulumi import ise:trustsec/securityGroup:SecurityGroup example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param SecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 is_read_only: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 propogate_to_apic: Optional[pulumi.Input[builtins.bool]] = None,
                 value: Optional[pulumi.Input[builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityGroupArgs.__new__(SecurityGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["is_read_only"] = is_read_only
            __props__.__dict__["name"] = name
            __props__.__dict__["propogate_to_apic"] = propogate_to_apic
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(SecurityGroup, __self__).__init__(
            'ise:trustsec/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            is_read_only: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            propogate_to_apic: Optional[pulumi.Input[builtins.bool]] = None,
            value: Optional[pulumi.Input[builtins.int]] = None) -> 'SecurityGroup':
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] description: Description
        :param pulumi.Input[builtins.bool] is_read_only: Read-only
        :param pulumi.Input[builtins.str] name: The name of the security group
        :param pulumi.Input[builtins.bool] propogate_to_apic: Propagate to APIC (ACI)
        :param pulumi.Input[builtins.int] value: `-1` to auto-generate - Range: `-1`-`65519`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityGroupState.__new__(_SecurityGroupState)

        __props__.__dict__["description"] = description
        __props__.__dict__["is_read_only"] = is_read_only
        __props__.__dict__["name"] = name
        __props__.__dict__["propogate_to_apic"] = propogate_to_apic
        __props__.__dict__["value"] = value
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="isReadOnly")
    def is_read_only(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Read-only
        """
        return pulumi.get(self, "is_read_only")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        The name of the security group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="propogateToApic")
    def propogate_to_apic(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Propagate to APIC (ACI)
        """
        return pulumi.get(self, "propogate_to_apic")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[builtins.int]:
        """
        `-1` to auto-generate - Range: `-1`-`65519`
        """
        return pulumi.get(self, "value")


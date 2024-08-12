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

__all__ = ['IdentitySourceSequenceArgs', 'IdentitySourceSequence']

@pulumi.input_type
class IdentitySourceSequenceArgs:
    def __init__(__self__, *,
                 break_on_store_fail: pulumi.Input[bool],
                 certificate_authentication_profile: pulumi.Input[str],
                 identity_sources: pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IdentitySourceSequence resource.
        :param pulumi.Input[bool] break_on_store_fail: Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        :param pulumi.Input[str] certificate_authentication_profile: Certificate Authentication Profile, empty if doesn't exist
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the identity source sequence
        """
        pulumi.set(__self__, "break_on_store_fail", break_on_store_fail)
        pulumi.set(__self__, "certificate_authentication_profile", certificate_authentication_profile)
        pulumi.set(__self__, "identity_sources", identity_sources)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="breakOnStoreFail")
    def break_on_store_fail(self) -> pulumi.Input[bool]:
        """
        Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        """
        return pulumi.get(self, "break_on_store_fail")

    @break_on_store_fail.setter
    def break_on_store_fail(self, value: pulumi.Input[bool]):
        pulumi.set(self, "break_on_store_fail", value)

    @property
    @pulumi.getter(name="certificateAuthenticationProfile")
    def certificate_authentication_profile(self) -> pulumi.Input[str]:
        """
        Certificate Authentication Profile, empty if doesn't exist
        """
        return pulumi.get(self, "certificate_authentication_profile")

    @certificate_authentication_profile.setter
    def certificate_authentication_profile(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_authentication_profile", value)

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]]:
        return pulumi.get(self, "identity_sources")

    @identity_sources.setter
    def identity_sources(self, value: pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]]):
        pulumi.set(self, "identity_sources", value)

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
        The name of the identity source sequence
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IdentitySourceSequenceState:
    def __init__(__self__, *,
                 break_on_store_fail: Optional[pulumi.Input[bool]] = None,
                 certificate_authentication_profile: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IdentitySourceSequence resources.
        :param pulumi.Input[bool] break_on_store_fail: Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        :param pulumi.Input[str] certificate_authentication_profile: Certificate Authentication Profile, empty if doesn't exist
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the identity source sequence
        """
        if break_on_store_fail is not None:
            pulumi.set(__self__, "break_on_store_fail", break_on_store_fail)
        if certificate_authentication_profile is not None:
            pulumi.set(__self__, "certificate_authentication_profile", certificate_authentication_profile)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if identity_sources is not None:
            pulumi.set(__self__, "identity_sources", identity_sources)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="breakOnStoreFail")
    def break_on_store_fail(self) -> Optional[pulumi.Input[bool]]:
        """
        Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        """
        return pulumi.get(self, "break_on_store_fail")

    @break_on_store_fail.setter
    def break_on_store_fail(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "break_on_store_fail", value)

    @property
    @pulumi.getter(name="certificateAuthenticationProfile")
    def certificate_authentication_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Certificate Authentication Profile, empty if doesn't exist
        """
        return pulumi.get(self, "certificate_authentication_profile")

    @certificate_authentication_profile.setter
    def certificate_authentication_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate_authentication_profile", value)

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
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]]]:
        return pulumi.get(self, "identity_sources")

    @identity_sources.setter
    def identity_sources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['IdentitySourceSequenceIdentitySourceArgs']]]]):
        pulumi.set(self, "identity_sources", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the identity source sequence
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class IdentitySourceSequence(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 break_on_store_fail: Optional[pulumi.Input[bool]] = None,
                 certificate_authentication_profile: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentitySourceSequenceIdentitySourceArgs', 'IdentitySourceSequenceIdentitySourceArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage an Identity Source Sequence.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.IdentitySourceSequence("example",
            name="Sequence1",
            description="My identity source sequence",
            break_on_store_fail=True,
            certificate_authentication_profile="Preloaded_Certificate_Profile",
            identity_sources=[{
                "name": "Internal Users",
                "order": 1,
            }])
        ```

        ## Import

        ```sh
        $ pulumi import ise:identitymanagement/identitySourceSequence:IdentitySourceSequence example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] break_on_store_fail: Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        :param pulumi.Input[str] certificate_authentication_profile: Certificate Authentication Profile, empty if doesn't exist
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the identity source sequence
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentitySourceSequenceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage an Identity Source Sequence.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_ise as ise

        example = ise.identitymanagement.IdentitySourceSequence("example",
            name="Sequence1",
            description="My identity source sequence",
            break_on_store_fail=True,
            certificate_authentication_profile="Preloaded_Certificate_Profile",
            identity_sources=[{
                "name": "Internal Users",
                "order": 1,
            }])
        ```

        ## Import

        ```sh
        $ pulumi import ise:identitymanagement/identitySourceSequence:IdentitySourceSequence example "76d24097-41c4-4558-a4d0-a8c07ac08470"
        ```

        :param str resource_name: The name of the resource.
        :param IdentitySourceSequenceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentitySourceSequenceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 break_on_store_fail: Optional[pulumi.Input[bool]] = None,
                 certificate_authentication_profile: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentitySourceSequenceIdentitySourceArgs', 'IdentitySourceSequenceIdentitySourceArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentitySourceSequenceArgs.__new__(IdentitySourceSequenceArgs)

            if break_on_store_fail is None and not opts.urn:
                raise TypeError("Missing required property 'break_on_store_fail'")
            __props__.__dict__["break_on_store_fail"] = break_on_store_fail
            if certificate_authentication_profile is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_authentication_profile'")
            __props__.__dict__["certificate_authentication_profile"] = certificate_authentication_profile
            __props__.__dict__["description"] = description
            if identity_sources is None and not opts.urn:
                raise TypeError("Missing required property 'identity_sources'")
            __props__.__dict__["identity_sources"] = identity_sources
            __props__.__dict__["name"] = name
        super(IdentitySourceSequence, __self__).__init__(
            'ise:identitymanagement/identitySourceSequence:IdentitySourceSequence',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            break_on_store_fail: Optional[pulumi.Input[bool]] = None,
            certificate_authentication_profile: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            identity_sources: Optional[pulumi.Input[Sequence[pulumi.Input[Union['IdentitySourceSequenceIdentitySourceArgs', 'IdentitySourceSequenceIdentitySourceArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'IdentitySourceSequence':
        """
        Get an existing IdentitySourceSequence resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] break_on_store_fail: Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        :param pulumi.Input[str] certificate_authentication_profile: Certificate Authentication Profile, empty if doesn't exist
        :param pulumi.Input[str] description: Description
        :param pulumi.Input[str] name: The name of the identity source sequence
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentitySourceSequenceState.__new__(_IdentitySourceSequenceState)

        __props__.__dict__["break_on_store_fail"] = break_on_store_fail
        __props__.__dict__["certificate_authentication_profile"] = certificate_authentication_profile
        __props__.__dict__["description"] = description
        __props__.__dict__["identity_sources"] = identity_sources
        __props__.__dict__["name"] = name
        return IdentitySourceSequence(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="breakOnStoreFail")
    def break_on_store_fail(self) -> pulumi.Output[bool]:
        """
        Do not access other stores in the sequence if a selected identity store cannot be accessed for authentication
        """
        return pulumi.get(self, "break_on_store_fail")

    @property
    @pulumi.getter(name="certificateAuthenticationProfile")
    def certificate_authentication_profile(self) -> pulumi.Output[str]:
        """
        Certificate Authentication Profile, empty if doesn't exist
        """
        return pulumi.get(self, "certificate_authentication_profile")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="identitySources")
    def identity_sources(self) -> pulumi.Output[Sequence['outputs.IdentitySourceSequenceIdentitySource']]:
        return pulumi.get(self, "identity_sources")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the identity source sequence
        """
        return pulumi.get(self, "name")


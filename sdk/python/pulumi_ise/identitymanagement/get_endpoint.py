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
    'GetEndpointResult',
    'AwaitableGetEndpointResult',
    'get_endpoint',
    'get_endpoint_output',
]

@pulumi.output_type
class GetEndpointResult:
    """
    A collection of values returned by getEndpoint.
    """
    def __init__(__self__, custom_attributes=None, description=None, group_id=None, id=None, identity_store=None, identity_store_id=None, mac=None, mdm_compliance_status=None, mdm_encrypted=None, mdm_enrolled=None, mdm_imei=None, mdm_jail_broken=None, mdm_manufacturer=None, mdm_model=None, mdm_os=None, mdm_phone_number=None, mdm_pinlock=None, mdm_reachable=None, mdm_serial=None, mdm_server_name=None, name=None, portal_user=None, profile_id=None, static_group_assignment=None, static_group_assignment_defined=None, static_profile_assignment=None, static_profile_assignment_defined=None):
        if custom_attributes and not isinstance(custom_attributes, dict):
            raise TypeError("Expected argument 'custom_attributes' to be a dict")
        pulumi.set(__self__, "custom_attributes", custom_attributes)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if group_id and not isinstance(group_id, str):
            raise TypeError("Expected argument 'group_id' to be a str")
        pulumi.set(__self__, "group_id", group_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if identity_store and not isinstance(identity_store, str):
            raise TypeError("Expected argument 'identity_store' to be a str")
        pulumi.set(__self__, "identity_store", identity_store)
        if identity_store_id and not isinstance(identity_store_id, str):
            raise TypeError("Expected argument 'identity_store_id' to be a str")
        pulumi.set(__self__, "identity_store_id", identity_store_id)
        if mac and not isinstance(mac, str):
            raise TypeError("Expected argument 'mac' to be a str")
        pulumi.set(__self__, "mac", mac)
        if mdm_compliance_status and not isinstance(mdm_compliance_status, bool):
            raise TypeError("Expected argument 'mdm_compliance_status' to be a bool")
        pulumi.set(__self__, "mdm_compliance_status", mdm_compliance_status)
        if mdm_encrypted and not isinstance(mdm_encrypted, bool):
            raise TypeError("Expected argument 'mdm_encrypted' to be a bool")
        pulumi.set(__self__, "mdm_encrypted", mdm_encrypted)
        if mdm_enrolled and not isinstance(mdm_enrolled, bool):
            raise TypeError("Expected argument 'mdm_enrolled' to be a bool")
        pulumi.set(__self__, "mdm_enrolled", mdm_enrolled)
        if mdm_imei and not isinstance(mdm_imei, str):
            raise TypeError("Expected argument 'mdm_imei' to be a str")
        pulumi.set(__self__, "mdm_imei", mdm_imei)
        if mdm_jail_broken and not isinstance(mdm_jail_broken, bool):
            raise TypeError("Expected argument 'mdm_jail_broken' to be a bool")
        pulumi.set(__self__, "mdm_jail_broken", mdm_jail_broken)
        if mdm_manufacturer and not isinstance(mdm_manufacturer, str):
            raise TypeError("Expected argument 'mdm_manufacturer' to be a str")
        pulumi.set(__self__, "mdm_manufacturer", mdm_manufacturer)
        if mdm_model and not isinstance(mdm_model, str):
            raise TypeError("Expected argument 'mdm_model' to be a str")
        pulumi.set(__self__, "mdm_model", mdm_model)
        if mdm_os and not isinstance(mdm_os, str):
            raise TypeError("Expected argument 'mdm_os' to be a str")
        pulumi.set(__self__, "mdm_os", mdm_os)
        if mdm_phone_number and not isinstance(mdm_phone_number, str):
            raise TypeError("Expected argument 'mdm_phone_number' to be a str")
        pulumi.set(__self__, "mdm_phone_number", mdm_phone_number)
        if mdm_pinlock and not isinstance(mdm_pinlock, bool):
            raise TypeError("Expected argument 'mdm_pinlock' to be a bool")
        pulumi.set(__self__, "mdm_pinlock", mdm_pinlock)
        if mdm_reachable and not isinstance(mdm_reachable, bool):
            raise TypeError("Expected argument 'mdm_reachable' to be a bool")
        pulumi.set(__self__, "mdm_reachable", mdm_reachable)
        if mdm_serial and not isinstance(mdm_serial, str):
            raise TypeError("Expected argument 'mdm_serial' to be a str")
        pulumi.set(__self__, "mdm_serial", mdm_serial)
        if mdm_server_name and not isinstance(mdm_server_name, str):
            raise TypeError("Expected argument 'mdm_server_name' to be a str")
        pulumi.set(__self__, "mdm_server_name", mdm_server_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if portal_user and not isinstance(portal_user, str):
            raise TypeError("Expected argument 'portal_user' to be a str")
        pulumi.set(__self__, "portal_user", portal_user)
        if profile_id and not isinstance(profile_id, str):
            raise TypeError("Expected argument 'profile_id' to be a str")
        pulumi.set(__self__, "profile_id", profile_id)
        if static_group_assignment and not isinstance(static_group_assignment, bool):
            raise TypeError("Expected argument 'static_group_assignment' to be a bool")
        pulumi.set(__self__, "static_group_assignment", static_group_assignment)
        if static_group_assignment_defined and not isinstance(static_group_assignment_defined, bool):
            raise TypeError("Expected argument 'static_group_assignment_defined' to be a bool")
        pulumi.set(__self__, "static_group_assignment_defined", static_group_assignment_defined)
        if static_profile_assignment and not isinstance(static_profile_assignment, bool):
            raise TypeError("Expected argument 'static_profile_assignment' to be a bool")
        pulumi.set(__self__, "static_profile_assignment", static_profile_assignment)
        if static_profile_assignment_defined and not isinstance(static_profile_assignment_defined, bool):
            raise TypeError("Expected argument 'static_profile_assignment_defined' to be a bool")
        pulumi.set(__self__, "static_profile_assignment_defined", static_profile_assignment_defined)

    @property
    @pulumi.getter(name="customAttributes")
    def custom_attributes(self) -> Mapping[str, builtins.str]:
        """
        Custom Attributes
        """
        return pulumi.get(self, "custom_attributes")

    @property
    @pulumi.getter
    def description(self) -> builtins.str:
        """
        Description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> builtins.str:
        """
        Identity Group ID
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="identityStore")
    def identity_store(self) -> builtins.str:
        """
        Identity Store
        """
        return pulumi.get(self, "identity_store")

    @property
    @pulumi.getter(name="identityStoreId")
    def identity_store_id(self) -> builtins.str:
        """
        Identity Store Id
        """
        return pulumi.get(self, "identity_store_id")

    @property
    @pulumi.getter
    def mac(self) -> builtins.str:
        """
        MAC address of the endpoint
        """
        return pulumi.get(self, "mac")

    @property
    @pulumi.getter(name="mdmComplianceStatus")
    def mdm_compliance_status(self) -> builtins.bool:
        """
        Mdm Compliance Status
        """
        return pulumi.get(self, "mdm_compliance_status")

    @property
    @pulumi.getter(name="mdmEncrypted")
    def mdm_encrypted(self) -> builtins.bool:
        """
        Mdm Encrypted
        """
        return pulumi.get(self, "mdm_encrypted")

    @property
    @pulumi.getter(name="mdmEnrolled")
    def mdm_enrolled(self) -> builtins.bool:
        """
        Mdm Enrolled
        """
        return pulumi.get(self, "mdm_enrolled")

    @property
    @pulumi.getter(name="mdmImei")
    def mdm_imei(self) -> builtins.str:
        """
        Mdm IMEI
        """
        return pulumi.get(self, "mdm_imei")

    @property
    @pulumi.getter(name="mdmJailBroken")
    def mdm_jail_broken(self) -> builtins.bool:
        """
        Mdm JailBroken
        """
        return pulumi.get(self, "mdm_jail_broken")

    @property
    @pulumi.getter(name="mdmManufacturer")
    def mdm_manufacturer(self) -> builtins.str:
        """
        Mdm Manufacturer
        """
        return pulumi.get(self, "mdm_manufacturer")

    @property
    @pulumi.getter(name="mdmModel")
    def mdm_model(self) -> builtins.str:
        """
        Mdm Model
        """
        return pulumi.get(self, "mdm_model")

    @property
    @pulumi.getter(name="mdmOs")
    def mdm_os(self) -> builtins.str:
        """
        Mdm OS
        """
        return pulumi.get(self, "mdm_os")

    @property
    @pulumi.getter(name="mdmPhoneNumber")
    def mdm_phone_number(self) -> builtins.str:
        """
        Mdm PhoneNumber
        """
        return pulumi.get(self, "mdm_phone_number")

    @property
    @pulumi.getter(name="mdmPinlock")
    def mdm_pinlock(self) -> builtins.bool:
        """
        Mdm Pinlock
        """
        return pulumi.get(self, "mdm_pinlock")

    @property
    @pulumi.getter(name="mdmReachable")
    def mdm_reachable(self) -> builtins.bool:
        """
        Mdm Reachable
        """
        return pulumi.get(self, "mdm_reachable")

    @property
    @pulumi.getter(name="mdmSerial")
    def mdm_serial(self) -> builtins.str:
        """
        Mdm Serial
        """
        return pulumi.get(self, "mdm_serial")

    @property
    @pulumi.getter(name="mdmServerName")
    def mdm_server_name(self) -> builtins.str:
        """
        Mdm Server Name
        """
        return pulumi.get(self, "mdm_server_name")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        """
        The name of the endpoint
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="portalUser")
    def portal_user(self) -> builtins.str:
        """
        Portal User
        """
        return pulumi.get(self, "portal_user")

    @property
    @pulumi.getter(name="profileId")
    def profile_id(self) -> builtins.str:
        """
        Profile ID
        """
        return pulumi.get(self, "profile_id")

    @property
    @pulumi.getter(name="staticGroupAssignment")
    def static_group_assignment(self) -> builtins.bool:
        """
        Static Group Assignment
        """
        return pulumi.get(self, "static_group_assignment")

    @property
    @pulumi.getter(name="staticGroupAssignmentDefined")
    def static_group_assignment_defined(self) -> builtins.bool:
        """
        staticGroupAssignmentDefined
        """
        return pulumi.get(self, "static_group_assignment_defined")

    @property
    @pulumi.getter(name="staticProfileAssignment")
    def static_profile_assignment(self) -> builtins.bool:
        """
        Static Profile Assignment
        """
        return pulumi.get(self, "static_profile_assignment")

    @property
    @pulumi.getter(name="staticProfileAssignmentDefined")
    def static_profile_assignment_defined(self) -> builtins.bool:
        """
        Static Profile Assignment Defined
        """
        return pulumi.get(self, "static_profile_assignment_defined")


class AwaitableGetEndpointResult(GetEndpointResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEndpointResult(
            custom_attributes=self.custom_attributes,
            description=self.description,
            group_id=self.group_id,
            id=self.id,
            identity_store=self.identity_store,
            identity_store_id=self.identity_store_id,
            mac=self.mac,
            mdm_compliance_status=self.mdm_compliance_status,
            mdm_encrypted=self.mdm_encrypted,
            mdm_enrolled=self.mdm_enrolled,
            mdm_imei=self.mdm_imei,
            mdm_jail_broken=self.mdm_jail_broken,
            mdm_manufacturer=self.mdm_manufacturer,
            mdm_model=self.mdm_model,
            mdm_os=self.mdm_os,
            mdm_phone_number=self.mdm_phone_number,
            mdm_pinlock=self.mdm_pinlock,
            mdm_reachable=self.mdm_reachable,
            mdm_serial=self.mdm_serial,
            mdm_server_name=self.mdm_server_name,
            name=self.name,
            portal_user=self.portal_user,
            profile_id=self.profile_id,
            static_group_assignment=self.static_group_assignment,
            static_group_assignment_defined=self.static_group_assignment_defined,
            static_profile_assignment=self.static_profile_assignment,
            static_profile_assignment_defined=self.static_profile_assignment_defined)


def get_endpoint(id: Optional[builtins.str] = None,
                 name: Optional[builtins.str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEndpointResult:
    """
    This data source can read the Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_endpoint(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the endpoint
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:identitymanagement/getEndpoint:getEndpoint', __args__, opts=opts, typ=GetEndpointResult).value

    return AwaitableGetEndpointResult(
        custom_attributes=pulumi.get(__ret__, 'custom_attributes'),
        description=pulumi.get(__ret__, 'description'),
        group_id=pulumi.get(__ret__, 'group_id'),
        id=pulumi.get(__ret__, 'id'),
        identity_store=pulumi.get(__ret__, 'identity_store'),
        identity_store_id=pulumi.get(__ret__, 'identity_store_id'),
        mac=pulumi.get(__ret__, 'mac'),
        mdm_compliance_status=pulumi.get(__ret__, 'mdm_compliance_status'),
        mdm_encrypted=pulumi.get(__ret__, 'mdm_encrypted'),
        mdm_enrolled=pulumi.get(__ret__, 'mdm_enrolled'),
        mdm_imei=pulumi.get(__ret__, 'mdm_imei'),
        mdm_jail_broken=pulumi.get(__ret__, 'mdm_jail_broken'),
        mdm_manufacturer=pulumi.get(__ret__, 'mdm_manufacturer'),
        mdm_model=pulumi.get(__ret__, 'mdm_model'),
        mdm_os=pulumi.get(__ret__, 'mdm_os'),
        mdm_phone_number=pulumi.get(__ret__, 'mdm_phone_number'),
        mdm_pinlock=pulumi.get(__ret__, 'mdm_pinlock'),
        mdm_reachable=pulumi.get(__ret__, 'mdm_reachable'),
        mdm_serial=pulumi.get(__ret__, 'mdm_serial'),
        mdm_server_name=pulumi.get(__ret__, 'mdm_server_name'),
        name=pulumi.get(__ret__, 'name'),
        portal_user=pulumi.get(__ret__, 'portal_user'),
        profile_id=pulumi.get(__ret__, 'profile_id'),
        static_group_assignment=pulumi.get(__ret__, 'static_group_assignment'),
        static_group_assignment_defined=pulumi.get(__ret__, 'static_group_assignment_defined'),
        static_profile_assignment=pulumi.get(__ret__, 'static_profile_assignment'),
        static_profile_assignment_defined=pulumi.get(__ret__, 'static_profile_assignment_defined'))
def get_endpoint_output(id: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        name: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetEndpointResult]:
    """
    This data source can read the Endpoint.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.identitymanagement.get_endpoint(id="76d24097-41c4-4558-a4d0-a8c07ac08470")
    ```


    :param builtins.str id: The id of the object
    :param builtins.str name: The name of the endpoint
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:identitymanagement/getEndpoint:getEndpoint', __args__, opts=opts, typ=GetEndpointResult)
    return __ret__.apply(lambda __response__: GetEndpointResult(
        custom_attributes=pulumi.get(__response__, 'custom_attributes'),
        description=pulumi.get(__response__, 'description'),
        group_id=pulumi.get(__response__, 'group_id'),
        id=pulumi.get(__response__, 'id'),
        identity_store=pulumi.get(__response__, 'identity_store'),
        identity_store_id=pulumi.get(__response__, 'identity_store_id'),
        mac=pulumi.get(__response__, 'mac'),
        mdm_compliance_status=pulumi.get(__response__, 'mdm_compliance_status'),
        mdm_encrypted=pulumi.get(__response__, 'mdm_encrypted'),
        mdm_enrolled=pulumi.get(__response__, 'mdm_enrolled'),
        mdm_imei=pulumi.get(__response__, 'mdm_imei'),
        mdm_jail_broken=pulumi.get(__response__, 'mdm_jail_broken'),
        mdm_manufacturer=pulumi.get(__response__, 'mdm_manufacturer'),
        mdm_model=pulumi.get(__response__, 'mdm_model'),
        mdm_os=pulumi.get(__response__, 'mdm_os'),
        mdm_phone_number=pulumi.get(__response__, 'mdm_phone_number'),
        mdm_pinlock=pulumi.get(__response__, 'mdm_pinlock'),
        mdm_reachable=pulumi.get(__response__, 'mdm_reachable'),
        mdm_serial=pulumi.get(__response__, 'mdm_serial'),
        mdm_server_name=pulumi.get(__response__, 'mdm_server_name'),
        name=pulumi.get(__response__, 'name'),
        portal_user=pulumi.get(__response__, 'portal_user'),
        profile_id=pulumi.get(__response__, 'profile_id'),
        static_group_assignment=pulumi.get(__response__, 'static_group_assignment'),
        static_group_assignment_defined=pulumi.get(__response__, 'static_group_assignment_defined'),
        static_profile_assignment=pulumi.get(__response__, 'static_profile_assignment'),
        static_profile_assignment_defined=pulumi.get(__response__, 'static_profile_assignment_defined')))

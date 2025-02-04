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

__all__ = [
    'GetDictionaryResult',
    'AwaitableGetDictionaryResult',
    'get_dictionary',
    'get_dictionary_output',
]

@pulumi.output_type
class GetDictionaryResult:
    """
    A collection of values returned by getDictionary.
    """
    def __init__(__self__, description=None, dictionary_attr_type=None, id=None, name=None, version=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if dictionary_attr_type and not isinstance(dictionary_attr_type, str):
            raise TypeError("Expected argument 'dictionary_attr_type' to be a str")
        pulumi.set(__self__, "dictionary_attr_type", dictionary_attr_type)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The description of the dictionary
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dictionaryAttrType")
    def dictionary_attr_type(self) -> str:
        """
        The dictionary attribute type
        """
        return pulumi.get(self, "dictionary_attr_type")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The id of the object
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The dictionary name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        The version of the dictionary
        """
        return pulumi.get(self, "version")


class AwaitableGetDictionaryResult(GetDictionaryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDictionaryResult(
            description=self.description,
            dictionary_attr_type=self.dictionary_attr_type,
            id=self.id,
            name=self.name,
            version=self.version)


def get_dictionary(id: Optional[str] = None,
                   name: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDictionaryResult:
    """
    This data source can read the Network Access Dictionary.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_dictionary(id="Dict1")
    ```


    :param str id: The id of the object
    :param str name: The dictionary name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ise:networkaccess/getDictionary:getDictionary', __args__, opts=opts, typ=GetDictionaryResult).value

    return AwaitableGetDictionaryResult(
        description=pulumi.get(__ret__, 'description'),
        dictionary_attr_type=pulumi.get(__ret__, 'dictionary_attr_type'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        version=pulumi.get(__ret__, 'version'))
def get_dictionary_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                          name: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetDictionaryResult]:
    """
    This data source can read the Network Access Dictionary.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ise as ise

    example = ise.networkaccess.get_dictionary(id="Dict1")
    ```


    :param str id: The id of the object
    :param str name: The dictionary name
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('ise:networkaccess/getDictionary:getDictionary', __args__, opts=opts, typ=GetDictionaryResult)
    return __ret__.apply(lambda __response__: GetDictionaryResult(
        description=pulumi.get(__response__, 'description'),
        dictionary_attr_type=pulumi.get(__response__, 'dictionary_attr_type'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        version=pulumi.get(__response__, 'version')))

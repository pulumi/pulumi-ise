# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

insecure: Optional[bool]
"""
Allow insecure HTTPS client. This can also be set as the ISE_INSECURE environment variable. Defaults to `true`.
"""

password: Optional[str]
"""
Password for the ISE instance. This can also be set as the ISE_PASSWORD environment variable.
"""

retries: Optional[int]
"""
Number of retries for REST API calls. This can also be set as the ISE_RETRIES environment variable. Defaults to `3`.
"""

url: Optional[str]
"""
URL of the Cisco ISE instance. This can also be set as the ISE_URL environment variable.
"""

username: Optional[str]
"""
Username for the ISE instance. This can also be set as the ISE_USERNAME environment variable.
"""


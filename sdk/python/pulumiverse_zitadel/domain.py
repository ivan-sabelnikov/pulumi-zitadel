# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DomainArgs', 'Domain']

@pulumi.input_type
class DomainArgs:
    def __init__(__self__, *,
                 is_primary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Domain resource.
        :param pulumi.Input[bool] is_primary: Is domain primary
        :param pulumi.Input[str] name: Name of the domain
        :param pulumi.Input[str] org_id: ID of the organization
        """
        DomainArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            is_primary=is_primary,
            name=name,
            org_id=org_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             is_primary: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'isPrimary' in kwargs:
            is_primary = kwargs['isPrimary']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']

        if is_primary is not None:
            _setter("is_primary", is_primary)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)

    @property
    @pulumi.getter(name="isPrimary")
    def is_primary(self) -> Optional[pulumi.Input[bool]]:
        """
        Is domain primary
        """
        return pulumi.get(self, "is_primary")

    @is_primary.setter
    def is_primary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_primary", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the domain
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)


@pulumi.input_type
class _DomainState:
    def __init__(__self__, *,
                 is_primary: Optional[pulumi.Input[bool]] = None,
                 is_verified: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 validation_type: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Domain resources.
        :param pulumi.Input[bool] is_primary: Is domain primary
        :param pulumi.Input[bool] is_verified: Is domain verified
        :param pulumi.Input[str] name: Name of the domain
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[int] validation_type: Validation type
        """
        _DomainState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            is_primary=is_primary,
            is_verified=is_verified,
            name=name,
            org_id=org_id,
            validation_type=validation_type,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             is_primary: Optional[pulumi.Input[bool]] = None,
             is_verified: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             org_id: Optional[pulumi.Input[str]] = None,
             validation_type: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if 'isPrimary' in kwargs:
            is_primary = kwargs['isPrimary']
        if 'isVerified' in kwargs:
            is_verified = kwargs['isVerified']
        if 'orgId' in kwargs:
            org_id = kwargs['orgId']
        if 'validationType' in kwargs:
            validation_type = kwargs['validationType']

        if is_primary is not None:
            _setter("is_primary", is_primary)
        if is_verified is not None:
            _setter("is_verified", is_verified)
        if name is not None:
            _setter("name", name)
        if org_id is not None:
            _setter("org_id", org_id)
        if validation_type is not None:
            _setter("validation_type", validation_type)

    @property
    @pulumi.getter(name="isPrimary")
    def is_primary(self) -> Optional[pulumi.Input[bool]]:
        """
        Is domain primary
        """
        return pulumi.get(self, "is_primary")

    @is_primary.setter
    def is_primary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_primary", value)

    @property
    @pulumi.getter(name="isVerified")
    def is_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Is domain verified
        """
        return pulumi.get(self, "is_verified")

    @is_verified.setter
    def is_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_verified", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the domain
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="validationType")
    def validation_type(self) -> Optional[pulumi.Input[int]]:
        """
        Validation type
        """
        return pulumi.get(self, "validation_type")

    @validation_type.setter
    def validation_type(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "validation_type", value)


class Domain(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_primary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource representing a domain of the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.Domain("default",
            org_id=data["zitadel_org"]["default"]["id"],
            is_primary=False)
        ```

        ## Import

        terraform The resource can be imported using the ID format `name[:org_id]`, e.g.

        ```sh
         $ pulumi import zitadel:index/domain:Domain imported 'example.com:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_primary: Is domain primary
        :param pulumi.Input[str] name: Name of the domain
        :param pulumi.Input[str] org_id: ID of the organization
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[DomainArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource representing a domain of the organization.

        ## Example Usage

        ```python
        import pulumi
        import pulumiverse_zitadel as zitadel

        default = zitadel.Domain("default",
            org_id=data["zitadel_org"]["default"]["id"],
            is_primary=False)
        ```

        ## Import

        terraform The resource can be imported using the ID format `name[:org_id]`, e.g.

        ```sh
         $ pulumi import zitadel:index/domain:Domain imported 'example.com:123456789012345678'
        ```

        :param str resource_name: The name of the resource.
        :param DomainArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DomainArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DomainArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 is_primary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DomainArgs.__new__(DomainArgs)

            __props__.__dict__["is_primary"] = is_primary
            __props__.__dict__["name"] = name
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["is_verified"] = None
            __props__.__dict__["validation_type"] = None
        super(Domain, __self__).__init__(
            'zitadel:index/domain:Domain',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            is_primary: Optional[pulumi.Input[bool]] = None,
            is_verified: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            validation_type: Optional[pulumi.Input[int]] = None) -> 'Domain':
        """
        Get an existing Domain resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] is_primary: Is domain primary
        :param pulumi.Input[bool] is_verified: Is domain verified
        :param pulumi.Input[str] name: Name of the domain
        :param pulumi.Input[str] org_id: ID of the organization
        :param pulumi.Input[int] validation_type: Validation type
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DomainState.__new__(_DomainState)

        __props__.__dict__["is_primary"] = is_primary
        __props__.__dict__["is_verified"] = is_verified
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["validation_type"] = validation_type
        return Domain(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="isPrimary")
    def is_primary(self) -> pulumi.Output[Optional[bool]]:
        """
        Is domain primary
        """
        return pulumi.get(self, "is_primary")

    @property
    @pulumi.getter(name="isVerified")
    def is_verified(self) -> pulumi.Output[bool]:
        """
        Is domain verified
        """
        return pulumi.get(self, "is_verified")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the domain
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the organization
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="validationType")
    def validation_type(self) -> pulumi.Output[int]:
        """
        Validation type
        """
        return pulumi.get(self, "validation_type")


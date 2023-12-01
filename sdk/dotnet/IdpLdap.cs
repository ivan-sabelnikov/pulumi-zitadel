// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Pulumiverse.Zitadel
{
    /// <summary>
    /// Resource representing an LDAP IDP on the instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Zitadel.IdpLdap("default", new()
    ///     {
    ///         BaseDn = "dc=example,dc=com",
    ///         BindDn = "cn=admin,dc=example,dc=com",
    ///         BindPassword = "Password1!",
    ///         FirstNameAttribute = "firstname",
    ///         IdAttribute = "uid",
    ///         IsAutoCreation = false,
    ///         IsAutoUpdate = true,
    ///         IsCreationAllowed = true,
    ///         IsLinkingAllowed = false,
    ///         LastNameAttribute = "lastname",
    ///         Servers = new[]
    ///         {
    ///             "ldaps://my.primary.server:389",
    ///             "ldaps://my.secondary.server:389",
    ///         },
    ///         StartTls = false,
    ///         Timeout = "10s",
    ///         UserBase = "dn",
    ///         UserFilters = new[]
    ///         {
    ///             "uid",
    ///             "email",
    ///         },
    ///         UserObjectClasses = new[]
    ///         {
    ///             "inetOrgPerson",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// terraform The resource can be imported using the ID format `&lt;id[:bind_password]&gt;`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import zitadel:index/idpLdap:IdpLdap imported '123456789012345678:b1nd_p4ssw0rd'
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/idpLdap:IdpLdap")]
    public partial class IdpLdap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// User attribute for the avatar url
        /// </summary>
        [Output("avatarUrlAttribute")]
        public Output<string?> AvatarUrlAttribute { get; private set; } = null!;

        /// <summary>
        /// Base DN for LDAP connections
        /// </summary>
        [Output("baseDn")]
        public Output<string> BaseDn { get; private set; } = null!;

        /// <summary>
        /// Bind DN for LDAP connections
        /// </summary>
        [Output("bindDn")]
        public Output<string> BindDn { get; private set; } = null!;

        /// <summary>
        /// Bind password for LDAP connections
        /// </summary>
        [Output("bindPassword")]
        public Output<string> BindPassword { get; private set; } = null!;

        /// <summary>
        /// User attribute for the display name
        /// </summary>
        [Output("displayNameAttribute")]
        public Output<string?> DisplayNameAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the email
        /// </summary>
        [Output("emailAttribute")]
        public Output<string?> EmailAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the email verified state
        /// </summary>
        [Output("emailVerifiedAttribute")]
        public Output<string?> EmailVerifiedAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the first name
        /// </summary>
        [Output("firstNameAttribute")]
        public Output<string?> FirstNameAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the id
        /// </summary>
        [Output("idAttribute")]
        public Output<string?> IdAttribute { get; private set; } = null!;

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Output("isAutoCreation")]
        public Output<bool> IsAutoCreation { get; private set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Output("isAutoUpdate")]
        public Output<bool> IsAutoUpdate { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Output("isCreationAllowed")]
        public Output<bool> IsCreationAllowed { get; private set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Output("isLinkingAllowed")]
        public Output<bool> IsLinkingAllowed { get; private set; } = null!;

        /// <summary>
        /// User attribute for the last name
        /// </summary>
        [Output("lastNameAttribute")]
        public Output<string?> LastNameAttribute { get; private set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// User attribute for the nick name
        /// </summary>
        [Output("nickNameAttribute")]
        public Output<string?> NickNameAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the phone
        /// </summary>
        [Output("phoneAttribute")]
        public Output<string?> PhoneAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the phone verified state
        /// </summary>
        [Output("phoneVerifiedAttribute")]
        public Output<string?> PhoneVerifiedAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the preferred language
        /// </summary>
        [Output("preferredLanguageAttribute")]
        public Output<string?> PreferredLanguageAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the preferred username
        /// </summary>
        [Output("preferredUsernameAttribute")]
        public Output<string?> PreferredUsernameAttribute { get; private set; } = null!;

        /// <summary>
        /// User attribute for the profile
        /// </summary>
        [Output("profileAttribute")]
        public Output<string?> ProfileAttribute { get; private set; } = null!;

        /// <summary>
        /// Servers to try in order for establishing LDAP connections
        /// </summary>
        [Output("servers")]
        public Output<ImmutableArray<string>> Servers { get; private set; } = null!;

        /// <summary>
        /// Wether to use StartTLS for LDAP connections
        /// </summary>
        [Output("startTls")]
        public Output<bool> StartTls { get; private set; } = null!;

        /// <summary>
        /// Timeout for LDAP connections
        /// </summary>
        [Output("timeout")]
        public Output<string> Timeout { get; private set; } = null!;

        /// <summary>
        /// User base for LDAP connections
        /// </summary>
        [Output("userBase")]
        public Output<string> UserBase { get; private set; } = null!;

        /// <summary>
        /// User filters for LDAP connections
        /// </summary>
        [Output("userFilters")]
        public Output<ImmutableArray<string>> UserFilters { get; private set; } = null!;

        /// <summary>
        /// User object classes for LDAP connections
        /// </summary>
        [Output("userObjectClasses")]
        public Output<ImmutableArray<string>> UserObjectClasses { get; private set; } = null!;


        /// <summary>
        /// Create a IdpLdap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdpLdap(string name, IdpLdapArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/idpLdap:IdpLdap", name, args ?? new IdpLdapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdpLdap(string name, Input<string> id, IdpLdapState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/idpLdap:IdpLdap", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
                AdditionalSecretOutputs =
                {
                    "bindPassword",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdpLdap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdpLdap Get(string name, Input<string> id, IdpLdapState? state = null, CustomResourceOptions? options = null)
        {
            return new IdpLdap(name, id, state, options);
        }
    }

    public sealed class IdpLdapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User attribute for the avatar url
        /// </summary>
        [Input("avatarUrlAttribute")]
        public Input<string>? AvatarUrlAttribute { get; set; }

        /// <summary>
        /// Base DN for LDAP connections
        /// </summary>
        [Input("baseDn", required: true)]
        public Input<string> BaseDn { get; set; } = null!;

        /// <summary>
        /// Bind DN for LDAP connections
        /// </summary>
        [Input("bindDn", required: true)]
        public Input<string> BindDn { get; set; } = null!;

        [Input("bindPassword", required: true)]
        private Input<string>? _bindPassword;

        /// <summary>
        /// Bind password for LDAP connections
        /// </summary>
        public Input<string>? BindPassword
        {
            get => _bindPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bindPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User attribute for the display name
        /// </summary>
        [Input("displayNameAttribute")]
        public Input<string>? DisplayNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the email
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// User attribute for the email verified state
        /// </summary>
        [Input("emailVerifiedAttribute")]
        public Input<string>? EmailVerifiedAttribute { get; set; }

        /// <summary>
        /// User attribute for the first name
        /// </summary>
        [Input("firstNameAttribute")]
        public Input<string>? FirstNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the id
        /// </summary>
        [Input("idAttribute")]
        public Input<string>? IdAttribute { get; set; }

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation", required: true)]
        public Input<bool> IsAutoCreation { get; set; } = null!;

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate", required: true)]
        public Input<bool> IsAutoUpdate { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed", required: true)]
        public Input<bool> IsCreationAllowed { get; set; } = null!;

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed", required: true)]
        public Input<bool> IsLinkingAllowed { get; set; } = null!;

        /// <summary>
        /// User attribute for the last name
        /// </summary>
        [Input("lastNameAttribute")]
        public Input<string>? LastNameAttribute { get; set; }

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User attribute for the nick name
        /// </summary>
        [Input("nickNameAttribute")]
        public Input<string>? NickNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the phone
        /// </summary>
        [Input("phoneAttribute")]
        public Input<string>? PhoneAttribute { get; set; }

        /// <summary>
        /// User attribute for the phone verified state
        /// </summary>
        [Input("phoneVerifiedAttribute")]
        public Input<string>? PhoneVerifiedAttribute { get; set; }

        /// <summary>
        /// User attribute for the preferred language
        /// </summary>
        [Input("preferredLanguageAttribute")]
        public Input<string>? PreferredLanguageAttribute { get; set; }

        /// <summary>
        /// User attribute for the preferred username
        /// </summary>
        [Input("preferredUsernameAttribute")]
        public Input<string>? PreferredUsernameAttribute { get; set; }

        /// <summary>
        /// User attribute for the profile
        /// </summary>
        [Input("profileAttribute")]
        public Input<string>? ProfileAttribute { get; set; }

        [Input("servers", required: true)]
        private InputList<string>? _servers;

        /// <summary>
        /// Servers to try in order for establishing LDAP connections
        /// </summary>
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        /// <summary>
        /// Wether to use StartTLS for LDAP connections
        /// </summary>
        [Input("startTls", required: true)]
        public Input<bool> StartTls { get; set; } = null!;

        /// <summary>
        /// Timeout for LDAP connections
        /// </summary>
        [Input("timeout", required: true)]
        public Input<string> Timeout { get; set; } = null!;

        /// <summary>
        /// User base for LDAP connections
        /// </summary>
        [Input("userBase", required: true)]
        public Input<string> UserBase { get; set; } = null!;

        [Input("userFilters", required: true)]
        private InputList<string>? _userFilters;

        /// <summary>
        /// User filters for LDAP connections
        /// </summary>
        public InputList<string> UserFilters
        {
            get => _userFilters ?? (_userFilters = new InputList<string>());
            set => _userFilters = value;
        }

        [Input("userObjectClasses", required: true)]
        private InputList<string>? _userObjectClasses;

        /// <summary>
        /// User object classes for LDAP connections
        /// </summary>
        public InputList<string> UserObjectClasses
        {
            get => _userObjectClasses ?? (_userObjectClasses = new InputList<string>());
            set => _userObjectClasses = value;
        }

        public IdpLdapArgs()
        {
        }
        public static new IdpLdapArgs Empty => new IdpLdapArgs();
    }

    public sealed class IdpLdapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User attribute for the avatar url
        /// </summary>
        [Input("avatarUrlAttribute")]
        public Input<string>? AvatarUrlAttribute { get; set; }

        /// <summary>
        /// Base DN for LDAP connections
        /// </summary>
        [Input("baseDn")]
        public Input<string>? BaseDn { get; set; }

        /// <summary>
        /// Bind DN for LDAP connections
        /// </summary>
        [Input("bindDn")]
        public Input<string>? BindDn { get; set; }

        [Input("bindPassword")]
        private Input<string>? _bindPassword;

        /// <summary>
        /// Bind password for LDAP connections
        /// </summary>
        public Input<string>? BindPassword
        {
            get => _bindPassword;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _bindPassword = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// User attribute for the display name
        /// </summary>
        [Input("displayNameAttribute")]
        public Input<string>? DisplayNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the email
        /// </summary>
        [Input("emailAttribute")]
        public Input<string>? EmailAttribute { get; set; }

        /// <summary>
        /// User attribute for the email verified state
        /// </summary>
        [Input("emailVerifiedAttribute")]
        public Input<string>? EmailVerifiedAttribute { get; set; }

        /// <summary>
        /// User attribute for the first name
        /// </summary>
        [Input("firstNameAttribute")]
        public Input<string>? FirstNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the id
        /// </summary>
        [Input("idAttribute")]
        public Input<string>? IdAttribute { get; set; }

        /// <summary>
        /// enable if a new account in ZITADEL should be created automatically on login with an external account
        /// </summary>
        [Input("isAutoCreation")]
        public Input<bool>? IsAutoCreation { get; set; }

        /// <summary>
        /// enable if a the ZITADEL account fields should be updated automatically on each login
        /// </summary>
        [Input("isAutoUpdate")]
        public Input<bool>? IsAutoUpdate { get; set; }

        /// <summary>
        /// enable if users should be able to create a new account in ZITADEL when using an external account
        /// </summary>
        [Input("isCreationAllowed")]
        public Input<bool>? IsCreationAllowed { get; set; }

        /// <summary>
        /// enable if users should be able to link an existing ZITADEL user with an external account
        /// </summary>
        [Input("isLinkingAllowed")]
        public Input<bool>? IsLinkingAllowed { get; set; }

        /// <summary>
        /// User attribute for the last name
        /// </summary>
        [Input("lastNameAttribute")]
        public Input<string>? LastNameAttribute { get; set; }

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// User attribute for the nick name
        /// </summary>
        [Input("nickNameAttribute")]
        public Input<string>? NickNameAttribute { get; set; }

        /// <summary>
        /// User attribute for the phone
        /// </summary>
        [Input("phoneAttribute")]
        public Input<string>? PhoneAttribute { get; set; }

        /// <summary>
        /// User attribute for the phone verified state
        /// </summary>
        [Input("phoneVerifiedAttribute")]
        public Input<string>? PhoneVerifiedAttribute { get; set; }

        /// <summary>
        /// User attribute for the preferred language
        /// </summary>
        [Input("preferredLanguageAttribute")]
        public Input<string>? PreferredLanguageAttribute { get; set; }

        /// <summary>
        /// User attribute for the preferred username
        /// </summary>
        [Input("preferredUsernameAttribute")]
        public Input<string>? PreferredUsernameAttribute { get; set; }

        /// <summary>
        /// User attribute for the profile
        /// </summary>
        [Input("profileAttribute")]
        public Input<string>? ProfileAttribute { get; set; }

        [Input("servers")]
        private InputList<string>? _servers;

        /// <summary>
        /// Servers to try in order for establishing LDAP connections
        /// </summary>
        public InputList<string> Servers
        {
            get => _servers ?? (_servers = new InputList<string>());
            set => _servers = value;
        }

        /// <summary>
        /// Wether to use StartTLS for LDAP connections
        /// </summary>
        [Input("startTls")]
        public Input<bool>? StartTls { get; set; }

        /// <summary>
        /// Timeout for LDAP connections
        /// </summary>
        [Input("timeout")]
        public Input<string>? Timeout { get; set; }

        /// <summary>
        /// User base for LDAP connections
        /// </summary>
        [Input("userBase")]
        public Input<string>? UserBase { get; set; }

        [Input("userFilters")]
        private InputList<string>? _userFilters;

        /// <summary>
        /// User filters for LDAP connections
        /// </summary>
        public InputList<string> UserFilters
        {
            get => _userFilters ?? (_userFilters = new InputList<string>());
            set => _userFilters = value;
        }

        [Input("userObjectClasses")]
        private InputList<string>? _userObjectClasses;

        /// <summary>
        /// User object classes for LDAP connections
        /// </summary>
        public InputList<string> UserObjectClasses
        {
            get => _userObjectClasses ?? (_userObjectClasses = new InputList<string>());
            set => _userObjectClasses = value;
        }

        public IdpLdapState()
        {
        }
        public static new IdpLdapState Empty => new IdpLdapState();
    }
}

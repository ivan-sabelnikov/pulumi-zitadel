// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Vavsab.Zitadel
{
    /// <summary>
    /// Resource representing the custom login policy of an organization.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Vavsab.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var loginPolicy = new Zitadel.LoginPolicy("loginPolicy", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserLogin = true,
    ///         AllowRegister = true,
    ///         AllowExternalIdp = true,
    ///         ForceMfa = false,
    ///         PasswordlessType = "PASSWORDLESS_TYPE_ALLOWED",
    ///         HidePasswordReset = false,
    ///         PasswordCheckLifetime = "240h0m0s",
    ///         ExternalLoginCheckLifetime = "240h0m0s",
    ///         MultiFactorCheckLifetime = "24h0m0s",
    ///         MfaInitSkipLifetime = "720h0m0s",
    ///         SecondFactorCheckLifetime = "24h0m0s",
    ///         IgnoreUnknownUsernames = true,
    ///         DefaultRedirectUri = "localhost:8080",
    ///         SecondFactors = new[]
    ///         {
    ///             "SECOND_FACTOR_TYPE_OTP",
    ///             "SECOND_FACTOR_TYPE_U2F",
    ///         },
    ///         MultiFactors = new[]
    ///         {
    ///             "MULTI_FACTOR_TYPE_U2F_WITH_VERIFICATION",
    ///         },
    ///         Idps = new[]
    ///         {
    ///             zitadel_org_idp_oidc.Oidc_idp.Id,
    ///             zitadel_org_idp_jwt.Jwt_idp.Id,
    ///         },
    ///         AllowDomainDiscovery = true,
    ///         DisableLoginWithEmail = true,
    ///         DisableLoginWithPhone = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/loginPolicy:LoginPolicy")]
    public partial class LoginPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
        /// </summary>
        [Output("allowDomainDiscovery")]
        public Output<bool?> AllowDomainDiscovery { get; private set; } = null!;

        /// <summary>
        /// defines if a user is allowed to add a defined identity provider. E.g. Google auth
        /// </summary>
        [Output("allowExternalIdp")]
        public Output<bool> AllowExternalIdp { get; private set; } = null!;

        /// <summary>
        /// defines if a person is allowed to register a user on this organisation
        /// </summary>
        [Output("allowRegister")]
        public Output<bool> AllowRegister { get; private set; } = null!;

        /// <summary>
        /// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
        /// </summary>
        [Output("defaultRedirectUri")]
        public Output<string> DefaultRedirectUri { get; private set; } = null!;

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified email address
        /// </summary>
        [Output("disableLoginWithEmail")]
        public Output<bool?> DisableLoginWithEmail { get; private set; } = null!;

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified phone number
        /// </summary>
        [Output("disableLoginWithPhone")]
        public Output<bool?> DisableLoginWithPhone { get; private set; } = null!;

        [Output("externalLoginCheckLifetime")]
        public Output<string> ExternalLoginCheckLifetime { get; private set; } = null!;

        /// <summary>
        /// defines if a user MUST use a multi factor to log in
        /// </summary>
        [Output("forceMfa")]
        public Output<bool> ForceMfa { get; private set; } = null!;

        /// <summary>
        /// defines if password reset link should be shown in the login screen
        /// </summary>
        [Output("hidePasswordReset")]
        public Output<bool> HidePasswordReset { get; private set; } = null!;

        /// <summary>
        /// allowed idps to login or register
        /// </summary>
        [Output("idps")]
        public Output<ImmutableArray<string>> Idps { get; private set; } = null!;

        /// <summary>
        /// defines if unknown username on login screen directly return an error or always display the password screen
        /// </summary>
        [Output("ignoreUnknownUsernames")]
        public Output<bool> IgnoreUnknownUsernames { get; private set; } = null!;

        [Output("mfaInitSkipLifetime")]
        public Output<string> MfaInitSkipLifetime { get; private set; } = null!;

        [Output("multiFactorCheckLifetime")]
        public Output<string> MultiFactorCheckLifetime { get; private set; } = null!;

        /// <summary>
        /// allowed multi factors
        /// </summary>
        [Output("multiFactors")]
        public Output<ImmutableArray<string>> MultiFactors { get; private set; } = null!;

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        [Output("passwordCheckLifetime")]
        public Output<string> PasswordCheckLifetime { get; private set; } = null!;

        /// <summary>
        /// defines if passwordless is allowed for users
        /// </summary>
        [Output("passwordlessType")]
        public Output<string> PasswordlessType { get; private set; } = null!;

        [Output("secondFactorCheckLifetime")]
        public Output<string> SecondFactorCheckLifetime { get; private set; } = null!;

        /// <summary>
        /// allowed second factors
        /// </summary>
        [Output("secondFactors")]
        public Output<ImmutableArray<string>> SecondFactors { get; private set; } = null!;

        /// <summary>
        /// defines if a user is allowed to login with his username and password
        /// </summary>
        [Output("userLogin")]
        public Output<bool> UserLogin { get; private set; } = null!;


        /// <summary>
        /// Create a LoginPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoginPolicy(string name, LoginPolicyArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/loginPolicy:LoginPolicy", name, args ?? new LoginPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoginPolicy(string name, Input<string> id, LoginPolicyState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/loginPolicy:LoginPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/vavsab",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoginPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoginPolicy Get(string name, Input<string> id, LoginPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new LoginPolicy(name, id, state, options);
        }
    }

    public sealed class LoginPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
        /// </summary>
        [Input("allowDomainDiscovery")]
        public Input<bool>? AllowDomainDiscovery { get; set; }

        /// <summary>
        /// defines if a user is allowed to add a defined identity provider. E.g. Google auth
        /// </summary>
        [Input("allowExternalIdp", required: true)]
        public Input<bool> AllowExternalIdp { get; set; } = null!;

        /// <summary>
        /// defines if a person is allowed to register a user on this organisation
        /// </summary>
        [Input("allowRegister", required: true)]
        public Input<bool> AllowRegister { get; set; } = null!;

        /// <summary>
        /// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
        /// </summary>
        [Input("defaultRedirectUri", required: true)]
        public Input<string> DefaultRedirectUri { get; set; } = null!;

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified email address
        /// </summary>
        [Input("disableLoginWithEmail")]
        public Input<bool>? DisableLoginWithEmail { get; set; }

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified phone number
        /// </summary>
        [Input("disableLoginWithPhone")]
        public Input<bool>? DisableLoginWithPhone { get; set; }

        [Input("externalLoginCheckLifetime", required: true)]
        public Input<string> ExternalLoginCheckLifetime { get; set; } = null!;

        /// <summary>
        /// defines if a user MUST use a multi factor to log in
        /// </summary>
        [Input("forceMfa", required: true)]
        public Input<bool> ForceMfa { get; set; } = null!;

        /// <summary>
        /// defines if password reset link should be shown in the login screen
        /// </summary>
        [Input("hidePasswordReset", required: true)]
        public Input<bool> HidePasswordReset { get; set; } = null!;

        [Input("idps")]
        private InputList<string>? _idps;

        /// <summary>
        /// allowed idps to login or register
        /// </summary>
        public InputList<string> Idps
        {
            get => _idps ?? (_idps = new InputList<string>());
            set => _idps = value;
        }

        /// <summary>
        /// defines if unknown username on login screen directly return an error or always display the password screen
        /// </summary>
        [Input("ignoreUnknownUsernames", required: true)]
        public Input<bool> IgnoreUnknownUsernames { get; set; } = null!;

        [Input("mfaInitSkipLifetime", required: true)]
        public Input<string> MfaInitSkipLifetime { get; set; } = null!;

        [Input("multiFactorCheckLifetime", required: true)]
        public Input<string> MultiFactorCheckLifetime { get; set; } = null!;

        [Input("multiFactors")]
        private InputList<string>? _multiFactors;

        /// <summary>
        /// allowed multi factors
        /// </summary>
        public InputList<string> MultiFactors
        {
            get => _multiFactors ?? (_multiFactors = new InputList<string>());
            set => _multiFactors = value;
        }

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        [Input("passwordCheckLifetime", required: true)]
        public Input<string> PasswordCheckLifetime { get; set; } = null!;

        /// <summary>
        /// defines if passwordless is allowed for users
        /// </summary>
        [Input("passwordlessType", required: true)]
        public Input<string> PasswordlessType { get; set; } = null!;

        [Input("secondFactorCheckLifetime", required: true)]
        public Input<string> SecondFactorCheckLifetime { get; set; } = null!;

        [Input("secondFactors")]
        private InputList<string>? _secondFactors;

        /// <summary>
        /// allowed second factors
        /// </summary>
        public InputList<string> SecondFactors
        {
            get => _secondFactors ?? (_secondFactors = new InputList<string>());
            set => _secondFactors = value;
        }

        /// <summary>
        /// defines if a user is allowed to login with his username and password
        /// </summary>
        [Input("userLogin", required: true)]
        public Input<bool> UserLogin { get; set; } = null!;

        public LoginPolicyArgs()
        {
        }
        public static new LoginPolicyArgs Empty => new LoginPolicyArgs();
    }

    public sealed class LoginPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// if set to true, the suffix (@domain.com) of an unknown username input on the login screen will be matched against the org domains and will redirect to the registration of that organisation on success.
        /// </summary>
        [Input("allowDomainDiscovery")]
        public Input<bool>? AllowDomainDiscovery { get; set; }

        /// <summary>
        /// defines if a user is allowed to add a defined identity provider. E.g. Google auth
        /// </summary>
        [Input("allowExternalIdp")]
        public Input<bool>? AllowExternalIdp { get; set; }

        /// <summary>
        /// defines if a person is allowed to register a user on this organisation
        /// </summary>
        [Input("allowRegister")]
        public Input<bool>? AllowRegister { get; set; }

        /// <summary>
        /// defines where the user will be redirected to if the login is started without app context (e.g. from mail)
        /// </summary>
        [Input("defaultRedirectUri")]
        public Input<string>? DefaultRedirectUri { get; set; }

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified email address
        /// </summary>
        [Input("disableLoginWithEmail")]
        public Input<bool>? DisableLoginWithEmail { get; set; }

        /// <summary>
        /// defines if user can additionally (to the loginname) be identified by their verified phone number
        /// </summary>
        [Input("disableLoginWithPhone")]
        public Input<bool>? DisableLoginWithPhone { get; set; }

        [Input("externalLoginCheckLifetime")]
        public Input<string>? ExternalLoginCheckLifetime { get; set; }

        /// <summary>
        /// defines if a user MUST use a multi factor to log in
        /// </summary>
        [Input("forceMfa")]
        public Input<bool>? ForceMfa { get; set; }

        /// <summary>
        /// defines if password reset link should be shown in the login screen
        /// </summary>
        [Input("hidePasswordReset")]
        public Input<bool>? HidePasswordReset { get; set; }

        [Input("idps")]
        private InputList<string>? _idps;

        /// <summary>
        /// allowed idps to login or register
        /// </summary>
        public InputList<string> Idps
        {
            get => _idps ?? (_idps = new InputList<string>());
            set => _idps = value;
        }

        /// <summary>
        /// defines if unknown username on login screen directly return an error or always display the password screen
        /// </summary>
        [Input("ignoreUnknownUsernames")]
        public Input<bool>? IgnoreUnknownUsernames { get; set; }

        [Input("mfaInitSkipLifetime")]
        public Input<string>? MfaInitSkipLifetime { get; set; }

        [Input("multiFactorCheckLifetime")]
        public Input<string>? MultiFactorCheckLifetime { get; set; }

        [Input("multiFactors")]
        private InputList<string>? _multiFactors;

        /// <summary>
        /// allowed multi factors
        /// </summary>
        public InputList<string> MultiFactors
        {
            get => _multiFactors ?? (_multiFactors = new InputList<string>());
            set => _multiFactors = value;
        }

        /// <summary>
        /// Id for the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("passwordCheckLifetime")]
        public Input<string>? PasswordCheckLifetime { get; set; }

        /// <summary>
        /// defines if passwordless is allowed for users
        /// </summary>
        [Input("passwordlessType")]
        public Input<string>? PasswordlessType { get; set; }

        [Input("secondFactorCheckLifetime")]
        public Input<string>? SecondFactorCheckLifetime { get; set; }

        [Input("secondFactors")]
        private InputList<string>? _secondFactors;

        /// <summary>
        /// allowed second factors
        /// </summary>
        public InputList<string> SecondFactors
        {
            get => _secondFactors ?? (_secondFactors = new InputList<string>());
            set => _secondFactors = value;
        }

        /// <summary>
        /// defines if a user is allowed to login with his username and password
        /// </summary>
        [Input("userLogin")]
        public Input<bool>? UserLogin { get; set; }

        public LoginPolicyState()
        {
        }
        public static new LoginPolicyState Empty => new LoginPolicyState();
    }
}

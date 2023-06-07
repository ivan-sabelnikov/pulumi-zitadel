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
    /// Resource representing a Google IDP on the instance.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zitadel = Pulumiverse.Zitadel;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var google = new Zitadel.IdpGoogle("google", new()
    ///     {
    ///         ClientId = "182902...",
    ///         ClientSecret = "GOCSPX-*****",
    ///         IsAutoCreation = false,
    ///         IsAutoUpdate = true,
    ///         IsCreationAllowed = true,
    ///         IsLinkingAllowed = false,
    ///         Scopes = new[]
    ///         {
    ///             "openid",
    ///             "profile",
    ///             "email",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/idpGoogle:IdpGoogle")]
    public partial class IdpGoogle : global::Pulumi.CustomResource
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

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
        /// Name of the IDP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;


        /// <summary>
        /// Create a IdpGoogle resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdpGoogle(string name, IdpGoogleArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/idpGoogle:IdpGoogle", name, args ?? new IdpGoogleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdpGoogle(string name, Input<string> id, IdpGoogleState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/idpGoogle:IdpGoogle", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/pulumiverse",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdpGoogle resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdpGoogle Get(string name, Input<string> id, IdpGoogleState? state = null, CustomResourceOptions? options = null)
        {
            return new IdpGoogle(name, id, state, options);
        }
    }

    public sealed class IdpGoogleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

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
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public IdpGoogleArgs()
        {
        }
        public static new IdpGoogleArgs Empty => new IdpGoogleArgs();
    }

    public sealed class IdpGoogleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

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
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        public IdpGoogleState()
        {
        }
        public static new IdpGoogleState Empty => new IdpGoogleState();
    }
}

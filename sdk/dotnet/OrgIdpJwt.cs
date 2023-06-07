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
    /// Resource representing a generic JWT IdP of the organization.
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
    ///     var jwtIdp = new Zitadel.OrgIdpJwt("jwtIdp", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         StylingType = "STYLING_TYPE_UNSPECIFIED",
    ///         JwtEndpoint = "https://jwtendpoint.com",
    ///         Issuer = "https://google.com",
    ///         KeysEndpoint = "https://jwtendpoint.com/keys",
    ///         HeaderName = "x-auth-token",
    ///         AutoRegister = false,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/orgIdpJwt:OrgIdpJwt")]
    public partial class OrgIdpJwt : global::Pulumi.CustomResource
    {
        /// <summary>
        /// auto register for users from this idp
        /// </summary>
        [Output("autoRegister")]
        public Output<bool> AutoRegister { get; private set; } = null!;

        /// <summary>
        /// the name of the header where the JWT is sent in, default is authorization
        /// </summary>
        [Output("headerName")]
        public Output<string> HeaderName { get; private set; } = null!;

        /// <summary>
        /// the issuer of the jwt (for validation)
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// the endpoint where the jwt can be extracted
        /// </summary>
        [Output("jwtEndpoint")]
        public Output<string> JwtEndpoint { get; private set; } = null!;

        /// <summary>
        /// the endpoint to the key (JWK) which are used to sign the JWT with
        /// </summary>
        [Output("keysEndpoint")]
        public Output<string> KeysEndpoint { get; private set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        /// </summary>
        [Output("stylingType")]
        public Output<string> StylingType { get; private set; } = null!;


        /// <summary>
        /// Create a OrgIdpJwt resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OrgIdpJwt(string name, OrgIdpJwtArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/orgIdpJwt:OrgIdpJwt", name, args ?? new OrgIdpJwtArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OrgIdpJwt(string name, Input<string> id, OrgIdpJwtState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/orgIdpJwt:OrgIdpJwt", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OrgIdpJwt resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OrgIdpJwt Get(string name, Input<string> id, OrgIdpJwtState? state = null, CustomResourceOptions? options = null)
        {
            return new OrgIdpJwt(name, id, state, options);
        }
    }

    public sealed class OrgIdpJwtArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// auto register for users from this idp
        /// </summary>
        [Input("autoRegister", required: true)]
        public Input<bool> AutoRegister { get; set; } = null!;

        /// <summary>
        /// the name of the header where the JWT is sent in, default is authorization
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// the issuer of the jwt (for validation)
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// the endpoint where the jwt can be extracted
        /// </summary>
        [Input("jwtEndpoint", required: true)]
        public Input<string> JwtEndpoint { get; set; } = null!;

        /// <summary>
        /// the endpoint to the key (JWK) which are used to sign the JWT with
        /// </summary>
        [Input("keysEndpoint", required: true)]
        public Input<string> KeysEndpoint { get; set; } = null!;

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        /// </summary>
        [Input("stylingType", required: true)]
        public Input<string> StylingType { get; set; } = null!;

        public OrgIdpJwtArgs()
        {
        }
        public static new OrgIdpJwtArgs Empty => new OrgIdpJwtArgs();
    }

    public sealed class OrgIdpJwtState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// auto register for users from this idp
        /// </summary>
        [Input("autoRegister")]
        public Input<bool>? AutoRegister { get; set; }

        /// <summary>
        /// the name of the header where the JWT is sent in, default is authorization
        /// </summary>
        [Input("headerName")]
        public Input<string>? HeaderName { get; set; }

        /// <summary>
        /// the issuer of the jwt (for validation)
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// the endpoint where the jwt can be extracted
        /// </summary>
        [Input("jwtEndpoint")]
        public Input<string>? JwtEndpoint { get; set; }

        /// <summary>
        /// the endpoint to the key (JWK) which are used to sign the JWT with
        /// </summary>
        [Input("keysEndpoint")]
        public Input<string>? KeysEndpoint { get; set; }

        /// <summary>
        /// Name of the IDP
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Some identity providers specify the styling of the button to their login, supported values: STYLING*TYPE*UNSPECIFIED, STYLING*TYPE*GOOGLE
        /// </summary>
        [Input("stylingType")]
        public Input<string>? StylingType { get; set; }

        public OrgIdpJwtState()
        {
        }
        public static new OrgIdpJwtState Empty => new OrgIdpJwtState();
    }
}

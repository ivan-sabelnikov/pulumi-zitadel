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
    public static class GetIdpGoogle
    {
        /// <summary>
        /// Datasource representing a Google IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetIdpGoogle.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdpGoogleResult> InvokeAsync(GetIdpGoogleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdpGoogleResult>("zitadel:index/getIdpGoogle:getIdpGoogle", args ?? new GetIdpGoogleArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing a Google IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Zitadel.GetIdpGoogle.Invoke(new()
        ///     {
        ///         Id = "123456789012345678",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIdpGoogleResult> Invoke(GetIdpGoogleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdpGoogleResult>("zitadel:index/getIdpGoogle:getIdpGoogle", args ?? new GetIdpGoogleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdpGoogleArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIdpGoogleArgs()
        {
        }
        public static new GetIdpGoogleArgs Empty => new GetIdpGoogleArgs();
    }

    public sealed class GetIdpGoogleInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIdpGoogleInvokeArgs()
        {
        }
        public static new GetIdpGoogleInvokeArgs Empty => new GetIdpGoogleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdpGoogleResult
    {
        /// <summary>
        /// client id generated by the identity provider
        /// </summary>
        public readonly string ClientId;
        /// <summary>
        /// client secret generated by the identity provider
        /// </summary>
        public readonly string ClientSecret;
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// enabled if a new account in ZITADEL are created automatically on login with an external account
        /// </summary>
        public readonly bool IsAutoCreation;
        /// <summary>
        /// enabled if a the ZITADEL account fields are updated automatically on each login
        /// </summary>
        public readonly bool IsAutoUpdate;
        /// <summary>
        /// enabled if users are able to create a new account in ZITADEL when using an external account
        /// </summary>
        public readonly bool IsCreationAllowed;
        /// <summary>
        /// enabled if users are able to link an existing ZITADEL user with an external account
        /// </summary>
        public readonly bool IsLinkingAllowed;
        /// <summary>
        /// Name of the IDP
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// the scopes requested by ZITADEL during the request on the identity provider
        /// </summary>
        public readonly ImmutableArray<string> Scopes;

        [OutputConstructor]
        private GetIdpGoogleResult(
            string clientId,

            string clientSecret,

            string id,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string name,

            ImmutableArray<string> scopes)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            Id = id;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            Name = name;
            Scopes = scopes;
        }
    }
}

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
    public static class GetIdpAzureAd
    {
        /// <summary>
        /// Datasource representing an Azure AD IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var azureAd = Zitadel.GetIdpAzureAd.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIdpAzureAdResult> InvokeAsync(GetIdpAzureAdArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetIdpAzureAdResult>("zitadel:index/getIdpAzureAd:getIdpAzureAd", args ?? new GetIdpAzureAdArgs(), options.WithDefaults());

        /// <summary>
        /// Datasource representing an Azure AD IDP on the instance.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zitadel = Pulumi.Zitadel;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var azureAd = Zitadel.GetIdpAzureAd.Invoke(new()
        ///     {
        ///         Id = "177073614158299139",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIdpAzureAdResult> Invoke(GetIdpAzureAdInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetIdpAzureAdResult>("zitadel:index/getIdpAzureAd:getIdpAzureAd", args ?? new GetIdpAzureAdInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdpAzureAdArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public string Id { get; set; } = null!;

        public GetIdpAzureAdArgs()
        {
        }
        public static new GetIdpAzureAdArgs Empty => new GetIdpAzureAdArgs();
    }

    public sealed class GetIdpAzureAdInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of this resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public GetIdpAzureAdInvokeArgs()
        {
        }
        public static new GetIdpAzureAdInvokeArgs Empty => new GetIdpAzureAdInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdpAzureAdResult
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
        /// automatically mark emails as verified
        /// </summary>
        public readonly bool EmailVerified;
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
        /// <summary>
        /// the azure ad tenant id
        /// </summary>
        public readonly string TenantId;
        /// <summary>
        /// the azure ad tenant type
        /// </summary>
        public readonly string TenantType;

        [OutputConstructor]
        private GetIdpAzureAdResult(
            string clientId,

            string clientSecret,

            bool emailVerified,

            string id,

            bool isAutoCreation,

            bool isAutoUpdate,

            bool isCreationAllowed,

            bool isLinkingAllowed,

            string name,

            ImmutableArray<string> scopes,

            string tenantId,

            string tenantType)
        {
            ClientId = clientId;
            ClientSecret = clientSecret;
            EmailVerified = emailVerified;
            Id = id;
            IsAutoCreation = isAutoCreation;
            IsAutoUpdate = isAutoUpdate;
            IsCreationAllowed = isCreationAllowed;
            IsLinkingAllowed = isLinkingAllowed;
            Name = name;
            Scopes = scopes;
            TenantId = tenantId;
            TenantType = tenantType;
        }
    }
}

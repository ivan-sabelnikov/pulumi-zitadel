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
    public static class GetDefaultOidcSettings
    {
        /// <summary>
        /// Datasource representing the default oidc settings.
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
        ///     var @default = Zitadel.GetDefaultOidcSettings.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oidcSettings"] = @default,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDefaultOidcSettingsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDefaultOidcSettingsResult>("zitadel:index/getDefaultOidcSettings:getDefaultOidcSettings", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Datasource representing the default oidc settings.
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
        ///     var @default = Zitadel.GetDefaultOidcSettings.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["oidcSettings"] = @default,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDefaultOidcSettingsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDefaultOidcSettingsResult>("zitadel:index/getDefaultOidcSettings:getDefaultOidcSettings", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetDefaultOidcSettingsResult
    {
        /// <summary>
        /// lifetime duration of access tokens
        /// </summary>
        public readonly string AccessTokenLifetime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// lifetime duration of id tokens
        /// </summary>
        public readonly string IdTokenLifetime;
        /// <summary>
        /// expiration duration of refresh tokens
        /// </summary>
        public readonly string RefreshTokenExpiration;
        /// <summary>
        /// expiration duration of idle refresh tokens
        /// </summary>
        public readonly string RefreshTokenIdleExpiration;

        [OutputConstructor]
        private GetDefaultOidcSettingsResult(
            string accessTokenLifetime,

            string id,

            string idTokenLifetime,

            string refreshTokenExpiration,

            string refreshTokenIdleExpiration)
        {
            AccessTokenLifetime = accessTokenLifetime;
            Id = id;
            IdTokenLifetime = idTokenLifetime;
            RefreshTokenExpiration = refreshTokenExpiration;
            RefreshTokenIdleExpiration = refreshTokenIdleExpiration;
        }
    }
}

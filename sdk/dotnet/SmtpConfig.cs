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
    /// Resource representing the SMTP configuration of an instance.
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
    ///     var smtp = new Zitadel.SmtpConfig("smtp", new()
    ///     {
    ///         Host = "localhost:25",
    ///         Password = "password",
    ///         SenderAddress = "address",
    ///         SenderName = "no-reply",
    ///         Tls = true,
    ///         User = "user",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/smtpConfig:SmtpConfig")]
    public partial class SmtpConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Host and port address to your SMTP server.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Password used to communicate with your SMTP server.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Address used to send emails.
        /// </summary>
        [Output("senderAddress")]
        public Output<string> SenderAddress { get; private set; } = null!;

        /// <summary>
        /// Sender name used to send emails.
        /// </summary>
        [Output("senderName")]
        public Output<string> SenderName { get; private set; } = null!;

        /// <summary>
        /// TLS used to communicate with your SMTP server.
        /// </summary>
        [Output("tls")]
        public Output<bool?> Tls { get; private set; } = null!;

        /// <summary>
        /// User used to communicate with your SMTP server.
        /// </summary>
        [Output("user")]
        public Output<string?> User { get; private set; } = null!;


        /// <summary>
        /// Create a SmtpConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SmtpConfig(string name, SmtpConfigArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/smtpConfig:SmtpConfig", name, args ?? new SmtpConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SmtpConfig(string name, Input<string> id, SmtpConfigState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/smtpConfig:SmtpConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SmtpConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SmtpConfig Get(string name, Input<string> id, SmtpConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new SmtpConfig(name, id, state, options);
        }
    }

    public sealed class SmtpConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host and port address to your SMTP server.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Password used to communicate with your SMTP server.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Address used to send emails.
        /// </summary>
        [Input("senderAddress", required: true)]
        public Input<string> SenderAddress { get; set; } = null!;

        /// <summary>
        /// Sender name used to send emails.
        /// </summary>
        [Input("senderName", required: true)]
        public Input<string> SenderName { get; set; } = null!;

        /// <summary>
        /// TLS used to communicate with your SMTP server.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// User used to communicate with your SMTP server.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SmtpConfigArgs()
        {
        }
        public static new SmtpConfigArgs Empty => new SmtpConfigArgs();
    }

    public sealed class SmtpConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host and port address to your SMTP server.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Password used to communicate with your SMTP server.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Address used to send emails.
        /// </summary>
        [Input("senderAddress")]
        public Input<string>? SenderAddress { get; set; }

        /// <summary>
        /// Sender name used to send emails.
        /// </summary>
        [Input("senderName")]
        public Input<string>? SenderName { get; set; }

        /// <summary>
        /// TLS used to communicate with your SMTP server.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// User used to communicate with your SMTP server.
        /// </summary>
        [Input("user")]
        public Input<string>? User { get; set; }

        public SmtpConfigState()
        {
        }
        public static new SmtpConfigState Empty => new SmtpConfigState();
    }
}

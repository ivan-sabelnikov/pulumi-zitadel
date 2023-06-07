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
    /// Resource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
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
    ///     var machineUser = new Zitadel.MachineUser("machineUser", new()
    ///     {
    ///         OrgId = zitadel_org.Org.Id,
    ///         UserName = "machine@localhost.com",
    ///         Description = "description",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [ZitadelResourceType("zitadel:index/machineUser:MachineUser")]
    public partial class MachineUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        /// </summary>
        [Output("accessTokenType")]
        public Output<string?> AccessTokenType { get; private set; } = null!;

        /// <summary>
        /// Description of the user
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Loginnames
        /// </summary>
        [Output("loginNames")]
        public Output<ImmutableArray<string>> LoginNames { get; private set; } = null!;

        /// <summary>
        /// Name of the machine user
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Preferred login name
        /// </summary>
        [Output("preferredLoginName")]
        public Output<string> PreferredLoginName { get; private set; } = null!;

        /// <summary>
        /// State of the user
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Username
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a MachineUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MachineUser(string name, MachineUserArgs args, CustomResourceOptions? options = null)
            : base("zitadel:index/machineUser:MachineUser", name, args ?? new MachineUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MachineUser(string name, Input<string> id, MachineUserState? state = null, CustomResourceOptions? options = null)
            : base("zitadel:index/machineUser:MachineUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MachineUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MachineUser Get(string name, Input<string> id, MachineUserState? state = null, CustomResourceOptions? options = null)
        {
            return new MachineUser(name, id, state, options);
        }
    }

    public sealed class MachineUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        /// </summary>
        [Input("accessTokenType")]
        public Input<string>? AccessTokenType { get; set; }

        /// <summary>
        /// Description of the user
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the machine user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Username
        /// </summary>
        [Input("userName", required: true)]
        public Input<string> UserName { get; set; } = null!;

        public MachineUserArgs()
        {
        }
        public static new MachineUserArgs Empty => new MachineUserArgs();
    }

    public sealed class MachineUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access token type, supported values: ACCESS*TOKEN*TYPE*BEARER, ACCESS*TOKEN*TYPE*JWT
        /// </summary>
        [Input("accessTokenType")]
        public Input<string>? AccessTokenType { get; set; }

        /// <summary>
        /// Description of the user
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("loginNames")]
        private InputList<string>? _loginNames;

        /// <summary>
        /// Loginnames
        /// </summary>
        public InputList<string> LoginNames
        {
            get => _loginNames ?? (_loginNames = new InputList<string>());
            set => _loginNames = value;
        }

        /// <summary>
        /// Name of the machine user
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of the organization
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Preferred login name
        /// </summary>
        [Input("preferredLoginName")]
        public Input<string>? PreferredLoginName { get; set; }

        /// <summary>
        /// State of the user
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Username
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public MachineUserState()
        {
        }
        public static new MachineUserState Empty => new MachineUserState();
    }
}

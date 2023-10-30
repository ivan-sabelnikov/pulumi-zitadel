// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Resource representing a GitHub IdP on the organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := zitadel.NewOrgIdpGithub(ctx, "default", &zitadel.OrgIdpGithubArgs{
//				OrgId:        pulumi.Any(data.Zitadel_org.Default.Id),
//				ClientId:     pulumi.String("86a165..."),
//				ClientSecret: pulumi.String("*****afdbac18"),
//				Scopes: pulumi.StringArray{
//					pulumi.String("openid"),
//					pulumi.String("profile"),
//					pulumi.String("email"),
//				},
//				IsLinkingAllowed:  pulumi.Bool(false),
//				IsCreationAllowed: pulumi.Bool(true),
//				IsAutoCreation:    pulumi.Bool(false),
//				IsAutoUpdate:      pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// terraform The resource can be imported using the ID format `<id[:org_id][:client_secret]>`, e.g.
//
// ```sh
//
//	$ pulumi import zitadel:index/orgIdpGithub:OrgIdpGithub imported '123456789012345678:123456789012345678:1234567890123456781234567890123456787890'
//
// ```
type OrgIdpGithub struct {
	pulumi.CustomResourceState

	// client id generated by the identity provider
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolOutput `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolOutput `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolOutput `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolOutput `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of the organization
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
}

// NewOrgIdpGithub registers a new resource with the given unique name, arguments, and options.
func NewOrgIdpGithub(ctx *pulumi.Context,
	name string, args *OrgIdpGithubArgs, opts ...pulumi.ResourceOption) (*OrgIdpGithub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.ClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'ClientSecret'")
	}
	if args.IsAutoCreation == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoCreation'")
	}
	if args.IsAutoUpdate == nil {
		return nil, errors.New("invalid value for required argument 'IsAutoUpdate'")
	}
	if args.IsCreationAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsCreationAllowed'")
	}
	if args.IsLinkingAllowed == nil {
		return nil, errors.New("invalid value for required argument 'IsLinkingAllowed'")
	}
	if args.ClientSecret != nil {
		args.ClientSecret = pulumi.ToSecret(args.ClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgIdpGithub
	err := ctx.RegisterResource("zitadel:index/orgIdpGithub:OrgIdpGithub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgIdpGithub gets an existing OrgIdpGithub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgIdpGithub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgIdpGithubState, opts ...pulumi.ResourceOption) (*OrgIdpGithub, error) {
	var resource OrgIdpGithub
	err := ctx.ReadResource("zitadel:index/orgIdpGithub:OrgIdpGithub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgIdpGithub resources.
type orgIdpGithubState struct {
	// client id generated by the identity provider
	ClientId *string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret *string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation *bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate *bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed *bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed *bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

type OrgIdpGithubState struct {
	// client id generated by the identity provider
	ClientId pulumi.StringPtrInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringPtrInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolPtrInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolPtrInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolPtrInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolPtrInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGithubState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGithubState)(nil)).Elem()
}

type orgIdpGithubArgs struct {
	// client id generated by the identity provider
	ClientId string `pulumi:"clientId"`
	// client secret generated by the identity provider
	ClientSecret string `pulumi:"clientSecret"`
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// Name of the IDP
	Name *string `pulumi:"name"`
	// ID of the organization
	OrgId *string `pulumi:"orgId"`
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
}

// The set of arguments for constructing a OrgIdpGithub resource.
type OrgIdpGithubArgs struct {
	// client id generated by the identity provider
	ClientId pulumi.StringInput
	// client secret generated by the identity provider
	ClientSecret pulumi.StringInput
	// enable if a new account in ZITADEL should be created automatically on login with an external account
	IsAutoCreation pulumi.BoolInput
	// enable if a the ZITADEL account fields should be updated automatically on each login
	IsAutoUpdate pulumi.BoolInput
	// enable if users should be able to create a new account in ZITADEL when using an external account
	IsCreationAllowed pulumi.BoolInput
	// enable if users should be able to link an existing ZITADEL user with an external account
	IsLinkingAllowed pulumi.BoolInput
	// Name of the IDP
	Name pulumi.StringPtrInput
	// ID of the organization
	OrgId pulumi.StringPtrInput
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
}

func (OrgIdpGithubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgIdpGithubArgs)(nil)).Elem()
}

type OrgIdpGithubInput interface {
	pulumi.Input

	ToOrgIdpGithubOutput() OrgIdpGithubOutput
	ToOrgIdpGithubOutputWithContext(ctx context.Context) OrgIdpGithubOutput
}

func (*OrgIdpGithub) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGithub)(nil)).Elem()
}

func (i *OrgIdpGithub) ToOrgIdpGithubOutput() OrgIdpGithubOutput {
	return i.ToOrgIdpGithubOutputWithContext(context.Background())
}

func (i *OrgIdpGithub) ToOrgIdpGithubOutputWithContext(ctx context.Context) OrgIdpGithubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGithubOutput)
}

func (i *OrgIdpGithub) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpGithub] {
	return pulumix.Output[*OrgIdpGithub]{
		OutputState: i.ToOrgIdpGithubOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpGithubArrayInput is an input type that accepts OrgIdpGithubArray and OrgIdpGithubArrayOutput values.
// You can construct a concrete instance of `OrgIdpGithubArrayInput` via:
//
//	OrgIdpGithubArray{ OrgIdpGithubArgs{...} }
type OrgIdpGithubArrayInput interface {
	pulumi.Input

	ToOrgIdpGithubArrayOutput() OrgIdpGithubArrayOutput
	ToOrgIdpGithubArrayOutputWithContext(context.Context) OrgIdpGithubArrayOutput
}

type OrgIdpGithubArray []OrgIdpGithubInput

func (OrgIdpGithubArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGithub)(nil)).Elem()
}

func (i OrgIdpGithubArray) ToOrgIdpGithubArrayOutput() OrgIdpGithubArrayOutput {
	return i.ToOrgIdpGithubArrayOutputWithContext(context.Background())
}

func (i OrgIdpGithubArray) ToOrgIdpGithubArrayOutputWithContext(ctx context.Context) OrgIdpGithubArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGithubArrayOutput)
}

func (i OrgIdpGithubArray) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpGithub] {
	return pulumix.Output[[]*OrgIdpGithub]{
		OutputState: i.ToOrgIdpGithubArrayOutputWithContext(ctx).OutputState,
	}
}

// OrgIdpGithubMapInput is an input type that accepts OrgIdpGithubMap and OrgIdpGithubMapOutput values.
// You can construct a concrete instance of `OrgIdpGithubMapInput` via:
//
//	OrgIdpGithubMap{ "key": OrgIdpGithubArgs{...} }
type OrgIdpGithubMapInput interface {
	pulumi.Input

	ToOrgIdpGithubMapOutput() OrgIdpGithubMapOutput
	ToOrgIdpGithubMapOutputWithContext(context.Context) OrgIdpGithubMapOutput
}

type OrgIdpGithubMap map[string]OrgIdpGithubInput

func (OrgIdpGithubMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGithub)(nil)).Elem()
}

func (i OrgIdpGithubMap) ToOrgIdpGithubMapOutput() OrgIdpGithubMapOutput {
	return i.ToOrgIdpGithubMapOutputWithContext(context.Background())
}

func (i OrgIdpGithubMap) ToOrgIdpGithubMapOutputWithContext(ctx context.Context) OrgIdpGithubMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgIdpGithubMapOutput)
}

func (i OrgIdpGithubMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpGithub] {
	return pulumix.Output[map[string]*OrgIdpGithub]{
		OutputState: i.ToOrgIdpGithubMapOutputWithContext(ctx).OutputState,
	}
}

type OrgIdpGithubOutput struct{ *pulumi.OutputState }

func (OrgIdpGithubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgIdpGithub)(nil)).Elem()
}

func (o OrgIdpGithubOutput) ToOrgIdpGithubOutput() OrgIdpGithubOutput {
	return o
}

func (o OrgIdpGithubOutput) ToOrgIdpGithubOutputWithContext(ctx context.Context) OrgIdpGithubOutput {
	return o
}

func (o OrgIdpGithubOutput) ToOutput(ctx context.Context) pulumix.Output[*OrgIdpGithub] {
	return pulumix.Output[*OrgIdpGithub]{
		OutputState: o.OutputState,
	}
}

// client id generated by the identity provider
func (o OrgIdpGithubOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o OrgIdpGithubOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o OrgIdpGithubOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o OrgIdpGithubOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o OrgIdpGithubOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o OrgIdpGithubOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o OrgIdpGithubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o OrgIdpGithubOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o OrgIdpGithubOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OrgIdpGithub) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

type OrgIdpGithubArrayOutput struct{ *pulumi.OutputState }

func (OrgIdpGithubArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgIdpGithub)(nil)).Elem()
}

func (o OrgIdpGithubArrayOutput) ToOrgIdpGithubArrayOutput() OrgIdpGithubArrayOutput {
	return o
}

func (o OrgIdpGithubArrayOutput) ToOrgIdpGithubArrayOutputWithContext(ctx context.Context) OrgIdpGithubArrayOutput {
	return o
}

func (o OrgIdpGithubArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*OrgIdpGithub] {
	return pulumix.Output[[]*OrgIdpGithub]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpGithubArrayOutput) Index(i pulumi.IntInput) OrgIdpGithubOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgIdpGithub {
		return vs[0].([]*OrgIdpGithub)[vs[1].(int)]
	}).(OrgIdpGithubOutput)
}

type OrgIdpGithubMapOutput struct{ *pulumi.OutputState }

func (OrgIdpGithubMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgIdpGithub)(nil)).Elem()
}

func (o OrgIdpGithubMapOutput) ToOrgIdpGithubMapOutput() OrgIdpGithubMapOutput {
	return o
}

func (o OrgIdpGithubMapOutput) ToOrgIdpGithubMapOutputWithContext(ctx context.Context) OrgIdpGithubMapOutput {
	return o
}

func (o OrgIdpGithubMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*OrgIdpGithub] {
	return pulumix.Output[map[string]*OrgIdpGithub]{
		OutputState: o.OutputState,
	}
}

func (o OrgIdpGithubMapOutput) MapIndex(k pulumi.StringInput) OrgIdpGithubOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgIdpGithub {
		return vs[0].(map[string]*OrgIdpGithub)[vs[1].(string)]
	}).(OrgIdpGithubOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGithubInput)(nil)).Elem(), &OrgIdpGithub{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGithubArrayInput)(nil)).Elem(), OrgIdpGithubArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgIdpGithubMapInput)(nil)).Elem(), OrgIdpGithubMap{})
	pulumi.RegisterOutputType(OrgIdpGithubOutput{})
	pulumi.RegisterOutputType(OrgIdpGithubArrayOutput{})
	pulumi.RegisterOutputType(OrgIdpGithubMapOutput{})
}

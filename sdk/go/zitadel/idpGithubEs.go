// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing a GitHub Enterprise IDP on the instance.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/vavsab/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewIdpGithubEs(ctx, "githubEs", &zitadel.IdpGithubEsArgs{
// 			AuthorizationEndpoint: pulumi.String("https://auth.endpoint"),
// 			ClientId:              pulumi.String("86a165..."),
// 			ClientSecret:          pulumi.String("*****afdbac18"),
// 			IsAutoCreation:        pulumi.Bool(false),
// 			IsAutoUpdate:          pulumi.Bool(true),
// 			IsCreationAllowed:     pulumi.Bool(true),
// 			IsLinkingAllowed:      pulumi.Bool(false),
// 			Scopes: pulumi.StringArray{
// 				pulumi.String("openid"),
// 				pulumi.String("profile"),
// 				pulumi.String("email"),
// 			},
// 			TokenEndpoint: pulumi.String("https://token.endpoint"),
// 			UserEndpoint:  pulumi.String("https://user.endpoint"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type IdpGithubEs struct {
	pulumi.CustomResourceState

	// the providers authorization endpoint
	AuthorizationEndpoint pulumi.StringOutput `pulumi:"authorizationEndpoint"`
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// the providers token endpoint
	TokenEndpoint pulumi.StringOutput `pulumi:"tokenEndpoint"`
	// the providers user endpoint
	UserEndpoint pulumi.StringOutput `pulumi:"userEndpoint"`
}

// NewIdpGithubEs registers a new resource with the given unique name, arguments, and options.
func NewIdpGithubEs(ctx *pulumi.Context,
	name string, args *IdpGithubEsArgs, opts ...pulumi.ResourceOption) (*IdpGithubEs, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AuthorizationEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'AuthorizationEndpoint'")
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
	if args.TokenEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'TokenEndpoint'")
	}
	if args.UserEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'UserEndpoint'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IdpGithubEs
	err := ctx.RegisterResource("zitadel:index/idpGithubEs:IdpGithubEs", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdpGithubEs gets an existing IdpGithubEs resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdpGithubEs(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdpGithubEsState, opts ...pulumi.ResourceOption) (*IdpGithubEs, error) {
	var resource IdpGithubEs
	err := ctx.ReadResource("zitadel:index/idpGithubEs:IdpGithubEs", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdpGithubEs resources.
type idpGithubEsState struct {
	// the providers authorization endpoint
	AuthorizationEndpoint *string `pulumi:"authorizationEndpoint"`
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
	// the providers token endpoint
	TokenEndpoint *string `pulumi:"tokenEndpoint"`
	// the providers user endpoint
	UserEndpoint *string `pulumi:"userEndpoint"`
}

type IdpGithubEsState struct {
	// the providers authorization endpoint
	AuthorizationEndpoint pulumi.StringPtrInput
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
	// the providers token endpoint
	TokenEndpoint pulumi.StringPtrInput
	// the providers user endpoint
	UserEndpoint pulumi.StringPtrInput
}

func (IdpGithubEsState) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGithubEsState)(nil)).Elem()
}

type idpGithubEsArgs struct {
	// the providers authorization endpoint
	AuthorizationEndpoint string `pulumi:"authorizationEndpoint"`
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes []string `pulumi:"scopes"`
	// the providers token endpoint
	TokenEndpoint string `pulumi:"tokenEndpoint"`
	// the providers user endpoint
	UserEndpoint string `pulumi:"userEndpoint"`
}

// The set of arguments for constructing a IdpGithubEs resource.
type IdpGithubEsArgs struct {
	// the providers authorization endpoint
	AuthorizationEndpoint pulumi.StringInput
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
	// the scopes requested by ZITADEL during the request on the identity provider
	Scopes pulumi.StringArrayInput
	// the providers token endpoint
	TokenEndpoint pulumi.StringInput
	// the providers user endpoint
	UserEndpoint pulumi.StringInput
}

func (IdpGithubEsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*idpGithubEsArgs)(nil)).Elem()
}

type IdpGithubEsInput interface {
	pulumi.Input

	ToIdpGithubEsOutput() IdpGithubEsOutput
	ToIdpGithubEsOutputWithContext(ctx context.Context) IdpGithubEsOutput
}

func (*IdpGithubEs) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGithubEs)(nil)).Elem()
}

func (i *IdpGithubEs) ToIdpGithubEsOutput() IdpGithubEsOutput {
	return i.ToIdpGithubEsOutputWithContext(context.Background())
}

func (i *IdpGithubEs) ToIdpGithubEsOutputWithContext(ctx context.Context) IdpGithubEsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubEsOutput)
}

// IdpGithubEsArrayInput is an input type that accepts IdpGithubEsArray and IdpGithubEsArrayOutput values.
// You can construct a concrete instance of `IdpGithubEsArrayInput` via:
//
//          IdpGithubEsArray{ IdpGithubEsArgs{...} }
type IdpGithubEsArrayInput interface {
	pulumi.Input

	ToIdpGithubEsArrayOutput() IdpGithubEsArrayOutput
	ToIdpGithubEsArrayOutputWithContext(context.Context) IdpGithubEsArrayOutput
}

type IdpGithubEsArray []IdpGithubEsInput

func (IdpGithubEsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGithubEs)(nil)).Elem()
}

func (i IdpGithubEsArray) ToIdpGithubEsArrayOutput() IdpGithubEsArrayOutput {
	return i.ToIdpGithubEsArrayOutputWithContext(context.Background())
}

func (i IdpGithubEsArray) ToIdpGithubEsArrayOutputWithContext(ctx context.Context) IdpGithubEsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubEsArrayOutput)
}

// IdpGithubEsMapInput is an input type that accepts IdpGithubEsMap and IdpGithubEsMapOutput values.
// You can construct a concrete instance of `IdpGithubEsMapInput` via:
//
//          IdpGithubEsMap{ "key": IdpGithubEsArgs{...} }
type IdpGithubEsMapInput interface {
	pulumi.Input

	ToIdpGithubEsMapOutput() IdpGithubEsMapOutput
	ToIdpGithubEsMapOutputWithContext(context.Context) IdpGithubEsMapOutput
}

type IdpGithubEsMap map[string]IdpGithubEsInput

func (IdpGithubEsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGithubEs)(nil)).Elem()
}

func (i IdpGithubEsMap) ToIdpGithubEsMapOutput() IdpGithubEsMapOutput {
	return i.ToIdpGithubEsMapOutputWithContext(context.Background())
}

func (i IdpGithubEsMap) ToIdpGithubEsMapOutputWithContext(ctx context.Context) IdpGithubEsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdpGithubEsMapOutput)
}

type IdpGithubEsOutput struct{ *pulumi.OutputState }

func (IdpGithubEsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdpGithubEs)(nil)).Elem()
}

func (o IdpGithubEsOutput) ToIdpGithubEsOutput() IdpGithubEsOutput {
	return o
}

func (o IdpGithubEsOutput) ToIdpGithubEsOutputWithContext(ctx context.Context) IdpGithubEsOutput {
	return o
}

// the providers authorization endpoint
func (o IdpGithubEsOutput) AuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.AuthorizationEndpoint }).(pulumi.StringOutput)
}

// client id generated by the identity provider
func (o IdpGithubEsOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// client secret generated by the identity provider
func (o IdpGithubEsOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// enable if a new account in ZITADEL should be created automatically on login with an external account
func (o IdpGithubEsOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.BoolOutput { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enable if a the ZITADEL account fields should be updated automatically on each login
func (o IdpGithubEsOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.BoolOutput { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enable if users should be able to create a new account in ZITADEL when using an external account
func (o IdpGithubEsOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.BoolOutput { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enable if users should be able to link an existing ZITADEL user with an external account
func (o IdpGithubEsOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.BoolOutput { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// Name of the IDP
func (o IdpGithubEsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// the scopes requested by ZITADEL during the request on the identity provider
func (o IdpGithubEsOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

// the providers token endpoint
func (o IdpGithubEsOutput) TokenEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.TokenEndpoint }).(pulumi.StringOutput)
}

// the providers user endpoint
func (o IdpGithubEsOutput) UserEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *IdpGithubEs) pulumi.StringOutput { return v.UserEndpoint }).(pulumi.StringOutput)
}

type IdpGithubEsArrayOutput struct{ *pulumi.OutputState }

func (IdpGithubEsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdpGithubEs)(nil)).Elem()
}

func (o IdpGithubEsArrayOutput) ToIdpGithubEsArrayOutput() IdpGithubEsArrayOutput {
	return o
}

func (o IdpGithubEsArrayOutput) ToIdpGithubEsArrayOutputWithContext(ctx context.Context) IdpGithubEsArrayOutput {
	return o
}

func (o IdpGithubEsArrayOutput) Index(i pulumi.IntInput) IdpGithubEsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdpGithubEs {
		return vs[0].([]*IdpGithubEs)[vs[1].(int)]
	}).(IdpGithubEsOutput)
}

type IdpGithubEsMapOutput struct{ *pulumi.OutputState }

func (IdpGithubEsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdpGithubEs)(nil)).Elem()
}

func (o IdpGithubEsMapOutput) ToIdpGithubEsMapOutput() IdpGithubEsMapOutput {
	return o
}

func (o IdpGithubEsMapOutput) ToIdpGithubEsMapOutputWithContext(ctx context.Context) IdpGithubEsMapOutput {
	return o
}

func (o IdpGithubEsMapOutput) MapIndex(k pulumi.StringInput) IdpGithubEsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdpGithubEs {
		return vs[0].(map[string]*IdpGithubEs)[vs[1].(string)]
	}).(IdpGithubEsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubEsInput)(nil)).Elem(), &IdpGithubEs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubEsArrayInput)(nil)).Elem(), IdpGithubEsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdpGithubEsMapInput)(nil)).Elem(), IdpGithubEsMap{})
	pulumi.RegisterOutputType(IdpGithubEsOutput{})
	pulumi.RegisterOutputType(IdpGithubEsArrayOutput{})
	pulumi.RegisterOutputType(IdpGithubEsMapOutput{})
}

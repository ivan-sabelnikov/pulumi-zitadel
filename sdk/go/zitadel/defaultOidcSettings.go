// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource representing the default oidc settings.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := zitadel.NewDefaultOidcSettings(ctx, "default", &zitadel.DefaultOidcSettingsArgs{
// 			AccessTokenLifetime:        pulumi.String("12h0m0s"),
// 			IdTokenLifetime:            pulumi.String("12h0m0s"),
// 			RefreshTokenExpiration:     pulumi.String("720h0m0s"),
// 			RefreshTokenIdleExpiration: pulumi.String("2160h0m0s"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DefaultOidcSettings struct {
	pulumi.CustomResourceState

	// lifetime duration of access tokens
	AccessTokenLifetime pulumi.StringOutput `pulumi:"accessTokenLifetime"`
	// lifetime duration of id tokens
	IdTokenLifetime pulumi.StringOutput `pulumi:"idTokenLifetime"`
	// expiration duration of refresh tokens
	RefreshTokenExpiration pulumi.StringOutput `pulumi:"refreshTokenExpiration"`
	// expiration duration of idle refresh tokens
	RefreshTokenIdleExpiration pulumi.StringOutput `pulumi:"refreshTokenIdleExpiration"`
}

// NewDefaultOidcSettings registers a new resource with the given unique name, arguments, and options.
func NewDefaultOidcSettings(ctx *pulumi.Context,
	name string, args *DefaultOidcSettingsArgs, opts ...pulumi.ResourceOption) (*DefaultOidcSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessTokenLifetime == nil {
		return nil, errors.New("invalid value for required argument 'AccessTokenLifetime'")
	}
	if args.IdTokenLifetime == nil {
		return nil, errors.New("invalid value for required argument 'IdTokenLifetime'")
	}
	if args.RefreshTokenExpiration == nil {
		return nil, errors.New("invalid value for required argument 'RefreshTokenExpiration'")
	}
	if args.RefreshTokenIdleExpiration == nil {
		return nil, errors.New("invalid value for required argument 'RefreshTokenIdleExpiration'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DefaultOidcSettings
	err := ctx.RegisterResource("zitadel:index/defaultOidcSettings:DefaultOidcSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultOidcSettings gets an existing DefaultOidcSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultOidcSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultOidcSettingsState, opts ...pulumi.ResourceOption) (*DefaultOidcSettings, error) {
	var resource DefaultOidcSettings
	err := ctx.ReadResource("zitadel:index/defaultOidcSettings:DefaultOidcSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultOidcSettings resources.
type defaultOidcSettingsState struct {
	// lifetime duration of access tokens
	AccessTokenLifetime *string `pulumi:"accessTokenLifetime"`
	// lifetime duration of id tokens
	IdTokenLifetime *string `pulumi:"idTokenLifetime"`
	// expiration duration of refresh tokens
	RefreshTokenExpiration *string `pulumi:"refreshTokenExpiration"`
	// expiration duration of idle refresh tokens
	RefreshTokenIdleExpiration *string `pulumi:"refreshTokenIdleExpiration"`
}

type DefaultOidcSettingsState struct {
	// lifetime duration of access tokens
	AccessTokenLifetime pulumi.StringPtrInput
	// lifetime duration of id tokens
	IdTokenLifetime pulumi.StringPtrInput
	// expiration duration of refresh tokens
	RefreshTokenExpiration pulumi.StringPtrInput
	// expiration duration of idle refresh tokens
	RefreshTokenIdleExpiration pulumi.StringPtrInput
}

func (DefaultOidcSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultOidcSettingsState)(nil)).Elem()
}

type defaultOidcSettingsArgs struct {
	// lifetime duration of access tokens
	AccessTokenLifetime string `pulumi:"accessTokenLifetime"`
	// lifetime duration of id tokens
	IdTokenLifetime string `pulumi:"idTokenLifetime"`
	// expiration duration of refresh tokens
	RefreshTokenExpiration string `pulumi:"refreshTokenExpiration"`
	// expiration duration of idle refresh tokens
	RefreshTokenIdleExpiration string `pulumi:"refreshTokenIdleExpiration"`
}

// The set of arguments for constructing a DefaultOidcSettings resource.
type DefaultOidcSettingsArgs struct {
	// lifetime duration of access tokens
	AccessTokenLifetime pulumi.StringInput
	// lifetime duration of id tokens
	IdTokenLifetime pulumi.StringInput
	// expiration duration of refresh tokens
	RefreshTokenExpiration pulumi.StringInput
	// expiration duration of idle refresh tokens
	RefreshTokenIdleExpiration pulumi.StringInput
}

func (DefaultOidcSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultOidcSettingsArgs)(nil)).Elem()
}

type DefaultOidcSettingsInput interface {
	pulumi.Input

	ToDefaultOidcSettingsOutput() DefaultOidcSettingsOutput
	ToDefaultOidcSettingsOutputWithContext(ctx context.Context) DefaultOidcSettingsOutput
}

func (*DefaultOidcSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultOidcSettings)(nil)).Elem()
}

func (i *DefaultOidcSettings) ToDefaultOidcSettingsOutput() DefaultOidcSettingsOutput {
	return i.ToDefaultOidcSettingsOutputWithContext(context.Background())
}

func (i *DefaultOidcSettings) ToDefaultOidcSettingsOutputWithContext(ctx context.Context) DefaultOidcSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultOidcSettingsOutput)
}

// DefaultOidcSettingsArrayInput is an input type that accepts DefaultOidcSettingsArray and DefaultOidcSettingsArrayOutput values.
// You can construct a concrete instance of `DefaultOidcSettingsArrayInput` via:
//
//          DefaultOidcSettingsArray{ DefaultOidcSettingsArgs{...} }
type DefaultOidcSettingsArrayInput interface {
	pulumi.Input

	ToDefaultOidcSettingsArrayOutput() DefaultOidcSettingsArrayOutput
	ToDefaultOidcSettingsArrayOutputWithContext(context.Context) DefaultOidcSettingsArrayOutput
}

type DefaultOidcSettingsArray []DefaultOidcSettingsInput

func (DefaultOidcSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultOidcSettings)(nil)).Elem()
}

func (i DefaultOidcSettingsArray) ToDefaultOidcSettingsArrayOutput() DefaultOidcSettingsArrayOutput {
	return i.ToDefaultOidcSettingsArrayOutputWithContext(context.Background())
}

func (i DefaultOidcSettingsArray) ToDefaultOidcSettingsArrayOutputWithContext(ctx context.Context) DefaultOidcSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultOidcSettingsArrayOutput)
}

// DefaultOidcSettingsMapInput is an input type that accepts DefaultOidcSettingsMap and DefaultOidcSettingsMapOutput values.
// You can construct a concrete instance of `DefaultOidcSettingsMapInput` via:
//
//          DefaultOidcSettingsMap{ "key": DefaultOidcSettingsArgs{...} }
type DefaultOidcSettingsMapInput interface {
	pulumi.Input

	ToDefaultOidcSettingsMapOutput() DefaultOidcSettingsMapOutput
	ToDefaultOidcSettingsMapOutputWithContext(context.Context) DefaultOidcSettingsMapOutput
}

type DefaultOidcSettingsMap map[string]DefaultOidcSettingsInput

func (DefaultOidcSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultOidcSettings)(nil)).Elem()
}

func (i DefaultOidcSettingsMap) ToDefaultOidcSettingsMapOutput() DefaultOidcSettingsMapOutput {
	return i.ToDefaultOidcSettingsMapOutputWithContext(context.Background())
}

func (i DefaultOidcSettingsMap) ToDefaultOidcSettingsMapOutputWithContext(ctx context.Context) DefaultOidcSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultOidcSettingsMapOutput)
}

type DefaultOidcSettingsOutput struct{ *pulumi.OutputState }

func (DefaultOidcSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultOidcSettings)(nil)).Elem()
}

func (o DefaultOidcSettingsOutput) ToDefaultOidcSettingsOutput() DefaultOidcSettingsOutput {
	return o
}

func (o DefaultOidcSettingsOutput) ToDefaultOidcSettingsOutputWithContext(ctx context.Context) DefaultOidcSettingsOutput {
	return o
}

// lifetime duration of access tokens
func (o DefaultOidcSettingsOutput) AccessTokenLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultOidcSettings) pulumi.StringOutput { return v.AccessTokenLifetime }).(pulumi.StringOutput)
}

// lifetime duration of id tokens
func (o DefaultOidcSettingsOutput) IdTokenLifetime() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultOidcSettings) pulumi.StringOutput { return v.IdTokenLifetime }).(pulumi.StringOutput)
}

// expiration duration of refresh tokens
func (o DefaultOidcSettingsOutput) RefreshTokenExpiration() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultOidcSettings) pulumi.StringOutput { return v.RefreshTokenExpiration }).(pulumi.StringOutput)
}

// expiration duration of idle refresh tokens
func (o DefaultOidcSettingsOutput) RefreshTokenIdleExpiration() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultOidcSettings) pulumi.StringOutput { return v.RefreshTokenIdleExpiration }).(pulumi.StringOutput)
}

type DefaultOidcSettingsArrayOutput struct{ *pulumi.OutputState }

func (DefaultOidcSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultOidcSettings)(nil)).Elem()
}

func (o DefaultOidcSettingsArrayOutput) ToDefaultOidcSettingsArrayOutput() DefaultOidcSettingsArrayOutput {
	return o
}

func (o DefaultOidcSettingsArrayOutput) ToDefaultOidcSettingsArrayOutputWithContext(ctx context.Context) DefaultOidcSettingsArrayOutput {
	return o
}

func (o DefaultOidcSettingsArrayOutput) Index(i pulumi.IntInput) DefaultOidcSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultOidcSettings {
		return vs[0].([]*DefaultOidcSettings)[vs[1].(int)]
	}).(DefaultOidcSettingsOutput)
}

type DefaultOidcSettingsMapOutput struct{ *pulumi.OutputState }

func (DefaultOidcSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultOidcSettings)(nil)).Elem()
}

func (o DefaultOidcSettingsMapOutput) ToDefaultOidcSettingsMapOutput() DefaultOidcSettingsMapOutput {
	return o
}

func (o DefaultOidcSettingsMapOutput) ToDefaultOidcSettingsMapOutputWithContext(ctx context.Context) DefaultOidcSettingsMapOutput {
	return o
}

func (o DefaultOidcSettingsMapOutput) MapIndex(k pulumi.StringInput) DefaultOidcSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultOidcSettings {
		return vs[0].(map[string]*DefaultOidcSettings)[vs[1].(string)]
	}).(DefaultOidcSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultOidcSettingsInput)(nil)).Elem(), &DefaultOidcSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultOidcSettingsArrayInput)(nil)).Elem(), DefaultOidcSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultOidcSettingsMapInput)(nil)).Elem(), DefaultOidcSettingsMap{})
	pulumi.RegisterOutputType(DefaultOidcSettingsOutput{})
	pulumi.RegisterOutputType(DefaultOidcSettingsArrayOutput{})
	pulumi.RegisterOutputType(DefaultOidcSettingsMapOutput{})
}

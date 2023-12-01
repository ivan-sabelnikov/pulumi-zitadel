// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel/internal"
)

// Datasource representing an LDAP IDP on the instance.
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
//			_, err := zitadel.LookupIdpLdap(ctx, &zitadel.LookupIdpLdapArgs{
//				Id: "123456789012345678",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIdpLdap(ctx *pulumi.Context, args *LookupIdpLdapArgs, opts ...pulumi.InvokeOption) (*LookupIdpLdapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIdpLdapResult
	err := ctx.Invoke("zitadel:index/getIdpLdap:getIdpLdap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIdpLdap.
type LookupIdpLdapArgs struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
}

// A collection of values returned by getIdpLdap.
type LookupIdpLdapResult struct {
	// User attribute for the avatar url
	AvatarUrlAttribute string `pulumi:"avatarUrlAttribute"`
	// Base DN for LDAP connections
	BaseDn string `pulumi:"baseDn"`
	// Bind DN for LDAP connections
	BindDn string `pulumi:"bindDn"`
	// Bind password for LDAP connections
	BindPassword string `pulumi:"bindPassword"`
	// User attribute for the display name
	DisplayNameAttribute string `pulumi:"displayNameAttribute"`
	// User attribute for the email
	EmailAttribute string `pulumi:"emailAttribute"`
	// User attribute for the email verified state
	EmailVerifiedAttribute string `pulumi:"emailVerifiedAttribute"`
	// User attribute for the first name
	FirstNameAttribute string `pulumi:"firstNameAttribute"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// User attribute for the id
	IdAttribute string `pulumi:"idAttribute"`
	// enabled if a new account in ZITADEL are created automatically on login with an external account
	IsAutoCreation bool `pulumi:"isAutoCreation"`
	// enabled if a the ZITADEL account fields are updated automatically on each login
	IsAutoUpdate bool `pulumi:"isAutoUpdate"`
	// enabled if users are able to create a new account in ZITADEL when using an external account
	IsCreationAllowed bool `pulumi:"isCreationAllowed"`
	// enabled if users are able to link an existing ZITADEL user with an external account
	IsLinkingAllowed bool `pulumi:"isLinkingAllowed"`
	// User attribute for the last name
	LastNameAttribute string `pulumi:"lastNameAttribute"`
	// Name of the IDP
	Name string `pulumi:"name"`
	// User attribute for the nick name
	NickNameAttribute string `pulumi:"nickNameAttribute"`
	// User attribute for the phone
	PhoneAttribute string `pulumi:"phoneAttribute"`
	// User attribute for the phone verified state
	PhoneVerifiedAttribute string `pulumi:"phoneVerifiedAttribute"`
	// User attribute for the preferred language
	PreferredLanguageAttribute string `pulumi:"preferredLanguageAttribute"`
	// User attribute for the preferred username
	PreferredUsernameAttribute string `pulumi:"preferredUsernameAttribute"`
	// User attribute for the profile
	ProfileAttribute string `pulumi:"profileAttribute"`
	// Servers to try in order for establishing LDAP connections
	Servers []string `pulumi:"servers"`
	// Wether to use StartTLS for LDAP connections
	StartTls bool `pulumi:"startTls"`
	// Timeout for LDAP connections
	Timeout string `pulumi:"timeout"`
	// User base for LDAP connections
	UserBase string `pulumi:"userBase"`
	// User filters for LDAP connections
	UserFilters []string `pulumi:"userFilters"`
	// User object classes for LDAP connections
	UserObjectClasses []string `pulumi:"userObjectClasses"`
}

func LookupIdpLdapOutput(ctx *pulumi.Context, args LookupIdpLdapOutputArgs, opts ...pulumi.InvokeOption) LookupIdpLdapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIdpLdapResult, error) {
			args := v.(LookupIdpLdapArgs)
			r, err := LookupIdpLdap(ctx, &args, opts...)
			var s LookupIdpLdapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIdpLdapResultOutput)
}

// A collection of arguments for invoking getIdpLdap.
type LookupIdpLdapOutputArgs struct {
	// The ID of this resource.
	Id pulumi.StringInput `pulumi:"id"`
}

func (LookupIdpLdapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpLdapArgs)(nil)).Elem()
}

// A collection of values returned by getIdpLdap.
type LookupIdpLdapResultOutput struct{ *pulumi.OutputState }

func (LookupIdpLdapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIdpLdapResult)(nil)).Elem()
}

func (o LookupIdpLdapResultOutput) ToLookupIdpLdapResultOutput() LookupIdpLdapResultOutput {
	return o
}

func (o LookupIdpLdapResultOutput) ToLookupIdpLdapResultOutputWithContext(ctx context.Context) LookupIdpLdapResultOutput {
	return o
}

func (o LookupIdpLdapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIdpLdapResult] {
	return pulumix.Output[LookupIdpLdapResult]{
		OutputState: o.OutputState,
	}
}

// User attribute for the avatar url
func (o LookupIdpLdapResultOutput) AvatarUrlAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.AvatarUrlAttribute }).(pulumi.StringOutput)
}

// Base DN for LDAP connections
func (o LookupIdpLdapResultOutput) BaseDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.BaseDn }).(pulumi.StringOutput)
}

// Bind DN for LDAP connections
func (o LookupIdpLdapResultOutput) BindDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.BindDn }).(pulumi.StringOutput)
}

// Bind password for LDAP connections
func (o LookupIdpLdapResultOutput) BindPassword() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.BindPassword }).(pulumi.StringOutput)
}

// User attribute for the display name
func (o LookupIdpLdapResultOutput) DisplayNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.DisplayNameAttribute }).(pulumi.StringOutput)
}

// User attribute for the email
func (o LookupIdpLdapResultOutput) EmailAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.EmailAttribute }).(pulumi.StringOutput)
}

// User attribute for the email verified state
func (o LookupIdpLdapResultOutput) EmailVerifiedAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.EmailVerifiedAttribute }).(pulumi.StringOutput)
}

// User attribute for the first name
func (o LookupIdpLdapResultOutput) FirstNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.FirstNameAttribute }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupIdpLdapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.Id }).(pulumi.StringOutput)
}

// User attribute for the id
func (o LookupIdpLdapResultOutput) IdAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.IdAttribute }).(pulumi.StringOutput)
}

// enabled if a new account in ZITADEL are created automatically on login with an external account
func (o LookupIdpLdapResultOutput) IsAutoCreation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) bool { return v.IsAutoCreation }).(pulumi.BoolOutput)
}

// enabled if a the ZITADEL account fields are updated automatically on each login
func (o LookupIdpLdapResultOutput) IsAutoUpdate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) bool { return v.IsAutoUpdate }).(pulumi.BoolOutput)
}

// enabled if users are able to create a new account in ZITADEL when using an external account
func (o LookupIdpLdapResultOutput) IsCreationAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) bool { return v.IsCreationAllowed }).(pulumi.BoolOutput)
}

// enabled if users are able to link an existing ZITADEL user with an external account
func (o LookupIdpLdapResultOutput) IsLinkingAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) bool { return v.IsLinkingAllowed }).(pulumi.BoolOutput)
}

// User attribute for the last name
func (o LookupIdpLdapResultOutput) LastNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.LastNameAttribute }).(pulumi.StringOutput)
}

// Name of the IDP
func (o LookupIdpLdapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.Name }).(pulumi.StringOutput)
}

// User attribute for the nick name
func (o LookupIdpLdapResultOutput) NickNameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.NickNameAttribute }).(pulumi.StringOutput)
}

// User attribute for the phone
func (o LookupIdpLdapResultOutput) PhoneAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.PhoneAttribute }).(pulumi.StringOutput)
}

// User attribute for the phone verified state
func (o LookupIdpLdapResultOutput) PhoneVerifiedAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.PhoneVerifiedAttribute }).(pulumi.StringOutput)
}

// User attribute for the preferred language
func (o LookupIdpLdapResultOutput) PreferredLanguageAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.PreferredLanguageAttribute }).(pulumi.StringOutput)
}

// User attribute for the preferred username
func (o LookupIdpLdapResultOutput) PreferredUsernameAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.PreferredUsernameAttribute }).(pulumi.StringOutput)
}

// User attribute for the profile
func (o LookupIdpLdapResultOutput) ProfileAttribute() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.ProfileAttribute }).(pulumi.StringOutput)
}

// Servers to try in order for establishing LDAP connections
func (o LookupIdpLdapResultOutput) Servers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) []string { return v.Servers }).(pulumi.StringArrayOutput)
}

// Wether to use StartTLS for LDAP connections
func (o LookupIdpLdapResultOutput) StartTls() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) bool { return v.StartTls }).(pulumi.BoolOutput)
}

// Timeout for LDAP connections
func (o LookupIdpLdapResultOutput) Timeout() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.Timeout }).(pulumi.StringOutput)
}

// User base for LDAP connections
func (o LookupIdpLdapResultOutput) UserBase() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) string { return v.UserBase }).(pulumi.StringOutput)
}

// User filters for LDAP connections
func (o LookupIdpLdapResultOutput) UserFilters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) []string { return v.UserFilters }).(pulumi.StringArrayOutput)
}

// User object classes for LDAP connections
func (o LookupIdpLdapResultOutput) UserObjectClasses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupIdpLdapResult) []string { return v.UserObjectClasses }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIdpLdapResultOutput{})
}

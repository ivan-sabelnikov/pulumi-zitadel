// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a human user situated under an organization, which then can be authorized through memberships or direct grants on other resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-zitadel/sdk/go/zitadel"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// 	"github.com/pulumiverse/pulumi-zitadel/sdk/go/zitadel"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		humanUserHumanUser, err := zitadel.LookupHumanUser(ctx, &GetHumanUserArgs{
// 			OrgId:  data.Zitadel_org.Org.Id,
// 			UserId: "177073614158299139",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("humanUser", humanUserHumanUser)
// 		return nil
// 	})
// }
// ```
func LookupHumanUser(ctx *pulumi.Context, args *LookupHumanUserArgs, opts ...pulumi.InvokeOption) (*LookupHumanUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupHumanUserResult
	err := ctx.Invoke("zitadel:index/getHumanUser:getHumanUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHumanUser.
type LookupHumanUserArgs struct {
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// The ID of this resource.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getHumanUser.
type LookupHumanUserResult struct {
	// Display name of the user
	DisplayName string `pulumi:"displayName"`
	// Email of the user
	Email string `pulumi:"email"`
	// First name of the user
	FirstName string `pulumi:"firstName"`
	// Gender of the user
	Gender string `pulumi:"gender"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Is the email verified of the user, can only be true if password of the user is set
	IsEmailVerified bool `pulumi:"isEmailVerified"`
	// Is the phone verified of the user
	IsPhoneVerified bool `pulumi:"isPhoneVerified"`
	// Last name of the user
	LastName string `pulumi:"lastName"`
	// Loginnames
	LoginNames []string `pulumi:"loginNames"`
	// Nick name of the user
	NickName string `pulumi:"nickName"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// Phone of the user
	Phone string `pulumi:"phone"`
	// Preferred language of the user
	PreferredLanguage string `pulumi:"preferredLanguage"`
	// Preferred login name
	PreferredLoginName string `pulumi:"preferredLoginName"`
	// State of the user
	State string `pulumi:"state"`
	// The ID of this resource.
	UserId string `pulumi:"userId"`
	// Username
	UserName string `pulumi:"userName"`
}

func LookupHumanUserOutput(ctx *pulumi.Context, args LookupHumanUserOutputArgs, opts ...pulumi.InvokeOption) LookupHumanUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupHumanUserResult, error) {
			args := v.(LookupHumanUserArgs)
			r, err := LookupHumanUser(ctx, &args, opts...)
			var s LookupHumanUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupHumanUserResultOutput)
}

// A collection of arguments for invoking getHumanUser.
type LookupHumanUserOutputArgs struct {
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
	// The ID of this resource.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupHumanUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHumanUserArgs)(nil)).Elem()
}

// A collection of values returned by getHumanUser.
type LookupHumanUserResultOutput struct{ *pulumi.OutputState }

func (LookupHumanUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHumanUserResult)(nil)).Elem()
}

func (o LookupHumanUserResultOutput) ToLookupHumanUserResultOutput() LookupHumanUserResultOutput {
	return o
}

func (o LookupHumanUserResultOutput) ToLookupHumanUserResultOutputWithContext(ctx context.Context) LookupHumanUserResultOutput {
	return o
}

// Display name of the user
func (o LookupHumanUserResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Email of the user
func (o LookupHumanUserResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.Email }).(pulumi.StringOutput)
}

// First name of the user
func (o LookupHumanUserResultOutput) FirstName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.FirstName }).(pulumi.StringOutput)
}

// Gender of the user
func (o LookupHumanUserResultOutput) Gender() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.Gender }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupHumanUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Is the email verified of the user, can only be true if password of the user is set
func (o LookupHumanUserResultOutput) IsEmailVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHumanUserResult) bool { return v.IsEmailVerified }).(pulumi.BoolOutput)
}

// Is the phone verified of the user
func (o LookupHumanUserResultOutput) IsPhoneVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupHumanUserResult) bool { return v.IsPhoneVerified }).(pulumi.BoolOutput)
}

// Last name of the user
func (o LookupHumanUserResultOutput) LastName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.LastName }).(pulumi.StringOutput)
}

// Loginnames
func (o LookupHumanUserResultOutput) LoginNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupHumanUserResult) []string { return v.LoginNames }).(pulumi.StringArrayOutput)
}

// Nick name of the user
func (o LookupHumanUserResultOutput) NickName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.NickName }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupHumanUserResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Phone of the user
func (o LookupHumanUserResultOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.Phone }).(pulumi.StringOutput)
}

// Preferred language of the user
func (o LookupHumanUserResultOutput) PreferredLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.PreferredLanguage }).(pulumi.StringOutput)
}

// Preferred login name
func (o LookupHumanUserResultOutput) PreferredLoginName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.PreferredLoginName }).(pulumi.StringOutput)
}

// State of the user
func (o LookupHumanUserResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.State }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupHumanUserResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Username
func (o LookupHumanUserResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHumanUserResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHumanUserResultOutput{})
}

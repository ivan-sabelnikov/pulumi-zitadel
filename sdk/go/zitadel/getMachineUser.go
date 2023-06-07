// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package zitadel

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource representing a serviceaccount situated under an organization, which then can be authorized through memberships or direct grants on other resources.
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
// 		machineUserMachineUser, err := zitadel.LookupMachineUser(ctx, &GetMachineUserArgs{
// 			OrgId:  data.Zitadel_org.Org.Id,
// 			UserId: "177073617463410691",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("machineUser", machineUserMachineUser)
// 		return nil
// 	})
// }
// ```
func LookupMachineUser(ctx *pulumi.Context, args *LookupMachineUserArgs, opts ...pulumi.InvokeOption) (*LookupMachineUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupMachineUserResult
	err := ctx.Invoke("zitadel:index/getMachineUser:getMachineUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMachineUser.
type LookupMachineUserArgs struct {
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// The ID of this resource.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getMachineUser.
type LookupMachineUserResult struct {
	// Access token type
	AccessTokenType string `pulumi:"accessTokenType"`
	// Description of the user
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Loginnames
	LoginNames []string `pulumi:"loginNames"`
	// Name of the machine user
	Name string `pulumi:"name"`
	// ID of the organization
	OrgId string `pulumi:"orgId"`
	// Preferred login name
	PreferredLoginName string `pulumi:"preferredLoginName"`
	// State of the user
	State string `pulumi:"state"`
	// The ID of this resource.
	UserId string `pulumi:"userId"`
	// Username
	UserName string `pulumi:"userName"`
}

func LookupMachineUserOutput(ctx *pulumi.Context, args LookupMachineUserOutputArgs, opts ...pulumi.InvokeOption) LookupMachineUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMachineUserResult, error) {
			args := v.(LookupMachineUserArgs)
			r, err := LookupMachineUser(ctx, &args, opts...)
			var s LookupMachineUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMachineUserResultOutput)
}

// A collection of arguments for invoking getMachineUser.
type LookupMachineUserOutputArgs struct {
	// ID of the organization
	OrgId pulumi.StringInput `pulumi:"orgId"`
	// The ID of this resource.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupMachineUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineUserArgs)(nil)).Elem()
}

// A collection of values returned by getMachineUser.
type LookupMachineUserResultOutput struct{ *pulumi.OutputState }

func (LookupMachineUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMachineUserResult)(nil)).Elem()
}

func (o LookupMachineUserResultOutput) ToLookupMachineUserResultOutput() LookupMachineUserResultOutput {
	return o
}

func (o LookupMachineUserResultOutput) ToLookupMachineUserResultOutputWithContext(ctx context.Context) LookupMachineUserResultOutput {
	return o
}

// Access token type
func (o LookupMachineUserResultOutput) AccessTokenType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.AccessTokenType }).(pulumi.StringOutput)
}

// Description of the user
func (o LookupMachineUserResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupMachineUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Loginnames
func (o LookupMachineUserResultOutput) LoginNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMachineUserResult) []string { return v.LoginNames }).(pulumi.StringArrayOutput)
}

// Name of the machine user
func (o LookupMachineUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// ID of the organization
func (o LookupMachineUserResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Preferred login name
func (o LookupMachineUserResultOutput) PreferredLoginName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.PreferredLoginName }).(pulumi.StringOutput)
}

// State of the user
func (o LookupMachineUserResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.State }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o LookupMachineUserResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.UserId }).(pulumi.StringOutput)
}

// Username
func (o LookupMachineUserResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMachineUserResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMachineUserResultOutput{})
}

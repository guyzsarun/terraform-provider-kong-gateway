// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Authentication struct {
	Mechanism types.String `tfsdk:"mechanism"`
	Password  types.String `tfsdk:"password"`
	Strategy  types.String `tfsdk:"strategy"`
	Tokenauth types.Bool   `tfsdk:"tokenauth"`
	User      types.String `tfsdk:"user"`
}

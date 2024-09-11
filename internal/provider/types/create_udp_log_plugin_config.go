// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateUDPLogPluginConfig struct {
	CustomFieldsByLua map[string]types.String `tfsdk:"custom_fields_by_lua"`
	Host              types.String            `tfsdk:"host"`
	Port              types.Int64             `tfsdk:"port"`
	Timeout           types.Number            `tfsdk:"timeout"`
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateRequestTransformerAdvancedPluginConfig struct {
	Add        *CreateRequestTransformerAdvancedPluginAdd     `tfsdk:"add"`
	Allow      *Allow                                         `tfsdk:"allow"`
	Append     *CreateRequestTransformerAdvancedPluginAdd     `tfsdk:"append"`
	DotsInKeys types.Bool                                     `tfsdk:"dots_in_keys"`
	HTTPMethod types.String                                   `tfsdk:"http_method"`
	Remove     *Add                                           `tfsdk:"remove"`
	Rename     *Add                                           `tfsdk:"rename"`
	Replace    *CreateRequestTransformerAdvancedPluginReplace `tfsdk:"replace"`
}

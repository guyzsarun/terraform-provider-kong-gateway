// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateResponseRatelimitingPluginConfig struct {
	BlockOnFirstViolation types.Bool                     `tfsdk:"block_on_first_violation"`
	FaultTolerant         types.Bool                     `tfsdk:"fault_tolerant"`
	HeaderName            types.String                   `tfsdk:"header_name"`
	HideClientHeaders     types.Bool                     `tfsdk:"hide_client_headers"`
	LimitBy               types.String                   `tfsdk:"limit_by"`
	Limits                map[string]types.String        `tfsdk:"limits"`
	Policy                types.String                   `tfsdk:"policy"`
	Redis                 *CreateRateLimitingPluginRedis `tfsdk:"redis"`
}

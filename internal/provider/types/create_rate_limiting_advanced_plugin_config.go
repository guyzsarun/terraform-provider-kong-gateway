// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateRateLimitingAdvancedPluginConfig struct {
	ConsumerGroups        []types.String `tfsdk:"consumer_groups"`
	DictionaryName        types.String   `tfsdk:"dictionary_name"`
	DisablePenalty        types.Bool     `tfsdk:"disable_penalty"`
	EnforceConsumerGroups types.Bool     `tfsdk:"enforce_consumer_groups"`
	ErrorCode             types.Number   `tfsdk:"error_code"`
	ErrorMessage          types.String   `tfsdk:"error_message"`
	HeaderName            types.String   `tfsdk:"header_name"`
	HideClientHeaders     types.Bool     `tfsdk:"hide_client_headers"`
	Identifier            types.String   `tfsdk:"identifier"`
	Limit                 []types.Number `tfsdk:"limit"`
	Namespace             types.String   `tfsdk:"namespace"`
	Path                  types.String   `tfsdk:"path"`
	Redis                 *Redis         `tfsdk:"redis"`
	RetryAfterJitterMax   types.Number   `tfsdk:"retry_after_jitter_max"`
	Strategy              types.String   `tfsdk:"strategy"`
	SyncRate              types.Number   `tfsdk:"sync_rate"`
	WindowSize            []types.Number `tfsdk:"window_size"`
	WindowType            types.String   `tfsdk:"window_type"`
}

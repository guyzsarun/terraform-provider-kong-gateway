// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateOauth2IntrospectionPluginConfig struct {
	Anonymous                  types.String            `tfsdk:"anonymous"`
	AuthorizationValue         types.String            `tfsdk:"authorization_value"`
	ConsumerBy                 types.String            `tfsdk:"consumer_by"`
	CustomClaimsForward        []types.String          `tfsdk:"custom_claims_forward"`
	CustomIntrospectionHeaders map[string]types.String `tfsdk:"custom_introspection_headers"`
	HideCredentials            types.Bool              `tfsdk:"hide_credentials"`
	IntrospectRequest          types.Bool              `tfsdk:"introspect_request"`
	IntrospectionURL           types.String            `tfsdk:"introspection_url"`
	Keepalive                  types.Int64             `tfsdk:"keepalive"`
	RunOnPreflight             types.Bool              `tfsdk:"run_on_preflight"`
	Timeout                    types.Int64             `tfsdk:"timeout"`
	TokenTypeHint              types.String            `tfsdk:"token_type_hint"`
	TTL                        types.Number            `tfsdk:"ttl"`
}

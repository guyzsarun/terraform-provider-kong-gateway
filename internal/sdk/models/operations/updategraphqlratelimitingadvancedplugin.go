// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateGraphqlratelimitingadvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                                string                                          `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateGraphqlRateLimitingAdvancedPlugin *shared.CreateGraphqlRateLimitingAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateGraphqlratelimitingadvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateGraphqlratelimitingadvancedPluginRequest) GetCreateGraphqlRateLimitingAdvancedPlugin() *shared.CreateGraphqlRateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateGraphqlRateLimitingAdvancedPlugin
}

type UpdateGraphqlratelimitingadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// GraphqlRateLimitingAdvanced plugin
	GraphqlRateLimitingAdvancedPlugin *shared.GraphqlRateLimitingAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateGraphqlratelimitingadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateGraphqlratelimitingadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateGraphqlratelimitingadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateGraphqlratelimitingadvancedPluginResponse) GetGraphqlRateLimitingAdvancedPlugin() *shared.GraphqlRateLimitingAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.GraphqlRateLimitingAdvancedPlugin
}

func (o *UpdateGraphqlratelimitingadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

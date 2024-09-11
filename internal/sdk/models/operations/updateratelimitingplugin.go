// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateRatelimitingPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                 string                           `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateRateLimitingPlugin *shared.CreateRateLimitingPlugin `request:"mediaType=application/json"`
}

func (o *UpdateRatelimitingPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateRatelimitingPluginRequest) GetCreateRateLimitingPlugin() *shared.CreateRateLimitingPlugin {
	if o == nil {
		return nil
	}
	return o.CreateRateLimitingPlugin
}

type UpdateRatelimitingPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// RateLimiting plugin
	RateLimitingPlugin *shared.RateLimitingPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRatelimitingPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRatelimitingPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRatelimitingPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRatelimitingPluginResponse) GetRateLimitingPlugin() *shared.RateLimitingPlugin {
	if o == nil {
		return nil
	}
	return o.RateLimitingPlugin
}

func (o *UpdateRatelimitingPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

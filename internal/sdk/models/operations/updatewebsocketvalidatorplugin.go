// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateWebsocketvalidatorPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                       string                                 `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateWebsocketValidatorPlugin *shared.CreateWebsocketValidatorPlugin `request:"mediaType=application/json"`
}

func (o *UpdateWebsocketvalidatorPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateWebsocketvalidatorPluginRequest) GetCreateWebsocketValidatorPlugin() *shared.CreateWebsocketValidatorPlugin {
	if o == nil {
		return nil
	}
	return o.CreateWebsocketValidatorPlugin
}

type UpdateWebsocketvalidatorPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// WebsocketValidator plugin
	WebsocketValidatorPlugin *shared.WebsocketValidatorPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateWebsocketvalidatorPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWebsocketvalidatorPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWebsocketvalidatorPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWebsocketvalidatorPluginResponse) GetWebsocketValidatorPlugin() *shared.WebsocketValidatorPlugin {
	if o == nil {
		return nil
	}
	return o.WebsocketValidatorPlugin
}

func (o *UpdateWebsocketvalidatorPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

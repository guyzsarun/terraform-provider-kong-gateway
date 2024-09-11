// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateStatsdadvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                   string                             `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateStatsdAdvancedPlugin *shared.CreateStatsdAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateStatsdadvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateStatsdadvancedPluginRequest) GetCreateStatsdAdvancedPlugin() *shared.CreateStatsdAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateStatsdAdvancedPlugin
}

type UpdateStatsdadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// StatsdAdvanced plugin
	StatsdAdvancedPlugin *shared.StatsdAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateStatsdadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateStatsdadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateStatsdadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateStatsdadvancedPluginResponse) GetStatsdAdvancedPlugin() *shared.StatsdAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.StatsdAdvancedPlugin
}

func (o *UpdateStatsdadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

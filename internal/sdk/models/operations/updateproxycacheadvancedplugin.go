// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateProxycacheadvancedPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                       string                                 `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateProxyCacheAdvancedPlugin *shared.CreateProxyCacheAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *UpdateProxycacheadvancedPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateProxycacheadvancedPluginRequest) GetCreateProxyCacheAdvancedPlugin() *shared.CreateProxyCacheAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.CreateProxyCacheAdvancedPlugin
}

type UpdateProxycacheadvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ProxyCacheAdvanced plugin
	ProxyCacheAdvancedPlugin *shared.ProxyCacheAdvancedPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateProxycacheadvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProxycacheadvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProxycacheadvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateProxycacheadvancedPluginResponse) GetProxyCacheAdvancedPlugin() *shared.ProxyCacheAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.ProxyCacheAdvancedPlugin
}

func (o *UpdateProxycacheadvancedPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

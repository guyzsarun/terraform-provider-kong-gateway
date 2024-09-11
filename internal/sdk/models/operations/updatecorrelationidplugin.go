// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateCorrelationidPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                  string                            `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateCorrelationIDPlugin *shared.CreateCorrelationIDPlugin `request:"mediaType=application/json"`
}

func (o *UpdateCorrelationidPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateCorrelationidPluginRequest) GetCreateCorrelationIDPlugin() *shared.CreateCorrelationIDPlugin {
	if o == nil {
		return nil
	}
	return o.CreateCorrelationIDPlugin
}

type UpdateCorrelationidPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// CorrelationId plugin
	CorrelationIDPlugin *shared.CorrelationIDPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateCorrelationidPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCorrelationidPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCorrelationidPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCorrelationidPluginResponse) GetCorrelationIDPlugin() *shared.CorrelationIDPlugin {
	if o == nil {
		return nil
	}
	return o.CorrelationIDPlugin
}

func (o *UpdateCorrelationidPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

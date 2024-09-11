// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateExittransformerPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID                    string                              `pathParam:"style=simple,explode=false,name=PluginId"`
	CreateExitTransformerPlugin *shared.CreateExitTransformerPlugin `request:"mediaType=application/json"`
}

func (o *UpdateExittransformerPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateExittransformerPluginRequest) GetCreateExitTransformerPlugin() *shared.CreateExitTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.CreateExitTransformerPlugin
}

type UpdateExittransformerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// ExitTransformer plugin
	ExitTransformerPlugin *shared.ExitTransformerPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateExittransformerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateExittransformerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateExittransformerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateExittransformerPluginResponse) GetExitTransformerPlugin() *shared.ExitTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.ExitTransformerPlugin
}

func (o *UpdateExittransformerPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

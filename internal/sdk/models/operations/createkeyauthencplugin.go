// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type CreateKeyauthencPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created KeyAuthEnc plugin
	KeyAuthEncPlugin *shared.KeyAuthEncPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateKeyauthencPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateKeyauthencPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateKeyauthencPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateKeyauthencPluginResponse) GetKeyAuthEncPlugin() *shared.KeyAuthEncPlugin {
	if o == nil {
		return nil
	}
	return o.KeyAuthEncPlugin
}

func (o *CreateKeyauthencPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateHmacAuthRequest struct {
	// ID of the HMAC-auth credential to lookup
	HMACAuthID string `pathParam:"style=simple,explode=false,name=HMACAuthId"`
	// Fields of the HMAC-auth credential that need to be updated
	HMACAuth shared.HMACAuthInput `request:"mediaType=application/json"`
}

func (o *UpdateHmacAuthRequest) GetHMACAuthID() string {
	if o == nil {
		return ""
	}
	return o.HMACAuthID
}

func (o *UpdateHmacAuthRequest) GetHMACAuth() shared.HMACAuthInput {
	if o == nil {
		return shared.HMACAuthInput{}
	}
	return o.HMACAuth
}

type UpdateHmacAuthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully updated HMAC-auth credential
	HMACAuth *shared.HMACAuth
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateHmacAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateHmacAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateHmacAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateHmacAuthResponse) GetHMACAuth() *shared.HMACAuth {
	if o == nil {
		return nil
	}
	return o.HMACAuth
}

func (o *UpdateHmacAuthResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpsertKeyAuthRequest struct {
	// ID of the API-key to lookup
	KeyAuthID string `pathParam:"style=simple,explode=false,name=KeyAuthId"`
	// Description of the API-key
	KeyAuth shared.KeyAuthInput `request:"mediaType=application/json"`
}

func (o *UpsertKeyAuthRequest) GetKeyAuthID() string {
	if o == nil {
		return ""
	}
	return o.KeyAuthID
}

func (o *UpsertKeyAuthRequest) GetKeyAuth() shared.KeyAuthInput {
	if o == nil {
		return shared.KeyAuthInput{}
	}
	return o.KeyAuth
}

type UpsertKeyAuthResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted API-key
	KeyAuth *shared.KeyAuth
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertKeyAuthResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertKeyAuthResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertKeyAuthResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertKeyAuthResponse) GetKeyAuth() *shared.KeyAuth {
	if o == nil {
		return nil
	}
	return o.KeyAuth
}

func (o *UpsertKeyAuthResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

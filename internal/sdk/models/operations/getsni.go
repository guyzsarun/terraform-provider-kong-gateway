// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type GetSniRequest struct {
	// ID or name of the SNI to lookup
	SNIIDOrName string `pathParam:"style=simple,explode=false,name=SNIIdOrName"`
}

func (o *GetSniRequest) GetSNIIDOrName() string {
	if o == nil {
		return ""
	}
	return o.SNIIDOrName
}

type GetSniResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully fetched SNI
	Sni *shared.Sni
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *GetSniResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSniResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSniResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSniResponse) GetSni() *shared.Sni {
	if o == nil {
		return nil
	}
	return o.Sni
}

func (o *GetSniResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

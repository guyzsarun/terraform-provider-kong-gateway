// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateCertificateRequest struct {
	// ID of the Certificate to lookup
	CertificateID string `pathParam:"style=simple,explode=false,name=CertificateId"`
	// Fields of the Certificate that need to be updated
	Certificate shared.CertificateInput `request:"mediaType=application/json"`
}

func (o *UpdateCertificateRequest) GetCertificateID() string {
	if o == nil {
		return ""
	}
	return o.CertificateID
}

func (o *UpdateCertificateRequest) GetCertificate() shared.CertificateInput {
	if o == nil {
		return shared.CertificateInput{}
	}
	return o.Certificate
}

type UpdateCertificateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully updated Certificate
	Certificate *shared.Certificate
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateCertificateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCertificateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCertificateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCertificateResponse) GetCertificate() *shared.Certificate {
	if o == nil {
		return nil
	}
	return o.Certificate
}

func (o *UpdateCertificateResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

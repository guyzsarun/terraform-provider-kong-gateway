// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateConsumerRequest struct {
	// ID or username of the Consumer to lookup
	ConsumerIDOrUsername string `pathParam:"style=simple,explode=false,name=ConsumerIdOrUsername"`
	// Fields of the Consumer that need to be updated
	Consumer shared.ConsumerInput `request:"mediaType=application/json"`
}

func (o *UpdateConsumerRequest) GetConsumerIDOrUsername() string {
	if o == nil {
		return ""
	}
	return o.ConsumerIDOrUsername
}

func (o *UpdateConsumerRequest) GetConsumer() shared.ConsumerInput {
	if o == nil {
		return shared.ConsumerInput{}
	}
	return o.Consumer
}

type UpdateConsumerResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully updated Consumer
	Consumer *shared.Consumer
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateConsumerResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateConsumerResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateConsumerResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateConsumerResponse) GetConsumer() *shared.Consumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *UpdateConsumerResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

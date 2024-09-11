// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpdateRouteRequest struct {
	// ID or name of the Route to lookup
	RouteIDOrName string `pathParam:"style=simple,explode=false,name=RouteIdOrName"`
	// Fields of the Route that need to be updated
	Route shared.RouteInput `request:"mediaType=application/json"`
}

func (o *UpdateRouteRequest) GetRouteIDOrName() string {
	if o == nil {
		return ""
	}
	return o.RouteIDOrName
}

func (o *UpdateRouteRequest) GetRoute() shared.RouteInput {
	if o == nil {
		return shared.RouteInput{}
	}
	return o.Route
}

type UpdateRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully updated Route
	Route *shared.Route
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateRouteResponse) GetRoute() *shared.Route {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *UpdateRouteResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

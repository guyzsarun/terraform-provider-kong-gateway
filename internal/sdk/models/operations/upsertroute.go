// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type UpsertRouteRequest struct {
	// ID or name of the Route to lookup
	RouteIDOrName string `pathParam:"style=simple,explode=false,name=RouteIdOrName"`
	// Description of the Route
	Route shared.RouteInput `request:"mediaType=application/json"`
}

func (o *UpsertRouteRequest) GetRouteIDOrName() string {
	if o == nil {
		return ""
	}
	return o.RouteIDOrName
}

func (o *UpsertRouteRequest) GetRoute() shared.RouteInput {
	if o == nil {
		return shared.RouteInput{}
	}
	return o.Route
}

type UpsertRouteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully upserted Route
	Route *shared.Route
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpsertRouteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpsertRouteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpsertRouteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpsertRouteResponse) GetRoute() *shared.Route {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *UpsertRouteResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

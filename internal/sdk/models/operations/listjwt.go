// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"net/http"
)

type ListJwtRequest struct {
	// Number of resources to be returned.
	Size *int64 `default:"100" queryParam:"style=form,explode=true,name=size"`
	// Offset from which to return the next set of resources. Use the value of the 'offset' field from the response of a list operation as input here to paginate through all the resources
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// A list of tags to filter the list of resources on. Multiple tags can be concatenated using ',' to mean AND or using '/' to mean OR.
	Tags *string `queryParam:"style=form,explode=true,name=tags"`
}

func (l ListJwtRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListJwtRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListJwtRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListJwtRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *ListJwtRequest) GetTags() *string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// ListJwtResponseBody - A successful response listing JWTs
type ListJwtResponseBody struct {
	Data []shared.Jwt `json:"data,omitempty"`
	// URI to the next page (may be null)
	Next *string `json:"next,omitempty"`
	// Offset is used to paginate through the API. Provide this value to the next list operation to fetch the next page
	Offset *string `json:"offset,omitempty"`
}

func (o *ListJwtResponseBody) GetData() []shared.Jwt {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *ListJwtResponseBody) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListJwtResponseBody) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

type ListJwtResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// A successful response listing JWTs
	Object *ListJwtResponseBody
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *ListJwtResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListJwtResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListJwtResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListJwtResponse) GetObject() *ListJwtResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

func (o *ListJwtResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}

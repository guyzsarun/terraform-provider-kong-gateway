// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type HmacAuthPluginAlgorithms string

const (
	HmacAuthPluginAlgorithmsHmacSha1   HmacAuthPluginAlgorithms = "hmac-sha1"
	HmacAuthPluginAlgorithmsHmacSha256 HmacAuthPluginAlgorithms = "hmac-sha256"
	HmacAuthPluginAlgorithmsHmacSha384 HmacAuthPluginAlgorithms = "hmac-sha384"
	HmacAuthPluginAlgorithmsHmacSha512 HmacAuthPluginAlgorithms = "hmac-sha512"
)

func (e HmacAuthPluginAlgorithms) ToPointer() *HmacAuthPluginAlgorithms {
	return &e
}
func (e *HmacAuthPluginAlgorithms) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "hmac-sha1":
		fallthrough
	case "hmac-sha256":
		fallthrough
	case "hmac-sha384":
		fallthrough
	case "hmac-sha512":
		*e = HmacAuthPluginAlgorithms(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HmacAuthPluginAlgorithms: %v", v)
	}
}

type HmacAuthPluginConfig struct {
	// A list of HMAC digest algorithms that the user wants to support. Allowed values are `hmac-sha1`, `hmac-sha256`, `hmac-sha384`, and `hmac-sha512`
	Algorithms []HmacAuthPluginAlgorithms `json:"algorithms,omitempty"`
	// An optional string (Consumer UUID or username) value to use as an “anonymous” consumer if authentication fails.
	Anonymous *string `json:"anonymous,omitempty"`
	// Clock skew in seconds to prevent replay attacks.
	ClockSkew *float64 `json:"clock_skew,omitempty"`
	// A list of headers that the client should at least use for HTTP signature creation.
	EnforceHeaders []string `json:"enforce_headers,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// A boolean value telling the plugin to enable body validation.
	ValidateRequestBody *bool `json:"validate_request_body,omitempty"`
}

func (o *HmacAuthPluginConfig) GetAlgorithms() []HmacAuthPluginAlgorithms {
	if o == nil {
		return nil
	}
	return o.Algorithms
}

func (o *HmacAuthPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *HmacAuthPluginConfig) GetClockSkew() *float64 {
	if o == nil {
		return nil
	}
	return o.ClockSkew
}

func (o *HmacAuthPluginConfig) GetEnforceHeaders() []string {
	if o == nil {
		return nil
	}
	return o.EnforceHeaders
}

func (o *HmacAuthPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *HmacAuthPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *HmacAuthPluginConfig) GetValidateRequestBody() *bool {
	if o == nil {
		return nil
	}
	return o.ValidateRequestBody
}

type HmacAuthPluginProtocols string

const (
	HmacAuthPluginProtocolsGrpc           HmacAuthPluginProtocols = "grpc"
	HmacAuthPluginProtocolsGrpcs          HmacAuthPluginProtocols = "grpcs"
	HmacAuthPluginProtocolsHTTP           HmacAuthPluginProtocols = "http"
	HmacAuthPluginProtocolsHTTPS          HmacAuthPluginProtocols = "https"
	HmacAuthPluginProtocolsTCP            HmacAuthPluginProtocols = "tcp"
	HmacAuthPluginProtocolsTLS            HmacAuthPluginProtocols = "tls"
	HmacAuthPluginProtocolsTLSPassthrough HmacAuthPluginProtocols = "tls_passthrough"
	HmacAuthPluginProtocolsUDP            HmacAuthPluginProtocols = "udp"
	HmacAuthPluginProtocolsWs             HmacAuthPluginProtocols = "ws"
	HmacAuthPluginProtocolsWss            HmacAuthPluginProtocols = "wss"
)

func (e HmacAuthPluginProtocols) ToPointer() *HmacAuthPluginProtocols {
	return &e
}
func (e *HmacAuthPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = HmacAuthPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HmacAuthPluginProtocols: %v", v)
	}
}

// HmacAuthPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type HmacAuthPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *HmacAuthPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type HmacAuthPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *HmacAuthPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HmacAuthPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type HmacAuthPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *HmacAuthPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// HmacAuthPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type HmacAuthPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *HmacAuthPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type HmacAuthPlugin struct {
	Config *HmacAuthPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"hmac-auth" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []HmacAuthPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *HmacAuthPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *HmacAuthPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *HmacAuthPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *HmacAuthPluginService `json:"service,omitempty"`
}

func (h HmacAuthPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HmacAuthPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HmacAuthPlugin) GetConfig() *HmacAuthPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *HmacAuthPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HmacAuthPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *HmacAuthPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HmacAuthPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *HmacAuthPlugin) GetName() *string {
	return types.String("hmac-auth")
}

func (o *HmacAuthPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *HmacAuthPlugin) GetProtocols() []HmacAuthPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *HmacAuthPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *HmacAuthPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *HmacAuthPlugin) GetConsumer() *HmacAuthPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *HmacAuthPlugin) GetConsumerGroup() *HmacAuthPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *HmacAuthPlugin) GetRoute() *HmacAuthPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *HmacAuthPlugin) GetService() *HmacAuthPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

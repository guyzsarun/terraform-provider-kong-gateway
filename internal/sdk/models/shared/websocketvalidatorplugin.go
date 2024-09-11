// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// WebsocketValidatorPluginType - The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
type WebsocketValidatorPluginType string

const (
	WebsocketValidatorPluginTypeDraft4 WebsocketValidatorPluginType = "draft4"
)

func (e WebsocketValidatorPluginType) ToPointer() *WebsocketValidatorPluginType {
	return &e
}
func (e *WebsocketValidatorPluginType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft4":
		*e = WebsocketValidatorPluginType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketValidatorPluginType: %v", v)
	}
}

type WebsocketValidatorPluginBinary struct {
	// Schema used to validate upstream-originated binary frames. The semantics of this field depend on the validation type set by `config.upstream.binary.type`.
	Schema string `json:"schema"`
	// The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
	Type WebsocketValidatorPluginType `json:"type"`
}

func (o *WebsocketValidatorPluginBinary) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *WebsocketValidatorPluginBinary) GetType() WebsocketValidatorPluginType {
	if o == nil {
		return WebsocketValidatorPluginType("")
	}
	return o.Type
}

// WebsocketValidatorPluginConfigType - The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
type WebsocketValidatorPluginConfigType string

const (
	WebsocketValidatorPluginConfigTypeDraft4 WebsocketValidatorPluginConfigType = "draft4"
)

func (e WebsocketValidatorPluginConfigType) ToPointer() *WebsocketValidatorPluginConfigType {
	return &e
}
func (e *WebsocketValidatorPluginConfigType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft4":
		*e = WebsocketValidatorPluginConfigType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketValidatorPluginConfigType: %v", v)
	}
}

type WebsocketValidatorPluginText struct {
	// Schema used to validate upstream-originated binary frames. The semantics of this field depend on the validation type set by `config.upstream.binary.type`.
	Schema string `json:"schema"`
	// The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
	Type WebsocketValidatorPluginConfigType `json:"type"`
}

func (o *WebsocketValidatorPluginText) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *WebsocketValidatorPluginText) GetType() WebsocketValidatorPluginConfigType {
	if o == nil {
		return WebsocketValidatorPluginConfigType("")
	}
	return o.Type
}

type WebsocketValidatorPluginClient struct {
	Binary *WebsocketValidatorPluginBinary `json:"binary,omitempty"`
	Text   *WebsocketValidatorPluginText   `json:"text,omitempty"`
}

func (o *WebsocketValidatorPluginClient) GetBinary() *WebsocketValidatorPluginBinary {
	if o == nil {
		return nil
	}
	return o.Binary
}

func (o *WebsocketValidatorPluginClient) GetText() *WebsocketValidatorPluginText {
	if o == nil {
		return nil
	}
	return o.Text
}

// WebsocketValidatorPluginConfigUpstreamType - The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
type WebsocketValidatorPluginConfigUpstreamType string

const (
	WebsocketValidatorPluginConfigUpstreamTypeDraft4 WebsocketValidatorPluginConfigUpstreamType = "draft4"
)

func (e WebsocketValidatorPluginConfigUpstreamType) ToPointer() *WebsocketValidatorPluginConfigUpstreamType {
	return &e
}
func (e *WebsocketValidatorPluginConfigUpstreamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft4":
		*e = WebsocketValidatorPluginConfigUpstreamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketValidatorPluginConfigUpstreamType: %v", v)
	}
}

type WebsocketValidatorPluginConfigBinary struct {
	// Schema used to validate upstream-originated binary frames. The semantics of this field depend on the validation type set by `config.upstream.binary.type`.
	Schema string `json:"schema"`
	// The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
	Type WebsocketValidatorPluginConfigUpstreamType `json:"type"`
}

func (o *WebsocketValidatorPluginConfigBinary) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *WebsocketValidatorPluginConfigBinary) GetType() WebsocketValidatorPluginConfigUpstreamType {
	if o == nil {
		return WebsocketValidatorPluginConfigUpstreamType("")
	}
	return o.Type
}

// WebsocketValidatorPluginConfigUpstreamTextType - The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
type WebsocketValidatorPluginConfigUpstreamTextType string

const (
	WebsocketValidatorPluginConfigUpstreamTextTypeDraft4 WebsocketValidatorPluginConfigUpstreamTextType = "draft4"
)

func (e WebsocketValidatorPluginConfigUpstreamTextType) ToPointer() *WebsocketValidatorPluginConfigUpstreamTextType {
	return &e
}
func (e *WebsocketValidatorPluginConfigUpstreamTextType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "draft4":
		*e = WebsocketValidatorPluginConfigUpstreamTextType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketValidatorPluginConfigUpstreamTextType: %v", v)
	}
}

type WebsocketValidatorPluginConfigText struct {
	// Schema used to validate upstream-originated binary frames. The semantics of this field depend on the validation type set by `config.upstream.binary.type`.
	Schema string `json:"schema"`
	// The corresponding validation library for `config.upstream.binary.schema`. Currently, only `draft4` is supported.
	Type WebsocketValidatorPluginConfigUpstreamTextType `json:"type"`
}

func (o *WebsocketValidatorPluginConfigText) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *WebsocketValidatorPluginConfigText) GetType() WebsocketValidatorPluginConfigUpstreamTextType {
	if o == nil {
		return WebsocketValidatorPluginConfigUpstreamTextType("")
	}
	return o.Type
}

type WebsocketValidatorPluginUpstream struct {
	Binary *WebsocketValidatorPluginConfigBinary `json:"binary,omitempty"`
	Text   *WebsocketValidatorPluginConfigText   `json:"text,omitempty"`
}

func (o *WebsocketValidatorPluginUpstream) GetBinary() *WebsocketValidatorPluginConfigBinary {
	if o == nil {
		return nil
	}
	return o.Binary
}

func (o *WebsocketValidatorPluginUpstream) GetText() *WebsocketValidatorPluginConfigText {
	if o == nil {
		return nil
	}
	return o.Text
}

type WebsocketValidatorPluginConfig struct {
	Client   *WebsocketValidatorPluginClient   `json:"client,omitempty"`
	Upstream *WebsocketValidatorPluginUpstream `json:"upstream,omitempty"`
}

func (o *WebsocketValidatorPluginConfig) GetClient() *WebsocketValidatorPluginClient {
	if o == nil {
		return nil
	}
	return o.Client
}

func (o *WebsocketValidatorPluginConfig) GetUpstream() *WebsocketValidatorPluginUpstream {
	if o == nil {
		return nil
	}
	return o.Upstream
}

type WebsocketValidatorPluginProtocols string

const (
	WebsocketValidatorPluginProtocolsGrpc           WebsocketValidatorPluginProtocols = "grpc"
	WebsocketValidatorPluginProtocolsGrpcs          WebsocketValidatorPluginProtocols = "grpcs"
	WebsocketValidatorPluginProtocolsHTTP           WebsocketValidatorPluginProtocols = "http"
	WebsocketValidatorPluginProtocolsHTTPS          WebsocketValidatorPluginProtocols = "https"
	WebsocketValidatorPluginProtocolsTCP            WebsocketValidatorPluginProtocols = "tcp"
	WebsocketValidatorPluginProtocolsTLS            WebsocketValidatorPluginProtocols = "tls"
	WebsocketValidatorPluginProtocolsTLSPassthrough WebsocketValidatorPluginProtocols = "tls_passthrough"
	WebsocketValidatorPluginProtocolsUDP            WebsocketValidatorPluginProtocols = "udp"
	WebsocketValidatorPluginProtocolsWs             WebsocketValidatorPluginProtocols = "ws"
	WebsocketValidatorPluginProtocolsWss            WebsocketValidatorPluginProtocols = "wss"
)

func (e WebsocketValidatorPluginProtocols) ToPointer() *WebsocketValidatorPluginProtocols {
	return &e
}
func (e *WebsocketValidatorPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = WebsocketValidatorPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for WebsocketValidatorPluginProtocols: %v", v)
	}
}

// WebsocketValidatorPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type WebsocketValidatorPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketValidatorPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type WebsocketValidatorPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketValidatorPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// WebsocketValidatorPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type WebsocketValidatorPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketValidatorPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// WebsocketValidatorPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type WebsocketValidatorPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *WebsocketValidatorPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type WebsocketValidatorPlugin struct {
	Config *WebsocketValidatorPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"websocket-validator" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []WebsocketValidatorPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *WebsocketValidatorPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *WebsocketValidatorPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *WebsocketValidatorPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *WebsocketValidatorPluginService `json:"service,omitempty"`
}

func (w WebsocketValidatorPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *WebsocketValidatorPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *WebsocketValidatorPlugin) GetConfig() *WebsocketValidatorPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *WebsocketValidatorPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *WebsocketValidatorPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *WebsocketValidatorPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *WebsocketValidatorPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *WebsocketValidatorPlugin) GetName() *string {
	return types.String("websocket-validator")
}

func (o *WebsocketValidatorPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *WebsocketValidatorPlugin) GetProtocols() []WebsocketValidatorPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *WebsocketValidatorPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *WebsocketValidatorPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *WebsocketValidatorPlugin) GetConsumer() *WebsocketValidatorPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *WebsocketValidatorPlugin) GetConsumerGroup() *WebsocketValidatorPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *WebsocketValidatorPlugin) GetRoute() *WebsocketValidatorPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *WebsocketValidatorPlugin) GetService() *WebsocketValidatorPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// SizeUnit - Size unit can be set either in `bytes`, `kilobytes`, or `megabytes` (default). This configuration is not available in versions prior to Kong Gateway 1.3 and Kong Gateway (OSS) 2.0.
type SizeUnit string

const (
	SizeUnitMegabytes SizeUnit = "megabytes"
	SizeUnitKilobytes SizeUnit = "kilobytes"
	SizeUnitBytes     SizeUnit = "bytes"
)

func (e SizeUnit) ToPointer() *SizeUnit {
	return &e
}
func (e *SizeUnit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "megabytes":
		fallthrough
	case "kilobytes":
		fallthrough
	case "bytes":
		*e = SizeUnit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SizeUnit: %v", v)
	}
}

type CreateRequestSizeLimitingPluginConfig struct {
	// Allowed request payload size in megabytes. Default is `128` megabytes (128000000 bytes).
	AllowedPayloadSize *int64 `json:"allowed_payload_size,omitempty"`
	// Set to `true` to ensure a valid `Content-Length` header exists before reading the request body.
	RequireContentLength *bool `json:"require_content_length,omitempty"`
	// Size unit can be set either in `bytes`, `kilobytes`, or `megabytes` (default). This configuration is not available in versions prior to Kong Gateway 1.3 and Kong Gateway (OSS) 2.0.
	SizeUnit *SizeUnit `json:"size_unit,omitempty"`
}

func (o *CreateRequestSizeLimitingPluginConfig) GetAllowedPayloadSize() *int64 {
	if o == nil {
		return nil
	}
	return o.AllowedPayloadSize
}

func (o *CreateRequestSizeLimitingPluginConfig) GetRequireContentLength() *bool {
	if o == nil {
		return nil
	}
	return o.RequireContentLength
}

func (o *CreateRequestSizeLimitingPluginConfig) GetSizeUnit() *SizeUnit {
	if o == nil {
		return nil
	}
	return o.SizeUnit
}

type CreateRequestSizeLimitingPluginProtocols string

const (
	CreateRequestSizeLimitingPluginProtocolsGrpc           CreateRequestSizeLimitingPluginProtocols = "grpc"
	CreateRequestSizeLimitingPluginProtocolsGrpcs          CreateRequestSizeLimitingPluginProtocols = "grpcs"
	CreateRequestSizeLimitingPluginProtocolsHTTP           CreateRequestSizeLimitingPluginProtocols = "http"
	CreateRequestSizeLimitingPluginProtocolsHTTPS          CreateRequestSizeLimitingPluginProtocols = "https"
	CreateRequestSizeLimitingPluginProtocolsTCP            CreateRequestSizeLimitingPluginProtocols = "tcp"
	CreateRequestSizeLimitingPluginProtocolsTLS            CreateRequestSizeLimitingPluginProtocols = "tls"
	CreateRequestSizeLimitingPluginProtocolsTLSPassthrough CreateRequestSizeLimitingPluginProtocols = "tls_passthrough"
	CreateRequestSizeLimitingPluginProtocolsUDP            CreateRequestSizeLimitingPluginProtocols = "udp"
	CreateRequestSizeLimitingPluginProtocolsWs             CreateRequestSizeLimitingPluginProtocols = "ws"
	CreateRequestSizeLimitingPluginProtocolsWss            CreateRequestSizeLimitingPluginProtocols = "wss"
)

func (e CreateRequestSizeLimitingPluginProtocols) ToPointer() *CreateRequestSizeLimitingPluginProtocols {
	return &e
}
func (e *CreateRequestSizeLimitingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRequestSizeLimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestSizeLimitingPluginProtocols: %v", v)
	}
}

// CreateRequestSizeLimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRequestSizeLimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestSizeLimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestSizeLimitingPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestSizeLimitingPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestSizeLimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRequestSizeLimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestSizeLimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestSizeLimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRequestSizeLimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestSizeLimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestSizeLimitingPlugin struct {
	Config *CreateRequestSizeLimitingPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"request-size-limiting" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRequestSizeLimitingPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateRequestSizeLimitingPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateRequestSizeLimitingPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRequestSizeLimitingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRequestSizeLimitingPluginService `json:"service,omitempty"`
}

func (c CreateRequestSizeLimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRequestSizeLimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRequestSizeLimitingPlugin) GetConfig() *CreateRequestSizeLimitingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateRequestSizeLimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRequestSizeLimitingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateRequestSizeLimitingPlugin) GetName() *string {
	return types.String("request-size-limiting")
}

func (o *CreateRequestSizeLimitingPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateRequestSizeLimitingPlugin) GetProtocols() []CreateRequestSizeLimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRequestSizeLimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRequestSizeLimitingPlugin) GetConsumer() *CreateRequestSizeLimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRequestSizeLimitingPlugin) GetConsumerGroup() *CreateRequestSizeLimitingPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateRequestSizeLimitingPlugin) GetRoute() *CreateRequestSizeLimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRequestSizeLimitingPlugin) GetService() *CreateRequestSizeLimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

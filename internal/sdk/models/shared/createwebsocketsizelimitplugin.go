// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreateWebsocketSizeLimitPluginConfig struct {
	ClientMaxPayload   *int64 `json:"client_max_payload,omitempty"`
	UpstreamMaxPayload *int64 `json:"upstream_max_payload,omitempty"`
}

func (o *CreateWebsocketSizeLimitPluginConfig) GetClientMaxPayload() *int64 {
	if o == nil {
		return nil
	}
	return o.ClientMaxPayload
}

func (o *CreateWebsocketSizeLimitPluginConfig) GetUpstreamMaxPayload() *int64 {
	if o == nil {
		return nil
	}
	return o.UpstreamMaxPayload
}

type CreateWebsocketSizeLimitPluginProtocols string

const (
	CreateWebsocketSizeLimitPluginProtocolsGrpc           CreateWebsocketSizeLimitPluginProtocols = "grpc"
	CreateWebsocketSizeLimitPluginProtocolsGrpcs          CreateWebsocketSizeLimitPluginProtocols = "grpcs"
	CreateWebsocketSizeLimitPluginProtocolsHTTP           CreateWebsocketSizeLimitPluginProtocols = "http"
	CreateWebsocketSizeLimitPluginProtocolsHTTPS          CreateWebsocketSizeLimitPluginProtocols = "https"
	CreateWebsocketSizeLimitPluginProtocolsTCP            CreateWebsocketSizeLimitPluginProtocols = "tcp"
	CreateWebsocketSizeLimitPluginProtocolsTLS            CreateWebsocketSizeLimitPluginProtocols = "tls"
	CreateWebsocketSizeLimitPluginProtocolsTLSPassthrough CreateWebsocketSizeLimitPluginProtocols = "tls_passthrough"
	CreateWebsocketSizeLimitPluginProtocolsUDP            CreateWebsocketSizeLimitPluginProtocols = "udp"
	CreateWebsocketSizeLimitPluginProtocolsWs             CreateWebsocketSizeLimitPluginProtocols = "ws"
	CreateWebsocketSizeLimitPluginProtocolsWss            CreateWebsocketSizeLimitPluginProtocols = "wss"
)

func (e CreateWebsocketSizeLimitPluginProtocols) ToPointer() *CreateWebsocketSizeLimitPluginProtocols {
	return &e
}
func (e *CreateWebsocketSizeLimitPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateWebsocketSizeLimitPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateWebsocketSizeLimitPluginProtocols: %v", v)
	}
}

// CreateWebsocketSizeLimitPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateWebsocketSizeLimitPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateWebsocketSizeLimitPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateWebsocketSizeLimitPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateWebsocketSizeLimitPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateWebsocketSizeLimitPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateWebsocketSizeLimitPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateWebsocketSizeLimitPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateWebsocketSizeLimitPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateWebsocketSizeLimitPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateWebsocketSizeLimitPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateWebsocketSizeLimitPlugin struct {
	Config *CreateWebsocketSizeLimitPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"websocket-size-limit" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateWebsocketSizeLimitPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateWebsocketSizeLimitPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateWebsocketSizeLimitPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateWebsocketSizeLimitPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateWebsocketSizeLimitPluginService `json:"service,omitempty"`
}

func (c CreateWebsocketSizeLimitPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateWebsocketSizeLimitPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateWebsocketSizeLimitPlugin) GetConfig() *CreateWebsocketSizeLimitPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateWebsocketSizeLimitPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateWebsocketSizeLimitPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateWebsocketSizeLimitPlugin) GetName() *string {
	return types.String("websocket-size-limit")
}

func (o *CreateWebsocketSizeLimitPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateWebsocketSizeLimitPlugin) GetProtocols() []CreateWebsocketSizeLimitPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateWebsocketSizeLimitPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateWebsocketSizeLimitPlugin) GetConsumer() *CreateWebsocketSizeLimitPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateWebsocketSizeLimitPlugin) GetConsumerGroup() *CreateWebsocketSizeLimitPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateWebsocketSizeLimitPlugin) GetRoute() *CreateWebsocketSizeLimitPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateWebsocketSizeLimitPlugin) GetService() *CreateWebsocketSizeLimitPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

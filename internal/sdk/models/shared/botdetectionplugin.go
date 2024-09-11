// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type BotDetectionPluginConfig struct {
	// An array of regular expressions that should be allowed. The regular expressions will be checked against the `User-Agent` header.
	Allow []string `json:"allow,omitempty"`
	// An array of regular expressions that should be denied. The regular expressions will be checked against the `User-Agent` header.
	Deny []string `json:"deny,omitempty"`
}

func (o *BotDetectionPluginConfig) GetAllow() []string {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *BotDetectionPluginConfig) GetDeny() []string {
	if o == nil {
		return nil
	}
	return o.Deny
}

type BotDetectionPluginProtocols string

const (
	BotDetectionPluginProtocolsGrpc           BotDetectionPluginProtocols = "grpc"
	BotDetectionPluginProtocolsGrpcs          BotDetectionPluginProtocols = "grpcs"
	BotDetectionPluginProtocolsHTTP           BotDetectionPluginProtocols = "http"
	BotDetectionPluginProtocolsHTTPS          BotDetectionPluginProtocols = "https"
	BotDetectionPluginProtocolsTCP            BotDetectionPluginProtocols = "tcp"
	BotDetectionPluginProtocolsTLS            BotDetectionPluginProtocols = "tls"
	BotDetectionPluginProtocolsTLSPassthrough BotDetectionPluginProtocols = "tls_passthrough"
	BotDetectionPluginProtocolsUDP            BotDetectionPluginProtocols = "udp"
	BotDetectionPluginProtocolsWs             BotDetectionPluginProtocols = "ws"
	BotDetectionPluginProtocolsWss            BotDetectionPluginProtocols = "wss"
)

func (e BotDetectionPluginProtocols) ToPointer() *BotDetectionPluginProtocols {
	return &e
}
func (e *BotDetectionPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = BotDetectionPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BotDetectionPluginProtocols: %v", v)
	}
}

// BotDetectionPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type BotDetectionPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *BotDetectionPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type BotDetectionPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *BotDetectionPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// BotDetectionPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type BotDetectionPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *BotDetectionPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// BotDetectionPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type BotDetectionPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *BotDetectionPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type BotDetectionPlugin struct {
	Config *BotDetectionPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"bot-detection" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []BotDetectionPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *BotDetectionPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *BotDetectionPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *BotDetectionPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *BotDetectionPluginService `json:"service,omitempty"`
}

func (b BotDetectionPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BotDetectionPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BotDetectionPlugin) GetConfig() *BotDetectionPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *BotDetectionPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *BotDetectionPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *BotDetectionPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *BotDetectionPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *BotDetectionPlugin) GetName() *string {
	return types.String("bot-detection")
}

func (o *BotDetectionPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *BotDetectionPlugin) GetProtocols() []BotDetectionPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *BotDetectionPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *BotDetectionPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *BotDetectionPlugin) GetConsumer() *BotDetectionPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *BotDetectionPlugin) GetConsumerGroup() *BotDetectionPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *BotDetectionPlugin) GetRoute() *BotDetectionPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *BotDetectionPlugin) GetService() *BotDetectionPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

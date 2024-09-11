// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type DegraphqlPluginConfig struct {
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	GraphqlServerPath *string `json:"graphql_server_path,omitempty"`
}

func (o *DegraphqlPluginConfig) GetGraphqlServerPath() *string {
	if o == nil {
		return nil
	}
	return o.GraphqlServerPath
}

type DegraphqlPluginProtocols string

const (
	DegraphqlPluginProtocolsGrpc           DegraphqlPluginProtocols = "grpc"
	DegraphqlPluginProtocolsGrpcs          DegraphqlPluginProtocols = "grpcs"
	DegraphqlPluginProtocolsHTTP           DegraphqlPluginProtocols = "http"
	DegraphqlPluginProtocolsHTTPS          DegraphqlPluginProtocols = "https"
	DegraphqlPluginProtocolsTCP            DegraphqlPluginProtocols = "tcp"
	DegraphqlPluginProtocolsTLS            DegraphqlPluginProtocols = "tls"
	DegraphqlPluginProtocolsTLSPassthrough DegraphqlPluginProtocols = "tls_passthrough"
	DegraphqlPluginProtocolsUDP            DegraphqlPluginProtocols = "udp"
	DegraphqlPluginProtocolsWs             DegraphqlPluginProtocols = "ws"
	DegraphqlPluginProtocolsWss            DegraphqlPluginProtocols = "wss"
)

func (e DegraphqlPluginProtocols) ToPointer() *DegraphqlPluginProtocols {
	return &e
}
func (e *DegraphqlPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = DegraphqlPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DegraphqlPluginProtocols: %v", v)
	}
}

// DegraphqlPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type DegraphqlPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *DegraphqlPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type DegraphqlPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *DegraphqlPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// DegraphqlPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type DegraphqlPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *DegraphqlPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// DegraphqlPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type DegraphqlPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *DegraphqlPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type DegraphqlPlugin struct {
	Config *DegraphqlPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"degraphql" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []DegraphqlPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *DegraphqlPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *DegraphqlPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *DegraphqlPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *DegraphqlPluginService `json:"service,omitempty"`
}

func (d DegraphqlPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(d, "", false)
}

func (d *DegraphqlPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &d, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *DegraphqlPlugin) GetConfig() *DegraphqlPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *DegraphqlPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *DegraphqlPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *DegraphqlPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *DegraphqlPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *DegraphqlPlugin) GetName() *string {
	return types.String("degraphql")
}

func (o *DegraphqlPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *DegraphqlPlugin) GetProtocols() []DegraphqlPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *DegraphqlPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *DegraphqlPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *DegraphqlPlugin) GetConsumer() *DegraphqlPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *DegraphqlPlugin) GetConsumerGroup() *DegraphqlPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *DegraphqlPlugin) GetRoute() *DegraphqlPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *DegraphqlPlugin) GetService() *DegraphqlPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

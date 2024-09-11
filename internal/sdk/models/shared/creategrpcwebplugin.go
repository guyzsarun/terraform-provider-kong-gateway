// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreateGrpcWebPluginConfig struct {
	// The value of the `Access-Control-Allow-Origin` header in the response to the gRPC-Web client.
	AllowOriginHeader *string `json:"allow_origin_header,omitempty"`
	// If set to `true` causes the plugin to pass the stripped request path to the upstream gRPC service.
	PassStrippedPath *bool `json:"pass_stripped_path,omitempty"`
	// If present, describes the gRPC types and methods. Required to support payload transcoding. When absent, the web client must use application/grpw-web+proto content.
	Proto *string `json:"proto,omitempty"`
}

func (o *CreateGrpcWebPluginConfig) GetAllowOriginHeader() *string {
	if o == nil {
		return nil
	}
	return o.AllowOriginHeader
}

func (o *CreateGrpcWebPluginConfig) GetPassStrippedPath() *bool {
	if o == nil {
		return nil
	}
	return o.PassStrippedPath
}

func (o *CreateGrpcWebPluginConfig) GetProto() *string {
	if o == nil {
		return nil
	}
	return o.Proto
}

type CreateGrpcWebPluginProtocols string

const (
	CreateGrpcWebPluginProtocolsGrpc           CreateGrpcWebPluginProtocols = "grpc"
	CreateGrpcWebPluginProtocolsGrpcs          CreateGrpcWebPluginProtocols = "grpcs"
	CreateGrpcWebPluginProtocolsHTTP           CreateGrpcWebPluginProtocols = "http"
	CreateGrpcWebPluginProtocolsHTTPS          CreateGrpcWebPluginProtocols = "https"
	CreateGrpcWebPluginProtocolsTCP            CreateGrpcWebPluginProtocols = "tcp"
	CreateGrpcWebPluginProtocolsTLS            CreateGrpcWebPluginProtocols = "tls"
	CreateGrpcWebPluginProtocolsTLSPassthrough CreateGrpcWebPluginProtocols = "tls_passthrough"
	CreateGrpcWebPluginProtocolsUDP            CreateGrpcWebPluginProtocols = "udp"
	CreateGrpcWebPluginProtocolsWs             CreateGrpcWebPluginProtocols = "ws"
	CreateGrpcWebPluginProtocolsWss            CreateGrpcWebPluginProtocols = "wss"
)

func (e CreateGrpcWebPluginProtocols) ToPointer() *CreateGrpcWebPluginProtocols {
	return &e
}
func (e *CreateGrpcWebPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateGrpcWebPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateGrpcWebPluginProtocols: %v", v)
	}
}

// CreateGrpcWebPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateGrpcWebPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateGrpcWebPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateGrpcWebPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateGrpcWebPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateGrpcWebPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateGrpcWebPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateGrpcWebPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateGrpcWebPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateGrpcWebPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateGrpcWebPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateGrpcWebPlugin struct {
	Config *CreateGrpcWebPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"grpc-web" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateGrpcWebPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateGrpcWebPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateGrpcWebPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateGrpcWebPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateGrpcWebPluginService `json:"service,omitempty"`
}

func (c CreateGrpcWebPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateGrpcWebPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateGrpcWebPlugin) GetConfig() *CreateGrpcWebPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateGrpcWebPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateGrpcWebPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateGrpcWebPlugin) GetName() *string {
	return types.String("grpc-web")
}

func (o *CreateGrpcWebPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateGrpcWebPlugin) GetProtocols() []CreateGrpcWebPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateGrpcWebPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateGrpcWebPlugin) GetConsumer() *CreateGrpcWebPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateGrpcWebPlugin) GetConsumerGroup() *CreateGrpcWebPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateGrpcWebPlugin) GetRoute() *CreateGrpcWebPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateGrpcWebPlugin) GetService() *CreateGrpcWebPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

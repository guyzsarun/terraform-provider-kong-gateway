// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreatePrometheusPluginConfig struct {
	// A boolean value that determines if bandwidth metrics should be collected. If enabled, `bandwidth_bytes` and `stream_sessions_total` metrics will be exported.
	BandwidthMetrics *bool `json:"bandwidth_metrics,omitempty"`
	// A boolean value that determines if latency metrics should be collected. If enabled, `kong_latency_ms`, `upstream_latency_ms` and `request_latency_ms` metrics will be exported.
	LatencyMetrics *bool `json:"latency_metrics,omitempty"`
	// A boolean value that determines if per-consumer metrics should be collected. If enabled, the `kong_http_requests_total` and `kong_bandwidth_bytes` metrics fill in the consumer label when available.
	PerConsumer *bool `json:"per_consumer,omitempty"`
	// A boolean value that determines if status code metrics should be collected. If enabled, `http_requests_total`, `stream_sessions_total` metrics will be exported.
	StatusCodeMetrics *bool `json:"status_code_metrics,omitempty"`
	// A boolean value that determines if upstream metrics should be collected. If enabled, `upstream_target_health` metric will be exported.
	UpstreamHealthMetrics *bool `json:"upstream_health_metrics,omitempty"`
}

func (o *CreatePrometheusPluginConfig) GetBandwidthMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.BandwidthMetrics
}

func (o *CreatePrometheusPluginConfig) GetLatencyMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.LatencyMetrics
}

func (o *CreatePrometheusPluginConfig) GetPerConsumer() *bool {
	if o == nil {
		return nil
	}
	return o.PerConsumer
}

func (o *CreatePrometheusPluginConfig) GetStatusCodeMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.StatusCodeMetrics
}

func (o *CreatePrometheusPluginConfig) GetUpstreamHealthMetrics() *bool {
	if o == nil {
		return nil
	}
	return o.UpstreamHealthMetrics
}

type CreatePrometheusPluginProtocols string

const (
	CreatePrometheusPluginProtocolsGrpc           CreatePrometheusPluginProtocols = "grpc"
	CreatePrometheusPluginProtocolsGrpcs          CreatePrometheusPluginProtocols = "grpcs"
	CreatePrometheusPluginProtocolsHTTP           CreatePrometheusPluginProtocols = "http"
	CreatePrometheusPluginProtocolsHTTPS          CreatePrometheusPluginProtocols = "https"
	CreatePrometheusPluginProtocolsTCP            CreatePrometheusPluginProtocols = "tcp"
	CreatePrometheusPluginProtocolsTLS            CreatePrometheusPluginProtocols = "tls"
	CreatePrometheusPluginProtocolsTLSPassthrough CreatePrometheusPluginProtocols = "tls_passthrough"
	CreatePrometheusPluginProtocolsUDP            CreatePrometheusPluginProtocols = "udp"
	CreatePrometheusPluginProtocolsWs             CreatePrometheusPluginProtocols = "ws"
	CreatePrometheusPluginProtocolsWss            CreatePrometheusPluginProtocols = "wss"
)

func (e CreatePrometheusPluginProtocols) ToPointer() *CreatePrometheusPluginProtocols {
	return &e
}
func (e *CreatePrometheusPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreatePrometheusPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreatePrometheusPluginProtocols: %v", v)
	}
}

// CreatePrometheusPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreatePrometheusPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePrometheusPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreatePrometheusPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePrometheusPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreatePrometheusPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreatePrometheusPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePrometheusPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreatePrometheusPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreatePrometheusPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreatePrometheusPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreatePrometheusPlugin struct {
	Config *CreatePrometheusPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"prometheus" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreatePrometheusPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreatePrometheusPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreatePrometheusPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreatePrometheusPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreatePrometheusPluginService `json:"service,omitempty"`
}

func (c CreatePrometheusPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePrometheusPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePrometheusPlugin) GetConfig() *CreatePrometheusPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreatePrometheusPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreatePrometheusPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreatePrometheusPlugin) GetName() *string {
	return types.String("prometheus")
}

func (o *CreatePrometheusPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreatePrometheusPlugin) GetProtocols() []CreatePrometheusPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreatePrometheusPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreatePrometheusPlugin) GetConsumer() *CreatePrometheusPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreatePrometheusPlugin) GetConsumerGroup() *CreatePrometheusPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreatePrometheusPlugin) GetRoute() *CreatePrometheusPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreatePrometheusPlugin) GetService() *CreatePrometheusPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

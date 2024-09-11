// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// Hash algorithm to be used for canary release.
//
// * `consumer`: The hash will be based on the consumer.
// * `ip`: The hash will be based on the client IP address.
// * `none`: No hash will be applied.
// * `allow`: Allows the specified groups to access the canary release.
// * `deny`: Denies the specified groups from accessing the canary release.
// * `header`: The hash will be based on the specified header value.
type Hash string

const (
	HashConsumer Hash = "consumer"
	HashIP       Hash = "ip"
	HashNone     Hash = "none"
	HashAllow    Hash = "allow"
	HashDeny     Hash = "deny"
	HashHeader   Hash = "header"
)

func (e Hash) ToPointer() *Hash {
	return &e
}
func (e *Hash) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "ip":
		fallthrough
	case "none":
		fallthrough
	case "allow":
		fallthrough
	case "deny":
		fallthrough
	case "header":
		*e = Hash(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hash: %v", v)
	}
}

type CreateCanaryPluginConfig struct {
	// A string representing an HTTP header name.
	CanaryByHeaderName *string `json:"canary_by_header_name,omitempty"`
	// The duration of the canary release in seconds.
	Duration *float64 `json:"duration,omitempty"`
	// The groups allowed to access the canary release.
	Groups []string `json:"groups,omitempty"`
	// Hash algorithm to be used for canary release.
	//
	// * `consumer`: The hash will be based on the consumer.
	// * `ip`: The hash will be based on the client IP address.
	// * `none`: No hash will be applied.
	// * `allow`: Allows the specified groups to access the canary release.
	// * `deny`: Denies the specified groups from accessing the canary release.
	// * `header`: The hash will be based on the specified header value.
	Hash *Hash `json:"hash,omitempty"`
	// A string representing an HTTP header name.
	HashHeader *string `json:"hash_header,omitempty"`
	// The percentage of traffic to be routed to the canary release.
	Percentage *float64 `json:"percentage,omitempty"`
	// Future time in seconds since epoch, when the canary release will start. Ignored when `percentage` is set, or when using `allow` or `deny` in `hash`.
	Start *float64 `json:"start,omitempty"`
	// The number of steps for the canary release.
	Steps *float64 `json:"steps,omitempty"`
	// Specifies whether to fallback to the upstream server if the canary release fails.
	UpstreamFallback *bool `json:"upstream_fallback,omitempty"`
	// A string representing a host name, such as example.com.
	UpstreamHost *string `json:"upstream_host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	UpstreamPort *int64 `json:"upstream_port,omitempty"`
	// The URI of the upstream server to be used for the canary release.
	UpstreamURI *string `json:"upstream_uri,omitempty"`
}

func (o *CreateCanaryPluginConfig) GetCanaryByHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.CanaryByHeaderName
}

func (o *CreateCanaryPluginConfig) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *CreateCanaryPluginConfig) GetGroups() []string {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *CreateCanaryPluginConfig) GetHash() *Hash {
	if o == nil {
		return nil
	}
	return o.Hash
}

func (o *CreateCanaryPluginConfig) GetHashHeader() *string {
	if o == nil {
		return nil
	}
	return o.HashHeader
}

func (o *CreateCanaryPluginConfig) GetPercentage() *float64 {
	if o == nil {
		return nil
	}
	return o.Percentage
}

func (o *CreateCanaryPluginConfig) GetStart() *float64 {
	if o == nil {
		return nil
	}
	return o.Start
}

func (o *CreateCanaryPluginConfig) GetSteps() *float64 {
	if o == nil {
		return nil
	}
	return o.Steps
}

func (o *CreateCanaryPluginConfig) GetUpstreamFallback() *bool {
	if o == nil {
		return nil
	}
	return o.UpstreamFallback
}

func (o *CreateCanaryPluginConfig) GetUpstreamHost() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamHost
}

func (o *CreateCanaryPluginConfig) GetUpstreamPort() *int64 {
	if o == nil {
		return nil
	}
	return o.UpstreamPort
}

func (o *CreateCanaryPluginConfig) GetUpstreamURI() *string {
	if o == nil {
		return nil
	}
	return o.UpstreamURI
}

type CreateCanaryPluginProtocols string

const (
	CreateCanaryPluginProtocolsGrpc           CreateCanaryPluginProtocols = "grpc"
	CreateCanaryPluginProtocolsGrpcs          CreateCanaryPluginProtocols = "grpcs"
	CreateCanaryPluginProtocolsHTTP           CreateCanaryPluginProtocols = "http"
	CreateCanaryPluginProtocolsHTTPS          CreateCanaryPluginProtocols = "https"
	CreateCanaryPluginProtocolsTCP            CreateCanaryPluginProtocols = "tcp"
	CreateCanaryPluginProtocolsTLS            CreateCanaryPluginProtocols = "tls"
	CreateCanaryPluginProtocolsTLSPassthrough CreateCanaryPluginProtocols = "tls_passthrough"
	CreateCanaryPluginProtocolsUDP            CreateCanaryPluginProtocols = "udp"
	CreateCanaryPluginProtocolsWs             CreateCanaryPluginProtocols = "ws"
	CreateCanaryPluginProtocolsWss            CreateCanaryPluginProtocols = "wss"
)

func (e CreateCanaryPluginProtocols) ToPointer() *CreateCanaryPluginProtocols {
	return &e
}
func (e *CreateCanaryPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateCanaryPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCanaryPluginProtocols: %v", v)
	}
}

// CreateCanaryPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateCanaryPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCanaryPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateCanaryPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCanaryPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCanaryPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateCanaryPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCanaryPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateCanaryPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateCanaryPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateCanaryPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateCanaryPlugin struct {
	Config *CreateCanaryPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"canary" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateCanaryPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateCanaryPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateCanaryPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateCanaryPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateCanaryPluginService `json:"service,omitempty"`
}

func (c CreateCanaryPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCanaryPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCanaryPlugin) GetConfig() *CreateCanaryPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateCanaryPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateCanaryPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateCanaryPlugin) GetName() *string {
	return types.String("canary")
}

func (o *CreateCanaryPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateCanaryPlugin) GetProtocols() []CreateCanaryPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateCanaryPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateCanaryPlugin) GetConsumer() *CreateCanaryPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateCanaryPlugin) GetConsumerGroup() *CreateCanaryPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateCanaryPlugin) GetRoute() *CreateCanaryPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateCanaryPlugin) GetService() *CreateCanaryPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

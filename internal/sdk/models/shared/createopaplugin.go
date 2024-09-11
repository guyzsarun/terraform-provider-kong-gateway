// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// OpaProtocol - The protocol to use when talking to Open Policy Agent (OPA) server. Allowed protocols are `http` and `https`.
type OpaProtocol string

const (
	OpaProtocolHTTP  OpaProtocol = "http"
	OpaProtocolHTTPS OpaProtocol = "https"
)

func (e OpaProtocol) ToPointer() *OpaProtocol {
	return &e
}
func (e *OpaProtocol) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "http":
		fallthrough
	case "https":
		*e = OpaProtocol(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OpaProtocol: %v", v)
	}
}

type CreateOpaPluginConfig struct {
	IncludeBodyInOpaInput *bool `json:"include_body_in_opa_input,omitempty"`
	// If set to true, the Kong Gateway Consumer object in use for the current request (if any) is included as input to OPA.
	IncludeConsumerInOpaInput *bool `json:"include_consumer_in_opa_input,omitempty"`
	// If set to true and the `Content-Type` header of the current request is `application/json`, the request body will be JSON decoded and the decoded struct is included as input to OPA.
	IncludeParsedJSONBodyInOpaInput *bool `json:"include_parsed_json_body_in_opa_input,omitempty"`
	// If set to true, the Kong Gateway Route object in use for the current request is included as input to OPA.
	IncludeRouteInOpaInput *bool `json:"include_route_in_opa_input,omitempty"`
	// If set to true, the Kong Gateway Service object in use for the current request is included as input to OPA.
	IncludeServiceInOpaInput *bool `json:"include_service_in_opa_input,omitempty"`
	// If set to true, the regex capture groups captured on the Kong Gateway Route's path field in the current request (if any) are included as input to OPA.
	IncludeURICapturesInOpaInput *bool `json:"include_uri_captures_in_opa_input,omitempty"`
	// A string representing a host name, such as example.com.
	OpaHost *string `json:"opa_host,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	OpaPath *string `json:"opa_path,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	OpaPort *int64 `json:"opa_port,omitempty"`
	// The protocol to use when talking to Open Policy Agent (OPA) server. Allowed protocols are `http` and `https`.
	OpaProtocol *OpaProtocol `json:"opa_protocol,omitempty"`
	// If set to true, the OPA certificate will be verified according to the CA certificates specified in lua_ssl_trusted_certificate.
	SslVerify *bool `json:"ssl_verify,omitempty"`
}

func (o *CreateOpaPluginConfig) GetIncludeBodyInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeBodyInOpaInput
}

func (o *CreateOpaPluginConfig) GetIncludeConsumerInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeConsumerInOpaInput
}

func (o *CreateOpaPluginConfig) GetIncludeParsedJSONBodyInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeParsedJSONBodyInOpaInput
}

func (o *CreateOpaPluginConfig) GetIncludeRouteInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeRouteInOpaInput
}

func (o *CreateOpaPluginConfig) GetIncludeServiceInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeServiceInOpaInput
}

func (o *CreateOpaPluginConfig) GetIncludeURICapturesInOpaInput() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeURICapturesInOpaInput
}

func (o *CreateOpaPluginConfig) GetOpaHost() *string {
	if o == nil {
		return nil
	}
	return o.OpaHost
}

func (o *CreateOpaPluginConfig) GetOpaPath() *string {
	if o == nil {
		return nil
	}
	return o.OpaPath
}

func (o *CreateOpaPluginConfig) GetOpaPort() *int64 {
	if o == nil {
		return nil
	}
	return o.OpaPort
}

func (o *CreateOpaPluginConfig) GetOpaProtocol() *OpaProtocol {
	if o == nil {
		return nil
	}
	return o.OpaProtocol
}

func (o *CreateOpaPluginConfig) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

type CreateOpaPluginProtocols string

const (
	CreateOpaPluginProtocolsGrpc           CreateOpaPluginProtocols = "grpc"
	CreateOpaPluginProtocolsGrpcs          CreateOpaPluginProtocols = "grpcs"
	CreateOpaPluginProtocolsHTTP           CreateOpaPluginProtocols = "http"
	CreateOpaPluginProtocolsHTTPS          CreateOpaPluginProtocols = "https"
	CreateOpaPluginProtocolsTCP            CreateOpaPluginProtocols = "tcp"
	CreateOpaPluginProtocolsTLS            CreateOpaPluginProtocols = "tls"
	CreateOpaPluginProtocolsTLSPassthrough CreateOpaPluginProtocols = "tls_passthrough"
	CreateOpaPluginProtocolsUDP            CreateOpaPluginProtocols = "udp"
	CreateOpaPluginProtocolsWs             CreateOpaPluginProtocols = "ws"
	CreateOpaPluginProtocolsWss            CreateOpaPluginProtocols = "wss"
)

func (e CreateOpaPluginProtocols) ToPointer() *CreateOpaPluginProtocols {
	return &e
}
func (e *CreateOpaPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateOpaPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateOpaPluginProtocols: %v", v)
	}
}

// CreateOpaPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateOpaPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpaPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOpaPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpaPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOpaPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateOpaPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpaPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateOpaPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateOpaPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateOpaPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateOpaPlugin struct {
	Config *CreateOpaPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"opa" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateOpaPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateOpaPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateOpaPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateOpaPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateOpaPluginService `json:"service,omitempty"`
}

func (c CreateOpaPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateOpaPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateOpaPlugin) GetConfig() *CreateOpaPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateOpaPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateOpaPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateOpaPlugin) GetName() *string {
	return types.String("opa")
}

func (o *CreateOpaPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateOpaPlugin) GetProtocols() []CreateOpaPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateOpaPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateOpaPlugin) GetConsumer() *CreateOpaPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateOpaPlugin) GetConsumerGroup() *CreateOpaPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateOpaPlugin) GetRoute() *CreateOpaPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateOpaPlugin) GetService() *CreateOpaPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

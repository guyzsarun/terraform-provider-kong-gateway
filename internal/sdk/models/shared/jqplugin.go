// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type JQPluginRequestJQProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *JQPluginRequestJQProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *JQPluginRequestJQProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *JQPluginRequestJQProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *JQPluginRequestJQProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *JQPluginRequestJQProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type JQPluginResponseJQProgramOptions struct {
	ASCIIOutput   *bool `json:"ascii_output,omitempty"`
	CompactOutput *bool `json:"compact_output,omitempty"`
	JoinOutput    *bool `json:"join_output,omitempty"`
	RawOutput     *bool `json:"raw_output,omitempty"`
	SortKeys      *bool `json:"sort_keys,omitempty"`
}

func (o *JQPluginResponseJQProgramOptions) GetASCIIOutput() *bool {
	if o == nil {
		return nil
	}
	return o.ASCIIOutput
}

func (o *JQPluginResponseJQProgramOptions) GetCompactOutput() *bool {
	if o == nil {
		return nil
	}
	return o.CompactOutput
}

func (o *JQPluginResponseJQProgramOptions) GetJoinOutput() *bool {
	if o == nil {
		return nil
	}
	return o.JoinOutput
}

func (o *JQPluginResponseJQProgramOptions) GetRawOutput() *bool {
	if o == nil {
		return nil
	}
	return o.RawOutput
}

func (o *JQPluginResponseJQProgramOptions) GetSortKeys() *bool {
	if o == nil {
		return nil
	}
	return o.SortKeys
}

type JQPluginConfig struct {
	RequestIfMediaType       []string                          `json:"request_if_media_type,omitempty"`
	RequestJqProgram         *string                           `json:"request_jq_program,omitempty"`
	RequestJqProgramOptions  *JQPluginRequestJQProgramOptions  `json:"request_jq_program_options,omitempty"`
	ResponseIfMediaType      []string                          `json:"response_if_media_type,omitempty"`
	ResponseIfStatusCode     []int64                           `json:"response_if_status_code,omitempty"`
	ResponseJqProgram        *string                           `json:"response_jq_program,omitempty"`
	ResponseJqProgramOptions *JQPluginResponseJQProgramOptions `json:"response_jq_program_options,omitempty"`
}

func (o *JQPluginConfig) GetRequestIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.RequestIfMediaType
}

func (o *JQPluginConfig) GetRequestJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.RequestJqProgram
}

func (o *JQPluginConfig) GetRequestJqProgramOptions() *JQPluginRequestJQProgramOptions {
	if o == nil {
		return nil
	}
	return o.RequestJqProgramOptions
}

func (o *JQPluginConfig) GetResponseIfMediaType() []string {
	if o == nil {
		return nil
	}
	return o.ResponseIfMediaType
}

func (o *JQPluginConfig) GetResponseIfStatusCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseIfStatusCode
}

func (o *JQPluginConfig) GetResponseJqProgram() *string {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgram
}

func (o *JQPluginConfig) GetResponseJqProgramOptions() *JQPluginResponseJQProgramOptions {
	if o == nil {
		return nil
	}
	return o.ResponseJqProgramOptions
}

type JQPluginProtocols string

const (
	JQPluginProtocolsGrpc           JQPluginProtocols = "grpc"
	JQPluginProtocolsGrpcs          JQPluginProtocols = "grpcs"
	JQPluginProtocolsHTTP           JQPluginProtocols = "http"
	JQPluginProtocolsHTTPS          JQPluginProtocols = "https"
	JQPluginProtocolsTCP            JQPluginProtocols = "tcp"
	JQPluginProtocolsTLS            JQPluginProtocols = "tls"
	JQPluginProtocolsTLSPassthrough JQPluginProtocols = "tls_passthrough"
	JQPluginProtocolsUDP            JQPluginProtocols = "udp"
	JQPluginProtocolsWs             JQPluginProtocols = "ws"
	JQPluginProtocolsWss            JQPluginProtocols = "wss"
)

func (e JQPluginProtocols) ToPointer() *JQPluginProtocols {
	return &e
}
func (e *JQPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = JQPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for JQPluginProtocols: %v", v)
	}
}

// JQPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type JQPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *JQPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JQPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *JQPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JQPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type JQPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *JQPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// JQPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type JQPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *JQPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type JQPlugin struct {
	Config *JQPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"jq" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []JQPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *JQPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *JQPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *JQPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *JQPluginService `json:"service,omitempty"`
}

func (j JQPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(j, "", false)
}

func (j *JQPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &j, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *JQPlugin) GetConfig() *JQPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *JQPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *JQPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *JQPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *JQPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *JQPlugin) GetName() *string {
	return types.String("jq")
}

func (o *JQPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *JQPlugin) GetProtocols() []JQPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *JQPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *JQPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *JQPlugin) GetConsumer() *JQPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *JQPlugin) GetConsumerGroup() *JQPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *JQPlugin) GetRoute() *JQPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *JQPlugin) GetService() *JQPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

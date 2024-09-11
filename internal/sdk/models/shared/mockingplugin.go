// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type MockingPluginConfig struct {
	// The contents of the specification file. You must use this option for hybrid or DB-less mode. You can include the full specification as part of the configuration. In Kong Manager, you can copy and paste the contents of the spec directly into the `Config.Api Specification` text field.
	APISpecification *string `json:"api_specification,omitempty"`
	// The path and name of the specification file loaded into Kong Gateway's database. You cannot use this option for DB-less or hybrid mode.
	APISpecificationFilename *string `json:"api_specification_filename,omitempty"`
	// The base path to be used for path match evaluation. This value is ignored if `include_base_path` is set to `false`.
	CustomBasePath *string `json:"custom_base_path,omitempty"`
	// Indicates whether to include the base path when performing path match evaluation.
	IncludeBasePath *bool `json:"include_base_path,omitempty"`
	// A global list of the HTTP status codes that can only be selected and returned.
	IncludedStatusCodes []int64 `json:"included_status_codes,omitempty"`
	// The maximum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be greater than the `min_delay_time`.
	MaxDelayTime *float64 `json:"max_delay_time,omitempty"`
	// The minimum value in seconds of delay time. Set this value when `random_delay` is enabled and you want to adjust the default. The value must be less than the `max_delay_time`.
	MinDelayTime *float64 `json:"min_delay_time,omitempty"`
	// Enables a random delay in the mocked response. Introduces delays to simulate real-time response times by APIs.
	RandomDelay *bool `json:"random_delay,omitempty"`
	// Randomly selects one example and returns it. This parameter requires the spec to have multiple examples configured.
	RandomExamples *bool `json:"random_examples,omitempty"`
	// Determines whether to randomly select an HTTP status code from the responses of the corresponding API method. The default value is `false`, which means the minimum HTTP status code is always selected and returned.
	RandomStatusCode *bool `json:"random_status_code,omitempty"`
}

func (o *MockingPluginConfig) GetAPISpecification() *string {
	if o == nil {
		return nil
	}
	return o.APISpecification
}

func (o *MockingPluginConfig) GetAPISpecificationFilename() *string {
	if o == nil {
		return nil
	}
	return o.APISpecificationFilename
}

func (o *MockingPluginConfig) GetCustomBasePath() *string {
	if o == nil {
		return nil
	}
	return o.CustomBasePath
}

func (o *MockingPluginConfig) GetIncludeBasePath() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeBasePath
}

func (o *MockingPluginConfig) GetIncludedStatusCodes() []int64 {
	if o == nil {
		return nil
	}
	return o.IncludedStatusCodes
}

func (o *MockingPluginConfig) GetMaxDelayTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxDelayTime
}

func (o *MockingPluginConfig) GetMinDelayTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MinDelayTime
}

func (o *MockingPluginConfig) GetRandomDelay() *bool {
	if o == nil {
		return nil
	}
	return o.RandomDelay
}

func (o *MockingPluginConfig) GetRandomExamples() *bool {
	if o == nil {
		return nil
	}
	return o.RandomExamples
}

func (o *MockingPluginConfig) GetRandomStatusCode() *bool {
	if o == nil {
		return nil
	}
	return o.RandomStatusCode
}

type MockingPluginProtocols string

const (
	MockingPluginProtocolsGrpc           MockingPluginProtocols = "grpc"
	MockingPluginProtocolsGrpcs          MockingPluginProtocols = "grpcs"
	MockingPluginProtocolsHTTP           MockingPluginProtocols = "http"
	MockingPluginProtocolsHTTPS          MockingPluginProtocols = "https"
	MockingPluginProtocolsTCP            MockingPluginProtocols = "tcp"
	MockingPluginProtocolsTLS            MockingPluginProtocols = "tls"
	MockingPluginProtocolsTLSPassthrough MockingPluginProtocols = "tls_passthrough"
	MockingPluginProtocolsUDP            MockingPluginProtocols = "udp"
	MockingPluginProtocolsWs             MockingPluginProtocols = "ws"
	MockingPluginProtocolsWss            MockingPluginProtocols = "wss"
)

func (e MockingPluginProtocols) ToPointer() *MockingPluginProtocols {
	return &e
}
func (e *MockingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = MockingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MockingPluginProtocols: %v", v)
	}
}

// MockingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type MockingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *MockingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type MockingPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *MockingPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// MockingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type MockingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *MockingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// MockingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type MockingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *MockingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type MockingPlugin struct {
	Config *MockingPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"mocking" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []MockingPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *MockingPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *MockingPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *MockingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *MockingPluginService `json:"service,omitempty"`
}

func (m MockingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MockingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MockingPlugin) GetConfig() *MockingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *MockingPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MockingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *MockingPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MockingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *MockingPlugin) GetName() *string {
	return types.String("mocking")
}

func (o *MockingPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *MockingPlugin) GetProtocols() []MockingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *MockingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *MockingPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *MockingPlugin) GetConsumer() *MockingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *MockingPlugin) GetConsumerGroup() *MockingPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *MockingPlugin) GetRoute() *MockingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *MockingPlugin) GetService() *MockingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

// CreateStatsdAdvancedPluginConsumerIdentifierDefault - The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are `custom_id`, `consumer_id`, `username`.
type CreateStatsdAdvancedPluginConsumerIdentifierDefault string

const (
	CreateStatsdAdvancedPluginConsumerIdentifierDefaultConsumerID CreateStatsdAdvancedPluginConsumerIdentifierDefault = "consumer_id"
	CreateStatsdAdvancedPluginConsumerIdentifierDefaultCustomID   CreateStatsdAdvancedPluginConsumerIdentifierDefault = "custom_id"
	CreateStatsdAdvancedPluginConsumerIdentifierDefaultUsername   CreateStatsdAdvancedPluginConsumerIdentifierDefault = "username"
)

func (e CreateStatsdAdvancedPluginConsumerIdentifierDefault) ToPointer() *CreateStatsdAdvancedPluginConsumerIdentifierDefault {
	return &e
}
func (e *CreateStatsdAdvancedPluginConsumerIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = CreateStatsdAdvancedPluginConsumerIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginConsumerIdentifierDefault: %v", v)
	}
}

type CreateStatsdAdvancedPluginConsumerIdentifier string

const (
	CreateStatsdAdvancedPluginConsumerIdentifierConsumerID CreateStatsdAdvancedPluginConsumerIdentifier = "consumer_id"
	CreateStatsdAdvancedPluginConsumerIdentifierCustomID   CreateStatsdAdvancedPluginConsumerIdentifier = "custom_id"
	CreateStatsdAdvancedPluginConsumerIdentifierUsername   CreateStatsdAdvancedPluginConsumerIdentifier = "username"
)

func (e CreateStatsdAdvancedPluginConsumerIdentifier) ToPointer() *CreateStatsdAdvancedPluginConsumerIdentifier {
	return &e
}
func (e *CreateStatsdAdvancedPluginConsumerIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = CreateStatsdAdvancedPluginConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginConsumerIdentifier: %v", v)
	}
}

type CreateStatsdAdvancedPluginName string

const (
	CreateStatsdAdvancedPluginNameKongLatency                CreateStatsdAdvancedPluginName = "kong_latency"
	CreateStatsdAdvancedPluginNameLatency                    CreateStatsdAdvancedPluginName = "latency"
	CreateStatsdAdvancedPluginNameRequestCount               CreateStatsdAdvancedPluginName = "request_count"
	CreateStatsdAdvancedPluginNameRequestPerUser             CreateStatsdAdvancedPluginName = "request_per_user"
	CreateStatsdAdvancedPluginNameRequestSize                CreateStatsdAdvancedPluginName = "request_size"
	CreateStatsdAdvancedPluginNameResponseSize               CreateStatsdAdvancedPluginName = "response_size"
	CreateStatsdAdvancedPluginNameStatusCount                CreateStatsdAdvancedPluginName = "status_count"
	CreateStatsdAdvancedPluginNameStatusCountPerUser         CreateStatsdAdvancedPluginName = "status_count_per_user"
	CreateStatsdAdvancedPluginNameUniqueUsers                CreateStatsdAdvancedPluginName = "unique_users"
	CreateStatsdAdvancedPluginNameUpstreamLatency            CreateStatsdAdvancedPluginName = "upstream_latency"
	CreateStatsdAdvancedPluginNameStatusCountPerWorkspace    CreateStatsdAdvancedPluginName = "status_count_per_workspace"
	CreateStatsdAdvancedPluginNameStatusCountPerUserPerRoute CreateStatsdAdvancedPluginName = "status_count_per_user_per_route"
	CreateStatsdAdvancedPluginNameShdictUsage                CreateStatsdAdvancedPluginName = "shdict_usage"
	CreateStatsdAdvancedPluginNameCacheDatastoreHitsTotal    CreateStatsdAdvancedPluginName = "cache_datastore_hits_total"
	CreateStatsdAdvancedPluginNameCacheDatastoreMissesTotal  CreateStatsdAdvancedPluginName = "cache_datastore_misses_total"
)

func (e CreateStatsdAdvancedPluginName) ToPointer() *CreateStatsdAdvancedPluginName {
	return &e
}
func (e *CreateStatsdAdvancedPluginName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong_latency":
		fallthrough
	case "latency":
		fallthrough
	case "request_count":
		fallthrough
	case "request_per_user":
		fallthrough
	case "request_size":
		fallthrough
	case "response_size":
		fallthrough
	case "status_count":
		fallthrough
	case "status_count_per_user":
		fallthrough
	case "unique_users":
		fallthrough
	case "upstream_latency":
		fallthrough
	case "status_count_per_workspace":
		fallthrough
	case "status_count_per_user_per_route":
		fallthrough
	case "shdict_usage":
		fallthrough
	case "cache_datastore_hits_total":
		fallthrough
	case "cache_datastore_misses_total":
		*e = CreateStatsdAdvancedPluginName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginName: %v", v)
	}
}

type CreateStatsdAdvancedPluginServiceIdentifier string

const (
	CreateStatsdAdvancedPluginServiceIdentifierServiceID         CreateStatsdAdvancedPluginServiceIdentifier = "service_id"
	CreateStatsdAdvancedPluginServiceIdentifierServiceName       CreateStatsdAdvancedPluginServiceIdentifier = "service_name"
	CreateStatsdAdvancedPluginServiceIdentifierServiceHost       CreateStatsdAdvancedPluginServiceIdentifier = "service_host"
	CreateStatsdAdvancedPluginServiceIdentifierServiceNameOrHost CreateStatsdAdvancedPluginServiceIdentifier = "service_name_or_host"
)

func (e CreateStatsdAdvancedPluginServiceIdentifier) ToPointer() *CreateStatsdAdvancedPluginServiceIdentifier {
	return &e
}
func (e *CreateStatsdAdvancedPluginServiceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = CreateStatsdAdvancedPluginServiceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginServiceIdentifier: %v", v)
	}
}

type CreateStatsdAdvancedPluginStatType string

const (
	CreateStatsdAdvancedPluginStatTypeCounter   CreateStatsdAdvancedPluginStatType = "counter"
	CreateStatsdAdvancedPluginStatTypeGauge     CreateStatsdAdvancedPluginStatType = "gauge"
	CreateStatsdAdvancedPluginStatTypeHistogram CreateStatsdAdvancedPluginStatType = "histogram"
	CreateStatsdAdvancedPluginStatTypeMeter     CreateStatsdAdvancedPluginStatType = "meter"
	CreateStatsdAdvancedPluginStatTypeSet       CreateStatsdAdvancedPluginStatType = "set"
	CreateStatsdAdvancedPluginStatTypeTimer     CreateStatsdAdvancedPluginStatType = "timer"
)

func (e CreateStatsdAdvancedPluginStatType) ToPointer() *CreateStatsdAdvancedPluginStatType {
	return &e
}
func (e *CreateStatsdAdvancedPluginStatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "counter":
		fallthrough
	case "gauge":
		fallthrough
	case "histogram":
		fallthrough
	case "meter":
		fallthrough
	case "set":
		fallthrough
	case "timer":
		*e = CreateStatsdAdvancedPluginStatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginStatType: %v", v)
	}
}

type CreateStatsdAdvancedPluginWorkspaceIdentifier string

const (
	CreateStatsdAdvancedPluginWorkspaceIdentifierWorkspaceID   CreateStatsdAdvancedPluginWorkspaceIdentifier = "workspace_id"
	CreateStatsdAdvancedPluginWorkspaceIdentifierWorkspaceName CreateStatsdAdvancedPluginWorkspaceIdentifier = "workspace_name"
)

func (e CreateStatsdAdvancedPluginWorkspaceIdentifier) ToPointer() *CreateStatsdAdvancedPluginWorkspaceIdentifier {
	return &e
}
func (e *CreateStatsdAdvancedPluginWorkspaceIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = CreateStatsdAdvancedPluginWorkspaceIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginWorkspaceIdentifier: %v", v)
	}
}

type CreateStatsdAdvancedPluginMetrics struct {
	ConsumerIdentifier  *CreateStatsdAdvancedPluginConsumerIdentifier  `json:"consumer_identifier,omitempty"`
	Name                CreateStatsdAdvancedPluginName                 `json:"name"`
	SampleRate          *float64                                       `json:"sample_rate,omitempty"`
	ServiceIdentifier   *CreateStatsdAdvancedPluginServiceIdentifier   `json:"service_identifier,omitempty"`
	StatType            CreateStatsdAdvancedPluginStatType             `json:"stat_type"`
	WorkspaceIdentifier *CreateStatsdAdvancedPluginWorkspaceIdentifier `json:"workspace_identifier,omitempty"`
}

func (o *CreateStatsdAdvancedPluginMetrics) GetConsumerIdentifier() *CreateStatsdAdvancedPluginConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *CreateStatsdAdvancedPluginMetrics) GetName() CreateStatsdAdvancedPluginName {
	if o == nil {
		return CreateStatsdAdvancedPluginName("")
	}
	return o.Name
}

func (o *CreateStatsdAdvancedPluginMetrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *CreateStatsdAdvancedPluginMetrics) GetServiceIdentifier() *CreateStatsdAdvancedPluginServiceIdentifier {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifier
}

func (o *CreateStatsdAdvancedPluginMetrics) GetStatType() CreateStatsdAdvancedPluginStatType {
	if o == nil {
		return CreateStatsdAdvancedPluginStatType("")
	}
	return o.StatType
}

func (o *CreateStatsdAdvancedPluginMetrics) GetWorkspaceIdentifier() *CreateStatsdAdvancedPluginWorkspaceIdentifier {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifier
}

// CreateStatsdAdvancedPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type CreateStatsdAdvancedPluginConcurrencyLimit int64

const (
	CreateStatsdAdvancedPluginConcurrencyLimitMinus1 CreateStatsdAdvancedPluginConcurrencyLimit = -1
	CreateStatsdAdvancedPluginConcurrencyLimitOne    CreateStatsdAdvancedPluginConcurrencyLimit = 1
)

func (e CreateStatsdAdvancedPluginConcurrencyLimit) ToPointer() *CreateStatsdAdvancedPluginConcurrencyLimit {
	return &e
}
func (e *CreateStatsdAdvancedPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = CreateStatsdAdvancedPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginConcurrencyLimit: %v", v)
	}
}

type CreateStatsdAdvancedPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *CreateStatsdAdvancedPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
	// Time in seconds before the initial retry is made for a failing batch.
	InitialRetryDelay *float64 `json:"initial_retry_delay,omitempty"`
	// Maximum number of entries that can be processed at a time.
	MaxBatchSize *int64 `json:"max_batch_size,omitempty"`
	// Maximum number of bytes that can be waiting on a queue, requires string content.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
	MaxCoalescingDelay *float64 `json:"max_coalescing_delay,omitempty"`
	// Maximum number of entries that can be waiting on the queue.
	MaxEntries *int64 `json:"max_entries,omitempty"`
	// Maximum time in seconds between retries, caps exponential backoff.
	MaxRetryDelay *float64 `json:"max_retry_delay,omitempty"`
	// Time in seconds before the queue gives up calling a failed handler for a batch.
	MaxRetryTime *float64 `json:"max_retry_time,omitempty"`
}

func (o *CreateStatsdAdvancedPluginQueue) GetConcurrencyLimit() *CreateStatsdAdvancedPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *CreateStatsdAdvancedPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *CreateStatsdAdvancedPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

// CreateStatsdAdvancedPluginServiceIdentifierDefault - The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are `service_name_or_host`, `service_id`, `service_name`, `service_host`.
type CreateStatsdAdvancedPluginServiceIdentifierDefault string

const (
	CreateStatsdAdvancedPluginServiceIdentifierDefaultServiceID         CreateStatsdAdvancedPluginServiceIdentifierDefault = "service_id"
	CreateStatsdAdvancedPluginServiceIdentifierDefaultServiceName       CreateStatsdAdvancedPluginServiceIdentifierDefault = "service_name"
	CreateStatsdAdvancedPluginServiceIdentifierDefaultServiceHost       CreateStatsdAdvancedPluginServiceIdentifierDefault = "service_host"
	CreateStatsdAdvancedPluginServiceIdentifierDefaultServiceNameOrHost CreateStatsdAdvancedPluginServiceIdentifierDefault = "service_name_or_host"
)

func (e CreateStatsdAdvancedPluginServiceIdentifierDefault) ToPointer() *CreateStatsdAdvancedPluginServiceIdentifierDefault {
	return &e
}
func (e *CreateStatsdAdvancedPluginServiceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "service_id":
		fallthrough
	case "service_name":
		fallthrough
	case "service_host":
		fallthrough
	case "service_name_or_host":
		*e = CreateStatsdAdvancedPluginServiceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginServiceIdentifierDefault: %v", v)
	}
}

// CreateStatsdAdvancedPluginWorkspaceIdentifierDefault - The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are `workspace_id`, `workspace_name`.
type CreateStatsdAdvancedPluginWorkspaceIdentifierDefault string

const (
	CreateStatsdAdvancedPluginWorkspaceIdentifierDefaultWorkspaceID   CreateStatsdAdvancedPluginWorkspaceIdentifierDefault = "workspace_id"
	CreateStatsdAdvancedPluginWorkspaceIdentifierDefaultWorkspaceName CreateStatsdAdvancedPluginWorkspaceIdentifierDefault = "workspace_name"
)

func (e CreateStatsdAdvancedPluginWorkspaceIdentifierDefault) ToPointer() *CreateStatsdAdvancedPluginWorkspaceIdentifierDefault {
	return &e
}
func (e *CreateStatsdAdvancedPluginWorkspaceIdentifierDefault) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "workspace_id":
		fallthrough
	case "workspace_name":
		*e = CreateStatsdAdvancedPluginWorkspaceIdentifierDefault(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginWorkspaceIdentifierDefault: %v", v)
	}
}

type CreateStatsdAdvancedPluginConfig struct {
	// List of status code ranges that are allowed to be logged in metrics.
	AllowStatusCodes []string `json:"allow_status_codes,omitempty"`
	// The default consumer identifier for metrics. This will take effect when a metric's consumer identifier is omitted. Allowed values are `custom_id`, `consumer_id`, `username`.
	ConsumerIdentifierDefault *CreateStatsdAdvancedPluginConsumerIdentifierDefault `json:"consumer_identifier_default,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Include the `hostname` in the `prefix` for each metric name.
	HostnameInPrefix *bool `json:"hostname_in_prefix,omitempty"`
	// List of Metrics to be logged.
	Metrics []CreateStatsdAdvancedPluginMetrics `json:"metrics,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// String to prefix to each metric's name.
	Prefix *string                          `json:"prefix,omitempty"`
	Queue  *CreateStatsdAdvancedPluginQueue `json:"queue,omitempty"`
	// The default service identifier for metrics. This will take effect when a metric's service identifier is omitted. Allowed values are `service_name_or_host`, `service_id`, `service_name`, `service_host`.
	ServiceIdentifierDefault *CreateStatsdAdvancedPluginServiceIdentifierDefault `json:"service_identifier_default,omitempty"`
	// Combine UDP packet up to the size configured. If zero (0), don't combine the UDP packet. Must be a number between 0 and 65507 (inclusive).
	UDPPacketSize *float64 `json:"udp_packet_size,omitempty"`
	// Use TCP instead of UDP.
	UseTCP *bool `json:"use_tcp,omitempty"`
	// The default workspace identifier for metrics. This will take effect when a metric's workspace identifier is omitted. Allowed values are `workspace_id`, `workspace_name`.
	WorkspaceIdentifierDefault *CreateStatsdAdvancedPluginWorkspaceIdentifierDefault `json:"workspace_identifier_default,omitempty"`
}

func (o *CreateStatsdAdvancedPluginConfig) GetAllowStatusCodes() []string {
	if o == nil {
		return nil
	}
	return o.AllowStatusCodes
}

func (o *CreateStatsdAdvancedPluginConfig) GetConsumerIdentifierDefault() *CreateStatsdAdvancedPluginConsumerIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifierDefault
}

func (o *CreateStatsdAdvancedPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateStatsdAdvancedPluginConfig) GetHostnameInPrefix() *bool {
	if o == nil {
		return nil
	}
	return o.HostnameInPrefix
}

func (o *CreateStatsdAdvancedPluginConfig) GetMetrics() []CreateStatsdAdvancedPluginMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *CreateStatsdAdvancedPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateStatsdAdvancedPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateStatsdAdvancedPluginConfig) GetQueue() *CreateStatsdAdvancedPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *CreateStatsdAdvancedPluginConfig) GetServiceIdentifierDefault() *CreateStatsdAdvancedPluginServiceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.ServiceIdentifierDefault
}

func (o *CreateStatsdAdvancedPluginConfig) GetUDPPacketSize() *float64 {
	if o == nil {
		return nil
	}
	return o.UDPPacketSize
}

func (o *CreateStatsdAdvancedPluginConfig) GetUseTCP() *bool {
	if o == nil {
		return nil
	}
	return o.UseTCP
}

func (o *CreateStatsdAdvancedPluginConfig) GetWorkspaceIdentifierDefault() *CreateStatsdAdvancedPluginWorkspaceIdentifierDefault {
	if o == nil {
		return nil
	}
	return o.WorkspaceIdentifierDefault
}

type CreateStatsdAdvancedPluginProtocols string

const (
	CreateStatsdAdvancedPluginProtocolsGrpc           CreateStatsdAdvancedPluginProtocols = "grpc"
	CreateStatsdAdvancedPluginProtocolsGrpcs          CreateStatsdAdvancedPluginProtocols = "grpcs"
	CreateStatsdAdvancedPluginProtocolsHTTP           CreateStatsdAdvancedPluginProtocols = "http"
	CreateStatsdAdvancedPluginProtocolsHTTPS          CreateStatsdAdvancedPluginProtocols = "https"
	CreateStatsdAdvancedPluginProtocolsTCP            CreateStatsdAdvancedPluginProtocols = "tcp"
	CreateStatsdAdvancedPluginProtocolsTLS            CreateStatsdAdvancedPluginProtocols = "tls"
	CreateStatsdAdvancedPluginProtocolsTLSPassthrough CreateStatsdAdvancedPluginProtocols = "tls_passthrough"
	CreateStatsdAdvancedPluginProtocolsUDP            CreateStatsdAdvancedPluginProtocols = "udp"
	CreateStatsdAdvancedPluginProtocolsWs             CreateStatsdAdvancedPluginProtocols = "ws"
	CreateStatsdAdvancedPluginProtocolsWss            CreateStatsdAdvancedPluginProtocols = "wss"
)

func (e CreateStatsdAdvancedPluginProtocols) ToPointer() *CreateStatsdAdvancedPluginProtocols {
	return &e
}
func (e *CreateStatsdAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateStatsdAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateStatsdAdvancedPluginProtocols: %v", v)
	}
}

// CreateStatsdAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateStatsdAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateStatsdAdvancedPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdAdvancedPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateStatsdAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateStatsdAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateStatsdAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateStatsdAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateStatsdAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateStatsdAdvancedPlugin struct {
	Config *CreateStatsdAdvancedPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"statsd-advanced" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateStatsdAdvancedPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateStatsdAdvancedPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateStatsdAdvancedPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateStatsdAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateStatsdAdvancedPluginService `json:"service,omitempty"`
}

func (c CreateStatsdAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateStatsdAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateStatsdAdvancedPlugin) GetConfig() *CreateStatsdAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateStatsdAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateStatsdAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateStatsdAdvancedPlugin) GetName() *string {
	return types.String("statsd-advanced")
}

func (o *CreateStatsdAdvancedPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateStatsdAdvancedPlugin) GetProtocols() []CreateStatsdAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateStatsdAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateStatsdAdvancedPlugin) GetConsumer() *CreateStatsdAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateStatsdAdvancedPlugin) GetConsumerGroup() *CreateStatsdAdvancedPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateStatsdAdvancedPlugin) GetRoute() *CreateStatsdAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateStatsdAdvancedPlugin) GetService() *CreateStatsdAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

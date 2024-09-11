// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type ClientErrorsSeverity string

const (
	ClientErrorsSeverityDebug   ClientErrorsSeverity = "debug"
	ClientErrorsSeverityInfo    ClientErrorsSeverity = "info"
	ClientErrorsSeverityNotice  ClientErrorsSeverity = "notice"
	ClientErrorsSeverityWarning ClientErrorsSeverity = "warning"
	ClientErrorsSeverityErr     ClientErrorsSeverity = "err"
	ClientErrorsSeverityCrit    ClientErrorsSeverity = "crit"
	ClientErrorsSeverityAlert   ClientErrorsSeverity = "alert"
	ClientErrorsSeverityEmerg   ClientErrorsSeverity = "emerg"
)

func (e ClientErrorsSeverity) ToPointer() *ClientErrorsSeverity {
	return &e
}
func (e *ClientErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = ClientErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClientErrorsSeverity: %v", v)
	}
}

type LogLevel string

const (
	LogLevelDebug   LogLevel = "debug"
	LogLevelInfo    LogLevel = "info"
	LogLevelNotice  LogLevel = "notice"
	LogLevelWarning LogLevel = "warning"
	LogLevelErr     LogLevel = "err"
	LogLevelCrit    LogLevel = "crit"
	LogLevelAlert   LogLevel = "alert"
	LogLevelEmerg   LogLevel = "emerg"
)

func (e LogLevel) ToPointer() *LogLevel {
	return &e
}
func (e *LogLevel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = LogLevel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LogLevel: %v", v)
	}
}

type ServerErrorsSeverity string

const (
	ServerErrorsSeverityDebug   ServerErrorsSeverity = "debug"
	ServerErrorsSeverityInfo    ServerErrorsSeverity = "info"
	ServerErrorsSeverityNotice  ServerErrorsSeverity = "notice"
	ServerErrorsSeverityWarning ServerErrorsSeverity = "warning"
	ServerErrorsSeverityErr     ServerErrorsSeverity = "err"
	ServerErrorsSeverityCrit    ServerErrorsSeverity = "crit"
	ServerErrorsSeverityAlert   ServerErrorsSeverity = "alert"
	ServerErrorsSeverityEmerg   ServerErrorsSeverity = "emerg"
)

func (e ServerErrorsSeverity) ToPointer() *ServerErrorsSeverity {
	return &e
}
func (e *ServerErrorsSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = ServerErrorsSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ServerErrorsSeverity: %v", v)
	}
}

type SuccessfulSeverity string

const (
	SuccessfulSeverityDebug   SuccessfulSeverity = "debug"
	SuccessfulSeverityInfo    SuccessfulSeverity = "info"
	SuccessfulSeverityNotice  SuccessfulSeverity = "notice"
	SuccessfulSeverityWarning SuccessfulSeverity = "warning"
	SuccessfulSeverityErr     SuccessfulSeverity = "err"
	SuccessfulSeverityCrit    SuccessfulSeverity = "crit"
	SuccessfulSeverityAlert   SuccessfulSeverity = "alert"
	SuccessfulSeverityEmerg   SuccessfulSeverity = "emerg"
)

func (e SuccessfulSeverity) ToPointer() *SuccessfulSeverity {
	return &e
}
func (e *SuccessfulSeverity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debug":
		fallthrough
	case "info":
		fallthrough
	case "notice":
		fallthrough
	case "warning":
		fallthrough
	case "err":
		fallthrough
	case "crit":
		fallthrough
	case "alert":
		fallthrough
	case "emerg":
		*e = SuccessfulSeverity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SuccessfulSeverity: %v", v)
	}
}

type CreateLogglyPluginConfig struct {
	ClientErrorsSeverity *ClientErrorsSeverity `json:"client_errors_severity,omitempty"`
	// Lua code as a key-value map
	CustomFieldsByLua map[string]any `json:"custom_fields_by_lua,omitempty"`
	// A string representing a host name, such as example.com.
	Host     *string   `json:"host,omitempty"`
	Key      *string   `json:"key,omitempty"`
	LogLevel *LogLevel `json:"log_level,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port                 *int64                `json:"port,omitempty"`
	ServerErrorsSeverity *ServerErrorsSeverity `json:"server_errors_severity,omitempty"`
	SuccessfulSeverity   *SuccessfulSeverity   `json:"successful_severity,omitempty"`
	Tags                 []string              `json:"tags,omitempty"`
	Timeout              *float64              `json:"timeout,omitempty"`
}

func (o *CreateLogglyPluginConfig) GetClientErrorsSeverity() *ClientErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ClientErrorsSeverity
}

func (o *CreateLogglyPluginConfig) GetCustomFieldsByLua() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomFieldsByLua
}

func (o *CreateLogglyPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateLogglyPluginConfig) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *CreateLogglyPluginConfig) GetLogLevel() *LogLevel {
	if o == nil {
		return nil
	}
	return o.LogLevel
}

func (o *CreateLogglyPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateLogglyPluginConfig) GetServerErrorsSeverity() *ServerErrorsSeverity {
	if o == nil {
		return nil
	}
	return o.ServerErrorsSeverity
}

func (o *CreateLogglyPluginConfig) GetSuccessfulSeverity() *SuccessfulSeverity {
	if o == nil {
		return nil
	}
	return o.SuccessfulSeverity
}

func (o *CreateLogglyPluginConfig) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateLogglyPluginConfig) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

type CreateLogglyPluginProtocols string

const (
	CreateLogglyPluginProtocolsGrpc           CreateLogglyPluginProtocols = "grpc"
	CreateLogglyPluginProtocolsGrpcs          CreateLogglyPluginProtocols = "grpcs"
	CreateLogglyPluginProtocolsHTTP           CreateLogglyPluginProtocols = "http"
	CreateLogglyPluginProtocolsHTTPS          CreateLogglyPluginProtocols = "https"
	CreateLogglyPluginProtocolsTCP            CreateLogglyPluginProtocols = "tcp"
	CreateLogglyPluginProtocolsTLS            CreateLogglyPluginProtocols = "tls"
	CreateLogglyPluginProtocolsTLSPassthrough CreateLogglyPluginProtocols = "tls_passthrough"
	CreateLogglyPluginProtocolsUDP            CreateLogglyPluginProtocols = "udp"
	CreateLogglyPluginProtocolsWs             CreateLogglyPluginProtocols = "ws"
	CreateLogglyPluginProtocolsWss            CreateLogglyPluginProtocols = "wss"
)

func (e CreateLogglyPluginProtocols) ToPointer() *CreateLogglyPluginProtocols {
	return &e
}
func (e *CreateLogglyPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateLogglyPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateLogglyPluginProtocols: %v", v)
	}
}

// CreateLogglyPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateLogglyPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateLogglyPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateLogglyPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateLogglyPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateLogglyPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateLogglyPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateLogglyPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateLogglyPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateLogglyPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateLogglyPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateLogglyPlugin struct {
	Config *CreateLogglyPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"loggly" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateLogglyPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateLogglyPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateLogglyPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateLogglyPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateLogglyPluginService `json:"service,omitempty"`
}

func (c CreateLogglyPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateLogglyPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateLogglyPlugin) GetConfig() *CreateLogglyPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateLogglyPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateLogglyPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateLogglyPlugin) GetName() *string {
	return types.String("loggly")
}

func (o *CreateLogglyPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateLogglyPlugin) GetProtocols() []CreateLogglyPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateLogglyPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateLogglyPlugin) GetConsumer() *CreateLogglyPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateLogglyPlugin) GetConsumerGroup() *CreateLogglyPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateLogglyPlugin) GetRoute() *CreateLogglyPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateLogglyPlugin) GetService() *CreateLogglyPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

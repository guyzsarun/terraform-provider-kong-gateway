// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type CreateResponseTransformerAdvancedPluginJSONTypes string

const (
	CreateResponseTransformerAdvancedPluginJSONTypesBoolean CreateResponseTransformerAdvancedPluginJSONTypes = "boolean"
	CreateResponseTransformerAdvancedPluginJSONTypesNumber  CreateResponseTransformerAdvancedPluginJSONTypes = "number"
	CreateResponseTransformerAdvancedPluginJSONTypesString  CreateResponseTransformerAdvancedPluginJSONTypes = "string"
)

func (e CreateResponseTransformerAdvancedPluginJSONTypes) ToPointer() *CreateResponseTransformerAdvancedPluginJSONTypes {
	return &e
}
func (e *CreateResponseTransformerAdvancedPluginJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerAdvancedPluginJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerAdvancedPluginJSONTypes: %v", v)
	}
}

type CreateResponseTransformerAdvancedPluginAdd struct {
	Headers   []string                                           `json:"headers,omitempty"`
	IfStatus  []string                                           `json:"if_status,omitempty"`
	JSON      []string                                           `json:"json,omitempty"`
	JSONTypes []CreateResponseTransformerAdvancedPluginJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginAdd) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerAdvancedPluginAdd) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

func (o *CreateResponseTransformerAdvancedPluginAdd) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerAdvancedPluginAdd) GetJSONTypes() []CreateResponseTransformerAdvancedPluginJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type CreateResponseTransformerAdvancedPluginAllow struct {
	JSON []string `json:"json,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginAllow) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type CreateResponseTransformerAdvancedPluginConfigJSONTypes string

const (
	CreateResponseTransformerAdvancedPluginConfigJSONTypesBoolean CreateResponseTransformerAdvancedPluginConfigJSONTypes = "boolean"
	CreateResponseTransformerAdvancedPluginConfigJSONTypesNumber  CreateResponseTransformerAdvancedPluginConfigJSONTypes = "number"
	CreateResponseTransformerAdvancedPluginConfigJSONTypesString  CreateResponseTransformerAdvancedPluginConfigJSONTypes = "string"
)

func (e CreateResponseTransformerAdvancedPluginConfigJSONTypes) ToPointer() *CreateResponseTransformerAdvancedPluginConfigJSONTypes {
	return &e
}
func (e *CreateResponseTransformerAdvancedPluginConfigJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerAdvancedPluginConfigJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerAdvancedPluginConfigJSONTypes: %v", v)
	}
}

type CreateResponseTransformerAdvancedPluginAppend struct {
	Headers   []string                                                 `json:"headers,omitempty"`
	IfStatus  []string                                                 `json:"if_status,omitempty"`
	JSON      []string                                                 `json:"json,omitempty"`
	JSONTypes []CreateResponseTransformerAdvancedPluginConfigJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginAppend) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerAdvancedPluginAppend) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

func (o *CreateResponseTransformerAdvancedPluginAppend) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerAdvancedPluginAppend) GetJSONTypes() []CreateResponseTransformerAdvancedPluginConfigJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type CreateResponseTransformerAdvancedPluginRemove struct {
	Headers  []string `json:"headers,omitempty"`
	IfStatus []string `json:"if_status,omitempty"`
	JSON     []string `json:"json,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginRemove) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerAdvancedPluginRemove) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

func (o *CreateResponseTransformerAdvancedPluginRemove) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type CreateResponseTransformerAdvancedPluginRename struct {
	Headers  []string `json:"headers,omitempty"`
	IfStatus []string `json:"if_status,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginRename) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerAdvancedPluginRename) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

type CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes string

const (
	CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypesBoolean CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes = "boolean"
	CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypesNumber  CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes = "number"
	CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypesString  CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes = "string"
)

func (e CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes) ToPointer() *CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes {
	return &e
}
func (e *CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "boolean":
		fallthrough
	case "number":
		fallthrough
	case "string":
		*e = CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes: %v", v)
	}
}

type CreateResponseTransformerAdvancedPluginReplace struct {
	// String with which to replace the entire response body.
	Body      *string                                                         `json:"body,omitempty"`
	Headers   []string                                                        `json:"headers,omitempty"`
	IfStatus  []string                                                        `json:"if_status,omitempty"`
	JSON      []string                                                        `json:"json,omitempty"`
	JSONTypes []CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes `json:"json_types,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginReplace) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateResponseTransformerAdvancedPluginReplace) GetHeaders() []string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateResponseTransformerAdvancedPluginReplace) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

func (o *CreateResponseTransformerAdvancedPluginReplace) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

func (o *CreateResponseTransformerAdvancedPluginReplace) GetJSONTypes() []CreateResponseTransformerAdvancedPluginConfigReplaceJSONTypes {
	if o == nil {
		return nil
	}
	return o.JSONTypes
}

type Transform struct {
	Functions []string `json:"functions,omitempty"`
	IfStatus  []string `json:"if_status,omitempty"`
	JSON      []string `json:"json,omitempty"`
}

func (o *Transform) GetFunctions() []string {
	if o == nil {
		return nil
	}
	return o.Functions
}

func (o *Transform) GetIfStatus() []string {
	if o == nil {
		return nil
	}
	return o.IfStatus
}

func (o *Transform) GetJSON() []string {
	if o == nil {
		return nil
	}
	return o.JSON
}

type CreateResponseTransformerAdvancedPluginConfig struct {
	Add    *CreateResponseTransformerAdvancedPluginAdd    `json:"add,omitempty"`
	Allow  *CreateResponseTransformerAdvancedPluginAllow  `json:"allow,omitempty"`
	Append *CreateResponseTransformerAdvancedPluginAppend `json:"append,omitempty"`
	// Whether dots (for example, `customers.info.phone`) should be treated as part of a property name or used to descend into nested JSON objects..
	DotsInKeys *bool                                           `json:"dots_in_keys,omitempty"`
	Remove     *CreateResponseTransformerAdvancedPluginRemove  `json:"remove,omitempty"`
	Rename     *CreateResponseTransformerAdvancedPluginRename  `json:"rename,omitempty"`
	Replace    *CreateResponseTransformerAdvancedPluginReplace `json:"replace,omitempty"`
	Transform  *Transform                                      `json:"transform,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetAdd() *CreateResponseTransformerAdvancedPluginAdd {
	if o == nil {
		return nil
	}
	return o.Add
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetAllow() *CreateResponseTransformerAdvancedPluginAllow {
	if o == nil {
		return nil
	}
	return o.Allow
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetAppend() *CreateResponseTransformerAdvancedPluginAppend {
	if o == nil {
		return nil
	}
	return o.Append
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetDotsInKeys() *bool {
	if o == nil {
		return nil
	}
	return o.DotsInKeys
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetRemove() *CreateResponseTransformerAdvancedPluginRemove {
	if o == nil {
		return nil
	}
	return o.Remove
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetRename() *CreateResponseTransformerAdvancedPluginRename {
	if o == nil {
		return nil
	}
	return o.Rename
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetReplace() *CreateResponseTransformerAdvancedPluginReplace {
	if o == nil {
		return nil
	}
	return o.Replace
}

func (o *CreateResponseTransformerAdvancedPluginConfig) GetTransform() *Transform {
	if o == nil {
		return nil
	}
	return o.Transform
}

type CreateResponseTransformerAdvancedPluginProtocols string

const (
	CreateResponseTransformerAdvancedPluginProtocolsGrpc           CreateResponseTransformerAdvancedPluginProtocols = "grpc"
	CreateResponseTransformerAdvancedPluginProtocolsGrpcs          CreateResponseTransformerAdvancedPluginProtocols = "grpcs"
	CreateResponseTransformerAdvancedPluginProtocolsHTTP           CreateResponseTransformerAdvancedPluginProtocols = "http"
	CreateResponseTransformerAdvancedPluginProtocolsHTTPS          CreateResponseTransformerAdvancedPluginProtocols = "https"
	CreateResponseTransformerAdvancedPluginProtocolsTCP            CreateResponseTransformerAdvancedPluginProtocols = "tcp"
	CreateResponseTransformerAdvancedPluginProtocolsTLS            CreateResponseTransformerAdvancedPluginProtocols = "tls"
	CreateResponseTransformerAdvancedPluginProtocolsTLSPassthrough CreateResponseTransformerAdvancedPluginProtocols = "tls_passthrough"
	CreateResponseTransformerAdvancedPluginProtocolsUDP            CreateResponseTransformerAdvancedPluginProtocols = "udp"
	CreateResponseTransformerAdvancedPluginProtocolsWs             CreateResponseTransformerAdvancedPluginProtocols = "ws"
	CreateResponseTransformerAdvancedPluginProtocolsWss            CreateResponseTransformerAdvancedPluginProtocols = "wss"
)

func (e CreateResponseTransformerAdvancedPluginProtocols) ToPointer() *CreateResponseTransformerAdvancedPluginProtocols {
	return &e
}
func (e *CreateResponseTransformerAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateResponseTransformerAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateResponseTransformerAdvancedPluginProtocols: %v", v)
	}
}

// CreateResponseTransformerAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateResponseTransformerAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseTransformerAdvancedPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseTransformerAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateResponseTransformerAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateResponseTransformerAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateResponseTransformerAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateResponseTransformerAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateResponseTransformerAdvancedPlugin struct {
	Config *CreateResponseTransformerAdvancedPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"response-transformer-advanced" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateResponseTransformerAdvancedPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateResponseTransformerAdvancedPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateResponseTransformerAdvancedPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateResponseTransformerAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateResponseTransformerAdvancedPluginService `json:"service,omitempty"`
}

func (c CreateResponseTransformerAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateResponseTransformerAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateResponseTransformerAdvancedPlugin) GetConfig() *CreateResponseTransformerAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateResponseTransformerAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateResponseTransformerAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateResponseTransformerAdvancedPlugin) GetName() *string {
	return types.String("response-transformer-advanced")
}

func (o *CreateResponseTransformerAdvancedPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateResponseTransformerAdvancedPlugin) GetProtocols() []CreateResponseTransformerAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateResponseTransformerAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateResponseTransformerAdvancedPlugin) GetConsumer() *CreateResponseTransformerAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateResponseTransformerAdvancedPlugin) GetConsumerGroup() *CreateResponseTransformerAdvancedPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateResponseTransformerAdvancedPlugin) GetRoute() *CreateResponseTransformerAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateResponseTransformerAdvancedPlugin) GetService() *CreateResponseTransformerAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

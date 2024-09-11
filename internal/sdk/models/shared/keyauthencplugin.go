// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type KeyAuthEncPluginConfig struct {
	// An optional string (consumer UUID or username) value to use as an “anonymous” consumer if authentication fails. If empty (default null), the request will fail with an authentication failure `4xx`. Note that this value must refer to the consumer `id` or `username` attribute, and **not** its `custom_id`.
	Anonymous *string `json:"anonymous,omitempty"`
	// An optional boolean value telling the plugin to show or hide the credential from the upstream service. If `true`, the plugin strips the credential from the request (i.e., the header, query string, or request body containing the key) before proxying it.
	HideCredentials *bool `json:"hide_credentials,omitempty"`
	// If enabled, the plugin reads the request body (if said request has one and its MIME type is supported) and tries to find the key in it. Supported MIME types: `application/www-form-urlencoded`, `application/json`, and `multipart/form-data`.
	KeyInBody *bool `json:"key_in_body,omitempty"`
	// If enabled (default), the plugin reads the request header and tries to find the key in it.
	KeyInHeader *bool `json:"key_in_header,omitempty"`
	// If enabled (default), the plugin reads the query parameter in the request and tries to find the key in it.
	KeyInQuery *bool `json:"key_in_query,omitempty"`
	// Describes an array of parameter names where the plugin will look for a key. The client must send the authentication key in one of those key names, and the plugin will try to read the credential from a header, request body, or query string parameter with the same name.  Key names may only contain [a-z], [A-Z], [0-9], [_] underscore, and [-] hyphen.
	KeyNames []string `json:"key_names,omitempty"`
	// When authentication fails the plugin sends `WWW-Authenticate` header with `realm` attribute value.
	Realm *string `json:"realm,omitempty"`
	// A boolean value that indicates whether the plugin should run (and try to authenticate) on `OPTIONS` preflight requests. If set to `false`, then `OPTIONS` requests are always allowed.
	RunOnPreflight *bool `json:"run_on_preflight,omitempty"`
}

func (o *KeyAuthEncPluginConfig) GetAnonymous() *string {
	if o == nil {
		return nil
	}
	return o.Anonymous
}

func (o *KeyAuthEncPluginConfig) GetHideCredentials() *bool {
	if o == nil {
		return nil
	}
	return o.HideCredentials
}

func (o *KeyAuthEncPluginConfig) GetKeyInBody() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInBody
}

func (o *KeyAuthEncPluginConfig) GetKeyInHeader() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInHeader
}

func (o *KeyAuthEncPluginConfig) GetKeyInQuery() *bool {
	if o == nil {
		return nil
	}
	return o.KeyInQuery
}

func (o *KeyAuthEncPluginConfig) GetKeyNames() []string {
	if o == nil {
		return nil
	}
	return o.KeyNames
}

func (o *KeyAuthEncPluginConfig) GetRealm() *string {
	if o == nil {
		return nil
	}
	return o.Realm
}

func (o *KeyAuthEncPluginConfig) GetRunOnPreflight() *bool {
	if o == nil {
		return nil
	}
	return o.RunOnPreflight
}

type KeyAuthEncPluginProtocols string

const (
	KeyAuthEncPluginProtocolsGrpc           KeyAuthEncPluginProtocols = "grpc"
	KeyAuthEncPluginProtocolsGrpcs          KeyAuthEncPluginProtocols = "grpcs"
	KeyAuthEncPluginProtocolsHTTP           KeyAuthEncPluginProtocols = "http"
	KeyAuthEncPluginProtocolsHTTPS          KeyAuthEncPluginProtocols = "https"
	KeyAuthEncPluginProtocolsTCP            KeyAuthEncPluginProtocols = "tcp"
	KeyAuthEncPluginProtocolsTLS            KeyAuthEncPluginProtocols = "tls"
	KeyAuthEncPluginProtocolsTLSPassthrough KeyAuthEncPluginProtocols = "tls_passthrough"
	KeyAuthEncPluginProtocolsUDP            KeyAuthEncPluginProtocols = "udp"
	KeyAuthEncPluginProtocolsWs             KeyAuthEncPluginProtocols = "ws"
	KeyAuthEncPluginProtocolsWss            KeyAuthEncPluginProtocols = "wss"
)

func (e KeyAuthEncPluginProtocols) ToPointer() *KeyAuthEncPluginProtocols {
	return &e
}
func (e *KeyAuthEncPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = KeyAuthEncPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for KeyAuthEncPluginProtocols: %v", v)
	}
}

// KeyAuthEncPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type KeyAuthEncPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthEncPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KeyAuthEncPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthEncPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KeyAuthEncPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type KeyAuthEncPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthEncPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// KeyAuthEncPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type KeyAuthEncPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *KeyAuthEncPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type KeyAuthEncPlugin struct {
	Config *KeyAuthEncPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"key-auth-enc" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []KeyAuthEncPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *KeyAuthEncPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *KeyAuthEncPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *KeyAuthEncPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *KeyAuthEncPluginService `json:"service,omitempty"`
}

func (k KeyAuthEncPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(k, "", false)
}

func (k *KeyAuthEncPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &k, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *KeyAuthEncPlugin) GetConfig() *KeyAuthEncPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *KeyAuthEncPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *KeyAuthEncPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *KeyAuthEncPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *KeyAuthEncPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *KeyAuthEncPlugin) GetName() *string {
	return types.String("key-auth-enc")
}

func (o *KeyAuthEncPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *KeyAuthEncPlugin) GetProtocols() []KeyAuthEncPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *KeyAuthEncPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *KeyAuthEncPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *KeyAuthEncPlugin) GetConsumer() *KeyAuthEncPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *KeyAuthEncPlugin) GetConsumerGroup() *KeyAuthEncPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *KeyAuthEncPlugin) GetRoute() *KeyAuthEncPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *KeyAuthEncPlugin) GetService() *KeyAuthEncPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/types"
)

type AiPromptGuardPluginConfig struct {
	// If true, will ignore all previous chat prompts from the conversation history.
	AllowAllConversationHistory *bool `json:"allow_all_conversation_history,omitempty"`
	// Array of valid regex patterns, or valid questions from the 'user' role in chat.
	AllowPatterns []string `json:"allow_patterns,omitempty"`
	// Array of invalid regex patterns, or invalid questions from the 'user' role in chat.
	DenyPatterns []string `json:"deny_patterns,omitempty"`
	// If true, will match all roles in addition to 'user' role in conversation history.
	MatchAllRoles *bool `json:"match_all_roles,omitempty"`
	// max allowed body size allowed to be introspected
	MaxRequestBodySize *int64 `json:"max_request_body_size,omitempty"`
}

func (o *AiPromptGuardPluginConfig) GetAllowAllConversationHistory() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAllConversationHistory
}

func (o *AiPromptGuardPluginConfig) GetAllowPatterns() []string {
	if o == nil {
		return nil
	}
	return o.AllowPatterns
}

func (o *AiPromptGuardPluginConfig) GetDenyPatterns() []string {
	if o == nil {
		return nil
	}
	return o.DenyPatterns
}

func (o *AiPromptGuardPluginConfig) GetMatchAllRoles() *bool {
	if o == nil {
		return nil
	}
	return o.MatchAllRoles
}

func (o *AiPromptGuardPluginConfig) GetMaxRequestBodySize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxRequestBodySize
}

type AiPromptGuardPluginProtocols string

const (
	AiPromptGuardPluginProtocolsGrpc           AiPromptGuardPluginProtocols = "grpc"
	AiPromptGuardPluginProtocolsGrpcs          AiPromptGuardPluginProtocols = "grpcs"
	AiPromptGuardPluginProtocolsHTTP           AiPromptGuardPluginProtocols = "http"
	AiPromptGuardPluginProtocolsHTTPS          AiPromptGuardPluginProtocols = "https"
	AiPromptGuardPluginProtocolsTCP            AiPromptGuardPluginProtocols = "tcp"
	AiPromptGuardPluginProtocolsTLS            AiPromptGuardPluginProtocols = "tls"
	AiPromptGuardPluginProtocolsTLSPassthrough AiPromptGuardPluginProtocols = "tls_passthrough"
	AiPromptGuardPluginProtocolsUDP            AiPromptGuardPluginProtocols = "udp"
	AiPromptGuardPluginProtocolsWs             AiPromptGuardPluginProtocols = "ws"
	AiPromptGuardPluginProtocolsWss            AiPromptGuardPluginProtocols = "wss"
)

func (e AiPromptGuardPluginProtocols) ToPointer() *AiPromptGuardPluginProtocols {
	return &e
}
func (e *AiPromptGuardPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = AiPromptGuardPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AiPromptGuardPluginProtocols: %v", v)
	}
}

// AiPromptGuardPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type AiPromptGuardPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptGuardPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptGuardPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type AiPromptGuardPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// AiPromptGuardPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type AiPromptGuardPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *AiPromptGuardPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type AiPromptGuardPlugin struct {
	Config *AiPromptGuardPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool   `json:"enabled,omitempty"`
	ID           *string `json:"id,omitempty"`
	InstanceName *string `json:"instance_name,omitempty"`
	name         *string `const:"ai-prompt-guard" json:"name,omitempty"`
	Ordering     any     `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []AiPromptGuardPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *AiPromptGuardPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *AiPromptGuardPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *AiPromptGuardPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *AiPromptGuardPluginService `json:"service,omitempty"`
}

func (a AiPromptGuardPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AiPromptGuardPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AiPromptGuardPlugin) GetConfig() *AiPromptGuardPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *AiPromptGuardPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AiPromptGuardPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *AiPromptGuardPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AiPromptGuardPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *AiPromptGuardPlugin) GetName() *string {
	return types.String("ai-prompt-guard")
}

func (o *AiPromptGuardPlugin) GetOrdering() any {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *AiPromptGuardPlugin) GetProtocols() []AiPromptGuardPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *AiPromptGuardPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AiPromptGuardPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AiPromptGuardPlugin) GetConsumer() *AiPromptGuardPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *AiPromptGuardPlugin) GetConsumerGroup() *AiPromptGuardPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *AiPromptGuardPlugin) GetRoute() *AiPromptGuardPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *AiPromptGuardPlugin) GetService() *AiPromptGuardPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

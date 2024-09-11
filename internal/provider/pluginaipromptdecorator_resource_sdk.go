// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginAiPromptDecoratorResourceModel) ToSharedCreateAiPromptDecoratorPlugin() *shared.CreateAiPromptDecoratorPlugin {
	var config *shared.CreateAiPromptDecoratorPluginConfig
	if r.Config != nil {
		maxRequestBodySize := new(int64)
		if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
			*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
		} else {
			maxRequestBodySize = nil
		}
		var prompts *shared.Prompts
		if r.Config.Prompts != nil {
			var append1 []shared.CreateAiPromptDecoratorPluginAppend = []shared.CreateAiPromptDecoratorPluginAppend{}
			for _, appendItem := range r.Config.Prompts.Append {
				var content string
				content = appendItem.Content.ValueString()

				role := new(shared.Role)
				if !appendItem.Role.IsUnknown() && !appendItem.Role.IsNull() {
					*role = shared.Role(appendItem.Role.ValueString())
				} else {
					role = nil
				}
				append1 = append(append1, shared.CreateAiPromptDecoratorPluginAppend{
					Content: content,
					Role:    role,
				})
			}
			var prepend []shared.Prepend = []shared.Prepend{}
			for _, prependItem := range r.Config.Prompts.Prepend {
				var content1 string
				content1 = prependItem.Content.ValueString()

				role1 := new(shared.CreateAiPromptDecoratorPluginRole)
				if !prependItem.Role.IsUnknown() && !prependItem.Role.IsNull() {
					*role1 = shared.CreateAiPromptDecoratorPluginRole(prependItem.Role.ValueString())
				} else {
					role1 = nil
				}
				prepend = append(prepend, shared.Prepend{
					Content: content1,
					Role:    role1,
				})
			}
			prompts = &shared.Prompts{
				Append:  append1,
				Prepend: prepend,
			}
		}
		config = &shared.CreateAiPromptDecoratorPluginConfig{
			MaxRequestBodySize: maxRequestBodySize,
			Prompts:            prompts,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering interface{}
	if !r.Ordering.IsUnknown() && !r.Ordering.IsNull() {
		_ = json.Unmarshal([]byte(r.Ordering.ValueString()), &ordering)
	}
	var protocols []shared.CreateAiPromptDecoratorPluginProtocols = []shared.CreateAiPromptDecoratorPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAiPromptDecoratorPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAiPromptDecoratorPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAiPromptDecoratorPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAiPromptDecoratorPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAiPromptDecoratorPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAiPromptDecoratorPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAiPromptDecoratorPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAiPromptDecoratorPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAiPromptDecoratorPluginService{
			ID: id3,
		}
	}
	out := shared.CreateAiPromptDecoratorPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *PluginAiPromptDecoratorResourceModel) RefreshFromSharedAiPromptDecoratorPlugin(resp *shared.AiPromptDecoratorPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAiPromptDecoratorPluginConfig{}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			if resp.Config.Prompts == nil {
				r.Config.Prompts = nil
			} else {
				r.Config.Prompts = &tfTypes.Prompts{}
				r.Config.Prompts.Append = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Append) > len(resp.Config.Prompts.Append) {
					r.Config.Prompts.Append = r.Config.Prompts.Append[:len(resp.Config.Prompts.Append)]
				}
				for appendCount, appendItem := range resp.Config.Prompts.Append {
					var append2 tfTypes.AiPromptDecoratorPluginAppend
					append2.Content = types.StringValue(appendItem.Content)
					if appendItem.Role != nil {
						append2.Role = types.StringValue(string(*appendItem.Role))
					} else {
						append2.Role = types.StringNull()
					}
					if appendCount+1 > len(r.Config.Prompts.Append) {
						r.Config.Prompts.Append = append(r.Config.Prompts.Append, append2)
					} else {
						r.Config.Prompts.Append[appendCount].Content = append2.Content
						r.Config.Prompts.Append[appendCount].Role = append2.Role
					}
				}
				r.Config.Prompts.Prepend = []tfTypes.AiPromptDecoratorPluginAppend{}
				if len(r.Config.Prompts.Prepend) > len(resp.Config.Prompts.Prepend) {
					r.Config.Prompts.Prepend = r.Config.Prompts.Prepend[:len(resp.Config.Prompts.Prepend)]
				}
				for prependCount, prependItem := range resp.Config.Prompts.Prepend {
					var prepend1 tfTypes.AiPromptDecoratorPluginAppend
					prepend1.Content = types.StringValue(prependItem.Content)
					if prependItem.Role != nil {
						prepend1.Role = types.StringValue(string(*prependItem.Role))
					} else {
						prepend1.Role = types.StringNull()
					}
					if prependCount+1 > len(r.Config.Prompts.Prepend) {
						r.Config.Prompts.Prepend = append(r.Config.Prompts.Prepend, prepend1)
					} else {
						r.Config.Prompts.Prepend[prependCount].Content = prepend1.Content
						r.Config.Prompts.Prepend[prependCount].Role = prepend1.Role
					}
				}
			}
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = types.StringNull()
		} else {
			orderingResult, _ := json.Marshal(resp.Ordering)
			r.Ordering = types.StringValue(string(orderingResult))
		}
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}

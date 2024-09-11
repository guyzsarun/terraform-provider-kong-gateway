// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginAiPromptTemplateResourceModel) ToSharedCreateAiPromptTemplatePlugin() *shared.CreateAiPromptTemplatePlugin {
	var config *shared.CreateAiPromptTemplatePluginConfig
	if r.Config != nil {
		allowUntemplatedRequests := new(bool)
		if !r.Config.AllowUntemplatedRequests.IsUnknown() && !r.Config.AllowUntemplatedRequests.IsNull() {
			*allowUntemplatedRequests = r.Config.AllowUntemplatedRequests.ValueBool()
		} else {
			allowUntemplatedRequests = nil
		}
		logOriginalRequest := new(bool)
		if !r.Config.LogOriginalRequest.IsUnknown() && !r.Config.LogOriginalRequest.IsNull() {
			*logOriginalRequest = r.Config.LogOriginalRequest.ValueBool()
		} else {
			logOriginalRequest = nil
		}
		maxRequestBodySize := new(int64)
		if !r.Config.MaxRequestBodySize.IsUnknown() && !r.Config.MaxRequestBodySize.IsNull() {
			*maxRequestBodySize = r.Config.MaxRequestBodySize.ValueInt64()
		} else {
			maxRequestBodySize = nil
		}
		var templates []shared.Templates = []shared.Templates{}
		for _, templatesItem := range r.Config.Templates {
			var name string
			name = templatesItem.Name.ValueString()

			var template string
			template = templatesItem.Template.ValueString()

			templates = append(templates, shared.Templates{
				Name:     name,
				Template: template,
			})
		}
		config = &shared.CreateAiPromptTemplatePluginConfig{
			AllowUntemplatedRequests: allowUntemplatedRequests,
			LogOriginalRequest:       logOriginalRequest,
			MaxRequestBodySize:       maxRequestBodySize,
			Templates:                templates,
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
	var protocols []shared.CreateAiPromptTemplatePluginProtocols = []shared.CreateAiPromptTemplatePluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAiPromptTemplatePluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAiPromptTemplatePluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAiPromptTemplatePluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAiPromptTemplatePluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAiPromptTemplatePluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAiPromptTemplatePluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAiPromptTemplatePluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAiPromptTemplatePluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAiPromptTemplatePluginService{
			ID: id3,
		}
	}
	out := shared.CreateAiPromptTemplatePlugin{
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

func (r *PluginAiPromptTemplateResourceModel) RefreshFromSharedAiPromptTemplatePlugin(resp *shared.AiPromptTemplatePlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAiPromptTemplatePluginConfig{}
			r.Config.AllowUntemplatedRequests = types.BoolPointerValue(resp.Config.AllowUntemplatedRequests)
			r.Config.LogOriginalRequest = types.BoolPointerValue(resp.Config.LogOriginalRequest)
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			r.Config.Templates = []tfTypes.Templates{}
			if len(r.Config.Templates) > len(resp.Config.Templates) {
				r.Config.Templates = r.Config.Templates[:len(resp.Config.Templates)]
			}
			for templatesCount, templatesItem := range resp.Config.Templates {
				var templates1 tfTypes.Templates
				templates1.Name = types.StringValue(templatesItem.Name)
				templates1.Template = types.StringValue(templatesItem.Template)
				if templatesCount+1 > len(r.Config.Templates) {
					r.Config.Templates = append(r.Config.Templates, templates1)
				} else {
					r.Config.Templates[templatesCount].Name = templates1.Name
					r.Config.Templates[templatesCount].Template = templates1.Template
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

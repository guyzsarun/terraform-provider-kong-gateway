// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *PluginJqResourceModel) ToSharedCreateJqPlugin() *shared.CreateJqPlugin {
	var config *shared.CreateJqPluginConfig
	if r.Config != nil {
		var requestIfMediaType []string = []string{}
		for _, requestIfMediaTypeItem := range r.Config.RequestIfMediaType {
			requestIfMediaType = append(requestIfMediaType, requestIfMediaTypeItem.ValueString())
		}
		requestJqProgram := new(string)
		if !r.Config.RequestJqProgram.IsUnknown() && !r.Config.RequestJqProgram.IsNull() {
			*requestJqProgram = r.Config.RequestJqProgram.ValueString()
		} else {
			requestJqProgram = nil
		}
		var requestJqProgramOptions *shared.RequestJqProgramOptions
		if r.Config.RequestJqProgramOptions != nil {
			asciiOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.ASCIIOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.ASCIIOutput.IsNull() {
				*asciiOutput = r.Config.RequestJqProgramOptions.ASCIIOutput.ValueBool()
			} else {
				asciiOutput = nil
			}
			compactOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.CompactOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.CompactOutput.IsNull() {
				*compactOutput = r.Config.RequestJqProgramOptions.CompactOutput.ValueBool()
			} else {
				compactOutput = nil
			}
			joinOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.JoinOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.JoinOutput.IsNull() {
				*joinOutput = r.Config.RequestJqProgramOptions.JoinOutput.ValueBool()
			} else {
				joinOutput = nil
			}
			rawOutput := new(bool)
			if !r.Config.RequestJqProgramOptions.RawOutput.IsUnknown() && !r.Config.RequestJqProgramOptions.RawOutput.IsNull() {
				*rawOutput = r.Config.RequestJqProgramOptions.RawOutput.ValueBool()
			} else {
				rawOutput = nil
			}
			sortKeys := new(bool)
			if !r.Config.RequestJqProgramOptions.SortKeys.IsUnknown() && !r.Config.RequestJqProgramOptions.SortKeys.IsNull() {
				*sortKeys = r.Config.RequestJqProgramOptions.SortKeys.ValueBool()
			} else {
				sortKeys = nil
			}
			requestJqProgramOptions = &shared.RequestJqProgramOptions{
				ASCIIOutput:   asciiOutput,
				CompactOutput: compactOutput,
				JoinOutput:    joinOutput,
				RawOutput:     rawOutput,
				SortKeys:      sortKeys,
			}
		}
		var responseIfMediaType []string = []string{}
		for _, responseIfMediaTypeItem := range r.Config.ResponseIfMediaType {
			responseIfMediaType = append(responseIfMediaType, responseIfMediaTypeItem.ValueString())
		}
		var responseIfStatusCode []int64 = []int64{}
		for _, responseIfStatusCodeItem := range r.Config.ResponseIfStatusCode {
			responseIfStatusCode = append(responseIfStatusCode, responseIfStatusCodeItem.ValueInt64())
		}
		responseJqProgram := new(string)
		if !r.Config.ResponseJqProgram.IsUnknown() && !r.Config.ResponseJqProgram.IsNull() {
			*responseJqProgram = r.Config.ResponseJqProgram.ValueString()
		} else {
			responseJqProgram = nil
		}
		var responseJqProgramOptions *shared.ResponseJqProgramOptions
		if r.Config.ResponseJqProgramOptions != nil {
			asciiOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.ASCIIOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.ASCIIOutput.IsNull() {
				*asciiOutput1 = r.Config.ResponseJqProgramOptions.ASCIIOutput.ValueBool()
			} else {
				asciiOutput1 = nil
			}
			compactOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.CompactOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.CompactOutput.IsNull() {
				*compactOutput1 = r.Config.ResponseJqProgramOptions.CompactOutput.ValueBool()
			} else {
				compactOutput1 = nil
			}
			joinOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.JoinOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.JoinOutput.IsNull() {
				*joinOutput1 = r.Config.ResponseJqProgramOptions.JoinOutput.ValueBool()
			} else {
				joinOutput1 = nil
			}
			rawOutput1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.RawOutput.IsUnknown() && !r.Config.ResponseJqProgramOptions.RawOutput.IsNull() {
				*rawOutput1 = r.Config.ResponseJqProgramOptions.RawOutput.ValueBool()
			} else {
				rawOutput1 = nil
			}
			sortKeys1 := new(bool)
			if !r.Config.ResponseJqProgramOptions.SortKeys.IsUnknown() && !r.Config.ResponseJqProgramOptions.SortKeys.IsNull() {
				*sortKeys1 = r.Config.ResponseJqProgramOptions.SortKeys.ValueBool()
			} else {
				sortKeys1 = nil
			}
			responseJqProgramOptions = &shared.ResponseJqProgramOptions{
				ASCIIOutput:   asciiOutput1,
				CompactOutput: compactOutput1,
				JoinOutput:    joinOutput1,
				RawOutput:     rawOutput1,
				SortKeys:      sortKeys1,
			}
		}
		config = &shared.CreateJqPluginConfig{
			RequestIfMediaType:       requestIfMediaType,
			RequestJqProgram:         requestJqProgram,
			RequestJqProgramOptions:  requestJqProgramOptions,
			ResponseIfMediaType:      responseIfMediaType,
			ResponseIfStatusCode:     responseIfStatusCode,
			ResponseJqProgram:        responseJqProgram,
			ResponseJqProgramOptions: responseJqProgramOptions,
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
	var protocols []shared.CreateJqPluginProtocols = []shared.CreateJqPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateJqPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateJqPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateJqPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateJqPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateJqPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateJqPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateJqPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateJqPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateJqPluginService{
			ID: id3,
		}
	}
	out := shared.CreateJqPlugin{
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

func (r *PluginJqResourceModel) RefreshFromSharedJqPlugin(resp *shared.JqPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateJqPluginConfig{}
			r.Config.RequestIfMediaType = []types.String{}
			for _, v := range resp.Config.RequestIfMediaType {
				r.Config.RequestIfMediaType = append(r.Config.RequestIfMediaType, types.StringValue(v))
			}
			r.Config.RequestJqProgram = types.StringPointerValue(resp.Config.RequestJqProgram)
			if resp.Config.RequestJqProgramOptions == nil {
				r.Config.RequestJqProgramOptions = nil
			} else {
				r.Config.RequestJqProgramOptions = &tfTypes.RequestJqProgramOptions{}
				r.Config.RequestJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.ASCIIOutput)
				r.Config.RequestJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.CompactOutput)
				r.Config.RequestJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.JoinOutput)
				r.Config.RequestJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.RawOutput)
				r.Config.RequestJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.RequestJqProgramOptions.SortKeys)
			}
			r.Config.ResponseIfMediaType = []types.String{}
			for _, v := range resp.Config.ResponseIfMediaType {
				r.Config.ResponseIfMediaType = append(r.Config.ResponseIfMediaType, types.StringValue(v))
			}
			r.Config.ResponseIfStatusCode = []types.Int64{}
			for _, v := range resp.Config.ResponseIfStatusCode {
				r.Config.ResponseIfStatusCode = append(r.Config.ResponseIfStatusCode, types.Int64Value(v))
			}
			r.Config.ResponseJqProgram = types.StringPointerValue(resp.Config.ResponseJqProgram)
			if resp.Config.ResponseJqProgramOptions == nil {
				r.Config.ResponseJqProgramOptions = nil
			} else {
				r.Config.ResponseJqProgramOptions = &tfTypes.RequestJqProgramOptions{}
				r.Config.ResponseJqProgramOptions.ASCIIOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.ASCIIOutput)
				r.Config.ResponseJqProgramOptions.CompactOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.CompactOutput)
				r.Config.ResponseJqProgramOptions.JoinOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.JoinOutput)
				r.Config.ResponseJqProgramOptions.RawOutput = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.RawOutput)
				r.Config.ResponseJqProgramOptions.SortKeys = types.BoolPointerValue(resp.Config.ResponseJqProgramOptions.SortKeys)
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

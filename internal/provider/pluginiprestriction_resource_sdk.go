// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginIPRestrictionResourceModel) ToSharedCreateIPRestrictionPlugin() *shared.CreateIPRestrictionPlugin {
	var config *shared.CreateIPRestrictionPluginConfig
	if r.Config != nil {
		var allow []string = []string{}
		for _, allowItem := range r.Config.Allow {
			allow = append(allow, allowItem.ValueString())
		}
		var deny []string = []string{}
		for _, denyItem := range r.Config.Deny {
			deny = append(deny, denyItem.ValueString())
		}
		message := new(string)
		if !r.Config.Message.IsUnknown() && !r.Config.Message.IsNull() {
			*message = r.Config.Message.ValueString()
		} else {
			message = nil
		}
		status := new(float64)
		if !r.Config.Status.IsUnknown() && !r.Config.Status.IsNull() {
			*status, _ = r.Config.Status.ValueBigFloat().Float64()
		} else {
			status = nil
		}
		config = &shared.CreateIPRestrictionPluginConfig{
			Allow:   allow,
			Deny:    deny,
			Message: message,
			Status:  status,
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
	var protocols []shared.CreateIPRestrictionPluginProtocols = []shared.CreateIPRestrictionPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateIPRestrictionPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateIPRestrictionPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateIPRestrictionPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateIPRestrictionPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateIPRestrictionPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateIPRestrictionPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateIPRestrictionPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateIPRestrictionPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateIPRestrictionPluginService{
			ID: id3,
		}
	}
	out := shared.CreateIPRestrictionPlugin{
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

func (r *PluginIPRestrictionResourceModel) RefreshFromSharedIPRestrictionPlugin(resp *shared.IPRestrictionPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateIPRestrictionPluginConfig{}
			r.Config.Allow = []types.String{}
			for _, v := range resp.Config.Allow {
				r.Config.Allow = append(r.Config.Allow, types.StringValue(v))
			}
			r.Config.Deny = []types.String{}
			for _, v := range resp.Config.Deny {
				r.Config.Deny = append(r.Config.Deny, types.StringValue(v))
			}
			r.Config.Message = types.StringPointerValue(resp.Config.Message)
			if resp.Config.Status != nil {
				r.Config.Status = types.NumberValue(big.NewFloat(float64(*resp.Config.Status)))
			} else {
				r.Config.Status = types.NumberNull()
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

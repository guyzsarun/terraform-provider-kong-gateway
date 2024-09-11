// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginHmacAuthDataSourceModel) RefreshFromSharedHmacAuthPlugin(resp *shared.HmacAuthPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateHmacAuthPluginConfig{}
			r.Config.Algorithms = []types.String{}
			for _, v := range resp.Config.Algorithms {
				r.Config.Algorithms = append(r.Config.Algorithms, types.StringValue(string(v)))
			}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			if resp.Config.ClockSkew != nil {
				r.Config.ClockSkew = types.NumberValue(big.NewFloat(float64(*resp.Config.ClockSkew)))
			} else {
				r.Config.ClockSkew = types.NumberNull()
			}
			r.Config.EnforceHeaders = []types.String{}
			for _, v := range resp.Config.EnforceHeaders {
				r.Config.EnforceHeaders = append(r.Config.EnforceHeaders, types.StringValue(v))
			}
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.ValidateRequestBody = types.BoolPointerValue(resp.Config.ValidateRequestBody)
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

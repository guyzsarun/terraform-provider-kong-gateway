// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginJWTDataSourceModel) RefreshFromSharedJWTPlugin(resp *shared.JWTPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateJWTPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.ClaimsToVerify = []types.String{}
			for _, v := range resp.Config.ClaimsToVerify {
				r.Config.ClaimsToVerify = append(r.Config.ClaimsToVerify, types.StringValue(string(v)))
			}
			r.Config.CookieNames = []types.String{}
			for _, v := range resp.Config.CookieNames {
				r.Config.CookieNames = append(r.Config.CookieNames, types.StringValue(v))
			}
			r.Config.HeaderNames = []types.String{}
			for _, v := range resp.Config.HeaderNames {
				r.Config.HeaderNames = append(r.Config.HeaderNames, types.StringValue(v))
			}
			r.Config.KeyClaimName = types.StringPointerValue(resp.Config.KeyClaimName)
			if resp.Config.MaximumExpiration != nil {
				r.Config.MaximumExpiration = types.NumberValue(big.NewFloat(float64(*resp.Config.MaximumExpiration)))
			} else {
				r.Config.MaximumExpiration = types.NumberNull()
			}
			r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
			r.Config.SecretIsBase64 = types.BoolPointerValue(resp.Config.SecretIsBase64)
			r.Config.URIParamNames = []types.String{}
			for _, v := range resp.Config.URIParamNames {
				r.Config.URIParamNames = append(r.Config.URIParamNames, types.StringValue(v))
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

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
)

func (r *VaultResourceModel) ToSharedVaultInput() *shared.VaultInput {
	var config interface{}
	if !r.Config.IsUnknown() && !r.Config.IsNull() {
		_ = json.Unmarshal([]byte(r.Config.ValueString()), &config)
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	prefix := new(string)
	if !r.Prefix.IsUnknown() && !r.Prefix.IsNull() {
		*prefix = r.Prefix.ValueString()
	} else {
		prefix = nil
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	out := shared.VaultInput{
		Config:      config,
		Description: description,
		Name:        name,
		Prefix:      prefix,
		Tags:        tags,
	}
	return &out
}

func (r *VaultResourceModel) RefreshFromSharedVault(resp *shared.Vault) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = types.StringNull()
		} else {
			configResult, _ := json.Marshal(resp.Config)
			r.Config = types.StringValue(string(configResult))
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Description = types.StringPointerValue(resp.Description)
		r.ID = types.StringPointerValue(resp.ID)
		r.Name = types.StringPointerValue(resp.Name)
		r.Prefix = types.StringPointerValue(resp.Prefix)
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}

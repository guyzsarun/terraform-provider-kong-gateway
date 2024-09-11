// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/shared"
	"math/big"
)

func (r *PluginAiProxyDataSourceModel) RefreshFromSharedAiProxyPlugin(resp *shared.AiProxyPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAiProxyPluginConfig{}
			if resp.Config.Auth == nil {
				r.Config.Auth = nil
			} else {
				r.Config.Auth = &tfTypes.Auth{}
				r.Config.Auth.AllowOverride = types.BoolPointerValue(resp.Config.Auth.AllowOverride)
				r.Config.Auth.AwsAccessKeyID = types.StringPointerValue(resp.Config.Auth.AwsAccessKeyID)
				r.Config.Auth.AwsSecretAccessKey = types.StringPointerValue(resp.Config.Auth.AwsSecretAccessKey)
				r.Config.Auth.AzureClientID = types.StringPointerValue(resp.Config.Auth.AzureClientID)
				r.Config.Auth.AzureClientSecret = types.StringPointerValue(resp.Config.Auth.AzureClientSecret)
				r.Config.Auth.AzureTenantID = types.StringPointerValue(resp.Config.Auth.AzureTenantID)
				r.Config.Auth.AzureUseManagedIdentity = types.BoolPointerValue(resp.Config.Auth.AzureUseManagedIdentity)
				r.Config.Auth.GcpServiceAccountJSON = types.StringPointerValue(resp.Config.Auth.GcpServiceAccountJSON)
				r.Config.Auth.GcpUseServiceAccount = types.BoolPointerValue(resp.Config.Auth.GcpUseServiceAccount)
				r.Config.Auth.HeaderName = types.StringPointerValue(resp.Config.Auth.HeaderName)
				r.Config.Auth.HeaderValue = types.StringPointerValue(resp.Config.Auth.HeaderValue)
				if resp.Config.Auth.ParamLocation != nil {
					r.Config.Auth.ParamLocation = types.StringValue(string(*resp.Config.Auth.ParamLocation))
				} else {
					r.Config.Auth.ParamLocation = types.StringNull()
				}
				r.Config.Auth.ParamName = types.StringPointerValue(resp.Config.Auth.ParamName)
				r.Config.Auth.ParamValue = types.StringPointerValue(resp.Config.Auth.ParamValue)
			}
			if resp.Config.Logging == nil {
				r.Config.Logging = nil
			} else {
				r.Config.Logging = &tfTypes.Logging{}
				r.Config.Logging.LogPayloads = types.BoolPointerValue(resp.Config.Logging.LogPayloads)
				r.Config.Logging.LogStatistics = types.BoolPointerValue(resp.Config.Logging.LogStatistics)
			}
			r.Config.MaxRequestBodySize = types.Int64PointerValue(resp.Config.MaxRequestBodySize)
			if resp.Config.Model == nil {
				r.Config.Model = nil
			} else {
				r.Config.Model = &tfTypes.Model{}
				r.Config.Model.Name = types.StringPointerValue(resp.Config.Model.Name)
				if resp.Config.Model.Options == nil {
					r.Config.Model.Options = nil
				} else {
					r.Config.Model.Options = &tfTypes.OptionsObj{}
					r.Config.Model.Options.AnthropicVersion = types.StringPointerValue(resp.Config.Model.Options.AnthropicVersion)
					r.Config.Model.Options.AzureAPIVersion = types.StringPointerValue(resp.Config.Model.Options.AzureAPIVersion)
					r.Config.Model.Options.AzureDeploymentID = types.StringPointerValue(resp.Config.Model.Options.AzureDeploymentID)
					r.Config.Model.Options.AzureInstance = types.StringPointerValue(resp.Config.Model.Options.AzureInstance)
					if resp.Config.Model.Options.Bedrock == nil {
						r.Config.Model.Options.Bedrock = nil
					} else {
						r.Config.Model.Options.Bedrock = &tfTypes.Bedrock{}
						r.Config.Model.Options.Bedrock.AwsRegion = types.StringPointerValue(resp.Config.Model.Options.Bedrock.AwsRegion)
					}
					if resp.Config.Model.Options.Gemini == nil {
						r.Config.Model.Options.Gemini = nil
					} else {
						r.Config.Model.Options.Gemini = &tfTypes.Gemini{}
						r.Config.Model.Options.Gemini.APIEndpoint = types.StringPointerValue(resp.Config.Model.Options.Gemini.APIEndpoint)
						r.Config.Model.Options.Gemini.LocationID = types.StringPointerValue(resp.Config.Model.Options.Gemini.LocationID)
						r.Config.Model.Options.Gemini.ProjectID = types.StringPointerValue(resp.Config.Model.Options.Gemini.ProjectID)
					}
					if resp.Config.Model.Options.InputCost != nil {
						r.Config.Model.Options.InputCost = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.InputCost)))
					} else {
						r.Config.Model.Options.InputCost = types.NumberNull()
					}
					if resp.Config.Model.Options.Llama2Format != nil {
						r.Config.Model.Options.Llama2Format = types.StringValue(string(*resp.Config.Model.Options.Llama2Format))
					} else {
						r.Config.Model.Options.Llama2Format = types.StringNull()
					}
					r.Config.Model.Options.MaxTokens = types.Int64PointerValue(resp.Config.Model.Options.MaxTokens)
					if resp.Config.Model.Options.MistralFormat != nil {
						r.Config.Model.Options.MistralFormat = types.StringValue(string(*resp.Config.Model.Options.MistralFormat))
					} else {
						r.Config.Model.Options.MistralFormat = types.StringNull()
					}
					if resp.Config.Model.Options.OutputCost != nil {
						r.Config.Model.Options.OutputCost = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.OutputCost)))
					} else {
						r.Config.Model.Options.OutputCost = types.NumberNull()
					}
					if resp.Config.Model.Options.Temperature != nil {
						r.Config.Model.Options.Temperature = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.Temperature)))
					} else {
						r.Config.Model.Options.Temperature = types.NumberNull()
					}
					r.Config.Model.Options.TopK = types.Int64PointerValue(resp.Config.Model.Options.TopK)
					if resp.Config.Model.Options.TopP != nil {
						r.Config.Model.Options.TopP = types.NumberValue(big.NewFloat(float64(*resp.Config.Model.Options.TopP)))
					} else {
						r.Config.Model.Options.TopP = types.NumberNull()
					}
					r.Config.Model.Options.UpstreamPath = types.StringPointerValue(resp.Config.Model.Options.UpstreamPath)
					r.Config.Model.Options.UpstreamURL = types.StringPointerValue(resp.Config.Model.Options.UpstreamURL)
				}
				if resp.Config.Model.Provider != nil {
					r.Config.Model.Provider = types.StringValue(string(*resp.Config.Model.Provider))
				} else {
					r.Config.Model.Provider = types.StringNull()
				}
			}
			r.Config.ModelNameHeader = types.BoolPointerValue(resp.Config.ModelNameHeader)
			if resp.Config.ResponseStreaming != nil {
				r.Config.ResponseStreaming = types.StringValue(string(*resp.Config.ResponseStreaming))
			} else {
				r.Config.ResponseStreaming = types.StringNull()
			}
			if resp.Config.RouteType != nil {
				r.Config.RouteType = types.StringValue(string(*resp.Config.RouteType))
			} else {
				r.Config.RouteType = types.StringNull()
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

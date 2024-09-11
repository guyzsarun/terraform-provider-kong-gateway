// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateStatsdPluginConfig struct {
	AllowStatusCodes           []types.String        `tfsdk:"allow_status_codes"`
	ConsumerIdentifierDefault  types.String          `tfsdk:"consumer_identifier_default"`
	FlushTimeout               types.Number          `tfsdk:"flush_timeout"`
	Host                       types.String          `tfsdk:"host"`
	HostnameInPrefix           types.Bool            `tfsdk:"hostname_in_prefix"`
	Metrics                    []StatsdPluginMetrics `tfsdk:"metrics"`
	Port                       types.Int64           `tfsdk:"port"`
	Prefix                     types.String          `tfsdk:"prefix"`
	Queue                      *Queue                `tfsdk:"queue"`
	QueueSize                  types.Int64           `tfsdk:"queue_size"`
	RetryCount                 types.Int64           `tfsdk:"retry_count"`
	ServiceIdentifierDefault   types.String          `tfsdk:"service_identifier_default"`
	TagStyle                   types.String          `tfsdk:"tag_style"`
	UDPPacketSize              types.Number          `tfsdk:"udp_packet_size"`
	UseTCP                     types.Bool            `tfsdk:"use_tcp"`
	WorkspaceIdentifierDefault types.String          `tfsdk:"workspace_identifier_default"`
}

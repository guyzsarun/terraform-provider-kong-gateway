// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Redis struct {
	ClusterMaxRedirections types.Int64     `tfsdk:"cluster_max_redirections"`
	ClusterNodes           []ClusterNodes  `tfsdk:"cluster_nodes"`
	ConnectTimeout         types.Int64     `tfsdk:"connect_timeout"`
	ConnectionIsProxied    types.Bool      `tfsdk:"connection_is_proxied"`
	Database               types.Int64     `tfsdk:"database"`
	Host                   types.String    `tfsdk:"host"`
	KeepaliveBacklog       types.Int64     `tfsdk:"keepalive_backlog"`
	KeepalivePoolSize      types.Int64     `tfsdk:"keepalive_pool_size"`
	Password               types.String    `tfsdk:"password"`
	Port                   types.Int64     `tfsdk:"port"`
	ReadTimeout            types.Int64     `tfsdk:"read_timeout"`
	SendTimeout            types.Int64     `tfsdk:"send_timeout"`
	SentinelMaster         types.String    `tfsdk:"sentinel_master"`
	SentinelNodes          []SentinelNodes `tfsdk:"sentinel_nodes"`
	SentinelPassword       types.String    `tfsdk:"sentinel_password"`
	SentinelRole           types.String    `tfsdk:"sentinel_role"`
	SentinelUsername       types.String    `tfsdk:"sentinel_username"`
	ServerName             types.String    `tfsdk:"server_name"`
	Ssl                    types.Bool      `tfsdk:"ssl"`
	SslVerify              types.Bool      `tfsdk:"ssl_verify"`
	Username               types.String    `tfsdk:"username"`
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Consul struct {
	Host    types.String `tfsdk:"host"`
	HTTPS   types.Bool   `tfsdk:"https"`
	KvPath  types.String `tfsdk:"kv_path"`
	Port    types.Int64  `tfsdk:"port"`
	Timeout types.Number `tfsdk:"timeout"`
	Token   types.String `tfsdk:"token"`
}

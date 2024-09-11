// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &PluginProxyCacheDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginProxyCacheDataSource{}

func NewPluginProxyCacheDataSource() datasource.DataSource {
	return &PluginProxyCacheDataSource{}
}

// PluginProxyCacheDataSource is the data source implementation.
type PluginProxyCacheDataSource struct {
	client *sdk.KongGateway
}

// PluginProxyCacheDataSourceModel describes the data model.
type PluginProxyCacheDataSourceModel struct {
	Config        *tfTypes.CreateProxyCachePluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                  `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                  `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                           `tfsdk:"created_at"`
	Enabled       types.Bool                            `tfsdk:"enabled"`
	ID            types.String                          `tfsdk:"id"`
	InstanceName  types.String                          `tfsdk:"instance_name"`
	Ordering      types.String                          `tfsdk:"ordering"`
	Protocols     []types.String                        `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                  `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                  `tfsdk:"service"`
	Tags          []types.String                        `tfsdk:"tags"`
	UpdatedAt     types.Int64                           `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginProxyCacheDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_proxy_cache"
}

// Schema defines the schema for the data source.
func (r *PluginProxyCacheDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginProxyCache DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cache_control": schema.BoolAttribute{
						Computed:    true,
						Description: `When enabled, respect the Cache-Control behaviors defined in RFC7234.`,
					},
					"cache_ttl": schema.Int64Attribute{
						Computed:    true,
						Description: `TTL, in seconds, of cache entities.`,
					},
					"content_type": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value.`,
					},
					"ignore_uri_case": schema.BoolAttribute{
						Computed: true,
					},
					"memory": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"dictionary_name": schema.StringAttribute{
								Computed:    true,
								Description: `The name of the shared dictionary in which to hold cache entities when the memory strategy is selected. Note that this dictionary currently must be defined manually in the Kong Nginx template.`,
							},
						},
					},
					"request_method": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Downstream request methods considered cacheable.`,
					},
					"response_code": schema.ListAttribute{
						Computed:    true,
						ElementType: types.Int64Type,
						Description: `Upstream response status code considered cacheable.`,
					},
					"response_headers": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"age": schema.BoolAttribute{
								Computed: true,
							},
							"x_cache_key": schema.BoolAttribute{
								Computed: true,
							},
							"x_cache_status": schema.BoolAttribute{
								Computed: true,
							},
						},
						Description: `Caching related diagnostic headers that should be included in cached responses`,
					},
					"storage_ttl": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of seconds to keep resources in the storage backend. This value is independent of ` + "`" + `cache_ttl` + "`" + ` or resource TTLs defined by Cache-Control behaviors.`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Description: `The backing data store in which to hold cache entities.`,
					},
					"vary_headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.`,
					},
					"vary_query_params": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *PluginProxyCacheDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongGateway)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.KongGateway, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PluginProxyCacheDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginProxyCacheDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	request := operations.GetProxycachePluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetProxycachePlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ProxyCachePlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedProxyCachePlugin(res.ProxyCachePlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

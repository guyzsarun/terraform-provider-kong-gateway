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
var _ datasource.DataSource = &PluginHTTPLogDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginHTTPLogDataSource{}

func NewPluginHTTPLogDataSource() datasource.DataSource {
	return &PluginHTTPLogDataSource{}
}

// PluginHTTPLogDataSource is the data source implementation.
type PluginHTTPLogDataSource struct {
	client *sdk.KongGateway
}

// PluginHTTPLogDataSourceModel describes the data model.
type PluginHTTPLogDataSourceModel struct {
	Config        *tfTypes.CreateHTTPLogPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer               `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer               `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                        `tfsdk:"created_at"`
	Enabled       types.Bool                         `tfsdk:"enabled"`
	ID            types.String                       `tfsdk:"id"`
	InstanceName  types.String                       `tfsdk:"instance_name"`
	Ordering      types.String                       `tfsdk:"ordering"`
	Protocols     []types.String                     `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer               `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer               `tfsdk:"service"`
	Tags          []types.String                     `tfsdk:"tags"`
	UpdatedAt     types.Int64                        `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginHTTPLogDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_http_log"
}

// Schema defines the schema for the data source.
func (r *PluginHTTPLogDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginHTTPLog DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"content_type": schema.StringAttribute{
						Computed:    true,
						Description: `Indicates the type of data sent. The only available option is ` + "`" + `application/json` + "`" + `.`,
					},
					"custom_fields_by_lua": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Lua code as a key-value map`,
					},
					"flush_timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `Optional time in seconds. If ` + "`" + `queue_size` + "`" + ` > 1, this is the max idle time before sending a log with less than ` + "`" + `queue_size` + "`" + ` records.`,
					},
					"headers": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `An optional table of headers included in the HTTP message to the upstream server. Values are indexed by header name, and each header name accepts a single string.`,
					},
					"http_endpoint": schema.StringAttribute{
						Computed:    true,
						Description: `A string representing a URL, such as https://example.com/path/to/resource?q=search.`,
					},
					"keepalive": schema.NumberAttribute{
						Computed:    true,
						Description: `An optional value in milliseconds that defines how long an idle connection will live before being closed.`,
					},
					"method": schema.StringAttribute{
						Computed:    true,
						Description: `An optional method used to send data to the HTTP server. Supported values are ` + "`" + `POST` + "`" + ` (default), ` + "`" + `PUT` + "`" + `, and ` + "`" + `PATCH` + "`" + `.`,
					},
					"queue": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"concurrency_limit": schema.Int64Attribute{
								Computed:    true,
								Description: `The number of of queue delivery timers. -1 indicates unlimited.`,
							},
							"initial_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Description: `Time in seconds before the initial retry is made for a failing batch.`,
							},
							"max_batch_size": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of entries that can be processed at a time.`,
							},
							"max_bytes": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of bytes that can be waiting on a queue, requires string content.`,
							},
							"max_coalescing_delay": schema.NumberAttribute{
								Computed:    true,
								Description: `Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.`,
							},
							"max_entries": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum number of entries that can be waiting on the queue.`,
							},
							"max_retry_delay": schema.NumberAttribute{
								Computed:    true,
								Description: `Maximum time in seconds between retries, caps exponential backoff.`,
							},
							"max_retry_time": schema.NumberAttribute{
								Computed:    true,
								Description: `Time in seconds before the queue gives up calling a failed handler for a batch.`,
							},
						},
					},
					"queue_size": schema.Int64Attribute{
						Computed:    true,
						Description: `Maximum number of log entries to be sent on each message to the upstream server.`,
					},
					"retry_count": schema.Int64Attribute{
						Computed:    true,
						Description: `Number of times to retry when sending data to the upstream server.`,
					},
					"timeout": schema.NumberAttribute{
						Computed:    true,
						Description: `An optional timeout in milliseconds when sending data to the upstream server.`,
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

func (r *PluginHTTPLogDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginHTTPLogDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginHTTPLogDataSourceModel
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

	request := operations.GetHttplogPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetHttplogPlugin(ctx, request)
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
	if !(res.HTTPLogPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedHTTPLogPlugin(res.HTTPLogPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

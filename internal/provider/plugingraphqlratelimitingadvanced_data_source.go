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
var _ datasource.DataSource = &PluginGraphqlRateLimitingAdvancedDataSource{}
var _ datasource.DataSourceWithConfigure = &PluginGraphqlRateLimitingAdvancedDataSource{}

func NewPluginGraphqlRateLimitingAdvancedDataSource() datasource.DataSource {
	return &PluginGraphqlRateLimitingAdvancedDataSource{}
}

// PluginGraphqlRateLimitingAdvancedDataSource is the data source implementation.
type PluginGraphqlRateLimitingAdvancedDataSource struct {
	client *sdk.KongGateway
}

// PluginGraphqlRateLimitingAdvancedDataSourceModel describes the data model.
type PluginGraphqlRateLimitingAdvancedDataSourceModel struct {
	Config        *tfTypes.CreateGraphqlRateLimitingAdvancedPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                                   `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                                   `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                                            `tfsdk:"created_at"`
	Enabled       types.Bool                                             `tfsdk:"enabled"`
	ID            types.String                                           `tfsdk:"id"`
	InstanceName  types.String                                           `tfsdk:"instance_name"`
	Ordering      types.String                                           `tfsdk:"ordering"`
	Protocols     []types.String                                         `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                                   `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                                   `tfsdk:"service"`
	Tags          []types.String                                         `tfsdk:"tags"`
	UpdatedAt     types.Int64                                            `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *PluginGraphqlRateLimitingAdvancedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_graphql_rate_limiting_advanced"
}

// Schema defines the schema for the data source.
func (r *PluginGraphqlRateLimitingAdvancedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginGraphqlRateLimitingAdvanced DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"cost_strategy": schema.StringAttribute{
						Computed:    true,
						Description: `Strategy to use to evaluate query costs. Either ` + "`" + `default` + "`" + ` or ` + "`" + `node_quantifier` + "`" + `.`,
					},
					"dictionary_name": schema.StringAttribute{
						Computed:    true,
						Description: `The shared dictionary where counters will be stored until the next sync cycle.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Description: `Optionally hide informative response headers. Available options: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
					},
					"identifier": schema.StringAttribute{
						Computed:    true,
						Description: `How to define the rate limit key. Can be ` + "`" + `ip` + "`" + `, ` + "`" + `credential` + "`" + `, ` + "`" + `consumer` + "`" + `.`,
					},
					"limit": schema.ListAttribute{
						Computed:    true,
						ElementType: types.NumberType,
						Description: `One or more requests-per-window limits to apply.`,
					},
					"max_cost": schema.NumberAttribute{
						Computed:    true,
						Description: `A defined maximum cost per query. 0 means unlimited.`,
					},
					"namespace": schema.StringAttribute{
						Computed:    true,
						Description: `The rate limiting namespace to use for this plugin instance. This namespace is used to share rate limiting counters across different instances. If it is not provided, a random UUID is generated. NOTE: For the plugin instances sharing the same namespace, all the configurations that are required for synchronizing counters, e.g. ` + "`" + `strategy` + "`" + `, ` + "`" + `redis` + "`" + `, ` + "`" + `sync_rate` + "`" + `, ` + "`" + `window_size` + "`" + `, ` + "`" + `dictionary_name` + "`" + `, need to be the same.`,
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"cluster_max_redirections": schema.Int64Attribute{
								Computed:    true,
								Description: `Maximum retry attempts for redirection.`,
							},
							"cluster_nodes": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ip": schema.StringAttribute{
											Computed:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
										},
									},
								},
								Description: `Cluster addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.`,
							},
							"connect_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"connection_is_proxied": schema.BoolAttribute{
								Computed:    true,
								Description: `If the connection to Redis is proxied (e.g. Envoy), set it ` + "`" + `true` + "`" + `. Set the ` + "`" + `host` + "`" + ` and ` + "`" + `port` + "`" + ` to point to the proxy address.`,
							},
							"database": schema.Int64Attribute{
								Computed:    true,
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"keepalive_backlog": schema.Int64Attribute{
								Computed:    true,
								Description: `Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return ` + "`" + `nil` + "`" + `. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than ` + "`" + `keepalive_pool_size` + "`" + `. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than ` + "`" + `keepalive_pool_size` + "`" + `.`,
							},
							"keepalive_pool_size": schema.Int64Attribute{
								Computed:    true,
								Description: `The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither ` + "`" + `keepalive_pool_size` + "`" + ` nor ` + "`" + `keepalive_backlog` + "`" + ` is specified, no pool is created. If ` + "`" + `keepalive_pool_size` + "`" + ` isn't specified but ` + "`" + `keepalive_backlog` + "`" + ` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
							},
							"read_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"send_timeout": schema.Int64Attribute{
								Computed:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
							},
							"sentinel_master": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_nodes": schema.ListNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Computed:    true,
											Description: `A string representing a host name, such as example.com.`,
										},
										"port": schema.Int64Attribute{
											Computed:    true,
											Description: `An integer representing a port number between 0 and 65535, inclusive.`,
										},
									},
								},
								Description: `Sentinel node addresses to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.`,
							},
							"sentinel_password": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.`,
							},
							"sentinel_role": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel role to use for Redis connections when the ` + "`" + `redis` + "`" + ` strategy is defined. Defining this value implies using Redis Sentinel.`,
							},
							"sentinel_username": schema.StringAttribute{
								Computed:    true,
								Description: `Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.`,
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Description: `If set to true, uses SSL to connect to Redis.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
					},
					"score_factor": schema.NumberAttribute{
						Computed:    true,
						Description: `A scoring factor to multiply (or divide) the cost. The ` + "`" + `score_factor` + "`" + ` must always be greater than 0.`,
					},
					"strategy": schema.StringAttribute{
						Computed:    true,
						Description: `The rate-limiting strategy to use for retrieving and incrementing the limits.`,
					},
					"sync_rate": schema.NumberAttribute{
						Computed:    true,
						Description: `How often to sync counter data to the central data store. A value of 0 results in synchronous behavior; a value of -1 ignores sync behavior entirely and only stores counters in node memory. A value greater than 0 syncs the counters in that many number of seconds.`,
					},
					"window_size": schema.ListAttribute{
						Computed:    true,
						ElementType: types.NumberType,
						Description: `One or more window sizes to apply a limit to (defined in seconds).`,
					},
					"window_type": schema.StringAttribute{
						Computed:    true,
						Description: `Sets the time window to either ` + "`" + `sliding` + "`" + ` or ` + "`" + `fixed` + "`" + `.`,
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

func (r *PluginGraphqlRateLimitingAdvancedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *PluginGraphqlRateLimitingAdvancedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PluginGraphqlRateLimitingAdvancedDataSourceModel
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

	request := operations.GetGraphqlratelimitingadvancedPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetGraphqlratelimitingadvancedPlugin(ctx, request)
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
	if !(res.GraphqlRateLimitingAdvancedPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedGraphqlRateLimitingAdvancedPlugin(res.GraphqlRateLimitingAdvancedPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

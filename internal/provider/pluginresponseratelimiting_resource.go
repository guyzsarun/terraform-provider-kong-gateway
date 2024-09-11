// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
	"github.com/kong/terraform-provider-kong-gateway/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PluginResponseRatelimitingResource{}
var _ resource.ResourceWithImportState = &PluginResponseRatelimitingResource{}

func NewPluginResponseRatelimitingResource() resource.Resource {
	return &PluginResponseRatelimitingResource{}
}

// PluginResponseRatelimitingResource defines the resource implementation.
type PluginResponseRatelimitingResource struct {
	client *sdk.KongGateway
}

// PluginResponseRatelimitingResourceModel describes the resource data model.
type PluginResponseRatelimitingResourceModel struct {
	Config        *tfTypes.CreateResponseRatelimitingPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                            `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                            `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                                     `tfsdk:"created_at"`
	Enabled       types.Bool                                      `tfsdk:"enabled"`
	ID            types.String                                    `tfsdk:"id"`
	InstanceName  types.String                                    `tfsdk:"instance_name"`
	Ordering      types.String                                    `tfsdk:"ordering"`
	Protocols     []types.String                                  `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                            `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                            `tfsdk:"service"`
	Tags          []types.String                                  `tfsdk:"tags"`
	UpdatedAt     types.Int64                                     `tfsdk:"updated_at"`
}

func (r *PluginResponseRatelimitingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_response_ratelimiting"
}

func (r *PluginResponseRatelimitingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginResponseRatelimiting Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"block_on_first_violation": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A boolean value that determines if the requests should be blocked as soon as one limit is being exceeded. This will block requests that are supposed to consume other limits too.`,
					},
					"fault_tolerant": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party datastore. If ` + "`" + `true` + "`" + `, requests will be proxied anyway, effectively disabling the rate-limiting function until the datastore is working again. If ` + "`" + `false` + "`" + `, then the clients will see ` + "`" + `500` + "`" + ` errors.`,
					},
					"header_name": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The name of the response header used to increment the counters.`,
					},
					"hide_client_headers": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Optionally hide informative response headers.`,
					},
					"limit_by": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The entity that will be used when aggregating the limits: ` + "`" + `consumer` + "`" + `, ` + "`" + `credential` + "`" + `, ` + "`" + `ip` + "`" + `. If the ` + "`" + `consumer` + "`" + ` or the ` + "`" + `credential` + "`" + ` cannot be determined, the system will always fallback to ` + "`" + `ip` + "`" + `. must be one of ["consumer", "credential", "ip"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"consumer",
								"credential",
								"ip",
							),
						},
					},
					"limits": schema.MapAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `A map that defines rate limits for the plugin.`,
						Validators: []validator.Map{
							mapvalidator.ValueStringsAre(validators.IsValidJSON()),
						},
					},
					"policy": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The rate-limiting policies to use for retrieving and incrementing the limits. must be one of ["local", "cluster", "redis"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"local",
								"cluster",
								"redis",
							),
						},
					},
					"redis": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"database": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `Database to use for the Redis connection when using the ` + "`" + `redis` + "`" + ` strategy`,
							},
							"host": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing a host name, such as example.com.`,
							},
							"password": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.`,
							},
							"port": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a port number between 0 and 65535, inclusive.`,
								Validators: []validator.Int64{
									int64validator.AtMost(65535),
								},
							},
							"server_name": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `A string representing an SNI (server name indication) value for TLS.`,
							},
							"ssl": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, uses SSL to connect to Redis.`,
							},
							"ssl_verify": schema.BoolAttribute{
								Computed:    true,
								Optional:    true,
								Description: `If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure ` + "`" + `lua_ssl_trusted_certificate` + "`" + ` in ` + "`" + `kong.conf` + "`" + ` to specify the CA (or server) certificate used by your Redis server. You may also need to configure ` + "`" + `lua_ssl_verify_depth` + "`" + ` accordingly.`,
							},
							"timeout": schema.Int64Attribute{
								Computed:    true,
								Optional:    true,
								Description: `An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.`,
								Validators: []validator.Int64{
									int64validator.AtMost(2147483646),
								},
							},
							"username": schema.StringAttribute{
								Computed:    true,
								Optional:    true,
								Description: `Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to ` + "`" + `default` + "`" + `.`,
							},
						},
						Description: `Redis configuration`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ordering": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: `Parsed as JSON.`,
				Validators: []validator.String{
					validators.IsValidJSON(),
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				Optional:    true,
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

func (r *PluginResponseRatelimitingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.KongGateway)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.KongGateway, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PluginResponseRatelimitingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PluginResponseRatelimitingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := data.ToSharedCreateResponseRatelimitingPlugin()
	res, err := r.client.Plugins.CreateResponseratelimitingPlugin(ctx, request)
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
	if res.StatusCode != 201 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ResponseRatelimitingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedResponseRatelimitingPlugin(res.ResponseRatelimitingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginResponseRatelimitingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PluginResponseRatelimitingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request := operations.GetResponseratelimitingPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetResponseratelimitingPlugin(ctx, request)
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
	if !(res.ResponseRatelimitingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedResponseRatelimitingPlugin(res.ResponseRatelimitingPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginResponseRatelimitingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PluginResponseRatelimitingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	createResponseRatelimitingPlugin := data.ToSharedCreateResponseRatelimitingPlugin()
	request := operations.UpdateResponseratelimitingPluginRequest{
		PluginID:                         pluginID,
		CreateResponseRatelimitingPlugin: createResponseRatelimitingPlugin,
	}
	res, err := r.client.Plugins.UpdateResponseratelimitingPlugin(ctx, request)
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
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.ResponseRatelimitingPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedResponseRatelimitingPlugin(res.ResponseRatelimitingPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginResponseRatelimitingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PluginResponseRatelimitingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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

	request := operations.DeleteResponseratelimitingPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.DeleteResponseratelimitingPlugin(ctx, request)
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
	if res.StatusCode != 204 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}

}

func (r *PluginResponseRatelimitingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

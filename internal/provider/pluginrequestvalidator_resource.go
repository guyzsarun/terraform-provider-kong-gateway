// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
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
	speakeasy_boolvalidators "github.com/kong/terraform-provider-kong-gateway/internal/validators/boolvalidators"
	speakeasy_objectvalidators "github.com/kong/terraform-provider-kong-gateway/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/kong/terraform-provider-kong-gateway/internal/validators/stringvalidators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &PluginRequestValidatorResource{}
var _ resource.ResourceWithImportState = &PluginRequestValidatorResource{}

func NewPluginRequestValidatorResource() resource.Resource {
	return &PluginRequestValidatorResource{}
}

// PluginRequestValidatorResource defines the resource implementation.
type PluginRequestValidatorResource struct {
	client *sdk.KongGateway
}

// PluginRequestValidatorResourceModel describes the resource data model.
type PluginRequestValidatorResourceModel struct {
	Config        *tfTypes.CreateRequestValidatorPluginConfig `tfsdk:"config"`
	Consumer      *tfTypes.ACLConsumer                        `tfsdk:"consumer"`
	ConsumerGroup *tfTypes.ACLConsumer                        `tfsdk:"consumer_group"`
	CreatedAt     types.Int64                                 `tfsdk:"created_at"`
	Enabled       types.Bool                                  `tfsdk:"enabled"`
	ID            types.String                                `tfsdk:"id"`
	InstanceName  types.String                                `tfsdk:"instance_name"`
	Ordering      types.String                                `tfsdk:"ordering"`
	Protocols     []types.String                              `tfsdk:"protocols"`
	Route         *tfTypes.ACLConsumer                        `tfsdk:"route"`
	Service       *tfTypes.ACLConsumer                        `tfsdk:"service"`
	Tags          []types.String                              `tfsdk:"tags"`
	UpdatedAt     types.Int64                                 `tfsdk:"updated_at"`
}

func (r *PluginRequestValidatorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_plugin_request_validator"
}

func (r *PluginRequestValidatorResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "PluginRequestValidator Resource",
		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"allowed_content_types": schema.ListAttribute{
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
						Description: `List of allowed content types. The value can be configured with the ` + "`" + `charset` + "`" + ` parameter. For example, ` + "`" + `application/json; charset=UTF-8` + "`" + `.`,
					},
					"body_schema": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `The request body schema specification. One of ` + "`" + `body_schema` + "`" + ` or ` + "`" + `parameter_schema` + "`" + ` must be specified.`,
					},
					"content_type_parameter_validation": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Determines whether to enable parameters validation of request content-type.`,
					},
					"parameter_schema": schema.ListNestedAttribute{
						Computed: true,
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Validators: []validator.Object{
								speakeasy_objectvalidators.NotNull(),
							},
							Attributes: map[string]schema.Attribute{
								"explode": schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Required when ` + "`" + `schema` + "`" + ` and ` + "`" + `style` + "`" + ` are set. When ` + "`" + `explode` + "`" + ` is ` + "`" + `true` + "`" + `, parameter values of type ` + "`" + `array` + "`" + ` or ` + "`" + `object` + "`" + ` generate separate parameters for each value of the array or key-value pair of the map. For other types of parameters, this property has no effect.`,
								},
								"in": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The location of the parameter. Not Null; must be one of ["query", "header", "path"]`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
										stringvalidator.OneOf(
											"query",
											"header",
											"path",
										),
									},
								},
								"name": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `The name of the parameter. Parameter names are case-sensitive, and correspond to the parameter name used by the ` + "`" + `in` + "`" + ` property. If ` + "`" + `in` + "`" + ` is ` + "`" + `path` + "`" + `, the ` + "`" + `name` + "`" + ` field MUST correspond to the named capture group from the configured ` + "`" + `route` + "`" + `. Not Null`,
									Validators: []validator.String{
										speakeasy_stringvalidators.NotNull(),
									},
								},
								"required": schema.BoolAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Determines whether this parameter is mandatory. Not Null`,
									Validators: []validator.Bool{
										speakeasy_boolvalidators.NotNull(),
									},
								},
								"schema": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Requred when ` + "`" + `style` + "`" + ` and ` + "`" + `explode` + "`" + ` are set. This is the schema defining the type used for the parameter. It is validated using ` + "`" + `draft4` + "`" + ` for JSON Schema draft 4 compliant validator. In addition to being a valid JSON Schema, the parameter schema MUST have a top-level ` + "`" + `type` + "`" + ` property to enable proper deserialization before validating.`,
								},
								"style": schema.StringAttribute{
									Computed:    true,
									Optional:    true,
									Description: `Required when ` + "`" + `schema` + "`" + ` and ` + "`" + `explode` + "`" + ` are set. Describes how the parameter value will be deserialized depending on the type of the parameter value. must be one of ["label", "form", "matrix", "simple", "spaceDelimited", "pipeDelimited", "deepObject"]`,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"label",
											"form",
											"matrix",
											"simple",
											"spaceDelimited",
											"pipeDelimited",
											"deepObject",
										),
									},
								},
							},
						},
						Description: `Array of parameter validator specification. One of ` + "`" + `body_schema` + "`" + ` or ` + "`" + `parameter_schema` + "`" + ` must be specified.`,
					},
					"verbose_response": schema.BoolAttribute{
						Computed:    true,
						Optional:    true,
						Description: `If enabled, the plugin returns more verbose and detailed validation errors.`,
					},
					"version": schema.StringAttribute{
						Computed:    true,
						Optional:    true,
						Description: `Which validator to use. Supported values are ` + "`" + `kong` + "`" + ` (default) for using Kong's own schema validator, or ` + "`" + `draft4` + "`" + ` for using a JSON Schema Draft 4-compliant validator. must be one of ["kong", "draft4"]`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"kong",
								"draft4",
							),
						},
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

func (r *PluginRequestValidatorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PluginRequestValidatorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *PluginRequestValidatorResourceModel
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

	request := data.ToSharedCreateRequestValidatorPlugin()
	res, err := r.client.Plugins.CreateRequestvalidatorPlugin(ctx, request)
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
	if !(res.RequestValidatorPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRequestValidatorPlugin(res.RequestValidatorPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginRequestValidatorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *PluginRequestValidatorResourceModel
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

	request := operations.GetRequestvalidatorPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.GetRequestvalidatorPlugin(ctx, request)
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
	if !(res.RequestValidatorPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRequestValidatorPlugin(res.RequestValidatorPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginRequestValidatorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *PluginRequestValidatorResourceModel
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

	createRequestValidatorPlugin := data.ToSharedCreateRequestValidatorPlugin()
	request := operations.UpdateRequestvalidatorPluginRequest{
		PluginID:                     pluginID,
		CreateRequestValidatorPlugin: createRequestValidatorPlugin,
	}
	res, err := r.client.Plugins.UpdateRequestvalidatorPlugin(ctx, request)
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
	if !(res.RequestValidatorPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedRequestValidatorPlugin(res.RequestValidatorPlugin)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *PluginRequestValidatorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *PluginRequestValidatorResourceModel
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

	request := operations.DeleteRequestvalidatorPluginRequest{
		PluginID: pluginID,
	}
	res, err := r.client.Plugins.DeleteRequestvalidatorPlugin(ctx, request)
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

func (r *PluginRequestValidatorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

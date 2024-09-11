// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-kong-gateway/internal/provider/types"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk"
	"github.com/kong/terraform-provider-kong-gateway/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GatewayConsumerGroupMemberResource{}
var _ resource.ResourceWithImportState = &GatewayConsumerGroupMemberResource{}

func NewGatewayConsumerGroupMemberResource() resource.Resource {
	return &GatewayConsumerGroupMemberResource{}
}

// GatewayConsumerGroupMemberResource defines the resource implementation.
type GatewayConsumerGroupMemberResource struct {
	client *sdk.KongGateway
}

// GatewayConsumerGroupMemberResourceModel describes the resource data model.
type GatewayConsumerGroupMemberResourceModel struct {
	ConsumerGroup   *tfTypes.ConsumerGroup `tfsdk:"consumer_group"`
	ConsumerGroupID types.String           `tfsdk:"consumer_group_id"`
	ConsumerID      types.String           `tfsdk:"consumer_id"`
	Consumers       []tfTypes.Consumer     `tfsdk:"consumers"`
}

func (r *GatewayConsumerGroupMemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_consumer_group_member"
}

func (r *GatewayConsumerGroupMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayConsumerGroupMember Resource",
		Attributes: map[string]schema.Attribute{
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"created_at": schema.Int64Attribute{
						Computed:    true,
						Description: `Unix epoch when the resource was created.`,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"tags": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"updated_at": schema.Int64Attribute{
						Computed:    true,
						Description: `Unix epoch when the resource was last updated.`,
					},
				},
			},
			"consumer_group_id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Requires replacement if changed.`,
			},
			"consumer_id": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: `Requires replacement if changed.`,
			},
			"consumers": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"created_at": schema.Int64Attribute{
							Computed:    true,
							Description: `Unix epoch when the resource was created.`,
						},
						"custom_id": schema.StringAttribute{
							Computed:    true,
							Description: `Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or ` + "`" + `username` + "`" + ` with the request.`,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `An optional set of strings associated with the Consumer for grouping and filtering.`,
						},
						"updated_at": schema.Int64Attribute{
							Computed:    true,
							Description: `Unix epoch when the resource was last updated.`,
						},
						"username": schema.StringAttribute{
							Computed:    true,
							Description: `The unique username of the Consumer. You must send either this field or ` + "`" + `custom_id` + "`" + ` with the request.`,
						},
					},
				},
			},
		},
	}
}

func (r *GatewayConsumerGroupMemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GatewayConsumerGroupMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *GatewayConsumerGroupMemberResourceModel
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

	var consumerGroupID string
	consumerGroupID = data.ConsumerGroupID.ValueString()

	requestBody := data.ToOperationsAddConsumerToGroupRequestBody()
	request := operations.AddConsumerToGroupRequest{
		ConsumerGroupID: consumerGroupID,
		RequestBody:     requestBody,
	}
	res, err := r.client.ConsumerGroups.AddConsumerToGroup(ctx, request)
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
	if !(res.Object != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromOperationsAddConsumerToGroupResponseBody(res.Object)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayConsumerGroupMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *GatewayConsumerGroupMemberResourceModel
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

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayConsumerGroupMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *GatewayConsumerGroupMemberResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GatewayConsumerGroupMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *GatewayConsumerGroupMemberResourceModel
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

	var consumerGroupID string
	consumerGroupID = data.ConsumerGroupID.ValueString()

	var consumerID string
	consumerID = data.ConsumerID.ValueString()

	request := operations.RemoveConsumerFromGroupRequest{
		ConsumerGroupID: consumerGroupID,
		ConsumerID:      consumerID,
	}
	res, err := r.client.ConsumerGroups.RemoveConsumerFromGroup(ctx, request)
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

func (r *GatewayConsumerGroupMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource gateway_consumer_group_member.")
}

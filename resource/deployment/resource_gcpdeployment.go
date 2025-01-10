package deploymentresource

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arcdeployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &GcpDeploymentResource{}

// var _ resource.ResourceWithImportState = &AwsDeploymentResource{}

type GcpDeploymentResource struct {
	loginCred utils.LoginCred
}

func NewGcpDeploymentResource() resource.Resource {
	return &GcpDeploymentResource{}
}

func (r *GcpDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gcp_deployment"
}

func (r *GcpDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	commonDeploymentSchema := DeploymentResourceModelSchema()
	gcpDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Resource identifier",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"instance_key": schema.StringAttribute{
			MarkdownDescription: "Path of the public key file which is needed to ssh into the deployed ArcEdge.",
			Required:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region where ArcEdge will be deployed.",
			Required:            true,
		},
		"byoip": schema.StringAttribute{
			MarkdownDescription: "A reserved IP which will be assigned to the deployed active ArcEdge.",
			Computed:            true,
			Optional:            true,
		},
		"byoip_backup": schema.StringAttribute{
			MarkdownDescription: "A reserved IP which will be assigned to the deployed standby ArcEdge.",
			Computed:            true,
			Optional:            true,
		},
		"zone": schema.StringAttribute{
			MarkdownDescription: "Zone where active ArcEdge will be deployed.",
			Optional:            true,
		},
		"backup_zone": schema.StringAttribute{
			MarkdownDescription: "Zone where standby ArcEdge will be deployed.",
			Optional:            true,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "Instance size of the ArcEdge deployed.",
			Required:            true,
		},
		"assign_public_ip": schema.BoolAttribute{
			Computed: true,
			Optional: true,
		},
		"networks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"project": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"network": schema.StringAttribute{
						Required: true,
					},
					"subnetwork": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			Required: true,
		},
		"network_tags": schema.ListAttribute{
			ElementType: types.StringType,
			Computed:    true,
			Optional:    true,
		},
	}

	for k, v := range commonDeploymentSchema {
		gcpDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: gcpDeploymentSchema,
	}
}

func (r *GcpDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	loginCred, ok := req.ProviderData.(utils.LoginCred)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *http.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.loginCred = loginCred
}

func (r *GcpDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arcdeployment.GcpDeploymentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(false, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	_, err := CreateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create AWS Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resDeployment, err := CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create AWS Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arcdeployment.GcpDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arcdeployment.GcpDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resDeployment, err := GetDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Re Refresh Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arcdeployment.GcpDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arcdeployment.GcpDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	_, err := UpdateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resDeployment, err := CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update AWS Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arcdeployment.GcpDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arcdeployment.GcpDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	err := DeleteDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	err = CheckDeletion(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
}

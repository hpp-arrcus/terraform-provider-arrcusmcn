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
)

var _ resource.Resource = &AzureDeploymentResource{}

// var _ resource.ResourceWithImportState = &AwsDeploymentResource{}

type AzureDeploymentResource struct {
	loginCred utils.LoginCred
}

func NewAzureDeploymentResource() resource.Resource {
	return &AzureDeploymentResource{}
}

func (r *AzureDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azure_deployment"
}

func (r *AzureDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	commonDeploymentSchema := DeploymentResourceModelSchema()
	azureDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource identifier",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region where ArcEdge will be deployed.",
			Required:            true,
		},
		"instance_key": schema.StringAttribute{
			MarkdownDescription: "Path of the public key file which is needed to ssh into the deployed ArcEdge.",
			Required:            true,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "Instance size of the ArcEdge deployed.",
			Required:            true,
		},
		"resource_group": schema.StringAttribute{
			MarkdownDescription: " Resource group where ArcEdge will be deployed.",
			Required:            true,
		},
		"vnet": schema.StringAttribute{
			MarkdownDescription: "Vnet where the ArcEdge will be deployed.",
			Required:            true,
		},
		"zone": schema.StringAttribute{
			MarkdownDescription: "Zone where active ArcEdge will be deployed.",
			Computed:            true,
			Optional:            true,
		},
		"backup_zone": schema.StringAttribute{
			MarkdownDescription: "Zone where standby ArcEdge will be deployed.",
			Computed:            true,
			Optional:            true,
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
		"enable_accelerated_networking": schema.BoolAttribute{
			MarkdownDescription: "Specifies if Accelerated Networking(SR-IOV) is enabled for the ArcEdge.",
			Required:            true,
		},
		"accelerated_networking_enabled": schema.BoolAttribute{
			Computed: true,
			Optional: true,
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
					"subnetwork": schema.StringAttribute{
						Required: true,
					},
				},
			},
			Optional: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		azureDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		Attributes: azureDeploymentSchema,
	}
}

func (r *AzureDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AzureDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arcdeployment.AzureDeploymentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(false, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resDeployment, err := CreateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Azure Deployment",
			err.Error(),
		)
		return
	}

	var res arcdeployment.AzureDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)

	resDeployment, err = CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Azure Deployment failed to create",
			err.Error(),
		)
		return
	}

	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AzureDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arcdeployment.AzureDeploymentResourceModel
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

	var res arcdeployment.AzureDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AzureDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arcdeployment.AzureDeploymentResourceModel
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
			"Unable to Re Update Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resDeployment, err := CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Azure Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arcdeployment.AzureDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AzureDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arcdeployment.AzureDeploymentResourceModel
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

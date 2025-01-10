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

var _ resource.Resource = &AwsDeploymentResource{}

// var _ resource.ResourceWithImportState = &AwsDeploymentResource{}

type AwsDeploymentResource struct {
	loginCred utils.LoginCred
}

func NewAwsDeploymentResource() resource.Resource {
	return &AwsDeploymentResource{}
}

func (r *AwsDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aws_deployment"
}

func (r *AwsDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	commonDeploymentSchema := DeploymentResourceModelSchema()
	awsDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed:            true,
			MarkdownDescription: "Resource identifier",
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "A unique name for an ArcEdge deployment running on AWS",
			Required:            true,
		},
		"instance_key": schema.StringAttribute{
			MarkdownDescription: "Instance key name which is needed to ssh into the deployed ArcEdge.",
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
		"vpc_id": schema.StringAttribute{
			MarkdownDescription: "VPC ID where the ArcEdge will be deployed.",
			Required:            true,
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
					"subnet_a": schema.StringAttribute{
						Required: true,
					},
					"subnet_b": schema.StringAttribute{
						Optional: true,
					},
					"private_subnet_route_table": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			Required: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		awsDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: awsDeploymentSchema,
	}
}

func (r *AwsDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *AwsDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arcdeployment.AwsDeploymentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(false)
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

	var res arcdeployment.AwsDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AwsDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arcdeployment.AwsDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(true)
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

	var res arcdeployment.AwsDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AwsDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arcdeployment.AwsDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true)
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

	var res arcdeployment.AwsDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *AwsDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arcdeployment.AwsDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true)
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

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

var _ resource.Resource = &KvmDeploymentResource{}

// var _ resource.ResourceWithImportState = &AwsDeploymentResource{}

type KvmDeploymentResource struct {
	loginCred utils.LoginCred
}

func NewKvmDeploymentResource() resource.Resource {
	return &KvmDeploymentResource{}
}

func (r *KvmDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kvm_deployment"
}

func (r *KvmDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	commonDeploymentSchema := DeploymentResourceModelSchema()
	kvmDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"vcpus": schema.Int64Attribute{
			Required: true,
		},
		"vm_memory": schema.Int64Attribute{
			Required: true,
		},
		"private_ip": schema.StringAttribute{
			Required: true,
		},
		"public_ip": schema.StringAttribute{
			Required: true,
		},
		"ssh_psw": schema.StringAttribute{
			Computed: true,
			Optional: true,
		},
		"public_ip_backup": schema.StringAttribute{
			Computed: true,
			Optional: true,
		},
		"private_ip_backup": schema.StringAttribute{
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
					"private_ip_default_gw": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"private_ip_cidr_mask": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"network": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
				},
			},
			Required: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		kvmDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		Attributes: kvmDeploymentSchema,
	}
}

func (r *KvmDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *KvmDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arcdeployment.KvmDeploymentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(false)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resDeployment, err := CreateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create KVM Deployment",
			err.Error(),
		)
		return
	}

	var res arcdeployment.KvmDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)

	resDeployment, err = CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"KVM Deployment failed to create",
			err.Error(),
		)
		return
	}

	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.SshPsw = data.SshPsw
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *KvmDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arcdeployment.KvmDeploymentResourceModel
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

	var res arcdeployment.KvmDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.SshPsw = data.SshPsw
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *KvmDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arcdeployment.KvmDeploymentResourceModel
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

	var res arcdeployment.KvmDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	res.SshPsw = data.SshPsw
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *KvmDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arcdeployment.KvmDeploymentResourceModel
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

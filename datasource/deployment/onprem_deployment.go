package datadeployment

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arcdeployment"
	deploymentresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/deployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSource = &KvmDeploymentData{}

type KvmDeploymentData struct {
	loginCred utils.LoginCred
}

func NewKvmDeploymentDataSource() datasource.DataSource {
	return &KvmDeploymentData{}
}

func (d *KvmDeploymentData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kvm_deployment"
}

func (d *KvmDeploymentData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	commonDeploymentSchema := deploymentresource.DeploymentDataModelSchema()
	kvmDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"vcpus": schema.Int64Attribute{
			Computed: true,
		},
		"vm_memory": schema.Int64Attribute{
			Computed: true,
		},
		"private_ip": schema.StringAttribute{
			Computed: true,
		},
		"public_ip": schema.StringAttribute{
			Computed: true,
		},
		"ssh_psw": schema.StringAttribute{
			Computed: true,
		},
		"public_ip_backup": schema.StringAttribute{
			Computed: true,
		},
		"private_ip_backup": schema.StringAttribute{
			Computed: true,
		},
		"assign_public_ip": schema.BoolAttribute{
			Computed: true,
		},
		"networks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"private_ip_default_gw": schema.StringAttribute{
						Computed: true,
					},
					"private_ip_cidr_mask": schema.StringAttribute{
						Computed: true,
					},
					"network": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		kvmDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: kvmDeploymentSchema,
	}
}

func (r *KvmDeploymentData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *KvmDeploymentData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arcdeployment.KvmDeploymentResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resDeployment, err := deploymentresource.GetDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

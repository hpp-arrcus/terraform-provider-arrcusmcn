package datadeployment

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arcdeployment"
	deploymentresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/deployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSource = &AzureDeploymentData{}

type AzureDeploymentData struct {
	loginCred utils.LoginCred
}

func NewAzureDeploymentDataSource() datasource.DataSource {
	return &AzureDeploymentData{}
}

func (d *AzureDeploymentData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azure_deployment"
}

func (d *AzureDeploymentData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	commonDeploymentSchema := deploymentresource.DeploymentDataModelSchema()
	azureDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"region": schema.StringAttribute{
			Computed: true,
		},
		"instance_key": schema.StringAttribute{
			Computed: true,
		},
		"instance_type": schema.StringAttribute{
			Computed: true,
		},
		"resource_group": schema.StringAttribute{
			Computed: true,
		},
		"vnet": schema.StringAttribute{
			Computed: true,
		},
		"zone": schema.StringAttribute{
			Computed: true,
		},
		"backup_zone": schema.StringAttribute{
			Computed: true,
		},
		"byoip": schema.StringAttribute{
			Computed: true,
		},
		"byoip_backup": schema.StringAttribute{
			Computed: true,
		},
		"enable_accelerated_networking": schema.BoolAttribute{
			Computed: true,
		},
		"accelerated_networking_enabled": schema.BoolAttribute{
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
					"subnetwork": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		azureDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: azureDeploymentSchema,
	}
}

func (r *AzureDeploymentData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AzureDeploymentData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arcdeployment.AzureDeploymentResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(false, false)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	js, _ := json.Marshal(*reqDeployment)
	log.Default().Println(string(js))
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

	var res arcdeployment.AzureDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

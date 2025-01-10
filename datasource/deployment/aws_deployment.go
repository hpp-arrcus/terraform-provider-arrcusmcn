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

var _ datasource.DataSource = &AwsDeploymentData{}

type AwsDeploymentData struct {
	loginCred utils.LoginCred
}

func NewAwsDeploymentDataSource() datasource.DataSource {
	return &AwsDeploymentData{}
}

func (d *AwsDeploymentData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aws_deployment"
}

func (d *AwsDeploymentData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	commonDeploymentSchema := deploymentresource.DeploymentDataModelSchema()
	awsDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			MarkdownDescription: "Aws Deployment name",
			Required:            true,
		},
		"instance_key": schema.StringAttribute{
			MarkdownDescription: "instance key",
			Computed:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "region",
			Computed:            true,
		},
		"byoip": schema.StringAttribute{
			MarkdownDescription: "bring your own ip address active acredge instance",
			Computed:            true,
		},
		"byoip_backup": schema.StringAttribute{
			MarkdownDescription: "bring your own ip address for backup arcedge instance",
			Computed:            true,
		},
		"zone": schema.StringAttribute{
			MarkdownDescription: "zone for active arcedge instance",
			Computed:            true,
		},
		"backup_zone": schema.StringAttribute{
			MarkdownDescription: "zone for backup arcedge instance",
			Computed:            true,
		},
		"vpc_id": schema.StringAttribute{
			MarkdownDescription: "vpc id",
			Computed:            true,
		},
		"instance_type": schema.StringAttribute{
			MarkdownDescription: "bring your own ip address",
			Computed:            true,
		},
		"assign_public_ip": schema.BoolAttribute{
			MarkdownDescription: "whether to assign public ip to arcedge instances",
			Computed:            true,
		},
		"networks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Computed: true,
					},
					"subnet_a": schema.StringAttribute{
						Computed: true,
					},
					"subnet_b": schema.StringAttribute{
						Computed: true,
					},
					"private_subnet_route_table": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
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

func (r *AwsDeploymentData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AwsDeploymentData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arcdeployment.AwsDeploymentResourceModel
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

	var res arcdeployment.AwsDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

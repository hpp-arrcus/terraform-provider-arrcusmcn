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

var _ datasource.DataSource = &OciDeploymentData{}

type OciDeploymentData struct {
	loginCred utils.LoginCred
}

func NewOciDeploymentDataSource() datasource.DataSource {
	return &OciDeploymentData{}
}

func (d *OciDeploymentData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oci_deployment"
}

func (d *OciDeploymentData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	commonDeploymentSchema := deploymentresource.DeploymentDataModelSchema()
	ociDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"name": schema.StringAttribute{
			Required: true,
		},
		"instance_key": schema.StringAttribute{
			Computed: true,
		},
		"region": schema.StringAttribute{
			Computed: true,
		},
		"availability_domain": schema.StringAttribute{
			Computed: true,
		},
		"backup_availability_domain": schema.StringAttribute{
			Computed: true,
		},
		"compartment": schema.StringAttribute{
			Computed: true,
		},
		"image_compartment": schema.StringAttribute{
			Computed: true,
		},
		"compute_shape": schema.StringAttribute{
			Computed: true,
		},
		"compute_cpus": schema.Int32Attribute{
			Computed: true,
		},
		"compute_memory_in_gbs": schema.Int32Attribute{
			Computed: true,
		},
		"enable_firewall": schema.BoolAttribute{
			Computed: true,
		},
		"byoip": schema.StringAttribute{
			Computed: true,
		},
		"byoip_backup": schema.StringAttribute{
			Computed: true,
		},
		"networks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"subnet_name": schema.StringAttribute{
						Computed: true,
					},
					"vcn_ocid": schema.StringAttribute{
						Computed: true,
					},
					"vcn_name": schema.StringAttribute{
						Computed: true,
					},
					"subnet_ocid": schema.StringAttribute{
						Computed: true,
					},
					"subnet_access": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
	}

	for k, v := range commonDeploymentSchema {
		ociDeploymentSchema[k] = v
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: ociDeploymentSchema,
	}
}

func (r *OciDeploymentData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *OciDeploymentData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arcdeployment.OciDeploymentResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(false, false)
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

	var res arcdeployment.OciDeploymentResourceModel
	diag = res.ToTerraformModel(resDeployment)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

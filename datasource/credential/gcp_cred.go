package datacredential

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arccredential"
	credentialresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/credential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

var _ datasource.DataSource = &GcpCredData{}

type GcpCredData struct {
	loginCred utils.LoginCred
}

func NewGcpCredDataSource() datasource.DataSource {
	return &GcpCredData{}
}

func (d *GcpCredData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gcp_cred"
}

func (d *GcpCredData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"account_key_file": schema.StringAttribute{
				Computed: true,
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Computed: true,
			},
			"private_key_id": schema.StringAttribute{
				Computed: true,
			},
			"private_key": schema.StringAttribute{
				Computed: true,
			},
			"client_email": schema.StringAttribute{
				Computed: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
			},
			"auth_uri": schema.StringAttribute{
				Computed: true,
			},
			"token_uri": schema.StringAttribute{
				Computed: true,
			},
			"auth_provider_x509_cert_url": schema.StringAttribute{
				Computed: true,
			},
			"client_x509_cert_url": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *GcpCredData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *GcpCredData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arccredential.GcpCredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(false)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resCred, err := credentialresource.GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arccredential.GcpCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

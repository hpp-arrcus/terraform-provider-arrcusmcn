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

var _ datasource.DataSource = &OciCredData{}

type OciCredData struct {
	loginCred utils.LoginCred
}

func NewOciCredDataSource() datasource.DataSource {
	return &OciCredData{}
}
func (d *OciCredData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oci_cred"
}

func (d *OciCredData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"user": schema.StringAttribute{
				Computed: true,
			},
			"identity_domain": schema.StringAttribute{
				Computed: true,
			},
			"tenancy": schema.StringAttribute{
				Computed: true,
			},
			"region": schema.StringAttribute{
				Computed: true,
			},
			"key_file": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "awscred_identifier",
			},
		},
	}
}

func (r *OciCredData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *OciCredData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arccredential.OciCredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(false)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resCred, err := credentialresource.GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arccredential.OciCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

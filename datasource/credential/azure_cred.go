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

var _ datasource.DataSource = &AzureCredData{}

type AzureCredData struct {
	loginCred utils.LoginCred
}

func NewAzureCredDataSource() datasource.DataSource {
	return &AzureCredData{}
}

func (d *AzureCredData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_azure_cred"
}

func (d *AzureCredData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"subscription_id": schema.StringAttribute{
				Computed: true,
			},
			"client_id": schema.StringAttribute{
				Computed: true,
			},
			"client_secret": schema.StringAttribute{
				Computed: true,
			},
			"tenant_id": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *AzureCredData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AzureCredData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arccredential.AzureCredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred := data.ToAOCredModel()
	resCred, err := credentialresource.GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh Azure Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arccredential.AzureCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh Azure Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

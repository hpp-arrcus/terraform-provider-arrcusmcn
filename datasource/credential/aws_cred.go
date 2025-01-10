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

var _ datasource.DataSource = &AwsCredData{}

type AwsCredData struct {
	loginCred utils.LoginCred
}

func NewAwsCredDataSource() datasource.DataSource {
	return &AwsCredData{}
}

func (d *AwsCredData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aws_cred"
}

func (d *AwsCredData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "Aws credential name",
				Required:            true,
			},
			"access_key": schema.StringAttribute{
				MarkdownDescription: "Aws access key",
				Computed:            true,
			},
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "Aws secret key",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "awscred_identifier",
			},
		},
	}
}

func (r *AwsCredData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *AwsCredData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arccredential.AwsCredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred := data.ToAOCredModel()
	resCred, err := credentialresource.GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh AWS Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arccredential.AwsCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh AWS Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

// import (
// 	"context"
// 	"errors"

// 	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
// 	resourceCred "github.com/Arrcus/terraform-provider-arrcusmcn/resource/credential"
// 	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas/credential"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func DataSourceAwsCred() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceAwsCredRead,
// 		Schema:      schemas.AwsCredentialDataSchema(),
// 	}
// }

// func dataSourceAwsCredRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	accessToken := m.(map[string]string)["access_token"]
// 	tenant := m.(map[string]string)["tenant"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	err := resourceCred.GetCred(ctx, baseURL, tenant, accessToken, d, models.ProvidersAws)
// 	if err != nil {
// 		return diag.FromErr(errors.New("Can't find AWS credential with given name."))
// 	}
// 	d.SetId(d.Get("name").(string))
// 	return diags
// }

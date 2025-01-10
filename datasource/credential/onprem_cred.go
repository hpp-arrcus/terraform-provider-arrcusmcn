package datacredential

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arccredential"
	credentialresource "github.com/Arrcus/terraform-provider-arrcusmcn/resource/credential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &KvmCredData{}

type KvmCredData struct {
	loginCred utils.LoginCred
}

func NewKvmCredDataSource() datasource.DataSource {
	return &KvmCredData{}
}

func (d *KvmCredData) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kvm_cred"
}

func (d *KvmCredData) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"server_ip": schema.StringAttribute{
				Computed: true,
			},
			"user_name": schema.StringAttribute{
				Computed: true,
			},
			"domain": schema.StringAttribute{
				Computed: true,
			},
			"ssh_key": schema.StringAttribute{
				Computed: true,
			},
			"data_if_name": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *KvmCredData) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *KvmCredData) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data arccredential.KvmCredentialResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, diag := data.ToAOCredModel(false)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resCred, err := credentialresource.GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh KVM Credential",
			"An unexpected error occurred while creating the resource read request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res arccredential.KvmCredentialResourceModel
	diag = res.ToTerraformModel(resCred)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

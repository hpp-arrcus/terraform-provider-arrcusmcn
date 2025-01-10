package credentialresource

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arccredential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &AzureCredentialResource{}
var _ resource.ResourceWithImportState = &AzureCredentialResource{}

func NewGcpCredentialResource() resource.Resource {
	return &GcpCredentialResource{}
}

type GcpCredentialResource struct {
	loginCred utils.LoginCred
}

func (r *GcpCredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gcp_cred"
}

func (r *GcpCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"account_key_file": schema.StringAttribute{
				Optional: true,
			},
			"type": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"project_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"private_key_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"private_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"client_email": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"client_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"auth_uri": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"token_uri": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"auth_provider_x509_cert_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"client_x509_cert_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *GcpCredentialResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *GcpCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arccredential.GcpCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create GCP Credential",
			err.Error(),
		)
		return
	}

	resCred, err := CreateCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create GCP Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.GcpCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create GCP Credential",
			err.Error(),
		)
		return
	}
	res.PrivateKey = types.StringValue(reqCred.Credentials.GcpCredentials.PrivateKey)
	res.AccountKeyFile = types.StringValue(data.AccountKeyFile.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arccredential.GcpCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			err.Error(),
		)
		return
	}
	resCred, err := GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.GcpCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh GCP Credential",
			err.Error(),
		)
		return
	}
	res.PrivateKey = types.StringValue(reqCred.Credentials.GcpCredentials.PrivateKey)
	res.AccountKeyFile = types.StringValue(data.AccountKeyFile.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arccredential.GcpCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update GCP Credential",
			err.Error(),
		)
		return
	}

	resCred, err := UpdateCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update GCP Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.GcpCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update GCP Credential",
			err.Error(),
		)
		return
	}
	res.PrivateKey = types.StringValue(reqCred.Credentials.GcpCredentials.PrivateKey)
	res.AccountKeyFile = types.StringValue(data.AccountKeyFile.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *GcpCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arccredential.GcpCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := DeleteCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Aws Credential",
			"An unexpected error occurred while attempting to delete the resource. "+
				"Please retry the operation or report this issue to the provider developers.\n\n"+
				"HTTP Error: "+err.Error(),
		)
		return
	}
}

func (r *GcpCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func loadAccountKeyFile(filePath string) (*models.GcpCredentials, error) {
	jsonString, err := utils.ReadTextFile(filePath)
	if err != nil {
		return nil, err
	}
	gcpCred := models.GcpCredentials{}
	err = json.Unmarshal([]byte(*jsonString), &gcpCred)
	if err != nil {
		return nil, err
	}
	return &gcpCred, nil
}

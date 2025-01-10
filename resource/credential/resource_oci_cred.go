package credentialresource

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arccredential"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.Resource = &OciCredentialResource{}
var _ resource.ResourceWithImportState = &OciCredentialResource{}

func NewOciCredentialResource() resource.Resource {
	return &OciCredentialResource{}
}

type OciCredentialResource struct {
	loginCred utils.LoginCred
}

func (r *OciCredentialResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oci_cred"
}

func (r *OciCredentialResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"user": schema.StringAttribute{
				Required: true,
			},
			"identity_domain": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"tenancy": schema.StringAttribute{
				Required: true,
			},
			"region": schema.StringAttribute{
				Required: true,
			},
			"key_file": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "awscred_identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *OciCredentialResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OciCredentialResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arccredential.OciCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OCI Credential",
			err.Error(),
		)
		return
	}

	resCred, err := CreateCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OCI Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.OciCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OCI Credential",
			err.Error(),
		)
		return
	}

	res.KeyFile = data.KeyFile
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciCredentialResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arccredential.OciCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			err.Error(),
		)
		return
	}

	resCred, err := GetCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqCred.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.OciCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Refresh OCI Credential",
			err.Error(),
		)
		return
	}

	res.KeyFile = data.KeyFile
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciCredentialResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arccredential.OciCredentialResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqCred, err := data.ToAOCredModel(true)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OCI Credential",
			err.Error(),
		)
		return
	}

	resCred, err := UpdateCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OCI Credential",
			err.Error(),
		)
		return
	}

	var res arccredential.OciCredentialResourceModel
	err = res.ToTerraformModel(resCred)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OCI Credential",
			err.Error(),
		)
		return
	}
	res.KeyFile = data.KeyFile
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciCredentialResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arccredential.OciCredentialResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	err := DeleteCred(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, data.Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete OCI Credential",
			err.Error(),
		)
		return
	}
}

func (r *OciCredentialResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

package deploymentresource

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Arrcus/terraform-provider-arrcusmcn/arcmodels/arcdeployment"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

var _ resource.Resource = &OciDeploymentResource{}

type OciDeploymentResource struct {
	loginCred utils.LoginCred
}

func NewOciDeploymentResource() resource.Resource {
	return &OciDeploymentResource{}
}

func (r *OciDeploymentResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_oci_deployment"
}

func (r *OciDeploymentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	commonDeploymentSchema := DeploymentResourceModelSchema()
	ociDeploymentSchema := map[string]schema.Attribute{
		"id": schema.StringAttribute{
			MarkdownDescription: "Resource identifier",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		// "name": schema.StringAttribute{
		// 	Required: true,
		// },
		"instance_key": schema.StringAttribute{
			MarkdownDescription: "Path of the public key file which is needed to ssh into the deployed ArcEdge.",
			Required:            true,
		},
		"region": schema.StringAttribute{
			MarkdownDescription: "Region where ArcEdge will be deployed.",
			Required:            true,
		},
		"availability_domain": schema.StringAttribute{
			MarkdownDescription: "Availability Domain where active ArcEdge will be deployed.",
			Required:            true,
		},
		"backup_availability_domain": schema.StringAttribute{
			MarkdownDescription: "Availability Domain where standby ArcEdge will be deployed.",
			Optional:            true,
			Computed:            true,
		},
		"compartment": schema.StringAttribute{
			MarkdownDescription: "Ocid of the compartment where ArcEdge will be deployed.",
			Required:            true,
		},
		"image_compartment": schema.StringAttribute{
			MarkdownDescription: "Ocid of compartment where the arcedge image located.",
			Required:            true,
		},
		"compute_shape": schema.StringAttribute{
			MarkdownDescription: "Instance size of the ArcEdge deployed.",
			Required:            true,
		},
		"compute_cpus": schema.Int32Attribute{
			MarkdownDescription: "Number of CPUs of the ArcEdge deployed. This is required if compute shape is flexible.",
			Required:            true,
		},
		"compute_memory_in_gbs": schema.Int32Attribute{
			MarkdownDescription: "Memory size of the ArcEdge deployed. This is required if compute shape is flexible.",
			Required:            true,
		},
		"enable_firewall": schema.BoolAttribute{
			MarkdownDescription: "",
			Optional:            true,
			Computed:            true,
		},
		"byoip": schema.StringAttribute{
			MarkdownDescription: "A reserved IP which will be assigned to the deployed active ArcEdge.",
			Optional:            true,
			Computed:            true,
		},
		"byoip_backup": schema.StringAttribute{
			MarkdownDescription: "A reserved IP which will be assigned to the deployed standby ArcEdge.",
			Optional:            true,
			Computed:            true,
		},
		"networks": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"subnet_name": schema.StringAttribute{
						Computed: true,
						Optional: true,
					},
					"vcn_ocid": schema.StringAttribute{
						Required: true,
					},
					"vcn_name": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"subnet_ocid": schema.StringAttribute{
						Required: true,
					},
					"subnet_access": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Required: true,
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

func (r *OciDeploymentResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *OciDeploymentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data arcdeployment.OciDeploymentResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(false, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	js, _ := json.Marshal(reqDeployment)
	log.Default().Println(string(js))
	_, err := CreateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OCI Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resDeployment, err := CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create OCI Deployment",
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
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciDeploymentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data arcdeployment.OciDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}
	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resDeployment, err := GetDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
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
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciDeploymentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data arcdeployment.OciDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	_, err := UpdateDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqDeployment)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	resDeployment, err := CheckDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name, "Creation")
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Update OCI Deployment",
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
	res.InstanceKey = data.InstanceKey
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *OciDeploymentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data arcdeployment.OciDeploymentResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqDeployment, diag := data.ToAODeploymentModel(true, true)
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	err := DeleteDeployment(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	err = CheckDeletion(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqDeployment.Name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Delete Deployment",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
}

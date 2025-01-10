package resource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"

	// schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas/credential"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &UserResource{}

type UserResource struct {
	loginCred utils.LoginCred
}

func NewUserResource() resource.Resource {
	return &UserResource{}
}

type UserResourceModel struct {
	UserName     types.String `tfsdk:"username"`
	SamlUserName types.String `tfsdk:"saml_username"`
	Name         types.String `tfsdk:"name"`
	Password     types.String `tfsdk:"password"`
	Email        types.String `tfsdk:"email"`
	IsDefault    types.Bool   `tfsdk:"is_default"`
	Roles        types.List   `tfsdk:"roles"`
}

func (m *UserResourceModel) ToAOUserModel() (*models.User, diag.Diagnostics) {
	user := models.User{
		Name:     m.Name.String(),
		Username: m.UserName.ValueStringPointer(),
		Password: m.Password.ValueStringPointer(),
		Email:    m.Email.ValueStringPointer(),
	}

	if !m.IsDefault.IsNull() && !m.IsDefault.IsUnknown() {
		user.IsDefault = m.IsDefault.ValueBoolPointer()
	}

	if !m.SamlUserName.IsNull() {
		user.SamlUsername = m.SamlUserName.String()
	}

	roles := make([]string, 0)
	diag := m.Roles.ElementsAs(context.Background(), roles, false)
	if diag.HasError() {
		return nil, diag
	}

	user.Roles = make([]models.Rolename, 0)
	for _, role := range roles {
		switch role {
		case string(models.RolenameArcOrchAdmin):
			user.Roles = append(user.Roles, models.RolenameArcOrchAdmin)
		case string(models.RolenameTenantAdmin):
			user.Roles = append(user.Roles, models.RolenameTenantAdmin)
		case string(models.RolenameTenantOperator):
			user.Roles = append(user.Roles, models.RolenameTenantOperator)
		case string(models.RolenameDashboardReader):
			user.Roles = append(user.Roles, models.RolenameDashboardReader)
		}
	}
	return &user, nil
}

func (m *UserResourceModel) ToTerraformModel(user *models.User) diag.Diagnostics {
	m.UserName = types.StringValue(*user.Username)
	m.Name = types.StringValue(user.Name)
	m.Email = types.StringValue(*user.Email)
	m.IsDefault = types.BoolValue(*user.IsDefault)
	m.SamlUserName = types.StringValue(user.SamlUsername)

	elements := make([]attr.Value, 0)
	for _, role := range user.Roles {
		elements = append(elements, types.StringValue(string(role)))
	}
	roleList, diag := types.ListValue(types.StringType, elements)
	if diag.HasError() {
		return diag
	}
	m.Roles = roleList
	return nil
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Required: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"password": schema.StringAttribute{
				Required: true,
			},
			"email": schema.StringAttribute{
				Required: true,
			},
			"is_default": schema.BoolAttribute{
				Computed: true,
			},
			"roles": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
		},
	}
}

func (r *UserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	loginCred, ok := req.ProviderData.(utils.LoginCred)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected LoginCred, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.loginCred = loginCred
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqUser, diag := data.ToAOUserModel()
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resUser, err := CreateUser(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqUser)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create User",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res UserResourceModel
	diag = res.ToTerraformModel(resUser)
	if diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqUser, diag := data.ToAOUserModel()
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resUser, err := ReadUser(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqUser.Username)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create User",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res UserResourceModel
	diag = res.ToTerraformModel(resUser)
	if diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqUser, diag := data.ToAOUserModel()
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	resUser, err := UpdateUser(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, reqUser)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create User",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	var res UserResourceModel
	diag = res.ToTerraformModel(resUser)
	if diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &res)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UserResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	reqUser, diag := data.ToAOUserModel()
	if diag != nil && diag.HasError() {
		resp.Diagnostics.Append(diag...)
		return
	}

	err := DeleteUser(ctx, r.loginCred.BaseUrl, r.loginCred.Tenant, r.loginCred.AuthorizationHeader, r.loginCred.AccessToken, *reqUser.Username)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create User",
			"An unexpected error occurred while creating the resource create request. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}
}

func CreateUser(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	user *models.User,
) (*models.User, error) {
	url := baseURL + "users/" + "?tenant=" + tenant

	res, err := utils.PostRequest(url, *user, authHeader, accessToken)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func ReadUser(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	username string,
) (*models.User, error) {
	tokens := strings.Split(username, "@")
	url := fmt.Sprintf("%susers/%s?tenant=%s", baseURL, tokens[0], tenant)
	res, err := utils.GetRequest(url, authHeader, accessToken)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}
	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func UpdateUser(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	user *models.User,
) (*models.User, error) {
	tokens := strings.Split(*user.Username, "@")
	url := fmt.Sprintf("%susers/%s?tenant=%s", baseURL, tokens[0], tenant)
	res, err := utils.PutRequest(url, *user, authHeader, accessToken)
	if err != nil {
		return nil, err
	}
	resBody, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 204 && res.StatusCode != 200 {
		return nil, errors.New(string(resBody))
	}
	resUser := models.User{}
	err = json.Unmarshal(resBody, &resUser)
	if err != nil {
		return nil, err
	}

	return &resUser, nil
}

func DeleteUser(
	ctx context.Context,
	baseURL string,
	tenant string,
	authHeader string,
	accessToken string,
	username string,
) error {
	tokens := strings.Split(username, "@")
	url := fmt.Sprintf("%susers/%s?tenant=%s", baseURL, tokens[0], tenant)
	err := utils.DeleteRequest(url, authHeader, accessToken)
	if err != nil {
		return err
	}

	return nil
}

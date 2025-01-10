package arccredential

import (
	"errors"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type OciCredentialResourceModel struct {
	Name           types.String `tfsdk:"name"`
	User           types.String `tfsdk:"user"`
	IdentityDomain types.String `tfsdk:"identity_domain"`
	Tenancy        types.String `tfsdk:"tenancy"`
	Region         types.String `tfsdk:"region"`
	KeyFile        types.String `tfsdk:"key_file"`
	Id             types.String `tfsdk:"id"`
}

func (m *OciCredentialResourceModel) ToAOCredModel(isResource bool) (*models.Credentials, error) {
	ociCred := models.OciCredentials{}
	ociCred.User = m.User.ValueString()
	ociCred.IdentityDomain = m.IdentityDomain.ValueString()
	ociCred.Tenancy = m.Tenancy.ValueString()
	ociCred.Region = m.Region.ValueString()

	if isResource {
		if !m.KeyFile.IsNull() && !m.KeyFile.IsUnknown() {
			ociKey, err := utils.ReadTextFile(m.KeyFile.ValueString())
			if err != nil {
				return nil, fmt.Errorf("Can't read oci key_file %s: %s", m.KeyFile.ValueString(), err)
			}
			ociCred.KeyFile = ociKey
		} else {
			return nil, errors.New("key_file is missing.")
		}
	}

	cred := models.Credentials{}
	cred.Name = m.Name.ValueStringPointer()
	cred.Credentials.OciCredentials = ociCred
	prod := models.ProvidersOci
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

func (m *OciCredentialResourceModel) ToTerraformModel(cred *models.Credentials) error {
	m.User = types.StringValue(cred.Credentials.OciCredentials.User)
	m.IdentityDomain = types.StringValue(cred.Credentials.OciCredentials.IdentityDomain)
	m.Tenancy = types.StringValue(cred.Credentials.OciCredentials.Tenancy)
	m.Region = types.StringValue(cred.Credentials.OciCredentials.Region)

	if cred.Name == nil {
		return fmt.Errorf("cloud credential can't have empty name")
	}
	m.Name = types.StringValue(*cred.Name)
	m.Id = types.StringValue(*cred.Name)
	return nil
}

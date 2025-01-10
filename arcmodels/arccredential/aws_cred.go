package arccredential

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AwsCredentialResourceModel struct {
	Name      types.String `tfsdk:"name"`
	AccessKey types.String `tfsdk:"access_key"`
	SecretKey types.String `tfsdk:"secret_key"`
	Id        types.String `tfsdk:"id"`
}

func (m *AwsCredentialResourceModel) ToAOCredModel() *models.Credentials {
	awsCred := models.AwsCredentials{}
	awsCred.AccessKey = m.AccessKey.ValueString()
	awsCred.SecretKey = m.SecretKey.ValueString()

	cred := models.Credentials{}
	cred.Name = m.Name.ValueStringPointer()
	cred.Credentials.AwsCredentials = awsCred
	prod := models.ProvidersAws
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred
}

func (m *AwsCredentialResourceModel) ToTerraformModel(cred *models.Credentials) error {
	m.AccessKey = types.StringValue(cred.Credentials.AwsCredentials.AccessKey)
	m.SecretKey = types.StringValue(cred.Credentials.AwsCredentials.SecretKey)
	if cred.Name == nil {
		return fmt.Errorf("cloud credential can't have empty name")
	}
	m.Name = types.StringValue(*cred.Name)
	m.Id = types.StringValue(*cred.Name)
	return nil
}

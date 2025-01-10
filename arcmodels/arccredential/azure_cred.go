package arccredential

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type AzureCredentialResourceModel struct {
	Name           types.String `tfsdk:"name"`
	SubscriptionId types.String `tfsdk:"subscription_id"`
	ClientId       types.String `tfsdk:"client_id"`
	ClientSecret   types.String `tfsdk:"client_secret"`
	TenantId       types.String `tfsdk:"tenant_id"`
	Id             types.String `tfsdk:"id"`
}

func (m *AzureCredentialResourceModel) ToAOCredModel() *models.Credentials {
	azureCred := models.AzureCredentials{}
	azureCred.AzureClientID = m.ClientId.ValueString()
	azureCred.AzureClientSecret = m.ClientSecret.ValueString()
	azureCred.AzureSubscriptionID = m.SubscriptionId.ValueString()
	azureCred.AzureTenantID = m.TenantId.ValueString()

	cred := models.Credentials{}
	cred.Name = m.Name.ValueStringPointer()
	cred.Credentials.AzureCredentials = azureCred
	prod := models.ProvidersAzure
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred
}

func (m *AzureCredentialResourceModel) ToTerraformModel(cred *models.Credentials) error {
	m.ClientId = types.StringValue(cred.Credentials.AzureCredentials.AzureClientID)
	m.ClientSecret = types.StringValue(cred.Credentials.AzureCredentials.AzureClientSecret)
	m.SubscriptionId = types.StringValue(cred.Credentials.AzureCredentials.AzureSubscriptionID)
	m.TenantId = types.StringValue(cred.Credentials.AzureCredentials.AzureTenantID)

	if cred.Name == nil {
		return fmt.Errorf("cloud credential can't have empty name")
	}
	m.Name = types.StringValue(*cred.Name)
	m.Id = types.StringValue(*cred.Name)
	return nil
}

package arccredential

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type GcpCredentialResourceModel struct {
	Name                    types.String `tfsdk:"name"`
	AccountKeyFile          types.String `tfsdk:"account_key_file"`
	Type                    types.String `tfsdk:"type"`
	ProjectId               types.String `tfsdk:"project_id"`
	PrivateKeyId            types.String `tfsdk:"private_key_id"`
	PrivateKey              types.String `tfsdk:"private_key"`
	ClientEmail             types.String `tfsdk:"client_email"`
	ClientId                types.String `tfsdk:"client_id"`
	AuthUri                 types.String `tfsdk:"auth_uri"`
	TokenUri                types.String `tfsdk:"token_uri"`
	AuthProviderX509CertUrl types.String `tfsdk:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       types.String `tfsdk:"client_x509_cert_url"`
	Id                      types.String `tfsdk:"id"`
}

func (m *GcpCredentialResourceModel) ToAOCredModel(isResource bool) (*models.Credentials, error) {
	cred := models.Credentials{
		Name: m.Name.ValueStringPointer(),
	}

	if isResource {
		if !m.AccountKeyFile.IsNull() && !m.AccountKeyFile.IsUnknown() {
			gcpCred, err := loadAccountKeyFile(m.AccountKeyFile.ValueString())
			if err != nil {
				return nil, err
			}
			cred.Credentials.GcpCredentials = *gcpCred
		} else {
			gcpCred := models.GcpCredentials{}
			if !m.Type.IsNull() && !m.Type.IsUnknown() {
				gcpCred.Type = m.Type.ValueString()
			} else {
				return nil, errors.New("type can't be empty if gcp account is not provided")
			}
			if !m.ProjectId.IsNull() && !m.ProjectId.IsUnknown() {
				gcpCred.ProjectID = m.ProjectId.ValueString()
			} else {
				return nil, errors.New("project_id can't be empty if gcp account is not provided")
			}
			if !m.PrivateKeyId.IsNull() && !m.PrivateKeyId.IsUnknown() {
				gcpCred.PrivateKeyID = m.PrivateKeyId.ValueString()
			} else {
				return nil, errors.New("private_key_id can't be empty if gcp account is not provided")
			}
			if !m.PrivateKey.IsNull() && !m.PrivateKey.IsUnknown() {
				gcpCred.PrivateKey = m.PrivateKey.ValueString()
			} else {
				return nil, errors.New("private_key can't be empty if gcp account is not provided")
			}
			if !m.ClientEmail.IsNull() && !m.ClientEmail.IsUnknown() {
				gcpCred.ClientEmail = m.ClientEmail.ValueString()
			} else {
				return nil, errors.New("client_email can't be empty if gcp account is not provided")
			}
			if !m.ClientId.IsNull() && !m.ClientId.IsUnknown() {
				gcpCred.ClientID = m.ClientId.ValueString()
			} else {
				return nil, errors.New("client_id can't be empty if gcp account is not provided")
			}
			if !m.AuthUri.IsNull() && !m.AuthUri.IsUnknown() {
				gcpCred.AuthURI = m.AuthUri.ValueString()
			} else {
				return nil, errors.New("auth_uri can't be empty if gcp account is not provided")
			}
			if !m.TokenUri.IsNull() && !m.TokenUri.IsUnknown() {
				gcpCred.TokenURI = m.TokenUri.ValueString()
			} else {
				return nil, errors.New("token uri can't be empty if gcp account is not provided")
			}
			if !m.AuthProviderX509CertUrl.IsNull() && !m.AuthProviderX509CertUrl.IsUnknown() {
				gcpCred.AuthProviderX509CertURL = m.AuthProviderX509CertUrl.ValueString()
			} else {
				return nil, errors.New("auth_provider_x509_cert_url can't be empty if gcp account is not provided")
			}
			if !m.ClientX509CertUrl.IsNull() && !m.ClientX509CertUrl.IsUnknown() {
				gcpCred.ClientX509CertURL = m.ClientX509CertUrl.ValueString()
			} else {
				return nil, errors.New("client_x509_cert_url can't be empty if gcp account is not provided")
			}
			cred.Credentials.GcpCredentials = gcpCred

		}
	}

	prod := models.ProvidersGcp
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

func (m *GcpCredentialResourceModel) ToTerraformModel(cred *models.Credentials) error {
	m.Type = types.StringValue(cred.Credentials.GcpCredentials.Type)
	m.ProjectId = types.StringValue(cred.Credentials.GcpCredentials.ProjectID)
	m.PrivateKeyId = types.StringValue(cred.Credentials.GcpCredentials.PrivateKeyID)
	m.PrivateKey = types.StringValue(cred.Credentials.GcpCredentials.PrivateKey)
	m.ClientEmail = types.StringValue(cred.Credentials.GcpCredentials.ClientEmail)
	m.ClientId = types.StringValue(cred.Credentials.GcpCredentials.ClientID)
	m.AuthUri = types.StringValue(cred.Credentials.GcpCredentials.AuthURI)
	m.TokenUri = types.StringValue(cred.Credentials.GcpCredentials.TokenURI)
	m.AuthProviderX509CertUrl = types.StringValue(cred.Credentials.GcpCredentials.AuthProviderX509CertURL)
	m.ClientX509CertUrl = types.StringValue(cred.Credentials.GcpCredentials.ClientX509CertURL)

	if cred.Name == nil {
		return fmt.Errorf("cloud credential can't have empty name")
	}
	m.Name = types.StringValue(*cred.Name)
	m.Id = types.StringValue(*cred.Name)
	return nil
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

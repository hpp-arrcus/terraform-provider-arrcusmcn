package arccredential

import (
	"context"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type KvmCredentialResourceModel struct {
	Name       types.String `tfsdk:"name"`
	ServerIp   types.String `tfsdk:"server_ip"`
	UserName   types.String `tfsdk:"user_name"`
	Domain     types.String `tfsdk:"domain"`
	SshKey     types.String `tfsdk:"ssh_key"`
	DataIfName types.List   `tfsdk:"data_if_name"`
	Id         types.String `tfsdk:"id"`
}

func (m *KvmCredentialResourceModel) ToAOCredModel(isResource bool) (*models.Credentials, diag.Diagnostics) {
	kvmCred := models.OnpremCredentials{}
	kvmCred.ServerIP = m.ServerIp.ValueString()
	kvmCred.UserName = m.UserName.ValueString()
	kvmCred.Domain = m.Domain.ValueString()

	if isResource {
		if !m.SshKey.IsNull() && !m.SshKey.IsUnknown() {
			sshKey, err := utils.ReadTextFile(m.SshKey.ValueString())
			if err != nil {
				return nil, diag.Diagnostics{diag.NewErrorDiagnostic("Invalid cloud credential", fmt.Sprintf("Can't read ssh key %s: %s", m.SshKey.ValueString(), err))}
			}
			kvmCred.SSHKey = *sshKey
		} else {
			return nil, diag.Diagnostics{diag.NewErrorDiagnostic("Invalid cloud credential", "ssh key is missing.")}
		}
	}

	if !m.DataIfName.IsNull() && !m.DataIfName.IsUnknown() {
		elements := make([]types.String, 0, len(m.DataIfName.Elements()))
		diag := m.DataIfName.ElementsAs(context.Background(), &elements, false)
		if diag.HasError() {
			return nil, diag
		}
		nameList := make([]string, 0)
		for _, element := range elements {
			nameList = append(nameList, element.ValueString())
		}
		kvmCred.DataIfName = nameList
	}

	cred := models.Credentials{}
	cred.Name = m.Name.ValueStringPointer()
	cred.Credentials.OnpremCredentials = kvmCred
	prod := models.ProvidersOnpremise
	cred.Provider = &prod
	return &cred, nil
}

func (m *KvmCredentialResourceModel) ToTerraformModel(cred *models.Credentials) diag.Diagnostics {
	m.ServerIp = types.StringValue(cred.Credentials.OnpremCredentials.ServerIP)
	m.UserName = types.StringValue(cred.Credentials.OnpremCredentials.UserName)
	m.Domain = types.StringValue(cred.Credentials.OnpremCredentials.Domain)
	if cred.Name == nil {
		return diag.Diagnostics{diag.NewErrorDiagnostic("Invalid cloud credential", "cloud credential can't have empty name")}
	}
	nameList := make([]attr.Value, 0)
	for _, ifName := range cred.Credentials.OnpremCredentials.DataIfName {
		nameList = append(nameList, types.StringValue(ifName))
	}

	ifNameValue, diag := types.ListValue(types.StringType, nameList)
	if diag.HasError() {
		return diag
	}

	m.DataIfName = ifNameValue
	m.Name = types.StringValue(*cred.Name)
	m.Id = types.StringValue(*cred.Name)
	return nil
}

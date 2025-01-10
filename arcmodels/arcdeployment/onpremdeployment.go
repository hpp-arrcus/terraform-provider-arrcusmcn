package arcdeployment

import (
	"context"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type KvmDeploymentResourceModel struct {
	DeploymentResourceModel
	Vcpus           types.Int64  `tfsdk:"vcpus"`
	VMMemory        types.Int64  `tfsdk:"vm_memory"`
	PrivateIp       types.String `tfsdk:"private_ip"`
	PublicIp        types.String `tfsdk:"public_ip"`
	SshPsw          types.String `tfsdk:"ssh_psw"`
	PublicIpBackup  types.String `tfsdk:"public_ip_backup"`
	PrivateIpBackup types.String `tfsdk:"private_ip_backup"`
	AssignPublicIp  types.Bool   `tfsdk:"assign_public_ip"`
	Networks        types.List   `tfsdk:"networks"`
	Id              types.String `tfsdk:"id"`
}

func (m *KvmDeploymentResourceModel) ToAODeploymentModel(includeComputed bool) (*models.Deployment, diag.Diagnostics) {
	deployment, diags := m.ToAOCommonDeploymentModel(models.ProvidersOnpremise, includeComputed)
	if diags.HasError() {
		return nil, diags
	}
	kvmDeployment := models.OnpremDeployment{
		Vcpus:     m.Vcpus.ValueInt64(),
		VMMemory:  m.VMMemory.ValueInt64(),
		PrivateIP: m.PrivateIp.ValueString(),
		PublicIP:  m.PublicIp.ValueString(),
		SSHPsw:    m.SshPsw.ValueString(),
	}

	if !m.PublicIpBackup.IsNull() && !m.PublicIpBackup.IsUnknown() {
		kvmDeployment.PublicIPBackup = m.PublicIpBackup.ValueString()
	}

	if !m.PrivateIpBackup.IsNull() && !m.PrivateIpBackup.IsUnknown() {
		kvmDeployment.PrivateIPBackup = m.PrivateIpBackup.ValueString()
	}

	kvmNetworks := make([]*models.KvmNetwork, 0)
	if !m.Networks.IsNull() && !m.Networks.IsUnknown() {
		networks := make([]KvmNetworkResourceModel, 0, len(m.Networks.Elements()))
		diag := m.Networks.ElementsAs(context.Background(), &networks, false)
		if diag.HasError() {
			return nil, diag.Errors()
		}
		for _, network := range networks {
			networkObj := models.KvmNetwork{}
			if !network.Name.IsNull() && !network.Name.IsUnknown() {
				networkObj.Name = network.Name.ValueString()
			}
			if !network.Network.IsNull() && !network.Network.IsUnknown() {
				networkObj.Network = network.Network.ValueString()
			}
			if !network.PrivateIPDefaultGw.IsNull() && !network.PrivateIPDefaultGw.IsUnknown() {
				networkObj.PrivateIPDefaultGw = network.PrivateIPDefaultGw.ValueString()
			}
			if !network.PrivateIPCidrMask.IsNull() && !network.PrivateIPCidrMask.IsUnknown() {
				networkObj.PrivateIPCidrMask = network.PrivateIPCidrMask.ValueString()
			}
			if diag.HasError() {
				return nil, diag.Errors()
			}

			kvmNetworks = append(kvmNetworks, &networkObj)
		}
	}

	if !m.AssignPublicIp.IsUnknown() {
		kvmDeployment.AssignPublicIP = m.AssignPublicIp.ValueBool()
	} else {
		kvmDeployment.AssignPublicIP = true
	}

	kvmDeployment.Networks = kvmNetworks
	deployment.OnpremDeployment = &kvmDeployment
	return deployment, nil
}

func (m *KvmDeploymentResourceModel) ToTerraformModel(deployment *models.Deployment) diag.Diagnostics {
	diags := m.ToTerraformCommonDeploymentModle(deployment)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	m.Vcpus = types.Int64Value(deployment.OnpremDeployment.Vcpus)
	m.VMMemory = types.Int64Value(deployment.OnpremDeployment.VMMemory)
	m.PrivateIp = types.StringValue(deployment.OnpremDeployment.PrivateIP)
	m.PublicIp = types.StringValue(deployment.OnpremDeployment.PublicIP)
	m.SshPsw = types.StringValue(deployment.OnpremDeployment.SSHPsw)
	m.PublicIpBackup = types.StringValue(deployment.OnpremDeployment.PublicIPBackup)
	m.AssignPublicIp = types.BoolValue(deployment.OnpremDeployment.AssignPublicIP)

	netList := make([]attr.Value, 0)
	for _, network := range deployment.OnpremDeployment.Networks {
		temp := &KvmNetworkResourceModel{}
		value, diags := temp.ToObjectValue(network)
		if diags != nil && diags.HasError() {
			return diags
		}
		netList = append(netList, value)
	}

	m.Networks, diags = types.ListValue(types.ObjectType{}.WithAttributeTypes(KvmNetworkResourceModel{}.AttrType()), netList)
	if diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	return nil
}

type KvmNetworkResourceModel struct {
	Name               types.String `tfsdk:"name"`
	PrivateIPDefaultGw types.String `tfsdk:"private_ip_default_gw"`
	PrivateIPCidrMask  types.String `tfsdk:"private_ip_cidr_mask"`
	Network            types.String `tfsdk:"network"`
}

func (a KvmNetworkResourceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"name":                  types.StringType,
		"private_ip_default_gw": types.StringType,
		"private_ip_cidr_mask":  types.StringType,
		"network":               types.StringType,
	}
}

func (a *KvmNetworkResourceModel) ToObjectValue(kvmNet *models.KvmNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	attrValues := map[string]attr.Value{
		"name":                  types.StringValue(kvmNet.Name),
		"private_ip_default_gw": types.StringValue(kvmNet.PrivateIPDefaultGw),
		"private_ip_cidr_mask":  types.StringValue(kvmNet.PrivateIPCidrMask),
		"network":               types.StringValue(kvmNet.Name),
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

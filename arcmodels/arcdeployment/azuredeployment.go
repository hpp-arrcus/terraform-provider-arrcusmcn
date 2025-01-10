package arcdeployment

import (
	"context"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AzureDeploymentResourceModel struct {
	DeploymentResourceModel
	Region                       types.String `tfsdk:"region"`
	InstanceKey                  types.String `tfsdk:"instance_key"`
	InstanceType                 types.String `tfsdk:"instance_type"`
	ResourceGroup                types.String `tfsdk:"resource_group"`
	Vnet                         types.String `tfsdk:"vnet"`
	Zone                         types.String `tfsdk:"zone"`
	BackupZone                   types.String `tfsdk:"backup_zone"`
	Byoip                        types.String `tfsdk:"byoip"`
	ByoipBackup                  types.String `tfsdk:"byoip_backup"`
	EnableAcceleratedNetworking  types.Bool   `tfsdk:"enable_accelerated_networking"`
	AcceleratedNetworkingEnabled types.Bool   `tfsdk:"accelerated_networking_enabled"`
	AssignPublicIp               types.Bool   `tfsdk:"assign_public_ip"`
	Networks                     types.List   `tfsdk:"networks"`
	Id                           types.String `tfsdk:"id"`
}

func (m *AzureDeploymentResourceModel) ToAODeploymentModel(includeComputed bool, isResource bool) (*models.Deployment, diag.Diagnostics) {
	deployment, diags := m.ToAOCommonDeploymentModel(models.ProvidersAzure, includeComputed)
	if diags.HasError() {
		return nil, diags
	}
	emptyString := ""
	var keyName *string = nil
	var publicKey *string = &emptyString
	var err error = nil
	if isResource {
		publicKey, err = utils.ReadTextFile(m.InstanceKey.ValueString())
		if err != nil {
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert azure model", err.Error())}
		}
		keyName, err = utils.GetPublicKeyName(*publicKey)
		if err != nil {
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert azure model", err.Error())}
		}
	}

	azureDeployment := models.AzureDeployment{
		Region: m.Region.ValueString(),
		InstanceKey: &models.InstanceKey{
			Name:    keyName,
			Content: *publicKey,
		},
		InstanceType:                m.InstanceType.ValueString(),
		ResourceGroup:               m.ResourceGroup.ValueString(),
		Vnet:                        m.Vnet.ValueString(),
		EnableAcceleratedNetworking: m.EnableAcceleratedNetworking.ValueBool(),
	}

	if !m.Zone.IsNull() && !m.Zone.IsUnknown() {
		azureDeployment.Zone = m.Zone.ValueString()
	}

	if !m.BackupZone.IsNull() && !m.BackupZone.IsUnknown() {
		azureDeployment.BackupZone = m.BackupZone.ValueString()
	}

	if !m.Byoip.IsNull() && !m.Byoip.IsUnknown() {
		azureDeployment.Byoip = m.Byoip.ValueString()
	}

	if !m.ByoipBackup.IsNull() && !m.ByoipBackup.IsUnknown() {
		azureDeployment.ByoipBackup = m.ByoipBackup.ValueString()
	}

	if !m.AcceleratedNetworkingEnabled.IsNull() && !m.AcceleratedNetworkingEnabled.IsUnknown() {
		azureDeployment.AcceleratedNetworkingEnabled = m.AcceleratedNetworkingEnabled.ValueBool()
	}

	if !m.AssignPublicIp.IsNull() && !m.AssignPublicIp.IsUnknown() {
		azureDeployment.AssignPublicIP = m.AssignPublicIp.ValueBool()
	}

	azureNetworks := make([]*models.AzureNetwork, 0)
	if !m.Networks.IsNull() && !m.Networks.IsUnknown() {
		networks := make([]AzureNetworkResourceModel, 0, len(m.Networks.Elements()))
		diag := m.Networks.ElementsAs(context.Background(), &networks, false)
		if diag.HasError() {
			return nil, diag.Errors()
		}
		for _, network := range networks {
			networkObj := models.AzureNetwork{}
			if !network.Name.IsNull() && !network.Name.IsUnknown() {
				networkObj.Name = network.Name.ValueString()
			}
			if !network.Subnetwork.IsNull() && !network.Subnetwork.IsUnknown() {
				networkObj.Subnetwork = network.Subnetwork.ValueString()
			}
			if diag.HasError() {
				return nil, diag.Errors()
			}

			azureNetworks = append(azureNetworks, &networkObj)
		}
	}
	azureDeployment.Networks = azureNetworks
	deployment.AzureDeployment = &azureDeployment
	return deployment, nil
}

func (m *AzureDeploymentResourceModel) ToTerraformModel(deployment *models.Deployment) diag.Diagnostics {
	diags := m.ToTerraformCommonDeploymentModle(deployment)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	m.Region = types.StringValue(deployment.AzureDeployment.Region)
	m.InstanceType = types.StringValue(deployment.AzureDeployment.InstanceType)
	m.ResourceGroup = types.StringValue(deployment.AzureDeployment.ResourceGroup)
	m.Vnet = types.StringValue(deployment.AzureDeployment.Vnet)
	m.Zone = types.StringValue(deployment.AzureDeployment.Zone)
	m.BackupZone = types.StringValue(deployment.AzureDeployment.BackupZone)
	m.Byoip = types.StringValue(deployment.AzureDeployment.Byoip)
	m.ByoipBackup = types.StringValue(deployment.AzureDeployment.ByoipBackup)
	m.EnableAcceleratedNetworking = types.BoolValue(deployment.AzureDeployment.EnableAcceleratedNetworking)
	m.AcceleratedNetworkingEnabled = types.BoolValue(deployment.AzureDeployment.AcceleratedNetworkingEnabled)
	m.AssignPublicIp = types.BoolValue(deployment.AzureDeployment.AssignPublicIP)

	netList := make([]attr.Value, 0)
	for _, network := range deployment.AzureDeployment.Networks {
		temp := &AzureNetworkResourceModel{}
		value, diags := temp.ToObjectValue(network)
		if diags != nil && diags.HasError() {
			return diags
		}
		netList = append(netList, value)
	}

	m.Networks, diags = types.ListValue(types.ObjectType{}.WithAttributeTypes(AzureNetworkResourceModel{}.AttrType()), netList)
	if diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	return nil
}

type AzureNetworkResourceModel struct {
	Name       types.String `tfsdk:"name"`
	Subnetwork types.String `tfsdk:"subnetwork"`
}

func (a AzureNetworkResourceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       types.StringType,
		"subnetwork": types.StringType,
	}
}

func (a *AzureNetworkResourceModel) ToObjectValue(azureNet *models.AzureNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	attrValues := map[string]attr.Value{
		"name":       types.StringValue(azureNet.Name),
		"subnetwork": types.StringValue(azureNet.Subnetwork),
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

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

type GcpDeploymentResourceModel struct {
	DeploymentResourceModel
	InstanceKey    types.String `tfsdk:"instance_key"`
	Region         types.String `tfsdk:"region"`
	Zone           types.String `tfsdk:"zone"`
	BackupZone     types.String `tfsdk:"backup_zone"`
	Byoip          types.String `tfsdk:"byoip"`
	ByoipBackup    types.String `tfsdk:"byoip_backup"`
	InstanceType   types.String `tfsdk:"instance_type"`
	AssignPublicIp types.Bool   `tfsdk:"assign_public_ip"`
	Networks       types.List   `tfsdk:"networks"`
	NetworkTags    types.List   `tfsdk:"network_tags"`
	Id             types.String `tfsdk:"id"`
}

func (m *GcpDeploymentResourceModel) ToAODeploymentModel(includeComputed bool, isResource bool) (*models.Deployment, diag.Diagnostics) {
	deployment, diags := m.ToAOCommonDeploymentModel(models.ProvidersGcp, includeComputed)
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
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert gcp model", err.Error())}
		}
		keyName, err = utils.GetPublicKeyName(*publicKey)
		if err != nil {
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert gcp model", err.Error())}
		}
	}
	gcpDeployment := models.GcpDeployment{
		Region: m.Region.ValueString(),
		InstanceKey: &models.InstanceKey{
			Name:    keyName,
			Content: *publicKey,
		},
		InstanceType: m.InstanceType.ValueString(),
	}

	if !m.Zone.IsNull() && !m.Zone.IsUnknown() {
		gcpDeployment.Zone = m.Zone.ValueString()
	}

	if !m.BackupZone.IsNull() && !m.BackupZone.IsUnknown() {
		gcpDeployment.BackupZone = m.BackupZone.ValueString()
	}

	if !m.Byoip.IsNull() && !m.Byoip.IsUnknown() {
		gcpDeployment.Byoip = m.Byoip.ValueString()
	}

	if !m.ByoipBackup.IsNull() && !m.ByoipBackup.IsUnknown() {
		gcpDeployment.ByoipBackup = m.ByoipBackup.ValueString()
	}

	gcpNetworks := make([]*models.GcpNetwork, 0)
	if !m.Networks.IsNull() && !m.Networks.IsUnknown() {
		networks := make([]GcpNetworkResourceModel, 0, len(m.Networks.Elements()))
		diag := m.Networks.ElementsAs(context.Background(), &networks, false)
		if diag.HasError() {
			return nil, diag.Errors()
		}
		for _, network := range networks {
			networkObj := models.GcpNetwork{}
			if !network.Name.IsNull() && !network.Name.IsUnknown() {
				networkObj.Name = network.Name.ValueString()
			}
			if !network.Network.IsNull() && !network.Network.IsUnknown() {
				networkObj.Network = network.Network.ValueString()
			}
			if !network.Subnetwork.IsNull() && !network.Subnetwork.IsUnknown() {
				networkObj.Subnetwork = network.Subnetwork.ValueString()
			}
			if diag.HasError() {
				return nil, diag.Errors()
			}

			gcpNetworks = append(gcpNetworks, &networkObj)
		}
	}

	if !m.AssignPublicIp.IsUnknown() {
		gcpDeployment.AssignPublicIP = m.AssignPublicIp.ValueBool()
	} else {
		gcpDeployment.AssignPublicIP = true
	}

	if !m.NetworkTags.IsNull() && !m.NetworkTags.IsUnknown() {
		elements := make([]types.String, 0, len(m.NetworkTags.Elements()))
		diag := m.NetworkTags.ElementsAs(context.Background(), &elements, false)
		if diag.HasError() {
			return nil, diag
		}
		tags := make([]string, 0)
		for _, element := range elements {
			tags = append(tags, element.ValueString())
		}
		gcpDeployment.NetworkTags = tags

	}
	gcpDeployment.Networks = gcpNetworks
	deployment.GcpDeployment = &gcpDeployment
	return deployment, nil
}

func (m *GcpDeploymentResourceModel) ToTerraformModel(deployment *models.Deployment) diag.Diagnostics {
	diags := m.ToTerraformCommonDeploymentModle(deployment)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	m.Region = types.StringValue(deployment.GcpDeployment.Region)
	m.InstanceType = types.StringValue(deployment.GcpDeployment.InstanceType)
	m.Zone = types.StringValue(deployment.GcpDeployment.Zone)
	m.BackupZone = types.StringValue(deployment.GcpDeployment.BackupZone)
	m.Byoip = types.StringValue(deployment.GcpDeployment.Byoip)
	m.ByoipBackup = types.StringValue(deployment.GcpDeployment.ByoipBackup)
	m.AssignPublicIp = types.BoolValue(deployment.GcpDeployment.AssignPublicIP)

	netList := make([]attr.Value, 0)
	for _, network := range deployment.GcpDeployment.Networks {
		temp := &GcpNetworkResourceModel{}
		value, diags := temp.ToObjectValue(network)
		if diags != nil && diags.HasError() {
			return diags
		}
		netList = append(netList, value)
	}

	m.Networks, diags = types.ListValue(types.ObjectType{}.WithAttributeTypes(GcpNetworkResourceModel{}.AttrType()), netList)
	if diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)

	tagList := make([]attr.Value, 0)
	for _, tag := range deployment.GcpDeployment.NetworkTags {
		tagList = append(tagList, types.StringValue(tag))
	}

	tagValues, diags := types.ListValue(types.StringType, tagList)
	if diags != nil && diags.HasError() {
		return diags
	}

	m.NetworkTags = tagValues
	return nil
}

type GcpNetworkResourceModel struct {
	Name       types.String `tfsdk:"name"`
	Project    types.String `tfsdk:"project"`
	Network    types.String `tfsdk:"network"`
	Subnetwork types.String `tfsdk:"subnetwork"`
}

func (a GcpNetworkResourceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"name":       types.StringType,
		"project":    types.StringType,
		"network":    types.StringType,
		"subnetwork": types.StringType,
	}
}

func (a *GcpNetworkResourceModel) ToObjectValue(gcpNet *models.GcpNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	attrValues := map[string]attr.Value{
		"name":       types.StringValue(gcpNet.Name),
		"project":    types.StringValue(gcpNet.Project),
		"network":    types.StringValue(gcpNet.Network),
		"subnetwork": types.StringValue(gcpNet.Subnetwork),
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

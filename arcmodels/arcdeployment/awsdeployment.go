package arcdeployment

import (
	"context"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AwsDeploymentResourceModel struct {
	DeploymentResourceModel
	InstanceKey    types.String `tfsdk:"instance_key"`
	Region         types.String `tfsdk:"region"`
	Byoip          types.String `tfsdk:"byoip"`
	ByoipBackup    types.String `tfsdk:"byoip_backup"`
	Zone           types.String `tfsdk:"zone"`
	BackupZone     types.String `tfsdk:"backup_zone"`
	VpcId          types.String `tfsdk:"vpc_id"`
	InstanceType   types.String `tfsdk:"instance_type"`
	AssignPublicIP types.Bool   `tfsdk:"assign_public_ip"`
	Networks       types.List   `tfsdk:"networks"`
	Id             types.String `tfsdk:"id"`
}

func (m *AwsDeploymentResourceModel) ToAODeploymentModel(includeComputed bool) (*models.Deployment, diag.Diagnostics) {
	deployment, diag := m.ToAOCommonDeploymentModel(models.ProvidersAws, includeComputed)
	if diag.HasError() {
		return nil, diag
	}
	awsDeployment := &models.AwsDeployment{
		InstanceKey:  m.InstanceKey.ValueString(),
		InstanceType: m.InstanceType.ValueString(),
		Region:       m.Region.ValueString(),
		VpcID:        m.VpcId.ValueString(),
		Zone:         m.Zone.ValueString(),
	}

	if !m.Byoip.IsUnknown() {
		awsDeployment.Byoip = m.Byoip.ValueString()
	}
	if !m.ByoipBackup.IsUnknown() {
		awsDeployment.ByoipBackup = m.ByoipBackup.ValueString()
	}
	if !m.BackupZone.IsUnknown() {
		awsDeployment.BackupZone = m.BackupZone.ValueString()
	}

	if !m.AssignPublicIP.IsUnknown() {
		awsDeployment.AssignPublicIP = m.AssignPublicIP.ValueBool()
	} else {
		awsDeployment.AssignPublicIP = true
	}

	awsNetworks := make([]*models.AwsNetwork, 0)
	if !m.Networks.IsNull() {
		networks := make([]AwsNetworkResourceModel, 0, len(m.Networks.Elements()))
		diag := m.Networks.ElementsAs(context.Background(), &networks, false)
		if diag.HasError() {
			return nil, diag.Errors()
		}
		for _, network := range networks {
			networkObj := models.AwsNetwork{}
			if !network.Name.IsNull() && !network.Name.IsUnknown() {
				networkObj.Name = network.Name.ValueString()
			}
			if !network.SubnetA.IsNull() && !network.SubnetA.IsUnknown() {
				networkObj.SubnetA = network.SubnetA.ValueString()
			}
			if !network.SubnetB.IsNull() && !network.SubnetB.IsUnknown() {
				networkObj.SubnetB = network.SubnetB.ValueString()
			}
			if !network.PrivateSubnetRouteTable.IsNull() && !network.PrivateSubnetRouteTable.IsUnknown() {
				networkObj.PrivateSubnetRouteTable = network.PrivateSubnetRouteTable.ValueString()
			}
			if diag.HasError() {
				return nil, diag.Errors()
			}

			awsNetworks = append(awsNetworks, &networkObj)
		}
	}
	awsDeployment.Networks = awsNetworks
	deployment.AwsDeployment = awsDeployment
	return deployment, nil
}

func (m *AwsDeploymentResourceModel) ToTerraformModel(deployment *models.Deployment) diag.Diagnostics {
	diags := m.ToTerraformCommonDeploymentModle(deployment)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	m.InstanceKey = types.StringValue(deployment.AwsDeployment.InstanceKey)
	m.Region = types.StringValue(deployment.AwsDeployment.Region)
	m.Byoip = types.StringValue(deployment.AwsDeployment.Byoip)
	m.ByoipBackup = types.StringValue(deployment.AwsDeployment.ByoipBackup)
	m.Zone = types.StringValue(deployment.AwsDeployment.Zone)
	m.BackupZone = types.StringValue(deployment.AwsDeployment.BackupZone)
	m.VpcId = types.StringValue(deployment.AwsDeployment.VpcID)
	m.InstanceType = types.StringValue(deployment.AwsDeployment.InstanceType)
	m.AssignPublicIP = types.BoolValue(deployment.AwsDeployment.AssignPublicIP)

	netList := make([]attr.Value, 0)

	for _, network := range deployment.AwsDeployment.Networks {
		temp := &AwsNetworkResourceModel{}
		value, diags := temp.ToObjectValue(network)
		if diags != nil && diags.HasError() {
			return diags
		}
		netList = append(netList, value)
	}

	m.Networks, diags = types.ListValue(types.ObjectType{}.WithAttributeTypes(AwsNetworkResourceModel{}.AttrType()), netList)
	if diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	return nil
}

type AwsNetworkResourceModel struct {
	Name                    types.String `tfsdk:"name"`
	SubnetA                 types.String `tfsdk:"subnet_a"`
	SubnetB                 types.String `tfsdk:"subnet_b"`
	PrivateSubnetRouteTable types.String `tfsdk:"private_subnet_route_table"`
}

func (a AwsNetworkResourceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"name":                       types.StringType,
		"subnet_a":                   types.StringType,
		"subnet_b":                   types.StringType,
		"private_subnet_route_table": types.StringType,
	}
}
func (a *AwsNetworkResourceModel) ToObjectValue(awsNet *models.AwsNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	attrValues := map[string]attr.Value{
		"name":                       types.StringValue(awsNet.Name),
		"subnet_a":                   types.StringValue(awsNet.SubnetA),
		"subnet_b":                   types.StringValue(awsNet.SubnetB),
		"private_subnet_route_table": types.StringValue(awsNet.PrivateSubnetRouteTable),
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

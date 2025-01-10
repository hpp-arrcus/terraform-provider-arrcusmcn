package arcdeployment

import (
	"context"
	"strconv"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type OciDeploymentResourceModel struct {
	DeploymentResourceModel
	InstanceKey              types.String `tfsdk:"instance_key"`
	Region                   types.String `tfsdk:"region"`
	AvailabilityDomain       types.String `tfsdk:"availability_domain"`
	BackupAvailabilityDomain types.String `tfsdk:"backup_availability_domain"`
	Compartment              types.String `tfsdk:"compartment"`
	ImageCompartment         types.String `tfsdk:"image_compartment"`
	ComputeShape             types.String `tfsdk:"compute_shape"`
	ComputeCpus              types.Int32  `tfsdk:"compute_cpus"`
	ComputeMemoryInGbs       types.Int32  `tfsdk:"compute_memory_in_gbs"`
	EnableFirewall           types.Bool   `tfsdk:"enable_firewall"`
	Byoip                    types.String `tfsdk:"byoip"`
	ByoipBackup              types.String `tfsdk:"byoip_backup"`
	Networks                 types.List   `tfsdk:"networks"`
	Id                       types.String `tfsdk:"id"`
}

func (m *OciDeploymentResourceModel) ToAODeploymentModel(includeComputed bool, isResource bool) (*models.Deployment, diag.Diagnostics) {
	deployment, diags := m.ToAOCommonDeploymentModel(models.ProvidersOci, includeComputed)
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
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert oci model", err.Error())}
		}
		keyName, err = utils.GetPublicKeyName(*publicKey)
		if err != nil {
			return nil, diag.Diagnostics{diag.NewWarningDiagnostic("failed to convert oci model", err.Error())}
		}
	}
	ociDeployment := models.OciDeployment{
		Region: m.Region.ValueString(),
		SSHAuthorizedKeys: &models.InstanceKey{
			Name:    keyName,
			Content: *publicKey,
		},
		AvailabilityDomain: m.AvailabilityDomain.ValueString(),
		Compartment:        m.Compartment.ValueString(),
		ImageCompartment:   m.ImageCompartment.ValueString(),
		ComputeShape:       m.ComputeShape.ValueString(),
	}

	if !m.ComputeMemoryInGbs.IsNull() && !m.ComputeMemoryInGbs.IsUnknown() {
		ociDeployment.ComputeMemoryInGbs = strconv.Itoa(int(m.ComputeMemoryInGbs.ValueInt32()))
	}

	if !m.ComputeCpus.IsNull() && !m.ComputeCpus.IsUnknown() {
		ociDeployment.ComputeCpus = strconv.Itoa(int(m.ComputeCpus.ValueInt32()))
	}

	if !m.Byoip.IsNull() && !m.Byoip.IsUnknown() {
		ociDeployment.Byoip = m.Byoip.ValueString()
	}

	if !m.BackupAvailabilityDomain.IsNull() && !m.BackupAvailabilityDomain.IsUnknown() {
		ociDeployment.BackupAvailabilityDomain = m.BackupAvailabilityDomain.ValueString()
	}

	if !m.ByoipBackup.IsNull() && !m.ByoipBackup.IsUnknown() {
		ociDeployment.ByoipBackup = m.ByoipBackup.ValueString()
	}

	if !m.EnableFirewall.IsNull() && !m.EnableFirewall.IsUnknown() {
		ociDeployment.EnableFirewall = m.EnableFirewall.ValueBoolPointer()
	} else {
		trueVal := true
		ociDeployment.EnableFirewall = &trueVal
	}

	ociNetworks := make([]*models.OciNetwork, 0)
	if !m.Networks.IsNull() && !m.Networks.IsUnknown() {
		networks := make([]OciNetworkResourceModel, 0, len(m.Networks.Elements()))
		diag := m.Networks.ElementsAs(context.Background(), &networks, false)
		if diag.HasError() {
			return nil, diag.Errors()
		}
		for _, network := range networks {
			networkObj := models.OciNetwork{}
			if !network.SubnetName.IsNull() && !network.SubnetName.IsUnknown() {
				networkObj.SubnetName = network.SubnetName.ValueString()
			}
			if !network.VcnOcid.IsNull() && !network.VcnOcid.IsUnknown() {
				networkObj.VcnOcid = network.VcnOcid.ValueString()
			}
			if !network.VcnName.IsNull() && !network.VcnName.IsUnknown() {
				networkObj.VcnName = network.VcnName.ValueString()
			}
			if !network.SubnetOcid.IsNull() && !network.SubnetOcid.IsUnknown() {
				networkObj.SubnetOcid = network.SubnetOcid.ValueString()
			}
			if !network.SubnetAccess.IsNull() && !network.SubnetAccess.IsUnknown() {
				networkObj.SubnetAccess = network.SubnetAccess.ValueString()
			}
			if diag.HasError() {
				return nil, diag.Errors()
			}

			ociNetworks = append(ociNetworks, &networkObj)
		}
	}

	ociDeployment.Networks = ociNetworks
	deployment.OciDeployment = &ociDeployment
	return deployment, nil
}

func (m *OciDeploymentResourceModel) ToTerraformModel(deployment *models.Deployment) diag.Diagnostics {
	diags := m.ToTerraformCommonDeploymentModle(deployment)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	m.Region = types.StringValue(deployment.OciDeployment.Region)
	m.AvailabilityDomain = types.StringValue(deployment.OciDeployment.AvailabilityDomain)
	m.BackupAvailabilityDomain = types.StringValue(deployment.OciDeployment.BackupAvailabilityDomain)
	m.Compartment = types.StringValue(deployment.OciDeployment.Compartment)
	m.ImageCompartment = types.StringValue(deployment.OciDeployment.ImageCompartment)
	m.ComputeShape = types.StringValue(deployment.OciDeployment.ComputeShape)
	cpus, err := strconv.ParseInt(deployment.OciDeployment.ComputeCpus, 10, 32)
	if err != nil {
		panic(err)
	}
	m.ComputeCpus = types.Int32Value(int32(cpus))
	mem, err := strconv.ParseInt(deployment.OciDeployment.ComputeMemoryInGbs, 10, 32)
	if err != nil {
		panic(err)
	}
	m.ComputeMemoryInGbs = types.Int32Value(int32(mem))
	m.EnableFirewall = types.BoolPointerValue(deployment.OciDeployment.EnableFirewall)
	m.Byoip = types.StringValue(deployment.OciDeployment.Byoip)
	m.ByoipBackup = types.StringValue(deployment.OciDeployment.ByoipBackup)

	netList := make([]attr.Value, 0)
	for _, network := range deployment.OciDeployment.Networks {
		temp := &OciNetworkResourceModel{}
		value, diags := temp.ToObjectValue(network)
		if diags != nil && diags.HasError() {
			return diags
		}
		netList = append(netList, value)
	}

	m.Networks, diags = types.ListValue(types.ObjectType{}.WithAttributeTypes(OciNetworkResourceModel{}.AttrType()), netList)
	if diags.HasError() {
		return diags
	}
	m.Id = types.StringValue(*deployment.Name)
	return nil
}

type OciNetworkResourceModel struct {
	SubnetName   types.String `tfsdk:"subnet_name"`
	VcnOcid      types.String `tfsdk:"vcn_ocid"`
	VcnName      types.String `tfsdk:"vcn_name"`
	SubnetOcid   types.String `tfsdk:"subnet_ocid"`
	SubnetAccess types.String `tfsdk:"subnet_access"`
}

func (a OciNetworkResourceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"subnet_name":   types.StringType,
		"vcn_ocid":      types.StringType,
		"vcn_name":      types.StringType,
		"subnet_ocid":   types.StringType,
		"subnet_access": types.StringType,
	}
}

func (a *OciNetworkResourceModel) ToObjectValue(ociNet *models.OciNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	attrValues := map[string]attr.Value{
		"subnet_name":   types.StringValue(ociNet.SubnetName),
		"vcn_ocid":      types.StringValue(ociNet.VcnOcid),
		"vcn_name":      types.StringValue(ociNet.VcnOcid),
		"subnet_ocid":   types.StringValue(ociNet.SubnetOcid),
		"subnet_access": types.StringValue(ociNet.SubnetAccess),
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

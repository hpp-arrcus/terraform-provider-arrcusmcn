package arcdeployment

import (
	"context"
	"log"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var ObjectAsOptions = basetypes.ObjectAsOptions{
	UnhandledNullAsEmpty:    true,
	UnhandledUnknownAsEmpty: false,
}

type DeploymentResourceModel struct {
	Name                    types.String  `tfsdk:"name"`
	CredentialId            types.String  `tfsdk:"credential_id"`
	CredentialName          types.String  `tfsdk:"credential_name"`
	EnableHighAvailability  types.Bool    `tfsdk:"enable_high_availability"`
	ArcOrchIP               types.String  `tfsdk:"arc_orch_ip"`
	Action                  types.String  `tfsdk:"action"`
	Status                  types.String  `tfsdk:"status"`
	StatusID                types.Int64   `tfsdk:"status_id"`
	ArcedgeARole            types.String  `tfsdk:"arcedge_a_role"`
	ArcedgeBRole            types.String  `tfsdk:"arcedge_b_role"`
	ArcedgeAStatus          types.String  `tfsdk:"arcedge_a_status"`
	ArcedgeAStatusId        types.Int64   `tfsdk:"arcedge_a_status_id"`
	ArcedgeBStatus          types.String  `tfsdk:"arcedge_b_status"`
	ArcedgeBStatusId        types.Int64   `tfsdk:"arcedge_b_status_id"`
	ArcedgeAIP              types.String  `tfsdk:"arcedge_a_ip"`
	ActiveIPGateway         types.String  `tfsdk:"active_ip_gateway"`
	ArcedgeAPrivateIP       types.String  `tfsdk:"arcedge_a_private_ip"`
	ArcedgeBIP              types.String  `tfsdk:"arcedge_b_ip"`
	ArcedgeBPrivateIP       types.String  `tfsdk:"arcedge_b_private_ip"`
	PrivateCidr             types.String  `tfsdk:"private_cidr"`
	IngressSg               types.String  `tfsdk:"ingress_sg"`
	HubNumber               types.Int64   `tfsdk:"hub_number"`
	CoordinatesLat          types.Float64 `tfsdk:"coordinates_lat"`
	CoordinatesLong         types.Float64 `tfsdk:"coordinates_long"`
	SourceImage             types.Object  `tfsdk:"source_image"`
	LatestAvailableImage    types.Object  `tfsdk:"latest_available_image"`
	ActiveNetworkInterfaces types.List    `tfsdk:"active_network_interfaces"`
	BackupNetworkInterfaces types.List    `tfsdk:"backup_network_interfaces"`
}

type ArcedgeImageModel struct {
	Name     types.String `tfsdk:"name"`
	Version  types.String `tfsdk:"version"`
	ImageId  types.String `tfsdk:"image_id"`
	Provider types.String `tfsdk:"provider"`
}

func (a ArcedgeImageModel) AttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"image_id": types.StringType,
		"version":  types.StringType,
		"provider": types.StringType,
		"name":     types.StringType,
	}
}

func (a ArcedgeImageModel) ToAttrValues(ai *models.ArcedgeImage) map[string]attr.Value {
	return map[string]attr.Value{
		"image_id": types.StringValue(*ai.ImageID),
		"version":  types.StringValue(ai.Version),
		"provider": types.StringValue(string(ai.Provider)),
		"name":     types.StringValue(ai.Name),
	}
}

func (a *ArcedgeImageModel) ToAOModel() *models.ArcedgeImage {
	aeImage := models.ArcedgeImage{}
	aeImage.ImageID = a.ImageId.ValueStringPointer()
	if !a.Name.IsNull() && !a.Name.IsUnknown() {
		aeImage.Name = a.Name.ValueString()
	}
	if !a.Version.IsNull() && !a.Version.IsUnknown() {
		aeImage.Version = a.Version.ValueString()
	}
	return &aeImage
}

type NetworkInterfacesModel struct {
	Name                 types.String `tfsdk:"name"`
	PrivateIpv4Address   types.String `tfsdk:"private_ipv4_address"`
	PublicIpv4Address    types.String `tfsdk:"public_ipv4_address"`
	GlobalIpv6Address    types.String `tfsdk:"global_ipv6_address"`
	LinklocalIpv6Address types.String `tfsdk:"linklocal_ipv6_address"`
	MacAddress           types.String `tfsdk:"mac_address"`
	AdapterType          types.String `tfsdk:"adapter_type"`
	PrivateIpv4Pfxlen    types.Int64  `tfsdk:"private_ipv4_pfxlen"`
	PublicIpv4Pfxlen     types.Int64  `tfsdk:"public_ipv4_pfxlen"`
	LinklocalIpv6Pfxlen  types.Int64  `tfsdk:"linklocal_ipv6_pfxlen"`
	GlobalIpv6Pfxlen     types.Int64  `tfsdk:"global_ipv6_pfxlen"`
	AwsInterface         types.Object `tfsdk:"aws_interface"`
	AzureInterface       types.Object `tfsdk:"azure_interface"`
	OciInterface         types.Object `tfsdk:"oci_interface"`
}

func (a *NetworkInterfacesModel) ToAOModel() (*models.NetworkInterface, diag.Diagnostics) {
	res := models.NetworkInterface{}
	ctx := context.Background()
	if !a.Name.IsNull() && !a.Name.IsUnknown() {
		res.Name = a.Name.ValueString()
	}
	if !a.PrivateIpv4Address.IsNull() && !a.PrivateIpv4Address.IsUnknown() {
		res.PrivateIPV4Address = a.PrivateIpv4Address.ValueString()
	}
	if !a.PublicIpv4Address.IsNull() && !a.PublicIpv4Address.IsUnknown() {
		res.PublicIPV4Address = a.PublicIpv4Address.ValueString()
	}
	if !a.GlobalIpv6Address.IsNull() && !a.GlobalIpv6Address.IsUnknown() {
		res.GlobalIPV6Address = a.GlobalIpv6Address.ValueString()
	}
	if !a.LinklocalIpv6Address.IsNull() && !a.LinklocalIpv6Address.IsUnknown() {
		res.LinklocalIPV6Address = a.LinklocalIpv6Address.ValueString()
	}
	if !a.MacAddress.IsNull() && !a.MacAddress.IsUnknown() {
		res.MacAddress = a.MacAddress.ValueString()
	}
	if !a.AdapterType.IsNull() && !a.AdapterType.IsUnknown() {
		res.AdapterType = a.AdapterType.ValueString()
	}
	if !a.PrivateIpv4Pfxlen.IsNull() && !a.PrivateIpv4Pfxlen.IsUnknown() {
		res.PrivateIPV4Pfxlen = a.PrivateIpv4Pfxlen.ValueInt64()
	}
	if !a.PublicIpv4Pfxlen.IsNull() && !a.PublicIpv4Pfxlen.IsUnknown() {
		res.PublicIPV4Pfxlen = a.PublicIpv4Pfxlen.ValueInt64()
	}
	if !a.LinklocalIpv6Pfxlen.IsNull() && !a.LinklocalIpv6Pfxlen.IsUnknown() {
		res.LinklocalIPV6Pfxlen = a.LinklocalIpv6Pfxlen.ValueInt64()
	}
	if !a.GlobalIpv6Pfxlen.IsNull() && !a.GlobalIpv6Pfxlen.IsUnknown() {
		res.GlobalIPV6Pfxlen = a.GlobalIpv6Pfxlen.ValueInt64()
	}
	if !a.AwsInterface.IsNull() && !a.AwsInterface.IsUnknown() {
		awsIfModel := AwsInterfaceModel{}
		diag := a.AwsInterface.As(ctx, &awsIfModel, ObjectAsOptions)
		if diag.HasError() {
			return nil, diag
		}
		m, diag := awsIfModel.ToAOModel()
		if diag.HasError() {
			return nil, diag
		}
		res.AwsInterface = m
	}
	if !a.AzureInterface.IsNull() && !a.AzureInterface.IsUnknown() {
		azureIfModel := AzureInterfaceModel{}
		diag := a.AzureInterface.As(ctx, &azureIfModel, ObjectAsOptions)
		if diag.HasError() {
			return nil, diag
		}
		m, diag := azureIfModel.ToAOModel()
		if diag.HasError() {
			return nil, diag
		}
		res.AzureInterface = m
	}
	if !a.OciInterface.IsNull() && !a.OciInterface.IsUnknown() {
		ociIfModel := OciInterfaceModel{}
		diag := a.OciInterface.As(ctx, &ociIfModel, ObjectAsOptions)
		if diag.HasError() {
			return nil, diag
		}
		m, diag := ociIfModel.ToAOModel()
		if diag.HasError() {
			return nil, diag
		}
		res.OciInterface = m
	}
	return &res, nil
}

func (a NetworkInterfacesModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"name":                   types.StringType,
		"private_ipv4_address":   types.StringType,
		"public_ipv4_address":    types.StringType,
		"global_ipv6_address":    types.StringType,
		"linklocal_ipv6_address": types.StringType,
		"mac_address":            types.StringType,
		"adapter_type":           types.StringType,
		"private_ipv4_pfxlen":    types.Int64Type,
		"public_ipv4_pfxlen":     types.Int64Type,
		"linklocal_ipv6_pfxlen":  types.Int64Type,
		"global_ipv6_pfxlen":     types.Int64Type,
		"aws_interface":          types.ObjectType{}.WithAttributeTypes(AwsInterfaceModel{}.AttrType()),
		"azure_interface":        types.ObjectType{}.WithAttributeTypes(AzureInterfaceModel{}.AttrType()),
		"oci_interface":          types.ObjectType{}.WithAttributeTypes(OciInterfaceModel{}.AttrType()),
	}
}

func (a NetworkInterfacesModel) ToObjectValue(netIf *models.NetworkInterface) (*basetypes.ObjectValue, diag.Diagnostics) {
	var attrValues map[string]attr.Value
	if netIf != nil {
		attrValues = map[string]attr.Value{
			"name":                   types.StringValue(netIf.Name),
			"private_ipv4_address":   types.StringValue(netIf.PrivateIPV4Address),
			"public_ipv4_address":    types.StringValue(netIf.PublicIPV4Address),
			"global_ipv6_address":    types.StringValue(netIf.GlobalIPV6Address),
			"linklocal_ipv6_address": types.StringValue(netIf.LinklocalIPV6Address),
			"mac_address":            types.StringValue(netIf.MacAddress),
			"adapter_type":           types.StringValue(netIf.AdapterType),
			"private_ipv4_pfxlen":    types.Int64Value(netIf.PrivateIPV4Pfxlen),
			"public_ipv4_pfxlen":     types.Int64Value(netIf.PublicIPV4Pfxlen),
			"linklocal_ipv6_pfxlen":  types.Int64Value(netIf.LinklocalIPV6Pfxlen),
			"global_ipv6_pfxlen":     types.Int64Value(netIf.GlobalIPV6Pfxlen),
		}

		if netIf.AwsInterface != nil {
			awsIfObjVal, diags := AwsInterfaceModel{}.ToObjectValue(netIf.AwsInterface)
			if diags != nil && diags.HasError() {
				return nil, diags
			}
			attrValues["aws_interface"] = awsIfObjVal
		} else {
			attrValues["aws_interface"] = types.ObjectNull(AwsInterfaceModel{}.AttrType())
		}

		if netIf.AzureInterface != nil {
			azIfObjVal, diags := AzureInterfaceModel{}.ToObjectValue(netIf.AzureInterface)
			if diags != nil && diags.HasError() {
				return nil, diags
			}
			attrValues["azure_interface"] = azIfObjVal
		} else {
			attrValues["azure_interface"] = types.ObjectNull(AzureInterfaceModel{}.AttrType())
		}

		if netIf.OciInterface != nil {
			ociIfObjVal, diags := OciInterfaceModel{}.ToObjectValue(netIf.OciInterface)
			if diags != nil && diags.HasError() {
				return nil, diags
			}
			attrValues["oci_interface"] = ociIfObjVal
		} else {
			attrValues["oci_interface"] = types.ObjectNull(OciInterfaceModel{}.AttrType())
		}

	} else {
		attrValues = map[string]attr.Value{
			"name":                   types.StringNull(),
			"private_ipv4_address":   types.StringNull(),
			"public_ipv4_address":    types.StringNull(),
			"global_ipv6_address":    types.StringNull(),
			"linklocal_ipv6_address": types.StringNull(),
			"mac_address":            types.StringNull(),
			"adapter_type":           types.StringNull(),
			"private_ipv4_pfxlen":    types.Int64Null(),
			"public_ipv4_pfxlen":     types.Int64Null(),
			"linklocal_ipv6_pfxlen":  types.Int64Null(),
			"global_ipv6_pfxlen":     types.Int64Null(),
			"aws_interface":          types.ObjectNull(AwsInterfaceModel{}.AttrType()),
			"azure_interface":        types.ObjectNull(AzureInterfaceModel{}.AttrType()),
			"oci_interface":          types.ObjectNull(OciInterfaceModel{}.AttrType()),
		}
	}

	res, diags := types.ObjectValue(a.AttrType(), attrValues)
	if diags != nil && diags.HasError() {
		return nil, diags
	}
	return &res, nil
}

type AwsInterfaceModel struct {
	InterfaceId    types.String `tfsdk:"interface_id"`
	RouteTableId   types.String `tfsdk:"route_table_id"`
	SubnetId       types.String `tfsdk:"subnet_id"`
	SecurityGroups types.List   `tfsdk:"security_groups"`
}

func (a *AwsInterfaceModel) ToAOModel() (*models.AwsInterface, diag.Diagnostics) {
	res := models.AwsInterface{}
	if !a.InterfaceId.IsNull() && !a.InterfaceId.IsUnknown() {
		res.InterfaceID = a.InterfaceId.ValueString()
	}
	if !a.SubnetId.IsNull() && !a.SubnetId.IsUnknown() {
		res.SubnetID = a.SubnetId.ValueString()
	}
	if !a.SecurityGroups.IsNull() && !a.SecurityGroups.IsUnknown() {
		elements := make([]types.String, 0, len(a.SecurityGroups.Elements()))
		diag := a.SecurityGroups.ElementsAs(context.Background(), &elements, false)
		if diag.HasError() {
			return nil, diag
		}
		nsg := make([]string, 0)
		for _, element := range elements {
			nsg = append(nsg, element.ValueString())
		}
		res.SecurityGroups = nsg
	}

	return &res, nil
}

func (a AwsInterfaceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"interface_id":    types.StringType,
		"route_table_id":  types.StringType,
		"subnet_id":       types.StringType,
		"security_groups": types.ListType{}.WithElementType(types.StringType),
	}
}

func (a AwsInterfaceModel) ToObjectValue(awsIf *models.AwsInterface) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	var attrValues map[string]attr.Value
	if awsIf != nil {
		sgList := make([]attr.Value, 0)
		for _, sg := range awsIf.SecurityGroups {
			sgList = append(sgList, types.StringValue(sg))
		}

		sgValues, diags := types.ListValue(types.StringType, sgList)
		if diags != nil && diags.HasError() {
			return nil, diags
		}

		attrValues = map[string]attr.Value{
			"interface_id":    types.StringValue(awsIf.InterfaceID),
			"route_table_id":  types.StringValue(awsIf.RouteTableID),
			"subnet_id":       types.StringValue(awsIf.SubnetID),
			"security_groups": sgValues,
		}
	} else {
		attrValues = map[string]attr.Value{
			"interface_id":    types.StringNull(),
			"route_table_id":  types.StringNull(),
			"subnet_id":       types.StringNull(),
			"security_groups": types.ListNull(types.StringType),
		}
	}
	res, diags := types.ObjectValue(attrTypes, attrValues)
	if diags != nil && diags.HasError() {
		return nil, diags
	}
	return &res, nil
}

type AzureInterfaceModel struct {
	InterfaceId    types.String `tfsdk:"interface_id"`
	RouteTableName types.String `tfsdk:"route_table_name"`
}

func (a *AzureInterfaceModel) ToAOModel() (*models.AzureInterface, diag.Diagnostics) {
	res := models.AzureInterface{}
	if !a.InterfaceId.IsNull() && !a.InterfaceId.IsUnknown() {
		res.InterfaceID = a.InterfaceId.ValueString()
	}
	if !a.RouteTableName.IsNull() && !a.RouteTableName.IsUnknown() {
		res.RouteTableName = a.RouteTableName.ValueString()
	}
	return &res, nil
}

func (a AzureInterfaceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"interface_id":     types.StringType,
		"route_table_name": types.StringType,
	}
}

func (a AzureInterfaceModel) ToObjectValue(azureIf *models.AzureInterface) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	var attrValues map[string]attr.Value
	if azureIf != nil {
		attrValues = map[string]attr.Value{
			"interface_id":     types.StringValue(azureIf.InterfaceID),
			"route_table_name": types.StringValue(azureIf.RouteTableName),
		}
	} else {
		attrValues = map[string]attr.Value{
			"interface_id":     types.StringNull(),
			"route_table_name": types.StringNull(),
		}
	}
	res, diags := types.ObjectValue(attrTypes, attrValues)
	if diags != nil && diags.HasError() {
		return nil, diags
	}
	return &res, nil
}

type OciInterfaceModel struct {
	VnicId                types.String `tfsdk:"vnic_id"`
	Network               types.Object `tfsdk:"network"`
	PublicIpId            types.String `tfsdk:"public_ip_id"`
	PrivateIpId           types.String `tfsdk:"private_ip_id"`
	SecondaryPrivateIp    types.String `tfsdk:"secondary_private_ip"`
	SecondaryPrivateIpId  types.String `tfsdk:"secondary_private_ip_id"`
	RouteTableId          types.String `tfsdk:"route_table_id"`
	NetworkSecurityGroups types.List   `tfsdk:"network_security_groups"`
}

func (a *OciInterfaceModel) ToAOModel() (*models.OciInterface, diag.Diagnostics) {
	res := models.OciInterface{}
	if !a.VnicId.IsNull() && !a.VnicId.IsUnknown() {
		res.VnicID = a.VnicId.ValueString()
	}
	if !a.PublicIpId.IsNull() && !a.PublicIpId.IsUnknown() {
		res.PublicIPID = a.PublicIpId.ValueString()
	}
	if !a.PrivateIpId.IsNull() && !a.PrivateIpId.IsUnknown() {
		res.PrivateIPID = a.PrivateIpId.ValueString()
	}
	if !a.SecondaryPrivateIp.IsNull() && !a.SecondaryPrivateIp.IsUnknown() {
		res.SecondaryPrivateIP = a.SecondaryPrivateIp.ValueString()
	}
	if !a.SecondaryPrivateIpId.IsNull() && !a.SecondaryPrivateIpId.IsUnknown() {
		res.SecondaryPrivateIPID = a.SecondaryPrivateIpId.ValueString()
	}
	if !a.RouteTableId.IsNull() && !a.RouteTableId.IsUnknown() {
		res.RouteTableID = a.RouteTableId.ValueString()
	}
	if !a.Network.IsNull() && !a.Network.IsUnknown() {
		networkModel := OciNetwork{}
		diag := a.Network.As(context.Background(), &networkModel, basetypes.ObjectAsOptions{
			UnhandledNullAsEmpty:    true,
			UnhandledUnknownAsEmpty: false,
		})
		if diag.HasError() {
			return nil, diag
		}
		res.Network = networkModel.ToAOModel()
	}
	if !a.NetworkSecurityGroups.IsNull() && !a.NetworkSecurityGroups.IsUnknown() {
		elements := make([]types.String, 0, len(a.NetworkSecurityGroups.Elements()))
		diag := a.NetworkSecurityGroups.ElementsAs(context.Background(), &elements, false)
		if diag.HasError() {
			return nil, diag
		}
		nsg := make([]string, 0)
		for _, element := range elements {
			nsg = append(nsg, element.ValueString())
		}
		res.NetworkSecurityGroups = nsg
	}

	return &res, nil
}

func (a OciInterfaceModel) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"vnic_id":                 types.StringType,
		"public_ip_id":            types.StringType,
		"private_ip_id":           types.StringType,
		"route_table_id":          types.StringType,
		"secondary_private_ip":    types.StringType,
		"secondary_private_ip_id": types.StringType,
		"network":                 types.ObjectType{}.WithAttributeTypes(OciNetwork{}.AttrType()),
		"network_security_groups": types.ListType{}.WithElementType(types.StringType),
	}
}
func (a OciInterfaceModel) ToObjectValue(ociIf *models.OciInterface) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	var attrValues map[string]attr.Value
	if ociIf != nil {
		netValue, diags := OciNetwork{}.ToObjectValue(ociIf.Network)
		if diags != nil && diags.HasError() {
			return nil, diags
		}

		sgList := make([]attr.Value, 0)
		for _, sg := range ociIf.NetworkSecurityGroups {
			sgList = append(sgList, types.StringValue(sg))
		}

		sgValues, diags := types.ListValue(types.StringType, sgList)
		if diags != nil && diags.HasError() {
			return nil, diags
		}

		attrValues = map[string]attr.Value{
			"vnic_id":                 types.StringValue(ociIf.VnicID),
			"public_ip_id":            types.StringValue(ociIf.PublicIPID),
			"private_ip_id":           types.StringValue(ociIf.PrivateIPID),
			"route_table_id":          types.StringValue(ociIf.RouteTableID),
			"secondary_private_ip":    types.StringValue(ociIf.SecondaryPrivateIP),
			"secondary_private_ip_id": types.StringValue(ociIf.SecondaryPrivateIPID),
			"network":                 netValue,
			"network_security_groups": sgValues,
		}
	} else {
		attrValues = map[string]attr.Value{
			"vnic_id":                 types.StringNull(),
			"public_ip_id":            types.StringNull(),
			"private_ip_id":           types.StringNull(),
			"route_table_id":          types.StringNull(),
			"secondary_private_ip":    types.StringNull(),
			"secondary_private_ip_id": types.StringNull(),
			"network":                 types.ObjectNull(OciNetwork{}.AttrType()),
			"network_security_groups": types.ListNull(types.StringType),
		}
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

type OciNetwork struct {
	SubnetName   types.String `tfsdk:"subnet_name"`
	VcnOcid      types.String `tfsdk:"vcn_ocid"`
	VcnName      types.String `tfsdk:"vcn_name"`
	SubnetOcid   types.String `tfsdk:"subnet_ocid"`
	SubnetAccess types.String `tfsdk:"subnet_access"`
}

func (a *OciNetwork) ToAOModel() *models.OciNetwork {
	res := models.OciNetwork{}
	if !a.SubnetName.IsNull() && !a.SubnetName.IsUnknown() {
		res.SubnetName = a.SubnetName.ValueString()
	}
	if !a.VcnOcid.IsNull() && !a.VcnOcid.IsUnknown() {
		res.VcnOcid = a.VcnOcid.ValueString()
	}
	if !a.VcnName.IsNull() && !a.VcnName.IsUnknown() {
		res.VcnName = a.VcnName.ValueString()
	}
	if !a.SubnetOcid.IsNull() && !a.SubnetOcid.IsUnknown() {
		res.SubnetOcid = a.SubnetOcid.ValueString()
	}
	if !a.SubnetAccess.IsNull() && !a.SubnetAccess.IsUnknown() {
		res.SubnetAccess = a.SubnetAccess.ValueString()
	}
	return &res
}

func (a OciNetwork) AttrType() map[string]attr.Type {
	return map[string]attr.Type{
		"subnet_name":   types.StringType,
		"vcn_ocid":      types.StringType,
		"vcn_name":      types.StringType,
		"subnet_ocid":   types.StringType,
		"subnet_access": types.StringType,
	}
}
func (a OciNetwork) ToObjectValue(ociNet *models.OciNetwork) (*basetypes.ObjectValue, diag.Diagnostics) {
	attrTypes := a.AttrType()
	var attrValues map[string]attr.Value
	if ociNet != nil {
		attrValues = map[string]attr.Value{
			"subnet_name":   types.StringValue(ociNet.SubnetName),
			"vcn_ocid":      types.StringValue(ociNet.VcnOcid),
			"vcn_name":      types.StringValue(ociNet.VcnName),
			"subnet_ocid":   types.StringValue(ociNet.SubnetOcid),
			"subnet_access": types.StringValue(ociNet.SubnetAccess),
		}
	} else {
		attrValues = map[string]attr.Value{
			"subnet_name":   types.StringNull(),
			"vcn_ocid":      types.StringNull(),
			"vcn_name":      types.StringNull(),
			"subnet_ocid":   types.StringNull(),
			"subnet_access": types.StringNull(),
		}
	}
	res, diag := types.ObjectValue(attrTypes, attrValues)
	if diag.HasError() {
		return nil, diag
	}
	return &res, nil
}

func (m *DeploymentResourceModel) ToAOCommonDeploymentModel(provider models.Providers, includeComputed bool) (*models.Deployment, diag.Diagnostics) {
	deployment := models.Deployment{}
	deployment.Name = m.Name.ValueStringPointer()
	deployment.Provider = provider
	deployment.CredentialName = m.CredentialName.ValueString()
	deployment.EnableHighAvailability = m.EnableHighAvailability.ValueBoolPointer()

	aeImageModel := ArcedgeImageModel{}
	diag := m.SourceImage.As(context.Background(), &aeImageModel, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: false,
	})
	if diag.HasError() {
		return nil, diag
	}
	deployment.SourceImage = aeImageModel.ToAOModel()

	if includeComputed {
		deployment.CredentialsID = m.CredentialId.ValueString()
		deployment.ArcOrchIP = m.ArcOrchIP.ValueString()
		deployment.Action = m.Action.ValueString()
		deployment.Status = m.Status.ValueString()
		deployment.StatusID = m.StatusID.ValueInt64()
		deployment.ArcedgeAStatus = m.ArcedgeAStatus.ValueString()
		deployment.ArcedgeAStatusID = m.ArcedgeAStatusId.ValueInt64()
		deployment.ArcedgeBStatus = m.ArcedgeBStatus.ValueString()
		deployment.ArcedgeBStatusID = m.ArcedgeBStatusId.ValueInt64()
		deployment.ArcedgeARole = models.ArcedgeRole(m.ArcedgeARole.ValueString())
		deployment.ArcedgeBRole = models.ArcedgeRole(m.ArcedgeBRole.ValueString())
		deployment.ArcedgeAIP = m.ArcedgeAIP.ValueString()
		deployment.ActiveIPGateway = m.ActiveIPGateway.ValueString()
		deployment.ArcedgeAPrivateIP = m.ArcedgeAPrivateIP.ValueString()
		deployment.ArcedgeBIP = m.ArcedgeBIP.ValueString()
		deployment.ArcedgeBPrivateIP = m.ArcedgeBPrivateIP.ValueString()
		deployment.PrivateCidr = m.PrivateCidr.ValueString()
		deployment.IngressSg = m.IngressSg.ValueString()
		deployment.HubNumber = m.HubNumber.ValueInt64()
		deployment.Coordinates = &models.Coordinates{
			Lat:  m.CoordinatesLat.ValueFloat64Pointer(),
			Long: m.CoordinatesLong.ValueFloat64Pointer(),
		}

		if !m.LatestAvailableImage.IsNull() {
			aeImageModel := ArcedgeImageModel{}
			m.LatestAvailableImage.As(context.Background(), aeImageModel, basetypes.ObjectAsOptions{
				UnhandledNullAsEmpty:    true,
				UnhandledUnknownAsEmpty: false,
			})
			deployment.LatestAvailableImage = aeImageModel.ToAOModel()
		}

		if !m.ActiveNetworkInterfaces.IsNull() {
			activeIfs := make([]*models.NetworkInterface, 0)
			networkIfs := make([]NetworkInterfacesModel, 0, len(m.ActiveNetworkInterfaces.Elements()))
			diag := m.ActiveNetworkInterfaces.ElementsAs(context.Background(), &networkIfs, false)
			if diag.HasError() {
				return nil, diag.Errors()
			}

			for _, networkIf := range networkIfs {
				networkObj, diag := networkIf.ToAOModel()
				if diag.HasError() {
					return nil, diag
				}
				activeIfs = append(activeIfs, networkObj)
			}

			deployment.ActiveNetworkInterfaces = activeIfs
		}

		if !m.BackupNetworkInterfaces.IsNull() {
			backupIfs := make([]*models.NetworkInterface, 0)
			networkIfs := make([]NetworkInterfacesModel, 0, len(m.BackupNetworkInterfaces.Elements()))
			diag := m.BackupNetworkInterfaces.ElementsAs(context.Background(), &networkIfs, false)
			if diag.HasError() {
				return nil, diag.Errors()
			}

			for _, networkIf := range networkIfs {
				networkObj, diag := networkIf.ToAOModel()
				if diag.HasError() {
					return nil, diag
				}
				backupIfs = append(backupIfs, networkObj)
			}

			deployment.BackupNetworkInterfaces = backupIfs
		}
	}
	return &deployment, nil
}

func (m *DeploymentResourceModel) ToTerraformCommonDeploymentModle(deployment *models.Deployment) diag.Diagnostics {
	m.Name = types.StringValue(*deployment.Name)
	m.CredentialId = types.StringValue(deployment.CredentialsID)
	m.CredentialName = types.StringValue(deployment.CredentialName)
	if deployment.EnableHighAvailability == nil {
		m.EnableHighAvailability = types.BoolValue(false)
	} else {
		m.EnableHighAvailability = types.BoolValue(*deployment.EnableHighAvailability)
	}
	m.ArcOrchIP = types.StringValue(deployment.ArcOrchIP)
	m.Action = types.StringValue(deployment.Action)
	m.Status = types.StringValue(deployment.Status)
	m.StatusID = types.Int64Value(deployment.StatusID)
	m.ArcedgeARole = types.StringValue(string(deployment.ArcedgeARole))
	m.ArcedgeBRole = types.StringValue(string(deployment.ArcedgeBRole))
	m.ArcedgeAStatus = types.StringValue(string(deployment.ArcedgeAStatus))
	m.ArcedgeBStatus = types.StringValue(string(deployment.ArcedgeBStatus))
	m.ArcedgeAStatusId = types.Int64Value(deployment.ArcedgeAStatusID)
	m.ArcedgeBStatusId = types.Int64Value(deployment.ArcedgeBStatusID)
	m.ArcedgeAIP = types.StringValue(deployment.ArcedgeAIP)
	m.ActiveIPGateway = types.StringValue(deployment.ActiveIPGateway)
	m.ArcedgeAPrivateIP = types.StringValue(deployment.ArcedgeAPrivateIP)
	m.ArcedgeBIP = types.StringValue(deployment.ArcedgeBIP)
	m.ArcedgeBPrivateIP = types.StringValue(deployment.ArcedgeBPrivateIP)
	m.PrivateCidr = types.StringValue(deployment.PrivateCidr)
	m.HubNumber = types.Int64Value(deployment.HubNumber)
	if deployment.Coordinates != nil {
		m.CoordinatesLat = types.Float64Value(*deployment.Coordinates.Lat)
		m.CoordinatesLong = types.Float64Value(*deployment.Coordinates.Long)
	}

	sourceImageObj, diags := types.ObjectValue(ArcedgeImageModel{}.AttrTypes(), ArcedgeImageModel{}.ToAttrValues(deployment.SourceImage))
	if diags != nil && diags.HasError() {
		log.Default().Println("a1!")
		return diags
	}
	m.SourceImage = sourceImageObj

	latestImageObj, diags := types.ObjectValue(ArcedgeImageModel{}.AttrTypes(), ArcedgeImageModel{}.ToAttrValues(deployment.LatestAvailableImage))
	if diags != nil && diags.HasError() {
		log.Default().Println("a2!")
		return diags
	}
	m.LatestAvailableImage = latestImageObj

	aIfList := make([]attr.Value, 0)
	for _, aIf := range deployment.ActiveNetworkInterfaces {
		aIfObj, diags := NetworkInterfacesModel{}.ToObjectValue(aIf)
		if diags != nil && diags.HasError() {
			return diags
		}

		aIfList = append(aIfList, aIfObj)
	}

	activeIfList, diags := types.ListValue(types.ObjectType{}.WithAttributeTypes(NetworkInterfacesModel{}.AttrType()), aIfList)
	if diags != nil && diags != nil && diags.HasError() {
		return diags
	}
	m.ActiveNetworkInterfaces = activeIfList

	bIfList := make([]attr.Value, 0)
	for _, bIf := range deployment.BackupNetworkInterfaces {
		bIfObj, diags := NetworkInterfacesModel{}.ToObjectValue(bIf)
		if diags != nil && diags.HasError() {
			return diags
		}

		bIfList = append(bIfList, bIfObj)
	}
	backupIfList, diags := types.ListValue(types.ObjectType{}.WithAttributeTypes(NetworkInterfacesModel{}.AttrType()), bIfList)
	if diags != nil && diags.HasError() {
		return diags
	}
	m.BackupNetworkInterfaces = backupIfList
	return nil
}

package deploymentresource

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var ObjectAsOptions = basetypes.ObjectAsOptions{
	UnhandledNullAsEmpty:    true,
	UnhandledUnknownAsEmpty: false,
}

func DeploymentResourceModelSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"credential_id": schema.StringAttribute{
			MarkdownDescription: "a unique identifier of the credential for the deployed Arcedge",
			Computed:            true,
		},
		"credential_name": schema.StringAttribute{
			MarkdownDescription: "The name of the credential where the ArcEdge will be deployed.",
			Required:            true,
		},
		"enable_high_availability": schema.BoolAttribute{
			MarkdownDescription: "Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.",
			Required:            true,
		},
		"arc_orch_ip": schema.StringAttribute{
			MarkdownDescription: "Public IP of the ArcEdge that has been deployed.",
			Computed:            true,
		},
		"action": schema.StringAttribute{
			MarkdownDescription: "Action initiated by the ArcOrch (Will be either CREATE/UPDATE/DELETE).",
			Computed:            true,
		},
		"status": schema.StringAttribute{
			MarkdownDescription: "Current status of the ArcEdge deployment (verbose).",
			Computed:            true,
		},
		"status_id": schema.Int64Attribute{
			MarkdownDescription: "Current status of the ArcEdge deployment (numerical).",
			Computed:            true,
		},
		"arcedge_a_role": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_b_role": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_a_status": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_a_status_id": schema.Int64Attribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_b_status": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_b_status_id": schema.Int64Attribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_a_ip": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"active_ip_gateway": schema.StringAttribute{
			MarkdownDescription: " Default gateway for the deployed ArcEdge.",
			Computed:            true,
		},
		"arcedge_a_private_ip": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_b_ip": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"arcedge_b_private_ip": schema.StringAttribute{
			MarkdownDescription: "",
			Computed:            true,
		},
		"private_cidr": schema.StringAttribute{
			MarkdownDescription: "CIDR block of the private subnet.",
			Computed:            true,
		},
		"ingress_sg": schema.StringAttribute{
			MarkdownDescription: "Security group created for the ArcEdge deployment.",
			Computed:            true,
		},
		"hub_number": schema.Int64Attribute{
			MarkdownDescription: " Hub number of the overlay network.",
			Computed:            true,
		},
		"coordinates_lat": schema.Float64Attribute{
			MarkdownDescription: "Latitude where ArcEdge is deployed",
			Computed:            true,
		},
		"coordinates_long": schema.Float64Attribute{
			MarkdownDescription: "Longitude where ArcEdge is deployed.",
			Computed:            true,
		},
		"source_image":              ArcedgeImageResourceSchema(),
		"latest_available_image":    ArcedgeImageDataSchema(),
		"active_network_interfaces": NetworkInterfaceDataSchema(),
		"backup_network_interfaces": NetworkInterfaceDataSchema(),
	}
}

func DeploymentDataModelSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"credential_id": schema.StringAttribute{
			Computed: true,
		},
		"credential_name": schema.StringAttribute{
			Computed: true,
		},
		"enable_high_availability": schema.BoolAttribute{
			Computed: true,
		},
		"arc_orch_ip": schema.StringAttribute{
			Computed: true,
		},
		"action": schema.StringAttribute{
			Computed: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
		"status_id": schema.Int64Attribute{
			Computed: true,
		},
		"arcedge_a_role": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_b_role": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_a_status": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_a_status_id": schema.Int64Attribute{
			Computed: true,
		},
		"arcedge_b_status": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_b_status_id": schema.Int64Attribute{
			Computed: true,
		},
		"arcedge_a_ip": schema.StringAttribute{
			Computed: true,
		},
		"active_ip_gateway": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_a_private_ip": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_b_ip": schema.StringAttribute{
			Computed: true,
		},
		"arcedge_b_private_ip": schema.StringAttribute{
			Computed: true,
		},
		"private_cidr": schema.StringAttribute{
			Computed: true,
		},
		"ingress_sg": schema.StringAttribute{
			Computed: true,
		},
		"hub_number": schema.Int64Attribute{
			Computed: true,
		},
		"coordinates_lat": schema.Float64Attribute{
			Computed: true,
		},
		"coordinates_long": schema.Float64Attribute{
			Computed: true,
		},
		"source_image":              ArcedgeDataResourceSchema(),
		"latest_available_image":    ArcedgeImageDataSchema(),
		"active_network_interfaces": NetworkInterfaceDataSchema(),
		"backup_network_interfaces": NetworkInterfaceDataSchema(),
	}
}

func ArcedgeImageResourceSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"image_id": schema.StringAttribute{
				MarkdownDescription: "",
				Required:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"provider": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
		Required: true,
	}
}

func ArcedgeDataResourceSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"image_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"provider": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
		Computed: true,
	}
}

func ArcedgeImageDataSchema() schema.SingleNestedAttribute {
	return schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"image_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"provider": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
		},
		Computed: true,
	}
}

func NetworkInterfaceDataSchema() schema.ListNestedAttribute {
	return schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Computed: true,
				},
				"private_ipv4_address": schema.StringAttribute{
					Computed: true,
				},
				"public_ipv4_address": schema.StringAttribute{
					Computed: true,
				},
				"global_ipv6_address": schema.StringAttribute{
					Computed: true,
				},
				"linklocal_ipv6_address": schema.StringAttribute{
					Computed: true,
				},
				"mac_address": schema.StringAttribute{
					Computed: true,
				},
				"adapter_type": schema.StringAttribute{
					Computed: true,
				},
				"private_ipv4_pfxlen": schema.Int64Attribute{
					Computed: true,
				},
				"public_ipv4_pfxlen": schema.Int64Attribute{
					Computed: true,
				},
				"linklocal_ipv6_pfxlen": schema.Int64Attribute{
					Computed: true,
				},
				"global_ipv6_pfxlen": schema.Int64Attribute{
					Computed: true,
				},
				"aws_interface": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							Computed: true,
						},
						"route_table_id": schema.StringAttribute{
							Computed: true,
						},
						"subnet_id": schema.StringAttribute{
							Computed: true,
						},
						"security_groups": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
					Computed: true,
				},
				"azure_interface": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"interface_id": schema.StringAttribute{
							Computed: true,
						},
						"route_table_name": schema.StringAttribute{
							Computed: true,
						},
					},
					Computed: true,
				},
				"oci_interface": schema.SingleNestedAttribute{
					Attributes: map[string]schema.Attribute{
						"vnic_id": schema.StringAttribute{
							Computed: true,
						},
						"public_ip_id": schema.StringAttribute{
							Computed: true,
						},
						"private_ip_id": schema.StringAttribute{
							Computed: true,
						},
						"secondary_private_ip": schema.StringAttribute{
							Computed: true,
						},
						"secondary_private_ip_id": schema.StringAttribute{
							Computed: true,
						},
						"route_table_id": schema.StringAttribute{
							Computed: true,
						},
						"network": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"subnet_name": schema.StringAttribute{
									Computed: true,
								},
								"vcn_ocid": schema.StringAttribute{
									Computed: true,
								},
								"vcn_name": schema.StringAttribute{
									Computed: true,
								},
								"subnet_ocid": schema.StringAttribute{
									Computed: true,
								},
								"subnet_access": schema.StringAttribute{
									Computed: true,
								},
							},
							Computed: true,
						},
						"network_security_groups": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
						},
					},
					Computed: true,
				},
			},
		},
		Computed: true,
	}
}

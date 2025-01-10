package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DeploymentResourceSchema() map[string]*schema.Schema {
	schema := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"credential_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"credential_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"enable_high_availability": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"arc_orch_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_a_role": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_role": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_b_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_a_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"active_ip_gateway": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_private_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_private_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_cidr": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ingress_sg": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"hub_number": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"coordinates_lat": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"coordinates_long": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"source_image": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: schemas.ArcedgeImageResourceSchema(),
			},
		},
		"latest_available_image": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.ArcedgeImageResourceSchema(),
			},
		},
		"active_network_interfaces": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.NetworkInterfaceResourceSchema(),
			},
		},
		"backup_network_interfaces": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.NetworkInterfaceResourceSchema(),
			},
		},
	}
	return schema
}

func DeploymentDataSchema() map[string]*schema.Schema {
	schema := map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"credential_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"credential_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"enable_high_availability": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"arc_orch_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"action": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_a_role": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_role": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_b_status": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_status_id": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"arcedge_a_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"active_ip_gateway": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_a_private_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"arcedge_b_private_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_cidr": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ingress_sg": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"hub_number": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"coordinates_lat": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"coordinates_long": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"source_image": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.ArcedgeImageResourceSchema(),
			},
		},
		"latest_available_image": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.ArcedgeImageResourceSchema(),
			},
		},
		"active_network_interfaces": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.NetworkInterfaceResourceSchema(),
			},
		},
		"backup_network_interfaces": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: schemas.NetworkInterfaceResourceSchema(),
			},
		},
	}
	return schema
}

func ToCommonDeploymentObj(d *schema.ResourceData, provider models.Providers, deployment *models.Deployment, includeComputed bool) error {
	deployment.Name = utils.StrPtr(d.Get("name").(string))
	deployment.Provider = provider
	deployment.CredentialName = *utils.StrPtr(d.Get("credential_name").(string))
	deployment.EnableHighAvailability = d.Get("enable_high_availability").(bool)
	// deployment.EnablePrivateSubnet = d.Get("enable_private_subnet").(bool)
	if _, exist := d.GetOkExists("source_image"); exist {
		sourceImageFieldFormat := fmt.Sprintf("%s.%d.%%s", "source_image", 0)
		deployment.SourceImage = &models.ArcedgeImage{
			ImageID: utils.StrPtr(d.Get(fmt.Sprintf(sourceImageFieldFormat, "image_id")).(string)),
		}
	}

	if includeComputed {
		deployment.CredentialsID = *utils.StrPtr(d.Get("credential_id").(string))
		deployment.ArcOrchIP = *utils.StrPtr(d.Get("arc_orch_ip").(string))
		deployment.Action = *utils.StrPtr(d.Get("action").(string))
		deployment.Status = *utils.StrPtr(d.Get("status").(string))
		deployment.StatusID = *utils.Int64Ptr(d.Get("status_id").(int))
		deployment.ArcedgeAStatus = *utils.StrPtr(d.Get("arcedge_a_status").(string))
		deployment.ArcedgeAStatusID = *utils.Int64Ptr(d.Get("arcedge_a_status_id").(int))
		deployment.ArcedgeBStatus = *utils.StrPtr(d.Get("arcedge_b_status").(string))
		deployment.ArcedgeBStatusID = *utils.Int64Ptr(d.Get("arcedge_b_status_id").(int))
		deployment.StatusID = *utils.Int64Ptr(d.Get("status_id").(int))
		deployment.Status = *utils.StrPtr(d.Get("status").(string))
		deployment.StatusID = *utils.Int64Ptr(d.Get("status_id").(int))
		deployment.ArcedgeARole = models.ArcedgeRole(d.Get("arcedge_a_role").(string))
		deployment.ArcedgeBRole = models.ArcedgeRole(d.Get("arcedge_b_role").(string))
		deployment.ArcedgeAIP = *utils.StrPtr(d.Get("arcedge_a_ip").(string))
		deployment.ActiveIPGateway = *utils.StrPtr(d.Get("active_ip_gateway").(string))
		deployment.ArcedgeAPrivateIP = *utils.StrPtr(d.Get("arcedge_a_private_ip").(string))
		deployment.ArcedgeBIP = *utils.StrPtr(d.Get("arcedge_b_ip").(string))
		deployment.ArcedgeBPrivateIP = *utils.StrPtr(d.Get("arcedge_b_private_ip").(string))
		deployment.PrivateCidr = *utils.StrPtr(d.Get("private_cidr").(string))
		deployment.IngressSg = *utils.StrPtr(d.Get("ingress_sg").(string))
		deployment.HubNumber = *utils.Int64Ptr(d.Get("hub_number").(int))
		deployment.Coordinates = &models.Coordinates{
			Lat:  utils.Float64Ptr(d.Get("coordinates_lat").(float64)),
			Long: utils.Float64Ptr(d.Get("coordinates_long").(float64)),
		}

		if _, exist := d.GetOkExists("latest_available_image"); exist {
			latestImageFieldFormat := fmt.Sprintf("%s.%d.%%s", "latest_available_image", 0)
			deployment.SourceImage = &models.ArcedgeImage{
				ImageID: utils.StrPtr(d.Get(fmt.Sprintf(latestImageFieldFormat, "image_id")).(string)),
			}
		}

		if netIf, exist := d.GetOkExists("active_network_interfaces"); exist {
			net := netIf.([]interface{})
			deployment.ActiveNetworkInterfaces = make([]*models.NetworkInterface, 0)
			for i, _ := range net {
				netFieldFormat := fmt.Sprintf("%s.%d.%%s", "active_network_interfaces", i)
				temp := models.NetworkInterface{}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ipv4_address")); exist {
					temp.PrivateIPV4Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "public_ipv4_address")); exist {
					temp.PublicIPV4Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "linklocal_ipv6_address")); exist {
					temp.LinklocalIPV6Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "global_ipv6_address")); exist {
					temp.GlobalIPV6Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "mac_address")); exist {
					temp.MacAddress = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "adapter_type")); exist {
					temp.AdapterType = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ipv4_pfxlen")); exist {
					temp.PrivateIPV4Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "public_ipv4_pfxlen")); exist {
					temp.PublicIPV4Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "linklocal_ipv6_pfxlen")); exist {
					temp.LinklocalIPV6Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "global_ipv6_pfxlen")); exist {
					temp.GlobalIPV6Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "azure_interface")); exist {
					azureIfs := f.([]interface{})
					if len(azureIfs) > 0 {
						azureIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "azure_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(azureIfFormat, "interface_id")); exist {
							temp.AzureInterface.InterfaceID = g.(string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "oci_interface")); exist {
					ociIfs := f.([]interface{})
					if len(ociIfs) > 0 {
						ociIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "oci_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "vnic_id")); exist {
							temp.OciInterface.VnicID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "public_ip_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "private_ip_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "route_table_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "network_security_groups")); exist {
							temp.OciInterface.NetworkSecurityGroups = g.([]string)
						}
						if _, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "network")); exist {
							ociIfNetFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(ociIfFormat, "network"), 0)
							ociIfNets := f.([]interface{})
							if len(ociIfNets) > 0 {
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_name")); exist {
									temp.OciInterface.Network.SubnetName = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "vcn_ocid")); exist {
									temp.OciInterface.Network.VcnOcid = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "vcn_name")); exist {
									temp.OciInterface.Network.VcnName = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_ocid")); exist {
									temp.OciInterface.Network.SubnetOcid = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_access")); exist {
									temp.OciInterface.Network.SubnetAccess = g.(string)
								}
							}
						}
					}
				}
				deployment.ActiveNetworkInterfaces = append(deployment.ActiveNetworkInterfaces, &temp)
			}
		}

		if netIf, exist := d.GetOkExists("backup_network_interfaces"); exist {
			net := netIf.([]interface{})
			deployment.BackupNetworkInterfaces = make([]*models.NetworkInterface, 0)
			for i, _ := range net {
				netFieldFormat := fmt.Sprintf("%s.%d.%%s", "backup_network_interfaces", i)
				temp := models.NetworkInterface{}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ipv4_address")); exist {
					temp.PrivateIPV4Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "public_ipv4_address")); exist {
					temp.PublicIPV4Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "linklocal_ipv6_address")); exist {
					temp.LinklocalIPV6Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "global_ipv6_address")); exist {
					temp.GlobalIPV6Address = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "mac_address")); exist {
					temp.MacAddress = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "adapter_type")); exist {
					temp.AdapterType = f.(string)
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ipv4_pfxlen")); exist {
					temp.PrivateIPV4Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "public_ipv4_pfxlen")); exist {
					temp.PublicIPV4Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "linklocal_ipv6_pfxlen")); exist {
					temp.LinklocalIPV6Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "global_ipv6_pfxlen")); exist {
					temp.GlobalIPV6Pfxlen = int64(f.(int))
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "aws_interface")); exist {
					awsIfs := f.([]interface{})
					if len(awsIfs) > 0 {
						awsIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "aws_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "interface_id")); exist {
							temp.AwsInterface.InterfaceID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(awsIfFormat, "security_groups")); exist {
							temp.AwsInterface.SecurityGroups = g.([]string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "azure_interface")); exist {
					azureIfs := f.([]interface{})
					if len(azureIfs) > 0 {
						azureIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "azure_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(azureIfFormat, "interface_id")); exist {
							temp.AzureInterface.InterfaceID = g.(string)
						}
					}
				}
				if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "oci_interface")); exist {
					ociIfs := f.([]interface{})
					if len(ociIfs) > 0 {
						ociIfFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(netFieldFormat, "oci_interface"), 0)
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "vnic_id")); exist {
							temp.OciInterface.VnicID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "public_ip_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "private_ip_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "route_table_id")); exist {
							temp.OciInterface.PublicIPID = g.(string)
						}
						if g, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "network_security_groups")); exist {
							temp.OciInterface.NetworkSecurityGroups = g.([]string)
						}
						if _, exist := d.GetOkExists(fmt.Sprintf(ociIfFormat, "network")); exist {
							ociIfNetFormat := fmt.Sprintf("%s.%d.%%s", fmt.Sprintf(ociIfFormat, "network"), 0)
							ociIfNets := f.([]interface{})
							if len(ociIfNets) > 0 {
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_name")); exist {
									temp.OciInterface.Network.SubnetName = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "vcn_ocid")); exist {
									temp.OciInterface.Network.VcnOcid = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "vcn_name")); exist {
									temp.OciInterface.Network.VcnName = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_ocid")); exist {
									temp.OciInterface.Network.SubnetOcid = g.(string)
								}
								if g, exist := d.GetOkExists(fmt.Sprintf(ociIfNetFormat, "subnet_access")); exist {
									temp.OciInterface.Network.SubnetAccess = g.(string)
								}
							}
						}
					}
				}
				deployment.BackupNetworkInterfaces = append(deployment.BackupNetworkInterfaces, &temp)
			}
		}

	}
	return nil
}

func UpdateCommonDeploymentResource(deployment *models.Deployment, d *schema.ResourceData) error {
	err := d.Set("arc_orch_ip", deployment.ArcOrchIP)
	if err != nil {
		return err
	}
	err = d.Set("action", deployment.Action)
	if err != nil {
		return err
	}
	err = d.Set("status", deployment.Status)
	if err != nil {
		return err
	}
	err = d.Set("status_id", int(deployment.StatusID))
	if err != nil {
		return err
	}
	err = d.Set("arcedge_a_status", deployment.ArcedgeAStatus)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_a_status_id", int(deployment.ArcedgeAStatusID))
	if err != nil {
		return err
	}
	err = d.Set("arcedge_b_status", deployment.ArcedgeBStatus)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_b_status_id", int(deployment.ArcedgeBStatusID))
	if err != nil {
		return err
	}
	err = d.Set("arcedge_a_role", deployment.ArcedgeARole)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_b_role", deployment.ArcedgeBRole)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_a_ip", deployment.ArcedgeAIP)
	if err != nil {
		return err
	}
	err = d.Set("active_ip_gateway", deployment.ActiveIPGateway)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_a_private_ip", deployment.ArcedgeAPrivateIP)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_b_ip", deployment.ArcedgeBIP)
	if err != nil {
		return err
	}
	err = d.Set("arcedge_b_private_ip", deployment.ArcedgeBPrivateIP)
	if err != nil {
		return err
	}
	err = d.Set("private_cidr", deployment.PrivateCidr)
	if err != nil {
		return err
	}
	err = d.Set("ingress_sg", deployment.IngressSg)
	if err != nil {
		return err
	}
	err = d.Set("hub_number", int(deployment.HubNumber))
	if err != nil {
		return err
	}
	err = d.Set("enable_high_availability", deployment.EnableHighAvailability)
	if err != nil {
		return err
	}
	// err = d.Set("enable_private_subnet", deployment.EnablePrivateSubnet)
	// if err != nil {
	// 	return err
	// }
	err = d.Set("credential_id", deployment.CredentialsID)
	if err != nil {
		return err
	}
	err = d.Set("credential_name", deployment.CredentialName)
	if err != nil {
		return err
	}

	latestImage := make(map[string]interface{})
	latestImage["image_id"] = deployment.LatestAvailableImage.ImageID
	latestImage["version"] = deployment.LatestAvailableImage.Version
	latestImage["provider"] = deployment.LatestAvailableImage.Provider
	latestImage["name"] = deployment.LatestAvailableImage.Name
	err = d.Set("latest_available_image", []map[string]interface{}{latestImage})
	if err != nil {
		return err
	}
	sourceImage := make(map[string]interface{})
	sourceImage["image_id"] = deployment.SourceImage.ImageID
	sourceImage["version"] = deployment.SourceImage.Version
	sourceImage["provider"] = deployment.SourceImage.Provider
	sourceImage["name"] = deployment.SourceImage.Name
	err = d.Set("source_image", []map[string]interface{}{sourceImage})
	if err != nil {
		return err
	}

	activeNetIfs := make([]interface{}, 0)
	for _, netif := range deployment.ActiveNetworkInterfaces {
		activeNetIfs = append(activeNetIfs, NetworkInterfaceResource(netif))
	}
	err = d.Set("active_network_interfaces", activeNetIfs)
	if err != nil {
		return err
	}

	backupNetIfs := make([]interface{}, 0)
	for _, netif := range deployment.BackupNetworkInterfaces {
		backupNetIfs = append(backupNetIfs, NetworkInterfaceResource(netif))
	}
	err = d.Set("backup_network_interfaces", backupNetIfs)
	if err != nil {
		return err
	}
	return nil
}

func NetworkInterfaceResource(networkinterface *models.NetworkInterface) map[string]interface{} {
	res := make(map[string]interface{})
	res["name"] = networkinterface.Name
	res["private_ipv4_address"] = networkinterface.PrivateIPV4Address
	res["public_ipv4_address"] = networkinterface.PublicIPV4Address
	res["linklocal_ipv6_address"] = networkinterface.LinklocalIPV6Address
	res["global_ipv6_address"] = networkinterface.GlobalIPV6Address
	res["mac_address"] = networkinterface.MacAddress
	res["adapter_type"] = networkinterface.AdapterType
	res["private_ipv4_pfxlen"] = networkinterface.PrivateIPV4Pfxlen
	res["public_ipv4_pfxlen"] = networkinterface.PublicIPV4Pfxlen
	res["linklocal_ipv6_pfxlen"] = networkinterface.LinklocalIPV6Pfxlen
	res["global_ipv6_pfxlen"] = networkinterface.GlobalIPV6Pfxlen

	if networkinterface.AwsInterface != nil {
		awsIf := make(map[string]interface{})
		awsIf["interface_id"] = networkinterface.AwsInterface.InterfaceID
		awsIf["security_groups"] = networkinterface.AwsInterface.SecurityGroups
	}

	if networkinterface.AzureInterface != nil {
		azureIf := make(map[string]interface{})
		azureIf["interface_id"] = networkinterface.AzureInterface.InterfaceID
	}

	if networkinterface.OciInterface != nil {
		ociIf := make(map[string]interface{})
		ociIf["vnic_id"] = networkinterface.OciInterface.VnicID
		ociIf["public_ip_id"] = networkinterface.OciInterface.PublicIPID
		ociIf["network_security_groups"] = networkinterface.OciInterface.NetworkSecurityGroups
		ociIf["route_table_id"] = networkinterface.OciInterface.RouteTableID
		ociIf["private_ip_id"] = networkinterface.OciInterface.PrivateIPID
		if networkinterface.OciInterface.Network != nil {
			ociNet := make(map[string]interface{})
			ociNet["subnet_name"] = networkinterface.OciInterface.Network.SubnetName
			ociNet["vcn_ocid"] = networkinterface.OciInterface.Network.VcnOcid
			ociNet["vcn_name"] = networkinterface.OciInterface.Network.VcnName
			ociNet["subnet_ocid"] = networkinterface.OciInterface.Network.SubnetOcid
			ociNet["subnet_access"] = networkinterface.OciInterface.Network.SubnetAccess
		}
	}
	return res
}

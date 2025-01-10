package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OciNetworkResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vcn": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subnet": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func NetworkInterfaceResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_ipv4_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"public_ipv4_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"global_ipv6_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"linklocal_ipv6_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"mac_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"adapter_type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_ipv4_pfxlen": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"public_ipv4_pfxlen": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"linklocal_ipv6_pfxlen": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"global_ipv6_pfxlen": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"aws_interface": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: AwsInterfaceSchema(),
			},
		},
		"azure_interface": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: AzureInterfaceSchema(),
			},
		},
		"oci_interface": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: OciNetworkSchema(),
			},
		},
	}
}

func AzureNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"subnetA": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subnetB": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"private_subnet_route_table": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func AwsInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"security_groups": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func AzureInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"interface_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func OciNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"subnet_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"vcn_ocid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vcn_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"subnet_ocid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subnet_access": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func OciInterfaceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vnic_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"network": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: OciNetworkSchema(),
			},
		},
		"public_ip_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_ip_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"route_table_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"network_security_groups": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

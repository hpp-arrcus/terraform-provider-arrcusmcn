package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OciDeploymentResourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentResourceSchema()
	for k, v := range OciDeploymentSchema() {
		deploymentSchema[k] = v
	}
	return deploymentSchema
}

func OciDeploymentDatasourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentDataSchema()
	for k, v := range OciDeploymentSchema() {
		v.Computed = true
		v.Required = false
		if v.Type == schema.TypeList {
			v.MaxItems = 0
			v.MinItems = 0
		}
		deploymentSchema[k] = v
	}
	return deploymentSchema
}

func OciDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"instance_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
		"availability_domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"backup_availability_domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"compartment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"image_compartment": {
			Type:     schema.TypeString,
			Required: true,
		},
		"compute_shape": {
			Type:     schema.TypeString,
			Required: true,
		},
		"compute_cpus": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"compute_memory_in_gbs": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"enable_firewall": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"byoip": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"byoip_backup": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"networks": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: OciNetworkSchema(),
			},
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

func ToOciDeploymentObj(d *schema.ResourceData, deployment *models.Deployment) error {
	publicKey, err := utils.ReadTextFile(d.Get("instance_key").(string))
	if err != nil {
		return err
	}
	keyName, err := utils.GetPublicKeyName(*publicKey)
	if err != nil {
		return err
	}
	deployment.OciDeployment = &models.OciDeployment{
		SSHAuthorizedKeys: &models.InstanceKey{
			Content: *publicKey,
			Name:    keyName,
		},
		Region:             *utils.StrPtr(d.Get("region").(string)),
		AvailabilityDomain: *utils.StrPtr(d.Get("availability_domain").(string)),
		Compartment:        *utils.StrPtr(d.Get("compartment").(string)),
		ImageCompartment:   *utils.StrPtr(d.Get("image_compartment").(string)),
		ComputeShape:       *utils.StrPtr(d.Get("compute_shape").(string)),
	}
	if f, exist := d.GetOkExists("backup_availability_domain"); exist {
		deployment.OciDeployment.BackupAvailabilityDomain = f.(string)
	}
	if f, exist := d.GetOkExists("compute_cpus"); exist {
		deployment.OciDeployment.ComputeCpus = f.(string)
	}
	if f, exist := d.GetOkExists("compute_memory_in_gbs"); exist {
		deployment.OciDeployment.ComputeMemoryInGbs = f.(string)
	}
	if f, exist := d.GetOkExists("enable_firewall"); exist {
		deployment.OciDeployment.EnableFirewall = utils.BoolPtr(f.(bool))
	} else {
		deployment.OciDeployment.EnableFirewall = utils.BoolPtr(true)
	}
	if f, exist := d.GetOkExists("byoip"); exist {
		deployment.OciDeployment.Byoip = f.(string)
	}
	if f, exist := d.GetOkExists("byoip_backup"); exist {
		deployment.OciDeployment.ByoipBackup = f.(string)
	}

	if netIf, exist := d.GetOkExists("networks"); exist {
		net := netIf.([]interface{})
		deployment.OciDeployment.Networks = make([]*models.OciNetwork, 0)
		for i, _ := range net {
			netFieldFormat := fmt.Sprintf("%s.%d.%%s", "networks", i)
			temp := models.OciNetwork{}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnet_name")); exist {
				temp.SubnetName = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "vcn_ocid")); exist {
				temp.VcnOcid = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "vcn_name")); exist {
				temp.VcnName = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnet_ocid")); exist {
				temp.SubnetOcid = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnet_access")); exist {
				temp.SubnetAccess = f.(string)
			}
			deployment.OciDeployment.Networks = append(deployment.OciDeployment.Networks, &temp)
		}
	}
	return nil
}

func UpdateOciDeploymentResource(deployment *models.OciDeployment, d *schema.ResourceData) error {
	// ociDeployment["instance_key"] = deployment.InstanceKey
	err := d.Set("region", deployment.Region)
	if err != nil {
		return err
	}

	err = d.Set("availability_domain", deployment.AvailabilityDomain)
	if err != nil {
		return err
	}

	err = d.Set("compartment", deployment.Compartment)
	if err != nil {
		return err
	}

	err = d.Set("image_compartment", deployment.ImageCompartment)
	if err != nil {
		return err
	}

	err = d.Set("compute_shape", deployment.ComputeShape)
	if err != nil {
		return err
	}

	err = d.Set("compute_cpus", deployment.ComputeCpus)
	if err != nil {
		return err
	}

	err = d.Set("compute_memory_in_gbs", deployment.ComputeMemoryInGbs)
	if err != nil {
		return err
	}

	err = d.Set("enable_firewall", *deployment.EnableFirewall)
	if err != nil {
		return err
	}

	err = d.Set("byoip", deployment.Byoip)
	if err != nil {
		return err
	}

	err = d.Set("byoip_backup", deployment.ByoipBackup)
	if err != nil {
		return err
	}

	netList := make([]map[string]interface{}, 0)
	for _, net := range deployment.Networks {
		tmp := make(map[string]interface{})
		tmp["subnet_name"] = net.SubnetName
		tmp["vcn_ocid"] = net.VcnOcid
		tmp["vcn_name"] = net.VcnOcid
		tmp["subnet_ocid"] = net.SubnetOcid
		tmp["subnet_access"] = net.SubnetAccess
		netList = append(netList, tmp)
	}
	err = d.Set("networks", netList)
	if err != nil {
		return err
	}

	return nil
}

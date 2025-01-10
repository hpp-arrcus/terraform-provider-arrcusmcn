package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OnpremDeploymentResourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentResourceSchema()
	for k, v := range OnpremDeploymentSchema() {
		deploymentSchema[k] = v
	}
	return deploymentSchema
}

func OnpremDeploymentDatasourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentDataSchema()
	for k, v := range OnpremDeploymentSchema() {
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

func OnpremDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"vcpus": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"vm_memory": {
			Type:     schema.TypeInt,
			Required: true,
		},
		"private_ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"public_ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ssh_psw": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"public_ip_backup": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"private_ip_backup": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"assign_public_ip": {
			Type:     schema.TypeBool,
			Optional: true,
		},
		"networks": {
			Type:     schema.TypeList,
			Required: true,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: KvmNetworkSchema(),
			},
		},
	}
}

func KvmNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"private_ip_default_gw": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"private_ip_cidr_mask": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"network": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func ToOnpremDeploymentObj(d *schema.ResourceData, deployment *models.Deployment) error {
	deployment.OnpremDeployment = &models.OnpremDeployment{
		VMCount:   1,
		Vcpus:     *utils.Int64Ptr(d.Get("vcpus").(int)),
		VMMemory:  *utils.Int64Ptr(d.Get("vm_memory").(int)),
		PrivateIP: *utils.StrPtr(d.Get("private_ip").(string)),
		PublicIP:  *utils.StrPtr(d.Get("public_ip").(string)),
		Prefix:    *utils.StrPtr(d.Get("name").(string)),
		SSHPsw:    *utils.StrPtr(d.Get("ssh_psw").(string)),
	}

	if f, exist := d.GetOkExists("assign_public_ip"); exist {
		deployment.OnpremDeployment.AssignPublicIP = f.(bool)
	} else {
		deployment.OnpremDeployment.AssignPublicIP = true
	}

	if netIf, exist := d.GetOkExists("networks"); exist {
		net := netIf.([]interface{})
		deployment.OnpremDeployment.Networks = make([]*models.KvmNetwork, 0)
		for i, _ := range net {
			netFieldFormat := fmt.Sprintf("%s.%d.%%s", "networks", i)
			temp := models.KvmNetwork{}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "name")); exist {
				temp.Name = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ip_default_gw")); exist {
				temp.PrivateIPDefaultGw = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_ip_cidr_mask")); exist {
				temp.PrivateIPCidrMask = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "network")); exist {
				temp.Network = f.(string)
			}
			deployment.OnpremDeployment.Networks = append(deployment.OnpremDeployment.Networks, &temp)
		}
	}
	return nil
}

func UpdateOnpremDeploymentResource(deployment *models.OnpremDeployment, d *schema.ResourceData) error {
	err := d.Set("vcpus", deployment.Vcpus)
	if err != nil {
		return err
	}

	err = d.Set("vm_memory", deployment.VMMemory)
	if err != nil {
		return err
	}

	err = d.Set("private_ip", deployment.PrivateIP)
	if err != nil {
		return err
	}

	err = d.Set("private_ip_backup", deployment.PrivateIPBackup)
	if err != nil {
		return err
	}

	err = d.Set("public_ip", deployment.PublicIP)
	if err != nil {
		return err
	}

	err = d.Set("public_ip_backup", deployment.PublicIPBackup)
	if err != nil {
		return err
	}

	err = d.Set("assign_public_ip", deployment.AssignPublicIP)
	if err != nil {
		return err
	}

	err = d.Set("ssh_psw", deployment.SSHPsw)
	if err != nil {
		return err
	}

	netList := make([]map[string]interface{}, 0)
	for _, net := range deployment.Networks {
		tmp := make(map[string]interface{})
		tmp["name"] = net.Name
		tmp["private_ip_default_gw"] = net.PrivateIPDefaultGw
		tmp["private_ip_cidr_mask"] = net.PrivateIPCidrMask
		tmp["network"] = net.Network
		netList = append(netList, tmp)
	}
	err = d.Set("networks", netList)
	if err != nil {
		return err
	}

	return nil
}

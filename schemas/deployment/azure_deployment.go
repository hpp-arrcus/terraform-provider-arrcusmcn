package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AzureDeploymentResourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentResourceSchema()
	for k, v := range AzureDeploymentSchema() {
		deploymentSchema[k] = v
	}

	return deploymentSchema
}

func AzureDeploymentDatasourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentDataSchema()
	for k, v := range AzureDeploymentSchema() {
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

func AzureDeploymentSchema() map[string]*schema.Schema {
	schema := map[string]*schema.Schema{
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
		"instance_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"instance_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resource_group": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vnet": {
			Type:     schema.TypeString,
			Required: true,
		},
		"zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"backup_zone": {
			Type:     schema.TypeString,
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
		"enable_accelerated_networking": {
			Type:     schema.TypeBool,
			Required: true,
		},
		"accelerated_networking_enabled": {
			Type:     schema.TypeBool,
			Computed: true,
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
				Schema: AzureNetworkSchema(),
			},
		},
	}

	return schema
}

func AzureNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"subnetwork": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func ToAzureDeploymentObj(d *schema.ResourceData, deployment *models.Deployment) error {
	publicKey, err := utils.ReadTextFile(d.Get("instance_key").(string))
	if err != nil {
		return err
	}
	keyName, err := utils.GetPublicKeyName(*publicKey)
	if err != nil {
		return err
	}
	deployment.AzureDeployment = &models.AzureDeployment{
		InstanceKey: &models.InstanceKey{
			Content: *publicKey,
			Name:    keyName,
		},
		InstanceType:                 *utils.StrPtr(d.Get("instance_type").(string)),
		Region:                       *utils.StrPtr(d.Get("region").(string)),
		Vnet:                         *utils.StrPtr(d.Get("vnet").(string)),
		ResourceGroup:                *utils.StrPtr(d.Get("resource_group").(string)),
		EnableAcceleratedNetworking:  d.Get("enable_accelerated_networking").(bool),
		AcceleratedNetworkingEnabled: d.Get("accelerated_networking_enabled").(bool),
	}
	if f, exist := d.GetOkExists("zone"); exist {
		deployment.AzureDeployment.Zone = f.(string)
	}
	if f, exist := d.GetOkExists("backup_zone"); exist {
		deployment.AzureDeployment.BackupZone = f.(string)
	}
	if f, exist := d.GetOkExists("byoip"); exist {
		deployment.AzureDeployment.Byoip = f.(string)
	}
	if f, exist := d.GetOkExists("byoip_backup"); exist {
		deployment.AzureDeployment.ByoipBackup = f.(string)
	}
	if f, exist := d.GetOkExists("assign_public_ip"); exist {
		deployment.AzureDeployment.AssignPublicIP = f.(bool)
	} else {
		deployment.AzureDeployment.AssignPublicIP = true
	}

	if netIf, exist := d.GetOkExists("networks"); exist {
		net := netIf.([]interface{})
		deployment.AzureDeployment.Networks = make([]*models.AzureNetwork, 0)
		for i, _ := range net {
			netFieldFormat := fmt.Sprintf("%s.%d.%%s", "networks", i)
			temp := models.AzureNetwork{}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "name")); exist {
				temp.Name = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnetwork")); exist {
				temp.Subnetwork = f.(string)
			}
			deployment.AzureDeployment.Networks = append(deployment.AzureDeployment.Networks, &temp)
		}
	}

	return nil
}

func UpdateAzureDeploymentResource(deployment *models.AzureDeployment, d *schema.ResourceData) error {
	err := d.Set("region", deployment.Region)
	if err != nil {
		return err
	}

	err = d.Set("vnet", deployment.Vnet)
	if err != nil {
		return err
	}

	err = d.Set("instance_type", deployment.InstanceType)
	if err != nil {
		return err
	}

	err = d.Set("resource_group", deployment.ResourceGroup)
	if err != nil {
		return err
	}

	err = d.Set("zone", deployment.Zone)
	if err != nil {
		return err
	}

	err = d.Set("backup_zone", deployment.BackupZone)
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

	err = d.Set("enable_accelerated_networking", deployment.EnableAcceleratedNetworking)
	if err != nil {
		return err
	}

	err = d.Set("accelerated_networking_enabled", deployment.AcceleratedNetworkingEnabled)
	if err != nil {
		return err
	}

	err = d.Set("assign_public_ip", deployment.AssignPublicIP)
	if err != nil {
		return err
	}

	netList := make([]map[string]interface{}, 0)
	for _, net := range deployment.Networks {
		tmp := make(map[string]interface{})
		tmp["name"] = net.Name
		tmp["subnetwork"] = net.Subnetwork
		netList = append(netList, tmp)
	}
	err = d.Set("networks", netList)
	if err != nil {
		return err
	}

	return nil
}

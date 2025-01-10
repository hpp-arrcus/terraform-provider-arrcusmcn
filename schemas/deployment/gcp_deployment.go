package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GcpDeploymentResourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentResourceSchema()
	for k, v := range GcpDeploymentSchema() {
		deploymentSchema[k] = v
	}
	return deploymentSchema
}

func GcpDeploymentDatasourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentDataSchema()
	for k, v := range GcpDeploymentSchema() {
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

func GcpDeploymentSchema() map[string]*schema.Schema {
	schema := map[string]*schema.Schema{
		"instance_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
		"zone": {
			Type:     schema.TypeString,
			Required: true,
		},
		"backup_zone": {
			Type:     schema.TypeString,
			Required: true,
		},
		"byoip": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"byoip_backup": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"instance_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
				Schema: GcpNetworkSchema(),
			},
		},
	}

	return schema
}

func GcpNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"network": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subnetwork": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func ToGcpDeploymentObj(d *schema.ResourceData, deployment *models.Deployment) error {
	publicKey, err := utils.ReadTextFile(d.Get("instance_key").(string))
	if err != nil {
		return err
	}
	keyName, err := utils.GetPublicKeyName(*publicKey)
	if err != nil {
		return err
	}
	deployment.GcpDeployment = &models.GcpDeployment{
		InstanceKey: &models.InstanceKey{
			Content: *publicKey,
			Name:    keyName,
		},
		InstanceType: *utils.StrPtr(d.Get("instance_type").(string)),
		Region:       *utils.StrPtr(d.Get("region").(string)),
		Zone:         *utils.StrPtr(d.Get("zone").(string)),
	}

	if f, exist := d.GetOkExists("backup_zone"); exist {
		deployment.GcpDeployment.BackupZone = f.(string)
	}
	if f, exist := d.GetOkExists("byoip"); exist {
		deployment.GcpDeployment.Byoip = f.(string)
	}
	if f, exist := d.GetOkExists("byoip_backup"); exist {
		deployment.GcpDeployment.ByoipBackup = f.(string)
	}
	if f, exist := d.GetOkExists("assign_public_ip"); exist {
		deployment.GcpDeployment.AssignPublicIP = f.(bool)
	} else {
		deployment.GcpDeployment.AssignPublicIP = true
	}
	if tags, exist := d.GetOkExists("network_tags"); exist {
		tagList := tags.([]string)
		deployment.GcpDeployment.NetworkTags = tagList
	}

	if netIf, exist := d.GetOkExists("networks"); exist {
		net := netIf.([]interface{})
		deployment.GcpDeployment.Networks = make([]*models.GcpNetwork, 0)
		for i, _ := range net {
			netFieldFormat := fmt.Sprintf("%s.%d.%%s", "networks", i)
			temp := models.GcpNetwork{}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "name")); exist {
				temp.Name = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "network")); exist {
				temp.Network = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnetwork")); exist {
				temp.Subnetwork = f.(string)
			}
			deployment.GcpDeployment.Networks = append(deployment.GcpDeployment.Networks, &temp)
		}
	}

	return nil
}

func UpdateGcpDeploymentResource(deployment *models.GcpDeployment, d *schema.ResourceData) error {
	err := d.Set("region", deployment.Region)
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

	err = d.Set("instance_type", deployment.InstanceType)
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

	err = d.Set("assign_public_ip", deployment.AssignPublicIP)
	if err != nil {
		return err
	}

	err = d.Set("network_tags", deployment.NetworkTags)
	if err != nil {
		return err
	}

	netList := make([]map[string]interface{}, 0)
	for _, net := range deployment.Networks {
		tmp := make(map[string]interface{})
		tmp["name"] = net.Name
		tmp["network"] = net.Network
		tmp["subnetwork"] = net.Subnetwork
		netList = append(netList, tmp)
	}
	err = d.Set("networks", netList)
	if err != nil {
		return err
	}

	return nil
}

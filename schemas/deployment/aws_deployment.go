package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsDeploymentResourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentResourceSchema()
	for k, v := range AwsDeploymentSchema() {
		deploymentSchema[k] = v
	}

	return deploymentSchema
}

func AwsDeploymentDatasourceSchema() map[string]*schema.Schema {
	deploymentSchema := DeploymentDataSchema()
	for k, v := range AwsDeploymentSchema() {
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

func AwsDeploymentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"instance_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"region": {
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
		"zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"backup_zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"vpc_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"instance_type": {
			Type:     schema.TypeString,
			Required: true,
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
				Schema: AwsNetworkSchema(),
			},
		},
	}
}

func AwsNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"subnet_a": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subnet_b": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"private_subnet_route_table": {
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func ToAwsDeploymentObj(d *schema.ResourceData, deployment *models.Deployment) error {
	deployment.AwsDeployment = &models.AwsDeployment{
		InstanceKey:  *utils.StrPtr(d.Get("instance_key").(string)),
		InstanceType: *utils.StrPtr(d.Get("instance_type").(string)),
		Region:       *utils.StrPtr(d.Get("region").(string)),
		VpcID:        *utils.StrPtr(d.Get("vpc_id").(string)),
		Zone:         *utils.StrPtr(d.Get("zone").(string)),
	}

	if f, exist := d.GetOkExists("byoip"); exist {
		deployment.AwsDeployment.Byoip = f.(string)
	}
	if f, exist := d.GetOkExists("byoip_backup"); exist {
		deployment.AwsDeployment.ByoipBackup = f.(string)
	}
	if f, exist := d.GetOkExists("backup_zone"); exist {
		deployment.AwsDeployment.BackupZone = f.(string)
	}
	if f, exist := d.GetOkExists("assign_public_ip"); exist {
		deployment.AwsDeployment.AssignPublicIP = f.(bool)
	} else {
		deployment.AwsDeployment.AssignPublicIP = true
	}

	if netIf, exist := d.GetOkExists("networks"); exist {
		net := netIf.([]interface{})
		deployment.AwsDeployment.Networks = make([]*models.AwsNetwork, 0)
		for i, _ := range net {
			netFieldFormat := fmt.Sprintf("%s.%d.%%s", "networks", i)
			temp := models.AwsNetwork{}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "name")); exist {
				temp.Name = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnet_a")); exist {
				temp.SubnetA = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "subnet_b")); exist {
				temp.SubnetB = f.(string)
			}
			if f, exist := d.GetOkExists(fmt.Sprintf(netFieldFormat, "private_subnet_route_table")); exist {
				temp.PrivateSubnetRouteTable = f.(string)
			}
			deployment.AwsDeployment.Networks = append(deployment.AwsDeployment.Networks, &temp)
		}
	}

	return nil
}

func UpdateAwsDeploymentResource(deployment *models.AwsDeployment, d *schema.ResourceData) error {
	err := d.Set("instance_key", deployment.InstanceKey)
	if err != nil {
		return err
	}
	err = d.Set("region", deployment.Region)
	if err != nil {
		return err
	}
	err = d.Set("vpc_id", deployment.VpcID)
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
	err = d.Set("zone", deployment.Zone)
	if err != nil {
		return err
	}
	err = d.Set("backup_zone", deployment.BackupZone)
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
		tmp["subnet_a"] = net.SubnetA
		tmp["subnet_b"] = net.SubnetB
		tmp["private_subnet_route_table"] = net.PrivateSubnetRouteTable
		netList = append(netList, tmp)
	}
	err = d.Set("networks", netList)
	if err != nil {
		return err
	}

	return nil
}

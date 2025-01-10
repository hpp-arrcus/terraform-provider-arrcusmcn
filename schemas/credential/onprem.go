package schemas

import (
	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OnpremCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"server_ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ssh_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"data_if_name": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func OnpremCredentialDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"server_ip": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"user_name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"ssh_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"data_if_name": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func ToOnpremCredSchema(cred *models.Credentials, d *schema.ResourceData) error {
	err := d.Set("server_ip", cred.Credentials.OnpremCredentials.ServerIP)
	if err != nil {
		return nil
	}
	err = d.Set("user_name", cred.Credentials.OnpremCredentials.UserName)
	if err != nil {
		return nil
	}
	err = d.Set("domain", cred.Credentials.OnpremCredentials.Domain)
	if err != nil {
		return nil
	}
	err = d.Set("data_if_name", cred.Credentials.OnpremCredentials.DataIfName)
	if err != nil {
		return nil
	}
	return nil
}

func ToOnpremCredObj(d *schema.ResourceData) (*models.Credentials, error) {
	onpremCred := models.OnpremCredentials{}
	if v, exists := d.GetOk("server_ip"); exists {
		onpremCred.ServerIP = *utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk("user_name"); exists {
		onpremCred.UserName = *utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk("ssh_key"); exists {
		onpremCred.SSHKey = *utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk("domain"); exists {
		onpremCred.Domain = *utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk("data_if_name"); exists {
		ifList := v.([]interface{})
		temp := make([]string, 0)
		for _, ifname := range ifList {
			temp = append(temp, ifname.(string))
		}
		onpremCred.DataIfName = temp
	}

	cred := models.Credentials{}
	cred.Name = utils.StrPtr(d.Get("name").(string))
	prod := models.ProvidersOnpremise
	cred.Provider = &prod
	cred.Credentials.OnpremCredentials = onpremCred
	return &cred, nil
}

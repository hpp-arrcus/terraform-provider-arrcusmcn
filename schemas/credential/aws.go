package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AwsCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"access_key": {
			Type:     schema.TypeString,
			Required: true,
		},
		"secret_key": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func AwsCredentialDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"access_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"secret_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func ToAwsCredSchema(cred *models.Credentials, d *schema.ResourceData) error {
	err := d.Set("access_key", cred.Credentials.AwsCredentials.AccessKey)
	if err != nil {
		return nil
	}
	err = d.Set("secret_key", cred.Credentials.AwsCredentials.SecretKey)
	if err != nil {
		return nil
	}
	return nil
}

func ToAwsCredObj(d *schema.ResourceData) (*models.Credentials, error) {
	awsCred := models.AwsCredentials{}
	if v, exists := d.GetOk("access_key"); exists {
		awsCred.AccessKey = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("access_key is missing.")
	}

	if v, exists := d.GetOk("secret_key"); exists {
		awsCred.SecretKey = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("secret_key is missing.")
	}

	cred := models.Credentials{}
	cred.Name = utils.StrPtr(d.Get("name").(string))
	cred.Credentials.AwsCredentials = awsCred
	prod := models.ProvidersAws
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

package schemas

import (
	"errors"
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func OciCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user": {
			Type:     schema.TypeString,
			Required: true,
		},
		"identity_domain": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenancy": {
			Type:     schema.TypeString,
			Required: true,
		},
		"region": {
			Type:     schema.TypeString,
			Required: true,
		},
		"key_file": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func OciCredentialDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"user": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"identity_domain": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tenancy": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"region": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"key_file": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func ToOciCredSchema(cred *models.Credentials, d *schema.ResourceData) error {
	err := d.Set("user", cred.Credentials.OciCredentials.User)
	if err != nil {
		return nil
	}
	err = d.Set("identity_domain", cred.Credentials.OciCredentials.IdentityDomain)
	if err != nil {
		return nil
	}

	err = d.Set("tenancy", cred.Credentials.OciCredentials.Tenancy)
	if err != nil {
		return nil
	}

	err = d.Set("region", cred.Credentials.OciCredentials.Region)
	if err != nil {
		return nil
	}
	return nil
}

func ToOciCredObj(d *schema.ResourceData) (*models.Credentials, error) {
	ociCred := models.OciCredentials{}
	if v, exists := d.GetOk("user"); exists {
		ociCred.User = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("user is missing.")
	}

	if v, exists := d.GetOk("identity_domain"); exists {
		ociCred.IdentityDomain = *utils.StrPtr(v.(string))
	} else {
		ociCred.IdentityDomain = "DEFAULT"
	}

	if v, exists := d.GetOk("tenancy"); exists {
		ociCred.Tenancy = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("tenancy is missing.")
	}

	if v, exists := d.GetOk("region"); exists {
		ociCred.Region = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("region is missing.")
	}

	if v, exists := d.GetOk("key_file"); exists {
		ociKey, err := utils.ReadTextFile(v.(string))
		if err != nil {
			return nil, fmt.Errorf("Can't read oci key_file %s: %s", v.(string), err)
		}
		ociCred.KeyFile = ociKey
	} else {
		return nil, errors.New("key_file is missing.")
	}

	cred := models.Credentials{}
	cred.Name = utils.StrPtr(d.Get("name").(string))
	cred.Credentials.OciCredentials = ociCred
	prod := models.ProvidersOci
	cred.Provider = &prod
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

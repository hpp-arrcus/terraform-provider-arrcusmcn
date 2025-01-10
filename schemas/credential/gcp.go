package schemas

import (
	"encoding/json"
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func GcpCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"account_key_file": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"project_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"private_key_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"private_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"client_email": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"auth_uri": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"token_uri": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"auth_provider_x509_cert_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"client_x509_cert_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func GcpCredentialDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"account_key_file": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"project_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_key_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"private_key": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_email": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"auth_uri": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"token_uri": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"auth_provider_x509_cert_url": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_x509_cert_url": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func ToGcpCredSchema(cred *models.Credentials, d *schema.ResourceData) error {
	err := d.Set("type", cred.Credentials.GcpCredentials.Type)
	if err != nil {
		return err
	}
	err = d.Set("project_id", cred.Credentials.GcpCredentials.ProjectID)
	if err != nil {
		return err
	}
	err = d.Set("private_key_id", cred.Credentials.GcpCredentials.PrivateKeyID)
	if err != nil {
		return err
	}
	err = d.Set("private_key", cred.Credentials.GcpCredentials.PrivateKey)
	if err != nil {
		return err
	}
	err = d.Set("client_email", cred.Credentials.GcpCredentials.ClientEmail)
	if err != nil {
		return err
	}
	err = d.Set("client_id", cred.Credentials.GcpCredentials.ClientID)
	if err != nil {
		return err
	}
	err = d.Set("auth_uri", cred.Credentials.GcpCredentials.AuthURI)
	if err != nil {
		return err
	}
	err = d.Set("token_uri", cred.Credentials.GcpCredentials.TokenURI)
	if err != nil {
		return err
	}
	err = d.Set("auth_provider_x509_cert_url", cred.Credentials.GcpCredentials.AuthProviderX509CertURL)
	if err != nil {
		return err
	}
	err = d.Set("client_x509_cert_url", cred.Credentials.GcpCredentials.ClientX509CertURL)
	if err != nil {
		return err
	}
	return nil
}

func ToGcpCredObj(d *schema.ResourceData) (*models.Credentials, error) {
	gcpCred, err := LoadAccountKeyFile(d)
	if err != nil {
		return nil, err
	}
	if gcpCred == nil {
		if v, exists := d.GetOk("type"); exists {
			gcpCred.Type = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("type is missing.")
		}

		if v, exists := d.GetOk("project_id"); exists {
			gcpCred.ProjectID = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("project_id is missing.")
		}

		if v, exists := d.GetOk("private_key_id"); exists {
			gcpCred.PrivateKeyID = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("private_key_id is missing.")
		}

		if v, exists := d.GetOk("private_key"); exists {
			gcpCred.PrivateKey = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("private_key is missing.")
		}

		if v, exists := d.GetOk("client_email"); exists {
			gcpCred.ClientEmail = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("client_email is missing.")
		}

		if v, exists := d.GetOk("client_id"); exists {
			gcpCred.ClientID = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("client_id is missing.")
		}

		if v, exists := d.GetOk("auth_uri"); exists {
			gcpCred.AuthURI = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("auth_uri is missing.")
		}

		if v, exists := d.GetOk("token_uri"); exists {
			gcpCred.TokenURI = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("token_uri is missing.")
		}

		if v, exists := d.GetOk("auth_provider_x509_cert_url"); exists {
			gcpCred.AuthProviderX509CertURL = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("auth_provider_x509_cert_url is missing.")
		}

		if v, exists := d.GetOk("client_x509_cert_url"); exists {
			gcpCred.ClientX509CertURL = *utils.StrPtr(v.(string))
		} else {
			return nil, errors.New("client_x509_cert_url is missing.")
		}
	}
	cred := models.Credentials{}
	cred.Name = utils.StrPtr(d.Get("name").(string))
	prod := models.ProvidersGcp
	cred.Provider = &prod
	cred.Credentials.GcpCredentials = *gcpCred
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

func LoadAccountKeyFile(d *schema.ResourceData) (*models.GcpCredentials, error) {
	if v, exist := d.GetOk("account_key_file"); exist {
		path := v.(string)
		jsonString, err := utils.ReadTextFile(path)
		if err != nil {
			return nil, err
		}
		gcpCred := models.GcpCredentials{}
		err = json.Unmarshal([]byte(*jsonString), &gcpCred)
		if err != nil {
			return nil, err
		}
		return &gcpCred, nil
	} else {
		return nil, errors.New("Account_key_file is not given.")
	}
}

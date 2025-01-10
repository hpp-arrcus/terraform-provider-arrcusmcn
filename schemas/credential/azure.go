package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AzureCredentialSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subscription_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"client_secret": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_id": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func AzureCredentialDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"subscription_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"client_secret": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"tenant_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func ToAzureCredSchema(cred *models.Credentials, d *schema.ResourceData) error {
	err := d.Set("subscription_id", cred.Credentials.AzureCredentials.AzureSubscriptionID)
	if err != nil {
		return nil
	}
	err = d.Set("client_id", cred.Credentials.AzureCredentials.AzureClientID)
	if err != nil {
		return nil
	}
	err = d.Set("client_secret", cred.Credentials.AzureCredentials.AzureClientSecret)
	if err != nil {
		return nil
	}
	err = d.Set("tenant_id", cred.Credentials.AzureCredentials.AzureTenantID)
	if err != nil {
		return nil
	}
	return err
}

func ToAzureCredObj(d *schema.ResourceData) (*models.Credentials, error) {
	azureCred := models.AzureCredentials{}
	if v, exists := d.GetOk("subscription_id"); exists {
		azureCred.AzureSubscriptionID = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("subscription_id is missing.")
	}

	if v, exists := d.GetOk("client_id"); exists {
		azureCred.AzureClientID = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("client_id is missing.")
	}

	if v, exists := d.GetOk("client_secret"); exists {
		azureCred.AzureClientSecret = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("client_secret is missing.")
	}

	if v, exists := d.GetOk("tenant_id"); exists {
		azureCred.AzureTenantID = *utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("tenant_id is missing.")
	}

	cred := models.Credentials{}
	cred.Name = utils.StrPtr(d.Get("name").(string))
	prod := models.ProvidersAzure
	cred.Provider = &prod
	cred.Credentials.AzureCredentials = azureCred
	cred.Credentials.DataIfName = make([]string, 0)
	return &cred, nil
}

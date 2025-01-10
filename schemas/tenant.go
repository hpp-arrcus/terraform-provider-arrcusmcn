package schemas

import (
	"errors"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"organization": {
			Type:     schema.TypeString,
			Required: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Required: true,
		},
		"defaultuser": {
			Type:     schema.TypeList,
			Required: true,
			MaxItems: 1,
			MinItems: 1,
			Elem: &schema.Resource{
				Schema: UserSchema(),
			},
		},
		"numdeployments": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"numconnections": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"is_default": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func TenantDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"organization": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"domain": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"defaultuser": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Resource{
				Schema: UserSchema(),
			},
		},
		"numdeployments": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"numconnections": {
			Type:     schema.TypeInt,
			Computed: true,
		},
		"is_default": {
			Type:     schema.TypeBool,
			Computed: true,
		},
	}
}

func ToTenantSchema(tenant *models.Tenant, d *schema.ResourceData) error {
	err := d.Set("name", *tenant.Name)
	if err != nil {
		return err
	}
	err = d.Set("organization", tenant.Organization)
	if err != nil {
		return err
	}
	err = d.Set("domain", *tenant.Domain)
	if err != nil {
		return err
	}
	err = d.Set("numdeployments", tenant.Numdeployments)
	if err != nil {
		return err
	}
	err = d.Set("numconnections", tenant.Numconnections)
	if err != nil {
		return err
	}
	err = d.Set("is_default", *tenant.IsDefault)
	if err != nil {
		return err
	}

	defaultUserList := make([]map[string]interface{}, 1)
	defaultUserList[0] = make(map[string]interface{})
	dfu := tenant.Defaultuser
	defaultUserList[0]["username"] = *dfu.Username
	defaultUserList[0]["name"] = dfu.Name
	defaultUserList[0]["password"] = *dfu.Password
	defaultUserList[0]["email"] = *dfu.Email
	defaultUserList[0]["is_default"] = *dfu.IsDefault
	defaultUserList[0]["roles"] = dfu.Roles
	err = d.Set("defaultuser", defaultUserList)
	if err != nil {
		return err
	}

	return nil
}

func ToTenantObj(d *schema.ResourceData) (*models.Tenant, error) {
	tenant := models.Tenant{}
	tenant.Defaultuser = &models.User{}
	if v, exists := d.GetOk("name"); exists {
		tenant.Name = utils.StrPtr(v.(string))
	} else {
		return nil, errors.New("name is missing.")
	}
	if v, exists := d.GetOk("organization"); exists {
		tenant.Organization = v.(string)
	}

	if v, exists := d.GetOk("domain"); exists {
		tenant.Domain = utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk("numdeployments"); exists {
		tenant.Numdeployments = int64(v.(int))
	}

	if v, exists := d.GetOk("numconnections"); exists {
		tenant.Numconnections = int64(v.(int))
	}

	if v, exists := d.GetOk("is_default"); exists {
		vptr := v.(bool)
		tenant.IsDefault = &vptr
	} else {
		tenant.IsDefault = nil
	}

	if _, exists := d.GetOk("defaultuser"); exists {
		user, err := ToUserObj(d, "defaultuser.0")
		if err != nil {
			return nil, err
		}
		tenant.Defaultuser = user
	} else {
		tenant.IsDefault = nil
	}

	return &tenant, nil
}

package schemas

import (
	"fmt"

	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func UserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:     schema.TypeString,
			Required: true,
		},
		"email": {
			Type:     schema.TypeString,
			Required: true,
		},
		"is_default": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"roles": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func UserDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"password": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"email": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"is_default": {
			Type:     schema.TypeBool,
			Computed: true,
		},
		"roles": {
			Type:     schema.TypeList,
			Computed: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
	}
}

func ToUserSchema(user *models.User, d *schema.ResourceData, prefix string) error {
	var prefixFormat string
	if prefix != "" {
		prefixFormat = prefix + "."
	} else {
		prefixFormat = ""
	}

	err := d.Set(fmt.Sprintf("%s%s", prefixFormat, "username"), *user.Username)
	if err != nil {
		return err
	}
	err = d.Set(fmt.Sprintf("%s%s", prefixFormat, "name"), user.Name)
	if err != nil {
		return err
	}
	err = d.Set(fmt.Sprintf("%s%s", prefixFormat, "email"), *user.Email)
	if err != nil {
		return err
	}
	err = d.Set(fmt.Sprintf("%s%s", prefixFormat, "password"), "")
	if err != nil {
		return err
	}
	err = d.Set(fmt.Sprintf("%s%s", prefixFormat, "is_default"), *user.IsDefault)
	if err != nil {
		return err
	}

	err = d.Set(fmt.Sprintf("%s%s", prefixFormat, "roles"), user.Roles)
	if err != nil {
		return err
	}
	return nil
}

func ToUserObj(d *schema.ResourceData, prefix string) (*models.User, error) {
	var prefixFormat string
	if prefix != "" {
		prefixFormat = prefix + "."
	} else {
		prefixFormat = ""
	}

	user := models.User{}
	if v, exists := d.GetOk(fmt.Sprintf("%s%s", prefixFormat, "username")); exists {
		user.Name = v.(string)
		user.Username = utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk(fmt.Sprintf("%s%s", prefixFormat, "email")); exists {
		user.Email = utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk(fmt.Sprintf("%s%s", prefixFormat, "password")); exists {
		user.Password = utils.StrPtr(v.(string))
	}

	if v, exists := d.GetOk(fmt.Sprintf("%s%s", prefixFormat, "is_default")); exists {
		user.IsDefault = utils.BoolPtr(v.(bool))
	}

	if v, exists := d.GetOk(fmt.Sprintf("%s%s", prefixFormat, "roles")); exists {
		roles := make([]models.Rolename, 0)
		rolesIf := v.([]interface{})
		for i, _ := range rolesIf {
			roleFormat := fmt.Sprintf("%s.%d", fmt.Sprintf("%s%s", prefixFormat, "roles"), i)
			switch d.Get(roleFormat).(string) {
			case string(models.RolenameArcOrchAdmin):
				roles = append(roles, models.RolenameArcOrchAdmin)
			case string(models.RolenameTenantAdmin):
				roles = append(roles, models.RolenameTenantAdmin)
			case string(models.RolenameTenantOperator):
				roles = append(roles, models.RolenameTenantOperator)
			case string(models.RolenameDashboardReader):
				roles = append(roles, models.RolenameDashboardReader)
			}
		}
		user.Roles = roles
	}

	return &user, nil
}

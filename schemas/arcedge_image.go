package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ArcedgeImageResourceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"image_id": {
			Type:     schema.TypeString,
			Required: true,
		},
		"provider": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

func ArcedgeImageDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"version": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"image_id": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"provider": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Computed: true,
		},
	}
}

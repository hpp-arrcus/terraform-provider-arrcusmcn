package datasource

// import (
// 	"context"

// 	resource "github.com/Arrcus/terraform-provider-arrcusmcn/resource"
// 	"github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func DataSourceTenant() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceTenantRead,
// 		Schema:      schemas.TenantDataSchema(),
// 	}
// }

// func dataSourceTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	tenant, err := resource.ReadTenant(ctx, baseURL, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	d.SetId(tenant.ID.String())
// 	return diags
// }

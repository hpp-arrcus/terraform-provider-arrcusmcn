package datasource

// import (
// 	"context"

// 	resource "github.com/Arrcus/terraform-provider-arrcusmcn/resource"
// 	"github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func DataSourceUser() *schema.Resource {
// 	return &schema.Resource{
// 		ReadContext: dataSourceUserRead,
// 		Schema:      schemas.UserDataSchema(),
// 	}
// }

// func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	tenant := m.(map[string]string)["tenant"]
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	user, err := resource.ReadUsers(ctx, baseURL, tenant, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	d.SetId(user.ID.String())
// 	return diags
// }

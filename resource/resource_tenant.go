package resource

// import (
// 	"context"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	"github.com/Arrcus/terraform-provider-arrcusmcn/models"
// 	schemas "github.com/Arrcus/terraform-provider-arrcusmcn/schemas"
// 	"github.com/Arrcus/terraform-provider-arrcusmcn/utils"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// func ResourceTenant() *schema.Resource {
// 	return &schema.Resource{
// 		CreateContext: resourceTenantCreate,
// 		ReadContext:   resourceTenantRead,
// 		UpdateContext: resourceTenantUpdate,
// 		DeleteContext: resourceTenantDelete,
// 		Schema:        schemas.TenantSchema(),
// 		Importer: &schema.ResourceImporter{
// 			StateContext: schema.ImportStatePassthroughContext,
// 		},
// 	}
// }

// func resourceTenantCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	err := CreateTenant(ctx, baseURL, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	return diags
// }

// func resourceTenantRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	_, err := ReadTenant(ctx, baseURL, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	return diags
// }

// func resourceTenantUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	if d.HasChange("name") {
// 		return diag.Errorf(fmt.Sprintf("name can't be changed"))
// 	}
// 	if d.HasChange("username") {
// 		return diag.Errorf("username can't be changed")
// 	}
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	err := UpdateTenant(ctx, baseURL, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	return diags
// }

// func resourceTenantDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
// 	var diags diag.Diagnostics
// 	accessToken := m.(map[string]string)["access_token"]
// 	baseURL := m.(map[string]string)["baseUrl"]
// 	err := DeleteTenant(ctx, baseURL, accessToken, d)
// 	if err != nil {
// 		return diag.FromErr(err)
// 	}
// 	return diags
// }

// func CreateTenant(
// 	ctx context.Context,
// 	baseURL string,
// 	accessToken string,
// 	d *schema.ResourceData,
// ) error {
// 	url := baseURL + "tenants"
// 	tenant, err := schemas.ToTenantObj(d)
// 	if err != nil {
// 		return err
// 	}

// 	res, err := utils.PostRequest(url, *tenant, accessToken)
// 	if err != nil {
// 		return err
// 	}
// 	resBody, _ := ioutil.ReadAll(res.Body)
// 	resTenant := models.Tenant{}
// 	err = json.Unmarshal(resBody, &resTenant)
// 	if err != nil {
// 		return err
// 	}
// 	d.SetId(resTenant.ID.String())
// 	return nil
// }

// func ReadTenant(
// 	ctx context.Context,
// 	baseURL string,
// 	accessToken string,
// 	d *schema.ResourceData,
// ) (*models.Tenant, error) {
// 	tenant, err := schemas.ToTenantObj(d)
// 	if err != nil {
// 		return nil, err
// 	}
// 	url := baseURL + "tenants/" + *tenant.Name
// 	res, err := utils.GetRequest(url, accessToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resBody, _ := ioutil.ReadAll(res.Body)
// 	if res.StatusCode != 200 {
// 		return nil, errors.New(string(resBody))
// 	}
// 	log.Println(string(resBody))
// 	resTenant := models.Tenant{}
// 	err = json.Unmarshal(resBody, &resTenant)
// 	if err != nil {
// 		return nil, err
// 	}
// 	log.Printf("%v\n", resTenant)
// 	err = schemas.ToTenantSchema(&resTenant, d)
// 	return &resTenant, nil
// }

// func UpdateTenant(
// 	ctx context.Context,
// 	baseURL string,
// 	accessToken string,
// 	d *schema.ResourceData,
// ) error {
// 	tenant, err := schemas.ToTenantObj(d)
// 	if err != nil {
// 		return err
// 	}
// 	url := baseURL + "tenants/" + *tenant.Name
// 	res, err := utils.PutRequest(url, *tenant, accessToken)
// 	if err != nil {
// 		return err
// 	}
// 	resBody, _ := ioutil.ReadAll(res.Body)
// 	if res.StatusCode != 204 && res.StatusCode != 200 {
// 		return errors.New(string(resBody))
// 	}
// 	resTenant := models.Tenant{}
// 	err = json.Unmarshal(resBody, &resTenant)
// 	if err != nil {
// 		return err
// 	}
// 	err = schemas.ToTenantSchema(&resTenant, d)

// 	return nil
// }

// func DeleteTenant(
// 	ctx context.Context,
// 	baseURL string,
// 	accessToken string,
// 	d *schema.ResourceData,
// ) error {
// 	tenant, err := schemas.ToTenantObj(d)
// 	if err != nil {
// 		return err
// 	}
// 	url := baseURL + "tenants/" + *tenant.Name
// 	err = utils.DeleteRequest(url, accessToken)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

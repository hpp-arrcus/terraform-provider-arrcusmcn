# <resource name> arrcusmcn_tenant

The arrcusmcn_tenant resource is for the use of creation and management of Arrcus MCN accounts.

## Example Usage

```hcl
resource "arrcusmcn_tenant" "tenant" {
  name = "example"
  organization = "The Example Co."
  domain = "example.com"
  defaultuser_name = "John Example"
  defaultuser_username = "admin@example.com"
  defaultuser_password = "example123"
  defaultuser_email = "john@example.com"
}
```

## Argument Reference

* `username` - (Required) Name for the tenant.
* `organization` - (Required) Organization name for the tenant.
* `domain` - (Required) Domain for the tenant Domain will be part of username for all users under the tenant.
* `defaultuser_name` - (Required) Full name for the tenant administrator.
* `defaultuser_username` - (Required) Name for the tenant administrator.
* `defaultuser_password` - (Required) Password used for login as the tenant administrator.
* `defaultuser_email` - (Required) Email for the tenant administrator.

## Attribute Reference

* `id` - a unique identifier for the resource
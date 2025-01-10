# <resource name> arcorch_azure_cred

The arcorch_azure_cred resource is for the use of creation and management of Azure accounts. The accounts created will be managed in the ArcOrchestrator. 

## Example Usage

```hcl
resource "arcorch_azure_cred" "azure_cred" {
  name = "azure_cred"
  subscription_id = "subscription_id"
  client_id = "client_id"
  client_secret = "client_secret"
  tenant_id = "tenant_id"
}
```

## Argument Reference

* `name` - (Required) A unique name for the given Azure credential.
* `subscription_id` - (Required) Subscription ID of the Azure account.
* `client_id` - (Required) Client ID associated with the application of the Azure account.
* `client_secret` - (Required) Client ID associated with the application of the Azure account.
* `tenant_id` - (Required) Tenant ID of the Azure account.

## Attribute Reference

* `id` - a unique identifier for the resource

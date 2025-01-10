# <resource name> arcorch_gcp_cred

The arcorch_gcp_cred data source provides details of a specific GCP credential created on the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arcorch_gcp_cred" "gcp_cred" {
  name = "gcp_cred"
}
```

## Argument Reference

* `name` - (Required) A unique name for the specific Gcp credential

## Attribute Reference

* `id` - A unique identifier for the resource
* `account_key_file` - Full path and filename of the gcp service account key file.
* `auth_provider_x509_cert_url` - The auth_provider_x509_cert_url field from gcp service account key file.
* `auth_uri` - The auth_uri field from gcp service account key file.
* `client_email` - The client_email field from gcp service account key file.
* `client_id` - The client_id field from gcp service account key file.
* `client_x509_cert_url` - The client_x509_cert_url field from gcp service account key file. 
* `private_key` - The private_key field from gcp service account key file.
* `private_key_id` - The private_key_id field from gcp service account key file.
* `project_id` - The project_id field from gcp service account key file.
* `token_uri` - The token_uri field from gcp service account key file.
* `type` - The type field from gcp service account key file.

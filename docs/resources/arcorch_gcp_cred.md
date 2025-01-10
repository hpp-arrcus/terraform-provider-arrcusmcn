# <resource name> arcorch_gcp_cred

The arcorch_gcp_cred resource is for the use of creation and management of GCP accounts. The accounts created will be managed in the ArcOrchestrator. 

To create an arcorch_gcp_cred, you can import a GCP service account key file by supplying the complete path and filename to the account_key_file field. Alternatively, you can manually complete all the necessary fields according to the requirements of arcorch_gcp_cred. More detailed information can be found at https://cloud.google.com/iam/docs/keys-create-delete.

## Example Usage

```hcl
resource "arcorch_gcp_cred" "gcp_cred" {
  name = "gcp-cred"
  account_key_file = "~/gcp_service_account_key.json"
  auth_provider_x509_cert_url = "auth_provider_x509_cert_url"
  auth_uri = "auth_uri"
  client_email = "client_email"
  client_id = "client_id"
  client_x509_cert_url = "client_x509_cert_url"
  private_key = "private_key"
  private_key_id = "private_key_id"
  project_id = "project_id"
  token_uri = "token_uri"
  type = "service_account"
}
```

## Argument Reference

* `name` - (Required) A unique name for the given GCP credential
* `account_key_file` - (Optional) Full path and filename of the gcp service account key fie. If account_key_file is given, other fields are not needed. 
* `auth_provider_x509_cert_url` - (Optional) The auth_provider_x509_cert_url field from gcp service account key fie. This is needed if account_key_file is not given.
* `auth_uri` - (Optional) The auth_uri field from gcp service account key fie. This is needed if account_key_file is not given.
* `client_email` - (Optional) The client_email field from gcp service account key fie. This is needed if account_key_file is not given.
* `client_id` - (Optional) The client_id field from gcp service account key fie. This is needed if account_key_file is not given.
* `client_x509_cert_url` - (Optional) The client_x509_cert_url field from gcp service account key fie. This is needed if account_key_file is not given.
* `private_key` - (Optional) The private_key field from gcp service account key fie. This is needed if account_key_file is not given.
* `private_key_id` - (Optional) The private_key_id field from gcp service account key fie. This is needed if account_key_file is not given.
* `project_id` - (Optional) The project_id field from gcp service account key fie. This is needed if account_key_file is not given.
* `token_uri` - (Optional) The token_uri field from gcp service account key fie. This is needed if account_key_file is not given.
* `type` - (Optional) The type field from gcp service account key fie. This is needed if account_key_file is not given.


## Attribute Reference

* `id` - a unique identifier for the resource

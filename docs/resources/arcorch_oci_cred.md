# <resource name> arcorch_oci_cred

The arcorch_oci_cred resource is for the use of creation and management of OCI accounts. The accounts created will be managed in the ArcOrchestrator. 

All of field arcorch_oci_cred needs can be found from oci configuration file. More detailed information can be found at https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm.

## Example Usage

```hcl
resource "arcorch_oci_cred" "oci_cred" {
  name = "oci-cred"
  user = "user-ocid"
  identity_domain = "DEFAULT"
  tenancy = "tenancy-ocid"
  region = "region"
  key_file = "~/oci-cred.pem"
}
```

## Argument Reference

* `name` - (Required) A unique name for the given AWS credential
* `user` - (Required) Ocid of user.
* `tenancy` - (Required) Ocid of tenancy.
* `region` - (Required) An Oracle Cloud Infrastructure region.
* `key_file` - (Required) Full path and filename of the private key.
* `identity_domain` - (Optional) identity domain. Default value is DEFAULT.


## Attribute Reference

* `id` - A unique identifier for the resource

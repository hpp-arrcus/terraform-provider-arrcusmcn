# <resource name> arcorch_oci_cred

The arcorch_oci_cred data source provides details of a specific OCI credential created on the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arcorch_oci_cred" "oci_cred" {
  name = "oci_cred"
}
```

## Argument Reference

* `name` - (Required) A unique name for the specific OCI credential

## Attribute Reference

* `id` - A unique identifier for the resource
* `user` - Ocid of user.
* `tenancy` - Ocid of tenancy.
* `region` - An Oracle Cloud Infrastructure region.
* `key_file` - Full path and filename of the private key.
* `identity_domain` - Identity domain.

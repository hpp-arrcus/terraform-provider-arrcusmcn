# <resource name> arrcusmcn_user

The arrcusmcn_user data source provides details of Arrcus MCN account used in Provder which currently logging in.

## Example Usage

```hcl
data "arrcusmcn_user" "user" {}

```

## Attribute Reference

* `id` - a unique identifier for the resource
* `username` - Name for the account
* `password` - Password for the account
* `email` - Email address for the account

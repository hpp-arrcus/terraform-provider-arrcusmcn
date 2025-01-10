# <resource name> arrcusmcn_aws_cred

The arrcusmcn_aws_cred resource is for the use of creation and management of AWS accounts. The accounts created will be managed in the ArcOrchestrator. 

## Example Usage

```hcl
resource "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
  access_key = "ABCDEFG"
  secret_key = "ABCDEFG"
}
```

## Argument Reference

* `name` - (Required) A unique name for the given AWS credential.
* `access_key` - (Required) Access key of an AWS account.
* `secret_key` - (Required) Secret key of an AWS account.

## Attribute Reference

* `id` - a unique identifier for the resource

# <resource name> arrcusmcn_aws_cred

The arrcusmcn_aws_cred data source provides details of a specific AWS credential created on the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
}
```

## Argument Reference

* `name` - (Required) a unique name for the specific AWS credential

## Attribute Reference

* `id` - a unique identifier for the resource
* `access_key` - Access key of an AWS account
* `secret_key` - Secret key of an AWS account

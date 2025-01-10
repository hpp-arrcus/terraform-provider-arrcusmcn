# <provider> Arrcus MCN Provider

Arrcus MCN provider is used to manage ArcEdge deployments running on different cloud platforms using Arrcus ArcOrchestrator.

## Example Usage

```hcl
provider "arrcusmcn" {
  endpoint = "https://arcorch.com"
  port = "443"
  username = "admin"
  password = "password"
  api_token = "abcdefg"
}
```

## Argument Reference

- `endpoint` - (Required) ArcOrchestrator ip address
- `port` - (Required) ArcOrchestrator port
- `username` - (Required) account username which will be used to login to ArcOrchestrator.
- `password` - (Optional) account password corresponding to given username. If password is not given, api_token is required.
- `api_token` - (Optional) api_key corresponding to given username. If api_token is not given, password is required.

# <resource name> arcorch_onprem_cred

The arcorch_onprem_cred data source provides details of a specific on-premise credential created on the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arcorch_onprem_cred" "onprem_cred" {
  name = "onprem_cred"
}
```

## Argument Reference

* `name` - (Required) A unique name for the specific on-premise credential

## Attribute Reference

* `id` - A unique identifier for the resource
* `server_ip` - IP address of the on-premise server.
* `user_name` - Username for SSH login to the on-premise server.
* `ssh_key` - Private key file for SSH login to the on-premise server. The on-premise server should have coresponding public key in the authorized_keys file.
* `data_if_name` - Hypervisor virtual function vf pools for ArcEdge interfaces.
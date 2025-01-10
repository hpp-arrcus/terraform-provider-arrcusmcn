# <resource name> arcorch_onprem_cred

The arcorch_onprem_cred resource is for the use of creation and management of onprem accounts. The accounts created will be managed in the ArcOrchestrator. 

## Example Usage

```hcl
resource "arcorch_onprem_cred" "onprem_cred" {
  name = "onprem-cred"
  server_ip = "1.2.3.4"
  user_name = "user"
  ssh_key = "~/onprem_cred.pem"
  data_if_name {
    "if_name_1",
    "if_name_2"
  }
}
```

## Argument Reference

* `name` - (Required) A unique name for the given onpremise credential
* `server_ip` - (Required) IP address of the on-premise server.
* `user_name` - (Required) Username for SSH login to the on-premise server.
* `ssh_key` - (Required) Private key file for SSH login to the on-premise server. The on-premise server should have coresponding public key in the authorized_keys file.
* `data_if_name` - (Optional) Hypervisor virtual function vf pools for ArcEdge interfaces.

## Attribute Reference

* `id` - A unique identifier for the resource

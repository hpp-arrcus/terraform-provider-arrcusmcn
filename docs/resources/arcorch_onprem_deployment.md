# <resource name> arcorch_onprem_deployment

arcorch_onprem_deployment resource will be used to create ArcEdge deployments in the desired onprem account. 

## Example Usage

```hcl
resource "arrcusmcn_onprem_deployment" "arrcusmcn_onprem" {
    credential_name = "onprem-cred"
    enable_high_availability = false
    enable_private_subnet = false
    name = "onprem-arcedge"
    vcpus = 2
    vm_memory = 4096
    private_ip = "172.16.102.115"
    subnet_mask = 24
    default_gw = "172.16.102.1"
    public_ip = "172.16.102.115"
    ssh_psw = "arrcus"
    source_image {
      image_id = "ArcEdge_5.2.1B.20230601.qcow2"
    }
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on onpremise server.
* `credential_name ` - (Required) The name of the onprem credential where the ArcEdge will be deployed.
* `vcpus ` - (Required) Number of CPUs of the ArcEdge deployed.
* `vm_memory ` - (Required) Memory size(MB) of the ArcEdge deployed.
* `public_ip ` - (Required) Public IP address assigned to the ArcEdge deployed.
* `private_ip ` - (Required) Private IP address assigned to the ArcEdge deployed.
* `subnet_mask ` - (Required) Subnet mask of the ArcEdge deployed.
* `default_gw ` - (Required) Default gateway of the ArcEdge deployed.
* `ssh_psw ` - (Optional) Password for ssh into the ArcEdge. If the field is not provided, ssh to the ArcEdge will be disabled.
* `source_image ` - (Required) Image of the ArcEdge.
* `image_id ` - (Required) The ocid of the image.
## Attribute Reference

* `credential_id` - A unique identifier of the credential for the deployed Arcedge
* `arc_orch_ip ` -  Public IP of the ArcEdge that has been deployed.
* `action ` -  Action initiated by the ArcOrch (Will be either CREATE/UPDATE/DELETE).
* `status ` -  Current status of the ArcEdge deployment (verbose).
* `status_id ` -  Current status of the ArcEdge deployment (numerical).
* `active_ip_gateway ` -  Default gateway for the deployed ArcEdge.
* `active_private_ip ` -  ArcEdge private IP.
* `backup_ip ` -  Backup ArcEdge public IP (when enable_high_availability is set to true).
* `backup_private_ip ` -  Backup ArcEdge private IP (when enable_high_availability is set to true).
* `private_cidr ` -  CIDR block of the private subnet.
* `ingress_sg ` -  Security group created for the ArcEdge deployment.
* `hub_number ` -  Hub number of the overlay network.
* `coordinates_lat ` -  Latitude where ArcEdge is deployed.
* `coordinates_long ` -  Longitude where ArcEdge is deployed.
* `latest_available_image ` - Latest version of available image of Arcedge.
* `network_interfaces ` - Arcedge network interface details.

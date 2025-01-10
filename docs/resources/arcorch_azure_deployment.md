# <resource name> arcorch_azure_deployment

arcorch_azure_deployment resource will be used to create ArcEdge deployments in the desired Azure account. 

## Example Usage

```hcl
resource "arcorch_azure_deployment" "arcorch_azure" {
  name = "azure-arcedge"
  instance_key = "~/arcedge.pub"
  enable_accelerated_networking = false
  enable_high_availability = false 
  instance_type = "Standard_B2s"
  region = "westus"
  resource_group = "resource_group"
  vnet = "vnet"
  networks = [
    {
      subnetwork = "default"
  },{
      subnetwork = "subnet-1"
  }
  ]
  source_image {
    image_id = "image_id"
  }
  credential_name = "azure-cred"
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on GCP
* `credential_name ` - (Required) The name of the GCP credential where the ArcEdge will be deployed.
* `region ` - (Required) Region where ArcEdge will be deployed.
* `resource_group ` - (Required) Resource group where ArcEdge will be deployed.
* `instance_type ` - (Required) Instance size of the ArcEdge deployed.
* `instance_key ` - (Required) Path of the public key file which is needed to ssh into the deployed ArcEdge. 
* `vnet ` - (Required) Vnet where the ArcEdge will be deployed.
* `networks` - (Required) Network setup for deployed ArcEdges.
  * `subnetwork` - (Required) Subnet for deployed ArcEdges.
* `enable_accelerated_networking ` - (Required) Specifies if Accelerated Networking(SR-IOV) is enabled for the ArcEdge.
* `enable_high_availability ` - (Required) Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `source_image ` - (Required) Image of the ArcEdge.
  * `image_id ` - (Required) The ID of the image.
* `zone` - (Optional) Zone where active ArcEdge will be deployed.
* `backup_zone` - (Optional) Zone where standby ArcEdge will be deployed.
* `byoip ` - (Optional) A reserved IP which will be assigned to the deployed ArcEdge.
* `byoip_backup` - (Optional) A reserved IP which will be assigned to the deployed standby ArcEdge.

## Attribute Reference

* `credential_id` - a unique identifier of the credential for the deployed Arcedge
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
* `network_interfaces ` - Details of arcedge network interfaces.

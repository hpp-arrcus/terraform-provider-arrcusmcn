# <resource name> arcorch_gcp_deployment

arcorch_gcp_deployment resource will be used to create ArcEdge deployments in the desired GCP account. 

## Example Usage

```hcl
resource "arrcusmcn_gcp_deployment" "arrcusmcn_gcp" {
  name = "gcp-arcedge"
  instance_key = "~/instance.pub"
  region = "us-west1"
  zone = "us-west1-c"
  backup_zone = "us-west1-a"
  instance_type = "n2-standard-2"
  networks = [
    {
      network = "network1"
      subnetwork = "subnet1"
    },
    {
      network = "network2"
      subnetwork = "subnet2"
    }
  ]
  source_image  {
   image_id = "image_id"
  }
  enable_high_availability = false
  credential_name = "gcp-cred"
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on GCP
* `credential_name ` - (Required) The name of the GCP credential where the ArcEdge will be deployed.
* `region ` - (Required) Region where ArcEdge will be deployed.
* `zone ` - (Required) Zone where ArcEdge will be deployed.
* `instance_type ` - (Required) Instance size of the ArcEdge deployed.
* `instance_key ` - (Required) Path of the public key file which is needed to ssh into the deployed ArcEdge. 
* `networks` - (Required) Network setup for deployed ArcEdges.
  * `subnet_a` - (Required) Subnet for deployed active ArcEdge.
  * `subnet_b` - (Optional) Subnet for deployed standby ArcEdge if high availability is enabled.
* `source_image ` - (Required) Image of the ArcEdge.
  * `image_id ` - (Required) The ID of the image.
* `enable_high_availability ` - (Required) Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `byoip ` - (Optional) A reserved IP which will be assigned to the deployed ArcEdge.
* `network_tags ` - (Optional) GCP config tags need to be applied.
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
* `network_interfaces ` - Arcedge network interface details.

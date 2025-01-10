# <resource name> arrcusmcn_aws_deployment

arrcusmcn_aws_deployment resource will be used to create ArcEdge deployments in the desired AWS account. 

## Example Usage

```hcl

resource "arrcusmcn_aws_cred" "arcorch_aws" {
  name = "aws-hub"
  credential_name = "aws-cred"
  region = "us-east-1"
  zone = "us-east-1a"
  backup_zone = "us-east-1c"
  vpc_id = "vpc_id"
  instance_key = "arcedge"
  instance_type = "t2.medium"
  source_image  {
    image_id = "ami-image_id"
  }
  networks = [
    {
      subnet_a = "subnet_a"
      subnet_b = "subnet_b"
    },
  ]
  enable_high_availability = true
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on AWS
* `credential_name ` - (Required) The name of the AWS credential where the ArcEdge will be deployed.
* `region ` - (Required) Region where ArcEdge will be deployed.
* `instance_type ` - (Required) Instance size of the ArcEdge deployed.
* `instance_key ` - (Required) Instance key name which is needed to ssh into the deployed ArcEdge.
* `vpc_id ` - (Required) VPC ID where the ArcEdge will be deployed.
* `networks` - (Required) Network setup for deployed ArcEdges.
  * `subnet_a` - (Required) Subnet for deployed active ArcEdge.
  * `subnet_b` - (Optional) Subnet for deployed standby ArcEdge if high availability is enabled.
* `enable_high_availability ` - (Required) Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `source_image ` - (Required) Image of the ArcEdge.
  * `image_id ` - (Required) The id of the image.
* `byoip ` - (Optional) A reserved IP which will be assigned to the deployed ArcEdge.
* `byoip_backup` - (Optional) A reserved IP which will be assigned to the deployed standby ArcEdge.
* `zone` - (Optional) Zone where active ArcEdge will be deployed.
* `backup_zone` - (Optional) Zone where standby ArcEdge will be deployed..

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


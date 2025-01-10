# <resource name> arrcusmcn_aws_deployment

The arrcusmcn_aws_deployment data source provides details of a specific ArcEdge created using the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arrcusmcn_aws_deployment" "arcorch_aws" {
  name = "aws_hub"
}
```

## Argument Reference

* `name ` - (Required) A unique name for specific ArcEdge deployment running on AWS

## Attribute Reference

* `id` - A unique identifier for the resource
* `arc_orch_ip ` -  Public IP of the ArcEdge that has been deployed.
* `action ` -  Action initiated by the ArcOrch (Will be either CREATE/UPDATE/DELETE).
* `status ` -  Current status of the ArcEdge deployment (verbose).
* `status_id ` -  Current status of the ArcEdge deployment (numerical).
* `active_ip_gateway ` -  Default gateway for the deployed ArcEdge.
* `active_private_ip ` -  ArcEdge private IP.
* `backup_ip ` -  Backup ArcEdge public IP (when enable_high_availability is set to true).
* `backup_private_ip ` -  Backup ArcEdge private IP (when enable_high_availability is set to true).
* `byoip ` - The reserved IP is assigned to the deployed ArcEdge.
* `credential_id` - A unique identifier of the credential for the deployed Arcedge.
* `credential_name ` - The name of the AWS credential where the ArcEdge deployed.
* `enable_high_availability ` - Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `public_subnet ` -  Public subnet (should contain an attached Internet Gateway).
* `private_subnet ` - Private subnet when the ArcEdge is deployed as a spoke.
* `private_cidr ` -  CIDR block of the private subnet.
* `ingress_sg ` -  Security group created for the ArcEdge deployment.
* `instance_key ` - Instance key name which is needed to ssh into the deployed ArcEdge.
* `instance_type ` - Instance size of the ArcEdge deployed.
* `hub_number ` -  Hub number of the overlay network.
* `region ` - Region where ArcEdge deployed.
* `coordinates_lat ` -  Latitude where ArcEdge is deployed.
* `coordinates_long ` -  Longitude where ArcEdge is deployed.
* `source_image ` - Image of the ArcEdge.
* `latest_available_image ` - Latest version of available image of Arcedge.
* `network_interfaces ` - Arcedge network interface details.
* `vpc_id ` - VPC ID where the ArcEdge deployed.

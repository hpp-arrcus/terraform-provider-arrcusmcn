# <resource name> arcorch_oci_deployment

The arcorch_oci_deployment data source provides details of a specific ArcEdge created using the Arrcus ArcOrchestrator.

## Example Usage

```hcl
data "arcorch_oci_deployment" "arcorch_oci" {
  name = "oci_hub"
}
```

## Argument Reference

* `name ` - (Required) A unique name for specific ArcEdge deployment running on OCI

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
* `credential_name ` - The name of the OCI credential where the ArcEdge deployed.
* `enable_high_availability ` - Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `private_cidr ` -  CIDR block of the private subnet.
* `ingress_sg ` -  Security group created for the ArcEdge deployment.
* `hub_number ` -  Hub number of the overlay network.
* `region ` - Region where ArcEdge deployed.
* `coordinates_lat ` -  Latitude where ArcEdge is deployed.
* `coordinates_long ` -  Longitude where ArcEdge is deployed.
* `source_image ` - Image of the ArcEdge.
* `latest_available_image ` - Latest version of available image of Arcedge.
* `network_interfaces ` -  Arcedge network interface details.
* `availability_domain ` - Availability Domain where ArcEdge deployed.
* `compartment ` - Ocid of the compartment where ArcEdge deployed.
* `compute_shape ` - Instance size of the ArcEdge deployed.
* `compute_cpus ` - Number of CPUs of the ArcEdge using flexible shape deployed.
* `compute_memory_in_gbs ` - Memory size of the ArcEdge using flexible shape deployed.
* `instance_key ` - Path of the public key file which is needed to ssh into the deployed ArcEdge. 
* `public_subnet ` - Public subnet.
  * `vcn ` - Ocid of the vcn for the subnet.
  * `ocid ` - Ocid of the subent.
* `public_subnet ` - Public subnet when the Arcedge deployed as spoke.
  * `vcn ` - Ocid of the vcn for the subnet.
  * `ocid ` - Ocid of the subent.
* `image_compartment ` - Ocid of compartment where the arcedge image located.
* `enable_firewall ` - Provide reachability to ArcEdge for protocols and ports, such as ssh.Default is true.

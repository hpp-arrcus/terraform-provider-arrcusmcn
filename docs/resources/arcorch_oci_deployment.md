# <resource name> arcorch_oci_deployment

arcorch_oci_deployment resource will be used to create ArcEdge deployments in the desired OCI account. 

## Example Usage

```hcl
resource "arrcusmcn_oci_deployment" "arrcusmcn_oci" {
    instance_key = "~/instance.pub"
    region = "us-sanjose-1"
    availability_domain = "uNwK:US-SANJOSE-1-AD-1"
    compartment = "compartment-ocid"
    networks = [
      {
        vcn_ocid = "vcn1_ocid"
        subnet_ocid = "subnet1_ocid"
      },
      {
        vcn_ocid = "vcn2_ocid"
        subnet_ocid = "subnet2_ocid"
      },
    ]
    compute_shape = "VM.Standard.E3.Flex"
    compute_cpus = "2"
    compute_memory_in_gbs = "4"
    image_compartment = "image-compartment-ocid"
    credential_name = "oci-cred"
    name = "oci-arceddge"
    source_image {
        image_id = "image-ocid"
    }
    enable_high_availability = false
}
```

## Argument Reference

* `name ` - (Required) A unique name for an ArcEdge deployment running on OCI
* `credential_name ` - (Required) The name of the OCI credential where the ArcEdge will be deployed.
* `region ` - (Required) Region where ArcEdge will be deployed.
* `compartment ` - (Required) Ocid of the compartment where ArcEdge will be deployed.
* `availability_domain ` - (Required) Availability Domain where active ArcEdge will be deployed.
* `backup_availability_domain` - (Optional) Availability Domain where standby ArcEdge will be deployed.
* `compute_shape ` - (Required) Instance size of the ArcEdge deployed.
* `compute_cpus ` - (Optional) Number of CPUs of the ArcEdge deployed. This is required if compute shape is flexible.
* `compute_memory_in_gbs ` - (Optional) Memory size of the ArcEdge deployed. This is required if compute shape is flexible.
* `instance_key ` - (Required) Path of the public key file which is needed to ssh into the deployed ArcEdge. 
* `networks` - (Required) Network setup for deployed ArcEdges.
  * `vcn ` - (Required) Ocid of the vcn for the subnet.
  * `ocid ` - (Required) Ocid of the subent.
* `image_compartment ` - (Required) Ocid of compartment where the arcedge image located.
* `source_image ` - (Required) Image of the ArcEdge.
  * `image_id ` - (Required) The ocid of the image.
* `enable_firewall ` - (Optional) Provide reachability to ArcEdge for protocols and ports, such as ssh. Default is true.
* `enable_high_availability ` - (Required) Set to true if high availability is desired. A pair of ArcEdge in active/standby mode will be deployed.
* `byoip ` - (Optional) A reserved IP which will be assigned to the deployed ArcEdge.
* `byoip_backup` - (Optional) A reserved IP which will be assigned to the deployed standby ArcEdge.
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

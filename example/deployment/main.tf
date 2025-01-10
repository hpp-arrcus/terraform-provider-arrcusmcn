terraform {
  required_providers {
    arrcusmcn = {
      version = "1.0.0"
      source = "arrcus.com/arrcus/arrcusmcn"
    }
  }
}

provider "arrcusmcn" {
  username = ""
  password = ""
  serverip = ""
  port = ""
}

data "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
}

resource "arrcusmcn_aws_deployment" "arrcusmcn_aws" {
  name = ""
  credentials_id = data.arrcusmcn_aws_cred.aws_cred.id
  public_subnet = ""
  region = ""
  vpc_id = ""
  instance_key = ""
  instance_type = ""
  private_subnet = ""
  enable_high_availability = false
  enable_private_subnet = false
}

output "arcedge" {
  value = arrcusmcn_aws_deployment.arrcusmcn_aws
}
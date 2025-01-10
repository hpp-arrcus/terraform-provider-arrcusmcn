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

resource "arrcusmcn_aws_cred" "aws_cred" {
  name = "aws_cred"
  access_key = ""
  secret_key = ""
}
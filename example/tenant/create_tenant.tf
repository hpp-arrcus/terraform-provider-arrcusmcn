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

resource "arrcusmcn_tenant" "tenant" {
  name = ""
  organization = ""
  domain = ""
  defaultuser_name = ""
  defaultuser_username = ""
  defaultuser_password = ""
  defaultuser_email = ""
  defaultuser_roles = ["TenantAdmin", "TenantOperator"]
}

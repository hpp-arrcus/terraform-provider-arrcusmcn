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

/*
Run terraform apply directly will create a new user with given info.
If you want to update user without create it, please follow:
1. Remove `resource "arrcusmcn_user" "user"` and then run terraform apply to get id of the user currently using.
2. Add `resource "arrcusmcn_user" "user"` with corresponding info.
3. Run `terraform import arrcusmcn_user.user {id}` to import existing user.
4. Run `terraform apply`
*/

resource "arrcusmcn_user" "user" {
  name = ""
  username = ""
  password = ""
  email = ""
  roles = ["TenantOperator"]
}

output "user" {
  value = resource.arrcusmcn_user.user
}

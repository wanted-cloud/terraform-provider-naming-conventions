terraform {
  required_providers {
    conventions = {
      source = "registry.terraform.io/wanted-cloud/naming-conventions"
    }
  }
}

provider "conventions" {}

output "timestamp" {
  value = provider::conventions::rfc3339_parse()
}
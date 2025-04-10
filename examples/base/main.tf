terraform {
  required_providers {
    conventions = {
      source = "github.com/wanted-cloud/conventions"
    }
  }
}

provider "conventions" {}

output "timestamp" {
  value = provider::conventions::rfc3339_parse()
}
terraform {
  required_providers {
    conventions = {
      source = "registry.terraform.io/wanted-cloud/naming-conventions"
    }
  }
}

provider "conventions" {
  metadata = {}
}

data "conventions_generate" "test" {
  //provider = conventions
  example_attribute = "example_value"
}

output "generate_test_output" {
  description = "Output from the naming-conventions_generate data source"
  value = provider::conventions::generate("azurerm_resource_group", {
    a = "test"
  })
}

output "data_generate_test_output" {
  description = "Output from the conventions_generate data source"
  value       = data.conventions_generate.test
}
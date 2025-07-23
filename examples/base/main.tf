terraform {
  required_providers {
    conventions = {
      source = "registry.terraform.io/wanted-cloud/naming-conventions"
    }
  }
}

provider "conventions" {
  metadata = {
    location = "gwc"
  }
  patterns = {
      azurerm_resource_group = {
        pattern = "{{.name}}-rg-{{.location}}"
      }
    }
}

data "conventions_generate" "test" {
  type = "azurerm_resource_group"
  name = "test"
}

output "generate_test_output" {
  description = "Output from the naming-conventions_generate data source"
  value = provider::conventions::generate("azurerm_resource_group", {
    name = "test"
    location = "eastus"
  })
}

output "data_generate_test_output" {
  description = "Output from the conventions_generate data source"
  value       = data.conventions_generate.test
}
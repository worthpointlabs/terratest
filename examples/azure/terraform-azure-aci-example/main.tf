# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AN AZURE CONTAINER REGISTRY
# This is an example of how to deploy an Azure Container Registry
# ---------------------------------------------------------------------------------------------------------------------

# ------------------------------------------------------------------------------
# CONFIGURE OUR AZURE CONNECTION
# ------------------------------------------------------------------------------

provider "azurerm" {
  version = "=2.22.0"
  features {}
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "rg" {
  name     = var.resource_group_name
  location = var.location
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY AN AZURE CONTAINER INSTANCE
# ---------------------------------------------------------------------------------------------------------------------

data "azurerm_client_config" "current" {
}

resource "azurerm_container_group" "aci" {
  name                = var.aci_name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name

  ip_address_type     = "public"
  dns_name_label      = var.aci_name
  os_type             = "Linux"

  container {
    name   = "hello-world"
    image  = "microsoft/aci-helloworld:latest"
    cpu    = "0.5"
    memory = "1.5"

    ports {
      port     = 443
      protocol = "TCP"
    }
  }

  tags = {
    Environment = "Development"
  }
}

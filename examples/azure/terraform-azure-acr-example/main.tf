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
# DEPLOY AN AZURE CONTAINER REGISTRY
# ---------------------------------------------------------------------------------------------------------------------

data "azurerm_client_config" "current" {
}

resource "azurerm_container_registry" "acr" {
  name                = var.acr_name
  location            = azurerm_resource_group.rg.location
  resource_group_name = azurerm_resource_group.rg.name

  sku                 = var.sku
  admin_enabled       = true

  tags = {
    Environment = "Development"
  }
}

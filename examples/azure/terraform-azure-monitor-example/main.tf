provider "azurerm" {
  version = "~>2.20"

  features {
    key_vault {
      purge_soft_delete_on_destroy = true
    }
  }
}

# Configure the Microsoft Azure Active Directory Provider
provider "azuread" {
  version = "=0.7.0"
}

# ---------------------------------------------------------------------------------------------------------------------
# PIN TERRAFORM VERSION TO >= 0.12
# The examples have been upgraded to 0.12 syntax
# ---------------------------------------------------------------------------------------------------------------------

terraform {
  required_version = ">= 0.12"
}

resource "random_string" "short" {
  length  = 3
  lower   = true
  upper   = false
  number  = false
  special = false
}

resource "random_string" "long" {
  length  = 6
  lower   = true
  upper   = false
  number  = false
  special = false
}

# ---------------------------------------------------------------------------------------------------------------------
# DEPLOY A RESOURCE GROUP
# See test/terraform_azure_example_test.go for how to write automated tests for this code.
# ---------------------------------------------------------------------------------------------------------------------

resource "azurerm_resource_group" "rg" {
  name     = format("%s-%s-%s", "terratest", random_string.short.result, "monitor")
  location = "East US"
}

data "azurerm_client_config" "current" {}

resource "azurerm_storage_account" "storage" {
  name                     = format("%s%s", random_string.long.result, "storage")
  resource_group_name      = azurerm_resource_group.rg.name
  location                 = azurerm_resource_group.rg.location
  account_tier             = "Standard"
  account_replication_type = "GRS"

  tags = {
    environment = "staging"
  }
}

resource "azurerm_key_vault" "keyVault" {
  name                        = format("%s-%s", random_string.short.result, "vault")
  location                    = azurerm_resource_group.rg.location
  resource_group_name         = azurerm_resource_group.rg.name
  enabled_for_disk_encryption = true
  tenant_id                   = data.azurerm_client_config.current.tenant_id
  soft_delete_enabled         = true
  # soft_delete_retention_days  = 7
  purge_protection_enabled = false

  sku_name = "standard"

  access_policy {
    tenant_id = data.azurerm_client_config.current.tenant_id
    object_id = data.azurerm_client_config.current.object_id

    key_permissions = [
      "create",
      "get",
      "list",
      "delete",
    ]

    secret_permissions = [
      "set",
      "get",
      "list",
      "delete",
    ]

    certificate_permissions = [
      "create",
      "delete",
      "deleteissuers",
      "get",
      "getissuers",
      "import",
      "list",
      "listissuers",
      "managecontacts",
      "manageissuers",
      "setissuers",
      "update",
    ]
  }

  network_acls {
    default_action = "Deny"
    bypass         = "AzureServices"
  }

  tags = {
    environment = "Testing"
  }
}

# https://www.terraform.io/docs/providers/azurerm/r/monitor_diagnostic_setting.html
resource "azurerm_monitor_diagnostic_setting" "diagnosticSetting" {
  name               = var.diagnosticSettingName
  target_resource_id = azurerm_key_vault.keyVault.id
  storage_account_id = azurerm_storage_account.storage.id

  log {
    category = "AuditEvent"
    enabled  = false

    retention_policy {
      enabled = false
    }
  }

  metric {
    category = "AllMetrics"

    retention_policy {
      enabled = false
    }
  }
}

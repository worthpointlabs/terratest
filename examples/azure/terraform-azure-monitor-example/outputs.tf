# output "vm_name" {
#   value = azurerm_virtual_machine.main.name
# }

output "resource_group_name" {
  value = azurerm_resource_group.rg.name
}

output "diagnostic_setting_name" {
  value = azurerm_monitor_diagnostic_setting.diagnosticSetting.name
}

output "diagnostic_setting_id" {
  value = azurerm_monitor_diagnostic_setting.diagnosticSetting.id
}

output "keyvault_id" {
  value = azurerm_key_vault.keyVault.id
}

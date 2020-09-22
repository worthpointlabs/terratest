# output "vm_name" {
#   value = azurerm_virtual_machine.main.name
# }

output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "diagnostic_setting_name" {
  value = azurerm_monitor_diagnostic_setting.main.name
}

output "diagnostic_setting_id" {
  value = azurerm_monitor_diagnostic_setting.main.id
}

output "keyvault_id" {
  value = azurerm_key_vault.main.id
}

# output "vm_name" {
#   value = azurerm_virtual_machine.main.name
# }

output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "diagnostic_setting_name" {
  value = azurerm_monitor_diagnostic_setting.main.name
}

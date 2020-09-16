# output "vm_name" {
#   value = azurerm_virtual_machine.main.name
# }

output "resource_group_name" {
  value = azurerm_resource_group.main.name
}

output "lb_name" {
  value = azurerm_lb.main.name
}

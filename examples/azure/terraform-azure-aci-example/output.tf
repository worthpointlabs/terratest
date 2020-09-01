output "ip_address" {
  value = azurerm_container_group.aci.ip_address
}

output "fqdn" {
  value = azurerm_container_group.aci.fqdn
}

output "subscription_id" {
  value = data.azurerm_client_config.current.subscription_id
}

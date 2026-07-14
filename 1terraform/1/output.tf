output "container_app_url" {
  value = "https://${azurerm_container_app.webapp.latest_revision_fqdn}"
}

output "database_fqdn" {
  value = azurerm_postgresql_flexible_server.main.fqdn
}

output "container_registry_login" {
  value = azurerm_container_registry.main.login_server
}

output "resource_group_name" {
  value = azurerm_resource_group.main.name
}
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 4.0"
    }
    random = {
      source  = "hashicorp/random"
      version = "~> 3.5"
    }
  }
}

provider "azurerm" {
  features {}
}

# Random suffix for globally unique names
resource "random_id" "suffix" {
  byte_length = 4
}

# Resource Group
resource "azurerm_resource_group" "main" {
  name     = "rg-golanag-dev-${random_id.suffix.hex}"
  location = "West Europe"
}

# Container Registry
resource "azurerm_container_registry" "main" {
  name                = "crgolanag${random_id.suffix.hex}"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location
  sku                 = "Basic"
  admin_enabled       = true
}

# PostgreSQL Flexible Server
resource "random_password" "db_password" {
  length  = 16
  special = false
}

resource "azurerm_postgresql_flexible_server" "main" {
  name                   = "pg-golanag-${random_id.suffix.hex}"
  resource_group_name    = azurerm_resource_group.main.name
  location               = azurerm_resource_group.main.location
  version                = "15"
  administrator_login    = "pgadmin"
  administrator_password = random_password.db_password.result
  sku_name               = "B_Standard_B1ms"
  storage_mb             = 32768
  
  # Dev environment: allow public access (restrict in prod)
  public_network_access_enabled = true
}

resource "azurerm_postgresql_flexible_server_database" "main" {
  name      = "golanag"
  server_id = azurerm_postgresql_flexible_server.main.id
}

# Log Analytics
resource "azurerm_log_analytics_workspace" "main" {
  name                = "log-golanag-${random_id.suffix.hex}"
  resource_group_name = azurerm_resource_group.main.name
  location            = azurerm_resource_group.main.location
  sku                 = "PerGB2018"
  retention_in_days   = 30
}

# Container Apps Environment
resource "azurerm_container_app_environment" "main" {
  name                       = "env-golanag-${random_id.suffix.hex}"
  resource_group_name        = azurerm_resource_group.main.name
  location                   = azurerm_resource_group.main.location
  log_analytics_workspace_id = azurerm_log_analytics_workspace.main.id
}

# Container App
resource "azurerm_container_app" "webapp" {
  name                         = "webapp-golanag"
  container_app_environment_id = azurerm_container_app_environment.main.id
  resource_group_name          = azurerm_resource_group.main.name
  revision_mode                = "Single"

  template {
    container {
      name   = "webapp"
      image  = "${azurerm_container_registry.main.login_server}/golanag:latest"
      cpu    = 0.25
      memory = "0.5Gi"
      
      env {
        name  = "DB_HOST"
        value = azurerm_postgresql_flexible_server.main.fqdn
      }
      env {
        name  = "DB_PORT"
        value = "5432"
      }
      env {
        name  = "DB_USER"
        value = "admin"
      }
      env {
        name  = "DB_NAME"
        value = "golanag"
      }
      env {
        name        = "DB_PASSWORD"
        secret_name = "db-password"
      }
    }
  }

  secret {
    name  = "db-password"
    value = random_password.db_password.result
  }

  ingress {
    external_enabled = true
    target_port      = 8080
    transport        = "http"

    traffic_weight {
      latest_revision = true
      percentage      = 100
    }
  }
}
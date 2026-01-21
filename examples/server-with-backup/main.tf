# examples/server-with-backup/main.tf
# Server with automatic backup

terraform {
  required_providers {
    vdsina = {
      source  = "scinfra-pro/vdsina"
      version = "~> 0.1.0"
    }
  }
}

provider "vdsina" {
  api_token = var.vdsina_api_token
  base_url  = var.vdsina_base_url
}

# Create server
resource "vdsina_server" "production" {
  datacenter  = var.datacenter_id
  server_plan = var.server_plan_id
  template    = var.template_id
  name        = "production-server"
}

# Create server backup
resource "vdsina_backup" "weekly" {
  server_id   = vdsina_server.production.id
  name        = "weekly-backup"
  autoprolong = true
}

output "server_ip" {
  description = "Server IP"
  value       = vdsina_server.production.ip
}

output "backup_id" {
  description = "Backup ID"
  value       = vdsina_backup.weekly.id
}

output "backup_status" {
  description = "Backup status"
  value       = vdsina_backup.weekly.status
}

# Test configuration for terraform-provider-vdsina

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

# ============================================
# Data Sources - read-only, safe to test
# ============================================

# 1. Datacenters
data "vdsina_datacenters" "all" {}

# 2. Server groups
data "vdsina_server_groups" "all" {}

# 3. Server plans
data "vdsina_server_plans" "standard" {
  group_id = 1
}

# 4. Templates
data "vdsina_templates" "all" {}

# 5. Account info
data "vdsina_account" "me" {}

# 6. Account balance
data "vdsina_account_balance" "current" {}

# 7. Account limits
data "vdsina_account_limits" "limits" {}

# 8. Existing servers
data "vdsina_servers" "all" {}

# 9. SSH keys
data "vdsina_ssh_keys" "all" {}

# 10. IPs
data "vdsina_ips" "all" {}

# ============================================
# Outputs
# ============================================

output "datacenters" {
  value = data.vdsina_datacenters.all.datacenters
}

output "server_groups" {
  value = data.vdsina_server_groups.all.groups
}

output "account_name" {
  value = data.vdsina_account.me.name
}

output "balance" {
  value = data.vdsina_account_balance.current.total
}

output "servers_count" {
  value = length(data.vdsina_servers.all.servers)
}

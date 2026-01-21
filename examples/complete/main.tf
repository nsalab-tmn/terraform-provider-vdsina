# examples/complete/main.tf
# Complete example of terraform-provider-vdsina usage

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
# Data Sources
# ============================================

# 1. Datacenters
data "vdsina_datacenters" "all" {}

# 2. Server groups
data "vdsina_server_groups" "all" {}

# 3. Server plans (for Standard servers group)
data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

# 4. OS Templates
data "vdsina_templates" "all" {}

# 5. List of servers
data "vdsina_servers" "all" {}

# 6. SSH keys
data "vdsina_ssh_keys" "all" {}

# 7. Account information
data "vdsina_account" "me" {}

# 8. Account balance
data "vdsina_account_balance" "current" {}

# 9. Account limits
data "vdsina_account_limits" "limits" {}

# 10. List of ISO images
data "vdsina_iso_list" "all" {}

# 11. List of backups
data "vdsina_backups" "all" {}

# 12. DNS zones
data "vdsina_dns_zones" "all" {}

# 13. IP addresses
data "vdsina_ips" "all" {}

# ============================================
# Resources
# ============================================

# 1. Existing server (will import)
resource "vdsina_server" "test" {
  datacenter  = 4   # Amsterdam
  server_plan = 1   # 2 RAM / 1 CPU / 40 NVMe
  template    = 23  # Ubuntu 24.04
  ssh_key     = vdsina_ssh_key.test.id
  host        = "test.terraform.local"
  name        = "terraform-test-server"
}

# 2. SSH Key (existing, will import)
resource "vdsina_ssh_key" "test" {
  name = "terraform-key"
  data = var.ssh_public_key
}

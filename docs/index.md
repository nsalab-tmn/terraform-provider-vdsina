---
page_title: "VDSina Provider"
subcategory: ""
description: |-
  Terraform provider for managing VDSina cloud resources.
---

# VDSina Provider

The VDSina provider is used to manage resources in the [VDSina](https://vdsina.com) cloud platform. It supports creating and managing VPS servers, SSH keys, DNS zones, ISO images, backups, and more.

## Example Usage

```hcl
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
}

# Create a server
resource "vdsina_server" "web" {
  datacenter  = 4   # Amsterdam
  server_plan = 1   # 2 RAM / 1 CPU / 40 NVMe
  template    = 23  # Ubuntu 24.04
  name        = "web-server"
}
```

## Authentication

The provider requires an API token for authentication. You can obtain a token from your VDSina account settings.

### Environment Variables

You can provide credentials via environment variables:

```bash
export VDSINA_API_TOKEN="your-api-token"
```

### Provider Configuration

```hcl
provider "vdsina" {
  api_token = var.vdsina_api_token
  base_url  = "https://userapi.vdsina.ru/v1"  # optional, this is the default
}
```

## Schema

### Required

- `api_token` (String, Sensitive) VDSina API token. Can also be set via `VDSINA_API_TOKEN` environment variable.

### Optional

- `base_url` (String) VDSina API base URL. Defaults to `https://userapi.vdsina.ru/v1`.

## Resources

- `vdsina_server` - VPS server management
- `vdsina_ssh_key` - SSH key management
- `vdsina_iso` - ISO image management
- `vdsina_backup` - Server backup management
- `vdsina_dns` - DNS zone management
- `vdsina_dns_record` - DNS record management
- `vdsina_server_reboot` - Server reboot action
- `vdsina_server_reinstall` - OS reinstall action
- `vdsina_server_password` - Password reset action
- `vdsina_server_plan_change` - Plan upgrade action
- `vdsina_server_start_prolong` - Server prolong action
- `vdsina_server_iso` - ISO attach/detach action
- `vdsina_backup_restore` - Backup restore action

## Data Sources

- `vdsina_datacenters` - List of datacenters
- `vdsina_server_groups` - Server plan groups
- `vdsina_server_plans` - Server plans (tariffs)
- `vdsina_templates` - OS templates
- `vdsina_servers` - List of servers
- `vdsina_ssh_keys` - List of SSH keys
- `vdsina_account` - Account information
- `vdsina_account_balance` - Account balance
- `vdsina_account_limits` - Account limits
- `vdsina_iso_list` - List of ISO images
- `vdsina_backups` - List of backups
- `vdsina_dns_zones` - List of DNS zones
- `vdsina_ips` - List of IP addresses

# Terraform Provider for VDSina

[![Release](https://img.shields.io/github/v/release/nsalab-tmn/terraform-provider-vdsina)](https://github.com/nsalab-tmn/terraform-provider-vdsina/releases)
[![Tests](https://github.com/nsalab-tmn/terraform-provider-vdsina/actions/workflows/test.yml/badge.svg)](https://github.com/nsalab-tmn/terraform-provider-vdsina/actions/workflows/test.yml)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-brightgreen.svg)](https://opensource.org/licenses/MPL-2.0)

Terraform provider for managing resources in [VDSina](https://vdsina.com) cloud platform.

This is a maintained fork of [scinfra-pro/terraform-provider-vdsina](https://github.com/scinfra-pro/terraform-provider-vdsina), published as `nsalab-tmn/vdsina`. It waits for new servers to become active with a public IP, waits for deleted servers to disappear, and no longer drops resources from state on transient API errors — see [CHANGELOG.md](CHANGELOG.md).

## Features

- **VPS Servers** - Create, update, delete servers with custom configurations
- **SSH Keys** - Manage SSH keys for server access
- **ISO Images** - Upload and attach custom ISO images
- **Backups** - Create and restore server backups
- **DNS Zones** - Manage DNS zones and records
- **Server Actions** - Reboot, reinstall, password reset, plan changes

## Installation

### Terraform Registry (Recommended)

```hcl
terraform {
  required_providers {
    vdsina = {
      source  = "nsalab-tmn/vdsina"
      version = "~> 0.3.0"
    }
  }
}

provider "vdsina" {
  api_token = var.vdsina_api_token
}
```

### Manual Installation

Download the appropriate binary from the [releases page](https://github.com/nsalab-tmn/terraform-provider-vdsina/releases) and place it in:

```
~/.terraform.d/plugins/registry.terraform.io/nsalab-tmn/vdsina/0.3.0/<os>_<arch>/
```

## Configuration

### Authentication

The provider requires an API token. Get yours from [VDSina account settings](https://my.vdsina.ru/account/api).

```hcl
provider "vdsina" {
  api_token = var.vdsina_api_token  # or use VDSINA_API_TOKEN env var
}
```

### Environment Variables

```bash
export VDSINA_API_TOKEN="your-api-token"
```

## Quick Start

```hcl
# Get available datacenters
data "vdsina_datacenters" "all" {}

# Get server plans for Standard group (id=1)
data "vdsina_server_plans" "standard" {
  group_id = 1
}

# Get OS templates
data "vdsina_templates" "all" {}

# Create SSH key
resource "vdsina_ssh_key" "main" {
  name = "my-key"
  data = file("~/.ssh/id_rsa.pub")
}

# Create a server
resource "vdsina_server" "web" {
  datacenter  = 4   # Amsterdam
  server_plan = 1   # 2 RAM / 1 CPU / 40 NVMe
  template    = 23  # Ubuntu 24.04
  ssh_key     = vdsina_ssh_key.main.id
  name        = "web-server"
}

output "server_ip" {
  value = vdsina_server.web.ip
}
```

## Resources

| Resource | Description |
|----------|-------------|
| `vdsina_server` | VPS server management |
| `vdsina_ssh_key` | SSH key management |
| `vdsina_iso` | ISO image management |
| `vdsina_backup` | Server backup management |
| `vdsina_dns` | DNS zone management |
| `vdsina_dns_record` | DNS record management |
| `vdsina_server_reboot` | Server reboot action |
| `vdsina_server_reinstall` | OS reinstall action |
| `vdsina_server_password` | Password management |
| `vdsina_server_plan_change` | Plan upgrade action |
| `vdsina_server_start_prolong` | Server prolong action |
| `vdsina_server_iso` | ISO attach/detach action |
| `vdsina_backup_restore` | Backup restore action |

## Data Sources

| Data Source | Description |
|-------------|-------------|
| `vdsina_datacenters` | List of datacenters |
| `vdsina_server_groups` | Server plan groups |
| `vdsina_server_plans` | Server plans by group |
| `vdsina_templates` | OS templates |
| `vdsina_servers` | List of servers |
| `vdsina_ssh_keys` | List of SSH keys |
| `vdsina_account` | Account information |
| `vdsina_account_balance` | Account balance |
| `vdsina_account_limits` | Account limits |
| `vdsina_iso_list` | List of ISO images |
| `vdsina_backups` | List of backups |
| `vdsina_dns_zones` | List of DNS zones |
| `vdsina_ips` | List of IP addresses |

## Development

### Requirements

- [Go](https://golang.org/doc/install) 1.21+
- [Terraform](https://www.terraform.io/downloads.html) 1.0+

### Building

```bash
make build
```

### Testing

```bash
make test
```

### Local Installation

```bash
make install
```

## License

This project is licensed under the Mozilla Public License 2.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

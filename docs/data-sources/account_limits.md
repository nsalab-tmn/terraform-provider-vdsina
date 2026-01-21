---
page_title: "vdsina_account_limits Data Source - terraform-provider-vdsina"
subcategory: "Account"
description: |-
  Retrieves VDSina account limits by service type.
---

# vdsina_account_limits (Data Source)

Use this data source to get information about your VDSina account limits for various service types.

Each limit includes:
- `max` - Maximum allowed quantity
- `now` - Currently used quantity
- `child_max` - Maximum per child account (for IP/GPU)

## Example Usage

```hcl
data "vdsina_account_limits" "limits" {}

output "server_limits" {
  value = {
    max = data.vdsina_account_limits.limits.server[0].max
    now = data.vdsina_account_limits.limits.server[0].now
  }
}
```

### Check remaining capacity

```hcl
data "vdsina_account_limits" "limits" {}

locals {
  servers_remaining = (
    data.vdsina_account_limits.limits.server[0].max -
    data.vdsina_account_limits.limits.server[0].now
  )
}

output "can_create_more_servers" {
  value = local.servers_remaining > 0
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

Each limit is a list with one element containing:

| Attribute | Type | Description |
|-----------|------|-------------|
| `max` | Number | Maximum allowed. |
| `now` | Number | Currently used. |
| `child_max` | Number | Maximum per child account. |

### Available Limits

| Attribute | Description |
|-----------|-------------|
| `server` | VPS servers limit. |
| `server_ip4` | Additional IPv4 addresses limit. |
| `server_ip6` | Additional IPv6 addresses limit. |
| `iso` | ISO images limit. |
| `backup` | Backups limit. |
| `ssl` | SSL certificates limit (may be empty if not available). |
| `domain` | Domains limit (may be empty if not available). |
| `dns` | DNS zones limit. |
| `extdisk_hdd` | External HDD disks limit. |
| `extdisk_nvme` | External NVMe disks limit. |
| `reserve_ip` | Reserved IP addresses limit. |
| `gpu` | GPU limit. |

~> **Note:** `ssl` and `domain` limits may return an empty list if the service is not available for your account.

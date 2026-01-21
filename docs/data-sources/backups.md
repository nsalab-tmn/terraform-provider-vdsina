---
page_title: "vdsina_backups Data Source - terraform-provider-vdsina"
subcategory: "Backup"
description: |-
  Retrieves list of backups in your VDSina account.
---

# vdsina_backups (Data Source)

Use this data source to get information about server backups in your VDSina account.

## Example Usage

### Get all backups

```hcl
data "vdsina_backups" "all" {}

output "backups" {
  value = data.vdsina_backups.all.backups
}
```

### Find active backups

```hcl
data "vdsina_backups" "all" {}

locals {
  active_backups = [
    for b in data.vdsina_backups.all.backups : b
    if b.status == "active"
  ]
}

output "active_backup_ids" {
  value = local.active_backups[*].id
}
```

### Find backups for specific server

```hcl
data "vdsina_backups" "all" {}

locals {
  server_backups = [
    for b in data.vdsina_backups.all.backups : b
    if b.server_id == 12345
  ]
}

output "server_backup_count" {
  value = length(local.server_backups)
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `backups` - List of backups. Each backup has the following attributes:

### Backup Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Backup ID. |
| `name` | String | Backup name. |
| `full_name` | String | Full name with size (e.g., "Backup #12345 – server (40 Gb)"). |
| `created` | String | Creation date and time. |
| `updated` | String | Last update date and time. |
| `end` | String | Expiration date. |
| `status` | String | Status: `new`, `active`, `block`, `notpaid`, `deleted`. |
| `status_text` | String | Status description (e.g., "Creation (9%) Processing"). |
| `datacenter_id` | Number | Datacenter ID where backup is stored. |
| `datacenter_name` | String | Datacenter name. |
| `datacenter_country` | String | Datacenter country code. |
| `server_id` | Number | Parent server ID (0 if not available). |
| `server_name` | String | Parent server name. |
| `can_update` | Boolean | Whether the backup can be renamed. |
| `can_prolong` | Boolean | Whether the backup can be prolonged. |
| `can_delete` | Boolean | Whether the backup can be deleted. |

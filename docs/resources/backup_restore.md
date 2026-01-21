# vdsina_backup_restore (Resource)

Restores a VDSina backup to a server.

~> **Warning:** This action will overwrite all data on the target server!

## Example Usage

### Restore Backup to Server

```hcl
data "vdsina_backups" "all" {}

resource "vdsina_backup_restore" "restore" {
  backup_id = data.vdsina_backups.all.backups[0].id
  server_id = vdsina_server.web.id
}
```

### Restore with Trigger

```hcl
resource "vdsina_backup_restore" "periodic" {
  backup_id = vdsina_backup.production.id
  server_id = vdsina_server.staging.id
  
  triggers = {
    restore_time = timestamp()
  }
}
```

## Schema

### Required

* `backup_id` - (Number, ForceNew) Backup ID to restore from.
* `server_id` - (Number, ForceNew) Server ID to restore to. Must be in the same datacenter as the backup.

### Optional

* `triggers` - (Map of String) Map of values that, when changed, will trigger a restore.

### Read-Only

* `id` - (String) Composite ID in format `backup_id:server_id`.
* `last_restored` - (String) Timestamp of the last restore action.

## Timeouts

* `create` - (Default 30 minutes) Used for backup restoration.

~> **Note:** The server must be in the same datacenter as the backup.

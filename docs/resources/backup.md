# vdsina_backup (Resource)

Manages a VDSina server backup. Creates a backup of the specified server.

~> **Note:** Backup creation is asynchronous. The provider will poll until the backup completes (up to 30 minutes).

## Example Usage

### Basic Backup

```hcl
resource "vdsina_backup" "weekly" {
  server_id = vdsina_server.web.id
}
```

### Named Backup with Auto-prolong

```hcl
resource "vdsina_server" "web" {
  datacenter  = 4
  server_plan = 1
  template    = 23
  name        = "web-server"
}

resource "vdsina_backup" "production" {
  server_id   = vdsina_server.web.id
  name        = "production-backup"
  autoprolong = true
}
```

### Create Server from Backup

```hcl
data "vdsina_backups" "all" {}

resource "vdsina_server" "restored" {
  datacenter  = 4
  server_plan = 1
  backup      = data.vdsina_backups.all.backups[0].id
  name        = "restored-from-backup"
}
```

## Schema

### Required

* `server_id` - (Number, ForceNew) ID of the server to backup.

### Optional

* `name` - (String) Backup name. If not specified, server name will be used.
* `autoprolong` - (Boolean) Enable automatic prolongation of backup service. Default: `true`.

### Read-Only

* `id` - (String) The ID of the backup.
* `full_name` - (String) Full backup service name.
* `status` - (String) Backup status: `new`, `active`, `block`, `notpaid`, `deleted`.
* `status_text` - (String) Status description.
* `datacenter_id` - (Number) Datacenter ID where backup is stored.
* `datacenter_name` - (String) Datacenter name.
* `created` - (String) Creation timestamp.
* `updated` - (String) Last update timestamp.
* `end` - (String) Expiration date.

## Import

Backups can be imported using the backup ID:

```bash
terraform import vdsina_backup.example 614084
```

## Timeouts

* `create` - (Default 30 minutes) Used for backup creation.
* `update` - (Default 5 minutes) Used for backup updates.
* `delete` - (Default 5 minutes) Used for backup deletion.

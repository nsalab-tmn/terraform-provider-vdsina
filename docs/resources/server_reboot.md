# vdsina_server_reboot (Resource)

Reboots a VDSina server. The reboot is triggered on resource creation and when triggers change.

## Example Usage

### Basic Reboot

```hcl
resource "vdsina_server_reboot" "web" {
  server_id = vdsina_server.web.id
}
```

### Soft Reboot with Trigger

```hcl
resource "vdsina_server_reboot" "app" {
  server_id = vdsina_server.app.id
  type      = "soft"
  
  triggers = {
    reboot_time = timestamp()
  }
}
```

### Hard Reboot

```hcl
resource "vdsina_server_reboot" "stuck" {
  server_id = vdsina_server.stuck.id
  type      = "hard"
}
```

## Schema

### Required

* `server_id` - (Number, ForceNew) Server ID to reboot.

### Optional

* `type` - (String) Reboot type: `soft` (graceful) or `hard` (force). Default: `soft`.
* `triggers` - (Map of String) Map of arbitrary keys and values that, when changed, will trigger a reboot.

### Read-Only

* `id` - (String) The server ID.
* `last_reboot` - (String) Timestamp of the last reboot.

~> **Warning:** Hard reboot may result in data loss. Use only when server is frozen.

# vdsina_server_prolong (Resource)

Prolongs and starts a VDSina server that was stopped due to payment issues.

## Example Usage

### Prolong Stopped Server

```hcl
resource "vdsina_server_prolong" "restart" {
  server_id = vdsina_server.web.id
}
```

## Schema

### Required

* `server_id` - (Number, ForceNew) Server ID to prolong and start.

### Read-Only

* `id` - (String) The server ID.
* `prolonged_at` - (String) Timestamp of the prolong action.

~> **Note:** This action requires sufficient account balance.

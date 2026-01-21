# vdsina_server_plan_change (Resource)

Changes the tariff plan of a VDSina server (upgrade only).

~> **Note:** Only upgrades are supported. Downgrades require creating a new server.

## Example Usage

### Upgrade Server Plan

```hcl
data "vdsina_server_plans" "standard" {
  group_id = 2
}

resource "vdsina_server_plan_change" "upgrade" {
  server_id   = vdsina_server.web.id
  server_plan = data.vdsina_server_plans.standard.plans[1].id  # Larger plan
}
```

## Schema

### Required

* `server_id` - (Number, ForceNew) Server ID to upgrade.
* `server_plan` - (Number) New server plan ID (must be an upgrade).

### Read-Only

* `id` - (String) The server ID.
* `plan_changed_at` - (String) Timestamp of the plan change.

~> **Warning:** Plan change may require server reboot and cause brief downtime.

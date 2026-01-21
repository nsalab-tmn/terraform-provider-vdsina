# vdsina_server_reinstall (Resource)

Reinstalls the operating system on a VDSina server.

~> **Warning:** All data on the server will be destroyed during reinstall!

## Example Usage

### Reinstall with Same Template

```hcl
resource "vdsina_server_reinstall" "upgrade" {
  server_id = vdsina_server.web.id
  template  = 23  # Ubuntu 24.04
}
```

### Reinstall with SSH Key

```hcl
resource "vdsina_server_reinstall" "fresh" {
  server_id = vdsina_server.app.id
  template  = 23
  ssh_key   = vdsina_ssh_key.deploy.id
  host      = "app.example.com"
}
```

## Schema

### Required

* `server_id` - (Number, ForceNew) Server ID to reinstall.
* `template` - (Number) OS template ID to install.

### Optional

* `ssh_key` - (Number) SSH key ID to add during reinstall.
* `host` - (String) New hostname for the server.
* `triggers` - (Map of String) Map of values that, when changed, will trigger reinstall.

### Read-Only

* `id` - (String) The server ID.
* `last_reinstall` - (String) Timestamp of the last reinstall.
* `password` - (String, Sensitive) New root password after reinstall.

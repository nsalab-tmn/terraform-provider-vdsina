# vdsina_server_iso (Resource)

Attaches or detaches an ISO image to/from a VDSina server.

## Example Usage

### Attach ISO to Server

```hcl
resource "vdsina_iso" "rescue" {
  url = "https://boot.netboot.xyz/ipxe/netboot.xyz.iso"
}

resource "vdsina_server_iso" "attach" {
  server_id = vdsina_server.web.id
  iso_id    = vdsina_iso.rescue.id
}
```

### Detach ISO

To detach an ISO, simply remove the `vdsina_server_iso` resource from your configuration and run `terraform apply`.

## Schema

### Required

* `server_id` - (Number, ForceNew) Server ID to attach ISO to.
* `iso_id` - (Number) ISO image ID to attach.

### Read-Only

* `id` - (String) The server ID.

~> **Note:** After attaching ISO, reboot the server to boot from it.

## Lifecycle

- **Create:** Attaches the ISO to the server.
- **Delete:** Detaches the ISO from the server.

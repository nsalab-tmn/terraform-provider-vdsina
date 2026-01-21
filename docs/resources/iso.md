# vdsina_iso (Resource)

Manages a VDSina ISO image. ISO is downloaded from a URL and stored in VDSina for server installation.

~> **Note:** ISO creation is asynchronous. The provider will poll until the download completes (up to 30 minutes).

## Example Usage

### Basic ISO

```hcl
resource "vdsina_iso" "netboot" {
  url = "https://boot.netboot.xyz/ipxe/netboot.xyz.iso"
}
```

### ISO with Server

```hcl
resource "vdsina_iso" "custom" {
  url = "https://example.com/my-custom.iso"
}

resource "vdsina_server" "app" {
  datacenter  = 4
  server_plan = 1
  iso         = vdsina_iso.custom.id
  name        = "custom-os-server"
}
```

## Schema

### Required

* `url` - (String, ForceNew) URL to download ISO from (http, https, ftp, ftps). Maximum file size: 10GB.

### Read-Only

* `id` - (String) The ID of the ISO.
* `name` - (String) ISO file name.
* `full_name` - (String) Full ISO service name.
* `status` - (String) ISO status: `new`, `active`, `block`, `notpaid`, `deleted`.
* `status_text` - (String) Status description.
* `created` - (String) Creation timestamp.
* `updated` - (String) Last update timestamp.
* `end` - (String) Expiration date.
* `file_size` - (String) File size (e.g., "841 Mb").
* `file_md5` - (String) File MD5 checksum.
* `attached` - (Boolean) Whether ISO is attached to any server.

## Import

ISO images can be imported using the ISO ID:

```bash
terraform import vdsina_iso.example 614417
```

## Timeouts

* `create` - (Default 30 minutes) Used for ISO download and creation.
* `delete` - (Default 5 minutes) Used for ISO deletion.

~> **Warning:** ISO cannot be deleted while attached to a server. Detach it first.

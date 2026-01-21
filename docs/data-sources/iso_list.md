---
page_title: "vdsina_iso_list Data Source - terraform-provider-vdsina"
subcategory: "ISO"
description: |-
  Retrieves list of ISO images in your VDSina account.
---

# vdsina_iso_list (Data Source)

Use this data source to get information about ISO images uploaded to your VDSina account.

ISO images can be attached to servers for OS installation or recovery purposes.

## Example Usage

### Get all ISO images

```hcl
data "vdsina_iso_list" "all" {}

output "isos" {
  value = data.vdsina_iso_list.all.isos
}
```

### Find attached ISO images

```hcl
data "vdsina_iso_list" "all" {}

locals {
  attached_isos = [
    for iso in data.vdsina_iso_list.all.isos : iso
    if iso.attached
  ]
}

output "attached_isos" {
  value = local.attached_isos[*].name
}
```

### Find ISO by name

```hcl
data "vdsina_iso_list" "all" {}

locals {
  my_iso = [
    for iso in data.vdsina_iso_list.all.isos : iso
    if can(regex("ubuntu", lower(iso.name)))
  ]
}

output "ubuntu_iso_id" {
  value = length(local.my_iso) > 0 ? local.my_iso[0].id : null
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `isos` - List of ISO images. Each ISO has the following attributes:

### ISO Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | ISO service ID. |
| `name` | String | ISO file name (e.g., "ubuntu-22.04-server.iso"). |
| `full_name` | String | Full name with ID (e.g., "ISO #12345 – ubuntu-22.04-server.iso"). |
| `created` | String | Creation date and time. |
| `updated` | String | Last update date and time. |
| `end` | String | Expiration date. |
| `status` | String | Status: `new`, `active`, `block`, `notpaid`, `deleted`. |
| `status_text` | String | Status description. |
| `file_size` | String | File size (e.g., "841 Mb"). |
| `file_md5` | String | MD5 checksum of the file. |
| `attached` | Boolean | Whether the ISO is currently attached to a server. |
| `server_id` | Number | ID of the server the ISO is attached to (0 if not attached). |
| `server_name` | String | Name of the server the ISO is attached to. |
| `can_delete` | Boolean | Whether the ISO can be deleted. |

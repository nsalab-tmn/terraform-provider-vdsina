---
page_title: "vdsina_ssh_keys Data Source - terraform-provider-vdsina"
subcategory: "SSH Key"
description: |-
  Retrieves list of SSH keys in your VDSina account.
---

# vdsina_ssh_keys (Data Source)

Use this data source to get information about SSH keys stored in your VDSina account.

SSH keys can be used for passwordless authentication when creating or reinstalling servers.

## Example Usage

### Get all SSH keys

```hcl
data "vdsina_ssh_keys" "all" {}

output "ssh_keys" {
  value = data.vdsina_ssh_keys.all.ssh_keys
}
```

### Find SSH key by name

```hcl
data "vdsina_ssh_keys" "all" {}

locals {
  my_key = [
    for k in data.vdsina_ssh_keys.all.ssh_keys : k
    if k.name == "my-laptop"
  ][0]
}

output "my_key_id" {
  value = local.my_key.id
}
```

### Use existing SSH key for server

```hcl
data "vdsina_ssh_keys" "all" {}
data "vdsina_datacenters" "all" {}
data "vdsina_server_plans" "standard" {
  group_id = 2
}
data "vdsina_templates" "all" {}

resource "vdsina_server" "web" {
  datacenter  = data.vdsina_datacenters.all.datacenters[0].id
  server_plan = data.vdsina_server_plans.standard.plans[0].id
  template    = data.vdsina_templates.all.templates[0].id
  ssh_key     = data.vdsina_ssh_keys.all.ssh_keys[0].id
  name        = "web-server"
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `ssh_keys` - List of SSH keys. Each key has the following attributes:

### SSH Key Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique SSH key ID. Use this when creating servers. |
| `name` | String | SSH key name. |

~> **Note:** The public key data is not included in the list response. Use the `vdsina_ssh_key` resource to manage keys with full data.

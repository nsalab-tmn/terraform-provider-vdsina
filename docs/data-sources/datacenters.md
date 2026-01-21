---
page_title: "vdsina_datacenters Data Source - terraform-provider-vdsina"
subcategory: "Helpers"
description: |-
  Retrieves list of available VDSina datacenters.
---

# vdsina_datacenters (Data Source)

Use this data source to get information about available VDSina datacenters.

Datacenters are geographical locations where you can deploy your VPS servers. Each datacenter has a unique ID that you'll need when creating servers.

## Example Usage

### Get all datacenters

```hcl
data "vdsina_datacenters" "all" {}

output "datacenters" {
  value = data.vdsina_datacenters.all.datacenters
}
```

### Use datacenter ID for server creation

```hcl
data "vdsina_datacenters" "all" {}

# Use first datacenter
locals {
  datacenter_id = data.vdsina_datacenters.all.datacenters[0].id
}

resource "vdsina_server" "example" {
  datacenter  = local.datacenter_id
  server_plan = 1
  template    = 1
  name        = "my-server"
}
```

### Find datacenter by country

```hcl
data "vdsina_datacenters" "all" {}

# Find Netherlands datacenter
locals {
  nl_datacenter = [
    for dc in data.vdsina_datacenters.all.datacenters : dc
    if dc.country == "nl" && dc.active
  ][0]
}

output "nl_datacenter_id" {
  value = local.nl_datacenter.id
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `datacenters` - List of datacenters. Each datacenter has the following attributes:

### Datacenter Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique datacenter ID. Use this when creating servers. |
| `name` | String | Human-readable datacenter name (e.g., "Amsterdam 1, Netherlands"). |
| `country` | String | Two-letter country code (e.g., "nl", "ru", "de"). |
| `active` | Boolean | Whether the datacenter is available for new server deployments. |

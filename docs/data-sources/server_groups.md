---
page_title: "vdsina_server_groups Data Source - terraform-provider-vdsina"
subcategory: "Helpers"
description: |-
  Retrieves list of VDSina server tariff plan groups.
---

# vdsina_server_groups (Data Source)

Use this data source to get information about available VDSina server groups (tariff plan categories).

Server groups categorize different types of servers: Standard, Hi-CPU, GPU, Dedicated, Eternal, etc. Each group contains multiple server plans with different specifications.

## Example Usage

### Get all server groups

```hcl
data "vdsina_server_groups" "all" {}

output "server_groups" {
  value = data.vdsina_server_groups.all.groups
}
```

### Get server plans for a specific group

```hcl
data "vdsina_server_groups" "all" {}

# Use first group (Standard servers) to get plans
data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

output "standard_plans" {
  value = data.vdsina_server_plans.standard.plans
}
```

### Find GPU server group

```hcl
data "vdsina_server_groups" "all" {}

locals {
  gpu_groups = [
    for g in data.vdsina_server_groups.all.groups : g
    if can(regex("GPU", g.name)) && g.active
  ]
}

output "gpu_group_ids" {
  value = [for g in local.gpu_groups : g.id]
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `groups` - List of server groups. Each group has the following attributes:

### Group Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique server group ID. Use this when fetching server plans. |
| `name` | String | Group name (e.g., "Standard servers", "Hi-CPU servers", "Eternal servers"). |
| `type` | String | Group type (e.g., "vds"). |
| `active` | Boolean | Whether the group is available for new server orders. |
| `description` | String | Detailed description of the group's specifications. |

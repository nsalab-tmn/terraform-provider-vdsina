---
page_title: "vdsina_server_plans Data Source - terraform-provider-vdsina"
subcategory: "Helpers"
description: |-
  Retrieves list of VDSina server tariff plans for a specific group.
---

# vdsina_server_plans (Data Source)

Use this data source to get information about available server plans (tariffs) for a specific server group.

Each plan includes CPU, RAM, disk, traffic specifications and pricing information.

## Example Usage

### Get plans for Standard servers group

```hcl
data "vdsina_server_groups" "all" {}

data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

output "available_plans" {
  value = data.vdsina_server_plans.standard.plans
}
```

### Find cheapest plan

```hcl
data "vdsina_server_groups" "all" {}

data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

locals {
  cheapest_plan = data.vdsina_server_plans.standard.plans[0]
}

output "cheapest_plan" {
  value = {
    name = local.cheapest_plan.name
    cost = local.cheapest_plan.cost
    cpu  = local.cheapest_plan.cpu
    ram  = local.cheapest_plan.ram
    disk = local.cheapest_plan.disk
  }
}
```

### Find plan with specific RAM

```hcl
data "vdsina_server_groups" "all" {}

data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

locals {
  plans_4gb_ram = [
    for p in data.vdsina_server_plans.standard.plans : p
    if p.ram >= 4 && p.active
  ]
}

output "plans_with_4gb_ram" {
  value = local.plans_4gb_ram
}
```

### Create server with specific plan

```hcl
data "vdsina_server_groups" "all" {}
data "vdsina_datacenters" "all" {}
data "vdsina_templates" "all" {}

data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

resource "vdsina_server" "web" {
  datacenter  = data.vdsina_datacenters.all.datacenters[0].id
  server_plan = data.vdsina_server_plans.standard.plans[0].id
  template    = data.vdsina_templates.all.templates[0].id
  name        = "web-server"
}
```

## Argument Reference

The following arguments are required:

- `group_id` - (Required) Server group ID to get plans for. Get group IDs from `vdsina_server_groups` data source.

## Attributes Reference

The following attributes are exported:

- `plans` - List of server plans. Each plan has the following attributes:

### Plan Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique plan ID. Use this when creating servers. |
| `name` | String | Plan name (e.g., "2 RAM / 1 CPU / 40 NVMe"). |
| `cost` | Number | Daily cost with your discounts applied. |
| `full_cost` | Number | Daily cost without discounts. |
| `period` | String | Billing period (usually "day"). |
| `min_money` | Number | Minimum balance required to order this plan. |
| `can_bonus` | Boolean | Whether bonus balance can be used for this plan. |
| `description` | String | Plan description. |
| `active` | Boolean | Whether the plan is active. |
| `enable` | Boolean | Whether the plan is available for ordering. |
| `has_params` | Boolean | Whether the plan supports custom configuration (constructor). |
| `backup_cost` | Number | Backup cost per GB per day. |
| `cpu` | Number | Number of CPU cores. |
| `ram` | Number | RAM size in GB. |
| `disk` | Number | Disk size in GB (NVMe). |
| `traffic` | Number | Monthly traffic in TB (1 = 1TB). |

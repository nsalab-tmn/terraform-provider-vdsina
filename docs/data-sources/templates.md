---
page_title: "vdsina_templates Data Source - terraform-provider-vdsina"
subcategory: "Helpers"
description: |-
  Retrieves list of available VDSina OS templates.
---

# vdsina_templates (Data Source)

Use this data source to get information about available OS templates for server installation.

Each template includes minimum system requirements and list of compatible server plans.

## Example Usage

### Get all templates

```hcl
data "vdsina_templates" "all" {}

output "templates_count" {
  value = length(data.vdsina_templates.all.templates)
}
```

### Find Ubuntu templates

```hcl
data "vdsina_templates" "all" {}

locals {
  ubuntu_templates = [
    for t in data.vdsina_templates.all.templates : t
    if can(regex("Ubuntu", t.name)) && t.active
  ]
}

output "ubuntu_templates" {
  value = local.ubuntu_templates[*].name
}
```

### Find template compatible with specific plan

```hcl
data "vdsina_templates" "all" {}
data "vdsina_server_plans" "standard" {
  group_id = 2
}

locals {
  plan_id = data.vdsina_server_plans.standard.plans[0].id
  
  compatible_templates = [
    for t in data.vdsina_templates.all.templates : t
    if contains(t.server_plans, plan_id) && t.active
  ]
}

output "compatible_templates" {
  value = local.compatible_templates[*].name
}
```

### Create server with specific template

```hcl
data "vdsina_templates" "all" {}
data "vdsina_datacenters" "all" {}
data "vdsina_server_groups" "all" {}
data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

locals {
  ubuntu_template = [
    for t in data.vdsina_templates.all.templates : t
    if t.name == "Ubuntu 24.04"
  ][0]
}

resource "vdsina_server" "web" {
  datacenter  = data.vdsina_datacenters.all.datacenters[0].id
  server_plan = data.vdsina_server_plans.standard.plans[0].id
  template    = local.ubuntu_template.id
  name        = "web-server"
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `templates` - List of OS templates. Each template has the following attributes:

### Template Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique template ID. Use this when creating or reinstalling servers. |
| `name` | String | Template name (e.g., "Ubuntu 24.04", "Debian 12", "Windows Server 2022"). |
| `active` | Boolean | Whether the template is available for installation. |
| `ssh_key` | Boolean | Whether SSH key injection is supported for this template. |
| `template_group` | Number | Template group ID for categorization. |
| `server_plans` | List of Number | List of compatible server plan IDs. |
| `cpu_min` | Number | Minimum required CPU cores. |
| `ram_min` | Number | Minimum required RAM in GB. |
| `disk_min` | Number | Minimum required disk size in GB. |

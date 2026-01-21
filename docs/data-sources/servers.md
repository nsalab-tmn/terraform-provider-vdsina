---
page_title: "vdsina_servers Data Source - terraform-provider-vdsina"
subcategory: "Server"
description: |-
  Retrieves list of VDSina servers in your account.
---

# vdsina_servers (Data Source)

Use this data source to get information about all VPS servers in your VDSina account.

## Example Usage

### Get all servers

```hcl
data "vdsina_servers" "all" {}

output "servers" {
  value = data.vdsina_servers.all.servers
}
```

### Find active servers

```hcl
data "vdsina_servers" "all" {}

locals {
  active_servers = [
    for s in data.vdsina_servers.all.servers : s
    if s.status == "active"
  ]
}

output "active_server_ips" {
  value = local.active_servers[*].ip
}
```

### Find servers in specific datacenter

```hcl
data "vdsina_servers" "all" {}

locals {
  amsterdam_servers = [
    for s in data.vdsina_servers.all.servers : s
    if can(regex("Amsterdam", s.datacenter_name))
  ]
}

output "amsterdam_servers" {
  value = local.amsterdam_servers[*].name
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

- `servers` - List of servers. Each server has the following attributes:

### Server Attributes

| Attribute | Type | Description |
|-----------|------|-------------|
| `id` | Number | Unique server ID. |
| `name` | String | Server name. |
| `full_name` | String | Full server name with plan info (e.g., "Server 2 RAM / 1 CPU / 40 NVMe #12345"). |
| `status` | String | Server status: `new`, `active`, `block`, `deleted`. |
| `ip` | String | Public IP address. |
| `datacenter_id` | Number | Datacenter ID. |
| `datacenter_name` | String | Datacenter name (e.g., "Amsterdam 1, Netherlands"). |
| `plan_id` | Number | Server plan (tariff) ID. |
| `plan_name` | String | Server plan name (e.g., "2 RAM / 1 CPU / 40 NVMe"). |
| `template_id` | Number | OS template ID. |
| `template_name` | String | OS template name (e.g., "Ubuntu 24.04"). |
| `created` | String | Creation date. |
| `end` | String | Expiration date. |

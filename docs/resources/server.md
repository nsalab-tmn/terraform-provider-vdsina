# vdsina_server (Resource)

Manages a VDSina VPS server including standard, constructor, and GPU configurations.

## Example Usage

### Basic VPS

```hcl
resource "vdsina_server" "web_server" {
  datacenter  = 4
  server_plan = 1
  template    = 23
  name        = "web-server"
}
```

### VPS with SSH Key

```hcl
data "vdsina_datacenters" "all" {}
data "vdsina_server_groups" "all" {}

data "vdsina_server_plans" "standard" {
  group_id = data.vdsina_server_groups.all.groups[0].id
}

data "vdsina_templates" "all" {}

resource "vdsina_ssh_key" "deploy" {
  name = "deploy-key"
  key  = file("~/.ssh/id_rsa.pub")
}

resource "vdsina_server" "app_server" {
  datacenter  = data.vdsina_datacenters.all.datacenters[0].id
  server_plan = data.vdsina_server_plans.standard.plans[0].id
  template    = data.vdsina_templates.all.templates[0].id
  ssh_key     = vdsina_ssh_key.deploy.id
  name        = "app-server"
  host        = "app.example.com"
}
```

### VPS from Backup

```hcl
data "vdsina_backups" "all" {}

resource "vdsina_server" "restored" {
  datacenter  = 4
  server_plan = 1
  backup      = data.vdsina_backups.all.backups[0].id
  name        = "restored-server"
}
```

## Schema

### Required

* `datacenter` - (Number) Datacenter ID where the server will be created. Use `vdsina_datacenters` data source to find available datacenters.
* `server_plan` - (Number) Server plan ID (tariff). Use `vdsina_server_plans` data source for available plans.

### Optional

* `name` - (String) Name of the server.
* `template` - (Number) OS template ID. Use `vdsina_templates` data source for available templates. Mutually exclusive with `backup` and `iso`.
* `ssh_key` - (Number) SSH key ID. Use `vdsina_ssh_keys` data source for available keys.
* `host` - (String) Hostname for the server.
* `backup` - (Number) Backup ID to restore from. Mutually exclusive with `template` and `iso`.
* `iso` - (Number) ISO ID to install from. Mutually exclusive with `template` and `backup`.
* `autoprolong` - (Boolean) Enable auto-renewal when balance is sufficient. Default: `true`.
* `cpu` - (Number) Number of CPU cores (for constructor plans).
* `ram` - (Number) RAM in GB (for constructor plans).
* `disk` - (Number) Disk size in GB (for constructor plans).
* `gpu` - (Number) Number of GPUs (for GPU plans).

### Read-Only

* `id` - (String) The ID of the server.
* `full_name` - (String) Full server name including plan info.
* `status` - (String) Current status of the server: `new`, `active`, `block`, `notpaid`, `deleted`.
* `status_text` - (String) Status description.
* `ip` - (String) Public IP address assigned to the server.
* `ip_local` - (String) Local/private IP address.
* `created` - (String) Server creation timestamp.
* `end` - (String) Expiration date.
* `datacenter_name` - (String) Datacenter name.
* `template_name` - (String) OS template name.
* `plan_name` - (String) Server plan name.

## Timeouts

* `create` - (Default `20m`) Creation waits until the server has status `active` and a public IP. Statuses `block`, `notpaid` and `deleted` fail the create immediately.
* `update` - (Default `5m`)
* `delete` - (Default `10m`) Deletion waits until the API returns 404 or status `deleted` for the server.

Reading the server removes it from state only when the API returns 404 or status `deleted`; any other API error fails the refresh instead.

## Import

Servers can be imported using the server ID:

```bash
terraform import vdsina_server.example 613517
```

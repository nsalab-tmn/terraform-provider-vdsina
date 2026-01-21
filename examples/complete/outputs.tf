# examples/complete/outputs.tf
# Outputs for Data Sources and Resources validation

# ============================================
# 1. vdsina_datacenters
# ============================================

output "datacenters" {
  description = "List of available datacenters"
  value       = data.vdsina_datacenters.all.datacenters
}

output "datacenters_count" {
  description = "Number of datacenters"
  value       = length(data.vdsina_datacenters.all.datacenters)
}

output "first_datacenter" {
  description = "First datacenter details"
  value = {
    id      = data.vdsina_datacenters.all.datacenters[0].id
    name    = data.vdsina_datacenters.all.datacenters[0].name
    country = data.vdsina_datacenters.all.datacenters[0].country
    active  = data.vdsina_datacenters.all.datacenters[0].active
  }
}

# ============================================
# 2. vdsina_server_groups
# ============================================

output "server_groups" {
  description = "List of server groups"
  value       = data.vdsina_server_groups.all.groups
}

output "server_groups_count" {
  description = "Number of server groups"
  value       = length(data.vdsina_server_groups.all.groups)
}

output "first_server_group" {
  description = "First server group details"
  value = {
    id          = data.vdsina_server_groups.all.groups[0].id
    name        = data.vdsina_server_groups.all.groups[0].name
    type        = data.vdsina_server_groups.all.groups[0].type
    active      = data.vdsina_server_groups.all.groups[0].active
    description = data.vdsina_server_groups.all.groups[0].description
  }
}

# ============================================
# 3. vdsina_server_plans
# ============================================

output "server_plans" {
  description = "List of server plans for Standard group"
  value       = data.vdsina_server_plans.standard.plans
}

output "server_plans_count" {
  description = "Number of server plans"
  value       = length(data.vdsina_server_plans.standard.plans)
}

output "cheapest_plan" {
  description = "Cheapest server plan"
  value = {
    id          = data.vdsina_server_plans.standard.plans[0].id
    name        = data.vdsina_server_plans.standard.plans[0].name
    cost        = data.vdsina_server_plans.standard.plans[0].cost
    cpu         = data.vdsina_server_plans.standard.plans[0].cpu
    ram         = data.vdsina_server_plans.standard.plans[0].ram
    disk        = data.vdsina_server_plans.standard.plans[0].disk
    backup_cost = data.vdsina_server_plans.standard.plans[0].backup_cost
  }
}

# ============================================
# 4. vdsina_templates
# ============================================

output "templates_count" {
  description = "Number of OS templates"
  value       = length(data.vdsina_templates.all.templates)
}

output "first_template" {
  description = "First OS template details"
  value = {
    id             = data.vdsina_templates.all.templates[0].id
    name           = data.vdsina_templates.all.templates[0].name
    active         = data.vdsina_templates.all.templates[0].active
    ssh_key        = data.vdsina_templates.all.templates[0].ssh_key
    template_group = data.vdsina_templates.all.templates[0].template_group
    cpu_min        = data.vdsina_templates.all.templates[0].cpu_min
    ram_min        = data.vdsina_templates.all.templates[0].ram_min
    disk_min       = data.vdsina_templates.all.templates[0].disk_min
    server_plans   = data.vdsina_templates.all.templates[0].server_plans
  }
}

# ============================================
# 5. vdsina_servers
# ============================================

output "servers_count" {
  description = "Number of servers"
  value       = length(data.vdsina_servers.all.servers)
}

output "servers" {
  description = "List of servers"
  value = [for s in data.vdsina_servers.all.servers : {
    id              = s.id
    name            = s.name
    status          = s.status
    ip              = s.ip
    datacenter_name = s.datacenter_name
    plan_name       = s.plan_name
  }]
}

# ============================================
# 6. vdsina_ssh_keys
# ============================================

output "ssh_keys" {
  description = "List of SSH keys"
  value       = data.vdsina_ssh_keys.all.ssh_keys
}

# ============================================
# 7. vdsina_account
# ============================================

output "account" {
  description = "Account information"
  value = {
    id                 = data.vdsina_account.me.account_id
    name               = data.vdsina_account.me.name
    created            = data.vdsina_account.me.created
    forecast           = data.vdsina_account.me.forecast
    can_add_service    = data.vdsina_account.me.can_add_service
    can_convert_to_cash = data.vdsina_account.me.can_convert_to_cash
  }
}

# ============================================
# 8. vdsina_account_balance
# ============================================

output "balance" {
  description = "Account balance"
  value = {
    real    = data.vdsina_account_balance.current.real
    bonus   = data.vdsina_account_balance.current.bonus
    partner = data.vdsina_account_balance.current.partner
    total   = data.vdsina_account_balance.current.total
  }
}

# ============================================
# 9. vdsina_account_limits
# ============================================

output "limits_server" {
  description = "Server limits"
  value       = data.vdsina_account_limits.limits.server
}

output "limits_backup" {
  description = "Backup limits"
  value       = data.vdsina_account_limits.limits.backup
}

output "limits_dns" {
  description = "DNS limits"
  value       = data.vdsina_account_limits.limits.dns
}

# ============================================
# 10. vdsina_iso_list
# ============================================

output "iso_list" {
  description = "List of ISO images"
  value = [for iso in data.vdsina_iso_list.all.isos : {
    id          = iso.id
    name        = iso.name
    status      = iso.status
    file_size   = iso.file_size
    attached    = iso.attached
    server_name = iso.server_name
  }]
}

# ============================================
# 11. vdsina_backups
# ============================================

output "backups" {
  description = "List of backups"
  value = [for b in data.vdsina_backups.all.backups : {
    id          = b.id
    name        = b.name
    status      = b.status
    server_name = b.server_name
    can_delete  = b.can_delete
  }]
}

# ============================================
# 12. vdsina_dns_zones
# ============================================

output "dns_zones" {
  description = "List of DNS zones"
  value = [for z in data.vdsina_dns_zones.all.zones : {
    id          = z.id
    name        = z.name
    status      = z.status
    real        = z.real
    can_delete  = z.can_delete
  }]
}

# ============================================
# 13. vdsina_ips
# ============================================

output "ips" {
  description = "List of IP addresses"
  value = [for ip in data.vdsina_ips.all.ips : {
    id                 = ip.id
    ip                 = ip.ip
    type               = ip.type
    host               = ip.host
    gateway            = ip.gateway
    netmask            = ip.netmask
    datacenter_name    = ip.datacenter_name
  }]
}

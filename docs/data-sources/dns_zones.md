# vdsina_dns_zones (Data Source)

Retrieves a list of DNS zones (domains) from your VDSina account.

## Example Usage

### Get all DNS zones

```hcl
data "vdsina_dns_zones" "all" {}

output "dns_zones" {
  value = data.vdsina_dns_zones.all.zones
}
```

### Check if DNS is properly configured

```hcl
data "vdsina_dns_zones" "all" {}

output "active_zones" {
  value = [for z in data.vdsina_dns_zones.all.zones : z.name if z.status == "active" && z.real]
}
```

## Argument Reference

This data source has no required arguments.

## Attributes Reference

The following attributes are exported:

### zones

A list of DNS zones. Each zone has the following attributes:

* `id` - (Integer) DNS zone ID.
* `name` - (String) Domain name (e.g., "example.com").
* `full_name` - (String) Full service name (e.g., "DNS #290676 – example.com").
* `created` - (String) Creation date.
* `updated` - (String) Last update date.
* `end` - (String) Service end date.
* `status` - (String) Zone status. Possible values: `new`, `active`, `block`, `notpaid`, `deleted`.
* `status_text` - (String) Human-readable status description.
* `real` - (Boolean) Whether DNS servers are correctly set up in domain NS records.
* `can_delete` - (Boolean) Whether this DNS zone can be deleted.

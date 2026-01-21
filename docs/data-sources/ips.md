# vdsina_ips (Data Source)

Retrieves a list of IP addresses from your VDSina account.

## Example Usage

### Get all IP addresses

```hcl
data "vdsina_ips" "all" {}

output "ips" {
  value = data.vdsina_ips.all.ips
}
```

### Filter IPv4 addresses only

```hcl
data "vdsina_ips" "all" {}

output "ipv4_addresses" {
  value = [for ip in data.vdsina_ips.all.ips : ip.ip if ip.type == "4"]
}
```

### Get IPs by datacenter

```hcl
data "vdsina_ips" "all" {}

output "amsterdam_ips" {
  value = [for ip in data.vdsina_ips.all.ips : ip.ip if ip.datacenter_country == "nl"]
}
```

## Argument Reference

This data source has no required arguments.

## Attributes Reference

The following attributes are exported:

### ips

A list of IP addresses. Each IP has the following attributes:

* `id` - (Integer) IP address ID.
* `ip` - (String) IP address (IPv4 or IPv6).
* `type` - (String) IP type: `4` for IPv4, `6` for IPv6.
* `host` - (String) Hostname associated with the IP address.
* `gateway` - (String) Gateway address.
* `netmask` - (String) Network mask.
* `is_net` - (Boolean) Whether this is a network (subnet).
* `datacenter_id` - (Integer) Datacenter ID where IP is located.
* `datacenter_name` - (String) Datacenter name.
* `datacenter_country` - (String) Datacenter country code (e.g., "nl", "ru").

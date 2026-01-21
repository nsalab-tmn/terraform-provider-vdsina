# vdsina_dns (Resource)

Manages a VDSina DNS zone. Creates a DNS zone (domain) in VDSina DNS hosting.

~> **Note:** DNS zones cannot be updated after creation. Changing the domain name requires recreation.

## Example Usage

### Basic DNS Zone

```hcl
resource "vdsina_dns" "example" {
  name = "example.com"
}
```

### DNS Zone with Default A Record

```hcl
resource "vdsina_dns" "website" {
  name = "mywebsite.com"
  ip   = "78.40.199.136"
}
```

~> **Note:** When `ip` is specified, VDSina automatically creates default A records for the root domain and `mx.domain`.

### Complete DNS Setup

```hcl
resource "vdsina_dns" "domain" {
  name = "example.com"
  ip   = vdsina_server.web.ip
}

resource "vdsina_dns_record" "www" {
  zone_id = vdsina_dns.domain.id
  host    = "www"
  type    = "CNAME"
  value   = "example.com"
}

resource "vdsina_dns_record" "mail" {
  zone_id  = vdsina_dns.domain.id
  host     = "@"
  type     = "MX"
  value    = "mail.example.com"
  priority = 10
}
```

## Schema

### Required

* `name` - (String, ForceNew) Domain name for the DNS zone (e.g., "example.com").

### Optional

* `ip` - (String, ForceNew) IP address for generating default A record.

### Read-Only

* `id` - (String) The ID of the DNS zone.
* `full_name` - (String) Full DNS service name (e.g., "DNS #290676 – example.com").
* `status` - (String) DNS zone status: `new`, `active`, `block`, `notpaid`, `deleted`.
* `status_text` - (String) Status description.
* `real` - (Boolean) Whether DNS servers are correctly set up in domain NS records.
* `can_delete` - (Boolean) Whether this DNS zone can be deleted.
* `created` - (String) Creation timestamp.
* `updated` - (String) Last update timestamp.
* `end` - (String) Expiration date.

## Import

DNS zones can be imported using the zone ID:

```bash
terraform import vdsina_dns.example 290676
```

## NS Records

After creating a DNS zone, configure your domain registrar to use VDSina nameservers:

- `ns1.vdsina.com`
- `ns2.vdsina.com`

# vdsina_dns_record (Resource)

Manages DNS records within a VDSina DNS zone.

## Example Usage

### A Record

```hcl
resource "vdsina_dns" "domain" {
  name = "example.com"
}

resource "vdsina_dns_record" "web" {
  zone_id = vdsina_dns.domain.id
  host    = "www"
  type    = "A"
  value   = "78.40.199.136"
}
```

### Root Domain A Record

```hcl
resource "vdsina_dns_record" "root" {
  zone_id = vdsina_dns.domain.id
  host    = "@"
  type    = "A"
  value   = "78.40.199.136"
}
```

### MX Record with Priority

```hcl
resource "vdsina_dns_record" "mail" {
  zone_id  = vdsina_dns.domain.id
  host     = "@"
  type     = "MX"
  value    = "mail.example.com"
  priority = 10
}
```

### CNAME Record

```hcl
resource "vdsina_dns_record" "alias" {
  zone_id = vdsina_dns.domain.id
  host    = "blog"
  type    = "CNAME"
  value   = "example.com"
}
```

### TXT Record (SPF)

```hcl
resource "vdsina_dns_record" "spf" {
  zone_id = vdsina_dns.domain.id
  host    = "@"
  type    = "TXT"
  value   = "v=spf1 include:_spf.google.com ~all"
}
```

### CAA Record

```hcl
resource "vdsina_dns_record" "caa" {
  zone_id = vdsina_dns.domain.id
  host    = "@"
  type    = "CAA"
  tag     = "issue"
  value   = "letsencrypt.org"
}
```

### Wildcard Record

```hcl
resource "vdsina_dns_record" "wildcard" {
  zone_id = vdsina_dns.domain.id
  host    = "*"
  type    = "A"
  value   = "78.40.199.136"
}
```

## Schema

### Required

* `zone_id` - (Number, ForceNew) DNS zone ID where the record will be created.
* `host` - (String, ForceNew) Hostname for the record. Use `@` for root domain, `*` for wildcard, or subdomain name.
* `type` - (String, ForceNew) Record type: `A`, `AAAA`, `CNAME`, `MX`, `NS`, `SRV`, `CAA`, `TXT`.
* `value` - (String) Record value (IP address for A/AAAA, domain for CNAME/MX/NS, text for TXT).

### Optional

* `priority` - (Number) Priority for MX and SRV records.
* `tag` - (String) Tag for CAA records: `issue`, `issuewild`, `iodef`, `unknown`.

### Read-Only

* `id` - (String) The ID of the DNS record.
* `timestamp` - (String) Record creation/update timestamp.

## Import

DNS records can be imported using the format `zone_id:record_id`:

```bash
terraform import vdsina_dns_record.example 290676:12345
```

## Notes

- The `host` field is normalized by the API. When you create a record with `host = "@"`, the API returns the full domain name with a trailing dot (e.g., `example.com.`).
- DNS propagation may take up to 600 seconds (default TTL).

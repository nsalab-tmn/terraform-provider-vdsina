# examples/dns-zone/main.tf
# DNS zone with records

terraform {
  required_providers {
    vdsina = {
      source  = "scinfra-pro/vdsina"
      version = "~> 0.2.0"
    }
  }
}

provider "vdsina" {
  api_token = var.vdsina_api_token
  base_url  = var.vdsina_base_url
}

# Create DNS zone
resource "vdsina_dns" "domain" {
  name = var.domain_name
  ip   = var.default_ip  # Creates default A record for root
}

# A record for www
resource "vdsina_dns_record" "www" {
  zone_id = vdsina_dns.domain.id
  host    = "www"
  type    = "A"
  value   = var.default_ip
}

# CNAME for blog
resource "vdsina_dns_record" "blog" {
  zone_id = vdsina_dns.domain.id
  host    = "blog"
  type    = "CNAME"
  value   = var.domain_name
}

# MX record for mail
resource "vdsina_dns_record" "mx" {
  zone_id  = vdsina_dns.domain.id
  host     = "@"
  type     = "MX"
  value    = "mail.${var.domain_name}"
  priority = 10
}

# TXT record for SPF
resource "vdsina_dns_record" "spf" {
  zone_id = vdsina_dns.domain.id
  host    = "@"
  type    = "TXT"
  value   = "v=spf1 a mx ~all"
}

output "dns_zone_id" {
  description = "DNS zone ID"
  value       = vdsina_dns.domain.id
}

output "nameservers" {
  description = "Configure these nameservers at your registrar"
  value       = ["ns1.vdsina.com", "ns2.vdsina.com"]
}

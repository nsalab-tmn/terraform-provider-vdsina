# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2026-01-21

### Added

- Initial release of VDSina Terraform Provider
- **Resources:**
  - `vdsina_server` - VPS server management with full CRUD
  - `vdsina_ssh_key` - SSH key management
  - `vdsina_iso` - ISO image upload and management
  - `vdsina_backup` - Server backup creation and management
  - `vdsina_dns` - DNS zone management
  - `vdsina_dns_record` - DNS record management (A, AAAA, CNAME, MX, NS, SRV, CAA, TXT)
  - `vdsina_server_reboot` - Server reboot action
  - `vdsina_server_reinstall` - OS reinstall action
  - `vdsina_server_password` - Password management
  - `vdsina_server_plan_change` - Plan upgrade action
  - `vdsina_server_start_prolong` - Server prolong action
  - `vdsina_server_iso` - ISO attach/detach action
  - `vdsina_backup_restore` - Backup restore action
- **Data Sources:**
  - `vdsina_datacenters` - List available datacenters
  - `vdsina_server_groups` - List server plan groups
  - `vdsina_server_plans` - List server plans by group
  - `vdsina_templates` - List OS templates
  - `vdsina_servers` - List all servers
  - `vdsina_ssh_keys` - List SSH keys
  - `vdsina_account` - Account information
  - `vdsina_account_balance` - Account balance
  - `vdsina_account_limits` - Account service limits
  - `vdsina_iso_list` - List ISO images
  - `vdsina_backups` - List backups
  - `vdsina_dns_zones` - List DNS zones
  - `vdsina_ips` - List IP addresses
- Full documentation for all resources and data sources
- CI/CD pipeline with GitHub Actions
- GPG-signed releases for Terraform Registry

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2026-09-17

First release of the `nsalab-tmn/vdsina` fork.

### Fixed

- `vdsina_server` create waits until the server is `active` with a public IP, so `ip` is known in the same apply (previously it returned right after `POST /server` with no IP). `block`, `notpaid` and `deleted` fail the create.
- `vdsina_server` delete waits until the server is gone (404 or status `deleted`).
- `vdsina_server` and `vdsina_ssh_key` read remove the resource from state only on 404 (or a `deleted` server); other API errors are returned instead of silently dropping the resource and planning a re-create.
- `IsNotFound`, `IsUnauthorized` and `IsForbidden` recognise wrapped API errors.
- `vdsina_server` `template`, `ssh_key` and `host` are Computed: when omitted, the values VDSina assigns (e.g. the default hostname) no longer force a replacement on the next plan.
- `vdsina_server` `autoprolong = false` is applied once the server is active and its failure is reported; the update sent right after `POST /server` did not take effect.

### Changed

- `vdsina_server` default timeouts: create 20m (was 10m), delete 10m (was 5m).
- Registry address `nsalab-tmn/vdsina`.

## [0.2.0] - 2026-01-21

### Changed

- Refactored all `d.Set()` calls to properly check and return errors
- Follows the same error handling pattern as terraform-provider-aeza
- Improved code quality for golangci-lint compliance

## [0.1.0] - 2026-01-21 (unreleased)

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

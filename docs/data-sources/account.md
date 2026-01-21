---
page_title: "vdsina_account Data Source - terraform-provider-vdsina"
subcategory: "Account"
description: |-
  Retrieves VDSina account information.
---

# vdsina_account (Data Source)

Use this data source to get information about your VDSina account, including creation date, shutdown forecast, and available permissions.

## Example Usage

```hcl
data "vdsina_account" "me" {}

output "account_info" {
  value = {
    id       = data.vdsina_account.me.account_id
    name     = data.vdsina_account.me.name
    forecast = data.vdsina_account.me.forecast
  }
}
```

### Check if can create services

```hcl
data "vdsina_account" "me" {}

output "can_create_servers" {
  value = data.vdsina_account.me.can_add_service
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

| Attribute | Type | Description |
|-----------|------|-------------|
| `account_id` | Number | Unique account ID. |
| `name` | String | Account name (usually email or username). |
| `created` | String | Account creation date and time. |
| `forecast` | String | Shutdown forecast date - the date until which there are enough funds to pay for all services. |
| `can_add_user` | Boolean | Whether you can create sub-users. |
| `can_add_service` | Boolean | Whether you can order new services (servers, etc.). |
| `can_convert_to_cash` | Boolean | Whether you can withdraw money from partner balance. |

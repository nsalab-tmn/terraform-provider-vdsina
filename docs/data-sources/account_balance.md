---
page_title: "vdsina_account_balance Data Source - terraform-provider-vdsina"
subcategory: "Account"
description: |-
  Retrieves VDSina account balance.
---

# vdsina_account_balance (Data Source)

Use this data source to get your VDSina account balance information.

The balance consists of three parts:
- **Real** - Main balance for paying services
- **Bonus** - Promotional bonus balance
- **Partner** - Partner program earnings

## Example Usage

```hcl
data "vdsina_account_balance" "current" {}

output "balance" {
  value = {
    real  = data.vdsina_account_balance.current.real
    bonus = data.vdsina_account_balance.current.bonus
    total = data.vdsina_account_balance.current.total
  }
}
```

### Check if balance is low

```hcl
data "vdsina_account_balance" "current" {}

output "low_balance_warning" {
  value = data.vdsina_account_balance.current.total < 10 ? "Warning: Low balance!" : "Balance OK"
}
```

## Argument Reference

This data source has no arguments.

## Attributes Reference

The following attributes are exported:

| Attribute | Type | Description |
|-----------|------|-------------|
| `real` | Number | Main account balance (USD or RUB depending on API). |
| `bonus` | Number | Bonus/promotional balance. |
| `partner` | Number | Partner program balance. |
| `total` | Number | Total balance (real + bonus + partner). |

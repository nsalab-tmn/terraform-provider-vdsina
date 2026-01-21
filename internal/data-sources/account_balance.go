// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/account_balance.go
package datasources

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func AccountBalanceDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina account balance",

		ReadContext: accountBalanceRead,

		Schema: map[string]*schema.Schema{
			"real": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Real (main) account balance",
			},
			"bonus": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Bonus account balance",
			},
			"partner": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Partner account balance",
			},
			"total": {
				Type:        schema.TypeFloat,
				Computed:    true,
				Description: "Total balance (real + bonus + partner)",
			},
		},
	}
}

func accountBalanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	balance, err := c.GetAccountBalance(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read account balance: %w", err))
	}

	real, _ := strconv.ParseFloat(balance.Real, 64)
	bonus, _ := strconv.ParseFloat(balance.Bonus, 64)
	partner, _ := strconv.ParseFloat(balance.Partner, 64)

	d.SetId("account_balance")

	if err := d.Set("real", real); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("bonus", bonus); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("partner", partner); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("total", real+bonus+partner); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

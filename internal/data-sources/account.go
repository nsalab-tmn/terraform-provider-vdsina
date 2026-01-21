// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/account.go
package datasources

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func AccountDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina account information",

		ReadContext: accountRead,

		Schema: map[string]*schema.Schema{
			"account_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Account ID",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account name (email)",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Account creation date",
			},
			"forecast": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Forecast shutdown date",
			},
			"can_add_user": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can add sub-users",
			},
			"can_add_service": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can create new services",
			},
			"can_convert_to_cash": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can convert bonus to real balance",
			},
		},
	}
}

func accountRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	account, err := c.GetAccount(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read account: %w", err))
	}

	d.SetId(strconv.Itoa(account.Account.ID))
	d.Set("account_id", account.Account.ID)
	d.Set("name", account.Account.Name)
	d.Set("created", account.Created)
	d.Set("forecast", account.Forecast)
	d.Set("can_add_user", account.Can.AddUser)
	d.Set("can_add_service", account.Can.AddService)
	d.Set("can_convert_to_cash", account.Can.ConvertToCash)

	return nil
}

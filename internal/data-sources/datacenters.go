// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/datacenters.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func DatacentersDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina datacenters",

		ReadContext: datacentersRead,

		Schema: map[string]*schema.Schema{
			"datacenters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of available datacenters",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Datacenter ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Datacenter name (e.g., 'Amsterdam 1', 'Moscow 1')",
						},
						"country": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Country code (e.g., 'nl', 'ru', 'de')",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the datacenter is active",
						},
					},
				},
			},
		},
	}
}

func datacentersRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	datacenters, err := c.GetDatacenters(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read datacenters: %w", err))
	}

	datacenterList := make([]map[string]interface{}, len(datacenters))
	for i, dc := range datacenters {
		datacenterList[i] = map[string]interface{}{
			"id":      dc.ID,
			"name":    dc.Name,
			"country": dc.Country,
			"active":  dc.Active,
		}
	}

	d.SetId("datacenters")
	if err := d.Set("datacenters", datacenterList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set datacenters in state: %w", err))
	}

	return nil
}

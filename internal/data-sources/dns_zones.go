// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/dns_zones.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func DNSZonesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of DNS zones",

		ReadContext: dnsZonesRead,

		Schema: map[string]*schema.Schema{
			"zones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of DNS zones",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "DNS zone ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full service name",
						},
						"created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation date",
						},
						"updated": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last update date",
						},
						"end": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End date",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status: new, active, block, notpaid, deleted",
						},
						"status_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status description",
						},
						"real": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "DNS servers correctly set up in domain NS records",
						},
						"can_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can this DNS zone be deleted",
						},
					},
				},
			},
		},
	}
}

func dnsZonesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	zones, err := c.GetDNSZones(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read DNS zones: %w", err))
	}

	d.SetId("dns_zones")

	zoneList := make([]map[string]interface{}, len(zones))
	for i, zone := range zones {
		zoneList[i] = map[string]interface{}{
			"id":          zone.ID,
			"name":        zone.Name,
			"full_name":   zone.FullName,
			"created":     zone.Created,
			"updated":     zone.Updated,
			"end":         zone.End,
			"status":      zone.Status,
			"status_text": zone.StatusText,
			"real":        zone.Real,
			"can_delete":  zone.Can.Delete,
		}
	}

	d.Set("zones", zoneList)

	return nil
}

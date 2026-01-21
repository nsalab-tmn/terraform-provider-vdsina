// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/servers.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ServersDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of VDSina servers",

		ReadContext: serversRead,

		Schema: map[string]*schema.Schema{
			"servers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of servers",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Server ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server name",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full server name with plan info",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server status (new, active, block, deleted)",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public IP address",
						},
						"datacenter_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Datacenter ID",
						},
						"datacenter_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Datacenter name",
						},
						"plan_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Server plan ID",
						},
						"plan_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server plan name",
						},
						"template_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "OS template ID",
						},
						"template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OS template name",
						},
						"created": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation date",
						},
						"end": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expiration date",
						},
					},
				},
			},
		},
	}
}

func serversRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	servers, err := c.GetServers(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read servers: %w", err))
	}

	serverList := make([]map[string]interface{}, len(servers))
	for i, s := range servers {
		serverList[i] = map[string]interface{}{
			"id":              s.ID,
			"name":            s.Name,
			"full_name":       s.FullName,
			"status":          s.Status,
			"ip":              s.IP.IP,
			"datacenter_id":   s.Datacenter.ID,
			"datacenter_name": s.Datacenter.Name,
			"plan_id":         s.ServerPlan.ID,
			"plan_name":       s.ServerPlan.Name,
			"template_id":     s.Template.ID,
			"template_name":   s.Template.Name,
			"created":         s.Created,
			"end":             s.End,
		}
	}

	d.SetId("servers")
	if err := d.Set("servers", serverList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set servers in state: %w", err))
	}

	return nil
}

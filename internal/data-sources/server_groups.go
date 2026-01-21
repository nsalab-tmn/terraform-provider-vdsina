// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/server_groups.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ServerGroupsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina server groups (Standard, AMD, Hi-CPU, GPU, Eternal)",

		ReadContext: serverGroupsRead,

		Schema: map[string]*schema.Schema{
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of server groups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Server group ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server group name (e.g., 'Standard', 'Hi-CPU', 'GPU')",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Server group type (e.g., 'vds')",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the group is active",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group description",
						},
					},
				},
			},
		},
	}
}

func serverGroupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	groups, err := c.GetServerGroups(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read server groups: %w", err))
	}

	groupList := make([]map[string]interface{}, len(groups))
	for i, g := range groups {
		groupList[i] = map[string]interface{}{
			"id":          g.ID,
			"name":        g.Name,
			"type":        g.Type,
			"active":      g.Active,
			"description": g.Description,
		}
	}

	d.SetId("server_groups")
	if err := d.Set("groups", groupList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set groups in state: %w", err))
	}

	return nil
}

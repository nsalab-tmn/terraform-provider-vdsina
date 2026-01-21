// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/server_plans.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ServerPlansDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina server plans by group",

		ReadContext: serverPlansRead,

		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Server group ID to get plans for",
			},
			"plans": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of server plans",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Plan ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Plan name (e.g., '2 RAM / 1 CPU / 40 NVMe')",
						},
						"cost": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Daily cost",
						},
						"full_cost": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Full cost without discount",
						},
						"period": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Billing period (day)",
						},
						"min_money": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Minimum balance required",
						},
						"can_bonus": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can use bonus balance",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Plan description",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the plan is active",
						},
						"enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the plan is enabled for ordering",
						},
						"has_params": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether plan supports custom configuration (constructor)",
						},
						"backup_cost": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Backup cost per GB per day",
						},
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of CPU cores",
						},
						"ram": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "RAM in GB",
						},
						"disk": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Disk size in GB",
						},
						"traffic": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Traffic limit in GB (0 = unlimited)",
						},
					},
				},
			},
		},
	}
}

func serverPlansRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	groupID := d.Get("group_id").(int)

	plans, err := c.GetServerPlans(ctx, groupID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read server plans for group %d: %w", groupID, err))
	}

	planList := make([]map[string]interface{}, len(plans))
	for i, p := range plans {
		planList[i] = map[string]interface{}{
			"id":          p.ID,
			"name":        p.Name,
			"cost":        p.Cost,
			"full_cost":   p.FullCost,
			"period":      p.Period,
			"min_money":   p.MinMoney,
			"can_bonus":   p.CanBonus,
			"description": p.Description,
			"active":      p.Active,
			"enable":      p.Enable,
			"has_params":  p.HasParams,
			"backup_cost": p.Backup.Cost,
			"cpu":         p.Data.CPU.Value,
			"ram":         p.Data.RAM.Value,
			"disk":        p.Data.Disk.Value,
			"traffic":     p.Data.Traff.Value,
		}
	}

	d.SetId(fmt.Sprintf("server_plans_%d", groupID))
	if err := d.Set("plans", planList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set plans in state: %w", err))
	}

	return nil
}

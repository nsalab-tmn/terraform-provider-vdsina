// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/backups.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func BackupsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of backups",

		ReadContext: backupsRead,

		Schema: map[string]*schema.Schema{
			"backups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of backups",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup name",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full backup name with size",
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
							Description: "Status description (e.g. 'Creation (9%) Processing')",
						},
						"datacenter_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Datacenter ID where backup is stored",
						},
						"datacenter_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Datacenter name",
						},
						"datacenter_country": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Datacenter country code",
						},
						"server_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Parent server ID (0 if not available)",
						},
						"server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parent server name",
						},
						"can_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can this backup be updated",
						},
						"can_prolong": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can this backup be prolonged",
						},
						"can_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can this backup be deleted",
						},
					},
				},
			},
		},
	}
}

func backupsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	backups, err := c.GetBackups(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read backups: %w", err))
	}

	d.SetId("backups")

	backupList := make([]map[string]interface{}, len(backups))
	for i, backup := range backups {
		backupMap := map[string]interface{}{
			"id":                 backup.ID,
			"name":               backup.Name,
			"full_name":          backup.FullName,
			"created":            backup.Created,
			"updated":            backup.Updated,
			"end":                backup.End,
			"status":             backup.Status,
			"status_text":        backup.StatusText,
			"datacenter_id":      backup.Datacenter.ID,
			"datacenter_name":    backup.Datacenter.Name,
			"datacenter_country": backup.Datacenter.Country,
			"can_update":         backup.Can.Update,
			"can_prolong":        backup.Can.Prolong,
			"can_delete":         backup.Can.Delete,
		}

		if backup.Server != nil {
			backupMap["server_id"] = backup.Server.ID
			backupMap["server_name"] = backup.Server.Name
		} else {
			backupMap["server_id"] = 0
			backupMap["server_name"] = ""
		}

		backupList[i] = backupMap
	}

	if err := d.Set("backups", backupList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set backups in state: %w", err))
	}

	return nil
}

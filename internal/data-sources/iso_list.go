// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/iso_list.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ISOListDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of ISO images",

		ReadContext: isoListRead,

		Schema: map[string]*schema.Schema{
			"isos": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of ISO images",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ISO service ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ISO file name",
						},
						"full_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Full ISO name with ID",
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
						"file_size": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "File size (e.g. '841 Mb')",
						},
						"file_md5": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "File MD5 checksum",
						},
						"attached": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is ISO attached to any server",
						},
						"server_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of server ISO is attached to (0 if not attached)",
						},
						"server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of server ISO is attached to",
						},
						"can_delete": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Can this ISO be deleted",
						},
					},
				},
			},
		},
	}
}

func isoListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	isos, err := c.GetISOs(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read ISO list: %w", err))
	}

	d.SetId("iso_list")

	isoList := make([]map[string]interface{}, len(isos))
	for i, iso := range isos {
		isoMap := map[string]interface{}{
			"id":          iso.ID,
			"name":        iso.Name,
			"full_name":   iso.FullName,
			"created":     iso.Created,
			"updated":     iso.Updated,
			"end":         iso.End,
			"status":      iso.Status,
			"status_text": iso.StatusText,
			"file_size":   iso.File.Size,
			"file_md5":    iso.File.MD5,
			"attached":    iso.Attached,
			"can_delete":  iso.Can.Delete,
		}

		if iso.Server != nil {
			isoMap["server_id"] = iso.Server.ID
			isoMap["server_name"] = iso.Server.Name
		} else {
			isoMap["server_id"] = 0
			isoMap["server_name"] = ""
		}

		isoList[i] = isoMap
	}

	if err := d.Set("isos", isoList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set isos in state: %w", err))
	}

	return nil
}

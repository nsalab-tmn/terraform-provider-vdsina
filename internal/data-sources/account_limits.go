// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/account_limits.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

func limitSchema(description string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: description,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"max": {
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Maximum allowed",
				},
				"now": {
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Currently used",
				},
				"child_max": {
					Type:        schema.TypeInt,
					Computed:    true,
					Description: "Maximum per child (for IP/GPU)",
				},
			},
		},
	}
}

func AccountLimitsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina account limits by service type",

		ReadContext: accountLimitsRead,

		Schema: map[string]*schema.Schema{
			"server":       limitSchema("VPS servers limit"),
			"server_ip4":   limitSchema("Additional IPv4 addresses limit"),
			"server_ip6":   limitSchema("Additional IPv6 addresses limit"),
			"iso":          limitSchema("ISO images limit"),
			"backup":       limitSchema("Backups limit"),
			"ssl":          limitSchema("SSL certificates limit"),
			"domain":       limitSchema("Domains limit"),
			"dns":          limitSchema("DNS zones limit"),
			"extdisk_hdd":  limitSchema("External HDD disks limit"),
			"extdisk_nvme": limitSchema("External NVMe disks limit"),
			"reserve_ip":   limitSchema("Reserved IP addresses limit"),
			"gpu":          limitSchema("GPU limit"),
		},
	}
}

func limitToMap(l models.LimitInfo) []map[string]interface{} {
	return []map[string]interface{}{{
		"max":       l.Max,
		"now":       l.Now,
		"child_max": l.ChildMax,
	}}
}

func accountLimitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	limits, err := c.GetAccountLimits(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read account limits: %w", err))
	}

	d.SetId("account_limits")

	if err := d.Set("server", limitToMap(limits.Server)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("server_ip4", limitToMap(limits.ServerIP4)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("server_ip6", limitToMap(limits.ServerIP6)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("iso", limitToMap(limits.ISO)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("backup", limitToMap(limits.Backup)); err != nil {
		return diag.FromErr(err)
	}

	if limits.SSL != nil {
		if err := d.Set("ssl", limitToMap(*limits.SSL)); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("ssl", []map[string]interface{}{}); err != nil {
			return diag.FromErr(err)
		}
	}

	if limits.Domain != nil {
		if err := d.Set("domain", limitToMap(*limits.Domain)); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("domain", []map[string]interface{}{}); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("dns", limitToMap(limits.DNS)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("extdisk_hdd", limitToMap(limits.ExtdiskHDD)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("extdisk_nvme", limitToMap(limits.ExtdiskNVMe)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("reserve_ip", limitToMap(limits.ReserveIP)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("gpu", limitToMap(limits.GPU)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

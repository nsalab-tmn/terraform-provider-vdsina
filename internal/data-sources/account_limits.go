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

func accountLimitsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	limits, err := c.GetAccountLimits(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read account limits: %w", err))
	}

	d.SetId("account_limits")

	d.Set("server", []map[string]interface{}{{
		"max":       limits.Server.Max,
		"now":       limits.Server.Now,
		"child_max": limits.Server.ChildMax,
	}})

	d.Set("server_ip4", []map[string]interface{}{{
		"max":       limits.ServerIP4.Max,
		"now":       limits.ServerIP4.Now,
		"child_max": limits.ServerIP4.ChildMax,
	}})

	d.Set("server_ip6", []map[string]interface{}{{
		"max":       limits.ServerIP6.Max,
		"now":       limits.ServerIP6.Now,
		"child_max": limits.ServerIP6.ChildMax,
	}})

	d.Set("iso", []map[string]interface{}{{
		"max":       limits.ISO.Max,
		"now":       limits.ISO.Now,
		"child_max": limits.ISO.ChildMax,
	}})

	d.Set("backup", []map[string]interface{}{{
		"max":       limits.Backup.Max,
		"now":       limits.Backup.Now,
		"child_max": limits.Backup.ChildMax,
	}})

	if limits.SSL != nil {
		d.Set("ssl", []map[string]interface{}{{
			"max":       limits.SSL.Max,
			"now":       limits.SSL.Now,
			"child_max": limits.SSL.ChildMax,
		}})
	} else {
		d.Set("ssl", []map[string]interface{}{})
	}

	if limits.Domain != nil {
		d.Set("domain", []map[string]interface{}{{
			"max":       limits.Domain.Max,
			"now":       limits.Domain.Now,
			"child_max": limits.Domain.ChildMax,
		}})
	} else {
		d.Set("domain", []map[string]interface{}{})
	}

	d.Set("dns", []map[string]interface{}{{
		"max":       limits.DNS.Max,
		"now":       limits.DNS.Now,
		"child_max": limits.DNS.ChildMax,
	}})

	d.Set("extdisk_hdd", []map[string]interface{}{{
		"max":       limits.ExtdiskHDD.Max,
		"now":       limits.ExtdiskHDD.Now,
		"child_max": limits.ExtdiskHDD.ChildMax,
	}})

	d.Set("extdisk_nvme", []map[string]interface{}{{
		"max":       limits.ExtdiskNVMe.Max,
		"now":       limits.ExtdiskNVMe.Now,
		"child_max": limits.ExtdiskNVMe.ChildMax,
	}})

	d.Set("reserve_ip", []map[string]interface{}{{
		"max":       limits.ReserveIP.Max,
		"now":       limits.ReserveIP.Now,
		"child_max": limits.ReserveIP.ChildMax,
	}})

	d.Set("gpu", []map[string]interface{}{{
		"max":       limits.GPU.Max,
		"now":       limits.GPU.Now,
		"child_max": limits.GPU.ChildMax,
	}})

	return nil
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/ips.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func IPsDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina IP addresses",

		ReadContext: ipsRead,

		Schema: map[string]*schema.Schema{
			"ips": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of IP addresses",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IP address ID",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address (IPv4 or IPv6)",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP type: '4' for IPv4, '6' for IPv6",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hostname associated with IP",
						},
						"gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway address",
						},
						"netmask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network mask",
						},
						"is_net": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this is a network (subnet)",
						},
						"datacenter_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Datacenter ID where IP is located",
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
					},
				},
			},
		},
	}
}

func ipsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	ips, err := c.GetIPs(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read IPs: %w", err))
	}

	ipList := make([]map[string]interface{}, len(ips))
	for i, ip := range ips {
		ipList[i] = map[string]interface{}{
			"id":                 ip.ID,
			"ip":                 ip.IP,
			"type":               ip.Type,
			"host":               ip.Host,
			"gateway":            ip.Gateway,
			"netmask":            ip.Netmask,
			"is_net":             ip.IsNet,
			"datacenter_id":      ip.Datacenter.ID,
			"datacenter_name":    ip.Datacenter.Name,
			"datacenter_country": ip.Datacenter.Country,
		}
	}

	d.SetId("ips")
	if err := d.Set("ips", ipList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set ips in state: %w", err))
	}

	return nil
}

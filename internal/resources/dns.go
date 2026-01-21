// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/dns.go
package resources

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

func DNSResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing DNS zones. Creates a DNS zone (domain) in VDSina DNS hosting.",

		CreateContext: dnsCreate,
		ReadContext:   dnsRead,
		DeleteContext: dnsDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain name for the DNS zone (e.g., example.com).",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "IP address for generating default A record. Optional.",
			},
			"full_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full DNS service name.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS zone status (new, active, block, notpaid, deleted).",
			},
			"status_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS zone status description.",
			},
			"real": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "DNS servers correctly set up in domain NS records.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS zone creation date.",
			},
			"updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS zone last update date.",
			},
			"end": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DNS zone service end date.",
			},
			"can_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can delete DNS zone.",
			},
		},
	}
}

func dnsCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	req := models.DNSZoneCreateRequest{
		Name: d.Get("name").(string),
	}
	if ip, ok := d.GetOk("ip"); ok {
		req.IP = ip.(string)
	}

	id, err := c.CreateDNSZone(ctx, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create DNS zone: %w", err))
	}

	d.SetId(strconv.Itoa(id))

	return dnsRead(ctx, d, meta)
}

func dnsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid DNS zone ID: %w", err))
	}

	zone, err := c.GetDNSZone(ctx, id)
	if err != nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("name", zone.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("full_name", zone.FullName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status", zone.Status); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status_text", zone.StatusText); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("real", zone.Real); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created", zone.Created); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated", zone.Updated); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("end", zone.End); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_delete", zone.Can.Delete); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func dnsDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid DNS zone ID: %w", err))
	}

	err = c.DeleteDNSZone(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete DNS zone %d: %w", id, err))
	}

	return nil
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/dns_record.go
package resources

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

func DNSRecordResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing DNS records within a DNS zone.",

		CreateContext: dnsRecordCreate,
		ReadContext:   dnsRecordRead,
		UpdateContext: dnsRecordUpdate,
		DeleteContext: dnsRecordDelete,

		Importer: &schema.ResourceImporter{
			StateContext: dnsRecordImportState,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "DNS zone ID where the record will be created.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Hostname for the record. Use '@' for root, '*' for wildcard, or subdomain name.",
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"A", "AAAA", "CNAME", "MX", "NS", "SRV", "CAA", "TXT",
				}, false),
				Description: "Record type: A, AAAA, CNAME, MX, NS, SRV, CAA, TXT.",
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Record value (IP address for A/AAAA, domain for CNAME/MX/NS, etc.).",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Priority for MX and SRV records (integer).",
			},
			"tag": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"issue", "issuewild", "iodef", "unknown", "",
				}, false),
				Description: "Tag for CAA records: issue, issuewild, iodef, unknown.",
			},
			"timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Record creation/update timestamp.",
			},
			"can_update": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Ability to update record.",
			},
			"can_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Ability to delete record.",
			},
		},
	}
}

func dnsRecordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	zoneID := d.Get("zone_id").(int)

	req := models.DNSRecordCreateRequest{
		Host:  d.Get("host").(string),
		Type:  d.Get("type").(string),
		Value: d.Get("value").(string),
	}

	if v, ok := d.GetOk("priority"); ok {
		priority := v.(int)
		req.Priority = &priority
	}
	if v, ok := d.GetOk("tag"); ok {
		tag := v.(string)
		req.Tag = &tag
	}

	recordID, err := c.CreateDNSRecord(ctx, zoneID, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create DNS record: %w", err))
	}

	d.SetId(fmt.Sprintf("%d:%d", zoneID, recordID))

	return dnsRecordRead(ctx, d, meta)
}

func dnsRecordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	zoneID, recordID, err := parseDNSRecordID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	var record *models.DNSRecord
	for i := 0; i < 5; i++ {
		record, err = c.GetDNSRecord(ctx, zoneID, recordID)
		if err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("zone_id", zoneID)
	d.Set("host", record.Host)
	d.Set("type", record.Type)
	d.Set("value", record.Value)
	d.Set("timestamp", record.Timestamp)
	d.Set("can_update", record.Can.Update)
	d.Set("can_delete", record.Can.Delete)

	if record.Priority != nil {
		d.Set("priority", *record.Priority)
	} else {
		d.Set("priority", 0)
	}
	if record.Tag != nil {
		d.Set("tag", *record.Tag)
	}

	return nil
}

func dnsRecordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	_, recordID, err := parseDNSRecordID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if d.HasChanges("value", "priority", "tag") {
		req := models.DNSRecordUpdateRequest{
			Value: d.Get("value").(string),
		}

		if v, ok := d.GetOk("priority"); ok {
			priority := v.(int)
			req.Priority = &priority
		}
		if v, ok := d.GetOk("tag"); ok {
			tag := v.(string)
			req.Tag = &tag
		}

		err := c.UpdateDNSRecord(ctx, recordID, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update DNS record %d: %w", recordID, err))
		}
	}

	return dnsRecordRead(ctx, d, meta)
}

func dnsRecordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	_, recordID, err := parseDNSRecordID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeleteDNSRecord(ctx, recordID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete DNS record %d: %w", recordID, err))
	}

	return nil
}

func parseDNSRecordID(id string) (int, int, error) {
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid DNS record ID format: %s (expected zoneID:recordID)", id)
	}

	zoneID, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid zone ID: %s", parts[0])
	}

	recordID, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid record ID: %s", parts[1])
	}

	return zoneID, recordID, nil
}

func dnsRecordImportState(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	zoneID, _, err := parseDNSRecordID(d.Id())
	if err != nil {
		return nil, err
	}

	d.Set("zone_id", zoneID)

	return []*schema.ResourceData{d}, nil
}

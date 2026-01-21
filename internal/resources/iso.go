// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/iso.go
package resources

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ISOResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing ISO images. ISO is downloaded from URL and stored in VDSina.",

		CreateContext: isoCreate,
		ReadContext:   isoRead,
		DeleteContext: isoDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "URL to download ISO from (http, https, ftp, ftps). Max 10GB.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ISO file name",
			},
			"full_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full ISO service name",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ISO status: new, active, block, notpaid, deleted",
			},
			"status_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status description",
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
		},
	}
}

func isoCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	url := d.Get("url").(string)

	key, err := c.StartISODownload(ctx, url)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to start ISO download: %w", err))
	}

	pollInterval := 5 * time.Second
	timeout := d.Timeout(schema.TimeoutCreate)
	deadline := time.Now().Add(timeout)

	for {
		if time.Now().After(deadline) {
			return diag.FromErr(fmt.Errorf("timeout waiting for ISO download to complete"))
		}

		status, description, err := c.GetISODownloadStatus(ctx, key)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to check ISO download status: %w", err))
		}

		switch status {
		case "done":
			goto createService
		case "error":
			return diag.FromErr(fmt.Errorf("ISO download failed: %s", description))
		case "processing":
			time.Sleep(pollInterval)
			continue
		default:
			return diag.FromErr(fmt.Errorf("unknown ISO download status: %s", status))
		}
	}

createService:
	id, err := c.CreateISOFromDownload(ctx, key)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create ISO service: %w", err))
	}

	d.SetId(strconv.Itoa(id))

	return isoRead(ctx, d, meta)
}

func isoRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid ISO ID: %s", d.Id()))
	}

	iso, err := c.GetISO(ctx, id)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", iso.Name)
	d.Set("full_name", iso.FullName)
	d.Set("status", iso.Status)
	d.Set("status_text", iso.StatusText)
	d.Set("created", iso.Created)
	d.Set("updated", iso.Updated)
	d.Set("end", iso.End)
	d.Set("file_size", iso.File.Size)
	d.Set("file_md5", iso.File.MD5)
	d.Set("attached", iso.Attached)

	return nil
}

func isoDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid ISO ID: %s", d.Id()))
	}

	iso, err := c.GetISO(ctx, id)
	if err != nil {
		d.SetId("")
		return nil
	}

	if iso.Attached {
		return diag.FromErr(fmt.Errorf("cannot delete ISO %d: it is attached to a server. Detach it first", id))
	}

	err = c.DeleteISO(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete ISO: %w", err))
	}

	d.SetId("")
	return nil
}

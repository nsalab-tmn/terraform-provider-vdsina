// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/ssh_key.go
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

func SSHKeyResource() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a VDSina SSH key",

		CreateContext: sshKeyCreate,
		ReadContext:   sshKeyRead,
		UpdateContext: sshKeyUpdate,
		DeleteContext: sshKeyDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SSH key name",
			},
			"data": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Public SSH key (e.g., contents of ~/.ssh/id_rsa.pub)",
			},
		},
	}
}

func sshKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	req := models.SSHKeyCreateRequest{
		Name: d.Get("name").(string),
		Data: d.Get("data").(string),
	}

	id, err := c.CreateSSHKey(ctx, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create SSH key: %w", err))
	}

	d.SetId(strconv.Itoa(id))

	return sshKeyRead(ctx, d, meta)
}

func sshKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid SSH key ID: %s", d.Id()))
	}

	key, err := c.GetSSHKey(ctx, id)
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", key.Name)
	if key.Data != "" {
		d.Set("data", key.Data)
	}

	return nil
}

func sshKeyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid SSH key ID: %s", d.Id()))
	}

	if d.HasChanges("name", "data") {
		req := models.SSHKeyUpdateRequest{}

		if d.HasChange("name") {
			req.Name = d.Get("name").(string)
		}
		if d.HasChange("data") {
			req.Data = d.Get("data").(string)
		}

		err := c.UpdateSSHKey(ctx, id, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update SSH key: %w", err))
		}
	}

	return sshKeyRead(ctx, d, meta)
}

func sshKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid SSH key ID: %s", d.Id()))
	}

	err = c.DeleteSSHKey(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete SSH key: %w", err))
	}

	d.SetId("")
	return nil
}

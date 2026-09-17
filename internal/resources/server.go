// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/server.go
package resources

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

func ServerResource() *schema.Resource {
	return &schema.Resource{
		Description: "Manages a VDSina VPS server",

		CreateContext: serverCreate,
		ReadContext:   serverRead,
		UpdateContext: serverUpdate,
		DeleteContext: serverDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"datacenter": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Datacenter ID where the server will be created",
			},
			"server_plan": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server plan ID (tariff)",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Server name (can be changed after creation)",
			},
			"template": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				Description:   "OS template ID (mutually exclusive with backup, iso)",
				ConflictsWith: []string{"backup", "iso"},
			},
			"ssh_key": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "SSH key ID to add to the server",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Hostname for the server",
			},
			"backup": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				Description:   "Backup ID to restore from (mutually exclusive with template, iso)",
				ConflictsWith: []string{"template", "iso"},
			},
			"iso": {
				Type:          schema.TypeInt,
				Optional:      true,
				ForceNew:      true,
				Description:   "ISO ID to install from (mutually exclusive with template, backup)",
				ConflictsWith: []string{"template", "backup"},
			},
			"cpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of CPU cores (for constructor plans)",
			},
			"ram": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "RAM in GB (for constructor plans)",
			},
			"disk": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Disk size in GB (for constructor plans)",
			},
			"gpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of GPUs (for GPU plans)",
			},
			"autoprolong": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable auto-renewal when balance is sufficient",
			},
			"full_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full server name including plan info",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server status (new, active, block, deleted)",
			},
			"status_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status description",
			},
			"ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public IP address",
			},
			"ip_local": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Local/private IP address",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation date",
			},
			"end": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Expiration date",
			},
			"datacenter_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Datacenter name",
			},
			"template_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS template name",
			},
			"plan_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server plan name",
			},
		},
	}
}

func serverCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	req := models.ServerCreateRequest{
		Datacenter: d.Get("datacenter").(int),
		ServerPlan: d.Get("server_plan").(int),
	}

	if v, ok := d.GetOk("template"); ok {
		req.Template = v.(int)
	}
	if v, ok := d.GetOk("ssh_key"); ok {
		req.SSHKey = v.(int)
	}
	if v, ok := d.GetOk("host"); ok {
		req.Host = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		req.Name = v.(string)
	}
	if v, ok := d.GetOk("backup"); ok {
		val := v.(int)
		req.Backup = &val
	}
	if v, ok := d.GetOk("iso"); ok {
		val := v.(int)
		req.ISO = &val
	}
	if v, ok := d.GetOk("cpu"); ok {
		req.CPU = v.(int)
	}
	if v, ok := d.GetOk("ram"); ok {
		req.RAM = v.(int)
	}
	if v, ok := d.GetOk("disk"); ok {
		req.Disk = v.(int)
	}
	if v, ok := d.GetOk("gpu"); ok {
		req.GPU = v.(int)
	}

	id, err := c.CreateServer(ctx, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create server: %w", err))
	}

	d.SetId(strconv.Itoa(id))

	autoprolong := d.Get("autoprolong").(bool)

	if !autoprolong {
		updateReq := models.ServerUpdateRequest{
			Autoprolong: "0",
		}
		time.Sleep(2 * time.Second)
		_ = c.UpdateServer(ctx, id, updateReq)
	}

	if err := waitForServerReady(ctx, c, id, d.Timeout(schema.TimeoutCreate)); err != nil {
		return diag.FromErr(err)
	}

	return serverRead(ctx, d, meta)
}

func serverRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid server ID: %s", d.Id()))
	}

	server, err := c.GetServer(ctx, id)
	if err != nil {
		if client.IsNotFound(err) {
			log.Printf("[WARN] VDSina server %d not found, removing from state", id)
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("failed to read server %d: %w", id, err))
	}

	if server.Status == "deleted" {
		log.Printf("[WARN] VDSina server %d is deleted, removing from state", id)
		d.SetId("")
		return nil
	}

	if err := d.Set("name", server.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("full_name", server.FullName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status", server.Status); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status_text", server.StatusText); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created", server.Created); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("end", server.End); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("host", server.Host); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("autoprolong", server.Autoprolong); err != nil {
		return diag.FromErr(err)
	}

	if server.IP.IP != "" {
		if err := d.Set("ip", server.IP.IP); err != nil {
			return diag.FromErr(err)
		}
	}
	if server.IPLocal.IP != "" {
		if err := d.Set("ip_local", server.IPLocal.IP); err != nil {
			return diag.FromErr(err)
		}
	}

	if err := d.Set("datacenter", server.Datacenter.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("datacenter_name", server.Datacenter.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("server_plan", server.ServerPlan.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("plan_name", server.ServerPlan.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template", server.Template.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("template_name", server.Template.Name); err != nil {
		return diag.FromErr(err)
	}

	if server.SSHKey != nil {
		if err := d.Set("ssh_key", server.SSHKey.ID); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid server ID: %s", d.Id()))
	}

	if d.HasChanges("name", "autoprolong") {
		req := models.ServerUpdateRequest{}

		if d.HasChange("name") {
			req.Name = d.Get("name").(string)
		}
		if d.HasChange("autoprolong") {
			if d.Get("autoprolong").(bool) {
				req.Autoprolong = "1"
			} else {
				req.Autoprolong = "0"
			}
		}

		err := c.UpdateServer(ctx, id, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update server: %w", err))
		}
	}

	return serverRead(ctx, d, meta)
}

func serverDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid server ID: %s", d.Id()))
	}

	err = c.DeleteServer(ctx, id)
	if err != nil {
		if client.IsNotFound(err) {
			d.SetId("")
			return nil
		}
		return diag.FromErr(fmt.Errorf("failed to delete server: %w", err))
	}

	if err := waitForServerDeleted(ctx, c, id, d.Timeout(schema.TimeoutDelete)); err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

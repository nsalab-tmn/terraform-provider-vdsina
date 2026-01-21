// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/server_actions.go
package resources

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func ServerRebootResource() *schema.Resource {
	return &schema.Resource{
		Description: "Reboots a VDSina server. The reboot is triggered on resource creation and when triggers change.",

		CreateContext: serverRebootCreate,
		ReadContext:   serverRebootRead,
		UpdateContext: serverRebootUpdate,
		DeleteContext: serverRebootDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to reboot",
			},
			"type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "soft",
				Description:  "Reboot type: 'soft' (graceful) or 'hard' (force)",
				ValidateFunc: validation.StringInSlice([]string{"soft", "hard"}, false),
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger a reboot",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_reboot": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last reboot",
			},
		},
	}
}

func serverRebootCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)
	rebootType := d.Get("type").(string)

	err := c.RebootServer(ctx, serverID, rebootType)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to reboot server: %w", err))
	}

	d.SetId(strconv.Itoa(serverID))
	if err := d.Set("last_reboot", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverRebootRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func serverRebootUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("triggers") || d.HasChange("type") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)
		rebootType := d.Get("type").(string)

		err := c.RebootServer(ctx, serverID, rebootType)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to reboot server: %w", err))
		}

		if err := d.Set("last_reboot", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverRebootDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func ServerReinstallResource() *schema.Resource {
	return &schema.Resource{
		Description: "Reinstalls OS on a VDSina server. WARNING: This will erase all data on the server!",

		CreateContext: serverReinstallCreate,
		ReadContext:   serverReinstallRead,
		UpdateContext: serverReinstallUpdate,
		DeleteContext: serverReinstallDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to reinstall",
			},
			"template": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "OS template ID for reinstallation",
			},
			"ssh_key": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "SSH key ID to add during reinstallation",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hostname to set during reinstallation",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger a reinstall",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_reinstall": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last reinstallation",
			},
		},
	}
}

func serverReinstallCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)
	templateID := d.Get("template").(int)
	sshKeyID := d.Get("ssh_key").(int)
	host := d.Get("host").(string)

	err := c.ReinstallServer(ctx, serverID, templateID, sshKeyID, host)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to reinstall server: %w", err))
	}

	d.SetId(strconv.Itoa(serverID))
	if err := d.Set("last_reinstall", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverReinstallRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func serverReinstallUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChanges("triggers", "template", "ssh_key", "host") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)
		templateID := d.Get("template").(int)
		sshKeyID := d.Get("ssh_key").(int)
		host := d.Get("host").(string)

		err := c.ReinstallServer(ctx, serverID, templateID, sshKeyID, host)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to reinstall server: %w", err))
		}

		if err := d.Set("last_reinstall", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverReinstallDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func ServerPasswordResource() *schema.Resource {
	return &schema.Resource{
		Description: "Manages root password for a VDSina server. Can retrieve current password or set a new one.",

		CreateContext: serverPasswordCreate,
		ReadContext:   serverPasswordRead,
		UpdateContext: serverPasswordUpdate,
		DeleteContext: serverPasswordDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Root password. If not set, retrieves current password. If set, changes password to the specified value.",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger password reset",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last password update/retrieval",
			},
		},
	}
}

func serverPasswordCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)

	if v, ok := d.GetOk("password"); ok && v.(string) != "" {
		err := c.SetServerPassword(ctx, serverID, v.(string))
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to set server password: %w", err))
		}
	} else {
		password, err := c.GetServerPassword(ctx, serverID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to get server password: %w", err))
		}
		if err := d.Set("password", password); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.Itoa(serverID))
	if err := d.Set("last_updated", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverPasswordRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid server ID: %s", d.Id()))
	}

	password, err := c.GetServerPassword(ctx, serverID)
	if err != nil {
		d.SetId("")
		return nil
	}

	if err := d.Set("password", password); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverPasswordUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChanges("password", "triggers") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)
		password := d.Get("password").(string)

		if password != "" {
			err := c.SetServerPassword(ctx, serverID, password)
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to set server password: %w", err))
			}
		} else {
			err := c.SetServerPassword(ctx, serverID, "")
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to reset server password: %w", err))
			}
			newPassword, err := c.GetServerPassword(ctx, serverID)
			if err != nil {
				return diag.FromErr(fmt.Errorf("failed to get new server password: %w", err))
			}
			if err := d.Set("password", newPassword); err != nil {
				return diag.FromErr(err)
			}
		}

		if err := d.Set("last_updated", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverPasswordDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func ServerPlanChangeResource() *schema.Resource {
	return &schema.Resource{
		Description: "Changes the tariff plan of a VDSina server. Only upgrade is supported (downgrade is not possible). The server will be restarted.",

		CreateContext: serverPlanChangeCreate,
		ReadContext:   serverPlanChangeRead,
		UpdateContext: serverPlanChangeUpdate,
		DeleteContext: serverPlanChangeDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to change plan",
			},
			"server_plan": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "New tariff plan ID (must be from the same tariff group, only upgrade)",
			},
			"cpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "New CPU count (for constructor tariffs)",
			},
			"ram": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "New RAM amount in GB (for constructor tariffs)",
			},
			"disk": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "New storage amount in GB (cannot be smaller than current, for constructor tariffs)",
			},
			"gpu": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "New GPU count (for constructor tariffs)",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger a plan change",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_changed": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last plan change",
			},
		},
	}
}

func serverPlanChangeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)

	req := client.ServerPlanChangeRequest{
		ServerPlan: d.Get("server_plan").(int),
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

	err := c.ChangeServerPlan(ctx, serverID, req)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to change server plan: %w", err))
	}

	d.SetId(strconv.Itoa(serverID))
	if err := d.Set("last_changed", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverPlanChangeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func serverPlanChangeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChanges("server_plan", "cpu", "ram", "disk", "gpu", "triggers") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)

		req := client.ServerPlanChangeRequest{
			ServerPlan: d.Get("server_plan").(int),
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

		err := c.ChangeServerPlan(ctx, serverID, req)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to change server plan: %w", err))
		}

		if err := d.Set("last_changed", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverPlanChangeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func ServerProlongResource() *schema.Resource {
	return &schema.Resource{
		Description: "Prolongs and starts a VDSina server. Extends the paid period and starts the server if it was stopped (e.g., due to disabled auto-renewal).",

		CreateContext: serverProlongCreate,
		ReadContext:   serverProlongRead,
		UpdateContext: serverProlongUpdate,
		DeleteContext: serverProlongDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to prolong and start",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger a prolong",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_prolonged": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last prolong action",
			},
		},
	}
}

func serverProlongCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)

	err := c.ProlongServer(ctx, serverID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to prolong server: %w", err))
	}

	d.SetId(strconv.Itoa(serverID))
	if err := d.Set("last_prolonged", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverProlongRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func serverProlongUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("triggers") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)

		err := c.ProlongServer(ctx, serverID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to prolong server: %w", err))
		}

		if err := d.Set("last_prolonged", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverProlongDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

func ServerISOResource() *schema.Resource {
	return &schema.Resource{
		Description: "Attaches an ISO image to a VDSina server. WARNING: The server will be restarted!",

		CreateContext: serverISOCreate,
		ReadContext:   serverISORead,
		UpdateContext: serverISOUpdate,
		DeleteContext: serverISODelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to attach ISO to",
			},
			"iso_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "ISO service ID to attach",
			},
			"last_attached": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last ISO attach",
			},
		},
	}
}

func serverISOCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)
	isoID := d.Get("iso_id").(int)

	err := c.AttachISO(ctx, serverID, isoID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to attach ISO: %w", err))
	}

	d.SetId(fmt.Sprintf("%d:%d", serverID, isoID))
	if err := d.Set("last_attached", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func serverISORead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	serverID := d.Get("server_id").(int)

	_, err := c.GetServer(ctx, serverID)
	if err != nil {
		d.SetId("")
		return nil
	}

	return nil
}

func serverISOUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("iso_id") {
		c := meta.(*client.Client)

		serverID := d.Get("server_id").(int)
		isoID := d.Get("iso_id").(int)

		err := c.AttachISO(ctx, serverID, isoID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to attach ISO: %w", err))
		}

		d.SetId(fmt.Sprintf("%d:%d", serverID, isoID))
		if err := d.Set("last_attached", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func serverISODelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	serverID := d.Get("server_id").(int)

	err := c.DetachISO(ctx, serverID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to detach ISO: %w", err))
	}

	d.SetId("")
	return nil
}

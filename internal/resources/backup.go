// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/backup.go
package resources

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

func BackupResource() *schema.Resource {
	return &schema.Resource{
		Description: "Resource for managing server backups. Creates a backup of the specified server.",

		CreateContext: backupCreate,
		ReadContext:   backupRead,
		UpdateContext: backupUpdate,
		DeleteContext: backupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "ID of the server to backup.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Backup name. If not specified, server name will be used.",
			},
			"autoprolong": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable automatic prolongation of backup service.",
			},
			"full_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full backup service name.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backup status (new, active, block, notpaid, deleted).",
			},
			"status_text": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backup status description.",
			},
			"datacenter_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Datacenter ID where backup is stored.",
			},
			"created": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backup creation date.",
			},
			"updated": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backup last update date.",
			},
			"end": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Backup service end date.",
			},
			"can_update": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can update backup.",
			},
			"can_prolong": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can prolong backup.",
			},
			"can_delete": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Can delete backup.",
			},
		},
	}
}

func backupCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	serverID := d.Get("server_id").(int)

	backup, err := c.CreateBackup(ctx, serverID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to create backup: %w", err))
	}

	d.SetId(strconv.Itoa(backup.ID))

	stateConf := &retry.StateChangeConf{
		Pending:    []string{"new"},
		Target:     []string{"active"},
		Refresh:    backupStatusRefreshFunc(ctx, c, backup.ID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForStateContext(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("error waiting for backup to become active: %w", err))
	}

	if name, ok := d.GetOk("name"); ok {
		autoprolong := "1"
		if !d.Get("autoprolong").(bool) {
			autoprolong = "0"
		}
		err = c.UpdateBackup(ctx, backup.ID, models.BackupUpdateRequest{
			Name:        name.(string),
			Autoprolong: autoprolong,
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update backup name: %w", err))
		}
	}

	return backupRead(ctx, d, meta)
}

func backupRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid backup ID: %w", err))
	}

	backup, err := c.GetBackup(ctx, id)
	if err != nil {
		d.SetId("")
		return nil
	}

	if backup.Server != nil {
		if err := d.Set("server_id", backup.Server.ID); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("name", backup.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("full_name", backup.FullName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status", backup.Status); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("status_text", backup.StatusText); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("datacenter_id", backup.Datacenter.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created", backup.Created); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("updated", backup.Updated); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("end", backup.End); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_update", backup.Can.Update); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_prolong", backup.Can.Prolong); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("can_delete", backup.Can.Delete); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func backupUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid backup ID: %w", err))
	}

	if d.HasChanges("name", "autoprolong") {
		autoprolong := "1"
		if !d.Get("autoprolong").(bool) {
			autoprolong = "0"
		}

		err = c.UpdateBackup(ctx, id, models.BackupUpdateRequest{
			Name:        d.Get("name").(string),
			Autoprolong: autoprolong,
		})
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to update backup %d: %w", id, err))
		}
	}

	return backupRead(ctx, d, meta)
}

func backupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(fmt.Errorf("invalid backup ID: %w", err))
	}

	err = c.DeleteBackup(ctx, id)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to delete backup %d: %w", id, err))
	}

	return nil
}

func backupStatusRefreshFunc(ctx context.Context, c *client.Client, id int) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		backup, err := c.GetBackup(ctx, id)
		if err != nil {
			return nil, "", err
		}
		return backup, backup.Status, nil
	}
}

func BackupRestoreResource() *schema.Resource {
	return &schema.Resource{
		Description: "Restores a VDSina backup to a server. WARNING: This will overwrite all data on the server!",

		CreateContext: backupRestoreCreate,
		ReadContext:   backupRestoreRead,
		UpdateContext: backupRestoreUpdate,
		DeleteContext: backupRestoreDelete,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Backup ID to restore from",
			},
			"server_id": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Server ID to restore to (must be in the same datacenter as backup)",
			},
			"triggers": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Map of arbitrary keys and values that, when changed, will trigger a restore",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"last_restored": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last restore action",
			},
		},
	}
}

func backupRestoreCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	backupID := d.Get("backup_id").(int)
	serverID := d.Get("server_id").(int)

	err := c.RestoreBackup(ctx, backupID, serverID)
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to restore backup: %w", err))
	}

	d.SetId(fmt.Sprintf("%d:%d", backupID, serverID))
	if err := d.Set("last_restored", time.Now().Format(time.RFC3339)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func backupRestoreRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func backupRestoreUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("triggers") {
		c := meta.(*client.Client)

		backupID := d.Get("backup_id").(int)
		serverID := d.Get("server_id").(int)

		err := c.RestoreBackup(ctx, backupID, serverID)
		if err != nil {
			return diag.FromErr(fmt.Errorf("failed to restore backup: %w", err))
		}

		if err := d.Set("last_restored", time.Now().Format(time.RFC3339)); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func backupRestoreDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}

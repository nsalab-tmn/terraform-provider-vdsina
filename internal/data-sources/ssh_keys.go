// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/ssh_keys.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func SSHKeysDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving list of VDSina SSH keys",

		ReadContext: sshKeysRead,

		Schema: map[string]*schema.Schema{
			"ssh_keys": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of SSH keys",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "SSH key ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSH key name",
						},
					},
				},
			},
		},
	}
}

func sshKeysRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	keys, err := c.GetSSHKeys(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read SSH keys: %w", err))
	}

	keyList := make([]map[string]interface{}, len(keys))
	for i, k := range keys {
		keyList[i] = map[string]interface{}{
			"id":   k.ID,
			"name": k.Name,
		}
	}

	d.SetId("ssh_keys")
	if err := d.Set("ssh_keys", keyList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set ssh_keys in state: %w", err))
	}

	return nil
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/data-sources/templates.go
package datasources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

func TemplatesDataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Data source for retrieving VDSina OS templates",

		ReadContext: templatesRead,

		Schema: map[string]*schema.Schema{
			"templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of OS templates",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Template ID",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template name (e.g., 'Ubuntu 24.04', 'Debian 12')",
						},
						"active": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the template is active",
						},
						"ssh_key": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether SSH key injection is supported",
						},
						"template_group": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Template group ID for categorization",
						},
						"server_plans": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of compatible server plan IDs",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"cpu_min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum required CPU cores",
						},
						"ram_min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum required RAM in GB",
						},
						"disk_min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Minimum required disk size in GB",
						},
					},
				},
			},
		},
	}
}

func templatesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	templates, err := c.GetTemplates(ctx)
	if err != nil {
		return diag.FromErr(fmt.Errorf("unable to read templates: %w", err))
	}

	templateList := make([]map[string]interface{}, len(templates))
	for i, t := range templates {
		templateList[i] = map[string]interface{}{
			"id":             t.ID,
			"name":           t.Name,
			"active":         t.Active,
			"ssh_key":        t.SSHKey,
			"template_group": t.TemplateGroup,
			"server_plans":   t.ServerPlans,
			"cpu_min":        t.Limits.CPU.Min,
			"ram_min":        t.Limits.RAM.Min,
			"disk_min":       t.Limits.Disk.Min,
		}
	}

	d.SetId("templates")
	if err := d.Set("templates", templateList); err != nil {
		return diag.FromErr(fmt.Errorf("unable to set templates in state: %w", err))
	}

	return nil
}

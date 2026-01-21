// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/provider/provider.go
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
	datasources "github.com/scinfra-pro/terraform-provider-vdsina/internal/data-sources"
	"github.com/scinfra-pro/terraform-provider-vdsina/internal/resources"
)

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VDSINA_API_TOKEN", nil),
				Description: "API Token for VDSina provider. Can be set via VDSINA_API_TOKEN env var.",
				Sensitive:   true,
			},
			"base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VDSINA_BASE_URL", "https://userapi.vdsina.ru/v1"),
				Description: "Base URL for VDSina API. Defaults to https://userapi.vdsina.ru/v1",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"vdsina_server":               resources.ServerResource(),
			"vdsina_ssh_key":              resources.SSHKeyResource(),
			"vdsina_server_reboot":        resources.ServerRebootResource(),
			"vdsina_server_reinstall":     resources.ServerReinstallResource(),
			"vdsina_server_password":      resources.ServerPasswordResource(),
			"vdsina_server_plan_change":   resources.ServerPlanChangeResource(),
			"vdsina_server_start_prolong": resources.ServerProlongResource(),
			"vdsina_server_iso":           resources.ServerISOResource(),
			"vdsina_iso":                  resources.ISOResource(),
			"vdsina_backup":               resources.BackupResource(),
			"vdsina_backup_restore":       resources.BackupRestoreResource(),
			"vdsina_dns":                  resources.DNSResource(),
			"vdsina_dns_record":           resources.DNSRecordResource(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"vdsina_datacenters":     datasources.DatacentersDataSource(),
			"vdsina_server_groups":   datasources.ServerGroupsDataSource(),
			"vdsina_server_plans":    datasources.ServerPlansDataSource(),
			"vdsina_templates":       datasources.TemplatesDataSource(),
			"vdsina_servers":         datasources.ServersDataSource(),
			"vdsina_ssh_keys":        datasources.SSHKeysDataSource(),
			"vdsina_account":         datasources.AccountDataSource(),
			"vdsina_account_balance": datasources.AccountBalanceDataSource(),
			"vdsina_account_limits":  datasources.AccountLimitsDataSource(),
			"vdsina_iso_list":        datasources.ISOListDataSource(),
			"vdsina_backups":         datasources.BackupsDataSource(),
			"vdsina_dns_zones":       datasources.DNSZonesDataSource(),
			"vdsina_ips":             datasources.IPsDataSource(),
		},

		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	apiToken := d.Get("api_token").(string)
	baseURL := d.Get("base_url").(string)

	var diags diag.Diagnostics

	if apiToken == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "API Token is required",
			Detail:   "The api_token must be provided. Set it in the provider configuration or via VDSINA_API_TOKEN environment variable.",
		})
		return nil, diags
	}

	c, err := client.NewClient(baseURL, apiToken)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create VDSina client",
			Detail:   err.Error(),
		})
		return nil, diags
	}

	return c, diags
}

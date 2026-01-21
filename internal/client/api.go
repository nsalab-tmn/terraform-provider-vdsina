// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/api.go
package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

// GetDatacenters returns list of datacenters
// API: GET /datacenter
func (c *Client) GetDatacenters(ctx context.Context) ([]models.Datacenter, error) {
	var response models.DatacentersResponse

	err := c.Get(ctx, "/datacenter", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get datacenters: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetServerGroups returns list of server groups
// API: GET /server-group
func (c *Client) GetServerGroups(ctx context.Context) ([]models.ServerGroup, error) {
	var response models.ServerGroupsResponse

	err := c.Get(ctx, "/server-group", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get server groups: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetServerPlans returns list of server plans for a group
// API: GET /server-plan/{groupID}
func (c *Client) GetServerPlans(ctx context.Context, groupID int) ([]models.ServerPlan, error) {
	var response models.ServerPlansResponse

	path := fmt.Sprintf("/server-plan/%d", groupID)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get server plans for group %d: %w", groupID, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetTemplates returns list of OS templates
// API: GET /template
func (c *Client) GetTemplates(ctx context.Context) ([]models.Template, error) {
	var response models.TemplatesResponse

	err := c.Get(ctx, "/template", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get templates: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// CreateServer creates a new server
// API: POST /server
func (c *Client) CreateServer(ctx context.Context, req models.ServerCreateRequest) (int, error) {
	var response models.ServerCreateResponse

	err := c.Post(ctx, "/server", req, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create server: %w", err)
	}

	if response.Status != "ok" {
		return 0, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// GetServer returns server information
// API: GET /server/{id}
func (c *Client) GetServer(ctx context.Context, id int) (*models.Server, error) {
	var response models.ServerResponse

	path := fmt.Sprintf("/server/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get server %d: %w", id, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// UpdateServer updates server settings
// API: PUT /server/{id}
func (c *Client) UpdateServer(ctx context.Context, id int, req models.ServerUpdateRequest) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server/%d", id)
	err := c.Put(ctx, path, req, &response)
	if err != nil {
		return fmt.Errorf("failed to update server %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// DeleteServer deletes a server
// API: DELETE /server/{id}
func (c *Client) DeleteServer(ctx context.Context, id int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server/%d", id)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete server %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetServers returns list of all servers
// API: GET /server
func (c *Client) GetServers(ctx context.Context) ([]models.Server, error) {
	var response models.ServersListResponse

	err := c.Get(ctx, "/server", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get servers: %w", err)
	}

	if response.Status != "ok" {
		if response.StatusMsg == "No Server information" {
			return []models.Server{}, nil
		}
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetSSHKeys returns list of SSH keys
// API: GET /ssh-key
func (c *Client) GetSSHKeys(ctx context.Context) ([]models.SSHKey, error) {
	var response models.SSHKeysResponse

	err := c.Get(ctx, "/ssh-key", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get SSH keys: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetSSHKey returns SSH key by ID
// API: GET /ssh-key/{id}
func (c *Client) GetSSHKey(ctx context.Context, id int) (*models.SSHKey, error) {
	var response models.SSHKeyResponse

	path := fmt.Sprintf("/ssh-key/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get SSH key %d: %w", id, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// CreateSSHKey creates a new SSH key
// API: POST /ssh-key
func (c *Client) CreateSSHKey(ctx context.Context, req models.SSHKeyCreateRequest) (int, error) {
	var response models.SSHKeyCreateResponse

	err := c.Post(ctx, "/ssh-key", req, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create SSH key: %w", err)
	}

	if response.Status != "ok" {
		return 0, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// UpdateSSHKey updates SSH key
// API: PUT /ssh-key/{id}
func (c *Client) UpdateSSHKey(ctx context.Context, id int, req models.SSHKeyUpdateRequest) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/ssh-key/%d", id)
	err := c.Put(ctx, path, req, &response)
	if err != nil {
		return fmt.Errorf("failed to update SSH key %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// DeleteSSHKey deletes SSH key
// API: DELETE /ssh-key/{id}
func (c *Client) DeleteSSHKey(ctx context.Context, id int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/ssh-key/%d", id)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete SSH key %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetAccount returns account information
// API: GET /account
func (c *Client) GetAccount(ctx context.Context) (*models.Account, error) {
	var response models.AccountResponse

	err := c.Get(ctx, "/account", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// GetAccountBalance returns account balance
// API: GET /account.balance
func (c *Client) GetAccountBalance(ctx context.Context) (*models.AccountBalance, error) {
	var response models.AccountBalanceResponse

	err := c.Get(ctx, "/account.balance", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get account balance: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// GetAccountLimits returns account limits by service types
// API: GET /account.limit
func (c *Client) GetAccountLimits(ctx context.Context) (*models.AccountLimits, error) {
	var response models.AccountLimitsResponse

	err := c.Get(ctx, "/account.limit", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get account limits: %w", err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	limits := &models.AccountLimits{
		Server:      response.Data.Server,
		ServerIP4:   response.Data.ServerIP4,
		ServerIP6:   response.Data.ServerIP6,
		ISO:         response.Data.ISO,
		Backup:      response.Data.Backup,
		DNS:         response.Data.DNS,
		ExtdiskHDD:  response.Data.ExtdiskHDD,
		ExtdiskNVMe: response.Data.ExtdiskNVMe,
		ReserveIP:   response.Data.ReserveIP,
		GPU:         response.Data.GPU,
	}

	if len(response.Data.SSL) > 0 && response.Data.SSL[0] == '{' {
		var ssl models.LimitInfo
		if err := json.Unmarshal(response.Data.SSL, &ssl); err == nil {
			limits.SSL = &ssl
		}
	}

	if len(response.Data.Domain) > 0 && response.Data.Domain[0] == '{' {
		var domain models.LimitInfo
		if err := json.Unmarshal(response.Data.Domain, &domain); err == nil {
			limits.Domain = &domain
		}
	}

	return limits, nil
}

// RebootServer reboots a server
// API: PUT /server.reboot/{id}
func (c *Client) RebootServer(ctx context.Context, id int, rebootType string) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.reboot/%d", id)
	body := map[string]string{"type": rebootType}

	err := c.Put(ctx, path, body, &response)
	if err != nil {
		return fmt.Errorf("failed to reboot server %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// ReinstallServer reinstalls OS on server
// API: PUT /server.reinstall/{id}
func (c *Client) ReinstallServer(ctx context.Context, id int, templateID int, sshKeyID int, host string) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.reinstall/%d", id)
	body := map[string]interface{}{
		"template": templateID,
	}
	if sshKeyID > 0 {
		body["ssh-key"] = sshKeyID
	}
	if host != "" {
		body["host"] = host
	}

	err := c.Put(ctx, path, body, &response)
	if err != nil {
		return fmt.Errorf("failed to reinstall server %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetServerPassword gets server password
// API: GET /server.password/{id}
func (c *Client) GetServerPassword(ctx context.Context, id int) (string, error) {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
		Data      struct {
			Password string `json:"password"`
		} `json:"data"`
	}

	path := fmt.Sprintf("/server.password/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return "", fmt.Errorf("failed to get server password %d: %w", id, err)
	}

	if response.Status != "ok" {
		return "", fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.Password, nil
}

// SetServerPassword sets new server password
// API: PUT /server.password/{id}
func (c *Client) SetServerPassword(ctx context.Context, id int, password string) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.password/%d", id)
	body := map[string]string{}
	if password != "" {
		body["password"] = password
	}

	err := c.Put(ctx, path, body, &response)
	if err != nil {
		return fmt.Errorf("failed to set server password %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetISOs returns list of ISO images
// API: GET /iso
func (c *Client) GetISOs(ctx context.Context) ([]models.ISO, error) {
	var response models.ISOListResponse

	err := c.Get(ctx, "/iso", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get ISOs: %w", err)
	}

	if response.Status != "ok" {
		if response.StatusMsg == "No ISO information" {
			return []models.ISO{}, nil
		}
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	if response.Data == nil {
		return []models.ISO{}, nil
	}

	return response.Data, nil
}

// GetISO returns ISO by ID
// API: GET /iso/{id}
func (c *Client) GetISO(ctx context.Context, id int) (*models.ISO, error) {
	var response models.ISOResponse

	path := fmt.Sprintf("/iso/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get ISO %d: %w", id, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// StartISODownload starts ISO download by URL
// API: POST /iso
func (c *Client) StartISODownload(ctx context.Context, url string) (string, error) {
	var response models.ISODownloadResponse

	body := map[string]string{"url": url}
	err := c.Post(ctx, "/iso", body, &response)
	if err != nil {
		return "", fmt.Errorf("failed to start ISO download: %w", err)
	}

	if response.Status != "ok" {
		return "", fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// GetISODownloadStatus checks ISO download status
// API: GET /iso/{KEY}
func (c *Client) GetISODownloadStatus(ctx context.Context, key string) (string, string, error) {
	var response models.ISODownloadStatusResponse

	path := fmt.Sprintf("/iso/%s", key)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return "", "", fmt.Errorf("failed to get ISO download status: %w", err)
	}

	if response.Status != "ok" {
		return "", "", fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.Status, response.Data.Description, nil
}

// CreateISOFromDownload creates ISO service from downloaded file
// API: POST /iso/{KEY}
func (c *Client) CreateISOFromDownload(ctx context.Context, key string) (int, error) {
	var response models.ISOCreateFromKeyResponse

	path := fmt.Sprintf("/iso/%s", key)
	err := c.Post(ctx, path, nil, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create ISO from download: %w", err)
	}

	if response.Status != "ok" {
		return 0, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// DeleteISO deletes ISO
// API: DELETE /iso/{id}
func (c *Client) DeleteISO(ctx context.Context, id int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/iso/%d", id)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete ISO %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetBackups returns list of backups
// API: GET /backup
func (c *Client) GetBackups(ctx context.Context) ([]models.Backup, error) {
	var response models.BackupListResponse

	err := c.Get(ctx, "/backup", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get backups: %w", err)
	}

	if response.Status != "ok" {
		if response.StatusMsg == "No Backup information" {
			return []models.Backup{}, nil
		}
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	if response.Data == nil {
		return []models.Backup{}, nil
	}

	return response.Data, nil
}

// GetBackup returns backup by ID
// API: GET /backup/{id}
func (c *Client) GetBackup(ctx context.Context, id int) (*models.Backup, error) {
	var response models.BackupResponse

	path := fmt.Sprintf("/backup/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get backup %d: %w", id, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// CreateBackup creates server backup
// API: POST /backup/{serverID}
func (c *Client) CreateBackup(ctx context.Context, serverID int) (*models.Backup, error) {
	var response models.BackupCreateResponse

	path := fmt.Sprintf("/backup/%d", serverID)
	err := c.Post(ctx, path, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create backup for server %d: %w", serverID, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	backups, err := c.GetBackups(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get backups after creation: %w", err)
	}

	for _, backup := range backups {
		if backup.Server != nil && backup.Server.ID == serverID && backup.Status == "new" {
			return &backup, nil
		}
	}

	return nil, fmt.Errorf("backup created but could not find it in the list")
}

// UpdateBackup updates backup
// API: PUT /backup/{backupID}
func (c *Client) UpdateBackup(ctx context.Context, id int, req models.BackupUpdateRequest) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/backup/%d", id)
	err := c.Put(ctx, path, req, &response)
	if err != nil {
		return fmt.Errorf("failed to update backup %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// DeleteBackup deletes backup
// API: DELETE /backup/{backupID}
func (c *Client) DeleteBackup(ctx context.Context, id int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/backup/%d", id)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete backup %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetDNSZones returns list of DNS zones
// API: GET /dns
func (c *Client) GetDNSZones(ctx context.Context) ([]models.DNSZone, error) {
	var response models.DNSZoneListResponse

	err := c.Get(ctx, "/dns", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get DNS zones: %w", err)
	}

	if response.Status != "ok" {
		if response.StatusMsg == "No DNS domains information" {
			return []models.DNSZone{}, nil
		}
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	if response.Data == nil {
		return []models.DNSZone{}, nil
	}

	return response.Data, nil
}

// GetDNSZone returns DNS zone by ID
// API: GET /dns/{id}
func (c *Client) GetDNSZone(ctx context.Context, id int) (*models.DNSZone, error) {
	var response models.DNSZoneResponse

	path := fmt.Sprintf("/dns/%d", id)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get DNS zone %d: %w", id, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return &response.Data, nil
}

// CreateDNSZone creates DNS zone
// API: POST /dns
func (c *Client) CreateDNSZone(ctx context.Context, req models.DNSZoneCreateRequest) (int, error) {
	var response models.DNSZoneCreateResponse

	err := c.Post(ctx, "/dns", req, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create DNS zone %s: %w", req.Name, err)
	}

	if response.Status != "ok" {
		return 0, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// DeleteDNSZone deletes DNS zone
// API: DELETE /dns/{id}
func (c *Client) DeleteDNSZone(ctx context.Context, id int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/dns/%d", id)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete DNS zone %d: %w", id, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetDNSRecords returns list of DNS records for zone
// API: GET /dns.record/{serviceID}
func (c *Client) GetDNSRecords(ctx context.Context, zoneID int) ([]models.DNSRecord, error) {
	var response models.DNSRecordListResponse

	path := fmt.Sprintf("/dns.record/%d", zoneID)
	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get DNS records for zone %d: %w", zoneID, err)
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data, nil
}

// GetDNSRecord returns DNS record by ID
// API: GET /dns.record/{serviceID} + search by ID
func (c *Client) GetDNSRecord(ctx context.Context, zoneID, recordID int) (*models.DNSRecord, error) {
	records, err := c.GetDNSRecords(ctx, zoneID)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		if record.ID == recordID {
			return &record, nil
		}
	}

	return nil, fmt.Errorf("DNS record %d not found in zone %d", recordID, zoneID)
}

// CreateDNSRecord creates DNS record
// API: POST /dns.record/{serviceID}
func (c *Client) CreateDNSRecord(ctx context.Context, zoneID int, req models.DNSRecordCreateRequest) (int, error) {
	var response models.DNSRecordCreateResponse

	path := fmt.Sprintf("/dns.record/%d", zoneID)
	err := c.Post(ctx, path, req, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create DNS record in zone %d: %w", zoneID, err)
	}

	if response.Status != "ok" {
		return 0, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return response.Data.ID, nil
}

// UpdateDNSRecord updates DNS record
// API: PUT /dns.record/{recordID}
func (c *Client) UpdateDNSRecord(ctx context.Context, recordID int, req models.DNSRecordUpdateRequest) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/dns.record/%d", recordID)
	err := c.Put(ctx, path, req, &response)
	if err != nil {
		return fmt.Errorf("failed to update DNS record %d: %w", recordID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// DeleteDNSRecord deletes DNS record
// API: DELETE /dns.record/{recordID}
func (c *Client) DeleteDNSRecord(ctx context.Context, recordID int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/dns.record/%d", recordID)
	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to delete DNS record %d: %w", recordID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// AttachISO attaches ISO image to server
// API: PUT /server.iso/{serverID}
func (c *Client) AttachISO(ctx context.Context, serverID int, isoID int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.iso/%d", serverID)
	body := map[string]int{"iso": isoID}

	err := c.Put(ctx, path, body, &response)
	if err != nil {
		return fmt.Errorf("failed to attach ISO %d to server %d: %w", isoID, serverID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// DetachISO detaches ISO image from server
// API: DELETE /server.iso/{serverID}
func (c *Client) DetachISO(ctx context.Context, serverID int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.iso/%d", serverID)

	err := c.Delete(ctx, path, &response)
	if err != nil {
		return fmt.Errorf("failed to detach ISO from server %d: %w", serverID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// RestoreBackup restores backup to server
// API: PUT /backup.restore/{backupID}
func (c *Client) RestoreBackup(ctx context.Context, backupID int, serverID int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/backup.restore/%d", backupID)
	body := map[string]int{"service": serverID}

	err := c.Put(ctx, path, body, &response)
	if err != nil {
		return fmt.Errorf("failed to restore backup %d to server %d: %w", backupID, serverID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// ProlongServer prolongs and starts server
// API: PUT /server.prolong/{serverID}
func (c *Client) ProlongServer(ctx context.Context, serverID int) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.prolong/%d", serverID)
	err := c.Put(ctx, path, nil, &response)
	if err != nil {
		return fmt.Errorf("failed to prolong server %d: %w", serverID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

type ServerPlanChangeRequest struct {
	ServerPlan int `json:"server-plan"`
	CPU        int `json:"cpu,omitempty"`
	RAM        int `json:"ram,omitempty"`
	Disk       int `json:"disk,omitempty"`
	GPU        int `json:"gpu,omitempty"`
}

// ChangeServerPlan changes server plan (upgrade only)
// API: PUT /server.plan/{serverID}
func (c *Client) ChangeServerPlan(ctx context.Context, serverID int, req ServerPlanChangeRequest) error {
	var response struct {
		Status    string `json:"status"`
		StatusMsg string `json:"status_msg"`
	}

	path := fmt.Sprintf("/server.plan/%d", serverID)
	err := c.Put(ctx, path, req, &response)
	if err != nil {
		return fmt.Errorf("failed to change server plan %d: %w", serverID, err)
	}

	if response.Status != "ok" {
		return fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	return nil
}

// GetIPs returns list of IP addresses
// API: GET /ip
func (c *Client) GetIPs(ctx context.Context) ([]models.IP, error) {
	var response models.IPListResponse

	err := c.Get(ctx, "/ip", &response)
	if err != nil {
		return nil, fmt.Errorf("failed to get IPs: %w", err)
	}

	if response.Status != "ok" {
		if response.StatusMsg == "No IP information" {
			return []models.IP{}, nil
		}
		return nil, fmt.Errorf("API returned status: %s - %s", response.Status, response.StatusMsg)
	}

	if response.Data == nil {
		return []models.IP{}, nil
	}

	return response.Data, nil
}

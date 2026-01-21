// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/backup.go
package models

type Backup struct {
	ID         int              `json:"id"`
	Name       string           `json:"name"`
	FullName   string           `json:"full_name"`
	Created    string           `json:"created"`
	Updated    string           `json:"updated"`
	End        string           `json:"end"`
	Status     string           `json:"status"`
	StatusText string           `json:"status_text"`
	Datacenter BackupDatacenter `json:"datacenter"`
	Server     *BackupServer    `json:"server"`
	Can        BackupCan        `json:"can"`
}

type BackupDatacenter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type BackupServer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BackupCan struct {
	Update  bool `json:"update"`
	Prolong bool `json:"prolong"`
	Delete  bool `json:"delete"`
}

type BackupListResponse struct {
	Status    string   `json:"status"`
	StatusMsg string   `json:"status_msg"`
	Data      []Backup `json:"data"`
}

type BackupResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      Backup `json:"data"`
}

type BackupCreateResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
}

type BackupUpdateRequest struct {
	Name        string `json:"name"`
	Autoprolong string `json:"autoprolong,omitempty"`
}

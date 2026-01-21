// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/server.go
package models

type Server struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	FullName    string `json:"full_name"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	End         string `json:"end"`
	Status      string `json:"status"`
	StatusText  string `json:"status_text"`
	Autoprolong bool   `json:"autoprolong"`
	Autorun     *bool  `json:"autorun"`
	Host        string `json:"host"`

	IP        ServerIP        `json:"ip"`
	IPLocal   ServerIPLocal   `json:"ip_local"`
	Data      ServerData      `json:"data"`
	ServerPlan ServerPlanRef  `json:"server-plan"`
	Template   ServerTemplate `json:"template"`
	Datacenter ServerDC       `json:"datacenter"`
	SSHKey     *ServerSSHKey  `json:"ssh-key,omitempty"`
	Can        ServerCan      `json:"can"`
	Bandwidth  ServerBandwidth `json:"bandwidth"`
}

type ServerIP struct {
	ID   int    `json:"id"`
	IP   string `json:"ip"`
	Type string `json:"type"`
}

type ServerIPLocal struct {
	IP      string `json:"ip"`
	Netmask string `json:"netmask"`
	MAC     string `json:"mac"`
}

type ServerData struct {
	CPU  ResourceInfo `json:"cpu"`
	RAM  ResourceInfo `json:"ram"`
	Disk ResourceInfo `json:"disk"`
}

type ServerPlanRef struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServerTemplate struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServerDC struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type ServerSSHKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServerCan struct {
	Reboot  bool `json:"reboot"`
	Update  bool `json:"update"`
	Delete  bool `json:"delete"`
	Prolong bool `json:"prolong"`
	Backup  bool `json:"backup"`
	IPLocal bool `json:"ip_local"`
}

type ServerBandwidth struct {
	CurrentMonth int64 `json:"current_month"`
	PastMonth    int64 `json:"past_month"`
}

type ServerCreateRequest struct {
	Datacenter int    `json:"datacenter"`
	ServerPlan int    `json:"server-plan"`
	Template   int    `json:"template,omitempty"`
	SSHKey     int    `json:"ssh-key,omitempty"`
	Host       string `json:"host,omitempty"`
	Name       string `json:"name,omitempty"`
	Backup     *int   `json:"backup,omitempty"`
	ISO        *int   `json:"iso,omitempty"`
	CPU        int    `json:"cpu,omitempty"`
	RAM        int    `json:"ram,omitempty"`
	Disk       int    `json:"disk,omitempty"`
	GPU        int    `json:"gpu,omitempty"`
}

type ServerUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Autoprolong string `json:"autoprolong,omitempty"`
}

type ServerResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      Server `json:"data"`
}

type ServerCreateResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

type ServersListResponse struct {
	Status    string   `json:"status"`
	StatusMsg string   `json:"status_msg"`
	Data      []Server `json:"data"`
}

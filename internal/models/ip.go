// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/ip.go
package models

type IP struct {
	ID         int          `json:"id"`
	IP         string       `json:"ip"`
	Type       string       `json:"type"`
	Host       string       `json:"host"`
	Gateway    string       `json:"gateway"`
	Netmask    string       `json:"netmask"`
	Datacenter IPDatacenter `json:"datacenter"`
	IsNet      bool         `json:"is_net"`
}

type IPDatacenter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type IPListResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []IP   `json:"data"`
}

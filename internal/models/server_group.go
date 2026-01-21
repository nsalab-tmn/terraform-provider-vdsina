// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/server_group.go
package models

type ServerGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
}

type ServerGroupsResponse struct {
	Status    string        `json:"status"`
	StatusMsg string        `json:"status_msg"`
	Data      []ServerGroup `json:"data"`
}

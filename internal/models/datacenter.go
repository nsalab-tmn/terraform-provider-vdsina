// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/datacenter.go
package models

type Datacenter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Active  bool   `json:"active"`
}

type DatacentersResponse struct {
	Status    string       `json:"status"`
	StatusMsg string       `json:"status_msg"`
	Data      []Datacenter `json:"data"`
}

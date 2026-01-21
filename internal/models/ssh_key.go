// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/ssh_key.go
package models

type SSHKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Data string `json:"data,omitempty"`
}

type SSHKeyCreateRequest struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type SSHKeyUpdateRequest struct {
	Name string `json:"name,omitempty"`
	Data string `json:"data,omitempty"`
}

type SSHKeysResponse struct {
	Status    string   `json:"status"`
	StatusMsg string   `json:"status_msg"`
	Data      []SSHKey `json:"data"`
}

type SSHKeyResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      SSHKey `json:"data"`
}

type SSHKeyCreateResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

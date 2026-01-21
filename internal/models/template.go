// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/template.go
package models

type Template struct {
	ID            int            `json:"id"`
	Name          string         `json:"name"`
	Active        bool           `json:"active"`
	SSHKey        bool           `json:"ssh-key"`
	TemplateGroup int            `json:"template_group"`
	ServerPlans   []int          `json:"server-plan"`
	Limits        TemplateLimits `json:"limits"`
}

type TemplateLimits struct {
	CPU  TemplateLimitValue `json:"cpu"`
	RAM  TemplateLimitValue `json:"ram"`
	Disk TemplateLimitValue `json:"disk"`
}

type TemplateLimitValue struct {
	Min int `json:"min"`
}

type TemplatesResponse struct {
	Status    string     `json:"status"`
	StatusMsg string     `json:"status_msg"`
	Data      []Template `json:"data"`
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/server_plan.go
package models

type ServerPlan struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Cost        float64 `json:"cost"`
	FullCost    float64 `json:"full_cost"`
	Period      string  `json:"period"`
	MinMoney    float64 `json:"min_money"`
	CanBonus    bool    `json:"can_bonus"`
	Description string  `json:"description"`
	Active      bool    `json:"active"`
	Enable      bool    `json:"enable"`
	HasParams   bool    `json:"has_params"`
	Selected    bool    `json:"selected"`

	Backup      BackupInfo        `json:"backup"`
	Data        ServerPlanData    `json:"data"`
	ServerGroup int               `json:"server-group"`
	Params      *ServerPlanParams `json:"params,omitempty"`
}

type BackupInfo struct {
	Cost     float64 `json:"cost"`
	FullCost float64 `json:"full_cost"`
	Period   string  `json:"period"`
	For      string  `json:"for"`
}

type ServerPlanData struct {
	CPU   ResourceInfo `json:"cpu"`
	RAM   ResourceInfo `json:"ram"`
	Disk  ResourceInfo `json:"disk"`
	Traff ResourceInfo `json:"traff"`
}

type ResourceInfo struct {
	Value int    `json:"value"`
	Bytes int64  `json:"bytes,omitempty"`
	For   string `json:"for"`
}

type ServerPlanParams struct {
	CPU  ParamRange `json:"cpu,omitempty"`
	RAM  ParamRange `json:"ram,omitempty"`
	Disk ParamRange `json:"disk,omitempty"`
	GPU  ParamRange `json:"gpu,omitempty"`
}

type ParamRange struct {
	Min  int     `json:"min"`
	Max  int     `json:"max"`
	Step int     `json:"step"`
	Cost float64 `json:"cost"`
}

type ServerPlansResponse struct {
	Status    string       `json:"status"`
	StatusMsg string       `json:"status_msg"`
	Data      []ServerPlan `json:"data"`
}

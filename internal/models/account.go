// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/account.go
package models

import "encoding/json"

type Account struct {
	Account  AccountInfo `json:"account"`
	Created  string      `json:"created"`
	Forecast string      `json:"forecast"`
	Can      AccountCan  `json:"can"`
}

type AccountInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AccountCan struct {
	AddUser       bool `json:"add_user"`
	AddService    bool `json:"add_service"`
	ConvertToCash bool `json:"convert_to_cash"`
}

type AccountBalance struct {
	Real    string `json:"real"`
	Bonus   string `json:"bonus"`
	Partner string `json:"partner"`
}

type AccountResponse struct {
	Status    string  `json:"status"`
	StatusMsg string  `json:"status_msg"`
	Data      Account `json:"data"`
}

type AccountBalanceResponse struct {
	Status    string         `json:"status"`
	StatusMsg string         `json:"status_msg"`
	Data      AccountBalance `json:"data"`
}

type LimitInfo struct {
	Max      int `json:"max"`
	Now      int `json:"now"`
	ChildMax int `json:"child_max,omitempty"`
}

type AccountLimitsRaw struct {
	Server      LimitInfo       `json:"server"`
	ServerIP4   LimitInfo       `json:"server-ip4"`
	ServerIP6   LimitInfo       `json:"server-ip6"`
	ISO         LimitInfo       `json:"iso"`
	Backup      LimitInfo       `json:"backup"`
	SSL         json.RawMessage `json:"ssl"`
	Domain      json.RawMessage `json:"domain"`
	DNS         LimitInfo       `json:"dns"`
	ExtdiskHDD  LimitInfo       `json:"extdisk-hdd"`
	ExtdiskNVMe LimitInfo       `json:"extdisk-nvme"`
	ReserveIP   LimitInfo       `json:"reserve-ip"`
	GPU         LimitInfo       `json:"gpu"`
}

type AccountLimits struct {
	Server      LimitInfo
	ServerIP4   LimitInfo
	ServerIP6   LimitInfo
	ISO         LimitInfo
	Backup      LimitInfo
	SSL         *LimitInfo
	Domain      *LimitInfo
	DNS         LimitInfo
	ExtdiskHDD  LimitInfo
	ExtdiskNVMe LimitInfo
	ReserveIP   LimitInfo
	GPU         LimitInfo
}

type AccountLimitsResponse struct {
	Status    string           `json:"status"`
	StatusMsg string           `json:"status_msg"`
	Data      AccountLimitsRaw `json:"data"`
}

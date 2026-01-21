// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/dns.go
package models

type DNSZone struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FullName   string     `json:"full_name"`
	Created    string     `json:"created"`
	Updated    string     `json:"updated"`
	End        string     `json:"end"`
	Status     string     `json:"status"`
	StatusText string     `json:"status_text"`
	Real       bool       `json:"real"`
	Can        DNSZoneCan `json:"can"`
}

type DNSZoneCan struct {
	Delete bool `json:"delete"`
}

type DNSZoneListResponse struct {
	Status    string    `json:"status"`
	StatusMsg string    `json:"status_msg"`
	Data      []DNSZone `json:"data"`
}

type DNSZoneResponse struct {
	Status    string  `json:"status"`
	StatusMsg string  `json:"status_msg"`
	Data      DNSZone `json:"data"`
}

type DNSZoneCreateRequest struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

type DNSZoneCreateResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

type DNSRecordCan struct {
	Update bool `json:"update"`
	Delete bool `json:"delete"`
}

type DNSRecord struct {
	ID        int          `json:"id"`
	Host      string       `json:"host"`
	Type      string       `json:"type"`
	Value     string       `json:"value"`
	Priority  *int         `json:"priority"`
	Tag       *string      `json:"tag"`
	Timestamp string       `json:"timestamp"`
	Can       DNSRecordCan `json:"can"`
}

type DNSRecordListResponse struct {
	Status    string      `json:"status"`
	StatusMsg string      `json:"status_msg"`
	Data      []DNSRecord `json:"data"`
}

type DNSRecordCreateRequest struct {
	Host     string  `json:"host"`
	Type     string  `json:"type"`
	Value    string  `json:"value"`
	Priority *int    `json:"priority,omitempty"`
	Tag      *string `json:"tag,omitempty"`
}

type DNSRecordCreateResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

type DNSRecordUpdateRequest struct {
	Value    string  `json:"value"`
	Priority *int    `json:"priority,omitempty"`
	Tag      *string `json:"tag,omitempty"`
}

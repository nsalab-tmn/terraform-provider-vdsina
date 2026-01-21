// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/models/iso.go
package models

type ISO struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	FullName   string     `json:"full_name"`
	Created    string     `json:"created"`
	Updated    string     `json:"updated"`
	End        string     `json:"end"`
	Status     string     `json:"status"`
	StatusText string     `json:"status_text"`
	File       ISOFile    `json:"file"`
	Attached   bool       `json:"attached"`
	Server     *ISOServer `json:"server"`
	Can        ISOCan     `json:"can"`
}

type ISOFile struct {
	Size string `json:"size"`
	MD5  string `json:"md5"`
}

type ISOServer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ISOCan struct {
	Delete bool `json:"delete"`
}

type ISOListResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      []ISO  `json:"data"`
}

type ISOResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      ISO    `json:"data"`
}

type ISODownloadResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID string `json:"id"`
	} `json:"data"`
}

type ISODownloadStatusResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"data"`
}

type ISOCreateFromKeyResponse struct {
	Status    string `json:"status"`
	StatusMsg string `json:"status_msg"`
	Data      struct {
		ID int `json:"id"`
	} `json:"data"`
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/client.go
package client

import (
	"net/http"
	"time"
)

type Client struct {
	host       string
	apiToken   string
	httpClient *http.Client
}

func NewClient(baseURL, apiToken string) (*Client, error) {
	return &Client{
		host:     baseURL,
		apiToken: apiToken,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}, nil
}

func (c *Client) GetHost() string {
	return c.host
}

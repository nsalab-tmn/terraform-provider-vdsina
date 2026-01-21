// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/logs.go
package client

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

func (c *Client) logRequest(method, url string, body []byte, headers http.Header) {
	log.Printf("[DEBUG] === HTTP REQUEST ===")
	log.Printf("[DEBUG] Method: %s", method)
	log.Printf("[DEBUG] URL: %s", url)

	log.Printf("[DEBUG] Headers:")
	for key, values := range headers {
		if strings.ToLower(key) == "authorization" {
			log.Printf("[DEBUG]   %s: [REDACTED]", key)
		} else {
			log.Printf("[DEBUG]   %s: %v", key, values)
		}
	}

	log.Printf("[DEBUG] Request Body:")
	if len(body) > 0 {
		if c.isJSONContent(headers) {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
				log.Printf("[DEBUG] %s", prettyJSON.String())
			} else {
				log.Printf("[DEBUG] %s", string(body))
			}
		} else {
			log.Printf("[DEBUG] %s", string(body))
		}
	} else {
		log.Printf("[DEBUG] [EMPTY]")
	}
	log.Printf("[DEBUG] ====================")
}

func (c *Client) logResponse(method, url string, status string, body []byte, headers http.Header, duration time.Duration) {
	log.Printf("[DEBUG] === HTTP RESPONSE ===")
	log.Printf("[DEBUG] Method: %s", method)
	log.Printf("[DEBUG] URL: %s", url)
	log.Printf("[DEBUG] Status: %s", status)
	log.Printf("[DEBUG] Duration: %v", duration)

	if len(headers) > 0 {
		log.Printf("[DEBUG] Response Headers:")
		for key, values := range headers {
			log.Printf("[DEBUG]   %s: %v", key, values)
		}
	}

	log.Printf("[DEBUG] Response Body:")
	if len(body) > 0 {
		if c.isJSONContent(headers) {
			var prettyJSON bytes.Buffer
			if err := json.Indent(&prettyJSON, body, "", "  "); err == nil {
				log.Printf("[DEBUG] %s", prettyJSON.String())
			} else {
				log.Printf("[DEBUG] %s (invalid JSON)", string(body))
			}
		} else {
			bodyStr := string(body)
			if len(bodyStr) > 1000 {
				bodyStr = bodyStr[:1000] + "... [TRUNCATED]"
			}
			log.Printf("[DEBUG] %s", bodyStr)
		}
	} else {
		log.Printf("[DEBUG] [EMPTY]")
	}
	log.Printf("[DEBUG] =====================")
}

func (c *Client) isJSONContent(headers http.Header) bool {
	contentType := headers.Get("Content-Type")
	return strings.Contains(contentType, "application/json") ||
		strings.Contains(contentType, "text/json") ||
		strings.Contains(contentType, "json")
}

func (c *Client) logError(method, url string, err error, duration time.Duration) {
	log.Printf("[ERROR] === HTTP ERROR ===")
	log.Printf("[ERROR] Method: %s", method)
	log.Printf("[ERROR] URL: %s", url)
	log.Printf("[ERROR] Error: %v", err)
	log.Printf("[ERROR] Duration: %v", duration)
	log.Printf("[ERROR] ==================")
}

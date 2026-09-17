// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/client/errors.go
package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type APIError struct {
	StatusCode  int
	Status      string
	StatusMsg   string
	Description string
	Data        map[string][]string
}

func (e *APIError) Error() string {
	if e.Description != "" {
		return fmt.Sprintf("API Error %d: %s - %s", e.StatusCode, e.StatusMsg, e.Description)
	}
	return fmt.Sprintf("API Error %d: %s", e.StatusCode, e.StatusMsg)
}

type ErrorResponse struct {
	Status      string              `json:"status"`
	StatusMsg   string              `json:"status_msg"`
	Description string              `json:"description"`
	Data        map[string][]string `json:"data"`
}

func parseAPIError(statusCode int, body []byte) error {
	var errResp ErrorResponse
	if err := json.Unmarshal(body, &errResp); err != nil {
		return &APIError{
			StatusCode:  statusCode,
			Status:      "error",
			StatusMsg:   fmt.Sprintf("HTTP %d", statusCode),
			Description: string(body),
		}
	}

	return &APIError{
		StatusCode:  statusCode,
		Status:      errResp.Status,
		StatusMsg:   errResp.StatusMsg,
		Description: errResp.Description,
		Data:        errResp.Data,
	}
}

// hasStatusCode unwraps err, since the client methods wrap API errors with %w.
func hasStatusCode(err error, statusCode int) bool {
	var apiErr *APIError
	return errors.As(err, &apiErr) && apiErr.StatusCode == statusCode
}

func IsNotFound(err error) bool {
	return hasStatusCode(err, 404)
}

func IsUnauthorized(err error) bool {
	return hasStatusCode(err, 401)
}

func IsForbidden(err error) bool {
	return hasStatusCode(err, 403)
}

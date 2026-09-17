// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package client

import (
	"errors"
	"fmt"
	"testing"
)

func TestStatusHelpersUnwrap(t *testing.T) {
	notFound := fmt.Errorf("failed to get server 42: %w", &APIError{StatusCode: 404})
	forbidden := fmt.Errorf("failed to get server 42: %w", &APIError{StatusCode: 403})
	unauthorized := fmt.Errorf("outer: %w", fmt.Errorf("inner: %w", &APIError{StatusCode: 401}))

	if !IsNotFound(notFound) {
		t.Error("IsNotFound(wrapped 404) = false, want true")
	}
	if IsNotFound(forbidden) {
		t.Error("IsNotFound(wrapped 403) = true, want false")
	}
	if !IsForbidden(forbidden) {
		t.Error("IsForbidden(wrapped 403) = false, want true")
	}
	if !IsUnauthorized(unauthorized) {
		t.Error("IsUnauthorized(double-wrapped 401) = false, want true")
	}
	if IsNotFound(errors.New("connection reset")) {
		t.Error("IsNotFound(non-API error) = true, want false")
	}
	if IsNotFound(nil) {
		t.Error("IsNotFound(nil) = true, want false")
	}
}

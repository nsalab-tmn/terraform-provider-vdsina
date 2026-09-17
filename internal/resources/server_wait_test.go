// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package resources

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

const notFoundBody = `{"status":"error","status_code":404,"status_msg":"","description":"Not Found (Service not found)","data":null}`

type apiReply struct {
	code int
	body string
}

func serverBody(status, ip string) apiReply {
	ipField := "null"
	if ip != "" {
		ipField = fmt.Sprintf(`{"id":1,"ip":%q,"type":"4"}`, ip)
	}
	return apiReply{
		code: http.StatusOK,
		body: fmt.Sprintf(`{"status":"ok","status_msg":"Server information","data":{"id":42,"status":%q,"status_text":"","ip":%s}}`, status, ipField),
	}
}

// newSequenceClient serves the replies in order for GET /server/42, repeating the last one.
func newSequenceClient(t *testing.T, replies ...apiReply) (*client.Client, *int32) {
	t.Helper()
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/server/42" {
			t.Errorf("unexpected request path %s", r.URL.Path)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		n := int(atomic.AddInt32(&calls, 1)) - 1
		if n >= len(replies) {
			n = len(replies) - 1
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(replies[n].code)
		_, _ = w.Write([]byte(replies[n].body))
	}))
	t.Cleanup(srv.Close)

	c, err := client.NewClient(srv.URL, "test-token")
	if err != nil {
		t.Fatal(err)
	}
	return c, &calls
}

func TestServerReadyRefresh(t *testing.T) {
	cases := []struct {
		name       string
		reply      apiReply
		wantState  string
		wantResult bool
		wantErr    bool
	}{
		{name: "new", reply: serverBody("new", ""), wantState: serverStateCreating, wantResult: true},
		{name: "active without ip", reply: serverBody("active", ""), wantState: serverStateWaitingForIP, wantResult: true},
		{name: "active with ip", reply: serverBody("active", "203.0.113.10"), wantState: serverStateReady, wantResult: true},
		{name: "blocked", reply: serverBody("block", ""), wantState: "block", wantResult: true, wantErr: true},
		{name: "not found", reply: apiReply{code: http.StatusNotFound, body: notFoundBody}},
		{name: "server error", reply: apiReply{code: http.StatusInternalServerError, body: `{"status":"error"}`}, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := newSequenceClient(t, tc.reply)
			result, state, err := serverReadyRefresh(context.Background(), c, 42)()

			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if state != tc.wantState {
				t.Errorf("state = %q, want %q", state, tc.wantState)
			}
			if (result != nil) != tc.wantResult {
				t.Errorf("result = %v, want non-nil %v", result, tc.wantResult)
			}
		})
	}
}

func TestServerDeletedRefresh(t *testing.T) {
	cases := []struct {
		name      string
		reply     apiReply
		wantState string
		wantErr   bool
	}{
		{name: "not found", reply: apiReply{code: http.StatusNotFound, body: notFoundBody}, wantState: serverStateDeleted},
		{name: "status deleted", reply: serverBody("deleted", ""), wantState: serverStateDeleted},
		{name: "still active", reply: serverBody("active", "203.0.113.10"), wantState: serverStateDeleting},
		{name: "server error", reply: apiReply{code: http.StatusInternalServerError, body: `{"status":"error"}`}, wantErr: true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c, _ := newSequenceClient(t, tc.reply)
			_, state, err := serverDeletedRefresh(context.Background(), c, 42)()

			if (err != nil) != tc.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tc.wantErr)
			}
			if state != tc.wantState {
				t.Errorf("state = %q, want %q", state, tc.wantState)
			}
		})
	}
}

func shortenServerWait(t *testing.T) {
	t.Helper()
	delay, poll := serverWaitDelay, serverWaitPollInterval
	serverWaitDelay, serverWaitPollInterval = 10*time.Millisecond, 10*time.Millisecond
	t.Cleanup(func() { serverWaitDelay, serverWaitPollInterval = delay, poll })
}

func TestWaitForServerReady(t *testing.T) {
	shortenServerWait(t)

	c, calls := newSequenceClient(t,
		serverBody("new", ""),
		serverBody("new", ""),
		serverBody("active", ""),
		serverBody("active", "203.0.113.10"),
	)

	if err := waitForServerReady(context.Background(), c, 42, 5*time.Second); err != nil {
		t.Fatalf("waitForServerReady: %v", err)
	}
	if got := atomic.LoadInt32(calls); got < 4 {
		t.Errorf("GET /server/42 called %d times, want at least 4", got)
	}
}

func TestWaitForServerReadyFailsOnBlockedServer(t *testing.T) {
	shortenServerWait(t)

	c, _ := newSequenceClient(t, serverBody("new", ""), serverBody("notpaid", ""))

	if err := waitForServerReady(context.Background(), c, 42, 5*time.Second); err == nil {
		t.Fatal("waitForServerReady succeeded for a notpaid server, want error")
	}
}

func TestWaitForServerDeleted(t *testing.T) {
	shortenServerWait(t)

	c, _ := newSequenceClient(t,
		serverBody("active", "203.0.113.10"),
		apiReply{code: http.StatusNotFound, body: notFoundBody},
	)

	if err := waitForServerDeleted(context.Background(), c, 42, 5*time.Second); err != nil {
		t.Fatalf("waitForServerDeleted: %v", err)
	}
}

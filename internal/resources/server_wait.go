// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/resources/server_wait.go
package resources

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/client"
)

const (
	serverStateCreating     = "creating"
	serverStateWaitingForIP = "waiting_for_ip"
	serverStateReady        = "ready"
	serverStateDeleting     = "deleting"
	serverStateDeleted      = "deleted"
)

// Variables so tests can shorten them.
var (
	serverWaitDelay        = 5 * time.Second
	serverWaitPollInterval = 10 * time.Second
)

// serverReadyRefresh reports a new server as ready once VDSina has activated it
// and assigned a public IP. POST /server returns before either happens.
func serverReadyRefresh(ctx context.Context, c *client.Client, id int) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		server, err := c.GetServer(ctx, id)
		if err != nil {
			if client.IsNotFound(err) {
				return nil, "", nil
			}
			return nil, "", err
		}

		switch server.Status {
		case "new":
			return server, serverStateCreating, nil
		case "active":
			if server.IP.IP == "" {
				return server, serverStateWaitingForIP, nil
			}
			return server, serverStateReady, nil
		default:
			return server, server.Status, fmt.Errorf("server %d entered status %q while waiting to become active: %s", id, server.Status, server.StatusText)
		}
	}
}

// serverDeletedRefresh reports a server as deleted once the API returns 404 or status "deleted".
func serverDeletedRefresh(ctx context.Context, c *client.Client, id int) retry.StateRefreshFunc {
	return func() (interface{}, string, error) {
		server, err := c.GetServer(ctx, id)
		if err != nil {
			if client.IsNotFound(err) {
				return id, serverStateDeleted, nil
			}
			return nil, "", err
		}

		if server.Status == "deleted" {
			return server, serverStateDeleted, nil
		}
		return server, serverStateDeleting, nil
	}
}

func waitForServerReady(ctx context.Context, c *client.Client, id int, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending:        []string{serverStateCreating, serverStateWaitingForIP},
		Target:         []string{serverStateReady},
		Refresh:        serverReadyRefresh(ctx, c, id),
		Timeout:        timeout,
		Delay:          serverWaitDelay,
		PollInterval:   serverWaitPollInterval,
		NotFoundChecks: 6,
	}

	if _, err := conf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("waiting for server %d to become active with a public IP: %w", id, err)
	}
	return nil
}

func waitForServerDeleted(ctx context.Context, c *client.Client, id int, timeout time.Duration) error {
	conf := &retry.StateChangeConf{
		Pending:      []string{serverStateDeleting},
		Target:       []string{serverStateDeleted},
		Refresh:      serverDeletedRefresh(ctx, c, id),
		Timeout:      timeout,
		Delay:        serverWaitDelay,
		PollInterval: serverWaitPollInterval,
	}

	if _, err := conf.WaitForStateContext(ctx); err != nil {
		return fmt.Errorf("waiting for server %d to be deleted: %w", id, err)
	}
	return nil
}

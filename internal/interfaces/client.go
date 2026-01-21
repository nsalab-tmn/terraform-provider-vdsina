// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

// internal/interfaces/client.go
package interfaces

import (
	"context"

	"github.com/scinfra-pro/terraform-provider-vdsina/internal/models"
)

type DataClient interface {
	GetDatacenters(ctx context.Context) ([]models.Datacenter, error)
	GetServerGroups(ctx context.Context) ([]models.ServerGroup, error)
	GetServerPlans(ctx context.Context, groupID int) ([]models.ServerPlan, error)
	GetTemplates(ctx context.Context) ([]models.Template, error)
}

type ResourceClient interface {
	DataClient
}

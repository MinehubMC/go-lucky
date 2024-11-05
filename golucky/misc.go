package golucky

import (
	"context"
	"fmt"
)

type HealthResult struct {
	Healthy bool `json:"healthy"`
	Details any  `json:"details"`
}

func (c *Client) CheckHealth(ctx context.Context) (*HealthResult, error) {
	return getRequest[HealthResult](ctx, fmt.Sprintf("%s/health", c.config.RestIp), c.config.AuthKey)
}

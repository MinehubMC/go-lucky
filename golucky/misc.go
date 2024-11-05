package golucky

import "fmt"

type HealthResult struct {
	Healthy bool `json:"healthy"`
	Details any  `json:"details"`
}

func (c *Client) CheckHealth() (*HealthResult, error) {
	return getRequest[HealthResult](fmt.Sprintf("%s/health", c.config.RestIp), c.config.AuthKey)
}

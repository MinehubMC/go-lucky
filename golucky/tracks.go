package golucky

import (
	"context"
	"fmt"
)

type Track struct {
	Name   string   `json:"name"`
	Groups []string `json:"groups"`
}

type newTrack struct {
	Name string `json:"name"`
}
type newGroups struct {
	Groups []string `json:"groups"`
}

func (c *Client) ListTracks(ctx context.Context) (*[]string, error) {
	return getRequest[[]string](ctx, fmt.Sprintf("%s/track", c.config.RestIp), c.config.AuthKey)
}

func (c *Client) CreateTrack(ctx context.Context, trackName string) (*Track, error) {
	return postRequestBody[Track](ctx, fmt.Sprintf("%s/track", c.config.RestIp), newTrack{Name: trackName}, c.config.AuthKey)
}

func (c *Client) GetTrack(ctx context.Context, trackName string) (*Track, error) {
	return getRequest[Track](ctx, fmt.Sprintf("%s/track/%s", c.config.RestIp, trackName), c.config.AuthKey)
}

func (c *Client) UpdateTrack(ctx context.Context, trackName string, groups []string) (*Track, error) {
	return patchRequestBody[Track](ctx, fmt.Sprintf("%s/track/%s", c.config.RestIp, trackName), newGroups{Groups: groups}, c.config.AuthKey)
}

func (c *Client) DeleteTrack(ctx context.Context, trackName string) error {
	return deleteRequestNoResponse(ctx, fmt.Sprintf("%s/track/%s", c.config.RestIp, trackName), c.config.AuthKey)
}
